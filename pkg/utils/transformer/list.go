package transformer

type Pagination struct {
	Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Pagination PaginationMeta `json:"pagination"`
	MetaData   interface{}    `json:"meta_data,omitempty"`
}

type PaginationMeta struct {
	Limit      int    `json:"limit,omitempty"`
	Page       int    `json:"page,omitempty"`
	Sort       string `json:"sort,omitempty"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}
