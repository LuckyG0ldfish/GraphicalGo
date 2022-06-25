package ui

import g "github.com/AllenDang/giu"

var ( 
	sashPos1 float32 = 200
	sashPos2 float32 = 200
	sashPos3 float32 = 200
	sashPos4 float32 = 100
)

func Build(wnd *g.MasterWindow) {
	wnd.Run(iterate)
}

func iterate() {
	g.SingleWindow().Layout(
		g.SplitLayout(g.DirectionHorizontal, sashPos1,
			g.Layout{
				g.Label("Left panel"),
				g.Row(g.Button("Button1"), g.Button("Button2")),
			},
			g.SplitLayout(g.DirectionVertical, sashPos2,
				g.Layout{},
				g.SplitLayout(g.DirectionHorizontal, sashPos3,
					g.Layout{},
					g.SplitLayout(g.DirectionVertical, sashPos4,
						g.Layout{},
						g.Layout{},
					),
				),
			),
		),
	)
}
