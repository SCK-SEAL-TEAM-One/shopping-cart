pipeline {
  agent any
  stages {
    stage('install dependency') {
      steps {
        sh 'make install_dependency_frontend'
      }
    }

    stage('code analysis') {
      parallel {
        stage('code analysis frontend') {
          steps {
            sh 'make code_analysis_frontend'
          }
        }

        stage('code analysis backend') {
          steps {
            sh 'make code_analysis_backend'
          }
        }

      }
    }

    stage('run unit test') {
      parallel {
        stage('code analysis frontend') {
          steps {
            sh 'make run_unittest_frontend'
          }
        }

        stage('code analysis backend') {
          steps {
            sh 'make run_unittest_backend'
            junit 'store-service/*.xml'
            script{
                def scannerHome = tool 'SonarQubeScanner';
                withSonarQubeEnv('SonarQubeScanner'){
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
          }
        }

      }
    }

    stage('setup test fixtures') {
      steps {
        sh 'docker-compose up -d store-database bank-gateway shipping-gateway'
      }
    }

    stage('run integration test') {
      steps {
        // sh 'make run_integratetest_backend'
        sh 'cd store-service && go test -tags=integration ./...'
      }
    }

    stage('build') {
      parallel {
        stage('build frontend') {
          steps {
            sh 'make build_frontend'
          }
        }

        stage('build backend') {
          steps {
            sh 'make build_backend'
          }
        }

      }
    }

    stage('run ATDD') {
      steps {
        sh 'make start_service'
        sh 'make run_newman'
        sh 'make run_robot'
        sh 'make stop_service'
      }
    }

  }
  post {
    always {
      robot outputPath: './', passThreshold: 100.0
      sh 'make stop_service'
      sh 'docker volume prune -f'
    }

  }
}