# ECOMMERCE MARKETPLACE



# Database Schema Migration and Dumping

- We are using the tool db mate for dumping the current schema and also perform migrations from now on

## Check connection & migration status

`docker compose run --rm dbmate status`

## Dump current remote schema to ./db/schema.sql

`docker compose run --rm dbmate dump`

## Create a new migration file (creates in ./db/migrations/ on your machine)

`docker compose run --rm dbmate new migration_name`
