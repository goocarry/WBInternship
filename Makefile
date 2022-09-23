APP_BIN = build/

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

clean:
	rm -rf ./build || true

devdb:
	export POSTGRESQL_URL='postgres://wbl0user:postgres@localhost:5432/wbl0?sslmode=disable'

testdb:
	export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/wbl0_test?sslmode=disable'
