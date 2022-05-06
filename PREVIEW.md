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
  "msg": null,
  "checksum": "162c563e7cef8f388435ab82085084a6003802e7",
  "checksum_debug": "open_failed:/path/to/favicon.ico",
  "client": "127.0.0.1",
  "server": "_",
  "host": "example.com",
  "referrer": "https://abc.example.com/",
  "referrer_host": "abc.example.com",
  "request_method": "GET",
  "request_uri": "/favicon.ico",
  "request_http_version": "1.1",
  "error_type": "open_failed",
  "error_details": "/path/to/favicon.ico"
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
  "checksum": "75e07e06eb2114694f3ebed7069ec861e553ec62",
  "checksum_debug": "fastcgi:Primary script unknown",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "upstream": "fastcgi://unix:/var/run/fpm.sock:",
  "upstream_host": "unix:",
  "request_method": "GET",
  "request_uri": "/login.php",
  "request_http_version": "1.1",
  "error_type": "fastcgi",
  "error_details": "Primary script unknown"
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
  "msg": null,
  "checksum": "e80aa6330977803f1b60a76249ba1886a8a84b96",
  "checksum_debug": "fastcgi:PHP message: PHP Warning:  file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "upstream": "fastcgi://unix:/var/run/fpm.sock:",
  "upstream_host": "unix:",
  "request_method": "GET",
  "request_uri": "/login.php",
  "request_http_version": "1.1",
  "error_type": "fastcgi",
  "error_details": "PHP message: PHP Warning:  file_exists(): open_basedir restriction in effect. File(/home/public_html/www/wp-content/themes/dynamic.css) is not within the allowed path(s): (/home/public_html:/usr/share/pear:/usr/share/php:/tmp:/usr/local/lib/php) in /home/public_html/public_html/wp-content/themes/includes/functions.php on line 238"
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
  "checksum": "11ffd7bfc6417dc3048454e771184ad6bd40844a",
  "checksum_debug": "access forbidden by rule",
  "client": "127.0.0.1",
  "server": "example.com",
  "host": "example.com",
  "request_method": "GET",
  "request_uri": "/.git/config",
  "request_http_version": "1.1",
  "error_type": "_"
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
  "checksum": "1b67e7926c55f397afa00659e2cdd01c4f5fee38",
  "checksum_debug": "this is exception might be happened",
  "error_type": "_"
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
  "checksum": "991b65b224110451fce1dea3769794234ddbf257",
  "checksum_debug": "ssl do handshake failed ssl error d ssl routines tls process client hello unsupported protocol while ssl handshaking",
  "client": "127.0.0.1",
  "server": "0.0.0.0:443",
  "error_type": "_"
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
  "checksum": "d6f4ed54991696a1973d19c5c04c3d4b4a14b4ef",
  "checksum_debug": "closed keepalive connection",
  "client": "127.0.0.1",
  "error_type": "_"
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
  "checksum": "c61ce49f75dad963fa96e97391155d0165442f02",
  "checksum_debug": "client closed connection while waiting for request",
  "client": "127.0.0.1",
  "server": "0.0.0.0:80",
  "error_type": "_"
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
  "msg": null,
  "checksum": "73e613b09048ab3cc222ffc791060392f9994eed",
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
  "error_type": "naxsi_exlog",
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
  "msg": null,
  "checksum": "b5e6e44b90feaeff959b01ce07cb862c07246092",
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
  "error_type": "naxsi_fmt",
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
  ]
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
  "checksum": "f5099d8be866addf727ef471a6e79679e3395991",
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
  "error_type": "_"
}
```

---

```txt
2021/02/03 15:16:17 [error] 6#6: *1 failed to load external Lua file "/usr/local/openresty/lualib/access_normal.lua": cannot open /usr/local/openresty/lualib/access_normal.lua: Permission denied, client: 192.168.88.93, server: _, request: "GET / HTTP/1.1", host: "192.168.88.220"
```

```json
{
  "time": "Wed, 03 Feb 2021 15:16:17 +0000",
  "level": "error",
  "pid": 6,
  "tid": 6,
  "cid": 1,
  "message": "failed to load external Lua file \"/usr/local/openresty/lualib/access_normal.lua\": cannot open /usr/local/openresty/lualib/access_normal.lua: Permission denied, client: 192.168.88.93, server: _, request: \"GET / HTTP/1.1\", host: \"192.168.88.220\"",
  "msg": "failed to load external lua file usr local openresty lualib access normal lua cannot open usr local openresty lualib access normal lua permission denied",
  "checksum": "92d4de28c03dbf595c1aaabacedcd88f4f293306",
  "checksum_debug": "failed to load external lua file usr local openresty lualib access normal lua cannot open usr local openresty lualib access normal lua permission denied",
  "client": "192.168.88.93",
  "server": "_",
  "host": "192.168.88.220",
  "request_method": "GET",
  "request_uri": "/",
  "request_http_version": "1.1",
  "error_type": "_"
}
```

---

```txt
2022/03/01 21:22:23 [info] 19#19: [ngx_pagespeed 1.13.35.2-0] Shutting down PageSpeed child
```

```json
{
  "time": "Tue, 01 Mar 2022 21:22:23 +0000",
  "level": "info",
  "pid": 19,
  "tid": 19,
  "message": "[ngx_pagespeed 1.13.35.2-0] Shutting down PageSpeed child",
  "msg": "ngx pagespeed shutting down pagespeed child",
  "checksum": "a3bdc753560b26e213e02f49b2e0c1a1235ca024",
  "checksum_debug": "ngx pagespeed shutting down pagespeed child",
  "error_type": "_"
}
```

---

```txt
2022/03/01 21:22:23 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
```

```json
{
  "time": "Tue, 01 Mar 2022 21:22:23 +0000",
  "level": "notice",
  "pid": 1,
  "tid": 1,
  "message": "getrlimit(RLIMIT_NOFILE): 1048576:1048576",
  "msg": "getrlimit rlimit nofile",
  "checksum": "f7d16cf02f2c32df17bb82ee79ed77239247a2ae",
  "checksum_debug": "getrlimit rlimit nofile",
  "error_type": "_"
}
```

---
