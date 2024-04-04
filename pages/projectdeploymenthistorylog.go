// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

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

// ProjectDeploymentHistoryLogService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewProjectDeploymentHistoryLogService] method instead.
type ProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentHistoryLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *ProjectDeploymentHistoryLogService) {
	r = &ProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *ProjectDeploymentHistoryLogService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentHistoryLogGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentHistoryLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentHistoryLogGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentHistoryLogGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ProjectDeploymentHistoryLogGetResponseEnvelope]
type projectDeploymentHistoryLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
