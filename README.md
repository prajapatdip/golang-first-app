
# Golang SDK application.

In this project I have created just simple customer details golang SDK application and also created docker image to package all the files in the container and to make it portable so it can run anywhere.


## Run Locally

Fetch the source code.

```
    git clone https://github.com/prajapatdip/golang-first-app.git
    cd golang-first-app
```

You can run this project by using two methods:

```
    1. By golang compiler.
    2. By building docker image and runnig it locally.
```

### 1. By golang compiler.

```
    cd app
    go build
    go ./app
```

### 2. By building docker image and runnig it locally.

```
    docker build --target build . -t customer-app
    docker build . -t customer-app
    docker run customer-app
```
