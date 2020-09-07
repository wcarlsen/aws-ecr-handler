package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// Initialize ECR session
func ECRSession() *ecr.ECR {
	config := &aws.Config{Region: aws.String("eu-west-1")}
	svc := ecr.New(session.New(), config)
	return svc
}

// Initialize assume role ECR session
// credentials expire after 15 min
func AssumeRoleECRSession(roleArn string) *ecr.ECR {
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, roleArn)
	svc := ecr.New(sess, &aws.Config{Credentials: creds})
	return svc
}
