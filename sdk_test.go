package limsdk

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/lim-team/LiMaoSDK/core"
	"github.com/stretchr/testify/assert"
)

func TestOnEvents(t *testing.T) {
	sdk := New(core.NewOptions(core.WithTest(true)))

	var waitG sync.WaitGroup
	waitG.Add(1)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)

	sdk.OnEvents(func(events []*core.Event) {
		assert.Equal(t, int64(1), events[0].EventID)
		cancel()
	})

	sdk.API.MockAddEvent(&core.Event{
		EventID: 1,
		Message: &core.Message{},
	})

	sdk.Start()

	<-timeoutCtx.Done()

	assert.Equal(t, context.Canceled, timeoutCtx.Err())

}
