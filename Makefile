docker_build:
	docker build -t dreamdata:latest .

docker_run: docker_build
	docker container run -it -p 8000:8000 --env-file .env dreamdata:latest

run_backend:
	go run cmd/*.go

run_frontend:
	cd frontend && npm start
