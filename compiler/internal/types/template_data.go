package types

type TemplateData struct {
	PackageName  string
	Messages     []Message
	Enums        []Enum
	HasTimestamp bool
}

type Message struct {
	MessageName string
	Fields      []Field
}

type Field struct {
	Name      string
	Type      string
	Number    int
	ProtoName string
}

type Service struct {
	SName   string
	Methods []Method
}

type Method struct {
	Name         string
	RequestType  string
	ResponseType string
}

type Enum struct {
	EnumName string
	Values   []EnumValue
}

type EnumValue struct {
	PName string
	Name  string
	Value int32
}
