// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelOrder struct {
}

func (*validateOpCancelOrder) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelOrder) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelOrderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelOrderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateOrder struct {
}

func (*validateOpCreateOrder) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateOrder) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateOrderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateOrderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateOutpost struct {
}

func (*validateOpCreateOutpost) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateOutpost) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateOutpostInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateOutpostInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSite struct {
}

func (*validateOpCreateSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteOutpost struct {
}

func (*validateOpDeleteOutpost) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteOutpost) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteOutpostInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteOutpostInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSite struct {
}

func (*validateOpDeleteSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCatalogItem struct {
}

func (*validateOpGetCatalogItem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCatalogItem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCatalogItemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCatalogItemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOrder struct {
}

func (*validateOpGetOrder) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOrder) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOrderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOrderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOutpost struct {
}

func (*validateOpGetOutpost) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOutpost) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOutpostInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOutpostInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOutpostInstanceTypes struct {
}

func (*validateOpGetOutpostInstanceTypes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOutpostInstanceTypes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOutpostInstanceTypesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOutpostInstanceTypesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSiteAddress struct {
}

func (*validateOpGetSiteAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSiteAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSiteAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSiteAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSite struct {
}

func (*validateOpGetSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSiteInput(input); err != nil {
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

type validateOpUpdateOutpost struct {
}

func (*validateOpUpdateOutpost) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateOutpost) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateOutpostInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateOutpostInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSiteAddress struct {
}

func (*validateOpUpdateSiteAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSiteAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSiteAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSiteAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSite struct {
}

func (*validateOpUpdateSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSiteRackPhysicalProperties struct {
}

func (*validateOpUpdateSiteRackPhysicalProperties) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSiteRackPhysicalProperties) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSiteRackPhysicalPropertiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSiteRackPhysicalPropertiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelOrderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelOrder{}, middleware.After)
}

func addOpCreateOrderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateOrder{}, middleware.After)
}

func addOpCreateOutpostValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateOutpost{}, middleware.After)
}

func addOpCreateSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSite{}, middleware.After)
}

func addOpDeleteOutpostValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteOutpost{}, middleware.After)
}

func addOpDeleteSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSite{}, middleware.After)
}

func addOpGetCatalogItemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCatalogItem{}, middleware.After)
}

func addOpGetOrderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOrder{}, middleware.After)
}

func addOpGetOutpostValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOutpost{}, middleware.After)
}

func addOpGetOutpostInstanceTypesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOutpostInstanceTypes{}, middleware.After)
}

func addOpGetSiteAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSiteAddress{}, middleware.After)
}

func addOpGetSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSite{}, middleware.After)
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

func addOpUpdateOutpostValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateOutpost{}, middleware.After)
}

func addOpUpdateSiteAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSiteAddress{}, middleware.After)
}

func addOpUpdateSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSite{}, middleware.After)
}

func addOpUpdateSiteRackPhysicalPropertiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSiteRackPhysicalProperties{}, middleware.After)
}

func validateAddress(v *types.Address) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Address"}
	if v.AddressLine1 == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddressLine1"))
	}
	if v.City == nil {
		invalidParams.Add(smithy.NewErrParamRequired("City"))
	}
	if v.StateOrRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateOrRegion"))
	}
	if v.PostalCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PostalCode"))
	}
	if v.CountryCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CountryCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelOrderInput(v *CancelOrderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelOrderInput"}
	if v.OrderId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OrderId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateOrderInput(v *CreateOrderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateOrderInput"}
	if v.OutpostIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostIdentifier"))
	}
	if v.LineItems == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LineItems"))
	}
	if len(v.PaymentOption) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PaymentOption"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateOutpostInput(v *CreateOutpostInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateOutpostInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSiteInput(v *CreateSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSiteInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.OperatingAddress != nil {
		if err := validateAddress(v.OperatingAddress); err != nil {
			invalidParams.AddNested("OperatingAddress", err.(smithy.InvalidParamsError))
		}
	}
	if v.ShippingAddress != nil {
		if err := validateAddress(v.ShippingAddress); err != nil {
			invalidParams.AddNested("ShippingAddress", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteOutpostInput(v *DeleteOutpostInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteOutpostInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSiteInput(v *DeleteSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSiteInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCatalogItemInput(v *GetCatalogItemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCatalogItemInput"}
	if v.CatalogItemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogItemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOrderInput(v *GetOrderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOrderInput"}
	if v.OrderId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OrderId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOutpostInput(v *GetOutpostInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOutpostInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOutpostInstanceTypesInput(v *GetOutpostInstanceTypesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOutpostInstanceTypesInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSiteAddressInput(v *GetSiteAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSiteAddressInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if len(v.AddressType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AddressType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSiteInput(v *GetSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSiteInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
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
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
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
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateOutpostInput(v *UpdateOutpostInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateOutpostInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSiteAddressInput(v *UpdateSiteAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSiteAddressInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if len(v.AddressType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AddressType"))
	}
	if v.Address == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Address"))
	} else if v.Address != nil {
		if err := validateAddress(v.Address); err != nil {
			invalidParams.AddNested("Address", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSiteInput(v *UpdateSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSiteInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSiteRackPhysicalPropertiesInput(v *UpdateSiteRackPhysicalPropertiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSiteRackPhysicalPropertiesInput"}
	if v.SiteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SiteId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
