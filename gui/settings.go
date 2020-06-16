package gui

import (
	"fmt"
	"github.com/AllenDang/giu"
)

func (g *GUI) settings() giu.Widget {
	widgets := giu.Layout{
		giu.InputText("TV Address", 0, &g.cfg.TvAddr),
		giu.InputText("Key", 0, &g.cfg.ClientKey),
	}

	for i, input := range g.cfg.TvInputs {
		widgets = append(widgets, giu.InputText(fmt.Sprintf("%s Label", input.Id), 0, &g.cfg.TvInputs[i].Alias))
	}

	widgets = append(widgets, giu.Button("Save", g.onClickSettingsSave))
	return topLevelTreeNode("Settings", false, widgets)
}

func (g *GUI) onClickSettingsSave() {
	g.cfg.Save()
}