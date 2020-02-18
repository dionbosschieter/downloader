#!/bin/sh -ex
# build all plugins you can find
for pluginPath in ./searchprovider/*
do
    name=$(basename $pluginPath)
    plugin_path="searchprovider/$name/$name.go"

    if [ ! -e "${plugin_path}" ];
    then
        continue
    fi

    go build -buildmode=plugin -o $name.so searchprovider/$name/$name.go
done

go build -o downloader main.go

# create appropriate folder for plugins
fpm -f -s dir -t deb -n godownloader downloader=/usr/local/bin/
