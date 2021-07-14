package sqs

import (
	`errors`
	`sync`

	`github.com/aws/aws-sdk-go-v2/aws`
	`github.com/aws/aws-sdk-go-v2/credentials`
	`github.com/aws/aws-sdk-go-v2/service/sqs`
	`github.com/storezhang/glog`
	`github.com/storezhang/pangu`
)

func newSqs(conf *pangu.Config, logger glog.Logger) (client *Client, err error) {
	panguConfig := new(panguConfig)
	if err = conf.Load(panguConfig); nil != err {
		return
	}

	accessKey := ""
	secretKey := ""
	region := ""
	if accessKey, err = getAccessKey(panguConfig.Aws); nil != err {
		return
	}
	if secretKey, err = getSecretKey(panguConfig.Aws); nil != err {
		return
	}
	if region, err = getRegion(panguConfig.Aws); nil != err {
		return
	}

	sqsConfig := panguConfig.Aws.Sqs
	options := sqs.Options{
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			sqsConfig.Session,
		)),
		// Logger: nil,
		Region: region,
	}

	sqsClient := sqs.New(options)
	queues := sqsConfig.Queues
	queueMap := make(map[string]*string, len(queues))
	for _, _queue := range queues {
		_queue := _queue
		queueMap[_queue.Label] = &_queue.Name
	}
	if "" != sqsConfig.Queue {
		queueMap[sqsConfig.Queue] = &sqsConfig.Queue
	}
	label := sqsConfig.Queue
	if "" == label {
		label = queues[0].Label
	}

	// 创建客户端
	client = &Client{
		client: sqsClient,

		defaultLabel:    label,
		queueMap:        queueMap,
		waitTimeSeconds: int32(sqsConfig.Wait),
		_queueUrlCache:  sync.Map{},
		logger:          logger,
	}

	return
}

func getAccessKey(aws awsConfig) (key string, err error) {
	key = aws.Credentials.Access
	if "" == key {
		key = aws.Sqs.Credentials.Access
	}
	if "" == key {
		err = errors.New("必须配置AccessKey")
	}

	return
}

func getSecretKey(aws awsConfig) (key string, err error) {
	key = aws.Credentials.Secret
	if "" == key {
		key = aws.Sqs.Credentials.Secret
	}
	if "" == key {
		err = errors.New("必须配置SecretKey")
	}

	return
}

func getRegion(aws awsConfig) (region string, err error) {
	region = aws.Region
	if "" == region {
		region = aws.Sqs.Region
	}
	if "" == region {
		err = errors.New("必须配置Region")
	}

	return
}
