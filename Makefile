up:
	ENV=dev go run cmd/server/main.go

swagger:
	swag init --parseDependency --parseInternal -g cmd/server/main.go

docker-build:
	docker build -t card-validator-api . && docker run --rm -p 8080:8080 card-validator-api
