// This module contains a collection of YANG definitions
// for Cisco IOS-XR ha-eem-policy package operational data.
// 
// This module contains definitions
// for the following management objects:
//   eem: EEM operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ha_eem_policy_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ha_eem_policy_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ha-eem-policy-oper eem}", reflect.TypeOf(Eem{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ha-eem-policy-oper:eem", reflect.TypeOf(Eem{}))
}

// Eem
// EEM operational data
type Eem struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // directory user.
    DirUser Eem_DirUser

    // list of environmental variables.
    EnvVariables Eem_EnvVariables

    // Refresh time.
    RefreshTime Eem_RefreshTime

    // list the registered policies.
    RegPolicies Eem_RegPolicies

    // list the available policies.
    AvlPolicies Eem_AvlPolicies
}

func (eem *Eem) GetFilter() yfilter.YFilter { return eem.YFilter }

func (eem *Eem) SetFilter(yf yfilter.YFilter) { eem.YFilter = yf }

func (eem *Eem) GetGoName(yname string) string {
    if yname == "dir-user" { return "DirUser" }
    if yname == "env-variables" { return "EnvVariables" }
    if yname == "refresh-time" { return "RefreshTime" }
    if yname == "reg-policies" { return "RegPolicies" }
    if yname == "avl-policies" { return "AvlPolicies" }
    return ""
}

func (eem *Eem) GetSegmentPath() string {
    return "Cisco-IOS-XR-ha-eem-policy-oper:eem"
}

func (eem *Eem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dir-user" {
        return &eem.DirUser
    }
    if childYangName == "env-variables" {
        return &eem.EnvVariables
    }
    if childYangName == "refresh-time" {
        return &eem.RefreshTime
    }
    if childYangName == "reg-policies" {
        return &eem.RegPolicies
    }
    if childYangName == "avl-policies" {
        return &eem.AvlPolicies
    }
    return nil
}

func (eem *Eem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dir-user"] = &eem.DirUser
    children["env-variables"] = &eem.EnvVariables
    children["refresh-time"] = &eem.RefreshTime
    children["reg-policies"] = &eem.RegPolicies
    children["avl-policies"] = &eem.AvlPolicies
    return children
}

func (eem *Eem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eem *Eem) GetBundleName() string { return "cisco_ios_xr" }

func (eem *Eem) GetYangName() string { return "eem" }

func (eem *Eem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eem *Eem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eem *Eem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eem *Eem) SetParent(parent types.Entity) { eem.parent = parent }

func (eem *Eem) GetParent() types.Entity { return eem.parent }

func (eem *Eem) GetParentYangName() string { return "Cisco-IOS-XR-ha-eem-policy-oper" }

// Eem_DirUser
// directory user
type Eem_DirUser struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // directory user library.
    Library Eem_DirUser_Library

    // directory user policy.
    Policy Eem_DirUser_Policy
}

func (dirUser *Eem_DirUser) GetFilter() yfilter.YFilter { return dirUser.YFilter }

func (dirUser *Eem_DirUser) SetFilter(yf yfilter.YFilter) { dirUser.YFilter = yf }

func (dirUser *Eem_DirUser) GetGoName(yname string) string {
    if yname == "library" { return "Library" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (dirUser *Eem_DirUser) GetSegmentPath() string {
    return "dir-user"
}

func (dirUser *Eem_DirUser) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "library" {
        return &dirUser.Library
    }
    if childYangName == "policy" {
        return &dirUser.Policy
    }
    return nil
}

func (dirUser *Eem_DirUser) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["library"] = &dirUser.Library
    children["policy"] = &dirUser.Policy
    return children
}

func (dirUser *Eem_DirUser) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dirUser *Eem_DirUser) GetBundleName() string { return "cisco_ios_xr" }

func (dirUser *Eem_DirUser) GetYangName() string { return "dir-user" }

func (dirUser *Eem_DirUser) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dirUser *Eem_DirUser) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dirUser *Eem_DirUser) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dirUser *Eem_DirUser) SetParent(parent types.Entity) { dirUser.parent = parent }

func (dirUser *Eem_DirUser) GetParent() types.Entity { return dirUser.parent }

func (dirUser *Eem_DirUser) GetParentYangName() string { return "eem" }

// Eem_DirUser_Library
// directory user library
type Eem_DirUser_Library struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // policy. The type is string.
    Policy interface{}

    // library. The type is string.
    Library interface{}
}

func (library *Eem_DirUser_Library) GetFilter() yfilter.YFilter { return library.YFilter }

func (library *Eem_DirUser_Library) SetFilter(yf yfilter.YFilter) { library.YFilter = yf }

func (library *Eem_DirUser_Library) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    if yname == "library" { return "Library" }
    return ""
}

func (library *Eem_DirUser_Library) GetSegmentPath() string {
    return "library"
}

func (library *Eem_DirUser_Library) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (library *Eem_DirUser_Library) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (library *Eem_DirUser_Library) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = library.Policy
    leafs["library"] = library.Library
    return leafs
}

func (library *Eem_DirUser_Library) GetBundleName() string { return "cisco_ios_xr" }

func (library *Eem_DirUser_Library) GetYangName() string { return "library" }

func (library *Eem_DirUser_Library) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (library *Eem_DirUser_Library) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (library *Eem_DirUser_Library) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (library *Eem_DirUser_Library) SetParent(parent types.Entity) { library.parent = parent }

func (library *Eem_DirUser_Library) GetParent() types.Entity { return library.parent }

func (library *Eem_DirUser_Library) GetParentYangName() string { return "dir-user" }

// Eem_DirUser_Policy
// directory user policy
type Eem_DirUser_Policy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // policy. The type is string.
    Policy interface{}

    // library. The type is string.
    Library interface{}
}

func (policy *Eem_DirUser_Policy) GetFilter() yfilter.YFilter { return policy.YFilter }

func (policy *Eem_DirUser_Policy) SetFilter(yf yfilter.YFilter) { policy.YFilter = yf }

func (policy *Eem_DirUser_Policy) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    if yname == "library" { return "Library" }
    return ""
}

func (policy *Eem_DirUser_Policy) GetSegmentPath() string {
    return "policy"
}

func (policy *Eem_DirUser_Policy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policy *Eem_DirUser_Policy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policy *Eem_DirUser_Policy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = policy.Policy
    leafs["library"] = policy.Library
    return leafs
}

func (policy *Eem_DirUser_Policy) GetBundleName() string { return "cisco_ios_xr" }

func (policy *Eem_DirUser_Policy) GetYangName() string { return "policy" }

func (policy *Eem_DirUser_Policy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policy *Eem_DirUser_Policy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policy *Eem_DirUser_Policy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policy *Eem_DirUser_Policy) SetParent(parent types.Entity) { policy.parent = parent }

func (policy *Eem_DirUser_Policy) GetParent() types.Entity { return policy.parent }

func (policy *Eem_DirUser_Policy) GetParentYangName() string { return "dir-user" }

// Eem_EnvVariables
// list of environmental variables
type Eem_EnvVariables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // environmental variables name and value . The type is slice of
    // Eem_EnvVariables_EnvVariable.
    EnvVariable []Eem_EnvVariables_EnvVariable
}

func (envVariables *Eem_EnvVariables) GetFilter() yfilter.YFilter { return envVariables.YFilter }

func (envVariables *Eem_EnvVariables) SetFilter(yf yfilter.YFilter) { envVariables.YFilter = yf }

func (envVariables *Eem_EnvVariables) GetGoName(yname string) string {
    if yname == "env-variable" { return "EnvVariable" }
    return ""
}

func (envVariables *Eem_EnvVariables) GetSegmentPath() string {
    return "env-variables"
}

func (envVariables *Eem_EnvVariables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "env-variable" {
        for _, c := range envVariables.EnvVariable {
            if envVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Eem_EnvVariables_EnvVariable{}
        envVariables.EnvVariable = append(envVariables.EnvVariable, child)
        return &envVariables.EnvVariable[len(envVariables.EnvVariable)-1]
    }
    return nil
}

func (envVariables *Eem_EnvVariables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range envVariables.EnvVariable {
        children[envVariables.EnvVariable[i].GetSegmentPath()] = &envVariables.EnvVariable[i]
    }
    return children
}

func (envVariables *Eem_EnvVariables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (envVariables *Eem_EnvVariables) GetBundleName() string { return "cisco_ios_xr" }

func (envVariables *Eem_EnvVariables) GetYangName() string { return "env-variables" }

func (envVariables *Eem_EnvVariables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (envVariables *Eem_EnvVariables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (envVariables *Eem_EnvVariables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (envVariables *Eem_EnvVariables) SetParent(parent types.Entity) { envVariables.parent = parent }

func (envVariables *Eem_EnvVariables) GetParent() types.Entity { return envVariables.parent }

func (envVariables *Eem_EnvVariables) GetParentYangName() string { return "eem" }

// Eem_EnvVariables_EnvVariable
// environmental variables name and value 
type Eem_EnvVariables_EnvVariable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Environmental variable name. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // variable name. The type is string.
    NameXr interface{}

    // value. The type is string.
    Value interface{}
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetFilter() yfilter.YFilter { return envVariable.YFilter }

func (envVariable *Eem_EnvVariables_EnvVariable) SetFilter(yf yfilter.YFilter) { envVariable.YFilter = yf }

func (envVariable *Eem_EnvVariables_EnvVariable) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "name-xr" { return "NameXr" }
    if yname == "value" { return "Value" }
    return ""
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetSegmentPath() string {
    return "env-variable" + "[name='" + fmt.Sprintf("%v", envVariable.Name) + "']"
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = envVariable.Name
    leafs["name-xr"] = envVariable.NameXr
    leafs["value"] = envVariable.Value
    return leafs
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetBundleName() string { return "cisco_ios_xr" }

func (envVariable *Eem_EnvVariables_EnvVariable) GetYangName() string { return "env-variable" }

func (envVariable *Eem_EnvVariables_EnvVariable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (envVariable *Eem_EnvVariables_EnvVariable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (envVariable *Eem_EnvVariables_EnvVariable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (envVariable *Eem_EnvVariables_EnvVariable) SetParent(parent types.Entity) { envVariable.parent = parent }

func (envVariable *Eem_EnvVariables_EnvVariable) GetParent() types.Entity { return envVariable.parent }

func (envVariable *Eem_EnvVariables_EnvVariable) GetParentYangName() string { return "env-variables" }

// Eem_RefreshTime
// Refresh time
type Eem_RefreshTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Event manager refresh-time . The type is interface{} with range:
    // 0..4294967295.
    Refreshtime interface{}
}

func (refreshTime *Eem_RefreshTime) GetFilter() yfilter.YFilter { return refreshTime.YFilter }

func (refreshTime *Eem_RefreshTime) SetFilter(yf yfilter.YFilter) { refreshTime.YFilter = yf }

func (refreshTime *Eem_RefreshTime) GetGoName(yname string) string {
    if yname == "refreshtime" { return "Refreshtime" }
    return ""
}

func (refreshTime *Eem_RefreshTime) GetSegmentPath() string {
    return "refresh-time"
}

func (refreshTime *Eem_RefreshTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (refreshTime *Eem_RefreshTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (refreshTime *Eem_RefreshTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["refreshtime"] = refreshTime.Refreshtime
    return leafs
}

func (refreshTime *Eem_RefreshTime) GetBundleName() string { return "cisco_ios_xr" }

func (refreshTime *Eem_RefreshTime) GetYangName() string { return "refresh-time" }

func (refreshTime *Eem_RefreshTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (refreshTime *Eem_RefreshTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (refreshTime *Eem_RefreshTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (refreshTime *Eem_RefreshTime) SetParent(parent types.Entity) { refreshTime.parent = parent }

func (refreshTime *Eem_RefreshTime) GetParent() types.Entity { return refreshTime.parent }

func (refreshTime *Eem_RefreshTime) GetParentYangName() string { return "eem" }

// Eem_RegPolicies
// list the registered policies
type Eem_RegPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // policy name and create time . The type is slice of
    // Eem_RegPolicies_RegPolicy.
    RegPolicy []Eem_RegPolicies_RegPolicy
}

func (regPolicies *Eem_RegPolicies) GetFilter() yfilter.YFilter { return regPolicies.YFilter }

func (regPolicies *Eem_RegPolicies) SetFilter(yf yfilter.YFilter) { regPolicies.YFilter = yf }

func (regPolicies *Eem_RegPolicies) GetGoName(yname string) string {
    if yname == "reg-policy" { return "RegPolicy" }
    return ""
}

func (regPolicies *Eem_RegPolicies) GetSegmentPath() string {
    return "reg-policies"
}

func (regPolicies *Eem_RegPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reg-policy" {
        for _, c := range regPolicies.RegPolicy {
            if regPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Eem_RegPolicies_RegPolicy{}
        regPolicies.RegPolicy = append(regPolicies.RegPolicy, child)
        return &regPolicies.RegPolicy[len(regPolicies.RegPolicy)-1]
    }
    return nil
}

func (regPolicies *Eem_RegPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range regPolicies.RegPolicy {
        children[regPolicies.RegPolicy[i].GetSegmentPath()] = &regPolicies.RegPolicy[i]
    }
    return children
}

func (regPolicies *Eem_RegPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regPolicies *Eem_RegPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (regPolicies *Eem_RegPolicies) GetYangName() string { return "reg-policies" }

func (regPolicies *Eem_RegPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regPolicies *Eem_RegPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regPolicies *Eem_RegPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regPolicies *Eem_RegPolicies) SetParent(parent types.Entity) { regPolicies.parent = parent }

func (regPolicies *Eem_RegPolicies) GetParent() types.Entity { return regPolicies.parent }

func (regPolicies *Eem_RegPolicies) GetParentYangName() string { return "eem" }

// Eem_RegPolicies_RegPolicy
// policy name and create time 
type Eem_RegPolicies_RegPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. policy name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // policy type. The type is string.
    Type interface{}

    // time created. The type is string.
    TimeCreated interface{}

    // policy name. The type is string.
    PolicyName interface{}

    // class. The type is string.
    Class interface{}

    // event type. The type is string.
    EventType interface{}

    // trap. The type is string.
    Trap interface{}

    // PersistTime . The type is interface{} with range: 0..4294967295.
    PersistTime interface{}

    // username. The type is string.
    Username interface{}

    // description. The type is string.
    Description interface{}
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetFilter() yfilter.YFilter { return regPolicy.YFilter }

func (regPolicy *Eem_RegPolicies_RegPolicy) SetFilter(yf yfilter.YFilter) { regPolicy.YFilter = yf }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "time-created" { return "TimeCreated" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "class" { return "Class" }
    if yname == "event-type" { return "EventType" }
    if yname == "trap" { return "Trap" }
    if yname == "persist-time" { return "PersistTime" }
    if yname == "username" { return "Username" }
    if yname == "description" { return "Description" }
    return ""
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetSegmentPath() string {
    return "reg-policy" + "[name='" + fmt.Sprintf("%v", regPolicy.Name) + "']"
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = regPolicy.Name
    leafs["type"] = regPolicy.Type
    leafs["time-created"] = regPolicy.TimeCreated
    leafs["policy-name"] = regPolicy.PolicyName
    leafs["class"] = regPolicy.Class
    leafs["event-type"] = regPolicy.EventType
    leafs["trap"] = regPolicy.Trap
    leafs["persist-time"] = regPolicy.PersistTime
    leafs["username"] = regPolicy.Username
    leafs["description"] = regPolicy.Description
    return leafs
}

func (regPolicy *Eem_RegPolicies_RegPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetYangName() string { return "reg-policy" }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regPolicy *Eem_RegPolicies_RegPolicy) SetParent(parent types.Entity) { regPolicy.parent = parent }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetParent() types.Entity { return regPolicy.parent }

func (regPolicy *Eem_RegPolicies_RegPolicy) GetParentYangName() string { return "reg-policies" }

// Eem_AvlPolicies
// list the available policies
type Eem_AvlPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // policy name and create time . The type is slice of
    // Eem_AvlPolicies_AvlPolicy.
    AvlPolicy []Eem_AvlPolicies_AvlPolicy
}

func (avlPolicies *Eem_AvlPolicies) GetFilter() yfilter.YFilter { return avlPolicies.YFilter }

func (avlPolicies *Eem_AvlPolicies) SetFilter(yf yfilter.YFilter) { avlPolicies.YFilter = yf }

func (avlPolicies *Eem_AvlPolicies) GetGoName(yname string) string {
    if yname == "avl-policy" { return "AvlPolicy" }
    return ""
}

func (avlPolicies *Eem_AvlPolicies) GetSegmentPath() string {
    return "avl-policies"
}

func (avlPolicies *Eem_AvlPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "avl-policy" {
        for _, c := range avlPolicies.AvlPolicy {
            if avlPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Eem_AvlPolicies_AvlPolicy{}
        avlPolicies.AvlPolicy = append(avlPolicies.AvlPolicy, child)
        return &avlPolicies.AvlPolicy[len(avlPolicies.AvlPolicy)-1]
    }
    return nil
}

func (avlPolicies *Eem_AvlPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range avlPolicies.AvlPolicy {
        children[avlPolicies.AvlPolicy[i].GetSegmentPath()] = &avlPolicies.AvlPolicy[i]
    }
    return children
}

func (avlPolicies *Eem_AvlPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (avlPolicies *Eem_AvlPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (avlPolicies *Eem_AvlPolicies) GetYangName() string { return "avl-policies" }

func (avlPolicies *Eem_AvlPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (avlPolicies *Eem_AvlPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (avlPolicies *Eem_AvlPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (avlPolicies *Eem_AvlPolicies) SetParent(parent types.Entity) { avlPolicies.parent = parent }

func (avlPolicies *Eem_AvlPolicies) GetParent() types.Entity { return avlPolicies.parent }

func (avlPolicies *Eem_AvlPolicies) GetParentYangName() string { return "eem" }

// Eem_AvlPolicies_AvlPolicy
// policy name and create time 
type Eem_AvlPolicies_AvlPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. System policy name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // policy type. The type is string.
    Type interface{}

    // time created. The type is string.
    TimeCreated interface{}

    // policy name. The type is string.
    PolicyName interface{}
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetFilter() yfilter.YFilter { return avlPolicy.YFilter }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) SetFilter(yf yfilter.YFilter) { avlPolicy.YFilter = yf }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "time-created" { return "TimeCreated" }
    if yname == "policy-name" { return "PolicyName" }
    return ""
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetSegmentPath() string {
    return "avl-policy" + "[name='" + fmt.Sprintf("%v", avlPolicy.Name) + "']"
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = avlPolicy.Name
    leafs["type"] = avlPolicy.Type
    leafs["time-created"] = avlPolicy.TimeCreated
    leafs["policy-name"] = avlPolicy.PolicyName
    return leafs
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetYangName() string { return "avl-policy" }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) SetParent(parent types.Entity) { avlPolicy.parent = parent }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetParent() types.Entity { return avlPolicy.parent }

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetParentYangName() string { return "avl-policies" }

