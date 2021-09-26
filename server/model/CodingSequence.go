package model

type CodingSequence struct {
	Timestamp  int `json:"timestamp"`
	ChangeData struct {
		From struct {
			Line int `json:"line"`
			Ch   int `json:"ch"`
		} `json:"from"`
		To struct {
			Line int `json:"line"`
			Ch   int `json:"ch"`
		} `json:"to"`
		Text    []string `json:"text"`
		Removed []string `json:"removed"`
		Origin  string   `json:"origin"`
	} `json:"changeData,omitempty"`
	Cursor struct {
		Line   int         `json:"line"`
		Ch     int         `json:"ch"`
		Sticky interface{} `json:"sticky"`
	} `json:"cursor,omitempty"`
}
