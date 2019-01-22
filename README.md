# go-webapp

# Requirements:
	1. golang
    2. mysql-server

# Installation:

### Download/install go:
	sudo tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin
	export GOPATH=$HOME/go

### Download mysql libs for go:
	go get -u github.com/go-sql-driver/mysql

### Navigate to ~/go/src
	git clone https://github.com/wescules/go-webapp.git
	cd go-webapp/migrations

### Initilize database
	go run migrations.go
	cd ..
    
### Run app
	go run main.go
