# Installation

## Prepare

For achieve better implementation consider following steps:

### Compile nginx

Increase `NGX_MAX_ERROR_STR` to larger value like 20480 on [ngx_log.h](https://github.com/nginx/nginx/blob/master/src/core/ngx_log.h) during compile.

### PHP-FPM

If you using PHP FPM, increase `log_limit` [log_limit](https://www.php.net/manual/en/install.fpm.configuration.php#log-limit) on fpm config to `1048576` for increase line wrapping.

## Usage

Just using `docker-compose`

```bash
docker-compose up -d
```
