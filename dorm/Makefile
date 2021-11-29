NAME = dorm
TAG= 1.0.0
INSTANCE = dorm
DB_INSTANCE = db

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 10713:10713 --link $(DB_INSTANCE) -e MYSQL_HOST="$(DB_INSTANCE)" -d $(NAME):$(TAG) # -e MYSQL_HOST="$(DB_INSTANCE)"