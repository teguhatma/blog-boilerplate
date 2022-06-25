# Blog-Boilerplate
---

You need to install these packages first.

1. Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
    ```bash
    $ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
    ```
2. Install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
    #### Mac OS
    ```bash
    brew install sqlc
    ```
    #### Ubuntu
    ```bash
    sudo snap install sqlc
    ```
    #### Go Install
    ##### Go >= 1.17
    ```bash
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```
    ##### Go < 1.17
    ```bash
    go get github.com/kyleconroy/sqlc/cmd/sqlc
    ```
3. Install docker by donwload the binary [here](https://www.docker.com/get-started/)

### Run the server
1. Run docker and do:
   ```bash
   docker-compose up
   ```
   or
   ```bash
   docker-compose up -d
   ```
2. Create schema
   ```bash
   make createschema
   ```
3. Make migration
   ```bash
   make migrateup
   ```
4. Run your server
   ```bash
   make server
   ```

the address -> http://localhost:8030/api/v1