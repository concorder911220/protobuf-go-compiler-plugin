# 💫 About Plugin:

This plugin allow you to generate initial microservices with proto files and templates.

# 📊 Types:

```go

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

```

# Supplying Data

## For modules:

```go
{
  Messages []Message
  Enums []Enum
  Services []Service
  HasTimestamp bool
  ModulePath string
  ModuleName string
}
```

## For methods:

```go
{
  ServiceName      string
  ServiceNameLower string
  MethodName       string
  RequestType      string
  ResponseType     string
  PackageName      string
  ModuleName       string
  ModulePath       string
  IsCommand        bool
}
```

# Command example

```bash
protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --gbtemplate_out=modules=true,out=../internal,tplpath=../templates_folder:. ./proto/*.proto
protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --gbtemplate_out=types=true,out=../internal,tplpath=../templates_folder:. ./proto/*.proto
```

## 📫 How to contribute:

```yaml
version: v1
plugins:
  - plugin: gbtemplate
    out: ../
    opt: modules=true,out=../internal,tplpath=../templates

  - plugin: gbtemplate
    out: ../
    opt: methods=true,out=../,tplpath=../templates
```
