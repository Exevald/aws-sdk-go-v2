// Code generated by smithy-go-codegen DO NOT EDIT.

package invoicing

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/invoicing/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetInvoiceProfile struct {
}

func (*validateOpBatchGetInvoiceProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetInvoiceProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetInvoiceProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetInvoiceProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateInvoiceUnit struct {
}

func (*validateOpCreateInvoiceUnit) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateInvoiceUnit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateInvoiceUnitInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateInvoiceUnitInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteInvoiceUnit struct {
}

func (*validateOpDeleteInvoiceUnit) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteInvoiceUnit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteInvoiceUnitInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteInvoiceUnitInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInvoiceUnit struct {
}

func (*validateOpGetInvoiceUnit) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInvoiceUnit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInvoiceUnitInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInvoiceUnitInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListInvoiceSummaries struct {
}

func (*validateOpListInvoiceSummaries) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListInvoiceSummaries) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListInvoiceSummariesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListInvoiceSummariesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateInvoiceUnit struct {
}

func (*validateOpUpdateInvoiceUnit) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateInvoiceUnit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateInvoiceUnitInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateInvoiceUnitInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetInvoiceProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetInvoiceProfile{}, middleware.After)
}

func addOpCreateInvoiceUnitValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateInvoiceUnit{}, middleware.After)
}

func addOpDeleteInvoiceUnitValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteInvoiceUnit{}, middleware.After)
}

func addOpGetInvoiceUnitValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInvoiceUnit{}, middleware.After)
}

func addOpListInvoiceSummariesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListInvoiceSummaries{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateInvoiceUnitValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateInvoiceUnit{}, middleware.After)
}

func validateBillingPeriod(v *types.BillingPeriod) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BillingPeriod"}
	if v.Month == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Month"))
	}
	if v.Year == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Year"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDateInterval(v *types.DateInterval) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DateInterval"}
	if v.StartDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartDate"))
	}
	if v.EndDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndDate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInvoiceSummariesFilter(v *types.InvoiceSummariesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvoiceSummariesFilter"}
	if v.TimeInterval != nil {
		if err := validateDateInterval(v.TimeInterval); err != nil {
			invalidParams.AddNested("TimeInterval", err.(smithy.InvalidParamsError))
		}
	}
	if v.BillingPeriod != nil {
		if err := validateBillingPeriod(v.BillingPeriod); err != nil {
			invalidParams.AddNested("BillingPeriod", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInvoiceSummariesSelector(v *types.InvoiceSummariesSelector) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvoiceSummariesSelector"}
	if len(v.ResourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceTag(v *types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceTagList(v []types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTagList"}
	for i := range v {
		if err := validateResourceTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetInvoiceProfileInput(v *BatchGetInvoiceProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetInvoiceProfileInput"}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateInvoiceUnitInput(v *CreateInvoiceUnitInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateInvoiceUnitInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.InvoiceReceiver == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvoiceReceiver"))
	}
	if v.Rule == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Rule"))
	}
	if v.ResourceTags != nil {
		if err := validateResourceTagList(v.ResourceTags); err != nil {
			invalidParams.AddNested("ResourceTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteInvoiceUnitInput(v *DeleteInvoiceUnitInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteInvoiceUnitInput"}
	if v.InvoiceUnitArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvoiceUnitArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInvoiceUnitInput(v *GetInvoiceUnitInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInvoiceUnitInput"}
	if v.InvoiceUnitArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvoiceUnitArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListInvoiceSummariesInput(v *ListInvoiceSummariesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInvoiceSummariesInput"}
	if v.Selector == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Selector"))
	} else if v.Selector != nil {
		if err := validateInvoiceSummariesSelector(v.Selector); err != nil {
			invalidParams.AddNested("Selector", err.(smithy.InvalidParamsError))
		}
	}
	if v.Filter != nil {
		if err := validateInvoiceSummariesFilter(v.Filter); err != nil {
			invalidParams.AddNested("Filter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.ResourceTags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceTags"))
	} else if v.ResourceTags != nil {
		if err := validateResourceTagList(v.ResourceTags); err != nil {
			invalidParams.AddNested("ResourceTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.ResourceTagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceTagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateInvoiceUnitInput(v *UpdateInvoiceUnitInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateInvoiceUnitInput"}
	if v.InvoiceUnitArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvoiceUnitArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
