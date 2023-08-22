package list

import "time"

type Result[Type any] struct {
	List          []Type
	Total         int64
	FilteredTotal int64
	Page          int64
	IsNext        bool
	IsPrev        bool
}

type DateResult[Type any] struct {
	Result[Type]
	StartDate time.Time
	EndDate   time.Time
}