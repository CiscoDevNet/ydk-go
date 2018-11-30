// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin operational.
// 
// This module defines operational data for
// sysadmin ship process.
// 
// Copyright(c) 2010-2016 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_ship

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_ship"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-ship stat}", reflect.TypeOf(Stat{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-ship:stat", reflect.TypeOf(Stat{}))
}

// Stat
// SHIP Info
type Stat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select ship client component. The type is slice of Stat_ShipComp.
    ShipComp []*Stat_ShipComp
}

func (stat *Stat) GetEntityData() *types.CommonEntityData {
    stat.EntityData.YFilter = stat.YFilter
    stat.EntityData.YangName = "stat"
    stat.EntityData.BundleName = "cisco_ios_xr"
    stat.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-ship"
    stat.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-ship:stat"
    stat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stat.EntityData.Children = types.NewOrderedMap()
    stat.EntityData.Children.Append("ship_comp", types.YChild{"ShipComp", nil})
    for i := range stat.ShipComp {
        stat.EntityData.Children.Append(types.GetSegmentPath(stat.ShipComp[i]), types.YChild{"ShipComp", stat.ShipComp[i]})
    }
    stat.EntityData.Leafs = types.NewOrderedMap()

    stat.EntityData.YListKeys = []string {}

    return &(stat.EntityData)
}

// Stat_ShipComp
// Select ship client component
type Stat_ShipComp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of component. The type is string.
    CompName interface{}

    // The type is slice of Stat_ShipComp_Process.
    Process []*Stat_ShipComp_Process
}

func (shipComp *Stat_ShipComp) GetEntityData() *types.CommonEntityData {
    shipComp.EntityData.YFilter = shipComp.YFilter
    shipComp.EntityData.YangName = "ship_comp"
    shipComp.EntityData.BundleName = "cisco_ios_xr"
    shipComp.EntityData.ParentYangName = "stat"
    shipComp.EntityData.SegmentPath = "ship_comp" + types.AddKeyToken(shipComp.CompName, "comp-name")
    shipComp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shipComp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shipComp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shipComp.EntityData.Children = types.NewOrderedMap()
    shipComp.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range shipComp.Process {
        shipComp.EntityData.Children.Append(types.GetSegmentPath(shipComp.Process[i]), types.YChild{"Process", shipComp.Process[i]})
    }
    shipComp.EntityData.Leafs = types.NewOrderedMap()
    shipComp.EntityData.Leafs.Append("comp-name", types.YLeaf{"CompName", shipComp.CompName})

    shipComp.EntityData.YListKeys = []string {"CompName"}

    return &(shipComp.EntityData)
}

// Stat_ShipComp_Process
type Stat_ShipComp_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    ProcessName interface{}

    // The type is slice of Stat_ShipComp_Process_Client.
    Client []*Stat_ShipComp_Process_Client
}

func (process *Stat_ShipComp_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "ship_comp"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessName, "process-name")
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range process.Client {
        process.EntityData.Children.Append(types.GetSegmentPath(process.Client[i]), types.YChild{"Client", process.Client[i]})
    }
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {"ProcessName"}

    return &(process.EntityData)
}

// Stat_ShipComp_Process_Client
type Stat_ShipComp_Process_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    ClientName interface{}

    // The type is slice of Stat_ShipComp_Process_Client_Cat.
    Cat []*Stat_ShipComp_Process_Client_Cat
}

func (client *Stat_ShipComp_Process_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "process"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientName, "client-name")
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("cat", types.YChild{"Cat", nil})
    for i := range client.Cat {
        client.EntityData.Children.Append(types.GetSegmentPath(client.Cat[i]), types.YChild{"Cat", client.Cat[i]})
    }
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", client.ClientName})

    client.EntityData.YListKeys = []string {"ClientName"}

    return &(client.EntityData)
}

// Stat_ShipComp_Process_Client_Cat
type Stat_ShipComp_Process_Client_Cat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CatName.
    CatName interface{}

    // The type is slice of Stat_ShipComp_Process_Client_Cat_Counter32b.
    Counter32b []*Stat_ShipComp_Process_Client_Cat_Counter32b
}

func (cat *Stat_ShipComp_Process_Client_Cat) GetEntityData() *types.CommonEntityData {
    cat.EntityData.YFilter = cat.YFilter
    cat.EntityData.YangName = "cat"
    cat.EntityData.BundleName = "cisco_ios_xr"
    cat.EntityData.ParentYangName = "client"
    cat.EntityData.SegmentPath = "cat" + types.AddKeyToken(cat.CatName, "cat-name")
    cat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cat.EntityData.Children = types.NewOrderedMap()
    cat.EntityData.Children.Append("counter-32b", types.YChild{"Counter32b", nil})
    for i := range cat.Counter32b {
        cat.EntityData.Children.Append(types.GetSegmentPath(cat.Counter32b[i]), types.YChild{"Counter32b", cat.Counter32b[i]})
    }
    cat.EntityData.Leafs = types.NewOrderedMap()
    cat.EntityData.Leafs.Append("cat-name", types.YLeaf{"CatName", cat.CatName})

    cat.EntityData.YListKeys = []string {"CatName"}

    return &(cat.EntityData)
}

// Stat_ShipComp_Process_Client_Cat_Counter32b
type Stat_ShipComp_Process_Client_Cat_Counter32b struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    CounterName interface{}

    // The type is interface{} with range: 0..4294967295.
    CounterValue interface{}

    // The type is interface{} with range: 0..4294967295.
    Watermark interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TimeStamp interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo1 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo2 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo3 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo4 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo5 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo6 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo7 interface{}

    // The type is interface{} with range: 0..4294967295.
    HistInfo8 interface{}
}

func (counter32b *Stat_ShipComp_Process_Client_Cat_Counter32b) GetEntityData() *types.CommonEntityData {
    counter32b.EntityData.YFilter = counter32b.YFilter
    counter32b.EntityData.YangName = "counter-32b"
    counter32b.EntityData.BundleName = "cisco_ios_xr"
    counter32b.EntityData.ParentYangName = "cat"
    counter32b.EntityData.SegmentPath = "counter-32b" + types.AddKeyToken(counter32b.CounterName, "counter-name")
    counter32b.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counter32b.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counter32b.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counter32b.EntityData.Children = types.NewOrderedMap()
    counter32b.EntityData.Leafs = types.NewOrderedMap()
    counter32b.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", counter32b.CounterName})
    counter32b.EntityData.Leafs.Append("counter-value", types.YLeaf{"CounterValue", counter32b.CounterValue})
    counter32b.EntityData.Leafs.Append("watermark", types.YLeaf{"Watermark", counter32b.Watermark})
    counter32b.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", counter32b.TimeStamp})
    counter32b.EntityData.Leafs.Append("hist-info1", types.YLeaf{"HistInfo1", counter32b.HistInfo1})
    counter32b.EntityData.Leafs.Append("hist-info2", types.YLeaf{"HistInfo2", counter32b.HistInfo2})
    counter32b.EntityData.Leafs.Append("hist-info3", types.YLeaf{"HistInfo3", counter32b.HistInfo3})
    counter32b.EntityData.Leafs.Append("hist-info4", types.YLeaf{"HistInfo4", counter32b.HistInfo4})
    counter32b.EntityData.Leafs.Append("hist-info5", types.YLeaf{"HistInfo5", counter32b.HistInfo5})
    counter32b.EntityData.Leafs.Append("hist-info6", types.YLeaf{"HistInfo6", counter32b.HistInfo6})
    counter32b.EntityData.Leafs.Append("hist-info7", types.YLeaf{"HistInfo7", counter32b.HistInfo7})
    counter32b.EntityData.Leafs.Append("hist-info8", types.YLeaf{"HistInfo8", counter32b.HistInfo8})

    counter32b.EntityData.YListKeys = []string {"CounterName"}

    return &(counter32b.EntityData)
}

// Stat_ShipComp_Process_Client_Cat_CatName
type Stat_ShipComp_Process_Client_Cat_CatName string

const (
    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_TRANSPORT_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_TRANSPORT_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_SERVER_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_SERVER_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_CLIENT_LIB_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_CLIENT_LIB_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_REPLICATION_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_REPLICATION_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_READER_REPLICA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_READER_REPLICA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_WRITER_REPLICA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_WRITER_REPLICA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_TRANSPORT_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_TRANSPORT_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_SERVER_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_SERVER_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_CLIENT_LIB_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_CLIENT_LIB_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_REPLICATION_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_REPLICATION_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_TRANSPORT_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_TRANSPORT_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_SERVER_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_SERVER_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_CLIENT_LIB_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_CLIENT_LIB_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_REPLICATION_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_REPLICATION_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_READER_REPLICA_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_READER_REPLICA_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_WRITER_REPLICA_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_WRITER_REPLICA_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_TRANSPORT_ERRORS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_TRANSPORT_ERRORS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_SERVER_ERRORS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_SERVER_ERRORS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_CLIENT_LIB_ERRORS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_CLIENT_LIB_ERRORS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_REPLICATION_ERRORS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_REPLICATION_ERRORS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_FEATURE_MA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "FEATURE_MA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_FEATURE_MA_COUNTERS_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "FEATURE_MA_COUNTERS_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_FEATURE_MA_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "FEATURE_MA_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_MA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_MA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_MA_COUNTERS_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_MA_COUNTERS_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_MA_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_MA_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_FEATURE_EA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "FEATURE_EA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_FEATURE_EA_COUNTERS_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "FEATURE_EA_COUNTERS_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_EA_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_EA_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_EA_COUNTERS_ERRORS Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_EA_COUNTERS_ERRORS"

    Stat_ShipComp_Process_Client_Cat_CatName_VIRTUAL_INTERFACE_EA_COUNTERS_FAILOVER Stat_ShipComp_Process_Client_Cat_CatName = "VIRTUAL_INTERFACE_EA_COUNTERS_FAILOVER"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_HISTOGRAM_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_HISTOGRAM_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_WATERMARK_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_WATERMARK_COUNTERS"

    Stat_ShipComp_Process_Client_Cat_CatName_SHIP_NODE_COUNTERS Stat_ShipComp_Process_Client_Cat_CatName = "SHIP_NODE_COUNTERS"
)

