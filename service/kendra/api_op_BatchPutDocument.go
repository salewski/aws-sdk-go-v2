// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchPutDocumentInput struct {
	_ struct{} `type:"structure"`

	// One or more documents to add to the index.
	//
	// Documents have the following file size limits.
	//
	//    * 5 MB total size for inline documents
	//
	//    * 50 MB total size for files from an S3 bucket
	//
	//    * 5 MB extracted text for any file
	//
	// For more information about file size and transaction per second quotas, see
	// Quotas (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
	//
	// Documents is a required field
	Documents []Document `min:"1" type:"list" required:"true"`

	// The identifier of the index to add the documents to. You need to create the
	// index first using the CreateIndex operation.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of a role that is allowed to run the BatchPutDocument
	// operation. For more information, see IAM Roles for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s BatchPutDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPutDocumentInput"}

	if s.Documents == nil {
		invalidParams.Add(aws.NewErrParamRequired("Documents"))
	}
	if s.Documents != nil && len(s.Documents) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Documents", 1))
	}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}
	if s.Documents != nil {
		for i, v := range s.Documents {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Documents", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchPutDocumentOutput struct {
	_ struct{} `type:"structure"`

	// A list of documents that were not added to the index because the document
	// failed a validation check. Each document contains an error message that indicates
	// why the document couldn't be added to the index.
	//
	// If there was an error adding a document to an index the error is reported
	// in your AWS CloudWatch log. For more information, see Monitoring Amazon Kendra
	// with Amazon CloudWatch Logs (https://docs.aws.amazon.com/kendra/latest/dg/cloudwatch-logs.html)
	FailedDocuments []BatchPutDocumentResponseFailedDocument `type:"list"`
}

// String returns the string representation
func (s BatchPutDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchPutDocument = "BatchPutDocument"

// BatchPutDocumentRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Adds one or more documents to an index.
//
// The BatchPutDocument operation enables you to ingest inline documents or
// a set of documents stored in an Amazon S3 bucket. Use this operation to ingest
// your text and unstructured text into an index, add custom attributes to the
// documents, and to attach an access control list to the documents added to
// the index.
//
// The documents are indexed asynchronously. You can see the progress of the
// batch using AWS CloudWatch. Any error messages related to processing the
// batch are sent to your AWS CloudWatch log.
//
//    // Example sending a request using BatchPutDocumentRequest.
//    req := client.BatchPutDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/BatchPutDocument
func (c *Client) BatchPutDocumentRequest(input *BatchPutDocumentInput) BatchPutDocumentRequest {
	op := &aws.Operation{
		Name:       opBatchPutDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchPutDocumentInput{}
	}

	req := c.newRequest(op, input, &BatchPutDocumentOutput{})

	return BatchPutDocumentRequest{Request: req, Input: input, Copy: c.BatchPutDocumentRequest}
}

// BatchPutDocumentRequest is the request type for the
// BatchPutDocument API operation.
type BatchPutDocumentRequest struct {
	*aws.Request
	Input *BatchPutDocumentInput
	Copy  func(*BatchPutDocumentInput) BatchPutDocumentRequest
}

// Send marshals and sends the BatchPutDocument API request.
func (r BatchPutDocumentRequest) Send(ctx context.Context) (*BatchPutDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchPutDocumentResponse{
		BatchPutDocumentOutput: r.Request.Data.(*BatchPutDocumentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchPutDocumentResponse is the response type for the
// BatchPutDocument API operation.
type BatchPutDocumentResponse struct {
	*BatchPutDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchPutDocument request.
func (r *BatchPutDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}