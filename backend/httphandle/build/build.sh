#!/bin/bash
GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o restapp gohttpexamples/sample4/delivery/restapplication