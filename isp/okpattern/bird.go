package okpattern

/*
鳥を表現する
*/

// Bird 鳥
type Bird struct{}

func NewBird() IBird {
	return Bird{}
}

func (b Bird) Eat() {} // 鳥として食べる
func (b Bird) Fly() {} // 鳥として飛ぶ
