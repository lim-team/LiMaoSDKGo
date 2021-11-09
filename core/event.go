package core

import (
	"time"

	"go.uber.org/zap"
)

type EventManager struct {
	opts           *Options
	currentEventID int64
	Log
	eventsChan     chan []*Event
	stopChan       chan bool
	eventsListener EventsListener
	mockEvents     []*Event // 模拟event 仅仅test模式有效
}

func NewEventManager(opts *Options) *EventManager {

	return &EventManager{
		opts:       opts,
		Log:        NewTLog("EventManager"),
		eventsChan: make(chan []*Event),
		stopChan:   make(chan bool),
		mockEvents: make([]*Event, 0),
	}
}

func (e *EventManager) loopEvent() {

	ticker := time.NewTicker(e.opts.EventLoopDuration)

	for {
		select {
		case <-ticker.C:
			getEventResp, err := e.getEvents()
			if err != nil {
				e.Warn("获取事件失败！", zap.Error(err))
				continue
			}
			if getEventResp != nil && getEventResp.Status == 1 {

				e.eventsChan <- getEventResp.Results
			}

		case <-e.stopChan:
			goto exit
		}
	}

exit:
	e.Info("退出事件监听")
}

func (e *EventManager) getEvents() (*getEventResp, error) {
	return e.opts.API.GetEvents(e.currentEventID)
}

func (e *EventManager) handleEvents() {
	for {
		select {
		case events := <-e.eventsChan:
			if e.eventsListener != nil {

				e.eventsListener(&EventResult{
					Events: events,
					ACK: func() {
						e.currentEventID = events[len(events)-1].EventID
					},
				})
			}
		case <-e.stopChan:
			goto exit

		}
	}
exit:
}

func (e *EventManager) SetEventsListener(eventsListener EventsListener) {
	e.eventsListener = eventsListener
}

func (e *EventManager) Start() {
	go e.handleEvents()
	go e.loopEvent()
}

func (e *EventManager) Stop() {
	close(e.stopChan)
}
