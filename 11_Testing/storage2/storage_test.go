package storage

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	var notifiedUser, notifiedMsg string
	saved := notifyUser
	defer func() { notifyUser = saved }()
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}
	const user = "joe@example.org"
	checkQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatal("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("notify user %s, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
