package model

type PageInfo struct {
	PageIndex int
	PageSize  int
}

func (p PageInfo) GetOffset() int {
	return (p.PageIndex - 1) * p.PageSize
}
