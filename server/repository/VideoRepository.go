package repository

import (
	"context"

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

func (r *VideoRepositoryImpl) Add(user *ent.User, title string, language *model.Language, comment string, codingSequence *[]model.CodingSequence) (*ent.Video, error) {
	return r.client.Video.Create().
		SetUser(user).
		SetTitle(title).
		SetLanguageTag(language).
		SetRecordingTime((*codingSequence)[len(*codingSequence)-1].Timestamp).
		SetComment(comment).
		SetCodingSequence(codingSequence).
		Save(r.context)
}

