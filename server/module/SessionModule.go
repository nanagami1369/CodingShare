package module

import (
	"errors"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/repository"
)

type SessionModuleImpl struct {
	repository repository.SessionRepository
	SessionKey string
}

func NewSessionModule(r repository.SessionRepository) SessionModule {
	return &SessionModuleImpl{
		repository: r,
		SessionKey: "session_id",
	}
}

func (m *SessionModuleImpl) Login(store sessions.Session, user *ent.User, dateOfExpiry time.Time) (string, error) {
	// セッションを作成しサーバーに保存
	newSession, err := m.repository.Set(user, dateOfExpiry)
	if err != nil {
		return "", err
	}
	token := newSession.ID.String()

	// セッションをクッキーに保存
	store.Set("session_id", token)
	store.Save()

	return token, nil
}

func (m *SessionModuleImpl) Get(store sessions.Session) (*ent.Session, *ent.User, error) {
	uuid, err := m.getSessionIdFromStore(store)
	if err != nil {
		return nil, nil, err
	}
	return m.repository.Get(uuid)
}

func (m *SessionModuleImpl) IsLogin(store sessions.Session) (bool, error) {
	uuid, err := m.getSessionIdFromStore(store)
	if err != nil {
		return false, err
	}
	return m.repository.Exists(uuid)
}

func (m *SessionModuleImpl) Logout(store sessions.Session) (*ent.User, error) {
	uuid, err := m.getSessionIdFromStore(store)
	if err != nil {
		return nil, err
	}
	// ログアウト処理はログイン済みのユーザーしかアクセスできない仕様なので存在確認はしない
	user, err := m.repository.Remove(uuid)
	if err != nil {
		return nil, err
	}
	store.Delete("session_id")
	return user, nil
}

func (m *SessionModuleImpl) getSessionIdFromStore(store sessions.Session) (uuid.UUID, error) {
	sessionValue := store.Get("session_id")
	var token string
	if session_id, ok := sessionValue.(string); ok {
		token = session_id
	} else {
		return uuid.Nil, errors.New("session can not read")
	}
	return uuid.Parse(token)
}