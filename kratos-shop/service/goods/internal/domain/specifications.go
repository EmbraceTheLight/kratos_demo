package domain

type SpecificationValue struct {
	ID     int64
	AttrId int64
	Value  string
	Sort   int32
}
type Specification struct {
	ID                  int64
	TypeID              int64
	Name                string
	Sort                int32
	Status              bool
	IsSKU               bool
	IsSelect            bool
	SpecificationValues []*SpecificationValue
}

func (b *Specification) IsTypeIDEmpty() bool {
	return b.TypeID == 0
}

func (b *Specification) IsValueEmpty() bool {
	return b.SpecificationValues == nil
}
