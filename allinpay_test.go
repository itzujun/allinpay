package allinpay

import (
	"os"
	"testing"
)

func TestPay_Request(t *testing.T) {
	pay, err := New(&Config{
		RequestUrl:  os.Getenv("AllinpayRequestUrl"),
		PrivateFile: os.Getenv("AllinpayPrivateFile"),
		PublicFile:  os.Getenv("AllinpayPublicFile"),
		Sysid:       os.Getenv("AllinpaySysid"),
		IsDebug:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	resp, err := pay.Request("MemberService", "createMember", map[string]interface{}{
		"bizUserId":  "test_user_1",
		"memberType": 3,
		"source":     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	t.Log(resp.CheckStatus())

	resp, err = pay.RequestAndCheckStatus("MemberService", "createMember", map[string]interface{}{
		"bizUserId":  "test_user_1",
		"memberType": 3,
		"source":     1,
	})
	t.Log(resp, err)
}

func TestMemberServiceSetRealName(t *testing.T) {
	pay, err := New(&Config{
		RequestUrl:  os.Getenv("AllinpayRequestUrl"),
		PrivateFile: os.Getenv("AllinpayPrivateFile"),
		PublicFile:  os.Getenv("AllinpayPublicFile"),
		Sysid:       os.Getenv("AllinpaySysid"),
		IsDebug:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	no, err := pay.GetConfig().Encrypt([]byte("111111111111111111"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := pay.RequestAndCheckStatus("MemberService", "setRealName", map[string]interface{}{
		"bizUserId":    "test_user_1",
		"isAuth":       true,
		"name":         "小李",
		"identityType": 1,
		"identityNo":   string(no),
	})
	t.Log(resp, err)
}
