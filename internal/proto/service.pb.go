// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: internal/proto/service.proto

package TestTaskKvadro

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Authorlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *Authorlist) Reset() {
	*x = Authorlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authorlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorlist) ProtoMessage() {}

func (x *Authorlist) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorlist.ProtoReflect.Descriptor instead.
func (*Authorlist) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Authorlist) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_internal_proto_service_proto protoreflect.FileDescriptor

var file_internal_proto_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b,
	0x22, 0x30, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x32, 0x6f, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x32, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x12, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x0a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4f, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x73, 0x61, 0x73, 0x2f, 0x54, 0x65,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x76, 0x61, 0x64, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_service_proto_rawDescOnce sync.Once
	file_internal_proto_service_proto_rawDescData = file_internal_proto_service_proto_rawDesc
)

func file_internal_proto_service_proto_rawDescGZIP() []byte {
	file_internal_proto_service_proto_rawDescOnce.Do(func() {
		file_internal_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_service_proto_rawDescData)
	})
	return file_internal_proto_service_proto_rawDescData
}

var file_internal_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_service_proto_goTypes = []interface{}{
	(*Author)(nil),     // 0: internal.Author
	(*Authorlist)(nil), // 1: internal.Authorlist
	(*Book)(nil),       // 2: internal.Book
	(*BookList)(nil),   // 3: internal.BookList
}
var file_internal_proto_service_proto_depIdxs = []int32{
	0, // 0: internal.Authorlist.authors:type_name -> internal.Author
	2, // 1: internal.BookList.books:type_name -> internal.Book
	0, // 2: internal.API.FindBook:input_type -> internal.Author
	2, // 3: internal.API.FindAuthor:input_type -> internal.Book
	3, // 4: internal.API.FindBook:output_type -> internal.BookList
	1, // 5: internal.API.FindAuthor:output_type -> internal.Authorlist
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_service_proto_init() }
func file_internal_proto_service_proto_init() {
	if File_internal_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authorlist); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_service_proto_goTypes,
		DependencyIndexes: file_internal_proto_service_proto_depIdxs,
		MessageInfos:      file_internal_proto_service_proto_msgTypes,
	}.Build()
	File_internal_proto_service_proto = out.File
	file_internal_proto_service_proto_rawDesc = nil
	file_internal_proto_service_proto_goTypes = nil
	file_internal_proto_service_proto_depIdxs = nil
}
