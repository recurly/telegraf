#!/usr/bin/env groovy
node {
    properties([buildDiscarder(logRotator(daysToKeepStr: '14'))])

    stage('Checkout') {
      deleteDir()
      checkout scm
      sh '''
         wget https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
         tar xvvf go1.10.2.linux-amd64.tar.gz
      '''
    }

    stage("Build") {
      echo "Testing and building."
      sh '''
        export GOPATH=${PWD}/go"
        export PATH="${GOPATH}/bin:${PATH}"
        make build
      '''
    }
}
