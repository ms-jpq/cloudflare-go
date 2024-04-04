// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureRules], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/posture", query.AccountID)
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

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) ListAutoPaging(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureRules] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, params DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureRules struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureRulesInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureRulesMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureRulesType `json:"type"`
	JSON devicePostureRulesJSON `json:"-"`
}

// devicePostureRulesJSON contains the JSON metadata for the struct
// [DevicePostureRules]
type devicePostureRulesJSON struct {
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

func (r *DevicePostureRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesJSON) RawJSON() string {
	return r.raw
}

// The value to be checked against.
type DevicePostureRulesInput struct {
	// Whether or not file exists
	Exists bool `json:"exists"`
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system"`
	// File path.
	Path string `json:"path"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string `json:"thumbprint"`
	// List ID.
	ID string `json:"id"`
	// Domain
	Domain string `json:"domain"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string `json:"os_version_extra"`
	// Version of OS
	Version string `json:"version"`
	// Enabled
	Enabled    bool        `json:"enabled"`
	CheckDisks interface{} `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn string `json:"cn"`
	// Compliance Status
	ComplianceStatus DevicePostureRulesInputComplianceStatus `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID string `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureRulesInputState `json:"state"`
	// Version Operator
	VersionOperator DevicePostureRulesInputVersionOperator `json:"versionOperator"`
	// Count Operator
	CountOperator DevicePostureRulesInputCountOperator `json:"countOperator"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureRulesInputRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureRulesInputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureRulesInputNetworkStatus `json:"network_status"`
	JSON          devicePostureRulesInputJSON          `json:"-"`
	union         DevicePostureRulesInputUnion
}

// devicePostureRulesInputJSON contains the JSON metadata for the struct
// [DevicePostureRulesInput]
type devicePostureRulesInputJSON struct {
	Exists           apijson.Field
	OperatingSystem  apijson.Field
	Path             apijson.Field
	Sha256           apijson.Field
	Thumbprint       apijson.Field
	ID               apijson.Field
	Domain           apijson.Field
	Operator         apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	Version          apijson.Field
	Enabled          apijson.Field
	CheckDisks       apijson.Field
	RequireAll       apijson.Field
	CertificateID    apijson.Field
	Cn               apijson.Field
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	LastSeen         apijson.Field
	OS               apijson.Field
	Overall          apijson.Field
	SensorConfig     apijson.Field
	State            apijson.Field
	VersionOperator  apijson.Field
	CountOperator    apijson.Field
	IssueCount       apijson.Field
	EidLastSeen      apijson.Field
	RiskLevel        apijson.Field
	ScoreOperator    apijson.Field
	TotalScore       apijson.Field
	ActiveThreats    apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	NetworkStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r devicePostureRulesInputJSON) RawJSON() string {
	return r.raw
}

func (r *DevicePostureRulesInput) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DevicePostureRulesInput) AsUnion() DevicePostureRulesInputUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by
// [zero_trust.DevicePostureRulesInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesTaniumInputRequest] or
// [zero_trust.DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureRulesInputUnion interface {
	implementsZeroTrustDevicePostureRulesInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DevicePostureRulesInputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesOSVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest{}),
		},
	)
}

type DevicePostureRulesInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                  `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesFileInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesFileInputRequest]
type devicePostureRulesInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                          `json:"domain"`
	JSON   devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                       `json:"os_version_extra"`
	JSON           devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesOSVersionInputRequest]
type devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesFirewallInputRequest]
type devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesSentineloneInputRequest]
type devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest]
type devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                              `json:"requireAll"`
	JSON       devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesApplicationInputRequest]
type devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                               `json:"cn,required"`
	JSON devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest]
type devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Compliance Status
type DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                    `json:"connection_id,required"`
	JSON         devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesIntuneInputRequest]
type devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Compliance Status
type DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                    `json:"issue_count,required"`
	JSON       devicePostureRulesInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesKolideInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesKolideInputRequest]
type devicePostureRulesInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Count Operator
type DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                   `json:"total_score"`
	JSON       devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesTaniumInputRequest]
type devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930                  `json:"operator"`
	JSON     devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Network status of device.
type DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Compliance Status
type DevicePostureRulesInputComplianceStatus string

const (
	DevicePostureRulesInputComplianceStatusCompliant     DevicePostureRulesInputComplianceStatus = "compliant"
	DevicePostureRulesInputComplianceStatusNoncompliant  DevicePostureRulesInputComplianceStatus = "noncompliant"
	DevicePostureRulesInputComplianceStatusUnknown       DevicePostureRulesInputComplianceStatus = "unknown"
	DevicePostureRulesInputComplianceStatusNotapplicable DevicePostureRulesInputComplianceStatus = "notapplicable"
	DevicePostureRulesInputComplianceStatusIngraceperiod DevicePostureRulesInputComplianceStatus = "ingraceperiod"
	DevicePostureRulesInputComplianceStatusError         DevicePostureRulesInputComplianceStatus = "error"
)

func (r DevicePostureRulesInputComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputComplianceStatusCompliant, DevicePostureRulesInputComplianceStatusNoncompliant, DevicePostureRulesInputComplianceStatusUnknown, DevicePostureRulesInputComplianceStatusNotapplicable, DevicePostureRulesInputComplianceStatusIngraceperiod, DevicePostureRulesInputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureRulesInputState string

const (
	DevicePostureRulesInputStateOnline  DevicePostureRulesInputState = "online"
	DevicePostureRulesInputStateOffline DevicePostureRulesInputState = "offline"
	DevicePostureRulesInputStateUnknown DevicePostureRulesInputState = "unknown"
)

func (r DevicePostureRulesInputState) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputStateOnline, DevicePostureRulesInputStateOffline, DevicePostureRulesInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureRulesInputVersionOperator string

const (
	DevicePostureRulesInputVersionOperatorLess            DevicePostureRulesInputVersionOperator = "<"
	DevicePostureRulesInputVersionOperatorLessOrEquals    DevicePostureRulesInputVersionOperator = "<="
	DevicePostureRulesInputVersionOperatorGreater         DevicePostureRulesInputVersionOperator = ">"
	DevicePostureRulesInputVersionOperatorGreaterOrEquals DevicePostureRulesInputVersionOperator = ">="
	DevicePostureRulesInputVersionOperatorEquals          DevicePostureRulesInputVersionOperator = "=="
)

func (r DevicePostureRulesInputVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputVersionOperatorLess, DevicePostureRulesInputVersionOperatorLessOrEquals, DevicePostureRulesInputVersionOperatorGreater, DevicePostureRulesInputVersionOperatorGreaterOrEquals, DevicePostureRulesInputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type DevicePostureRulesInputCountOperator string

const (
	DevicePostureRulesInputCountOperatorLess            DevicePostureRulesInputCountOperator = "<"
	DevicePostureRulesInputCountOperatorLessOrEquals    DevicePostureRulesInputCountOperator = "<="
	DevicePostureRulesInputCountOperatorGreater         DevicePostureRulesInputCountOperator = ">"
	DevicePostureRulesInputCountOperatorGreaterOrEquals DevicePostureRulesInputCountOperator = ">="
	DevicePostureRulesInputCountOperatorEquals          DevicePostureRulesInputCountOperator = "=="
)

func (r DevicePostureRulesInputCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputCountOperatorLess, DevicePostureRulesInputCountOperatorLessOrEquals, DevicePostureRulesInputCountOperatorGreater, DevicePostureRulesInputCountOperatorGreaterOrEquals, DevicePostureRulesInputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureRulesInputRiskLevel string

const (
	DevicePostureRulesInputRiskLevelLow      DevicePostureRulesInputRiskLevel = "low"
	DevicePostureRulesInputRiskLevelMedium   DevicePostureRulesInputRiskLevel = "medium"
	DevicePostureRulesInputRiskLevelHigh     DevicePostureRulesInputRiskLevel = "high"
	DevicePostureRulesInputRiskLevelCritical DevicePostureRulesInputRiskLevel = "critical"
)

func (r DevicePostureRulesInputRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputRiskLevelLow, DevicePostureRulesInputRiskLevelMedium, DevicePostureRulesInputRiskLevelHigh, DevicePostureRulesInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureRulesInputScoreOperator string

const (
	DevicePostureRulesInputScoreOperatorLess            DevicePostureRulesInputScoreOperator = "<"
	DevicePostureRulesInputScoreOperatorLessOrEquals    DevicePostureRulesInputScoreOperator = "<="
	DevicePostureRulesInputScoreOperatorGreater         DevicePostureRulesInputScoreOperator = ">"
	DevicePostureRulesInputScoreOperatorGreaterOrEquals DevicePostureRulesInputScoreOperator = ">="
	DevicePostureRulesInputScoreOperatorEquals          DevicePostureRulesInputScoreOperator = "=="
)

func (r DevicePostureRulesInputScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputScoreOperatorLess, DevicePostureRulesInputScoreOperatorLessOrEquals, DevicePostureRulesInputScoreOperatorGreater, DevicePostureRulesInputScoreOperatorGreaterOrEquals, DevicePostureRulesInputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DevicePostureRulesInputNetworkStatus string

const (
	DevicePostureRulesInputNetworkStatusConnected     DevicePostureRulesInputNetworkStatus = "connected"
	DevicePostureRulesInputNetworkStatusDisconnected  DevicePostureRulesInputNetworkStatus = "disconnected"
	DevicePostureRulesInputNetworkStatusDisconnecting DevicePostureRulesInputNetworkStatus = "disconnecting"
	DevicePostureRulesInputNetworkStatusConnecting    DevicePostureRulesInputNetworkStatus = "connecting"
)

func (r DevicePostureRulesInputNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputNetworkStatusConnected, DevicePostureRulesInputNetworkStatusDisconnected, DevicePostureRulesInputNetworkStatusDisconnecting, DevicePostureRulesInputNetworkStatusConnecting:
		return true
	}
	return false
}

type DevicePostureRulesMatch struct {
	Platform DevicePostureRulesMatchPlatform `json:"platform"`
	JSON     devicePostureRulesMatchJSON     `json:"-"`
}

// devicePostureRulesMatchJSON contains the JSON metadata for the struct
// [DevicePostureRulesMatch]
type devicePostureRulesMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRulesMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesMatchJSON) RawJSON() string {
	return r.raw
}

type DevicePostureRulesMatchPlatform string

const (
	DevicePostureRulesMatchPlatformWindows DevicePostureRulesMatchPlatform = "windows"
	DevicePostureRulesMatchPlatformMac     DevicePostureRulesMatchPlatform = "mac"
	DevicePostureRulesMatchPlatformLinux   DevicePostureRulesMatchPlatform = "linux"
	DevicePostureRulesMatchPlatformAndroid DevicePostureRulesMatchPlatform = "android"
	DevicePostureRulesMatchPlatformIos     DevicePostureRulesMatchPlatform = "ios"
)

func (r DevicePostureRulesMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureRulesMatchPlatformWindows, DevicePostureRulesMatchPlatformMac, DevicePostureRulesMatchPlatformLinux, DevicePostureRulesMatchPlatformAndroid, DevicePostureRulesMatchPlatformIos:
		return true
	}
	return false
}

// The type of device posture rule.
type DevicePostureRulesType string

const (
	DevicePostureRulesTypeFile              DevicePostureRulesType = "file"
	DevicePostureRulesTypeApplication       DevicePostureRulesType = "application"
	DevicePostureRulesTypeTanium            DevicePostureRulesType = "tanium"
	DevicePostureRulesTypeGateway           DevicePostureRulesType = "gateway"
	DevicePostureRulesTypeWARP              DevicePostureRulesType = "warp"
	DevicePostureRulesTypeDiskEncryption    DevicePostureRulesType = "disk_encryption"
	DevicePostureRulesTypeSentinelone       DevicePostureRulesType = "sentinelone"
	DevicePostureRulesTypeCarbonblack       DevicePostureRulesType = "carbonblack"
	DevicePostureRulesTypeFirewall          DevicePostureRulesType = "firewall"
	DevicePostureRulesTypeOSVersion         DevicePostureRulesType = "os_version"
	DevicePostureRulesTypeDomainJoined      DevicePostureRulesType = "domain_joined"
	DevicePostureRulesTypeClientCertificate DevicePostureRulesType = "client_certificate"
	DevicePostureRulesTypeUniqueClientID    DevicePostureRulesType = "unique_client_id"
	DevicePostureRulesTypeKolide            DevicePostureRulesType = "kolide"
	DevicePostureRulesTypeTaniumS2s         DevicePostureRulesType = "tanium_s2s"
	DevicePostureRulesTypeCrowdstrikeS2s    DevicePostureRulesType = "crowdstrike_s2s"
	DevicePostureRulesTypeIntune            DevicePostureRulesType = "intune"
	DevicePostureRulesTypeWorkspaceOne      DevicePostureRulesType = "workspace_one"
	DevicePostureRulesTypeSentineloneS2s    DevicePostureRulesType = "sentinelone_s2s"
)

func (r DevicePostureRulesType) IsKnown() bool {
	switch r {
	case DevicePostureRulesTypeFile, DevicePostureRulesTypeApplication, DevicePostureRulesTypeTanium, DevicePostureRulesTypeGateway, DevicePostureRulesTypeWARP, DevicePostureRulesTypeDiskEncryption, DevicePostureRulesTypeSentinelone, DevicePostureRulesTypeCarbonblack, DevicePostureRulesTypeFirewall, DevicePostureRulesTypeOSVersion, DevicePostureRulesTypeDomainJoined, DevicePostureRulesTypeClientCertificate, DevicePostureRulesTypeUniqueClientID, DevicePostureRulesTypeKolide, DevicePostureRulesTypeTaniumS2s, DevicePostureRulesTypeCrowdstrikeS2s, DevicePostureRulesTypeIntune, DevicePostureRulesTypeWorkspaceOne, DevicePostureRulesTypeSentineloneS2s:
		return true
	}
	return false
}

// operator
type UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 string

const (
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less            UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals    UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater         UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals          UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "=="
)

func (r UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals:
		return true
	}
	return false
}

// Operating system
type UnnamedSchemaRef41885dd46b9e0294254c49305a273681 string

const (
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "windows"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux   UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "linux"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac     UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "mac"
)

func (r UnnamedSchemaRef41885dd46b9e0294254c49305a273681) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac:
		return true
	}
	return false
}

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
	AccountID param.Field[string] `path:"account_id,required"`
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
	Input param.Field[DevicePostureNewParamsInputUnion] `json:"input"`
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

func (r DevicePostureNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsTypeFile, DevicePostureNewParamsTypeApplication, DevicePostureNewParamsTypeTanium, DevicePostureNewParamsTypeGateway, DevicePostureNewParamsTypeWARP, DevicePostureNewParamsTypeDiskEncryption, DevicePostureNewParamsTypeSentinelone, DevicePostureNewParamsTypeCarbonblack, DevicePostureNewParamsTypeFirewall, DevicePostureNewParamsTypeOSVersion, DevicePostureNewParamsTypeDomainJoined, DevicePostureNewParamsTypeClientCertificate, DevicePostureNewParamsTypeUniqueClientID, DevicePostureNewParamsTypeKolide, DevicePostureNewParamsTypeTaniumS2s, DevicePostureNewParamsTypeCrowdstrikeS2s, DevicePostureNewParamsTypeIntune, DevicePostureNewParamsTypeWorkspaceOne, DevicePostureNewParamsTypeSentineloneS2s:
		return true
	}
	return false
}

// The value to be checked against.
type DevicePostureNewParamsInput struct {
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system"`
	// File path.
	Path param.Field[string] `json:"path"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// List ID.
	ID param.Field[string] `json:"id"`
	// Domain
	Domain param.Field[string] `json:"domain"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
	// Version of OS
	Version param.Field[string] `json:"version"`
	// Enabled
	Enabled    param.Field[bool]        `json:"enabled"`
	CheckDisks param.Field[interface{}] `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn"`
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureNewParamsInputState] `json:"state"`
	// Version Operator
	VersionOperator param.Field[DevicePostureNewParamsInputVersionOperator] `json:"versionOperator"`
	// Count Operator
	CountOperator param.Field[DevicePostureNewParamsInputCountOperator] `json:"countOperator"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureNewParamsInputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureNewParamsInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureNewParamsInputNetworkStatus] `json:"network_status"`
}

func (r DevicePostureNewParamsInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInput) implementsZeroTrustDevicePostureNewParamsInputUnion() {}

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
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest],
// [DevicePostureNewParamsInput].
type DevicePostureNewParamsInputUnion interface {
	implementsZeroTrustDevicePostureNewParamsInputUnion()
}

type DevicePostureNewParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator,required"`
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

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

type DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

type DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
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

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
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

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
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

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

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

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Count Operator
type DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

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

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

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
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureNewParamsInputUnion() {
}

// Network status of device.
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Compliance Status
type DevicePostureNewParamsInputComplianceStatus string

const (
	DevicePostureNewParamsInputComplianceStatusCompliant     DevicePostureNewParamsInputComplianceStatus = "compliant"
	DevicePostureNewParamsInputComplianceStatusNoncompliant  DevicePostureNewParamsInputComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputComplianceStatusUnknown       DevicePostureNewParamsInputComplianceStatus = "unknown"
	DevicePostureNewParamsInputComplianceStatusNotapplicable DevicePostureNewParamsInputComplianceStatus = "notapplicable"
	DevicePostureNewParamsInputComplianceStatusIngraceperiod DevicePostureNewParamsInputComplianceStatus = "ingraceperiod"
	DevicePostureNewParamsInputComplianceStatusError         DevicePostureNewParamsInputComplianceStatus = "error"
)

func (r DevicePostureNewParamsInputComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputComplianceStatusCompliant, DevicePostureNewParamsInputComplianceStatusNoncompliant, DevicePostureNewParamsInputComplianceStatusUnknown, DevicePostureNewParamsInputComplianceStatusNotapplicable, DevicePostureNewParamsInputComplianceStatusIngraceperiod, DevicePostureNewParamsInputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewParamsInputState string

const (
	DevicePostureNewParamsInputStateOnline  DevicePostureNewParamsInputState = "online"
	DevicePostureNewParamsInputStateOffline DevicePostureNewParamsInputState = "offline"
	DevicePostureNewParamsInputStateUnknown DevicePostureNewParamsInputState = "unknown"
)

func (r DevicePostureNewParamsInputState) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputStateOnline, DevicePostureNewParamsInputStateOffline, DevicePostureNewParamsInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureNewParamsInputVersionOperator string

const (
	DevicePostureNewParamsInputVersionOperatorLess            DevicePostureNewParamsInputVersionOperator = "<"
	DevicePostureNewParamsInputVersionOperatorLessOrEquals    DevicePostureNewParamsInputVersionOperator = "<="
	DevicePostureNewParamsInputVersionOperatorGreater         DevicePostureNewParamsInputVersionOperator = ">"
	DevicePostureNewParamsInputVersionOperatorGreaterOrEquals DevicePostureNewParamsInputVersionOperator = ">="
	DevicePostureNewParamsInputVersionOperatorEquals          DevicePostureNewParamsInputVersionOperator = "=="
)

func (r DevicePostureNewParamsInputVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputVersionOperatorLess, DevicePostureNewParamsInputVersionOperatorLessOrEquals, DevicePostureNewParamsInputVersionOperatorGreater, DevicePostureNewParamsInputVersionOperatorGreaterOrEquals, DevicePostureNewParamsInputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type DevicePostureNewParamsInputCountOperator string

const (
	DevicePostureNewParamsInputCountOperatorLess            DevicePostureNewParamsInputCountOperator = "<"
	DevicePostureNewParamsInputCountOperatorLessOrEquals    DevicePostureNewParamsInputCountOperator = "<="
	DevicePostureNewParamsInputCountOperatorGreater         DevicePostureNewParamsInputCountOperator = ">"
	DevicePostureNewParamsInputCountOperatorGreaterOrEquals DevicePostureNewParamsInputCountOperator = ">="
	DevicePostureNewParamsInputCountOperatorEquals          DevicePostureNewParamsInputCountOperator = "=="
)

func (r DevicePostureNewParamsInputCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputCountOperatorLess, DevicePostureNewParamsInputCountOperatorLessOrEquals, DevicePostureNewParamsInputCountOperatorGreater, DevicePostureNewParamsInputCountOperatorGreaterOrEquals, DevicePostureNewParamsInputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewParamsInputRiskLevel string

const (
	DevicePostureNewParamsInputRiskLevelLow      DevicePostureNewParamsInputRiskLevel = "low"
	DevicePostureNewParamsInputRiskLevelMedium   DevicePostureNewParamsInputRiskLevel = "medium"
	DevicePostureNewParamsInputRiskLevelHigh     DevicePostureNewParamsInputRiskLevel = "high"
	DevicePostureNewParamsInputRiskLevelCritical DevicePostureNewParamsInputRiskLevel = "critical"
)

func (r DevicePostureNewParamsInputRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputRiskLevelLow, DevicePostureNewParamsInputRiskLevelMedium, DevicePostureNewParamsInputRiskLevelHigh, DevicePostureNewParamsInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureNewParamsInputScoreOperator string

const (
	DevicePostureNewParamsInputScoreOperatorLess            DevicePostureNewParamsInputScoreOperator = "<"
	DevicePostureNewParamsInputScoreOperatorLessOrEquals    DevicePostureNewParamsInputScoreOperator = "<="
	DevicePostureNewParamsInputScoreOperatorGreater         DevicePostureNewParamsInputScoreOperator = ">"
	DevicePostureNewParamsInputScoreOperatorGreaterOrEquals DevicePostureNewParamsInputScoreOperator = ">="
	DevicePostureNewParamsInputScoreOperatorEquals          DevicePostureNewParamsInputScoreOperator = "=="
)

func (r DevicePostureNewParamsInputScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputScoreOperatorLess, DevicePostureNewParamsInputScoreOperatorLessOrEquals, DevicePostureNewParamsInputScoreOperatorGreater, DevicePostureNewParamsInputScoreOperatorGreaterOrEquals, DevicePostureNewParamsInputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DevicePostureNewParamsInputNetworkStatus string

const (
	DevicePostureNewParamsInputNetworkStatusConnected     DevicePostureNewParamsInputNetworkStatus = "connected"
	DevicePostureNewParamsInputNetworkStatusDisconnected  DevicePostureNewParamsInputNetworkStatus = "disconnected"
	DevicePostureNewParamsInputNetworkStatusDisconnecting DevicePostureNewParamsInputNetworkStatus = "disconnecting"
	DevicePostureNewParamsInputNetworkStatusConnecting    DevicePostureNewParamsInputNetworkStatus = "connecting"
)

func (r DevicePostureNewParamsInputNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputNetworkStatusConnected, DevicePostureNewParamsInputNetworkStatusDisconnected, DevicePostureNewParamsInputNetworkStatusDisconnecting, DevicePostureNewParamsInputNetworkStatusConnecting:
		return true
	}
	return false
}

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

func (r DevicePostureNewParamsMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsMatchPlatformWindows, DevicePostureNewParamsMatchPlatformMac, DevicePostureNewParamsMatchPlatformLinux, DevicePostureNewParamsMatchPlatformAndroid, DevicePostureNewParamsMatchPlatformIos:
		return true
	}
	return false
}

type DevicePostureNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRules                                        `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

func (r DevicePostureNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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
	Input param.Field[DevicePostureUpdateParamsInputUnion] `json:"input"`
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

func (r DevicePostureUpdateParamsType) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsTypeFile, DevicePostureUpdateParamsTypeApplication, DevicePostureUpdateParamsTypeTanium, DevicePostureUpdateParamsTypeGateway, DevicePostureUpdateParamsTypeWARP, DevicePostureUpdateParamsTypeDiskEncryption, DevicePostureUpdateParamsTypeSentinelone, DevicePostureUpdateParamsTypeCarbonblack, DevicePostureUpdateParamsTypeFirewall, DevicePostureUpdateParamsTypeOSVersion, DevicePostureUpdateParamsTypeDomainJoined, DevicePostureUpdateParamsTypeClientCertificate, DevicePostureUpdateParamsTypeUniqueClientID, DevicePostureUpdateParamsTypeKolide, DevicePostureUpdateParamsTypeTaniumS2s, DevicePostureUpdateParamsTypeCrowdstrikeS2s, DevicePostureUpdateParamsTypeIntune, DevicePostureUpdateParamsTypeWorkspaceOne, DevicePostureUpdateParamsTypeSentineloneS2s:
		return true
	}
	return false
}

// The value to be checked against.
type DevicePostureUpdateParamsInput struct {
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system"`
	// File path.
	Path param.Field[string] `json:"path"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// List ID.
	ID param.Field[string] `json:"id"`
	// Domain
	Domain param.Field[string] `json:"domain"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
	// Version of OS
	Version param.Field[string] `json:"version"`
	// Enabled
	Enabled    param.Field[bool]        `json:"enabled"`
	CheckDisks param.Field[interface{}] `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn"`
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureUpdateParamsInputState] `json:"state"`
	// Version Operator
	VersionOperator param.Field[DevicePostureUpdateParamsInputVersionOperator] `json:"versionOperator"`
	// Count Operator
	CountOperator param.Field[DevicePostureUpdateParamsInputCountOperator] `json:"countOperator"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureUpdateParamsInputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureUpdateParamsInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureUpdateParamsInputNetworkStatus] `json:"network_status"`
}

func (r DevicePostureUpdateParamsInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInput) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {}

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
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest],
// [DevicePostureUpdateParamsInput].
type DevicePostureUpdateParamsInputUnion interface {
	implementsZeroTrustDevicePostureUpdateParamsInputUnion()
}

type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

type DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

type DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
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

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
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

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

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

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Count Operator
type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

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

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

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
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureUpdateParamsInputUnion() {
}

// Network status of device.
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Compliance Status
type DevicePostureUpdateParamsInputComplianceStatus string

const (
	DevicePostureUpdateParamsInputComplianceStatusCompliant     DevicePostureUpdateParamsInputComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputComplianceStatusNoncompliant  DevicePostureUpdateParamsInputComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputComplianceStatusUnknown       DevicePostureUpdateParamsInputComplianceStatus = "unknown"
	DevicePostureUpdateParamsInputComplianceStatusNotapplicable DevicePostureUpdateParamsInputComplianceStatus = "notapplicable"
	DevicePostureUpdateParamsInputComplianceStatusIngraceperiod DevicePostureUpdateParamsInputComplianceStatus = "ingraceperiod"
	DevicePostureUpdateParamsInputComplianceStatusError         DevicePostureUpdateParamsInputComplianceStatus = "error"
)

func (r DevicePostureUpdateParamsInputComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputComplianceStatusCompliant, DevicePostureUpdateParamsInputComplianceStatusNoncompliant, DevicePostureUpdateParamsInputComplianceStatusUnknown, DevicePostureUpdateParamsInputComplianceStatusNotapplicable, DevicePostureUpdateParamsInputComplianceStatusIngraceperiod, DevicePostureUpdateParamsInputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateParamsInputState string

const (
	DevicePostureUpdateParamsInputStateOnline  DevicePostureUpdateParamsInputState = "online"
	DevicePostureUpdateParamsInputStateOffline DevicePostureUpdateParamsInputState = "offline"
	DevicePostureUpdateParamsInputStateUnknown DevicePostureUpdateParamsInputState = "unknown"
)

func (r DevicePostureUpdateParamsInputState) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputStateOnline, DevicePostureUpdateParamsInputStateOffline, DevicePostureUpdateParamsInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureUpdateParamsInputVersionOperator string

const (
	DevicePostureUpdateParamsInputVersionOperatorLess            DevicePostureUpdateParamsInputVersionOperator = "<"
	DevicePostureUpdateParamsInputVersionOperatorLessOrEquals    DevicePostureUpdateParamsInputVersionOperator = "<="
	DevicePostureUpdateParamsInputVersionOperatorGreater         DevicePostureUpdateParamsInputVersionOperator = ">"
	DevicePostureUpdateParamsInputVersionOperatorGreaterOrEquals DevicePostureUpdateParamsInputVersionOperator = ">="
	DevicePostureUpdateParamsInputVersionOperatorEquals          DevicePostureUpdateParamsInputVersionOperator = "=="
)

func (r DevicePostureUpdateParamsInputVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputVersionOperatorLess, DevicePostureUpdateParamsInputVersionOperatorLessOrEquals, DevicePostureUpdateParamsInputVersionOperatorGreater, DevicePostureUpdateParamsInputVersionOperatorGreaterOrEquals, DevicePostureUpdateParamsInputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type DevicePostureUpdateParamsInputCountOperator string

const (
	DevicePostureUpdateParamsInputCountOperatorLess            DevicePostureUpdateParamsInputCountOperator = "<"
	DevicePostureUpdateParamsInputCountOperatorLessOrEquals    DevicePostureUpdateParamsInputCountOperator = "<="
	DevicePostureUpdateParamsInputCountOperatorGreater         DevicePostureUpdateParamsInputCountOperator = ">"
	DevicePostureUpdateParamsInputCountOperatorGreaterOrEquals DevicePostureUpdateParamsInputCountOperator = ">="
	DevicePostureUpdateParamsInputCountOperatorEquals          DevicePostureUpdateParamsInputCountOperator = "=="
)

func (r DevicePostureUpdateParamsInputCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputCountOperatorLess, DevicePostureUpdateParamsInputCountOperatorLessOrEquals, DevicePostureUpdateParamsInputCountOperatorGreater, DevicePostureUpdateParamsInputCountOperatorGreaterOrEquals, DevicePostureUpdateParamsInputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateParamsInputRiskLevel string

const (
	DevicePostureUpdateParamsInputRiskLevelLow      DevicePostureUpdateParamsInputRiskLevel = "low"
	DevicePostureUpdateParamsInputRiskLevelMedium   DevicePostureUpdateParamsInputRiskLevel = "medium"
	DevicePostureUpdateParamsInputRiskLevelHigh     DevicePostureUpdateParamsInputRiskLevel = "high"
	DevicePostureUpdateParamsInputRiskLevelCritical DevicePostureUpdateParamsInputRiskLevel = "critical"
)

func (r DevicePostureUpdateParamsInputRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputRiskLevelLow, DevicePostureUpdateParamsInputRiskLevelMedium, DevicePostureUpdateParamsInputRiskLevelHigh, DevicePostureUpdateParamsInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureUpdateParamsInputScoreOperator string

const (
	DevicePostureUpdateParamsInputScoreOperatorLess            DevicePostureUpdateParamsInputScoreOperator = "<"
	DevicePostureUpdateParamsInputScoreOperatorLessOrEquals    DevicePostureUpdateParamsInputScoreOperator = "<="
	DevicePostureUpdateParamsInputScoreOperatorGreater         DevicePostureUpdateParamsInputScoreOperator = ">"
	DevicePostureUpdateParamsInputScoreOperatorGreaterOrEquals DevicePostureUpdateParamsInputScoreOperator = ">="
	DevicePostureUpdateParamsInputScoreOperatorEquals          DevicePostureUpdateParamsInputScoreOperator = "=="
)

func (r DevicePostureUpdateParamsInputScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputScoreOperatorLess, DevicePostureUpdateParamsInputScoreOperatorLessOrEquals, DevicePostureUpdateParamsInputScoreOperatorGreater, DevicePostureUpdateParamsInputScoreOperatorGreaterOrEquals, DevicePostureUpdateParamsInputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DevicePostureUpdateParamsInputNetworkStatus string

const (
	DevicePostureUpdateParamsInputNetworkStatusConnected     DevicePostureUpdateParamsInputNetworkStatus = "connected"
	DevicePostureUpdateParamsInputNetworkStatusDisconnected  DevicePostureUpdateParamsInputNetworkStatus = "disconnected"
	DevicePostureUpdateParamsInputNetworkStatusDisconnecting DevicePostureUpdateParamsInputNetworkStatus = "disconnecting"
	DevicePostureUpdateParamsInputNetworkStatusConnecting    DevicePostureUpdateParamsInputNetworkStatus = "connecting"
)

func (r DevicePostureUpdateParamsInputNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputNetworkStatusConnected, DevicePostureUpdateParamsInputNetworkStatusDisconnected, DevicePostureUpdateParamsInputNetworkStatusDisconnecting, DevicePostureUpdateParamsInputNetworkStatusConnecting:
		return true
	}
	return false
}

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

func (r DevicePostureUpdateParamsMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsMatchPlatformWindows, DevicePostureUpdateParamsMatchPlatformMac, DevicePostureUpdateParamsMatchPlatformLinux, DevicePostureUpdateParamsMatchPlatformAndroid, DevicePostureUpdateParamsMatchPlatformIos:
		return true
	}
	return false
}

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRules                                        `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

func (r DevicePostureUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r DevicePostureDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureDeleteResponse                               `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRules                                        `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)

func (r DevicePostureGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
