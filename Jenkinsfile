#!/usr/bin/env groovy
node {
    properties([buildDiscarder(logRotator(daysToKeepStr: '14'))])

    stage('Checkout') {
      deleteDir()
      checkout scm
    }
    stage("Build") {
      sh 'make build'
    }
    stage("Test") {
      sh 'make test'
    }
    stage("Release") {
      sh 'make test'
    }
}
