# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/datamesh/flightinner.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1.datamesh import domaindata_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindata__pb2
from kuscia.proto.api.v1alpha1.datamesh import domaindatasource_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2
from kuscia.proto.api.v1alpha1.datamesh import flightdm_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_flightdm__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4kuscia/proto/api/v1alpha1/datamesh/flightinner.proto\x12\"kuscia.proto.api.v1alpha1.datamesh\x1a\x33kuscia/proto/api/v1alpha1/datamesh/domaindata.proto\x1a\x39kuscia/proto/api/v1alpha1/datamesh/domaindatasource.proto\x1a\x31kuscia/proto/api/v1alpha1/datamesh/flightdm.proto\"\xef\x01\n\x14\x43ommandDataMeshQuery\x12I\n\x05query\x18\x01 \x01(\x0b\x32:.kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataQuery\x12\x42\n\ndomaindata\x18\x02 \x01(\x0b\x32..kuscia.proto.api.v1alpha1.datamesh.DomainData\x12H\n\ndatasource\x18\x03 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.datamesh.DomainDataSource\"\xf2\x01\n\x15\x43ommandDataMeshUpdate\x12K\n\x06update\x18\x01 \x01(\x0b\x32;.kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataUpdate\x12\x42\n\ndomaindata\x18\x02 \x01(\x0b\x32..kuscia.proto.api.v1alpha1.datamesh.DomainData\x12H\n\ndatasource\x18\x03 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.datamesh.DomainDataSourceB\\\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datameshb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.datamesh.flightinner_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datamesh'
  _globals['_COMMANDDATAMESHQUERY']._serialized_start=256
  _globals['_COMMANDDATAMESHQUERY']._serialized_end=495
  _globals['_COMMANDDATAMESHUPDATE']._serialized_start=498
  _globals['_COMMANDDATAMESHUPDATE']._serialized_end=740
# @@protoc_insertion_point(module_scope)
