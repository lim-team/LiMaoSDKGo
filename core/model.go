package core

import "github.com/gookit/goutil/maputil"

type EventsListener func(eventResult *EventResult)

type Event struct {
	EventID     int64        `json:"event_id"`
	Message     *Message     `json:"message"`
	InlineQuery *InlineQuery `json:"inline_query"`
}

type Payload map[string]any

func (p Payload) Text() (PayloadText, error) {

	result := maputil.Data(p)

	contentType := ContentType(result.Int("type"))
	if contentType != Text {
		return PayloadText{}, ErrorContentType
	}

	entitieResultObjObj := result.Get("entities")
	entities := make([]*Entitiy, 0)

	if entitieResultObjObj != nil {
		entitieResultObjs := entitieResultObjObj.([]interface{})
		for _, entitieResultObj := range entitieResultObjs {
			entitieResults := entitieResultObj.([]map[string]any)
			for _, entitieResult := range entitieResults {
				entitieResultData := maputil.Data(entitieResult)
				entities = append(entities, &Entitiy{
					Length: entitieResultData.Int("length"),
					Offset: int(entitieResultData.Int("offset")),
					Type:   entitieResultData.Str("type"),
				})
			}
		}

	}

	return PayloadText{
		PayloadBase: PayloadBase{
			Type:     contentType,
			Entities: entities,
		},
		Content: result.Str("content"),
	}, nil
}

type Message struct {
	MessageID  int64  `json:"message_id"`  // 服务端的消息ID(全局唯一)
	MessageSeq uint32 `json:"message_seq"` // 消息序列号 （用户唯一，有序递增）
	FromUID    string `json:"from_uid"`    // 发送者UID
	Timestamp  int32  `json:"timestamp"`   // 服务器消息时间戳(10位，到秒)
	Channel
	Payload Payload `json:"payload"` // 消息正文
}

type PayloadBase struct {
	Type     ContentType `json:"type"`
	Entities []*Entitiy  `json:"entities"`
}

type PayloadText struct {
	PayloadBase
	Content string `json:"content"`
}

func NewPayloadText(content string, entities ...*Entitiy) PayloadText {

	return PayloadText{
		PayloadBase: PayloadBase{
			Type:     Text,
			Entities: entities,
		},
		Content: content,
	}
}

func (p PayloadText) Payload() Payload {
	mp := map[string]any{
		"type":    p.Type,
		"content": p.Content,
	}
	if len(p.Entities) > 0 {
		mp["entities"] = p.Entities
	}
	return Payload(mp)
}

type Entitiy struct {
	Length int    `json:"length"`
	Offset int    `json:"offset"`
	Type   string `json:"type"`
}

type MessageReq struct {
	Channel         // 接受频道
	Payload Payload `json:"payload"` // 消息正文 // 消息负载
}

type MessageResp struct {
	MessageID   int64  `json:"message_id"`    // 消息ID
	ClientMsgNo string `json:"client_msg_no"` // 客户端消息唯一编号
	MessageSeq  uint32 `json:"message_seq"`   // 消息序号
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
