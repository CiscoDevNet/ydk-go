// NETCONF Monitoring Module.
// All elements in this module are read-only.
// 
// Copyright (c) 2010 IETF Trust and the persons identified as
// authors of the code. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD
// License set forth in Section 4.c of the IETF Trust's
// Legal Provisions Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6022; see
// the RFC itself for full legal notices.
package netconf_monitoring

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package netconf_monitoring"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring netconf-state}", reflect.TypeOf(NetconfState{}))
    ydk.RegisterEntity("ietf-netconf-monitoring:netconf-state", reflect.TypeOf(NetconfState{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring get-schema}", reflect.TypeOf(GetSchema{}))
    ydk.RegisterEntity("ietf-netconf-monitoring:get-schema", reflect.TypeOf(GetSchema{}))
}

type NetconfBeep struct {
}

func (id NetconfBeep) String() string {
	return "ietf-netconf-monitoring:netconf-beep"
}

type NetconfSsh struct {
}

func (id NetconfSsh) String() string {
	return "ietf-netconf-monitoring:netconf-ssh"
}

type Rnc struct {
}

func (id Rnc) String() string {
	return "ietf-netconf-monitoring:rnc"
}

type Yin struct {
}

func (id Yin) String() string {
	return "ietf-netconf-monitoring:yin"
}

type Rng struct {
}

func (id Rng) String() string {
	return "ietf-netconf-monitoring:rng"
}

type Xsd struct {
}

func (id Xsd) String() string {
	return "ietf-netconf-monitoring:xsd"
}

type NetconfSoapOverBeep struct {
}

func (id NetconfSoapOverBeep) String() string {
	return "ietf-netconf-monitoring:netconf-soap-over-beep"
}

type NetconfTls struct {
}

func (id NetconfTls) String() string {
	return "ietf-netconf-monitoring:netconf-tls"
}

type Yang struct {
}

func (id Yang) String() string {
	return "ietf-netconf-monitoring:yang"
}

type SchemaFormat struct {
}

func (id SchemaFormat) String() string {
	return "ietf-netconf-monitoring:schema-format"
}

type NetconfSoapOverHttps struct {
}

func (id NetconfSoapOverHttps) String() string {
	return "ietf-netconf-monitoring:netconf-soap-over-https"
}

type Transport struct {
}

func (id Transport) String() string {
	return "ietf-netconf-monitoring:transport"
}

// NetconfDatastoreType represents Enumeration of possible NETCONF datastore types.
type NetconfDatastoreType string

const (
    NetconfDatastoreType_running NetconfDatastoreType = "running"

    NetconfDatastoreType_candidate NetconfDatastoreType = "candidate"

    NetconfDatastoreType_startup NetconfDatastoreType = "startup"
)

// NetconfState
// The netconf-state container is the root of the monitoring
// data model.
type NetconfState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains the list of NETCONF capabilities supported by the server.
    Capabilities NetconfState_Capabilities

    // Contains the list of NETCONF configuration datastores.
    Datastores NetconfState_Datastores

    // Contains the list of data model schemas supported by the server.
    Schemas NetconfState_Schemas

    // The sessions container includes session-specific data for NETCONF
    // management sessions.  The session list MUST include all currently active
    // NETCONF sessions.
    Sessions NetconfState_Sessions

    // Statistical data pertaining to the NETCONF server.
    Statistics NetconfState_Statistics
}

func (netconfState *NetconfState) GetFilter() yfilter.YFilter { return netconfState.YFilter }

func (netconfState *NetconfState) SetFilter(yf yfilter.YFilter) { netconfState.YFilter = yf }

func (netconfState *NetconfState) GetGoName(yname string) string {
    if yname == "capabilities" { return "Capabilities" }
    if yname == "datastores" { return "Datastores" }
    if yname == "schemas" { return "Schemas" }
    if yname == "sessions" { return "Sessions" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (netconfState *NetconfState) GetSegmentPath() string {
    return "ietf-netconf-monitoring:netconf-state"
}

func (netconfState *NetconfState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capabilities" {
        return &netconfState.Capabilities
    }
    if childYangName == "datastores" {
        return &netconfState.Datastores
    }
    if childYangName == "schemas" {
        return &netconfState.Schemas
    }
    if childYangName == "sessions" {
        return &netconfState.Sessions
    }
    if childYangName == "statistics" {
        return &netconfState.Statistics
    }
    return nil
}

func (netconfState *NetconfState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capabilities"] = &netconfState.Capabilities
    children["datastores"] = &netconfState.Datastores
    children["schemas"] = &netconfState.Schemas
    children["sessions"] = &netconfState.Sessions
    children["statistics"] = &netconfState.Statistics
    return children
}

func (netconfState *NetconfState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconfState *NetconfState) GetBundleName() string { return "ietf" }

func (netconfState *NetconfState) GetYangName() string { return "netconf-state" }

func (netconfState *NetconfState) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (netconfState *NetconfState) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (netconfState *NetconfState) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (netconfState *NetconfState) SetParent(parent types.Entity) { netconfState.parent = parent }

func (netconfState *NetconfState) GetParent() types.Entity { return netconfState.parent }

func (netconfState *NetconfState) GetParentYangName() string { return "ietf-netconf-monitoring" }

// NetconfState_Capabilities
// Contains the list of NETCONF capabilities supported by the
// server.
type NetconfState_Capabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of NETCONF capabilities supported by the server. The type is slice of
    // string.
    Capability []interface{}
}

func (capabilities *NetconfState_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *NetconfState_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *NetconfState_Capabilities) GetGoName(yname string) string {
    if yname == "capability" { return "Capability" }
    return ""
}

func (capabilities *NetconfState_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *NetconfState_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capabilities *NetconfState_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capabilities *NetconfState_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["capability"] = capabilities.Capability
    return leafs
}

func (capabilities *NetconfState_Capabilities) GetBundleName() string { return "ietf" }

func (capabilities *NetconfState_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *NetconfState_Capabilities) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (capabilities *NetconfState_Capabilities) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (capabilities *NetconfState_Capabilities) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (capabilities *NetconfState_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *NetconfState_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *NetconfState_Capabilities) GetParentYangName() string { return "netconf-state" }

// NetconfState_Datastores
// Contains the list of NETCONF configuration datastores.
type NetconfState_Datastores struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of NETCONF configuration datastores supported by the NETCONF server
    // and related information. The type is slice of
    // NetconfState_Datastores_Datastore.
    Datastore []NetconfState_Datastores_Datastore
}

func (datastores *NetconfState_Datastores) GetFilter() yfilter.YFilter { return datastores.YFilter }

func (datastores *NetconfState_Datastores) SetFilter(yf yfilter.YFilter) { datastores.YFilter = yf }

func (datastores *NetconfState_Datastores) GetGoName(yname string) string {
    if yname == "datastore" { return "Datastore" }
    return ""
}

func (datastores *NetconfState_Datastores) GetSegmentPath() string {
    return "datastores"
}

func (datastores *NetconfState_Datastores) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "datastore" {
        for _, c := range datastores.Datastore {
            if datastores.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfState_Datastores_Datastore{}
        datastores.Datastore = append(datastores.Datastore, child)
        return &datastores.Datastore[len(datastores.Datastore)-1]
    }
    return nil
}

func (datastores *NetconfState_Datastores) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range datastores.Datastore {
        children[datastores.Datastore[i].GetSegmentPath()] = &datastores.Datastore[i]
    }
    return children
}

func (datastores *NetconfState_Datastores) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (datastores *NetconfState_Datastores) GetBundleName() string { return "ietf" }

func (datastores *NetconfState_Datastores) GetYangName() string { return "datastores" }

func (datastores *NetconfState_Datastores) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (datastores *NetconfState_Datastores) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (datastores *NetconfState_Datastores) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (datastores *NetconfState_Datastores) SetParent(parent types.Entity) { datastores.parent = parent }

func (datastores *NetconfState_Datastores) GetParent() types.Entity { return datastores.parent }

func (datastores *NetconfState_Datastores) GetParentYangName() string { return "netconf-state" }

// NetconfState_Datastores_Datastore
// List of NETCONF configuration datastores supported by
// the NETCONF server and related information.
type NetconfState_Datastores_Datastore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the datastore associated with this list
    // entry. The type is NetconfDatastoreType.
    Name interface{}

    // The NETCONF <lock> and <partial-lock> operations allow a client to lock
    // specific resources in a datastore.  The NETCONF server will prevent changes
    // to the locked resources by all sessions except the one that acquired the
    // lock(s).  Monitoring information is provided for each datastore entry
    // including details such as the session that acquired the lock, the type of
    // lock (global or partial) and the list of locked resources.  Multiple locks
    // per datastore are supported.
    Locks NetconfState_Datastores_Datastore_Locks
}

func (datastore *NetconfState_Datastores_Datastore) GetFilter() yfilter.YFilter { return datastore.YFilter }

func (datastore *NetconfState_Datastores_Datastore) SetFilter(yf yfilter.YFilter) { datastore.YFilter = yf }

func (datastore *NetconfState_Datastores_Datastore) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "locks" { return "Locks" }
    return ""
}

func (datastore *NetconfState_Datastores_Datastore) GetSegmentPath() string {
    return "datastore" + "[name='" + fmt.Sprintf("%v", datastore.Name) + "']"
}

func (datastore *NetconfState_Datastores_Datastore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "locks" {
        return &datastore.Locks
    }
    return nil
}

func (datastore *NetconfState_Datastores_Datastore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["locks"] = &datastore.Locks
    return children
}

func (datastore *NetconfState_Datastores_Datastore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = datastore.Name
    return leafs
}

func (datastore *NetconfState_Datastores_Datastore) GetBundleName() string { return "ietf" }

func (datastore *NetconfState_Datastores_Datastore) GetYangName() string { return "datastore" }

func (datastore *NetconfState_Datastores_Datastore) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (datastore *NetconfState_Datastores_Datastore) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (datastore *NetconfState_Datastores_Datastore) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (datastore *NetconfState_Datastores_Datastore) SetParent(parent types.Entity) { datastore.parent = parent }

func (datastore *NetconfState_Datastores_Datastore) GetParent() types.Entity { return datastore.parent }

func (datastore *NetconfState_Datastores_Datastore) GetParentYangName() string { return "datastores" }

// NetconfState_Datastores_Datastore_Locks
// The NETCONF <lock> and <partial-lock> operations allow
// a client to lock specific resources in a datastore.  The
// NETCONF server will prevent changes to the locked
// resources by all sessions except the one that acquired
// the lock(s).
// 
// Monitoring information is provided for each datastore
// entry including details such as the session that acquired
// the lock, the type of lock (global or partial) and the
// list of locked resources.  Multiple locks per datastore
// are supported.
// This type is a presence type.
type NetconfState_Datastores_Datastore_Locks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Present if the global lock is set.
    GlobalLock NetconfState_Datastores_Datastore_Locks_GlobalLock

    // List of partial locks. The type is slice of
    // NetconfState_Datastores_Datastore_Locks_PartialLock.
    PartialLock []NetconfState_Datastores_Datastore_Locks_PartialLock
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetFilter() yfilter.YFilter { return locks.YFilter }

func (locks *NetconfState_Datastores_Datastore_Locks) SetFilter(yf yfilter.YFilter) { locks.YFilter = yf }

func (locks *NetconfState_Datastores_Datastore_Locks) GetGoName(yname string) string {
    if yname == "global-lock" { return "GlobalLock" }
    if yname == "partial-lock" { return "PartialLock" }
    return ""
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetSegmentPath() string {
    return "locks"
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-lock" {
        return &locks.GlobalLock
    }
    if childYangName == "partial-lock" {
        for _, c := range locks.PartialLock {
            if locks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfState_Datastores_Datastore_Locks_PartialLock{}
        locks.PartialLock = append(locks.PartialLock, child)
        return &locks.PartialLock[len(locks.PartialLock)-1]
    }
    return nil
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-lock"] = &locks.GlobalLock
    for i := range locks.PartialLock {
        children[locks.PartialLock[i].GetSegmentPath()] = &locks.PartialLock[i]
    }
    return children
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetBundleName() string { return "ietf" }

func (locks *NetconfState_Datastores_Datastore_Locks) GetYangName() string { return "locks" }

func (locks *NetconfState_Datastores_Datastore_Locks) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (locks *NetconfState_Datastores_Datastore_Locks) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (locks *NetconfState_Datastores_Datastore_Locks) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (locks *NetconfState_Datastores_Datastore_Locks) SetParent(parent types.Entity) { locks.parent = parent }

func (locks *NetconfState_Datastores_Datastore_Locks) GetParent() types.Entity { return locks.parent }

func (locks *NetconfState_Datastores_Datastore_Locks) GetParentYangName() string { return "datastore" }

// NetconfState_Datastores_Datastore_Locks_GlobalLock
// Present if the global lock is set.
type NetconfState_Datastores_Datastore_Locks_GlobalLock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The session ID of the session that has locked this resource.  Both a global
    // lock and a partial lock MUST contain the NETCONF session-id.  If the lock
    // is held by a session that is not managed by the NETCONF server (e.g., a CLI
    // session), a session id of 0 (zero) is reported. The type is interface{}
    // with range: 0..4294967295. This attribute is mandatory.
    LockedBySession interface{}

    // The date and time of when the resource was locked. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    // This attribute is mandatory.
    LockedTime interface{}
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetFilter() yfilter.YFilter { return globalLock.YFilter }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) SetFilter(yf yfilter.YFilter) { globalLock.YFilter = yf }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetGoName(yname string) string {
    if yname == "locked-by-session" { return "LockedBySession" }
    if yname == "locked-time" { return "LockedTime" }
    return ""
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetSegmentPath() string {
    return "global-lock"
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["locked-by-session"] = globalLock.LockedBySession
    leafs["locked-time"] = globalLock.LockedTime
    return leafs
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetBundleName() string { return "ietf" }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetYangName() string { return "global-lock" }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) SetParent(parent types.Entity) { globalLock.parent = parent }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetParent() types.Entity { return globalLock.parent }

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetParentYangName() string { return "locks" }

// NetconfState_Datastores_Datastore_Locks_PartialLock
// List of partial locks.
type NetconfState_Datastores_Datastore_Locks_PartialLock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the lock id returned in the <partial-lock>
    // response. The type is interface{} with range: 0..4294967295.
    LockId interface{}

    // The session ID of the session that has locked this resource.  Both a global
    // lock and a partial lock MUST contain the NETCONF session-id.  If the lock
    // is held by a session that is not managed by the NETCONF server (e.g., a CLI
    // session), a session id of 0 (zero) is reported. The type is interface{}
    // with range: 0..4294967295. This attribute is mandatory.
    LockedBySession interface{}

    // The date and time of when the resource was locked. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    // This attribute is mandatory.
    LockedTime interface{}

    // The xpath expression that was used to request the lock.  The select
    // expression indicates the original intended scope of the lock. The type is
    // slice of string.
    Select []interface{}

    // The list of instance-identifiers (i.e., the locked nodes).  The scope of
    // the partial lock is defined by the list of locked nodes. The type is slice
    // of string.
    LockedNode []interface{}
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetFilter() yfilter.YFilter { return partialLock.YFilter }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) SetFilter(yf yfilter.YFilter) { partialLock.YFilter = yf }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetGoName(yname string) string {
    if yname == "lock-id" { return "LockId" }
    if yname == "locked-by-session" { return "LockedBySession" }
    if yname == "locked-time" { return "LockedTime" }
    if yname == "select" { return "Select" }
    if yname == "locked-node" { return "LockedNode" }
    return ""
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetSegmentPath() string {
    return "partial-lock" + "[lock-id='" + fmt.Sprintf("%v", partialLock.LockId) + "']"
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lock-id"] = partialLock.LockId
    leafs["locked-by-session"] = partialLock.LockedBySession
    leafs["locked-time"] = partialLock.LockedTime
    leafs["select"] = partialLock.Select
    leafs["locked-node"] = partialLock.LockedNode
    return leafs
}

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetBundleName() string { return "ietf" }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetYangName() string { return "partial-lock" }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) SetParent(parent types.Entity) { partialLock.parent = parent }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetParent() types.Entity { return partialLock.parent }

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetParentYangName() string { return "locks" }

// NetconfState_Schemas
// Contains the list of data model schemas supported by the
// server.
type NetconfState_Schemas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of data model schemas supported by the server. The type is slice of
    // NetconfState_Schemas_Schema.
    Schema []NetconfState_Schemas_Schema
}

func (schemas *NetconfState_Schemas) GetFilter() yfilter.YFilter { return schemas.YFilter }

func (schemas *NetconfState_Schemas) SetFilter(yf yfilter.YFilter) { schemas.YFilter = yf }

func (schemas *NetconfState_Schemas) GetGoName(yname string) string {
    if yname == "schema" { return "Schema" }
    return ""
}

func (schemas *NetconfState_Schemas) GetSegmentPath() string {
    return "schemas"
}

func (schemas *NetconfState_Schemas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "schema" {
        for _, c := range schemas.Schema {
            if schemas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfState_Schemas_Schema{}
        schemas.Schema = append(schemas.Schema, child)
        return &schemas.Schema[len(schemas.Schema)-1]
    }
    return nil
}

func (schemas *NetconfState_Schemas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range schemas.Schema {
        children[schemas.Schema[i].GetSegmentPath()] = &schemas.Schema[i]
    }
    return children
}

func (schemas *NetconfState_Schemas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (schemas *NetconfState_Schemas) GetBundleName() string { return "ietf" }

func (schemas *NetconfState_Schemas) GetYangName() string { return "schemas" }

func (schemas *NetconfState_Schemas) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (schemas *NetconfState_Schemas) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (schemas *NetconfState_Schemas) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (schemas *NetconfState_Schemas) SetParent(parent types.Entity) { schemas.parent = parent }

func (schemas *NetconfState_Schemas) GetParent() types.Entity { return schemas.parent }

func (schemas *NetconfState_Schemas) GetParentYangName() string { return "netconf-state" }

// NetconfState_Schemas_Schema
// List of data model schemas supported by the server.
type NetconfState_Schemas_Schema struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Identifier to uniquely reference the schema.  The
    // identifier is used in the <get-schema> operation and may be used for other
    // purposes such as file retrieval.  For modeling languages that support or
    // require a data model name (e.g., YANG module name) the identifier MUST
    // match that name.  For YANG data models, the identifier is the name of the
    // module or submodule.  In other cases, an identifier such as a filename MAY
    // be used instead. The type is string.
    Identifier interface{}

    // This attribute is a key. Version of the schema supported.  Multiple
    // versions MAY be supported simultaneously by a NETCONF server.  Each version
    // MUST be reported individually in the schema list, i.e., with same
    // identifier, possibly different location, but different version.  For YANG
    // data models, version is the value of the most recent YANG 'revision'
    // statement in the module or submodule, or the empty string if no 'revision'
    // statement is present. The type is string.
    Version interface{}

    // This attribute is a key. The data modeling language the schema is written
    // in (currently xsd, yang, yin, rng, or rnc). For YANG data models, 'yang'
    // format MUST be supported and 'yin' format MAY also be provided. The type is
    // one of the following: RncYinRngXsdYang.
    Format interface{}

    // The XML namespace defined by the data model.  For YANG data models, this is
    // the module's namespace. If the list entry describes a submodule, this field
    // contains the namespace of the module to which the submodule belongs. The
    // type is string. This attribute is mandatory.
    Namespace interface{}

    // One or more locations from which the schema can be retrieved.  This list
    // SHOULD contain at least one entry per schema.  A schema entry may be
    // located on a remote file system (e.g., reference to file system for ftp
    // retrieval) or retrieved directly from a server supporting the <get-schema>
    // operation (denoted by the value 'NETCONF'). The type is one of the
    // following types: slice of  
    // :go:struct:`NetconfState_Schemas_Schema_Location
    // <ydk/models/netconf_monitoring/NetconfState_Schemas_Schema_Location>`, or
    // slice of string.
    Location []interface{}
}

func (schema *NetconfState_Schemas_Schema) GetFilter() yfilter.YFilter { return schema.YFilter }

func (schema *NetconfState_Schemas_Schema) SetFilter(yf yfilter.YFilter) { schema.YFilter = yf }

func (schema *NetconfState_Schemas_Schema) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    if yname == "version" { return "Version" }
    if yname == "format" { return "Format" }
    if yname == "namespace" { return "Namespace" }
    if yname == "location" { return "Location" }
    return ""
}

func (schema *NetconfState_Schemas_Schema) GetSegmentPath() string {
    return "schema" + "[identifier='" + fmt.Sprintf("%v", schema.Identifier) + "']" + "[version='" + fmt.Sprintf("%v", schema.Version) + "']" + "[format='" + fmt.Sprintf("%v", schema.Format) + "']"
}

func (schema *NetconfState_Schemas_Schema) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (schema *NetconfState_Schemas_Schema) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (schema *NetconfState_Schemas_Schema) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier"] = schema.Identifier
    leafs["version"] = schema.Version
    leafs["format"] = schema.Format
    leafs["namespace"] = schema.Namespace
    leafs["location"] = schema.Location
    return leafs
}

func (schema *NetconfState_Schemas_Schema) GetBundleName() string { return "ietf" }

func (schema *NetconfState_Schemas_Schema) GetYangName() string { return "schema" }

func (schema *NetconfState_Schemas_Schema) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (schema *NetconfState_Schemas_Schema) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (schema *NetconfState_Schemas_Schema) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (schema *NetconfState_Schemas_Schema) SetParent(parent types.Entity) { schema.parent = parent }

func (schema *NetconfState_Schemas_Schema) GetParent() types.Entity { return schema.parent }

func (schema *NetconfState_Schemas_Schema) GetParentYangName() string { return "schemas" }

// NetconfState_Schemas_Schema_Location represents <get-schema> operation (denoted by the value 'NETCONF').
type NetconfState_Schemas_Schema_Location string

const (
    NetconfState_Schemas_Schema_Location_NETCONF NetconfState_Schemas_Schema_Location = "NETCONF"
)

// NetconfState_Sessions
// The sessions container includes session-specific data for
// NETCONF management sessions.  The session list MUST include
// all currently active NETCONF sessions.
type NetconfState_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All NETCONF sessions managed by the NETCONF server MUST be reported in this
    // list. The type is slice of NetconfState_Sessions_Session.
    Session []NetconfState_Sessions_Session
}

func (sessions *NetconfState_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *NetconfState_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *NetconfState_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *NetconfState_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *NetconfState_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfState_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *NetconfState_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *NetconfState_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *NetconfState_Sessions) GetBundleName() string { return "ietf" }

func (sessions *NetconfState_Sessions) GetYangName() string { return "sessions" }

func (sessions *NetconfState_Sessions) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sessions *NetconfState_Sessions) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sessions *NetconfState_Sessions) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sessions *NetconfState_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *NetconfState_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *NetconfState_Sessions) GetParentYangName() string { return "netconf-state" }

// NetconfState_Sessions_Session
// All NETCONF sessions managed by the NETCONF server
// MUST be reported in this list.
type NetconfState_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Unique identifier for the session.  This value is
    // the NETCONF session identifier, as defined in RFC 4741. The type is
    // interface{} with range: 1..4294967295.
    SessionId interface{}

    // Identifies the transport for each session, e.g., 'netconf-ssh',
    // 'netconf-soap', etc. The type is one of the following:
    // NetconfBeepNetconfSshNetconfSoapOverBeepNetconfTlsNetconfSoapOverHttps.
    // This attribute is mandatory.
    Transport interface{}

    // The username is the client identity that was authenticated by the NETCONF
    // transport protocol.  The algorithm used to derive the username is NETCONF
    // transport protocol specific and in addition specific to the authentication
    // mechanism used by the NETCONF transport protocol. The type is string. This
    // attribute is mandatory.
    Username interface{}

    // Host identifier of the NETCONF client.  The value returned is
    // implementation specific (e.g., hostname, IPv4 address, IPv6 address). The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string with length: 1..253.
    SourceHost interface{}

    // Time at the server at which the session was established. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}). This
    // attribute is mandatory.
    LoginTime interface{}

    // Number of correct <rpc> messages received. The type is interface{} with
    // range: 0..4294967295.
    InRpcs interface{}

    // Number of messages received when an <rpc> message was expected, that were
    // not correct <rpc> messages.  This includes XML parse errors and errors on
    // the rpc layer. The type is interface{} with range: 0..4294967295.
    InBadRpcs interface{}

    // Number of <rpc-reply> messages sent that contained an <rpc-error> element.
    // The type is interface{} with range: 0..4294967295.
    OutRpcErrors interface{}

    // Number of <notification> messages sent. The type is interface{} with range:
    // 0..4294967295.
    OutNotifications interface{}
}

func (session *NetconfState_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *NetconfState_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *NetconfState_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "transport" { return "Transport" }
    if yname == "username" { return "Username" }
    if yname == "source-host" { return "SourceHost" }
    if yname == "login-time" { return "LoginTime" }
    if yname == "in-rpcs" { return "InRpcs" }
    if yname == "in-bad-rpcs" { return "InBadRpcs" }
    if yname == "out-rpc-errors" { return "OutRpcErrors" }
    if yname == "out-notifications" { return "OutNotifications" }
    return ""
}

func (session *NetconfState_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *NetconfState_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *NetconfState_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *NetconfState_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["transport"] = session.Transport
    leafs["username"] = session.Username
    leafs["source-host"] = session.SourceHost
    leafs["login-time"] = session.LoginTime
    leafs["in-rpcs"] = session.InRpcs
    leafs["in-bad-rpcs"] = session.InBadRpcs
    leafs["out-rpc-errors"] = session.OutRpcErrors
    leafs["out-notifications"] = session.OutNotifications
    return leafs
}

func (session *NetconfState_Sessions_Session) GetBundleName() string { return "ietf" }

func (session *NetconfState_Sessions_Session) GetYangName() string { return "session" }

func (session *NetconfState_Sessions_Session) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (session *NetconfState_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (session *NetconfState_Sessions_Session) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (session *NetconfState_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *NetconfState_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *NetconfState_Sessions_Session) GetParentYangName() string { return "sessions" }

// NetconfState_Statistics
// Statistical data pertaining to the NETCONF server.
type NetconfState_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Date and time at which the management subsystem was started. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NetconfStartTime interface{}

    // Number of sessions silently dropped because an invalid <hello> message was
    // received.  This includes <hello> messages with a 'session-id' attribute,
    // bad namespace, and bad capability declarations. The type is interface{}
    // with range: 0..4294967295.
    InBadHellos interface{}

    // Number of sessions started.  This counter is incremented when a <hello>
    // message with a <session-id> is sent.  'in-sessions' - 'in-bad-hellos' =   
    // 'number of correctly started netconf sessions'. The type is interface{}
    // with range: 0..4294967295.
    InSessions interface{}

    // Number of sessions that were abnormally terminated, e.g., due to idle
    // timeout or transport close.  This counter is not incremented when a session
    // is properly closed by a <close-session> operation, or killed by a
    // <kill-session> operation. The type is interface{} with range:
    // 0..4294967295.
    DroppedSessions interface{}

    // Number of correct <rpc> messages received. The type is interface{} with
    // range: 0..4294967295.
    InRpcs interface{}

    // Number of messages received when an <rpc> message was expected, that were
    // not correct <rpc> messages.  This includes XML parse errors and errors on
    // the rpc layer. The type is interface{} with range: 0..4294967295.
    InBadRpcs interface{}

    // Number of <rpc-reply> messages sent that contained an <rpc-error> element.
    // The type is interface{} with range: 0..4294967295.
    OutRpcErrors interface{}

    // Number of <notification> messages sent. The type is interface{} with range:
    // 0..4294967295.
    OutNotifications interface{}
}

func (statistics *NetconfState_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *NetconfState_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *NetconfState_Statistics) GetGoName(yname string) string {
    if yname == "netconf-start-time" { return "NetconfStartTime" }
    if yname == "in-bad-hellos" { return "InBadHellos" }
    if yname == "in-sessions" { return "InSessions" }
    if yname == "dropped-sessions" { return "DroppedSessions" }
    if yname == "in-rpcs" { return "InRpcs" }
    if yname == "in-bad-rpcs" { return "InBadRpcs" }
    if yname == "out-rpc-errors" { return "OutRpcErrors" }
    if yname == "out-notifications" { return "OutNotifications" }
    return ""
}

func (statistics *NetconfState_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *NetconfState_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *NetconfState_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *NetconfState_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["netconf-start-time"] = statistics.NetconfStartTime
    leafs["in-bad-hellos"] = statistics.InBadHellos
    leafs["in-sessions"] = statistics.InSessions
    leafs["dropped-sessions"] = statistics.DroppedSessions
    leafs["in-rpcs"] = statistics.InRpcs
    leafs["in-bad-rpcs"] = statistics.InBadRpcs
    leafs["out-rpc-errors"] = statistics.OutRpcErrors
    leafs["out-notifications"] = statistics.OutNotifications
    return leafs
}

func (statistics *NetconfState_Statistics) GetBundleName() string { return "ietf" }

func (statistics *NetconfState_Statistics) GetYangName() string { return "statistics" }

func (statistics *NetconfState_Statistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (statistics *NetconfState_Statistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (statistics *NetconfState_Statistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (statistics *NetconfState_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *NetconfState_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *NetconfState_Statistics) GetParentYangName() string { return "netconf-state" }

// GetSchema
// This operation is used to retrieve a schema from the
// NETCONF server.
// 
// Positive Response:
//   The NETCONF server returns the requested schema.
// 
// Negative Response:
//   If requested schema does not exist, the <error-tag> is
//   'invalid-value'.
// 
//   If more than one schema matches the requested parameters, the
//   <error-tag> is 'operation-failed', and <error-app-tag> is
//   'data-not-unique'.
type GetSchema struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input GetSchema_Input

    
    Output GetSchema_Output
}

func (getSchema *GetSchema) GetFilter() yfilter.YFilter { return getSchema.YFilter }

func (getSchema *GetSchema) SetFilter(yf yfilter.YFilter) { getSchema.YFilter = yf }

func (getSchema *GetSchema) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (getSchema *GetSchema) GetSegmentPath() string {
    return "ietf-netconf-monitoring:get-schema"
}

func (getSchema *GetSchema) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &getSchema.Input
    }
    if childYangName == "output" {
        return &getSchema.Output
    }
    return nil
}

func (getSchema *GetSchema) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &getSchema.Input
    children["output"] = &getSchema.Output
    return children
}

func (getSchema *GetSchema) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (getSchema *GetSchema) GetBundleName() string { return "ietf" }

func (getSchema *GetSchema) GetYangName() string { return "get-schema" }

func (getSchema *GetSchema) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (getSchema *GetSchema) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (getSchema *GetSchema) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (getSchema *GetSchema) SetParent(parent types.Entity) { getSchema.parent = parent }

func (getSchema *GetSchema) GetParent() types.Entity { return getSchema.parent }

func (getSchema *GetSchema) GetParentYangName() string { return "ietf-netconf-monitoring" }

// GetSchema_Input
type GetSchema_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Identifier for the schema list entry. The type is string. This attribute is
    // mandatory.
    Identifier interface{}

    // Version of the schema requested.  If this parameter is not present, and
    // more than one version of the schema exists on the server, a
    // 'data-not-unique' error is returned, as described above. The type is
    // string.
    Version interface{}

    // The data modeling language of the schema.  If this parameter is not
    // present, and more than one formats of the schema exists on the server, a
    // 'data-not-unique' error is returned, as described above. The type is one of
    // the following: RncYinRngXsdYang.
    Format interface{}
}

func (input *GetSchema_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *GetSchema_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *GetSchema_Input) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    if yname == "version" { return "Version" }
    if yname == "format" { return "Format" }
    return ""
}

func (input *GetSchema_Input) GetSegmentPath() string {
    return "input"
}

func (input *GetSchema_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *GetSchema_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *GetSchema_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier"] = input.Identifier
    leafs["version"] = input.Version
    leafs["format"] = input.Format
    return leafs
}

func (input *GetSchema_Input) GetBundleName() string { return "ietf" }

func (input *GetSchema_Input) GetYangName() string { return "input" }

func (input *GetSchema_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *GetSchema_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *GetSchema_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *GetSchema_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *GetSchema_Input) GetParent() types.Entity { return input.parent }

func (input *GetSchema_Input) GetParentYangName() string { return "get-schema" }

// GetSchema_Output
type GetSchema_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains the schema content. The type is string.
    Data interface{}
}

func (output *GetSchema_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *GetSchema_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *GetSchema_Output) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    return ""
}

func (output *GetSchema_Output) GetSegmentPath() string {
    return "output"
}

func (output *GetSchema_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *GetSchema_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *GetSchema_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data"] = output.Data
    return leafs
}

func (output *GetSchema_Output) GetBundleName() string { return "ietf" }

func (output *GetSchema_Output) GetYangName() string { return "output" }

func (output *GetSchema_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *GetSchema_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *GetSchema_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *GetSchema_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *GetSchema_Output) GetParent() types.Entity { return output.parent }

func (output *GetSchema_Output) GetParentYangName() string { return "get-schema" }

