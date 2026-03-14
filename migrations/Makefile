up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://postgres:secret@localhost:5432/golang_sql?sslmode=disable goose up
down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://postgres:secret@localhost:5432/golang_sql?sslmode=disable goose down
createMigrationFile:
	goose create initial_databases sql
