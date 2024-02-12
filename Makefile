build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/server

build_docker: build
	docker build -t golango_rinha .

run_docker: build_docker
	docker run -p=3000:3000 golango_rinha

up: build_docker
	docker compose up --build

