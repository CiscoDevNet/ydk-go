// This module contains a collection of YANG definitions
// for virtual service configurational data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package virtual_service_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package virtual_service_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-virtual-service-cfg virtual-service-cfg-data}", reflect.TypeOf(VirtualServiceCfgData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data", reflect.TypeOf(VirtualServiceCfgData{}))
}

// VirtualServiceCfgData
// Virtual Service configurational data
type VirtualServiceCfgData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application list.
    Apps VirtualServiceCfgData_Apps

    // App-hosting control.
    Controls VirtualServiceCfgData_Controls
}

func (virtualServiceCfgData *VirtualServiceCfgData) GetEntityData() *types.CommonEntityData {
    virtualServiceCfgData.EntityData.YFilter = virtualServiceCfgData.YFilter
    virtualServiceCfgData.EntityData.YangName = "virtual-service-cfg-data"
    virtualServiceCfgData.EntityData.BundleName = "cisco_ios_xe"
    virtualServiceCfgData.EntityData.ParentYangName = "Cisco-IOS-XE-virtual-service-cfg"
    virtualServiceCfgData.EntityData.SegmentPath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data"
    virtualServiceCfgData.EntityData.AbsolutePath = virtualServiceCfgData.EntityData.SegmentPath
    virtualServiceCfgData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    virtualServiceCfgData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    virtualServiceCfgData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    virtualServiceCfgData.EntityData.Children = types.NewOrderedMap()
    virtualServiceCfgData.EntityData.Children.Append("apps", types.YChild{"Apps", &virtualServiceCfgData.Apps})
    virtualServiceCfgData.EntityData.Children.Append("controls", types.YChild{"Controls", &virtualServiceCfgData.Controls})
    virtualServiceCfgData.EntityData.Leafs = types.NewOrderedMap()

    virtualServiceCfgData.EntityData.YListKeys = []string {}

    return &(virtualServiceCfgData.EntityData)
}

// VirtualServiceCfgData_Apps
// Application list
type VirtualServiceCfgData_Apps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application info. The type is slice of VirtualServiceCfgData_Apps_App.
    App []*VirtualServiceCfgData_Apps_App
}

func (apps *VirtualServiceCfgData_Apps) GetEntityData() *types.CommonEntityData {
    apps.EntityData.YFilter = apps.YFilter
    apps.EntityData.YangName = "apps"
    apps.EntityData.BundleName = "cisco_ios_xe"
    apps.EntityData.ParentYangName = "virtual-service-cfg-data"
    apps.EntityData.SegmentPath = "apps"
    apps.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/" + apps.EntityData.SegmentPath
    apps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    apps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    apps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    apps.EntityData.Children = types.NewOrderedMap()
    apps.EntityData.Children.Append("app", types.YChild{"App", nil})
    for i := range apps.App {
        apps.EntityData.Children.Append(types.GetSegmentPath(apps.App[i]), types.YChild{"App", apps.App[i]})
    }
    apps.EntityData.Leafs = types.NewOrderedMap()

    apps.EntityData.YListKeys = []string {}

    return &(apps.EntityData)
}

// VirtualServiceCfgData_Apps_App
// Application info
type VirtualServiceCfgData_Apps_App struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Application name. The type is string with length:
    // 1..128.
    ApplicationName interface{}

    // Application Network Resource Information.
    ApplicationNetworkResource VirtualServiceCfgData_Apps_App_ApplicationNetworkResource

    // Application Resource profile.
    ApplicationResourceProfile VirtualServiceCfgData_Apps_App_ApplicationResourceProfile

    // Application attached device.
    ApplicationAttachedDevice VirtualServiceCfgData_Apps_App_ApplicationAttachedDevice
}

func (app *VirtualServiceCfgData_Apps_App) GetEntityData() *types.CommonEntityData {
    app.EntityData.YFilter = app.YFilter
    app.EntityData.YangName = "app"
    app.EntityData.BundleName = "cisco_ios_xe"
    app.EntityData.ParentYangName = "apps"
    app.EntityData.SegmentPath = "app" + types.AddKeyToken(app.ApplicationName, "application-name")
    app.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/apps/" + app.EntityData.SegmentPath
    app.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    app.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    app.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    app.EntityData.Children = types.NewOrderedMap()
    app.EntityData.Children.Append("application-network-resource", types.YChild{"ApplicationNetworkResource", &app.ApplicationNetworkResource})
    app.EntityData.Children.Append("application-resource-profile", types.YChild{"ApplicationResourceProfile", &app.ApplicationResourceProfile})
    app.EntityData.Children.Append("application-attached-device", types.YChild{"ApplicationAttachedDevice", &app.ApplicationAttachedDevice})
    app.EntityData.Leafs = types.NewOrderedMap()
    app.EntityData.Leafs.Append("application-name", types.YLeaf{"ApplicationName", app.ApplicationName})

    app.EntityData.YListKeys = []string {"ApplicationName"}

    return &(app.EntityData)
}

// VirtualServiceCfgData_Apps_App_ApplicationNetworkResource
// Application Network Resource Information
type VirtualServiceCfgData_Apps_App_ApplicationNetworkResource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vnic gateway. The type is string with length: 1..2.
    VnicGateway0 interface{}

    // VirtualPortGroup guest interface name as number in range of 0 .. 3. The
    // type is string with length: 1..1.
    VirtualportgroupGuestInterfaceName1 interface{}

    // Guest interface IP address. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpAddress1 interface{}

    // Guest IP netmask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpNetmask1 interface{}

    // VirtualPortGroup application default gateway IP address. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupApplicationDefaultGateway1 interface{}

    // Name server IP address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Nameserver0 interface{}

    // VirtualPortGroup default gateway guest interface. The type is interface{}
    // with range: 0..255.
    VirtualportgroupGuestInterfaceDefaultGateway1 interface{}

    // Vnic gateway. The type is string with length: 1..2.
    VnicGateway1 interface{}

    // VirtualPortGroup guest interface name as number in range of 0 .. 3. The
    // type is string with length: 1..1.
    VirtualportgroupGuestInterfaceName2 interface{}

    // Guest interface IP address. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpAddress2 interface{}

    // Guest IP netmask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpNetmask2 interface{}

    // VirtualPortGroup application qateway IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupApplicationGateway2 interface{}

    // Name server IP address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Nameserver1 interface{}

    // VirtualPortGroup default gateway guest interface. The type is interface{}
    // with range: 0..255.
    VirtualportgroupGuestInterfaceDefaultGateway2 interface{}

    // Vnic gateway. The type is string with length: 1..2.
    VnicGateway2 interface{}

    // VirtualPortGroup guest interface name as number in range of 0 .. 3. The
    // type is string with length: 1..1.
    VirtualportgroupGuestInterfaceName3 interface{}

    // Guest interface IP address. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpAddress3 interface{}

    // Guest IP netmask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpNetmask3 interface{}

    // VirtualPortGroup application gateway IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupApplicationGateway3 interface{}

    // Name server IP address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Nameserver2 interface{}

    // VirtualPortGroup default gateway guest interface. The type is interface{}
    // with range: 0..255.
    VirtualportgroupGuestInterfaceDefaultGateway3 interface{}

    // Vnic gateway. The type is string with length: 1..2.
    VnicGateway3 interface{}

    // VirtualPortGroup guest interface name as number in range of 0 .. 3. The
    // type is string with length: 1..1.
    VirtualportgroupGuestInterfaceName4 interface{}

    // Guest interface IP address. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpAddress4 interface{}

    // Guest IP netmask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupGuestIpNetmask4 interface{}

    // VirtualPortGroup application gateway IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VirtualportgroupApplicationGateway4 interface{}

    // Name server IP address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Nameserver3 interface{}

    // VirtualPortGroup default gateway guest interface. The type is interface{}
    // with range: 0..255.
    VirtualportgroupGuestInterfaceDefaultGateway4 interface{}

    // Management port guest interface name as number in range of 0 .. 3. The type
    // is string with length: 1..1.
    ManagementInterfaceName interface{}

    // Guest management port interface IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ManagementGuestIpAddress interface{}

    // Guest management port interface IP netmask. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ManagementGuestIpNetmask interface{}

    // Management port application gateway IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ManagementApplicationGateway interface{}

    // Name server IP address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Nameseerver4 interface{}

    // Management port default gateway guest interface. The type is interface{}
    // with range: 0..255.
    ManagementGuestInterfaceDefaultGateway interface{}

    // Application MAC address.
    ApplicationMacAddress VirtualServiceCfgData_Apps_App_ApplicationNetworkResource_ApplicationMacAddress
}

func (applicationNetworkResource *VirtualServiceCfgData_Apps_App_ApplicationNetworkResource) GetEntityData() *types.CommonEntityData {
    applicationNetworkResource.EntityData.YFilter = applicationNetworkResource.YFilter
    applicationNetworkResource.EntityData.YangName = "application-network-resource"
    applicationNetworkResource.EntityData.BundleName = "cisco_ios_xe"
    applicationNetworkResource.EntityData.ParentYangName = "app"
    applicationNetworkResource.EntityData.SegmentPath = "application-network-resource"
    applicationNetworkResource.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/apps/app/" + applicationNetworkResource.EntityData.SegmentPath
    applicationNetworkResource.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    applicationNetworkResource.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    applicationNetworkResource.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    applicationNetworkResource.EntityData.Children = types.NewOrderedMap()
    applicationNetworkResource.EntityData.Children.Append("application-mac-address", types.YChild{"ApplicationMacAddress", &applicationNetworkResource.ApplicationMacAddress})
    applicationNetworkResource.EntityData.Leafs = types.NewOrderedMap()
    applicationNetworkResource.EntityData.Leafs.Append("vnic-gateway-0", types.YLeaf{"VnicGateway0", applicationNetworkResource.VnicGateway0})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-name-1", types.YLeaf{"VirtualportgroupGuestInterfaceName1", applicationNetworkResource.VirtualportgroupGuestInterfaceName1})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-address-1", types.YLeaf{"VirtualportgroupGuestIpAddress1", applicationNetworkResource.VirtualportgroupGuestIpAddress1})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-netmask-1", types.YLeaf{"VirtualportgroupGuestIpNetmask1", applicationNetworkResource.VirtualportgroupGuestIpNetmask1})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-application-default-gateway-1", types.YLeaf{"VirtualportgroupApplicationDefaultGateway1", applicationNetworkResource.VirtualportgroupApplicationDefaultGateway1})
    applicationNetworkResource.EntityData.Leafs.Append("nameserver-0", types.YLeaf{"Nameserver0", applicationNetworkResource.Nameserver0})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-default-gateway-1", types.YLeaf{"VirtualportgroupGuestInterfaceDefaultGateway1", applicationNetworkResource.VirtualportgroupGuestInterfaceDefaultGateway1})
    applicationNetworkResource.EntityData.Leafs.Append("vnic-gateway-1", types.YLeaf{"VnicGateway1", applicationNetworkResource.VnicGateway1})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-name-2", types.YLeaf{"VirtualportgroupGuestInterfaceName2", applicationNetworkResource.VirtualportgroupGuestInterfaceName2})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-address-2", types.YLeaf{"VirtualportgroupGuestIpAddress2", applicationNetworkResource.VirtualportgroupGuestIpAddress2})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-netmask-2", types.YLeaf{"VirtualportgroupGuestIpNetmask2", applicationNetworkResource.VirtualportgroupGuestIpNetmask2})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-application-gateway-2", types.YLeaf{"VirtualportgroupApplicationGateway2", applicationNetworkResource.VirtualportgroupApplicationGateway2})
    applicationNetworkResource.EntityData.Leafs.Append("nameserver-1", types.YLeaf{"Nameserver1", applicationNetworkResource.Nameserver1})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-default-gateway-2", types.YLeaf{"VirtualportgroupGuestInterfaceDefaultGateway2", applicationNetworkResource.VirtualportgroupGuestInterfaceDefaultGateway2})
    applicationNetworkResource.EntityData.Leafs.Append("vnic-gateway-2", types.YLeaf{"VnicGateway2", applicationNetworkResource.VnicGateway2})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-name-3", types.YLeaf{"VirtualportgroupGuestInterfaceName3", applicationNetworkResource.VirtualportgroupGuestInterfaceName3})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-address-3", types.YLeaf{"VirtualportgroupGuestIpAddress3", applicationNetworkResource.VirtualportgroupGuestIpAddress3})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-netmask-3", types.YLeaf{"VirtualportgroupGuestIpNetmask3", applicationNetworkResource.VirtualportgroupGuestIpNetmask3})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-application-gateway-3", types.YLeaf{"VirtualportgroupApplicationGateway3", applicationNetworkResource.VirtualportgroupApplicationGateway3})
    applicationNetworkResource.EntityData.Leafs.Append("nameserver2", types.YLeaf{"Nameserver2", applicationNetworkResource.Nameserver2})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-default-gateway-3", types.YLeaf{"VirtualportgroupGuestInterfaceDefaultGateway3", applicationNetworkResource.VirtualportgroupGuestInterfaceDefaultGateway3})
    applicationNetworkResource.EntityData.Leafs.Append("vnic-gateway-3", types.YLeaf{"VnicGateway3", applicationNetworkResource.VnicGateway3})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-name-4", types.YLeaf{"VirtualportgroupGuestInterfaceName4", applicationNetworkResource.VirtualportgroupGuestInterfaceName4})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-address-4", types.YLeaf{"VirtualportgroupGuestIpAddress4", applicationNetworkResource.VirtualportgroupGuestIpAddress4})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-ip-netmask-4", types.YLeaf{"VirtualportgroupGuestIpNetmask4", applicationNetworkResource.VirtualportgroupGuestIpNetmask4})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-application-gateway-4", types.YLeaf{"VirtualportgroupApplicationGateway4", applicationNetworkResource.VirtualportgroupApplicationGateway4})
    applicationNetworkResource.EntityData.Leafs.Append("nameserver-3", types.YLeaf{"Nameserver3", applicationNetworkResource.Nameserver3})
    applicationNetworkResource.EntityData.Leafs.Append("virtualportgroup-guest-interface-default-gateway-4", types.YLeaf{"VirtualportgroupGuestInterfaceDefaultGateway4", applicationNetworkResource.VirtualportgroupGuestInterfaceDefaultGateway4})
    applicationNetworkResource.EntityData.Leafs.Append("management-interface-name", types.YLeaf{"ManagementInterfaceName", applicationNetworkResource.ManagementInterfaceName})
    applicationNetworkResource.EntityData.Leafs.Append("management-guest-ip-address", types.YLeaf{"ManagementGuestIpAddress", applicationNetworkResource.ManagementGuestIpAddress})
    applicationNetworkResource.EntityData.Leafs.Append("management-guest-ip-netmask", types.YLeaf{"ManagementGuestIpNetmask", applicationNetworkResource.ManagementGuestIpNetmask})
    applicationNetworkResource.EntityData.Leafs.Append("management-application-gateway", types.YLeaf{"ManagementApplicationGateway", applicationNetworkResource.ManagementApplicationGateway})
    applicationNetworkResource.EntityData.Leafs.Append("nameseerver4", types.YLeaf{"Nameseerver4", applicationNetworkResource.Nameseerver4})
    applicationNetworkResource.EntityData.Leafs.Append("management-guest-interface-default-gateway", types.YLeaf{"ManagementGuestInterfaceDefaultGateway", applicationNetworkResource.ManagementGuestInterfaceDefaultGateway})

    applicationNetworkResource.EntityData.YListKeys = []string {}

    return &(applicationNetworkResource.EntityData)
}

// VirtualServiceCfgData_Apps_App_ApplicationNetworkResource_ApplicationMacAddress
// Application MAC address
type VirtualServiceCfgData_Apps_App_ApplicationNetworkResource_ApplicationMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // MAC interface name. The type is string with length: 0..32.
    MacInterfaceName interface{}
}

func (applicationMacAddress *VirtualServiceCfgData_Apps_App_ApplicationNetworkResource_ApplicationMacAddress) GetEntityData() *types.CommonEntityData {
    applicationMacAddress.EntityData.YFilter = applicationMacAddress.YFilter
    applicationMacAddress.EntityData.YangName = "application-mac-address"
    applicationMacAddress.EntityData.BundleName = "cisco_ios_xe"
    applicationMacAddress.EntityData.ParentYangName = "application-network-resource"
    applicationMacAddress.EntityData.SegmentPath = "application-mac-address"
    applicationMacAddress.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/apps/app/application-network-resource/" + applicationMacAddress.EntityData.SegmentPath
    applicationMacAddress.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    applicationMacAddress.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    applicationMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    applicationMacAddress.EntityData.Children = types.NewOrderedMap()
    applicationMacAddress.EntityData.Leafs = types.NewOrderedMap()
    applicationMacAddress.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", applicationMacAddress.MacAddress})
    applicationMacAddress.EntityData.Leafs.Append("mac-interface-name", types.YLeaf{"MacInterfaceName", applicationMacAddress.MacInterfaceName})

    applicationMacAddress.EntityData.YListKeys = []string {}

    return &(applicationMacAddress.EntityData)
}

// VirtualServiceCfgData_Apps_App_ApplicationResourceProfile
// Application Resource profile
type VirtualServiceCfgData_Apps_App_ApplicationResourceProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource profile name. The type is string with length: 0..64.
    ProfileName interface{}

    // Amount of VCPUs allocated for the application. The type is interface{} with
    // range: 0..65535.
    Vcpu interface{}

    // Amount of reserved cpu in unit. The type is interface{} with range:
    // 0..20000. The default value is 0.
    CpuUnits interface{}

    // Amount of reserved memory in MB. The type is interface{} with range:
    // 0..4096. The default value is 0.
    MemoryCapacityMb interface{}

    // Disk size in MB. The type is interface{} with range: 0..8192. The default
    // value is 0.
    DiskSizeMb interface{}

    // Resource package profile name. The type is string with length: 0..64.
    PkgProfileName interface{}
}

func (applicationResourceProfile *VirtualServiceCfgData_Apps_App_ApplicationResourceProfile) GetEntityData() *types.CommonEntityData {
    applicationResourceProfile.EntityData.YFilter = applicationResourceProfile.YFilter
    applicationResourceProfile.EntityData.YangName = "application-resource-profile"
    applicationResourceProfile.EntityData.BundleName = "cisco_ios_xe"
    applicationResourceProfile.EntityData.ParentYangName = "app"
    applicationResourceProfile.EntityData.SegmentPath = "application-resource-profile"
    applicationResourceProfile.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/apps/app/" + applicationResourceProfile.EntityData.SegmentPath
    applicationResourceProfile.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    applicationResourceProfile.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    applicationResourceProfile.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    applicationResourceProfile.EntityData.Children = types.NewOrderedMap()
    applicationResourceProfile.EntityData.Leafs = types.NewOrderedMap()
    applicationResourceProfile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", applicationResourceProfile.ProfileName})
    applicationResourceProfile.EntityData.Leafs.Append("vcpu", types.YLeaf{"Vcpu", applicationResourceProfile.Vcpu})
    applicationResourceProfile.EntityData.Leafs.Append("cpu-units", types.YLeaf{"CpuUnits", applicationResourceProfile.CpuUnits})
    applicationResourceProfile.EntityData.Leafs.Append("memory-capacity-mb", types.YLeaf{"MemoryCapacityMb", applicationResourceProfile.MemoryCapacityMb})
    applicationResourceProfile.EntityData.Leafs.Append("disk-size-mb", types.YLeaf{"DiskSizeMb", applicationResourceProfile.DiskSizeMb})
    applicationResourceProfile.EntityData.Leafs.Append("pkg-profile-name", types.YLeaf{"PkgProfileName", applicationResourceProfile.PkgProfileName})

    applicationResourceProfile.EntityData.YListKeys = []string {}

    return &(applicationResourceProfile.EntityData)
}

// VirtualServiceCfgData_Apps_App_ApplicationAttachedDevice
// Application attached device
type VirtualServiceCfgData_Apps_App_ApplicationAttachedDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // device name. The type is string with length: 0..32.
    DeviceName interface{}
}

func (applicationAttachedDevice *VirtualServiceCfgData_Apps_App_ApplicationAttachedDevice) GetEntityData() *types.CommonEntityData {
    applicationAttachedDevice.EntityData.YFilter = applicationAttachedDevice.YFilter
    applicationAttachedDevice.EntityData.YangName = "application-attached-device"
    applicationAttachedDevice.EntityData.BundleName = "cisco_ios_xe"
    applicationAttachedDevice.EntityData.ParentYangName = "app"
    applicationAttachedDevice.EntityData.SegmentPath = "application-attached-device"
    applicationAttachedDevice.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/apps/app/" + applicationAttachedDevice.EntityData.SegmentPath
    applicationAttachedDevice.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    applicationAttachedDevice.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    applicationAttachedDevice.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    applicationAttachedDevice.EntityData.Children = types.NewOrderedMap()
    applicationAttachedDevice.EntityData.Leafs = types.NewOrderedMap()
    applicationAttachedDevice.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", applicationAttachedDevice.DeviceName})

    applicationAttachedDevice.EntityData.YListKeys = []string {}

    return &(applicationAttachedDevice.EntityData)
}

// VirtualServiceCfgData_Controls
// App-hosting control
// This type is a presence type.
type VirtualServiceCfgData_Controls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Application hosting Infra enable status. The type is interface{} with
    // range: 0..255. The default value is 0.
    ApplicationHostingInfraEnableStatue interface{}
}

func (controls *VirtualServiceCfgData_Controls) GetEntityData() *types.CommonEntityData {
    controls.EntityData.YFilter = controls.YFilter
    controls.EntityData.YangName = "controls"
    controls.EntityData.BundleName = "cisco_ios_xe"
    controls.EntityData.ParentYangName = "virtual-service-cfg-data"
    controls.EntityData.SegmentPath = "controls"
    controls.EntityData.AbsolutePath = "Cisco-IOS-XE-virtual-service-cfg:virtual-service-cfg-data/" + controls.EntityData.SegmentPath
    controls.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    controls.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    controls.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    controls.EntityData.Children = types.NewOrderedMap()
    controls.EntityData.Leafs = types.NewOrderedMap()
    controls.EntityData.Leafs.Append("application-hosting-infra-enable-statue", types.YLeaf{"ApplicationHostingInfraEnableStatue", controls.ApplicationHostingInfraEnableStatue})

    controls.EntityData.YListKeys = []string {}

    return &(controls.EntityData)
}

