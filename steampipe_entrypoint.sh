#!/bin/bash

set -eoux pipefail

rsync -chavzP --stats /local/.steampipe/ ~/.steampipe/

mkdir --parents ~/.nomad
cp -v /secrets/.nomad/credentials ~/.nomad/credentials

mkdir --parents ~/.steampipe/db/12.1.0/postgres

# rm -rvf ~/.steampipe/db/12.1.0/data
# cp -v /secrets/postgres/.passwd ~/.steampipe/db/12.1.0/postgres/.passwd
steampipe service start 2>&1 | tee service_start.log

host="$(\
  grep 'Host(s):' service_start.log \
    | cut -d':' -f2- \
    | tr ',' $'\n' \
    | grep -Ev 'localhost|127.0.0.1' \
    | awk '{print $1;}' \
)"
port="$(grep 'Port:' service_start.log | awk '{print $2;}')"
database="$(grep 'Database:' service_start.log | awk '{print $2;}')"
user="$(grep 'User:' service_start.log | awk '{print $2;}')"
password="$(grep 'Password:' service_start.log | awk '{print $2;}')"
echo "postgres://$user:$password@$host:$port/$database?sslmode=disable" | tee /alloc/database_uri.txt

# steampipe service restart --force
tail --follow ~/.steampipe/logs/*.log
#| grep -P '\[(WARN|ERROR)\]' &

# ls -lah ~/.steampipe/internal
service_pid="$(jq -r '.Pid' ~/.steampipe/internal/steampipe.json)"
echo "Now waiting for service process (pid: $service_pid) to exit..."
tail --pid="$service_pid" -f /dev/null
