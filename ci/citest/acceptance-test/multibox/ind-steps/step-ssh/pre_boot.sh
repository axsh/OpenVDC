#!/bin/bash

private_key="${CACHE_DIR}/${BRANCH}/sshkey_${vm_name}"
public_key="${private_key}.pub"

(
    $starting_step "Create the key pair and setup ssh config for ${user}"
    [[ -f ${CACHE_DIR}/${BRANCH}/sshkey_${vm_name} ]]
    $skip_step_if_already_done; set -ex
    add_user_key "${ci_user}"
    mv ${NODE_DIR}/sshkey "${private_key}"
    mv ${NODE_DIR}/sshkey.pub "${public_key}"
) ; prev_cmd_failed

(
    # If the remains from older code are present, the private key will exist
    # without the public key. In that case we need to generate it.
    # Once enough time has passed for this to no longer occur, this step can
    # be deleted.
    $starting_step "Generate public key for ${user}"
    [[ -f ${CACHE_DIR}/${BRANCH}/sshkey_${vm_name}.pub ]]
    $skip_step_if_already_done; set -ex
    ssh-keygen -y -f "${private_key}" > "${public_key}"
) ; prev_cmd_failed

(
    $starting_step "Install authorized ssh key for ${user} on ${vm_name}"
    sudo bash -c "[ -f ${TMP_ROOT}/${user}/.ssh/authorized_keys ]"
    $skip_step_if_already_done; set -ex
    install_user_key "${ci_user}"
) ; prev_cmd_failed

# Quick hack for now. Should be changed in the seed image instead
(
    $starting_step "Disable DNS on sshd"
    grep -qw "UseDNS yes" "${TMP_ROOT}/etc/ssh/sshd_config"
    $skip_step_if_already_done; set -ex
    sudo sed -i 's/UseDNS yes/UseDNS no/g' "${TMP_ROOT}/etc/ssh/sshd_config"
) ; prev_cmd_failed
