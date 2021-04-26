package ngpattern

// IAction 生物の行動を表す
type IAction interface {
	Eat()
	Run()
	Speak()
	Fly()
}
