#!/bin/bash
set -euo pipefail

useless_backup_state='terraform.tfstate.backup'
if [ -e $useless_backup_state ]; then
  rm $useless_backup_state
fi

terraform init \
   -backend-config="project=nyt-interview-sreenu-yedavalli" \
   -backend-config="bucket=nyt-hello-tf" \`
   -backend-config="path=terraform.tfstate" \
   -backend-config="credentials=${TF_VAR_credential_path}"
