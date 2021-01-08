package gen_route

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/imports"
	"html/template"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

const matchCommentRequestMap = "@RequestMap\\(([a-z]*=\"[\\S\\s]*\").*\\)"

type MethodRef struct {
	Name       string
	MethodType string
	Path       string
}

type ControllerRef struct {
	Name        string
	RequestPath string
	PackageName string
	Methods     []MethodRef
}

func Template(controllers []ControllerRef) []byte {
	helperFuncs := template.FuncMap{
		"makeCtrVarName": func(str string) string {
			str = strings.ToLower(string(str[0])) + str[1:]
			return str
		},
		"firstLower":       FirstLower,
		"makeGroupVarName": MakeGroupVarName,
	}
	for _, v := range controllers {
		fmt.Println(v.Methods)
	}
	temp := `package gen_build
import (
	
	"github.com/gin-gonic/gin"
)
//This code is automatically generated for AmountRoute. Please do not change it
func AmountRoute(router *gin.Engine ) *gin.Engine {
	{{range $i, $v := .Controllers}}
	var {{$v.Name | makeCtrVarName}} = {{$v.PackageName}}.{{$v.Name}}{}
	{{$v.RequestPath | makeGroupVarName}}:=router.Group("{{$v.RequestPath}}")
	{ 
		{{range $e, $m := $v.Methods}}
		{{$v.RequestPath | makeGroupVarName}}.{{$m.MethodType}}("{{$m.Path}}",{{$v.Name | makeCtrVarName}}.{{$m.Name}})
		{{end}}	
	}
	{{end}}

	return router
}`
	out := bytes.NewBufferString("")
	// Parse the template and pass it the helper functions
	t := template.Must(template.New("go_route_gen.tmpl").Funcs(helperFuncs).Parse(temp))
	// Execute the template and pass it the metadata we collected before
	t.Execute(out, map[string][]ControllerRef{"Controllers": controllers})
	return out.Bytes()
}

func FirstLower(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}
func MakeGroupVarName(str string) string {
	str = strings.Trim(str, "/")
	str = strings.ToLower(string(str[0])) + str[1:] + "Group"

	return str
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPath string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {

		if fi.IsDir() {
			dirs = append(dirs, dirPath+PthSep+fi.Name())
			newfiles, _ := getAllFiles(dirPath + PthSep + fi.Name())
			files = append(files, newfiles...)
		} else {
			files = append(files, dirPath+PthSep+fi.Name())
		}

	}

	return files, nil
}

func Gen(filePath string) {
	var file []string
	file, err := getAllFiles(filePath)
	fmt.Println(file)
	return
	ctrColl := make([]ControllerRef, 0)
	for _, v := range file {
		controllers := Parse(v)
		for _, v := range controllers {
			if len(v.Methods) < 1 {
				continue
			}
			ctrColl = append(ctrColl, v)
		}
	}

	res := Template(ctrColl)
	formattedCode, err := imports.Process("./struct_app/gen_build/route.go", res, &imports.Options{Comments: true})
	if err != nil {
		fmt.Printf("cannot format source code, might be an error in template: %s\n", err)
	}
	ioutil.WriteFile("./struct_app/gen_build/route.go", formattedCode, 0777)
}
func Parse(file string) []ControllerRef {
	controllers := make([]ControllerRef, 0)

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	//ast.Print(fset, f)
	//return controllers
	/***
	解析Controller 属性
	*/

	for _, v := range f.Decls {
		fn, ok := v.(*ast.GenDecl)
		if !ok {
			continue
		}
		if fn.Doc == nil {
			continue
		}
		fnDoc := fn.Doc.List
		structComment := make([]string, 0)
		controllerRefItem := ControllerRef{}
		for _, v := range fnDoc {
			if len(structComment) == 2 {
				break
			}
			if v.Text == "//@Controller" {
				structComment = append(structComment, v.Text)
				continue
			}
			l, _ := regexp.MatchString(matchCommentRequestMap, v.Text)
			if l {
				controllerRefItem.RequestPath = parseRequestMap(v.Text).Path.Value
				structComment = append(structComment, v.Text)
				continue
			}
		}
		if len(structComment) == 2 {
			controllerRefItem.Name = fn.Specs[0].(*ast.TypeSpec).Name.Obj.Name
			controllerRefItem.PackageName = f.Name.Name
			if ('A' <= controllerRefItem.Name[0] && controllerRefItem.Name[0] <= 'Z') == false {
				continue
			}
			controllers = append(controllers, controllerRefItem)
		}

	}
	for _, v := range f.Decls {
		fn, ok := v.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if fn.Recv != nil {
			structName := fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
			methodComment := fn.Doc.Text()
			for k, c := range controllers {
				if structName == c.Name {
					u := fn.Name.Name[0]
					if ('A' <= u && u <= 'Z') == false {
						continue
					}
					l, _ := regexp.MatchString(matchCommentRequestMap, methodComment)
					if !l {
						continue
					}
					methodRef := MethodRef{}
					r := parseRequestMap(methodComment)
					methodRef.Name = fn.Name.Name
					if r.Path.Value != "" {
						methodRef.Path = r.Path.Value
					} else {
						methodRef.Path = fn.Name.Name
					}
					methodRef.MethodType = r.Method.Value

					controllers[k].Methods = append(c.Methods, methodRef)
				}
			}
		}
	}

	return controllers
}

type RequestMapValue struct {
	Path   RequestMapItem
	Method RequestMapItem
	Group  RequestMapItem
}
type RequestMapItem struct {
	Key   string
	Value string
}

func parseRequestMap(text string) RequestMapValue {
	//eg @RequestMap("method"="post" ...) max haven ten key=>value
	res := RequestMapValue{}
	t := reflect.TypeOf(res)
	st := regexp.MustCompile("([a-z]*)=([^,)]*)").FindAllString(text, t.NumField())
	for _, v := range st {
		item := RequestMapItem{}
		value := strings.Split(v, "=")
		item.Key = value[0]
		item.Value = strings.Trim(value[1], "\"")
		if value[0] == "method" {
			res.Method = item
			res.Method.Value = strings.ToUpper(item.Value)
		}
		if value[0] == "path" {
			res.Path = item
		}
		if value[0] == "group" {
			res.Group = item
		}
	}

	return res
}
