#!/bin/sh

SCRIPT_DIR=$(
    cd $(dirname $0)
    pwd
)
mysql -u$DB_USER -h$DB_IP -P$DB_PORT -p$DB_PASSWORD -e "DROP DATABASE IF EXISTS account;
CREATE DATABASE account;"
