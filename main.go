package main

import (
	"code.google.com/p/go-opencv/opencv"
	"log"
)

func main() {
	log.Printf("Running!")

	capture := opencv.NewFileCapture("/dev/video0")

	image := capture.QueryFrame()
	if image != nil {
		SaveImage("test.jpg", image, 0)
		log.Printf("WE HAVE IMAGE :)")
	} else {
		log.Printf("NO IMAGE :(")
	}

	image.Release()

	capture.Release()
}
