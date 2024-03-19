// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// NetflowTopService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewNetflowTopService] method instead.
type NetflowTopService struct {
	Options []option.RequestOption
}

// NewNetflowTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetflowTopService(opts ...option.RequestOption) (r *NetflowTopService) {
	r = &NetflowTopService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by network traffic (NetFlows) over a given
// time period. Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *NetflowTopService) Ases(ctx context.Context, query NetflowTopAsesParams, opts ...option.RequestOption) (res *NetflowTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetflowTopAsesResponseEnvelope
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by network traffic (NetFlows) over a given time period.
// Visit https://en.wikipedia.org/wiki/NetFlow for more information.
func (r *NetflowTopService) Locations(ctx context.Context, query NetflowTopLocationsParams, opts ...option.RequestOption) (res *NetflowTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetflowTopLocationsResponseEnvelope
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetflowTopAsesResponse struct {
	Top0 []NetflowTopAsesResponseTop0 `json:"top_0,required"`
	JSON netflowTopAsesResponseJSON   `json:"-"`
}

// netflowTopAsesResponseJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponse]
type netflowTopAsesResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesResponseTop0 struct {
	ClientASN    float64                        `json:"clientASN,required"`
	ClientAsName string                         `json:"clientASName,required"`
	Value        string                         `json:"value,required"`
	JSON         netflowTopAsesResponseTop0JSON `json:"-"`
}

// netflowTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseTop0]
type netflowTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NetflowTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponse struct {
	Top0 []NetflowTopLocationsResponseTop0 `json:"top_0,required"`
	JSON netflowTopLocationsResponseJSON   `json:"-"`
}

// netflowTopLocationsResponseJSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponse]
type netflowTopLocationsResponseJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                              `json:"clientCountryName,required"`
	Value               string                              `json:"value,required"`
	JSON                netflowTopLocationsResponseTop0JSON `json:"-"`
}

// netflowTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponseTop0]
type netflowTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]NetflowTopAsesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[NetflowTopAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopAsesParams]'s query parameters as `url.Values`.
func (r NetflowTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetflowTopAsesParamsDateRange string

const (
	NetflowTopAsesParamsDateRange1d         NetflowTopAsesParamsDateRange = "1d"
	NetflowTopAsesParamsDateRange2d         NetflowTopAsesParamsDateRange = "2d"
	NetflowTopAsesParamsDateRange7d         NetflowTopAsesParamsDateRange = "7d"
	NetflowTopAsesParamsDateRange14d        NetflowTopAsesParamsDateRange = "14d"
	NetflowTopAsesParamsDateRange28d        NetflowTopAsesParamsDateRange = "28d"
	NetflowTopAsesParamsDateRange12w        NetflowTopAsesParamsDateRange = "12w"
	NetflowTopAsesParamsDateRange24w        NetflowTopAsesParamsDateRange = "24w"
	NetflowTopAsesParamsDateRange52w        NetflowTopAsesParamsDateRange = "52w"
	NetflowTopAsesParamsDateRange1dControl  NetflowTopAsesParamsDateRange = "1dControl"
	NetflowTopAsesParamsDateRange2dControl  NetflowTopAsesParamsDateRange = "2dControl"
	NetflowTopAsesParamsDateRange7dControl  NetflowTopAsesParamsDateRange = "7dControl"
	NetflowTopAsesParamsDateRange14dControl NetflowTopAsesParamsDateRange = "14dControl"
	NetflowTopAsesParamsDateRange28dControl NetflowTopAsesParamsDateRange = "28dControl"
	NetflowTopAsesParamsDateRange12wControl NetflowTopAsesParamsDateRange = "12wControl"
	NetflowTopAsesParamsDateRange24wControl NetflowTopAsesParamsDateRange = "24wControl"
)

// Format results are returned in.
type NetflowTopAsesParamsFormat string

const (
	NetflowTopAsesParamsFormatJson NetflowTopAsesParamsFormat = "JSON"
	NetflowTopAsesParamsFormatCsv  NetflowTopAsesParamsFormat = "CSV"
)

type NetflowTopAsesResponseEnvelope struct {
	Result  NetflowTopAsesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    netflowTopAsesResponseEnvelopeJSON `json:"-"`
}

// netflowTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseEnvelope]
type netflowTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]NetflowTopLocationsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[NetflowTopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopLocationsParams]'s query parameters as
// `url.Values`.
func (r NetflowTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetflowTopLocationsParamsDateRange string

const (
	NetflowTopLocationsParamsDateRange1d         NetflowTopLocationsParamsDateRange = "1d"
	NetflowTopLocationsParamsDateRange2d         NetflowTopLocationsParamsDateRange = "2d"
	NetflowTopLocationsParamsDateRange7d         NetflowTopLocationsParamsDateRange = "7d"
	NetflowTopLocationsParamsDateRange14d        NetflowTopLocationsParamsDateRange = "14d"
	NetflowTopLocationsParamsDateRange28d        NetflowTopLocationsParamsDateRange = "28d"
	NetflowTopLocationsParamsDateRange12w        NetflowTopLocationsParamsDateRange = "12w"
	NetflowTopLocationsParamsDateRange24w        NetflowTopLocationsParamsDateRange = "24w"
	NetflowTopLocationsParamsDateRange52w        NetflowTopLocationsParamsDateRange = "52w"
	NetflowTopLocationsParamsDateRange1dControl  NetflowTopLocationsParamsDateRange = "1dControl"
	NetflowTopLocationsParamsDateRange2dControl  NetflowTopLocationsParamsDateRange = "2dControl"
	NetflowTopLocationsParamsDateRange7dControl  NetflowTopLocationsParamsDateRange = "7dControl"
	NetflowTopLocationsParamsDateRange14dControl NetflowTopLocationsParamsDateRange = "14dControl"
	NetflowTopLocationsParamsDateRange28dControl NetflowTopLocationsParamsDateRange = "28dControl"
	NetflowTopLocationsParamsDateRange12wControl NetflowTopLocationsParamsDateRange = "12wControl"
	NetflowTopLocationsParamsDateRange24wControl NetflowTopLocationsParamsDateRange = "24wControl"
)

// Format results are returned in.
type NetflowTopLocationsParamsFormat string

const (
	NetflowTopLocationsParamsFormatJson NetflowTopLocationsParamsFormat = "JSON"
	NetflowTopLocationsParamsFormatCsv  NetflowTopLocationsParamsFormat = "CSV"
)

type NetflowTopLocationsResponseEnvelope struct {
	Result  NetflowTopLocationsResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    netflowTopLocationsResponseEnvelopeJSON `json:"-"`
}

// netflowTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetflowTopLocationsResponseEnvelope]
type netflowTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
