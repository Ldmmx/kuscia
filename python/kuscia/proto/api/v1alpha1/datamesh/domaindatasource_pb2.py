# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/datamesh/domaindatasource.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n9kuscia/proto/api/v1alpha1/datamesh/domaindatasource.proto\x12\"kuscia.proto.api.v1alpha1.datamesh\x1a&kuscia/proto/api/v1alpha1/common.proto\"o\n\x1cQueryDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x15\n\rdatasource_id\x18\x02 \x01(\t\"\x96\x01\n\x1dQueryDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x42\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.datamesh.DomainDataSource\"\xc2\x01\n\x10\x44omainDataSource\x12\x15\n\rdatasource_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x0e\n\x06status\x18\x04 \x01(\t\x12@\n\x04info\x18\x05 \x01(\x0b\x32\x32.kuscia.proto.api.v1alpha1.datamesh.DataSourceInfo\x12\x10\n\x08info_key\x18\x06 \x01(\t\x12\x17\n\x0f\x61\x63\x63\x65ss_directly\x18\x07 \x01(\x08\"\xb2\x02\n\x0e\x44\x61taSourceInfo\x12H\n\x07localfs\x18\x01 \x01(\x0b\x32\x37.kuscia.proto.api.v1alpha1.datamesh.LocalDataSourceInfo\x12\x42\n\x03oss\x18\x02 \x01(\x0b\x32\x35.kuscia.proto.api.v1alpha1.datamesh.OssDataSourceInfo\x12L\n\x08\x64\x61tabase\x18\x03 \x01(\x0b\x32:.kuscia.proto.api.v1alpha1.datamesh.DatabaseDataSourceInfo\x12\x44\n\x04odps\x18\x04 \x01(\x0b\x32\x36.kuscia.proto.api.v1alpha1.datamesh.OdpsDataSourceInfo\"#\n\x13LocalDataSourceInfo\x12\x0c\n\x04path\x18\x01 \x01(\t\"\xb3\x01\n\x11OssDataSourceInfo\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x0e\n\x06\x62ucket\x18\x02 \x01(\t\x12\x0e\n\x06prefix\x18\x03 \x01(\t\x12\x15\n\raccess_key_id\x18\x04 \x01(\t\x12\x19\n\x11\x61\x63\x63\x65ss_key_secret\x18\x05 \x01(\t\x12\x13\n\x0bvirtualhost\x18\x06 \x01(\x08\x12\x0f\n\x07version\x18\x07 \x01(\t\x12\x14\n\x0cstorage_type\x18\x08 \x01(\t\"\\\n\x16\x44\x61tabaseDataSourceInfo\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x0c\n\x04user\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x10\n\x08\x64\x61tabase\x18\x04 \x01(\t\"i\n\x12OdpsDataSourceInfo\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x0f\n\x07project\x18\x02 \x01(\t\x12\x15\n\raccess_key_id\x18\x03 \x01(\t\x12\x19\n\x11\x61\x63\x63\x65ss_key_secret\x18\x04 \x01(\t2\xb8\x01\n\x17\x44omainDataSourceService\x12\x9c\x01\n\x15QueryDomainDataSource\x12@.kuscia.proto.api.v1alpha1.datamesh.QueryDomainDataSourceRequest\x1a\x41.kuscia.proto.api.v1alpha1.datamesh.QueryDomainDataSourceResponseB\\\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datameshb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.datamesh.domaindatasource_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datamesh'
  _globals['_QUERYDOMAINDATASOURCEREQUEST']._serialized_start=137
  _globals['_QUERYDOMAINDATASOURCEREQUEST']._serialized_end=248
  _globals['_QUERYDOMAINDATASOURCERESPONSE']._serialized_start=251
  _globals['_QUERYDOMAINDATASOURCERESPONSE']._serialized_end=401
  _globals['_DOMAINDATASOURCE']._serialized_start=404
  _globals['_DOMAINDATASOURCE']._serialized_end=598
  _globals['_DATASOURCEINFO']._serialized_start=601
  _globals['_DATASOURCEINFO']._serialized_end=907
  _globals['_LOCALDATASOURCEINFO']._serialized_start=909
  _globals['_LOCALDATASOURCEINFO']._serialized_end=944
  _globals['_OSSDATASOURCEINFO']._serialized_start=947
  _globals['_OSSDATASOURCEINFO']._serialized_end=1126
  _globals['_DATABASEDATASOURCEINFO']._serialized_start=1128
  _globals['_DATABASEDATASOURCEINFO']._serialized_end=1220
  _globals['_ODPSDATASOURCEINFO']._serialized_start=1222
  _globals['_ODPSDATASOURCEINFO']._serialized_end=1327
  _globals['_DOMAINDATASOURCESERVICE']._serialized_start=1330
  _globals['_DOMAINDATASOURCESERVICE']._serialized_end=1514
# @@protoc_insertion_point(module_scope)
