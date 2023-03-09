#!/bin/sh
echo "TELEGRAM_APITOKEN=$1" >> bin/.env
echo "DEBUG=$2" >> bin/.env
./bin/main