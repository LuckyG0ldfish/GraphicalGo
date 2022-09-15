package subelements

type Button struct {
	name 		string 
	pos 		int
	function 	func()

	xRight		int
	yBot		int

	xLeft		int
	yTop		int
}

func CreateButton(name string, pos int, fun func()) *Button{
	var but Button 
	but.name = name 
	but.pos = pos
	but.function = fun 

	yVar := 10 + pos*Dist
	mix := pro.Can.XRight
	max := pro.Win.XWidth
	
	but.xLeft = mix+5
	but.xRight = max-5

	but.yTop = yVar
	but.yBot = yVar+20

	pro.Obj.Pressables = append(pro.Obj.Pressables, &but)

 	return &but
}

func (but *Button) Press() {
	but.function()
}

func (but *Button) GetPos() int {
	return but.pos
}

func (but *Button) GetName() string {
	return but.name
}

// func (but *Button) SetName(name string){
// 	but.name = name 
// }

// func (but *Button) SetFunc(fun func) {
// 	but.func = fun 
// }

func ObjectButton(Pro Project) { 
	Pro.Can.Dragables = append(Pro.Can.Dragables, CreateObject("test", 200)) 
}

func (but *Button) GetXRight() int{
	return but.xRight
}

func (but *Button) GetYBot() int{
	return but.yBot
}

func (but *Button) GetXLeft() int{
	return but.xLeft
}

func (but *Button) GetYTop() int{
	return but.yTop
}