up:
	docker-compose up --build -d

down:
	docker-compose down


run:
	nodemon --exec go run main.go --signal SIGTERM

r:
	go run main.go
	