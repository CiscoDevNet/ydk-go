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
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains a list of protocol capability URIs.
    Capabilities RestconfState_Capabilities

    // Container representing the notification event streams supported by the
    // server.
    Streams RestconfState_Streams
}

func (restconfState *RestconfState) GetFilter() yfilter.YFilter { return restconfState.YFilter }

func (restconfState *RestconfState) SetFilter(yf yfilter.YFilter) { restconfState.YFilter = yf }

func (restconfState *RestconfState) GetGoName(yname string) string {
    if yname == "capabilities" { return "Capabilities" }
    if yname == "streams" { return "Streams" }
    return ""
}

func (restconfState *RestconfState) GetSegmentPath() string {
    return "ietf-restconf-monitoring:restconf-state"
}

func (restconfState *RestconfState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capabilities" {
        return &restconfState.Capabilities
    }
    if childYangName == "streams" {
        return &restconfState.Streams
    }
    return nil
}

func (restconfState *RestconfState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capabilities"] = &restconfState.Capabilities
    children["streams"] = &restconfState.Streams
    return children
}

func (restconfState *RestconfState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (restconfState *RestconfState) GetBundleName() string { return "ietf" }

func (restconfState *RestconfState) GetYangName() string { return "restconf-state" }

func (restconfState *RestconfState) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (restconfState *RestconfState) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (restconfState *RestconfState) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (restconfState *RestconfState) SetParent(parent types.Entity) { restconfState.parent = parent }

func (restconfState *RestconfState) GetParent() types.Entity { return restconfState.parent }

func (restconfState *RestconfState) GetParentYangName() string { return "ietf-restconf-monitoring" }

// RestconfState_Capabilities
// Contains a list of protocol capability URIs
type RestconfState_Capabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A RESTCONF protocol capability URI. The type is slice of string.
    Capability []interface{}
}

func (capabilities *RestconfState_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *RestconfState_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *RestconfState_Capabilities) GetGoName(yname string) string {
    if yname == "capability" { return "Capability" }
    return ""
}

func (capabilities *RestconfState_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *RestconfState_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capabilities *RestconfState_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capabilities *RestconfState_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["capability"] = capabilities.Capability
    return leafs
}

func (capabilities *RestconfState_Capabilities) GetBundleName() string { return "ietf" }

func (capabilities *RestconfState_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *RestconfState_Capabilities) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (capabilities *RestconfState_Capabilities) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (capabilities *RestconfState_Capabilities) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (capabilities *RestconfState_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *RestconfState_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *RestconfState_Capabilities) GetParentYangName() string { return "restconf-state" }

// RestconfState_Streams
// Container representing the notification event streams
// supported by the server.
type RestconfState_Streams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry describes an event stream supported by the server. The type is
    // slice of RestconfState_Streams_Stream.
    Stream []RestconfState_Streams_Stream
}

func (streams *RestconfState_Streams) GetFilter() yfilter.YFilter { return streams.YFilter }

func (streams *RestconfState_Streams) SetFilter(yf yfilter.YFilter) { streams.YFilter = yf }

func (streams *RestconfState_Streams) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    return ""
}

func (streams *RestconfState_Streams) GetSegmentPath() string {
    return "streams"
}

func (streams *RestconfState_Streams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stream" {
        for _, c := range streams.Stream {
            if streams.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RestconfState_Streams_Stream{}
        streams.Stream = append(streams.Stream, child)
        return &streams.Stream[len(streams.Stream)-1]
    }
    return nil
}

func (streams *RestconfState_Streams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range streams.Stream {
        children[streams.Stream[i].GetSegmentPath()] = &streams.Stream[i]
    }
    return children
}

func (streams *RestconfState_Streams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (streams *RestconfState_Streams) GetBundleName() string { return "ietf" }

func (streams *RestconfState_Streams) GetYangName() string { return "streams" }

func (streams *RestconfState_Streams) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (streams *RestconfState_Streams) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (streams *RestconfState_Streams) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (streams *RestconfState_Streams) SetParent(parent types.Entity) { streams.parent = parent }

func (streams *RestconfState_Streams) GetParent() types.Entity { return streams.parent }

func (streams *RestconfState_Streams) GetParentYangName() string { return "restconf-state" }

// RestconfState_Streams_Stream
// Each entry describes an event stream supported by
// the server.
type RestconfState_Streams_Stream struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    Access []RestconfState_Streams_Stream_Access
}

func (stream *RestconfState_Streams_Stream) GetFilter() yfilter.YFilter { return stream.YFilter }

func (stream *RestconfState_Streams_Stream) SetFilter(yf yfilter.YFilter) { stream.YFilter = yf }

func (stream *RestconfState_Streams_Stream) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "replay-support" { return "ReplaySupport" }
    if yname == "replay-log-creation-time" { return "ReplayLogCreationTime" }
    if yname == "access" { return "Access" }
    return ""
}

func (stream *RestconfState_Streams_Stream) GetSegmentPath() string {
    return "stream" + "[name='" + fmt.Sprintf("%v", stream.Name) + "']"
}

func (stream *RestconfState_Streams_Stream) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range stream.Access {
            if stream.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RestconfState_Streams_Stream_Access{}
        stream.Access = append(stream.Access, child)
        return &stream.Access[len(stream.Access)-1]
    }
    return nil
}

func (stream *RestconfState_Streams_Stream) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stream.Access {
        children[stream.Access[i].GetSegmentPath()] = &stream.Access[i]
    }
    return children
}

func (stream *RestconfState_Streams_Stream) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = stream.Name
    leafs["description"] = stream.Description
    leafs["replay-support"] = stream.ReplaySupport
    leafs["replay-log-creation-time"] = stream.ReplayLogCreationTime
    return leafs
}

func (stream *RestconfState_Streams_Stream) GetBundleName() string { return "ietf" }

func (stream *RestconfState_Streams_Stream) GetYangName() string { return "stream" }

func (stream *RestconfState_Streams_Stream) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (stream *RestconfState_Streams_Stream) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (stream *RestconfState_Streams_Stream) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (stream *RestconfState_Streams_Stream) SetParent(parent types.Entity) { stream.parent = parent }

func (stream *RestconfState_Streams_Stream) GetParent() types.Entity { return stream.parent }

func (stream *RestconfState_Streams_Stream) GetParentYangName() string { return "streams" }

// RestconfState_Streams_Stream_Access
// The server will create an entry in this list for each
// encoding format that is supported for this stream.
// The media type 'text/event-stream' is expected
// for all event streams. This list identifies the
// sub-types supported for this stream.
type RestconfState_Streams_Stream_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the secondary encoding format within the
    // 'text/event-stream' encoding used by all streams. The type 'xml' is
    // supported for XML encoding.  The type 'json' is supported for JSON
    // encoding. The type is string.
    Encoding interface{}

    // Contains a URL that represents the entry point for establishing
    // notification delivery via server sent events. The type is string. This
    // attribute is mandatory.
    Location interface{}
}

func (access *RestconfState_Streams_Stream_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *RestconfState_Streams_Stream_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *RestconfState_Streams_Stream_Access) GetGoName(yname string) string {
    if yname == "encoding" { return "Encoding" }
    if yname == "location" { return "Location" }
    return ""
}

func (access *RestconfState_Streams_Stream_Access) GetSegmentPath() string {
    return "access" + "[encoding='" + fmt.Sprintf("%v", access.Encoding) + "']"
}

func (access *RestconfState_Streams_Stream_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (access *RestconfState_Streams_Stream_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (access *RestconfState_Streams_Stream_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encoding"] = access.Encoding
    leafs["location"] = access.Location
    return leafs
}

func (access *RestconfState_Streams_Stream_Access) GetBundleName() string { return "ietf" }

func (access *RestconfState_Streams_Stream_Access) GetYangName() string { return "access" }

func (access *RestconfState_Streams_Stream_Access) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (access *RestconfState_Streams_Stream_Access) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (access *RestconfState_Streams_Stream_Access) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (access *RestconfState_Streams_Stream_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *RestconfState_Streams_Stream_Access) GetParent() types.Entity { return access.parent }

func (access *RestconfState_Streams_Stream_Access) GetParentYangName() string { return "stream" }

