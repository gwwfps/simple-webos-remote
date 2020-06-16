package gui

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func topLevelTreeNode(label string, opened bool, children ...giu.Widget) giu.Widget {
	flags := imgui.TreeNodeFlagsCollapsingHeader
	if opened {
		flags = flags | imgui.TreeNodeFlagsDefaultOpen
	}
	return giu.TreeNode(label, giu.TreeNodeFlags(flags), children)
}
