package context

import (
	"fmt"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// g "github.com/AllenDang/giu"
)

func calcNewSize(d []elements.Element) (w, h int) {
	if len(d) == 0 {
		return 0, 0 
	}
	
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
	
	if len(temp) != 0 {
		w, h := calcNewSize(temp)
		d.SetXRight(d.GetXLeft() + 30 + w)
		d.SetYBot(d.GetYTop() + 40 + h)
		adjustInternalPositioning(d)
	} else {
		d.GetBaseWidth()
		d.SetXRight(d.GetXLeft() + d.GetBaseWidth())
		d.SetYBot(d.GetYTop() + d.GetBaseHeight())
	}

	if d.GetParent() != nil {
		NotifyOfSizeChange(d.GetParent())
	} else {
		updateOverviewElements()
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

func removeElementByID(e []*elements.Element, i int) []*elements.Element {
	// empty or last removing 
	if len(e) < 2 {
		return make([]*elements.Element, 0)
	}

	ret := make([]*elements.Element, len(e)-1)
	tempI := 0 
	for in, v := range e {
		if in != i {
			ret[tempI] = v
			tempI ++ 
		}
	}

	return ret
}

func findID(e []*elements.Element, el *elements.Element) (bool, int) {
	for i, v := range e {
		if v == el {
			return true, i
		}
	}
	return false, -1
}

func RemoveElement(e []elements.Element, el elements.Element) []elements.Element {
	// empty or last removing 
	if len(e) < 2 {
		fmt.Println("empty")
		return make([]elements.Element, 0)
	}
	ret := make([]elements.Element, len(e)-1)
	tempI := 0 
	for _, v := range e {
		if v.GetID() != el.GetID() {
			ret[tempI] = v
			tempI ++ 
		}
	}
	fmt.Println("created")
	return ret
}

func RecursiveLevelChange(startlevel int, el elements.Element) {
	el.SetLevel(startlevel)

	for _, v := range el.GetSubelements() {
		RecursiveLevelChange(startlevel + 1 , v)
	}
}

func updateOverviewElements() {
	pro := GetPro()
	pro.Over.Pressables = make([]OverViewElement, 0)
	for _, v := range pro.Level1 {
		recursiveAddToOverview(v)
	}
}

func recursiveAddToOverview(e elements.Element) {
	NewOverViewElement(e.GetName(), e.GetLevel())
	for _, v := range e.GetSubelements() {
		recursiveAddToOverview(v)
	}
}

