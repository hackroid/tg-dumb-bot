#!/bin/sh
echo "TELEGRAM_APITOKEN=$1" >> .env
echo "DEBUG=$2" >> .env
./bin/main