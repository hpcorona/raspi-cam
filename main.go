package main

import (
	"code.google.com/p/go-opencv/opencv"
	"log"
)

func main() {
	log.Printf("Running!")

	capture := opencv.NewCameraCapture(opencv.CV_CAP_ANY)

	image := capture.QueryFrame()
	if image != nil {
		opencv.SaveImage("test.pxm", image, opencv.CV_IMWRITE_PXM_BINARY)
		log.Printf("WE HAVE IMAGE :)")
	} else {
		log.Printf("NO IMAGE :(")
	}

	capture.Release()
}
