package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

const ECRPolicy = `
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Sid": "Allow pull from all",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::738063116313:root"
        }
      },
      "Action": [
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "ecr:BatchCheckLayerAvailability"
      ]
    }
  ]
}
`

// Policy document
type PolicuDocument struct {
	Version   string           `json:"Version"`
	Statement []StatementEntry `json:"Statement"`
}

// Statement entry
type StatementEntry struct {
	Sid       string    `json:"Sid"`
	Effect    string    `json:"Effect"`
	Principal Principal `json:"Principal"`
	Action    []string  `json:"Action"`
}

type Principal struct {
	Arn string `json:"AWS"`
}

// Set ECR policy input
func SetECRPolicyInput(repoName string, accountId string) *ecr.SetRepositoryPolicyInput {
	policy := PolicuDocument{Version: "2008-10-17", Statement: []StatementEntry{{Sid: "Allow pull from oxygen account", Effect: "Allow", Principal: Principal{"arn:aws:iam::738063116313:root"}, Action: []string{"ecr:GetDownloadUrlForLayer",
		"ecr:BatchGetImage",
		"ecr:BatchCheckLayerAvailability"}}}}
	pol, err := json.Marshal(&policy)

	if err != nil {

	}

	input := &ecr.SetRepositoryPolicyInput{
		Force:          aws.Bool(true),
		PolicyText:     aws.String(string(pol)),
		RegistryId:     aws.String(accountId),
		RepositoryName: aws.String(repoName),
	}
	return input
}

// Set ECR policy
func SetECRPolicy(repoName string, accountId string) (*ecr.SetRepositoryPolicyOutput, error) {
	input := SetECRPolicyInput(repoName, accountId)
	svc := ECRSession()
	return svc.SetRepositoryPolicy(input)
}
