package model

type Header struct {
	VideoID       int       `json:"videoId"` // DB側から設定
	UserID        string    `json:"userId"`  // DB側から設定
	Name          string    `json:"name"`    // DB側から設定
	Title         string    `json:"title"`
	Language      *Language `json:"language"`
	UploadTime    int64     `json:"uploadTime"` // DB側で設定
	RecordingTime int       `json:"recordingTime"`
	Comment       string    `json:"comment"`
}
