// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ems package configuration.
// 
// This module contains definitions
// for the following management objects:
//   grpc: GRPC configruation
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_ems_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_ems_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-ems-cfg grpc}", reflect.TypeOf(Grpc{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-ems-cfg:grpc", reflect.TypeOf(Grpc{}))
}

// Grpc
// GRPC configruation
type Grpc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server listening port. The type is interface{} with range: 10000..57999.
    Port interface{}

    // Server vrf name. The type is string.
    Vrf interface{}

    // Enable GRPC. The type is interface{}.
    Enable interface{}

    // Maximum concurrent requests per user. The type is interface{} with range:
    // 1..32.
    MaxRequestPerUser interface{}

    // No TLS. The type is interface{}.
    NoTls interface{}

    // Trustpoint Name. The type is string.
    TlsTrustpoint interface{}

    // Address family identifier type. The type is string.
    AddressFamily interface{}

    // TLS mutual authentication. The type is interface{}.
    TlsMutual interface{}

    // Maximum concurrent requests in total. The type is interface{} with range:
    // 1..256.
    MaxRequestTotal interface{}

    // Service Layer.
    ServiceLayer Grpc_ServiceLayer
}

func (grpc *Grpc) GetEntityData() *types.CommonEntityData {
    grpc.EntityData.YFilter = grpc.YFilter
    grpc.EntityData.YangName = "grpc"
    grpc.EntityData.BundleName = "cisco_ios_xr"
    grpc.EntityData.ParentYangName = "Cisco-IOS-XR-man-ems-cfg"
    grpc.EntityData.SegmentPath = "Cisco-IOS-XR-man-ems-cfg:grpc"
    grpc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grpc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grpc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grpc.EntityData.Children = types.NewOrderedMap()
    grpc.EntityData.Children.Append("service-layer", types.YChild{"ServiceLayer", &grpc.ServiceLayer})
    grpc.EntityData.Leafs = types.NewOrderedMap()
    grpc.EntityData.Leafs.Append("port", types.YLeaf{"Port", grpc.Port})
    grpc.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", grpc.Vrf})
    grpc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", grpc.Enable})
    grpc.EntityData.Leafs.Append("max-request-per-user", types.YLeaf{"MaxRequestPerUser", grpc.MaxRequestPerUser})
    grpc.EntityData.Leafs.Append("no-tls", types.YLeaf{"NoTls", grpc.NoTls})
    grpc.EntityData.Leafs.Append("tls-trustpoint", types.YLeaf{"TlsTrustpoint", grpc.TlsTrustpoint})
    grpc.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", grpc.AddressFamily})
    grpc.EntityData.Leafs.Append("tls-mutual", types.YLeaf{"TlsMutual", grpc.TlsMutual})
    grpc.EntityData.Leafs.Append("max-request-total", types.YLeaf{"MaxRequestTotal", grpc.MaxRequestTotal})

    grpc.EntityData.YListKeys = []string {}

    return &(grpc.EntityData)
}

// Grpc_ServiceLayer
// Service Layer
type Grpc_ServiceLayer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ServiceLayer. The type is interface{}.
    Enable interface{}
}

func (serviceLayer *Grpc_ServiceLayer) GetEntityData() *types.CommonEntityData {
    serviceLayer.EntityData.YFilter = serviceLayer.YFilter
    serviceLayer.EntityData.YangName = "service-layer"
    serviceLayer.EntityData.BundleName = "cisco_ios_xr"
    serviceLayer.EntityData.ParentYangName = "grpc"
    serviceLayer.EntityData.SegmentPath = "service-layer"
    serviceLayer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceLayer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceLayer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceLayer.EntityData.Children = types.NewOrderedMap()
    serviceLayer.EntityData.Leafs = types.NewOrderedMap()
    serviceLayer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", serviceLayer.Enable})

    serviceLayer.EntityData.YListKeys = []string {}

    return &(serviceLayer.EntityData)
}

