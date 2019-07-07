build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

bash:
	docker exec -it go_test_go_test_1 bash
