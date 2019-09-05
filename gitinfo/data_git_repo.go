package gitinfo

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func dataSourceGitRepo() *schema.Resource {
	return &schema.Resource{

		Read: readGitInfo,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"branch_full": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"branch_short": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"sha_full": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"sha_short": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readGitInfo(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	headRef, err := r.Head()
	if err != nil {
		return err
	}

	sha := headRef.Hash().String()

	branches, err := r.Branches()
	if err != nil {
		return err
	}
	var branch string
	branches.ForEach(func(b *plumbing.Reference) error {
		if b.Hash().String() == sha {
			branch = b.Name().String()
		}
		return nil
	})

	d.SetId(path)
	d.Set("branch_full", branch)
	d.Set("branch_short", strings.TrimPrefix(branch, "refs/heads/"))
	d.Set("sha_full", sha)
	d.Set("sha_short", sha[:7])

	return nil
}
