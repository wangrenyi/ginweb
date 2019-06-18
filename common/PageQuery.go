package common

type PageQuery struct {
	PageIndex uint
	PageSize  uint
	OrderBy   string //aa,asc|bb,desc
	Criteria  map[string]string
}
