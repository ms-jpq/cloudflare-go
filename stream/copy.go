// File generated from our OpenAPI spec by Stainless.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CopyService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCopyService] method instead.
type CopyService struct {
	Options []option.RequestOption
}

// NewCopyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCopyService(opts ...option.RequestOption) (r *CopyService) {
	r = &CopyService{}
	r.Options = opts
	return
}

// Uploads a video to Stream from a provided URL.
func (r *CopyService) New(ctx context.Context, params CopyNewParams, opts ...option.RequestOption) (res *StreamVideos, err error) {
	opts = append(r.Options[:], opts...)
	var env CopyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/copy", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CopyNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A video's URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                `json:"thumbnailTimestampPct"`
	Watermark             param.Field[CopyNewParamsWatermark] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
	// Comma-separated key-value pairs following the TUS protocol specification. Values
	// are Base-64 encoded. Supported keys: `name`, `requiresignedurls`,
	// `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`.
	UploadMetadata param.Field[string] `header:"Upload-Metadata"`
}

func (r CopyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CopyNewParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r CopyNewParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CopyNewResponseEnvelope struct {
	Errors   []CopyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CopyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamVideos                      `json:"result,required"`
	// Whether the API call was successful
	Success CopyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    copyNewResponseEnvelopeJSON    `json:"-"`
}

// copyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CopyNewResponseEnvelope]
type copyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CopyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r copyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CopyNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    copyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// copyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CopyNewResponseEnvelopeErrors]
type copyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CopyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r copyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CopyNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    copyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// copyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CopyNewResponseEnvelopeMessages]
type copyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CopyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r copyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CopyNewResponseEnvelopeSuccess bool

const (
	CopyNewResponseEnvelopeSuccessTrue CopyNewResponseEnvelopeSuccess = true
)
