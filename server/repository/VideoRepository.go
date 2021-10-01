package repository

import (
	"context"
	"time"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/video"
	"github.com/nanagami1369/CodingShare/model"
)

func NewVideoRepository(context context.Context, client *ent.Client) VideoRepository {
	return &VideoRepositoryImpl{
		context: context,
		client:  client,
	}
}

type VideoRepositoryImpl struct {
	context context.Context
	client  *ent.Client
}

func (r *VideoRepositoryImpl) Add(user *ent.User, title string, language *model.Language, comment string, codingSequence *[]model.CodingSequence) (*model.Video, error) {
	video, err := r.client.Video.Create().
		SetUser(user).
		SetTitle(title).
		SetLanguageTag(language).
		SetRecordingTime((*codingSequence)[len(*codingSequence)-1].Timestamp).
		SetComment(comment).
		SetCodingSequence(codingSequence).
		Save(r.context)
	if err != nil {
		return nil, err
	}
	return &model.Video{
		Header: model.Header{
			UserID:        user.UserID,
			Name:          user.UserID,
			Title:         video.Title,
			Language:      video.LanguageTag,
			UploadTime:    time.Now().UTC().UnixNano() / 1e6,
			RecordingTime: video.RecordingTime,
			Comment:       video.Comment,
		},
		Value: video.CodingSequence,
	}, err
}

func (r *VideoRepositoryImpl) FindOne(id int) (*model.Video, error) {
	video, err := r.client.Video.Query().Where(video.ID(id)).Only(r.context)
	if err != nil {
		return nil, err
	}
	user, err := video.QueryUser().First(r.context)
	if err != nil {
		return nil, err
	}
	return &model.Video{
		Header: model.Header{
			UserID:        user.UserID,
			Name:          user.UserID,
			Title:         video.Title,
			Language:      video.LanguageTag,
			UploadTime:    time.Now().UTC().UnixNano() / 1e6,
			RecordingTime: video.RecordingTime,
			Comment:       video.Comment,
		},
		Value: video.CodingSequence,
	}, err
}

func (r *VideoRepositoryImpl) FindFromTitle(title string) ([]*model.Video, error) {
	searchedVideos, err := r.client.Video.Query().
		Where(video.TitleContains(title)).
		All(r.context)
	if err != nil {
		return nil, err
	}
	videos := make([]*model.Video, len(searchedVideos))
	for _, video := range searchedVideos {
		user, err := video.QueryUser().First(r.context)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &model.Video{
			Header: model.Header{
				UserID:        user.UserID,
				Name:          user.UserID,
				Title:         video.Title,
				Language:      video.LanguageTag,
				UploadTime:    time.Now().UTC().UnixNano() / 1e6,
				RecordingTime: video.RecordingTime,
				Comment:       video.Comment,
			},
			Value: video.CodingSequence})
	}
	return videos, nil
}
func (r *VideoRepositoryImpl) Exists(id int) (bool, error) {
	return r.client.Video.Query().Where(video.ID(id)).Exist(r.context)
}
