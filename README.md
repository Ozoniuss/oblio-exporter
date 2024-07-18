# Oblio exporter

In Romania, every limited responsibility company must download all issued invoices to customers and store them for up to 5 years, due to proper government spending in digitalization. That means my mom, who essentially learned to use the computer in the last 10 years because it was a must for owning a company, is faced with needing a resilient storage solution where she can back up all those invoices, because the issuer's platform only keeps them for two months.

Now, imagine what would mean downloading every single bill, and storing it on her 6 year-old cheap laptop alongisde hundreds of others. And expecting them to be there 5 years from now. This solution downloads all issued bills from the platform my mom uses, and uploads them to Backblaze.

## Running the solution

Set your environment variables:

```bash
export OBLIO_CLIENT_ID=
export OBLIO_CLIENT_SECRET=
export OBLIO_CLIENT_CIF=
```

and simply run and specifiy the month for which you want to download bills:

```
please select billing month (e.g. 6 for june) > 5
```

Issued invoices will download for all CIFs.

## Known issues

-   Running the solution for the first time panics during authorization. Everything seems to work as expected after running it for a second time.
