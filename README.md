# inventory-system

## Migrate

migrate -database "postgres://postgres:postgres@localhost:5435/postgres?sslmode=disable" -path db/migrations up

## build flow

1. goa gen

goa gen inventory-system/api/svc/design  

2. goa example

goa example inventory-system/api/svc/design  

3. goa build (when modify sth, just do this and restart app)

go build ./cmd/inventory_system