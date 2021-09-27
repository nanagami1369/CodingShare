package model

type Header struct {
	UserID        string    `json:"userId"` // DB側から設定
	Name          string    `json:"name"`   // DB側から設定
	Title         string    `json:"title"`
	Language      *Language `json:"language"`
	UploadTime    int64     `json:"uploadTime"` // DB側で設定
	RecordingTime int       `json:"recordingTime"`
	Comment       string    `json:"comment"`
}
