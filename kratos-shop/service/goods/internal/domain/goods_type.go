package domain

type GoodsType struct {
	ID        int64
	Name      string
	TypeCode  string
	AliasName string
	IsVirtual bool
	Desc      string
	Sort      int32
}
