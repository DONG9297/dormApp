NAME = register
TAG= 1.0.1
INSTANCE = register
DB_INSTANCE = db

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 10711:10711 --link $(DB_INSTANCE) -d $(NAME):$(TAG)