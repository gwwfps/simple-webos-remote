package gui

import (
	"github.com/AllenDang/giu"
	"image"
	"image/color"
)

type pointerAreaWidget struct {
}

func (w *pointerAreaWidget) Build() {
	pos := giu.GetCursorPos()
	end := pos.Add(image.Pt(960, 540))
	canvas := giu.GetCanvas()
	canvas.AddRectFilled(pos, end, color.RGBA{255, 255, 255, 255}, 0, 0)
}

func (g *GUI) pointer() giu.Widget {
	return topLevelTreeNode("Pointer", false, &pointerAreaWidget{})
}
