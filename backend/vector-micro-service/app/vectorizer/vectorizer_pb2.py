# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: vectorizer.proto
# Protobuf Python Version: 6.31.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    0,
    '',
    'vectorizer.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10vectorizer.proto\x12\nvectorizer\"\x8c\x01\n\x08PDFChunk\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07\x63hat_id\x18\x02 \x01(\t\x12\x0f\n\x07\x66ile_id\x18\x03 \x01(\t\x12\x10\n\x08\x66ilename\x18\x04 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x05 \x01(\x0c\x12\x16\n\x0eis_first_chunk\x18\x06 \x01(\x08\x12\x15\n\ris_last_chunk\x18\x07 \x01(\x08\"J\n\x11VectorizeResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x13\n\x0b\x63hunk_count\x18\x03 \x01(\x05\"e\n\x0cQueryRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07\x63hat_id\x18\x02 \x01(\t\x12\x0f\n\x07\x66ile_id\x18\x03 \x01(\t\x12\r\n\x05top_k\x18\x04 \x01(\x05\x12\x13\n\x0bquery_texts\x18\x05 \x03(\t\"[\n\rQueryResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\x12(\n\x07results\x18\x03 \x03(\x0b\x32\x17.vectorizer.QueryResult\"[\n\x0bQueryResult\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08\x64ocument\x18\x02 \x01(\t\x12\x0e\n\x06source\x18\x03 \x01(\t\x12\x0c\n\x04page\x18\x04 \x01(\x05\x12\x10\n\x08\x64istance\x18\x05 \x01(\x01\"A\n\x0c\x43ountRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07\x63hat_id\x18\x02 \x01(\t\x12\x0f\n\x07\x66ile_id\x18\x03 \x01(\t\"@\n\rCountResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\r\n\x05\x63ount\x18\x03 \x01(\x05\"B\n\rDeleteRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07\x63hat_id\x18\x02 \x01(\t\x12\x0f\n\x07\x66ile_id\x18\x03 \x01(\t\"2\n\x0e\x44\x65leteResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t2\xb2\x02\n\x11VectorizerService\x12K\n\x12UploadAndVectorize\x12\x14.vectorizer.PDFChunk\x1a\x1d.vectorizer.VectorizeResponse(\x01\x12\x43\n\x0cQueryVectors\x12\x18.vectorizer.QueryRequest\x1a\x19.vectorizer.QueryResponse\x12\x43\n\x0c\x43ountVectors\x12\x18.vectorizer.CountRequest\x1a\x19.vectorizer.CountResponse\x12\x46\n\rDeleteVectors\x12\x19.vectorizer.DeleteRequest\x1a\x1a.vectorizer.DeleteResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'vectorizer_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_PDFCHUNK']._serialized_start=33
  _globals['_PDFCHUNK']._serialized_end=173
  _globals['_VECTORIZERESPONSE']._serialized_start=175
  _globals['_VECTORIZERESPONSE']._serialized_end=249
  _globals['_QUERYREQUEST']._serialized_start=251
  _globals['_QUERYREQUEST']._serialized_end=352
  _globals['_QUERYRESPONSE']._serialized_start=354
  _globals['_QUERYRESPONSE']._serialized_end=445
  _globals['_QUERYRESULT']._serialized_start=447
  _globals['_QUERYRESULT']._serialized_end=538
  _globals['_COUNTREQUEST']._serialized_start=540
  _globals['_COUNTREQUEST']._serialized_end=605
  _globals['_COUNTRESPONSE']._serialized_start=607
  _globals['_COUNTRESPONSE']._serialized_end=671
  _globals['_DELETEREQUEST']._serialized_start=673
  _globals['_DELETEREQUEST']._serialized_end=739
  _globals['_DELETERESPONSE']._serialized_start=741
  _globals['_DELETERESPONSE']._serialized_end=791
  _globals['_VECTORIZERSERVICE']._serialized_start=794
  _globals['_VECTORIZERSERVICE']._serialized_end=1100
# @@protoc_insertion_point(module_scope)
