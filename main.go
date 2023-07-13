package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

var width int = 600
var height int = 600

type Color struct {
	r uint8
	g uint8
	b uint8
}

type Pixel struct {
	X int
	Y int
	C Color
}

type Display struct {
	window   *sdl.Window
	surface  *sdl.Surface
	renderer *sdl.Renderer
	texture  *sdl.Texture
	pixels   []byte
	pitch    int
}

func (d Display) SetPixel(p Pixel) {
	index := p.Y*(d.pitch) + p.X*4
	d.pixels[index] = p.C.b
	d.pixels[index+1] = p.C.g
	d.pixels[index+2] = p.C.r
	d.pixels[index+3] = 255

}

func MakeDisplay() Display {
	window, err := sdl.CreateWindow("help", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, int32(width), int32(height))
	pixels, pitch, _ := tex.Lock(nil)

	return Display{window, surface, renderer, tex, pixels, pitch}
}

func TracePath(r Ray, depth int) {

}

func main() {
	defer sdl.Quit()
	displayobject := MakeDisplay()
	origin := V3{0, 0, 0}
	topleftrenderplane := V3{-1, -1, -1}

	//pixels, pitch, _ := displayobject.texture.Lock(nil)

	for true {

		for x := 0; x < width; x++ {
			px := float64(x) / float64(width)
			xscanvector := V3Add(topleftrenderplane, V3MultByFloat(V3{2, 0, 0}, px))

			for y := 0; y < height; y++ {
				py := float64(y) / float64(height)

				fullscan := V3Add(xscanvector, V3MultByFloat(V3{0, 2, 0}, py))

				newray := CreateRay(origin, fullscan)
				fmt.Println(fullscan)
				fmt.Println(newray)

			}
		}

		displayobject.texture.Update(nil, unsafe.Pointer(&displayobject.pixels[0]), displayobject.pitch)
		// displayobject.texture.Unlock()

		displayobject.renderer.Copy(displayobject.texture, nil, nil)
		displayobject.window.UpdateSurface()
	}
	displayobject.window.Destroy()
}
