module bookstore-users-api

go 1.15

require bookstore-oauth-go v0.0.0

replace bookstore-oauth-go v0.0.0 => ./../bookstore-oauth-go

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/mercadolibre/golang-restclient v0.0.0-20170701022150-51958130a0a0 // indirect
	go.uber.org/zap v1.19.0
)
