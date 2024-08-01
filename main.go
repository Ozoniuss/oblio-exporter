package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	config "oblio-exporter/config"
	"oblio-exporter/httputil"
	monthpkg "oblio-exporter/month"
)

type AuthResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope     string `json:"scope"`
}

type ListResponse struct {
	Data []ListResponseData `json:"data"`
}

type ListResponseData struct {
	Einvoice   string `json:"einvoice"`
	SeriesName string `json:"seriesName"`
	Number     string `json:"number"`
}

type requester struct {
	client *http.Client
	token  string
	logger *slog.Logger

	config config.OblioConfig
}

// authorize generates a bearer token for the requester that can be used for
// subsequent requests
func (r *requester) authorize() error {
	{
		urldata := url.Values{}
		urldata.Set("client_id", r.config.ClientId)
		urldata.Set("client_secret", r.config.ClientSecret)

		encoded := urldata.Encode()

		req, err := http.NewRequest("POST", "https://www.oblio.eu/api/authorize/token", strings.NewReader(encoded))
		if err != nil {
			return fmt.Errorf("creating request: %s", err.Error())
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("sending request: %s", err.Error())
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return errors.New("unknown error occured")
			}
			return fmt.Errorf("got bad response: %s", string(body))
		}

		d := json.NewDecoder(response.Body)
		var authResp AuthResponse
		err = d.Decode(&authResp)
		if err != nil {
			body, err := io.ReadAll(response.Body)
			slog.Debug("got auth response", "body", string(body), "err", err)
			return fmt.Errorf("decoding response body: %s", err.Error())
		}

		r.token = authResp.Token
		return nil
	}
}

func (r *requester) downloadInvoiceForCif(month int, cif string) error {

	req, err := http.NewRequest("GET", "https://www.oblio.eu/api/docs/invoice/list", nil)
	if err != nil {
		return fmt.Errorf("creating request: %s", err.Error())
	}
	req.Header.Add("Authorization", "Bearer "+r.token)

	query := req.URL.Query()
	query.Add("cif", "RO"+cif)
	query.Add("issuedBefore", monthpkg.FormatEndOfMonth(time.Month(month)))
	query.Add("issuedAfter", monthpkg.FormatBeginningOfMonth(time.Month(month)))
	req.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return errors.New("unknown error occured")
		}
		return fmt.Errorf("got bad response: %s", string(body))
	}

	d := json.NewDecoder(response.Body)
	var listResp ListResponse
	err = d.Decode(&listResp)
	if err != nil {
		return fmt.Errorf("decoding response body: %s", err.Error())
	}

	for _, individualReport := range listResp.Data {
		generatedFilename := individualReport.SeriesName + individualReport.Number
		r.logger.Debug("generated filename", "filename", generatedFilename)
		filename, err := r.downloadFile(generatedFilename, individualReport.Einvoice)
		if err != nil {
			fmt.Printf("could not download report for einvoice %s (%s): %s\n", individualReport.Einvoice, generatedFilename, err.Error())
		} else {
			fmt.Printf("successfully downloaded report: file %s\n", filename)
		}
	}

	return nil
}

// downloadFile makes a request to the einvoice url and downloads the report.
//
// generatedFilename is used if the response does not include a proper filename
// that can be used to name the report.
func (r *requester) downloadFile(generatedFilename, url string) (string, error) {

	resp, err := r.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error accessing file: %s", err.Error())
	}
	defer resp.Body.Close()

	filename, err := httputil.GetFileNameFromHeader(resp.Header.Get("Content-Disposition"))
	if err != nil {
		fmt.Printf("could not get filename from header: %s\n", err.Error())
		fmt.Printf("creating filename from series and number instead: %s\n", generatedFilename)
		filename = generatedFilename
	}
	// Create blank file
	f, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("could not create file: %s", err.Error())
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Printf("could not write report for %s: %s", filename, err.Error())
	}
	return filename, nil
}

func run() error {
	config, err := config.NewOblioConfig()
	if err != nil {
		return fmt.Errorf("could not read config: %s", err.Error())
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	r := requester{
		client: http.DefaultClient,
		config: config,
		logger: logger,
	}
	fmt.Println("connecting to oblio...")
	err = r.authorize()
	if err != nil {
		return fmt.Errorf("could not connect to oblio: %s", err.Error())
	}
	fmt.Printf("connected to oblio account %s\n", r.config.ClientId)

	month := monthpkg.GetMonthInput()
	M := time.Month(month)
	fmt.Printf("exporting report for %s...\n", M)

	for _, cif := range config.CIFs {
		fmt.Printf("downloading invoice for cif %s...\n", cif)
		err := r.downloadInvoiceForCif(month, cif)
		if err != nil {
			return fmt.Errorf("could not list report: %s", err.Error())
		}
		fmt.Printf("finished downloading invoice for cif %s\n", cif)
	}

	return nil
}

func main() {

	json2 := `
	{
		"mytype": "hello",
		"lol": 4
	}
	`
	r := strings.NewReader(json2)

	type S struct {
		Mytype string `json:"mytype"`
	}

	var s S
	d := json.NewDecoder(r)
	err := d.Decode(&s)
	fmt.Println(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)

	err = run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
