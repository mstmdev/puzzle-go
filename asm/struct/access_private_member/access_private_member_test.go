package access_private_member

import (
	"testing"

	"github.com/mstmdev/puzzle-go/asm/struct/model"
)

func TestAccessPrivateMember(t *testing.T) {
	testCases := []struct {
		name     string
		age      int
		password string
	}{
		{"admin", 23, "123456"},
		{"user", 25, "resu"},
		{"anonymous", 0, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			account := model.NewAccount(tc.name, tc.age, tc.password)
			actualName := account.Name
			actualAge := GetAge(account)
			actualPassword := GetPassword(account)
			if tc.age != actualAge || tc.password != actualPassword {
				t.Errorf("expect => name:%s age:%d password:%s actual=> name:%s age:%d password:%s", tc.name, tc.age, tc.password, actualName, actualAge, actualPassword)
			} else {
				t.Logf("name:%s age:%d password:%s", actualName, actualAge, actualPassword)
			}
		})
	}
}
