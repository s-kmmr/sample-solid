package ngpattern

/*
鳥を表現する
*/

// Bird 鳥
type Bird struct{}

func NewBird() IAction {
	return Bird{}
}
func (b Bird) Eat()   {}
func (b Bird) Run()   {} // 鳥が走る？
func (b Bird) Speak() {} // 鳥が喋る？
func (b Bird) Fly()   {}
