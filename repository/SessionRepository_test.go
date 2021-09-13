package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/enttest"
	"github.com/nanagami1369/CodingShare/ent/user"
)

func TestSet(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	user, _ := client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	dateOfExpiry := time.Now().AddDate(0, 0, 3)
	t.Log(dateOfExpiry)
	repository := NewSessionRepository(context, client)
	t.Run("セッションが登録できるか", func(t *testing.T) {
		session, err := repository.Set(user, dateOfExpiry)
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		t.Logf("set session :%v", session)
	})
	t.Run("現在時刻が有効期限と同じタイミングで失敗するか", func(t *testing.T) {
		_, err := repository.Set(user, time.Now())
		expected := errors.New("set session err date of expiry is before now")
		if fmt.Sprint(err) != fmt.Sprint(expected) {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", expected, err)
		}
	})
	t.Run("有効期限切れのタイミングで失敗するか", func(t *testing.T) {
		beforeNow := time.Now().Add(-100)
		_, err := repository.Set(user, beforeNow)
		expected := errors.New("set session err date of expiry is before now")
		if fmt.Sprint(err) != fmt.Sprint(expected) {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", expected, err)
		}
	})
}

func TestGet(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	user, _ := client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	dateOfExpiry := time.Now().AddDate(0, 0, 3)
	sampleSession, _ := client.Session.Create().
		SetUser(user).
		SetDateOfExpiry(dateOfExpiry).
		Save(context)
	testUuid := sampleSession.ID
	repository := NewSessionRepository(context, client)
	t.Run("Setしたセッションが取得できるか", func(t *testing.T) {
		// Setしたセッションが存在するか判定する
		session, err := repository.Get(testUuid)
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		t.Logf("searched user :%v", session)
	})
}

func TestRemove(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	user, _ := client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	dateOfExpiry := time.Now().AddDate(0, 0, 3)
	sampleSession, _ := client.Session.Create().
		SetUser(user).
		SetDateOfExpiry(dateOfExpiry).
		Save(context)
	testUuid := sampleSession.ID
	repository := NewSessionRepository(context, client)
	t.Run("セッションを削除できるか", func(t *testing.T) {

		// Setしたセッションが存在するか判定する
		err := repository.Remove(testUuid)
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
	})

}

func TestExistFromSession(t *testing.T) {
		// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	// set sample data
	user, _ := client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	dateOfExpiry := time.Now().AddDate(0, 0, 3)
	sampleSession, _ := client.Session.Create().
		SetUser(user).
		SetDateOfExpiry(dateOfExpiry).
		Save(context)
	testUuid := sampleSession.ID
	repository := NewSessionRepository(context, client)
t.Run("存在する場合", func(t *testing.T) {
		result, err := repository.Exists(testUuid)
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
		result, err := repository.Exists(uuid.New())
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
