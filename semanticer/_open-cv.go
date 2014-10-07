package semanticer

import (
	cv "github.com/hybridgroup/go-opencv/opencv"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/opencv"
)

import "runtime"

func Camera() {
	// Arrange that main.main runs on main thread.
	// See http://code.google.com/p/go-wiki/wiki/LockOSThread
	runtime.LockOSThread()

	gbot := gobot.NewGobot()

	window := opencv.NewWindowDriver("window")
	camera := opencv.NewCameraDriver("camera", 0)

	work := func() {
		gobot.On(camera.Events["Frame"], func(data interface{}) {
			window.ShowImage(data.(*cv.IplImage))
		})
	}

	gbot.Robots = append(gbot.Robots,
		gobot.NewRobot("cameraBot", []gobot.Connection{}, []gobot.Device{window, camera}, work))

	gbot.Start()
}
