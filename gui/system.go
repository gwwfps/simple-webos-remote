package gui

import "github.com/AllenDang/giu"

func (g *GUI) system() giu.Widget {
	var widgets giu.Layout
	if g.tvm.Connected() {
		widgets = append(widgets, giu.Line(
			giu.Button("Power Off", g.onClickSystemPowerOff),
			giu.Button("Info", g.onClickSystemInfo),
			giu.Button("Settings", g.onClickSystemSettings),
		))
	} else {
		widgets = append(widgets, giu.Button("Power On", g.onClickSystemPowerOn))
	}
	return giu.Layout{
		topLevelTreeNode("Power", true, widgets),
	}
}

func (g *GUI) onClickSystemPowerOn() {
	err := g.tvm.PowerOn()
	g.updateErr(err)
}

func (g *GUI) onClickSystemPowerOff() {
	err := g.tvm.PowerOff()
	g.updateErr(err)
}

func (g *GUI) onClickSystemInfo() {
	err := g.tvm.OpenInfo()
	g.updateErr(err)
}

func (g *GUI) onClickSystemSettings() {
	err := g.tvm.OpenSettings()
	g.updateErr(err)
}
