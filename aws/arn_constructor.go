package aws

// Construct arn from accountId
func ArnConstructor(accountId string) string {
	return "arn:aws:iam::" + accountId + ":root"
}
