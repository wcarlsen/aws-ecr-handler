package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// Create ECR repository input
func CreateECRInput(repoName string) *ecr.CreateRepositoryInput {
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repoName),
		ImageScanningConfiguration: &ecr.ImageScanningConfiguration{
			ScanOnPush: aws.Bool(true),
		},
		ImageTagMutability: aws.String("IMMUTABLE"),
	}
	return input
}

// Create ECR repository
func CreateECR(repoName string) (*ecr.CreateRepositoryOutput, error) {
	input := CreateECRInput(repoName)
	svc := ECRSession()
	return svc.CreateRepository(input)
}
