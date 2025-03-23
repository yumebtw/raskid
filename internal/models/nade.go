package models

type Nade struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Map         string `json:"map"`
	Vector      string `json:"vector"`
	Type        string `json:"type"`
	Usage       string `json:"usage"`
	Position    string `json:"position"`
	VideoLink   string `json:"video_link"`
	Description string `json:"description"`
}
