
all:
	go build -C src -tags netgo -ldflags '-s -w' -o ../app

run:
	make && ./app

