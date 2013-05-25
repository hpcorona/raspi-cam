package main

import (
	cv "github.com/hpcorona/gopencv"
	"log"
)

func main() {
	log.Printf("Running!")

	capture := cv.NewCaptureFromCAM(cv.CV_CAP_ANY)
	defer capture.Release()

	image := capture.QueryFrame()
	if image != nil {
		win := cv.NewWindow("Image")
		defer win.Destroy()
		win.ShowImage(image)
		//win.Resize(640, 480)
		cv.WaitKey(0)

		//cv.SaveImage("test.png", image, 0)
		log.Printf("WE HAVE IMAGE :)")
	} else {
		log.Printf("NO IMAGE :(")
	}
}
