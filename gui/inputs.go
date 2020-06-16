package gui

import "github.com/AllenDang/giu"

func (g *GUI) inputs() giu.Widget {
	var buttons giu.Layout
	for _, input := range g.cfg.TvInputs {
		buttons = append(buttons, giu.Button(input.Alias, g.buildOnClickInputsHandler(input.Id)))
	}
	return topLevelTreeNode("Inputs", true, giu.Line(buttons...))
}

func (g *GUI) buildOnClickInputsHandler(id string) func() {
	return func() {
		err := g.tvm.TV().TvSwitchInput(id)
		g.updateErr(err)
	}
}
