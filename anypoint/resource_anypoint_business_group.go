package anypoint

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rhoegg/go-anypoint/anypointplatform"
	"log"
)

func resourceAnypointBusinessGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnypointBusinessGroupCreate,
		Read:   resourceAnypointBusinessGroupRead,
		Delete: resourceAnypointBusinessGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
				ForceNew:     true,
			},

			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAnypointBusinessGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*anypointplatform.Client)

	// Build up our creation options
	opts := &anypointplatform.BusinessGroupCreateRequest{
		Name: d.Get("name").(string),
	}

	log.Printf("[DEBUG] Business group create configuration: %#v", opts)
	bg, _, err := client.BusinessGroup.Create(context.Background(), opts)
	if err != nil {
		return fmt.Errorf("Error creating Business group: %s", err)
	}

	d.SetId(bg.ID)
	log.Printf("[INFO] Business Group: %s", bg.ID)

	return resourceAnypointBusinessGroupRead(d, meta)
}

func resourceAnypointBusinessGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*anypointplatform.Client)

	bg, resp, err := client.BusinessGroup.Get(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error retrieving Business Group: %s", err)
	}
	if resp.StatusCode != 404 {
		d.Set("name", bg.Name)
		d.Set("clientId", bg.ClientID)
	} else {
		d.SetId("")
	}

	return nil
}

func resourceAnypointBusinessGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*anypointplatform.Client)

	log.Printf("[INFO] Deleting Business group: %s", d.Id())
	_, err := client.BusinessGroup.Delete(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting BusinessGroup: %s", err)
	}

	d.SetId("")
	return nil
}
