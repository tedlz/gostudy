package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "job@example.com"
	usage[user] = 980000000 // 模拟 980MB 的使用情况

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubString)
	}
}

func TestCheckQuotaNotifiesUserFix(t *testing.T) {
	// 保存和恢复原始的 notifyUser
	saved := notifyUser
	defer func() { notifyUser = saved }()

	// 使用虚假的 notifyUser 测试
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "job@example.com"
	usage[user] = 980000000 // 模拟 980MB 的使用情况

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubString)
	}
}
