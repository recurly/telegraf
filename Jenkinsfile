#!/usr/bin/env groovy
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'make build'
            }
        }
        stage('Release Docker Image') {
            steps {
                sh 'make push-latest'
                sh 'make push-version'
            }
        }
    }
}
