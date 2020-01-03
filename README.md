<!-- 
Autogenerated by Typical-Go from template 'README.tmpl'. 
DO NOT EDIT THIS FILE!
-->

# Typical REST Server

Example of typical and scalable RESTful API Server for Go

### Usage

| Bash Snippet | Usage |
|---|---|
|`typical-rest-server`|Run the application|
|`typical-rest-server route`|Print available API Routes|

### Configuration

| Name | Type | Default | Required |
|---|---|---|:---:|
|APP_ADDRESS|string|:8089|Yes|
|PG_DBNAME|string||Yes|
|PG_HOST|string|localhost||
|PG_PASSWORD|string|pgpass|Yes|
|PG_PORT|int|5432||
|PG_USER|string|postgres|Yes|
|REDIS_DB|int|0||
|REDIS_DIAL_TIMEOUT|Duration|5s|Yes|
|REDIS_HOST|string|localhost|Yes|
|REDIS_IDLE_CHECK_FREQUENCY|Duration|1m|Yes|
|REDIS_IDLE_TIMEOUT|Duration|5m|Yes|
|REDIS_MAX_CONN_AGE|Duration|30m|Yes|
|REDIS_PASSWORD|string|redispass||
|REDIS_POOL_SIZE|int|20|Yes|
|REDIS_PORT|string|6379|Yes|
|REDIS_READ_WRITE_TIMEOUT|Duration|3s|Yes|
|SERVER_DEBUG|bool|false||

----

## Development Guide

### Prerequisite

Install [Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)

### Build & Run

Use `./typicalw run` to build and run the project.

### Testing

Use `./typicalw test` to test the project.

### Release the distribution

Use `./typicalw release` to make the release. [Learn More](https://typical-go.github.io/learn-more/build-tool/release-distribution.html)

### Others
| Bash Snippet | Usage |
|---|---|
|`./typicalw docker`|Docker utility|
|`./typicalw docker compose`|Generate docker-compose.yaml|
|`./typicalw docker up`|Spin up docker containers according docker-compose|
|`./typicalw docker down`|Take down all docker containers according docker-compose|
|`./typicalw docker wipe`|Kill all running docker container|
|`./typicalw readme`|Generate README Documentation|
|`./typicalw rails`|Rails-like generation|
|`./typicalw rails scaffold`|Generate CRUD API|
|`./typicalw rails repository`|Generate Repository from tablename|
|`./typicalw redis`|Redis Tool|
|`./typicalw redis console`|Redis Interactive|
|`./typicalw postgres`|Postgres Database Tool|
|`./typicalw postgres create`|Create New Database|
|`./typicalw postgres drop`|Drop Database|
|`./typicalw postgres migrate`|Migrate Database|
|`./typicalw postgres rollback`|Rollback Database|
|`./typicalw postgres seed`|Data seeding|
|`./typicalw postgres reset`|Reset Database|
|`./typicalw postgres console`|PostgreSQL Interactive|
