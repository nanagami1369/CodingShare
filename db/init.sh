#!/bin/sh

SCRIPT_DIR=$(
    cd $(dirname $0)
    pwd
)
mysql -u$DB_USER -h$DB_IP -P$DB_PORT -p$DB_PASSWORD <$SCRIPT_DIR/schema.sql
mysql -u$DB_USER -h$DB_IP -P$DB_PORT -p$DB_PASSWORD <$SCRIPT_DIR/sample_data.sql
mysql -u$DB_USER -h$DB_IP -P$DB_PORT -p$DB_PASSWORD -e "use $DB_NAME;SELECT * FROM users;SELECT * FROM admins"
