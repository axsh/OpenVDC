# This is a little trick to allow the rpmbuild command to define a suffix for
# development (non stable) versions.
%define release 1

Name: openvdc
Version: 0.1%{?dev_release_suffix:dev.%{dev_release_suffix}}
Release: %{release}%{?dist}
Summary: Metapackage that depends on all other OpenVDC packages.
Vendor: Axsh Co. LTD <dev@axsh.net>
URL: http://openvdc.org
Source: https://github.com/axsh/openvdc
License: LGPLv3

BuildArch: x86_64

BuildRequires: rpmdevtools lxc-devel git
BuildRequires: golang >= 1.6

Requires: mesosphere-zookeeper mesos
%{systemd_requires}
Requires: openvdc-cli
Requires: openvdc-executor
Requires: openvdc-scheduler

## This will not work with rpm v. 4.11 (which is what the jenkins vm has!)
##  rpm -v --querytags will give a list of acceptable tags. Suggests: *is*
## an acceptable tag for 4.12, apparently.
#Suggests: mesosphere-zookeeper mesos

%description
An empty metapackage that depends on all OpenVDC services. Just a conventient way to install everything at once on a single machine.

%files
# Metapackage, so no files!


%build
# rpmbuild resets $PATH so ensure to have "$GOPATH/bin".
export PATH="$PATH:${GOPATH}/bin"
cd "${GOPATH}/src/github.com/axsh/openvdc"
(
  VERSION=%{version} ./build.sh
)
cd "${GOPATH}/src/github.com/axsh/openvdc/ci/acceptance-test/tests"
go test -tags=acceptance -c -o openvdc-acceptance-test

%install
cd "${GOPATH}/src/github.com/axsh/openvdc"
mkdir -p "$RPM_BUILD_ROOT"/opt/axsh/openvdc/bin
mkdir -p "$RPM_BUILD_ROOT"%{_unitdir}
mkdir -p "$RPM_BUILD_ROOT"/usr/bin
ln -sf /opt/axsh/openvdc/bin/openvdc  "$RPM_BUILD_ROOT"/usr/bin
cp openvdc "$RPM_BUILD_ROOT"/opt/axsh/openvdc/bin
cp openvdc-executor "$RPM_BUILD_ROOT"/opt/axsh/openvdc/bin
cp openvdc-scheduler "$RPM_BUILD_ROOT"/opt/axsh/openvdc/bin
cp ci/acceptance-test/tests/openvdc-acceptance-test "$RPM_BUILD_ROOT"/opt/axsh/openvdc/bin
cp pkg/rhel/openvdc-scheduler.service "$RPM_BUILD_ROOT"%{_unitdir}
mkdir -p "${RPM_BUILD_ROOT}/etc/openvdc"
cp pkg/conf/executor.toml "${RPM_BUILD_ROOT}/etc/openvdc/"
cp pkg/conf/scheduler.toml "${RPM_BUILD_ROOT}/etc/openvdc/"

%package cli
Summary: OpenVDC cli

%description cli
The OpenVDC commandline interface.

%files cli
%dir /opt/axsh/openvdc
%dir /opt/axsh/openvdc/bin
/usr/bin/openvdc
/opt/axsh/openvdc/bin/openvdc

%package executor
Summary: OpenVDC executor
Requires: lxc

%description executor
This is a 'stub'. An appropriate message must be substituted at some point.

%files executor
%dir /opt/axsh/openvdc
%dir /opt/axsh/openvdc/bin
/opt/axsh/openvdc/bin/openvdc-executor
%dir /etc/openvdc
%config(noreplace) /etc/openvdc/executor.toml

%package scheduler
Summary: OpenVDC scheduler

%description scheduler
This is a 'stub'. An appropriate message must be substituted at some point.

%files scheduler
%dir /opt/axsh/openvdc
%dir /opt/axsh/openvdc/bin
/opt/axsh/openvdc/bin/openvdc-scheduler
%{_unitdir}/openvdc-scheduler.service
%config(noreplace) /etc/openvdc/scheduler.toml

%post
%{systemd_post openvdc-scheduler.service}

%postun
%{systemd_postun openvdc-scheduler.service}

%preun
%{systemd_preun openvdc-scheduler.service}

%package acceptance-test
Summary: The OpenVDC acceptance test used in its CI process.
Requires: openvdc-cli

%description acceptance-test
An acceptance test designed to run on a specifically designed environment. The environment building scripts can be found in the OpenVDC source code repository. The average OpenVDC user will not need to install this.

%files acceptance-test
%dir /opt/axsh/openvdc
%dir /opt/axsh/openvdc/bin
/opt/axsh/openvdc/bin/openvdc-acceptance-test
