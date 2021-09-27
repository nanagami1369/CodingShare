package module

import (
	"errors"
	"time"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

type VideoModuleImpl struct {
	repository repository.VideoRepository
}

func NewVideoModule(r repository.VideoRepository) VideoModule {
	return &VideoModuleImpl{
		repository: r,
	}
}

func (m *VideoModuleImpl) Save(request *model.SaveVideoRequest, user *ent.User) (*ent.Video, error) {
	if request.Header.Title == "" {
		return nil, errors.New("video save request err title property is not set")
	}
	if request.Header.Language == nil {
		return nil, errors.New("video save request err lang property is not set")
	}
	if request.Value == nil || len(*request.Value) == 0 {
		return nil, errors.New("video save request err coding sequence is not set")
	}
	return m.repository.Add(user, request.Header.Title, request.Header.Language, request.Header.Comment, request.Value)
}

func (m *VideoModuleImpl) Load(id int) (*model.Video, error) {
	isExists, err := m.repository.Exists(id)
	if err != nil {
		return nil, err
	}
	if !isExists {
		return nil, nil
	}
	dbVideo, user, err := m.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return &model.Video{
		Header: model.Header{
			UserID:        user.UserID,
			Name:          user.UserID,
			Title:         dbVideo.Title,
			Language:      dbVideo.LanguageTag,
			UploadTime:    time.Now().UTC().UnixNano() / 1e6,
			RecordingTime: dbVideo.RecordingTime,
			Comment:       dbVideo.Comment,
		},
		Value: dbVideo.CodingSequence,
	}, nil
}
