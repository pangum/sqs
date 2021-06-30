package sqs

import (
	`context`
	`errors`

	`github.com/aws/aws-sdk-go-v2/aws`
	`github.com/aws/aws-sdk-go-v2/credentials`
	`github.com/aws/aws-sdk-go-v2/service/sqs`
	`github.com/storezhang/pangu`
)

func newSqs(config *pangu.Config) (client *Client, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
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
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey)),
		Logger:      nil,
		Region:      region,
	}
	sqsClient := sqs.New(options)
	// 获取连接地址
	var urlRsp *sqs.GetQueueUrlOutput
	if urlRsp, err = sqsClient.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: &panguConfig.Aws.Sqs.Queue,
	}); nil != err {
		return
	}
	client = &Client{
		client:   sqsClient,
		queueUrl: *urlRsp.QueueUrl,
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
