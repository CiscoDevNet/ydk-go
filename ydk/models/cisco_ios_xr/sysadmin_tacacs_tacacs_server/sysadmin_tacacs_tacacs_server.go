// This module defines the TACACS+ data model.
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
    Host []TacacsServer_Host

    
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
    tacacsServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacsServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacsServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacsServer.EntityData.Children = make(map[string]types.YChild)
    tacacsServer.EntityData.Children["host"] = types.YChild{"Host", nil}
    for i := range tacacsServer.Host {
        tacacsServer.EntityData.Children[types.GetSegmentPath(&tacacsServer.Host[i])] = types.YChild{"Host", &tacacsServer.Host[i]}
    }
    tacacsServer.EntityData.Children["Cisco-IOS-XR-sysadmin-tacacs-show-tacacs:requests"] = types.YChild{"Requests", &tacacsServer.Requests}
    tacacsServer.EntityData.Children["Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authentication"] = types.YChild{"TestAuthentication", &tacacsServer.TestAuthentication}
    tacacsServer.EntityData.Children["Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authorization"] = types.YChild{"TestAuthorization", &tacacsServer.TestAuthorization}
    tacacsServer.EntityData.Children["Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-accounting"] = types.YChild{"TestAccounting", &tacacsServer.TestAccounting}
    tacacsServer.EntityData.Leafs = make(map[string]types.YLeaf)
    tacacsServer.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", tacacsServer.Timeout}
    tacacsServer.EntityData.Leafs["key"] = types.YLeaf{"Key", tacacsServer.Key}
    return &(tacacsServer.EntityData)
}

// TacacsServer_Host
type TacacsServer_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    host.EntityData.SegmentPath = "host" + "[ip='" + fmt.Sprintf("%v", host.Ip) + "']" + "[port='" + fmt.Sprintf("%v", host.Port) + "']"
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = make(map[string]types.YChild)
    host.EntityData.Leafs = make(map[string]types.YLeaf)
    host.EntityData.Leafs["ip"] = types.YLeaf{"Ip", host.Ip}
    host.EntityData.Leafs["port"] = types.YLeaf{"Port", host.Port}
    host.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", host.Timeout}
    host.EntityData.Leafs["key"] = types.YLeaf{"Key", host.Key}
    return &(host.EntityData)
}

// TacacsServer_Requests
type TacacsServer_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of TacacsServer_Requests_Ipv4.
    Ipv4 []TacacsServer_Requests_Ipv4
}

func (requests *TacacsServer_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "tacacs-server"
    requests.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-show-tacacs:requests"
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = make(map[string]types.YChild)
    requests.EntityData.Children["ipv4"] = types.YChild{"Ipv4", nil}
    for i := range requests.Ipv4 {
        requests.EntityData.Children[types.GetSegmentPath(&requests.Ipv4[i])] = types.YChild{"Ipv4", &requests.Ipv4[i]}
    }
    requests.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(requests.EntityData)
}

// TacacsServer_Requests_Ipv4
type TacacsServer_Requests_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Server IPv4 address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    ipv4.EntityData.SegmentPath = "ipv4" + "[addr='" + fmt.Sprintf("%v", ipv4.Addr) + "']" + "[port='" + fmt.Sprintf("%v", ipv4.Port) + "']"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["addr"] = types.YLeaf{"Addr", ipv4.Addr}
    ipv4.EntityData.Leafs["port"] = types.YLeaf{"Port", ipv4.Port}
    ipv4.EntityData.Leafs["opens"] = types.YLeaf{"Opens", ipv4.Opens}
    ipv4.EntityData.Leafs["closes"] = types.YLeaf{"Closes", ipv4.Closes}
    ipv4.EntityData.Leafs["aborts"] = types.YLeaf{"Aborts", ipv4.Aborts}
    ipv4.EntityData.Leafs["errors"] = types.YLeaf{"Errors", ipv4.Errors}
    ipv4.EntityData.Leafs["packets_in"] = types.YLeaf{"PacketsIn", ipv4.PacketsIn}
    ipv4.EntityData.Leafs["packets_out"] = types.YLeaf{"PacketsOut", ipv4.PacketsOut}
    return &(ipv4.EntityData)
}

// TacacsServer_TestAuthentication
// This type is a presence type.
type TacacsServer_TestAuthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication. The type is string.
    Authentication interface{}
}

func (testAuthentication *TacacsServer_TestAuthentication) GetEntityData() *types.CommonEntityData {
    testAuthentication.EntityData.YFilter = testAuthentication.YFilter
    testAuthentication.EntityData.YangName = "test-authentication"
    testAuthentication.EntityData.BundleName = "cisco_ios_xr"
    testAuthentication.EntityData.ParentYangName = "tacacs-server"
    testAuthentication.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authentication"
    testAuthentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAuthentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAuthentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAuthentication.EntityData.Children = make(map[string]types.YChild)
    testAuthentication.EntityData.Leafs = make(map[string]types.YLeaf)
    testAuthentication.EntityData.Leafs["authentication"] = types.YLeaf{"Authentication", testAuthentication.Authentication}
    return &(testAuthentication.EntityData)
}

// TacacsServer_TestAuthorization
// This type is a presence type.
type TacacsServer_TestAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authorization. The type is string.
    Authorization interface{}
}

func (testAuthorization *TacacsServer_TestAuthorization) GetEntityData() *types.CommonEntityData {
    testAuthorization.EntityData.YFilter = testAuthorization.YFilter
    testAuthorization.EntityData.YangName = "test-authorization"
    testAuthorization.EntityData.BundleName = "cisco_ios_xr"
    testAuthorization.EntityData.ParentYangName = "tacacs-server"
    testAuthorization.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-authorization"
    testAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAuthorization.EntityData.Children = make(map[string]types.YChild)
    testAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    testAuthorization.EntityData.Leafs["authorization"] = types.YLeaf{"Authorization", testAuthorization.Authorization}
    return &(testAuthorization.EntityData)
}

// TacacsServer_TestAccounting
// This type is a presence type.
type TacacsServer_TestAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting. The type is string.
    Accounting interface{}
}

func (testAccounting *TacacsServer_TestAccounting) GetEntityData() *types.CommonEntityData {
    testAccounting.EntityData.YFilter = testAccounting.YFilter
    testAccounting.EntityData.YangName = "test-accounting"
    testAccounting.EntityData.BundleName = "cisco_ios_xr"
    testAccounting.EntityData.ParentYangName = "tacacs-server"
    testAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-tacacs-test-tacacs:test-accounting"
    testAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    testAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    testAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    testAccounting.EntityData.Children = make(map[string]types.YChild)
    testAccounting.EntityData.Leafs = make(map[string]types.YLeaf)
    testAccounting.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", testAccounting.Accounting}
    return &(testAccounting.EntityData)
}

