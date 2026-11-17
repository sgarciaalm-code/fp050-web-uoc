pipeline {
    agent any

    options {
        disableConcurrentBuilds()
        timeout(time: 30, unit: 'MINUTES')
    }

    stages {
        stage('Verificar repositorio') {
            steps {
                echo 'Comprobando archivos del proyecto'

                sh 'test -f main.go'
                sh 'test -f Dockerfile'
                sh 'test -f k8s/deployment.yaml'
                sh 'test -f k8s/service.yaml'
            }
        }

        stage('Construir imagen Docker') {
            steps {
                echo 'Construyendo nueva imagen Docker'

                sh 'docker build -t fp050-web-uoc:$BUILD_NUMBER .'
            }
        }

        stage('Cargar imagen en Minikube') {
            steps {
                echo 'Importando imagen al cluster Kubernetes'

                sh '''
                    set -eu

                    docker save "fp050-web-uoc:${BUILD_NUMBER}" | \
                    docker exec -i minikube ctr -n k8s.io images import -

                    docker exec minikube ctr -n k8s.io images ls | \
                    grep "fp050-web-uoc:${BUILD_NUMBER}"
                '''
            }
        }

        stage('Desplegar en Kubernetes') {
            steps {
                echo 'Actualizando las replicas de Kubernetes'

                sh '''
                    set -eu

                    kubectl apply -f k8s/service.yaml

                    sed "s|image: fp050-web-uoc:.*|image: fp050-web-uoc:${BUILD_NUMBER}|" \
                    k8s/deployment.yaml | kubectl apply -f -

                    kubectl rollout status deployment/fp050-web-uoc --timeout=240s
                '''
            }
        }

        stage('Verificar despliegue') {
            steps {
                echo 'Comprobando estado final'

                sh 'kubectl get deployment fp050-web-uoc'
                sh 'kubectl get pods -l app=fp050-web-uoc'
                sh 'kubectl get svc fp050-web-uoc'
            }
        }
    }

    post {
        success {
            echo 'Despliegue automatico completado correctamente'
        }

        failure {
            echo 'Error durante el despliegue automatico'
        }
    }
}