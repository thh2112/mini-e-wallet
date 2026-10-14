package response

type Pagination[T any] struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	Items      []T   `json:"items"`
}

func Paginate[T any](page, limit, total int64, items []T) *Pagination[T] {
	if items == nil {
		items = make([]T, 0)
	}

	var totalPages int64

	if limit > 0 {
		totalPages = total / limit
		if total%limit != 0 {
			totalPages++
		}
	}

	return &Pagination[T]{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Items:      items,
	}

}
