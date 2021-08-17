BINARY=engine
dev:
	go test -v -cover -covermode=atomic ./...

test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go


unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker compose up

start_container:
	docker build --tag diffme-api ./deployments/Dockerfile.server && docker diffme-api

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint