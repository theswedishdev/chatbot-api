package types

type ChuckNorrisResponse struct {
	Categories []string `json:"categories,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	IconURL    string   `json:"icon_url,omitempty"`
	ID         string   `json:"id,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
	URL        string   `json:"url,omitempty"`
	Value      string   `json:"value,omitempty"`
}
