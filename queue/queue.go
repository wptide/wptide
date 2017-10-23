package queue

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsManager struct {
	Session  *session.Session
	SQS      *sqs.SQS
	QueueUrl *string
}

type Task struct {
	ResponseAPIEndpoint string   `json:"response_api_endpoint"`
	SourceURL           string   `json:"source_url"`
	SourceType          string   `json:"source_type"`
	Audits              *[]Audit `json:"audits"`
	Force               bool     `json:"force"`
}

type Audit struct {
	Type    string        `json:"type"`
	Options *AuditOption  `json:"options"`
	Scoring *AuditScoring `json:"scoring"`
}

type AuditOption struct {
	Standard string `json:"standard"`
	Report   string `json:"report"`
}

type AuditScoring struct {
	WeightingsFile string `json:"weightings_file"`
}

func (mgr SqsManager) getStandardPluginTask(sourceUrl, sourceType string) *Task {
	return &Task{
		SourceURL:  sourceUrl,
		SourceType: sourceType,
		Audits: &[]Audit{
			{
				Type: "phpcs",
				Options: &AuditOption{
					Standard: "tide-default",
					Report:   "json",
				},
				Scoring: &AuditScoring{
					WeightingsFile: "tide/weightings.json",
				},
			},
		},
	}
}

func (mgr SqsManager) SendToSQS(sourceUrl, sourceType string) error {
	taskEncoded, _ := json.Marshal(mgr.getStandardPluginTask(sourceUrl, sourceType))

	_, err := mgr.SQS.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(taskEncoded)),
		QueueUrl:     mgr.QueueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}

func getQueueUrl(svc *sqs.SQS, name string) (string, error) {
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})

	if err != nil {
		return "", err
	}
	return *result.QueueUrl, nil
}

func getSession(region, key, secret string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	return sess, err
}

func New(region, key, secret, queue string) *SqsManager {

	sess, _ := getSession(region, key, secret);
	svc := sqs.New(sess)
	queueUrl, _ := getQueueUrl(svc, queue)

	return &SqsManager{
		Session:  sess,
		SQS:      svc,
		QueueUrl: &queueUrl,
	}
}
