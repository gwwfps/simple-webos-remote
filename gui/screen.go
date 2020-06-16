package gui

import "github.com/AllenDang/giu"

func (g *GUI) screen() giu.Widget {
	return topLevelTreeNode("Screens", true,
		giu.Checkbox("Enabled", &g.screenEnabled, g.onChangeScreenEnabled),
	)
}

func (g *GUI) onChangeScreenEnabled() {
	var err error
	if g.screenEnabled {
		err = g.tvm.EnableScreen()
	} else {
		err = g.tvm.DisableScreen()
	}
	g.updateErr(err)
}
