docker-up:
	docker-compose up

docker-down:
	docker-compose down 
db-create:
docker exec -it postgres createdb --username=root --owner=root PRODKOFE

postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mfarah -d postgres:14-alpine