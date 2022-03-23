#!/bin/sh

migrate -path /migrations -database $SQLSERVER_CONNECTIONSTRING up
