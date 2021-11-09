package limsdk

import (
	"context"
	"testing"
	"time"

	"github.com/lim-team/LiMaoSDK/core"
	"github.com/stretchr/testify/assert"
)

func TestOnEvents(t *testing.T) {
	sdk := New(core.NewOptions(core.WithTest(true)))

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)

	go sdk.OnEvents(func(eventResult *core.EventResult) {
		assert.Equal(t, int64(1), eventResult.Events[0].EventID)
		eventResult.ACK()
		cancel()
	})

	sdk.API.MockAddEvent(&core.Event{
		EventID: 1,
		Message: &core.Message{},
	})

	<-timeoutCtx.Done()

	assert.Equal(t, context.Canceled, timeoutCtx.Err())

}
