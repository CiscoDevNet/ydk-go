// This module contains a collection of YANG definitions
// for Cisco IOS-XR ha-eem package configuration.
// 
// This module contains definitions
// for the following management objects:
//   event-manager: Event manager configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ha_eem_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ha_eem_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ha-eem-cfg event-manager}", reflect.TypeOf(EventManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ha-eem-cfg:event-manager", reflect.TypeOf(EventManager{}))
}

// EventManagerChecksum represents Event manager checksum
type EventManagerChecksum string

const (
    // Use SHA-1 checksum
    EventManagerChecksum_sha_1 EventManagerChecksum = "sha-1"

    // Use MD5 checksum
    EventManagerChecksum_md5 EventManagerChecksum = "md5"
)

// EventManagerPolicySec represents Event manager policy sec
type EventManagerPolicySec string

const (
    // Cisco Signature
    EventManagerPolicySec_rsa_2048 EventManagerPolicySec = "rsa-2048"

    // Trust Signature
    EventManagerPolicySec_trust EventManagerPolicySec = "trust"
)

// EventManagerPolicyMode represents Event manager policy mode
type EventManagerPolicyMode string

const (
    // Cisco Signature
    EventManagerPolicyMode_cisco EventManagerPolicyMode = "cisco"

    // Trust Signature
    EventManagerPolicyMode_trust EventManagerPolicyMode = "trust"
)

// EventManagerPolicy represents Event manager policy
type EventManagerPolicy string

const (
    // Event manager system policy
    EventManagerPolicy_system EventManagerPolicy = "system"

    // Event manager user policy
    EventManagerPolicy_user EventManagerPolicy = "user"
)

// EventManager
// Event manager configuration
type EventManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set refresh time (in seconds) for policy username's AAA taskmap. The type
    // is interface{} with range: 10..4294967295. Units are second. The default
    // value is 1800.
    RefreshTime interface{}

    // Enable suspend policy scheduling. The type is bool.
    ScheduleSuspend interface{}

    // Set event manager user policy directory. The type is string.
    DirectoryUserPolicy interface{}

    // Path of the user policy library directory. The type is string.
    DirectoryUserLibrary interface{}

    // Register an event manager policy.
    Policies EventManager_Policies

    // scheduler classs type.
    SchedulerScript EventManager_SchedulerScript

    // Set an event manager global variable for event manager policies.
    Environments EventManager_Environments
}

func (eventManager *EventManager) GetEntityData() *types.CommonEntityData {
    eventManager.EntityData.YFilter = eventManager.YFilter
    eventManager.EntityData.YangName = "event-manager"
    eventManager.EntityData.BundleName = "cisco_ios_xr"
    eventManager.EntityData.ParentYangName = "Cisco-IOS-XR-ha-eem-cfg"
    eventManager.EntityData.SegmentPath = "Cisco-IOS-XR-ha-eem-cfg:event-manager"
    eventManager.EntityData.AbsolutePath = eventManager.EntityData.SegmentPath
    eventManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventManager.EntityData.Children = types.NewOrderedMap()
    eventManager.EntityData.Children.Append("policies", types.YChild{"Policies", &eventManager.Policies})
    eventManager.EntityData.Children.Append("scheduler-script", types.YChild{"SchedulerScript", &eventManager.SchedulerScript})
    eventManager.EntityData.Children.Append("environments", types.YChild{"Environments", &eventManager.Environments})
    eventManager.EntityData.Leafs = types.NewOrderedMap()
    eventManager.EntityData.Leafs.Append("refresh-time", types.YLeaf{"RefreshTime", eventManager.RefreshTime})
    eventManager.EntityData.Leafs.Append("schedule-suspend", types.YLeaf{"ScheduleSuspend", eventManager.ScheduleSuspend})
    eventManager.EntityData.Leafs.Append("directory-user-policy", types.YLeaf{"DirectoryUserPolicy", eventManager.DirectoryUserPolicy})
    eventManager.EntityData.Leafs.Append("directory-user-library", types.YLeaf{"DirectoryUserLibrary", eventManager.DirectoryUserLibrary})

    eventManager.EntityData.YListKeys = []string {}

    return &(eventManager.EntityData)
}

// EventManager_Policies
// Register an event manager policy
type EventManager_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the policy file. The type is slice of EventManager_Policies_Policy.
    Policy []*EventManager_Policies_Policy
}

func (policies *EventManager_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "event-manager"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/" + policies.EntityData.SegmentPath
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range policies.Policy {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.Policy[i]), types.YChild{"Policy", policies.Policy[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// EventManager_Policies_Policy
// Name of the policy file
type EventManager_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the policy file. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PolicyName interface{}

    // A configured username. The type is string. This attribute is mandatory.
    Username interface{}

    // Time of validity (in seconds) for cached AAA taskmap of username (default
    // is 3600). The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory. Units are second.
    PersistTime interface{}

    // Event manager type of this policy. The type is EventManagerPolicy.
    PolicyType interface{}

    // Specify Embedded Event Manager policy checksum. The type is
    // EventManagerChecksum.
    ChecksumType interface{}

    // CheckSum Value. The type is string.
    CheckSumValue interface{}

    // Specify Embedded Event Manager policy security mode. The type is
    // EventManagerPolicyMode.
    PolicySecurityMode interface{}

    // Event Manager policy security Level. The type is EventManagerPolicySec.
    PolicySecurityLevel interface{}
}

func (policy *EventManager_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.PolicyName, "policy-name")
    policy.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/policies/" + policy.EntityData.SegmentPath
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policy.PolicyName})
    policy.EntityData.Leafs.Append("username", types.YLeaf{"Username", policy.Username})
    policy.EntityData.Leafs.Append("persist-time", types.YLeaf{"PersistTime", policy.PersistTime})
    policy.EntityData.Leafs.Append("policy-type", types.YLeaf{"PolicyType", policy.PolicyType})
    policy.EntityData.Leafs.Append("checksum-type", types.YLeaf{"ChecksumType", policy.ChecksumType})
    policy.EntityData.Leafs.Append("check-sum-value", types.YLeaf{"CheckSumValue", policy.CheckSumValue})
    policy.EntityData.Leafs.Append("policy-security-mode", types.YLeaf{"PolicySecurityMode", policy.PolicySecurityMode})
    policy.EntityData.Leafs.Append("policy-security-level", types.YLeaf{"PolicySecurityLevel", policy.PolicySecurityLevel})

    policy.EntityData.YListKeys = []string {"PolicyName"}

    return &(policy.EntityData)
}

// EventManager_SchedulerScript
// scheduler classs type
type EventManager_SchedulerScript struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // scheduler thread classs .
    ThreadClasses EventManager_SchedulerScript_ThreadClasses
}

func (schedulerScript *EventManager_SchedulerScript) GetEntityData() *types.CommonEntityData {
    schedulerScript.EntityData.YFilter = schedulerScript.YFilter
    schedulerScript.EntityData.YangName = "scheduler-script"
    schedulerScript.EntityData.BundleName = "cisco_ios_xr"
    schedulerScript.EntityData.ParentYangName = "event-manager"
    schedulerScript.EntityData.SegmentPath = "scheduler-script"
    schedulerScript.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/" + schedulerScript.EntityData.SegmentPath
    schedulerScript.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedulerScript.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedulerScript.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedulerScript.EntityData.Children = types.NewOrderedMap()
    schedulerScript.EntityData.Children.Append("thread-classes", types.YChild{"ThreadClasses", &schedulerScript.ThreadClasses})
    schedulerScript.EntityData.Leafs = types.NewOrderedMap()

    schedulerScript.EntityData.YListKeys = []string {}

    return &(schedulerScript.EntityData)
}

// EventManager_SchedulerScript_ThreadClasses
// scheduler thread classs 
type EventManager_SchedulerScript_ThreadClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // scheduler classs type argument. The type is slice of
    // EventManager_SchedulerScript_ThreadClasses_ThreadClass.
    ThreadClass []*EventManager_SchedulerScript_ThreadClasses_ThreadClass
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetEntityData() *types.CommonEntityData {
    threadClasses.EntityData.YFilter = threadClasses.YFilter
    threadClasses.EntityData.YangName = "thread-classes"
    threadClasses.EntityData.BundleName = "cisco_ios_xr"
    threadClasses.EntityData.ParentYangName = "scheduler-script"
    threadClasses.EntityData.SegmentPath = "thread-classes"
    threadClasses.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/scheduler-script/" + threadClasses.EntityData.SegmentPath
    threadClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threadClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threadClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threadClasses.EntityData.Children = types.NewOrderedMap()
    threadClasses.EntityData.Children.Append("thread-class", types.YChild{"ThreadClass", nil})
    for i := range threadClasses.ThreadClass {
        threadClasses.EntityData.Children.Append(types.GetSegmentPath(threadClasses.ThreadClass[i]), types.YChild{"ThreadClass", threadClasses.ThreadClass[i]})
    }
    threadClasses.EntityData.Leafs = types.NewOrderedMap()

    threadClasses.EntityData.YListKeys = []string {}

    return &(threadClasses.EntityData)
}

// EventManager_SchedulerScript_ThreadClasses_ThreadClass
// scheduler classs type argument
type EventManager_SchedulerScript_ThreadClasses_ThreadClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the global variable. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ThreadClassName interface{}

    // number of scheduler threads. The type is interface{} with range: 1..5. This
    // attribute is mandatory.
    NumThreads interface{}
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetEntityData() *types.CommonEntityData {
    threadClass.EntityData.YFilter = threadClass.YFilter
    threadClass.EntityData.YangName = "thread-class"
    threadClass.EntityData.BundleName = "cisco_ios_xr"
    threadClass.EntityData.ParentYangName = "thread-classes"
    threadClass.EntityData.SegmentPath = "thread-class" + types.AddKeyToken(threadClass.ThreadClassName, "thread-class-name")
    threadClass.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/scheduler-script/thread-classes/" + threadClass.EntityData.SegmentPath
    threadClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threadClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threadClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threadClass.EntityData.Children = types.NewOrderedMap()
    threadClass.EntityData.Leafs = types.NewOrderedMap()
    threadClass.EntityData.Leafs.Append("thread-class-name", types.YLeaf{"ThreadClassName", threadClass.ThreadClassName})
    threadClass.EntityData.Leafs.Append("num-threads", types.YLeaf{"NumThreads", threadClass.NumThreads})

    threadClass.EntityData.YListKeys = []string {"ThreadClassName"}

    return &(threadClass.EntityData)
}

// EventManager_Environments
// Set an event manager global variable for event
// manager policies
type EventManager_Environments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the global variable. The type is slice of
    // EventManager_Environments_Environment.
    Environment []*EventManager_Environments_Environment
}

func (environments *EventManager_Environments) GetEntityData() *types.CommonEntityData {
    environments.EntityData.YFilter = environments.YFilter
    environments.EntityData.YangName = "environments"
    environments.EntityData.BundleName = "cisco_ios_xr"
    environments.EntityData.ParentYangName = "event-manager"
    environments.EntityData.SegmentPath = "environments"
    environments.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/" + environments.EntityData.SegmentPath
    environments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environments.EntityData.Children = types.NewOrderedMap()
    environments.EntityData.Children.Append("environment", types.YChild{"Environment", nil})
    for i := range environments.Environment {
        environments.EntityData.Children.Append(types.GetSegmentPath(environments.Environment[i]), types.YChild{"Environment", environments.Environment[i]})
    }
    environments.EntityData.Leafs = types.NewOrderedMap()

    environments.EntityData.YListKeys = []string {}

    return &(environments.EntityData)
}

// EventManager_Environments_Environment
// Name of the global variable
type EventManager_Environments_Environment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the global variable. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    EnvironmentName interface{}

    // Value of the global variable. The type is string. This attribute is
    // mandatory.
    EnvironmentValue interface{}
}

func (environment *EventManager_Environments_Environment) GetEntityData() *types.CommonEntityData {
    environment.EntityData.YFilter = environment.YFilter
    environment.EntityData.YangName = "environment"
    environment.EntityData.BundleName = "cisco_ios_xr"
    environment.EntityData.ParentYangName = "environments"
    environment.EntityData.SegmentPath = "environment" + types.AddKeyToken(environment.EnvironmentName, "environment-name")
    environment.EntityData.AbsolutePath = "Cisco-IOS-XR-ha-eem-cfg:event-manager/environments/" + environment.EntityData.SegmentPath
    environment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environment.EntityData.Children = types.NewOrderedMap()
    environment.EntityData.Leafs = types.NewOrderedMap()
    environment.EntityData.Leafs.Append("environment-name", types.YLeaf{"EnvironmentName", environment.EnvironmentName})
    environment.EntityData.Leafs.Append("environment-value", types.YLeaf{"EnvironmentValue", environment.EnvironmentValue})

    environment.EntityData.YListKeys = []string {"EnvironmentName"}

    return &(environment.EntityData)
}

