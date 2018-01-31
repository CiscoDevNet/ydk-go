// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-lib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   aaa: Authentication, Authorization and Accounting
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_lib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_lib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-lib-cfg aaa}", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-lib-cfg:aaa", reflect.TypeOf(Aaa{}))
}

// Aaa
// Authentication, Authorization and Accounting
type Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This class is used for setting the default taskgroup to be used for remote
    // server authentication. The type is string.
    DefaultTaskgroup interface{}

    // Enable LI RADIUS Feature. The type is interface{}.
    Intercept interface{}

    // AAA accounting.
    Accountings Aaa_Accountings

    // AAA authorization.
    Authorizations Aaa_Authorizations

    // Configuration related to 'update' accounting.
    AccountingUpdate Aaa_AccountingUpdate

    // AAA banner.
    Banner Aaa_Banner

    // AAA authentication.
    Authentications Aaa_Authentications

    // AAA subscriber.
    AaaSubscriber Aaa_AaaSubscriber

    // AAA Mobile.
    AaaMobile Aaa_AaaMobile

    // AAA Dot1x.
    AaaDot1X Aaa_AaaDot1X

    // AAA RADIUS attribute configurations.
    RadiusAttribute Aaa_RadiusAttribute

    // AAA group definitions.
    ServerGroups Aaa_ServerGroups

    // Configure local usernames.
    Usernames Aaa_Usernames

    // Specify a taskgroup to inherit from.
    Taskgroups Aaa_Taskgroups

    // Specify a Usergroup to inherit from.
    Usergroups Aaa_Usergroups

    // Diameter base protocol.
    Diameter Aaa_Diameter

    // Remote Access Dial-In User Service.
    Radius Aaa_Radius

    // Modify TACACS+ query parameters.
    Tacacs Aaa_Tacacs
}

func (aaa *Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Aaa) GetGoName(yname string) string {
    if yname == "default-taskgroup" { return "DefaultTaskgroup" }
    if yname == "intercept" { return "Intercept" }
    if yname == "accountings" { return "Accountings" }
    if yname == "authorizations" { return "Authorizations" }
    if yname == "accounting-update" { return "AccountingUpdate" }
    if yname == "banner" { return "Banner" }
    if yname == "authentications" { return "Authentications" }
    if yname == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber" { return "AaaSubscriber" }
    if yname == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile" { return "AaaMobile" }
    if yname == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x" { return "AaaDot1X" }
    if yname == "Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute" { return "RadiusAttribute" }
    if yname == "Cisco-IOS-XR-aaa-locald-cfg:server-groups" { return "ServerGroups" }
    if yname == "Cisco-IOS-XR-aaa-locald-cfg:usernames" { return "Usernames" }
    if yname == "Cisco-IOS-XR-aaa-locald-cfg:taskgroups" { return "Taskgroups" }
    if yname == "Cisco-IOS-XR-aaa-locald-cfg:usergroups" { return "Usergroups" }
    if yname == "Cisco-IOS-XR-aaa-diameter-cfg:diameter" { return "Diameter" }
    if yname == "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius" { return "Radius" }
    if yname == "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs" { return "Tacacs" }
    return ""
}

func (aaa *Aaa) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-lib-cfg:aaa"
}

func (aaa *Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accountings" {
        return &aaa.Accountings
    }
    if childYangName == "authorizations" {
        return &aaa.Authorizations
    }
    if childYangName == "accounting-update" {
        return &aaa.AccountingUpdate
    }
    if childYangName == "banner" {
        return &aaa.Banner
    }
    if childYangName == "authentications" {
        return &aaa.Authentications
    }
    if childYangName == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber" {
        return &aaa.AaaSubscriber
    }
    if childYangName == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile" {
        return &aaa.AaaMobile
    }
    if childYangName == "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x" {
        return &aaa.AaaDot1X
    }
    if childYangName == "Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute" {
        return &aaa.RadiusAttribute
    }
    if childYangName == "Cisco-IOS-XR-aaa-locald-cfg:server-groups" {
        return &aaa.ServerGroups
    }
    if childYangName == "Cisco-IOS-XR-aaa-locald-cfg:usernames" {
        return &aaa.Usernames
    }
    if childYangName == "Cisco-IOS-XR-aaa-locald-cfg:taskgroups" {
        return &aaa.Taskgroups
    }
    if childYangName == "Cisco-IOS-XR-aaa-locald-cfg:usergroups" {
        return &aaa.Usergroups
    }
    if childYangName == "Cisco-IOS-XR-aaa-diameter-cfg:diameter" {
        return &aaa.Diameter
    }
    if childYangName == "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius" {
        return &aaa.Radius
    }
    if childYangName == "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs" {
        return &aaa.Tacacs
    }
    return nil
}

func (aaa *Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accountings"] = &aaa.Accountings
    children["authorizations"] = &aaa.Authorizations
    children["accounting-update"] = &aaa.AccountingUpdate
    children["banner"] = &aaa.Banner
    children["authentications"] = &aaa.Authentications
    children["Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber"] = &aaa.AaaSubscriber
    children["Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile"] = &aaa.AaaMobile
    children["Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x"] = &aaa.AaaDot1X
    children["Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute"] = &aaa.RadiusAttribute
    children["Cisco-IOS-XR-aaa-locald-cfg:server-groups"] = &aaa.ServerGroups
    children["Cisco-IOS-XR-aaa-locald-cfg:usernames"] = &aaa.Usernames
    children["Cisco-IOS-XR-aaa-locald-cfg:taskgroups"] = &aaa.Taskgroups
    children["Cisco-IOS-XR-aaa-locald-cfg:usergroups"] = &aaa.Usergroups
    children["Cisco-IOS-XR-aaa-diameter-cfg:diameter"] = &aaa.Diameter
    children["Cisco-IOS-XR-aaa-protocol-radius-cfg:radius"] = &aaa.Radius
    children["Cisco-IOS-XR-aaa-tacacs-cfg:tacacs"] = &aaa.Tacacs
    return children
}

func (aaa *Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-taskgroup"] = aaa.DefaultTaskgroup
    leafs["intercept"] = aaa.Intercept
    return leafs
}

func (aaa *Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Aaa) GetYangName() string { return "aaa" }

func (aaa *Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Aaa) GetParentYangName() string { return "Cisco-IOS-XR-aaa-lib-cfg" }

// Aaa_Accountings
// AAA accounting
type Aaa_Accountings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_Accountings_Accounting.
    Accounting []Aaa_Accountings_Accounting
}

func (accountings *Aaa_Accountings) GetFilter() yfilter.YFilter { return accountings.YFilter }

func (accountings *Aaa_Accountings) SetFilter(yf yfilter.YFilter) { accountings.YFilter = yf }

func (accountings *Aaa_Accountings) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (accountings *Aaa_Accountings) GetSegmentPath() string {
    return "accountings"
}

func (accountings *Aaa_Accountings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        for _, c := range accountings.Accounting {
            if accountings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Accountings_Accounting{}
        accountings.Accounting = append(accountings.Accounting, child)
        return &accountings.Accounting[len(accountings.Accounting)-1]
    }
    return nil
}

func (accountings *Aaa_Accountings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accountings.Accounting {
        children[accountings.Accounting[i].GetSegmentPath()] = &accountings.Accounting[i]
    }
    return children
}

func (accountings *Aaa_Accountings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accountings *Aaa_Accountings) GetBundleName() string { return "cisco_ios_xr" }

func (accountings *Aaa_Accountings) GetYangName() string { return "accountings" }

func (accountings *Aaa_Accountings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountings *Aaa_Accountings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountings *Aaa_Accountings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountings *Aaa_Accountings) SetParent(parent types.Entity) { accountings.parent = parent }

func (accountings *Aaa_Accountings) GetParent() types.Entity { return accountings.parent }

func (accountings *Aaa_Accountings) GetParentYangName() string { return "aaa" }

// Aaa_Accountings_Accounting
// Configurations related to accounting
type Aaa_Accountings_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. exec:Account exec sessions, commands: Account CLI
    // commands. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // rpfailover. The type is AaaAccountingRpFailover.
    RpFailover interface{}

    // Broadcast. The type is AaaAccountingBroadcast.
    Broadcast interface{}

    // Stop only/Start Stop. The type is AaaAccounting.
    TypeXr interface{}

    // Method Type. The type is AaaMethod.
    Method1 interface{}

    // Method Type. The type is AaaMethod.
    Method2 interface{}

    // Method Type. The type is AaaMethod.
    Method3 interface{}

    // Method Type. The type is AaaMethod.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (accounting *Aaa_Accountings_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Aaa_Accountings_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Aaa_Accountings_Accounting) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "rp-failover" { return "RpFailover" }
    if yname == "broadcast" { return "Broadcast" }
    if yname == "type-xr" { return "TypeXr" }
    if yname == "method1" { return "Method1" }
    if yname == "method2" { return "Method2" }
    if yname == "method3" { return "Method3" }
    if yname == "method4" { return "Method4" }
    if yname == "server-group-name1" { return "ServerGroupName1" }
    if yname == "server-group-name2" { return "ServerGroupName2" }
    if yname == "server-group-name3" { return "ServerGroupName3" }
    if yname == "server-group-name4" { return "ServerGroupName4" }
    return ""
}

func (accounting *Aaa_Accountings_Accounting) GetSegmentPath() string {
    return "accounting" + "[type='" + fmt.Sprintf("%v", accounting.Type) + "']" + "[listname='" + fmt.Sprintf("%v", accounting.Listname) + "']"
}

func (accounting *Aaa_Accountings_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Aaa_Accountings_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Aaa_Accountings_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = accounting.Type
    leafs["listname"] = accounting.Listname
    leafs["rp-failover"] = accounting.RpFailover
    leafs["broadcast"] = accounting.Broadcast
    leafs["type-xr"] = accounting.TypeXr
    leafs["method1"] = accounting.Method1
    leafs["method2"] = accounting.Method2
    leafs["method3"] = accounting.Method3
    leafs["method4"] = accounting.Method4
    leafs["server-group-name1"] = accounting.ServerGroupName1
    leafs["server-group-name2"] = accounting.ServerGroupName2
    leafs["server-group-name3"] = accounting.ServerGroupName3
    leafs["server-group-name4"] = accounting.ServerGroupName4
    return leafs
}

func (accounting *Aaa_Accountings_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Aaa_Accountings_Accounting) GetYangName() string { return "accounting" }

func (accounting *Aaa_Accountings_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Aaa_Accountings_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Aaa_Accountings_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Aaa_Accountings_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Aaa_Accountings_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Aaa_Accountings_Accounting) GetParentYangName() string { return "accountings" }

// Aaa_Authorizations
// AAA authorization
type Aaa_Authorizations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_Authorizations_Authorization.
    Authorization []Aaa_Authorizations_Authorization
}

func (authorizations *Aaa_Authorizations) GetFilter() yfilter.YFilter { return authorizations.YFilter }

func (authorizations *Aaa_Authorizations) SetFilter(yf yfilter.YFilter) { authorizations.YFilter = yf }

func (authorizations *Aaa_Authorizations) GetGoName(yname string) string {
    if yname == "authorization" { return "Authorization" }
    return ""
}

func (authorizations *Aaa_Authorizations) GetSegmentPath() string {
    return "authorizations"
}

func (authorizations *Aaa_Authorizations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authorization" {
        for _, c := range authorizations.Authorization {
            if authorizations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Authorizations_Authorization{}
        authorizations.Authorization = append(authorizations.Authorization, child)
        return &authorizations.Authorization[len(authorizations.Authorization)-1]
    }
    return nil
}

func (authorizations *Aaa_Authorizations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authorizations.Authorization {
        children[authorizations.Authorization[i].GetSegmentPath()] = &authorizations.Authorization[i]
    }
    return children
}

func (authorizations *Aaa_Authorizations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authorizations *Aaa_Authorizations) GetBundleName() string { return "cisco_ios_xr" }

func (authorizations *Aaa_Authorizations) GetYangName() string { return "authorizations" }

func (authorizations *Aaa_Authorizations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorizations *Aaa_Authorizations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorizations *Aaa_Authorizations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorizations *Aaa_Authorizations) SetParent(parent types.Entity) { authorizations.parent = parent }

func (authorizations *Aaa_Authorizations) GetParent() types.Entity { return authorizations.parent }

func (authorizations *Aaa_Authorizations) GetParentYangName() string { return "aaa" }

// Aaa_Authorizations_Authorization
// Configurations related to authorization
type Aaa_Authorizations_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. network: Authorize IKE requests, commands:
    // Authorize CLI commands. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. List name for AAA authorization. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Type. The type is AaaMethod.
    Method1 interface{}

    // Method Type. The type is AaaMethod.
    Method2 interface{}

    // Method Type. The type is AaaMethod.
    Method3 interface{}

    // Method Type. The type is AaaMethod.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (authorization *Aaa_Authorizations_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *Aaa_Authorizations_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *Aaa_Authorizations_Authorization) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method1" { return "Method1" }
    if yname == "method2" { return "Method2" }
    if yname == "method3" { return "Method3" }
    if yname == "method4" { return "Method4" }
    if yname == "server-group-name1" { return "ServerGroupName1" }
    if yname == "server-group-name2" { return "ServerGroupName2" }
    if yname == "server-group-name3" { return "ServerGroupName3" }
    if yname == "server-group-name4" { return "ServerGroupName4" }
    return ""
}

func (authorization *Aaa_Authorizations_Authorization) GetSegmentPath() string {
    return "authorization" + "[type='" + fmt.Sprintf("%v", authorization.Type) + "']" + "[listname='" + fmt.Sprintf("%v", authorization.Listname) + "']"
}

func (authorization *Aaa_Authorizations_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorization *Aaa_Authorizations_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorization *Aaa_Authorizations_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = authorization.Type
    leafs["listname"] = authorization.Listname
    leafs["method1"] = authorization.Method1
    leafs["method2"] = authorization.Method2
    leafs["method3"] = authorization.Method3
    leafs["method4"] = authorization.Method4
    leafs["server-group-name1"] = authorization.ServerGroupName1
    leafs["server-group-name2"] = authorization.ServerGroupName2
    leafs["server-group-name3"] = authorization.ServerGroupName3
    leafs["server-group-name4"] = authorization.ServerGroupName4
    return leafs
}

func (authorization *Aaa_Authorizations_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *Aaa_Authorizations_Authorization) GetYangName() string { return "authorization" }

func (authorization *Aaa_Authorizations_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *Aaa_Authorizations_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *Aaa_Authorizations_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *Aaa_Authorizations_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *Aaa_Authorizations_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *Aaa_Authorizations_Authorization) GetParentYangName() string { return "authorizations" }

// Aaa_AccountingUpdate
// Configuration related to 'update' accounting
// This type is a presence type.
type Aaa_AccountingUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // newinfo/periodic. The type is AaaAccountingUpdate. This attribute is
    // mandatory.
    Type interface{}

    // Periodic update interval in minutes. The type is interface{} with range:
    // 0..35791394. Units are minute.
    PeriodicInterval interface{}
}

func (accountingUpdate *Aaa_AccountingUpdate) GetFilter() yfilter.YFilter { return accountingUpdate.YFilter }

func (accountingUpdate *Aaa_AccountingUpdate) SetFilter(yf yfilter.YFilter) { accountingUpdate.YFilter = yf }

func (accountingUpdate *Aaa_AccountingUpdate) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    return ""
}

func (accountingUpdate *Aaa_AccountingUpdate) GetSegmentPath() string {
    return "accounting-update"
}

func (accountingUpdate *Aaa_AccountingUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accountingUpdate *Aaa_AccountingUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accountingUpdate *Aaa_AccountingUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = accountingUpdate.Type
    leafs["periodic-interval"] = accountingUpdate.PeriodicInterval
    return leafs
}

func (accountingUpdate *Aaa_AccountingUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (accountingUpdate *Aaa_AccountingUpdate) GetYangName() string { return "accounting-update" }

func (accountingUpdate *Aaa_AccountingUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountingUpdate *Aaa_AccountingUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountingUpdate *Aaa_AccountingUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountingUpdate *Aaa_AccountingUpdate) SetParent(parent types.Entity) { accountingUpdate.parent = parent }

func (accountingUpdate *Aaa_AccountingUpdate) GetParent() types.Entity { return accountingUpdate.parent }

func (accountingUpdate *Aaa_AccountingUpdate) GetParentYangName() string { return "aaa" }

// Aaa_Banner
// AAA banner
type Aaa_Banner struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA login banner. The type is string.
    Login interface{}
}

func (banner *Aaa_Banner) GetFilter() yfilter.YFilter { return banner.YFilter }

func (banner *Aaa_Banner) SetFilter(yf yfilter.YFilter) { banner.YFilter = yf }

func (banner *Aaa_Banner) GetGoName(yname string) string {
    if yname == "login" { return "Login" }
    return ""
}

func (banner *Aaa_Banner) GetSegmentPath() string {
    return "banner"
}

func (banner *Aaa_Banner) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (banner *Aaa_Banner) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (banner *Aaa_Banner) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["login"] = banner.Login
    return leafs
}

func (banner *Aaa_Banner) GetBundleName() string { return "cisco_ios_xr" }

func (banner *Aaa_Banner) GetYangName() string { return "banner" }

func (banner *Aaa_Banner) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (banner *Aaa_Banner) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (banner *Aaa_Banner) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (banner *Aaa_Banner) SetParent(parent types.Entity) { banner.parent = parent }

func (banner *Aaa_Banner) GetParent() types.Entity { return banner.parent }

func (banner *Aaa_Banner) GetParentYangName() string { return "aaa" }

// Aaa_Authentications
// AAA authentication
type Aaa_Authentications struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_Authentications_Authentication.
    Authentication []Aaa_Authentications_Authentication
}

func (authentications *Aaa_Authentications) GetFilter() yfilter.YFilter { return authentications.YFilter }

func (authentications *Aaa_Authentications) SetFilter(yf yfilter.YFilter) { authentications.YFilter = yf }

func (authentications *Aaa_Authentications) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (authentications *Aaa_Authentications) GetSegmentPath() string {
    return "authentications"
}

func (authentications *Aaa_Authentications) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        for _, c := range authentications.Authentication {
            if authentications.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Authentications_Authentication{}
        authentications.Authentication = append(authentications.Authentication, child)
        return &authentications.Authentication[len(authentications.Authentication)-1]
    }
    return nil
}

func (authentications *Aaa_Authentications) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authentications.Authentication {
        children[authentications.Authentication[i].GetSegmentPath()] = &authentications.Authentication[i]
    }
    return children
}

func (authentications *Aaa_Authentications) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentications *Aaa_Authentications) GetBundleName() string { return "cisco_ios_xr" }

func (authentications *Aaa_Authentications) GetYangName() string { return "authentications" }

func (authentications *Aaa_Authentications) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentications *Aaa_Authentications) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentications *Aaa_Authentications) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentications *Aaa_Authentications) SetParent(parent types.Entity) { authentications.parent = parent }

func (authentications *Aaa_Authentications) GetParent() types.Entity { return authentications.parent }

func (authentications *Aaa_Authentications) GetParentYangName() string { return "aaa" }

// Aaa_Authentications_Authentication
// Configurations related to authentication
type Aaa_Authentications_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. login: Authenticate login sessions, ppp:
    // Authenticate ppp sessions. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. List name for AAA authentication. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Type. The type is AaaMethod.
    Method1 interface{}

    // Method Type. The type is AaaMethod.
    Method2 interface{}

    // Method Type. The type is AaaMethod.
    Method3 interface{}

    // Method Type. The type is AaaMethod.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (authentication *Aaa_Authentications_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Aaa_Authentications_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Aaa_Authentications_Authentication) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method1" { return "Method1" }
    if yname == "method2" { return "Method2" }
    if yname == "method3" { return "Method3" }
    if yname == "method4" { return "Method4" }
    if yname == "server-group-name1" { return "ServerGroupName1" }
    if yname == "server-group-name2" { return "ServerGroupName2" }
    if yname == "server-group-name3" { return "ServerGroupName3" }
    if yname == "server-group-name4" { return "ServerGroupName4" }
    return ""
}

func (authentication *Aaa_Authentications_Authentication) GetSegmentPath() string {
    return "authentication" + "[type='" + fmt.Sprintf("%v", authentication.Type) + "']" + "[listname='" + fmt.Sprintf("%v", authentication.Listname) + "']"
}

func (authentication *Aaa_Authentications_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Aaa_Authentications_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Aaa_Authentications_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = authentication.Type
    leafs["listname"] = authentication.Listname
    leafs["method1"] = authentication.Method1
    leafs["method2"] = authentication.Method2
    leafs["method3"] = authentication.Method3
    leafs["method4"] = authentication.Method4
    leafs["server-group-name1"] = authentication.ServerGroupName1
    leafs["server-group-name2"] = authentication.ServerGroupName2
    leafs["server-group-name3"] = authentication.ServerGroupName3
    leafs["server-group-name4"] = authentication.ServerGroupName4
    return leafs
}

func (authentication *Aaa_Authentications_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Aaa_Authentications_Authentication) GetYangName() string { return "authentication" }

func (authentication *Aaa_Authentications_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Aaa_Authentications_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Aaa_Authentications_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Aaa_Authentications_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Aaa_Authentications_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Aaa_Authentications_Authentication) GetParentYangName() string { return "authentications" }

// Aaa_AaaSubscriber
// AAA subscriber
type Aaa_AaaSubscriber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA authorization policy.
    PolicyIfAuthors Aaa_AaaSubscriber_PolicyIfAuthors

    // AAA accounting.
    Accountings Aaa_AaaSubscriber_Accountings

    // Set accounting parameters for Service.
    ServiceAccounting Aaa_AaaSubscriber_ServiceAccounting

    // AAA authorization prepaid.
    PrepaidAuthors Aaa_AaaSubscriber_PrepaidAuthors

    // AAA authorization.
    Authorizations Aaa_AaaSubscriber_Authorizations

    // AAA authentication.
    Authentications Aaa_AaaSubscriber_Authentications
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetFilter() yfilter.YFilter { return aaaSubscriber.YFilter }

func (aaaSubscriber *Aaa_AaaSubscriber) SetFilter(yf yfilter.YFilter) { aaaSubscriber.YFilter = yf }

func (aaaSubscriber *Aaa_AaaSubscriber) GetGoName(yname string) string {
    if yname == "policy-if-authors" { return "PolicyIfAuthors" }
    if yname == "accountings" { return "Accountings" }
    if yname == "service-accounting" { return "ServiceAccounting" }
    if yname == "prepaid-authors" { return "PrepaidAuthors" }
    if yname == "authorizations" { return "Authorizations" }
    if yname == "authentications" { return "Authentications" }
    return ""
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber"
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-if-authors" {
        return &aaaSubscriber.PolicyIfAuthors
    }
    if childYangName == "accountings" {
        return &aaaSubscriber.Accountings
    }
    if childYangName == "service-accounting" {
        return &aaaSubscriber.ServiceAccounting
    }
    if childYangName == "prepaid-authors" {
        return &aaaSubscriber.PrepaidAuthors
    }
    if childYangName == "authorizations" {
        return &aaaSubscriber.Authorizations
    }
    if childYangName == "authentications" {
        return &aaaSubscriber.Authentications
    }
    return nil
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-if-authors"] = &aaaSubscriber.PolicyIfAuthors
    children["accountings"] = &aaaSubscriber.Accountings
    children["service-accounting"] = &aaaSubscriber.ServiceAccounting
    children["prepaid-authors"] = &aaaSubscriber.PrepaidAuthors
    children["authorizations"] = &aaaSubscriber.Authorizations
    children["authentications"] = &aaaSubscriber.Authentications
    return children
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetBundleName() string { return "cisco_ios_xr" }

func (aaaSubscriber *Aaa_AaaSubscriber) GetYangName() string { return "aaa-subscriber" }

func (aaaSubscriber *Aaa_AaaSubscriber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaaSubscriber *Aaa_AaaSubscriber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaaSubscriber *Aaa_AaaSubscriber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaaSubscriber *Aaa_AaaSubscriber) SetParent(parent types.Entity) { aaaSubscriber.parent = parent }

func (aaaSubscriber *Aaa_AaaSubscriber) GetParent() types.Entity { return aaaSubscriber.parent }

func (aaaSubscriber *Aaa_AaaSubscriber) GetParentYangName() string { return "aaa" }

// Aaa_AaaSubscriber_PolicyIfAuthors
// AAA authorization policy
type Aaa_AaaSubscriber_PolicyIfAuthors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor.
    PolicyIfAuthor []Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetFilter() yfilter.YFilter { return policyIfAuthors.YFilter }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) SetFilter(yf yfilter.YFilter) { policyIfAuthors.YFilter = yf }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetGoName(yname string) string {
    if yname == "policy-if-author" { return "PolicyIfAuthor" }
    return ""
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetSegmentPath() string {
    return "policy-if-authors"
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-if-author" {
        for _, c := range policyIfAuthors.PolicyIfAuthor {
            if policyIfAuthors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor{}
        policyIfAuthors.PolicyIfAuthor = append(policyIfAuthors.PolicyIfAuthor, child)
        return &policyIfAuthors.PolicyIfAuthor[len(policyIfAuthors.PolicyIfAuthor)-1]
    }
    return nil
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyIfAuthors.PolicyIfAuthor {
        children[policyIfAuthors.PolicyIfAuthor[i].GetSegmentPath()] = &policyIfAuthors.PolicyIfAuthor[i]
    }
    return children
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetBundleName() string { return "cisco_ios_xr" }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetYangName() string { return "policy-if-authors" }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) SetParent(parent types.Entity) { policyIfAuthors.parent = parent }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetParent() types.Entity { return policyIfAuthors.parent }

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor
// Configurations related to authorization
type Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetFilter() yfilter.YFilter { return policyIfAuthor.YFilter }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) SetFilter(yf yfilter.YFilter) { policyIfAuthor.YFilter = yf }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetSegmentPath() string {
    return "policy-if-author" + "[type='" + fmt.Sprintf("%v", policyIfAuthor.Type) + "']" + "[listname='" + fmt.Sprintf("%v", policyIfAuthor.Listname) + "']"
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = policyIfAuthor.Type
    leafs["listname"] = policyIfAuthor.Listname
    leafs["method"] = policyIfAuthor.Method
    leafs["server-group-name"] = policyIfAuthor.ServerGroupName
    return leafs
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetBundleName() string { return "cisco_ios_xr" }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetYangName() string { return "policy-if-author" }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) SetParent(parent types.Entity) { policyIfAuthor.parent = parent }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetParent() types.Entity { return policyIfAuthor.parent }

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetParentYangName() string { return "policy-if-authors" }

// Aaa_AaaSubscriber_Accountings
// AAA accounting
type Aaa_AaaSubscriber_Accountings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_AaaSubscriber_Accountings_Accounting.
    Accounting []Aaa_AaaSubscriber_Accountings_Accounting
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetFilter() yfilter.YFilter { return accountings.YFilter }

func (accountings *Aaa_AaaSubscriber_Accountings) SetFilter(yf yfilter.YFilter) { accountings.YFilter = yf }

func (accountings *Aaa_AaaSubscriber_Accountings) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetSegmentPath() string {
    return "accountings"
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        for _, c := range accountings.Accounting {
            if accountings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaSubscriber_Accountings_Accounting{}
        accountings.Accounting = append(accountings.Accounting, child)
        return &accountings.Accounting[len(accountings.Accounting)-1]
    }
    return nil
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accountings.Accounting {
        children[accountings.Accounting[i].GetSegmentPath()] = &accountings.Accounting[i]
    }
    return children
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetBundleName() string { return "cisco_ios_xr" }

func (accountings *Aaa_AaaSubscriber_Accountings) GetYangName() string { return "accountings" }

func (accountings *Aaa_AaaSubscriber_Accountings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountings *Aaa_AaaSubscriber_Accountings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountings *Aaa_AaaSubscriber_Accountings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountings *Aaa_AaaSubscriber_Accountings) SetParent(parent types.Entity) { accountings.parent = parent }

func (accountings *Aaa_AaaSubscriber_Accountings) GetParent() types.Entity { return accountings.parent }

func (accountings *Aaa_AaaSubscriber_Accountings) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_Accountings_Accounting
// Configurations related to accounting
type Aaa_AaaSubscriber_Accountings_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set accounting lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Broadcast. The type is AaaAccountingBroadcast. This attribute is mandatory.
    Broadcast interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "broadcast" { return "Broadcast" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetSegmentPath() string {
    return "accounting" + "[type='" + fmt.Sprintf("%v", accounting.Type) + "']" + "[listname='" + fmt.Sprintf("%v", accounting.Listname) + "']"
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = accounting.Type
    leafs["listname"] = accounting.Listname
    leafs["broadcast"] = accounting.Broadcast
    leafs["method"] = accounting.Method
    leafs["server-group-name"] = accounting.ServerGroupName
    return leafs
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetYangName() string { return "accounting" }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetParentYangName() string { return "accountings" }

// Aaa_AaaSubscriber_ServiceAccounting
// Set accounting parameters for Service
type Aaa_AaaSubscriber_ServiceAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Send extended/brief service accounting records. The type is
    // AaaServiceAccounting.
    Type interface{}
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetFilter() yfilter.YFilter { return serviceAccounting.YFilter }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) SetFilter(yf yfilter.YFilter) { serviceAccounting.YFilter = yf }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    return ""
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetSegmentPath() string {
    return "service-accounting"
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = serviceAccounting.Type
    return leafs
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetYangName() string { return "service-accounting" }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) SetParent(parent types.Entity) { serviceAccounting.parent = parent }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetParent() types.Entity { return serviceAccounting.parent }

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_PrepaidAuthors
// AAA authorization prepaid
type Aaa_AaaSubscriber_PrepaidAuthors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor.
    PrepaidAuthor []Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetFilter() yfilter.YFilter { return prepaidAuthors.YFilter }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) SetFilter(yf yfilter.YFilter) { prepaidAuthors.YFilter = yf }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetGoName(yname string) string {
    if yname == "prepaid-author" { return "PrepaidAuthor" }
    return ""
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetSegmentPath() string {
    return "prepaid-authors"
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepaid-author" {
        for _, c := range prepaidAuthors.PrepaidAuthor {
            if prepaidAuthors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor{}
        prepaidAuthors.PrepaidAuthor = append(prepaidAuthors.PrepaidAuthor, child)
        return &prepaidAuthors.PrepaidAuthor[len(prepaidAuthors.PrepaidAuthor)-1]
    }
    return nil
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prepaidAuthors.PrepaidAuthor {
        children[prepaidAuthors.PrepaidAuthor[i].GetSegmentPath()] = &prepaidAuthors.PrepaidAuthor[i]
    }
    return children
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetBundleName() string { return "cisco_ios_xr" }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetYangName() string { return "prepaid-authors" }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) SetParent(parent types.Entity) { prepaidAuthors.parent = parent }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetParent() types.Entity { return prepaidAuthors.parent }

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor
// Configurations related to authorization
type Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetFilter() yfilter.YFilter { return prepaidAuthor.YFilter }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) SetFilter(yf yfilter.YFilter) { prepaidAuthor.YFilter = yf }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetSegmentPath() string {
    return "prepaid-author" + "[type='" + fmt.Sprintf("%v", prepaidAuthor.Type) + "']" + "[listname='" + fmt.Sprintf("%v", prepaidAuthor.Listname) + "']"
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = prepaidAuthor.Type
    leafs["listname"] = prepaidAuthor.Listname
    leafs["method"] = prepaidAuthor.Method
    leafs["server-group-name"] = prepaidAuthor.ServerGroupName
    return leafs
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetBundleName() string { return "cisco_ios_xr" }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetYangName() string { return "prepaid-author" }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) SetParent(parent types.Entity) { prepaidAuthor.parent = parent }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetParent() types.Entity { return prepaidAuthor.parent }

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetParentYangName() string { return "prepaid-authors" }

// Aaa_AaaSubscriber_Authorizations
// AAA authorization
type Aaa_AaaSubscriber_Authorizations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_Authorizations_Authorization.
    Authorization []Aaa_AaaSubscriber_Authorizations_Authorization
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetFilter() yfilter.YFilter { return authorizations.YFilter }

func (authorizations *Aaa_AaaSubscriber_Authorizations) SetFilter(yf yfilter.YFilter) { authorizations.YFilter = yf }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetGoName(yname string) string {
    if yname == "authorization" { return "Authorization" }
    return ""
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetSegmentPath() string {
    return "authorizations"
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authorization" {
        for _, c := range authorizations.Authorization {
            if authorizations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaSubscriber_Authorizations_Authorization{}
        authorizations.Authorization = append(authorizations.Authorization, child)
        return &authorizations.Authorization[len(authorizations.Authorization)-1]
    }
    return nil
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authorizations.Authorization {
        children[authorizations.Authorization[i].GetSegmentPath()] = &authorizations.Authorization[i]
    }
    return children
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetBundleName() string { return "cisco_ios_xr" }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetYangName() string { return "authorizations" }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorizations *Aaa_AaaSubscriber_Authorizations) SetParent(parent types.Entity) { authorizations.parent = parent }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetParent() types.Entity { return authorizations.parent }

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_Authorizations_Authorization
// Configurations related to authorization
type Aaa_AaaSubscriber_Authorizations_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetSegmentPath() string {
    return "authorization" + "[type='" + fmt.Sprintf("%v", authorization.Type) + "']" + "[listname='" + fmt.Sprintf("%v", authorization.Listname) + "']"
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = authorization.Type
    leafs["listname"] = authorization.Listname
    leafs["method"] = authorization.Method
    leafs["server-group-name"] = authorization.ServerGroupName
    return leafs
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetYangName() string { return "authorization" }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetParentYangName() string { return "authorizations" }

// Aaa_AaaSubscriber_Authentications
// AAA authentication
type Aaa_AaaSubscriber_Authentications struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_AaaSubscriber_Authentications_Authentication.
    Authentication []Aaa_AaaSubscriber_Authentications_Authentication
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetFilter() yfilter.YFilter { return authentications.YFilter }

func (authentications *Aaa_AaaSubscriber_Authentications) SetFilter(yf yfilter.YFilter) { authentications.YFilter = yf }

func (authentications *Aaa_AaaSubscriber_Authentications) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetSegmentPath() string {
    return "authentications"
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        for _, c := range authentications.Authentication {
            if authentications.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaSubscriber_Authentications_Authentication{}
        authentications.Authentication = append(authentications.Authentication, child)
        return &authentications.Authentication[len(authentications.Authentication)-1]
    }
    return nil
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authentications.Authentication {
        children[authentications.Authentication[i].GetSegmentPath()] = &authentications.Authentication[i]
    }
    return children
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetBundleName() string { return "cisco_ios_xr" }

func (authentications *Aaa_AaaSubscriber_Authentications) GetYangName() string { return "authentications" }

func (authentications *Aaa_AaaSubscriber_Authentications) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentications *Aaa_AaaSubscriber_Authentications) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentications *Aaa_AaaSubscriber_Authentications) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentications *Aaa_AaaSubscriber_Authentications) SetParent(parent types.Entity) { authentications.parent = parent }

func (authentications *Aaa_AaaSubscriber_Authentications) GetParent() types.Entity { return authentications.parent }

func (authentications *Aaa_AaaSubscriber_Authentications) GetParentYangName() string { return "aaa-subscriber" }

// Aaa_AaaSubscriber_Authentications_Authentication
// Configurations related to authentication
type Aaa_AaaSubscriber_Authentications_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set authentication lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authentication list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetSegmentPath() string {
    return "authentication" + "[type='" + fmt.Sprintf("%v", authentication.Type) + "']" + "[listname='" + fmt.Sprintf("%v", authentication.Listname) + "']"
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = authentication.Type
    leafs["listname"] = authentication.Listname
    leafs["method"] = authentication.Method
    leafs["server-group-name"] = authentication.ServerGroupName
    return leafs
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetYangName() string { return "authentication" }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetParentYangName() string { return "authentications" }

// Aaa_AaaMobile
// AAA Mobile
type Aaa_AaaMobile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA Mobile Accounting.
    Mobiles Aaa_AaaMobile_Mobiles
}

func (aaaMobile *Aaa_AaaMobile) GetFilter() yfilter.YFilter { return aaaMobile.YFilter }

func (aaaMobile *Aaa_AaaMobile) SetFilter(yf yfilter.YFilter) { aaaMobile.YFilter = yf }

func (aaaMobile *Aaa_AaaMobile) GetGoName(yname string) string {
    if yname == "mobiles" { return "Mobiles" }
    return ""
}

func (aaaMobile *Aaa_AaaMobile) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile"
}

func (aaaMobile *Aaa_AaaMobile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mobiles" {
        return &aaaMobile.Mobiles
    }
    return nil
}

func (aaaMobile *Aaa_AaaMobile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mobiles"] = &aaaMobile.Mobiles
    return children
}

func (aaaMobile *Aaa_AaaMobile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaaMobile *Aaa_AaaMobile) GetBundleName() string { return "cisco_ios_xr" }

func (aaaMobile *Aaa_AaaMobile) GetYangName() string { return "aaa-mobile" }

func (aaaMobile *Aaa_AaaMobile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaaMobile *Aaa_AaaMobile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaaMobile *Aaa_AaaMobile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaaMobile *Aaa_AaaMobile) SetParent(parent types.Entity) { aaaMobile.parent = parent }

func (aaaMobile *Aaa_AaaMobile) GetParent() types.Entity { return aaaMobile.parent }

func (aaaMobile *Aaa_AaaMobile) GetParentYangName() string { return "aaa" }

// Aaa_AaaMobile_Mobiles
// AAA Mobile Accounting
type Aaa_AaaMobile_Mobiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_AaaMobile_Mobiles_Mobile.
    Mobile []Aaa_AaaMobile_Mobiles_Mobile
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetFilter() yfilter.YFilter { return mobiles.YFilter }

func (mobiles *Aaa_AaaMobile_Mobiles) SetFilter(yf yfilter.YFilter) { mobiles.YFilter = yf }

func (mobiles *Aaa_AaaMobile_Mobiles) GetGoName(yname string) string {
    if yname == "mobile" { return "Mobile" }
    return ""
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetSegmentPath() string {
    return "mobiles"
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mobile" {
        for _, c := range mobiles.Mobile {
            if mobiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaMobile_Mobiles_Mobile{}
        mobiles.Mobile = append(mobiles.Mobile, child)
        return &mobiles.Mobile[len(mobiles.Mobile)-1]
    }
    return nil
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mobiles.Mobile {
        children[mobiles.Mobile[i].GetSegmentPath()] = &mobiles.Mobile[i]
    }
    return children
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetBundleName() string { return "cisco_ios_xr" }

func (mobiles *Aaa_AaaMobile_Mobiles) GetYangName() string { return "mobiles" }

func (mobiles *Aaa_AaaMobile_Mobiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobiles *Aaa_AaaMobile_Mobiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobiles *Aaa_AaaMobile_Mobiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobiles *Aaa_AaaMobile_Mobiles) SetParent(parent types.Entity) { mobiles.parent = parent }

func (mobiles *Aaa_AaaMobile_Mobiles) GetParent() types.Entity { return mobiles.parent }

func (mobiles *Aaa_AaaMobile_Mobiles) GetParentYangName() string { return "aaa-mobile" }

// Aaa_AaaMobile_Mobiles_Mobile
// Configurations related to accounting
type Aaa_AaaMobile_Mobiles_Mobile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Broadcast. The type is AaaAccountingBroadcast. This attribute is mandatory.
    Broadcast interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetFilter() yfilter.YFilter { return mobile.YFilter }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) SetFilter(yf yfilter.YFilter) { mobile.YFilter = yf }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetGoName(yname string) string {
    if yname == "listname" { return "Listname" }
    if yname == "broadcast" { return "Broadcast" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetSegmentPath() string {
    return "mobile" + "[listname='" + fmt.Sprintf("%v", mobile.Listname) + "']"
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["listname"] = mobile.Listname
    leafs["broadcast"] = mobile.Broadcast
    leafs["method"] = mobile.Method
    leafs["server-group-name"] = mobile.ServerGroupName
    return leafs
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetBundleName() string { return "cisco_ios_xr" }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetYangName() string { return "mobile" }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) SetParent(parent types.Entity) { mobile.parent = parent }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetParent() types.Entity { return mobile.parent }

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetParentYangName() string { return "mobiles" }

// Aaa_AaaDot1X
// AAA Dot1x
type Aaa_AaaDot1X struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA authentication.
    Authentications Aaa_AaaDot1X_Authentications
}

func (aaaDot1X *Aaa_AaaDot1X) GetFilter() yfilter.YFilter { return aaaDot1X.YFilter }

func (aaaDot1X *Aaa_AaaDot1X) SetFilter(yf yfilter.YFilter) { aaaDot1X.YFilter = yf }

func (aaaDot1X *Aaa_AaaDot1X) GetGoName(yname string) string {
    if yname == "authentications" { return "Authentications" }
    return ""
}

func (aaaDot1X *Aaa_AaaDot1X) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x"
}

func (aaaDot1X *Aaa_AaaDot1X) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentications" {
        return &aaaDot1X.Authentications
    }
    return nil
}

func (aaaDot1X *Aaa_AaaDot1X) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentications"] = &aaaDot1X.Authentications
    return children
}

func (aaaDot1X *Aaa_AaaDot1X) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaaDot1X *Aaa_AaaDot1X) GetBundleName() string { return "cisco_ios_xr" }

func (aaaDot1X *Aaa_AaaDot1X) GetYangName() string { return "aaa-dot1x" }

func (aaaDot1X *Aaa_AaaDot1X) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaaDot1X *Aaa_AaaDot1X) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaaDot1X *Aaa_AaaDot1X) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaaDot1X *Aaa_AaaDot1X) SetParent(parent types.Entity) { aaaDot1X.parent = parent }

func (aaaDot1X *Aaa_AaaDot1X) GetParent() types.Entity { return aaaDot1X.parent }

func (aaaDot1X *Aaa_AaaDot1X) GetParentYangName() string { return "aaa" }

// Aaa_AaaDot1X_Authentications
// AAA authentication
type Aaa_AaaDot1X_Authentications struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_AaaDot1X_Authentications_Authentication.
    Authentication []Aaa_AaaDot1X_Authentications_Authentication
}

func (authentications *Aaa_AaaDot1X_Authentications) GetFilter() yfilter.YFilter { return authentications.YFilter }

func (authentications *Aaa_AaaDot1X_Authentications) SetFilter(yf yfilter.YFilter) { authentications.YFilter = yf }

func (authentications *Aaa_AaaDot1X_Authentications) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (authentications *Aaa_AaaDot1X_Authentications) GetSegmentPath() string {
    return "authentications"
}

func (authentications *Aaa_AaaDot1X_Authentications) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        for _, c := range authentications.Authentication {
            if authentications.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_AaaDot1X_Authentications_Authentication{}
        authentications.Authentication = append(authentications.Authentication, child)
        return &authentications.Authentication[len(authentications.Authentication)-1]
    }
    return nil
}

func (authentications *Aaa_AaaDot1X_Authentications) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authentications.Authentication {
        children[authentications.Authentication[i].GetSegmentPath()] = &authentications.Authentication[i]
    }
    return children
}

func (authentications *Aaa_AaaDot1X_Authentications) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentications *Aaa_AaaDot1X_Authentications) GetBundleName() string { return "cisco_ios_xr" }

func (authentications *Aaa_AaaDot1X_Authentications) GetYangName() string { return "authentications" }

func (authentications *Aaa_AaaDot1X_Authentications) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentications *Aaa_AaaDot1X_Authentications) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentications *Aaa_AaaDot1X_Authentications) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentications *Aaa_AaaDot1X_Authentications) SetParent(parent types.Entity) { authentications.parent = parent }

func (authentications *Aaa_AaaDot1X_Authentications) GetParent() types.Entity { return authentications.parent }

func (authentications *Aaa_AaaDot1X_Authentications) GetParentYangName() string { return "aaa-dot1x" }

// Aaa_AaaDot1X_Authentications_Authentication
// Configurations related to authentication
type Aaa_AaaDot1X_Authentications_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set authentication lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authentication list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "listname" { return "Listname" }
    if yname == "method" { return "Method" }
    if yname == "server-group-name" { return "ServerGroupName" }
    return ""
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetSegmentPath() string {
    return "authentication" + "[type='" + fmt.Sprintf("%v", authentication.Type) + "']" + "[listname='" + fmt.Sprintf("%v", authentication.Listname) + "']"
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = authentication.Type
    leafs["listname"] = authentication.Listname
    leafs["method"] = authentication.Method
    leafs["server-group-name"] = authentication.ServerGroupName
    return leafs
}

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetYangName() string { return "authentication" }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Aaa_AaaDot1X_Authentications_Authentication) GetParentYangName() string { return "authentications" }

// Aaa_RadiusAttribute
// AAA RADIUS attribute configurations
type Aaa_RadiusAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute.
    NasPortId Aaa_RadiusAttribute_NasPortId

    // AAA calling station id attribute.
    CallingStation Aaa_RadiusAttribute_CallingStation

    // AAA called station id attribute.
    CalledStation Aaa_RadiusAttribute_CalledStation

    // AAA nas-port-id attribute.
    NasPort Aaa_RadiusAttribute_NasPort

    // AAA nas-port-id attribute format.
    FormatOthers Aaa_RadiusAttribute_FormatOthers
}

func (radiusAttribute *Aaa_RadiusAttribute) GetFilter() yfilter.YFilter { return radiusAttribute.YFilter }

func (radiusAttribute *Aaa_RadiusAttribute) SetFilter(yf yfilter.YFilter) { radiusAttribute.YFilter = yf }

func (radiusAttribute *Aaa_RadiusAttribute) GetGoName(yname string) string {
    if yname == "nas-port-id" { return "NasPortId" }
    if yname == "calling-station" { return "CallingStation" }
    if yname == "called-station" { return "CalledStation" }
    if yname == "nas-port" { return "NasPort" }
    if yname == "format-others" { return "FormatOthers" }
    return ""
}

func (radiusAttribute *Aaa_RadiusAttribute) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute"
}

func (radiusAttribute *Aaa_RadiusAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nas-port-id" {
        return &radiusAttribute.NasPortId
    }
    if childYangName == "calling-station" {
        return &radiusAttribute.CallingStation
    }
    if childYangName == "called-station" {
        return &radiusAttribute.CalledStation
    }
    if childYangName == "nas-port" {
        return &radiusAttribute.NasPort
    }
    if childYangName == "format-others" {
        return &radiusAttribute.FormatOthers
    }
    return nil
}

func (radiusAttribute *Aaa_RadiusAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nas-port-id"] = &radiusAttribute.NasPortId
    children["calling-station"] = &radiusAttribute.CallingStation
    children["called-station"] = &radiusAttribute.CalledStation
    children["nas-port"] = &radiusAttribute.NasPort
    children["format-others"] = &radiusAttribute.FormatOthers
    return children
}

func (radiusAttribute *Aaa_RadiusAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radiusAttribute *Aaa_RadiusAttribute) GetBundleName() string { return "cisco_ios_xr" }

func (radiusAttribute *Aaa_RadiusAttribute) GetYangName() string { return "radius-attribute" }

func (radiusAttribute *Aaa_RadiusAttribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radiusAttribute *Aaa_RadiusAttribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radiusAttribute *Aaa_RadiusAttribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radiusAttribute *Aaa_RadiusAttribute) SetParent(parent types.Entity) { radiusAttribute.parent = parent }

func (radiusAttribute *Aaa_RadiusAttribute) GetParent() types.Entity { return radiusAttribute.parent }

func (radiusAttribute *Aaa_RadiusAttribute) GetParentYangName() string { return "aaa" }

// Aaa_RadiusAttribute_NasPortId
// AAA nas-port-id attribute
type Aaa_RadiusAttribute_NasPortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_NasPortId_Formats
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetFilter() yfilter.YFilter { return nasPortId.YFilter }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) SetFilter(yf yfilter.YFilter) { nasPortId.YFilter = yf }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetGoName(yname string) string {
    if yname == "formats" { return "Formats" }
    return ""
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetSegmentPath() string {
    return "nas-port-id"
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "formats" {
        return &nasPortId.Formats
    }
    return nil
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["formats"] = &nasPortId.Formats
    return children
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetBundleName() string { return "cisco_ios_xr" }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetYangName() string { return "nas-port-id" }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) SetParent(parent types.Entity) { nasPortId.parent = parent }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetParent() types.Entity { return nasPortId.parent }

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetParentYangName() string { return "radius-attribute" }

// Aaa_RadiusAttribute_NasPortId_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_NasPortId_Formats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_NasPortId_Formats_Format.
    Format []Aaa_RadiusAttribute_NasPortId_Formats_Format
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetFilter() yfilter.YFilter { return formats.YFilter }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) SetFilter(yf yfilter.YFilter) { formats.YFilter = yf }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetGoName(yname string) string {
    if yname == "format" { return "Format" }
    return ""
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetSegmentPath() string {
    return "formats"
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format" {
        for _, c := range formats.Format {
            if formats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_RadiusAttribute_NasPortId_Formats_Format{}
        formats.Format = append(formats.Format, child)
        return &formats.Format[len(formats.Format)-1]
    }
    return nil
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range formats.Format {
        children[formats.Format[i].GetSegmentPath()] = &formats.Format[i]
    }
    return children
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetBundleName() string { return "cisco_ios_xr" }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetYangName() string { return "formats" }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) SetParent(parent types.Entity) { formats.parent = parent }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetParent() types.Entity { return formats.parent }

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetParentYangName() string { return "nas-port-id" }

// Aaa_RadiusAttribute_NasPortId_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_NasPortId_Formats_Format struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetFilter() yfilter.YFilter { return format.YFilter }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) SetFilter(yf yfilter.YFilter) { format.YFilter = yf }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "format-name" { return "FormatName" }
    return ""
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetSegmentPath() string {
    return "format" + "[type='" + fmt.Sprintf("%v", format.Type) + "']"
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = format.Type
    leafs["format-name"] = format.FormatName
    return leafs
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetBundleName() string { return "cisco_ios_xr" }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetYangName() string { return "format" }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) SetParent(parent types.Entity) { format.parent = parent }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetParent() types.Entity { return format.parent }

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetParentYangName() string { return "formats" }

// Aaa_RadiusAttribute_CallingStation
// AAA calling station id attribute
type Aaa_RadiusAttribute_CallingStation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_CallingStation_Formats
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetFilter() yfilter.YFilter { return callingStation.YFilter }

func (callingStation *Aaa_RadiusAttribute_CallingStation) SetFilter(yf yfilter.YFilter) { callingStation.YFilter = yf }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetGoName(yname string) string {
    if yname == "formats" { return "Formats" }
    return ""
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetSegmentPath() string {
    return "calling-station"
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "formats" {
        return &callingStation.Formats
    }
    return nil
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["formats"] = &callingStation.Formats
    return children
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetBundleName() string { return "cisco_ios_xr" }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetYangName() string { return "calling-station" }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (callingStation *Aaa_RadiusAttribute_CallingStation) SetParent(parent types.Entity) { callingStation.parent = parent }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetParent() types.Entity { return callingStation.parent }

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetParentYangName() string { return "radius-attribute" }

// Aaa_RadiusAttribute_CallingStation_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_CallingStation_Formats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_CallingStation_Formats_Format.
    Format []Aaa_RadiusAttribute_CallingStation_Formats_Format
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetFilter() yfilter.YFilter { return formats.YFilter }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) SetFilter(yf yfilter.YFilter) { formats.YFilter = yf }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetGoName(yname string) string {
    if yname == "format" { return "Format" }
    return ""
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetSegmentPath() string {
    return "formats"
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format" {
        for _, c := range formats.Format {
            if formats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_RadiusAttribute_CallingStation_Formats_Format{}
        formats.Format = append(formats.Format, child)
        return &formats.Format[len(formats.Format)-1]
    }
    return nil
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range formats.Format {
        children[formats.Format[i].GetSegmentPath()] = &formats.Format[i]
    }
    return children
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetBundleName() string { return "cisco_ios_xr" }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetYangName() string { return "formats" }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) SetParent(parent types.Entity) { formats.parent = parent }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetParent() types.Entity { return formats.parent }

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetParentYangName() string { return "calling-station" }

// Aaa_RadiusAttribute_CallingStation_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_CallingStation_Formats_Format struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetFilter() yfilter.YFilter { return format.YFilter }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) SetFilter(yf yfilter.YFilter) { format.YFilter = yf }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "format-name" { return "FormatName" }
    return ""
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetSegmentPath() string {
    return "format" + "[type='" + fmt.Sprintf("%v", format.Type) + "']"
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = format.Type
    leafs["format-name"] = format.FormatName
    return leafs
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetBundleName() string { return "cisco_ios_xr" }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetYangName() string { return "format" }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) SetParent(parent types.Entity) { format.parent = parent }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetParent() types.Entity { return format.parent }

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetParentYangName() string { return "formats" }

// Aaa_RadiusAttribute_CalledStation
// AAA called station id attribute
type Aaa_RadiusAttribute_CalledStation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_CalledStation_Formats
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetFilter() yfilter.YFilter { return calledStation.YFilter }

func (calledStation *Aaa_RadiusAttribute_CalledStation) SetFilter(yf yfilter.YFilter) { calledStation.YFilter = yf }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetGoName(yname string) string {
    if yname == "formats" { return "Formats" }
    return ""
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetSegmentPath() string {
    return "called-station"
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "formats" {
        return &calledStation.Formats
    }
    return nil
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["formats"] = &calledStation.Formats
    return children
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetBundleName() string { return "cisco_ios_xr" }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetYangName() string { return "called-station" }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (calledStation *Aaa_RadiusAttribute_CalledStation) SetParent(parent types.Entity) { calledStation.parent = parent }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetParent() types.Entity { return calledStation.parent }

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetParentYangName() string { return "radius-attribute" }

// Aaa_RadiusAttribute_CalledStation_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_CalledStation_Formats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_CalledStation_Formats_Format.
    Format []Aaa_RadiusAttribute_CalledStation_Formats_Format
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetFilter() yfilter.YFilter { return formats.YFilter }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) SetFilter(yf yfilter.YFilter) { formats.YFilter = yf }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetGoName(yname string) string {
    if yname == "format" { return "Format" }
    return ""
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetSegmentPath() string {
    return "formats"
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format" {
        for _, c := range formats.Format {
            if formats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_RadiusAttribute_CalledStation_Formats_Format{}
        formats.Format = append(formats.Format, child)
        return &formats.Format[len(formats.Format)-1]
    }
    return nil
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range formats.Format {
        children[formats.Format[i].GetSegmentPath()] = &formats.Format[i]
    }
    return children
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetBundleName() string { return "cisco_ios_xr" }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetYangName() string { return "formats" }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) SetParent(parent types.Entity) { formats.parent = parent }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetParent() types.Entity { return formats.parent }

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetParentYangName() string { return "called-station" }

// Aaa_RadiusAttribute_CalledStation_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_CalledStation_Formats_Format struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetFilter() yfilter.YFilter { return format.YFilter }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) SetFilter(yf yfilter.YFilter) { format.YFilter = yf }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "format-name" { return "FormatName" }
    return ""
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetSegmentPath() string {
    return "format" + "[type='" + fmt.Sprintf("%v", format.Type) + "']"
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = format.Type
    leafs["format-name"] = format.FormatName
    return leafs
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetBundleName() string { return "cisco_ios_xr" }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetYangName() string { return "format" }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) SetParent(parent types.Entity) { format.parent = parent }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetParent() types.Entity { return format.parent }

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetParentYangName() string { return "formats" }

// Aaa_RadiusAttribute_NasPort
// AAA nas-port-id attribute
type Aaa_RadiusAttribute_NasPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    FormatExtendeds Aaa_RadiusAttribute_NasPort_FormatExtendeds
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetFilter() yfilter.YFilter { return nasPort.YFilter }

func (nasPort *Aaa_RadiusAttribute_NasPort) SetFilter(yf yfilter.YFilter) { nasPort.YFilter = yf }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetGoName(yname string) string {
    if yname == "format-extendeds" { return "FormatExtendeds" }
    return ""
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetSegmentPath() string {
    return "nas-port"
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format-extendeds" {
        return &nasPort.FormatExtendeds
    }
    return nil
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["format-extendeds"] = &nasPort.FormatExtendeds
    return children
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetBundleName() string { return "cisco_ios_xr" }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetYangName() string { return "nas-port" }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nasPort *Aaa_RadiusAttribute_NasPort) SetParent(parent types.Entity) { nasPort.parent = parent }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetParent() types.Entity { return nasPort.parent }

func (nasPort *Aaa_RadiusAttribute_NasPort) GetParentYangName() string { return "radius-attribute" }

// Aaa_RadiusAttribute_NasPort_FormatExtendeds
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_NasPort_FormatExtendeds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nas-port-id extended attribute. The type is slice of
    // Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended.
    FormatExtended []Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetFilter() yfilter.YFilter { return formatExtendeds.YFilter }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) SetFilter(yf yfilter.YFilter) { formatExtendeds.YFilter = yf }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetGoName(yname string) string {
    if yname == "format-extended" { return "FormatExtended" }
    return ""
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetSegmentPath() string {
    return "format-extendeds"
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format-extended" {
        for _, c := range formatExtendeds.FormatExtended {
            if formatExtendeds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended{}
        formatExtendeds.FormatExtended = append(formatExtendeds.FormatExtended, child)
        return &formatExtendeds.FormatExtended[len(formatExtendeds.FormatExtended)-1]
    }
    return nil
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range formatExtendeds.FormatExtended {
        children[formatExtendeds.FormatExtended[i].GetSegmentPath()] = &formatExtendeds.FormatExtended[i]
    }
    return children
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetBundleName() string { return "cisco_ios_xr" }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetYangName() string { return "format-extendeds" }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) SetParent(parent types.Entity) { formatExtendeds.parent = parent }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetParent() types.Entity { return formatExtendeds.parent }

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetParentYangName() string { return "nas-port" }

// Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended
// nas-port-id extended attribute
type Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. format type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Value interface{}

    // This attribute is a key. AAA nas-port attribute format. The type is
    // interface{} with range: 0..45.
    Type interface{}

    // A 32 character string representing the format to be used. The type is
    // string with length: 1..32.
    FormatIdentifier interface{}
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetFilter() yfilter.YFilter { return formatExtended.YFilter }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) SetFilter(yf yfilter.YFilter) { formatExtended.YFilter = yf }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "type" { return "Type" }
    if yname == "format-identifier" { return "FormatIdentifier" }
    return ""
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetSegmentPath() string {
    return "format-extended" + "[value='" + fmt.Sprintf("%v", formatExtended.Value) + "']" + "[type='" + fmt.Sprintf("%v", formatExtended.Type) + "']"
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = formatExtended.Value
    leafs["type"] = formatExtended.Type
    leafs["format-identifier"] = formatExtended.FormatIdentifier
    return leafs
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetBundleName() string { return "cisco_ios_xr" }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetYangName() string { return "format-extended" }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) SetParent(parent types.Entity) { formatExtended.parent = parent }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetParent() types.Entity { return formatExtended.parent }

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetParentYangName() string { return "format-extendeds" }

// Aaa_RadiusAttribute_FormatOthers
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_FormatOthers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Other configs. The type is slice of
    // Aaa_RadiusAttribute_FormatOthers_FormatOther.
    FormatOther []Aaa_RadiusAttribute_FormatOthers_FormatOther
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetFilter() yfilter.YFilter { return formatOthers.YFilter }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) SetFilter(yf yfilter.YFilter) { formatOthers.YFilter = yf }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetGoName(yname string) string {
    if yname == "format-other" { return "FormatOther" }
    return ""
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetSegmentPath() string {
    return "format-others"
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "format-other" {
        for _, c := range formatOthers.FormatOther {
            if formatOthers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_RadiusAttribute_FormatOthers_FormatOther{}
        formatOthers.FormatOther = append(formatOthers.FormatOther, child)
        return &formatOthers.FormatOther[len(formatOthers.FormatOther)-1]
    }
    return nil
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range formatOthers.FormatOther {
        children[formatOthers.FormatOther[i].GetSegmentPath()] = &formatOthers.FormatOther[i]
    }
    return children
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetBundleName() string { return "cisco_ios_xr" }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetYangName() string { return "format-others" }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) SetParent(parent types.Entity) { formatOthers.parent = parent }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetParent() types.Entity { return formatOthers.parent }

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetParentYangName() string { return "radius-attribute" }

// Aaa_RadiusAttribute_FormatOthers_FormatOther
// Other configs
type Aaa_RadiusAttribute_FormatOthers_FormatOther struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    NasPortTypeName interface{}

    // Argument1. The type is string.
    AttributeConfig1 interface{}

    // Argument2. The type is string.
    AttributeConfig2 interface{}

    // Argument3. The type is string.
    AttributeConfig3 interface{}

    // Argument4. The type is string.
    AttributeConfig4 interface{}

    // Argument5. The type is string.
    AttributeConfig5 interface{}

    // Argument6. The type is string.
    AttributeConfig6 interface{}

    // Argument7. The type is string.
    AttributeConfig7 interface{}

    // Argument8. The type is string.
    AttributeConfig8 interface{}

    // Argument9. The type is string.
    AttributeConfig9 interface{}

    // Argument10. The type is string.
    AttributeConfig10 interface{}

    // Argument11. The type is string.
    AttributeConfig11 interface{}

    // Argument12. The type is string.
    AttributeConfig12 interface{}

    // Argument13. The type is string.
    AttributeConfig13 interface{}

    // Argument14. The type is string.
    AttributeConfig14 interface{}

    // Argument15. The type is string.
    AttributeConfig15 interface{}

    // Argument16. The type is string.
    AttributeConfig16 interface{}

    // Argument17. The type is string.
    AttributeConfig17 interface{}

    // Argument18. The type is string.
    AttributeConfig18 interface{}

    // Argument19. The type is interface{} with range: 1..253.
    AttributeConfig19 interface{}
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetFilter() yfilter.YFilter { return formatOther.YFilter }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) SetFilter(yf yfilter.YFilter) { formatOther.YFilter = yf }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetGoName(yname string) string {
    if yname == "nas-port-type-name" { return "NasPortTypeName" }
    if yname == "attribute-config1" { return "AttributeConfig1" }
    if yname == "attribute-config2" { return "AttributeConfig2" }
    if yname == "attribute-config3" { return "AttributeConfig3" }
    if yname == "attribute-config4" { return "AttributeConfig4" }
    if yname == "attribute-config5" { return "AttributeConfig5" }
    if yname == "attribute-config6" { return "AttributeConfig6" }
    if yname == "attribute-config7" { return "AttributeConfig7" }
    if yname == "attribute-config8" { return "AttributeConfig8" }
    if yname == "attribute-config9" { return "AttributeConfig9" }
    if yname == "attribute-config10" { return "AttributeConfig10" }
    if yname == "attribute-config11" { return "AttributeConfig11" }
    if yname == "attribute-config12" { return "AttributeConfig12" }
    if yname == "attribute-config13" { return "AttributeConfig13" }
    if yname == "attribute-config14" { return "AttributeConfig14" }
    if yname == "attribute-config15" { return "AttributeConfig15" }
    if yname == "attribute-config16" { return "AttributeConfig16" }
    if yname == "attribute-config17" { return "AttributeConfig17" }
    if yname == "attribute-config18" { return "AttributeConfig18" }
    if yname == "attribute-config19" { return "AttributeConfig19" }
    return ""
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetSegmentPath() string {
    return "format-other" + "[nas-port-type-name='" + fmt.Sprintf("%v", formatOther.NasPortTypeName) + "']"
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nas-port-type-name"] = formatOther.NasPortTypeName
    leafs["attribute-config1"] = formatOther.AttributeConfig1
    leafs["attribute-config2"] = formatOther.AttributeConfig2
    leafs["attribute-config3"] = formatOther.AttributeConfig3
    leafs["attribute-config4"] = formatOther.AttributeConfig4
    leafs["attribute-config5"] = formatOther.AttributeConfig5
    leafs["attribute-config6"] = formatOther.AttributeConfig6
    leafs["attribute-config7"] = formatOther.AttributeConfig7
    leafs["attribute-config8"] = formatOther.AttributeConfig8
    leafs["attribute-config9"] = formatOther.AttributeConfig9
    leafs["attribute-config10"] = formatOther.AttributeConfig10
    leafs["attribute-config11"] = formatOther.AttributeConfig11
    leafs["attribute-config12"] = formatOther.AttributeConfig12
    leafs["attribute-config13"] = formatOther.AttributeConfig13
    leafs["attribute-config14"] = formatOther.AttributeConfig14
    leafs["attribute-config15"] = formatOther.AttributeConfig15
    leafs["attribute-config16"] = formatOther.AttributeConfig16
    leafs["attribute-config17"] = formatOther.AttributeConfig17
    leafs["attribute-config18"] = formatOther.AttributeConfig18
    leafs["attribute-config19"] = formatOther.AttributeConfig19
    return leafs
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetBundleName() string { return "cisco_ios_xr" }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetYangName() string { return "format-other" }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) SetParent(parent types.Entity) { formatOther.parent = parent }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetParent() types.Entity { return formatOther.parent }

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetParentYangName() string { return "format-others" }

// Aaa_ServerGroups
// AAA group definitions
type Aaa_ServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DIAMETER server group definition.
    DiameterServerGroups Aaa_ServerGroups_DiameterServerGroups

    // RADIUS server group definition.
    RadiusServerGroups Aaa_ServerGroups_RadiusServerGroups

    // TACACS+ server-group definition.
    TacacsServerGroups Aaa_ServerGroups_TacacsServerGroups
}

func (serverGroups *Aaa_ServerGroups) GetFilter() yfilter.YFilter { return serverGroups.YFilter }

func (serverGroups *Aaa_ServerGroups) SetFilter(yf yfilter.YFilter) { serverGroups.YFilter = yf }

func (serverGroups *Aaa_ServerGroups) GetGoName(yname string) string {
    if yname == "Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups" { return "DiameterServerGroups" }
    if yname == "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups" { return "RadiusServerGroups" }
    if yname == "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups" { return "TacacsServerGroups" }
    return ""
}

func (serverGroups *Aaa_ServerGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-cfg:server-groups"
}

func (serverGroups *Aaa_ServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups" {
        return &serverGroups.DiameterServerGroups
    }
    if childYangName == "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups" {
        return &serverGroups.RadiusServerGroups
    }
    if childYangName == "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups" {
        return &serverGroups.TacacsServerGroups
    }
    return nil
}

func (serverGroups *Aaa_ServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups"] = &serverGroups.DiameterServerGroups
    children["Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups"] = &serverGroups.RadiusServerGroups
    children["Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups"] = &serverGroups.TacacsServerGroups
    return children
}

func (serverGroups *Aaa_ServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serverGroups *Aaa_ServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroups *Aaa_ServerGroups) GetYangName() string { return "server-groups" }

func (serverGroups *Aaa_ServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroups *Aaa_ServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroups *Aaa_ServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroups *Aaa_ServerGroups) SetParent(parent types.Entity) { serverGroups.parent = parent }

func (serverGroups *Aaa_ServerGroups) GetParent() types.Entity { return serverGroups.parent }

func (serverGroups *Aaa_ServerGroups) GetParentYangName() string { return "aaa" }

// Aaa_ServerGroups_DiameterServerGroups
// DIAMETER server group definition
type Aaa_ServerGroups_DiameterServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DIAMETER server group name. The type is slice of
    // Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup.
    DiameterServerGroup []Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetFilter() yfilter.YFilter { return diameterServerGroups.YFilter }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) SetFilter(yf yfilter.YFilter) { diameterServerGroups.YFilter = yf }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetGoName(yname string) string {
    if yname == "diameter-server-group" { return "DiameterServerGroup" }
    return ""
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups"
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diameter-server-group" {
        for _, c := range diameterServerGroups.DiameterServerGroup {
            if diameterServerGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup{}
        diameterServerGroups.DiameterServerGroup = append(diameterServerGroups.DiameterServerGroup, child)
        return &diameterServerGroups.DiameterServerGroup[len(diameterServerGroups.DiameterServerGroup)-1]
    }
    return nil
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diameterServerGroups.DiameterServerGroup {
        children[diameterServerGroups.DiameterServerGroup[i].GetSegmentPath()] = &diameterServerGroups.DiameterServerGroup[i]
    }
    return children
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetYangName() string { return "diameter-server-groups" }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) SetParent(parent types.Entity) { diameterServerGroups.parent = parent }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetParent() types.Entity { return diameterServerGroups.parent }

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetParentYangName() string { return "server-groups" }

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup
// DIAMETER server group name
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. DIAMETER server group name. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // List of DIAMETER servers present in the group.
    Servers Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetFilter() yfilter.YFilter { return diameterServerGroup.YFilter }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) SetFilter(yf yfilter.YFilter) { diameterServerGroup.YFilter = yf }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetGoName(yname string) string {
    if yname == "server-group-name" { return "ServerGroupName" }
    if yname == "servers" { return "Servers" }
    return ""
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetSegmentPath() string {
    return "diameter-server-group" + "[server-group-name='" + fmt.Sprintf("%v", diameterServerGroup.ServerGroupName) + "']"
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "servers" {
        return &diameterServerGroup.Servers
    }
    return nil
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["servers"] = &diameterServerGroup.Servers
    return children
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-group-name"] = diameterServerGroup.ServerGroupName
    return leafs
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetYangName() string { return "diameter-server-group" }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) SetParent(parent types.Entity) { diameterServerGroup.parent = parent }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetParent() types.Entity { return diameterServerGroup.parent }

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetParentYangName() string { return "diameter-server-groups" }

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers
// List of DIAMETER servers present in the group
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server.
    Server []Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetYangName() string { return "servers" }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetParent() types.Entity { return servers.parent }

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetParentYangName() string { return "diameter-server-group" }

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. Name for the diameter peer configuration. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PeerName interface{}
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "peer-name" { return "PeerName" }
    return ""
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetSegmentPath() string {
    return "server" + "[ordering-index='" + fmt.Sprintf("%v", server.OrderingIndex) + "']" + "[peer-name='" + fmt.Sprintf("%v", server.PeerName) + "']"
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = server.OrderingIndex
    leafs["peer-name"] = server.PeerName
    return leafs
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetYangName() string { return "server" }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetParentYangName() string { return "servers" }

// Aaa_ServerGroups_RadiusServerGroups
// RADIUS server group definition
type Aaa_ServerGroups_RadiusServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS server group name. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup.
    RadiusServerGroup []Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetFilter() yfilter.YFilter { return radiusServerGroups.YFilter }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) SetFilter(yf yfilter.YFilter) { radiusServerGroups.YFilter = yf }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetGoName(yname string) string {
    if yname == "radius-server-group" { return "RadiusServerGroup" }
    return ""
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups"
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "radius-server-group" {
        for _, c := range radiusServerGroups.RadiusServerGroup {
            if radiusServerGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup{}
        radiusServerGroups.RadiusServerGroup = append(radiusServerGroups.RadiusServerGroup, child)
        return &radiusServerGroups.RadiusServerGroup[len(radiusServerGroups.RadiusServerGroup)-1]
    }
    return nil
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range radiusServerGroups.RadiusServerGroup {
        children[radiusServerGroups.RadiusServerGroup[i].GetSegmentPath()] = &radiusServerGroups.RadiusServerGroup[i]
    }
    return children
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetYangName() string { return "radius-server-groups" }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) SetParent(parent types.Entity) { radiusServerGroups.parent = parent }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetParent() types.Entity { return radiusServerGroups.parent }

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetParentYangName() string { return "server-groups" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup
// RADIUS server group name
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RADIUS server group name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // This indicates the length of time (in minutes) for which RADIUS servers
    // present in this group remain marked as dead. The type is interface{} with
    // range: 1..1440. Units are minute.
    DeadTime interface{}

    // Specify interface for source address in RADIUS packets. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Specify VRF name of RADIUS group. The type is string.
    Vrf interface{}

    // List of filters in server group.
    Accounting Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting

    // List of RADIUS servers present in the group.
    Servers Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers

    // List of private RADIUS servers present in the group.
    PrivateServers Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers

    // Radius throttling options.
    ServerGroupThrottle Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle

    // Radius load-balancing options.
    LoadBalance Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance

    // List of filters in server group.
    Authorization Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetFilter() yfilter.YFilter { return radiusServerGroup.YFilter }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) SetFilter(yf yfilter.YFilter) { radiusServerGroup.YFilter = yf }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetGoName(yname string) string {
    if yname == "server-group-name" { return "ServerGroupName" }
    if yname == "dead-time" { return "DeadTime" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "vrf" { return "Vrf" }
    if yname == "accounting" { return "Accounting" }
    if yname == "servers" { return "Servers" }
    if yname == "private-servers" { return "PrivateServers" }
    if yname == "server-group-throttle" { return "ServerGroupThrottle" }
    if yname == "load-balance" { return "LoadBalance" }
    if yname == "authorization" { return "Authorization" }
    return ""
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetSegmentPath() string {
    return "radius-server-group" + "[server-group-name='" + fmt.Sprintf("%v", radiusServerGroup.ServerGroupName) + "']"
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &radiusServerGroup.Accounting
    }
    if childYangName == "servers" {
        return &radiusServerGroup.Servers
    }
    if childYangName == "private-servers" {
        return &radiusServerGroup.PrivateServers
    }
    if childYangName == "server-group-throttle" {
        return &radiusServerGroup.ServerGroupThrottle
    }
    if childYangName == "load-balance" {
        return &radiusServerGroup.LoadBalance
    }
    if childYangName == "authorization" {
        return &radiusServerGroup.Authorization
    }
    return nil
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &radiusServerGroup.Accounting
    children["servers"] = &radiusServerGroup.Servers
    children["private-servers"] = &radiusServerGroup.PrivateServers
    children["server-group-throttle"] = &radiusServerGroup.ServerGroupThrottle
    children["load-balance"] = &radiusServerGroup.LoadBalance
    children["authorization"] = &radiusServerGroup.Authorization
    return children
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-group-name"] = radiusServerGroup.ServerGroupName
    leafs["dead-time"] = radiusServerGroup.DeadTime
    leafs["source-interface"] = radiusServerGroup.SourceInterface
    leafs["vrf"] = radiusServerGroup.Vrf
    return leafs
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetYangName() string { return "radius-server-group" }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) SetParent(parent types.Entity) { radiusServerGroup.parent = parent }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetParent() types.Entity { return radiusServerGroup.parent }

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetParentYangName() string { return "radius-server-groups" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting
// List of filters in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a filter in server group.
    Request Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request

    // Specify a filter in server group.
    Reply Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetGoName(yname string) string {
    if yname == "request" { return "Request" }
    if yname == "reply" { return "Reply" }
    return ""
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request" {
        return &accounting.Request
    }
    if childYangName == "reply" {
        return &accounting.Reply
    }
    return nil
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["request"] = &accounting.Request
    children["reply"] = &accounting.Reply
    return children
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetYangName() string { return "accounting" }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "attribute-list-name" { return "AttributeListName" }
    return ""
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetSegmentPath() string {
    return "request"
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = request.Action
    leafs["attribute-list-name"] = request.AttributeListName
    return leafs
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetYangName() string { return "request" }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetParent() types.Entity { return request.parent }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetParentYangName() string { return "accounting" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "attribute-list-name" { return "AttributeListName" }
    return ""
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetSegmentPath() string {
    return "reply"
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = reply.Action
    leafs["attribute-list-name"] = reply.AttributeListName
    return leafs
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetYangName() string { return "reply" }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetParentYangName() string { return "accounting" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers
// List of RADIUS servers present in the group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server.
    Server []Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetYangName() string { return "servers" }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetParent() types.Entity { return servers.parent }

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auth-port-number" { return "AuthPortNumber" }
    if yname == "acct-port-number" { return "AcctPortNumber" }
    return ""
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetSegmentPath() string {
    return "server" + "[ordering-index='" + fmt.Sprintf("%v", server.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", server.IpAddress) + "']" + "[auth-port-number='" + fmt.Sprintf("%v", server.AuthPortNumber) + "']" + "[acct-port-number='" + fmt.Sprintf("%v", server.AcctPortNumber) + "']"
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = server.OrderingIndex
    leafs["ip-address"] = server.IpAddress
    leafs["auth-port-number"] = server.AuthPortNumber
    leafs["acct-port-number"] = server.AcctPortNumber
    return leafs
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetYangName() string { return "server" }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetParentYangName() string { return "servers" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers
// List of private RADIUS servers present in the
// group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A private server to include in the server group. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer.
    PrivateServer []Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetFilter() yfilter.YFilter { return privateServers.YFilter }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) SetFilter(yf yfilter.YFilter) { privateServers.YFilter = yf }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetGoName(yname string) string {
    if yname == "private-server" { return "PrivateServer" }
    return ""
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetSegmentPath() string {
    return "private-servers"
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "private-server" {
        for _, c := range privateServers.PrivateServer {
            if privateServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer{}
        privateServers.PrivateServer = append(privateServers.PrivateServer, child)
        return &privateServers.PrivateServer[len(privateServers.PrivateServer)-1]
    }
    return nil
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range privateServers.PrivateServer {
        children[privateServers.PrivateServer[i].GetSegmentPath()] = &privateServers.PrivateServer[i]
    }
    return children
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetBundleName() string { return "cisco_ios_xr" }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetYangName() string { return "private-servers" }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) SetParent(parent types.Entity) { privateServers.parent = parent }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetParent() types.Entity { return privateServers.parent }

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer
// A private server to include in the server
// group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    PrivateTimeout interface{}

    // Ignore Accounting port. The type is bool.
    IgnoreAccountingPort interface{}

    // Number of times to retransmit a request to the RADIUS server. The type is
    // interface{} with range: 1..100. The default value is 3.
    PrivateRetransmit interface{}

    // Idle time for the radius Server. The type is interface{} with range: 1..60.
    // The default value is 5.
    IdleTime interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    PrivateKey interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Ignore authentication Port. The type is bool.
    IgnoreAuthPort interface{}
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetFilter() yfilter.YFilter { return privateServer.YFilter }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) SetFilter(yf yfilter.YFilter) { privateServer.YFilter = yf }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auth-port-number" { return "AuthPortNumber" }
    if yname == "acct-port-number" { return "AcctPortNumber" }
    if yname == "private-timeout" { return "PrivateTimeout" }
    if yname == "ignore-accounting-port" { return "IgnoreAccountingPort" }
    if yname == "private-retransmit" { return "PrivateRetransmit" }
    if yname == "idle-time" { return "IdleTime" }
    if yname == "private-key" { return "PrivateKey" }
    if yname == "username" { return "Username" }
    if yname == "ignore-auth-port" { return "IgnoreAuthPort" }
    return ""
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetSegmentPath() string {
    return "private-server" + "[ordering-index='" + fmt.Sprintf("%v", privateServer.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", privateServer.IpAddress) + "']" + "[auth-port-number='" + fmt.Sprintf("%v", privateServer.AuthPortNumber) + "']" + "[acct-port-number='" + fmt.Sprintf("%v", privateServer.AcctPortNumber) + "']"
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = privateServer.OrderingIndex
    leafs["ip-address"] = privateServer.IpAddress
    leafs["auth-port-number"] = privateServer.AuthPortNumber
    leafs["acct-port-number"] = privateServer.AcctPortNumber
    leafs["private-timeout"] = privateServer.PrivateTimeout
    leafs["ignore-accounting-port"] = privateServer.IgnoreAccountingPort
    leafs["private-retransmit"] = privateServer.PrivateRetransmit
    leafs["idle-time"] = privateServer.IdleTime
    leafs["private-key"] = privateServer.PrivateKey
    leafs["username"] = privateServer.Username
    leafs["ignore-auth-port"] = privateServer.IgnoreAuthPort
    return leafs
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetBundleName() string { return "cisco_ios_xr" }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetYangName() string { return "private-server" }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) SetParent(parent types.Entity) { privateServer.parent = parent }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetParent() types.Entity { return privateServer.parent }

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetParentYangName() string { return "private-servers" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle
// Radius throttling options
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // To flow control the number of access requests sent to a radius server. The
    // type is interface{} with range: 0..65535. The default value is 0.
    Access interface{}

    // To flow control the number of accounting requests sent to a radius server.
    // The type is interface{} with range: 0..65535. The default value is 0.
    Accounting interface{}

    // Specify the number of timeouts exceeding which a throttled access request
    // is dropped. The type is interface{} with range: 1..10. The default value is
    // 3.
    AccessTimeout interface{}
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetFilter() yfilter.YFilter { return serverGroupThrottle.YFilter }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) SetFilter(yf yfilter.YFilter) { serverGroupThrottle.YFilter = yf }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    if yname == "accounting" { return "Accounting" }
    if yname == "access-timeout" { return "AccessTimeout" }
    return ""
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetSegmentPath() string {
    return "server-group-throttle"
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access"] = serverGroupThrottle.Access
    leafs["accounting"] = serverGroupThrottle.Accounting
    leafs["access-timeout"] = serverGroupThrottle.AccessTimeout
    return leafs
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetYangName() string { return "server-group-throttle" }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) SetParent(parent types.Entity) { serverGroupThrottle.parent = parent }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetParent() types.Entity { return serverGroupThrottle.parent }

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance
// Radius load-balancing options
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Method by which the next host will be picked.
    Method Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetFilter() yfilter.YFilter { return loadBalance.YFilter }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) SetFilter(yf yfilter.YFilter) { loadBalance.YFilter = yf }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    return ""
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetSegmentPath() string {
    return "load-balance"
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "method" {
        return &loadBalance.Method
    }
    return nil
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["method"] = &loadBalance.Method
    return children
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetBundleName() string { return "cisco_ios_xr" }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetYangName() string { return "load-balance" }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) SetParent(parent types.Entity) { loadBalance.parent = parent }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetParent() types.Entity { return loadBalance.parent }

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method
// Method by which the next host will be picked
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Batch size for selection of the server.
    Name Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetFilter() yfilter.YFilter { return method.YFilter }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) SetFilter(yf yfilter.YFilter) { method.YFilter = yf }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetSegmentPath() string {
    return "method"
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "name" {
        return &method.Name
    }
    return nil
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["name"] = &method.Name
    return children
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetBundleName() string { return "cisco_ios_xr" }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetYangName() string { return "method" }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) SetParent(parent types.Entity) { method.parent = parent }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetParent() types.Entity { return method.parent }

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetParentYangName() string { return "load-balance" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name
// Batch size for selection of the server
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pick the server with the least transactions outstanding. The type is
    // interface{} with range: -2147483648..2147483647. The default value is 4.
    LeastOutstanding interface{}

    // Batch size for selection of the server. The type is interface{} with range:
    // 1..1500. The default value is 25.
    BatchSize interface{}

    // Disable preferred server for this Server Group. The type is bool. The
    // default value is true.
    IgnorePreferredServer interface{}
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetFilter() yfilter.YFilter { return name.YFilter }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) SetFilter(yf yfilter.YFilter) { name.YFilter = yf }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetGoName(yname string) string {
    if yname == "least-outstanding" { return "LeastOutstanding" }
    if yname == "batch-size" { return "BatchSize" }
    if yname == "ignore-preferred-server" { return "IgnorePreferredServer" }
    return ""
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetSegmentPath() string {
    return "name"
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["least-outstanding"] = name.LeastOutstanding
    leafs["batch-size"] = name.BatchSize
    leafs["ignore-preferred-server"] = name.IgnorePreferredServer
    return leafs
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetBundleName() string { return "cisco_ios_xr" }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetYangName() string { return "name" }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) SetParent(parent types.Entity) { name.parent = parent }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetParent() types.Entity { return name.parent }

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetParentYangName() string { return "method" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization
// List of filters in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a filter in server group.
    Request Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request

    // Specify a filter in server group.
    Reply Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetGoName(yname string) string {
    if yname == "request" { return "Request" }
    if yname == "reply" { return "Reply" }
    return ""
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetSegmentPath() string {
    return "authorization"
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request" {
        return &authorization.Request
    }
    if childYangName == "reply" {
        return &authorization.Reply
    }
    return nil
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["request"] = &authorization.Request
    children["reply"] = &authorization.Reply
    return children
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetYangName() string { return "authorization" }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetParentYangName() string { return "radius-server-group" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "attribute-list-name" { return "AttributeListName" }
    return ""
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetSegmentPath() string {
    return "request"
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = request.Action
    leafs["attribute-list-name"] = request.AttributeListName
    return leafs
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetYangName() string { return "request" }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetParent() types.Entity { return request.parent }

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetParentYangName() string { return "authorization" }

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "attribute-list-name" { return "AttributeListName" }
    return ""
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetSegmentPath() string {
    return "reply"
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = reply.Action
    leafs["attribute-list-name"] = reply.AttributeListName
    return leafs
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetYangName() string { return "reply" }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetParentYangName() string { return "authorization" }

// Aaa_ServerGroups_TacacsServerGroups
// TACACS+ server-group definition
type Aaa_ServerGroups_TacacsServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TACACS+ Server group name. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup.
    TacacsServerGroup []Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetFilter() yfilter.YFilter { return tacacsServerGroups.YFilter }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) SetFilter(yf yfilter.YFilter) { tacacsServerGroups.YFilter = yf }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetGoName(yname string) string {
    if yname == "tacacs-server-group" { return "TacacsServerGroup" }
    return ""
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups"
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tacacs-server-group" {
        for _, c := range tacacsServerGroups.TacacsServerGroup {
            if tacacsServerGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup{}
        tacacsServerGroups.TacacsServerGroup = append(tacacsServerGroups.TacacsServerGroup, child)
        return &tacacsServerGroups.TacacsServerGroup[len(tacacsServerGroups.TacacsServerGroup)-1]
    }
    return nil
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tacacsServerGroups.TacacsServerGroup {
        children[tacacsServerGroups.TacacsServerGroup[i].GetSegmentPath()] = &tacacsServerGroups.TacacsServerGroup[i]
    }
    return children
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetYangName() string { return "tacacs-server-groups" }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) SetParent(parent types.Entity) { tacacsServerGroups.parent = parent }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetParent() types.Entity { return tacacsServerGroups.parent }

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetParentYangName() string { return "server-groups" }

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup
// TACACS+ Server group name
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TACACS+ Server group name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // Specify VRF name of TACACS group. The type is string.
    Vrf interface{}

    // Specify a TACACS+ server.
    Servers Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers

    // List of private TACACS servers present in the group.
    PrivateServers Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetFilter() yfilter.YFilter { return tacacsServerGroup.YFilter }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) SetFilter(yf yfilter.YFilter) { tacacsServerGroup.YFilter = yf }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetGoName(yname string) string {
    if yname == "server-group-name" { return "ServerGroupName" }
    if yname == "vrf" { return "Vrf" }
    if yname == "servers" { return "Servers" }
    if yname == "private-servers" { return "PrivateServers" }
    return ""
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetSegmentPath() string {
    return "tacacs-server-group" + "[server-group-name='" + fmt.Sprintf("%v", tacacsServerGroup.ServerGroupName) + "']"
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "servers" {
        return &tacacsServerGroup.Servers
    }
    if childYangName == "private-servers" {
        return &tacacsServerGroup.PrivateServers
    }
    return nil
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["servers"] = &tacacsServerGroup.Servers
    children["private-servers"] = &tacacsServerGroup.PrivateServers
    return children
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-group-name"] = tacacsServerGroup.ServerGroupName
    leafs["vrf"] = tacacsServerGroup.Vrf
    return leafs
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetYangName() string { return "tacacs-server-group" }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) SetParent(parent types.Entity) { tacacsServerGroup.parent = parent }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetParent() types.Entity { return tacacsServerGroup.parent }

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetParentYangName() string { return "tacacs-server-groups" }

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers
// Specify a TACACS+ server
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server.
    Server []Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetYangName() string { return "servers" }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetParent() types.Entity { return servers.parent }

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetParentYangName() string { return "tacacs-server-group" }

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetSegmentPath() string {
    return "server" + "[ordering-index='" + fmt.Sprintf("%v", server.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", server.IpAddress) + "']"
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = server.OrderingIndex
    leafs["ip-address"] = server.IpAddress
    return leafs
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetYangName() string { return "server" }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetParentYangName() string { return "servers" }

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers
// List of private TACACS servers present in the
// group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A private server to include in the server group. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer.
    PrivateServer []Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetFilter() yfilter.YFilter { return privateServers.YFilter }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) SetFilter(yf yfilter.YFilter) { privateServers.YFilter = yf }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetGoName(yname string) string {
    if yname == "private-server" { return "PrivateServer" }
    return ""
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetSegmentPath() string {
    return "private-servers"
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "private-server" {
        for _, c := range privateServers.PrivateServer {
            if privateServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer{}
        privateServers.PrivateServer = append(privateServers.PrivateServer, child)
        return &privateServers.PrivateServer[len(privateServers.PrivateServer)-1]
    }
    return nil
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range privateServers.PrivateServer {
        children[privateServers.PrivateServer[i].GetSegmentPath()] = &privateServers.PrivateServer[i]
    }
    return children
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetBundleName() string { return "cisco_ios_xr" }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetYangName() string { return "private-servers" }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) SetParent(parent types.Entity) { privateServers.parent = parent }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetParent() types.Entity { return privateServers.parent }

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetParentYangName() string { return "tacacs-server-group" }

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer
// A private server to include in the server
// group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port number (standard 49). The type is interface{}
    // with range: 1..65535.
    PortNumber interface{}

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetFilter() yfilter.YFilter { return privateServer.YFilter }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) SetFilter(yf yfilter.YFilter) { privateServer.YFilter = yf }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "key" { return "Key" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetSegmentPath() string {
    return "private-server" + "[ordering-index='" + fmt.Sprintf("%v", privateServer.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", privateServer.IpAddress) + "']" + "[port-number='" + fmt.Sprintf("%v", privateServer.PortNumber) + "']"
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = privateServer.OrderingIndex
    leafs["ip-address"] = privateServer.IpAddress
    leafs["port-number"] = privateServer.PortNumber
    leafs["key"] = privateServer.Key
    leafs["timeout"] = privateServer.Timeout
    return leafs
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetBundleName() string { return "cisco_ios_xr" }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetYangName() string { return "private-server" }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) SetParent(parent types.Entity) { privateServer.parent = parent }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetParent() types.Entity { return privateServer.parent }

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetParentYangName() string { return "private-servers" }

// Aaa_Usernames
// Configure local usernames
type Aaa_Usernames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local username. The type is slice of Aaa_Usernames_Username.
    Username []Aaa_Usernames_Username
}

func (usernames *Aaa_Usernames) GetFilter() yfilter.YFilter { return usernames.YFilter }

func (usernames *Aaa_Usernames) SetFilter(yf yfilter.YFilter) { usernames.YFilter = yf }

func (usernames *Aaa_Usernames) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    return ""
}

func (usernames *Aaa_Usernames) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-cfg:usernames"
}

func (usernames *Aaa_Usernames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "username" {
        for _, c := range usernames.Username {
            if usernames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usernames_Username{}
        usernames.Username = append(usernames.Username, child)
        return &usernames.Username[len(usernames.Username)-1]
    }
    return nil
}

func (usernames *Aaa_Usernames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usernames.Username {
        children[usernames.Username[i].GetSegmentPath()] = &usernames.Username[i]
    }
    return children
}

func (usernames *Aaa_Usernames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usernames *Aaa_Usernames) GetBundleName() string { return "cisco_ios_xr" }

func (usernames *Aaa_Usernames) GetYangName() string { return "usernames" }

func (usernames *Aaa_Usernames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usernames *Aaa_Usernames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usernames *Aaa_Usernames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usernames *Aaa_Usernames) SetParent(parent types.Entity) { usernames.parent = parent }

func (usernames *Aaa_Usernames) GetParent() types.Entity { return usernames.parent }

func (usernames *Aaa_Usernames) GetParentYangName() string { return "aaa" }

// Aaa_Usernames_Username
// Local username
type Aaa_Usernames_Username struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the users in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. Username. The type is string.
    Name interface{}

    // Specify the secret for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Secret interface{}

    // Specify the password for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Specify the usergroup to which this user belongs.
    UsergroupUnderUsernames Aaa_Usernames_Username_UsergroupUnderUsernames
}

func (username *Aaa_Usernames_Username) GetFilter() yfilter.YFilter { return username.YFilter }

func (username *Aaa_Usernames_Username) SetFilter(yf yfilter.YFilter) { username.YFilter = yf }

func (username *Aaa_Usernames_Username) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "name" { return "Name" }
    if yname == "secret" { return "Secret" }
    if yname == "password" { return "Password" }
    if yname == "usergroup-under-usernames" { return "UsergroupUnderUsernames" }
    return ""
}

func (username *Aaa_Usernames_Username) GetSegmentPath() string {
    return "username" + "[ordering-index='" + fmt.Sprintf("%v", username.OrderingIndex) + "']" + "[name='" + fmt.Sprintf("%v", username.Name) + "']"
}

func (username *Aaa_Usernames_Username) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup-under-usernames" {
        return &username.UsergroupUnderUsernames
    }
    return nil
}

func (username *Aaa_Usernames_Username) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["usergroup-under-usernames"] = &username.UsergroupUnderUsernames
    return children
}

func (username *Aaa_Usernames_Username) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = username.OrderingIndex
    leafs["name"] = username.Name
    leafs["secret"] = username.Secret
    leafs["password"] = username.Password
    return leafs
}

func (username *Aaa_Usernames_Username) GetBundleName() string { return "cisco_ios_xr" }

func (username *Aaa_Usernames_Username) GetYangName() string { return "username" }

func (username *Aaa_Usernames_Username) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (username *Aaa_Usernames_Username) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (username *Aaa_Usernames_Username) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (username *Aaa_Usernames_Username) SetParent(parent types.Entity) { username.parent = parent }

func (username *Aaa_Usernames_Username) GetParent() types.Entity { return username.parent }

func (username *Aaa_Usernames_Username) GetParentYangName() string { return "usernames" }

// Aaa_Usernames_Username_UsergroupUnderUsernames
// Specify the usergroup to which this user
// belongs
type Aaa_Usernames_Username_UsergroupUnderUsernames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is slice of
    // Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername.
    UsergroupUnderUsername []Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetFilter() yfilter.YFilter { return usergroupUnderUsernames.YFilter }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) SetFilter(yf yfilter.YFilter) { usergroupUnderUsernames.YFilter = yf }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetGoName(yname string) string {
    if yname == "usergroup-under-username" { return "UsergroupUnderUsername" }
    return ""
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetSegmentPath() string {
    return "usergroup-under-usernames"
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup-under-username" {
        for _, c := range usergroupUnderUsernames.UsergroupUnderUsername {
            if usergroupUnderUsernames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername{}
        usergroupUnderUsernames.UsergroupUnderUsername = append(usergroupUnderUsernames.UsergroupUnderUsername, child)
        return &usergroupUnderUsernames.UsergroupUnderUsername[len(usergroupUnderUsernames.UsergroupUnderUsername)-1]
    }
    return nil
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usergroupUnderUsernames.UsergroupUnderUsername {
        children[usergroupUnderUsernames.UsergroupUnderUsername[i].GetSegmentPath()] = &usergroupUnderUsernames.UsergroupUnderUsername[i]
    }
    return children
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetYangName() string { return "usergroup-under-usernames" }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) SetParent(parent types.Entity) { usergroupUnderUsernames.parent = parent }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetParent() types.Entity { return usergroupUnderUsernames.parent }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetParentYangName() string { return "username" }

// Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
// Name of the usergroup
type Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the usergroup. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetFilter() yfilter.YFilter { return usergroupUnderUsername.YFilter }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) SetFilter(yf yfilter.YFilter) { usergroupUnderUsername.YFilter = yf }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetSegmentPath() string {
    return "usergroup-under-username" + "[name='" + fmt.Sprintf("%v", usergroupUnderUsername.Name) + "']"
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = usergroupUnderUsername.Name
    return leafs
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetYangName() string { return "usergroup-under-username" }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) SetParent(parent types.Entity) { usergroupUnderUsername.parent = parent }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetParent() types.Entity { return usergroupUnderUsername.parent }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetParentYangName() string { return "usergroup-under-usernames" }

// Aaa_Taskgroups
// Specify a taskgroup to inherit from
type Aaa_Taskgroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Taskgroup name. The type is slice of Aaa_Taskgroups_Taskgroup.
    Taskgroup []Aaa_Taskgroups_Taskgroup
}

func (taskgroups *Aaa_Taskgroups) GetFilter() yfilter.YFilter { return taskgroups.YFilter }

func (taskgroups *Aaa_Taskgroups) SetFilter(yf yfilter.YFilter) { taskgroups.YFilter = yf }

func (taskgroups *Aaa_Taskgroups) GetGoName(yname string) string {
    if yname == "taskgroup" { return "Taskgroup" }
    return ""
}

func (taskgroups *Aaa_Taskgroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-cfg:taskgroups"
}

func (taskgroups *Aaa_Taskgroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup" {
        for _, c := range taskgroups.Taskgroup {
            if taskgroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup{}
        taskgroups.Taskgroup = append(taskgroups.Taskgroup, child)
        return &taskgroups.Taskgroup[len(taskgroups.Taskgroup)-1]
    }
    return nil
}

func (taskgroups *Aaa_Taskgroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskgroups.Taskgroup {
        children[taskgroups.Taskgroup[i].GetSegmentPath()] = &taskgroups.Taskgroup[i]
    }
    return children
}

func (taskgroups *Aaa_Taskgroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskgroups *Aaa_Taskgroups) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroups *Aaa_Taskgroups) GetYangName() string { return "taskgroups" }

func (taskgroups *Aaa_Taskgroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroups *Aaa_Taskgroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroups *Aaa_Taskgroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroups *Aaa_Taskgroups) SetParent(parent types.Entity) { taskgroups.parent = parent }

func (taskgroups *Aaa_Taskgroups) GetParent() types.Entity { return taskgroups.parent }

func (taskgroups *Aaa_Taskgroups) GetParentYangName() string { return "aaa" }

// Aaa_Taskgroups_Taskgroup
// Taskgroup name
type Aaa_Taskgroups_Taskgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Taskgroup name. The type is string.
    Name interface{}

    // Description for the task group. The type is string.
    Description interface{}

    // Specify a taskgroup to inherit from.
    TaskgroupUnderTaskgroups Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups

    // Specify task IDs to be part of this group.
    Tasks Aaa_Taskgroups_Taskgroup_Tasks
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetFilter() yfilter.YFilter { return taskgroup.YFilter }

func (taskgroup *Aaa_Taskgroups_Taskgroup) SetFilter(yf yfilter.YFilter) { taskgroup.YFilter = yf }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "taskgroup-under-taskgroups" { return "TaskgroupUnderTaskgroups" }
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetSegmentPath() string {
    return "taskgroup" + "[name='" + fmt.Sprintf("%v", taskgroup.Name) + "']"
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup-under-taskgroups" {
        return &taskgroup.TaskgroupUnderTaskgroups
    }
    if childYangName == "tasks" {
        return &taskgroup.Tasks
    }
    return nil
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["taskgroup-under-taskgroups"] = &taskgroup.TaskgroupUnderTaskgroups
    children["tasks"] = &taskgroup.Tasks
    return children
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = taskgroup.Name
    leafs["description"] = taskgroup.Description
    return leafs
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetYangName() string { return "taskgroup" }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) SetParent(parent types.Entity) { taskgroup.parent = parent }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetParent() types.Entity { return taskgroup.parent }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetParentYangName() string { return "taskgroups" }

// Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups
// Specify a taskgroup to inherit from
type Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the task group to include. The type is slice of
    // Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup.
    TaskgroupUnderTaskgroup []Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetFilter() yfilter.YFilter { return taskgroupUnderTaskgroups.YFilter }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) SetFilter(yf yfilter.YFilter) { taskgroupUnderTaskgroups.YFilter = yf }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetGoName(yname string) string {
    if yname == "taskgroup-under-taskgroup" { return "TaskgroupUnderTaskgroup" }
    return ""
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetSegmentPath() string {
    return "taskgroup-under-taskgroups"
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup-under-taskgroup" {
        for _, c := range taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup {
            if taskgroupUnderTaskgroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup{}
        taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup = append(taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup, child)
        return &taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup[len(taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup)-1]
    }
    return nil
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup {
        children[taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup[i].GetSegmentPath()] = &taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup[i]
    }
    return children
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetYangName() string { return "taskgroup-under-taskgroups" }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) SetParent(parent types.Entity) { taskgroupUnderTaskgroups.parent = parent }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetParent() types.Entity { return taskgroupUnderTaskgroups.parent }

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetParentYangName() string { return "taskgroup" }

// Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup
// Name of the task group to include
type Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the task group to include. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetFilter() yfilter.YFilter { return taskgroupUnderTaskgroup.YFilter }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) SetFilter(yf yfilter.YFilter) { taskgroupUnderTaskgroup.YFilter = yf }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetSegmentPath() string {
    return "taskgroup-under-taskgroup" + "[name='" + fmt.Sprintf("%v", taskgroupUnderTaskgroup.Name) + "']"
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = taskgroupUnderTaskgroup.Name
    return leafs
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetYangName() string { return "taskgroup-under-taskgroup" }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) SetParent(parent types.Entity) { taskgroupUnderTaskgroup.parent = parent }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetParent() types.Entity { return taskgroupUnderTaskgroup.parent }

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetParentYangName() string { return "taskgroup-under-taskgroups" }

// Aaa_Taskgroups_Taskgroup_Tasks
// Specify task IDs to be part of this group
type Aaa_Taskgroups_Taskgroup_Tasks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Task ID to be included. The type is slice of
    // Aaa_Taskgroups_Taskgroup_Tasks_Task.
    Task []Aaa_Taskgroups_Taskgroup_Tasks_Task
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetGoName(yname string) string {
    if yname == "task" { return "Task" }
    return ""
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "task" {
        for _, c := range tasks.Task {
            if tasks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup_Tasks_Task{}
        tasks.Task = append(tasks.Task, child)
        return &tasks.Task[len(tasks.Task)-1]
    }
    return nil
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tasks.Task {
        children[tasks.Task[i].GetSegmentPath()] = &tasks.Task[i]
    }
    return children
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetParentYangName() string { return "taskgroup" }

// Aaa_Taskgroups_Taskgroup_Tasks_Task
// Task ID to be included
type Aaa_Taskgroups_Taskgroup_Tasks_Task struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This specifies the operation permitted for this
    // task eg: read/write/execute/debug. The type is AaaLocaldTaskClass.
    Type interface{}

    // This attribute is a key. Task ID to which permission is to be granted
    // (please use class AllTasks to get a list of valid task IDs). The type is
    // string.
    TaskId interface{}
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetFilter() yfilter.YFilter { return task.YFilter }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) SetFilter(yf yfilter.YFilter) { task.YFilter = yf }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "task-id" { return "TaskId" }
    return ""
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetSegmentPath() string {
    return "task" + "[type='" + fmt.Sprintf("%v", task.Type) + "']" + "[task-id='" + fmt.Sprintf("%v", task.TaskId) + "']"
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = task.Type
    leafs["task-id"] = task.TaskId
    return leafs
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetBundleName() string { return "cisco_ios_xr" }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetYangName() string { return "task" }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) SetParent(parent types.Entity) { task.parent = parent }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetParent() types.Entity { return task.parent }

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetParentYangName() string { return "tasks" }

// Aaa_Usergroups
// Specify a Usergroup to inherit from
type Aaa_Usergroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usergroup name. The type is slice of Aaa_Usergroups_Usergroup.
    Usergroup []Aaa_Usergroups_Usergroup
}

func (usergroups *Aaa_Usergroups) GetFilter() yfilter.YFilter { return usergroups.YFilter }

func (usergroups *Aaa_Usergroups) SetFilter(yf yfilter.YFilter) { usergroups.YFilter = yf }

func (usergroups *Aaa_Usergroups) GetGoName(yname string) string {
    if yname == "usergroup" { return "Usergroup" }
    return ""
}

func (usergroups *Aaa_Usergroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-cfg:usergroups"
}

func (usergroups *Aaa_Usergroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup" {
        for _, c := range usergroups.Usergroup {
            if usergroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup{}
        usergroups.Usergroup = append(usergroups.Usergroup, child)
        return &usergroups.Usergroup[len(usergroups.Usergroup)-1]
    }
    return nil
}

func (usergroups *Aaa_Usergroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usergroups.Usergroup {
        children[usergroups.Usergroup[i].GetSegmentPath()] = &usergroups.Usergroup[i]
    }
    return children
}

func (usergroups *Aaa_Usergroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usergroups *Aaa_Usergroups) GetBundleName() string { return "cisco_ios_xr" }

func (usergroups *Aaa_Usergroups) GetYangName() string { return "usergroups" }

func (usergroups *Aaa_Usergroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroups *Aaa_Usergroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroups *Aaa_Usergroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroups *Aaa_Usergroups) SetParent(parent types.Entity) { usergroups.parent = parent }

func (usergroups *Aaa_Usergroups) GetParent() types.Entity { return usergroups.parent }

func (usergroups *Aaa_Usergroups) GetParentYangName() string { return "aaa" }

// Aaa_Usergroups_Usergroup
// Usergroup name
type Aaa_Usergroups_Usergroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Usergroup name. The type is string.
    Name interface{}

    // Description for the user group. The type is string.
    Description interface{}

    // Task group associated with this group.
    TaskgroupUnderUsergroups Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups

    // User group to be inherited by this group.
    UsergroupUnderUsergroups Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups
}

func (usergroup *Aaa_Usergroups_Usergroup) GetFilter() yfilter.YFilter { return usergroup.YFilter }

func (usergroup *Aaa_Usergroups_Usergroup) SetFilter(yf yfilter.YFilter) { usergroup.YFilter = yf }

func (usergroup *Aaa_Usergroups_Usergroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "taskgroup-under-usergroups" { return "TaskgroupUnderUsergroups" }
    if yname == "usergroup-under-usergroups" { return "UsergroupUnderUsergroups" }
    return ""
}

func (usergroup *Aaa_Usergroups_Usergroup) GetSegmentPath() string {
    return "usergroup" + "[name='" + fmt.Sprintf("%v", usergroup.Name) + "']"
}

func (usergroup *Aaa_Usergroups_Usergroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup-under-usergroups" {
        return &usergroup.TaskgroupUnderUsergroups
    }
    if childYangName == "usergroup-under-usergroups" {
        return &usergroup.UsergroupUnderUsergroups
    }
    return nil
}

func (usergroup *Aaa_Usergroups_Usergroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["taskgroup-under-usergroups"] = &usergroup.TaskgroupUnderUsergroups
    children["usergroup-under-usergroups"] = &usergroup.UsergroupUnderUsergroups
    return children
}

func (usergroup *Aaa_Usergroups_Usergroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = usergroup.Name
    leafs["description"] = usergroup.Description
    return leafs
}

func (usergroup *Aaa_Usergroups_Usergroup) GetBundleName() string { return "cisco_ios_xr" }

func (usergroup *Aaa_Usergroups_Usergroup) GetYangName() string { return "usergroup" }

func (usergroup *Aaa_Usergroups_Usergroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroup *Aaa_Usergroups_Usergroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroup *Aaa_Usergroups_Usergroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroup *Aaa_Usergroups_Usergroup) SetParent(parent types.Entity) { usergroup.parent = parent }

func (usergroup *Aaa_Usergroups_Usergroup) GetParent() types.Entity { return usergroup.parent }

func (usergroup *Aaa_Usergroups_Usergroup) GetParentYangName() string { return "usergroups" }

// Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups
// Task group associated with this group
type Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the task group. The type is slice of
    // Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup.
    TaskgroupUnderUsergroup []Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetFilter() yfilter.YFilter { return taskgroupUnderUsergroups.YFilter }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) SetFilter(yf yfilter.YFilter) { taskgroupUnderUsergroups.YFilter = yf }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetGoName(yname string) string {
    if yname == "taskgroup-under-usergroup" { return "TaskgroupUnderUsergroup" }
    return ""
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetSegmentPath() string {
    return "taskgroup-under-usergroups"
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup-under-usergroup" {
        for _, c := range taskgroupUnderUsergroups.TaskgroupUnderUsergroup {
            if taskgroupUnderUsergroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup{}
        taskgroupUnderUsergroups.TaskgroupUnderUsergroup = append(taskgroupUnderUsergroups.TaskgroupUnderUsergroup, child)
        return &taskgroupUnderUsergroups.TaskgroupUnderUsergroup[len(taskgroupUnderUsergroups.TaskgroupUnderUsergroup)-1]
    }
    return nil
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskgroupUnderUsergroups.TaskgroupUnderUsergroup {
        children[taskgroupUnderUsergroups.TaskgroupUnderUsergroup[i].GetSegmentPath()] = &taskgroupUnderUsergroups.TaskgroupUnderUsergroup[i]
    }
    return children
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetYangName() string { return "taskgroup-under-usergroups" }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) SetParent(parent types.Entity) { taskgroupUnderUsergroups.parent = parent }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetParent() types.Entity { return taskgroupUnderUsergroups.parent }

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetParentYangName() string { return "usergroup" }

// Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup
// Name of the task group
type Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the task group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetFilter() yfilter.YFilter { return taskgroupUnderUsergroup.YFilter }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) SetFilter(yf yfilter.YFilter) { taskgroupUnderUsergroup.YFilter = yf }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetSegmentPath() string {
    return "taskgroup-under-usergroup" + "[name='" + fmt.Sprintf("%v", taskgroupUnderUsergroup.Name) + "']"
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = taskgroupUnderUsergroup.Name
    return leafs
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetYangName() string { return "taskgroup-under-usergroup" }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) SetParent(parent types.Entity) { taskgroupUnderUsergroup.parent = parent }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetParent() types.Entity { return taskgroupUnderUsergroup.parent }

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetParentYangName() string { return "taskgroup-under-usergroups" }

// Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups
// User group to be inherited by this group
type Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the user group. The type is slice of
    // Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup.
    UsergroupUnderUsergroup []Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetFilter() yfilter.YFilter { return usergroupUnderUsergroups.YFilter }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) SetFilter(yf yfilter.YFilter) { usergroupUnderUsergroups.YFilter = yf }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetGoName(yname string) string {
    if yname == "usergroup-under-usergroup" { return "UsergroupUnderUsergroup" }
    return ""
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetSegmentPath() string {
    return "usergroup-under-usergroups"
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup-under-usergroup" {
        for _, c := range usergroupUnderUsergroups.UsergroupUnderUsergroup {
            if usergroupUnderUsergroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup{}
        usergroupUnderUsergroups.UsergroupUnderUsergroup = append(usergroupUnderUsergroups.UsergroupUnderUsergroup, child)
        return &usergroupUnderUsergroups.UsergroupUnderUsergroup[len(usergroupUnderUsergroups.UsergroupUnderUsergroup)-1]
    }
    return nil
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usergroupUnderUsergroups.UsergroupUnderUsergroup {
        children[usergroupUnderUsergroups.UsergroupUnderUsergroup[i].GetSegmentPath()] = &usergroupUnderUsergroups.UsergroupUnderUsergroup[i]
    }
    return children
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetYangName() string { return "usergroup-under-usergroups" }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) SetParent(parent types.Entity) { usergroupUnderUsergroups.parent = parent }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetParent() types.Entity { return usergroupUnderUsergroups.parent }

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetParentYangName() string { return "usergroup" }

// Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup
// Name of the user group
type Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the user group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetFilter() yfilter.YFilter { return usergroupUnderUsergroup.YFilter }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) SetFilter(yf yfilter.YFilter) { usergroupUnderUsergroup.YFilter = yf }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetSegmentPath() string {
    return "usergroup-under-usergroup" + "[name='" + fmt.Sprintf("%v", usergroupUnderUsergroup.Name) + "']"
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = usergroupUnderUsergroup.Name
    return leafs
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetYangName() string { return "usergroup-under-usergroup" }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) SetParent(parent types.Entity) { usergroupUnderUsergroup.parent = parent }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetParent() types.Entity { return usergroupUnderUsergroup.parent }

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetParentYangName() string { return "usergroup-under-usergroups" }

// Aaa_Diameter
// Diameter base protocol
type Aaa_Diameter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify interface for source address in DIAMETER packets. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Start diameter policy-if.
    Gy Aaa_Diameter_Gy

    // Origin sub commands.
    Origin Aaa_Diameter_Origin

    // Start diameter Nas.
    Nas Aaa_Diameter_Nas

    // TLS sub commands.
    DiameterTls Aaa_Diameter_DiameterTls

    // List of diameter peers.
    Peers Aaa_Diameter_Peers

    // Attribute list configuration for test command.
    Diams Aaa_Diameter_Diams

    // Start diameter policy-if.
    Gx Aaa_Diameter_Gx

    // Timers used for the peer.
    DiameterTimer Aaa_Diameter_DiameterTimer

    // Vendor specific.
    Vendor Aaa_Diameter_Vendor
}

func (diameter *Aaa_Diameter) GetFilter() yfilter.YFilter { return diameter.YFilter }

func (diameter *Aaa_Diameter) SetFilter(yf yfilter.YFilter) { diameter.YFilter = yf }

func (diameter *Aaa_Diameter) GetGoName(yname string) string {
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "gy" { return "Gy" }
    if yname == "origin" { return "Origin" }
    if yname == "nas" { return "Nas" }
    if yname == "diameter-tls" { return "DiameterTls" }
    if yname == "peers" { return "Peers" }
    if yname == "diams" { return "Diams" }
    if yname == "gx" { return "Gx" }
    if yname == "diameter-timer" { return "DiameterTimer" }
    if yname == "vendor" { return "Vendor" }
    return ""
}

func (diameter *Aaa_Diameter) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-diameter-cfg:diameter"
}

func (diameter *Aaa_Diameter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gy" {
        return &diameter.Gy
    }
    if childYangName == "origin" {
        return &diameter.Origin
    }
    if childYangName == "nas" {
        return &diameter.Nas
    }
    if childYangName == "diameter-tls" {
        return &diameter.DiameterTls
    }
    if childYangName == "peers" {
        return &diameter.Peers
    }
    if childYangName == "diams" {
        return &diameter.Diams
    }
    if childYangName == "gx" {
        return &diameter.Gx
    }
    if childYangName == "diameter-timer" {
        return &diameter.DiameterTimer
    }
    if childYangName == "vendor" {
        return &diameter.Vendor
    }
    return nil
}

func (diameter *Aaa_Diameter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["gy"] = &diameter.Gy
    children["origin"] = &diameter.Origin
    children["nas"] = &diameter.Nas
    children["diameter-tls"] = &diameter.DiameterTls
    children["peers"] = &diameter.Peers
    children["diams"] = &diameter.Diams
    children["gx"] = &diameter.Gx
    children["diameter-timer"] = &diameter.DiameterTimer
    children["vendor"] = &diameter.Vendor
    return children
}

func (diameter *Aaa_Diameter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-interface"] = diameter.SourceInterface
    return leafs
}

func (diameter *Aaa_Diameter) GetBundleName() string { return "cisco_ios_xr" }

func (diameter *Aaa_Diameter) GetYangName() string { return "diameter" }

func (diameter *Aaa_Diameter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameter *Aaa_Diameter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameter *Aaa_Diameter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameter *Aaa_Diameter) SetParent(parent types.Entity) { diameter.parent = parent }

func (diameter *Aaa_Diameter) GetParent() types.Entity { return diameter.parent }

func (diameter *Aaa_Diameter) GetParentYangName() string { return "aaa" }

// Aaa_Diameter_Gy
// Start diameter policy-if
type Aaa_Diameter_Gy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set retransmit count. The type is interface{} with range: 1..10.
    Retransmit interface{}

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}

    // Transaction timer value. The type is interface{} with range: 6..1000.
    TxTimer interface{}
}

func (gy *Aaa_Diameter_Gy) GetFilter() yfilter.YFilter { return gy.YFilter }

func (gy *Aaa_Diameter_Gy) SetFilter(yf yfilter.YFilter) { gy.YFilter = yf }

func (gy *Aaa_Diameter_Gy) GetGoName(yname string) string {
    if yname == "retransmit" { return "Retransmit" }
    if yname == "dest-host" { return "DestHost" }
    if yname == "tx-timer" { return "TxTimer" }
    return ""
}

func (gy *Aaa_Diameter_Gy) GetSegmentPath() string {
    return "gy"
}

func (gy *Aaa_Diameter_Gy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gy *Aaa_Diameter_Gy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gy *Aaa_Diameter_Gy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retransmit"] = gy.Retransmit
    leafs["dest-host"] = gy.DestHost
    leafs["tx-timer"] = gy.TxTimer
    return leafs
}

func (gy *Aaa_Diameter_Gy) GetBundleName() string { return "cisco_ios_xr" }

func (gy *Aaa_Diameter_Gy) GetYangName() string { return "gy" }

func (gy *Aaa_Diameter_Gy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gy *Aaa_Diameter_Gy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gy *Aaa_Diameter_Gy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gy *Aaa_Diameter_Gy) SetParent(parent types.Entity) { gy.parent = parent }

func (gy *Aaa_Diameter_Gy) GetParent() types.Entity { return gy.parent }

func (gy *Aaa_Diameter_Gy) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Origin
// Origin sub commands
type Aaa_Diameter_Origin struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Realm String. The type is string.
    Realm interface{}

    // Host name in FQDN format. The type is string.
    Host interface{}
}

func (origin *Aaa_Diameter_Origin) GetFilter() yfilter.YFilter { return origin.YFilter }

func (origin *Aaa_Diameter_Origin) SetFilter(yf yfilter.YFilter) { origin.YFilter = yf }

func (origin *Aaa_Diameter_Origin) GetGoName(yname string) string {
    if yname == "realm" { return "Realm" }
    if yname == "host" { return "Host" }
    return ""
}

func (origin *Aaa_Diameter_Origin) GetSegmentPath() string {
    return "origin"
}

func (origin *Aaa_Diameter_Origin) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (origin *Aaa_Diameter_Origin) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (origin *Aaa_Diameter_Origin) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["realm"] = origin.Realm
    leafs["host"] = origin.Host
    return leafs
}

func (origin *Aaa_Diameter_Origin) GetBundleName() string { return "cisco_ios_xr" }

func (origin *Aaa_Diameter_Origin) GetYangName() string { return "origin" }

func (origin *Aaa_Diameter_Origin) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (origin *Aaa_Diameter_Origin) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (origin *Aaa_Diameter_Origin) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (origin *Aaa_Diameter_Origin) SetParent(parent types.Entity) { origin.parent = parent }

func (origin *Aaa_Diameter_Origin) GetParent() types.Entity { return origin.parent }

func (origin *Aaa_Diameter_Origin) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Nas
// Start diameter Nas
type Aaa_Diameter_Nas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}
}

func (nas *Aaa_Diameter_Nas) GetFilter() yfilter.YFilter { return nas.YFilter }

func (nas *Aaa_Diameter_Nas) SetFilter(yf yfilter.YFilter) { nas.YFilter = yf }

func (nas *Aaa_Diameter_Nas) GetGoName(yname string) string {
    if yname == "dest-host" { return "DestHost" }
    return ""
}

func (nas *Aaa_Diameter_Nas) GetSegmentPath() string {
    return "nas"
}

func (nas *Aaa_Diameter_Nas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nas *Aaa_Diameter_Nas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nas *Aaa_Diameter_Nas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-host"] = nas.DestHost
    return leafs
}

func (nas *Aaa_Diameter_Nas) GetBundleName() string { return "cisco_ios_xr" }

func (nas *Aaa_Diameter_Nas) GetYangName() string { return "nas" }

func (nas *Aaa_Diameter_Nas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nas *Aaa_Diameter_Nas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nas *Aaa_Diameter_Nas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nas *Aaa_Diameter_Nas) SetParent(parent types.Entity) { nas.parent = parent }

func (nas *Aaa_Diameter_Nas) GetParent() types.Entity { return nas.parent }

func (nas *Aaa_Diameter_Nas) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_DiameterTls
// TLS sub commands
type Aaa_Diameter_DiameterTls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trustpoint label to be used. The type is string.
    Trustpoint interface{}
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetFilter() yfilter.YFilter { return diameterTls.YFilter }

func (diameterTls *Aaa_Diameter_DiameterTls) SetFilter(yf yfilter.YFilter) { diameterTls.YFilter = yf }

func (diameterTls *Aaa_Diameter_DiameterTls) GetGoName(yname string) string {
    if yname == "trustpoint" { return "Trustpoint" }
    return ""
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetSegmentPath() string {
    return "diameter-tls"
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trustpoint"] = diameterTls.Trustpoint
    return leafs
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetBundleName() string { return "cisco_ios_xr" }

func (diameterTls *Aaa_Diameter_DiameterTls) GetYangName() string { return "diameter-tls" }

func (diameterTls *Aaa_Diameter_DiameterTls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameterTls *Aaa_Diameter_DiameterTls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameterTls *Aaa_Diameter_DiameterTls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameterTls *Aaa_Diameter_DiameterTls) SetParent(parent types.Entity) { diameterTls.parent = parent }

func (diameterTls *Aaa_Diameter_DiameterTls) GetParent() types.Entity { return diameterTls.parent }

func (diameterTls *Aaa_Diameter_DiameterTls) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Peers
// List of diameter peers
type Aaa_Diameter_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diameter peer instance. The type is slice of Aaa_Diameter_Peers_Peer.
    Peer []Aaa_Diameter_Peers_Peer
}

func (peers *Aaa_Diameter_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Aaa_Diameter_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Aaa_Diameter_Peers) GetGoName(yname string) string {
    if yname == "peer" { return "Peer" }
    return ""
}

func (peers *Aaa_Diameter_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Aaa_Diameter_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peers.Peer {
            if peers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_Peers_Peer{}
        peers.Peer = append(peers.Peer, child)
        return &peers.Peer[len(peers.Peer)-1]
    }
    return nil
}

func (peers *Aaa_Diameter_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peers.Peer {
        children[peers.Peer[i].GetSegmentPath()] = &peers.Peer[i]
    }
    return children
}

func (peers *Aaa_Diameter_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *Aaa_Diameter_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Aaa_Diameter_Peers) GetYangName() string { return "peers" }

func (peers *Aaa_Diameter_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Aaa_Diameter_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Aaa_Diameter_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Aaa_Diameter_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Aaa_Diameter_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Aaa_Diameter_Peers) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Peers_Peer
// Diameter peer instance
type Aaa_Diameter_Peers_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name for the diameter peer configuration. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PeerName interface{}

    // Destination host information. The type is string.
    HostDestination interface{}

    // IPv4 address of diameter server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Realm to which the peer belongs to. The type is string.
    RealmDestination interface{}

    // Specify a Diameter transport protocol. The type is interface{} with range:
    // 1..65535.
    TcpTransport interface{}

    // Specify interface for source address in DIAMETER packets. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // IPv6 address of diameter server. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Specify a Diameter security type. The type is interface{} with range: 0..1.
    TlsTransport interface{}

    // VRF the peer belongs to. The type is string.
    VrfIp interface{}

    // Timers used for the peer.
    PeerTimer Aaa_Diameter_Peers_Peer_PeerTimer

    // Peer Type.
    PeerType Aaa_Diameter_Peers_Peer_PeerType
}

func (peer *Aaa_Diameter_Peers_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Aaa_Diameter_Peers_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Aaa_Diameter_Peers_Peer) GetGoName(yname string) string {
    if yname == "peer-name" { return "PeerName" }
    if yname == "host-destination" { return "HostDestination" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "realm-destination" { return "RealmDestination" }
    if yname == "tcp-transport" { return "TcpTransport" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "tls-transport" { return "TlsTransport" }
    if yname == "vrf-ip" { return "VrfIp" }
    if yname == "peer-timer" { return "PeerTimer" }
    if yname == "peer-type" { return "PeerType" }
    return ""
}

func (peer *Aaa_Diameter_Peers_Peer) GetSegmentPath() string {
    return "peer" + "[peer-name='" + fmt.Sprintf("%v", peer.PeerName) + "']"
}

func (peer *Aaa_Diameter_Peers_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-timer" {
        return &peer.PeerTimer
    }
    if childYangName == "peer-type" {
        return &peer.PeerType
    }
    return nil
}

func (peer *Aaa_Diameter_Peers_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-timer"] = &peer.PeerTimer
    children["peer-type"] = &peer.PeerType
    return children
}

func (peer *Aaa_Diameter_Peers_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-name"] = peer.PeerName
    leafs["host-destination"] = peer.HostDestination
    leafs["ipv4-address"] = peer.Ipv4Address
    leafs["realm-destination"] = peer.RealmDestination
    leafs["tcp-transport"] = peer.TcpTransport
    leafs["source-interface"] = peer.SourceInterface
    leafs["ipv6-address"] = peer.Ipv6Address
    leafs["tls-transport"] = peer.TlsTransport
    leafs["vrf-ip"] = peer.VrfIp
    return leafs
}

func (peer *Aaa_Diameter_Peers_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Aaa_Diameter_Peers_Peer) GetYangName() string { return "peer" }

func (peer *Aaa_Diameter_Peers_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Aaa_Diameter_Peers_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Aaa_Diameter_Peers_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Aaa_Diameter_Peers_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Aaa_Diameter_Peers_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Aaa_Diameter_Peers_Peer) GetParentYangName() string { return "peers" }

// Aaa_Diameter_Peers_Peer_PeerTimer
// Timers used for the peer
type Aaa_Diameter_Peers_Peer_PeerTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Transaction interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Connection interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Watchdog interface{}
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetFilter() yfilter.YFilter { return peerTimer.YFilter }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) SetFilter(yf yfilter.YFilter) { peerTimer.YFilter = yf }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetGoName(yname string) string {
    if yname == "transaction" { return "Transaction" }
    if yname == "connection" { return "Connection" }
    if yname == "watchdog" { return "Watchdog" }
    return ""
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetSegmentPath() string {
    return "peer-timer"
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transaction"] = peerTimer.Transaction
    leafs["connection"] = peerTimer.Connection
    leafs["watchdog"] = peerTimer.Watchdog
    return leafs
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetBundleName() string { return "cisco_ios_xr" }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetYangName() string { return "peer-timer" }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) SetParent(parent types.Entity) { peerTimer.parent = parent }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetParent() types.Entity { return peerTimer.parent }

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetParentYangName() string { return "peer" }

// Aaa_Diameter_Peers_Peer_PeerType
// Peer Type
type Aaa_Diameter_Peers_Peer_PeerType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled. The type is bool.
    Server interface{}
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetFilter() yfilter.YFilter { return peerType.YFilter }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) SetFilter(yf yfilter.YFilter) { peerType.YFilter = yf }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetSegmentPath() string {
    return "peer-type"
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server"] = peerType.Server
    return leafs
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetBundleName() string { return "cisco_ios_xr" }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetYangName() string { return "peer-type" }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) SetParent(parent types.Entity) { peerType.parent = parent }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetParent() types.Entity { return peerType.parent }

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetParentYangName() string { return "peer" }

// Aaa_Diameter_Diams
// Attribute list configuration for test command
type Aaa_Diameter_Diams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // attribute list configuration. The type is slice of Aaa_Diameter_Diams_Diam.
    Diam []Aaa_Diameter_Diams_Diam
}

func (diams *Aaa_Diameter_Diams) GetFilter() yfilter.YFilter { return diams.YFilter }

func (diams *Aaa_Diameter_Diams) SetFilter(yf yfilter.YFilter) { diams.YFilter = yf }

func (diams *Aaa_Diameter_Diams) GetGoName(yname string) string {
    if yname == "diam" { return "Diam" }
    return ""
}

func (diams *Aaa_Diameter_Diams) GetSegmentPath() string {
    return "diams"
}

func (diams *Aaa_Diameter_Diams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diam" {
        for _, c := range diams.Diam {
            if diams.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_Diams_Diam{}
        diams.Diam = append(diams.Diam, child)
        return &diams.Diam[len(diams.Diam)-1]
    }
    return nil
}

func (diams *Aaa_Diameter_Diams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diams.Diam {
        children[diams.Diam[i].GetSegmentPath()] = &diams.Diam[i]
    }
    return children
}

func (diams *Aaa_Diameter_Diams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diams *Aaa_Diameter_Diams) GetBundleName() string { return "cisco_ios_xr" }

func (diams *Aaa_Diameter_Diams) GetYangName() string { return "diams" }

func (diams *Aaa_Diameter_Diams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diams *Aaa_Diameter_Diams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diams *Aaa_Diameter_Diams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diams *Aaa_Diameter_Diams) SetParent(parent types.Entity) { diams.parent = parent }

func (diams *Aaa_Diameter_Diams) GetParent() types.Entity { return diams.parent }

func (diams *Aaa_Diameter_Diams) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Diams_Diam
// attribute list configuration
type Aaa_Diameter_Diams_Diam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. attribute list number. The type is interface{}
    // with range: 0..99.
    ListId interface{}

    // Specify an attribute definition.
    DiamAttrDefs Aaa_Diameter_Diams_Diam_DiamAttrDefs
}

func (diam *Aaa_Diameter_Diams_Diam) GetFilter() yfilter.YFilter { return diam.YFilter }

func (diam *Aaa_Diameter_Diams_Diam) SetFilter(yf yfilter.YFilter) { diam.YFilter = yf }

func (diam *Aaa_Diameter_Diams_Diam) GetGoName(yname string) string {
    if yname == "list-id" { return "ListId" }
    if yname == "diam-attr-defs" { return "DiamAttrDefs" }
    return ""
}

func (diam *Aaa_Diameter_Diams_Diam) GetSegmentPath() string {
    return "diam" + "[list-id='" + fmt.Sprintf("%v", diam.ListId) + "']"
}

func (diam *Aaa_Diameter_Diams_Diam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diam-attr-defs" {
        return &diam.DiamAttrDefs
    }
    return nil
}

func (diam *Aaa_Diameter_Diams_Diam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["diam-attr-defs"] = &diam.DiamAttrDefs
    return children
}

func (diam *Aaa_Diameter_Diams_Diam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["list-id"] = diam.ListId
    return leafs
}

func (diam *Aaa_Diameter_Diams_Diam) GetBundleName() string { return "cisco_ios_xr" }

func (diam *Aaa_Diameter_Diams_Diam) GetYangName() string { return "diam" }

func (diam *Aaa_Diameter_Diams_Diam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diam *Aaa_Diameter_Diams_Diam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diam *Aaa_Diameter_Diams_Diam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diam *Aaa_Diameter_Diams_Diam) SetParent(parent types.Entity) { diam.parent = parent }

func (diam *Aaa_Diameter_Diams_Diam) GetParent() types.Entity { return diam.parent }

func (diam *Aaa_Diameter_Diams_Diam) GetParentYangName() string { return "diams" }

// Aaa_Diameter_Diams_Diam_DiamAttrDefs
// Specify an attribute definition
type Aaa_Diameter_Diams_Diam_DiamAttrDefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // vendor id. The type is slice of
    // Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef.
    DiamAttrDef []Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetFilter() yfilter.YFilter { return diamAttrDefs.YFilter }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) SetFilter(yf yfilter.YFilter) { diamAttrDefs.YFilter = yf }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetGoName(yname string) string {
    if yname == "diam-attr-def" { return "DiamAttrDef" }
    return ""
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetSegmentPath() string {
    return "diam-attr-defs"
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diam-attr-def" {
        for _, c := range diamAttrDefs.DiamAttrDef {
            if diamAttrDefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef{}
        diamAttrDefs.DiamAttrDef = append(diamAttrDefs.DiamAttrDef, child)
        return &diamAttrDefs.DiamAttrDef[len(diamAttrDefs.DiamAttrDef)-1]
    }
    return nil
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diamAttrDefs.DiamAttrDef {
        children[diamAttrDefs.DiamAttrDef[i].GetSegmentPath()] = &diamAttrDefs.DiamAttrDef[i]
    }
    return children
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetBundleName() string { return "cisco_ios_xr" }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetYangName() string { return "diam-attr-defs" }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) SetParent(parent types.Entity) { diamAttrDefs.parent = parent }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetParent() types.Entity { return diamAttrDefs.parent }

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetParentYangName() string { return "diam" }

// Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef
// vendor id
type Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. value for vendor id. The type is interface{} with
    // range: 0..4294967295.
    VendorId interface{}

    // This attribute is a key. enter attribute id. The type is interface{} with
    // range: 1..65535.
    AttributeId interface{}

    // attr subcommands.
    DiamAttrValue Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetFilter() yfilter.YFilter { return diamAttrDef.YFilter }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) SetFilter(yf yfilter.YFilter) { diamAttrDef.YFilter = yf }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetGoName(yname string) string {
    if yname == "vendor-id" { return "VendorId" }
    if yname == "attribute-id" { return "AttributeId" }
    if yname == "diam-attr-value" { return "DiamAttrValue" }
    return ""
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetSegmentPath() string {
    return "diam-attr-def" + "[vendor-id='" + fmt.Sprintf("%v", diamAttrDef.VendorId) + "']" + "[attribute-id='" + fmt.Sprintf("%v", diamAttrDef.AttributeId) + "']"
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diam-attr-value" {
        return &diamAttrDef.DiamAttrValue
    }
    return nil
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["diam-attr-value"] = &diamAttrDef.DiamAttrValue
    return children
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vendor-id"] = diamAttrDef.VendorId
    leafs["attribute-id"] = diamAttrDef.AttributeId
    return leafs
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetBundleName() string { return "cisco_ios_xr" }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetYangName() string { return "diam-attr-def" }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) SetParent(parent types.Entity) { diamAttrDef.parent = parent }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetParent() types.Entity { return diamAttrDef.parent }

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetParentYangName() string { return "diam-attr-defs" }

// Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue
// attr subcommands
type Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String type. The type is string.
    TypeString interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TypeIpv4Address interface{}

    // Binary type. The type is string.
    TypeBinary interface{}

    // Boolean type. The type is interface{} with range: 0..4294967295.
    TypeBoolean interface{}

    // Enumeration type. The type is interface{} with range: 0..4294967295.
    TypeEnum interface{}

    // Grouped attribute. The type is interface{} with range: 0..99.
    TypeGrouped interface{}

    // Numeric type. The type is interface{} with range: 0..4294967295.
    TypeUlong interface{}

    // Diameter-identity type. The type is string.
    TypeIdentity interface{}

    // Dataypte of attribute. The type is interface{} with range: 0..23.
    DataType interface{}

    // Is mandatory?. The type is interface{} with range: 0..1.
    Mandatory interface{}
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetFilter() yfilter.YFilter { return diamAttrValue.YFilter }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) SetFilter(yf yfilter.YFilter) { diamAttrValue.YFilter = yf }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetGoName(yname string) string {
    if yname == "type-string" { return "TypeString" }
    if yname == "type-ipv4-address" { return "TypeIpv4Address" }
    if yname == "type-binary" { return "TypeBinary" }
    if yname == "type-boolean" { return "TypeBoolean" }
    if yname == "type-enum" { return "TypeEnum" }
    if yname == "type-grouped" { return "TypeGrouped" }
    if yname == "type-ulong" { return "TypeUlong" }
    if yname == "type-identity" { return "TypeIdentity" }
    if yname == "data-type" { return "DataType" }
    if yname == "mandatory" { return "Mandatory" }
    return ""
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetSegmentPath() string {
    return "diam-attr-value"
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type-string"] = diamAttrValue.TypeString
    leafs["type-ipv4-address"] = diamAttrValue.TypeIpv4Address
    leafs["type-binary"] = diamAttrValue.TypeBinary
    leafs["type-boolean"] = diamAttrValue.TypeBoolean
    leafs["type-enum"] = diamAttrValue.TypeEnum
    leafs["type-grouped"] = diamAttrValue.TypeGrouped
    leafs["type-ulong"] = diamAttrValue.TypeUlong
    leafs["type-identity"] = diamAttrValue.TypeIdentity
    leafs["data-type"] = diamAttrValue.DataType
    leafs["mandatory"] = diamAttrValue.Mandatory
    return leafs
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetBundleName() string { return "cisco_ios_xr" }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetYangName() string { return "diam-attr-value" }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) SetParent(parent types.Entity) { diamAttrValue.parent = parent }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetParent() types.Entity { return diamAttrValue.parent }

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetParentYangName() string { return "diam-attr-def" }

// Aaa_Diameter_Gx
// Start diameter policy-if
type Aaa_Diameter_Gx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set retransmit count. The type is interface{} with range: 1..10.
    Retransmit interface{}

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}

    // Transaction timer value. The type is interface{} with range: 6..1000.
    TxTimer interface{}
}

func (gx *Aaa_Diameter_Gx) GetFilter() yfilter.YFilter { return gx.YFilter }

func (gx *Aaa_Diameter_Gx) SetFilter(yf yfilter.YFilter) { gx.YFilter = yf }

func (gx *Aaa_Diameter_Gx) GetGoName(yname string) string {
    if yname == "retransmit" { return "Retransmit" }
    if yname == "dest-host" { return "DestHost" }
    if yname == "tx-timer" { return "TxTimer" }
    return ""
}

func (gx *Aaa_Diameter_Gx) GetSegmentPath() string {
    return "gx"
}

func (gx *Aaa_Diameter_Gx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gx *Aaa_Diameter_Gx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gx *Aaa_Diameter_Gx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retransmit"] = gx.Retransmit
    leafs["dest-host"] = gx.DestHost
    leafs["tx-timer"] = gx.TxTimer
    return leafs
}

func (gx *Aaa_Diameter_Gx) GetBundleName() string { return "cisco_ios_xr" }

func (gx *Aaa_Diameter_Gx) GetYangName() string { return "gx" }

func (gx *Aaa_Diameter_Gx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gx *Aaa_Diameter_Gx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gx *Aaa_Diameter_Gx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gx *Aaa_Diameter_Gx) SetParent(parent types.Entity) { gx.parent = parent }

func (gx *Aaa_Diameter_Gx) GetParent() types.Entity { return gx.parent }

func (gx *Aaa_Diameter_Gx) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_DiameterTimer
// Timers used for the peer
type Aaa_Diameter_DiameterTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Transaction interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Connection interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Watchdog interface{}
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetFilter() yfilter.YFilter { return diameterTimer.YFilter }

func (diameterTimer *Aaa_Diameter_DiameterTimer) SetFilter(yf yfilter.YFilter) { diameterTimer.YFilter = yf }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetGoName(yname string) string {
    if yname == "transaction" { return "Transaction" }
    if yname == "connection" { return "Connection" }
    if yname == "watchdog" { return "Watchdog" }
    return ""
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetSegmentPath() string {
    return "diameter-timer"
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transaction"] = diameterTimer.Transaction
    leafs["connection"] = diameterTimer.Connection
    leafs["watchdog"] = diameterTimer.Watchdog
    return leafs
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetBundleName() string { return "cisco_ios_xr" }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetYangName() string { return "diameter-timer" }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameterTimer *Aaa_Diameter_DiameterTimer) SetParent(parent types.Entity) { diameterTimer.parent = parent }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetParent() types.Entity { return diameterTimer.parent }

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Vendor
// Vendor specific
type Aaa_Diameter_Vendor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Supported vendors.
    Supported Aaa_Diameter_Vendor_Supported
}

func (vendor *Aaa_Diameter_Vendor) GetFilter() yfilter.YFilter { return vendor.YFilter }

func (vendor *Aaa_Diameter_Vendor) SetFilter(yf yfilter.YFilter) { vendor.YFilter = yf }

func (vendor *Aaa_Diameter_Vendor) GetGoName(yname string) string {
    if yname == "supported" { return "Supported" }
    return ""
}

func (vendor *Aaa_Diameter_Vendor) GetSegmentPath() string {
    return "vendor"
}

func (vendor *Aaa_Diameter_Vendor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "supported" {
        return &vendor.Supported
    }
    return nil
}

func (vendor *Aaa_Diameter_Vendor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["supported"] = &vendor.Supported
    return children
}

func (vendor *Aaa_Diameter_Vendor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vendor *Aaa_Diameter_Vendor) GetBundleName() string { return "cisco_ios_xr" }

func (vendor *Aaa_Diameter_Vendor) GetYangName() string { return "vendor" }

func (vendor *Aaa_Diameter_Vendor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vendor *Aaa_Diameter_Vendor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vendor *Aaa_Diameter_Vendor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vendor *Aaa_Diameter_Vendor) SetParent(parent types.Entity) { vendor.parent = parent }

func (vendor *Aaa_Diameter_Vendor) GetParent() types.Entity { return vendor.parent }

func (vendor *Aaa_Diameter_Vendor) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Vendor_Supported
// Supported vendors
type Aaa_Diameter_Vendor_Supported struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Cisco attribute support. The type is bool.
    Cisco interface{}

    // 3GPP attribute support. The type is bool.
    Threegpp interface{}

    // Etsi attribute support. The type is bool.
    Etsi interface{}

    // Vodafone attribute support. The type is bool.
    Vodafone interface{}
}

func (supported *Aaa_Diameter_Vendor_Supported) GetFilter() yfilter.YFilter { return supported.YFilter }

func (supported *Aaa_Diameter_Vendor_Supported) SetFilter(yf yfilter.YFilter) { supported.YFilter = yf }

func (supported *Aaa_Diameter_Vendor_Supported) GetGoName(yname string) string {
    if yname == "cisco" { return "Cisco" }
    if yname == "threegpp" { return "Threegpp" }
    if yname == "etsi" { return "Etsi" }
    if yname == "vodafone" { return "Vodafone" }
    return ""
}

func (supported *Aaa_Diameter_Vendor_Supported) GetSegmentPath() string {
    return "supported"
}

func (supported *Aaa_Diameter_Vendor_Supported) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (supported *Aaa_Diameter_Vendor_Supported) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (supported *Aaa_Diameter_Vendor_Supported) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cisco"] = supported.Cisco
    leafs["threegpp"] = supported.Threegpp
    leafs["etsi"] = supported.Etsi
    leafs["vodafone"] = supported.Vodafone
    return leafs
}

func (supported *Aaa_Diameter_Vendor_Supported) GetBundleName() string { return "cisco_ios_xr" }

func (supported *Aaa_Diameter_Vendor_Supported) GetYangName() string { return "supported" }

func (supported *Aaa_Diameter_Vendor_Supported) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (supported *Aaa_Diameter_Vendor_Supported) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (supported *Aaa_Diameter_Vendor_Supported) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (supported *Aaa_Diameter_Vendor_Supported) SetParent(parent types.Entity) { supported.parent = parent }

func (supported *Aaa_Diameter_Vendor_Supported) GetParent() types.Entity { return supported.parent }

func (supported *Aaa_Diameter_Vendor_Supported) GetParentYangName() string { return "vendor" }

// Aaa_Radius
// Remote Access Dial-In User Service
type Aaa_Radius struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of times to retransmit a request to the RADIUS server(0-Disable).
    // The type is interface{} with range: 0..100. The default value is 3.
    Retransmit interface{}

    // This indicates the length of time (in minutes) for which a RADIUS server
    // remains marked as dead. The type is interface{} with range: 1..1440. Units
    // are minute.
    DeadTime interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAccountingPort interface{}

    // Idle time for RADIUS server. The type is interface{} with range: 1..1000.
    // The default value is 5.
    IdleTime interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAuthPort interface{}

    // List of RADIUS servers.
    Hosts Aaa_Radius_Hosts

    // RADIUS server dead criteria.
    DeadCriteria Aaa_Radius_DeadCriteria

    // disallow null-username.
    Disallow Aaa_Radius_Disallow

    // IPv6 configuration.
    Ipv6 Aaa_Radius_Ipv6

    // RADIUS dynamic authorization.
    DynamicAuthorization Aaa_Radius_DynamicAuthorization

    // Radius load-balancing options.
    LoadBalanceOptions Aaa_Radius_LoadBalanceOptions

    // List of VRFs.
    Vrfs Aaa_Radius_Vrfs

    // Radius throttling options.
    Throttle Aaa_Radius_Throttle

    // Unknown VSA (Vendor Specific Attribute) ignore configuration for RADIUS
    // server.
    Vsa Aaa_Radius_Vsa

    // IPv4 configuration.
    Ipv4 Aaa_Radius_Ipv4

    // attribute.
    RadiusAttribute Aaa_Radius_RadiusAttribute

    // Table of attribute list.
    Attributes Aaa_Radius_Attributes

    // Source port.
    SourcePort Aaa_Radius_SourcePort
}

func (radius *Aaa_Radius) GetFilter() yfilter.YFilter { return radius.YFilter }

func (radius *Aaa_Radius) SetFilter(yf yfilter.YFilter) { radius.YFilter = yf }

func (radius *Aaa_Radius) GetGoName(yname string) string {
    if yname == "retransmit" { return "Retransmit" }
    if yname == "dead-time" { return "DeadTime" }
    if yname == "key" { return "Key" }
    if yname == "timeout" { return "Timeout" }
    if yname == "ignore-accounting-port" { return "IgnoreAccountingPort" }
    if yname == "idle-time" { return "IdleTime" }
    if yname == "username" { return "Username" }
    if yname == "ignore-auth-port" { return "IgnoreAuthPort" }
    if yname == "hosts" { return "Hosts" }
    if yname == "dead-criteria" { return "DeadCriteria" }
    if yname == "disallow" { return "Disallow" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "dynamic-authorization" { return "DynamicAuthorization" }
    if yname == "load-balance-options" { return "LoadBalanceOptions" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "throttle" { return "Throttle" }
    if yname == "vsa" { return "Vsa" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "radius-attribute" { return "RadiusAttribute" }
    if yname == "attributes" { return "Attributes" }
    if yname == "source-port" { return "SourcePort" }
    return ""
}

func (radius *Aaa_Radius) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius"
}

func (radius *Aaa_Radius) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hosts" {
        return &radius.Hosts
    }
    if childYangName == "dead-criteria" {
        return &radius.DeadCriteria
    }
    if childYangName == "disallow" {
        return &radius.Disallow
    }
    if childYangName == "ipv6" {
        return &radius.Ipv6
    }
    if childYangName == "dynamic-authorization" {
        return &radius.DynamicAuthorization
    }
    if childYangName == "load-balance-options" {
        return &radius.LoadBalanceOptions
    }
    if childYangName == "vrfs" {
        return &radius.Vrfs
    }
    if childYangName == "throttle" {
        return &radius.Throttle
    }
    if childYangName == "vsa" {
        return &radius.Vsa
    }
    if childYangName == "ipv4" {
        return &radius.Ipv4
    }
    if childYangName == "radius-attribute" {
        return &radius.RadiusAttribute
    }
    if childYangName == "attributes" {
        return &radius.Attributes
    }
    if childYangName == "source-port" {
        return &radius.SourcePort
    }
    return nil
}

func (radius *Aaa_Radius) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hosts"] = &radius.Hosts
    children["dead-criteria"] = &radius.DeadCriteria
    children["disallow"] = &radius.Disallow
    children["ipv6"] = &radius.Ipv6
    children["dynamic-authorization"] = &radius.DynamicAuthorization
    children["load-balance-options"] = &radius.LoadBalanceOptions
    children["vrfs"] = &radius.Vrfs
    children["throttle"] = &radius.Throttle
    children["vsa"] = &radius.Vsa
    children["ipv4"] = &radius.Ipv4
    children["radius-attribute"] = &radius.RadiusAttribute
    children["attributes"] = &radius.Attributes
    children["source-port"] = &radius.SourcePort
    return children
}

func (radius *Aaa_Radius) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retransmit"] = radius.Retransmit
    leafs["dead-time"] = radius.DeadTime
    leafs["key"] = radius.Key
    leafs["timeout"] = radius.Timeout
    leafs["ignore-accounting-port"] = radius.IgnoreAccountingPort
    leafs["idle-time"] = radius.IdleTime
    leafs["username"] = radius.Username
    leafs["ignore-auth-port"] = radius.IgnoreAuthPort
    return leafs
}

func (radius *Aaa_Radius) GetBundleName() string { return "cisco_ios_xr" }

func (radius *Aaa_Radius) GetYangName() string { return "radius" }

func (radius *Aaa_Radius) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radius *Aaa_Radius) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radius *Aaa_Radius) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radius *Aaa_Radius) SetParent(parent types.Entity) { radius.parent = parent }

func (radius *Aaa_Radius) GetParent() types.Entity { return radius.parent }

func (radius *Aaa_Radius) GetParentYangName() string { return "aaa" }

// Aaa_Radius_Hosts
// List of RADIUS servers
type Aaa_Radius_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance of a RADIUS server. The type is slice of Aaa_Radius_Hosts_Host.
    Host []Aaa_Radius_Hosts_Host
}

func (hosts *Aaa_Radius_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Aaa_Radius_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Aaa_Radius_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Aaa_Radius_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Aaa_Radius_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Aaa_Radius_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Aaa_Radius_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Aaa_Radius_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Aaa_Radius_Hosts) GetYangName() string { return "hosts" }

func (hosts *Aaa_Radius_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Aaa_Radius_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Aaa_Radius_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Aaa_Radius_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Aaa_Radius_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Aaa_Radius_Hosts) GetParentYangName() string { return "radius" }

// Aaa_Radius_Hosts_Host
// Instance of a RADIUS server
type Aaa_Radius_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}

    // Number of times to retransmit a request to the RADIUS server. The type is
    // interface{} with range: 1..100. The default value is 3.
    HostRetransmit interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    HostTimeout interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    HostKey interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAccountingPort interface{}

    // Idle time for RADIUS server. The type is interface{} with range: 1..1000.
    // The default value is 5.
    IdleTime interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAuthPort interface{}
}

func (host *Aaa_Radius_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Aaa_Radius_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Aaa_Radius_Hosts_Host) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auth-port-number" { return "AuthPortNumber" }
    if yname == "acct-port-number" { return "AcctPortNumber" }
    if yname == "host-retransmit" { return "HostRetransmit" }
    if yname == "host-timeout" { return "HostTimeout" }
    if yname == "host-key" { return "HostKey" }
    if yname == "ignore-accounting-port" { return "IgnoreAccountingPort" }
    if yname == "idle-time" { return "IdleTime" }
    if yname == "username" { return "Username" }
    if yname == "ignore-auth-port" { return "IgnoreAuthPort" }
    return ""
}

func (host *Aaa_Radius_Hosts_Host) GetSegmentPath() string {
    return "host" + "[ordering-index='" + fmt.Sprintf("%v", host.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[auth-port-number='" + fmt.Sprintf("%v", host.AuthPortNumber) + "']" + "[acct-port-number='" + fmt.Sprintf("%v", host.AcctPortNumber) + "']"
}

func (host *Aaa_Radius_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *Aaa_Radius_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *Aaa_Radius_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = host.OrderingIndex
    leafs["ip-address"] = host.IpAddress
    leafs["auth-port-number"] = host.AuthPortNumber
    leafs["acct-port-number"] = host.AcctPortNumber
    leafs["host-retransmit"] = host.HostRetransmit
    leafs["host-timeout"] = host.HostTimeout
    leafs["host-key"] = host.HostKey
    leafs["ignore-accounting-port"] = host.IgnoreAccountingPort
    leafs["idle-time"] = host.IdleTime
    leafs["username"] = host.Username
    leafs["ignore-auth-port"] = host.IgnoreAuthPort
    return leafs
}

func (host *Aaa_Radius_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Aaa_Radius_Hosts_Host) GetYangName() string { return "host" }

func (host *Aaa_Radius_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Aaa_Radius_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Aaa_Radius_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Aaa_Radius_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Aaa_Radius_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Aaa_Radius_Hosts_Host) GetParentYangName() string { return "hosts" }

// Aaa_Radius_DeadCriteria
// RADIUS server dead criteria
type Aaa_Radius_DeadCriteria struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of consecutive timeouts the router must experience in order to
    // mark the server as dead. All transmissions, including the initial transmit
    // and all retransmits, will be counted. The type is interface{} with range:
    // 1..100.
    Tries interface{}

    // The minimum amount of time which must elapse since the router last received
    // a valid RADIUS packet from the server prior to marking it dead. The type is
    // interface{} with range: 1..120.
    Time interface{}
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetFilter() yfilter.YFilter { return deadCriteria.YFilter }

func (deadCriteria *Aaa_Radius_DeadCriteria) SetFilter(yf yfilter.YFilter) { deadCriteria.YFilter = yf }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetGoName(yname string) string {
    if yname == "tries" { return "Tries" }
    if yname == "time" { return "Time" }
    return ""
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetSegmentPath() string {
    return "dead-criteria"
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tries"] = deadCriteria.Tries
    leafs["time"] = deadCriteria.Time
    return leafs
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetBundleName() string { return "cisco_ios_xr" }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetYangName() string { return "dead-criteria" }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (deadCriteria *Aaa_Radius_DeadCriteria) SetParent(parent types.Entity) { deadCriteria.parent = parent }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetParent() types.Entity { return deadCriteria.parent }

func (deadCriteria *Aaa_Radius_DeadCriteria) GetParentYangName() string { return "radius" }

// Aaa_Radius_Disallow
// disallow null-username
type Aaa_Radius_Disallow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disallow null-username. The type is interface{} with range:
    // -2147483648..2147483647.
    NullUsername interface{}
}

func (disallow *Aaa_Radius_Disallow) GetFilter() yfilter.YFilter { return disallow.YFilter }

func (disallow *Aaa_Radius_Disallow) SetFilter(yf yfilter.YFilter) { disallow.YFilter = yf }

func (disallow *Aaa_Radius_Disallow) GetGoName(yname string) string {
    if yname == "null-username" { return "NullUsername" }
    return ""
}

func (disallow *Aaa_Radius_Disallow) GetSegmentPath() string {
    return "disallow"
}

func (disallow *Aaa_Radius_Disallow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (disallow *Aaa_Radius_Disallow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (disallow *Aaa_Radius_Disallow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["null-username"] = disallow.NullUsername
    return leafs
}

func (disallow *Aaa_Radius_Disallow) GetBundleName() string { return "cisco_ios_xr" }

func (disallow *Aaa_Radius_Disallow) GetYangName() string { return "disallow" }

func (disallow *Aaa_Radius_Disallow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (disallow *Aaa_Radius_Disallow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (disallow *Aaa_Radius_Disallow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (disallow *Aaa_Radius_Disallow) SetParent(parent types.Entity) { disallow.parent = parent }

func (disallow *Aaa_Radius_Disallow) GetParent() types.Entity { return disallow.parent }

func (disallow *Aaa_Radius_Disallow) GetParentYangName() string { return "radius" }

// Aaa_Radius_Ipv6
// IPv6 configuration
type Aaa_Radius_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // AaaDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv6 *Aaa_Radius_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Aaa_Radius_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Aaa_Radius_Ipv6) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (ipv6 *Aaa_Radius_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Aaa_Radius_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Aaa_Radius_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Aaa_Radius_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ipv6.Dscp
    return leafs
}

func (ipv6 *Aaa_Radius_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Aaa_Radius_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Aaa_Radius_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Aaa_Radius_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Aaa_Radius_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Aaa_Radius_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Aaa_Radius_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Aaa_Radius_Ipv6) GetParentYangName() string { return "radius" }

// Aaa_Radius_DynamicAuthorization
// RADIUS dynamic authorization
type Aaa_Radius_DynamicAuthorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore option for server key or session key. The type is AaaSelectKey.
    Ignore interface{}

    // Specify the COA server port to listen on. The type is interface{} with
    // range: 1000..5000.
    Port interface{}

    // RADIUS  dynamic  authorization  type. The type is AaaAuthentication.
    AuthenticationType interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}

    // Client data.
    Clients Aaa_Radius_DynamicAuthorization_Clients
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetFilter() yfilter.YFilter { return dynamicAuthorization.YFilter }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) SetFilter(yf yfilter.YFilter) { dynamicAuthorization.YFilter = yf }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetGoName(yname string) string {
    if yname == "ignore" { return "Ignore" }
    if yname == "port" { return "Port" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "server-key" { return "ServerKey" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetSegmentPath() string {
    return "dynamic-authorization"
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients" {
        return &dynamicAuthorization.Clients
    }
    return nil
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clients"] = &dynamicAuthorization.Clients
    return children
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ignore"] = dynamicAuthorization.Ignore
    leafs["port"] = dynamicAuthorization.Port
    leafs["authentication-type"] = dynamicAuthorization.AuthenticationType
    leafs["server-key"] = dynamicAuthorization.ServerKey
    return leafs
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetYangName() string { return "dynamic-authorization" }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) SetParent(parent types.Entity) { dynamicAuthorization.parent = parent }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetParent() types.Entity { return dynamicAuthorization.parent }

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetParentYangName() string { return "radius" }

// Aaa_Radius_DynamicAuthorization_Clients
// Client data
type Aaa_Radius_DynamicAuthorization_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client data. The type is slice of
    // Aaa_Radius_DynamicAuthorization_Clients_Client.
    Client []Aaa_Radius_DynamicAuthorization_Clients_Client

    // Client data. The type is slice of
    // Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName.
    ClientVrfName []Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    if yname == "client-vrf-name" { return "ClientVrfName" }
    return ""
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_DynamicAuthorization_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    if childYangName == "client-vrf-name" {
        for _, c := range clients.ClientVrfName {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName{}
        clients.ClientVrfName = append(clients.ClientVrfName, child)
        return &clients.ClientVrfName[len(clients.ClientVrfName)-1]
    }
    return nil
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    for i := range clients.ClientVrfName {
        children[clients.ClientVrfName[i].GetSegmentPath()] = &clients.ClientVrfName[i]
    }
    return children
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetYangName() string { return "clients" }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetParentYangName() string { return "dynamic-authorization" }

// Aaa_Radius_DynamicAuthorization_Clients_Client
// Client data
type Aaa_Radius_DynamicAuthorization_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of COA client. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "server-key" { return "ServerKey" }
    return ""
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetSegmentPath() string {
    return "client" + "[ip-address='" + fmt.Sprintf("%v", client.IpAddress) + "']"
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = client.IpAddress
    leafs["server-key"] = client.ServerKey
    return leafs
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetYangName() string { return "client" }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetParentYangName() string { return "clients" }

// Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName
// Client data
type Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // This attribute is a key. IP address of COA client. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetFilter() yfilter.YFilter { return clientVrfName.YFilter }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) SetFilter(yf yfilter.YFilter) { clientVrfName.YFilter = yf }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "server-key" { return "ServerKey" }
    return ""
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetSegmentPath() string {
    return "client-vrf-name" + "[vrf-name='" + fmt.Sprintf("%v", clientVrfName.VrfName) + "']" + "[ip-address='" + fmt.Sprintf("%v", clientVrfName.IpAddress) + "']"
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = clientVrfName.VrfName
    leafs["ip-address"] = clientVrfName.IpAddress
    leafs["server-key"] = clientVrfName.ServerKey
    return leafs
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetBundleName() string { return "cisco_ios_xr" }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetYangName() string { return "client-vrf-name" }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) SetParent(parent types.Entity) { clientVrfName.parent = parent }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetParent() types.Entity { return clientVrfName.parent }

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetParentYangName() string { return "clients" }

// Aaa_Radius_LoadBalanceOptions
// Radius load-balancing options
type Aaa_Radius_LoadBalanceOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Method by which the next host will be picked.
    LoadBalanceMethod Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetFilter() yfilter.YFilter { return loadBalanceOptions.YFilter }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) SetFilter(yf yfilter.YFilter) { loadBalanceOptions.YFilter = yf }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetGoName(yname string) string {
    if yname == "load-balance-method" { return "LoadBalanceMethod" }
    return ""
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetSegmentPath() string {
    return "load-balance-options"
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-balance-method" {
        return &loadBalanceOptions.LoadBalanceMethod
    }
    return nil
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["load-balance-method"] = &loadBalanceOptions.LoadBalanceMethod
    return children
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetBundleName() string { return "cisco_ios_xr" }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetYangName() string { return "load-balance-options" }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) SetParent(parent types.Entity) { loadBalanceOptions.parent = parent }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetParent() types.Entity { return loadBalanceOptions.parent }

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetParentYangName() string { return "radius" }

// Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod
// Method by which the next host will be picked
type Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Batch size for selection of the server.
    BatchSize Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetFilter() yfilter.YFilter { return loadBalanceMethod.YFilter }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) SetFilter(yf yfilter.YFilter) { loadBalanceMethod.YFilter = yf }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetGoName(yname string) string {
    if yname == "batch-size" { return "BatchSize" }
    return ""
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetSegmentPath() string {
    return "load-balance-method"
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "batch-size" {
        return &loadBalanceMethod.BatchSize
    }
    return nil
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["batch-size"] = &loadBalanceMethod.BatchSize
    return children
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetBundleName() string { return "cisco_ios_xr" }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetYangName() string { return "load-balance-method" }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) SetParent(parent types.Entity) { loadBalanceMethod.parent = parent }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetParent() types.Entity { return loadBalanceMethod.parent }

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetParentYangName() string { return "load-balance-options" }

// Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize
// Batch size for selection of the server
type Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Batch size for selection of the server. The type is interface{} with range:
    // 1..1500. The default value is 25.
    BatchSize interface{}

    // Disable preferred server for this Server Group. The type is bool. The
    // default value is true.
    IgnorePreferredServer interface{}
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetFilter() yfilter.YFilter { return batchSize.YFilter }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) SetFilter(yf yfilter.YFilter) { batchSize.YFilter = yf }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetGoName(yname string) string {
    if yname == "batch-size" { return "BatchSize" }
    if yname == "ignore-preferred-server" { return "IgnorePreferredServer" }
    return ""
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetSegmentPath() string {
    return "batch-size"
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["batch-size"] = batchSize.BatchSize
    leafs["ignore-preferred-server"] = batchSize.IgnorePreferredServer
    return leafs
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetBundleName() string { return "cisco_ios_xr" }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetYangName() string { return "batch-size" }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) SetParent(parent types.Entity) { batchSize.parent = parent }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetParent() types.Entity { return batchSize.parent }

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetParentYangName() string { return "load-balance-method" }

// Aaa_Radius_Vrfs
// List of VRFs
type Aaa_Radius_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRF. The type is slice of Aaa_Radius_Vrfs_Vrf.
    Vrf []Aaa_Radius_Vrfs_Vrf
}

func (vrfs *Aaa_Radius_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Aaa_Radius_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Aaa_Radius_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Aaa_Radius_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Aaa_Radius_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Aaa_Radius_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Aaa_Radius_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Aaa_Radius_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Aaa_Radius_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Aaa_Radius_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Aaa_Radius_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Aaa_Radius_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Aaa_Radius_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Aaa_Radius_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Aaa_Radius_Vrfs) GetParentYangName() string { return "radius" }

// Aaa_Radius_Vrfs_Vrf
// A VRF
type Aaa_Radius_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. Specify 'default' for defalut VRF. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in RADIUS packets. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Aaa_Radius_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["source-interface"] = vrf.SourceInterface
    return leafs
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Aaa_Radius_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Aaa_Radius_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Aaa_Radius_Throttle
// Radius throttling options
type Aaa_Radius_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // To flow control the number of access requests sent to a radius server. The
    // type is interface{} with range: 0..65535. The default value is 0.
    Access interface{}

    // To flow control the number of accounting requests sent to a radius server.
    // The type is interface{} with range: 0..65535. The default value is 0.
    Accounting interface{}

    // Specify the number of timeouts exceeding which a throttled access request
    // is dropped. The type is interface{} with range: 1..10. The default value is
    // 3.
    AccessTimeout interface{}
}

func (throttle *Aaa_Radius_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *Aaa_Radius_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *Aaa_Radius_Throttle) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    if yname == "accounting" { return "Accounting" }
    if yname == "access-timeout" { return "AccessTimeout" }
    return ""
}

func (throttle *Aaa_Radius_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *Aaa_Radius_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *Aaa_Radius_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *Aaa_Radius_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access"] = throttle.Access
    leafs["accounting"] = throttle.Accounting
    leafs["access-timeout"] = throttle.AccessTimeout
    return leafs
}

func (throttle *Aaa_Radius_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *Aaa_Radius_Throttle) GetYangName() string { return "throttle" }

func (throttle *Aaa_Radius_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *Aaa_Radius_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *Aaa_Radius_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *Aaa_Radius_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *Aaa_Radius_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *Aaa_Radius_Throttle) GetParentYangName() string { return "radius" }

// Aaa_Radius_Vsa
// Unknown VSA (Vendor Specific Attribute) ignore
// configuration for RADIUS server
type Aaa_Radius_Vsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Vendor Specific Attribute.
    Attribute Aaa_Radius_Vsa_Attribute
}

func (vsa *Aaa_Radius_Vsa) GetFilter() yfilter.YFilter { return vsa.YFilter }

func (vsa *Aaa_Radius_Vsa) SetFilter(yf yfilter.YFilter) { vsa.YFilter = yf }

func (vsa *Aaa_Radius_Vsa) GetGoName(yname string) string {
    if yname == "attribute" { return "Attribute" }
    return ""
}

func (vsa *Aaa_Radius_Vsa) GetSegmentPath() string {
    return "vsa"
}

func (vsa *Aaa_Radius_Vsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attribute" {
        return &vsa.Attribute
    }
    return nil
}

func (vsa *Aaa_Radius_Vsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attribute"] = &vsa.Attribute
    return children
}

func (vsa *Aaa_Radius_Vsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vsa *Aaa_Radius_Vsa) GetBundleName() string { return "cisco_ios_xr" }

func (vsa *Aaa_Radius_Vsa) GetYangName() string { return "vsa" }

func (vsa *Aaa_Radius_Vsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vsa *Aaa_Radius_Vsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vsa *Aaa_Radius_Vsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vsa *Aaa_Radius_Vsa) SetParent(parent types.Entity) { vsa.parent = parent }

func (vsa *Aaa_Radius_Vsa) GetParent() types.Entity { return vsa.parent }

func (vsa *Aaa_Radius_Vsa) GetParentYangName() string { return "radius" }

// Aaa_Radius_Vsa_Attribute
// Vendor Specific Attribute
type Aaa_Radius_Vsa_Attribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore the VSA.
    Ignore Aaa_Radius_Vsa_Attribute_Ignore
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetFilter() yfilter.YFilter { return attribute.YFilter }

func (attribute *Aaa_Radius_Vsa_Attribute) SetFilter(yf yfilter.YFilter) { attribute.YFilter = yf }

func (attribute *Aaa_Radius_Vsa_Attribute) GetGoName(yname string) string {
    if yname == "ignore" { return "Ignore" }
    return ""
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetSegmentPath() string {
    return "attribute"
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ignore" {
        return &attribute.Ignore
    }
    return nil
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ignore"] = &attribute.Ignore
    return children
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetBundleName() string { return "cisco_ios_xr" }

func (attribute *Aaa_Radius_Vsa_Attribute) GetYangName() string { return "attribute" }

func (attribute *Aaa_Radius_Vsa_Attribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attribute *Aaa_Radius_Vsa_Attribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attribute *Aaa_Radius_Vsa_Attribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attribute *Aaa_Radius_Vsa_Attribute) SetParent(parent types.Entity) { attribute.parent = parent }

func (attribute *Aaa_Radius_Vsa_Attribute) GetParent() types.Entity { return attribute.parent }

func (attribute *Aaa_Radius_Vsa_Attribute) GetParentYangName() string { return "vsa" }

// Aaa_Radius_Vsa_Attribute_Ignore
// Ignore the VSA
type Aaa_Radius_Vsa_Attribute_Ignore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore the VSA and no prefix will reject the unknown VSA. The type is
    // interface{}.
    Unknown interface{}
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetFilter() yfilter.YFilter { return ignore.YFilter }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) SetFilter(yf yfilter.YFilter) { ignore.YFilter = yf }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetGoName(yname string) string {
    if yname == "unknown" { return "Unknown" }
    return ""
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetSegmentPath() string {
    return "ignore"
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown"] = ignore.Unknown
    return leafs
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetBundleName() string { return "cisco_ios_xr" }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetYangName() string { return "ignore" }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) SetParent(parent types.Entity) { ignore.parent = parent }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetParent() types.Entity { return ignore.parent }

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetParentYangName() string { return "attribute" }

// Aaa_Radius_Ipv4
// IPv4 configuration
type Aaa_Radius_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // AaaDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Aaa_Radius_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Aaa_Radius_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Aaa_Radius_Ipv4) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (ipv4 *Aaa_Radius_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Aaa_Radius_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Aaa_Radius_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Aaa_Radius_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ipv4.Dscp
    return leafs
}

func (ipv4 *Aaa_Radius_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Aaa_Radius_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Aaa_Radius_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Aaa_Radius_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Aaa_Radius_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Aaa_Radius_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Aaa_Radius_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Aaa_Radius_Ipv4) GetParentYangName() string { return "radius" }

// Aaa_Radius_RadiusAttribute
// attribute
type Aaa_Radius_RadiusAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Filter-Id attribute configuration.
    FilterId11 Aaa_Radius_RadiusAttribute_FilterId11

    // Acct-Session-Id attribute(44).
    AcctMultiSessionId Aaa_Radius_RadiusAttribute_AcctMultiSessionId

    // Acct-Session-Id attribute(44).
    AcctSessionId Aaa_Radius_RadiusAttribute_AcctSessionId
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetFilter() yfilter.YFilter { return radiusAttribute.YFilter }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) SetFilter(yf yfilter.YFilter) { radiusAttribute.YFilter = yf }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetGoName(yname string) string {
    if yname == "filter-id-11" { return "FilterId11" }
    if yname == "acct-multi-session-id" { return "AcctMultiSessionId" }
    if yname == "acct-session-id" { return "AcctSessionId" }
    return ""
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetSegmentPath() string {
    return "radius-attribute"
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "filter-id-11" {
        return &radiusAttribute.FilterId11
    }
    if childYangName == "acct-multi-session-id" {
        return &radiusAttribute.AcctMultiSessionId
    }
    if childYangName == "acct-session-id" {
        return &radiusAttribute.AcctSessionId
    }
    return nil
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["filter-id-11"] = &radiusAttribute.FilterId11
    children["acct-multi-session-id"] = &radiusAttribute.AcctMultiSessionId
    children["acct-session-id"] = &radiusAttribute.AcctSessionId
    return children
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetBundleName() string { return "cisco_ios_xr" }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetYangName() string { return "radius-attribute" }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) SetParent(parent types.Entity) { radiusAttribute.parent = parent }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetParent() types.Entity { return radiusAttribute.parent }

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetParentYangName() string { return "radius" }

// Aaa_Radius_RadiusAttribute_FilterId11
// Filter-Id attribute configuration
type Aaa_Radius_RadiusAttribute_FilterId11 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the attribute default direction.
    Defaults Aaa_Radius_RadiusAttribute_FilterId11_Defaults
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetFilter() yfilter.YFilter { return filterId11.YFilter }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) SetFilter(yf yfilter.YFilter) { filterId11.YFilter = yf }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetGoName(yname string) string {
    if yname == "defaults" { return "Defaults" }
    return ""
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetSegmentPath() string {
    return "filter-id-11"
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "defaults" {
        return &filterId11.Defaults
    }
    return nil
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["defaults"] = &filterId11.Defaults
    return children
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetBundleName() string { return "cisco_ios_xr" }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetYangName() string { return "filter-id-11" }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) SetParent(parent types.Entity) { filterId11.parent = parent }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetParent() types.Entity { return filterId11.parent }

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetParentYangName() string { return "radius-attribute" }

// Aaa_Radius_RadiusAttribute_FilterId11_Defaults
// Set the attribute default direction
type Aaa_Radius_RadiusAttribute_FilterId11_Defaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Filtering is applied to ingress(inbound)/egress(outbound) packets only. The
    // type is AaaDirection.
    Direction interface{}
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetFilter() yfilter.YFilter { return defaults.YFilter }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) SetFilter(yf yfilter.YFilter) { defaults.YFilter = yf }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    return ""
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetSegmentPath() string {
    return "defaults"
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = defaults.Direction
    return leafs
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetBundleName() string { return "cisco_ios_xr" }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetYangName() string { return "defaults" }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) SetParent(parent types.Entity) { defaults.parent = parent }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetParent() types.Entity { return defaults.parent }

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetParentYangName() string { return "filter-id-11" }

// Aaa_Radius_RadiusAttribute_AcctMultiSessionId
// Acct-Session-Id attribute(44)
type Aaa_Radius_RadiusAttribute_AcctMultiSessionId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend Acct-Session-Id attribute with Nas-Port-Id.
    IncludeParentSessionId Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetFilter() yfilter.YFilter { return acctMultiSessionId.YFilter }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) SetFilter(yf yfilter.YFilter) { acctMultiSessionId.YFilter = yf }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetGoName(yname string) string {
    if yname == "include-parent-session-id" { return "IncludeParentSessionId" }
    return ""
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetSegmentPath() string {
    return "acct-multi-session-id"
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "include-parent-session-id" {
        return &acctMultiSessionId.IncludeParentSessionId
    }
    return nil
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["include-parent-session-id"] = &acctMultiSessionId.IncludeParentSessionId
    return children
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetBundleName() string { return "cisco_ios_xr" }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetYangName() string { return "acct-multi-session-id" }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) SetParent(parent types.Entity) { acctMultiSessionId.parent = parent }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetParent() types.Entity { return acctMultiSessionId.parent }

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetParentYangName() string { return "radius-attribute" }

// Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId
// Prepend Acct-Session-Id attribute with
// Nas-Port-Id
type Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // false/true. The type is AaaConfig.
    Config interface{}
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetFilter() yfilter.YFilter { return includeParentSessionId.YFilter }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) SetFilter(yf yfilter.YFilter) { includeParentSessionId.YFilter = yf }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    return ""
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetSegmentPath() string {
    return "include-parent-session-id"
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config"] = includeParentSessionId.Config
    return leafs
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetBundleName() string { return "cisco_ios_xr" }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetYangName() string { return "include-parent-session-id" }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) SetParent(parent types.Entity) { includeParentSessionId.parent = parent }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetParent() types.Entity { return includeParentSessionId.parent }

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetParentYangName() string { return "acct-multi-session-id" }

// Aaa_Radius_RadiusAttribute_AcctSessionId
// Acct-Session-Id attribute(44)
type Aaa_Radius_RadiusAttribute_AcctSessionId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend Acct-Session-Id attribute with Nas-Port-Id.
    PrependNasPortId Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetFilter() yfilter.YFilter { return acctSessionId.YFilter }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) SetFilter(yf yfilter.YFilter) { acctSessionId.YFilter = yf }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetGoName(yname string) string {
    if yname == "prepend-nas-port-id" { return "PrependNasPortId" }
    return ""
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetSegmentPath() string {
    return "acct-session-id"
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-nas-port-id" {
        return &acctSessionId.PrependNasPortId
    }
    return nil
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prepend-nas-port-id"] = &acctSessionId.PrependNasPortId
    return children
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetBundleName() string { return "cisco_ios_xr" }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetYangName() string { return "acct-session-id" }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) SetParent(parent types.Entity) { acctSessionId.parent = parent }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetParent() types.Entity { return acctSessionId.parent }

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetParentYangName() string { return "radius-attribute" }

// Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId
// Prepend Acct-Session-Id attribute with
// Nas-Port-Id
type Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // false/true. The type is AaaConfig.
    Config interface{}
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetFilter() yfilter.YFilter { return prependNasPortId.YFilter }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) SetFilter(yf yfilter.YFilter) { prependNasPortId.YFilter = yf }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    return ""
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetSegmentPath() string {
    return "prepend-nas-port-id"
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config"] = prependNasPortId.Config
    return leafs
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetBundleName() string { return "cisco_ios_xr" }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetYangName() string { return "prepend-nas-port-id" }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) SetParent(parent types.Entity) { prependNasPortId.parent = parent }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetParent() types.Entity { return prependNasPortId.parent }

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetParentYangName() string { return "acct-session-id" }

// Aaa_Radius_Attributes
// Table of attribute list
type Aaa_Radius_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Attribute list name. The type is slice of Aaa_Radius_Attributes_Attribute.
    Attribute []Aaa_Radius_Attributes_Attribute
}

func (attributes *Aaa_Radius_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *Aaa_Radius_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *Aaa_Radius_Attributes) GetGoName(yname string) string {
    if yname == "attribute" { return "Attribute" }
    return ""
}

func (attributes *Aaa_Radius_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *Aaa_Radius_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attribute" {
        for _, c := range attributes.Attribute {
            if attributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute{}
        attributes.Attribute = append(attributes.Attribute, child)
        return &attributes.Attribute[len(attributes.Attribute)-1]
    }
    return nil
}

func (attributes *Aaa_Radius_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attributes.Attribute {
        children[attributes.Attribute[i].GetSegmentPath()] = &attributes.Attribute[i]
    }
    return children
}

func (attributes *Aaa_Radius_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *Aaa_Radius_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *Aaa_Radius_Attributes) GetYangName() string { return "attributes" }

func (attributes *Aaa_Radius_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *Aaa_Radius_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *Aaa_Radius_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *Aaa_Radius_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *Aaa_Radius_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *Aaa_Radius_Attributes) GetParentYangName() string { return "radius" }

// Aaa_Radius_Attributes_Attribute
// Attribute list name
type Aaa_Radius_Attributes_Attribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute list name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AttributeListName interface{}

    // Specify RADIUS attribute. The type is string.
    Attribute interface{}

    // Vendor Specific Attribute.
    VendorIds Aaa_Radius_Attributes_Attribute_VendorIds
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetFilter() yfilter.YFilter { return attribute.YFilter }

func (attribute *Aaa_Radius_Attributes_Attribute) SetFilter(yf yfilter.YFilter) { attribute.YFilter = yf }

func (attribute *Aaa_Radius_Attributes_Attribute) GetGoName(yname string) string {
    if yname == "attribute-list-name" { return "AttributeListName" }
    if yname == "attribute" { return "Attribute" }
    if yname == "vendor-ids" { return "VendorIds" }
    return ""
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetSegmentPath() string {
    return "attribute" + "[attribute-list-name='" + fmt.Sprintf("%v", attribute.AttributeListName) + "']"
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vendor-ids" {
        return &attribute.VendorIds
    }
    return nil
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vendor-ids"] = &attribute.VendorIds
    return children
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-list-name"] = attribute.AttributeListName
    leafs["attribute"] = attribute.Attribute
    return leafs
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetBundleName() string { return "cisco_ios_xr" }

func (attribute *Aaa_Radius_Attributes_Attribute) GetYangName() string { return "attribute" }

func (attribute *Aaa_Radius_Attributes_Attribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attribute *Aaa_Radius_Attributes_Attribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attribute *Aaa_Radius_Attributes_Attribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attribute *Aaa_Radius_Attributes_Attribute) SetParent(parent types.Entity) { attribute.parent = parent }

func (attribute *Aaa_Radius_Attributes_Attribute) GetParent() types.Entity { return attribute.parent }

func (attribute *Aaa_Radius_Attributes_Attribute) GetParentYangName() string { return "attributes" }

// Aaa_Radius_Attributes_Attribute_VendorIds
// Vendor Specific Attribute
type Aaa_Radius_Attributes_Attribute_VendorIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Vendor ID of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId.
    VendorId []Aaa_Radius_Attributes_Attribute_VendorIds_VendorId
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetFilter() yfilter.YFilter { return vendorIds.YFilter }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) SetFilter(yf yfilter.YFilter) { vendorIds.YFilter = yf }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetGoName(yname string) string {
    if yname == "vendor-id" { return "VendorId" }
    return ""
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetSegmentPath() string {
    return "vendor-ids"
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vendor-id" {
        for _, c := range vendorIds.VendorId {
            if vendorIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute_VendorIds_VendorId{}
        vendorIds.VendorId = append(vendorIds.VendorId, child)
        return &vendorIds.VendorId[len(vendorIds.VendorId)-1]
    }
    return nil
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vendorIds.VendorId {
        children[vendorIds.VendorId[i].GetSegmentPath()] = &vendorIds.VendorId[i]
    }
    return children
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetBundleName() string { return "cisco_ios_xr" }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetYangName() string { return "vendor-ids" }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) SetParent(parent types.Entity) { vendorIds.parent = parent }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetParent() types.Entity { return vendorIds.parent }

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetParentYangName() string { return "attribute" }

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId
// Vendor ID of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Vendor Id of vsa. The type is interface{} with
    // range: -2147483648..2147483647.
    VendorId interface{}

    // Vendor Type of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType.
    VendorType []Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetFilter() yfilter.YFilter { return vendorId.YFilter }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) SetFilter(yf yfilter.YFilter) { vendorId.YFilter = yf }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetGoName(yname string) string {
    if yname == "vendor-id" { return "VendorId" }
    if yname == "vendor-type" { return "VendorType" }
    return ""
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetSegmentPath() string {
    return "vendor-id" + "[vendor-id='" + fmt.Sprintf("%v", vendorId.VendorId) + "']"
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vendor-type" {
        for _, c := range vendorId.VendorType {
            if vendorId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType{}
        vendorId.VendorType = append(vendorId.VendorType, child)
        return &vendorId.VendorType[len(vendorId.VendorType)-1]
    }
    return nil
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vendorId.VendorType {
        children[vendorId.VendorType[i].GetSegmentPath()] = &vendorId.VendorType[i]
    }
    return children
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vendor-id"] = vendorId.VendorId
    return leafs
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetBundleName() string { return "cisco_ios_xr" }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetYangName() string { return "vendor-id" }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) SetParent(parent types.Entity) { vendorId.parent = parent }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetParent() types.Entity { return vendorId.parent }

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetParentYangName() string { return "vendor-ids" }

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType
// Vendor Type of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Vendor Type of vsa. The type is interface{} with
    // range: -2147483648..2147483647.
    VendorType interface{}

    // Attribute Name of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName.
    AttributeName []Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetFilter() yfilter.YFilter { return vendorType.YFilter }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) SetFilter(yf yfilter.YFilter) { vendorType.YFilter = yf }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetGoName(yname string) string {
    if yname == "vendor-type" { return "VendorType" }
    if yname == "attribute-name" { return "AttributeName" }
    return ""
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetSegmentPath() string {
    return "vendor-type" + "[vendor-type='" + fmt.Sprintf("%v", vendorType.VendorType) + "']"
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attribute-name" {
        for _, c := range vendorType.AttributeName {
            if vendorType.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName{}
        vendorType.AttributeName = append(vendorType.AttributeName, child)
        return &vendorType.AttributeName[len(vendorType.AttributeName)-1]
    }
    return nil
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vendorType.AttributeName {
        children[vendorType.AttributeName[i].GetSegmentPath()] = &vendorType.AttributeName[i]
    }
    return children
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vendor-type"] = vendorType.VendorType
    return leafs
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetBundleName() string { return "cisco_ios_xr" }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetYangName() string { return "vendor-type" }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) SetParent(parent types.Entity) { vendorType.parent = parent }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetParent() types.Entity { return vendorType.parent }

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetParentYangName() string { return "vendor-id" }

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName
// Attribute Name of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Name of vsa. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AttributeName interface{}

    // AttributeName of vsa is absent. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent.
    AttributeNameAbsent []Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetFilter() yfilter.YFilter { return attributeName.YFilter }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) SetFilter(yf yfilter.YFilter) { attributeName.YFilter = yf }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetGoName(yname string) string {
    if yname == "attribute-name" { return "AttributeName" }
    if yname == "attribute-name-absent" { return "AttributeNameAbsent" }
    return ""
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetSegmentPath() string {
    return "attribute-name" + "[attribute-name='" + fmt.Sprintf("%v", attributeName.AttributeName) + "']"
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attribute-name-absent" {
        for _, c := range attributeName.AttributeNameAbsent {
            if attributeName.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent{}
        attributeName.AttributeNameAbsent = append(attributeName.AttributeNameAbsent, child)
        return &attributeName.AttributeNameAbsent[len(attributeName.AttributeNameAbsent)-1]
    }
    return nil
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attributeName.AttributeNameAbsent {
        children[attributeName.AttributeNameAbsent[i].GetSegmentPath()] = &attributeName.AttributeNameAbsent[i]
    }
    return children
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-name"] = attributeName.AttributeName
    return leafs
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetBundleName() string { return "cisco_ios_xr" }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetYangName() string { return "attribute-name" }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) SetParent(parent types.Entity) { attributeName.parent = parent }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetParent() types.Entity { return attributeName.parent }

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetParentYangName() string { return "vendor-type" }

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent
// AttributeName of vsa is absent
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AttributeName of vsa is absent. The type is
    // interface{} with range: -2147483648..2147483647.
    AttributeNameAbsent interface{}

    // AttributeName of vsa is present. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent.
    AttributeNamePresent []Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetFilter() yfilter.YFilter { return attributeNameAbsent.YFilter }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) SetFilter(yf yfilter.YFilter) { attributeNameAbsent.YFilter = yf }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetGoName(yname string) string {
    if yname == "attribute-name-absent" { return "AttributeNameAbsent" }
    if yname == "attribute-name-present" { return "AttributeNamePresent" }
    return ""
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetSegmentPath() string {
    return "attribute-name-absent" + "[attribute-name-absent='" + fmt.Sprintf("%v", attributeNameAbsent.AttributeNameAbsent) + "']"
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attribute-name-present" {
        for _, c := range attributeNameAbsent.AttributeNamePresent {
            if attributeNameAbsent.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent{}
        attributeNameAbsent.AttributeNamePresent = append(attributeNameAbsent.AttributeNamePresent, child)
        return &attributeNameAbsent.AttributeNamePresent[len(attributeNameAbsent.AttributeNamePresent)-1]
    }
    return nil
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attributeNameAbsent.AttributeNamePresent {
        children[attributeNameAbsent.AttributeNamePresent[i].GetSegmentPath()] = &attributeNameAbsent.AttributeNamePresent[i]
    }
    return children
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-name-absent"] = attributeNameAbsent.AttributeNameAbsent
    return leafs
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetBundleName() string { return "cisco_ios_xr" }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetYangName() string { return "attribute-name-absent" }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) SetParent(parent types.Entity) { attributeNameAbsent.parent = parent }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetParent() types.Entity { return attributeNameAbsent.parent }

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetParentYangName() string { return "attribute-name" }

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent
// AttributeName of vsa is present
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AttributeName of vsa is present. The type is
    // interface{} with range: -2147483648..2147483647.
    AttributeNamePresent interface{}
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetFilter() yfilter.YFilter { return attributeNamePresent.YFilter }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) SetFilter(yf yfilter.YFilter) { attributeNamePresent.YFilter = yf }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetGoName(yname string) string {
    if yname == "attribute-name-present" { return "AttributeNamePresent" }
    return ""
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetSegmentPath() string {
    return "attribute-name-present" + "[attribute-name-present='" + fmt.Sprintf("%v", attributeNamePresent.AttributeNamePresent) + "']"
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-name-present"] = attributeNamePresent.AttributeNamePresent
    return leafs
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetBundleName() string { return "cisco_ios_xr" }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetYangName() string { return "attribute-name-present" }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) SetParent(parent types.Entity) { attributeNamePresent.parent = parent }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetParent() types.Entity { return attributeNamePresent.parent }

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetParentYangName() string { return "attribute-name-absent" }

// Aaa_Radius_SourcePort
// Source port
type Aaa_Radius_SourcePort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable source-port extend . The type is interface{}.
    Extended interface{}
}

func (sourcePort *Aaa_Radius_SourcePort) GetFilter() yfilter.YFilter { return sourcePort.YFilter }

func (sourcePort *Aaa_Radius_SourcePort) SetFilter(yf yfilter.YFilter) { sourcePort.YFilter = yf }

func (sourcePort *Aaa_Radius_SourcePort) GetGoName(yname string) string {
    if yname == "extended" { return "Extended" }
    return ""
}

func (sourcePort *Aaa_Radius_SourcePort) GetSegmentPath() string {
    return "source-port"
}

func (sourcePort *Aaa_Radius_SourcePort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourcePort *Aaa_Radius_SourcePort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourcePort *Aaa_Radius_SourcePort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["extended"] = sourcePort.Extended
    return leafs
}

func (sourcePort *Aaa_Radius_SourcePort) GetBundleName() string { return "cisco_ios_xr" }

func (sourcePort *Aaa_Radius_SourcePort) GetYangName() string { return "source-port" }

func (sourcePort *Aaa_Radius_SourcePort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourcePort *Aaa_Radius_SourcePort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourcePort *Aaa_Radius_SourcePort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourcePort *Aaa_Radius_SourcePort) SetParent(parent types.Entity) { sourcePort.parent = parent }

func (sourcePort *Aaa_Radius_SourcePort) GetParent() types.Entity { return sourcePort.parent }

func (sourcePort *Aaa_Radius_SourcePort) GetParentYangName() string { return "radius" }

// Aaa_Tacacs
// Modify TACACS+ query parameters
type Aaa_Tacacs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Use a single connection for all sessions for a given TACACS+ server. The
    // type is bool. The default value is false.
    SingleConnect interface{}

    // IPv6 configuration.
    Ipv6 Aaa_Tacacs_Ipv6

    // Specify a TACACS+ server.
    Hosts Aaa_Tacacs_Hosts

    // IPv4 configuration.
    Ipv4 Aaa_Tacacs_Ipv4

    // List of VRFs.
    Vrfs Aaa_Tacacs_Vrfs
}

func (tacacs *Aaa_Tacacs) GetFilter() yfilter.YFilter { return tacacs.YFilter }

func (tacacs *Aaa_Tacacs) SetFilter(yf yfilter.YFilter) { tacacs.YFilter = yf }

func (tacacs *Aaa_Tacacs) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    if yname == "timeout" { return "Timeout" }
    if yname == "single-connect" { return "SingleConnect" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "hosts" { return "Hosts" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (tacacs *Aaa_Tacacs) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs"
}

func (tacacs *Aaa_Tacacs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &tacacs.Ipv6
    }
    if childYangName == "hosts" {
        return &tacacs.Hosts
    }
    if childYangName == "ipv4" {
        return &tacacs.Ipv4
    }
    if childYangName == "vrfs" {
        return &tacacs.Vrfs
    }
    return nil
}

func (tacacs *Aaa_Tacacs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &tacacs.Ipv6
    children["hosts"] = &tacacs.Hosts
    children["ipv4"] = &tacacs.Ipv4
    children["vrfs"] = &tacacs.Vrfs
    return children
}

func (tacacs *Aaa_Tacacs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key"] = tacacs.Key
    leafs["timeout"] = tacacs.Timeout
    leafs["single-connect"] = tacacs.SingleConnect
    return leafs
}

func (tacacs *Aaa_Tacacs) GetBundleName() string { return "cisco_ios_xr" }

func (tacacs *Aaa_Tacacs) GetYangName() string { return "tacacs" }

func (tacacs *Aaa_Tacacs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tacacs *Aaa_Tacacs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tacacs *Aaa_Tacacs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tacacs *Aaa_Tacacs) SetParent(parent types.Entity) { tacacs.parent = parent }

func (tacacs *Aaa_Tacacs) GetParent() types.Entity { return tacacs.parent }

func (tacacs *Aaa_Tacacs) GetParentYangName() string { return "aaa" }

// Aaa_Tacacs_Ipv6
// IPv6 configuration
type Aaa_Tacacs_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // TacacsDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Aaa_Tacacs_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Aaa_Tacacs_Ipv6) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ipv6.Dscp
    return leafs
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Aaa_Tacacs_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Aaa_Tacacs_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Aaa_Tacacs_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Aaa_Tacacs_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Aaa_Tacacs_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Aaa_Tacacs_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Aaa_Tacacs_Ipv6) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Hosts
// Specify a TACACS+ server
type Aaa_Tacacs_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the TACACS+ servers. The type is slice of Aaa_Tacacs_Hosts_Host.
    Host []Aaa_Tacacs_Hosts_Host
}

func (hosts *Aaa_Tacacs_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Aaa_Tacacs_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Aaa_Tacacs_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Aaa_Tacacs_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Aaa_Tacacs_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Aaa_Tacacs_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Aaa_Tacacs_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Aaa_Tacacs_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Aaa_Tacacs_Hosts) GetYangName() string { return "hosts" }

func (hosts *Aaa_Tacacs_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Aaa_Tacacs_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Aaa_Tacacs_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Aaa_Tacacs_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Aaa_Tacacs_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Aaa_Tacacs_Hosts) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Hosts_Host
// One of the TACACS+ servers
type Aaa_Tacacs_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port number (standard 49). The type is interface{}
    // with range: 1..65535.
    PortNumber interface{}

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Use a single connection for all sessions for a given TACACS+ server. The
    // type is bool. The default value is false.
    SingleConnect interface{}
}

func (host *Aaa_Tacacs_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Aaa_Tacacs_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Aaa_Tacacs_Hosts_Host) GetGoName(yname string) string {
    if yname == "ordering-index" { return "OrderingIndex" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "key" { return "Key" }
    if yname == "timeout" { return "Timeout" }
    if yname == "single-connect" { return "SingleConnect" }
    return ""
}

func (host *Aaa_Tacacs_Hosts_Host) GetSegmentPath() string {
    return "host" + "[ordering-index='" + fmt.Sprintf("%v", host.OrderingIndex) + "']" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[port-number='" + fmt.Sprintf("%v", host.PortNumber) + "']"
}

func (host *Aaa_Tacacs_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *Aaa_Tacacs_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *Aaa_Tacacs_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ordering-index"] = host.OrderingIndex
    leafs["ip-address"] = host.IpAddress
    leafs["port-number"] = host.PortNumber
    leafs["key"] = host.Key
    leafs["timeout"] = host.Timeout
    leafs["single-connect"] = host.SingleConnect
    return leafs
}

func (host *Aaa_Tacacs_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Aaa_Tacacs_Hosts_Host) GetYangName() string { return "host" }

func (host *Aaa_Tacacs_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Aaa_Tacacs_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Aaa_Tacacs_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Aaa_Tacacs_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Aaa_Tacacs_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Aaa_Tacacs_Hosts_Host) GetParentYangName() string { return "hosts" }

// Aaa_Tacacs_Ipv4
// IPv4 configuration
type Aaa_Tacacs_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // TacacsDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Aaa_Tacacs_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Aaa_Tacacs_Ipv4) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ipv4.Dscp
    return leafs
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Aaa_Tacacs_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Aaa_Tacacs_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Aaa_Tacacs_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Aaa_Tacacs_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Aaa_Tacacs_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Aaa_Tacacs_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Aaa_Tacacs_Ipv4) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Vrfs
// List of VRFs
type Aaa_Tacacs_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRF. The type is slice of Aaa_Tacacs_Vrfs_Vrf.
    Vrf []Aaa_Tacacs_Vrfs_Vrf
}

func (vrfs *Aaa_Tacacs_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Aaa_Tacacs_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Aaa_Tacacs_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Aaa_Tacacs_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Aaa_Tacacs_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Aaa_Tacacs_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Aaa_Tacacs_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Aaa_Tacacs_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Aaa_Tacacs_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Aaa_Tacacs_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Aaa_Tacacs_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Aaa_Tacacs_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Aaa_Tacacs_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Aaa_Tacacs_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Aaa_Tacacs_Vrfs) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Vrfs_Vrf
// A VRF
type Aaa_Tacacs_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. Specify 'default' for default VRF. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in TACACS+ packets. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["source-interface"] = vrf.SourceInterface
    return leafs
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

