package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/urfave/cli"
	"io/ioutil"
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
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a function",
			Action: func(c *cli.Context) error {
				fmt.Println("create")
				fileContents, err := ioutil.ReadFile("./testNode.zip")
				if err != nil {
					fmt.Println(err)
					return nil
				}

				encodedFileContents := base64.StdEncoding.EncodeToString([]byte(fileContents))
				fmt.Println(encodedFileContents)
				params := &lambda.CreateFunctionInput{
					Code: &lambda.FunctionCode{
						ZipFile: []byte(fileContents),
					},
					FunctionName: aws.String(functionName),
					Handler:      aws.String("testNode.handler"),
					Role:         aws.String("arn:aws:iam::259931312128:role/lambda"),
					Runtime:      aws.String("nodejs6.10"),
					Description:  aws.String("a function that does something cool"),
				}
				resp, err := service.CreateFunction(params)
				if err != nil {
					fmt.Println(err.Error())
					return nil
				}
				fmt.Println(resp)
				return nil
			},
		},
	}

	app.Run(os.Args)

}
