pipeline {
    agent any
    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '5', daysToKeepStr: '', numToKeepStr: '5')
    }
    stages {
        stage('check docker engine') {
            steps {
                sh '''
                docker -v
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

        
        // stage('build microservice presence') {
        //     steps {
        //         sh '''
        //         cd API.Presence && dotnet build -o builds && scp 
        //         '''
        //     }
        // }

        stage('pre build docker image') {
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