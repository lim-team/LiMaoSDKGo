package limsdk

import (
	"github.com/lim-team/LiMaoSDK/core"
)

type LiMaoSDK struct {
	eventManager *core.EventManager
	stopChan     chan bool
	API          core.API
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
}

func (l *LiMaoSDK) SendMessage(channel *core.Channel, payload *core.Payload) error {

	return nil
}

func (l *LiMaoSDK) Run() {
	l.eventManager.Start()

	<-l.stopChan
}

func (l *LiMaoSDK) Stop() {
	l.eventManager.Stop()
	close(l.stopChan)
}

func (l *LiMaoSDK) Start() {
	l.eventManager.Start()
}
