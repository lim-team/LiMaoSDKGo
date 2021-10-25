package core

type EventsListener func(events []*Event)

type Event struct {
	EventID int64    `json:"event_id"`
	Message *Message `json:"message"`
}

type Message struct {
	MessageID  int64  `json:"message_id"`  // 服务端的消息ID(全局唯一)
	MessageSeq uint32 `json:"message_seq"` // 消息序列号 （用户唯一，有序递增）
	FromUID    string `json:"from_uid"`    // 发送者UID
	Timestamp  int32  `json:"timestamp"`   // 服务器消息时间戳(10位，到秒)
	Channel
	Payload Payload `json:"payload"` // 消息正文
}

type Payload struct {
	Type        int        `json:"type"`
	Content     string     `json:"content"`
	Entities    []*Entitiy `json:"entities"`
	PayloadText            // 文本消息
}

type PayloadText struct {
	Content string `json:"content"`
}

type Entitiy struct {
	Length int    `json:"length"`
	Offset int    `json:"offset"`
	Type   string `json:"type"`
}

type MessageSeq struct {
}

type Channel struct {
	ChannelID   string `json:"channel_id,omitempty"`
	ChannelType uint8  `json:"channel_type,omitempty"`
}

type getEventResp struct {
	Status  int      `json:"status"`
	Results []*Event `json:"results"`
}
