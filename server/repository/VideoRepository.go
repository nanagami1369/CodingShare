package repository

import (
	"context"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
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
			VideoID:       video.ID,
			UserID:        user.UserID,
			Name:          user.UserID,
			Title:         video.Title,
			Language:      video.LanguageTag,
			UploadTime:    video.UploadTime.UTC().UnixNano() / 1e6,
			RecordingTime: video.RecordingTime,
			Comment:       video.Comment,
		},
		Value: video.CodingSequence,
	}, err
}

func (r *VideoRepositoryImpl) FindOne(id int) (*model.Video, error) {
	video, err := r.client.Video.Query().Where(video.And(video.ID(id), video.IsRemoved(false))).Only(r.context)
	if err != nil {
		return nil, err
	}
	user, err := video.QueryUser().First(r.context)
	if err != nil {
		return nil, err
	}
	return &model.Video{
		Header: model.Header{
			VideoID:       video.ID,
			UserID:        user.UserID,
			Name:          user.UserID,
			Title:         video.Title,
			Language:      video.LanguageTag,
			UploadTime:    video.UploadTime.UTC().UnixNano() / 1e6,
			RecordingTime: video.RecordingTime,
			Comment:       video.Comment,
		},
		Value: video.CodingSequence,
	}, err
}

func (r *VideoRepositoryImpl) FindFromTitle(title string) ([]*model.Video, error) {
	searchedVideos, err := r.client.Video.Query().
		Where(
			video.And(
				video.TitleContains(title),
				video.IsRemoved(false),
			),
		).
		All(r.context)
	if err != nil {
		return nil, err
	}
	videos := make([]*model.Video, 0)
	for _, video := range searchedVideos {
		user, err := video.QueryUser().First(r.context)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &model.Video{
			Header: model.Header{
				VideoID:       video.ID,
				UserID:        user.UserID,
				Name:          user.UserID,
				Title:         video.Title,
				Language:      video.LanguageTag,
				UploadTime:    video.UploadTime.UTC().UnixNano() / 1e6,
				RecordingTime: video.RecordingTime,
				Comment:       video.Comment,
			},
			Value: video.CodingSequence})
	}
	return videos, nil
}

func (r *VideoRepositoryImpl) FindFromUserId(id string) ([]*model.Video, error) {
	exists, err := r.client.User.Query().Where(user.UserID(id)).Exist(r.context)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	user, err := r.client.User.Query().Where(user.UserID(id)).First(r.context)
	if err != nil {
		return nil, err
	}
	searchedVideos, err := r.client.User.QueryVideos(user).Where(video.IsRemoved(false)).All(r.context)
	if err != nil {
		return nil, err
	}
	videos := make([]*model.Video, 0)
	for _, video := range searchedVideos {
		if err != nil {
			return nil, err
		}
		videos = append(videos, &model.Video{
			Header: model.Header{
				VideoID:       video.ID,
				UserID:        user.UserID,
				Name:          user.UserID,
				Title:         video.Title,
				Language:      video.LanguageTag,
				UploadTime:    video.UploadTime.UTC().UnixNano() / 1e6,
				RecordingTime: video.RecordingTime,
				Comment:       video.Comment,
			},
			Value: video.CodingSequence})
	}
	return videos, nil

}

func (r *VideoRepositoryImpl) Exists(id int) (bool, error) {
	return r.client.Video.Query().Where(video.And(video.ID(id), video.IsRemoved(false))).Exist(r.context)
}
