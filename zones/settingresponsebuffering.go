// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingResponseBufferingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingResponseBufferingService] method instead.
type SettingResponseBufferingService struct {
	Options []option.RequestOption
}

// NewSettingResponseBufferingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingResponseBufferingService(opts ...option.RequestOption) (r *SettingResponseBufferingService) {
	r = &SettingResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Edit(ctx context.Context, params SettingResponseBufferingEditParams, opts ...option.RequestOption) (res *ZoneSettingBuffering, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Get(ctx context.Context, query SettingResponseBufferingGetParams, opts ...option.RequestOption) (res *ZoneSettingBuffering, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBufferingJSON `json:"-"`
}

// zoneSettingBufferingJSON contains the JSON metadata for the struct
// [ZoneSettingBuffering]
type zoneSettingBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingBuffering) implementsZonesSettingEditResponse() {}

func (r ZoneSettingBuffering) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingBufferingID string

const (
	ZoneSettingBufferingIDResponseBuffering ZoneSettingBufferingID = "response_buffering"
)

func (r ZoneSettingBufferingID) IsKnown() bool {
	switch r {
	case ZoneSettingBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingBufferingValue string

const (
	ZoneSettingBufferingValueOn  ZoneSettingBufferingValue = "on"
	ZoneSettingBufferingValueOff ZoneSettingBufferingValue = "off"
)

func (r ZoneSettingBufferingValue) IsKnown() bool {
	switch r {
	case ZoneSettingBufferingValueOn, ZoneSettingBufferingValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBufferingEditable bool

const (
	ZoneSettingBufferingEditableTrue  ZoneSettingBufferingEditable = true
	ZoneSettingBufferingEditableFalse ZoneSettingBufferingEditable = false
)

func (r ZoneSettingBufferingEditable) IsKnown() bool {
	switch r {
	case ZoneSettingBufferingEditableTrue, ZoneSettingBufferingEditableFalse:
		return true
	}
	return false
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingBufferingParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingBufferingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingBufferingValue] `json:"value,required"`
}

func (r ZoneSettingBufferingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingBufferingParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingResponseBufferingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingResponseBufferingEditParamsValue] `json:"value,required"`
}

func (r SettingResponseBufferingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingResponseBufferingEditParamsValue string

const (
	SettingResponseBufferingEditParamsValueOn  SettingResponseBufferingEditParamsValue = "on"
	SettingResponseBufferingEditParamsValueOff SettingResponseBufferingEditParamsValue = "off"
)

func (r SettingResponseBufferingEditParamsValue) IsKnown() bool {
	switch r {
	case SettingResponseBufferingEditParamsValueOn, SettingResponseBufferingEditParamsValueOff:
		return true
	}
	return false
}

type SettingResponseBufferingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingBuffering                             `json:"result"`
	JSON   settingResponseBufferingEditResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingResponseBufferingEditResponseEnvelope]
type settingResponseBufferingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingResponseBufferingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingBuffering                            `json:"result"`
	JSON   settingResponseBufferingGetResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingResponseBufferingGetResponseEnvelope]
type settingResponseBufferingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
