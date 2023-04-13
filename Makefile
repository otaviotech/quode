test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

testcoverage:
	make test && make coverage

migrate:
	migrate -path=internal/infra/db/sql/migrations -database "postgres://quode:quode@localhost:5432/quode?sslmode=disable" up

migratetest:
	migrate -path=internal/infra/db/sql/migrations -database "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable" up

migratedown:
	migrate -path=internal/infra/db/sql/migrations -database "postgres://quode:quode@localhost:5432/quode?sslmode=disable" down 

migratetestdown:
	migrate -path=internal/infra/db/sql/migrations -database "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable" down 

createmigration:
	migrate create -ext sql -dir=internal/infra/db/sql/migrations -seq $(name)

start:
	go run cmd/main.go cmd/wire_gen.go

.PHONY: test coverage testcoverage migrate migratedown createmigration start