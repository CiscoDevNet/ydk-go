// This module implements entities for inter sdr connect
// feature
package sysadmin_sdr_connect

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_sdr_connect"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sdr-connect sdr-connect}", reflect.TypeOf(SdrConnect{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdr-connect:sdr-connect", reflect.TypeOf(SdrConnect{}))
}

// SdrConnect
type SdrConnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Connect two Secure Domain Routers.
    Connect SdrConnect_Connect
}

func (sdrConnect *SdrConnect) GetEntityData() *types.CommonEntityData {
    sdrConnect.EntityData.YFilter = sdrConnect.YFilter
    sdrConnect.EntityData.YangName = "sdr-connect"
    sdrConnect.EntityData.BundleName = "cisco_ios_xr"
    sdrConnect.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdr-connect"
    sdrConnect.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdr-connect:sdr-connect"
    sdrConnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrConnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrConnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrConnect.EntityData.Children = types.NewOrderedMap()
    sdrConnect.EntityData.Children.Append("connect", types.YChild{"Connect", &sdrConnect.Connect})
    sdrConnect.EntityData.Leafs = types.NewOrderedMap()

    sdrConnect.EntityData.YListKeys = []string {}

    return &(sdrConnect.EntityData)
}

// SdrConnect_Connect
// Connect two Secure Domain Routers
type SdrConnect_Connect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secure Domain Router connect config. The type is slice of
    // SdrConnect_Connect_Sdr.
    Sdr []*SdrConnect_Connect_Sdr
}

func (connect *SdrConnect_Connect) GetEntityData() *types.CommonEntityData {
    connect.EntityData.YFilter = connect.YFilter
    connect.EntityData.YangName = "connect"
    connect.EntityData.BundleName = "cisco_ios_xr"
    connect.EntityData.ParentYangName = "sdr-connect"
    connect.EntityData.SegmentPath = "connect"
    connect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connect.EntityData.Children = types.NewOrderedMap()
    connect.EntityData.Children.Append("sdr", types.YChild{"Sdr", nil})
    for i := range connect.Sdr {
        connect.EntityData.Children.Append(types.GetSegmentPath(connect.Sdr[i]), types.YChild{"Sdr", connect.Sdr[i]})
    }
    connect.EntityData.Leafs = types.NewOrderedMap()

    connect.EntityData.YListKeys = []string {}

    return &(connect.EntityData)
}

// SdrConnect_Connect_Sdr
// Secure Domain Router connect config
type SdrConnect_Connect_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Secure Domain Router , 30 max
    // characters. The type is string with pattern: [a-zA-Z0-9_-]{1,30}. This
    // attribute is mandatory.
    SdrName interface{}

    // This attribute is a key. Name of the Remote Secure Domain Router , 30 max
    // characters. The type is string with pattern: [a-zA-Z0-9_-]{1,30}. This
    // attribute is mandatory.
    RemoteSdrName interface{}

    // This attribute is a key. Index unique for each SDR connection pair. The
    // type is interface{} with range: 1..15. This attribute is mandatory.
    CsiId interface{}
}

func (sdr *SdrConnect_Connect_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "connect"
    sdr.EntityData.SegmentPath = "sdr" + types.AddKeyToken(sdr.SdrName, "sdr-name") + types.AddKeyToken(sdr.RemoteSdrName, "remote-sdr-name") + types.AddKeyToken(sdr.CsiId, "csi-id")
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = types.NewOrderedMap()
    sdr.EntityData.Leafs = types.NewOrderedMap()
    sdr.EntityData.Leafs.Append("sdr-name", types.YLeaf{"SdrName", sdr.SdrName})
    sdr.EntityData.Leafs.Append("remote-sdr-name", types.YLeaf{"RemoteSdrName", sdr.RemoteSdrName})
    sdr.EntityData.Leafs.Append("csi-id", types.YLeaf{"CsiId", sdr.CsiId})

    sdr.EntityData.YListKeys = []string {"SdrName", "RemoteSdrName", "CsiId"}

    return &(sdr.EntityData)
}

