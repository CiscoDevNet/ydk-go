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
    EntityData types.CommonEntityData
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

func (eem *Eem) GetEntityData() *types.CommonEntityData {
    eem.EntityData.YFilter = eem.YFilter
    eem.EntityData.YangName = "eem"
    eem.EntityData.BundleName = "cisco_ios_xr"
    eem.EntityData.ParentYangName = "Cisco-IOS-XR-ha-eem-policy-oper"
    eem.EntityData.SegmentPath = "Cisco-IOS-XR-ha-eem-policy-oper:eem"
    eem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eem.EntityData.Children = make(map[string]types.YChild)
    eem.EntityData.Children["dir-user"] = types.YChild{"DirUser", &eem.DirUser}
    eem.EntityData.Children["env-variables"] = types.YChild{"EnvVariables", &eem.EnvVariables}
    eem.EntityData.Children["refresh-time"] = types.YChild{"RefreshTime", &eem.RefreshTime}
    eem.EntityData.Children["reg-policies"] = types.YChild{"RegPolicies", &eem.RegPolicies}
    eem.EntityData.Children["avl-policies"] = types.YChild{"AvlPolicies", &eem.AvlPolicies}
    eem.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eem.EntityData)
}

// Eem_DirUser
// directory user
type Eem_DirUser struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // directory user library.
    Library Eem_DirUser_Library

    // directory user policy.
    Policy Eem_DirUser_Policy
}

func (dirUser *Eem_DirUser) GetEntityData() *types.CommonEntityData {
    dirUser.EntityData.YFilter = dirUser.YFilter
    dirUser.EntityData.YangName = "dir-user"
    dirUser.EntityData.BundleName = "cisco_ios_xr"
    dirUser.EntityData.ParentYangName = "eem"
    dirUser.EntityData.SegmentPath = "dir-user"
    dirUser.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dirUser.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dirUser.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dirUser.EntityData.Children = make(map[string]types.YChild)
    dirUser.EntityData.Children["library"] = types.YChild{"Library", &dirUser.Library}
    dirUser.EntityData.Children["policy"] = types.YChild{"Policy", &dirUser.Policy}
    dirUser.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dirUser.EntityData)
}

// Eem_DirUser_Library
// directory user library
type Eem_DirUser_Library struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policy. The type is string.
    Policy interface{}

    // library. The type is string.
    Library interface{}
}

func (library *Eem_DirUser_Library) GetEntityData() *types.CommonEntityData {
    library.EntityData.YFilter = library.YFilter
    library.EntityData.YangName = "library"
    library.EntityData.BundleName = "cisco_ios_xr"
    library.EntityData.ParentYangName = "dir-user"
    library.EntityData.SegmentPath = "library"
    library.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    library.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    library.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    library.EntityData.Children = make(map[string]types.YChild)
    library.EntityData.Leafs = make(map[string]types.YLeaf)
    library.EntityData.Leafs["policy"] = types.YLeaf{"Policy", library.Policy}
    library.EntityData.Leafs["library"] = types.YLeaf{"Library", library.Library}
    return &(library.EntityData)
}

// Eem_DirUser_Policy
// directory user policy
type Eem_DirUser_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policy. The type is string.
    Policy interface{}

    // library. The type is string.
    Library interface{}
}

func (policy *Eem_DirUser_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "dir-user"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    policy.EntityData.Leafs["policy"] = types.YLeaf{"Policy", policy.Policy}
    policy.EntityData.Leafs["library"] = types.YLeaf{"Library", policy.Library}
    return &(policy.EntityData)
}

// Eem_EnvVariables
// list of environmental variables
type Eem_EnvVariables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // environmental variables name and value . The type is slice of
    // Eem_EnvVariables_EnvVariable.
    EnvVariable []Eem_EnvVariables_EnvVariable
}

func (envVariables *Eem_EnvVariables) GetEntityData() *types.CommonEntityData {
    envVariables.EntityData.YFilter = envVariables.YFilter
    envVariables.EntityData.YangName = "env-variables"
    envVariables.EntityData.BundleName = "cisco_ios_xr"
    envVariables.EntityData.ParentYangName = "eem"
    envVariables.EntityData.SegmentPath = "env-variables"
    envVariables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    envVariables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    envVariables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    envVariables.EntityData.Children = make(map[string]types.YChild)
    envVariables.EntityData.Children["env-variable"] = types.YChild{"EnvVariable", nil}
    for i := range envVariables.EnvVariable {
        envVariables.EntityData.Children[types.GetSegmentPath(&envVariables.EnvVariable[i])] = types.YChild{"EnvVariable", &envVariables.EnvVariable[i]}
    }
    envVariables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(envVariables.EntityData)
}

// Eem_EnvVariables_EnvVariable
// environmental variables name and value 
type Eem_EnvVariables_EnvVariable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Environmental variable name. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // variable name. The type is string.
    NameXr interface{}

    // value. The type is string.
    Value interface{}
}

func (envVariable *Eem_EnvVariables_EnvVariable) GetEntityData() *types.CommonEntityData {
    envVariable.EntityData.YFilter = envVariable.YFilter
    envVariable.EntityData.YangName = "env-variable"
    envVariable.EntityData.BundleName = "cisco_ios_xr"
    envVariable.EntityData.ParentYangName = "env-variables"
    envVariable.EntityData.SegmentPath = "env-variable" + "[name='" + fmt.Sprintf("%v", envVariable.Name) + "']"
    envVariable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    envVariable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    envVariable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    envVariable.EntityData.Children = make(map[string]types.YChild)
    envVariable.EntityData.Leafs = make(map[string]types.YLeaf)
    envVariable.EntityData.Leafs["name"] = types.YLeaf{"Name", envVariable.Name}
    envVariable.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", envVariable.NameXr}
    envVariable.EntityData.Leafs["value"] = types.YLeaf{"Value", envVariable.Value}
    return &(envVariable.EntityData)
}

// Eem_RefreshTime
// Refresh time
type Eem_RefreshTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event manager refresh-time . The type is interface{} with range:
    // 0..4294967295.
    Refreshtime interface{}
}

func (refreshTime *Eem_RefreshTime) GetEntityData() *types.CommonEntityData {
    refreshTime.EntityData.YFilter = refreshTime.YFilter
    refreshTime.EntityData.YangName = "refresh-time"
    refreshTime.EntityData.BundleName = "cisco_ios_xr"
    refreshTime.EntityData.ParentYangName = "eem"
    refreshTime.EntityData.SegmentPath = "refresh-time"
    refreshTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    refreshTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    refreshTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    refreshTime.EntityData.Children = make(map[string]types.YChild)
    refreshTime.EntityData.Leafs = make(map[string]types.YLeaf)
    refreshTime.EntityData.Leafs["refreshtime"] = types.YLeaf{"Refreshtime", refreshTime.Refreshtime}
    return &(refreshTime.EntityData)
}

// Eem_RegPolicies
// list the registered policies
type Eem_RegPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policy name and create time . The type is slice of
    // Eem_RegPolicies_RegPolicy.
    RegPolicy []Eem_RegPolicies_RegPolicy
}

func (regPolicies *Eem_RegPolicies) GetEntityData() *types.CommonEntityData {
    regPolicies.EntityData.YFilter = regPolicies.YFilter
    regPolicies.EntityData.YangName = "reg-policies"
    regPolicies.EntityData.BundleName = "cisco_ios_xr"
    regPolicies.EntityData.ParentYangName = "eem"
    regPolicies.EntityData.SegmentPath = "reg-policies"
    regPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regPolicies.EntityData.Children = make(map[string]types.YChild)
    regPolicies.EntityData.Children["reg-policy"] = types.YChild{"RegPolicy", nil}
    for i := range regPolicies.RegPolicy {
        regPolicies.EntityData.Children[types.GetSegmentPath(&regPolicies.RegPolicy[i])] = types.YChild{"RegPolicy", &regPolicies.RegPolicy[i]}
    }
    regPolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(regPolicies.EntityData)
}

// Eem_RegPolicies_RegPolicy
// policy name and create time 
type Eem_RegPolicies_RegPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. policy name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // policy type. The type is string.
    Type_ interface{}

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

func (regPolicy *Eem_RegPolicies_RegPolicy) GetEntityData() *types.CommonEntityData {
    regPolicy.EntityData.YFilter = regPolicy.YFilter
    regPolicy.EntityData.YangName = "reg-policy"
    regPolicy.EntityData.BundleName = "cisco_ios_xr"
    regPolicy.EntityData.ParentYangName = "reg-policies"
    regPolicy.EntityData.SegmentPath = "reg-policy" + "[name='" + fmt.Sprintf("%v", regPolicy.Name) + "']"
    regPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regPolicy.EntityData.Children = make(map[string]types.YChild)
    regPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    regPolicy.EntityData.Leafs["name"] = types.YLeaf{"Name", regPolicy.Name}
    regPolicy.EntityData.Leafs["type"] = types.YLeaf{"Type_", regPolicy.Type_}
    regPolicy.EntityData.Leafs["time-created"] = types.YLeaf{"TimeCreated", regPolicy.TimeCreated}
    regPolicy.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", regPolicy.PolicyName}
    regPolicy.EntityData.Leafs["class"] = types.YLeaf{"Class", regPolicy.Class}
    regPolicy.EntityData.Leafs["event-type"] = types.YLeaf{"EventType", regPolicy.EventType}
    regPolicy.EntityData.Leafs["trap"] = types.YLeaf{"Trap", regPolicy.Trap}
    regPolicy.EntityData.Leafs["persist-time"] = types.YLeaf{"PersistTime", regPolicy.PersistTime}
    regPolicy.EntityData.Leafs["username"] = types.YLeaf{"Username", regPolicy.Username}
    regPolicy.EntityData.Leafs["description"] = types.YLeaf{"Description", regPolicy.Description}
    return &(regPolicy.EntityData)
}

// Eem_AvlPolicies
// list the available policies
type Eem_AvlPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policy name and create time . The type is slice of
    // Eem_AvlPolicies_AvlPolicy.
    AvlPolicy []Eem_AvlPolicies_AvlPolicy
}

func (avlPolicies *Eem_AvlPolicies) GetEntityData() *types.CommonEntityData {
    avlPolicies.EntityData.YFilter = avlPolicies.YFilter
    avlPolicies.EntityData.YangName = "avl-policies"
    avlPolicies.EntityData.BundleName = "cisco_ios_xr"
    avlPolicies.EntityData.ParentYangName = "eem"
    avlPolicies.EntityData.SegmentPath = "avl-policies"
    avlPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    avlPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    avlPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    avlPolicies.EntityData.Children = make(map[string]types.YChild)
    avlPolicies.EntityData.Children["avl-policy"] = types.YChild{"AvlPolicy", nil}
    for i := range avlPolicies.AvlPolicy {
        avlPolicies.EntityData.Children[types.GetSegmentPath(&avlPolicies.AvlPolicy[i])] = types.YChild{"AvlPolicy", &avlPolicies.AvlPolicy[i]}
    }
    avlPolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(avlPolicies.EntityData)
}

// Eem_AvlPolicies_AvlPolicy
// policy name and create time 
type Eem_AvlPolicies_AvlPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. System policy name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // policy type. The type is string.
    Type_ interface{}

    // time created. The type is string.
    TimeCreated interface{}

    // policy name. The type is string.
    PolicyName interface{}
}

func (avlPolicy *Eem_AvlPolicies_AvlPolicy) GetEntityData() *types.CommonEntityData {
    avlPolicy.EntityData.YFilter = avlPolicy.YFilter
    avlPolicy.EntityData.YangName = "avl-policy"
    avlPolicy.EntityData.BundleName = "cisco_ios_xr"
    avlPolicy.EntityData.ParentYangName = "avl-policies"
    avlPolicy.EntityData.SegmentPath = "avl-policy" + "[name='" + fmt.Sprintf("%v", avlPolicy.Name) + "']"
    avlPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    avlPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    avlPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    avlPolicy.EntityData.Children = make(map[string]types.YChild)
    avlPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    avlPolicy.EntityData.Leafs["name"] = types.YLeaf{"Name", avlPolicy.Name}
    avlPolicy.EntityData.Leafs["type"] = types.YLeaf{"Type_", avlPolicy.Type_}
    avlPolicy.EntityData.Leafs["time-created"] = types.YLeaf{"TimeCreated", avlPolicy.TimeCreated}
    avlPolicy.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", avlPolicy.PolicyName}
    return &(avlPolicy.EntityData)
}

