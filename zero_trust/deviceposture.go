// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePostureService] method
// instead.
type DevicePostureService struct {
	Options      []option.RequestOption
	Integrations *DevicePostureIntegrationService
}

// NewDevicePostureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePostureService(opts ...option.RequestOption) (r *DevicePostureService) {
	r = &DevicePostureService{}
	r.Options = opts
	r.Integrations = NewDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *[]TeamsDevicesDevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, body DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamsDevicesDevicePostureRules struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input TeamsDevicesDevicePostureRulesInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []TeamsDevicesDevicePostureRulesMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type TeamsDevicesDevicePostureRulesType `json:"type"`
	JSON teamsDevicesDevicePostureRulesJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesJSON contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRules]
type teamsDevicesDevicePostureRulesJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesJSON) RawJSON() string {
	return r.raw
}

// The value to be checked against.
//
// Union satisfied by
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest],
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest]
// or
// [zero_trust.TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest].
type TeamsDevicesDevicePostureRulesInput interface {
	implementsZeroTrustTeamsDevicesDevicePostureRulesInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamsDevicesDevicePostureRulesInput)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest{}),
		},
	)
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                              `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON contains the
// JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                      `json:"domain"`
	JSON   teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                   `json:"os_version_extra"`
	JSON           teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown5 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown6 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown7 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown8 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown9 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                          `json:"requireAll"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                           `json:"cn,required"`
	JSON teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Compliance Status
type TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
	ConnectionID    apijson.Field
	LastSeen        apijson.Field
	Operator        apijson.Field
	OS              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	State           apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown15 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown16 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown17 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown18 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown19 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline  TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown25 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown26 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown27 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown28 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown29 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Compliance Status
type TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError         TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                `json:"issue_count,required"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Count Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown35 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown36 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown37 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown38 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown39 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                               `json:"total_score"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown45 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown46 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown47 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown48 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown49 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow      TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium   TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh     TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown55 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown56 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown57 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown58 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown59 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustTeamsDevicesDevicePostureRulesInput() {
}

// Network status of device.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown65 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown66 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown67 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown68 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown69 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type TeamsDevicesDevicePostureRulesMatch struct {
	Platform TeamsDevicesDevicePostureRulesMatchPlatform `json:"platform"`
	JSON     teamsDevicesDevicePostureRulesMatchJSON     `json:"-"`
}

// teamsDevicesDevicePostureRulesMatchJSON contains the JSON metadata for the
// struct [TeamsDevicesDevicePostureRulesMatch]
type teamsDevicesDevicePostureRulesMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicePostureRulesMatchJSON) RawJSON() string {
	return r.raw
}

type TeamsDevicesDevicePostureRulesMatchPlatform string

const (
	TeamsDevicesDevicePostureRulesMatchPlatformWindows TeamsDevicesDevicePostureRulesMatchPlatform = "windows"
	TeamsDevicesDevicePostureRulesMatchPlatformMac     TeamsDevicesDevicePostureRulesMatchPlatform = "mac"
	TeamsDevicesDevicePostureRulesMatchPlatformLinux   TeamsDevicesDevicePostureRulesMatchPlatform = "linux"
	TeamsDevicesDevicePostureRulesMatchPlatformAndroid TeamsDevicesDevicePostureRulesMatchPlatform = "android"
	TeamsDevicesDevicePostureRulesMatchPlatformIos     TeamsDevicesDevicePostureRulesMatchPlatform = "ios"
)

// The type of device posture rule.
type TeamsDevicesDevicePostureRulesType string

const (
	TeamsDevicesDevicePostureRulesTypeFile              TeamsDevicesDevicePostureRulesType = "file"
	TeamsDevicesDevicePostureRulesTypeApplication       TeamsDevicesDevicePostureRulesType = "application"
	TeamsDevicesDevicePostureRulesTypeTanium            TeamsDevicesDevicePostureRulesType = "tanium"
	TeamsDevicesDevicePostureRulesTypeGateway           TeamsDevicesDevicePostureRulesType = "gateway"
	TeamsDevicesDevicePostureRulesTypeWARP              TeamsDevicesDevicePostureRulesType = "warp"
	TeamsDevicesDevicePostureRulesTypeDiskEncryption    TeamsDevicesDevicePostureRulesType = "disk_encryption"
	TeamsDevicesDevicePostureRulesTypeSentinelone       TeamsDevicesDevicePostureRulesType = "sentinelone"
	TeamsDevicesDevicePostureRulesTypeCarbonblack       TeamsDevicesDevicePostureRulesType = "carbonblack"
	TeamsDevicesDevicePostureRulesTypeFirewall          TeamsDevicesDevicePostureRulesType = "firewall"
	TeamsDevicesDevicePostureRulesTypeOSVersion         TeamsDevicesDevicePostureRulesType = "os_version"
	TeamsDevicesDevicePostureRulesTypeDomainJoined      TeamsDevicesDevicePostureRulesType = "domain_joined"
	TeamsDevicesDevicePostureRulesTypeClientCertificate TeamsDevicesDevicePostureRulesType = "client_certificate"
	TeamsDevicesDevicePostureRulesTypeUniqueClientID    TeamsDevicesDevicePostureRulesType = "unique_client_id"
	TeamsDevicesDevicePostureRulesTypeKolide            TeamsDevicesDevicePostureRulesType = "kolide"
	TeamsDevicesDevicePostureRulesTypeTaniumS2s         TeamsDevicesDevicePostureRulesType = "tanium_s2s"
	TeamsDevicesDevicePostureRulesTypeCrowdstrikeS2s    TeamsDevicesDevicePostureRulesType = "crowdstrike_s2s"
	TeamsDevicesDevicePostureRulesTypeIntune            TeamsDevicesDevicePostureRulesType = "intune"
	TeamsDevicesDevicePostureRulesTypeWorkspaceOne      TeamsDevicesDevicePostureRulesType = "workspace_one"
	TeamsDevicesDevicePostureRulesTypeSentineloneS2s    TeamsDevicesDevicePostureRulesType = "sentinelone_s2s"
)

type DevicePostureDeleteResponse struct {
	// API UUID.
	ID   string                          `json:"id"`
	JSON devicePostureDeleteResponseJSON `json:"-"`
}

// devicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [DevicePostureDeleteResponse]
type devicePostureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureNewParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureNewParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureNewParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureNewParamsType string

const (
	DevicePostureNewParamsTypeFile              DevicePostureNewParamsType = "file"
	DevicePostureNewParamsTypeApplication       DevicePostureNewParamsType = "application"
	DevicePostureNewParamsTypeTanium            DevicePostureNewParamsType = "tanium"
	DevicePostureNewParamsTypeGateway           DevicePostureNewParamsType = "gateway"
	DevicePostureNewParamsTypeWARP              DevicePostureNewParamsType = "warp"
	DevicePostureNewParamsTypeDiskEncryption    DevicePostureNewParamsType = "disk_encryption"
	DevicePostureNewParamsTypeSentinelone       DevicePostureNewParamsType = "sentinelone"
	DevicePostureNewParamsTypeCarbonblack       DevicePostureNewParamsType = "carbonblack"
	DevicePostureNewParamsTypeFirewall          DevicePostureNewParamsType = "firewall"
	DevicePostureNewParamsTypeOSVersion         DevicePostureNewParamsType = "os_version"
	DevicePostureNewParamsTypeDomainJoined      DevicePostureNewParamsType = "domain_joined"
	DevicePostureNewParamsTypeClientCertificate DevicePostureNewParamsType = "client_certificate"
	DevicePostureNewParamsTypeUniqueClientID    DevicePostureNewParamsType = "unique_client_id"
	DevicePostureNewParamsTypeKolide            DevicePostureNewParamsType = "kolide"
	DevicePostureNewParamsTypeTaniumS2s         DevicePostureNewParamsType = "tanium_s2s"
	DevicePostureNewParamsTypeCrowdstrikeS2s    DevicePostureNewParamsType = "crowdstrike_s2s"
	DevicePostureNewParamsTypeIntune            DevicePostureNewParamsType = "intune"
	DevicePostureNewParamsTypeWorkspaceOne      DevicePostureNewParamsType = "workspace_one"
	DevicePostureNewParamsTypeSentineloneS2s    DevicePostureNewParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureNewParamsInput interface {
	implementsZeroTrustDevicePostureNewParamsInput()
}

type DevicePostureNewParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown75 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown76 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown77 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown78 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown79 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureNewParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Count Operator
type DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown115 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown116 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown117 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown118 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown119 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Network status of device.
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureNewParamsMatch struct {
	Platform param.Field[DevicePostureNewParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureNewParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureNewParamsMatchPlatform string

const (
	DevicePostureNewParamsMatchPlatformWindows DevicePostureNewParamsMatchPlatform = "windows"
	DevicePostureNewParamsMatchPlatformMac     DevicePostureNewParamsMatchPlatform = "mac"
	DevicePostureNewParamsMatchPlatformLinux   DevicePostureNewParamsMatchPlatform = "linux"
	DevicePostureNewParamsMatchPlatformAndroid DevicePostureNewParamsMatchPlatform = "android"
	DevicePostureNewParamsMatchPlatformIos     DevicePostureNewParamsMatchPlatform = "ios"
)

type DevicePostureNewResponseEnvelope struct {
	Errors   []DevicePostureNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDevicePostureRules             `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureNewResponseEnvelope]
type devicePostureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePostureNewResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureNewResponseEnvelopeErrors]
type devicePostureNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePostureNewResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureNewResponseEnvelopeMessages]
type devicePostureNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

type DevicePostureUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureUpdateParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureUpdateParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureUpdateParamsType string

const (
	DevicePostureUpdateParamsTypeFile              DevicePostureUpdateParamsType = "file"
	DevicePostureUpdateParamsTypeApplication       DevicePostureUpdateParamsType = "application"
	DevicePostureUpdateParamsTypeTanium            DevicePostureUpdateParamsType = "tanium"
	DevicePostureUpdateParamsTypeGateway           DevicePostureUpdateParamsType = "gateway"
	DevicePostureUpdateParamsTypeWARP              DevicePostureUpdateParamsType = "warp"
	DevicePostureUpdateParamsTypeDiskEncryption    DevicePostureUpdateParamsType = "disk_encryption"
	DevicePostureUpdateParamsTypeSentinelone       DevicePostureUpdateParamsType = "sentinelone"
	DevicePostureUpdateParamsTypeCarbonblack       DevicePostureUpdateParamsType = "carbonblack"
	DevicePostureUpdateParamsTypeFirewall          DevicePostureUpdateParamsType = "firewall"
	DevicePostureUpdateParamsTypeOSVersion         DevicePostureUpdateParamsType = "os_version"
	DevicePostureUpdateParamsTypeDomainJoined      DevicePostureUpdateParamsType = "domain_joined"
	DevicePostureUpdateParamsTypeClientCertificate DevicePostureUpdateParamsType = "client_certificate"
	DevicePostureUpdateParamsTypeUniqueClientID    DevicePostureUpdateParamsType = "unique_client_id"
	DevicePostureUpdateParamsTypeKolide            DevicePostureUpdateParamsType = "kolide"
	DevicePostureUpdateParamsTypeTaniumS2s         DevicePostureUpdateParamsType = "tanium_s2s"
	DevicePostureUpdateParamsTypeCrowdstrikeS2s    DevicePostureUpdateParamsType = "crowdstrike_s2s"
	DevicePostureUpdateParamsTypeIntune            DevicePostureUpdateParamsType = "intune"
	DevicePostureUpdateParamsTypeWorkspaceOne      DevicePostureUpdateParamsType = "workspace_one"
	DevicePostureUpdateParamsTypeSentineloneS2s    DevicePostureUpdateParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureUpdateParamsInput interface {
	implementsZeroTrustDevicePostureUpdateParamsInput()
}

type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown145 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown146 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown147 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown148 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown149 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Count Operator
type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown185 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown186 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown187 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown188 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown189 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Network status of device.
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureUpdateParamsMatch struct {
	Platform param.Field[DevicePostureUpdateParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureUpdateParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureUpdateParamsMatchPlatform string

const (
	DevicePostureUpdateParamsMatchPlatformWindows DevicePostureUpdateParamsMatchPlatform = "windows"
	DevicePostureUpdateParamsMatchPlatformMac     DevicePostureUpdateParamsMatchPlatform = "mac"
	DevicePostureUpdateParamsMatchPlatformLinux   DevicePostureUpdateParamsMatchPlatform = "linux"
	DevicePostureUpdateParamsMatchPlatformAndroid DevicePostureUpdateParamsMatchPlatform = "android"
	DevicePostureUpdateParamsMatchPlatformIos     DevicePostureUpdateParamsMatchPlatform = "ios"
)

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []DevicePostureUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDevicePostureRules                `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePostureUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelope]
type devicePostureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelopeErrors]
type devicePostureUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    devicePostureUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePostureUpdateResponseEnvelopeMessages]
type devicePostureUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

type DevicePostureListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePostureListResponseEnvelope struct {
	Errors   []DevicePostureListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamsDevicesDevicePostureRules            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePostureListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePostureListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePostureListResponseEnvelopeJSON       `json:"-"`
}

// devicePostureListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureListResponseEnvelope]
type devicePostureListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    devicePostureListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureListResponseEnvelopeErrors]
type devicePostureListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureListResponseEnvelopeMessages]
type devicePostureListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureListResponseEnvelopeSuccess bool

const (
	DevicePostureListResponseEnvelopeSuccessTrue DevicePostureListResponseEnvelopeSuccess = true
)

type DevicePostureListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       devicePostureListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePostureListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePostureListResponseEnvelopeResultInfo]
type devicePostureListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePostureDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []DevicePostureDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelope]
type devicePostureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelopeErrors]
type devicePostureDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    devicePostureDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePostureDeleteResponseEnvelopeMessages]
type devicePostureDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

type DevicePostureGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePostureGetResponseEnvelope struct {
	Errors   []DevicePostureGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDevicePostureRules             `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureGetResponseEnvelope]
type devicePostureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePostureGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureGetResponseEnvelopeErrors]
type devicePostureGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePostureGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureGetResponseEnvelopeMessages]
type devicePostureGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)
