package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	saved := notifyUser
	defer func() { notifyUser = saved }()
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	// ...模拟已使用980MB的情况...

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+"want substring %q", notifiedMsg, wantSubstring)
	}
}
