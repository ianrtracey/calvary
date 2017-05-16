package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/urfave/cli"
	"os"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	}))
	service := lambda.New(sess)

	app := cli.NewApp()
	app.Name = "Calvary"
	app.Usage = "Your team's cli"

	var functionName string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "function name",
			Destination: &functionName,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "lists all functions available",
			Action: func(c *cli.Context) error {
				params := &lambda.ListFunctionsInput{}

				resp, err := service.ListFunctions(params)

				if err != nil {
					fmt.Println(err.Error())
					return nil
				}
				fmt.Println(resp)
				return nil
			},
		},
		{
			Name:    "invoke",
			Aliases: []string{"i"},
			Usage:   "invokes a function",
			Action: func(c *cli.Context) error {
				fmt.Println(functionName)
				invokeParams := &lambda.InvokeInput{
					FunctionName: aws.String(functionName),
					Payload:      []byte(nil),
				}
				invokeResp, err := service.Invoke(invokeParams)
				if err != nil {
					fmt.Println(err.Error())
					return nil
				}
				payload := string(invokeResp.Payload[:])
				fmt.Println(payload)
				return nil
			},
		},
	}

	app.Run(os.Args)

}
