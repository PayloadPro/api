#!/bin/bash

echo "Waiting for servers to be up"
sleep 10

HOSTPARAMS="--host cockroach-1 --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

$SQL -e "CREATE DATABASE IF NOT EXISTS payloadpro;"
$SQL -e "CREATE USER IF NOT EXISTS pp;"

$SQL -d payloadpro -e "CREATE TABLE IF NOT EXISTS 
    bins(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name VARCHAR(32),
        description VARCHAR(512),
        created TIMESTAMPTZ DEFAULT NOW(),
        remote_addr VARCHAR,
        INDEX (created)
    );"

$SQL -d payloadpro -e "CREATE TABLE IF NOT EXISTS 
    requests(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        bin UUID NOT NULL REFERENCES bins (id),
        method VARCHAR(7),
        content_type VARCHAR,
        content_length INT,
        created TIMESTAMPTZ DEFAULT NOW(),
        remote_addr VARCHAR,
        protocol VARCHAR,
        user_agent VARCHAR,
        body BYTES,
        INDEX (bin)
    );"

$SQL -d payloadpro -e "CREATE TABLE IF NOT EXISTS 
    stats(
        bin UUID PRIMARY KEY NOT NULL REFERENCES bins (id),
        total INT,
        get INT,
        post INT,
        put INT,
        patch INT,
        options INT,
        head INT,
        delete INT
    );"

$SQL -e "GRANT ALL ON payloadpro.bins, payloadpro.requests, payloadpro.stats TO pp;"
