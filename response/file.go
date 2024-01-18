package response

type SimpleFileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Create  string `json:"create"`
	Url     string `json:"url"`
}
