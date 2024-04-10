// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AddressMapIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressMapIPService] method
// instead.
type AddressMapIPService struct {
	Options []option.RequestOption
}

// NewAddressMapIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapIPService(opts ...option.RequestOption) (r *AddressMapIPService) {
	r = &AddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AddressMapIPService) Update(ctx context.Context, addressMapID string, ipAddress string, params AddressMapIPUpdateParams, opts ...option.RequestOption) (res *AddressMapIPUpdateResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapIPUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", params.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an IP from a particular address map.
func (r *AddressMapIPService) Delete(ctx context.Context, addressMapID string, ipAddress string, params AddressMapIPDeleteParams, opts ...option.RequestOption) (res *AddressMapIPDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapIPDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", params.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [addressing.AddressMapIPUpdateResponseUnknown],
// [addressing.AddressMapIPUpdateResponseArray] or [shared.UnionString].
type AddressMapIPUpdateResponseUnion interface {
	ImplementsAddressingAddressMapIPUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressMapIPUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressMapIPUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressMapIPUpdateResponseArray []interface{}

func (r AddressMapIPUpdateResponseArray) ImplementsAddressingAddressMapIPUpdateResponseUnion() {}

// Union satisfied by [addressing.AddressMapIPDeleteResponseUnknown],
// [addressing.AddressMapIPDeleteResponseArray] or [shared.UnionString].
type AddressMapIPDeleteResponseUnion interface {
	ImplementsAddressingAddressMapIPDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressMapIPDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressMapIPDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressMapIPDeleteResponseArray []interface{}

func (r AddressMapIPDeleteResponseArray) ImplementsAddressingAddressMapIPDeleteResponseUnion() {}

type AddressMapIPUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapIPUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   AddressMapIPUpdateResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressMapIPUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressMapIPUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapIPUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressMapIPUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapIPUpdateResponseEnvelope]
type addressMapIPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPUpdateResponseEnvelopeSuccess bool

const (
	AddressMapIPUpdateResponseEnvelopeSuccessTrue AddressMapIPUpdateResponseEnvelopeSuccess = true
)

func (r AddressMapIPUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       addressMapIPUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapIPUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressMapIPUpdateResponseEnvelopeResultInfo]
type addressMapIPUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapIPDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapIPDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapIPDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   AddressMapIPDeleteResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressMapIPDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressMapIPDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapIPDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapIPDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapIPDeleteResponseEnvelope]
type addressMapIPDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPDeleteResponseEnvelopeSuccess bool

const (
	AddressMapIPDeleteResponseEnvelopeSuccessTrue AddressMapIPDeleteResponseEnvelopeSuccess = true
)

func (r AddressMapIPDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       addressMapIPDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapIPDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressMapIPDeleteResponseEnvelopeResultInfo]
type addressMapIPDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
