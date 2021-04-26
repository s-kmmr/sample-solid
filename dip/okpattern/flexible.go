package okpattern

import (
	"golang.org/x/xerrors"
)

type Flexible struct {
}

func NewFlexible() IService {
	return Flexible{}
}

// Notice IServiceの実装：種類によって通知を分ける
func (f Flexible) Notice(kind KindType) error {
	switch kind {
	case typeA:
		println("notice typeA")
		return nil
	case typeB:
		println("notice typeB")
		return nil
	}
	return xerrors.Errorf("don't match kind: %v", kind)
}
