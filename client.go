package sqs

import (
	`context`
	`sync`
	`time`

	`github.com/aws/aws-sdk-go-v2/service/sqs`
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox/field`
)

// Client Sqs客户端封装
type Client struct {
	client *sqs.Client

	defaultLabel    string
	queueMap        map[string]*string
	waitTimeSeconds int32
	_queueUrlCache  sync.Map
	logger          glog.Logger
}

func (c *Client) Url(ctx context.Context, label string) (url *string, err error) {
	if cache, ok := c._queueUrlCache.Load(label); ok {
		url = cache.(*string)
	}
	if nil != url {
		return
	}

	var urlRsp *sqs.GetQueueUrlOutput
	if urlRsp, err = c.client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: c.queueMap[label],
	}); nil != err {
		return
	}
	url = urlRsp.QueueUrl
	c._queueUrlCache.Store(label, url)

	return
}

func (c *Client) Send(ctx context.Context, body string, opts ...optionSend) (output *SendOutput, err error) {
	options := defaultOptionsSend(c.defaultLabel)
	for _, opt := range opts {
		opt.applySend(options)
	}

	var url *string
	if url, err = c.Url(ctx, options.label); nil != err {
		return
	}

	var rsp *sqs.SendMessageOutput
	if rsp, err = c.client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:             &body,
		QueueUrl:                url,
		DelaySeconds:            options.delaySeconds,
		MessageAttributes:       options.messageAttributes,
		MessageSystemAttributes: options.messageSystemAttributes,
	}, options.fns...); nil != err {
		return
	}
	output = &SendOutput{SendMessageOutput: rsp}

	return
}

func (c *Client) Receive(ctx context.Context, opts ...optionReceive) (output *ReceiveOutput, err error) {
	options := defaultOptionsReceive(c.defaultLabel, c.waitTimeSeconds)
	for _, opt := range opts {
		opt.applyReceive(options)
	}

	var url *string
	if url, err = c.Url(ctx, options.label); nil != err {
		return
	}

	var rsp *sqs.ReceiveMessageOutput
	if rsp, err = c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              url,
		AttributeNames:        options.attributeNames,
		MaxNumberOfMessages:   0,
		MessageAttributeNames: options.messageAttributeNames,
		VisibilityTimeout:     options.visibilityTimeout,
		WaitTimeSeconds:       options.waitTimeSeconds,
	}, options.fns...); nil != err {
		return
	}
	output = &ReceiveOutput{ReceiveMessageOutput: rsp}

	return
}

func (c *Client) HandleReceive(ctx context.Context, handler Handler, opts ...optionReceive) (err error) {
	options := defaultOptionsReceive(c.defaultLabel, c.waitTimeSeconds)
	for _, opt := range opts {
		opt.applyReceive(options)
	}

	var url *string
	if url, err = c.Url(ctx, options.label); nil != err {
		return
	}

	var rsp *sqs.ReceiveMessageOutput
	for ; ; {
		if rsp, err = c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:              url,
			AttributeNames:        options.attributeNames,
			MaxNumberOfMessages:   1,
			MessageAttributeNames: options.messageAttributeNames,
			VisibilityTimeout:     options.visibilityTimeout,
			WaitTimeSeconds:       options.waitTimeSeconds,
		}, options.fns...); nil != err {
			return
		}

		if 1 != len(rsp.Messages) {
			continue
		}

		// 并行消费，加快消费速度
		go c.handleReceive(ctx, url, handler, rsp.Messages[0], options)
	}
}

func (c *Client) handleReceive(ctx context.Context, url *string, handler Handler, message types.Message, options *optionsReceive) {
	var (
		status ConsumeStatus
		err    error
	)

	// 重试
	for times := 0; times < options.maxRetry; times++ {
		if status, err = handler.OnMessage(withConsumeContext(ctx), &Message{Message: message}); nil == err {
			break
		} else {
			c.logger.Warn("消费出错，重试", field.Stringp("messageId", message.MessageId), field.Int("times", times))
			time.Sleep(options.retryDuration)
		}
	}
	if nil != err {
		_, err = c.client.ChangeMessageVisibility(ctx, &sqs.ChangeMessageVisibilityInput{
			QueueUrl:          url,
			ReceiptHandle:     message.ReceiptHandle,
			VisibilityTimeout: int32(options.retryDuration),
		})
		c.logger.Warn(
			"达到最大重试次数，消费仍然出错，改变消息可见性等待下一次消费",
			field.Stringp("messageId", message.MessageId),
			field.Duration("visibility", options.retryDuration),
		)

		return
	}

	switch status {
	case ConsumeStatusSuccess: // 消费成功，删除消息，不然会重复消费
		_, err = c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      url,
			ReceiptHandle: message.ReceiptHandle,
		})
	case ConsumeStatusLater: // 延迟消费，改变消息可见性，使其在指定的时间内再次被消费
		_, err = c.client.ChangeMessageVisibility(ctx, &sqs.ChangeMessageVisibilityInput{
			QueueUrl:          url,
			ReceiptHandle:     message.ReceiptHandle,
			VisibilityTimeout: int32(delay(ctx)),
		})
	case ConsumeStatusUnknown: // 默认状态，改变消息可见性，使前可以立即被消费
		_, err = c.client.ChangeMessageVisibility(ctx, &sqs.ChangeMessageVisibilityInput{
			QueueUrl:          url,
			ReceiptHandle:     message.ReceiptHandle,
			VisibilityTimeout: int32(time.Second),
		})
	}
}
