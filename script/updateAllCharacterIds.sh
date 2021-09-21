#!/bin/bash
set -o xtrace

PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

rm /tmp/allcharacters.json

cd /app/script

go run updateAllCharacterIds.go /tmp/allcharacters.json /app/static/total

if [ $? -ne 0 ]; then
   echo "Update failed, exiting"
   exit 1
fi

if [[ -f /tmp/allcharacters.json ]]; then
    mv /tmp/allcharacters.json /app/static/allcharacters.json
fi
