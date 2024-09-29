# hello - test app

## init

`go mod init hello && go mod tidy && go get .`

## build

`go build -o hello .`

## run

`PORT=8000 ENV=local ./hello`

## deploy locally

### prepare

`./get_envs.sh`

### dev

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.dev`

### stg

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.stg`

### prod

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.prod`
