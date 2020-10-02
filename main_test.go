package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func TestValidateArgs(t *testing.T) {
	testArgs := []string{"taskdiff", "sample-app:1", "sample-app:2"}
	if err := validateArgs(testArgs); err != nil {
		t.Error(err)
	}
}

func TestDescribeTaskDefinition(t *testing.T) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ecs.New(sess)

	_, err := describeTaskDefinition("sample-app", svc)
	if err != nil {
		t.Error(err)
	}
}
