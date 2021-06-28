#!/bin/sh -ex

go build -o downloader main.go

# create appropriate folder for plugins
fpm -f -s dir -t deb -n godownloader downloader=/usr/local/bin/
