// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sts provides the client and types for making API
// requests to AWS Security Token Service.
//
// The AWS Security Token Service (STS) is a web service that enables you to
// request temporary, limited-privilege credentials for AWS Identity and Access
// Management (IAM) users or for users that you authenticate (federated users).
// This guide provides descriptions of the STS API. For more detailed information
// about using this service, go to Temporary Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html).
//
// As an alternative to using the API, you can use one of the AWS SDKs, which
// consist of libraries and sample code for various programming languages and
// platforms (Java, Ruby, .NET, iOS, Android, etc.). The SDKs provide a convenient
// way to create programmatic access to STS. For example, the SDKs take care
// of cryptographically signing requests, managing errors, and retrying requests
// automatically. For information about the AWS SDKs, including how to download
// and install them, see the Tools for Amazon Web Services page (http://aws.amazon.com/tools/).
//
// For information about setting up signatures and authorization through the
// API, go to Signing AWS API Requests (http://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html)
// in the AWS General Reference. For general information about the Query API,
// go to Making Query Requests (http://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html)
// in Using IAM. For information about using security tokens with other AWS
// products, go to AWS Services That Work with IAM (http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html)
// in the IAM User Guide.
//
// If you're new to AWS and need additional technical information about a specific
// AWS product, you can find the product's technical documentation at http://aws.amazon.com/documentation/
// (http://aws.amazon.com/documentation/).
//
// Endpoints
//
// The AWS Security Token Service (STS) has a default endpoint of http://sts.amazonaws.com
// that maps to the US East (N. Virginia) region. Additional regions are available
// and are activated by default. For more information, see Activating and Deactivating
// AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the IAM User Guide.
//
// For information about STS endpoints, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#sts_region)
// in the AWS General Reference.
//
// Recording API requests
//
// STS supports AWS CloudTrail, which is a service that records AWS calls for
// your AWS account and delivers log files to an Amazon S3 bucket. By using
// information collected by CloudTrail, you can determine what requests were
// successfully made to STS, who made the request, when it was made, and so
// on. To learn more about CloudTrail, including how to turn it on and find
// your log files, see the AWS CloudTrail User Guide (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html).
//
// See http://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15 for more information on this service.
//
// See sts package documentation for more information.
// http://docs.aws.amazon.com/sdk-for-go/api/service/sts/
//
// Using the Client
//
// To contact AWS Security Token Service with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// http://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// http://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Security Token Service client STS for more
// information on creating client for this service.
// http://docs.aws.amazon.com/sdk-for-go/api/service/sts/#New
package sts
