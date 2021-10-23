package module

import (
	"errors"

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

func (m *VideoModuleImpl) Save(request *model.SaveVideoRequest, user *ent.User) (*model.Video, error) {
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
	return m.repository.FindOne(id)
}

func (m *VideoModuleImpl) Search(keyWord string) ([]*model.Video, error) {
	return m.repository.FindFromTitle(keyWord)
}

func (m *VideoModuleImpl) GetUserVideos(id string) ([]*model.Video, error) {
	return m.repository.FindFromUserId(id)
}

func (m *VideoModuleImpl) Remove(id int, loginUser *ent.User) error {
	isExist, err := m.repository.Exists(id)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("not is video")
	}
	video, err := m.repository.FindOne(id)
	if err != nil {
		return err
	}
	if loginUser.UserID != video.Header.UserID {
		return errors.New("this video is not for the loggin user")
	}
	return m.repository.Remove(id)
}
