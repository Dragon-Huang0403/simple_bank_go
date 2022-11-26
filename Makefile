cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

postgres: cmd-exists-docker
	docker run --name ${DATABASE_PROJECTNAME} -p 5432:5432 -e POSTGRES_USER=${DATABASE_USER} -e POSTGRES_PASSWORD=${DATABASE_PASS} -d postgres

createdb: cmd-exists-docker
	docker exec -it ${DATABASE_PROJECTNAME} createdb --username=${DATABASE_USER} --owner=${DATABASE_USER} ${DATABASE_TABLENAME}

dropdb: cmd-exists-docker
	docker exec -it ${DATABASE_PROJECTNAME} dropdb ${DATABASE_TABLENAME}

migrateup: cmd-exists-migrate
	migrate -path db/migration -database "${DATABASE_URL}" -verbose up

migratedown: cmd-exists-migrate
	migrate -path db/migration -database "${DATABASE_URL}" -verbose down

createmigrate: cmd-exists-migrate
	migrate create -ext sql -dir db/migration -seq schema

sqlc: cmd-exists-sqlc
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown createmigrate sqlc

ifneq (,$(wildcard ./.env))
    include .env
    export
endif
