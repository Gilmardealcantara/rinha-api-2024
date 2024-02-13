build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/server

build_docker: build
	# docker build -t golango_rinha .
	docker buildx build --platform linux/amd64 -t gilmardealcantara/golango_rinha:latest .

run_docker: build_docker
	# docker run -p=3000:3000 golango_rinha
	docker run -p=3000:3000 gilmardealcantara/golango_rinha:latest 
	

up: build_docker
	docker compose up --build 

db:
	export PGPASSWORD='123' && psql -d rinha -h localhost -p 5433 -U admin
