package list

import "time"

type Result[Type any] struct {
	List          []Type `json:"list"`
	Total         int64  `json:"total"`
	FilteredTotal int64  `json:"filteredTotal"`
	Page          int64  `json:"page"`
	IsNext        bool   `json:"isNext"`
	IsPrev        bool   `json:"isPrev"`
}

type DateResult[Type any] struct {
	Result[Type]
	StartDate time.Time
	EndDate   time.Time
}
