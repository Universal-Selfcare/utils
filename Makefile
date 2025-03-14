include .env

.PHONY: migrate
run/api:
	go run . -db-dsn=${UNIVERSAL_SELFCARE_DB_DSN} 

