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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePubsubLiteReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubLiteReservationCreate,
		Read:   resourcePubsubLiteReservationRead,
		Update: resourcePubsubLiteReservationUpdate,
		Delete: resourcePubsubLiteReservationDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubLiteReservationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Name of the reservation.`,
			},
			"throughput_capacity": {
				Type:     schema.TypeInt,
				Required: true,
				Description: `The reserved throughput capacity. Every unit of throughput capacity is
equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
messages.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The region of the pubsub lite reservation.`,
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

func resourcePubsubLiteReservationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	throughputCapacityProp, err := expandPubsubLiteReservationThroughputCapacity(d.Get("throughput_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("throughput_capacity"); !isEmptyValue(reflect.ValueOf(throughputCapacityProp)) && (ok || !reflect.DeepEqual(v, throughputCapacityProp)) {
		obj["throughputCapacity"] = throughputCapacityProp
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{region}}/reservations?reservationId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Reservation: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Reservation: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/reservations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Reservation %q: %#v", d.Id(), res)

	return resourcePubsubLiteReservationRead(d, meta)
}

func resourcePubsubLiteReservationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{region}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("PubsubLiteReservation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	if err := d.Set("throughput_capacity", flattenPubsubLiteReservationThroughputCapacity(res["throughputCapacity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	return nil
}

func resourcePubsubLiteReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	throughputCapacityProp, err := expandPubsubLiteReservationThroughputCapacity(d.Get("throughput_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("throughput_capacity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, throughputCapacityProp)) {
		obj["throughputCapacity"] = throughputCapacityProp
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{region}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Reservation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("throughput_capacity") {
		updateMask = append(updateMask, "throughputCapacity")
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

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Reservation %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Reservation %q: %#v", d.Id(), res)
	}

	return resourcePubsubLiteReservationRead(d, meta)
}

func resourcePubsubLiteReservationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{region}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Reservation %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Reservation")
	}

	log.Printf("[DEBUG] Finished deleting Reservation %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubLiteReservationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/reservations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/reservations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubLiteReservationThroughputCapacity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandPubsubLiteReservationThroughputCapacity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
