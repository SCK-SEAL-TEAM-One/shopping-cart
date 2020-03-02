fontend: install_dependency_fontend code_analys_fontend run_unittest_fontend build_fontend
backend: code_analys_backend run_unittest_backend build_backend

run_robot: 
	robot atdd/ui/shopping_cart_success.robot

run_newman: 
	newman run atdd/api/shopping_cart_success.json

install_dependency_fontend:
	cd store-web && npm install

code_analys_fontend:
	cd store-web && npm run lint

run_unittest_fontend:
	cd store-web && npm test

build_fontend:
	docker-compose build store-web

code_analys_backend:
	cd store-service && go vet ./...

run_unittest_backend:
	cd store-service && go test ./...

run_integratetest_backend:
	cd store-service && docker-compose up -d store-database
	sleep 12
	cd store-service && go test -tags=integration ./...
	cd store-service && docker-compose down

build_backend:
	docker-compose build store-service

integration_test_backend:
	cat tearup/init.sql | docker exec -i store-database /usr/bin/mysql -u sealteam --password=sckshuhari toy
	cd store-service && go test -tags=integration ./...
	docker-compose down