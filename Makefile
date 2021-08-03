.SILENT:

# Посмотреть Golang версию.
goversion:
	go version

# Собрать бинарник и положить в папку bin.
build: goversion
	go build -o ./bin/akina ./cmd/akina/main.go

# Собрать бинарник и запустить локально.
build_run_local: build
	./bin/akina

# Запустить локально.
run_local:
	./bin/akina

# Собрать бинарник приложения и создать docker image с этим бинарником.
build_docker_img: build
	docker build -t akinaimg ./

# Запустить docker container.
run_docker_ctr:
	docker run --name akinactr akinaimg

# Запустить существующий docker container.
start_docker_ctr:
	docker start akinactr

# Остановить запущенный docker container.
stop_docker_ctr:
	docker stop akinactr

# Heroku -----------------------------------------------------------|
# heroku login
# heroku create art07akina
# heroku container:login
# Создать docker image для heroku из уже существующего.
build_docker_img_for_heroku: build_docker_img
	docker tag akinaimg registry.heroku.com/art07akina/worker
# Загрузить docker image на heroku.
send_docker_img_to_heroku:
	docker push registry.heroku.com/art07akina/worker
# heroku container:release worker --app art07akina
# heroku ps:scale worker=1 --app art07akina
# Запустить лог.
akina_log:
	heroku logs --tail --app art07akina