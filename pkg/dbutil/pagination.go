package dbutil

type Pagination struct {
	Limit  int
	Offset int
}

func NewPagination(limit, offset int) Pagination {
	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

type PaginatedResult[T interface{}] struct {
	Data  []T
	Total int
}
