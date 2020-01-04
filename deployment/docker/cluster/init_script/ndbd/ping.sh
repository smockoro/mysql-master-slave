#!/bin/sh

while ! mysqladmin ping -h management1 --silent; do
  sleep 1
done
