build-app:
	go mod download
	go build -o server server.go
run-app:
	./server
docker-build:
	docker build -t app .
docker-run:
	docker run app
