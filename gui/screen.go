package gui

import "github.com/AllenDang/giu"

func (g *GUI) screen() giu.Widget {
	enabled := false
	return topLevelTreeNode("Screens", true,
		giu.Checkbox("Enabled", &enabled, g.onChangeScreenEnabled),
	)
}

func (g *GUI) onChangeScreenEnabled() {

}
