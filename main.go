package main

import (
	cv "github.com/hpcorona/gopencv"
	"log"
)

func main() {
	log.Printf("Running!")

	win := cv.NewWindow("Image")
	defer win.Destroy()

	//capture := cv.NewCaptureFromCAM(cv.CV_CAP_ANY)
	capture := cv.NewCameraCapture(cv.CV_CAP_ANY)
	capture.SetProperty(cv.CV_CAP_PROP_FRAME_WIDTH, 320)
	capture.SetProperty(cv.CV_CAP_PROP_FRAME_HEIGHT, 180)
	//capture.SetProperty(cv.CV_CAP_PROP_FRAME_WIDTH, 1280)
	//capture.SetProperty(cv.CV_CAP_PROP_FRAME_HEIGHT, 720)
	defer capture.Release()

	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				image := capture.QueryFrame()
				if image != nil {
					win.ShowImage(image)
					log.Printf("WE HAVE IMAGE :)")
				} else {
					log.Printf("NO IMAGE :(")
				}
			}
		}
	}()

	cv.WaitKey(0)

	quit <- true
}
