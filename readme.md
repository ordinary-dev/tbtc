# TBTC

*Take back the certificates!*

This program can turn the certificates that traefik stores in the acme.json file
into a regular pair of .pem files.

This is useful if you need certificates somewhere else, for example for a mail server,
but the traefik format is not supported.

At the moment, the program just runs every 3 hours.

## Usage
```yml
services:
  tbtc:
    image: ghcr.io/ordinary-dev/tbtc
    volumes:
      - traefik:/traefik
      - certificates:/data
    environment:
      # Required
      - TBTC_TARGET_DOMAIN=example.com
      # Default: acme.json
      - TBTC_ACME_FILE_PATH=/traefik/acme.json
      # Default: fullchain.pem
      - TBTC_CERTIFICATE_FILE_PATH=/data/fullchain.pem
      # Default: privkey.pem
      - TBTC_KEY_FILE_PATH=/data/privkey.pem
    restart: unless-stopped
```
