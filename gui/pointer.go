package gui

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"image/color"
)

var keyMapping = map[glfw.Key]string{
	glfw.KeyLeft:   "LEFT",
	glfw.KeyRight:  "RIGHT",
	glfw.KeyUp:     "UP",
	glfw.KeyDown:   "DOWN",
	glfw.KeyEnter:  "ENTER",
	glfw.KeyEscape: "BACK",
}

func (g *GUI) buildPointerAreaWidget() {
	width := 960
	height := 540
	c := color.RGBA{255, 255, 255, 100}

	pos := giu.GetCursorScreenPos()
	end := pos.Add(image.Pt(width, height))
	canvas := giu.GetCanvas()

	giu.InvisibleButton("pointer", float32(width), float32(height), func() {
		g.updatePointerPos(pos)
		err := g.tvm.ClickPointer()
		g.updateErr(err)
	}).Build()
	if giu.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorNone)
		g.updatePointerPos(pos)
		g.checkKeyboardInput()
		c.A = 200

		if giu.IsMouseDoubleClicked(giu.MouseButtonLeft) {
			err := g.tvm.DialPointerSocket()
			g.updateErr(err)
		}
	}

	canvas.AddRectFilled(pos, end, c, 0, 0)
	pointerPos := pos.Add(g.pointerPos)
	pointerSize := 10
	pointerColor := color.RGBA{0, 0, 0, 255}
	canvas.AddLine(pointerPos.Sub(image.Pt(pointerSize, 0)), pointerPos.Add(image.Pt(pointerSize, 0)), pointerColor, 1)
	canvas.AddLine(pointerPos.Sub(image.Pt(0, pointerSize)), pointerPos.Add(image.Pt(0, pointerSize)), pointerColor, 1)
}

func (g *GUI) updatePointerPos(cursor image.Point) {
	oldPos := g.pointerPos

	x, y := glfw.GetCurrentContext().GetCursorPos()
	x -= float64(cursor.X)
	y -= float64(cursor.Y)

	g.pointerPos = image.Pt(int(x), int(y))

	err := g.tvm.UpdatePointerPos(g.pointerPos.Sub(oldPos).Mul(int(g.cfg.PointerSens)))
	g.updateErr(err)
}

func (g *GUI) pointer() giu.Widget {
	return giu.Custom(g.buildPointerAreaWidget)
}

func (g *GUI) checkKeyboardInput() {
	for key, button := range keyMapping {
		action := glfw.GetCurrentContext().GetKey(key)
		if action == glfw.Press {
			g.keysPressed[key] = true
		} else if action == glfw.Release {
			if pressed, ok := g.keysPressed[key]; ok && pressed {
				g.keysPressed[key] = false
				err := g.tvm.PressButton(button)
				if err != nil {
					g.updateErr(err)
					break
				}
			}
		}
	}
}
