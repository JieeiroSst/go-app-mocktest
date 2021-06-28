# Win

  SET WEB_URL=localhost:4040
  SET WEB_MYSQL_DSN=root:1234@tcp(localhost:3306)/db?parseTime=True

# Linux

  export WEB_PORT=:4040
  export WEB_MYSQL_DSN=root:1234@tcp(localhost:3306)/db?parseTime=True
  
# Mockgen
  
  mockgen -source=./domain/article.go -destination=./integration_test/mocks_test.go -package=integration_test
