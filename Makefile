lint:
	golangci-lint run

dev:
	export POSTGRES_URL='postgres://wbl0user:postgres@localhost:5432/wbl0?sslmode=disable' && go run cmd/main.go

test:
	export POSTGRES_URL='postgres://postgres:postgres@localhost:5432/wbl0_test?sslmode=disable' && go run cmd/main.go

pub:
	./publisher/publisher

batch: 
	./batch.sh
