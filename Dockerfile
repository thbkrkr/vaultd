FROM scratch
COPY vaultd /vaultd
ENTRYPOINT ["/vaultd"]
