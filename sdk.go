package limsdk

import (
	"sync"

	"github.com/lim-team/LiMaoSDK/core"
)

type LiMaoSDK struct {
	eventManager *core.EventManager
	stopChan     chan bool
	API          core.API
	eventOnce    sync.Once
}

func New(opts *core.Options) *LiMaoSDK {
	opts.API = core.NewAPI(opts)
	return &LiMaoSDK{
		eventManager: core.NewEventManager(opts),
		stopChan:     make(chan bool),
		API:          opts.API,
	}
}

func (l *LiMaoSDK) OnEvents(listener core.EventsListener) {

	l.eventManager.SetEventsListener(listener)

	l.eventOnce.Do(func() {
		l.eventManager.Start()
	})

	<-l.stopChan
}

func (l *LiMaoSDK) SendMessage(channel *core.Channel, payload *core.Payload) error {

	return nil
}

func (l *LiMaoSDK) AnswerInlineQuery(result *core.InlineQueryResult) error {

	return l.API.AnswerInlineQuery(result)
}

func (l *LiMaoSDK) StopEvents() {
	close(l.stopChan)
}
