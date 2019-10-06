// This module contains monitoring information for the
// RESTCONF protocol.
// 
// Copyright (c) 2016 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC XXXX; see
// the RFC itself for full legal notices.
package restconf_monitoring

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package restconf_monitoring"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-restconf-monitoring restconf-state}", reflect.TypeOf(RestconfState{}))
    ydk.RegisterEntity("ietf-restconf-monitoring:restconf-state", reflect.TypeOf(RestconfState{}))
}

// RestconfState
// Contains RESTCONF protocol monitoring information.
type RestconfState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains a list of protocol capability URIs.
    Capabilities RestconfState_Capabilities

    // Container representing the notification event streams supported by the
    // server.
    Streams RestconfState_Streams
}

func (restconfState *RestconfState) GetEntityData() *types.CommonEntityData {
    restconfState.EntityData.YFilter = restconfState.YFilter
    restconfState.EntityData.YangName = "restconf-state"
    restconfState.EntityData.BundleName = "ietf"
    restconfState.EntityData.ParentYangName = "ietf-restconf-monitoring"
    restconfState.EntityData.SegmentPath = "ietf-restconf-monitoring:restconf-state"
    restconfState.EntityData.AbsolutePath = restconfState.EntityData.SegmentPath
    restconfState.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    restconfState.EntityData.NamespaceTable = ietf.GetNamespaces()
    restconfState.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    restconfState.EntityData.Children = types.NewOrderedMap()
    restconfState.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &restconfState.Capabilities})
    restconfState.EntityData.Children.Append("streams", types.YChild{"Streams", &restconfState.Streams})
    restconfState.EntityData.Leafs = types.NewOrderedMap()

    restconfState.EntityData.YListKeys = []string {}

    return &(restconfState.EntityData)
}

// RestconfState_Capabilities
// Contains a list of protocol capability URIs
type RestconfState_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A RESTCONF protocol capability URI. The type is slice of string.
    Capability []interface{}
}

func (capabilities *RestconfState_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "ietf"
    capabilities.EntityData.ParentYangName = "restconf-state"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.AbsolutePath = "ietf-restconf-monitoring:restconf-state/" + capabilities.EntityData.SegmentPath
    capabilities.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    capabilities.EntityData.NamespaceTable = ietf.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Leafs = types.NewOrderedMap()
    capabilities.EntityData.Leafs.Append("capability", types.YLeaf{"Capability", capabilities.Capability})

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

// RestconfState_Streams
// Container representing the notification event streams
// supported by the server.
type RestconfState_Streams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry describes an event stream supported by the server. The type is
    // slice of RestconfState_Streams_Stream.
    Stream []*RestconfState_Streams_Stream
}

func (streams *RestconfState_Streams) GetEntityData() *types.CommonEntityData {
    streams.EntityData.YFilter = streams.YFilter
    streams.EntityData.YangName = "streams"
    streams.EntityData.BundleName = "ietf"
    streams.EntityData.ParentYangName = "restconf-state"
    streams.EntityData.SegmentPath = "streams"
    streams.EntityData.AbsolutePath = "ietf-restconf-monitoring:restconf-state/" + streams.EntityData.SegmentPath
    streams.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    streams.EntityData.NamespaceTable = ietf.GetNamespaces()
    streams.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    streams.EntityData.Children = types.NewOrderedMap()
    streams.EntityData.Children.Append("stream", types.YChild{"Stream", nil})
    for i := range streams.Stream {
        streams.EntityData.Children.Append(types.GetSegmentPath(streams.Stream[i]), types.YChild{"Stream", streams.Stream[i]})
    }
    streams.EntityData.Leafs = types.NewOrderedMap()

    streams.EntityData.YListKeys = []string {}

    return &(streams.EntityData)
}

// RestconfState_Streams_Stream
// Each entry describes an event stream supported by
// the server.
type RestconfState_Streams_Stream struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The stream name. The type is string.
    Name interface{}

    // Description of stream content. The type is string.
    Description interface{}

    // Indicates if replay buffer supported for this stream. If 'true', then the
    // server MUST support the 'start-time' and 'stop-time' query parameters for
    // this stream. The type is bool. The default value is false.
    ReplaySupport interface{}

    // Indicates the time the replay log for this stream was created. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ReplayLogCreationTime interface{}

    // The server will create an entry in this list for each encoding format that
    // is supported for this stream. The media type 'text/event-stream' is
    // expected for all event streams. This list identifies the sub-types
    // supported for this stream. The type is slice of
    // RestconfState_Streams_Stream_Access.
    Access []*RestconfState_Streams_Stream_Access
}

func (stream *RestconfState_Streams_Stream) GetEntityData() *types.CommonEntityData {
    stream.EntityData.YFilter = stream.YFilter
    stream.EntityData.YangName = "stream"
    stream.EntityData.BundleName = "ietf"
    stream.EntityData.ParentYangName = "streams"
    stream.EntityData.SegmentPath = "stream" + types.AddKeyToken(stream.Name, "name")
    stream.EntityData.AbsolutePath = "ietf-restconf-monitoring:restconf-state/streams/" + stream.EntityData.SegmentPath
    stream.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    stream.EntityData.NamespaceTable = ietf.GetNamespaces()
    stream.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    stream.EntityData.Children = types.NewOrderedMap()
    stream.EntityData.Children.Append("access", types.YChild{"Access", nil})
    for i := range stream.Access {
        stream.EntityData.Children.Append(types.GetSegmentPath(stream.Access[i]), types.YChild{"Access", stream.Access[i]})
    }
    stream.EntityData.Leafs = types.NewOrderedMap()
    stream.EntityData.Leafs.Append("name", types.YLeaf{"Name", stream.Name})
    stream.EntityData.Leafs.Append("description", types.YLeaf{"Description", stream.Description})
    stream.EntityData.Leafs.Append("replay-support", types.YLeaf{"ReplaySupport", stream.ReplaySupport})
    stream.EntityData.Leafs.Append("replay-log-creation-time", types.YLeaf{"ReplayLogCreationTime", stream.ReplayLogCreationTime})

    stream.EntityData.YListKeys = []string {"Name"}

    return &(stream.EntityData)
}

// RestconfState_Streams_Stream_Access
// The server will create an entry in this list for each
// encoding format that is supported for this stream.
// The media type 'text/event-stream' is expected
// for all event streams. This list identifies the
// sub-types supported for this stream.
type RestconfState_Streams_Stream_Access struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is the secondary encoding format within the
    // 'text/event-stream' encoding used by all streams. The type 'xml' is
    // supported for XML encoding. The type 'json' is supported for JSON encoding.
    // The type is string.
    Encoding interface{}

    // Contains a URL that represents the entry point for establishing
    // notification delivery via server sent events. The type is string. This
    // attribute is mandatory.
    Location interface{}
}

func (access *RestconfState_Streams_Stream_Access) GetEntityData() *types.CommonEntityData {
    access.EntityData.YFilter = access.YFilter
    access.EntityData.YangName = "access"
    access.EntityData.BundleName = "ietf"
    access.EntityData.ParentYangName = "stream"
    access.EntityData.SegmentPath = "access" + types.AddKeyToken(access.Encoding, "encoding")
    access.EntityData.AbsolutePath = "ietf-restconf-monitoring:restconf-state/streams/stream/" + access.EntityData.SegmentPath
    access.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    access.EntityData.NamespaceTable = ietf.GetNamespaces()
    access.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    access.EntityData.Children = types.NewOrderedMap()
    access.EntityData.Leafs = types.NewOrderedMap()
    access.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", access.Encoding})
    access.EntityData.Leafs.Append("location", types.YLeaf{"Location", access.Location})

    access.EntityData.YListKeys = []string {"Encoding"}

    return &(access.EntityData)
}

