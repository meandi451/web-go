pipeline {
    agent any

    stages {
        stage('Construir imagen') {
            steps {
                sh 'docker build -t web-go:latest .'
            }
        }
        stage('Desplegar contenedor') {
            steps {
                sh 'docker rm -f web-go || true'
                sh 'docker run -d -p 8080:8080 --name web-go web-go:latest'
            }
        }
    }
}
