package settlements

type GetSettlementsParams struct {
}

type GetAllOrdersParams struct {
	Limit  int `query:"limit"`
	Page   int `query:"page"`
	Offset int
}

func (p *GetAllOrdersParams) Validate() {
	if p.Limit == 0 {
		p.Limit = 5
	}

	if p.Page > 1 {
		p.Offset = (p.Page - 1) * p.Limit
	}
}
