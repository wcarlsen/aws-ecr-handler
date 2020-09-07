package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// Delete ECR repository input
func DeleteECRInput(repoName string, accountId string) *ecr.DeleteRepositoryInput {
	input := &ecr.DeleteRepositoryInput{
		RepositoryName: aws.String(repoName),
		Force:          aws.Bool(true),
		RegistryId:     aws.String(accountId),
	}
	return input
}

// Delete ECR repository
func DeleteECR(repoName string, accountId string) (*ecr.DeleteRepositoryOutput, error) {
	input := DeleteECRInput(repoName, accountId)
	svc := ECRSession()
	return svc.DeleteRepository(input)
}
