# gorm-performance

This is the repository for performance test of gorm.

##Requirement
Docker Desktop

※The maintainer use codespaces.

## Local Setup

```
# migration
$ make migrate-up

# Execute measurement
$ cd app
$ go run main.go insert
```


## Variation
"SkipDefaultTransaction" and "PrepareStmt" are the settings of gorm.Config.
```
$ cd app
$ go run main.go insert -r ${REPEAT COUNT} -s ${boolean: SkipDefaultTransaction} -p ${boolean: PrepareStmt}
```