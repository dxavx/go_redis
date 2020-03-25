#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm go_redis_test_api