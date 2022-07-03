## Generate Swagger 

https://github.com/gofiber/swagger
https://github.com/swaggo/swag
https://github.com/swaggo/swag#declarative-comments-format

## Generate Doc Api

## Issue Swagger
### if swag init command not running
```sh
$ cd ~/go/bin
$ ls
```

### you should see "swag" then run
```sh
$ export PATH="/Users/XXX/go/bin:$PATH"
$ source ~/.zshrc
$ swag -v
```
### Step 1 check swag
```sh
$ cd ~/go/bin
$ ls
```
### Step 2 Export path in Project
```sh
export PATH="/Users/XXX/go/bin:$PATH" //in project
swag -v 
```

### Step 3  Generate the Swagger Specification.
```sh
swag init -g main.go --output docs/ 
```


