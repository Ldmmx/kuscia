# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/datamesh/flightdm.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1.datamesh import domaindata_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindata__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1kuscia/proto/api/v1alpha1/datamesh/flightdm.proto\x12\"kuscia.proto.api.v1alpha1.datamesh\x1a\x33kuscia/proto/api/v1alpha1/datamesh/domaindata.proto\"*\n\x0f\x43SVWriteOptions\x12\x17\n\x0f\x66ield_delimiter\x18\x01 \x01(\t\"i\n\x10\x46ileWriteOptions\x12J\n\x0b\x63sv_options\x18\x02 \x01(\x0b\x32\x33.kuscia.proto.api.v1alpha1.datamesh.CSVWriteOptionsH\x00\x42\t\n\x07Options\"3\n\x1a\x43ommandGetDomainDataSchema\x12\x15\n\rdomaindata_id\x18\x01 \x01(\t\"\xf1\x01\n\x16\x43ommandDomainDataQuery\x12\x15\n\rdomaindata_id\x18\x01 \x01(\t\x12\x0f\n\x07\x63olumns\x18\x02 \x03(\t\x12\x45\n\x0c\x63ontent_type\x18\x03 \x01(\x0e\x32/.kuscia.proto.api.v1alpha1.datamesh.ContentType\x12P\n\x12\x66ile_write_options\x18\x04 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.datamesh.FileWriteOptions\x12\x16\n\x0epartition_spec\x18\x05 \x01(\t\"\xbd\x03\n\x17\x43ommandDomainDataUpdate\x12\x15\n\rdomaindata_id\x18\x01 \x01(\t\x12W\n\x12\x64omaindata_request\x18\x02 \x01(\x0b\x32;.kuscia.proto.api.v1alpha1.datamesh.CreateDomainDataRequest\x12\x45\n\x0c\x63ontent_type\x18\x03 \x01(\x0e\x32/.kuscia.proto.api.v1alpha1.datamesh.ContentType\x12P\n\x12\x66ile_write_options\x18\x04 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.datamesh.FileWriteOptions\x12\x64\n\rextra_options\x18\x05 \x03(\x0b\x32M.kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataUpdate.ExtraOptionsEntry\x1a\x33\n\x11\x45xtraOptionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01**\n\x0b\x43ontentType\x12\t\n\x05Table\x10\x00\x12\x07\n\x03RAW\x10\x01\x12\x07\n\x03\x43SV\x10\x02\x42\\\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datameshb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.datamesh.flightdm_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datamesh'
  _globals['_COMMANDDOMAINDATAUPDATE_EXTRAOPTIONSENTRY']._options = None
  _globals['_COMMANDDOMAINDATAUPDATE_EXTRAOPTIONSENTRY']._serialized_options = b'8\001'
  _globals['_CONTENTTYPE']._serialized_start=1038
  _globals['_CONTENTTYPE']._serialized_end=1080
  _globals['_CSVWRITEOPTIONS']._serialized_start=142
  _globals['_CSVWRITEOPTIONS']._serialized_end=184
  _globals['_FILEWRITEOPTIONS']._serialized_start=186
  _globals['_FILEWRITEOPTIONS']._serialized_end=291
  _globals['_COMMANDGETDOMAINDATASCHEMA']._serialized_start=293
  _globals['_COMMANDGETDOMAINDATASCHEMA']._serialized_end=344
  _globals['_COMMANDDOMAINDATAQUERY']._serialized_start=347
  _globals['_COMMANDDOMAINDATAQUERY']._serialized_end=588
  _globals['_COMMANDDOMAINDATAUPDATE']._serialized_start=591
  _globals['_COMMANDDOMAINDATAUPDATE']._serialized_end=1036
  _globals['_COMMANDDOMAINDATAUPDATE_EXTRAOPTIONSENTRY']._serialized_start=985
  _globals['_COMMANDDOMAINDATAUPDATE_EXTRAOPTIONSENTRY']._serialized_end=1036
# @@protoc_insertion_point(module_scope)
