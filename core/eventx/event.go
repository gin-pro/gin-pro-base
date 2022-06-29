package eventx

import "context"

func InitEventBus(ctx context.Context, cancelFunc context.CancelFunc) *EventBus {
	return &EventBus{
		Subscribers: map[string]DataChannelSlice{},
		Ctx:         ctx,
		Cancel:      cancelFunc,
	}
}
