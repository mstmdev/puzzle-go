package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestReflect_Struct_OutputMethodsAndFields(t *testing.T) {
	account := getAccount()

	tp := reflect.TypeOf(account)
	ty := tp.Elem()
	v := reflect.ValueOf(account).Elem()

	// output all methods
	for i := 0; i < tp.NumMethod(); i++ {
		m := tp.Method(i)
		t.Logf("index=%d func name=%s pkgpath=%s", m.Index, m.Name, m.PkgPath)
	}

	// output all fields
	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)
		fv := v.Field(i)
		var value any = fv.String()
		switch fv.Kind() {
		case reflect.Int:
			value = fv.Int()
		case reflect.Bool:
			value = fv.Bool()
		case reflect.Struct:
			value = fmt.Sprintf("%v", fv)
		}

		t.Logf("index=%d field name=%s value=%v pkgpath=%s tag=%s anonymous=%v exported=%v offset=%d", f.Index, f.Name, value, f.PkgPath, f.Tag.Get("json"), f.Anonymous, f.IsExported(), f.Offset)
	}
}

func TestReflect_Struct_CallMethod(t *testing.T) {
	account := getAccount()
	expectSuccess := account.success + 1
	expectFail := account.fail + 1
	tp := reflect.TypeOf(account)
	loginMethod, ok := tp.MethodByName("Login")
	if !ok {
		t.Errorf("get Login method failed")
		return
	}

	// login success
	result := loginMethod.Func.Call([]reflect.Value{reflect.ValueOf(account), reflect.ValueOf("admin"), reflect.ValueOf("123456")})
	success := result[0].Bool()
	if success && expectSuccess == account.success {
		t.Logf("login success =>%d", account.success)
	} else {
		t.Errorf("expect login success:%d, but actual:%d", expectSuccess, account.success)
	}

	// login fail
	result = loginMethod.Func.Call([]reflect.Value{reflect.ValueOf(account), reflect.ValueOf("admin"), reflect.ValueOf("000000")})
	success = result[0].Bool()
	if !success && expectFail == account.fail {
		t.Logf("login failed =>%d", account.success)
	} else {
		t.Errorf("expect login fail:%d, but actual:%d", expectFail, account.fail)
	}
}

func getAccount() *Account {
	account := &Account{
		UserName:   "admin",
		Password:   "123456",
		IsEnabled:  true,
		CreateTime: time.Now(),
		success:    1,
		fail:       2,
	}
	return account
}
