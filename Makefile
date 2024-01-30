migrate-create:
	migrate create -ext sql -dir internal/database/migrations -seq ${NAME}

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/database/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path internal/database/migrations down

migrate-version:
	migrate -database ${POSTGRESQL_URL} -path internal/database/migrations version

migrate-force:
	migrate -database ${POSTGRESQL_URL} -path internal/database/migrations force ${VERSION}