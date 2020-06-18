package gui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/gwwfps/simple-webos-remote/config"
	"github.com/gwwfps/simple-webos-remote/tvmanager"
	"image"
	"time"
)

type GUI struct {
	tvm *tvmanager.TVManager
	cfg *config.Config

	youtubeLink   string
	screenEnabled bool
	pointerPos    image.Point
	keysPressed   map[glfw.Key]bool

	lastErr error
}

func New(tvm *tvmanager.TVManager, cfg *config.Config) *GUI {
	return &GUI{
		tvm: tvm,
		cfg: cfg,

		keysPressed: make(map[glfw.Key]bool),
	}
}

func (g *GUI) Run() {
	wnd := giu.NewMasterWindow("Remote", 1600, 600, giu.MasterWindowFlagsNotResizable, g.loadFont)
	g.style()
	go g.refresh()
	wnd.Main(g.loop)
}

func (g *GUI) loop() {
	widgets := giu.Layout{
		g.status(),
		g.system(),
	}

	if g.tvm.Connected() {
		widgets = append(widgets,
			g.inputs(),
			g.media(),
			g.apps(),
		)
	}

	widgets = append(widgets, g.settings())

	var windowFlags giu.WindowFlags = imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove | imgui.WindowFlagsNoResize
	giu.WindowV("Controls", nil, windowFlags, 0, 0, 640, 600, widgets)
	if g.tvm.Connected() {
		giu.WindowV("Pointer", nil, windowFlags, 640, 0, 960, 600, giu.Layout{g.pointer()})
	}
}

func (g *GUI) refresh() {
	ticker := time.NewTicker(time.Second)
	for {
		giu.Update()
		<-ticker.C
	}
}

func (g *GUI) status() giu.Widget {
	text := "Connected"
	if !g.tvm.Connected() {
		if g.tvm.ConnectionErr != nil {
			text = fmt.Sprintf("Cannot connect: %s", g.tvm.ConnectionErr.Error())
		} else if g.tvm.Connecting {
			text = "Connecting..."
		}
	}

	lastError := ""
	if g.lastErr != nil {
		lastError = g.lastErr.Error()
	}

	return giu.Layout{
		giu.Label(text),
		giu.PopupModal("Error", giu.Layout{
			giu.Label(lastError),
			giu.Button("OK", imgui.CloseCurrentPopup),
		}),
	}
}

func (g *GUI) updateErr(err error) {
	g.lastErr = err
	if err != nil {
		giu.Update()
		giu.OpenPopup("Error")
	}
}
