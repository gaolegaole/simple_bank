database DSL <https://dbdiagram.io/d/62e48905f31da965e8437d56>
golang-migrate/migrate <https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>
### windows makefile
1. 下载mingw64
2. 设置path
3. 修改ming32-make.exe 为make.exe
### 生成migrate
migrate -path db/migrations create -dir=db/migration -ext=sql -seq=1 init

### sqlc
go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.13.0
latest(v1.14.0) 需要go1.18
### gomock
go install github.com/golang/mock/mockgen@v1.6.0
mockgen -package mockdb -destination db/mock/store.go github.com/gaolegaole/simple_bank/db/sqlc Store