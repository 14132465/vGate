package data

import (
	"strings"
	"testing"
)

func TestNoDecoderMsg_MsgSnId(t *testing.T) {
	n := NoDecoderMsg{SnId: 7, SessionId: 1, Msg: `{}`}
	if n.MsgSnId() != 7 {
		t.Fatalf("MsgSnId: got %d want 7", n.MsgSnId())
	}
}

func TestServerDecoder_ValidJSON(t *testing.T) {
	inner := `{"cmd":"request","topic":"t","content":{}}`
	nd := NoDecoderMsg{SessionId: 99, SnId: 1, Msg: inner}
	ws, err := ServerDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.Cmd != Request {
		t.Fatalf("Cmd: got %q", ws.Cmd)
	}
	if ws.SessionId != 0 {
		t.Fatalf("ServerDecoder should not set SessionId from outer: got %d", ws.SessionId)
	}
}

func TestServerDecoder_InvalidJSON(t *testing.T) {
	nd := NoDecoderMsg{SessionId: 55, SnId: 1, Msg: `{not json`}
	ws, err := ServerDecoder(nd)
	if err == nil {
		t.Fatal("expected error")
	}
	if ws.Cmd != Unknown {
		t.Fatalf("Cmd: got %q want %q", ws.Cmd, Unknown)
	}
	if ws.SessionId != 55 {
		t.Fatalf("SessionId: got %d want 55", ws.SessionId)
	}
	if !strings.Contains(string(ws.Content), "not json") {
		t.Fatalf("Content should preserve raw msg")
	}
}

func TestGateDecoder_RequestUsesOuterSessionId(t *testing.T) {
	inner := `{"cmd":"request","topic":"t","content":{}}`
	nd := NoDecoderMsg{SessionId: 1001, SnId: 2, Msg: inner}
	ws, err := GateDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.SessionId != 1001 {
		t.Fatalf("SessionId: got %d want 1001", ws.SessionId)
	}
}

func TestGateDecoder_ResponseKeepsInnerSessionId(t *testing.T) {
	inner := `{"cmd":"response","topic":"t","sessionId":777,"content":{}}`
	nd := NoDecoderMsg{SessionId: 9999, SnId: 3, Msg: inner}
	ws, err := GateDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.SessionId != 777 {
		t.Fatalf("Response should keep inner sessionId: got %d want 777", ws.SessionId)
	}
}

func TestGateDecoder_NoticeKeepsInnerSessionId(t *testing.T) {
	inner := `{"cmd":"notice","topic":"t","sessionId":3,"content":{}}`
	nd := NoDecoderMsg{SessionId: 888, SnId: 4, Msg: inner}
	ws, err := GateDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.SessionId != 3 {
		t.Fatalf("Notice sessionId: got %d want 3", ws.SessionId)
	}
}

func TestGateDecoder_HeartbeatUsesOuterSessionId(t *testing.T) {
	inner := `{"cmd":"heartbeat"}`
	nd := NoDecoderMsg{SessionId: 2002, SnId: 5, Msg: inner}
	ws, err := GateDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.SessionId != 2002 {
		t.Fatalf("SessionId: got %d want 2002", ws.SessionId)
	}
}

func TestGateDecoder_UnknownCmdMarksUnknown(t *testing.T) {
	inner := `{"cmd":"weird","topic":"x"}`
	nd := NoDecoderMsg{SessionId: 3003, SnId: 6, Msg: inner}
	ws, err := GateDecoder(nd)
	if err != nil {
		t.Fatal(err)
	}
	if ws.Cmd != Unknown {
		t.Fatalf("Cmd: got %q want %q", ws.Cmd, Unknown)
	}
	if ws.SessionId != 3003 {
		t.Fatalf("SessionId: got %d want 3003", ws.SessionId)
	}
}
