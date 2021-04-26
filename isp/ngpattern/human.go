package ngpattern

/*
人間を表現してみる
*/

// Human 人間
type Human struct{}

func NewHuman() IAction {
	return Human{}
}

func (h Human) Eat()   {}
func (h Human) Run()   {}
func (h Human) Speak() {}
func (h Human) Fly()   {} // 人間が飛ぶ？
