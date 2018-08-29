## Mock Directory Structure

- $GOPATH
  - src
    - connector-go.git 
    - connector-go-mock.git

## Initial Setup
```
$ ./prebuild.sh
```

## Run using local connector-go.git directory
```
$ ./local-1/run.sh
```
in case of errors with vendors directory missing make sure that ./prebuild.sh command is called.

## Run using github connector-go.git url
```
$ ./sample-1/run.sh
```
