# hello - test service

## init

`go mod init`

## build

`go build main.go`

## prepare deployment

`./get_envs.sh`

## deploy

### dev

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.dev`

### stg

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.stg`

### prod

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.prod`
