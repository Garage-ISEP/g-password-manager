package models

const (
	PK_SECRET = "pk_secret"
	PK_GROUP  = "pk_group"
)

type SecretEntry struct {
	Pk            string `json:"pk"`
	Sk            string `json:"sk"`
	Name          string `json:"name" required:"true"`
	Secret        string `json:"secret" required:"true"`
	Url           string `json:"url"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Group         string `json:"group" required:"true"`
	Description   string `json:"description"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	CreatedById   string `json:"createdBy"`
	CreatedByName string `json:"createdByName"`
	UpdatedById   string `json:"updatedBy"`
	UpdatedByName string `json:"updatedByName"`
}

type GroupEntry struct {
	Pk      string `json:"pk"`
	Sk      string `json:"group"`
	Context string `json:"context"`
}
