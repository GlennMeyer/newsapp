build:
	docker build -t newsapp .
compose:
	docker-compose up --build --force-recreate
debug:
	docker run --rm -v ${CURDIR}:/usr/src/app -w /usr/src/app --name debug -it golang:1.13.0 /bin/bash
killapi:
	docker kill newsapp_api_1
killdb:
	docker kill newsapp_db_1
killdebug:
	docker kill debug
killrun:
	docker kill api
push:
	docker push glennmeyer/newsapp:latest
run:
	docker run --rm -p 80:8080 --name api newsapp
stopapi:
	docker stop newsapp_api_1
stopdb:
	docker stop newsapp_db_1
stopdebug:
	docker stop debug
stoprun:
	docker stop api
tag:
	docker tag newsapp glennmeyer/newsapp:latest
