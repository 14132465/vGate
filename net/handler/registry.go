package handler

import (
	"fmt"
	"sync"
)

// Registry 处理器注册中心
type Registry struct {
	MsgHandlerCreates map[string]MsgHandlerCreate // topic -> MsgHandlerCreate
	mu                sync.RWMutex
}

// NewRegistry 创建注册中心
func NewRegistry() *Registry {
	return &Registry{
		MsgHandlerCreates: make(map[string]MsgHandlerCreate),
	}
}

// Register 注册处理器
func (r *Registry) Register(handlerCreate MsgHandlerCreate) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	topic := handlerCreate.Topic
	if _, exists := r.MsgHandlerCreates[topic]; exists {
		return fmt.Errorf("MsgHandlerCreate for topic %s already registered", topic)
	}

	r.MsgHandlerCreates[topic] = handlerCreate
	return nil
}

// GetMsgHandlerCreate 根据主题获取处理器
func (r *Registry) GetMsgHandlerCreate(topic string) (MsgHandlerCreate, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	MsgHandlerCreate, ok := r.MsgHandlerCreates[topic]
	return MsgHandlerCreate, ok
}

// Unregister 注销处理器
func (r *Registry) Unregister(topic string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.MsgHandlerCreates[topic]; !exists {
		return fmt.Errorf("MsgHandlerCreate for topic %s not found", topic)
	}

	delete(r.MsgHandlerCreates, topic)
	return nil
}

// ListTopics 列出所有已注册的主题
func (r *Registry) ListTopics() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	topics := make([]string, 0, len(r.MsgHandlerCreates))
	for topic := range r.MsgHandlerCreates {
		topics = append(topics, topic)
	}
	return topics
}
