// This module contains a collection of YANG definitions
// for Cisco IOS-XR rgmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   redundancy-group-manager: Redundancy Group Manager
//     Configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package rgmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rgmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-rgmgr-cfg redundancy-group-manager}", reflect.TypeOf(RedundancyGroupManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager", reflect.TypeOf(RedundancyGroupManager{}))
}

// IccpMode represents Iccp mode
type IccpMode string

const (
    // Run the ICCP group in Singleton mode
    IccpMode_singleton IccpMode = "singleton"
)

// RedundancyGroupManager
// Redundancy Group Manager Configuration
type RedundancyGroupManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable redundancy group manager. The type is interface{}.
    Enable interface{}

    // MR-APS groups.
    Aps RedundancyGroupManager_Aps

    // ICCP configuration.
    Iccp RedundancyGroupManager_Iccp
}

func (redundancyGroupManager *RedundancyGroupManager) GetEntityData() *types.CommonEntityData {
    redundancyGroupManager.EntityData.YFilter = redundancyGroupManager.YFilter
    redundancyGroupManager.EntityData.YangName = "redundancy-group-manager"
    redundancyGroupManager.EntityData.BundleName = "cisco_ios_xr"
    redundancyGroupManager.EntityData.ParentYangName = "Cisco-IOS-XR-rgmgr-cfg"
    redundancyGroupManager.EntityData.SegmentPath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager"
    redundancyGroupManager.EntityData.AbsolutePath = redundancyGroupManager.EntityData.SegmentPath
    redundancyGroupManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancyGroupManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancyGroupManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancyGroupManager.EntityData.Children = types.NewOrderedMap()
    redundancyGroupManager.EntityData.Children.Append("aps", types.YChild{"Aps", &redundancyGroupManager.Aps})
    redundancyGroupManager.EntityData.Children.Append("iccp", types.YChild{"Iccp", &redundancyGroupManager.Iccp})
    redundancyGroupManager.EntityData.Leafs = types.NewOrderedMap()
    redundancyGroupManager.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", redundancyGroupManager.Enable})

    redundancyGroupManager.EntityData.YListKeys = []string {}

    return &(redundancyGroupManager.EntityData)
}

// RedundancyGroupManager_Aps
// MR-APS groups
type RedundancyGroupManager_Aps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default SONET controller backup configuration.
    DefaultRedundancyGroup RedundancyGroupManager_Aps_DefaultRedundancyGroup

    // Redundancy Group Table.
    Groups RedundancyGroupManager_Aps_Groups
}

func (aps *RedundancyGroupManager_Aps) GetEntityData() *types.CommonEntityData {
    aps.EntityData.YFilter = aps.YFilter
    aps.EntityData.YangName = "aps"
    aps.EntityData.BundleName = "cisco_ios_xr"
    aps.EntityData.ParentYangName = "redundancy-group-manager"
    aps.EntityData.SegmentPath = "aps"
    aps.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/" + aps.EntityData.SegmentPath
    aps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aps.EntityData.Children = types.NewOrderedMap()
    aps.EntityData.Children.Append("default-redundancy-group", types.YChild{"DefaultRedundancyGroup", &aps.DefaultRedundancyGroup})
    aps.EntityData.Children.Append("groups", types.YChild{"Groups", &aps.Groups})
    aps.EntityData.Leafs = types.NewOrderedMap()

    aps.EntityData.YListKeys = []string {}

    return &(aps.EntityData)
}

// RedundancyGroupManager_Aps_DefaultRedundancyGroup
// Default SONET controller backup configuration
type RedundancyGroupManager_Aps_DefaultRedundancyGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of remote peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Backup interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    BackupInterfaceName interface{}
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetEntityData() *types.CommonEntityData {
    defaultRedundancyGroup.EntityData.YFilter = defaultRedundancyGroup.YFilter
    defaultRedundancyGroup.EntityData.YangName = "default-redundancy-group"
    defaultRedundancyGroup.EntityData.BundleName = "cisco_ios_xr"
    defaultRedundancyGroup.EntityData.ParentYangName = "aps"
    defaultRedundancyGroup.EntityData.SegmentPath = "default-redundancy-group"
    defaultRedundancyGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/aps/" + defaultRedundancyGroup.EntityData.SegmentPath
    defaultRedundancyGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultRedundancyGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultRedundancyGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultRedundancyGroup.EntityData.Children = types.NewOrderedMap()
    defaultRedundancyGroup.EntityData.Leafs = types.NewOrderedMap()
    defaultRedundancyGroup.EntityData.Leafs.Append("next-hop-address", types.YLeaf{"NextHopAddress", defaultRedundancyGroup.NextHopAddress})
    defaultRedundancyGroup.EntityData.Leafs.Append("backup-interface-name", types.YLeaf{"BackupInterfaceName", defaultRedundancyGroup.BackupInterfaceName})

    defaultRedundancyGroup.EntityData.YListKeys = []string {}

    return &(defaultRedundancyGroup.EntityData)
}

// RedundancyGroupManager_Aps_Groups
// Redundancy Group Table
type RedundancyGroupManager_Aps_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group Configuration. The type is slice of
    // RedundancyGroupManager_Aps_Groups_Group.
    Group []*RedundancyGroupManager_Aps_Groups_Group
}

func (groups *RedundancyGroupManager_Aps_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "aps"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/aps/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// RedundancyGroupManager_Aps_Groups_Group
// Redundancy Group Configuration
type RedundancyGroupManager_Aps_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The redundancy group ID. The type is interface{}
    // with range: 1..32.
    GroupId interface{}

    // Controller configuration.
    Controllers RedundancyGroupManager_Aps_Groups_Group_Controllers
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id")
    group.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/aps/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("controllers", types.YChild{"Controllers", &group.Controllers})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})

    group.EntityData.YListKeys = []string {"GroupId"}

    return &(group.EntityData)
}

// RedundancyGroupManager_Aps_Groups_Group_Controllers
// Controller configuration
type RedundancyGroupManager_Aps_Groups_Group_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller.
    Controller []*RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "group"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/aps/groups/group/" + controllers.EntityData.SegmentPath
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller
// none
type RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Controller Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    ControllerName interface{}

    // IPv4 address of remote peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Backup interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    BackupInterfaceName interface{}
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/aps/groups/group/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("next-hop-address", types.YLeaf{"NextHopAddress", controller.NextHopAddress})
    controller.EntityData.Leafs.Append("backup-interface-name", types.YLeaf{"BackupInterfaceName", controller.BackupInterfaceName})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// RedundancyGroupManager_Iccp
// ICCP configuration
type RedundancyGroupManager_Iccp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group Table Configuration.
    IccpGroups RedundancyGroupManager_Iccp_IccpGroups
}

func (iccp *RedundancyGroupManager_Iccp) GetEntityData() *types.CommonEntityData {
    iccp.EntityData.YFilter = iccp.YFilter
    iccp.EntityData.YangName = "iccp"
    iccp.EntityData.BundleName = "cisco_ios_xr"
    iccp.EntityData.ParentYangName = "redundancy-group-manager"
    iccp.EntityData.SegmentPath = "iccp"
    iccp.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/" + iccp.EntityData.SegmentPath
    iccp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccp.EntityData.Children = types.NewOrderedMap()
    iccp.EntityData.Children.Append("iccp-groups", types.YChild{"IccpGroups", &iccp.IccpGroups})
    iccp.EntityData.Leafs = types.NewOrderedMap()

    iccp.EntityData.YListKeys = []string {}

    return &(iccp.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups
// Redundancy Group Table Configuration
type RedundancyGroupManager_Iccp_IccpGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group Configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup.
    IccpGroup []*RedundancyGroupManager_Iccp_IccpGroups_IccpGroup
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetEntityData() *types.CommonEntityData {
    iccpGroups.EntityData.YFilter = iccpGroups.YFilter
    iccpGroups.EntityData.YangName = "iccp-groups"
    iccpGroups.EntityData.BundleName = "cisco_ios_xr"
    iccpGroups.EntityData.ParentYangName = "iccp"
    iccpGroups.EntityData.SegmentPath = "iccp-groups"
    iccpGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/" + iccpGroups.EntityData.SegmentPath
    iccpGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpGroups.EntityData.Children = types.NewOrderedMap()
    iccpGroups.EntityData.Children.Append("iccp-group", types.YChild{"IccpGroup", nil})
    for i := range iccpGroups.IccpGroup {
        iccpGroups.EntityData.Children.Append(types.GetSegmentPath(iccpGroups.IccpGroup[i]), types.YChild{"IccpGroup", iccpGroups.IccpGroup[i]})
    }
    iccpGroups.EntityData.Leafs = types.NewOrderedMap()

    iccpGroups.EntityData.YListKeys = []string {}

    return &(iccpGroups.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup
// Redundancy Group Configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The redundancy icc group number. The type is
    // interface{} with range: 1..4294967295.
    GroupNumber interface{}

    // ICCP isolation recovery delay. The type is interface{} with range: 30..600.
    // Units are second.
    IsolationRecoveryDelay interface{}

    // ICCP mode. The type is IccpMode.
    Mode interface{}

    // ICCP backbone configuration.
    Backbones RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones

    // ICCP member configuration.
    Members RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members

    // nV Satellite configuration.
    NvSatellite RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite

    // Multi-chassis Link Aggregation Control Protocol commands.
    Mlacp RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetEntityData() *types.CommonEntityData {
    iccpGroup.EntityData.YFilter = iccpGroup.YFilter
    iccpGroup.EntityData.YangName = "iccp-group"
    iccpGroup.EntityData.BundleName = "cisco_ios_xr"
    iccpGroup.EntityData.ParentYangName = "iccp-groups"
    iccpGroup.EntityData.SegmentPath = "iccp-group" + types.AddKeyToken(iccpGroup.GroupNumber, "group-number")
    iccpGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/" + iccpGroup.EntityData.SegmentPath
    iccpGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpGroup.EntityData.Children = types.NewOrderedMap()
    iccpGroup.EntityData.Children.Append("backbones", types.YChild{"Backbones", &iccpGroup.Backbones})
    iccpGroup.EntityData.Children.Append("members", types.YChild{"Members", &iccpGroup.Members})
    iccpGroup.EntityData.Children.Append("Cisco-IOS-XR-icpe-infra-cfg:nv-satellite", types.YChild{"NvSatellite", &iccpGroup.NvSatellite})
    iccpGroup.EntityData.Children.Append("Cisco-IOS-XR-bundlemgr-cfg:mlacp", types.YChild{"Mlacp", &iccpGroup.Mlacp})
    iccpGroup.EntityData.Leafs = types.NewOrderedMap()
    iccpGroup.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", iccpGroup.GroupNumber})
    iccpGroup.EntityData.Leafs.Append("isolation-recovery-delay", types.YLeaf{"IsolationRecoveryDelay", iccpGroup.IsolationRecoveryDelay})
    iccpGroup.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", iccpGroup.Mode})

    iccpGroup.EntityData.YListKeys = []string {"GroupNumber"}

    return &(iccpGroup.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones
// ICCP backbone configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICCP backbone interface configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone.
    Backbone []*RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetEntityData() *types.CommonEntityData {
    backbones.EntityData.YFilter = backbones.YFilter
    backbones.EntityData.YangName = "backbones"
    backbones.EntityData.BundleName = "cisco_ios_xr"
    backbones.EntityData.ParentYangName = "iccp-group"
    backbones.EntityData.SegmentPath = "backbones"
    backbones.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/" + backbones.EntityData.SegmentPath
    backbones.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backbones.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backbones.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backbones.EntityData.Children = types.NewOrderedMap()
    backbones.EntityData.Children.Append("backbone", types.YChild{"Backbone", nil})
    for i := range backbones.Backbone {
        backbones.EntityData.Children.Append(types.GetSegmentPath(backbones.Backbone[i]), types.YChild{"Backbone", backbones.Backbone[i]})
    }
    backbones.EntityData.Leafs = types.NewOrderedMap()

    backbones.EntityData.YListKeys = []string {}

    return &(backbones.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone
// ICCP backbone interface configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    BackboneName interface{}
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetEntityData() *types.CommonEntityData {
    backbone.EntityData.YFilter = backbone.YFilter
    backbone.EntityData.YangName = "backbone"
    backbone.EntityData.BundleName = "cisco_ios_xr"
    backbone.EntityData.ParentYangName = "backbones"
    backbone.EntityData.SegmentPath = "backbone" + types.AddKeyToken(backbone.BackboneName, "backbone-name")
    backbone.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/backbones/" + backbone.EntityData.SegmentPath
    backbone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backbone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backbone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backbone.EntityData.Children = types.NewOrderedMap()
    backbone.EntityData.Leafs = types.NewOrderedMap()
    backbone.EntityData.Leafs.Append("backbone-name", types.YLeaf{"BackboneName", backbone.BackboneName})

    backbone.EntityData.YListKeys = []string {"BackboneName"}

    return &(backbone.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members
// ICCP member configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICCP member configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member.
    Member []*RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "cisco_ios_xr"
    members.EntityData.ParentYangName = "iccp-group"
    members.EntityData.SegmentPath = "members"
    members.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/" + members.EntityData.SegmentPath
    members.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    members.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    members.EntityData.Children = types.NewOrderedMap()
    members.EntityData.Children.Append("member", types.YChild{"Member", nil})
    for i := range members.Member {
        members.EntityData.Children.Append(types.GetSegmentPath(members.Member[i]), types.YChild{"Member", members.Member[i]})
    }
    members.EntityData.Leafs = types.NewOrderedMap()

    members.EntityData.YListKeys = []string {}

    return &(members.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member
// ICCP member configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "cisco_ios_xr"
    member.EntityData.ParentYangName = "members"
    member.EntityData.SegmentPath = "member" + types.AddKeyToken(member.NeighborAddress, "neighbor-address")
    member.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/members/" + member.EntityData.SegmentPath
    member.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    member.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    member.EntityData.Children = types.NewOrderedMap()
    member.EntityData.Leafs = types.NewOrderedMap()
    member.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", member.NeighborAddress})

    member.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(member.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite
// nV Satellite configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optional identifier for this system. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemMac interface{}
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "iccp-group"
    nvSatellite.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite"
    nvSatellite.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/" + nvSatellite.EntityData.SegmentPath
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs.Append("system-mac", types.YLeaf{"SystemMac", nvSatellite.SystemMac})

    nvSatellite.EntityData.YListKeys = []string {}

    return &(nvSatellite.EntityData)
}

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp
// Multi-chassis Link Aggregation Control Protocol
// commands
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds to wait before assuming mLACP peer is down. The type is
    // interface{} with range: 0..65534.
    ConnectTimeout interface{}

    // Unique LACP identifier for this system. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemMac interface{}

    // Unique identifier for this system in the ICCP Group. The type is
    // interface{} with range: 0..7.
    Node interface{}

    // Priority for this system. Lower value is higher priority. The type is
    // interface{} with range: 1..65535.
    SystemPriority interface{}
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetEntityData() *types.CommonEntityData {
    mlacp.EntityData.YFilter = mlacp.YFilter
    mlacp.EntityData.YangName = "mlacp"
    mlacp.EntityData.BundleName = "cisco_ios_xr"
    mlacp.EntityData.ParentYangName = "iccp-group"
    mlacp.EntityData.SegmentPath = "Cisco-IOS-XR-bundlemgr-cfg:mlacp"
    mlacp.EntityData.AbsolutePath = "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager/iccp/iccp-groups/iccp-group/" + mlacp.EntityData.SegmentPath
    mlacp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mlacp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mlacp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mlacp.EntityData.Children = types.NewOrderedMap()
    mlacp.EntityData.Leafs = types.NewOrderedMap()
    mlacp.EntityData.Leafs.Append("connect-timeout", types.YLeaf{"ConnectTimeout", mlacp.ConnectTimeout})
    mlacp.EntityData.Leafs.Append("system-mac", types.YLeaf{"SystemMac", mlacp.SystemMac})
    mlacp.EntityData.Leafs.Append("node", types.YLeaf{"Node", mlacp.Node})
    mlacp.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", mlacp.SystemPriority})

    mlacp.EntityData.YListKeys = []string {}

    return &(mlacp.EntityData)
}

