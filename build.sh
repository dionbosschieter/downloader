#!/bin/sh
# todo: convert to makefile
go build downloader.go settings.go telebot.go transmission.go

fpm -f -s dir -t deb -n godownloader downloader=/usr/local/bin/
