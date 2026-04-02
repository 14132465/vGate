package data

import (
	"testing"
)

func TestSessionManager_GetAddRemove(t *testing.T) {
	sm := SessionManagerInstance
	s := &Session{UUID: 0, Status: 0, Conn: nil}
	out := sm.AddSession(s)
	if out.UUID <= 0 {
		t.Fatal("expected assigned UUID")
	}
	id := out.UUID
	defer sm.RemoveSession(id)

	got := sm.GetSession(id)
	if got == nil || got.UUID != id {
		t.Fatal("GetSession after AddSession")
	}
	sm.RemoveSession(id)
	if sm.GetSession(id) != nil {
		t.Fatal("expected removed")
	}
}

func TestSessionManager_UpdateStatus(t *testing.T) {
	sm := SessionManagerInstance
	s := &Session{UUID: 0, Status: 0, Conn: nil}
	out := sm.AddSession(s)
	id := out.UUID
	defer sm.RemoveSession(id)

	sm.UpdateSessionStatus(id, 2)
	got := sm.GetSession(id)
	if got.Status != 2 {
		t.Fatalf("Status: got %d want 2", got.Status)
	}
}

func TestSessionManager_ChangeId(t *testing.T) {
	sm := SessionManagerInstance
	s := &Session{UUID: 0, Status: 1, Conn: nil}
	out := sm.AddSession(s)
	oldID := out.UUID
	newID := int64(9102837465)
	defer sm.RemoveSession(newID)

	sm.ChangeId(oldID, newID)
	if sm.GetSession(oldID) != nil {
		t.Fatal("old id should be gone")
	}
	got := sm.GetSession(newID)
	if got == nil || got.UUID != newID {
		t.Fatal("expected session at new id")
	}
}

func TestServerManager_AddGetRemove(t *testing.T) {
	mgr := ServerManagerInstance
	srv := &Server{UUID: 0, Status: 1, Conn: nil}
	out := mgr.AddServer(srv)
	if out.UUID <= 0 {
		t.Fatal("expected assigned UUID")
	}
	id := out.UUID
	defer mgr.RemoveServer(id)

	if g := mgr.GetServerOnly(id); g == nil || g.UUID != id {
		t.Fatal("GetServerOnly")
	}
	mgr.RemoveServer(id)
	if mgr.GetServerOnly(id) != nil {
		t.Fatal("expected removed")
	}
}

func TestServerManager_GetAndCreateServer_FromSession(t *testing.T) {
	sessMgr := SessionManagerInstance
	srvMgr := ServerManagerInstance

	const sid int64 = 777700001
	sessMgr.sessionMap[sid] = &Session{UUID: sid, Status: 1, Conn: nil}
	defer sessMgr.RemoveSession(sid)
	defer srvMgr.RemoveServer(sid)

	got := srvMgr.GetAndCreateServer(sid)
	if got == nil || got.UUID != sid || got.Status != 1 {
		t.Fatalf("GetAndCreateServer: %+v", got)
	}
}

func TestServerManager_UpdateStatusAndGetAlls(t *testing.T) {
	mgr := ServerManagerInstance
	s1 := &Server{UUID: 0, Status: 0, Conn: nil}
	s2 := &Server{UUID: 0, Status: 0, Conn: nil}
	a := mgr.AddServer(s1)
	b := mgr.AddServer(s2)
	defer mgr.RemoveServer(a.UUID)
	defer mgr.RemoveServer(b.UUID)

	mgr.UpdateServerStatus(a.UUID, 2)
	if mgr.GetServerOnly(a.UUID).Status != 2 {
		t.Fatal("UpdateServerStatus")
	}
	all := mgr.GetAlls()
	if len(all) < 2 {
		t.Fatalf("GetAlls len: %d", len(all))
	}
}
