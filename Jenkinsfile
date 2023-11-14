pipeline {
    agent any
    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '5', daysToKeepStr: '', numToKeepStr: '5')
    }
    stages {
        stage('check docker engine') {
            steps {
                sh '''
                docker -version
                '''
            }
        }

        stage('check repo path') {
            steps {
                sh '''
                pwd
                ls -l
                '''
            }
        }

        stage('build docker image') {
            steps {
                sh '''
                docker build -t khilmi-aminudin/todo-app .
                '''
            }
        }

        stage('pree build docker image') {
            // when {
            //     branch "fix-*"
            // }
            steps {
                sh '''
                echo "build success"
                '''
            }
        }
    }
}