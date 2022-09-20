start-pyroscope:
	docker run -it -p 4040:4040 pyroscope/pyroscope:latest server

run:
	go run main.go

poke-fibo:
	curl localhost:8080/fibo/45