# DRpcClient

[![Issues](https://img.shields.io/github/issues/DizoftTeam/drpc)](https://github.com/DizoftTeam/drpc/issues)
[![Forks](https://img.shields.io/github/forks/DizoftTeam/drpc)](https://github.com/DizoftTeam/drpc/network/members)
[![Stars](https://img.shields.io/github/stars/DizoftTeam/drpc)](https://github.com/DizoftTeam/drpc/stargazers)
[![License](https://img.shields.io/github/license/DizoftTeam/drpc)](https://github.com/DizoftTeam/drpc/blob/master/LICENSE)
[![AUR Votes](https://img.shields.io/aur/version/drpc)](https://img.shields.io/aur/votes/drpc?label=AUR%20votes)
[![AUR Version](https://img.shields.io/aur/version/drpc)](https://aur.archlinux.org/packages/drpc/)

Command line client for json rpc

## Installation

### ArchLinux

You can install it from [AUR](https://aur.archlinux.org/packages/drpc/)

### Manual (Debian like. Redhat)

Requirements:

* [Go](https://golang.org/)

```bash
git clone https://github.com/DizoftTeam/drpc.git
cd drpc
go build -o drpc main.go
# For all users
sudo ln -s /home/<user>/path/to/drpc /usr/local/bin/drpc
```

## Example

### 1.3.0

* Add file config support

Example of config

> Note: in this config we are using yaml override cause it usability

```yaml
_notify: &notify
  name: _override_name
  url: http://api.gvozdika.dizoft.ru/rpc
  method: _override_method
  params: { }
  headers: { }

_base: &base
  <<: *notify
  id: 1

requests:
  # Авторизация
  - <<: *base
    name: login
    method: v1.users.login
    params:
      login: test
      password: test

  # Выход из системы
  - <<: *base
    name: logout
    method: v1.users.logout
    params:
      token: $token # Require from cmd
```

Syntax is

```bash
drpc -file=DrpcFileConfig.yaml some_method_name [someParamName=someParamValue]
```

#### $token

It means that you should provide this argument in command line!

```bash
drpc -file=drpc.yaml logout token=SomeTokenHere
```

### 1.2.0

* Show response headers by `-v` or `-verbose`
* Show spent time for request
* Fixed empty response of nil (server is down as example)

```bash
drpc request 1 http://localhost -method=v1.refs.list -v
```

Example of response

```
--->
{}
---
{
  "id": ,
  "jsonrpc": "2.0",
  "method": "v1.refs.list",
  "params": {}
}
<---
Time: 61ms

[Connection] keep-alive
[X-Powered-By] PHP/7.4.12
[Access-Control-Allow-Credentials] true
[X-Debug-Tag] 5fb27bdf00a24
[Vary] Accept-Encoding
[Server] nginx/1.18.0
[Access-Control-Expose-Headers] 
[X-Debug-Duration] 44
[X-Debug-Link] /debug/default/view?tag=5fb27bdf00a24
[Date] Mon, 16 Nov 2020 13:17:19 GMT
[Content-Type] application/json; charset=UTF-8
---
{
  "id": 1,
  "jsonrpc": "2.0",
  "result": []
}
```

### 1.1.0

* Send request with headers

```bash
drpc request 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```

### 1.0.0

```bash
drpc -url='http://localhost:8089' -method='version' -params='{}' -id=1
```

### TODO

* Отправка запроса и получение заголовков

> -v - verbose

```bash
drpcclient request -v 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```

* Генерация `curl` запроса

```bash
drpcclient curl 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```
