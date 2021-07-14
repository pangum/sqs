package sqs

import (
	`errors`
	`sync`

	`github.com/aws/aws-sdk-go-v2/aws`
	`github.com/aws/aws-sdk-go-v2/credentials`
	`github.com/aws/aws-sdk-go-v2/service/sqs`
	`github.com/storezhang/pangu`
)

func newSqs(conf *pangu.Config) (client *Client, err error) {
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

	options := sqs.Options{
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			panguConfig.Aws.Sqs.Session,
		)),
		// Logger: nil,
		Region: region,
	}

	sqsClient := sqs.New(options)
	queues := panguConfig.Aws.Sqs.Queues
	queueMap := make(map[string]*string, len(queues))
	for _, queue := range queues {
		queueMap[queue.Label] = &queue.Name
	}
	// 创建客户端
	client = &Client{
		client: sqsClient,

		defaultLabel:    queues[0].Label,
		queueMap:        queueMap,
		waitTimeSeconds: int32(panguConfig.Aws.Sqs.Wait),
		_queueUrlCache:  sync.Map{},
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
