package reflect

import "time"

type Account struct {
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	IsEnabled  bool      `json:"is_enabled"`
	CreateTime time.Time `json:"create_time"`
	success    int       `json:"success"`
	fail       int       `json:"fail"`
}

func (a *Account) Login(userName string, password string) bool {
	success := a.UserName == userName && a.Password == password
	if success {
		a.success++
	} else {
		a.fail++
	}
	return success
}

func (a *Account) Disable() {
	a.IsEnabled = false
}
