package gui

import (
	"fmt"
	"github.com/AllenDang/giu"
)

func (g *GUI) settings() giu.Widget {
	widgets := giu.Layout{
		giu.InputText("TV Address", 0, &g.cfg.TvAddr),
		giu.InputText("TV Mac", 0, &g.cfg.TvMac),
		giu.InputText("Key", 0, &g.cfg.ClientKey),
		giu.InputText("Mutter Screen Name", 0, &g.cfg.ScreenName),
		giu.InputText("Sonos Address", 0, &g.cfg.SonosAddr),
		giu.DragIntV("Pointer Sensitivity", &g.cfg.PointerSens, 1.0, 1, 10, "%d"),
	}

	for i, input := range g.cfg.TvInputs {
		widgets = append(widgets, giu.InputText(fmt.Sprintf("%s Label", input.Id), 0, &g.cfg.TvInputs[i].Alias))
	}

	widgets = append(widgets, giu.Button("Save", g.onClickSettingsSave))
	return topLevelTreeNode("Config", false, widgets)
}

func (g *GUI) onClickSettingsSave() {
	g.cfg.Save()
}
