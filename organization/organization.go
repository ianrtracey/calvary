package organization

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	org "github.com/aws/aws-sdk-go/service/organizations"
)

func Create() string {
	sess := session.Must(session.NewSession())
	svc := org.New(sess)
	params := &org.CreateOrganizationInput{
		FeatureSet: aws.String("CONSOLIDATED_BILLING"),
	}
	resp, err := svc.CreateOrganization(params)
	if err != nil {
		fmt.Println(err.Error())
		return "failed"
	}
	fmt.Println(resp)
	return "success"
}

func Invite(accountName string, email string) {
	sess := session.Must(session.NewSession())

	svc := org.New(sess)

	params := &org.CreateAccountInput{
		AccountName: aws.String(accountName), // Required
		Email:       aws.String(email),       // Required
	}

	resp, err := svc.CreateAccount(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

}
