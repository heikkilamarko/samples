#!/bin/sh

DB_NAME=todos

docker-compose exec -T sql-server /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $SA_PASSWORD -Q "CREATE DATABASE $DB_NAME;"
