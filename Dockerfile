FROM alpine:3.15.0
# copy over the binary from the first stage
COPY go-hello /bin/
WORKDIR "/"
ENTRYPOINT ["/bin/ash" ]
