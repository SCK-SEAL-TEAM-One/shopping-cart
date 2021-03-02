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
            script {
              def root = tool type: 'go', name: 'Go1.15.6'
              withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]){
                sh 'make code_analysis_backend'
              }
            }
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
            script{
              def root = tool type: 'go', name: 'Go1.15.6'
              withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]){
                sh 'go get github.com/jstemmer/go-junit-report'
                sh 'cd store-service && go test -v -coverprofile=coverage.out ./... 2>&1 | /var/lib/jenkins/go/bin/go-junit-report > coverage.xml'
                junit 'store-service/*.xml'
              }
              def scannerHome = tool 'SonarQubeScanner';
              withSonarQubeEnv('SonarQubeScanner'){
                sh "${scannerHome}/bin/sonar-scanner"
              }
            }
            // sh 'make run_unittest_backend'
            // junit 'store-service/*.xml'
            // script{
            //     def scannerHome = tool 'SonarQubeScanner';
            //     withSonarQubeEnv('SonarQubeScanner'){
            //         sh "${scannerHome}/bin/sonar-scanner"
            //     }
            // }
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
        script{
          def root = tool type: 'go', name: 'Go1.15.6'
          withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]){
            sh 'make run_integratetest_backend'
          }
        }
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
        // sh 'make stop_service'
      }
    }

    stage('run Performance Testing') {
      steps {
        sh 'make run_performance_test_k6'
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