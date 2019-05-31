// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG definitions
// for Cisco IOS-XR syadmin NTP configuration.
// 
// This module contains definitions
// for the following management objects:
// NTP configuration data
// 
// Copyright (c) 2013-2016 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package ntp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ntp"))
    ydk.RegisterEntity("{http://cisco.com/calvados/ntp ntp}", reflect.TypeOf(Ntp{}))
    ydk.RegisterEntity("ntp:ntp", reflect.TypeOf(Ntp{}))
    ydk.RegisterEntity("{http://cisco.com/calvados/ntp clock-action}", reflect.TypeOf(ClockAction{}))
    ydk.RegisterEntity("ntp:clock-action", reflect.TypeOf(ClockAction{}))
}

// Ntp
type Ntp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of interface{} with range: 1..65534.
    TrustedKey []interface{}

    // The type is interface{}.
    Authenticate interface{}

    // The type is slice of Ntp_Peer.
    Peer []*Ntp_Peer

    // The type is slice of Ntp_Server.
    Server []*Ntp_Server

    // The type is slice of Ntp_AuthenticationKey.
    AuthenticationKey []*Ntp_AuthenticationKey

    
    Trace Ntp_Trace
}

func (ntp *Ntp) GetEntityData() *types.CommonEntityData {
    ntp.EntityData.YFilter = ntp.YFilter
    ntp.EntityData.YangName = "ntp"
    ntp.EntityData.BundleName = "cisco_ios_xr"
    ntp.EntityData.ParentYangName = "ntp"
    ntp.EntityData.SegmentPath = "ntp:ntp"
    ntp.EntityData.AbsolutePath = ntp.EntityData.SegmentPath
    ntp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntp.EntityData.Children = types.NewOrderedMap()
    ntp.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range ntp.Peer {
        ntp.EntityData.Children.Append(types.GetSegmentPath(ntp.Peer[i]), types.YChild{"Peer", ntp.Peer[i]})
    }
    ntp.EntityData.Children.Append("server", types.YChild{"Server", nil})
    for i := range ntp.Server {
        ntp.EntityData.Children.Append(types.GetSegmentPath(ntp.Server[i]), types.YChild{"Server", ntp.Server[i]})
    }
    ntp.EntityData.Children.Append("authentication-key", types.YChild{"AuthenticationKey", nil})
    for i := range ntp.AuthenticationKey {
        ntp.EntityData.Children.Append(types.GetSegmentPath(ntp.AuthenticationKey[i]), types.YChild{"AuthenticationKey", ntp.AuthenticationKey[i]})
    }
    ntp.EntityData.Children.Append("trace", types.YChild{"Trace", &ntp.Trace})
    ntp.EntityData.Leafs = types.NewOrderedMap()
    ntp.EntityData.Leafs.Append("trusted-key", types.YLeaf{"TrustedKey", ntp.TrustedKey})
    ntp.EntityData.Leafs.Append("authenticate", types.YLeaf{"Authenticate", ntp.Authenticate})

    ntp.EntityData.YListKeys = []string {}

    return &(ntp.EntityData)
}

// Ntp_Peer
type Ntp_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: 1..4.
    Version interface{}

    // The type is interface{} with range: 1..65534.
    KeyId interface{}

    // The type is interface{}.
    Prefer interface{}
}

func (peer *Ntp_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "ntp"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Name, "name")
    peer.EntityData.AbsolutePath = "ntp:ntp/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("name", types.YLeaf{"Name", peer.Name})
    peer.EntityData.Leafs.Append("version", types.YLeaf{"Version", peer.Version})
    peer.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", peer.KeyId})
    peer.EntityData.Leafs.Append("prefer", types.YLeaf{"Prefer", peer.Prefer})

    peer.EntityData.YListKeys = []string {"Name"}

    return &(peer.EntityData)
}

// Ntp_Server
type Ntp_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: 1..4.
    Version interface{}

    // The type is interface{} with range: 1..65534.
    KeyId interface{}

    // The type is interface{}.
    Prefer interface{}
}

func (server *Ntp_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "ntp"
    server.EntityData.SegmentPath = "server" + types.AddKeyToken(server.Name, "name")
    server.EntityData.AbsolutePath = "ntp:ntp/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("name", types.YLeaf{"Name", server.Name})
    server.EntityData.Leafs.Append("version", types.YLeaf{"Version", server.Version})
    server.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", server.KeyId})
    server.EntityData.Leafs.Append("prefer", types.YLeaf{"Prefer", server.Prefer})

    server.EntityData.YListKeys = []string {"Name"}

    return &(server.EntityData)
}

// Ntp_AuthenticationKey
type Ntp_AuthenticationKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..65534.
    KeyNumber interface{}

    // The type is Md5Keyword. This attribute is mandatory.
    Md5Keyword interface{}

    // The type is Encryption.
    Encryption interface{}

    // The type is string with length: 0..32. This attribute is mandatory.
    Keyname interface{}
}

func (authenticationKey *Ntp_AuthenticationKey) GetEntityData() *types.CommonEntityData {
    authenticationKey.EntityData.YFilter = authenticationKey.YFilter
    authenticationKey.EntityData.YangName = "authentication-key"
    authenticationKey.EntityData.BundleName = "cisco_ios_xr"
    authenticationKey.EntityData.ParentYangName = "ntp"
    authenticationKey.EntityData.SegmentPath = "authentication-key" + types.AddKeyToken(authenticationKey.KeyNumber, "key-number")
    authenticationKey.EntityData.AbsolutePath = "ntp:ntp/" + authenticationKey.EntityData.SegmentPath
    authenticationKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationKey.EntityData.Children = types.NewOrderedMap()
    authenticationKey.EntityData.Leafs = types.NewOrderedMap()
    authenticationKey.EntityData.Leafs.Append("key-number", types.YLeaf{"KeyNumber", authenticationKey.KeyNumber})
    authenticationKey.EntityData.Leafs.Append("md5-keyword", types.YLeaf{"Md5Keyword", authenticationKey.Md5Keyword})
    authenticationKey.EntityData.Leafs.Append("encryption", types.YLeaf{"Encryption", authenticationKey.Encryption})
    authenticationKey.EntityData.Leafs.Append("keyname", types.YLeaf{"Keyname", authenticationKey.Keyname})

    authenticationKey.EntityData.YListKeys = []string {"KeyNumber"}

    return &(authenticationKey.EntityData)
}

// Ntp_AuthenticationKey_Encryption
type Ntp_AuthenticationKey_Encryption string

const (
    Ntp_AuthenticationKey_Encryption_clear Ntp_AuthenticationKey_Encryption = "clear"

    Ntp_AuthenticationKey_Encryption_encrypted Ntp_AuthenticationKey_Encryption = "encrypted"
)

// Ntp_AuthenticationKey_Md5Keyword
type Ntp_AuthenticationKey_Md5Keyword string

const (
    Ntp_AuthenticationKey_Md5Keyword_md5 Ntp_AuthenticationKey_Md5Keyword = "md5"
)

// Ntp_Trace
type Ntp_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    NtpHelper Ntp_Trace_NtpHelper
}

func (trace *Ntp_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "ntp"
    trace.EntityData.SegmentPath = "trace"
    trace.EntityData.AbsolutePath = "ntp:ntp/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("ntp_helper", types.YChild{"NtpHelper", &trace.NtpHelper})
    trace.EntityData.Leafs = types.NewOrderedMap()

    trace.EntityData.YListKeys = []string {}

    return &(trace.EntityData)
}

// Ntp_Trace_NtpHelper
type Ntp_Trace_NtpHelper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Ntp_Trace_NtpHelper_Trace.
    Trace []*Ntp_Trace_NtpHelper_Trace
}

func (ntpHelper *Ntp_Trace_NtpHelper) GetEntityData() *types.CommonEntityData {
    ntpHelper.EntityData.YFilter = ntpHelper.YFilter
    ntpHelper.EntityData.YangName = "ntp_helper"
    ntpHelper.EntityData.BundleName = "cisco_ios_xr"
    ntpHelper.EntityData.ParentYangName = "trace"
    ntpHelper.EntityData.SegmentPath = "ntp_helper"
    ntpHelper.EntityData.AbsolutePath = "ntp:ntp/trace/" + ntpHelper.EntityData.SegmentPath
    ntpHelper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntpHelper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntpHelper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntpHelper.EntityData.Children = types.NewOrderedMap()
    ntpHelper.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range ntpHelper.Trace {
        ntpHelper.EntityData.Children.Append(types.GetSegmentPath(ntpHelper.Trace[i]), types.YChild{"Trace", ntpHelper.Trace[i]})
    }
    ntpHelper.EntityData.Leafs = types.NewOrderedMap()

    ntpHelper.EntityData.YListKeys = []string {}

    return &(ntpHelper.EntityData)
}

// Ntp_Trace_NtpHelper_Trace
// show traceable processes
type Ntp_Trace_NtpHelper_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Ntp_Trace_NtpHelper_Trace_Location.
    Location []*Ntp_Trace_NtpHelper_Trace_Location
}

func (trace *Ntp_Trace_NtpHelper_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "ntp_helper"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "ntp:ntp/trace/ntp_helper/" + trace.EntityData.SegmentPath
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

// Ntp_Trace_NtpHelper_Trace_Location
type Ntp_Trace_NtpHelper_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Ntp_Trace_NtpHelper_Trace_Location_AllOptions.
    AllOptions []*Ntp_Trace_NtpHelper_Trace_Location_AllOptions
}

func (location *Ntp_Trace_NtpHelper_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "ntp:ntp/trace/ntp_helper/trace/" + location.EntityData.SegmentPath
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

// Ntp_Trace_NtpHelper_Trace_Location_AllOptions
type Ntp_Trace_NtpHelper_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Ntp_Trace_NtpHelper_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Ntp_Trace_NtpHelper_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Ntp_Trace_NtpHelper_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "ntp:ntp/trace/ntp_helper/trace/location/" + allOptions.EntityData.SegmentPath
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

// Ntp_Trace_NtpHelper_Trace_Location_AllOptions_TraceBlocks
type Ntp_Trace_NtpHelper_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Ntp_Trace_NtpHelper_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "ntp:ntp/trace/ntp_helper/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// ClockAction
type ClockAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Clock ClockAction_Clock
}

func (clockAction *ClockAction) GetEntityData() *types.CommonEntityData {
    clockAction.EntityData.YFilter = clockAction.YFilter
    clockAction.EntityData.YangName = "clock-action"
    clockAction.EntityData.BundleName = "cisco_ios_xr"
    clockAction.EntityData.ParentYangName = "ntp"
    clockAction.EntityData.SegmentPath = "ntp:clock-action"
    clockAction.EntityData.AbsolutePath = clockAction.EntityData.SegmentPath
    clockAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockAction.EntityData.Children = types.NewOrderedMap()
    clockAction.EntityData.Children.Append("clock", types.YChild{"Clock", &clockAction.Clock})
    clockAction.EntityData.Leafs = types.NewOrderedMap()

    clockAction.EntityData.YListKeys = []string {}

    return &(clockAction.EntityData)
}

// ClockAction_Clock
type ClockAction_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Action ClockAction_Clock_Action
}

func (clock *ClockAction_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "clock-action"
    clock.EntityData.SegmentPath = "clock"
    clock.EntityData.AbsolutePath = "ntp:clock-action/" + clock.EntityData.SegmentPath
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("Action", types.YChild{"Action", &clock.Action})
    clock.EntityData.Leafs = types.NewOrderedMap()

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// ClockAction_Clock_Action
type ClockAction_Clock_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (action *ClockAction_Clock_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "Action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "clock"
    action.EntityData.SegmentPath = "Action"
    action.EntityData.AbsolutePath = "ntp:clock-action/clock/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Leafs = types.NewOrderedMap()

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

