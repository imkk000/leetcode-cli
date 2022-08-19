#!/bin/bash

sudo id > /dev/null

FILE='main.tmp'
LOCAL_FILE='/usr/local/bin/lcc'

go build -ldflags="-w -s" -o $FILE .

if [ ! -f "$FILE" ]; then
    echo "$FILE does not exist."
fi

sudo mv $FILE $LOCAL_FILE
