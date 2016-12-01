#!/bin/bash

if mount | grep -q "${TMP_ROOT}" ; then
    umount-seed-image
fi

(
    $starting_step "Deploy seed image for ${vm_name}"
    [ -f "$(vm_image)" ]
    $skip_step_if_already_done; set -xe
    cd "${ENV_ROOTDIR}"
    tar -Sxzf "${SEED_IMAGE}"
    mv box-disk1.raw "${NODE_DIR}"
    rm box-disk1.rpm-qa
) ; prev_cmd_failed

(
    $starting_step "Mount temporary root folder for ${vm_name}"
    mount | grep -q "${TMP_ROOT}"
    $skip_step_if_already_done
    mkdir -p "${TMP_ROOT}"
    mount-partition --sudo "$(vm_image)" 1 "${TMP_ROOT}"
) ; prev_cmd_failed

(
    # Doing this step in init.sh so it happens before we start running yum.
    # If we install software using yum, their configuration timestampts will
    # be newer than the ones in guestroot, causing those files to not be
    # transfered because of the -u flag.
    $starting_step "Synching guestroot for ${vm_name}"
    # This step is set to false by default and we rely on rsyncs -u flag
    # to take care of keeping the files updated
    false
    $skip_step_if_already_done; set -ex
    sudo rsync -ruv "${NODE_DIR}/guestroot/" "${TMP_ROOT}"
) ; prev_cmd_failed
