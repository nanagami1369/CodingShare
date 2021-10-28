#!/bin/sh
sed -i "s#{{SERVER_NAME}}#$SERVER_NAME#g;" /etc/nginx/nginx.conf

nginx -g 'daemon off;'
