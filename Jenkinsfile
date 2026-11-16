pipeline {
    agent any
    stages {
        stage('Elegir entorno') {
            steps {
                script {
                    if (env.GIT_BRANCH == 'origin/main') {
                        env.NS = 'produccion'
                        env.PUERTO = '8083'
                    } else {
                        env.NS = 'pruebas'
                        env.PUERTO = '8082'
                    }
                }
            }
        }
        stage('Construir imagen') {
            steps {
                sh 'docker build -t web-go:$NS .'
            }
        }
        stage('Cargar imagen en Minikube') {
            steps {
                sh 'minikube image load web-go:$NS'
            }
        }
        stage('Desplegar en Kubernetes') {
            steps {
                sh 'kubectl apply -n $NS -f k8s/'
                sh 'kubectl set image deployment/web-go web-go=web-go:$NS -n $NS'
                sh 'kubectl rollout restart deployment/web-go -n $NS'
                sh 'kubectl rollout status deployment/web-go -n $NS'
            }
        }
        stage('Abrir puerto') {
            steps {
                sh """
                    NODEPORT=\$(kubectl get svc web-go -n \$NS -o jsonpath='{.spec.ports[0].nodePort}')
                    docker rm -f entrada-\$NS || true
                    docker run -d --restart unless-stopped --network host --name entrada-\$NS alpine/socat tcp-listen:\$PUERTO,fork,reuseaddr tcp:\$(minikube ip):\$NODEPORT
                """
            }
        }
    }
}
