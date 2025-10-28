#!/bin/bash
set -e

module_prefix="plana"
package_name="$1"
version="$2"

mkdir -p ./tmp
curl -s https://ba.pokeguy.dev/$package_name/decompiled/$version/decompiled.zip -o ./tmp/decompiled.zip
unzip ./tmp/decompiled.zip -d ./tmp "DiffableCs/BlueArchive/**"
go run github.com/arisu-archive/bluearchive-togo/cmd/ba-togo@latest convert -i ./tmp/DiffableCs/BlueArchive -o protos -n "MX.NetworkProtocol" \
    -m "github.com/arisu-archive/$module_prefix-protos" \
    --map "FlatData=github.com/arisu-archive/$module_prefix-flatbuffers/go/flatdata"
echo -n "$version" > version.txt
rm -rf ./tmp
golangci-lint run --fix ./...
