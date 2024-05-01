package pagination

import "github.com/cilloparch/cillop/types/list"

type P struct {
	Page  *int `query:"page" validate:"omitempty,gt=0"`
	Limit *int `query:"limit" validate:"omitempty,gt=0"`
}

func (r *P) Default(maxPage ...int) {
	mxPage := 1000
	if len(maxPage) > 0 {
		mxPage = maxPage[0]
	}
	if r.Page == nil || *r.Page <= 0 {
		r.Page = new(int)
		*r.Page = 1
	}
	if r.Limit == nil || *r.Limit <= 0 || *r.Limit > mxPage {
		r.Limit = new(int)
		*r.Limit = 10
	}
}

func (r *P) ToListConfig() list.Config {
	r.Default()
	return list.Config{
		Offset: *r.Limit * (*r.Page - 1),
		Limit:  *r.Limit,
	}
}
