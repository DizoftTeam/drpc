# DRpcClient

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
