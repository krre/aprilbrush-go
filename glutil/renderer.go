package glutil

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	renderer := Renderer{}

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	gl.ClearColor(0.6, 0.6, 0.6, 1.0)

	return &renderer
}

func (renderer *Renderer) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
