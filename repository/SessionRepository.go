package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/session"
)

type SessionRepositoryImpl struct {
	context context.Context
	client  *ent.Client
}

func NewSessionRepository(context context.Context, client *ent.Client) SessionRepository {
	return &SessionRepositoryImpl{context: context, client: client}
}

func (r *SessionRepositoryImpl) Set(user *ent.User, dateOfExpiry time.Time) (*ent.Session, error) {
	now := time.Now()
	if dateOfExpiry.Before(now) {
		return nil, errors.New("set session err date of expiry is before now")
	}
	return r.client.Session.Create().
		SetUser(user).
		SetDateOfExpiry(dateOfExpiry).
		Save(r.context)
}

func (r *SessionRepositoryImpl) Get(uuid uuid.UUID) (*ent.Session, *ent.User, error) {
	session, err := r.client.Session.Query().Where(session.ID(uuid)).Only(r.context)
	if err != nil {
		return nil, nil, err
	}
	user, err := session.QueryUser().Only(r.context)
	if err != nil {
		return nil, nil, err
	}
	return session, user, err
}

func (r *SessionRepositoryImpl) Remove(uuid uuid.UUID) (*ent.User, error) {
	user, err := r.client.Session.Query().Where(session.ID(uuid)).QueryUser().Only(r.context)
	if err != nil {
		return nil, err
	}
	err = r.client.Session.DeleteOneID(uuid).Exec(r.context)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *SessionRepositoryImpl) Exists(uuid uuid.UUID) (bool, error) {
	return r.client.Session.Query().Where(session.ID(uuid)).Exist(r.context)
}
