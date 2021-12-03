NAME = login
TAG= 1.0.1
INSTANCE = login
DB_INSTANCE = db

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 10712:10712 --link $(DB_INSTANCE) -d $(NAME):$(TAG)