package core

type EventsListener func(eventResult *EventResult)

type Event struct {
	EventID     int64        `json:"event_id"`
	Message     *Message     `json:"message"`
	InlineQuery *InlineQuery `json:"inline_query"`
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

type MessageReq struct {
	Channel         // 接受频道
	Payload Payload // 消息负载
}

type Channel struct {
	ChannelID   string `json:"channel_id,omitempty"`
	ChannelType uint8  `json:"channel_type,omitempty"`
}

type getEventResp struct {
	Status  int      `json:"status"`
	Results []*Event `json:"results"`
}

type EventResult struct {
	Events []*Event
}

type InlineQuery struct {
	SID         string `json:"sid"`
	ChannelID   string `json:"channel_id"`
	ChannelType uint8  `json:"channel_type"`
	FromUID     string `json:"from_uid"` // 发送者uid
	Query       string `json:"query"`    // 查询关键字
	Offset      string `json:"offset"`   // 偏移量
}

type ResultType string

const (
	ResultTypeGIF ResultType = "gif"
)

type InlineQueryResult struct {
	InlineQuerySID string `json:"inline_query_sid"`
	// 结果类型
	Type ResultType `json:"type"`
	// 结果ID
	ID         string      `json:"id"`
	Results    interface{} `json:"results"`
	NextOffset string      `json:"next_offset"` // 下一次偏移量
}

// gif 结果
type GifResult struct {
	URL string `json:"url"` // gif完整路径
	// option
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
