#!/usr/bin/env bash

username="user"
password="password"
dbs="hrple,habit"

while getopts u:p:dbs: flag
do
    case "${flag}" in
        u) username=${OPTARG};;
        p) password=${OPTARG};;
		dbs) dbs=${OPTARG};;
    esac
done
docker run --name pg \
-e POSTGRES_USER=${username} \
-e POSTGRES_PASSWORD=${password} \
-e POSTGRES_MULTIPLE_DATABASES=${dbs} \
-p 5432:5432 \
-v $(pwd)/scripts/pg-init-scripts:/docker-entrypoint-initdb.d \
-d --rm postgres
exit 0
