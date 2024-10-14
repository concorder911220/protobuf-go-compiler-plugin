package generator

type SupplyData struct {
	TypeData   TypeData
	ModulePath string
	OutputPath string
}

type TypeData struct {
	PackageName string
	Messages    []Message
	Enums       []Enum
	Services    []Service
}

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
	Name         string
	Type         string
	Number       int
	ProtoName    string
	HasTimestamp bool
	// Optional  string
}

type Service struct {
	PackageName string
	SName       string
	Methods     []Method
}

type Method struct {
	PackageName  string
	PName        string
	MName        string
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
