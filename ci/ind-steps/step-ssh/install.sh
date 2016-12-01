#!/bin/bash

(
    $starting_step "Create the public key and setup ssh config"
    [[ -f ${TOYOCHO_ROOTDIR}/root@kemumaki ]]
    $skip_step_if_already_done; set -ex
    cat <<KEY > ${TOYOCHO_ROOTDIR}/root@kemumaki
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEApDCEWeLcLXxXKAy7+hHOwcPofvvsJ0APAAZpPF6rrrfbp+0I
wugnVnwM5WTZK5Js4JuzKTvR9xrdCSq2Sh4StjCVMLPUULZxBd5k615sPeTK/J9o
MXQqD+iucGrYX3TMQL216WRmakqqm1q7dO5EbNMpngbKZDinanNvdBZGYh4BidUV
lFgImtL/xS0NlIXUJQudEv7xKVcBY3TOAi0xhZO6vOVzP+1zy+Pwqksm/2RzOFOe
/FjYs/GS/D/kPJV3I+v/3JFYweTJc5BUcWK4271ObQnIVY9qXBvLzvJa7KX0pou3
D6+oijfPg7RbsD/iCpjKYwWDCqFXBlcKLPr8VQIDAQABAoIBAQCZQuyYVwo9rz5Q
FSr2r962wl9pAVGcBdC6rkFXZ+uMKPVyF/HAtiHaOetzyaJqaMEXHF+t0rgYmEvR
fbwxDUdcJ5drocFDIjn4R3MevcH+OG4+R7Jjz1JgDUufhiy5VfY/TJbc6Kos98uv
Po+TA8J9btL/Psl3qeakmGJmE1DZzjdWOuDxXpbxWBn+GyYegBqK8OxMoeLvLjwE
tX8eF98HVgkrCRirIbCLNhOcOPt2t0c/Osmi0srZxzkdIMS+pcWRs8NMWSphwiGo
xzODkPJOodC3baYKrM6fuUH8eoUISct75f9NE6yaC3IoPxr9x7g211EjhX0MuxmL
jkeMxqBVAoGBANehs10qbqQCY7c2vZf1NGlslp66qyr/+xWFM8CdaamZcQJnjNaz
QJbZ6eV1gU2hQyNKINVoXnaNXln80TecdxcvAvzlFO75NV7req0XOMU9idSCN19P
9mV5E+TbsSC0CfFY25cco6pXmYqjNjeo5jt7gdZXmi76XeIfCS7lzeUTAoGBAMLt
aoZ8Hf5XmcgwebG/8xOS2eguMdrFYIttbYYmRPBHrQ9609k0vYUEMcOWdQ9H2PvX
NRkSkW+vlASod07T1u61Oc5PW+DQdmOm/DBBk6VBugnboT8JBpC5q2ziFkPWoiYN
/HlfTji8FgFnYs4gFlborZLo3y4KT+0LQot5qw33AoGBANUVwwTPGLTUc6uq2aKf
umJv30wOFXYRrhKvJdwy8iaaLhX9NC06yBoKT9vjyZpoQMtPxrB9SdQHkXSFpE47
PWhYmbBFxPD+reIV+42vA/fN/zVUVNnIUCoogyNGgnxZzfWFJMLol0eO93kMHljU
fGfNfhnNSd0gpsP8fXuttlCnAoGAZrxU0a6m7D9lWMJb0Vg2mBFk8d8u/Xvbx4CU
IgkXeoB8cNUuxKLsZCfyxoANEFGwxgMGBEHqDzA91jqoLw6tkEBJyDfBV/rm+tyu
vAxD8exzTfDIyn1KTYreE9QRIg4EhgQHFj0p7/MJshG4XKVCwOl/WheMUSq0N14g
43vrsL8CgYEAlxpCyOj6li4GfOW1ocuvgubs76VGQfdFCpIsYFm7kgf9ntMoQ6zd
FdqrUqBCsJ8BCFhKtBGNW/NSlL9aNG9WawjU5aWI+g28o3cQSpXvCPtZJFNQ6gjY
1ENqhAaBs6ph9rXu9VRt97o9Sudgn6qcRehKwS+jOp4y0hybg2M0d20=
-----END RSA PRIVATE KEY-----
KEY
    chmod 600 ${TOYOCHO_ROOTDIR}/root@kemumaki

) ; prev_cmd_failed

(
    $starting_step "Install authorized ssh key for ${vm_name}"
    sudo chroot ${TMP_ROOT} /bin/bash -c \
         "[[ -f /root/.ssh/authorized_keys ]]"
    $skip_step_if_already_done; set -xe
    sudo chroot ${TMP_ROOT} /bin/bash -ex <<EOS
mkdir -p -m 700 /root/.ssh
chmod 700 /root/.ssh
cat <<CFG > /root/.ssh/authorized_keys
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCkMIRZ4twtfFcoDLv6Ec7Bw+h+++wnQA8ABmk8Xquut9un7QjC6CdWfAzlZNkrkmzgm7MpO9H3Gt0JKrZKHhK2MJUws9RQtnEF3mTrXmw95Mr8n2gxdCoP6K5wathfdMxAvbXpZGZqSqqbWrt07kRs0ymeBspkOKdqc290FkZiHgGJ1RWUWAia0v/FLQ2UhdQlC50S/vEpVwFjdM4CLTGFk7q85XM/7XPL4/CqSyb/ZHM4U578WNiz8ZL8P+Q8lXcj6//ckVjB5MlzkFRxYrjbvU5tCchVj2pcG8vO8lrspfSmi7cPr6iKN8+DtFuwP+IKmMpjBYMKoVcGVwos+vxV root@localhost.localdomain
CFG

chmod 600 /root/.ssh/authorized_keys
  
sed -i \
-e 's,^PermitRootLogin .*,PermitRootLogin yes,' \
-e 's,^PasswordAuthentication .*,PasswordAuthentication yes,' \
-e 's,^DenyUsers.root,#DenyUsers root,' \
\
/etc/ssh/sshd_config
EOS
  
) ; prev_cmd_failed

if [[ ${vna} == true ]]; then
    (
        $starting_step "Install autorized key for jenkins"
        [[ $(sudo wc -l ${TMP_ROOT} /root/.ssh/authorized_keys | awk '{ print $1 }') -eq 2 ]]
        $skip_step_if_already_done; set -xe
        sudo chroot ${TMP_ROOT} /bin/bash -ex <<EOS
cat <<CFG >> /root/.ssh/authorized_keys
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAtN9MMMPtRMY9wGRXvivF5mkSfY+/ZN7wfLPu1JMtCgxN9TgEmc9ag95e+EcWke4e6qiYFFfarU7OtMSIu4qKjXnolbcjzbm+dHBrQKsK4rZEggOtORANcsZRWGYhE78fvGzR4Bpfs3Gy0ko1BLEOATBPTreai2T5168vKrJG0tdA/77pvAwi41kppqQXlULuU3R20ynWZRsrX8JPb9BJIz/jusVskDX5U3Waw3jUi2qSfi+gBaVKb3uFk/ctUXec/etRWut1oiodn/Gd8WvrlWFuG/0Ob0aqg8bO51Yl657Kf3llFNvKYUPOfk0g8VvwHGWtEDjQ28/6RteQRiPMEw== root@kemumaki
CFG
EOS
    ) ; prev_cmd_failed
fi
