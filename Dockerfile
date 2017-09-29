FROM alpine:3.6

RUN apk --no-cache add bash curl

COPY vaultd /usr/local/bin/vaultd
COPY bin/crypt /usr/local/bin/crypt
