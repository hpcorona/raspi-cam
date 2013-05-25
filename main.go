package main

import (
	"code.google.com/p/go-opencv/opencv"
	"log"
)

func main() {
	log.Printf("Running!")

	capture := opencv.NewCameraCapture(opencv.CV_CAP_ANY)
	defer capture.Release()

	image := capture.QueryFrame()
	if image != nil {
		win := opencv.NewWindow("Image")
		defer win.Destroy()
		win.ShowImage(image)
		//win.Resize(640, 480)
		opencv.WaitKey(0)

		//opencv.SaveImage("test.png", image, 0)
		log.Printf("WE HAVE IMAGE :)")
	} else {
		log.Printf("NO IMAGE :(")
	}
}
