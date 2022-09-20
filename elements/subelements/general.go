package subelements

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

func calcNewSize(d []elements.Drawable) (w, h int) {
	w = 0
	h = 0
	tempX := 0 
	tempY := 0

	for i, v := range d {
		if i % 2 != 0 {
			tempX = tempX + (v.GetXRight() - v.GetXLeft())
			if tempX > w {
				w = tempX 
			}
			if tempY > (v.GetYTop() - v.GetYBot()) {
				h = h + tempY + 10
			} else {
				h = h + (v.GetYTop() - v.GetYBot()) + 10
			}
		} else {
			tempX = 30 + (v.GetXRight() - v.GetXLeft())
			tempY = (v.GetYTop() - v.GetYBot())
		}
	}
	return
}

func NotifyOfSizeChange(d elements.Drawable) {
	w, h := calcNewSize(d.GetSubelements())

	d.SetXRight(d.GetXLeft() + 30 + w)
	d.SetYBot(d.GetYTop() + 40 + h)

	if d.GetParent() != nil {
		NotifyOfSizeChange(d.GetParent())
	} 
}