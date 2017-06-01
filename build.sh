#!/bin/bash
GOOS=windows GOARCH=amd64 go build github.com/hansnow/smms
GOOS=linux GOARCH=amd64 go build github.com/hansnow/smms
GOOS=darwin GOARCH=amd64 go build github.com/hansnow/smms