#!/bin/bash

set -ex -o pipefail

whereami="$(cd "$(dirname $(readlink -f "$0"))" && pwd -P)"

BUILD_ENV_PATH=${1:?"ERROR: env file is not given."}
if [[ -n "${BUILD_ENV_PATH}" && ! -f "${BUILD_ENV_PATH}" ]]; then
  echo "ERROR: Can't find the file: ${BUILD_ENV_PATH}" >&2
  exit 1
fi

set -a
. ${BUILD_ENV_PATH}
set +a

DATA_DIR="${DATA_DIR:-/data2}"
CACHE_DIR="/data/openvdc-ci/branches"

repo_and_tag="openvdc/acceptance-test:${BRANCH}.${RELEASE_SUFFIX}"

function cleanup() {
  if [[ -z "${LEAVE_CONTAINER}" || "${LEAVE_CONTAINER}" == "0" ]]; then
    # Clean up containers
    # Images don't need to be cleaned up. Removing them immediately would slow down
    # builds and they can be garbage collected later.
    for CID in $(sudo docker ps -af ancestor="${repo_and_tag}" --format "{{.ID}}"); do
      sudo docker rm "${CID}"
    done
  else
    echo "LEAVE_CONTAINER was set and not 0. Skip container cleanup."
  fi

  # Give the newly created cache to this user instead of root so we can clean it up
  # later without needing sudo
  local user=$(/usr/bin/id -run)
  sudo chown -R $user:$user "${CACHE_DIR}"/"${BRANCH}"
}
trap "cleanup" EXIT

sudo docker build -t "${repo_and_tag}" --build-arg BRANCH="${BRANCH}" \
                                  --build-arg RELEASE_SUFFIX="${RELEASE_SUFFIX}" \
                                  --build-arg REBUILD="${REBUILD}" \
                                  "${whereami}"

sudo docker run --privileged -v "${DATA_DIR}":/data "${repo_and_tag}"
