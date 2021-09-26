package model

type SaveVideoRequest struct {
	Header struct {
		Title    string   `json:"title"`
		Language Language `json:"language"`
		Comment  string   `json:"comment"`
	} `json:"header"`
	Value []CodingSequence
}
