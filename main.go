package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/kylelemons/godebug/diff"
)

func validateArgs(args []string) error {
	argLength := len(args)
	if argLength != 3 {
		return errors.New("Two arguments must be specified")
	}

	return nil
}

func describeTaskDefinition(name string, svc *ecs.ECS) (*ecs.DescribeTaskDefinitionOutput, error) {
	input := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(name),
	}

	result, err := svc.DescribeTaskDefinition(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	if err := validateArgs(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	srcTaskDef := os.Args[1]
	destTaskDef := os.Args[2]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ecs.New(sess)

	srcResult, err := describeTaskDefinition(srcTaskDef, svc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	destResult, err := describeTaskDefinition(destTaskDef, svc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	diff := diff.Diff(srcResult.GoString(), destResult.GoString())
	fmt.Println(diff)
}
