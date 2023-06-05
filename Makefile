build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go

run: build
	docker-compose up --build server

migrate:
	migrate -path ./schema -database "mysql://root:<PASSWORD>@tcp(localhost:3306)/task-management" up