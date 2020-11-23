# Go Quick API Development

MySQL and Fiber based quick api development kit

*I will continue to improve while working on the go in my free time.*

**NOTE: Currently using effective not available.**

# Create Models

Example model file

```go
package models

type User struct {
    ID       uint   `db:"id"`
    Mail     string `db:"mail"`
    Password string `db:"password"`
    Token    string `db:"token"`
}

func (tbl *User) TableName() string {
    return "<YourTableName>"
}
```

# q CLI Guides

Show help commands

```
go run q -h
```

## Generate repositories from models dir

Auto generated repositories from `./models/*` directory in models

```
go run q.go repo
```

Example model using

```go
new(models.<ModelName>).All() // get all data on table
new(models.<ModelName>).FindFromId(1) // get id = 1 data on table

// Todo repo functions
// .Where()
// .Update()
// .Delete()
```

# How to use

```
go mod download
air
```

# TODO

- [ ] Code generation from models
    - [ ] Repository template 
        - [x] All
        - [x] FindFromId
        - [ ] Update
        - [ ] Delete
        - [ ] Where Condition
- [ ] Add mail class
- [ ] Re-structured MVVM (repository) sturcutured
- [ ] Manage structured via sqlite
