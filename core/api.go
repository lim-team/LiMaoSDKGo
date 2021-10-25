package core

import "fmt"

type API interface {
	GetEvents(offsetEventID int64) (*getEventResp, error)
	MockAddEvent(event *Event)
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

	url := fmt.Sprintf("%s/robots/%s/%s/getEvents", a.opts.APIURL, a.opts.AppID, a.opts.AppKey)
	resp, err := Get(url, map[string]string{
		"event_id": fmt.Sprintf("%d", offsetEventID),
	}, nil)
	if err != nil {
		return nil, err
	}
	var getEventResp *getEventResp
	err = ReadJsonByByte([]byte(resp.Body), &getEventResp)
	if err != nil {
		return nil, err
	}
	return getEventResp, nil
}

func (a *LIMAPI) MockAddEvent(event *Event) {
	a.mockEvents = append(a.mockEvents, event)
}
