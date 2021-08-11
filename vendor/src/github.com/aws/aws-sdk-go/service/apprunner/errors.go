// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apprunner

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServiceErrorException for service response error code
	// "InternalServiceErrorException".
	//
	// An unexpected service exception occurred.
	ErrCodeInternalServiceErrorException = "InternalServiceErrorException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// One or more input parameters aren't valid. Refer to the API action's document
	// page, correct the input parameters, and try the action again.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// You can't perform this action when the resource is in its current state.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A resource doesn't exist for the specified Amazon Resource Name (ARN) in
	// your AWS account.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// App Runner can't create this resource. You've reached your account quota
	// for this resource type.
	//
	// For App Runner per-resource quotas, see AWS App Runner endpoints and quotas
	// (https://docs.aws.amazon.com/general/latest/gr/apprunner.html) in the AWS
	// General Reference.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServiceErrorException": newErrorInternalServiceErrorException,
	"InvalidRequestException":       newErrorInvalidRequestException,
	"InvalidStateException":         newErrorInvalidStateException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
}