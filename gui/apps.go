package gui

import "github.com/AllenDang/giu"

func (g *GUI) apps() giu.Widget {
	return topLevelTreeNode("Apps", true,
		giu.Line(
			giu.Button("Youtube", g.onClickAppsYoutube),
			giu.Button("Twitch", g.onClickAppsTwitch),
			giu.Button("Netflix", g.onClickAppsNetflix),
			giu.Button("Prime", g.onClickAppsPrime),
			giu.Button("AppleTV", g.onCLickAppsAppleTV),
		),
		giu.InputText("Youtube Link", 0, &g.youtubeLink))
}

func (g *GUI) onClickAppsYoutube() {

}

func (g *GUI) onClickAppsTwitch() {

}

func (g *GUI) onClickAppsNetflix() {

}

func (g *GUI) onClickAppsPrime() {

}

func (g *GUI) onCLickAppsAppleTV() {

}
