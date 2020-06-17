package gui

import (
	"github.com/AllenDang/giu"
)

func (g *GUI) media() giu.Widget {
	return topLevelTreeNode("Media", true,
		giu.Line(
			giu.Button("Vol-", g.onClickMediaVolDown),
			giu.Button("Vol+", g.onClickMediaVolUp),
			giu.Button("Mute", g.onClickMediaMute),
			giu.Button("Unmute", g.onClickMediaUnmute),
			giu.Button("|>", g.onClickMediaPlay),
			giu.Button("||", g.onClickMediaPause),
		),
	)
}

func (g *GUI) onClickMediaVolDown() {
	err := g.tvm.VolDown()
	g.updateErr(err)
}

func (g *GUI) onClickMediaVolUp() {
	err := g.tvm.VolUp()
	g.updateErr(err)
}

func (g *GUI) onClickMediaMute() {
	err := g.tvm.Mute()
	g.updateErr(err)
}

func (g *GUI) onClickMediaUnmute() {
	err := g.tvm.Unmute()
	g.updateErr(err)
}

func (g *GUI) onClickMediaPlay() {
	err := g.tvm.PlayMedia()
	g.updateErr(err)
}

func (g *GUI) onClickMediaPause() {
	err := g.tvm.PauseMedia()
	g.updateErr(err)
}
