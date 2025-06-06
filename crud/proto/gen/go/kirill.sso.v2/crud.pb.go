// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: crud.proto

package kirill_sso_v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrencyRate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Rate          string                 `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"` // курс в рублях (double = float64)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyRate) Reset() {
	*x = CurrencyRate{}
	mi := &file_crud_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRate) ProtoMessage() {}

func (x *CurrencyRate) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRate.ProtoReflect.Descriptor instead.
func (*CurrencyRate) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRate) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CurrencyRate) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *CurrencyRate) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type CurrencyRates struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Rates         []*CurrencyRate        `protobuf:"bytes,2,rep,name=rates,proto3" json:"rates,omitempty"` // repeated = список с курсами валют
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyRates) Reset() {
	*x = CurrencyRates{}
	mi := &file_crud_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRates) ProtoMessage() {}

func (x *CurrencyRates) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRates.ProtoReflect.Descriptor instead.
func (*CurrencyRates) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyRates) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CurrencyRates) GetRates() []*CurrencyRate {
	if x != nil {
		return x.Rates
	}
	return nil
}

type CreateCurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Rate          string                 `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCurrencyRequest) Reset() {
	*x = CreateCurrencyRequest{}
	mi := &file_crud_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCurrencyRequest) ProtoMessage() {}

func (x *CreateCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCurrencyRequest.ProtoReflect.Descriptor instead.
func (*CreateCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCurrencyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateCurrencyRequest) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *CreateCurrencyRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type CreateCurrencyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CreatedRate   *CurrencyRate          `protobuf:"bytes,1,opt,name=created_rate,json=createdRate,proto3" json:"created_rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCurrencyResponse) Reset() {
	*x = CreateCurrencyResponse{}
	mi := &file_crud_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCurrencyResponse) ProtoMessage() {}

func (x *CreateCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCurrencyResponse.ProtoReflect.Descriptor instead.
func (*CreateCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCurrencyResponse) GetCreatedRate() *CurrencyRate {
	if x != nil {
		return x.CreatedRate
	}
	return nil
}

type UpdateCurrencyRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CurrencyUpdate *CurrencyRate          `protobuf:"bytes,1,opt,name=currency_update,json=currencyUpdate,proto3" json:"currency_update,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateCurrencyRequest) Reset() {
	*x = UpdateCurrencyRequest{}
	mi := &file_crud_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrencyRequest) ProtoMessage() {}

func (x *UpdateCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrencyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCurrencyRequest) GetCurrencyUpdate() *CurrencyRate {
	if x != nil {
		return x.CurrencyUpdate
	}
	return nil
}

type UpdateCurrencyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCurrencyResponse) Reset() {
	*x = UpdateCurrencyResponse{}
	mi := &file_crud_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrencyResponse) ProtoMessage() {}

func (x *UpdateCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrencyResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCurrencyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateCurrencyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteCurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCurrencyRequest) Reset() {
	*x = DeleteCurrencyRequest{}
	mi := &file_crud_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCurrencyRequest) ProtoMessage() {}

func (x *DeleteCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCurrencyRequest.ProtoReflect.Descriptor instead.
func (*DeleteCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCurrencyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type DeleteCurrencyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCurrencyResponse) Reset() {
	*x = DeleteCurrencyResponse{}
	mi := &file_crud_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCurrencyResponse) ProtoMessage() {}

func (x *DeleteCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCurrencyResponse.ProtoReflect.Descriptor instead.
func (*DeleteCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCurrencyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrencyRequest) Reset() {
	*x = GetCurrencyRequest{}
	mi := &file_crud_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyRequest) ProtoMessage() {}

func (x *GetCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyRequest.ProtoReflect.Descriptor instead.
func (*GetCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{8}
}

func (x *GetCurrencyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetCurrencyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      *CurrencyRate          `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrencyResponse) Reset() {
	*x = GetCurrencyResponse{}
	mi := &file_crud_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyResponse) ProtoMessage() {}

func (x *GetCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyResponse.ProtoReflect.Descriptor instead.
func (*GetCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{9}
}

func (x *GetCurrencyResponse) GetCurrency() *CurrencyRate {
	if x != nil {
		return x.Currency
	}
	return nil
}

var File_crud_proto protoreflect.FileDescriptor

const file_crud_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"crud.proto\x12\x04crud\"L\n" +
	"\fCurrencyRate\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x12\n" +
	"\x04rate\x18\x02 \x01(\tR\x04rate\x12\x14\n" +
	"\x05value\x18\x03 \x01(\x01R\x05value\"M\n" +
	"\rCurrencyRates\x12\x12\n" +
	"\x04date\x18\x01 \x01(\tR\x04date\x12(\n" +
	"\x05rates\x18\x02 \x03(\v2\x12.crud.CurrencyRateR\x05rates\"U\n" +
	"\x15CreateCurrencyRequest\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x12\n" +
	"\x04rate\x18\x02 \x01(\tR\x04rate\x12\x14\n" +
	"\x05value\x18\x03 \x01(\x01R\x05value\"O\n" +
	"\x16CreateCurrencyResponse\x125\n" +
	"\fcreated_rate\x18\x01 \x01(\v2\x12.crud.CurrencyRateR\vcreatedRate\"T\n" +
	"\x15UpdateCurrencyRequest\x12;\n" +
	"\x0fcurrency_update\x18\x01 \x01(\v2\x12.crud.CurrencyRateR\x0ecurrencyUpdate\"L\n" +
	"\x16UpdateCurrencyResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"+\n" +
	"\x15DeleteCurrencyRequest\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\"2\n" +
	"\x16DeleteCurrencyResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"(\n" +
	"\x12GetCurrencyRequest\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\"E\n" +
	"\x13GetCurrencyResponse\x12.\n" +
	"\bcurrency\x18\x01 \x01(\v2\x12.crud.CurrencyRateR\bcurrency2\xa9\x02\n" +
	"\x04Crud\x12K\n" +
	"\x0eCreateCurrency\x12\x1b.crud.CreateCurrencyRequest\x1a\x1c.crud.CreateCurrencyResponse\x12K\n" +
	"\x0eUpdateCurrency\x12\x1b.crud.UpdateCurrencyRequest\x1a\x1c.crud.UpdateCurrencyResponse\x12C\n" +
	"\x06Delete\x12\x1b.crud.DeleteCurrencyRequest\x1a\x1c.crud.DeleteCurrencyResponse\x12B\n" +
	"\vGetCurrency\x12\x18.crud.GetCurrencyRequest\x1a\x19.crud.GetCurrencyResponseB\x0fZ\rkirill.sso.v2b\x06proto3"

var (
	file_crud_proto_rawDescOnce sync.Once
	file_crud_proto_rawDescData []byte
)

func file_crud_proto_rawDescGZIP() []byte {
	file_crud_proto_rawDescOnce.Do(func() {
		file_crud_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_crud_proto_rawDesc), len(file_crud_proto_rawDesc)))
	})
	return file_crud_proto_rawDescData
}

var file_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_crud_proto_goTypes = []any{
	(*CurrencyRate)(nil),           // 0: crud.CurrencyRate
	(*CurrencyRates)(nil),          // 1: crud.CurrencyRates
	(*CreateCurrencyRequest)(nil),  // 2: crud.CreateCurrencyRequest
	(*CreateCurrencyResponse)(nil), // 3: crud.CreateCurrencyResponse
	(*UpdateCurrencyRequest)(nil),  // 4: crud.UpdateCurrencyRequest
	(*UpdateCurrencyResponse)(nil), // 5: crud.UpdateCurrencyResponse
	(*DeleteCurrencyRequest)(nil),  // 6: crud.DeleteCurrencyRequest
	(*DeleteCurrencyResponse)(nil), // 7: crud.DeleteCurrencyResponse
	(*GetCurrencyRequest)(nil),     // 8: crud.GetCurrencyRequest
	(*GetCurrencyResponse)(nil),    // 9: crud.GetCurrencyResponse
}
var file_crud_proto_depIdxs = []int32{
	0, // 0: crud.CurrencyRates.rates:type_name -> crud.CurrencyRate
	0, // 1: crud.CreateCurrencyResponse.created_rate:type_name -> crud.CurrencyRate
	0, // 2: crud.UpdateCurrencyRequest.currency_update:type_name -> crud.CurrencyRate
	0, // 3: crud.GetCurrencyResponse.currency:type_name -> crud.CurrencyRate
	2, // 4: crud.Crud.CreateCurrency:input_type -> crud.CreateCurrencyRequest
	4, // 5: crud.Crud.UpdateCurrency:input_type -> crud.UpdateCurrencyRequest
	6, // 6: crud.Crud.Delete:input_type -> crud.DeleteCurrencyRequest
	8, // 7: crud.Crud.GetCurrency:input_type -> crud.GetCurrencyRequest
	3, // 8: crud.Crud.CreateCurrency:output_type -> crud.CreateCurrencyResponse
	5, // 9: crud.Crud.UpdateCurrency:output_type -> crud.UpdateCurrencyResponse
	7, // 10: crud.Crud.Delete:output_type -> crud.DeleteCurrencyResponse
	9, // 11: crud.Crud.GetCurrency:output_type -> crud.GetCurrencyResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_crud_proto_init() }
func file_crud_proto_init() {
	if File_crud_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crud_proto_rawDesc), len(file_crud_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_proto_goTypes,
		DependencyIndexes: file_crud_proto_depIdxs,
		MessageInfos:      file_crud_proto_msgTypes,
	}.Build()
	File_crud_proto = out.File
	file_crud_proto_goTypes = nil
	file_crud_proto_depIdxs = nil
}
