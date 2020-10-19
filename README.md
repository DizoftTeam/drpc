# DRpcClient

Command line client for json rpc

## Installation

Coming soon

## Example

```bash
drpcclient -url='http://localhost:8089' -method='version' -params='{}' -id=1
```

### TODO

* Отправка запроса с заголовками

```bash
drpcclient request 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```

* Отправка запроса и получение заголовков

> -v - verbose

```bash
drpcclient request -v 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```

* Генерация `curl` запроса

```bash
drpcclient curl 1 http://localhost -headers='{"X-User-Token": "<guid>"}' -method=users.login -params='{"login": "user", "password": "admin"}'
```
