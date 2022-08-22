#!/bin/bash

sudo id >/dev/null

FILE='main.tmp'
LOCAL_FILE='/usr/local/bin/lcc'
GO_MODULE_NAME=$(grep '^module' go.mod | awk '{printf "%s", $2}')
BUILD_FLAGS="-w -s -X $GO_MODULE_NAME/config.Debug=false"

# bump version before build
./dev/version/bump

echo go module name is $GO_MODULE_NAME
echo build binary from $FILE to $LOCAL_FILE
echo build with flags $BUILD_FLAGS
go build -ldflags="$BUILD_FLAGS" -o $FILE .

if [ ! -f "$FILE" ]; then
  echo $FILE does not exist
fi

echo move file to $LOCAL_FILE
sudo mv $FILE $LOCAL_FILE
