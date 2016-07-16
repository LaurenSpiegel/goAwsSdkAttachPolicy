package main

	import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"

)

func main() {

	// Switch out with real accessKey, real secretKey
	creds := credentials.NewStaticCredentials("WEJIQR0VJM32U1WZF37U","zN7piffWOzYRVBHw6o2qNXlM3EKLUOMpmhw27VUs", "")
	// Switch out with real accountId
	accountId := "203475832551"
	userName := "Ethan"
	policyName := "useTheBucket"
	policyArn := `arn:aws:iam::` + accountId + `:policy/` + policyName

	iamConfig := aws.NewConfig().
	WithCredentials(creds).
	WithEndpoint("http://127.0.0.1:8600").
	WithRegion("us-west-1")

  iamSess := session.New(iamConfig)
  iamSvc := iam.New(iamSess)

// Create User

  createUserParams := &iam.CreateUserInput{
      UserName: aws.String(userName), // Required
  }
  createResp, err := iamSvc.CreateUser(createUserParams)

  if err != nil {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
      return
  }

  // Pretty-print the response data.
  fmt.Println(createResp)



// Create Access Key for User

  createAccessKeyParams := &iam.CreateAccessKeyInput{
      UserName: aws.String(userName),
  }
  accessKeyResp, err := iamSvc.CreateAccessKey(createAccessKeyParams)

  if err != nil {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
      return
  }

  // Pretty-print the response data.
  fmt.Println(accessKeyResp)

	

	// Create Policy
	
	thePolicy := `{ "Version": "2012-10-17",
	"Statement": [
		{
			"Sid": "permissionsOnBucket",
			"Action": [
			"s3:GetBucketAcl",
			"s3:ListBucket",
			"s3:ListBucketMultipartUploads"
			],
			"Effect": "Allow",
			"Resource": "arn:aws:s3:::bestbucket"
		},
		{
			"Sid": "permissionsOnObjectsInBucket",
			"Action": [
			"s3:AbortMultipartUpload",
			"s3:DeleteObject",
			"s3:GetObject",
			"s3:GetObjectAcl",
			"s3:ListMultipartUploadParts",
			"s3:PutObject",
			"s3:PutObjectAcl"
			],
			"Effect": "Allow",
			"Resource": "arn:aws:s3:::bestbucket/*"
		}
	]
}`

	policyParams := &iam.CreatePolicyInput{
			PolicyName:     aws.String(policyName),     // Required
	    PolicyDocument: aws.String(thePolicy), // Required
	}

	policyResp, err := iamSvc.CreatePolicy(policyParams)

	if err != nil {
	    // Print the error, cast err to awserr.Error to get the Code and
	    // Message from an error.
	    fmt.Println(err.Error())
	    return
	}

	// Pretty-print the response data.
	fmt.Println(policyResp)



	// Attach Policy to User

	attachParams := &iam.AttachUserPolicyInput{
	    PolicyArn: aws.String(policyArn),      // Required
	    UserName:  aws.String(userName), // Required
	}
	attachResp, err := iamSvc.AttachUserPolicy(attachParams)

	if err != nil {
	    // Print the error, cast err to awserr.Error to get the Code and
	    // Message from an error.
	    fmt.Println(err.Error())
	    return
	}

	// Pretty-print the response data.
	fmt.Println(attachResp)


	// Create Bucket

	s3config := aws.NewConfig().
		WithCredentials(creds).
		WithEndpoint("http://127.0.0.1:8000").
		WithRegion("us-west-1").
		WithS3ForcePathStyle(true)

	  s3sess := session.New(s3config)
    s3svc := s3.New(s3sess)

    bucket := "bestbucket"


	result, err := s3svc.CreateBucket(&s3.CreateBucketInput{
	    Bucket: &bucket,
	})

	if err != nil {
	    fmt.Println("Failed to create bucket", err)
	    return
	}

	fmt.Println(result)

}