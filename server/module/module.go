package module

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/model"
)

type UserAccountModule interface {
	Confirm(request *model.LoginRequest) (*ent.User, error)
	SignIn(request *model.SignInRequest) (*ent.User, error)
}

type SessionModule interface {
	Login(store sessions.Session, user *ent.User, dateOfExpiry time.Time) (string, error)
	IsLogin(store sessions.Session) (bool, error)
	Logout(store sessions.Session) (*ent.User, error)
	Get(store sessions.Session) (*ent.Session, *ent.User, error)
}

type VideoModule interface {
	Save(request *model.SaveVideoRequest, user *ent.User) (*model.Video, error)
	Load(id int) (*model.Video, error)
	Search(keyWord string) ([]*model.Video, error)
	GetUserVideos(id string) ([]*model.Video, error)
}
