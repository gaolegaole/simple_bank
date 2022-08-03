# https://www.ruanyifeng.com/blog/2015/02/make.html
# <target> : <prerequisites> 
# [tab]  <commands>
# 前置条件可以是一个make命令，会自动执行生成前置条件的make

# result.txt: source.txt
# 	cp source.txt result.txt
# source.txt:
# 	echo "this is the source" > source.txt

postgres:
	docker run --name pg12 -p 5432:5432 -e POSTGRES_PASSWORD=123123 -d postgres:12-alpine
createdb:
	docker exec -it pg12 createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it pg12 dropdb --username=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123123@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://postgres:123123@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://postgres:123123@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://postgres:123123@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
# 生成model CRUD
sqlc:
	docker run --rm -v E:\code\simple_bank:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gaolegaole/simple_bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock