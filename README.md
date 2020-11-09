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
