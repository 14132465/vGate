package data

import (
	"testing"
)

func TestBuildSubscriptionMsg(t *testing.T) {
	m := BuildSubscriptionMsg("top", "s1", "k")
	if m.Cmd != Subscription || m.Topic != "top" || m.ServerName != "s1" || m.SecretKey != "k" {
		t.Fatalf("%+v", m)
	}
}

func TestBuildUnSubscriptionMsg(t *testing.T) {
	m := BuildUnSubscriptionMsg("top", "s2")
	if m.Cmd != UnSubscription || m.Topic != "top" || m.ServerName != "s2" {
		t.Fatalf("%+v", m)
	}
}

func TestBuildNoticeMsg(t *testing.T) {
	m := BuildNoticeMsg("sk", "n", []byte(`{"a":1}`))
	if m.Cmd != Notice || m.Topic != "n" || m.SecretKey != "sk" {
		t.Fatalf("%+v", m)
	}
	if string(m.Content) != `{"a":1}` {
		t.Fatalf("Content: %s", m.Content)
	}
}

func TestBuildRequestMsg(t *testing.T) {
	m := BuildRequestMsg(10, "r", []byte(`{}`))
	if m.Cmd != Request || m.Topic != "r" || m.SessionId != 10 {
		t.Fatalf("%+v", m)
	}
}

func TestBuildResponseMsg(t *testing.T) {
	m := BuildResponseMsg(11, "r2", []byte(`[]`))
	if m.Cmd != Response || m.Topic != "r2" || m.SessionId != 11 {
		t.Fatalf("%+v", m)
	}
	if string(m.Content) != `[]` {
		t.Fatalf("Content: %s", m.Content)
	}
}
