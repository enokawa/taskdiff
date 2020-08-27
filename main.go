package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/kylelemons/godebug/diff"
)

func main() {
	args1 := os.Args[1]
	args2 := os.Args[2]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ecs.New(sess)

	src := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(args1),
	}

	srcResult, err := svc.DescribeTaskDefinition(src)
	if err != nil {
		fmt.Println(err)
	}

	dest := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(args2),
	}

	destResult, err := svc.DescribeTaskDefinition(dest)
	if err != nil {
		fmt.Println(err)
	}

	diff := diff.Diff(srcResult.GoString(), destResult.GoString())
	fmt.Println(diff)
}
