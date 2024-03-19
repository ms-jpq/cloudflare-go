// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RateLimitService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRateLimitService] method instead.
type RateLimitService struct {
	Options []option.RequestOption
}

// NewRateLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRateLimitService(opts ...option.RequestOption) (r *RateLimitService) {
	r = &RateLimitService{}
	r.Options = opts
	return
}

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
func (r *RateLimitService) New(ctx context.Context, zoneIdentifier string, body RateLimitNewParams, opts ...option.RequestOption) (res *RateLimitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the rate limits for a zone.
func (r *RateLimitService) List(ctx context.Context, zoneIdentifier string, query RateLimitListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RateLimitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
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

// Fetches the rate limits for a zone.
func (r *RateLimitService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RateLimitListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RateLimitListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing rate limit.
func (r *RateLimitService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RateLimitDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rate limit.
func (r *RateLimitService) Edit(ctx context.Context, zoneIdentifier string, id string, body RateLimitEditParams, opts ...option.RequestOption) (res *RateLimitEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a rate limit.
func (r *RateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RateLimitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [rate_limits.RateLimitNewResponseUnknown] or
// [shared.UnionString].
type RateLimitNewResponse interface {
	ImplementsRateLimitsRateLimitNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitListResponse struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RateLimitListResponseAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RateLimitListResponseBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RateLimitListResponseMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                   `json:"threshold"`
	JSON      rateLimitListResponseJSON `json:"-"`
}

// rateLimitListResponseJSON contains the JSON metadata for the struct
// [RateLimitListResponse]
type rateLimitListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitListResponseAction struct {
	// The action to perform.
	Mode RateLimitListResponseActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RateLimitListResponseActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                         `json:"timeout"`
	JSON    rateLimitListResponseActionJSON `json:"-"`
}

// rateLimitListResponseActionJSON contains the JSON metadata for the struct
// [RateLimitListResponseAction]
type rateLimitListResponseActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseActionJSON) RawJSON() string {
	return r.raw
}

// The action to perform.
type RateLimitListResponseActionMode string

const (
	RateLimitListResponseActionModeSimulate         RateLimitListResponseActionMode = "simulate"
	RateLimitListResponseActionModeBan              RateLimitListResponseActionMode = "ban"
	RateLimitListResponseActionModeChallenge        RateLimitListResponseActionMode = "challenge"
	RateLimitListResponseActionModeJsChallenge      RateLimitListResponseActionMode = "js_challenge"
	RateLimitListResponseActionModeManagedChallenge RateLimitListResponseActionMode = "managed_challenge"
)

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitListResponseActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                  `json:"content_type"`
	JSON        rateLimitListResponseActionResponseJSON `json:"-"`
}

// rateLimitListResponseActionResponseJSON contains the JSON metadata for the
// struct [RateLimitListResponseActionResponse]
type rateLimitListResponseActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseActionResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitListResponseBypass struct {
	Name RateLimitListResponseBypassName `json:"name"`
	// The URL to bypass.
	Value string                          `json:"value"`
	JSON  rateLimitListResponseBypassJSON `json:"-"`
}

// rateLimitListResponseBypassJSON contains the JSON metadata for the struct
// [RateLimitListResponseBypass]
type rateLimitListResponseBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseBypassJSON) RawJSON() string {
	return r.raw
}

type RateLimitListResponseBypassName string

const (
	RateLimitListResponseBypassNameURL RateLimitListResponseBypassName = "url"
)

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitListResponseMatch struct {
	Headers  []RateLimitListResponseMatchHeader `json:"headers"`
	Request  RateLimitListResponseMatchRequest  `json:"request"`
	Response RateLimitListResponseMatchResponse `json:"response"`
	JSON     rateLimitListResponseMatchJSON     `json:"-"`
}

// rateLimitListResponseMatchJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatch]
type rateLimitListResponseMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseMatchJSON) RawJSON() string {
	return r.raw
}

type RateLimitListResponseMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RateLimitListResponseMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                               `json:"value"`
	JSON  rateLimitListResponseMatchHeaderJSON `json:"-"`
}

// rateLimitListResponseMatchHeaderJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchHeader]
type rateLimitListResponseMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseMatchHeaderJSON) RawJSON() string {
	return r.raw
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitListResponseMatchHeadersOp string

const (
	RateLimitListResponseMatchHeadersOpEq RateLimitListResponseMatchHeadersOp = "eq"
	RateLimitListResponseMatchHeadersOpNe RateLimitListResponseMatchHeadersOp = "ne"
)

type RateLimitListResponseMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []RateLimitListResponseMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                `json:"url"`
	JSON rateLimitListResponseMatchRequestJSON `json:"-"`
}

// rateLimitListResponseMatchRequestJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchRequest]
type rateLimitListResponseMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseMatchRequestJSON) RawJSON() string {
	return r.raw
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitListResponseMatchRequestMethod string

const (
	RateLimitListResponseMatchRequestMethodGet    RateLimitListResponseMatchRequestMethod = "GET"
	RateLimitListResponseMatchRequestMethodPost   RateLimitListResponseMatchRequestMethod = "POST"
	RateLimitListResponseMatchRequestMethodPut    RateLimitListResponseMatchRequestMethod = "PUT"
	RateLimitListResponseMatchRequestMethodDelete RateLimitListResponseMatchRequestMethod = "DELETE"
	RateLimitListResponseMatchRequestMethodPatch  RateLimitListResponseMatchRequestMethod = "PATCH"
	RateLimitListResponseMatchRequestMethodHead   RateLimitListResponseMatchRequestMethod = "HEAD"
	RateLimitListResponseMatchRequestMethod_All   RateLimitListResponseMatchRequestMethod = "_ALL_"
)

type RateLimitListResponseMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                   `json:"origin_traffic"`
	JSON          rateLimitListResponseMatchResponseJSON `json:"-"`
}

// rateLimitListResponseMatchResponseJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchResponse]
type rateLimitListResponseMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RateLimitListResponseMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitListResponseMatchResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponse struct {
	// The unique identifier of the rate limit.
	ID   string                      `json:"id"`
	JSON rateLimitDeleteResponseJSON `json:"-"`
}

// rateLimitDeleteResponseJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponse]
type rateLimitDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [rate_limits.RateLimitEditResponseUnknown] or
// [shared.UnionString].
type RateLimitEditResponse interface {
	ImplementsRateLimitsRateLimitEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [rate_limits.RateLimitGetResponseUnknown] or
// [shared.UnionString].
type RateLimitGetResponse interface {
	ImplementsRateLimitsRateLimitGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RateLimitNewResponseEnvelope struct {
	Errors   []RateLimitNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitNewResponseEnvelopeJSON    `json:"-"`
}

// rateLimitNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitNewResponseEnvelope]
type rateLimitNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RateLimitNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    rateLimitNewResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RateLimitNewResponseEnvelopeErrors]
type rateLimitNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RateLimitNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rateLimitNewResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitNewResponseEnvelopeMessages]
type rateLimitNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitNewResponseEnvelopeSuccess bool

const (
	RateLimitNewResponseEnvelopeSuccessTrue RateLimitNewResponseEnvelopeSuccess = true
)

type RateLimitListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RateLimitListParams]'s query parameters as `url.Values`.
func (r RateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateLimitDeleteResponseEnvelope struct {
	Errors   []RateLimitDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitDeleteResponseEnvelopeJSON    `json:"-"`
}

// rateLimitDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseEnvelope]
type rateLimitDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rateLimitDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RateLimitDeleteResponseEnvelopeErrors]
type rateLimitDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    rateLimitDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitDeleteResponseEnvelopeMessages]
type rateLimitDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitDeleteResponseEnvelopeSuccess bool

const (
	RateLimitDeleteResponseEnvelopeSuccessTrue RateLimitDeleteResponseEnvelopeSuccess = true
)

type RateLimitEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RateLimitEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RateLimitEditResponseEnvelope struct {
	Errors   []RateLimitEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitEditResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitEditResponseEnvelopeJSON    `json:"-"`
}

// rateLimitEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitEditResponseEnvelope]
type rateLimitEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RateLimitEditResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    rateLimitEditResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RateLimitEditResponseEnvelopeErrors]
type rateLimitEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RateLimitEditResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rateLimitEditResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitEditResponseEnvelopeMessages]
type rateLimitEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitEditResponseEnvelopeSuccess bool

const (
	RateLimitEditResponseEnvelopeSuccessTrue RateLimitEditResponseEnvelopeSuccess = true
)

type RateLimitGetResponseEnvelope struct {
	Errors   []RateLimitGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitGetResponseEnvelopeJSON    `json:"-"`
}

// rateLimitGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelope]
type rateLimitGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RateLimitGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    rateLimitGetResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelopeErrors]
type rateLimitGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RateLimitGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rateLimitGetResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitGetResponseEnvelopeMessages]
type rateLimitGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitGetResponseEnvelopeSuccess bool

const (
	RateLimitGetResponseEnvelopeSuccessTrue RateLimitGetResponseEnvelopeSuccess = true
)
