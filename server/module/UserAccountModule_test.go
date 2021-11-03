package module

import (
	"context"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/enttest"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

func TestCheckSignInRequestvalidation(t *testing.T) {
	t.Run("正常値テスト", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "nanagami1369outlockexsample",
			RowPassword: "00000000",
			AccountType: user.AccountTypeGeneral,
		}, nil)
	})
	t.Run("UserId値無し", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "",
			RowPassword: "00000000",
			AccountType: user.AccountTypeGeneral,
		}, errors.New("sign in request error id is empty"))
	})
	t.Run("ユーザーIDがファイル", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "file",
			RowPassword: "00000000",
			AccountType: user.AccountTypeGeneral,
		}, errors.New("sign in request error file cannot be used for request id"))
	})
	t.Run("パスワード値無し", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "nanagami1369outlockexsample",
			RowPassword: "",
			AccountType: user.AccountTypeGeneral,
		}, errors.New("sign in request error password is 8 characters or more"))
	})
	t.Run("パスワード7文字", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "nanagami1369outlockexsample",
			RowPassword: "0000000",
			AccountType: user.AccountTypeGeneral,
		}, errors.New("sign in request error password is 8 characters or more"))
	})
	t.Run("AccountType値無し", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "nanagami1369outlockexsample",
			RowPassword: "00000000",
		}, errors.New("sign in request error accountType is empty"))
	})
	t.Run("学生なのに学籍番号無し", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "nanagami1369outlockexsample",
			RowPassword: "00000000",
			AccountType: user.AccountTypeStudent,
		}, errors.New("sign in request error request is student ,but request have not a student Number"))
	})
	t.Run("user idに半角英数以外を割り当てる", func(t *testing.T) {
		testCheckSignInRequestvalidation(t, &model.SignInRequest{
			Id:          "日本語",
			RowPassword: "00000000",
			AccountType: user.AccountTypeTeacher,
		}, errors.New("sign in request error only half-width alphanumeric characters can be used in the user id"))
	})

}

func testCheckSignInRequestvalidation(t *testing.T, in *model.SignInRequest, expected error) {
	t.Helper()
	if err := checkSignInRequestValidation(in); getErrorMessage(err) != getErrorMessage(expected) {
		t.Errorf("チェックに失敗しました \n"+
			"expected:%v\n"+
			"actual  :%v", expected, err)
	}
}

func TestSignIn(t *testing.T) {
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

	uam := NewUserAccountModule(repository.NewUserAccountRepository(context, client))

	t.Run("正常値テスト(学生の場合)", func(t *testing.T) {
		studentNumber := 1121141
		request := &model.SignInRequest{
			Id:            "1121141",
			RowPassword:   "00000000",
			AccountType:   user.AccountTypeStudent,
			StudentNumber: &studentNumber}
		_, err := uam.SignIn(request)
		if err != nil {
			t.Fatalf(err.Error())
		}
		// 登録したデータを取得できるか試す
		_, err = client.User.Query().Where(user.UserID(request.Id)).Only(context)
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
	t.Run("正常値テスト(学生以外の場合)", func(t *testing.T) {
		request := &model.SignInRequest{
			Id:          "tanaka",
			RowPassword: "00000000",
			AccountType: user.AccountTypeGeneral,
		}
		_, err := uam.SignIn(request)
		if err != nil {
			t.Fatalf(err.Error())
		}
		// 登録したデータを取得できるか試す
		_, err = client.User.Query().Where(user.UserID(request.Id)).Only(context)
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
	t.Run("異常値テスト(存在するアカウントを登録しようとする場合)", func(t *testing.T) {
		request := &model.SignInRequest{
			Id:          "1821141",
			RowPassword: "00000000",
			AccountType: user.AccountTypeGeneral,
		}
		expected := "sign in request error request id is Exists"
		_, err := uam.SignIn(request)
		if getErrorMessage(err) != expected {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", expected, err)
		}
	})
	t.Run("異常値テスト(存在する学籍番号を登録しようとする場合)", func(t *testing.T) {
		studentNumber := 1821141
		request := &model.SignInRequest{
			Id:            "1821142",
			RowPassword:   "00000000",
			AccountType:   user.AccountTypeStudent,
			StudentNumber: &studentNumber,
		}
		expected := "sign in request error request student number is Exists"
		_, err := uam.SignIn(request)
		if getErrorMessage(err) != expected {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", expected, err)
		}
	})
	t.Run("異常値テスト(学籍番号があるのに、学生では無い場合)", func(t *testing.T) {
		studentNumber := 1821141
		request := &model.SignInRequest{
			Id:            "1821142",
			RowPassword:   "00000000",
			AccountType:   user.AccountTypeTeacher,
			StudentNumber: &studentNumber,
		}
		expected := "sign in request error request has student number,but request is not student"
		_, err := uam.SignIn(request)
		if getErrorMessage(err) != expected {
			t.Errorf("チェックに失敗しました \n"+
				"expected:%v\n"+
				"actual  :%v", expected, err)
		}
	})
}

func getErrorMessage(err error) string {
	if err == nil {
		return "nil"
	} else {
		return err.Error()
	}
}
