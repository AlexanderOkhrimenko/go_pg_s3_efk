#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm go_pg_s3_efk_api
docker image rm go_pg_s3_efk_worker