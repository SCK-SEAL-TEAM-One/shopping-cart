frontend: install_dependency_frontend code_analys_frontend run_unittest_frontend build_frontend
backend: code_analys_backend run_unittest_backend build_backend run_integratetest_backend

run_robot: 
	robot atdd/ui/shopping_cart_success.robot

run_newman: 
	newman run atdd/api/shopping_cart_success.json -e atdd/api/environment/local_environment.json -d atdd/api/data/shopping_cart_success.json

install_dependency_frontend:
	cd store-web && npm install

code_analys_frontend:
	cd store-web && npm run lint

run_unittest_frontend:
	cd store-web && npm test

build_frontend:
	docker-compose build store-web

code_analys_backend:
	cd store-service && go vet ./...

run_unittest_backend:
	cd store-service && go test ./...

# ทำการ docker-compose up store-database ก่อน
run_integratetest_backend:
	docker-compose up -d store-database
	sleep 12
	cat tearup/init.sql | docker exec -i store-database /usr/bin/mysql -u sealteam --password=sckshuhari toy
	cd store-service && go test -tags=integration ./...
	docker-compose down

build_backend:
	docker-compose build store-service

start_service:
	docker-compose up -d

stop_service:
	docker-compose down
