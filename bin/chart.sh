#!/usr/bin/env bash

# exit immediately when a command fails
set -e
# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail

NAME="capsule"
VERSION="${1}"
echo VERSION: ${VERSION}
echo "Building HELM3 chart for ${NAME} ${VERSION} version"
echo "HELM3 version $(helm version)"

# Creating a new dir in the CI build environment
CHART_TEMP_DIR="target"
mkdir -p "${CHART_TEMP_DIR}"

cp -R charts/${NAME} "${CHART_TEMP_DIR}/${NAME}"
# We need to bump Helm version during chart package to support chart OCI dependencies
curl -sSLf https://raw.githubusercontent.com/helm/helm/v3.16.3/scripts/get-helm-3 | bash -s -- --version v3.16.3 --no-sudo
helm package "${CHART_TEMP_DIR}/${NAME}" \
    --app-version=${VERSION} \
    --version=${VERSION} \
    --dependency-update \
    --destination=${CHART_TEMP_DIR}
