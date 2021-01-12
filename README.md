# Reddit clone
A simple project emulate reddit's *vote function*

## Prerequisite
You can start it by 2 ways using **docker** or **without docker**.
* **Docker**
    * Install docker follow this article ([docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04))
    * *(optional)* You can install `docker-compose` if you want to run with `docker-compose`. ([docker-compose](https://docs.docker.com/compose/install/))
* **Without docker**
    * **Go:** I using `go 1.14` for the project. ([Install](https://golang.org/doc/install))
    * **React**: I using ReactJS to render frontend. So you need install it to run it. ([Install](https://reactjs.org/docs/create-a-new-react-app.html#recommended-toolchains))

## Installation
### Prepare
* Clone source code
```
$ git clone https://github.com/vietanhduong/reddit-clone.git
$ cd reddit-clone
```
* `.env` file:
``` .env
PORT=8080										# default port is 8080. You can override it.
SECRET_KEY=your-secret                          # This secret to generate JWT token
ADMIN_USERNAME=admin							# default username admin
ADMIN_PASSWORD=e10adc3949ba59abbe56e057f20f883e # default password md5(123456) 
```

### With Docker
```shell
$ docker build . -t reddit-clone 
$ docker run -p <port-out>:8080 --env-file=.env reddit-clone
```

### Without Docker
``` shell
# Dowload dependences
$ go mod download

# Install node module and build
$ cd client && yarn install && yarn build && cd ..

$ go build . -o main && ./main
```
### *Explanation*
In this project I combine frontend and backend in single container. In other words Go server will serve both side (client side and server side) by using `rice`. So `rice` will  convert static file to binary and embed it to main before `build`.

## Usage
* In this project I don't use database. I store data in memory and all data will lose when you shutdown server. You can implement database if you want by add connection to `repository` file.
* You can create new account *(need implement)*. When you start project there is only one account (admin) in here. Default username is `vietanhduong` and default password is `md5(vietanhdeptrai)`. You can override it by `.env` file.
* **Only admin role (`admin=true`) can create topic.**
* **You need login to vote topic.**
*  I store topics in memory and implement `priority queue` to store topics and count vote.
* API `/api/topics` will return first 10 topics have most vote. 

