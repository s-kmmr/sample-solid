package okpattern

import "golang.org/x/xerrors"

type Stable struct {
	i IService
}

func NewStable(_i IService) Stable {
	return Stable{
		i: _i,
	}
}

func (s Stable) Notice() error {
	if err := s.i.Notice(typeA); err != nil {
		return xerrors.Errorf("Notice(): %w", err)
	}
	return nil

}
