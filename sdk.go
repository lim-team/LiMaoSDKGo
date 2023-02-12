package limsdk

import (
	"sync"

	"github.com/lim-team/LiMaoSDKGo/core"
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

func (l *LiMaoSDK) SendMessage(channel *core.Channel, payload core.Payload) (*core.MessageResp, error) {

	return l.API.SendMessage(&core.MessageReq{
		Channel: core.Channel{
			ChannelID:   channel.ChannelID,
			ChannelType: channel.ChannelType,
		},
		Payload: payload,
	})
}

func (l *LiMaoSDK) AnswerInlineQuery(result *core.InlineQueryResult) error {

	return l.API.AnswerInlineQuery(result)
}

func (l *LiMaoSDK) Typing(channelID string, channelType uint8) error {

	return l.API.Typing(channelID, channelType)
}

func (l *LiMaoSDK) StopEvents() {
	close(l.stopChan)
}
