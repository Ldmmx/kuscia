��                                                           !     $     )     .  
   3     >  
   C     N  `   b     �     �  H   �  �     �   �     �  H   �     �  	   	  }     O   �  �   �  �   �     ~     �  z  �     	     	     	     "	     %	     (	     -	     0	     5	     :	  
   ?	     J	  
   O	     Z	  �   w	      
     	
  a   
  �   n
  �   0        [     %   `     �  s   �  Q   �  �   Q  �                -g -k -p -q -x 1080 80 8082 8083 9091 GRPC/GRPCS HTTP HTTP/HTTPS Kuscia 端口介绍 下面是需要用户感知的端口信息，按是否需要暴露给合作方，分为两类： 协议 否 否。该类端口仅需局域网内使用，无需暴露给合作方。 在实际场景中，为了确保 Kuscia 运行在一个安全的网络环境中，用户需要根据本地网络防火墙，管理 Kuscia 对外暴露给合作机构的端口。 如果采用 Docker 方式部署，那么在部署的过程中，为了能够跨域和域内访问 Kuscia 节点，需要将 Kuscia 节点的部分内部端口通过端口映射的方式暴露在宿主机上。 是 是。该类端口需要通过公网或专线能够让合作方访问。 是否需要暴露给合作方 端口号 节点 KusciaAPI 的访问端口，可参考[如何使用 KusciaAPI](../reference/apis/summary_cn.md#如何使用-kuscia-api) 节点 Metrics 指标采集端口，可参考 [Kuscia 监控](./kuscia_monitor) 节点之间的认证鉴权端口。在创建节点之间路由时需要指定，可参考[创建节点路由](../reference/apis/domainroute_cn.md#请求createdomainrouterequest) 访问节点中应用的端口。例如：可通过此端口访问 Serving 服务进行预测打分，可参考[使用 SecretFlow Serving 进行预测](../tutorial/run_sf_serving_with_api_cn.md#使用-secretflow-serving-进行预测) 说明 部署脚本对应参数 Project-Id-Version: Kuscia 
Report-Msgid-Bugs-To: 
POT-Creation-Date: 2025-03-10 20:35+0800
PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE
Last-Translator: FULL NAME <EMAIL@ADDRESS>
Language: en
Language-Team: en <LL@li.org>
Plural-Forms: nplurals=2; plural=(n != 1);
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 8bit
Generated-By: Babel 2.17.0
 -g -k -p -q -x 1080 80 8082 8083 9091 GRPC/GRPCS HTTP HTTP/HTTPS Introduction to Kuscia Ports The following port information requires user awareness, categorized into two types based on whether they need to be exposed to partners: Protocol No No. These ports are only used within the local network and do not need to be exposed to partners. In practical scenarios, to ensure Kuscia operates in a secure network environment, users need to manage the ports exposed by Kuscia to partner organizations based on the local network firewall. If deploying via Docker, to enable cross-domain and intra-domain access to Kuscia nodes, certain internal ports of the Kuscia nodes must be exposed on the host machine through port mapping during deployment. Yes Yes. These ports must be accessible to partners via the public internet or dedicated lines. Whether need to be exposed to partner Port Access port for the node's KusciaAPI. Refer to [How to Use KusciaAPI](../reference/apis/summary_cn.md#-kuscia-api). Port for collecting node metrics. Refer to [Kuscia Monitoring](./kuscia_monitor). Authentication and authorization port between nodes. Must be specified when creating inter-node routes. Refer to [Create DomianRoute](../reference/apis/domainroute_cn.md#createdomainrouterequest). Port for accessing applications on the node. For example, this port can be used to access Serving services for prediction scoring. Refer to [Using SecretFlow Serving for Prediction](../tutorial/run_sf_serving_with_api_cn.md#-secretflow-serving-). Description Deployment Script Parameter 