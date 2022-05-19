#!/bin/sh

docker build -t showdownjs . && docker run -p 8080:8080 showdownjs