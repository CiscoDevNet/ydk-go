// This module contains definitions
// for the Calvados model objects.
// 
// This module defines the TACACS+ data model.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_tacacs_tacacs_server

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_tacacs_tacacs_server"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-tacacs-tacacs-server tacacs-server}", reflect.TypeOf(TacacsServer{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server", reflect.TypeOf(TacacsServer{}))
}

// TacacsServer
type TacacsServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 1..1000.
    Timeout interface{}

    // The type is string.
    Key interface{}

    // The type is slice of TacacsServer_Host.
    Host []*TacacsServer_Host

    
    Requests TacacsServer_Requests

    
    TestAuthentication TacacsServer_TestAuthentication

    
    TestAuthorization TacacsServer_TestAuthorization

    
    TestAccounting TacacsServer_TestAccounting
}

func (tacacsServer *TacacsServer) GetEntityData() *types.CommonEntityData {
    tacacsServer.EntityData.YFilter = tacacsServer.YFilter
    tacacsServer.EntityData.YangName = "tacacs-server"
    tacacsServer.EntityData.BundleName = "cisco_ios_xr"
    tacacsServer.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server"
    tacacsServer.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server"
    tacacsServer.EntityData.AbsolutePath = tacacsServer.EntityData.SegmentPath
    tacacsServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacsServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacsServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacsServer.EntityData.Children = types.NewOrderedMap()
    tacacsServer.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range tacacsServer.Host {
        tacacsServer.EntityData.Children.Append(types.GetSegmentPath(tacacsServer.Host[i]), types.YChild{"Host", tacacsServer.Host[i]})
    }
    tacacsServer.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-tacacs-show-tacacs:requests", types.YChild{"Requests", &tacacsServer.Requests})
    tacacsServer.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authentication", types.YChild{"TestAuthentication", &tacacsServer.TestAuthentication})
    tacacsServer.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authorization", types.YChild{"TestAuthorization", &tacacsServer.TestAuthorization})
    tacacsServer.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-accounting", types.YChild{"TestAccounting", &tacacsServer.TestAccounting})
    tacacsServer.EntityData.Leafs = types.NewOrderedMap()
    tacacsServer.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", tacacsServer.Timeout})
    tacacsServer.EntityData.Leafs.Append("key", types.YLeaf{"Key", tacacsServer.Key})

    tacacsServer.EntityData.YListKeys = []string {}

    return &(tacacsServer.EntityData)
}

// TacacsServer_Host
type TacacsServer_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // This attribute is a key. The type is interface{} with range: 1..65535.
    Port interface{}

    // The type is interface{} with range: 1..1000.
    Timeout interface{}

    // The type is string.
    Key interface{}
}

func (host *TacacsServer_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "tacacs-server"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.Ip, "ip") + types.AddKeyToken(host.Port, "port")
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", host.Ip})
    host.EntityData.Leafs.Append("port", types.YLeaf{"Port", host.Port})
    host.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", host.Timeout})
    host.EntityData.Leafs.Append("key", types.YLeaf{"Key", host.Key})

    host.EntityData.YListKeys = []string {"Ip", "Port"}

    return &(host.EntityData)
}

// TacacsServer_Requests
type TacacsServer_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of TacacsServer_Requests_Ipv4.
    Ipv4 []*TacacsServer_Requests_Ipv4
}

func (requests *TacacsServer_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "tacacs-server"
    requests.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-show-tacacs:requests"
    requests.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/" + requests.EntityData.SegmentPath
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = types.NewOrderedMap()
    requests.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", nil})
    for i := range requests.Ipv4 {
        requests.EntityData.Children.Append(types.GetSegmentPath(requests.Ipv4[i]), types.YChild{"Ipv4", requests.Ipv4[i]})
    }
    requests.EntityData.Leafs = types.NewOrderedMap()

    requests.EntityData.YListKeys = []string {}

    return &(requests.EntityData)
}

// TacacsServer_Requests_Ipv4
type TacacsServer_Requests_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Server IPv4 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Addr interface{}

    // This attribute is a key. Server Port. The type is interface{} with range:
    // 0..65535.
    Port interface{}

    // Socket open count. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // Socket close count. The type is interface{} with range: 0..4294967295.
    Closes interface{}

    // Socket abort count. The type is interface{} with range: 0..4294967295.
    Aborts interface{}

    // Socket error count. The type is interface{} with range: 0..4294967295.
    Errors interface{}

    // Packets received count. The type is interface{} with range: 0..4294967295.
    PacketsIn interface{}

    // Packets transmitted count. The type is interface{} with range:
    // 0..4294967295.
    PacketsOut interface{}
}

func (ipv4 *TacacsServer_Requests_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "requests"
    ipv4.EntityData.SegmentPath = "ipv4" + types.AddKeyToken(ipv4.Addr, "addr") + types.AddKeyToken(ipv4.Port, "port")
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/Cisco-IOS-XR-sysadmin-tacacs-show-tacacs:requests/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", ipv4.Addr})
    ipv4.EntityData.Leafs.Append("port", types.YLeaf{"Port", ipv4.Port})
    ipv4.EntityData.Leafs.Append("opens", types.YLeaf{"Opens", ipv4.Opens})
    ipv4.EntityData.Leafs.Append("closes", types.YLeaf{"Closes", ipv4.Closes})
    ipv4.EntityData.Leafs.Append("aborts", types.YLeaf{"Aborts", ipv4.Aborts})
    ipv4.EntityData.Leafs.Append("errors", types.YLeaf{"Errors", ipv4.Errors})
    ipv4.EntityData.Leafs.Append("packets_in", types.YLeaf{"PacketsIn", ipv4.PacketsIn})
    ipv4.EntityData.Leafs.Append("packets_out", types.YLeaf{"PacketsOut", ipv4.PacketsOut})

    ipv4.EntityData.YListKeys = []string {"Addr", "Port"}

    return &(ipv4.EntityData)
}

// TacacsServer_TestAuthentication
// This type is a presence type.
type TacacsServer_TestAuthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Authentication. The type is string.
    Authentication interface{}
}

func (testAuthentication *TacacsServer_TestAuthentication) GetEntityData() *types.CommonEntityData {
    testAuthentication.EntityData.YFilter = testAuthentication.YFilter
    testAuthentication.EntityData.YangName = "test-authentication"
    testAuthentication.EntityData.BundleName = "cisco_ios_xr"
    testAuthentication.EntityData.ParentYangName = "tacacs-server"
    testAuthentication.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authentication"
    testAuthentication.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/" + testAuthentication.EntityData.SegmentPath
    testAuthentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAuthentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAuthentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAuthentication.EntityData.Children = types.NewOrderedMap()
    testAuthentication.EntityData.Leafs = types.NewOrderedMap()
    testAuthentication.EntityData.Leafs.Append("authentication", types.YLeaf{"Authentication", testAuthentication.Authentication})

    testAuthentication.EntityData.YListKeys = []string {}

    return &(testAuthentication.EntityData)
}

// TacacsServer_TestAuthorization
// This type is a presence type.
type TacacsServer_TestAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Authorization. The type is string.
    Authorization interface{}
}

func (testAuthorization *TacacsServer_TestAuthorization) GetEntityData() *types.CommonEntityData {
    testAuthorization.EntityData.YFilter = testAuthorization.YFilter
    testAuthorization.EntityData.YangName = "test-authorization"
    testAuthorization.EntityData.BundleName = "cisco_ios_xr"
    testAuthorization.EntityData.ParentYangName = "tacacs-server"
    testAuthorization.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authorization"
    testAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/" + testAuthorization.EntityData.SegmentPath
    testAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAuthorization.EntityData.Children = types.NewOrderedMap()
    testAuthorization.EntityData.Leafs = types.NewOrderedMap()
    testAuthorization.EntityData.Leafs.Append("authorization", types.YLeaf{"Authorization", testAuthorization.Authorization})

    testAuthorization.EntityData.YListKeys = []string {}

    return &(testAuthorization.EntityData)
}

// TacacsServer_TestAccounting
// This type is a presence type.
type TacacsServer_TestAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Accounting. The type is string.
    Accounting interface{}
}

func (testAccounting *TacacsServer_TestAccounting) GetEntityData() *types.CommonEntityData {
    testAccounting.EntityData.YFilter = testAccounting.YFilter
    testAccounting.EntityData.YangName = "test-accounting"
    testAccounting.EntityData.BundleName = "cisco_ios_xr"
    testAccounting.EntityData.ParentYangName = "tacacs-server"
    testAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-accounting"
    testAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-tacacs-tacacs-server:tacacs-server/" + testAccounting.EntityData.SegmentPath
    testAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAccounting.EntityData.Children = types.NewOrderedMap()
    testAccounting.EntityData.Leafs = types.NewOrderedMap()
    testAccounting.EntityData.Leafs.Append("accounting", types.YLeaf{"Accounting", testAccounting.Accounting})

    testAccounting.EntityData.YListKeys = []string {}

    return &(testAccounting.EntityData)
}

