package dto

type NoveltyReportReq struct {
	KeyWords []string `json:"keyWords"`
	Title    string   `json:"title"`
	CL       string   `json:"CL"`
}
