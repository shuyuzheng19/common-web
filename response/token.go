package response

type TokenResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
	Create string `json:"create"`
}
