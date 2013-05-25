
GOHOME=$HOME/gocode
PATH=$PATH:$HOME/go/bin
	
all:
	git pull
	go clean
	go build -o raspi-cam
	./raspi-cam
