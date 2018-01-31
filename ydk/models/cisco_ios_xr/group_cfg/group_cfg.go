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
    parent types.Entity
    YFilter yfilter.YFilter

    // Group config definition. The type is slice of Groups_Group.
    Group []Groups_Group
}

func (groups *Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Groups) GetSegmentPath() string {
    return "Cisco-IOS-XR-group-cfg:groups"
}

func (groups *Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Groups) GetYangName() string { return "groups" }

func (groups *Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Groups) GetParent() types.Entity { return groups.parent }

func (groups *Groups) GetParentYangName() string { return "Cisco-IOS-XR-group-cfg" }

// Groups_Group
// Group config definition
type Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with length: 0..32.
    GroupName interface{}
}

func (group *Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Groups_Group) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    return ""
}

func (group *Groups_Group) GetSegmentPath() string {
    return "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
}

func (group *Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name"] = group.GroupName
    return leafs
}

func (group *Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Groups_Group) GetYangName() string { return "group" }

func (group *Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Groups_Group) GetParentYangName() string { return "groups" }

// ApplyGroups
// apply groups
type ApplyGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // apply-group name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ApplyGroup interface{}
}

func (applyGroups *ApplyGroups) GetFilter() yfilter.YFilter { return applyGroups.YFilter }

func (applyGroups *ApplyGroups) SetFilter(yf yfilter.YFilter) { applyGroups.YFilter = yf }

func (applyGroups *ApplyGroups) GetGoName(yname string) string {
    if yname == "apply-group" { return "ApplyGroup" }
    return ""
}

func (applyGroups *ApplyGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-group-cfg:apply-groups"
}

func (applyGroups *ApplyGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (applyGroups *ApplyGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (applyGroups *ApplyGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["apply-group"] = applyGroups.ApplyGroup
    return leafs
}

func (applyGroups *ApplyGroups) GetBundleName() string { return "cisco_ios_xr" }

func (applyGroups *ApplyGroups) GetYangName() string { return "apply-groups" }

func (applyGroups *ApplyGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (applyGroups *ApplyGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (applyGroups *ApplyGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (applyGroups *ApplyGroups) SetParent(parent types.Entity) { applyGroups.parent = parent }

func (applyGroups *ApplyGroups) GetParent() types.Entity { return applyGroups.parent }

func (applyGroups *ApplyGroups) GetParentYangName() string { return "Cisco-IOS-XR-group-cfg" }

