// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDevicePostureNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Posture.New(context.TODO(), zero_trust.DevicePostureNewParams{
		AccountID:   cloudflare.F("699d98642c564d2e855e9661899b7252"),
		Name:        cloudflare.F("Admin Serial Numbers"),
		Type:        cloudflare.F(zero_trust.DevicePostureNewParamsTypeFile),
		Description: cloudflare.F("The rule for admin serial numbers"),
		Expiration:  cloudflare.F("1h"),
		Input: cloudflare.F[zero_trust.DevicePostureNewParamsInputUnion](zero_trust.DevicePostureNewParamsInputTeamsDevicesFileInputRequest{
			Exists:          cloudflare.F(true),
			OperatingSystem: cloudflare.F(zero_trust.UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux),
			Path:            cloudflare.F("/bin/cat"),
			Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
			Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
		}),
		Match: cloudflare.F([]zero_trust.DevicePostureNewParamsMatch{{
			Platform: cloudflare.F(zero_trust.DevicePostureNewParamsMatchPlatformWindows),
		}, {
			Platform: cloudflare.F(zero_trust.DevicePostureNewParamsMatchPlatformWindows),
		}, {
			Platform: cloudflare.F(zero_trust.DevicePostureNewParamsMatchPlatformWindows),
		}}),
		Schedule: cloudflare.F("1h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Posture.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureUpdateParams{
			AccountID:   cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(zero_trust.DevicePostureUpdateParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[zero_trust.DevicePostureUpdateParamsInputUnion](zero_trust.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(zero_trust.UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			}),
			Match: cloudflare.F([]zero_trust.DevicePostureUpdateParamsMatch{{
				Platform: cloudflare.F(zero_trust.DevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(zero_trust.DevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(zero_trust.DevicePostureUpdateParamsMatchPlatformWindows),
			}}),
			Schedule: cloudflare.F("1h"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Posture.List(context.TODO(), zero_trust.DevicePostureListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Posture.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureDeleteParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Body:      cloudflare.F[any](map[string]interface{}{}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Posture.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureGetParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
