- Perspectives
  This plugin is generating methods, types and interface and service from proto files.
  After running this command, you will see templates folder which has module.tmpl, service_gen.tmpl and types.tmpl, app.tmpl, {method}.tmpl
  Of course you can modify this template but you should consider templates have their own types.

package types

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
}

type Service struct {
SName string
Methods []Method
}

type Method struct {
Name string
RequestType string
ResponseType string
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

f.e:

message Vehicle {
int32 id = 1;
string vin = 2;
int32 year = 3;
string make = 4;
string model = 5;
string plate_number = 6;
string image_url = 7;
string thumb_url = 8;
VehicleOwnership ownership = 9;
string nags_id = 10;
VehicleType vehicle_type = 11;
int32 number = 12; // Vehicle number in the policy (to vehicle if multiple vehicles in the policy)
string style = 13;
int32 make_id = 14;
int32 model_id = 15;
}

Please consider VehicleType, because in current insurance repo, it is defined as VehicleType type = 11;
But it should be defined as vehicle_type in this plugin.
We need to investigate about this issue.
