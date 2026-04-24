pipeline {
    agent any
    tools {
    go 'go-tool'
    }
    environment {
            // Set GOPATH and append its bin folder to the PATH
            GOPATH = "${WORKSPACE}/go"
            PATH = "${GOPATH}/bin:${env.PATH}"
        }
    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'go mod download'
//                 sh 'go mod tidy'
                sh 'go get github.com/onsi/ginkgo/v2/ginkgo'
            }
        }
        stage('Ginkgo') {
                    steps {
                        sh 'ginkgo -v ./address-api-tests'
                    }
                }
    }
}

