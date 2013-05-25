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
	fps := int(capture.GetProperty(cv.CV_CAP_PROP_FPS))
	if fps <= 0 {
		fps = 33
	}
	//capture.SetProperty(cv.CV_CAP_PROP_FRAME_WIDTH, 1280)
	//capture.SetProperty(cv.CV_CAP_PROP_FRAME_HEIGHT, 720)
	defer capture.Release()

	for {
		image := capture.QueryFrame()
		if image != nil {
			win.ShowImage(image)
			log.Printf("WE HAVE IMAGE :)")
		} else {
			log.Printf("NO IMAGE :(")
		}

		key := cv.WaitKey(1000 / fps)
		if key == 27 {
			break
		}
	}
}
