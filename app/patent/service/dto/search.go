package dto

import "go-admin/common/dto"

type SimpleSearchReq struct {
	dto.Pagination

	Query string `json:"Query"`
	DB    string `json:"DB"`
}
