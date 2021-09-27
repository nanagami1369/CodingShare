package model

type Video struct {
	Header Header            `json:"header"`
	Value  *[]CodingSequence `json:"value"`
}
