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
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring get-schema}", reflect.TypeOf(GetSchema{}))
    ydk.RegisterEntity("ietf-netconf-monitoring:get-schema", reflect.TypeOf(GetSchema{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring netconf-state}", reflect.TypeOf(NetconfState{}))
    ydk.RegisterEntity("ietf-netconf-monitoring:netconf-state", reflect.TypeOf(NetconfState{}))
}

type Transport struct {
}

func (id Transport) String() string {
	return "ietf-netconf-monitoring:transport"
}

type NetconfSsh struct {
}

func (id NetconfSsh) String() string {
	return "ietf-netconf-monitoring:netconf-ssh"
}

type NetconfSoapOverBeep struct {
}

func (id NetconfSoapOverBeep) String() string {
	return "ietf-netconf-monitoring:netconf-soap-over-beep"
}

type NetconfSoapOverHttps struct {
}

func (id NetconfSoapOverHttps) String() string {
	return "ietf-netconf-monitoring:netconf-soap-over-https"
}

type NetconfBeep struct {
}

func (id NetconfBeep) String() string {
	return "ietf-netconf-monitoring:netconf-beep"
}

type NetconfTls struct {
}

func (id NetconfTls) String() string {
	return "ietf-netconf-monitoring:netconf-tls"
}

type SchemaFormat struct {
}

func (id SchemaFormat) String() string {
	return "ietf-netconf-monitoring:schema-format"
}

type Xsd struct {
}

func (id Xsd) String() string {
	return "ietf-netconf-monitoring:xsd"
}

type Yang struct {
}

func (id Yang) String() string {
	return "ietf-netconf-monitoring:yang"
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

type Rnc struct {
}

func (id Rnc) String() string {
	return "ietf-netconf-monitoring:rnc"
}

// NetconfDatastoreType represents Enumeration of possible NETCONF datastore types.
type NetconfDatastoreType string

const (
    NetconfDatastoreType_running NetconfDatastoreType = "running"

    NetconfDatastoreType_candidate NetconfDatastoreType = "candidate"

    NetconfDatastoreType_startup NetconfDatastoreType = "startup"
)

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetSchema_Input

    
    Output GetSchema_Output
}

func (getSchema *GetSchema) GetEntityData() *types.CommonEntityData {
    getSchema.EntityData.YFilter = getSchema.YFilter
    getSchema.EntityData.YangName = "get-schema"
    getSchema.EntityData.BundleName = "ietf"
    getSchema.EntityData.ParentYangName = "ietf-netconf-monitoring"
    getSchema.EntityData.SegmentPath = "ietf-netconf-monitoring:get-schema"
    getSchema.EntityData.AbsolutePath = getSchema.EntityData.SegmentPath
    getSchema.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    getSchema.EntityData.NamespaceTable = ietf.GetNamespaces()
    getSchema.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    getSchema.EntityData.Children = types.NewOrderedMap()
    getSchema.EntityData.Children.Append("input", types.YChild{"Input", &getSchema.Input})
    getSchema.EntityData.Children.Append("output", types.YChild{"Output", &getSchema.Output})
    getSchema.EntityData.Leafs = types.NewOrderedMap()

    getSchema.EntityData.YListKeys = []string {}

    return &(getSchema.EntityData)
}

// GetSchema_Input
type GetSchema_Input struct {
    EntityData types.CommonEntityData
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
    // the following: XsdYangYinRngRnc.
    Format interface{}
}

func (input *GetSchema_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "get-schema"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf-monitoring:get-schema/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", input.Identifier})
    input.EntityData.Leafs.Append("version", types.YLeaf{"Version", input.Version})
    input.EntityData.Leafs.Append("format", types.YLeaf{"Format", input.Format})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetSchema_Output
type GetSchema_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains the schema content. The type is string.
    Data interface{}
}

func (output *GetSchema_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "get-schema"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "ietf-netconf-monitoring:get-schema/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("data", types.YLeaf{"Data", output.Data})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// NetconfState
// The netconf-state container is the root of the monitoring
// data model.
type NetconfState struct {
    EntityData types.CommonEntityData
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

func (netconfState *NetconfState) GetEntityData() *types.CommonEntityData {
    netconfState.EntityData.YFilter = netconfState.YFilter
    netconfState.EntityData.YangName = "netconf-state"
    netconfState.EntityData.BundleName = "ietf"
    netconfState.EntityData.ParentYangName = "ietf-netconf-monitoring"
    netconfState.EntityData.SegmentPath = "ietf-netconf-monitoring:netconf-state"
    netconfState.EntityData.AbsolutePath = netconfState.EntityData.SegmentPath
    netconfState.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    netconfState.EntityData.NamespaceTable = ietf.GetNamespaces()
    netconfState.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    netconfState.EntityData.Children = types.NewOrderedMap()
    netconfState.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &netconfState.Capabilities})
    netconfState.EntityData.Children.Append("datastores", types.YChild{"Datastores", &netconfState.Datastores})
    netconfState.EntityData.Children.Append("schemas", types.YChild{"Schemas", &netconfState.Schemas})
    netconfState.EntityData.Children.Append("sessions", types.YChild{"Sessions", &netconfState.Sessions})
    netconfState.EntityData.Children.Append("statistics", types.YChild{"Statistics", &netconfState.Statistics})
    netconfState.EntityData.Leafs = types.NewOrderedMap()

    netconfState.EntityData.YListKeys = []string {}

    return &(netconfState.EntityData)
}

// NetconfState_Capabilities
// Contains the list of NETCONF capabilities supported by the
// server.
type NetconfState_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of NETCONF capabilities supported by the server. The type is slice of
    // string.
    Capability []interface{}
}

func (capabilities *NetconfState_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "ietf"
    capabilities.EntityData.ParentYangName = "netconf-state"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/" + capabilities.EntityData.SegmentPath
    capabilities.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    capabilities.EntityData.NamespaceTable = ietf.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Leafs = types.NewOrderedMap()
    capabilities.EntityData.Leafs.Append("capability", types.YLeaf{"Capability", capabilities.Capability})

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

// NetconfState_Datastores
// Contains the list of NETCONF configuration datastores.
type NetconfState_Datastores struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of NETCONF configuration datastores supported by the NETCONF server
    // and related information. The type is slice of
    // NetconfState_Datastores_Datastore.
    Datastore []*NetconfState_Datastores_Datastore
}

func (datastores *NetconfState_Datastores) GetEntityData() *types.CommonEntityData {
    datastores.EntityData.YFilter = datastores.YFilter
    datastores.EntityData.YangName = "datastores"
    datastores.EntityData.BundleName = "ietf"
    datastores.EntityData.ParentYangName = "netconf-state"
    datastores.EntityData.SegmentPath = "datastores"
    datastores.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/" + datastores.EntityData.SegmentPath
    datastores.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    datastores.EntityData.NamespaceTable = ietf.GetNamespaces()
    datastores.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    datastores.EntityData.Children = types.NewOrderedMap()
    datastores.EntityData.Children.Append("datastore", types.YChild{"Datastore", nil})
    for i := range datastores.Datastore {
        datastores.EntityData.Children.Append(types.GetSegmentPath(datastores.Datastore[i]), types.YChild{"Datastore", datastores.Datastore[i]})
    }
    datastores.EntityData.Leafs = types.NewOrderedMap()

    datastores.EntityData.YListKeys = []string {}

    return &(datastores.EntityData)
}

// NetconfState_Datastores_Datastore
// List of NETCONF configuration datastores supported by
// the NETCONF server and related information.
type NetconfState_Datastores_Datastore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (datastore *NetconfState_Datastores_Datastore) GetEntityData() *types.CommonEntityData {
    datastore.EntityData.YFilter = datastore.YFilter
    datastore.EntityData.YangName = "datastore"
    datastore.EntityData.BundleName = "ietf"
    datastore.EntityData.ParentYangName = "datastores"
    datastore.EntityData.SegmentPath = "datastore" + types.AddKeyToken(datastore.Name, "name")
    datastore.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/datastores/" + datastore.EntityData.SegmentPath
    datastore.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    datastore.EntityData.NamespaceTable = ietf.GetNamespaces()
    datastore.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    datastore.EntityData.Children = types.NewOrderedMap()
    datastore.EntityData.Children.Append("locks", types.YChild{"Locks", &datastore.Locks})
    datastore.EntityData.Leafs = types.NewOrderedMap()
    datastore.EntityData.Leafs.Append("name", types.YLeaf{"Name", datastore.Name})

    datastore.EntityData.YListKeys = []string {"Name"}

    return &(datastore.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Present if the global lock is set.
    GlobalLock NetconfState_Datastores_Datastore_Locks_GlobalLock

    // List of partial locks. The type is slice of
    // NetconfState_Datastores_Datastore_Locks_PartialLock.
    PartialLock []*NetconfState_Datastores_Datastore_Locks_PartialLock
}

func (locks *NetconfState_Datastores_Datastore_Locks) GetEntityData() *types.CommonEntityData {
    locks.EntityData.YFilter = locks.YFilter
    locks.EntityData.YangName = "locks"
    locks.EntityData.BundleName = "ietf"
    locks.EntityData.ParentYangName = "datastore"
    locks.EntityData.SegmentPath = "locks"
    locks.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/datastores/datastore/" + locks.EntityData.SegmentPath
    locks.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    locks.EntityData.NamespaceTable = ietf.GetNamespaces()
    locks.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    locks.EntityData.Children = types.NewOrderedMap()
    locks.EntityData.Children.Append("global-lock", types.YChild{"GlobalLock", &locks.GlobalLock})
    locks.EntityData.Children.Append("partial-lock", types.YChild{"PartialLock", nil})
    for i := range locks.PartialLock {
        locks.EntityData.Children.Append(types.GetSegmentPath(locks.PartialLock[i]), types.YChild{"PartialLock", locks.PartialLock[i]})
    }
    locks.EntityData.Leafs = types.NewOrderedMap()

    locks.EntityData.YListKeys = []string {}

    return &(locks.EntityData)
}

// NetconfState_Datastores_Datastore_Locks_GlobalLock
// Present if the global lock is set.
type NetconfState_Datastores_Datastore_Locks_GlobalLock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The session ID of the session that has locked this resource.  Both a global
    // lock and a partial lock MUST contain the NETCONF session-id.  If the lock
    // is held by a session that is not managed by the NETCONF server (e.g., a CLI
    // session), a session id of 0 (zero) is reported. The type is interface{}
    // with range: 0..4294967295. This attribute is mandatory.
    LockedBySession interface{}

    // The date and time of when the resource was locked. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    // This attribute is mandatory.
    LockedTime interface{}
}

func (globalLock *NetconfState_Datastores_Datastore_Locks_GlobalLock) GetEntityData() *types.CommonEntityData {
    globalLock.EntityData.YFilter = globalLock.YFilter
    globalLock.EntityData.YangName = "global-lock"
    globalLock.EntityData.BundleName = "ietf"
    globalLock.EntityData.ParentYangName = "locks"
    globalLock.EntityData.SegmentPath = "global-lock"
    globalLock.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/datastores/datastore/locks/" + globalLock.EntityData.SegmentPath
    globalLock.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    globalLock.EntityData.NamespaceTable = ietf.GetNamespaces()
    globalLock.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    globalLock.EntityData.Children = types.NewOrderedMap()
    globalLock.EntityData.Leafs = types.NewOrderedMap()
    globalLock.EntityData.Leafs.Append("locked-by-session", types.YLeaf{"LockedBySession", globalLock.LockedBySession})
    globalLock.EntityData.Leafs.Append("locked-time", types.YLeaf{"LockedTime", globalLock.LockedTime})

    globalLock.EntityData.YListKeys = []string {}

    return &(globalLock.EntityData)
}

// NetconfState_Datastores_Datastore_Locks_PartialLock
// List of partial locks.
type NetconfState_Datastores_Datastore_Locks_PartialLock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
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

func (partialLock *NetconfState_Datastores_Datastore_Locks_PartialLock) GetEntityData() *types.CommonEntityData {
    partialLock.EntityData.YFilter = partialLock.YFilter
    partialLock.EntityData.YangName = "partial-lock"
    partialLock.EntityData.BundleName = "ietf"
    partialLock.EntityData.ParentYangName = "locks"
    partialLock.EntityData.SegmentPath = "partial-lock" + types.AddKeyToken(partialLock.LockId, "lock-id")
    partialLock.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/datastores/datastore/locks/" + partialLock.EntityData.SegmentPath
    partialLock.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    partialLock.EntityData.NamespaceTable = ietf.GetNamespaces()
    partialLock.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    partialLock.EntityData.Children = types.NewOrderedMap()
    partialLock.EntityData.Leafs = types.NewOrderedMap()
    partialLock.EntityData.Leafs.Append("lock-id", types.YLeaf{"LockId", partialLock.LockId})
    partialLock.EntityData.Leafs.Append("locked-by-session", types.YLeaf{"LockedBySession", partialLock.LockedBySession})
    partialLock.EntityData.Leafs.Append("locked-time", types.YLeaf{"LockedTime", partialLock.LockedTime})
    partialLock.EntityData.Leafs.Append("select", types.YLeaf{"Select", partialLock.Select})
    partialLock.EntityData.Leafs.Append("locked-node", types.YLeaf{"LockedNode", partialLock.LockedNode})

    partialLock.EntityData.YListKeys = []string {"LockId"}

    return &(partialLock.EntityData)
}

// NetconfState_Schemas
// Contains the list of data model schemas supported by the
// server.
type NetconfState_Schemas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of data model schemas supported by the server. The type is slice of
    // NetconfState_Schemas_Schema.
    Schema []*NetconfState_Schemas_Schema
}

func (schemas *NetconfState_Schemas) GetEntityData() *types.CommonEntityData {
    schemas.EntityData.YFilter = schemas.YFilter
    schemas.EntityData.YangName = "schemas"
    schemas.EntityData.BundleName = "ietf"
    schemas.EntityData.ParentYangName = "netconf-state"
    schemas.EntityData.SegmentPath = "schemas"
    schemas.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/" + schemas.EntityData.SegmentPath
    schemas.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    schemas.EntityData.NamespaceTable = ietf.GetNamespaces()
    schemas.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    schemas.EntityData.Children = types.NewOrderedMap()
    schemas.EntityData.Children.Append("schema", types.YChild{"Schema", nil})
    for i := range schemas.Schema {
        schemas.EntityData.Children.Append(types.GetSegmentPath(schemas.Schema[i]), types.YChild{"Schema", schemas.Schema[i]})
    }
    schemas.EntityData.Leafs = types.NewOrderedMap()

    schemas.EntityData.YListKeys = []string {}

    return &(schemas.EntityData)
}

// NetconfState_Schemas_Schema
// List of data model schemas supported by the server.
type NetconfState_Schemas_Schema struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // one of the following: XsdYangYinRngRnc.
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

func (schema *NetconfState_Schemas_Schema) GetEntityData() *types.CommonEntityData {
    schema.EntityData.YFilter = schema.YFilter
    schema.EntityData.YangName = "schema"
    schema.EntityData.BundleName = "ietf"
    schema.EntityData.ParentYangName = "schemas"
    schema.EntityData.SegmentPath = "schema" + types.AddKeyToken(schema.Identifier, "identifier") + types.AddKeyToken(schema.Version, "version") + types.AddKeyToken(schema.Format, "format")
    schema.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/schemas/" + schema.EntityData.SegmentPath
    schema.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    schema.EntityData.NamespaceTable = ietf.GetNamespaces()
    schema.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    schema.EntityData.Children = types.NewOrderedMap()
    schema.EntityData.Leafs = types.NewOrderedMap()
    schema.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", schema.Identifier})
    schema.EntityData.Leafs.Append("version", types.YLeaf{"Version", schema.Version})
    schema.EntityData.Leafs.Append("format", types.YLeaf{"Format", schema.Format})
    schema.EntityData.Leafs.Append("namespace", types.YLeaf{"Namespace", schema.Namespace})
    schema.EntityData.Leafs.Append("location", types.YLeaf{"Location", schema.Location})

    schema.EntityData.YListKeys = []string {"Identifier", "Version", "Format"}

    return &(schema.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All NETCONF sessions managed by the NETCONF server MUST be reported in this
    // list. The type is slice of NetconfState_Sessions_Session.
    Session []*NetconfState_Sessions_Session
}

func (sessions *NetconfState_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "ietf"
    sessions.EntityData.ParentYangName = "netconf-state"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sessions.EntityData.NamespaceTable = ietf.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// NetconfState_Sessions_Session
// All NETCONF sessions managed by the NETCONF server
// MUST be reported in this list.
type NetconfState_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique identifier for the session.  This value is
    // the NETCONF session identifier, as defined in RFC 4741. The type is
    // interface{} with range: 1..4294967295.
    SessionId interface{}

    // Identifies the transport for each session, e.g., 'netconf-ssh',
    // 'netconf-soap', etc. The type is one of the following:
    // NetconfSshNetconfSoapOverBeepNetconfSoapOverHttpsNetconfBeepNetconfTls.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string with length: 1..253.
    SourceHost interface{}

    // Time at the server at which the session was established. The type is string
    // with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    // This attribute is mandatory.
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

func (session *NetconfState_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "ietf"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/sessions/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    session.EntityData.NamespaceTable = ietf.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", session.Transport})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("source-host", types.YLeaf{"SourceHost", session.SourceHost})
    session.EntityData.Leafs.Append("login-time", types.YLeaf{"LoginTime", session.LoginTime})
    session.EntityData.Leafs.Append("in-rpcs", types.YLeaf{"InRpcs", session.InRpcs})
    session.EntityData.Leafs.Append("in-bad-rpcs", types.YLeaf{"InBadRpcs", session.InBadRpcs})
    session.EntityData.Leafs.Append("out-rpc-errors", types.YLeaf{"OutRpcErrors", session.OutRpcErrors})
    session.EntityData.Leafs.Append("out-notifications", types.YLeaf{"OutNotifications", session.OutNotifications})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

// NetconfState_Statistics
// Statistical data pertaining to the NETCONF server.
type NetconfState_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Date and time at which the management subsystem was started. The type is
    // string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
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

func (statistics *NetconfState_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "ietf"
    statistics.EntityData.ParentYangName = "netconf-state"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "ietf-netconf-monitoring:netconf-state/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    statistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("netconf-start-time", types.YLeaf{"NetconfStartTime", statistics.NetconfStartTime})
    statistics.EntityData.Leafs.Append("in-bad-hellos", types.YLeaf{"InBadHellos", statistics.InBadHellos})
    statistics.EntityData.Leafs.Append("in-sessions", types.YLeaf{"InSessions", statistics.InSessions})
    statistics.EntityData.Leafs.Append("dropped-sessions", types.YLeaf{"DroppedSessions", statistics.DroppedSessions})
    statistics.EntityData.Leafs.Append("in-rpcs", types.YLeaf{"InRpcs", statistics.InRpcs})
    statistics.EntityData.Leafs.Append("in-bad-rpcs", types.YLeaf{"InBadRpcs", statistics.InBadRpcs})
    statistics.EntityData.Leafs.Append("out-rpc-errors", types.YLeaf{"OutRpcErrors", statistics.OutRpcErrors})
    statistics.EntityData.Leafs.Append("out-notifications", types.YLeaf{"OutNotifications", statistics.OutNotifications})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

