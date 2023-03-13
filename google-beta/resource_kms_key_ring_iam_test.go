// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package google

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const DEFAULT_KMS_TEST_LOCATION = "us-central1"

func TestAccKmsKeyRingIamBinding(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyDecrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Binding creation
				Config: testAccKmsKeyRingIamBinding_basic(projectId, orgId, billingAccount, account, keyRingName, roleId),
				Check: testAccCheckGoogleKmsKeyRingIam(t, keyRingId.keyRingId(), roleId, []string{
					fmt.Sprintf("serviceAccount:%s@%s.iam.gserviceaccount.com", account, projectId),
				}),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s %s", keyRingId.terraformId(), roleId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccKmsKeyRingIamBinding_update(projectId, orgId, billingAccount, account, keyRingName, roleId),
				Check: testAccCheckGoogleKmsKeyRingIam(t, keyRingId.keyRingId(), roleId, []string{
					fmt.Sprintf("serviceAccount:%s@%s.iam.gserviceaccount.com", account, projectId),
					fmt.Sprintf("serviceAccount:%s-2@%s.iam.gserviceaccount.com", account, projectId),
				}),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s %s", keyRingId.terraformId(), roleId),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccKmsKeyRingIamBinding_withCondition(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyDecrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	conditionTitle := "expires_after_2019_12_31"

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKmsKeyRingIamBinding_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s %s %s", keyRingId.terraformId(), roleId, conditionTitle),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccKmsKeyRingIamMember(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyEncrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccKmsKeyRingIamMember_basic(projectId, orgId, billingAccount, account, keyRingName, roleId),
				Check: testAccCheckGoogleKmsKeyRingIam(t, keyRingId.keyRingId(), roleId, []string{
					fmt.Sprintf("serviceAccount:%s@%s.iam.gserviceaccount.com", account, projectId),
				}),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_member.foo",
				ImportStateId:     fmt.Sprintf("%s %s serviceAccount:%s@%s.iam.gserviceaccount.com", keyRingId.terraformId(), roleId, account, projectId),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccKmsKeyRingIamMember_withCondition(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyEncrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	conditionTitle := "expires_after_2019_12_31"

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKmsKeyRingIamMember_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_member.foo",
				ImportStateId:     fmt.Sprintf("%s %s serviceAccount:%s@%s.iam.gserviceaccount.com %s", keyRingId.terraformId(), roleId, account, projectId, conditionTitle),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccKmsKeyRingIamPolicy(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyEncrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKmsKeyRingIamPolicy_basic(projectId, orgId, billingAccount, account, keyRingName, roleId),
				Check: testAccCheckGoogleKmsKeyRingIam(t, keyRingId.keyRingId(), roleId, []string{
					fmt.Sprintf("serviceAccount:%s@%s.iam.gserviceaccount.com", account, projectId),
				}),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_policy.foo",
				ImportStateId:     keyRingId.terraformId(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccKmsKeyRingIamPolicy_withCondition(t *testing.T) {
	t.Parallel()

	orgId := GetTestOrgFromEnv(t)
	projectId := fmt.Sprintf("tf-test-%d", RandInt(t))
	billingAccount := GetTestBillingAccountFromEnv(t)
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	roleId := "roles/cloudkms.cryptoKeyEncrypter"
	keyRingName := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	conditionTitle := "expires_after_2019_12_31"

	keyRingId := &kmsKeyRingId{
		Project:  projectId,
		Location: DEFAULT_KMS_TEST_LOCATION,
		Name:     keyRingName,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKmsKeyRingIamPolicy_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle),
			},
			{
				ResourceName:      "google_kms_key_ring_iam_policy.foo",
				ImportStateId:     keyRingId.terraformId(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGoogleKmsKeyRingIam(t *testing.T, keyRingId, role string, members []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := GoogleProviderConfig(t)
		p, err := config.NewKmsClient(config.UserAgent).Projects.Locations.KeyRings.GetIamPolicy(keyRingId).Do()
		if err != nil {
			return err
		}

		for _, binding := range p.Bindings {
			if binding.Role == role {
				sort.Strings(members)
				sort.Strings(binding.Members)

				if reflect.DeepEqual(members, binding.Members) {
					return nil
				}

				return fmt.Errorf("Binding found but expected members is %v, got %v", members, binding.Members)
			}
		}

		return fmt.Errorf("No binding for role %q", role)
	}
}

// We are using a custom role since iam_binding is authoritative on the member list and
// we want to avoid removing members from an existing role to prevent unwanted side effects.
func testAccKmsKeyRingIamBinding_basic(projectId, orgId, billingAccount, account, keyRingName, roleId string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "us-central1"
  name     = "%s"
}

resource "google_kms_key_ring_iam_binding" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  role        = "%s"
  members     = ["serviceAccount:${google_service_account.test_account.email}"]
}
`, projectId, orgId, billingAccount, account, keyRingName, roleId)
}

func testAccKmsKeyRingIamBinding_update(projectId, orgId, billingAccount, account, keyRingName, roleId string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_service_account" "test_account_2" {
  project      = google_project_service.iam.project
  account_id   = "%s-2"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "%s"
  name     = "%s"
}

resource "google_kms_key_ring_iam_binding" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  role        = "%s"
  members = [
    "serviceAccount:${google_service_account.test_account.email}",
    "serviceAccount:${google_service_account.test_account_2.email}",
  ]
}
`, projectId, orgId, billingAccount, account, account, DEFAULT_KMS_TEST_LOCATION, keyRingName, roleId)
}

func testAccKmsKeyRingIamBinding_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "us-central1"
  name     = "%s"
}

resource "google_kms_key_ring_iam_binding" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  role        = "%s"
  members     = ["serviceAccount:${google_service_account.test_account.email}"]
  condition {
    title       = "%s"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle)
}

func testAccKmsKeyRingIamMember_basic(projectId, orgId, billingAccount, account, keyRingName, roleId string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "%s"
  name     = "%s"
}

resource "google_kms_key_ring_iam_member" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  role        = "%s"
  member      = "serviceAccount:${google_service_account.test_account.email}"
}
`, projectId, orgId, billingAccount, account, DEFAULT_KMS_TEST_LOCATION, keyRingName, roleId)
}

func testAccKmsKeyRingIamMember_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "%s"
  name     = "%s"
}

resource "google_kms_key_ring_iam_member" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  role        = "%s"
  member      = "serviceAccount:${google_service_account.test_account.email}"
  condition {
    title       = "%s"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, projectId, orgId, billingAccount, account, DEFAULT_KMS_TEST_LOCATION, keyRingName, roleId, conditionTitle)
}

func testAccKmsKeyRingIamPolicy_basic(projectId, orgId, billingAccount, account, keyRingName, roleId string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "%s"
  name     = "%s"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%s"

    members = ["serviceAccount:${google_service_account.test_account.email}"]
  }
}

resource "google_kms_key_ring_iam_policy" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, projectId, orgId, billingAccount, account, DEFAULT_KMS_TEST_LOCATION, keyRingName, roleId)
}

func testAccKmsKeyRingIamPolicy_withCondition(projectId, orgId, billingAccount, account, keyRingName, roleId, conditionTitle string) string {
	return fmt.Sprintf(`
resource "google_project" "test_project" {
  name            = "Test project"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "kms" {
  project = google_project.test_project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "iam" {
  project = google_project_service.kms.project
  service = "iam.googleapis.com"
}

resource "google_service_account" "test_account" {
  project      = google_project_service.iam.project
  account_id   = "%s"
  display_name = "Kms Key Ring Iam Testing Account"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.iam.project
  location = "%s"
  name     = "%s"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%s"

    members = ["serviceAccount:${google_service_account.test_account.email}"]
    condition {
      title       = "%s"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
    }
  }
}

resource "google_kms_key_ring_iam_policy" "foo" {
  key_ring_id = google_kms_key_ring.key_ring.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, projectId, orgId, billingAccount, account, DEFAULT_KMS_TEST_LOCATION, keyRingName, roleId, conditionTitle)
}
