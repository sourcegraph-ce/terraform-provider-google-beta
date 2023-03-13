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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeMachineImageIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccComputeMachineImageIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeMachineImageIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccComputeMachineImageIamPolicy_emptyBinding(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_withConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_withAndWithoutConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamMember_withConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamMember_withAndWithoutConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_compute_machine_image_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
		},
	})
}

func testAccComputeMachineImageIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeMachineImageIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeMachineImageIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeMachineImageIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeMachineImageIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeMachineImageIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeMachineImageIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_machine_image_iam_binding" "foo2" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_machine_image_iam_binding" "foo3" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeMachineImageIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeMachineImageIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_machine_image_iam_member" "foo2" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_machine_image_iam_member" "foo3" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeMachineImageIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
