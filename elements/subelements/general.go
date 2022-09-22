package subelements

import (
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// g "github.com/AllenDang/giu"
)

func calcNewSize(d []elements.Element) (w, h int) {
	w = 0
	h = 0
	tempX := 0
	tempY := 0
	prev := false

	for i, v := range d {
		if i%2 != 0 {
			tempX = tempX + (v.GetXRight() - v.GetXLeft())
			if tempX > w {
				w = tempX
			}
			if tempY > (v.GetYBot() - v.GetYTop()) {
				h = h + tempY + 10
			} else {
				h = h + (v.GetYBot() - v.GetYTop()) + 10
			}
			prev = false
		} else {
			tempX = 30 + (v.GetXRight() - v.GetXLeft())
			tempY = (v.GetYBot() - v.GetYTop())
			prev = true
		}
	}
	if prev {
		if tempX > w {
			w = tempX
		}
		h = h + tempY
	}
	return
}

func NotifyOfSizeChange(d elements.Element) {
	if d == nil {
		return
	}
	temp := d.GetSubelements()
	if len(temp) == 0 {
		return
	}
	w, h := calcNewSize(temp)
	d.SetXRight(d.GetXLeft() + 30 + w)
	d.SetYBot(d.GetYTop() + 40 + h)
	adjustInternalPositioning(d)
	if d.GetParent() != nil {
		NotifyOfSizeChange(d.GetParent())
	}
}

func adjustInternalPositioning(d elements.Element) {
	xdist := 10
	ydist := 10
	xpos := d.GetXLeft() + xdist
	ypos := d.GetYTop() + ydist*3
	prevX := 0
	prevY := 0

	for i, v := range d.GetSubelements() {
		if i%2 == 0 {
			adjustSubElement(xpos, ypos, v)
			adjustInternalPositioning(v)
			prevX = v.GetXRight()
			prevY = v.GetYBot()
		} else {
			adjustSubElement(prevX+xdist, ypos, v)
			adjustInternalPositioning(v)
			prevX = 0
			if prevY > v.GetYBot() {
				ypos = prevY + ydist
			} else {
				ypos = v.GetYBot() + ydist
			}
		}
	}
}

func adjustSubElement(xleft int, yTop int, d elements.Element) {
	xdist := d.GetXRight() - d.GetXLeft()
	ydist := d.GetYBot() - d.GetYTop()

	d.SetXLeft(xleft)
	d.SetYTop(yTop)
	d.SetXRight(xleft + xdist)
	d.SetYBot(yTop + ydist)
}
