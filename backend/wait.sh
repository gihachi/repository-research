#!/bin/sh

container_name=$1
container_port=$2

# user name and pass is not required

until mysqladmin ping --host ${container_name} --port ${container_port} --silent; do
    sleep 2
done

fresh main.go -c runner.conf
