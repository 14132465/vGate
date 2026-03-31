package data

import (
	"net/http"
	"sync"
	"sync/atomic"
)

// Session会话结构体，包含客户端ID、会话状态、HTTP请求和响应对象等信息
type Session struct {
	UUID   int64 //客户端ID
	Status int8  //会话状态 0：未连接 1：已连接 2：已断开
	Resp   *http.ResponseWriter
	Req    *http.Request
}

type SessionManager struct {
	sessionMap map[int64]*Session //会话映射表，存储所有客户端的会话信息
	uuid       atomic.Int64
	mutex      sync.RWMutex
}

// 全局会话管理器实例
var SessionManagerInstance *SessionManager

// 初始化会话管理器实例
func init() {
	SessionManagerInstance = &SessionManager{
		sessionMap: make(map[int64]*Session),
	}
}

// 根据客户端ID获取会话信息
func (sm *SessionManager) GetSession(uuid int64) *Session {
	defer sm.mutex.RLocker().Unlock()
	sm.mutex.RLocker().Lock()
	if session, ok := sm.sessionMap[uuid]; ok {
		return session
	}
	return nil
}

// 添加会话信息
func (sm *SessionManager) AddSession(session *Session) *Session {

	if session.UUID <= 0 {
		id := sm.uuid.Add(1)
		session.UUID = id
		defer sm.mutex.Unlock()
		sm.mutex.Lock()
		sm.sessionMap[session.UUID] = session
	} else {
		//客户端ID已存在
	}
	return session

}

// 移除会话信息
func (sm *SessionManager) RemoveSession(uuid int64) {
	defer sm.mutex.Unlock()
	sm.mutex.Lock()
	delete(sm.sessionMap, uuid)
}

// 更新会话状态
func (sm *SessionManager) UpdateSessionStatus(uuid int64, status int8) {
	defer sm.mutex.Unlock()
	sm.mutex.Lock()
	if session, ok := sm.sessionMap[uuid]; ok {
		session.Status = status
	}
}
