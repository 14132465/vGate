package data

import (
	"encoding/json"
	"testing"
)

func TestBaseMsgGetters(t *testing.T) {
	raw := json.RawMessage(`{"x":1}`)
	m := &BaseMsg{Cmd: Request, Topic: "t1", Content: raw}
	if m.GetCmd() != Request {
		t.Fatalf("GetCmd: got %q want %q", m.GetCmd(), Request)
	}
	if m.GetTopic() != "t1" {
		t.Fatalf("GetTopic: got %q want %q", m.GetTopic(), "t1")
	}
	if string(m.GetContent()) != string(raw) {
		t.Fatalf("GetContent mismatch")
	}
}

func TestHeartbeatMsgSingleton(t *testing.T) {
	a := HeartbeatMsg()
	b := HeartbeatMsg()
	if a != b {
		t.Fatal("HeartbeatMsg should return same instance")
	}
	if a.Cmd != Heartbeat {
		t.Fatalf("Cmd: got %q want %q", a.Cmd, Heartbeat)
	}
}

func TestCustomMessageMarshalJSON_OmitsHiddenFields(t *testing.T) {
	c := CustomMessage{
		WsMsg: WsMsg{
			BaseMsg: BaseMsg{
				Cmd:     Request,
				Topic:   "topic-a",
				Content: json.RawMessage(`{}`),
			},
			ServerName: "srv",
			SessionId:  42,
			SecretKey:  "secret",
		},
		HideFields: []string{"secretKey", "sessionId"},
	}
	b, err := json.Marshal(&c)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if _, ok := m["secretKey"]; ok {
		t.Error("secretKey should be hidden")
	}
	if _, ok := m["sessionId"]; ok {
		t.Error("sessionId should be hidden")
	}
	// WsMsg embeds BaseMsg; MarshalJSON emits it as a nested object (field name "BaseMsg").
	if _, ok := m["cmd"]; ok {
		t.Error("cmd should not be at top level; embedded BaseMsg is nested")
	}
}
