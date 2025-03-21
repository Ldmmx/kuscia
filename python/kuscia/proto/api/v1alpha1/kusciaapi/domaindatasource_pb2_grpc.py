# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from kuscia.proto.api.v1alpha1.kusciaapi import domaindatasource_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2


class DomainDataSourceServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/CreateDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceResponse.FromString,
                )
        self.QueryDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/QueryDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.FromString,
                )
        self.UpdateDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/UpdateDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceResponse.FromString,
                )
        self.DeleteDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/DeleteDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceResponse.FromString,
                )
        self.BatchQueryDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/BatchQueryDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceResponse.FromString,
                )
        self.ListDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/ListDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceResponse.FromString,
                )


class DomainDataSourceServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchQueryDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DomainDataSourceServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceResponse.SerializeToString,
            ),
            'QueryDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.SerializeToString,
            ),
            'UpdateDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceResponse.SerializeToString,
            ),
            'DeleteDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceResponse.SerializeToString,
            ),
            'BatchQueryDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchQueryDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceResponse.SerializeToString,
            ),
            'ListDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.ListDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DomainDataSourceService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/CreateDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.CreateDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/QueryDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/UpdateDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.UpdateDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/DeleteDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.DeleteDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchQueryDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/BatchQueryDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.BatchQueryDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceService/ListDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_kusciaapi_dot_domaindatasource__pb2.ListDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
