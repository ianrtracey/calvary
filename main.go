package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	s "github.com/ianrtracey/calvary/service"
	"github.com/urfave/cli"
	"os"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	}))
	service := lambda.New(sess)

	// params := &lambda.ListFunctionsInput{}
	//
	// resp, err := service.ListFunctions(params)
	//
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	invokeParams := &lambda.InvokeInput{
		FunctionName: aws.String("testing-func"),
		Payload:      []byte(nil),
	}

	invokeResp, err := service.Invoke(invokeParams)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(invokeResp)
	payload := string(invokeResp.Payload[:])
	fmt.Println(payload)
	result := service.Add(1, 2)

	app := cli.NewApp()
	app.Name = "Calvary"
	app.Usage = "Your team's cli"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "lists all functions available",
			Action: func(c *cli.Context) error {
				fmt.Println("list")
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}
	app.Run(os.Args)

}
