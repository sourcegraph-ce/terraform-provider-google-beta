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
	log "github.com/sourcegraph-ce/logrus"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMonitoringService() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringServiceCreate,
		Read:   resourceMonitoringServiceRead,
		Update: resourceMonitoringServiceUpdate,
		Delete: resourceMonitoringServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name used for UI elements listing this Service.`,
			},
			"service_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z0-9\-]+$`),
				Description: `An optional service ID to use. If not given, the server will generate a
service ID.`,
			},
			"telemetry": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for how to query telemetry on a Service.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_name": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The full name of the resource that defines this service.
Formatted as described in
https://cloud.google.com/apis/design/resource_names.`,
						},
					},
				},
			},
			"user_labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels which have been used to annotate the service. Label keys must start
with a letter. Label keys and values may contain lowercase letters,
numbers, underscores, and dashes. Label keys and values have a maximum
length of 63 characters, and must be less than 128 bytes in size. Up to 64
label entries may be stored. For labels which do not have a semantic value,
the empty string may be supplied for the label value.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The full resource name for this service. The syntax is:
projects/[PROJECT_ID]/services/[SERVICE_ID].`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceMonitoringServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	userLabelsProp, err := expandMonitoringServiceUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); ok || !reflect.DeepEqual(v, userLabelsProp) {
		obj["userLabels"] = userLabelsProp
	}
	telemetryProp, err := expandMonitoringServiceTelemetry(d.Get("telemetry"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("telemetry"); !isEmptyValue(reflect.ValueOf(telemetryProp)) && (ok || !reflect.DeepEqual(v, telemetryProp)) {
		obj["telemetry"] = telemetryProp
	}
	nameProp, err := expandMonitoringServiceServiceId(d.Get("service_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_id"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceMonitoringServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/services?serviceId={{service_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Service: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Service: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isMonitoringConcurrentEditError)
	if err != nil {
		return fmt.Errorf("Error creating Service: %s", err)
	}
	if err := d.Set("name", flattenMonitoringServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Service %q: %#v", d.Id(), res)

	return resourceMonitoringServiceRead(d, meta)
}

func resourceMonitoringServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Service: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil, isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringService %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	if err := d.Set("name", flattenMonitoringServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringServiceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("user_labels", flattenMonitoringServiceUserLabels(res["userLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("telemetry", flattenMonitoringServiceTelemetry(res["telemetry"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("service_id", flattenMonitoringServiceServiceId(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	return nil
}

func resourceMonitoringServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Service: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	userLabelsProp, err := expandMonitoringServiceUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); ok || !reflect.DeepEqual(v, userLabelsProp) {
		obj["userLabels"] = userLabelsProp
	}
	telemetryProp, err := expandMonitoringServiceTelemetry(d.Get("telemetry"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("telemetry"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, telemetryProp)) {
		obj["telemetry"] = telemetryProp
	}

	obj, err = resourceMonitoringServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Service %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("user_labels") {
		updateMask = append(updateMask, "userLabels")
	}

	if d.HasChange("telemetry") {
		updateMask = append(updateMask, "telemetry")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate), isMonitoringConcurrentEditError)

	if err != nil {
		return fmt.Errorf("Error updating Service %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Service %q: %#v", d.Id(), res)
	}

	return resourceMonitoringServiceRead(d, meta)
}

func resourceMonitoringServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Service: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Service %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, "Service")
	}

	log.Printf("[DEBUG] Finished deleting Service %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringServiceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceUserLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceTelemetry(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_name"] =
		flattenMonitoringServiceTelemetryResourceName(original["resourceName"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringServiceTelemetryResourceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceServiceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandMonitoringServiceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceUserLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringServiceTelemetry(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceName, err := expandMonitoringServiceTelemetryResourceName(original["resource_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceName); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceName"] = transformedResourceName
	}

	return transformed, nil
}

func expandMonitoringServiceTelemetryResourceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceServiceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceMonitoringServiceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Currently only CUSTOM service types can be created, but the
	// custom identifier block does not actually have fields right now.
	// Set to empty to indicate manually-created service type is CUSTOM.
	if _, ok := obj["custom"]; !ok {
		obj["custom"] = map[string]interface{}{}
	}
	// Name/Service ID is a query parameter only
	delete(obj, "name")

	return obj, nil
}
