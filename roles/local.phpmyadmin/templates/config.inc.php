<?php
	for($i = 1; $i <= 3; $i++){
		$cfg["Servers"][$i]["connect_type"] = "tcp";
		$cfg["Servers"][$i]["host"] = "127.0.0.1";
		$cfg["Servers"][$i]["port"] = 33060 + $i;

		$cfg["Servers"][$i]["auth_type"] = "config";
		$cfg["Servers"][$i]["user"] = "{{ bot_user }}";
		$cfg["Servers"][$i]["password"] = "{{ bot_user }}";
	}
?>
