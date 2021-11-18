# For this you need to have docker installed on you computer.

# Build docker image
make docker_build

# Run dreamdata image in container
make docker_run

# To run without docker you need to have npm installed and Go installed on you computer.

# Run go backend
make run_backend

# Run react frontend
make run_frontend