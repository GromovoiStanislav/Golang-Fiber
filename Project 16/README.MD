## Fiber with MySQL (Gorm) Auth REST API example

### Setup:

```
go mod init fiber-example

go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/dgrijalva/jwt-go/v4
```

### Start the server:

```
go run .
```
