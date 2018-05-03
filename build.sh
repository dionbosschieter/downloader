#!/bin/sh -e
# todo: convert to makefile
go build -o downloader main.go

fpm -f -s dir -t deb -n godownloader downloader=/usr/local/bin/
