#!/bin/bash
set -eou pipefail

apt-get update
apt-get install -y lsb-release wget software-properties-common gnupg
bash -c "$(wget -O - https://apt.llvm.org/llvm.sh)"
