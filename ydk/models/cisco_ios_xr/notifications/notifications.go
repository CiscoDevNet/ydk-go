// Conversion of the 'ncEvent' XSD in the 
// NETCONF Notifications RFC.
package notifications

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package notifications"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:notification:1.0 create-subscription}", reflect.TypeOf(CreateSubscription{}))
    ydk.RegisterEntity("notifications:create-subscription", reflect.TypeOf(CreateSubscription{}))
}

// CreateSubscription
// The command to create a notification subscription. It
// takes as argument the name of the notification stream
// and filter. Both of those options limit the content of
// the subscription. In addition, there are two time-related
// parameters, startTime and stopTime, which can be used to 
// select the time interval of interest to the notification
// replay feature.
type CreateSubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CreateSubscription_Input
}

func (createSubscription *CreateSubscription) GetEntityData() *types.CommonEntityData {
    createSubscription.EntityData.YFilter = createSubscription.YFilter
    createSubscription.EntityData.YangName = "create-subscription"
    createSubscription.EntityData.BundleName = "cisco_ios_xr"
    createSubscription.EntityData.ParentYangName = "notifications"
    createSubscription.EntityData.SegmentPath = "notifications:create-subscription"
    createSubscription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createSubscription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createSubscription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createSubscription.EntityData.Children = types.NewOrderedMap()
    createSubscription.EntityData.Children.Append("input", types.YChild{"Input", &createSubscription.Input})
    createSubscription.EntityData.Leafs = types.NewOrderedMap()

    createSubscription.EntityData.YListKeys = []string {}

    return &(createSubscription.EntityData)
}

// CreateSubscription_Input
type CreateSubscription_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An optional parameter that indicates which stream of  events is of
    // interest. If not present, then events in the default NETCONF stream will be
    // sent. The type is string. The default value is NETCONF.
    Stream interface{}

    // An optional parameter that indicates which subset of all possible events is
    // of interest. The format of this parameter is the same as that of the filter
    // parameter in the NETCONF protocol operations. If not present, all events
    // not precluded by other parameters will  be sent. The type is string.
    Filter interface{}

    // A parameter used to trigger the replay feature and indicates that the
    // replay should start at the time specified. If start time is not present,
    // this is not a replay subscription. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // An optional parameter used with the optional replay feature to indicate the
    // newest notifications of interest. If stop time is not present, the
    // notifications will continue until the subscription is terminated. Must be
    // used with startTime. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}
}

func (input *CreateSubscription_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "create-subscription"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", input.Stream})
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", input.StartTime})
    input.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", input.StopTime})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

