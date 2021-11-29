NAME = login
TAG= 1.0.0
INSTANCE = login
DB_INSTANCE = db

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 10712:10712 --link $(DB_INSTANCE) -e MYSQL_HOST="$(DB_INSTANCE)" -d $(NAME):$(TAG) # -e MYSQL_HOST="$(DB_INSTANCE)"