package model

import "github.com/nanagami1369/CodingShare/ent/user"

type SignInRequest struct {
	Id            string           `json:"id"`
	RowPassword   string           `json:"password"`
	AccountType   user.AccountType // プログラム側で設定する事
	StudentNumber *int             // 生徒のアカウントは自分で作成させないのでプログラム側で設定する事
}
