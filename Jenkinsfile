pipeline {
    agent any

    stages {
        stage('Verificar repositorio') {
            steps {
                echo 'Comprobando archivos del proyecto'
                sh 'test -f main.go'
                sh 'test -f Dockerfile'
            }
        }

        stage('Comprobar Docker') {
            steps {
                echo 'Comprobando instalacion de Docker'
                sh 'docker --version'
                sh 'docker info'
            }
        }

        stage('Construir imagen Docker') {
            steps {
                echo 'Construyendo imagen de la aplicacion Go'
                sh 'docker build -t fp050-web-uoc:$BUILD_NUMBER .'
            }
        }

        stage('Verificar imagen') {
            steps {
                echo 'Comprobando imagen generada'
                sh 'docker image inspect fp050-web-uoc:$BUILD_NUMBER'
            }
        }
    }
}