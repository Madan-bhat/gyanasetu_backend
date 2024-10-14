BIN=gyanasetu-server 
run: 
	go run main.go

dev: $(air)
	air main.go


	
migrate-up: $(docker) 
	docker compose run --rm  migrate up 

migrate-down: $(docker)
	docker compose run --rm migrate down 

generate: $(docker) 
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate 

build:
	go build -o $(BIN) .
	
help:
	@echo "run - starts the server"
	@echo "dev - starts the dev server"
	@echo "migrate-up - performs create migrations(required go-migrate)"
	@echo "migrate-down - performs drop migrations(required go-migrate)"
	@echo "generate - performs sql to go generation(required sqlc)"
