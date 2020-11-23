// https://medium.com/justforfunc/understanding-go-programs-with-go-parser-c4e88a6edb87

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	modelPath    string = "models"
	templatePath string = "templates"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ParseModel struct {
	Path    string
	Name    string
	Table   string
	Columns []string
}

func GenerateRepositoryName(pm *ParseModel) string {
	return modelPath + "/" + strings.ToLower(pm.Name) + "_repo.go"
}

func ParseRepositoriesFromModelsDirectory(debug bool) bool {
	/*
		1. Read All models file
		2. Generate repositories
	*/
	var models []string
	err := filepath.Walk(modelPath, func(path string, info os.FileInfo, err error) error {
		if path != modelPath && !strings.Contains(path, "_repo") {
			models = append(models, path)

			if debug {
				fmt.Println("Find Model File: \t", path)
			}
		}

		return nil
	})

	check(err)

	// tokenizer AST
	parseModels := make([]*ParseModel, 0)
	for _, modelFile := range models {
		parseModel := new(ParseModel)

		parseModel.Path = modelFile

		dat, err := ioutil.ReadFile(modelFile)
		check(err)

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", string(dat), parser.AllErrors)
		check(err)

		// Parsing Model file
		for _, declare := range f.Decls {
			if typeDecl, ok := declare.(*ast.GenDecl); ok { // Get struct info
				findModelName := typeDecl.Specs[0].(*ast.TypeSpec).Name
				parseModel.Name = findModelName.Name

				fields := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List

				var fieldStrings []string
				for _, field := range fields {
					fieldStrings = append(fieldStrings, field.Names[0].Name)
				}

				parseModel.Columns = fieldStrings
			} else if funcDecl, ok := declare.(*ast.FuncDecl); ok { // Get tableName string from TableName() return stmt
				if funcDecl.Name.String() == "TableName" {
					if hasReturn, ok := funcDecl.Body.List[0].(*ast.ReturnStmt); ok {
						findTableName := string(dat[hasReturn.Results[0].Pos() : hasReturn.End()-2])
						parseModel.Table = findTableName
					}
				}
			}
		}

		parseModels = append(parseModels, parseModel)
		parseModel = nil //free mem
	}

	// Parse Repository
	for _, pm := range parseModels {
		if debug {
			fmt.Println("ModelName: \t\t", pm.Name)
			fmt.Println("TableName: \t\t", pm.Table)
			fmt.Println("Columns: \t\t", pm.Columns, "\n")
		}

		pmFileRead, err := ioutil.ReadFile(templatePath + "/repository.tpl")
		check(err)

		templatePm, err := template.New("ParseModel").Parse(string(pmFileRead))
		check(err)

		var buff bytes.Buffer

		err = templatePm.Execute(&buff, pm)
		check(err)

		f, err := os.Create(GenerateRepositoryName(pm))
		check(err)
		defer f.Close()

		f.Sync()

		w := bufio.NewWriter(f)
		_, errWrite := w.WriteString(buff.String())
		check(errWrite)

		w.Flush()
	}

	parseModels = nil //free mem
	return true
}

func main() {
	argsWithoutProg := os.Args[1:]

	if argsWithoutProg[0] == "repo" {
		ParseRepositoriesFromModelsDirectory(len(argsWithoutProg) == 2 && argsWithoutProg[1] == "dbg")
		fmt.Println("Generated repository files!")
	}
}
