package health

type StatusResponse struct {
	Status string `json:"status"`
}

func (s StatusResponse) IsSuccess() bool {
	return s.Status == "UP"
}
