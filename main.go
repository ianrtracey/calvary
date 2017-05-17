package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/ianrtracey/calvary/deployment"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	var fileName string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "function name",
			Destination: &functionName,
		},
		cli.StringFlag{
			Name:        "file",
			Usage:       "file name to upload",
			Destination: &fileName,
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
				if len(fileName) <= 0 {
					fmt.Println("create: file required")
					return nil
				}
				if len(functionName) <= 0 {
					fmt.Println("create: functionName required")
					return nil
				}
				fileContents, err := ioutil.ReadFile(fileName)
				if err != nil {
					fmt.Println(err)
					return nil
				}

				functionHandlerFile := strings.Split(filepath.Base(fileName), ".")[0]
				params := &lambda.CreateFunctionInput{
					Code: &lambda.FunctionCode{
						ZipFile: []byte(fileContents),
					},
					FunctionName: aws.String(functionName),
					Handler:      aws.String(fmt.Sprintf("%s.handler", functionHandlerFile)),
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
		{
			Name:  "init",
			Usage: "initialize a function file",
			Action: func(c *cli.Context) error {
				// not sure if this is the right way to handle this
				if len(functionName) <= 0 {
					fmt.Println("init: functionName required")
					return errors.New("init: functionName required")
				}
				nodeScaffolding := deployment.GetNodeFunctionFileScaffolding()
				file, err := os.Create(fmt.Sprintf("%s.js", functionName))
				defer file.Close()
				if err != nil {
					fmt.Println("error creating file")
					return nil
				}
				fmt.Fprintf(file, nodeScaffolding)
				return nil
			},
		},
	}

	app.Run(os.Args)

}
