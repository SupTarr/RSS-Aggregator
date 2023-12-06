#!/bin/sh

# Create new table by goose
cd sql/schema
goose postgres postgres://postgres:postgres@winhost:5432/rss-aggregator up | echo

# Create go function from sql by sqlc
cd ../..
sqlc generate | echo "$(date --rfc-3339='ns') Success to generate by sqlc"