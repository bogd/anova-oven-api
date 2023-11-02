FROM golang:1.21-alpine as builder

RUN mkdir -p /oven-client/src

WORKDIR /oven-client

COPY src/go.mod src/go.sum /oven-client/src/

WORKDIR /oven-client/src

RUN \
  go mod download -x 

COPY src .

RUN go build -o ../bin/oven-client -v .

# Final image
# FROM scratch	
FROM alpine:latest

LABEL maintainer="bogd"

COPY --from=builder /oven-client/bin/oven-client /oven-client
  
ENTRYPOINT ["/oven-client"]

