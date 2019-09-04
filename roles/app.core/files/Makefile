REF?=origin/master

.ONESHELL:

.PHONY: all
all: build restart

.PHONY: build
build:
	cd app
	rm -rf *
	git fetch
	git reset --hard $(REF)
	make

.PHONY: restart
restart:
	sudo systemctl restart isucon9-app
