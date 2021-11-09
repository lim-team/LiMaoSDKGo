package core

import "fmt"

type API interface {
	GetEvents(offsetEventID int64) (*getEventResp, error)
	MockAddEvent(event *Event)
	AnswerInlineQuery(result *InlineQueryResult) error
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

	data, err := a.request("getEvents", map[string]string{
		"event_id": fmt.Sprintf("%d", offsetEventID),
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

func (a *LIMAPI) request(method string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/robots/%s/%s/%s", a.opts.APIURL, a.opts.AppID, a.opts.AppKey, method)

	resp, err := Post(url, []byte(ToJson(payload)), nil)
	if err != nil {
		return nil, err
	}
	return []byte(resp.Body), nil
}
