{% for item in groups['all'] %}
{% if 'mysql_forward' in hostvars[item] %}
<?php
	$i++;
	$cfg["Servers"][$i]["host"] = "127.0.0.1";
	$cfg["Servers"][$i]["port"] = "{{ hostvars[item]['mysql_forward'] }}";
	$cfg["Servers"][$i]["auth_type"] = "config";
	$cfg["Servers"][$i]["user"] = "{{ bot_user }}";
	$cfg["Servers"][$i]["password"] = "{{ bot_user }}";
?>
{% endif %}
{% endfor %}
