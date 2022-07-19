package access_private_member

import (
	"github.com/mstmdev/puzzle-go/asm/struct/model"
)

func GetAge(account model.Account) (result int)

func GetPassword(account model.Account) (result string)
