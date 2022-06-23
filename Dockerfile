FROM alpine:3.15.0
# copy over the binary from the first stage
COPY go-hello /helloworld/helloworld
WORKDIR “/helloworld”
ENTRYPOINT [ “/helloworld/helloworld” ]