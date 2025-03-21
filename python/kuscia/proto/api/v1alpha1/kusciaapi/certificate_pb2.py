# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/certificate.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n5kuscia/proto/api/v1alpha1/kusciaapi/certificate.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"\xd4\x01\n\x17GenerateKeyCertsRequest\x12\x13\n\x0b\x63ommon_name\x18\x01 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x02 \x01(\t\x12\x14\n\x0corganization\x18\x03 \x01(\t\x12\x19\n\x11organization_unit\x18\x04 \x01(\t\x12\x10\n\x08locality\x18\x05 \x01(\t\x12\x10\n\x08province\x18\x06 \x01(\t\x12\x16\n\x0estreet_address\x18\x07 \x01(\t\x12\x14\n\x0c\x64uration_sec\x18\x08 \x01(\x03\x12\x10\n\x08key_type\x18\t \x01(\t\"n\n\x18GenerateKeyCertsResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x0b\n\x03key\x18\x02 \x01(\t\x12\x12\n\ncert_chain\x18\x03 \x03(\t2\xa6\x01\n\x12\x43\x65rtificateService\x12\x8f\x01\n\x10GenerateKeyCerts\x12<.kuscia.proto.api.v1alpha1.kusciaapi.GenerateKeyCertsRequest\x1a=.kuscia.proto.api.v1alpha1.kusciaapi.GenerateKeyCertsResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.kusciaapi.certificate_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _globals['_GENERATEKEYCERTSREQUEST']._serialized_start=135
  _globals['_GENERATEKEYCERTSREQUEST']._serialized_end=347
  _globals['_GENERATEKEYCERTSRESPONSE']._serialized_start=349
  _globals['_GENERATEKEYCERTSRESPONSE']._serialized_end=459
  _globals['_CERTIFICATESERVICE']._serialized_start=462
  _globals['_CERTIFICATESERVICE']._serialized_end=628
# @@protoc_insertion_point(module_scope)
