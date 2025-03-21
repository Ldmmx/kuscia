# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/domain.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0kuscia/proto/api/v1alpha1/kusciaapi/domain.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"\xde\x01\n\x13\x43reateDomainRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04role\x18\x03 \x01(\t\x12\x0c\n\x04\x63\x65rt\x18\x04 \x01(\t\x12\x44\n\x0b\x61uth_center\x18\x05 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.AuthCenter\x12\x18\n\x10master_domain_id\x18\x06 \x01(\t\"I\n\x14\x43reateDomainResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"b\n\x13\x44\x65leteDomainRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"I\n\x14\x44\x65leteDomainResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"a\n\x12QueryDomainRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"\x94\x01\n\x13QueryDomainResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12J\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32<.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainResponseData\"\xdf\x03\n\x17QueryDomainResponseData\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x0c\n\x04role\x18\x02 \x01(\t\x12\x0c\n\x04\x63\x65rt\x18\x03 \x01(\t\x12\x46\n\rnode_statuses\x18\x04 \x03(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.NodeStatus\x12U\n\x15\x64\x65ploy_token_statuses\x18\x05 \x03(\x0b\x32\x36.kuscia.proto.api.v1alpha1.kusciaapi.DeployTokenStatus\x12\x62\n\x0b\x61nnotations\x18\x06 \x03(\x0b\x32M.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainResponseData.AnnotationsEntry\x12\x44\n\x0b\x61uth_center\x18\x07 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.AuthCenter\x12\x18\n\x10master_domain_id\x18\x08 \x01(\t\x1a\x32\n\x10\x41nnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"v\n\nNodeStatus\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x1b\n\x13last_heartbeat_time\x18\x04 \x01(\t\x12\x1c\n\x14last_transition_time\x18\x05 \x01(\t\"\xde\x01\n\x13UpdateDomainRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04role\x18\x03 \x01(\t\x12\x0c\n\x04\x63\x65rt\x18\x04 \x01(\t\x12\x44\n\x0b\x61uth_center\x18\x05 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.AuthCenter\x12\x18\n\x10master_domain_id\x18\x06 \x01(\t\"I\n\x14UpdateDomainResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"g\n\x17\x42\x61tchQueryDomainRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x12\n\ndomain_ids\x18\x02 \x03(\t\"\x9e\x01\n\x18\x42\x61tchQueryDomainResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12O\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x41.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainResponseData\"\\\n\x1c\x42\x61tchQueryDomainResponseData\x12<\n\x07\x64omains\x18\x01 \x03(\x0b\x32+.kuscia.proto.api.v1alpha1.kusciaapi.Domain\"\x7f\n\x06\x44omain\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x0c\n\x04role\x18\x02 \x01(\t\x12\x0c\n\x04\x63\x65rt\x18\x03 \x01(\t\x12\x46\n\rnode_statuses\x18\x04 \x03(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.NodeStatus\"O\n\x11\x44\x65ployTokenStatus\x12\r\n\x05token\x18\x01 \x01(\t\x12\r\n\x05state\x18\x02 \x01(\t\x12\x1c\n\x14last_transition_time\x18\x03 \x01(\t\"C\n\nAuthCenter\x12\x1b\n\x13\x61uthentication_type\x18\x01 \x01(\t\x12\x18\n\x10token_gen_method\x18\x02 \x01(\t2\xb6\x05\n\rDomainService\x12\x83\x01\n\x0c\x43reateDomain\x12\x38.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRequest\x1a\x39.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainResponse\x12\x80\x01\n\x0bQueryDomain\x12\x37.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRequest\x1a\x38.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainResponse\x12\x83\x01\n\x0cUpdateDomain\x12\x38.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainRequest\x1a\x39.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainResponse\x12\x83\x01\n\x0c\x44\x65leteDomain\x12\x38.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainRequest\x1a\x39.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainResponse\x12\x8f\x01\n\x10\x42\x61tchQueryDomain\x12<.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRequest\x1a=.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.kusciaapi.domain_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _globals['_QUERYDOMAINRESPONSEDATA_ANNOTATIONSENTRY']._options = None
  _globals['_QUERYDOMAINRESPONSEDATA_ANNOTATIONSENTRY']._serialized_options = b'8\001'
  _globals['_CREATEDOMAINREQUEST']._serialized_start=130
  _globals['_CREATEDOMAINREQUEST']._serialized_end=352
  _globals['_CREATEDOMAINRESPONSE']._serialized_start=354
  _globals['_CREATEDOMAINRESPONSE']._serialized_end=427
  _globals['_DELETEDOMAINREQUEST']._serialized_start=429
  _globals['_DELETEDOMAINREQUEST']._serialized_end=527
  _globals['_DELETEDOMAINRESPONSE']._serialized_start=529
  _globals['_DELETEDOMAINRESPONSE']._serialized_end=602
  _globals['_QUERYDOMAINREQUEST']._serialized_start=604
  _globals['_QUERYDOMAINREQUEST']._serialized_end=701
  _globals['_QUERYDOMAINRESPONSE']._serialized_start=704
  _globals['_QUERYDOMAINRESPONSE']._serialized_end=852
  _globals['_QUERYDOMAINRESPONSEDATA']._serialized_start=855
  _globals['_QUERYDOMAINRESPONSEDATA']._serialized_end=1334
  _globals['_QUERYDOMAINRESPONSEDATA_ANNOTATIONSENTRY']._serialized_start=1284
  _globals['_QUERYDOMAINRESPONSEDATA_ANNOTATIONSENTRY']._serialized_end=1334
  _globals['_NODESTATUS']._serialized_start=1336
  _globals['_NODESTATUS']._serialized_end=1454
  _globals['_UPDATEDOMAINREQUEST']._serialized_start=1457
  _globals['_UPDATEDOMAINREQUEST']._serialized_end=1679
  _globals['_UPDATEDOMAINRESPONSE']._serialized_start=1681
  _globals['_UPDATEDOMAINRESPONSE']._serialized_end=1754
  _globals['_BATCHQUERYDOMAINREQUEST']._serialized_start=1756
  _globals['_BATCHQUERYDOMAINREQUEST']._serialized_end=1859
  _globals['_BATCHQUERYDOMAINRESPONSE']._serialized_start=1862
  _globals['_BATCHQUERYDOMAINRESPONSE']._serialized_end=2020
  _globals['_BATCHQUERYDOMAINRESPONSEDATA']._serialized_start=2022
  _globals['_BATCHQUERYDOMAINRESPONSEDATA']._serialized_end=2114
  _globals['_DOMAIN']._serialized_start=2116
  _globals['_DOMAIN']._serialized_end=2243
  _globals['_DEPLOYTOKENSTATUS']._serialized_start=2245
  _globals['_DEPLOYTOKENSTATUS']._serialized_end=2324
  _globals['_AUTHCENTER']._serialized_start=2326
  _globals['_AUTHCENTER']._serialized_end=2393
  _globals['_DOMAINSERVICE']._serialized_start=2396
  _globals['_DOMAINSERVICE']._serialized_end=3090
# @@protoc_insertion_point(module_scope)
