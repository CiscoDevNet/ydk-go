// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// The Directory Services (DS).
// 
// Copyright(c) 2010-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_ds

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_ds"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-ds services}", reflect.TypeOf(Services{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-ds:services", reflect.TypeOf(Services{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-ds services-stats}", reflect.TypeOf(ServicesStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-ds:services-stats", reflect.TypeOf(ServicesStats{}))
}

// ProcessIssuRole
type ProcessIssuRole string

const (
    ProcessIssuRole_Primary ProcessIssuRole = "Primary"

    ProcessIssuRole_Secondary ProcessIssuRole = "Secondary"

    ProcessIssuRole_Tertiary ProcessIssuRole = "Tertiary"

    ProcessIssuRole_Unknown ProcessIssuRole = "Unknown"
)

// ProcessRole
type ProcessRole string

const (
    ProcessRole_NoRole ProcessRole = "NoRole"

    ProcessRole_Active ProcessRole = "Active"

    ProcessRole_Standby ProcessRole = "Standby"

    ProcessRole_None ProcessRole = "None"

    ProcessRole_Unknown ProcessRole = "Unknown"
)

// ServiceScope
type ServiceScope string

const (
    ServiceScope_None ServiceScope = "None"

    ServiceScope_Rack ServiceScope = "Rack"

    ServiceScope_System ServiceScope = "System"

    ServiceScope_Node ServiceScope = "Node"
)

// Services
// Directory Services Entries
type Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Services_AllLocations.
    AllLocations []*Services_AllLocations
}

func (services *Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-ds"
    services.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-ds:services"
    services.EntityData.AbsolutePath = services.EntityData.SegmentPath
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = types.NewOrderedMap()
    services.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range services.AllLocations {
        services.EntityData.Children.Append(types.GetSegmentPath(services.AllLocations[i]), types.YChild{"AllLocations", services.AllLocations[i]})
    }
    services.EntityData.Leafs = types.NewOrderedMap()

    services.EntityData.YListKeys = []string {}

    return &(services.EntityData)
}

// Services_AllLocations
type Services_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node Location. The type is string.
    Location interface{}

    // The type is slice of Services_AllLocations_Services.
    Services []*Services_AllLocations_Services
}

func (allLocations *Services_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "services"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services/" + allLocations.EntityData.SegmentPath
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Children.Append("services", types.YChild{"Services", nil})
    for i := range allLocations.Services {
        allLocations.EntityData.Children.Append(types.GetSegmentPath(allLocations.Services[i]), types.YChild{"Services", allLocations.Services[i]})
    }
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

// Services_AllLocations_Services
type Services_AllLocations_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the service. The type is string.
    Name interface{}

    // endpoint info for a service in DS. The type is slice of
    // Services_AllLocations_Services_Endpoint.
    Endpoint []*Services_AllLocations_Services_Endpoint

    // clients registered for a service in DS. The type is slice of
    // Services_AllLocations_Services_Registrations.
    Registrations []*Services_AllLocations_Services_Registrations
}

func (services *Services_AllLocations_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "all-locations"
    services.EntityData.SegmentPath = "services" + types.AddKeyToken(services.Name, "name")
    services.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services/all-locations/" + services.EntityData.SegmentPath
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = types.NewOrderedMap()
    services.EntityData.Children.Append("endpoint", types.YChild{"Endpoint", nil})
    for i := range services.Endpoint {
        types.SetYListKey(services.Endpoint[i], i)
        services.EntityData.Children.Append(types.GetSegmentPath(services.Endpoint[i]), types.YChild{"Endpoint", services.Endpoint[i]})
    }
    services.EntityData.Children.Append("registrations", types.YChild{"Registrations", nil})
    for i := range services.Registrations {
        types.SetYListKey(services.Registrations[i], i)
        services.EntityData.Children.Append(types.GetSegmentPath(services.Registrations[i]), types.YChild{"Registrations", services.Registrations[i]})
    }
    services.EntityData.Leafs = types.NewOrderedMap()
    services.EntityData.Leafs.Append("name", types.YLeaf{"Name", services.Name})

    services.EntityData.YListKeys = []string {"Name"}

    return &(services.EntityData)
}

// Services_AllLocations_Services_Endpoint
// endpoint info for a service in DS
type Services_AllLocations_Services_Endpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is ServiceScope.
    Scope interface{}

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}

    // The type is ProcessRole.
    Role interface{}

    // The type is ProcessIssuRole.
    IssuRole interface{}

    // Ethernet address of the node hosting the endpoint. The type is string.
    Node interface{}
}

func (endpoint *Services_AllLocations_Services_Endpoint) GetEntityData() *types.CommonEntityData {
    endpoint.EntityData.YFilter = endpoint.YFilter
    endpoint.EntityData.YangName = "endpoint"
    endpoint.EntityData.BundleName = "cisco_ios_xr"
    endpoint.EntityData.ParentYangName = "services"
    endpoint.EntityData.SegmentPath = "endpoint" + types.AddNoKeyToken(endpoint)
    endpoint.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services/all-locations/services/" + endpoint.EntityData.SegmentPath
    endpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endpoint.EntityData.Children = types.NewOrderedMap()
    endpoint.EntityData.Leafs = types.NewOrderedMap()
    endpoint.EntityData.Leafs.Append("scope", types.YLeaf{"Scope", endpoint.Scope})
    endpoint.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", endpoint.Ip})
    endpoint.EntityData.Leafs.Append("port", types.YLeaf{"Port", endpoint.Port})
    endpoint.EntityData.Leafs.Append("role", types.YLeaf{"Role", endpoint.Role})
    endpoint.EntityData.Leafs.Append("issu_role", types.YLeaf{"IssuRole", endpoint.IssuRole})
    endpoint.EntityData.Leafs.Append("node", types.YLeaf{"Node", endpoint.Node})

    endpoint.EntityData.YListKeys = []string {}

    return &(endpoint.EntityData)
}

// Services_AllLocations_Services_Registrations
// clients registered for a service in DS
type Services_AllLocations_Services_Registrations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    Client interface{}

    // The type is interface{} with range: 0..4294967295.
    Pid interface{}
}

func (registrations *Services_AllLocations_Services_Registrations) GetEntityData() *types.CommonEntityData {
    registrations.EntityData.YFilter = registrations.YFilter
    registrations.EntityData.YangName = "registrations"
    registrations.EntityData.BundleName = "cisco_ios_xr"
    registrations.EntityData.ParentYangName = "services"
    registrations.EntityData.SegmentPath = "registrations" + types.AddNoKeyToken(registrations)
    registrations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services/all-locations/services/" + registrations.EntityData.SegmentPath
    registrations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registrations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registrations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registrations.EntityData.Children = types.NewOrderedMap()
    registrations.EntityData.Leafs = types.NewOrderedMap()
    registrations.EntityData.Leafs.Append("client", types.YLeaf{"Client", registrations.Client})
    registrations.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", registrations.Pid})

    registrations.EntityData.YListKeys = []string {}

    return &(registrations.EntityData)
}

// ServicesStats
// Directory Services Stats
type ServicesStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ds ServicesStats_Ds

    // The type is slice of ServicesStats_AllLocations.
    AllLocations []*ServicesStats_AllLocations
}

func (servicesStats *ServicesStats) GetEntityData() *types.CommonEntityData {
    servicesStats.EntityData.YFilter = servicesStats.YFilter
    servicesStats.EntityData.YangName = "services-stats"
    servicesStats.EntityData.BundleName = "cisco_ios_xr"
    servicesStats.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-ds"
    servicesStats.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-ds:services-stats"
    servicesStats.EntityData.AbsolutePath = servicesStats.EntityData.SegmentPath
    servicesStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicesStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicesStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicesStats.EntityData.Children = types.NewOrderedMap()
    servicesStats.EntityData.Children.Append("ds", types.YChild{"Ds", &servicesStats.Ds})
    servicesStats.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range servicesStats.AllLocations {
        servicesStats.EntityData.Children.Append(types.GetSegmentPath(servicesStats.AllLocations[i]), types.YChild{"AllLocations", servicesStats.AllLocations[i]})
    }
    servicesStats.EntityData.Leafs = types.NewOrderedMap()

    servicesStats.EntityData.YListKeys = []string {}

    return &(servicesStats.EntityData)
}

// ServicesStats_Ds
type ServicesStats_Ds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of ServicesStats_Ds_Trace.
    Trace []*ServicesStats_Ds_Trace
}

func (ds *ServicesStats_Ds) GetEntityData() *types.CommonEntityData {
    ds.EntityData.YFilter = ds.YFilter
    ds.EntityData.YangName = "ds"
    ds.EntityData.BundleName = "cisco_ios_xr"
    ds.EntityData.ParentYangName = "services-stats"
    ds.EntityData.SegmentPath = "ds"
    ds.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/" + ds.EntityData.SegmentPath
    ds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ds.EntityData.Children = types.NewOrderedMap()
    ds.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range ds.Trace {
        ds.EntityData.Children.Append(types.GetSegmentPath(ds.Trace[i]), types.YChild{"Trace", ds.Trace[i]})
    }
    ds.EntityData.Leafs = types.NewOrderedMap()

    ds.EntityData.YListKeys = []string {}

    return &(ds.EntityData)
}

// ServicesStats_Ds_Trace
// show traceable processes
type ServicesStats_Ds_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of ServicesStats_Ds_Trace_Location.
    Location []*ServicesStats_Ds_Trace_Location
}

func (trace *ServicesStats_Ds_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "ds"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/ds/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// ServicesStats_Ds_Trace_Location
type ServicesStats_Ds_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of ServicesStats_Ds_Trace_Location_AllOptions.
    AllOptions []*ServicesStats_Ds_Trace_Location_AllOptions
}

func (location *ServicesStats_Ds_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/ds/trace/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// ServicesStats_Ds_Trace_Location_AllOptions
type ServicesStats_Ds_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // ServicesStats_Ds_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*ServicesStats_Ds_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *ServicesStats_Ds_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/ds/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// ServicesStats_Ds_Trace_Location_AllOptions_TraceBlocks
type ServicesStats_Ds_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *ServicesStats_Ds_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/ds/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// ServicesStats_AllLocations
type ServicesStats_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of ServicesStats_AllLocations_Stats.
    Stats []*ServicesStats_AllLocations_Stats
}

func (allLocations *ServicesStats_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "services-stats"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/" + allLocations.EntityData.SegmentPath
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Children.Append("stats", types.YChild{"Stats", nil})
    for i := range allLocations.Stats {
        allLocations.EntityData.Children.Append(types.GetSegmentPath(allLocations.Stats[i]), types.YChild{"Stats", allLocations.Stats[i]})
    }
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

// ServicesStats_AllLocations_Stats
type ServicesStats_AllLocations_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the service. The type is string.
    Name interface{}

    // number of endpoints published for this service. The type is interface{}
    // with range: 0..4294967295.
    Published interface{}

    // number of endpoints deleted for this service. The type is interface{} with
    // range: 0..4294967295.
    Deleted interface{}

    // number of endpoints modified for this service. The type is interface{} with
    // range: 0..4294967295.
    Modified interface{}

    // number of clients registered for this service. The type is interface{} with
    // range: 0..4294967295.
    Registered interface{}

    // number of clients un-registered for this service. The type is interface{}
    // with range: 0..4294967295.
    Unregistered interface{}

    // number of clients notified for this service. The type is interface{} with
    // range: 0..4294967295.
    Notifications interface{}

    // number of remote service updates sent to remote nodes. The type is
    // interface{} with range: 0..4294967295.
    RemoteSent interface{}

    // number of remote service received from remote nodes. The type is
    // interface{} with range: 0..4294967295.
    RemoteRecv interface{}
}

func (stats *ServicesStats_AllLocations_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "all-locations"
    stats.EntityData.SegmentPath = "stats" + types.AddKeyToken(stats.Name, "name")
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-ds:services-stats/all-locations/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("name", types.YLeaf{"Name", stats.Name})
    stats.EntityData.Leafs.Append("published", types.YLeaf{"Published", stats.Published})
    stats.EntityData.Leafs.Append("deleted", types.YLeaf{"Deleted", stats.Deleted})
    stats.EntityData.Leafs.Append("modified", types.YLeaf{"Modified", stats.Modified})
    stats.EntityData.Leafs.Append("registered", types.YLeaf{"Registered", stats.Registered})
    stats.EntityData.Leafs.Append("unregistered", types.YLeaf{"Unregistered", stats.Unregistered})
    stats.EntityData.Leafs.Append("notifications", types.YLeaf{"Notifications", stats.Notifications})
    stats.EntityData.Leafs.Append("remote_sent", types.YLeaf{"RemoteSent", stats.RemoteSent})
    stats.EntityData.Leafs.Append("remote_recv", types.YLeaf{"RemoteRecv", stats.RemoteRecv})

    stats.EntityData.YListKeys = []string {"Name"}

    return &(stats.EntityData)
}

