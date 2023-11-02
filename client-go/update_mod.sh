#!/bin/bash

docker container run --rm \
  -v $(pwd)/src:/usr/src/anova-oven-api \
  -w /usr/src/anova-oven-api \
  golang:1.21-alpine \
  go mod tidy
  
  
