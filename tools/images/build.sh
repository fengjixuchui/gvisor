#!/bin/bash

# Copyright 2019 The gVisor Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is responsible for building a new GCP image that: 1) has nested
# virtualization enabled, and 2) has been completely set up with the
# image_setup.sh script. This script should be idempotent, as we memoize the
# setup script with a hash and check for that name.

set -xeou pipefail

# Parameters.
declare -r USERNAME=${USERNAME:-test}
declare -r IMAGE_PROJECT=${IMAGE_PROJECT:-ubuntu-os-cloud}
declare -r IMAGE_FAMILY=${IMAGE_FAMILY:-ubuntu-1604-lts}
declare -r ZONE=${ZONE:-us-central1-f}

# Random names.
declare -r DISK_NAME=$(mktemp -u disk-XXXXXX | tr A-Z a-z)
declare -r SNAPSHOT_NAME=$(mktemp -u snapshot-XXXXXX | tr A-Z a-z)
declare -r INSTANCE_NAME=$(mktemp -u build-XXXXXX | tr A-Z a-z)

# Hash inputs in order to memoize the produced image.
declare -r SETUP_HASH=$( (echo ${USERNAME} ${IMAGE_PROJECT} ${IMAGE_FAMILY} && cat "$@") | sha256sum - | cut -d' ' -f1 | cut -c 1-16)
declare -r IMAGE_NAME=${IMAGE_FAMILY:-image-}${SETUP_HASH}

# Does the image already exist? Skip the build.
declare -r existing=$(gcloud compute images list --filter="name=(${IMAGE_NAME})" --format="value(name)")
if ! [[ -z "${existing}" ]]; then
  echo "${existing}"
  exit 0
fi

# gcloud has path errors; is this a result of being a genrule?
export PATH=${PATH:-/bin:/usr/bin:/usr/local/bin}

# Start a unique instance. Note that this instance will have a unique persistent
# disk as it's boot disk with the same name as the instance.
gcloud compute instances create \
    --quiet \
    --image-project "${IMAGE_PROJECT}" \
    --image-family "${IMAGE_FAMILY}" \
    --boot-disk-size "200GB" \
    --zone "${ZONE}" \
    "${INSTANCE_NAME}" >/dev/null
function cleanup {
    gcloud compute instances delete --quiet --zone "${ZONE}" "${INSTANCE_NAME}"
}
trap cleanup EXIT

# Wait for the instance to become available (up to 5 minutes).
declare timeout=300
declare success=0
declare internal=""
declare -r start=$(date +%s)
declare -r end=$((${start}+${timeout}))
while [[ "$(date +%s)" -lt "${end}" ]] && [[ "${success}" -lt 3 ]]; do
  if gcloud compute ssh --zone "${internal}" "${ZONE}" "${USERNAME}"@"${INSTANCE_NAME}" -- env - true 2>/dev/null; then
    success=$((${success}+1))
  elif gcloud compute ssh --zone --internal-ip "${ZONE}" "${USERNAME}"@"${INSTANCE_NAME}" -- env - true 2>/dev/null; then
    success=$((${success}+1))
    internal="--internal-ip"
  fi
done

if [[ "${success}" -eq "0" ]]; then
  echo "connect timed out after ${timeout} seconds."
  exit 1
fi

# Run the install scripts provided.
for arg; do
  gcloud compute ssh --zone "${internal}" "${ZONE}" "${USERNAME}"@"${INSTANCE_NAME}" -- sudo bash - <"${arg}" >/dev/null
done

# Stop the instance; required before creating an image.
gcloud compute instances stop --quiet --zone "${ZONE}" "${INSTANCE_NAME}" >/dev/null

# Create a snapshot of the instance disk.
gcloud compute disks snapshot \
    --quiet \
    --zone "${ZONE}" \
    --snapshot-names="${SNAPSHOT_NAME}" \
    "${INSTANCE_NAME}" >/dev/null

# Create the disk image.
gcloud compute images create \
    --quiet \
    --source-snapshot="${SNAPSHOT_NAME}" \
    --licenses="https://www.googleapis.com/compute/v1/projects/vm-options/global/licenses/enable-vmx" \
    "${IMAGE_NAME}" >/dev/null

# Finish up.
echo "${IMAGE_NAME}"
