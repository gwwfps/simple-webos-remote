package gui

import "github.com/AllenDang/giu"

func (g *GUI) power() giu.Widget {
	var button giu.Widget
	if g.tvm.Connected() {
		button = giu.Button("Off", g.onClickPowerOff)
	} else {
		button = giu.Button("On", g.onClickPowerOn)
	}
	return giu.Layout{
		topLevelTreeNode("Power", true, button),
	}
}

func (g *GUI) onClickPowerOn() {
	err := g.tvm.PowerOn()
	g.updateErr(err)
}

func (g *GUI) onClickPowerOff() {
	err := g.tvm.PowerOff()
	g.updateErr(err)
}
