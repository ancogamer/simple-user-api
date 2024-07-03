run:
	docker-compose up -d
	go run .
cover:
	go test -race -v -failfast -coverprofile=coverage.out ./...  ; go tool cover -html=coverage.out -o coverage.html ; open coverage.html 
deps: 
	docker compose up -d