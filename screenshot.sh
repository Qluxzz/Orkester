#!/bin/bash

docker compose up -d

sleep 5s

node take_screenshots.mjs

docker compose down