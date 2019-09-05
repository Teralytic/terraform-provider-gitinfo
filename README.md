# terraform-provider-gitinfo

Use Git context (current branch name and SHA) in Terraform scripts.

## Use Case

We store Terraform module files in the same repo as our code.
CI builds a package (for instance, a docker image) from a commit and tags it with the git SHA.
When the master Terraform script applies the module, it builds a deployment specifying the docker image using the module SHA.
In order to do that, we use `terraform-provider-gitinfo` to get easy access to the git SHA.

## Terraform Example

```terraform
provider "gitinfo" {}

data "gitinfo_repo" "root" {
  path = "${path.root}"
  # In a module you would use "${path.module}"
}
```

Once this data resource is defined, you can use any of the following

```terraform
  "${data.gitinfo_repo.root.branch_full}"   # refs/heads/master
  "${data.gitinfo_repo.root.branch_short}"  # master
  "${data.gitinfo_repo.root.sha_full}"      # f78a498022d81d5b933ab68801f0b4387aa197d5
  "${data.gitinfo_repo.root.sha_short}"     # f78a498
```

## Installation

* Download `terraform-provider-gitinfo` binary from [Github](https://github.com/teralytic/terraform-provider-gitinfo/releases)
* Unzip the zip file
* Move `terraform-provider-gitinfo` binary to `$HOME/.terraform.d/plugins` directory

```bash
mkdir -p $HOME/.terraform.d/plugins
mv terraform-provider-gitinfo $HOME/.terraform.d/plugins/terraform-provider-gitinfo
```

* Run `terraform init` in your terraform project

```bash
terraform init
```

## Credits

This code based on Matthias Bartelmeß's https://github.com/fourplusone/terraform-provider-git.
(I didn't simply fork because `terraform-provider-git` because it serves a different purpose.
That one is about committing files to a configured repo.
This provider provides information without knowing a remote repo URL.)