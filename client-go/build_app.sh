#!/bin/bash

docker container run --rm \
  -v $(pwd)/src:/usr/src/anova-oven-api \
  -v $(pwd)/binaries:/usr/bin/anova-oven-api \
  -v .env \
  -v /tmp/go:/go \
  -v /tmp/gocache:/root/.cache/go-build \
  -w /usr/src/anova-oven-api \
  golang:1.21-alpine \
  go build -v -o /usr/bin/anova-oven-api/anova-oven-api .
  

