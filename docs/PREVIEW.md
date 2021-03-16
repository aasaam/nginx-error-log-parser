# Preview

This is sample of error log and parsed result

```txt
2020/01/02 03:04:05 [error] 7#8: *851624 open() "/path/to/favicon.ico" failed (2: No such file or directory), client: 127.0.0.1, server: _, request: "GET /favicon.ico HTTP/1.1", host: "example.com", referrer: "https://abc.example.com/"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "error",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "open() \"/path/to/favicon.ico\" failed (2: No such file or directory), client: 127.0.0.1, server: _, request: \"GET /favicon.ico HTTP/1.1\", host: \"example.com\", referrer: \"https://abc.example.com/\"",
  "msg": "",
  "checksum": "6e8951cd8f2c66b76daf2966d9f230f9",
  "checksum_debug": "open_failed:/path/to/favicon.ico",
  "client": "127.0.0.1",
  "server": "_",
  "host": "example.com",
  "upstream": "",
  "upstream_host": "",
  "referrer": "https://abc.example.com/",
  "referrer_host": "abc.example.com",
  "request_method": "GET",
  "request_uri": "/favicon.ico",
  "request_http_version": "1.1",
  "error_type": "open_failed",
  "error_details": "/path/to/favicon.ico",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [error] 7#8: *851624 FastCGI sent in stderr: "Primary script unknown" while reading response header from upstream, client: 127.0.0.1, server: example.com, request: "GET /login.php HTTP/1.1", upstream: "fastcgi://unix:/var/run/fpm.sock:", host: "example.com"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "error",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "FastCGI sent in stderr: \"Primary script unknown\" while reading response header from upstream, client: 127.0.0.1, server: example.com, request: \"GET /login.php HTTP/1.1\", upstream: \"fastcgi://unix:/var/run/fpm.sock:\", host: \"example.com\"",
  "msg": "while reading response header from upstream",
  "checksum": "9deb25b8b553f8183ce142146f927f7d",
  "checksum_debug": "fastcgi_error:Primary script unknown",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "upstream": "fastcgi://unix:/var/run/fpm.sock:",
  "upstream_host": "unix:",
  "referrer": "",
  "referrer_host": "",
  "request_method": "GET",
  "request_uri": "/login.php",
  "request_http_version": "1.1",
  "error_type": "fastcgi_error",
  "error_details": "Primary script unknown",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [error] 7#8: *851624 FastCGI sent in stderr: "PHP message: PHP Warning: file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238", client: 127.0.0.1, server: example.com, request: "GET /login.php HTTP/1.1", upstream: "fastcgi://unix:/var/run/fpm.sock:", host: "example.com"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "error",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "FastCGI sent in stderr: \"PHP message: PHP Warning:  file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238\", client: 127.0.0.1, server: example.com, request: \"GET /login.php HTTP/1.1\", upstream: \"fastcgi://unix:/var/run/fpm.sock:\", host: \"example.com\"",
  "msg": "",
  "checksum": "1492348079ec45554d5b945a86198edf",
  "checksum_debug": "fastcgi_error:PHP message: PHP Warning:  file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "upstream": "fastcgi://unix:/var/run/fpm.sock:",
  "upstream_host": "unix:",
  "referrer": "",
  "referrer_host": "",
  "request_method": "GET",
  "request_uri": "/login.php",
  "request_http_version": "1.1",
  "error_type": "fastcgi_error",
  "error_details": "PHP message: PHP Warning:  file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [error] 7#8: *851624 access forbidden by rule, client: 127.0.0.1, server: example.com, request: "GET /.git/config HTTP/1.1", host: "example.com"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "error",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "access forbidden by rule, client: 127.0.0.1, server: example.com, request: \"GET /.git/config HTTP/1.1\", host: \"example.com\"",
  "msg": "access forbidden by rule",
  "checksum": "7f2de4fc0ac7f5e58d16f31f04250e9f",
  "checksum_debug": "access forbidden by rule",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "GET",
  "request_uri": "/.git/config",
  "request_http_version": "1.1",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 this is exception might be happened
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "this is exception might be happened",
  "msg": "this is exception might be happened",
  "checksum": "7b4f0887031ba7a59b6091367e14378c",
  "checksum_debug": "this is exception might be happened",
  "client": "",
  "server": "",
  "host": "",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "",
  "request_uri": "",
  "request_http_version": "",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 SSL_do_handshake() failed (SSL: error:1417D102:SSL routines:tls_process_client_hello:unsupported protocol) while SSL handshaking, client: 127.0.0.1, server: 0.0.0.0:443
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "SSL_do_handshake() failed (SSL: error:1417D102:SSL routines:tls_process_client_hello:unsupported protocol) while SSL handshaking, client: 127.0.0.1, server: 0.0.0.0:443",
  "msg": "ssl do handshake failed ssl error d ssl routines tls process client hello unsupported protocol while ssl handshaking",
  "checksum": "affec6346bb221ca0c8fc03fef4e67bc",
  "checksum_debug": "ssl do handshake failed ssl error d ssl routines tls process client hello unsupported protocol while ssl handshaking",
  "client": "127.0.0.1",
  "server": "0.0.0.0:443",
  "host": "",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "",
  "request_uri": "",
  "request_http_version": "",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 client 127.0.0.1 closed keepalive connection
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "client 127.0.0.1 closed keepalive connection",
  "msg": "closed keepalive connection",
  "checksum": "69a90b22fc459bd77904145a7bd6bc43",
  "checksum_debug": "closed keepalive connection",
  "client": "127.0.0.1",
  "server": "",
  "host": "",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "",
  "request_uri": "",
  "request_http_version": "",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 client closed connection while waiting for request, client: 127.0.0.1, server: 0.0.0.0:80
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "client closed connection while waiting for request, client: 127.0.0.1, server: 0.0.0.0:80",
  "msg": "client closed connection while waiting for request",
  "checksum": "49283931bd300c3d9ec2274df08c8126",
  "checksum_debug": "client closed connection while waiting for request",
  "client": "127.0.0.1",
  "server": "0.0.0.0:80",
  "host": "",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "",
  "request_uri": "",
  "request_http_version": "",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 NAXSI_EXLOG: ip=127.0.0.1&server=sub.example.com&uri=%2Findex.php&id=1013&zone=ARGS&var_name=sid&content=147%27%5B0%5D, client: 127.0.0.1, server: 0.0.0.0, request: "GET /index.php HTTP/2.0", upstream: "http://127.0.0.1:80/index.php", host: "example.com", referrer: "https://www.example.com/page.html"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "NAXSI_EXLOG: ip=127.0.0.1&server=sub.example.com&uri=%2Findex.php&id=1013&zone=ARGS&var_name=sid&content=147%27%5B0%5D, client: 127.0.0.1, server: 0.0.0.0, request: \"GET /index.php HTTP/2.0\", upstream: \"http://127.0.0.1:80/index.php\", host: \"example.com\", referrer: \"https://www.example.com/page.html\"",
  "msg": "",
  "checksum": "4a74d64e8d2e65640b238d7a3424eaa3",
  "checksum_debug": "naxsi_exlog:sub.example.com:/index.php:1013:ARGS:sid:147'[0]",
  "client": "127.0.0.1",
  "server": "0.0.0.0",
  "host": "example.com",
  "upstream": "http://127.0.0.1:80/index.php",
  "upstream_host": "127.0.0.1:80",
  "referrer": "https://www.example.com/page.html",
  "referrer_host": "www.example.com",
  "request_method": "GET",
  "request_uri": "/index.php",
  "request_http_version": "2.0",
  "error_type": "",
  "error_details": "",
  "naxsi": "exlog",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "127.0.0.1",
  "naxsi_exlog_server": "sub.example.com",
  "naxsi_exlog_uri": "/index.php",
  "naxsi_exlog_id": "1013",
  "naxsi_exlog_zone": "ARGS",
  "naxsi_exlog_var_name": "sid",
  "naxsi_exlog_content": "147'[0]"
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 NAXSI_FMT: ip=127.0.0.1&server=sub.example.com&uri=/index.php&learning=0&vers=0.56&total_processed=1024&total_blocked=128&block=1&cscore0=$SQL&score0=4&cscore1=$XSS&score1=8&zone0=ARGS&id0=1013&var_name0=sid, client: 127.0.0.1, server: 0.0.0.0, request: "GET /index.php HTTP/2.0", upstream: "http://127.0.0.1:80/index.php", host: "example.com", referrer: "https://www.example.com/page.html"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "NAXSI_FMT: ip=127.0.0.1&server=sub.example.com&uri=/index.php&learning=0&vers=0.56&total_processed=1024&total_blocked=128&block=1&cscore0=$SQL&score0=4&cscore1=$XSS&score1=8&zone0=ARGS&id0=1013&var_name0=sid, client: 127.0.0.1, server: 0.0.0.0, request: \"GET /index.php HTTP/2.0\", upstream: \"http://127.0.0.1:80/index.php\", host: \"example.com\", referrer: \"https://www.example.com/page.html\"",
  "msg": "",
  "checksum": "5ae1cded04479557afa80f26cdd6d082",
  "checksum_debug": "naxsi_fmt:sub.example.com:/index.php:1:$SQL:4:1013:sid:ARGS:$XSS:8",
  "client": "127.0.0.1",
  "server": "0.0.0.0",
  "host": "example.com",
  "upstream": "http://127.0.0.1:80/index.php",
  "upstream_host": "127.0.0.1:80",
  "referrer": "https://www.example.com/page.html",
  "referrer_host": "www.example.com",
  "request_method": "GET",
  "request_uri": "/index.php",
  "request_http_version": "2.0",
  "error_type": "",
  "error_details": "",
  "naxsi": "fmt",
  "naxsi_fmt_ip": "127.0.0.1",
  "naxsi_fmt_server": "sub.example.com",
  "naxsi_fmt_uri": "/index.php",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "0.56",
  "naxsi_fmt_block": true,
  "naxsi_fmt_total_processed": 1024,
  "naxsi_fmt_total_blocked": 128,
  "naxsi_fmt_items": [
    {
      "zone": "ARGS",
      "id": "1013",
      "var_name": "sid",
      "cscore": "$SQL",
      "score": "4"
    },
    {
      "zone": "",
      "id": "",
      "var_name": "",
      "cscore": "$XSS",
      "score": "8"
    }
  ],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/01/02 03:04:05 [info] 7#8: *851624 client canceled stream 1 while sending to client, client: 127.0.0.1, server: 0.0.0.0, request: "GET /index.php HTTP/2.0", upstream: "http://127.0.0.1:80/index.php", host: "example.com", referrer: "https://www.example.com/page.html"
```

```json
{
  "time": "Thu, 02 Jan 2020 03:04:05 +0000",
  "level": "info",
  "pid": 7,
  "tid": 8,
  "cid": 851624,
  "message": "client canceled stream 1 while sending to client, client: 127.0.0.1, server: 0.0.0.0, request: \"GET /index.php HTTP/2.0\", upstream: \"http://127.0.0.1:80/index.php\", host: \"example.com\", referrer: \"https://www.example.com/page.html\"",
  "msg": "client canceled stream while sending to client",
  "checksum": "48a853990a333c7ddb7df0015e329595",
  "checksum_debug": "client canceled stream while sending to client",
  "client": "127.0.0.1",
  "server": "0.0.0.0",
  "host": "example.com",
  "upstream": "http://127.0.0.1:80/index.php",
  "upstream_host": "127.0.0.1:80",
  "referrer": "https://www.example.com/page.html",
  "referrer_host": "www.example.com",
  "request_method": "GET",
  "request_uri": "/index.php",
  "request_http_version": "2.0",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---

```txt
2020/12/31 15:34:35 [error] 6#6: *1 failed to load external Lua file "/usr/local/openresty/lualib/access_normal.lua": cannot open /usr/local/openresty/lualib/access_normal.lua: Permission denied, client: 192.168.88.93, server: _, request: "GET / HTTP/1.1", host: "192.168.88.220"
```

```json
{
  "time": "Thu, 31 Dec 2020 15:34:35 +0000",
  "level": "error",
  "pid": 6,
  "tid": 6,
  "cid": 1,
  "message": "failed to load external Lua file \"/usr/local/openresty/lualib/access_normal.lua\": cannot open /usr/local/openresty/lualib/access_normal.lua: Permission denied, client: 192.168.88.93, server: _, request: \"GET / HTTP/1.1\", host: \"192.168.88.220\"",
  "msg": "failed to load external lua file usr local openresty lualib access normal lua cannot open usr local openresty lualib access normal lua permission denied",
  "checksum": "dfd3a923ba8a2a753d2f9ec7b9dffa63",
  "checksum_debug": "failed to load external lua file usr local openresty lualib access normal lua cannot open usr local openresty lualib access normal lua permission denied",
  "client": "192.168.88.93",
  "server": "_",
  "host": "192.168.88.220",
  "upstream": "",
  "upstream_host": "",
  "referrer": "",
  "referrer_host": "",
  "request_method": "GET",
  "request_uri": "/",
  "request_http_version": "1.1",
  "error_type": "",
  "error_details": "",
  "naxsi": "",
  "naxsi_fmt_ip": "",
  "naxsi_fmt_server": "",
  "naxsi_fmt_uri": "",
  "naxsi_fmt_learning": false,
  "naxsi_fmt_vers": "",
  "naxsi_fmt_block": false,
  "naxsi_fmt_total_processed": 0,
  "naxsi_fmt_total_blocked": 0,
  "naxsi_fmt_items": [],
  "naxsi_exlog_ip": "",
  "naxsi_exlog_server": "",
  "naxsi_exlog_uri": "",
  "naxsi_exlog_id": "",
  "naxsi_exlog_zone": "",
  "naxsi_exlog_var_name": "",
  "naxsi_exlog_content": ""
}
```

---
