package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/kylelemons/godebug/diff"
)

func validateArgs() error {
	argLength := len(os.Args)
	if argLength != 3 {
		fmt.Println("error Argments.")
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := validateArgs(); err != nil {
		fmt.Println("validate error")
		os.Exit(1)
	}

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
