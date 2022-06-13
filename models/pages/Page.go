package pages

type Page struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int64       `json:"total"`
	List      interface{} `json:"list"`
}

func (p *Page) SetPage(pageIndex int, pageSize int, total int64, list interface{}) *Page {
	p.PageIndex = pageIndex
	p.PageSize = pageSize
	p.Total = total
	p.List = list
	return p
}

func (p Page) GetPage() Page {
	return p
}
