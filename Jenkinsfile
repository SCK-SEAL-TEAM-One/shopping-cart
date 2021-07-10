pipeline {
  agent any
  stages {
    stage('Install Dependency') {
      steps {
        sh 'make install_dependency_frontend'
      }
    }

    stage('Code Analysis') {
      parallel {
        stage('Frontend') {
          steps {
            sh 'make code_analysis_frontend'
          }
        }

        stage('Backend') {
          steps {
            script {
              def root = tool type: 'go', name: 'Go1.16.4'
              withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]){
                sh 'make code_analysis_backend'
              }
            }
          }
        }

      }
    }

    stage('Run Unit Testing') {
      parallel {
        stage('Frontend') {
          steps {
            sh 'make run_unittest_frontend'
          }
        }

        stage('Backend') {
          steps {
            script{
              def root = tool type: 'go', name: 'Go1.16.4'
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

    stage('Setup Test Fixtures') {
      steps {
        sh 'docker-compose up -d store-database bank-gateway shipping-gateway'
      }
    }

    stage('Run Integration Testing') {
      steps {
        script{
          def root = tool type: 'go', name: 'Go1.16.4'
          withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]){
            sh 'make run_integratetest_backend'
          }
        }
      }
    }

    stage('Build Docker Images') {
      parallel {
        stage('Build Frontend') {
          steps {
            sh 'make build_frontend'
          }
        }

        stage('Build Backend') {
          steps {
            sh 'make build_backend'
          }
        }

      }
    }

    stage('Run ATDD') {
      steps {
        sh 'make start_service'
        sh 'make run_newman'
        sh 'make run_robot'
        // sh 'make stop_service'
      }
    }

    stage('Run Performance Testing') {
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