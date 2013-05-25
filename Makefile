
all:
	git update
	go clean
	go build -o raspi-cam
	./raspi-cam
