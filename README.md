# fzu-bilibili

## [接口文档](https://documenter.getpostman.com/view/18742402/UVkvHXpc)

## config.yaml
```
server:
  port : 8000
  version: 1.0
  secretKey: something-very-secret
  adminJwtSecret: admin-secret

sql:
  host : 127.0.0.1
  port : 3306
  username : root
  password : 123456
  database : GOLANG

redis:
  Addr : localhost:6379
  password : 123456
  DB : 0

OSS:
  EndPoint :
  AccessKey :
  SecretKey :
  Bucket :
  Region :
  ```
