package core

import "errors"

// ContentType 正文类型
type ContentType int

const (
	// ---------- 聊天类 ----------

	// Text 文本消息
	Text ContentType = 1 // 文本消息

)

var (
	ErrorContentType = errors.New("错误类型")
)
