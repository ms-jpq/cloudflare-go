// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// ScriptBindingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptBindingService] method
// instead.
type ScriptBindingService struct {
	Options []option.RequestOption
}

// NewScriptBindingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptBindingService(opts ...option.RequestOption) (r *ScriptBindingService) {
	r = &ScriptBindingService{}
	r.Options = opts
	return
}

// List the bindings for a Workers script.
func (r *ScriptBindingService) Get(ctx context.Context, query ScriptBindingGetParams, opts ...option.RequestOption) (res *[]WorkersBinding, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptBindingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [workers.WorkersBindingWorkersKVNamespaceBinding] or
// [workers.WorkersBindingWorkersWasmModuleBinding].
type WorkersBinding interface {
	implementsWorkersWorkersBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkersBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkersBindingWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkersBindingWorkersWasmModuleBinding{}),
		},
	)
}

type WorkersBindingWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersBindingWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersBindingWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersBindingWorkersKVNamespaceBindingJSON contains the JSON metadata for the
// struct [WorkersBindingWorkersKVNamespaceBinding]
type workersBindingWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersBindingWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersBindingWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkersBindingWorkersKVNamespaceBinding) implementsWorkersWorkersBinding() {}

// The class of resource that the binding provides.
type WorkersBindingWorkersKVNamespaceBindingType string

const (
	WorkersBindingWorkersKVNamespaceBindingTypeKVNamespace WorkersBindingWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r WorkersBindingWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case WorkersBindingWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type WorkersBindingWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersBindingWorkersWasmModuleBindingType `json:"type,required"`
	JSON workersBindingWorkersWasmModuleBindingJSON `json:"-"`
}

// workersBindingWorkersWasmModuleBindingJSON contains the JSON metadata for the
// struct [WorkersBindingWorkersWasmModuleBinding]
type workersBindingWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersBindingWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersBindingWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkersBindingWorkersWasmModuleBinding) implementsWorkersWorkersBinding() {}

// The class of resource that the binding provides.
type WorkersBindingWorkersWasmModuleBindingType string

const (
	WorkersBindingWorkersWasmModuleBindingTypeWasmModule WorkersBindingWorkersWasmModuleBindingType = "wasm_module"
)

func (r WorkersBindingWorkersWasmModuleBindingType) IsKnown() bool {
	switch r {
	case WorkersBindingWorkersWasmModuleBindingTypeWasmModule:
		return true
	}
	return false
}

type ScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptBindingGetResponseEnvelope struct {
	Errors   []ScriptBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersBinding                           `json:"result,required"`
	// Whether the API call was successful
	Success ScriptBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptBindingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptBindingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptBindingGetResponseEnvelope]
type scriptBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptBindingGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    scriptBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptBindingGetResponseEnvelopeErrors]
type scriptBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptBindingGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    scriptBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptBindingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptBindingGetResponseEnvelopeMessages]
type scriptBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptBindingGetResponseEnvelopeSuccess bool

const (
	ScriptBindingGetResponseEnvelopeSuccessTrue ScriptBindingGetResponseEnvelopeSuccess = true
)

func (r ScriptBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
