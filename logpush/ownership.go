// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// OwnershipService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewOwnershipService] method instead.
type OwnershipService struct {
	Options []option.RequestOption
}

// NewOwnershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOwnershipService(opts ...option.RequestOption) (r *OwnershipService) {
	r = &OwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *OwnershipService) New(ctx context.Context, params OwnershipNewParams, opts ...option.RequestOption) (res *OwnershipNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OwnershipNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates ownership challenge of the destination.
func (r *OwnershipService) Validate(ctx context.Context, params OwnershipValidateParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854, err error) {
	opts = append(r.Options[:], opts...)
	var env OwnershipValidateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership/validate", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OwnershipNewResponse struct {
	Filename string                   `json:"filename"`
	Message  string                   `json:"message"`
	Valid    bool                     `json:"valid"`
	JSON     ownershipNewResponseJSON `json:"-"`
}

// ownershipNewResponseJSON contains the JSON metadata for the struct
// [OwnershipNewResponse]
type ownershipNewResponseJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r OwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   OwnershipNewResponse                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success OwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipNewResponseEnvelopeJSON    `json:"-"`
}

// ownershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelope]
type ownershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipNewResponseEnvelopeSuccess bool

const (
	OwnershipNewResponseEnvelopeSuccessTrue OwnershipNewResponseEnvelopeSuccess = true
)

func (r OwnershipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r OwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipValidateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   shared.UnnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success OwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipValidateResponseEnvelopeJSON    `json:"-"`
}

// ownershipValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipValidateResponseEnvelope]
type ownershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipValidateResponseEnvelopeSuccess bool

const (
	OwnershipValidateResponseEnvelopeSuccessTrue OwnershipValidateResponseEnvelopeSuccess = true
)

func (r OwnershipValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
