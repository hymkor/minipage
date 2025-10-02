package outline

import (
	"strings"
	"testing"
)

func TestHeader(t *testing.T) {
	source := `
SQL-Bless
=========

サポートコマンド
---------------

### scoop インストーラーを使用する場合
`
	expect := []*Header{
		&Header{Level: 1, Title: "SQL-Bless", ID: "sql-bless"},
		&Header{Level: 2, Title: "サポートコマンド", ID: "%E3%82%B5%E3%83%9D%E3%83%BC%E3%83%88%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89"},
		&Header{Level: 3, Title: "scoop インストーラーを使用する場合", ID: "scoop-%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%A9%E3%83%BC%E3%82%92%E4%BD%BF%E7%94%A8%E3%81%99%E3%82%8B%E5%A0%B4%E5%90%88"},
	}
	result, err := makeOutline(strings.NewReader(source))
	if err != nil {
		t.Fatal(err.Error())
	}
	for i := range expect {
		if result[i].Level != expect[i].Level {
			t.Fatalf("Level: expect %d, but %d", expect[i].Level, result[i].Level)
		}
		if result[i].Title != expect[i].Title {
			t.Fatalf("Title: expect %#v, but %#v", expect[i].Title, result[i].Title)
		}
		if result[i].ID != expect[i].ID {
			t.Fatalf("ID: expect %#v, but %#v", expect[i].ID, result[i].ID)
		}
	}
}
