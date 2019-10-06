// Conversion of the 'manageEvent' XSD in the NETCONF
// Notifications RFC.
package nc_notifications

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nc_notifications"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netmod:notification netconf}", reflect.TypeOf(Netconf{}))
    ydk.RegisterEntity("nc-notifications:netconf", reflect.TypeOf(Netconf{}))
}

// Netconf
// Top-level element in the notification namespace
type Netconf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of event streams supported by the system. When a query is issued,
    // the returned set of streams is determined based on user privileges.
    Streams Netconf_Streams
}

func (netconf *Netconf) GetEntityData() *types.CommonEntityData {
    netconf.EntityData.YFilter = netconf.YFilter
    netconf.EntityData.YangName = "netconf"
    netconf.EntityData.BundleName = "cisco_ios_xr"
    netconf.EntityData.ParentYangName = "nc-notifications"
    netconf.EntityData.SegmentPath = "nc-notifications:netconf"
    netconf.EntityData.AbsolutePath = netconf.EntityData.SegmentPath
    netconf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconf.EntityData.Children = types.NewOrderedMap()
    netconf.EntityData.Children.Append("streams", types.YChild{"Streams", &netconf.Streams})
    netconf.EntityData.Leafs = types.NewOrderedMap()

    netconf.EntityData.YListKeys = []string {}

    return &(netconf.EntityData)
}

// Netconf_Streams
// The list of event streams supported by the system. When
// a query is issued, the returned set of streams is
// determined based on user privileges.
type Netconf_Streams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stream name, description and other information. The type is slice of
    // Netconf_Streams_Stream.
    Stream []*Netconf_Streams_Stream
}

func (streams *Netconf_Streams) GetEntityData() *types.CommonEntityData {
    streams.EntityData.YFilter = streams.YFilter
    streams.EntityData.YangName = "streams"
    streams.EntityData.BundleName = "cisco_ios_xr"
    streams.EntityData.ParentYangName = "netconf"
    streams.EntityData.SegmentPath = "streams"
    streams.EntityData.AbsolutePath = "nc-notifications:netconf/" + streams.EntityData.SegmentPath
    streams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    streams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    streams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    streams.EntityData.Children = types.NewOrderedMap()
    streams.EntityData.Children.Append("stream", types.YChild{"Stream", nil})
    for i := range streams.Stream {
        streams.EntityData.Children.Append(types.GetSegmentPath(streams.Stream[i]), types.YChild{"Stream", streams.Stream[i]})
    }
    streams.EntityData.Leafs = types.NewOrderedMap()

    streams.EntityData.YListKeys = []string {}

    return &(streams.EntityData)
}

// Netconf_Streams_Stream
// Stream name, description and other information.
type Netconf_Streams_Stream struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the event stream. If this is the
    // default NETCONF stream, this must have the value 'NETCONF'. The type is
    // string.
    Name interface{}

    // A description of the event stream, including such information as the type
    // of events that are sent over this stream. The type is string. This
    // attribute is mandatory.
    Description interface{}

    // A description of the event stream, including such information as the type
    // of events that are sent over this stream. The type is bool. This attribute
    // is mandatory.
    ReplaySupport interface{}

    // The timestamp of the creation of the log used to support the replay
    // function on this stream. Note that this might be earlier then the earliest
    // available notification in the log. This object is updated if the log resets
    // for some reason.  This object MUST be present if replay is supported. The
    // type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ReplayLogCreationTime interface{}
}

func (stream *Netconf_Streams_Stream) GetEntityData() *types.CommonEntityData {
    stream.EntityData.YFilter = stream.YFilter
    stream.EntityData.YangName = "stream"
    stream.EntityData.BundleName = "cisco_ios_xr"
    stream.EntityData.ParentYangName = "streams"
    stream.EntityData.SegmentPath = "stream" + types.AddKeyToken(stream.Name, "name")
    stream.EntityData.AbsolutePath = "nc-notifications:netconf/streams/" + stream.EntityData.SegmentPath
    stream.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stream.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stream.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stream.EntityData.Children = types.NewOrderedMap()
    stream.EntityData.Leafs = types.NewOrderedMap()
    stream.EntityData.Leafs.Append("name", types.YLeaf{"Name", stream.Name})
    stream.EntityData.Leafs.Append("description", types.YLeaf{"Description", stream.Description})
    stream.EntityData.Leafs.Append("replaySupport", types.YLeaf{"ReplaySupport", stream.ReplaySupport})
    stream.EntityData.Leafs.Append("replayLogCreationTime", types.YLeaf{"ReplayLogCreationTime", stream.ReplayLogCreationTime})

    stream.EntityData.YListKeys = []string {"Name"}

    return &(stream.EntityData)
}

