#!/usr/bin/env groovy
node {
    properties([buildDiscarder(logRotator(daysToKeepStr: '14'))])

    stage('Checkout') {
      deleteDir()
      checkout scm
    }

    stage("Build") {
      echo "Testing and building."
      sh 'make build'
    }

    stage("Release") {
      sh 'make publish'
    }
}
