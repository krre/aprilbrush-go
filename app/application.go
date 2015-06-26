package app

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/krre/aprilbrush-go/brushengine"
	"github.com/krre/aprilbrush-go/glutil"
)

type Application struct {
	window   *glfw.Window
	renderer *glutil.Renderer
}

func NewApplication() *Application {
	application := Application{}

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()
	screenWidth := videoMode.Width
	screenHeight := videoMode.Height
	screenRatio := 0.8

	width := int(screenRatio * float64(screenWidth))
	height := int(screenRatio * float64(screenHeight))
	x := int((screenWidth - width) / 2)
	y := int((screenHeight - height) / 2)

	application.window, err = glfw.CreateWindow(width, height, "AprilBrush", nil, nil)
	if err != nil {
		panic(err)
	}
	application.window.SetPos(x, y)
	application.window.MakeContextCurrent()

	application.window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		mouseButtonState := application.window.GetMouseButton(glfw.MouseButtonLeft)
		if mouseButtonState == glfw.Press {
			brushengine.Paint(xpos, ypos)
		}
	})

	application.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			if mods == glfw.ModControl {
				switch key {
				case glfw.KeyN:
					fmt.Println("New")
				case glfw.KeyO:
					fmt.Println("Open")
				case glfw.KeyS:
					fmt.Println("Save")
				case glfw.KeyD:
					fmt.Println("Save as")
				case glfw.KeyQ:
					application.window.SetShouldClose(true)
				}
			} else {

			}
		}
	})

	camera := glutil.NewCamera()
	scene := glutil.NewScene()
	application.renderer = glutil.NewRenderer(scene, camera)

	return &application
}

func (application *Application) Run() {
	defer glfw.Terminate()

	for !application.window.ShouldClose() {
		application.renderer.Render()

		// Do OpenGL stuff
		application.window.SwapBuffers()
		glfw.PollEvents()
	}
}
