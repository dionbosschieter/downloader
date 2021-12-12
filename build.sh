#!/bin/sh -ex

CGO_ENABLED=0 go build -o downloader main.go
