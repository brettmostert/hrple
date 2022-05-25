#!/usr/bin/env bash

username="user"
password="password"
db="hrple"

while getopts u:p:db: flag
do
    case "${flag}" in
        u) username=${OPTARG};;
        p) password=${OPTARG};;
		db) db=${OPTARG};;
    esac
done

docker run --name pg -e POSTGRES_USER=${username} -e POSTGRES_PASSWORD=${password} -e POSTGRES_DB=${db} -p 5432:5432 -d --rm postgres
exit 0
