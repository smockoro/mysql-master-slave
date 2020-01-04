#!/bin/sh

container_name="ndb"

hostlist=`cat ../docker-compose.yml | grep container_name | grep ${container_name} | cut -d " " -f 6`
host_counter=`count -l ${hostlist}`

while ! mysqladmin ping -h management1 --silent; do
  sleep 1
done

# ndb host list
for host in ${hostlist}; do
  while ! mysqladmin ping -h ${host} --silent; do
    sleep 1
  done
done
