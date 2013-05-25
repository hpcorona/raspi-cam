
all:
	go get -u github.com/hpcorona/gopencv
	git pull
	go build -o raspi-cam
	./raspi-cam
