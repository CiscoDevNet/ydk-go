// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ems package configuration.
// 
// This module contains definitions
// for the following management objects:
//   grpc: GRPC configruation
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// Dscp represents Dscp
type Dscp string

const (
    // Default (DSCP 000000)
    Dscp_default_ Dscp = "default"

    // AF11 (DSCP 001010)
    Dscp_af11 Dscp = "af11"

    // AF12 (DSCP 001100)
    Dscp_af12 Dscp = "af12"

    // AF13 (DSCP 001110)
    Dscp_af13 Dscp = "af13"

    // AF21 (DSCP 010010)
    Dscp_af21 Dscp = "af21"

    // AF22 (DSCP 010100)
    Dscp_af22 Dscp = "af22"

    // AF23 (DSCP 010110)
    Dscp_af23 Dscp = "af23"

    // AF31 (DSCP 011010)
    Dscp_af31 Dscp = "af31"

    // AF32 (DSCP 011100)
    Dscp_af32 Dscp = "af32"

    // AF33 (DSCP 011110)
    Dscp_af33 Dscp = "af33"

    // AF41 (DSCP 100010)
    Dscp_af41 Dscp = "af41"

    // AF42 (DSCP 100100)
    Dscp_af42 Dscp = "af42"

    // AF43 (DSCP 100110)
    Dscp_af43 Dscp = "af43"

    // CS1 (Precedence 1) (DSCP 001000)
    Dscp_cs1 Dscp = "cs1"

    // CS2 (Precedence 2) (DSCP 010000)
    Dscp_cs2 Dscp = "cs2"

    // CS3 (Precedence 3) (DSCP 011000)
    Dscp_cs3 Dscp = "cs3"

    // CS4 (Precedence 4) (DSCP 100000)
    Dscp_cs4 Dscp = "cs4"

    // CS5 (Precedence 5) (DSCP 101000)
    Dscp_cs5 Dscp = "cs5"

    // CS6 (Precedence 6) (DSCP 110000)
    Dscp_cs6 Dscp = "cs6"

    // CS7 (Precedence 7) (DSCP 111000)
    Dscp_cs7 Dscp = "cs7"

    // EF (DSCP 101110)
    Dscp_ef Dscp = "ef"
)

// GrpCTlsCipherDefault represents Grp c tls cipher default
type GrpCTlsCipherDefault string

const (
    // Default disable all ciphers
    GrpCTlsCipherDefault_disable GrpCTlsCipherDefault = "disable"

    // Default enable all ciphers
    GrpCTlsCipherDefault_enable GrpCTlsCipherDefault = "enable"
)

// Grpc
// GRPC configruation
type Grpc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server listening port. The type is interface{} with range: 10000..57999.
    Port interface{}

    // Server vrf name. The type is string.
    Vrf interface{}

    // Maximum number of streaming gRPCs. The type is interface{} with range:
    // 1..128. The default value is 32.
    MaxStreams interface{}

    // Enable GRPC. The type is interface{}.
    Enable interface{}

    // Maximum number of streaming gRPCs per user. The type is interface{} with
    // range: 1..128. The default value is 32.
    MaxStreamsPerUser interface{}

    // Maximum concurrent requests per user. The type is interface{} with range:
    // 1..32.
    MaxRequestPerUser interface{}

    // No TLS. The type is interface{}.
    NoTls interface{}

    // Trustpoint Name. The type is string.
    TlsTrustpoint interface{}

    // QoS marking DSCP to be set on transmitted gRPC. The type is one of the
    // following types: enumeration Dscp, or int with range: 0..63.
    Dscp interface{}

    // Address family identifier type. The type is string.
    AddressFamily interface{}

    // TLS mutual authentication. The type is interface{}.
    TlsMutual interface{}

    // Maximum concurrent requests in total. The type is interface{} with range:
    // 1..256.
    MaxRequestTotal interface{}

    // Service Layer.
    ServiceLayer Grpc_ServiceLayer

    // TLS ciphers.
    TlsCipher Grpc_TlsCipher

    // Transport Layer Security (TLS).
    Tls Grpc_Tls
}

func (grpc *Grpc) GetEntityData() *types.CommonEntityData {
    grpc.EntityData.YFilter = grpc.YFilter
    grpc.EntityData.YangName = "grpc"
    grpc.EntityData.BundleName = "cisco_ios_xr"
    grpc.EntityData.ParentYangName = "Cisco-IOS-XR-man-ems-cfg"
    grpc.EntityData.SegmentPath = "Cisco-IOS-XR-man-ems-cfg:grpc"
    grpc.EntityData.AbsolutePath = grpc.EntityData.SegmentPath
    grpc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grpc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grpc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grpc.EntityData.Children = types.NewOrderedMap()
    grpc.EntityData.Children.Append("service-layer", types.YChild{"ServiceLayer", &grpc.ServiceLayer})
    grpc.EntityData.Children.Append("tls-cipher", types.YChild{"TlsCipher", &grpc.TlsCipher})
    grpc.EntityData.Children.Append("tls", types.YChild{"Tls", &grpc.Tls})
    grpc.EntityData.Leafs = types.NewOrderedMap()
    grpc.EntityData.Leafs.Append("port", types.YLeaf{"Port", grpc.Port})
    grpc.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", grpc.Vrf})
    grpc.EntityData.Leafs.Append("max-streams", types.YLeaf{"MaxStreams", grpc.MaxStreams})
    grpc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", grpc.Enable})
    grpc.EntityData.Leafs.Append("max-streams-per-user", types.YLeaf{"MaxStreamsPerUser", grpc.MaxStreamsPerUser})
    grpc.EntityData.Leafs.Append("max-request-per-user", types.YLeaf{"MaxRequestPerUser", grpc.MaxRequestPerUser})
    grpc.EntityData.Leafs.Append("no-tls", types.YLeaf{"NoTls", grpc.NoTls})
    grpc.EntityData.Leafs.Append("tls-trustpoint", types.YLeaf{"TlsTrustpoint", grpc.TlsTrustpoint})
    grpc.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", grpc.Dscp})
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
    serviceLayer.EntityData.AbsolutePath = "Cisco-IOS-XR-man-ems-cfg:grpc/" + serviceLayer.EntityData.SegmentPath
    serviceLayer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceLayer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceLayer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceLayer.EntityData.Children = types.NewOrderedMap()
    serviceLayer.EntityData.Leafs = types.NewOrderedMap()
    serviceLayer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", serviceLayer.Enable})

    serviceLayer.EntityData.YListKeys = []string {}

    return &(serviceLayer.EntityData)
}

// Grpc_TlsCipher
// TLS ciphers
type Grpc_TlsCipher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default of all ciphers. The type is GrpCTlsCipherDefault.
    Default interface{}

    // Enable ciphers if default is disabled. The type is string.
    Enable interface{}

    // Disable ciphers if default is enabled. The type is string.
    Disable interface{}
}

func (tlsCipher *Grpc_TlsCipher) GetEntityData() *types.CommonEntityData {
    tlsCipher.EntityData.YFilter = tlsCipher.YFilter
    tlsCipher.EntityData.YangName = "tls-cipher"
    tlsCipher.EntityData.BundleName = "cisco_ios_xr"
    tlsCipher.EntityData.ParentYangName = "grpc"
    tlsCipher.EntityData.SegmentPath = "tls-cipher"
    tlsCipher.EntityData.AbsolutePath = "Cisco-IOS-XR-man-ems-cfg:grpc/" + tlsCipher.EntityData.SegmentPath
    tlsCipher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlsCipher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlsCipher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlsCipher.EntityData.Children = types.NewOrderedMap()
    tlsCipher.EntityData.Leafs = types.NewOrderedMap()
    tlsCipher.EntityData.Leafs.Append("default", types.YLeaf{"Default", tlsCipher.Default})
    tlsCipher.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tlsCipher.Enable})
    tlsCipher.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", tlsCipher.Disable})

    tlsCipher.EntityData.YListKeys = []string {}

    return &(tlsCipher.EntityData)
}

// Grpc_Tls
// Transport Layer Security (TLS)
type Grpc_Tls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable TLS. The type is interface{}.
    Enable interface{}
}

func (tls *Grpc_Tls) GetEntityData() *types.CommonEntityData {
    tls.EntityData.YFilter = tls.YFilter
    tls.EntityData.YangName = "tls"
    tls.EntityData.BundleName = "cisco_ios_xr"
    tls.EntityData.ParentYangName = "grpc"
    tls.EntityData.SegmentPath = "tls"
    tls.EntityData.AbsolutePath = "Cisco-IOS-XR-man-ems-cfg:grpc/" + tls.EntityData.SegmentPath
    tls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tls.EntityData.Children = types.NewOrderedMap()
    tls.EntityData.Leafs = types.NewOrderedMap()
    tls.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tls.Enable})

    tls.EntityData.YListKeys = []string {}

    return &(tls.EntityData)
}

