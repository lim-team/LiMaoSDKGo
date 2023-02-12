package core

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
)

type API interface {
	GetEvents(offsetEventID int64) (*getEventResp, error)
	MockAddEvent(event *Event)
	// AnswerInlineQuery 响应查询
	AnswerInlineQuery(result *InlineQueryResult) error
	// SendMessage 发送消息
	SendMessage(message *MessageReq) (*MessageResp, error)
	// Typing 输入中
	// channelID 输入中显示在的频道ID
	Typing(channelID string, channelType uint8) error
}

type LIMAPI struct {
	opts *Options
	Log
	mockEvents []*Event
}

func NewAPI(opts *Options) API {
	return &LIMAPI{
		opts:       opts,
		Log:        NewTLog("API"),
		mockEvents: make([]*Event, 0),
	}
}

// GetEvents 获取事件
func (a *LIMAPI) GetEvents(offsetEventID int64) (*getEventResp, error) {

	if a.opts.Test {
		return &getEventResp{
			Status:  1,
			Results: a.mockEvents,
		}, nil
	}

	data, err := a.request("getEvents", map[string]interface{}{
		"event_id": offsetEventID,
	})
	if err != nil {
		return nil, err
	}
	var getEventResp *getEventResp
	err = ReadJsonByByte(data, &getEventResp)
	if err != nil {
		return nil, err
	}
	return getEventResp, nil
}

func (a *LIMAPI) MockAddEvent(event *Event) {
	a.mockEvents = append(a.mockEvents, event)
}

func (a *LIMAPI) AnswerInlineQuery(result *InlineQueryResult) error {
	_, err := a.request("answerInlineQuery", result)
	return err
}

func (a *LIMAPI) SendMessage(message *MessageReq) (*MessageResp, error) {
	resp, err := a.request("sendMessage", message)
	if err != nil {
		return nil, err
	}
	result := gjson.ParseBytes(resp)
	messageID := result.Get("message_id").Int()
	messageSeq := result.Get("message_seq").Uint()
	return &MessageResp{
		MessageID:   messageID,
		MessageSeq:  uint32(messageSeq),
		ClientMsgNo: result.Get("client_msg_no").Str,
	}, err
}

func (a *LIMAPI) Typing(channelID string, channelType uint8) error {
	_, err := a.request("typing", map[string]interface{}{
		"channel_id":   channelID,
		"channel_type": channelType,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *LIMAPI) request(method string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/robots/%s/%s/%s", a.opts.APIURL, a.opts.RobotID, a.opts.AppKey, method)
	resp, err := Post(url, []byte(ToJson(payload)), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status:%d", resp.StatusCode)
	}
	return []byte(resp.Body), nil
}
