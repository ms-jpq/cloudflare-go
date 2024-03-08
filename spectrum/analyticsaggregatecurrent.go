// File generated from our OpenAPI spec by Stainless.

package spectrum

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AnalyticsAggregateCurrentService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAnalyticsAggregateCurrentService] method instead.
type AnalyticsAggregateCurrentService struct {
	Options []option.RequestOption
}

// NewAnalyticsAggregateCurrentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAnalyticsAggregateCurrentService(opts ...option.RequestOption) (r *AnalyticsAggregateCurrentService) {
	r = &AnalyticsAggregateCurrentService{}
	r.Options = opts
	return
}

// Retrieves analytics aggregated from the last minute of usage on Spectrum
// applications underneath a given zone.
func (r *AnalyticsAggregateCurrentService) Get(ctx context.Context, zone string, query AnalyticsAggregateCurrentGetParams, opts ...option.RequestOption) (res *[]AnalyticsAggregateCurrentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsAggregateCurrentGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/analytics/aggregate/current", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsAggregateCurrentGetResponse = interface{}

type AnalyticsAggregateCurrentGetParams struct {
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppIDParam param.Field[string] `query:"app_id_param"`
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppID param.Field[string] `query:"appID"`
	// Co-location identifier.
	ColoName param.Field[string] `query:"colo_name"`
}

// URLQuery serializes [AnalyticsAggregateCurrentGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsAggregateCurrentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsAggregateCurrentGetResponseEnvelope struct {
	Errors   []AnalyticsAggregateCurrentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsAggregateCurrentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []AnalyticsAggregateCurrentGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsAggregateCurrentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AnalyticsAggregateCurrentGetResponseEnvelope]
type analyticsAggregateCurrentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AnalyticsAggregateCurrentGetResponseEnvelopeErrors]
type analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AnalyticsAggregateCurrentGetResponseEnvelopeMessages]
type analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsAggregateCurrentGetResponseEnvelopeSuccess bool

const (
	AnalyticsAggregateCurrentGetResponseEnvelopeSuccessTrue AnalyticsAggregateCurrentGetResponseEnvelopeSuccess = true
)
