- Perspectives
  This plugin is generating methods, types and interface and service from proto files.
  After running this command, you will see templates folder which has module.tmpl, service_gen.tmpl and types.tmpl, app.tmpl, {method}.tmpl
  Of course you can modify this template but you should consider templates have their own types.

package generator

type SupplyData struct {
TypeData TypeData
MetaInfo MetaInfo
}

type MetaInfo struct {
ModuleName string
ModulePath string
OutputPath string
}

type TypeData struct {
PackageName string
Messages []Message
Enums []Enum
Services []Service
}

type TemplateData struct {
PackageName string
Messages []Message
Enums []Enum
HasTimestamp bool
}

type Message struct {
MessageName string
Fields []Field
}

type Field struct {
Name string
Type string
Number int
ProtoName string
HasTimestamp bool
// Optional string
}

type Service struct {
PackageName string
SName string
Methods []Method
}

type Method struct {
PackageName string
PName string
MName string
RequestType string
ResponseType string
IsCommand bool
}

type Enum struct {
EnumName string
Values []EnumValue
}

type EnumValue struct {
PName string
Name string
Value int32
}

- How to use this plugin?

1. set your plugin application PATH as your environment PATH.
2. copy buf.gen.yaml and buf.yaml files to your proto directory.
   You can sepcify your output directory in buf.yaml file by using opt: methods=true,out=../.
3. run buf build
4. run buf generate
5. This command will generate template folders files in templates folder for each app, methods, services, types and modules.
   These templates are default templates. You can modify them as you want.

6. After you modify template, you can also run buf generate command again so you can get updated go files.

- Becareful: There is one consideration while making proto schema.
  You should define field name "snake_case".

protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --gbtemplate*out=modules=true,out=../internal,tplpath=../templates_folder:. ./proto/*.proto
protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --gbtemplate*out=types=true,out=../internal,tplpath=../templates_folder:. ./proto/*.proto
