pipeline {
    agent any
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/sschrs/go_api.git'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o output_binary .'
            }
        }
        stage('Run') {
            steps {
                sh './output_binary'
            }
        }
    }
}
