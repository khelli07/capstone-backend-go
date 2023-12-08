#!/bin/bash

docker run -d -p 80:$1 --env-file .env $2

# bash dockerrun.sh 9999 myimage:latest