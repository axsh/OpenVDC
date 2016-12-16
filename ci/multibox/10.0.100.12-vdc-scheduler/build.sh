#!/bin/bash

export ENV_ROOTDIR="$(cd "$(dirname $(readlink -f "$0"))/.." && pwd -P)"
export NODE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TMP_ROOT="${NODE_DIR}/tmp_root"

. "${ENV_ROOTDIR}/config.source"
. "${NODE_DIR}/vmspec.conf"
. "${ENV_ROOTDIR}/ind-steps/common.source"

scheduler=true

IND_STEPS=(
    "box"
    "ssh"
    "epel"
    "lxc"
    "mesosphere"
    "mesos"
    "zookeeper"
)

build "${IND_STEPS[@]}"
