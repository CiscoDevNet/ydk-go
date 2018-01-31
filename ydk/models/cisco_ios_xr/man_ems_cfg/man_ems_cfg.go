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
    parent types.Entity
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

    // Address family identifier type. The type is string.
    AddressFamily interface{}

    // Maximum concurrent requests in total. The type is interface{} with range:
    // 1..256.
    MaxRequestTotal interface{}

    // Service Layer.
    ServiceLayer Grpc_ServiceLayer

    // Transport Layer Security (TLS).
    Tls Grpc_Tls
}

func (grpc *Grpc) GetFilter() yfilter.YFilter { return grpc.YFilter }

func (grpc *Grpc) SetFilter(yf yfilter.YFilter) { grpc.YFilter = yf }

func (grpc *Grpc) GetGoName(yname string) string {
    if yname == "port" { return "Port" }
    if yname == "vrf" { return "Vrf" }
    if yname == "enable" { return "Enable" }
    if yname == "max-request-per-user" { return "MaxRequestPerUser" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "max-request-total" { return "MaxRequestTotal" }
    if yname == "service-layer" { return "ServiceLayer" }
    if yname == "tls" { return "Tls" }
    return ""
}

func (grpc *Grpc) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-ems-cfg:grpc"
}

func (grpc *Grpc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-layer" {
        return &grpc.ServiceLayer
    }
    if childYangName == "tls" {
        return &grpc.Tls
    }
    return nil
}

func (grpc *Grpc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-layer"] = &grpc.ServiceLayer
    children["tls"] = &grpc.Tls
    return children
}

func (grpc *Grpc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port"] = grpc.Port
    leafs["vrf"] = grpc.Vrf
    leafs["enable"] = grpc.Enable
    leafs["max-request-per-user"] = grpc.MaxRequestPerUser
    leafs["address-family"] = grpc.AddressFamily
    leafs["max-request-total"] = grpc.MaxRequestTotal
    return leafs
}

func (grpc *Grpc) GetBundleName() string { return "cisco_ios_xr" }

func (grpc *Grpc) GetYangName() string { return "grpc" }

func (grpc *Grpc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (grpc *Grpc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (grpc *Grpc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (grpc *Grpc) SetParent(parent types.Entity) { grpc.parent = parent }

func (grpc *Grpc) GetParent() types.Entity { return grpc.parent }

func (grpc *Grpc) GetParentYangName() string { return "Cisco-IOS-XR-man-ems-cfg" }

// Grpc_ServiceLayer
// Service Layer
type Grpc_ServiceLayer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ServiceLayer. The type is interface{}.
    Enable interface{}
}

func (serviceLayer *Grpc_ServiceLayer) GetFilter() yfilter.YFilter { return serviceLayer.YFilter }

func (serviceLayer *Grpc_ServiceLayer) SetFilter(yf yfilter.YFilter) { serviceLayer.YFilter = yf }

func (serviceLayer *Grpc_ServiceLayer) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (serviceLayer *Grpc_ServiceLayer) GetSegmentPath() string {
    return "service-layer"
}

func (serviceLayer *Grpc_ServiceLayer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceLayer *Grpc_ServiceLayer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceLayer *Grpc_ServiceLayer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = serviceLayer.Enable
    return leafs
}

func (serviceLayer *Grpc_ServiceLayer) GetBundleName() string { return "cisco_ios_xr" }

func (serviceLayer *Grpc_ServiceLayer) GetYangName() string { return "service-layer" }

func (serviceLayer *Grpc_ServiceLayer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceLayer *Grpc_ServiceLayer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceLayer *Grpc_ServiceLayer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceLayer *Grpc_ServiceLayer) SetParent(parent types.Entity) { serviceLayer.parent = parent }

func (serviceLayer *Grpc_ServiceLayer) GetParent() types.Entity { return serviceLayer.parent }

func (serviceLayer *Grpc_ServiceLayer) GetParentYangName() string { return "grpc" }

// Grpc_Tls
// Transport Layer Security (TLS)
type Grpc_Tls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trustpoint Name. The type is string.
    Trustpoint interface{}

    // Enable TLS. The type is interface{}.
    Enable interface{}
}

func (tls *Grpc_Tls) GetFilter() yfilter.YFilter { return tls.YFilter }

func (tls *Grpc_Tls) SetFilter(yf yfilter.YFilter) { tls.YFilter = yf }

func (tls *Grpc_Tls) GetGoName(yname string) string {
    if yname == "trustpoint" { return "Trustpoint" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (tls *Grpc_Tls) GetSegmentPath() string {
    return "tls"
}

func (tls *Grpc_Tls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tls *Grpc_Tls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tls *Grpc_Tls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trustpoint"] = tls.Trustpoint
    leafs["enable"] = tls.Enable
    return leafs
}

func (tls *Grpc_Tls) GetBundleName() string { return "cisco_ios_xr" }

func (tls *Grpc_Tls) GetYangName() string { return "tls" }

func (tls *Grpc_Tls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tls *Grpc_Tls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tls *Grpc_Tls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tls *Grpc_Tls) SetParent(parent types.Entity) { tls.parent = parent }

func (tls *Grpc_Tls) GetParent() types.Entity { return tls.parent }

func (tls *Grpc_Tls) GetParentYangName() string { return "grpc" }

