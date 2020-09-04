// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateToken struct {
}

func (*validateOpCreateToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterClient struct {
}

func (*validateOpRegisterClient) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterClient) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterClientInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterClientInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDeviceAuthorization struct {
}

func (*validateOpStartDeviceAuthorization) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDeviceAuthorization) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDeviceAuthorizationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDeviceAuthorizationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateToken{}, middleware.After)
}

func addOpRegisterClientValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterClient{}, middleware.After)
}

func addOpStartDeviceAuthorizationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDeviceAuthorization{}, middleware.After)
}

func validateOpCreateTokenInput(v *CreateTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTokenInput"}
	if v.ClientId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientId"))
	}
	if v.GrantType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GrantType"))
	}
	if v.DeviceCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceCode"))
	}
	if v.ClientSecret == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientSecret"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterClientInput(v *RegisterClientInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterClientInput"}
	if v.ClientType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientType"))
	}
	if v.ClientName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDeviceAuthorizationInput(v *StartDeviceAuthorizationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDeviceAuthorizationInput"}
	if v.ClientSecret == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientSecret"))
	}
	if v.StartUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartUrl"))
	}
	if v.ClientId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}