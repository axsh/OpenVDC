#!groovy

import org.jenkinsci.plugins.workflow.steps.FlowInterruptedException
// http://stackoverflow.com/questions/37425064/how-to-use-environment-variables-in-a-groovy-function-using-a-jenkinsfile
import groovy.transform.Field
@Field final BUILD_OS_TARGETS=['el7']

@Field buildParams = [
  "BUILD_OS": "all",
  "REBUILD": "false",
  "LEAVE_CONTAINER": "0",
]
def ask_build_parameter = { ->
  return input(message: "Build Parameters", id: "build_params",
    parameters:[
      [$class: 'ChoiceParameterDefinition',
        choices: "all\n" + BUILD_OS_TARGETS.join("\n"), description: 'Target OS name', name: 'BUILD_OS'],
      [$class: 'ChoiceParameterDefinition',
        choices: "false\ntrue", description: 'Rebuild cache image', name: 'REBUILD'],
      [$class: 'ChoiceParameterDefinition',
        choices: "0\n1", description: 'Leave container after build for debugging.', name: 'LEAVE_CONTAINER'],
    ])
}

// Environment variables supplied by Jenkins system configuration:
// env.REPO_BASE_DIR
// env.BUILD_CACHE_DIR
def write_build_env(label) {
  def build_env="""# These parameters are read from bash and docker --env-file.
# So do not use single or double quote for the value part.
LEAVE_CONTAINER=${buildParams.LEAVE_CONTAINER}
REPO_BASE_DIR=${env.REPO_BASE_DIR ?: ''}
BUILD_CACHE_DIR=${env.BUILD_CACHE_DIR ?: ''}
BUILD_OS=${label}
REBUILD=${buildParams.REBUILD}
RELEASE_SUFFIX=${RELEASE_SUFFIX}
# https://issues.jenkins-ci.org/browse/JENKINS-30252
GIT_BRANCH=${env.BRANCH_NAME}
BRANCH_NAME=${env.BRANCH_NAME}
BRANCH=${env.BRANCH_NAME}
SHA=${SHA}
BOX_DISKSPACE=${env.BOX_DISKSPACE}
BOX_MEMORY=${env.BOX_MEMORY}
GOVC_DATACENTER=${env.GOVC_DATACENTER}
GOVC_INSECURE=${env.GOVC_INSECURE}
GOVC_URL=${env.GOVC_URL}
ISO=${env.ISO}
ISO_DATASTORE=${env.ISO_DATASTORE}
IP_ADDRESS=${env.IP_ADDRESS}
OS=${env.OS}
"""
  writeFile(file: "build.env", text: build_env)
}

def checkout_and_merge() {
    checkout scm
    sh "git -c \"user.name=Axsh Bot\" -c \"user.email=dev@axsh.net\" merge origin/master"
}

@Field RELEASE_SUFFIX=null
@Field SHA=null

def stage_unit_test(label) {
  node(label) {
    stage "Units Tests ${label}"
    checkout_and_merge()
    write_build_env(label)
    sh "./ci/citest/unit-tests/unit-tests.sh ./build.env"
  }
}

def stage_rpmbuild(label) {
  node(label) {
    stage "RPM Build ${label}"
    checkout_and_merge()
    write_build_env(label)
    sh "./ci/citest/rpmbuild/rpmbuild.sh ./build.env"
  }
}

def stage_acceptance(label) {
  node("multibox") {
    stage "Acceptance Test ${label}"
    checkout_and_merge()
    write_build_env(label)
    sh "./ci/citest/acceptance-test/build_and_run_in_docker.sh ./build.env"
  }
}

def stage_acceptance_esxi(label) {
  node("esxi") {
    stage "Acceptance Test Esxi ${label}"
    checkout_and_merge()
    write_build_env(label)
    sh "./ci/citest/acceptance-test-esxi/build.sh ./build.env"
  }
}

node() {
    stage "Checkout"
    try {
      timeout(time: 10, unit :"SECONDS") {
        buildParams = ask_build_parameter()
      }
    }catch(org.jenkinsci.plugins.workflow.steps.FlowInterruptedException err) {
      // Only ignore errors for timeout.
    }
    checkout scm
    // http://stackoverflow.com/questions/36507410/is-it-possible-to-capture-the-stdout-from-the-sh-dsl-command-in-the-pipeline
    // https://issues.jenkins-ci.org/browse/JENKINS-26133
    RELEASE_SUFFIX=sh(returnStdout: true, script: "./deployment/packagebuild/gen-dev-build-tag.sh").trim()
    SHA=sh(returnStdout: true, script: "git rev-parse --verify HEAD").trim()
}


build_nodes=BUILD_OS_TARGETS.clone()
if( buildParams.BUILD_OS != "all" ){
  build_nodes=[BUILD_OS]
}

// Using .each{} hits "a CPS-transformed closure is not yet supported (JENKINS-26481)"
for( label in build_nodes) {
  //stage_unit_test(label)
  stage_rpmbuild(label)
  //stage_acceptance(label)
  stage_acceptance_esxi(label)
}
