#!/usr/bin/env bash
# wait-for-it.sh

set -e

# Store the first argument as the host
host="$1"

# Shift all arguments, so $@ contains the command to run
shift

# Store the rest of the arguments as the command to execute
cmd="$@"

# Loop until the MySQL server is available
until mysql -h "$host" -u root -proot -e 'SELECT 1'; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd
