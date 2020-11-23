# Example Model Parser

```go
package models

type User struct {
        ID              uint    `db:"id"`
        Mail            string  `db:"mail"`
        Password        string  `db:"password"`
        Token           string  `db:"token"`
}

func (tbl *User) TableName() string {
        return "users"
}
```

output

```
*ast.File
        *ast.Ident
        *ast.GenDecl
                *ast.TypeSpec
                        *ast.Ident
                        *ast.StructType
                                *ast.FieldList
                                        *ast.Field
                                                *ast.Ident
                                                *ast.Ident
                                                *ast.BasicLit
                                        *ast.Field
                                                *ast.Ident
                                                *ast.Ident
                                                *ast.BasicLit
                                        *ast.Field
                                                *ast.Ident
                                                *ast.Ident
                                                *ast.BasicLit
                                        *ast.Field
                                                *ast.Ident
                                                *ast.Ident
                                                *ast.BasicLit
        *ast.FuncDecl
                *ast.FieldList
                        *ast.Field
                                *ast.Ident
                                *ast.StarExpr
                                        *ast.Ident
                *ast.Ident
                *ast.FuncType
                        *ast.FieldList
                        *ast.FieldList
                                *ast.Field
                                        *ast.Ident
                *ast.BlockStmt
                        *ast.ReturnStmt
                                *ast.BasicLit
```
