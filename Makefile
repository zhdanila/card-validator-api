up:
	ENV=dev go run cmd/server/main.go

swagger:
	swag init --parseDependency --parseInternal -g cmd/server/main.go
