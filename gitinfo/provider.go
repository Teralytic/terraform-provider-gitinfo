package gitinfo

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns the Gitinfo Provider
func Provider() *schema.Provider {
	p := schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"gitinfo_repo": dataSourceGitRepo(),
		},
		ResourcesMap: map[string]*schema.Resource{
			// No resources
		},
		Schema: map[string]*schema.Schema{
			// No provider configuration
		},
	}

	p.ConfigureFunc = configureProviderFunc(&p)

	return &p
}

func configureProviderFunc(p *schema.Provider) schema.ConfigureFunc {
	return func(r *schema.ResourceData) (interface{}, error) {
		return nil, nil
	}
}
