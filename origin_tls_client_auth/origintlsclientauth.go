// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// OriginTLSClientAuthService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOriginTLSClientAuthService]
// method instead.
type OriginTLSClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *HostnameService
	Settings  *SettingService
}

// NewOriginTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthService(opts ...option.RequestOption) (r *OriginTLSClientAuthService) {
	r = &OriginTLSClientAuthService{}
	r.Options = opts
	r.Hostnames = NewHostnameService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *OriginTLSClientAuthService) New(ctx context.Context, params OriginTLSClientAuthNewParams, opts ...option.RequestOption) (res *OriginTLSClientAuthNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *OriginTLSClientAuthService) List(ctx context.Context, query OriginTLSClientAuthListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OriginTLSClientAuthListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Certificates
func (r *OriginTLSClientAuthService) ListAutoPaging(ctx context.Context, query OriginTLSClientAuthListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OriginTLSClientAuthListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Certificate
func (r *OriginTLSClientAuthService) Delete(ctx context.Context, certificateID string, body OriginTLSClientAuthDeleteParams, opts ...option.RequestOption) (res *OriginTLSClientAuthDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", body.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Certificate Details
func (r *OriginTLSClientAuthService) Get(ctx context.Context, certificateID string, query OriginTLSClientAuthGetParams, opts ...option.RequestOption) (res *OriginTLSClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthNewResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthNewResponse interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginTLSClientAuthListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                              `json:"private_key"`
	JSON       originTLSClientAuthListResponseJSON `json:"-"`
}

// originTLSClientAuthListResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthListResponse]
type originTLSClientAuthListResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthDeleteResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthDeleteResponse interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthGetResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthGetResponse interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginTLSClientAuthNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r OriginTLSClientAuthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthNewResponseEnvelope struct {
	Errors   []OriginTLSClientAuthNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthNewResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthNewResponseEnvelope]
type originTLSClientAuthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthNewResponseEnvelopeErrors]
type originTLSClientAuthNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthNewResponseEnvelopeMessages]
type originTLSClientAuthNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthNewResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthNewResponseEnvelopeSuccessTrue OriginTLSClientAuthNewResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthDeleteResponseEnvelope struct {
	Errors   []OriginTLSClientAuthDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthDeleteResponseEnvelope]
type originTLSClientAuthDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeErrors]
type originTLSClientAuthDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeMessages]
type originTLSClientAuthDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue OriginTLSClientAuthDeleteResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthGetResponseEnvelope]
type originTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthGetResponseEnvelopeErrors]
type originTLSClientAuthGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthGetResponseEnvelopeMessages]
type originTLSClientAuthGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthGetResponseEnvelopeSuccessTrue OriginTLSClientAuthGetResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
