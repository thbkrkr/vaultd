FROM alpine:3.6

RUN apk --no-cache add bash curl
COPY vaultd /usr/local/bin/vaultd
COPY bin/v /usr/local/bin/v
ENTRYPOINT ["v"]
CMD ["vaultd"]

ONBUILD COPY data /data