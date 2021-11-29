NAME = nodejs
TAG= 1.0.0
INSTANCE = nodejs
DB_INSTANCE = db
LOGIN_APP = login
REGISTER_APP = register
DORM_APP = dorm

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 8080:8080 --link $(DORM_APP) --link $(REGISTER_APP) --link $(LOGIN_APP) --link $(DB_INSTANCE) -e MYSQL_HOST="$(DB_INSTANCE)" -d $(NAME):$(TAG)
