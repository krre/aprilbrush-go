package main

import (
	"github.com/krre/aprilbrush-go/app"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	app := app.NewApplication()
	app.Run()
}
