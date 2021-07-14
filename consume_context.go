package sqs

import (
	`context`
	`time`
)

const (
	consumeContextKey = "sqs.consume"
	defaultDelay      = 5 * time.Second
)

type consumeContext struct {
	// 延迟时间
	delay time.Duration
}

// WithDelay 设置延迟时间
func WithDelay(ctx context.Context, delay time.Duration) {
	if consumeContext, ok := ctx.Value(consumeContextKey).(*consumeContext); ok {
		consumeContext.delay = delay
	}
}

func withConsumeContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, consumeContextKey, &consumeContext{
		delay: defaultDelay,
	})
}

func delay(ctx context.Context) (delay time.Duration) {
	if consumeContext, ok := ctx.Value(consumeContextKey).(*consumeContext); ok {
		delay = consumeContext.delay
	} else {
		delay = defaultDelay
	}

	return
}
