#!/bin/bash

rm -fr output
mkdir -p output
cp -r conf output/
CGO_ENASBLED=0 GOOS=linux GOARCH=amd64 go build -o ./output/gosdocker
