package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"encoding/json"
	"fmt"
	"strings"
	"errors"
	"github.com/xwp/go-tide/src/message"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// SqsProvider represents an SQS queue.
type SqsProvider struct {
	session   *session.Session
	// Use sqsiface instead of sqs.SQS to benefit from the interface.
	sqs       sqsiface.SQSAPI
	QueueUrl  *string
	QueueName *string
}

// Send implements the required interface method to be a MessageProvider.
//
// This method sends a new SQS SendMessageInput message to SQS.
func (mgr SqsProvider) SendMessage(msg *message.Message) error {

	// Encode the task to send as the message body.
	taskEncoded, _ := json.Marshal(msg)

	// Create the message object.
	messageInput := &sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(taskEncoded)),
		QueueUrl:     mgr.QueueUrl,
	}

	// Change message if .fifo queue
	var messageGroupId string = fmt.Sprintf("%s-%s", msg.RequestClient, msg.Slug)
	if strings.HasSuffix(*mgr.QueueName, ".fifo") {
		messageInput.MessageGroupId = &messageGroupId
	}

	// Send the message and check for errors.
	_, err := mgr.sqs.SendMessage(messageInput)

	if err != nil {
		return err
	}

	return nil
}

// Receive implements the required interface method to be a MessageProvider.
//
// This method sends a ReceiveMessageInput message to SQS and converts the message into a *task.Task object.
func (mgr SqsProvider) GetNextMessage() (*message.Message, error) {
	var returnMessage message.Message

	// Prepare the message
	messageInput := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            mgr.QueueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(300), // 300 seconds : 5 minutes
		WaitTimeSeconds:     aws.Int64(0),
	}

	// Retrieve the message from SQS
	result, err := mgr.sqs.ReceiveMessage(messageInput)

	if err != nil {
		return nil, err
	}

	// Attempt to unmarshal the message body into the returnTask.
	if len(result.Messages) != 0 {
		body := result.Messages[0].Body
		err = json.Unmarshal([]byte(*body), &returnMessage)

		// Return the queue receipt so that the message can be deleted.
		returnMessage.ExternalRef = result.Messages[0].ReceiptHandle
		return &returnMessage, err
	}

	return nil, errors.New("could not retrieve message")
}

// Delete implements the required interface method to be a MessageProvider.
//
// This method deletes a message from the queue.
func (mgr SqsProvider) DeleteMessage(reference *string) error {
	_, err := mgr.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      mgr.QueueUrl,
		ReceiptHandle: reference,
	})

	if err != nil {
		return err
	}

	return nil
}

// getQueueUrl requests the queueURL from SQS.
func getQueueUrl(svc sqsiface.SQSAPI, name string) (string, error) {
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})

	if err != nil {
		return "", err
	}
	return *result.QueueUrl, nil
}

// getSession establishes a new SQS session.
func getSession(region, key, secret string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	return sess, err
}

// NewSqsProvider is a convenience method to return a new *SqsProvider instance.
func NewSqsProvider(region, key, secret, queue string) *SqsProvider {

	sess, _ := getSession(region, key, secret);
	svc := sqs.New(sess)
	queueUrl, _ := getQueueUrl(svc, queue)

	return &SqsProvider{
		session:   sess,
		sqs:       svc,
		QueueUrl:  &queueUrl,
		QueueName: &queue,
	}
}
