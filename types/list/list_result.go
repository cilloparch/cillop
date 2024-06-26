package list

import "time"

// Result is a struct that defines the result of a list
type Result[Type any] struct {

	// List is a slice of items
	List []Type `json:"list"`

	// Total is the total number of items
	Total int64 `json:"total"`

	// FilteredTotal is the total number of items after filtering
	FilteredTotal int64 `json:"filteredTotal"`

	// Page is the current page
	Page int64 `json:"page"`

	// IsNext is a boolean that indicates if there is a next page
	IsNext bool `json:"isNext"`

	// IsPrev is a boolean that indicates if there is a previous page
	IsPrev bool `json:"isPrev"`
}

// DateResult is a struct that defines the result of a list with dates
type DateResult[Type any] struct {
	Result[Type]

	// StartDate is the start date of the list
	StartDate time.Time

	// EndDate is the end date of the list
	EndDate time.Time
}

func NewListResult[Type any](items []Type, total int64, filteredTotal int64, listConfig Config) *Result[Type] {
	return &Result[Type]{
		List:          items,
		Total:         total,
		FilteredTotal: filteredTotal,
		Page:          int64(listConfig.Offset/listConfig.Limit + 1),
		IsNext:        total > int64(listConfig.Offset+listConfig.Limit),
		IsPrev:        listConfig.Offset > 0,
	}
}

func NewDateListResult[Type any](items []Type, total int64, filteredTotal int64, listConfig Config, startDate time.Time, endDate time.Time) *DateResult[Type] {
	return &DateResult[Type]{
		Result:    *NewListResult[Type](items, total, filteredTotal, listConfig),
		StartDate: startDate,
		EndDate:   endDate,
	}
}
