package gui

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/gwwfps/simple-webos-remote/tvmanager"
)

func (g *GUI) apps() giu.Widget {
	return topLevelTreeNode("Apps", true,
		giu.Line(
			giu.Button("Youtube", g.onClickAppsYoutube),
			giu.Button("Twitch", g.onClickAppsTwitch),
			giu.Button("DAZN", g.onClickAppsDAZN),
		),
		giu.Line(
			giu.Button("Netflix", g.onClickAppsNetflix),
			giu.Button("Prime", g.onClickAppsPrime),
			giu.Button("AppleTV", g.onClickAppsAppleTV),
		),
		giu.InputTextV("Youtube Link", 0, &g.youtubeLink, imgui.InputTextFlagsEnterReturnsTrue, nil, g.onChangeAppsYoutubeLink))
}

func (g *GUI) onClickAppsYoutube() {
	err := g.tvm.LaunchApp(tvmanager.AppIdYoutube)
	g.updateErr(err)
}

func (g *GUI) onClickAppsTwitch() {
	err := g.tvm.LaunchApp(tvmanager.AppIdTwitch)
	g.updateErr(err)
}

func (g *GUI) onClickAppsNetflix() {
	err := g.tvm.LaunchApp(tvmanager.AppIdNetflix)
	g.updateErr(err)
}

func (g *GUI) onClickAppsPrime() {
	err := g.tvm.LaunchApp(tvmanager.AppIdPrime)
	g.updateErr(err)
}

func (g *GUI) onClickAppsAppleTV() {
	err := g.tvm.LaunchApp(tvmanager.AppIdAppleTV)
	g.updateErr(err)
}

func (g *GUI) onClickAppsDAZN() {
	err := g.tvm.LaunchApp(tvmanager.AppIdAppleDAZN)
	g.updateErr(err)
}

func (g *GUI) onChangeAppsYoutubeLink() {
	err := g.tvm.OpenYoutubeURL(g.youtubeLink)
	g.updateErr(err)
}
