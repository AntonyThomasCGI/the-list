
build:
	docker build -t the-list .

run:
	docker run -p 127.0.0.1:8080:8080 -it --rm --name the-list-app the-list

