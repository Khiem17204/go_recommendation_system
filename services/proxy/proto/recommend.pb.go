// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/recommend.proto

package recommend

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

type RecommendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cards         []string               `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendRequest) Reset() {
	*x = RecommendRequest{}
	mi := &file_proto_recommend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendRequest) ProtoMessage() {}

func (x *RecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendRequest.ProtoReflect.Descriptor instead.
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendRequest) GetCards() []string {
	if x != nil {
		return x.Cards
	}
	return nil
}

type CardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Cards         []string               `protobuf:"bytes,2,rep,name=cards,proto3" json:"cards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CardResponse) Reset() {
	*x = CardResponse{}
	mi := &file_proto_recommend_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardResponse) ProtoMessage() {}

func (x *CardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardResponse.ProtoReflect.Descriptor instead.
func (*CardResponse) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *CardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CardResponse) GetCards() []string {
	if x != nil {
		return x.Cards
	}
	return nil
}

type DeckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Decks         []string               `protobuf:"bytes,2,rep,name=decks,proto3" json:"decks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeckResponse) Reset() {
	*x = DeckResponse{}
	mi := &file_proto_recommend_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckResponse) ProtoMessage() {}

func (x *DeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckResponse.ProtoReflect.Descriptor instead.
func (*DeckResponse) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{2}
}

func (x *DeckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeckResponse) GetDecks() []string {
	if x != nil {
		return x.Decks
	}
	return nil
}

var File_proto_recommend_proto protoreflect.FileDescriptor

const file_proto_recommend_proto_rawDesc = "" +
	"\n" +
	"\x15proto/recommend.proto\x12\trecommend\"(\n" +
	"\x10RecommendRequest\x12\x14\n" +
	"\x05cards\x18\x01 \x03(\tR\x05cards\"<\n" +
	"\fCardResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x14\n" +
	"\x05cards\x18\x02 \x03(\tR\x05cards\"<\n" +
	"\fDeckResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x14\n" +
	"\x05decks\x18\x02 \x03(\tR\x05decks2\xa5\x01\n" +
	"\x15RecommendationService\x12E\n" +
	"\rCardRecommend\x12\x1b.recommend.RecommendRequest\x1a\x17.recommend.CardResponse\x12E\n" +
	"\rDeckRecommend\x12\x1b.recommend.RecommendRequest\x1a\x17.recommend.DeckResponseBOZMgithub.com/Khiem17204/go_recommendation_system/services/proxy/proto/recommendb\x06proto3"

var (
	file_proto_recommend_proto_rawDescOnce sync.Once
	file_proto_recommend_proto_rawDescData []byte
)

func file_proto_recommend_proto_rawDescGZIP() []byte {
	file_proto_recommend_proto_rawDescOnce.Do(func() {
		file_proto_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_recommend_proto_rawDesc), len(file_proto_recommend_proto_rawDesc)))
	})
	return file_proto_recommend_proto_rawDescData
}

var file_proto_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_recommend_proto_goTypes = []any{
	(*RecommendRequest)(nil), // 0: recommend.RecommendRequest
	(*CardResponse)(nil),     // 1: recommend.CardResponse
	(*DeckResponse)(nil),     // 2: recommend.DeckResponse
}
var file_proto_recommend_proto_depIdxs = []int32{
	0, // 0: recommend.RecommendationService.CardRecommend:input_type -> recommend.RecommendRequest
	0, // 1: recommend.RecommendationService.DeckRecommend:input_type -> recommend.RecommendRequest
	1, // 2: recommend.RecommendationService.CardRecommend:output_type -> recommend.CardResponse
	2, // 3: recommend.RecommendationService.DeckRecommend:output_type -> recommend.DeckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_recommend_proto_init() }
func file_proto_recommend_proto_init() {
	if File_proto_recommend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_recommend_proto_rawDesc), len(file_proto_recommend_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_recommend_proto_goTypes,
		DependencyIndexes: file_proto_recommend_proto_depIdxs,
		MessageInfos:      file_proto_recommend_proto_msgTypes,
	}.Build()
	File_proto_recommend_proto = out.File
	file_proto_recommend_proto_goTypes = nil
	file_proto_recommend_proto_depIdxs = nil
}
