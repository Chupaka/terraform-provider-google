// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceStorageBucketAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageBucketAccessControlCreate,
		Read:   resourceStorageBucketAccessControlRead,
		Update: resourceStorageBucketAccessControlUpdate,
		Delete: resourceStorageBucketAccessControlDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageBucketAccessControlImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the bucket.`,
			},
			"entity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The entity holding the permission, in one of the following forms:
  user-userId
  user-email
  group-groupId
  group-email
  domain-domain
  project-team-projectId
  allUsers
  allAuthenticatedUsers
Examples:
  The user liz@example.com would be user-liz@example.com.
  The group example@googlegroups.com would be
  group-example@googlegroups.com.
  To refer to all members of the Google Apps for Business domain
  example.com, the entity would be domain-example.com.`,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"OWNER", "READER", "WRITER", ""}, false),
				Description:  `The access permission for the entity.`,
			},
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The domain associated with the entity.`,
			},
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The email address associated with the entity.`,
			},
		},
	}
}

func resourceStorageBucketAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	lockName, err := replaceVars(d, config, "storage/buckets/{{bucket}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BucketAccessControl: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BucketAccessControl: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BucketAccessControl %q: %#v", d.Id(), res)

	return resourceStorageBucketAccessControlRead(d, meta)
}

func resourceStorageBucketAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("StorageBucketAccessControl %q", d.Id()))
	}

	if err := d.Set("bucket", flattenStorageBucketAccessControlBucket(res["bucket"], d, config)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("domain", flattenStorageBucketAccessControlDomain(res["domain"], d, config)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("email", flattenStorageBucketAccessControlEmail(res["email"], d, config)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("entity", flattenStorageBucketAccessControlEntity(res["entity"], d, config)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("role", flattenStorageBucketAccessControlRole(res["role"], d, config)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}

	return nil
}

func resourceStorageBucketAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	lockName, err := replaceVars(d, config, "storage/buckets/{{bucket}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BucketAccessControl %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating BucketAccessControl %q: %s", d.Id(), err)
	}

	return resourceStorageBucketAccessControlRead(d, meta)
}

func resourceStorageBucketAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "storage/buckets/{{bucket}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BucketAccessControl %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BucketAccessControl")
	}

	log.Printf("[DEBUG] Finished deleting BucketAccessControl %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageBucketAccessControlImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<bucket>[^/]+)/(?P<entity>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenStorageBucketAccessControlBucket(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenStorageBucketAccessControlDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenStorageBucketAccessControlEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenStorageBucketAccessControlEntity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenStorageBucketAccessControlRole(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandStorageBucketAccessControlBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlEntity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
