provider "gitinfo" {}


data "gitinfo_repo" "module" {
  path = "${path.root}"
  # path = "${path.module}"
}

resource "null_resource" "show" {
  triggers = {
    run_every_time = "${uuid()}"
  }

  provisioner "local-exec" {
    command = <<-COMMAND
      echo Full Branch: ${data.gitinfo_repo.module.branch_full} &&
      echo Short Branch: ${data.gitinfo_repo.module.branch_short} &&
      echo Full SHA: ${data.gitinfo_repo.module.sha_full} &&
      echo Short SHA: ${data.gitinfo_repo.module.sha_short}
    COMMAND
  }
}

