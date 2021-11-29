NAME = db
TAG= 1.0.0
INSTANCE = db

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 3336:3306 -e MYSQL_ROOT_PASSWORD=123456 -d $(NAME):$(TAG)