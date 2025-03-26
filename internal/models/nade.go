package models

type Nade struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Map         string `json:"map"`
	Vector      string `json:"vector"`
	Usage       string `json:"usage"`
	Description string `json:"description"`
	Link        string `json:"video_link"`
	Position    string `json:"position"`
	Class       string `json:"type"`
}
