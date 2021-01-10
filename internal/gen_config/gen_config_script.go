package gen_config

import (
	"bytes"
	"fmt"
	"gen/internal/helper"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/imports"
	"html/template"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type configureStructFields struct {
	FiledName string
	FiledType string
	Tag string


}
type configureStruct struct {

	StructName string
	YamlKey  string
	AttrName string
	StructFields []configureStructFields
	PackageName string
}
var allFiles []string

const matchCommentConfigStruct ="@ConfigStruct\\((key=\"[\\S\\s]*\")?\\)"
func GenConfig()  {
	allFiles,_=helper.GetAllFiles(".")

	parseFn()
}

type function struct {
	FunctionName string
	ReturnName string
	ReturnAttr string
}
func parseFn()  {
	fu:=make([]function,0)
	configureStructList:=parse(allFiles)
	for _,file:=range allFiles{
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset,file, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		funs:=make([]function,0)
		if f.Comments==nil{
			continue
		}
		if  !strings.Contains(f.Comments[0].Text(),"+build wireinject") {
			continue
		}


		for _, v := range f.Decls {
			fn, ok := v.(*ast.FuncDecl)
			if !ok {
				continue
			}

			if fn.Type.Results ==nil || len(fn.Type.Results.List)!=1 {
				continue
			}
			if fn.Doc==nil {
				continue
			}
			for _,d:= range fn.Doc.List {
				if d.Text!="//@ConfigEntity"{
					continue
				}else {
					fun:=function{}
					fun.ReturnName=fn.Type.Results.List[0].Type.(*ast.Ident).Name
					fun.FunctionName=fn.Name.Name
					funs=append(funs,fun)
					break
				}
			}
		}


		for _,c:=range configureStructList{
			for _,f:=range funs{
				if c.StructName==f.ReturnName {
					f.ReturnAttr=c.AttrName
					fu=append(fu,f)
				}
			}

		}
	}
	templateFunc(configureStructList,fu)

}
func templateFunc(configStructs []configureStruct,fns []function)  {
	helperFuncs := template.FuncMap{
		"raw": func(str string) template.HTML {

			return template.HTML(str)
		},
	}
	temp:=`
//Code generated by Gen. DO NOT EDIT.
//Code generated by Gen. DO NOT EDIT.
//Code generated by Gen. DO NOT EDIT.
package {{.Package}}
	type Manger struct {
		{{range $i, $v := .ConfigStructs}}
			{{if ne $v.YamlKey "###"}} 
			 {{if ne $v.PackageName $.Package }} 
			 	{{$v.AttrName}} {{$v.PackageName}}.{{$v.StructName}}  {{$v.YamlKey | raw}}
			 {{else}}
				{{$v.AttrName}} {{$v.StructName}} {{$v.YamlKey | raw}}
			 {{end}}
			{{end}}
		{{end}}
	}
    
	var CfgManger *Manger
	func AmountConfig() (*Manger, error) {
		conf, err := ioutil.ReadFile("./config/application.yaml")
		if err != nil {
			panic(err)
		}
		sys := &Manger{}
		err = yaml.Unmarshal(conf, sys)
		CfgManger = sys
		return sys, err
	}
	{{range $i, $v := .Fns}}
	func {{$v.FunctionName}} (){{$v.ReturnName}}  {
		return CfgManger.{{$v.ReturnAttr}}
	}
	{{end}}
`

	out := bytes.NewBufferString("")
	// Parse the template and pass it the helper functions
	t := template.Must(template.New("go_config_gen.tmpl").Funcs(helperFuncs).Parse(temp))
	// Execute the template and pass it the metadata we collected before
	t.Execute(out, map[string]interface{}{"ConfigStructs": configStructs,"Fns":fns,"Package":"config"})
	formattedCode,err := imports.Process("config/config_gen.go",out.Bytes(), &imports.Options{Comments: true,FormatOnly:false})
	if err != nil {
		fmt.Printf("cannot format source code, might be an error in template: %s\n", err)
	}
	ioutil.WriteFile("config/config_gen_target.go", formattedCode, 0777)
}


func parse(files []string) []configureStruct {
	configureStructList:=make([]configureStruct,0)
	for _,fs:=range files {
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset,fs, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		//ast.Print(fset,f)
		packageName:=f.Name.Name
		for _, v := range f.Decls {
			fn, ok := v.(*ast.GenDecl)
			if !ok {
				continue
			}
			if fn.Doc == nil {
				continue
			}

			for _,e:=range fn.Doc.List {
				l, _ := regexp.MatchString(matchCommentConfigStruct, e.Text)
				if !l {
					continue
				}
				spec,ok:=fn.Specs[0].(*ast.TypeSpec)
				if !ok{
					continue
				}
				st,ok:=spec.Type.(*ast.StructType)
				if !ok {
					continue
				}
				csfsList:=make([]configureStructFields,0)
				for _,v:=range st.Fields.List{
					csfs:=configureStructFields{}
					var ft string //filed type
					ty,ok:=	v.Type.(*ast.Ident)
					if !ok {
						ay,ok:=	v.Type.(*ast.ArrayType)
						if !ok {
							my,ok:=	v.Type.(*ast.MapType)
							if !ok {
								panic("not allowed struct []")
							}
							ft=fmt.Sprintf("map[%s]%s",my.Key,my.Value)
						}else {
							ft="[]"+ay.Elt.(*ast.Ident).Name
						}
					}else {
						ft=ty.Name
					}

					csfs.Tag=v.Tag.Value
					csfs.FiledType=ft
					//TODO
					// WireGenFile,WireGenFilesssss []string //inge two or more file on same line
					csfs.FiledName=v.Names[0].Name
					csfsList=append(csfsList,csfs)
				}
				cfgStruct:=configureStruct{}
				cfgStruct.PackageName=packageName
				cfgStruct.YamlKey=spec.Name.Name
				if 	ParseConfigStruct(e.Text).Key !=""{
					cfgStruct.YamlKey=ParseConfigStruct(e.Text).Key
				}

				cfgStruct.AttrName=cfgStruct.YamlKey
				if cfgStruct.YamlKey!="###" {
					cfgStruct.YamlKey=fmt.Sprintf("`yaml:\"%s\"`",cfgStruct.YamlKey)
				}
				cfgStruct.StructName=spec.Name.Name
				cfgStruct.StructFields=csfsList
				configureStructList=append(configureStructList,cfgStruct)
				break
			}
		}
	}
	return configureStructList
}

type InjectConfigStruct struct {
	Key string `json:"key"`
}

func ParseConfigStruct(text string) InjectConfigStruct {
	res := InjectConfigStruct{}
	t := reflect.TypeOf(res)
	st := regexp.MustCompile("([a-z]*)=([^,)]*)").FindAllString(text, t.NumField())
	if len(st)< 1{
		return res
	}
	value := strings.Split(st[0], "=")
	if len(value)< 1{
		return res
	}
	res.Key =  strings.Trim(value[1], "\"")
	return res
}

