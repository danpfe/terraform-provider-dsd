package tfpdsd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func resourceReservedName() *schema.Resource {
	return &schema.Resource{
		Description: "Reserves a name.",

		CreateContext: resourceReservedNameCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"something": {
				Description: "Something something?",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceReservedNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	in_data := map[string]string{
		"something": "Something else",
	}

	in_data_json, err := json.Marshal(in_data)
	if err != nil {
		return diag.FromErr(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/reserve_name", meta.(string)), bytes.NewBuffer(in_data_json))
	if err != nil {
		return diag.FromErr(err)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var parsed_response DsdResponse
	if err := json.Unmarshal(body, &parsed_response); err != nil {
		fmt.Println("Shit's broke, yo!")
	}

	d.SetId(strconv.FormatInt(parsed_response.ID, 10))

	var diags diag.Diagnostics
	return diags
}

type DsdResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
