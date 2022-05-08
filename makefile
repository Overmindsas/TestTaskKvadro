create:
	docker build -t testkvado .

run:
	docker run --name db -p 8000:3306 -d testkvado

