// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccVertexAITensorboard_vertexAiTensorboardExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckVertexAITensorboardDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAITensorboard_vertexAiTensorboardExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_tensorboard.tensorboard",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "project"},
			},
		},
	})
}

func testAccVertexAITensorboard_vertexAiTensorboardExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_tensorboard" "tensorboard" {
  display_name = "terraform%{random_suffix}"
  description  = "sample description"
  labels       = {
    "key1" : "value1",
    "key2" : "value2"
  }
  region       = "us-central1"
}
`, context)
}

func TestAccVertexAITensorboard_vertexAiTensorboardFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"kms_key_name":  BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckVertexAITensorboardDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAITensorboard_vertexAiTensorboardFullExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_tensorboard.tensorboard",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "project"},
			},
		},
	})
}

func testAccVertexAITensorboard_vertexAiTensorboardFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_tensorboard" "tensorboard" {
  display_name = "terraform%{random_suffix}"
  description  = "sample description"
  labels       = {
    "key1" : "value1",
    "key2" : "value2"
  }
  region       = "us-central1"
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  depends_on = [google_kms_crypto_key_iam_member.crypto_key]
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}
`, context)
}

func testAccCheckVertexAITensorboardDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_tensorboard" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{VertexAIBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("VertexAITensorboard still exists at %s", url)
			}
		}

		return nil
	}
}
