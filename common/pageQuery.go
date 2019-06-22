package common

type PageQuery struct {
	PageIndex int
	PageSize  int
	OrderBy   string //aa,asc|bb,desc
	Criteria  map[string]string
}
