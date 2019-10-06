// This module contains a collection of YANG definitions for
// monitoring Unified Threat Defense (UTD).
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package utd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package utd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-utd-oper utd-oper-data}", reflect.TypeOf(UtdOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-utd-oper:utd-oper-data", reflect.TypeOf(UtdOperData{}))
}

// UtdOperStatusVal represents Unified Threat Defense (UTD) operational status
type UtdOperStatusVal string

const (
    // Unified Threat Defense (UTD) operational status is unknown - Unable to determine status
    UtdOperStatusVal_utd_oper_status_unknown UtdOperStatusVal = "utd-oper-status-unknown"

    // Unified Threat Defense (UTD) operational status is green - Working as expected
    UtdOperStatusVal_utd_oper_status_green UtdOperStatusVal = "utd-oper-status-green"

    // Unified Threat Defense (UTD) operational status is yellow - Minor problem
    UtdOperStatusVal_utd_oper_status_yellow UtdOperStatusVal = "utd-oper-status-yellow"

    // Unified Threat Defense (UTD) operational status is red - Major problem
    UtdOperStatusVal_utd_oper_status_red UtdOperStatusVal = "utd-oper-status-red"

    // Unified Threat Defense (UTD) operational status is down - Communication has been lost
    UtdOperStatusVal_utd_oper_status_down UtdOperStatusVal = "utd-oper-status-down"
)

// UtdOperData
// Unified Threat Defense (UTD) operational data
type UtdOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unified Threat Defense (UTD) engine status.
    UtdEngineStatus UtdOperData_UtdEngineStatus

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) update
    // status.
    UtdIpsUpdateStatus UtdOperData_UtdIpsUpdateStatus

    // Unified Threat Defense (UTD) URL-Filtering (URLF) update status.
    UtdUrlfUpdateStatus UtdOperData_UtdUrlfUpdateStatus
}

func (utdOperData *UtdOperData) GetEntityData() *types.CommonEntityData {
    utdOperData.EntityData.YFilter = utdOperData.YFilter
    utdOperData.EntityData.YangName = "utd-oper-data"
    utdOperData.EntityData.BundleName = "cisco_ios_xe"
    utdOperData.EntityData.ParentYangName = "Cisco-IOS-XE-utd-oper"
    utdOperData.EntityData.SegmentPath = "Cisco-IOS-XE-utd-oper:utd-oper-data"
    utdOperData.EntityData.AbsolutePath = utdOperData.EntityData.SegmentPath
    utdOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utdOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utdOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utdOperData.EntityData.Children = types.NewOrderedMap()
    utdOperData.EntityData.Children.Append("utd-engine-status", types.YChild{"UtdEngineStatus", &utdOperData.UtdEngineStatus})
    utdOperData.EntityData.Children.Append("utd-ips-update-status", types.YChild{"UtdIpsUpdateStatus", &utdOperData.UtdIpsUpdateStatus})
    utdOperData.EntityData.Children.Append("utd-urlf-update-status", types.YChild{"UtdUrlfUpdateStatus", &utdOperData.UtdUrlfUpdateStatus})
    utdOperData.EntityData.Leafs = types.NewOrderedMap()

    utdOperData.EntityData.YListKeys = []string {}

    return &(utdOperData.EntityData)
}

// UtdOperData_UtdEngineStatus
// Unified Threat Defense (UTD) engine status
// This type is a presence type.
type UtdOperData_UtdEngineStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Engine version. The type is string.
    Version interface{}

    // Profile. The type is string.
    Profile interface{}

    // Overall status. The type is UtdOperStatusVal.
    Status interface{}

    // Overall status reason. The type is string.
    Reason interface{}

    // Percentage of memory used. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are percent.
    MemoryUsage interface{}

    // Status of memory usage. The type is UtdOperStatusVal.
    MemoryStatus interface{}

    // Status of engine instances. The type is slice of
    // UtdOperData_UtdEngineStatus_UtdEngineInstanceStatus.
    UtdEngineInstanceStatus []*UtdOperData_UtdEngineStatus_UtdEngineInstanceStatus
}

func (utdEngineStatus *UtdOperData_UtdEngineStatus) GetEntityData() *types.CommonEntityData {
    utdEngineStatus.EntityData.YFilter = utdEngineStatus.YFilter
    utdEngineStatus.EntityData.YangName = "utd-engine-status"
    utdEngineStatus.EntityData.BundleName = "cisco_ios_xe"
    utdEngineStatus.EntityData.ParentYangName = "utd-oper-data"
    utdEngineStatus.EntityData.SegmentPath = "utd-engine-status"
    utdEngineStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/" + utdEngineStatus.EntityData.SegmentPath
    utdEngineStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utdEngineStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utdEngineStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utdEngineStatus.EntityData.Children = types.NewOrderedMap()
    utdEngineStatus.EntityData.Children.Append("utd-engine-instance-status", types.YChild{"UtdEngineInstanceStatus", nil})
    for i := range utdEngineStatus.UtdEngineInstanceStatus {
        utdEngineStatus.EntityData.Children.Append(types.GetSegmentPath(utdEngineStatus.UtdEngineInstanceStatus[i]), types.YChild{"UtdEngineInstanceStatus", utdEngineStatus.UtdEngineInstanceStatus[i]})
    }
    utdEngineStatus.EntityData.Leafs = types.NewOrderedMap()
    utdEngineStatus.EntityData.Leafs.Append("version", types.YLeaf{"Version", utdEngineStatus.Version})
    utdEngineStatus.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", utdEngineStatus.Profile})
    utdEngineStatus.EntityData.Leafs.Append("status", types.YLeaf{"Status", utdEngineStatus.Status})
    utdEngineStatus.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", utdEngineStatus.Reason})
    utdEngineStatus.EntityData.Leafs.Append("memory-usage", types.YLeaf{"MemoryUsage", utdEngineStatus.MemoryUsage})
    utdEngineStatus.EntityData.Leafs.Append("memory-status", types.YLeaf{"MemoryStatus", utdEngineStatus.MemoryStatus})

    utdEngineStatus.EntityData.YListKeys = []string {}

    return &(utdEngineStatus.EntityData)
}

// UtdOperData_UtdEngineStatus_UtdEngineInstanceStatus
// Status of engine instances
type UtdOperData_UtdEngineStatus_UtdEngineInstanceStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Engine instance ID. The type is interface{} with
    // range: 0..4294967295.
    Id interface{}

    // Engine instance running. The type is bool.
    Running interface{}

    // Engine instance status. The type is UtdOperStatusVal.
    Status interface{}

    // Engine instance status reason. The type is string.
    Reason interface{}
}

func (utdEngineInstanceStatus *UtdOperData_UtdEngineStatus_UtdEngineInstanceStatus) GetEntityData() *types.CommonEntityData {
    utdEngineInstanceStatus.EntityData.YFilter = utdEngineInstanceStatus.YFilter
    utdEngineInstanceStatus.EntityData.YangName = "utd-engine-instance-status"
    utdEngineInstanceStatus.EntityData.BundleName = "cisco_ios_xe"
    utdEngineInstanceStatus.EntityData.ParentYangName = "utd-engine-status"
    utdEngineInstanceStatus.EntityData.SegmentPath = "utd-engine-instance-status" + types.AddKeyToken(utdEngineInstanceStatus.Id, "id")
    utdEngineInstanceStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/utd-engine-status/" + utdEngineInstanceStatus.EntityData.SegmentPath
    utdEngineInstanceStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utdEngineInstanceStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utdEngineInstanceStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utdEngineInstanceStatus.EntityData.Children = types.NewOrderedMap()
    utdEngineInstanceStatus.EntityData.Leafs = types.NewOrderedMap()
    utdEngineInstanceStatus.EntityData.Leafs.Append("id", types.YLeaf{"Id", utdEngineInstanceStatus.Id})
    utdEngineInstanceStatus.EntityData.Leafs.Append("running", types.YLeaf{"Running", utdEngineInstanceStatus.Running})
    utdEngineInstanceStatus.EntityData.Leafs.Append("status", types.YLeaf{"Status", utdEngineInstanceStatus.Status})
    utdEngineInstanceStatus.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", utdEngineInstanceStatus.Reason})

    utdEngineInstanceStatus.EntityData.YListKeys = []string {"Id"}

    return &(utdEngineInstanceStatus.EntityData)
}

// UtdOperData_UtdIpsUpdateStatus
// Unified Threat Defense (UTD) Intrusion Prevention System (IPS) update status
// This type is a presence type.
type UtdOperData_UtdIpsUpdateStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Intrusion Prevention System (IPS) update status.
    IpsUpdateStatus UtdOperData_UtdIpsUpdateStatus_IpsUpdateStatus
}

func (utdIpsUpdateStatus *UtdOperData_UtdIpsUpdateStatus) GetEntityData() *types.CommonEntityData {
    utdIpsUpdateStatus.EntityData.YFilter = utdIpsUpdateStatus.YFilter
    utdIpsUpdateStatus.EntityData.YangName = "utd-ips-update-status"
    utdIpsUpdateStatus.EntityData.BundleName = "cisco_ios_xe"
    utdIpsUpdateStatus.EntityData.ParentYangName = "utd-oper-data"
    utdIpsUpdateStatus.EntityData.SegmentPath = "utd-ips-update-status"
    utdIpsUpdateStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/" + utdIpsUpdateStatus.EntityData.SegmentPath
    utdIpsUpdateStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utdIpsUpdateStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utdIpsUpdateStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utdIpsUpdateStatus.EntityData.Children = types.NewOrderedMap()
    utdIpsUpdateStatus.EntityData.Children.Append("ips-update-status", types.YChild{"IpsUpdateStatus", &utdIpsUpdateStatus.IpsUpdateStatus})
    utdIpsUpdateStatus.EntityData.Leafs = types.NewOrderedMap()

    utdIpsUpdateStatus.EntityData.YListKeys = []string {}

    return &(utdIpsUpdateStatus.EntityData)
}

// UtdOperData_UtdIpsUpdateStatus_IpsUpdateStatus
// Intrusion Prevention System (IPS) update status
type UtdOperData_UtdIpsUpdateStatus_IpsUpdateStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Time of last attempted update. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateTime interface{}

    // Status of last attempted update. The type is UtdUpdateStatusVal.
    LastUpdateStatus interface{}

    // Reason for last attempted update failure. The type is string.
    LastUpdateReason interface{}

    // Time of last successful update. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastSuccessfulUpdateTime interface{}
}

func (ipsUpdateStatus *UtdOperData_UtdIpsUpdateStatus_IpsUpdateStatus) GetEntityData() *types.CommonEntityData {
    ipsUpdateStatus.EntityData.YFilter = ipsUpdateStatus.YFilter
    ipsUpdateStatus.EntityData.YangName = "ips-update-status"
    ipsUpdateStatus.EntityData.BundleName = "cisco_ios_xe"
    ipsUpdateStatus.EntityData.ParentYangName = "utd-ips-update-status"
    ipsUpdateStatus.EntityData.SegmentPath = "ips-update-status"
    ipsUpdateStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/utd-ips-update-status/" + ipsUpdateStatus.EntityData.SegmentPath
    ipsUpdateStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsUpdateStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsUpdateStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsUpdateStatus.EntityData.Children = types.NewOrderedMap()
    ipsUpdateStatus.EntityData.Leafs = types.NewOrderedMap()
    ipsUpdateStatus.EntityData.Leafs.Append("version", types.YLeaf{"Version", ipsUpdateStatus.Version})
    ipsUpdateStatus.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", ipsUpdateStatus.LastUpdateTime})
    ipsUpdateStatus.EntityData.Leafs.Append("last-update-status", types.YLeaf{"LastUpdateStatus", ipsUpdateStatus.LastUpdateStatus})
    ipsUpdateStatus.EntityData.Leafs.Append("last-update-reason", types.YLeaf{"LastUpdateReason", ipsUpdateStatus.LastUpdateReason})
    ipsUpdateStatus.EntityData.Leafs.Append("last-successful-update-time", types.YLeaf{"LastSuccessfulUpdateTime", ipsUpdateStatus.LastSuccessfulUpdateTime})

    ipsUpdateStatus.EntityData.YListKeys = []string {}

    return &(ipsUpdateStatus.EntityData)
}

// UtdOperData_UtdUrlfUpdateStatus
// Unified Threat Defense (UTD) URL-Filtering (URLF) update status
// This type is a presence type.
type UtdOperData_UtdUrlfUpdateStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // URL-Filtering (URLF) update status.
    UrlfUpdateStatus UtdOperData_UtdUrlfUpdateStatus_UrlfUpdateStatus
}

func (utdUrlfUpdateStatus *UtdOperData_UtdUrlfUpdateStatus) GetEntityData() *types.CommonEntityData {
    utdUrlfUpdateStatus.EntityData.YFilter = utdUrlfUpdateStatus.YFilter
    utdUrlfUpdateStatus.EntityData.YangName = "utd-urlf-update-status"
    utdUrlfUpdateStatus.EntityData.BundleName = "cisco_ios_xe"
    utdUrlfUpdateStatus.EntityData.ParentYangName = "utd-oper-data"
    utdUrlfUpdateStatus.EntityData.SegmentPath = "utd-urlf-update-status"
    utdUrlfUpdateStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/" + utdUrlfUpdateStatus.EntityData.SegmentPath
    utdUrlfUpdateStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utdUrlfUpdateStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utdUrlfUpdateStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utdUrlfUpdateStatus.EntityData.Children = types.NewOrderedMap()
    utdUrlfUpdateStatus.EntityData.Children.Append("urlf-update-status", types.YChild{"UrlfUpdateStatus", &utdUrlfUpdateStatus.UrlfUpdateStatus})
    utdUrlfUpdateStatus.EntityData.Leafs = types.NewOrderedMap()

    utdUrlfUpdateStatus.EntityData.YListKeys = []string {}

    return &(utdUrlfUpdateStatus.EntityData)
}

// UtdOperData_UtdUrlfUpdateStatus_UrlfUpdateStatus
// URL-Filtering (URLF) update status
type UtdOperData_UtdUrlfUpdateStatus_UrlfUpdateStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Time of last attempted update. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateTime interface{}

    // Status of last attempted update. The type is UtdUpdateStatusVal.
    LastUpdateStatus interface{}

    // Reason for last attempted update failure. The type is string.
    LastUpdateReason interface{}

    // Time of last successful update. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastSuccessfulUpdateTime interface{}
}

func (urlfUpdateStatus *UtdOperData_UtdUrlfUpdateStatus_UrlfUpdateStatus) GetEntityData() *types.CommonEntityData {
    urlfUpdateStatus.EntityData.YFilter = urlfUpdateStatus.YFilter
    urlfUpdateStatus.EntityData.YangName = "urlf-update-status"
    urlfUpdateStatus.EntityData.BundleName = "cisco_ios_xe"
    urlfUpdateStatus.EntityData.ParentYangName = "utd-urlf-update-status"
    urlfUpdateStatus.EntityData.SegmentPath = "urlf-update-status"
    urlfUpdateStatus.EntityData.AbsolutePath = "Cisco-IOS-XE-utd-oper:utd-oper-data/utd-urlf-update-status/" + urlfUpdateStatus.EntityData.SegmentPath
    urlfUpdateStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    urlfUpdateStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    urlfUpdateStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    urlfUpdateStatus.EntityData.Children = types.NewOrderedMap()
    urlfUpdateStatus.EntityData.Leafs = types.NewOrderedMap()
    urlfUpdateStatus.EntityData.Leafs.Append("version", types.YLeaf{"Version", urlfUpdateStatus.Version})
    urlfUpdateStatus.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", urlfUpdateStatus.LastUpdateTime})
    urlfUpdateStatus.EntityData.Leafs.Append("last-update-status", types.YLeaf{"LastUpdateStatus", urlfUpdateStatus.LastUpdateStatus})
    urlfUpdateStatus.EntityData.Leafs.Append("last-update-reason", types.YLeaf{"LastUpdateReason", urlfUpdateStatus.LastUpdateReason})
    urlfUpdateStatus.EntityData.Leafs.Append("last-successful-update-time", types.YLeaf{"LastSuccessfulUpdateTime", urlfUpdateStatus.LastSuccessfulUpdateTime})

    urlfUpdateStatus.EntityData.YListKeys = []string {}

    return &(urlfUpdateStatus.EntityData)
}

