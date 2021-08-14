package repository

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/enttest"
	"github.com/nanagami1369/CodingShare/ent/user"
)

func TestFindOne(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	client.User.Create().
		SetUserID("tanaka").
		SetPassword("sample").
		SetAccountType(user.AccountTypeTeacher).
		Save(context)
	client.User.Create().
		SetUserID("suzuki").
		SetPassword("sample").
		SetAccountType(user.AccountTypeGeneral).
		Save(context)

	repository := NewUserAccountRepository(context, client)
	// 登録したアカウントが存在するか判定する
	user, err := repository.FindOne("suzuki")
	if err != nil {
		t.Fatalf("err:%#v", err)
	}
	t.Logf("searched user :%v", user)
}

func TestExist(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	client.User.Create().
		SetUserID("tanaka").
		SetPassword("sample").
		SetAccountType(user.AccountTypeTeacher).
		Save(context)
	client.User.Create().
		SetUserID("suzuki").
		SetPassword("sample").
		SetAccountType(user.AccountTypeGeneral).
		Save(context)

	repository := NewUserAccountRepository(context, client)
	t.Run("存在する場合", func(t *testing.T) {
		result, err := repository.Exists("suzuki")
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		if result != true {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", true, false)
		}
	})
	t.Run("存在しない場合", func(t *testing.T) {
		result, err := repository.Exists("okamoto")
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		if result != false {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", false, true)
		}
	})

}
