package pagination

import "github.com/cilloparch/cillop/v2/types/list"

// P is the pagination struct
type P struct {

	// Page is the page number
	Page *int `query:"page" validate:"omitempty,gt=0"`

	// Limit is the number of items per page
	Limit *int `query:"limit" validate:"omitempty,gt=0"`
}

// Default sets the default values for the pagination
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

// ToListConfig converts the pagination to a list configuration
func (r *P) ToListConfig() list.Config {
	r.Default()
	return list.Config{
		Offset: *r.Limit * (*r.Page - 1),
		Limit:  *r.Limit,
	}
}
