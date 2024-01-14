
all:
	go build -C src -o app

run:
	make && ./app

