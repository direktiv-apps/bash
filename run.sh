#!/bin/sh

docker build -t bash . && docker run -p 8080:8080 bash