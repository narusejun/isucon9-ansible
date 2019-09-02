.PHONY: all
all: build restart

.PHONY: build
build:
	cd app
	git fetch
	git reset --hard origin/master
	go build -o app

.PHONY: restart
restart:
	sudo systemctl restart isucon9-app
