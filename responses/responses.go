package responses

type DefaultResponse struct {
	Content interface{} `json:"content"`
}

type MetaData struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
}

type DefaultArrayResponse struct {
	Content  interface{} `json:"content"`
	MetaData MetaData    `json:"meta_data"`
}

func NewArrayResponse(i interface{}) DefaultArrayResponse {
	return DefaultArrayResponse{
		Content: i,
		MetaData: MetaData{
			CurrentPage: 1,
			PerPage:     10,
			TotalPages:  1,
			TotalCount:  1,
		},
	}
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
