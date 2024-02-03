package model

type Band struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Description *string `json:"description,omitempty"`
}

type BandInput struct {
	Name        string  `json:"name"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Description *string `json:"description,omitempty"`
}
