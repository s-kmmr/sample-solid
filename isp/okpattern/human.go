package okpattern

/*
人間を表現してみる
*/

// Human 人間
type Human struct{}

func NewHuman() IHuman {
	return Human{}
}

func (o Human) Eat()   {} // 人間として食べる
func (o Human) Run()   {} // 人間として走る
func (o Human) Speak() {} // 人間として話す
