package sqs

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/xwp/go-tide/src/message"
	"github.com/pkg/errors"
	"encoding/json"
)

type mockSqs struct {
	sqsiface.SQSAPI
	sendMessageOutput   *sqs.SendMessageOutput
	deleteMessageOutput *sqs.DeleteMessageOutput
}

var (
	// Provider to mock fifo queue.
	testQueue    string      = "test.fifo"
	testQueueUrl string      = "http://sqsurl/test.fifo"
	testProvider SqsProvider = SqsProvider{
		session:   &session.Session{},
		sqs:       &mockSqs{},
		QueueName: &testQueue,
		QueueUrl:  &testQueueUrl,
	}

	// Provider to mock non-fifo queue.
	testNonFifoQueue    string      = "test"
	testNonFifoQueueUrl string      = "http://sqsurl/test"
	testNonFifoProvider SqsProvider = SqsProvider{
		session:   &session.Session{},
		sqs:       &mockSqs{},
		QueueName: &testNonFifoQueue,
		QueueUrl:  &testNonFifoQueueUrl,
	}

	// Provider to mock an error response.
	failQueue    string      = "fail.fifo"
	failQueueUrl string      = "http://sqsurl/fail.fifo"
	failProvider SqsProvider = SqsProvider{
		session:   &session.Session{},
		sqs:       &mockSqs{},
		QueueName: &failQueue,
		QueueUrl:  &failQueueUrl,
	}

	// Provider to mock an empty response.
	emptyQueue    string      = "empty.fifo"
	emptyQueueUrl string      = "http://sqsurl/empty.fifo"
	emptyProvider SqsProvider = SqsProvider{
		session:   &session.Session{},
		sqs:       &mockSqs{},
		QueueName: &emptyQueue,
		QueueUrl:  &emptyQueueUrl,
	}
)

func (m mockSqs) DeleteMessage(in *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {

	if *in.ReceiptHandle == "fail-id" {
		return m.deleteMessageOutput, errors.New("Something went wrong.")
	}
	// contains filtered or unexported fields - so will be an empty struct of sqs.DeleteMessageOutput
	// https://docs.aws.amazon.com/sdk-for-go/api/service/sqs/#DeleteMessageOutput
	return m.deleteMessageOutput, nil
}

func (m mockSqs) ReceiveMessage(in *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {

	var messages []*sqs.Message

	if *in.QueueUrl != failQueueUrl {

		if *in.QueueUrl != emptyQueueUrl {
			fake := message.Message{
				Title: "Success!",
			}

			bodyBytes, _ := json.Marshal(fake)
			body := string(bodyBytes)
			msg := &sqs.Message{
				Body: &body,
			}
			messages = append(messages, msg)
		}
	} else {
		return nil, errors.New("Something went wrong.")
	}

	out := &sqs.ReceiveMessageOutput{
		Messages: messages,
	}

	return out, nil
}

func (m mockSqs) SendMessage(in *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {

	var msg *message.Message
	err := json.Unmarshal([]byte(*in.MessageBody), &msg)
	if err != nil {
		return m.sendMessageOutput, err
	}

	if msg.Title == "FAIL" {
		return m.sendMessageOutput, errors.New("Something went wrong.")
	}

	return m.sendMessageOutput, nil
}

func (m mockSqs) GetQueueUrl(in *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	url := "http://sqsurl/" + *in.QueueName
	return &sqs.GetQueueUrlOutput{
		QueueUrl: &url,
	}, nil
}

func TestSqsProvider_SendMessage(t *testing.T) {
	type args struct {
		msg *message.Message
	}
	tests := []struct {
		name    string
		mgr     SqsProvider
		args    args
		wantErr bool
	}{
		{
			name: "Test Simple Message - FIFO",
			mgr:  testProvider,
			args: args{
				&message.Message{},
			},
			wantErr: false,
		},
		{
			name: "Test Simple Message - Non-FIFO",
			mgr:  testNonFifoProvider,
			args: args{
				&message.Message{},
			},
			wantErr: false,
		},
		{
			name: "Test Failed Message",
			mgr:  testNonFifoProvider,
			args: args{
				&message.Message{
					Title: "FAIL",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mgr.SendMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SqsProvider.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSqsProvider_GetNextMessage(t *testing.T) {
	tests := []struct {
		name    string
		mgr     SqsProvider
		want    *message.Message
		wantErr bool
	}{
		{
			name: "Get Simple Message",
			mgr:  testProvider,
			want: &message.Message{
				Title: "Success!",
			},
			wantErr: false,
		},
		{
			name:    "Failed Message",
			mgr:     failProvider,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty Message",
			mgr:     emptyProvider,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mgr.GetNextMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("SqsProvider.GetNextMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqsProvider.GetNextMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqsProvider_DeleteMessage(t *testing.T) {
	type args struct {
		reference *string
	}

	successId := "success-id"
	failId := "fail-id"

	tests := []struct {
		name    string
		mgr     SqsProvider
		args    args
		wantErr bool
	}{
		{
			name: "Test Delete Message",
			mgr:  testProvider,
			args: args{
				reference: &successId,
			},
			wantErr: false,
		},
		{
			name: "Test Delete Message Error",
			mgr:  testProvider,
			args: args{
				reference: &failId,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mgr.DeleteMessage(tt.args.reference); (err != nil) != tt.wantErr {
				t.Errorf("SqsProvider.DeleteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getSession(t *testing.T) {
	type args struct {
		region string
		key    string
		secret string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			name: "Dummy Session",
			args: args{},
			want: reflect.TypeOf(&session.Session{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getSession(tt.args.region, tt.args.key, tt.args.secret)
			if !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("getSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSqsProvider(t *testing.T) {
	type args struct {
		region string
		key    string
		secret string
		queue  string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			name: "Test Provider Client",
			args: args{
				region: "us-west-2",
				key:    "random-key",
				secret: "so-secret",
				queue:  "test.fifo",
			},
			want: reflect.TypeOf(&SqsProvider{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSqsProvider(tt.args.region, tt.args.key, tt.args.secret, tt.args.queue); !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("NewSqsProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getQueueUrl(t *testing.T) {
	type args struct {
		svc  sqsiface.SQSAPI
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test getQueueUrl",
			args: args{
				svc:  &mockSqs{},
				name: "test.fifo",
			},
			want: "http://sqsurl/test.fifo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getQueueUrl(tt.args.svc, tt.args.name)
			if got != tt.want {
				t.Errorf("getQueueUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
