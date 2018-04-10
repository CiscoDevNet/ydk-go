// This module contains IOS-XR group YANG data 
// for flexible cli groups 
//     
// Copyright (c) 2013-2016 by Cisco Systems, Inc.
// All rights reserved.
package group_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package group_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-group-cfg groups}", reflect.TypeOf(Groups{}))
    ydk.RegisterEntity("Cisco-IOS-XR-group-cfg:groups", reflect.TypeOf(Groups{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-group-cfg apply-groups}", reflect.TypeOf(ApplyGroups{}))
    ydk.RegisterEntity("Cisco-IOS-XR-group-cfg:apply-groups", reflect.TypeOf(ApplyGroups{}))
}

// Groups
// config groups
type Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group config definition. The type is slice of Groups_Group.
    Group []Groups_Group
}

func (groups *Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "Cisco-IOS-XR-group-cfg"
    groups.EntityData.SegmentPath = "Cisco-IOS-XR-group-cfg:groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Groups_Group
// Group config definition
type Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with length: 0..32.
    GroupName interface{}
}

func (group *Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", group.GroupName}
    return &(group.EntityData)
}

// ApplyGroups
// apply groups
type ApplyGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // apply-group name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ApplyGroup interface{}
}

func (applyGroups *ApplyGroups) GetEntityData() *types.CommonEntityData {
    applyGroups.EntityData.YFilter = applyGroups.YFilter
    applyGroups.EntityData.YangName = "apply-groups"
    applyGroups.EntityData.BundleName = "cisco_ios_xr"
    applyGroups.EntityData.ParentYangName = "Cisco-IOS-XR-group-cfg"
    applyGroups.EntityData.SegmentPath = "Cisco-IOS-XR-group-cfg:apply-groups"
    applyGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applyGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applyGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applyGroups.EntityData.Children = make(map[string]types.YChild)
    applyGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    applyGroups.EntityData.Leafs["apply-group"] = types.YLeaf{"ApplyGroup", applyGroups.ApplyGroup}
    return &(applyGroups.EntityData)
}

