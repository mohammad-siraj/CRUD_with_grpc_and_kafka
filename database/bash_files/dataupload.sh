#!bin/sh
docker cp SQL_files/car.sql postgresql:/car.sql
docker exec -u root postgresql psql root root -f car.sql