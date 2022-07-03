Swagger 

swag init command not running
$ cd ~/go/bin
$ ls
you should see "swag" then run
$ export PATH="/Users/XXX/go/bin:$PATH"
$ source ~/.zshrc
$ swag -v

Step 1 check swag
$ cd ~/go/bin
$ ls

Step 2 Export path in Project
export PATH="/Users/XXX/go/bin:$PATH" //in project
swag -v 

Step 3  Generate the Swagger Specification.
swag init -g main.go --output docs/ 

swagger Authentication for apis
https://github.com/swaggo/gin-swagger/issues/90

Generate Swagger 

https://medium.com/geekculture/tutorial-generate-swagger-specification-and-swaggerui-for-go-fiber-web-framework-6c787d1672de

https://github.com/gofiber/swagger
https://github.com/swaggo/swag
https://github.com/swaggo/swag#declarative-comments-format