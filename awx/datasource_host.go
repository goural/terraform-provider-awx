package awx

import (
	"strconv"

	awxgo "github.com/davidfischer-ch/awx-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInventoryRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of this inventory",
			},
			"id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Id of the ansible inventory",
			},
		},
	}
}

func dataSourceHostRead(d *schema.ResourceData, meta interface{}) error {
	awx := meta.(*awxgo.AWX)
	awxService := awx.HostService
	_, res, err := awxService.ListHosts(map[string]string{
		"name": d.Get("name").(string)})
	if err != nil {
		return err
	}
	if len(res.Results) == 0 {
		return nil
	}
	d.SetId(strconv.Itoa(res.Results[0].ID))
	d = setHostSourceData(d, res.Results[0])
	return nil
}

func setHostSourceData(d *schema.ResourceData, r *awxgo.Host) *schema.ResourceData {
	d.Set("name", r.Name)
	d.Set("id", r.ID)
	return d
}
