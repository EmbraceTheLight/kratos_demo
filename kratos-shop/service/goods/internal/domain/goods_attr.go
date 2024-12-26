package domain

type AttrGroup struct {
	ID     int64
	TypeID int64
	Title  string
	Desc   string
	Status bool
	Sort   int32
}

func (p *AttrGroup) IsTypeIDEmpty() bool {
	return p.TypeID == 0
}

type GoodsAttr struct {
	ID             int64
	TypeID         int64
	GroupID        int64
	Title          string
	Sort           int32
	Status         bool
	Desc           string
	GoodsAttrValue []*GoodsAttrValue
}

func (p *GoodsAttr) IsTypeIDEmpty() bool {
	return p.TypeID == 0
}

type GoodsAttrValue struct {
	ID      int64
	AttrID  int64
	GroupID int64
	Value   string
}

func (p *GoodsAttrValue) IsValueEmpty() bool {
	return p.Value == ""
}
