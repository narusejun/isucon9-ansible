kataru() {
	sudo cat /var/log/nginx/access.log | kataribe -f /etc/kataribe.toml
}

katarazu() {
	echo -n | sudo tee /var/log/nginx/access.log
}
