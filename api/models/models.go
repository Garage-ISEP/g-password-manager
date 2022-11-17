package models

type SecretEntry struct {
	Pk          string `json:"pk"`
	Sk          string `json:"sk" required:"true"`
	Secret      string `json:"secret" required:"true"`
	Url         string `json:"url"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Group       string `json:"group" required:"true"`
	Description string `json:"description"`
}

type GroupEntry struct {
	Pk      string `json:"pk"`
	Sk      string `json:"group"`
	Context string `json:"context"`
}
