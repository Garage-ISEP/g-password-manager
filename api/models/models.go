package models

const (
	PK_SECRET = "pk_secret"
	PK_GROUP  = "pk_group"
)

type Secret struct {
	Pk            string   `json:"pk"`
	Sk            string   `json:"sk"`
	Name          string   `json:"name" required:"true"`
	Url           string   `json:"url"`
	Email         string   `json:"email"`
	Username      string   `json:"username"`
	WritingGroups []string `json:"writingGroups" required:"true"`
	ReadingGroups []string `json:"readingGroups" required:"true"`
	Description   string   `json:"description"`

	Secret string `json:"secret" required:"true"`

	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	CreatedById   string `json:"createdBy"`
	CreatedByName string `json:"createdByName"`
	UpdatedById   string `json:"updatedBy"`
	UpdatedByName string `json:"updatedByName"`
}

// (Group Id, Group Name)
type Group struct {
	Pk   string `json:"pk"`
	Sk   string `json:"id"`
	Name string `json:"name"`
}
