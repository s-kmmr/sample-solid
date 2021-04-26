package okpattern

type IService interface {
	Notice(kind KindType) error
}
