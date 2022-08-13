<div align="center">
  <a href="https://github.com/mohammad-siraj/CRUD_operations_with_grpc">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>
  <h3 align="center">CRUD_operations_with_grpc</h3>
  <p align="center">
    CRUD api implementation in golang with grpc for car database hosted in a dokcer container (Postgres) and kafka for publishing it on a topic
    <br />
    <br />
  </p>
  <p align="left">
  </p>
</div>
<!-- ABOUT THE PROJECT -->

## Running the database container:
To start the docker container for Postgres instance by running the dokcer_compose file in 'dockerfiles/docker_compose.yaml':

```console
CRUD_grpc\dockerfiles> docker compose -f docker_compose.yaml up
```
## Table schema and Data upload in the container for cars
To upload the table schema and data for cars to  Postgres instance use the bash file in  'database/bash_files':

```console
CRUD_grpc\dockerfiles\bash_files> bash dataupload.sh
```
## Protobuf schema
To genrate grpc protocall we need protobuf schema for go lang code genartion from '.proto' file.To generate the go lang code we need to execute the following the terminal

```console
CRUD_grpc> protoc --go_out=crud_proto  --go-grpc_out=crud_proto crud_proto/crud_proto.proto
```

## Requirements:
* docker >= 17.12.0+
* docker-compose
* gorilla/mux
* lib/pq
* grpc
* protobuf
* kafka
* gcc

## About The Project
This is a CRUD grpc server written in go lang on a car (params: model,make,year) database hosted in a Docker container running Postgres instance and using grpc package for grpc server hosting and used postman to verify its behaviour along with protobuf for code generation
