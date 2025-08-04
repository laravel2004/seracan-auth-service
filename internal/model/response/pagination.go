package response

type PaginationResult struct {
	TotalData   int64       `json:"total_data"`
	TotalPage   int         `json:"total_page"`
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	Data        interface{} `json:"data"`
}
