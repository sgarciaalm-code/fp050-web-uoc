pipeline {
    agent any

    stages {
        stage('Verificar repositorio') {
            steps {
                echo 'Codigo descargado desde GitHub'
                sh 'git --version'
                sh 'test -f main.go'
                sh 'test -f Dockerfile'
            }
        }

        stage('Preparar despliegue') {
            steps {
                echo 'Pipeline inicial funcionando correctamente'
            }
        }
    }
}