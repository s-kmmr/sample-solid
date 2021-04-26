package ngpattern

import "golang.org/x/xerrors"

type Stable struct {
	f Flexible
}

func NewStable(_f Flexible) Stable {
	return Stable{
		f: _f,
	}
}

func (s Stable) Notice() error {
	if err := s.f.Notice(typeA); err != nil {
		return xerrors.Errorf("Notice(): %w", err)
	}
	return nil

}
