package config

import (
	"errors"
	"os"
	"strings"
)

type OblioConfig struct {
	ClientId     string
	ClientSecret string
	CIFs         []string
}

func NewOblioConfig() (OblioConfig, error) {
	clientId, ok := os.LookupEnv("OBLIO_CLIENT_ID")
	if !ok {
		return OblioConfig{}, errors.New("client id not set")
	}
	clientSecret, ok := os.LookupEnv("OBLIO_CLIENT_SECRET")
	if !ok {
		return OblioConfig{}, errors.New("client secret not set")
	}
	cifval, ok := os.LookupEnv("OBLIO_CLIENT_CIF")
	if !ok {
		return OblioConfig{}, errors.New("client CIFs not set")
	}
	cifs := strings.Split(cifval, ",")
	return OblioConfig{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		CIFs:         cifs,
	}, nil
}
