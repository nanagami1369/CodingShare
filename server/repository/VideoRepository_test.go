package repository

import (
	"context"
	"encoding/json"

	"testing"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/enttest"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/model"
)

func TestAdd(t *testing.T) {
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
	repository := NewVideoRepository(context, client)
	t.Run("ビデオが登録できるか", func(t *testing.T) {
		request := &model.SaveVideoRequest{}
		err := json.Unmarshal([]byte(jsonString), request)
		if err != nil {
			t.Error(err)
			return
		}
		title := request.Header.Title
		lang := request.Header.Language
		comment := request.Header.Comment
		video, err := repository.Add(user, title, lang, comment, request.Value)
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		user, err := video.QueryUser().First(context)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("set video header:",
			user.UserID,
			video.Title,
			video.LanguageTag.Name,
			video.UploadTime,
			video.RecordingTime,
			video.Comment)
	})
}

func TestFindOneFromVideo(t *testing.T) {
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
	request := &model.SaveVideoRequest{}
	err := json.Unmarshal([]byte(jsonString), request)
	if err != nil {
		t.Error(err)
		return
	}
	video, _ := client.Video.Create().
		SetUser(user).
		SetTitle(request.Header.Title).
		SetLanguageTag(request.Header.Language).
		SetRecordingTime((*request.Value)[len(*request.Value)-1].Timestamp).
		SetComment(request.Header.Comment).
		SetCodingSequence(request.Value).
		Save(context)
	repository := NewVideoRepository(context, client)
	// 登録したアカウントが存在するか判定する
	t.Run("ビデオが取得できるか", func(t *testing.T) {
		searchId := video.ID
		video, user, err := repository.FindOne(searchId)
		if err != nil {
			t.Fatalf("err:%#v", err)
		}
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("searched video :",
			user.UserID,
			video.Title,
			video.LanguageTag.Name,
			video.UploadTime,
			video.RecordingTime,
			video.Comment)
	})
}

func TestExistFromVideo(t *testing.T) {
	// openDB
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	context := context.Background()
	user, _ := client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	request := &model.SaveVideoRequest{}
	err := json.Unmarshal([]byte(jsonString), request)
	if err != nil {
		t.Error(err)
		return
	}
	video, _ := client.Video.Create().
		SetUser(user).
		SetTitle(request.Header.Title).
		SetLanguageTag(request.Header.Language).
		SetRecordingTime((*request.Value)[len(*request.Value)-1].Timestamp).
		SetComment(request.Header.Comment).
		SetCodingSequence(request.Value).
		Save(context)
	testId := video.ID
	repository := NewVideoRepository(context, client)
	t.Run("存在する場合", func(t *testing.T) {
		result, err := repository.Exists(testId)
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
		result, err := repository.Exists(1000000)
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

const jsonString = `{
    "header": {
        "title": "test",
        "language": { "tag": "text/x-csrc", "name": "C言語" },
        "comment": "一般的なC言語のハローワールドです"
    },
    "value": [
        {
            "timestamp": 0,
            "changeData": {
                "from": { "line": 0, "ch": 0 },
                "to": { "line": 0, "ch": 0 },
                "text": [""],
                "removed": [""],
                "origin": "input"
            },
            "cursor": { "line": 0, "ch": 0, "sticky": null }
        },
        {
            "timestamp": 1053,
            "changeData": {
                "from": { "line": 0, "ch": 0, "sticky": null },
                "to": { "line": 0, "ch": 0, "sticky": null },
                "text": ["#"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 1, "sticky": null }
        },
        {
            "timestamp": 1224,
            "changeData": {
                "from": { "line": 0, "ch": 1, "sticky": null },
                "to": { "line": 0, "ch": 1, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 2, "sticky": null }
        },
        {
            "timestamp": 1368,
            "changeData": {
                "from": { "line": 0, "ch": 2, "sticky": null },
                "to": { "line": 0, "ch": 2, "sticky": null },
                "text": ["n"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 3, "sticky": null }
        },
        {
            "timestamp": 1508,
            "changeData": {
                "from": { "line": 0, "ch": 3, "sticky": null },
                "to": { "line": 0, "ch": 3, "sticky": null },
                "text": ["c"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 1598,
            "changeData": {
                "from": { "line": 0, "ch": 4, "sticky": null },
                "to": { "line": 0, "ch": 4, "sticky": null },
                "text": ["l"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 5, "sticky": null }
        },
        {
            "timestamp": 1778,
            "changeData": {
                "from": { "line": 0, "ch": 5, "sticky": null },
                "to": { "line": 0, "ch": 5, "sticky": null },
                "text": ["u"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 6, "sticky": null }
        },
        {
            "timestamp": 1860,
            "changeData": {
                "from": { "line": 0, "ch": 6, "sticky": null },
                "to": { "line": 0, "ch": 6, "sticky": null },
                "text": ["d"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 7, "sticky": null }
        },
        {
            "timestamp": 2022,
            "changeData": {
                "from": { "line": 0, "ch": 7, "sticky": null },
                "to": { "line": 0, "ch": 7, "sticky": null },
                "text": ["e"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 8, "sticky": null }
        },
        {
            "timestamp": 2142,
            "changeData": {
                "from": { "line": 0, "ch": 8, "sticky": null },
                "to": { "line": 0, "ch": 8, "sticky": null },
                "text": [" "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 9, "sticky": null }
        },
        {
            "timestamp": 2967,
            "changeData": {
                "from": { "line": 0, "ch": 9, "sticky": null },
                "to": { "line": 0, "ch": 9, "sticky": null },
                "text": ["<"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 3099,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": null },
                "to": { "line": 0, "ch": 10, "sticky": null },
                "text": ["."],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 4429,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": 1 },
                "to": { "line": 0, "ch": 11, "sticky": "before" },
                "text": [""],
                "removed": ["."],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 4903,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": null },
                "to": { "line": 0, "ch": 10, "sticky": null },
                "text": [">"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 5921,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": "after" },
                "to": { "line": 0, "ch": 10, "sticky": "after" },
                "text": ["H"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 6035,
            "changeData": {
                "from": { "line": 0, "ch": 11, "sticky": null },
                "to": { "line": 0, "ch": 11, "sticky": null },
                "text": ["e"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 12, "sticky": null }
        },
        {
            "timestamp": 6207,
            "changeData": {
                "from": { "line": 0, "ch": 12, "sticky": null },
                "to": { "line": 0, "ch": 12, "sticky": null },
                "text": ["l"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 6380,
            "changeData": {
                "from": { "line": 0, "ch": 13, "sticky": null },
                "to": { "line": 0, "ch": 13, "sticky": null },
                "text": ["l"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 6548,
            "changeData": {
                "from": { "line": 0, "ch": 14, "sticky": null },
                "to": { "line": 0, "ch": 14, "sticky": null },
                "text": ["o"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 15, "sticky": null }
        },
        {
            "timestamp": 7862,
            "changeData": {
                "from": { "line": 0, "ch": 14, "sticky": 1 },
                "to": { "line": 0, "ch": 15, "sticky": null },
                "text": [""],
                "removed": ["o"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 8005,
            "changeData": {
                "from": { "line": 0, "ch": 13, "sticky": 1 },
                "to": { "line": 0, "ch": 14, "sticky": null },
                "text": [""],
                "removed": ["l"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 8158,
            "changeData": {
                "from": { "line": 0, "ch": 12, "sticky": 1 },
                "to": { "line": 0, "ch": 13, "sticky": null },
                "text": [""],
                "removed": ["l"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 12, "sticky": null }
        },
        {
            "timestamp": 8305,
            "changeData": {
                "from": { "line": 0, "ch": 11, "sticky": 1 },
                "to": { "line": 0, "ch": 12, "sticky": null },
                "text": [""],
                "removed": ["e"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 8464,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": 1 },
                "to": { "line": 0, "ch": 11, "sticky": null },
                "text": [""],
                "removed": ["H"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 8577,
            "changeData": {
                "from": { "line": 0, "ch": 10, "sticky": null },
                "to": { "line": 0, "ch": 10, "sticky": null },
                "text": ["s"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 8721,
            "changeData": {
                "from": { "line": 0, "ch": 11, "sticky": null },
                "to": { "line": 0, "ch": 11, "sticky": null },
                "text": ["t"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 12, "sticky": null }
        },
        {
            "timestamp": 8937,
            "changeData": {
                "from": { "line": 0, "ch": 12, "sticky": null },
                "to": { "line": 0, "ch": 12, "sticky": null },
                "text": ["d"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 9152,
            "changeData": {
                "from": { "line": 0, "ch": 13, "sticky": null },
                "to": { "line": 0, "ch": 13, "sticky": null },
                "text": ["u"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 9468,
            "changeData": {
                "from": { "line": 0, "ch": 14, "sticky": null },
                "to": { "line": 0, "ch": 14, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 15, "sticky": null }
        },
        {
            "timestamp": 9838,
            "changeData": {
                "from": { "line": 0, "ch": 14, "sticky": 1 },
                "to": { "line": 0, "ch": 15, "sticky": null },
                "text": [""],
                "removed": ["i"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 10000,
            "changeData": {
                "from": { "line": 0, "ch": 13, "sticky": 1 },
                "to": { "line": 0, "ch": 14, "sticky": null },
                "text": [""],
                "removed": ["u"],
                "origin": "+delete"
            },
            "cursor": { "line": 0, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 10318,
            "changeData": {
                "from": { "line": 0, "ch": 13, "sticky": null },
                "to": { "line": 0, "ch": 13, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 10482,
            "changeData": {
                "from": { "line": 0, "ch": 14, "sticky": null },
                "to": { "line": 0, "ch": 14, "sticky": null },
                "text": ["o"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 15, "sticky": null }
        },
        {
            "timestamp": 10735,
            "changeData": {
                "from": { "line": 0, "ch": 15, "sticky": null },
                "to": { "line": 0, "ch": 15, "sticky": null },
                "text": ["."],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 16, "sticky": null }
        },
        {
            "timestamp": 10984,
            "changeData": {
                "from": { "line": 0, "ch": 16, "sticky": null },
                "to": { "line": 0, "ch": 16, "sticky": null },
                "text": ["h"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 0, "ch": 17, "sticky": null }
        },
        {
            "timestamp": 11539,
            "changeData": {
                "from": { "line": 0, "ch": 18, "sticky": "before" },
                "to": { "line": 0, "ch": 18, "sticky": "before" },
                "text": ["", ""],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 1, "ch": 0, "sticky": null }
        },
        {
            "timestamp": 11708,
            "changeData": {
                "from": { "line": 1, "ch": 0, "sticky": null },
                "to": { "line": 1, "ch": 0, "sticky": null },
                "text": ["", ""],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 0, "sticky": null }
        },
        {
            "timestamp": 12225,
            "changeData": {
                "from": { "line": 2, "ch": 0, "sticky": null },
                "to": { "line": 2, "ch": 0, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 1, "sticky": null }
        },
        {
            "timestamp": 12360,
            "changeData": {
                "from": { "line": 2, "ch": 1, "sticky": null },
                "to": { "line": 2, "ch": 1, "sticky": null },
                "text": ["n"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 2, "sticky": null }
        },
        {
            "timestamp": 12449,
            "changeData": {
                "from": { "line": 2, "ch": 2, "sticky": null },
                "to": { "line": 2, "ch": 2, "sticky": null },
                "text": ["t"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 3, "sticky": null }
        },
        {
            "timestamp": 12554,
            "changeData": {
                "from": { "line": 2, "ch": 3, "sticky": null },
                "to": { "line": 2, "ch": 3, "sticky": null },
                "text": [" "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 12723,
            "changeData": {
                "from": { "line": 2, "ch": 4, "sticky": null },
                "to": { "line": 2, "ch": 4, "sticky": null },
                "text": ["m"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 5, "sticky": null }
        },
        {
            "timestamp": 12957,
            "changeData": {
                "from": { "line": 2, "ch": 5, "sticky": null },
                "to": { "line": 2, "ch": 5, "sticky": null },
                "text": ["a"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 6, "sticky": null }
        },
        {
            "timestamp": 13099,
            "changeData": {
                "from": { "line": 2, "ch": 6, "sticky": null },
                "to": { "line": 2, "ch": 6, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 7, "sticky": null }
        },
        {
            "timestamp": 13258,
            "changeData": {
                "from": { "line": 2, "ch": 7, "sticky": null },
                "to": { "line": 2, "ch": 7, "sticky": null },
                "text": ["n"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 8, "sticky": null }
        },
        {
            "timestamp": 13659,
            "changeData": {
                "from": { "line": 2, "ch": 8, "sticky": null },
                "to": { "line": 2, "ch": 8, "sticky": null },
                "text": ["()"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 9 }
        },
        {
            "timestamp": 14084,
            "changeData": {
                "from": { "line": 2, "ch": 9 },
                "to": { "line": 2, "ch": 9 },
                "text": ["v"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 14209,
            "changeData": {
                "from": { "line": 2, "ch": 10, "sticky": null },
                "to": { "line": 2, "ch": 10, "sticky": null },
                "text": ["o"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 14315,
            "changeData": {
                "from": { "line": 2, "ch": 11, "sticky": null },
                "to": { "line": 2, "ch": 11, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 12, "sticky": null }
        },
        {
            "timestamp": 14407,
            "changeData": {
                "from": { "line": 2, "ch": 12, "sticky": null },
                "to": { "line": 2, "ch": 12, "sticky": null },
                "text": ["d"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 14917,
            "changeData": {
                "from": { "line": 2, "ch": 14, "sticky": "before" },
                "to": { "line": 2, "ch": 14, "sticky": "before" },
                "text": [" "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 15, "sticky": null }
        },
        {
            "timestamp": 15318,
            "changeData": {
                "from": { "line": 2, "ch": 15, "sticky": null },
                "to": { "line": 2, "ch": 15, "sticky": null },
                "text": ["{}"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 2, "ch": 16 }
        },
        {
            "timestamp": 15526,
            "changeData": {
                "from": { "line": 2, "ch": 16 },
                "to": { "line": 2, "ch": 16 },
                "text": ["", "", ""],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 15527,
            "changeData": {
                "from": { "line": 3, "ch": 0, "sticky": null },
                "to": { "line": 3, "ch": 0, "sticky": null },
                "text": ["    "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 16121,
            "changeData": {
                "from": { "line": 3, "ch": 4, "sticky": null },
                "to": { "line": 3, "ch": 4, "sticky": null },
                "text": ["p"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 5, "sticky": null }
        },
        {
            "timestamp": 16254,
            "changeData": {
                "from": { "line": 3, "ch": 5, "sticky": null },
                "to": { "line": 3, "ch": 5, "sticky": null },
                "text": ["r"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 6, "sticky": null }
        },
        {
            "timestamp": 16423,
            "changeData": {
                "from": { "line": 3, "ch": 6, "sticky": null },
                "to": { "line": 3, "ch": 6, "sticky": null },
                "text": ["i"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 7, "sticky": null }
        },
        {
            "timestamp": 16605,
            "changeData": {
                "from": { "line": 3, "ch": 7, "sticky": null },
                "to": { "line": 3, "ch": 7, "sticky": null },
                "text": ["n"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 8, "sticky": null }
        },
        {
            "timestamp": 16716,
            "changeData": {
                "from": { "line": 3, "ch": 8, "sticky": null },
                "to": { "line": 3, "ch": 8, "sticky": null },
                "text": ["t"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 9, "sticky": null }
        },
        {
            "timestamp": 17157,
            "changeData": {
                "from": { "line": 3, "ch": 9, "sticky": null },
                "to": { "line": 3, "ch": 9, "sticky": null },
                "text": ["f"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 17595,
            "changeData": {
                "from": { "line": 3, "ch": 10, "sticky": null },
                "to": { "line": 3, "ch": 10, "sticky": null },
                "text": ["()"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 11 }
        },
        {
            "timestamp": 18640,
            "changeData": {
                "from": { "line": 3, "ch": 11 },
                "to": { "line": 3, "ch": 11 },
                "text": ["\"\""],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 12 }
        },
        {
            "timestamp": 20564,
            "changeData": {
                "from": { "line": 3, "ch": 12 },
                "to": { "line": 3, "ch": 12 },
                "text": ["H"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 13, "sticky": null }
        },
        {
            "timestamp": 20666,
            "changeData": {
                "from": { "line": 3, "ch": 13, "sticky": null },
                "to": { "line": 3, "ch": 13, "sticky": null },
                "text": ["e"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 14, "sticky": null }
        },
        {
            "timestamp": 20917,
            "changeData": {
                "from": { "line": 3, "ch": 14, "sticky": null },
                "to": { "line": 3, "ch": 14, "sticky": null },
                "text": ["l"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 15, "sticky": null }
        },
        {
            "timestamp": 21088,
            "changeData": {
                "from": { "line": 3, "ch": 15, "sticky": null },
                "to": { "line": 3, "ch": 15, "sticky": null },
                "text": ["l"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 16, "sticky": null }
        },
        {
            "timestamp": 21258,
            "changeData": {
                "from": { "line": 3, "ch": 16, "sticky": null },
                "to": { "line": 3, "ch": 16, "sticky": null },
                "text": ["o"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 17, "sticky": null }
        },
        {
            "timestamp": 21436,
            "changeData": {
                "from": { "line": 3, "ch": 17, "sticky": null },
                "to": { "line": 3, "ch": 17, "sticky": null },
                "text": [" "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 18, "sticky": null }
        },
        {
            "timestamp": 21994,
            "changeData": {
                "from": { "line": 3, "ch": 18, "sticky": null },
                "to": { "line": 3, "ch": 18, "sticky": null },
                "text": ["ｓ"],
                "removed": [""],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 19, "sticky": null }
        },
        {
            "timestamp": 22102,
            "changeData": {
                "from": { "line": 3, "ch": 18, "sticky": null },
                "to": { "line": 3, "ch": 19, "sticky": null },
                "text": ["せ"],
                "removed": ["ｓ"],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 19, "sticky": null }
        },
        {
            "timestamp": 22391,
            "changeData": {
                "from": { "line": 3, "ch": 19, "sticky": null },
                "to": { "line": 3, "ch": 19, "sticky": null },
                "text": ["い"],
                "removed": [""],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 20, "sticky": null }
        },
        {
            "timestamp": 22761,
            "changeData": {
                "from": { "line": 3, "ch": 19, "sticky": null },
                "to": { "line": 3, "ch": 20, "sticky": null },
                "text": [""],
                "removed": ["い"],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 19, "sticky": null }
        },
        {
            "timestamp": 22963,
            "changeData": {
                "from": { "line": 3, "ch": 19, "sticky": null },
                "to": { "line": 3, "ch": 19, "sticky": null },
                "text": ["ｋ"],
                "removed": [""],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 20, "sticky": null }
        },
        {
            "timestamp": 23068,
            "changeData": {
                "from": { "line": 3, "ch": 19, "sticky": null },
                "to": { "line": 3, "ch": 20, "sticky": null },
                "text": ["か"],
                "removed": ["ｋ"],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 20, "sticky": null }
        },
        {
            "timestamp": 23178,
            "changeData": {
                "from": { "line": 3, "ch": 20, "sticky": null },
                "to": { "line": 3, "ch": 20, "sticky": null },
                "text": ["い"],
                "removed": [""],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 21, "sticky": null }
        },
        {
            "timestamp": 23240,
            "changeData": {
                "from": { "line": 3, "ch": 18, "sticky": null },
                "to": { "line": 3, "ch": 21, "sticky": null },
                "text": ["世界"],
                "removed": ["せかい"],
                "origin": "*compose"
            },
            "cursor": { "line": 3, "ch": 20, "sticky": null }
        },
        {
            "timestamp": 24169,
            "changeData": {
                "from": { "line": 3, "ch": 22, "sticky": "before" },
                "to": { "line": 3, "ch": 22, "sticky": "before" },
                "text": [";"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 3, "ch": 23, "sticky": null }
        },
        {
            "timestamp": 24509,
            "changeData": {
                "from": { "line": 3, "ch": 23, "sticky": null },
                "to": { "line": 3, "ch": 23, "sticky": null },
                "text": ["", ""],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 24509,
            "changeData": {
                "from": { "line": 4, "ch": 0, "sticky": null },
                "to": { "line": 4, "ch": 0, "sticky": null },
                "text": ["    "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 4, "sticky": null }
        },
        {
            "timestamp": 24674,
            "changeData": {
                "from": { "line": 4, "ch": 4, "sticky": null },
                "to": { "line": 4, "ch": 4, "sticky": null },
                "text": ["r"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 5, "sticky": null }
        },
        {
            "timestamp": 24746,
            "changeData": {
                "from": { "line": 4, "ch": 5, "sticky": null },
                "to": { "line": 4, "ch": 5, "sticky": null },
                "text": ["e"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 6, "sticky": null }
        },
        {
            "timestamp": 24899,
            "changeData": {
                "from": { "line": 4, "ch": 6, "sticky": null },
                "to": { "line": 4, "ch": 6, "sticky": null },
                "text": ["t"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 7, "sticky": null }
        },
        {
            "timestamp": 25066,
            "changeData": {
                "from": { "line": 4, "ch": 7, "sticky": null },
                "to": { "line": 4, "ch": 7, "sticky": null },
                "text": ["u"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 8, "sticky": null }
        },
        {
            "timestamp": 25225,
            "changeData": {
                "from": { "line": 4, "ch": 8, "sticky": null },
                "to": { "line": 4, "ch": 8, "sticky": null },
                "text": ["r"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 9, "sticky": null }
        },
        {
            "timestamp": 25334,
            "changeData": {
                "from": { "line": 4, "ch": 9, "sticky": null },
                "to": { "line": 4, "ch": 9, "sticky": null },
                "text": ["n"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 10, "sticky": null }
        },
        {
            "timestamp": 25668,
            "changeData": {
                "from": { "line": 4, "ch": 10, "sticky": null },
                "to": { "line": 4, "ch": 10, "sticky": null },
                "text": [" "],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 11, "sticky": null }
        },
        {
            "timestamp": 26125,
            "changeData": {
                "from": { "line": 4, "ch": 11, "sticky": null },
                "to": { "line": 4, "ch": 11, "sticky": null },
                "text": ["0"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 12, "sticky": null }
        },
        {
            "timestamp": 26490,
            "changeData": {
                "from": { "line": 4, "ch": 12, "sticky": null },
                "to": { "line": 4, "ch": 12, "sticky": null },
                "text": [";"],
                "removed": [""],
                "origin": "+input"
            },
            "cursor": { "line": 4, "ch": 13, "sticky": null }
        },
        { "timestamp": 27757 }
    ]
}`
