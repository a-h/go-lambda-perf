# go-lambda-perf

## Tasks

### cli-build

dir: hello-world-cli

```
go build
```

### cli-time

Uses zsh's built in time feature to list the time in miliseconds.

dir: hello-world-cli
requires: cli-build
env: TIMEFMT=%mE

```
zsh -c "time ./hello-world-cli"
```

### lambda-deploy

dir: cdk

```
cdk deploy --hotswap
```
