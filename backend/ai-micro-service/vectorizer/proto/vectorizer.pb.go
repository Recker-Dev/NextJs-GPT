// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: vectorizer.proto

package vectorizerpb

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

type PDFChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FileId        string                 `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename      string                 `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Data          []byte                 `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	IsFirstChunk  bool                   `protobuf:"varint,6,opt,name=is_first_chunk,json=isFirstChunk,proto3" json:"is_first_chunk,omitempty"`
	IsLastChunk   bool                   `protobuf:"varint,7,opt,name=is_last_chunk,json=isLastChunk,proto3" json:"is_last_chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PDFChunk) Reset() {
	*x = PDFChunk{}
	mi := &file_vectorizer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PDFChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDFChunk) ProtoMessage() {}

func (x *PDFChunk) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDFChunk.ProtoReflect.Descriptor instead.
func (*PDFChunk) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{0}
}

func (x *PDFChunk) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PDFChunk) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *PDFChunk) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *PDFChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *PDFChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PDFChunk) GetIsFirstChunk() bool {
	if x != nil {
		return x.IsFirstChunk
	}
	return false
}

func (x *PDFChunk) GetIsLastChunk() bool {
	if x != nil {
		return x.IsLastChunk
	}
	return false
}

type VectorizeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ChunkCount    int32                  `protobuf:"varint,3,opt,name=chunk_count,json=chunkCount,proto3" json:"chunk_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VectorizeResponse) Reset() {
	*x = VectorizeResponse{}
	mi := &file_vectorizer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VectorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorizeResponse) ProtoMessage() {}

func (x *VectorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorizeResponse.ProtoReflect.Descriptor instead.
func (*VectorizeResponse) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{1}
}

func (x *VectorizeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VectorizeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VectorizeResponse) GetChunkCount() int32 {
	if x != nil {
		return x.ChunkCount
	}
	return 0
}

type QueryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FileId        string                 `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	TopK          int32                  `protobuf:"varint,4,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	QueryTexts    []string               `protobuf:"bytes,5,rep,name=query_texts,json=queryTexts,proto3" json:"query_texts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	mi := &file_vectorizer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *QueryRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *QueryRequest) GetTopK() int32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

func (x *QueryRequest) GetQueryTexts() []string {
	if x != nil {
		return x.QueryTexts
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Results       []*QueryResult         `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	mi := &file_vectorizer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{3}
}

func (x *QueryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QueryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryResponse) GetResults() []*QueryResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type QueryResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Document      string                 `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	Source        string                 `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Page          int32                  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Distance      float64                `protobuf:"fixed64,5,opt,name=distance,proto3" json:"distance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	mi := &file_vectorizer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{4}
}

func (x *QueryResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryResult) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *QueryResult) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *QueryResult) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryResult) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type CountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FileId        string                 `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	mi := &file_vectorizer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{5}
}

func (x *CountRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CountRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *CountRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type CountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count         int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	mi := &file_vectorizer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{6}
}

func (x *CountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FileId        string                 `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_vectorizer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *DeleteRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_vectorizer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vectorizer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_vectorizer_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_vectorizer_proto protoreflect.FileDescriptor

const file_vectorizer_proto_rawDesc = "" +
	"\n" +
	"\x10vectorizer.proto\x12\n" +
	"vectorizer\"\xcf\x01\n" +
	"\bPDFChunk\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x17\n" +
	"\afile_id\x18\x03 \x01(\tR\x06fileId\x12\x1a\n" +
	"\bfilename\x18\x04 \x01(\tR\bfilename\x12\x12\n" +
	"\x04data\x18\x05 \x01(\fR\x04data\x12$\n" +
	"\x0eis_first_chunk\x18\x06 \x01(\bR\fisFirstChunk\x12\"\n" +
	"\ris_last_chunk\x18\a \x01(\bR\visLastChunk\"h\n" +
	"\x11VectorizeResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x1f\n" +
	"\vchunk_count\x18\x03 \x01(\x05R\n" +
	"chunkCount\"\x8f\x01\n" +
	"\fQueryRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x17\n" +
	"\afile_id\x18\x03 \x01(\tR\x06fileId\x12\x13\n" +
	"\x05top_k\x18\x04 \x01(\x05R\x04topK\x12\x1f\n" +
	"\vquery_texts\x18\x05 \x03(\tR\n" +
	"queryTexts\"v\n" +
	"\rQueryResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x121\n" +
	"\aresults\x18\x03 \x03(\v2\x17.vectorizer.QueryResultR\aresults\"\x81\x01\n" +
	"\vQueryResult\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bdocument\x18\x02 \x01(\tR\bdocument\x12\x16\n" +
	"\x06source\x18\x03 \x01(\tR\x06source\x12\x12\n" +
	"\x04page\x18\x04 \x01(\x05R\x04page\x12\x1a\n" +
	"\bdistance\x18\x05 \x01(\x01R\bdistance\"Y\n" +
	"\fCountRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x17\n" +
	"\afile_id\x18\x03 \x01(\tR\x06fileId\"Y\n" +
	"\rCountResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x14\n" +
	"\x05count\x18\x03 \x01(\x05R\x05count\"Z\n" +
	"\rDeleteRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x17\n" +
	"\afile_id\x18\x03 \x01(\tR\x06fileId\"D\n" +
	"\x0eDeleteResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xb2\x02\n" +
	"\x11VectorizerService\x12K\n" +
	"\x12UploadAndVectorize\x12\x14.vectorizer.PDFChunk\x1a\x1d.vectorizer.VectorizeResponse(\x01\x12C\n" +
	"\fQueryVectors\x12\x18.vectorizer.QueryRequest\x1a\x19.vectorizer.QueryResponse\x12C\n" +
	"\fCountVectors\x12\x18.vectorizer.CountRequest\x1a\x19.vectorizer.CountResponse\x12F\n" +
	"\rDeleteVectors\x12\x19.vectorizer.DeleteRequest\x1a\x1a.vectorizer.DeleteResponseB\x1fZ\x1dvectorizer/proto;vectorizerpbb\x06proto3"

var (
	file_vectorizer_proto_rawDescOnce sync.Once
	file_vectorizer_proto_rawDescData []byte
)

func file_vectorizer_proto_rawDescGZIP() []byte {
	file_vectorizer_proto_rawDescOnce.Do(func() {
		file_vectorizer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vectorizer_proto_rawDesc), len(file_vectorizer_proto_rawDesc)))
	})
	return file_vectorizer_proto_rawDescData
}

var file_vectorizer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vectorizer_proto_goTypes = []any{
	(*PDFChunk)(nil),          // 0: vectorizer.PDFChunk
	(*VectorizeResponse)(nil), // 1: vectorizer.VectorizeResponse
	(*QueryRequest)(nil),      // 2: vectorizer.QueryRequest
	(*QueryResponse)(nil),     // 3: vectorizer.QueryResponse
	(*QueryResult)(nil),       // 4: vectorizer.QueryResult
	(*CountRequest)(nil),      // 5: vectorizer.CountRequest
	(*CountResponse)(nil),     // 6: vectorizer.CountResponse
	(*DeleteRequest)(nil),     // 7: vectorizer.DeleteRequest
	(*DeleteResponse)(nil),    // 8: vectorizer.DeleteResponse
}
var file_vectorizer_proto_depIdxs = []int32{
	4, // 0: vectorizer.QueryResponse.results:type_name -> vectorizer.QueryResult
	0, // 1: vectorizer.VectorizerService.UploadAndVectorize:input_type -> vectorizer.PDFChunk
	2, // 2: vectorizer.VectorizerService.QueryVectors:input_type -> vectorizer.QueryRequest
	5, // 3: vectorizer.VectorizerService.CountVectors:input_type -> vectorizer.CountRequest
	7, // 4: vectorizer.VectorizerService.DeleteVectors:input_type -> vectorizer.DeleteRequest
	1, // 5: vectorizer.VectorizerService.UploadAndVectorize:output_type -> vectorizer.VectorizeResponse
	3, // 6: vectorizer.VectorizerService.QueryVectors:output_type -> vectorizer.QueryResponse
	6, // 7: vectorizer.VectorizerService.CountVectors:output_type -> vectorizer.CountResponse
	8, // 8: vectorizer.VectorizerService.DeleteVectors:output_type -> vectorizer.DeleteResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vectorizer_proto_init() }
func file_vectorizer_proto_init() {
	if File_vectorizer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vectorizer_proto_rawDesc), len(file_vectorizer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vectorizer_proto_goTypes,
		DependencyIndexes: file_vectorizer_proto_depIdxs,
		MessageInfos:      file_vectorizer_proto_msgTypes,
	}.Build()
	File_vectorizer_proto = out.File
	file_vectorizer_proto_goTypes = nil
	file_vectorizer_proto_depIdxs = nil
}
