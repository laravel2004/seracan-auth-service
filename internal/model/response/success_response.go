package response

import "github.com/morkid/paginate"

type SuccessResponse struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    *PaginationMap `json:"meta,omitempty"`
}

type PaginationMap struct {
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	MaxPage    int64 `json:"max_page"`
	TotalPages int64 `json:"total_pages"`
	Total      int64 `json:"total"`
	Last       bool  `json:"last"`
	First      bool  `json:"first"`
	Visible    int64 `json:"visible"`
}

func NewSuccessPaginationResponse(message string, data paginate.Page) *SuccessResponse {
	meta := &PaginationMap{
		Page:       data.Page,
		Size:       data.Size,
		MaxPage:    data.MaxPage,
		TotalPages: data.TotalPages,
		Total:      data.Total,
		Last:       data.Last,
		First:      data.First,
		Visible:    data.Visible,
	}
	return &SuccessResponse{
		Message: message,
		Data:    data.Items,
		Meta:    meta,
	}
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}
}
