package api

type PaginationQueryItems struct {
	Before string
	After  string
	Limit  int
}

type PaginationOption func(*PaginationQueryItems)

func Before(beforeId string) PaginationOption {
	return func(pqi *PaginationQueryItems) {
		pqi.Before = beforeId
	}
}

func After(afterId string) PaginationOption {
	return func(pqi *PaginationQueryItems) {
		pqi.After = afterId
	}
}

func Limit(pageSize int) PaginationOption {
	return func(pqi *PaginationQueryItems) {
		pqi.Limit = pageSize
	}
}
