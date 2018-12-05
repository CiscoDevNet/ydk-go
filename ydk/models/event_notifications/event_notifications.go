// This module contains conceptual YANG specifications
// for NETCONF Event Notifications.
package event_notifications

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package event_notifications"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications establish-subscription}", reflect.TypeOf(EstablishSubscription{}))
    ydk.RegisterEntity("ietf-event-notifications:establish-subscription", reflect.TypeOf(EstablishSubscription{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications create-subscription}", reflect.TypeOf(CreateSubscription{}))
    ydk.RegisterEntity("ietf-event-notifications:create-subscription", reflect.TypeOf(CreateSubscription{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications modify-subscription}", reflect.TypeOf(ModifySubscription{}))
    ydk.RegisterEntity("ietf-event-notifications:modify-subscription", reflect.TypeOf(ModifySubscription{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications delete-subscription}", reflect.TypeOf(DeleteSubscription{}))
    ydk.RegisterEntity("ietf-event-notifications:delete-subscription", reflect.TypeOf(DeleteSubscription{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications streams}", reflect.TypeOf(Streams{}))
    ydk.RegisterEntity("ietf-event-notifications:streams", reflect.TypeOf(Streams{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications filters}", reflect.TypeOf(Filters{}))
    ydk.RegisterEntity("ietf-event-notifications:filters", reflect.TypeOf(Filters{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications subscription-config}", reflect.TypeOf(SubscriptionConfig{}))
    ydk.RegisterEntity("ietf-event-notifications:subscription-config", reflect.TypeOf(SubscriptionConfig{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-event-notifications subscriptions}", reflect.TypeOf(Subscriptions{}))
    ydk.RegisterEntity("ietf-event-notifications:subscriptions", reflect.TypeOf(Subscriptions{}))
}

type ErrorNoSuchOption struct {
}

func (id ErrorNoSuchOption) String() string {
	return "ietf-event-notifications:error-no-such-option"
}

type Stream struct {
}

func (id Stream) String() string {
	return "ietf-event-notifications:stream"
}

type ErrorNoSuchSubscription struct {
}

func (id ErrorNoSuchSubscription) String() string {
	return "ietf-event-notifications:error-no-such-subscription"
}

type NoResources struct {
}

func (id NoResources) String() string {
	return "ietf-event-notifications:no-resources"
}

type Inactive struct {
}

func (id Inactive) String() string {
	return "ietf-event-notifications:inactive"
}

type Suspended struct {
}

func (id Suspended) String() string {
	return "ietf-event-notifications:suspended"
}

type Encodings struct {
}

func (id Encodings) String() string {
	return "ietf-event-notifications:encodings"
}

type EncodeJson struct {
}

func (id EncodeJson) String() string {
	return "ietf-event-notifications:encode-json"
}

type Transport struct {
}

func (id Transport) String() string {
	return "ietf-event-notifications:transport"
}

type InternalError struct {
}

func (id InternalError) String() string {
	return "ietf-event-notifications:internal-error"
}

type ErrorOther struct {
}

func (id ErrorOther) String() string {
	return "ietf-event-notifications:error-other"
}

type Other struct {
}

func (id Other) String() string {
	return "ietf-event-notifications:other"
}

type InError struct {
}

func (id InError) String() string {
	return "ietf-event-notifications:in-error"
}

type ErrorInsufficientResources struct {
}

func (id ErrorInsufficientResources) String() string {
	return "ietf-event-notifications:error-insufficient-resources"
}

type Netconf struct {
}

func (id Netconf) String() string {
	return "ietf-event-notifications:netconf"
}

type ErrorConfiguredSubscription struct {
}

func (id ErrorConfiguredSubscription) String() string {
	return "ietf-event-notifications:error-configured-subscription"
}

type SubscriptionResult struct {
}

func (id SubscriptionResult) String() string {
	return "ietf-event-notifications:subscription-result"
}

type Error struct {
}

func (id Error) String() string {
	return "ietf-event-notifications:error"
}

type Active struct {
}

func (id Active) String() string {
	return "ietf-event-notifications:active"
}

type NETCONF struct {
}

func (id NETCONF) String() string {
	return "ietf-event-notifications:NETCONF"
}

type Ok struct {
}

func (id Ok) String() string {
	return "ietf-event-notifications:ok"
}

type SubscriptionStreamStatus struct {
}

func (id SubscriptionStreamStatus) String() string {
	return "ietf-event-notifications:subscription-stream-status"
}

type EncodeXml struct {
}

func (id EncodeXml) String() string {
	return "ietf-event-notifications:encode-xml"
}

type SubscriptionDeleted struct {
}

func (id SubscriptionDeleted) String() string {
	return "ietf-event-notifications:subscription-deleted"
}

type SubscriptionErrors struct {
}

func (id SubscriptionErrors) String() string {
	return "ietf-event-notifications:subscription-errors"
}

// PushSource represents being sent by the publisher.
type PushSource string

const (
    // Notifications will be sent from a specific interface on
    // a publisher
    PushSource_interface_originated PushSource = "interface-originated"

    // Notifications will be sent from a specific address on a
    // publisher
    PushSource_address_originated PushSource = "address-originated"
)

// EstablishSubscription
// This RPC allows a subscriber to create
// (and possibly negotiate) a subscription on its own behalf.
// If successful, the subscription
// remains in effect for the duration of the subscriber's
// association with the publisher, or until the subscription
// is terminated by virtue of a delete-subscription request.
// In case an error (as indicated by subscription-result)
// is returned, the subscription is
// not created.  In that case, the RPC output
// MAY include suggested parameter settings
// that would have a high likelihood of succeeding in a
// subsequent establish-subscription request.
type EstablishSubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EstablishSubscription_Input

    
    Output EstablishSubscription_Output
}

func (establishSubscription *EstablishSubscription) GetEntityData() *types.CommonEntityData {
    establishSubscription.EntityData.YFilter = establishSubscription.YFilter
    establishSubscription.EntityData.YangName = "establish-subscription"
    establishSubscription.EntityData.BundleName = "ietf"
    establishSubscription.EntityData.ParentYangName = "ietf-event-notifications"
    establishSubscription.EntityData.SegmentPath = "ietf-event-notifications:establish-subscription"
    establishSubscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    establishSubscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    establishSubscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    establishSubscription.EntityData.Children = types.NewOrderedMap()
    establishSubscription.EntityData.Children.Append("input", types.YChild{"Input", &establishSubscription.Input})
    establishSubscription.EntityData.Children.Append("output", types.YChild{"Output", &establishSubscription.Output})
    establishSubscription.EntityData.Leafs = types.NewOrderedMap()

    establishSubscription.EntityData.YListKeys = []string {}

    return &(establishSubscription.EntityData)
}

// EstablishSubscription_Input
type EstablishSubscription_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}

    // The push update's IP packet transport priority. This is made visible across
    // network hops to receiver. The transport priority is shared for all
    // receivers of a given subscription. The type is interface{} with range:
    // 0..63. The default value is 0.
    Dscp interface{}

    // Relative priority for a subscription.   Allows an underlying transport
    // layer perform informed load balance allocations between various
    // subscriptions. The type is interface{} with range: 0..255.
    SubscriptionPriority interface{}

    // Provides the Subscription ID of a parent subscription without which this
    // subscription should not exist. In other words, there is no reason to stream
    // these objects if another subscription is missing. The type is string.
    SubscriptionDependency interface{}
}

func (input *EstablishSubscription_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "establish-subscription"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", input.Stream})
    input.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", input.Encoding})
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", input.FilterRef})
    input.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", input.SubtreeFilter})
    input.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", input.XpathFilter})
    input.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", input.StartTime})
    input.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", input.StopTime})
    input.EntityData.Leafs.Append("period", types.YLeaf{"Period", input.Period})
    input.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", input.AnchorTime})
    input.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", input.NoSynchOnStart})
    input.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", input.DampeningPeriod})
    input.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", input.ExcludedChange})
    input.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", input.Dscp})
    input.EntityData.Leafs.Append("subscription-priority", types.YLeaf{"SubscriptionPriority", input.SubscriptionPriority})
    input.EntityData.Leafs.Append("subscription-dependency", types.YLeaf{"SubscriptionDependency", input.SubscriptionDependency})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EstablishSubscription_Output
type EstablishSubscription_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether subscription is operational, or if a problem was
    // encountered. The type is one of the following:
    // ErrorErrorNoSuchOptionErrorNoSuchSubscriptionErrorOtherErrorInsufficientResourcesErrorConfiguredSubscriptionErrorDataNotAuthorizedOk.
    // This attribute is mandatory.
    SubscriptionResult interface{}

    // Identifier used for this subscription. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    SubscriptionId interface{}

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}

    // The push update's IP packet transport priority. This is made visible across
    // network hops to receiver. The transport priority is shared for all
    // receivers of a given subscription. The type is interface{} with range:
    // 0..63. The default value is 0.
    Dscp interface{}

    // Relative priority for a subscription.   Allows an underlying transport
    // layer perform informed load balance allocations between various
    // subscriptions. The type is interface{} with range: 0..255.
    SubscriptionPriority interface{}

    // Provides the Subscription ID of a parent subscription without which this
    // subscription should not exist. In other words, there is no reason to stream
    // these objects if another subscription is missing. The type is string.
    SubscriptionDependency interface{}
}

func (output *EstablishSubscription_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "establish-subscription"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("subscription-result", types.YLeaf{"SubscriptionResult", output.SubscriptionResult})
    output.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", output.SubscriptionId})
    output.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", output.Stream})
    output.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", output.Encoding})
    output.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", output.Filter})
    output.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", output.FilterRef})
    output.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", output.SubtreeFilter})
    output.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", output.XpathFilter})
    output.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", output.StartTime})
    output.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", output.StopTime})
    output.EntityData.Leafs.Append("period", types.YLeaf{"Period", output.Period})
    output.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", output.AnchorTime})
    output.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", output.NoSynchOnStart})
    output.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", output.DampeningPeriod})
    output.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", output.ExcludedChange})
    output.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", output.Dscp})
    output.EntityData.Leafs.Append("subscription-priority", types.YLeaf{"SubscriptionPriority", output.SubscriptionPriority})
    output.EntityData.Leafs.Append("subscription-dependency", types.YLeaf{"SubscriptionDependency", output.SubscriptionDependency})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// CreateSubscription
// This operation initiates an event notification subscription
// that will send asynchronous event notifications to the
// initiator of the command until the association terminates.
// It is not possible to modify or delete a subscription
// that was created using this operation.  It is included for
// reasons of backward compatibility with RFC 5277
// implementations.
type CreateSubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CreateSubscription_Input
}

func (createSubscription *CreateSubscription) GetEntityData() *types.CommonEntityData {
    createSubscription.EntityData.YFilter = createSubscription.YFilter
    createSubscription.EntityData.YangName = "create-subscription"
    createSubscription.EntityData.BundleName = "ietf"
    createSubscription.EntityData.ParentYangName = "ietf-event-notifications"
    createSubscription.EntityData.SegmentPath = "ietf-event-notifications:create-subscription"
    createSubscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    createSubscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    createSubscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

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

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush. The default value is NETCONF.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}
}

func (input *CreateSubscription_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "create-subscription"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", input.Stream})
    input.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", input.Encoding})
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", input.StartTime})
    input.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", input.StopTime})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ModifySubscription
// This RPC allows a subscriber to modify a subscription
// that was previously created using establish-subscription.
// If successful, the subscription
// remains in effect for the duration of the subscriber's
// association with the publisher, or until the subscription
// is terminated by virtue of a delete-subscription request.
// In case an error is returned (as indicated by
// subscription-result), the subscription is
// not modified and the original subscription parameters
// remain in effect.  In that case, the rpc error response
// MAY include suggested parameter settings
// that would have a high likelihood of succeeding in a
// subsequent modify-subscription request.
type ModifySubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ModifySubscription_Input

    
    Output ModifySubscription_Output
}

func (modifySubscription *ModifySubscription) GetEntityData() *types.CommonEntityData {
    modifySubscription.EntityData.YFilter = modifySubscription.YFilter
    modifySubscription.EntityData.YangName = "modify-subscription"
    modifySubscription.EntityData.BundleName = "ietf"
    modifySubscription.EntityData.ParentYangName = "ietf-event-notifications"
    modifySubscription.EntityData.SegmentPath = "ietf-event-notifications:modify-subscription"
    modifySubscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    modifySubscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    modifySubscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    modifySubscription.EntityData.Children = types.NewOrderedMap()
    modifySubscription.EntityData.Children.Append("input", types.YChild{"Input", &modifySubscription.Input})
    modifySubscription.EntityData.Children.Append("output", types.YChild{"Output", &modifySubscription.Output})
    modifySubscription.EntityData.Leafs = types.NewOrderedMap()

    modifySubscription.EntityData.YListKeys = []string {}

    return &(modifySubscription.EntityData)
}

// ModifySubscription_Input
type ModifySubscription_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifier to use for this subscription. The type is interface{} with
    // range: 0..4294967295.
    SubscriptionId interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}
}

func (input *ModifySubscription_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "modify-subscription"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", input.SubscriptionId})
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", input.FilterRef})
    input.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", input.SubtreeFilter})
    input.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", input.XpathFilter})
    input.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", input.StartTime})
    input.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", input.StopTime})
    input.EntityData.Leafs.Append("period", types.YLeaf{"Period", input.Period})
    input.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", input.AnchorTime})
    input.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", input.NoSynchOnStart})
    input.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", input.DampeningPeriod})
    input.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", input.ExcludedChange})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ModifySubscription_Output
type ModifySubscription_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether subscription is operational, or if a problem was
    // encountered. The type is one of the following:
    // ErrorErrorNoSuchOptionErrorNoSuchSubscriptionErrorOtherErrorInsufficientResourcesErrorConfiguredSubscriptionErrorDataNotAuthorizedOk.
    // This attribute is mandatory.
    SubscriptionResult interface{}

    // Identifier used for this subscription. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    SubscriptionId interface{}

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}

    // The push update's IP packet transport priority. This is made visible across
    // network hops to receiver. The transport priority is shared for all
    // receivers of a given subscription. The type is interface{} with range:
    // 0..63. The default value is 0.
    Dscp interface{}

    // Relative priority for a subscription.   Allows an underlying transport
    // layer perform informed load balance allocations between various
    // subscriptions. The type is interface{} with range: 0..255.
    SubscriptionPriority interface{}

    // Provides the Subscription ID of a parent subscription without which this
    // subscription should not exist. In other words, there is no reason to stream
    // these objects if another subscription is missing. The type is string.
    SubscriptionDependency interface{}
}

func (output *ModifySubscription_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "modify-subscription"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("subscription-result", types.YLeaf{"SubscriptionResult", output.SubscriptionResult})
    output.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", output.SubscriptionId})
    output.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", output.Stream})
    output.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", output.Encoding})
    output.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", output.Filter})
    output.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", output.FilterRef})
    output.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", output.SubtreeFilter})
    output.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", output.XpathFilter})
    output.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", output.StartTime})
    output.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", output.StopTime})
    output.EntityData.Leafs.Append("period", types.YLeaf{"Period", output.Period})
    output.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", output.AnchorTime})
    output.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", output.NoSynchOnStart})
    output.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", output.DampeningPeriod})
    output.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", output.ExcludedChange})
    output.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", output.Dscp})
    output.EntityData.Leafs.Append("subscription-priority", types.YLeaf{"SubscriptionPriority", output.SubscriptionPriority})
    output.EntityData.Leafs.Append("subscription-dependency", types.YLeaf{"SubscriptionDependency", output.SubscriptionDependency})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DeleteSubscription
// This RPC allows a subscriber to delete a subscription that
// was previously created using establish-subscription.
type DeleteSubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input DeleteSubscription_Input

    
    Output DeleteSubscription_Output
}

func (deleteSubscription *DeleteSubscription) GetEntityData() *types.CommonEntityData {
    deleteSubscription.EntityData.YFilter = deleteSubscription.YFilter
    deleteSubscription.EntityData.YangName = "delete-subscription"
    deleteSubscription.EntityData.BundleName = "ietf"
    deleteSubscription.EntityData.ParentYangName = "ietf-event-notifications"
    deleteSubscription.EntityData.SegmentPath = "ietf-event-notifications:delete-subscription"
    deleteSubscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    deleteSubscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    deleteSubscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    deleteSubscription.EntityData.Children = types.NewOrderedMap()
    deleteSubscription.EntityData.Children.Append("input", types.YChild{"Input", &deleteSubscription.Input})
    deleteSubscription.EntityData.Children.Append("output", types.YChild{"Output", &deleteSubscription.Output})
    deleteSubscription.EntityData.Leafs = types.NewOrderedMap()

    deleteSubscription.EntityData.YListKeys = []string {}

    return &(deleteSubscription.EntityData)
}

// DeleteSubscription_Input
type DeleteSubscription_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifier of the subscription that is to be deleted. Only subscriptions
    // that were created using establish-subscription can be deleted via this RPC.
    // The type is interface{} with range: 0..4294967295. This attribute is
    // mandatory.
    SubscriptionId interface{}
}

func (input *DeleteSubscription_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "delete-subscription"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", input.SubscriptionId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DeleteSubscription_Output
type DeleteSubscription_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether subscription is operational, or if a problem was
    // encountered. The type is one of the following:
    // ErrorErrorNoSuchOptionErrorNoSuchSubscriptionErrorOtherErrorInsufficientResourcesErrorConfiguredSubscriptionErrorDataNotAuthorizedOk.
    // This attribute is mandatory.
    SubscriptionResult interface{}
}

func (output *DeleteSubscription_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "delete-subscription"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("subscription-result", types.YLeaf{"SubscriptionResult", output.SubscriptionResult})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Streams
// This container contains a leaf list of built-in
// streams that are provided by the system.
type Streams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifies the built-in streams that are supported by the system.  Built-in
    // streams are associated with their own identities, each of which carries a
    // special semantics. In case configurable custom streams are supported, as
    // indicated by the custom-stream identity, the configuration of those custom
    // streams is provided         separately. The type is slice of ['NETCONF',
    // 'CustomStream', 'YangPush'].
    Stream []interface{}
}

func (streams *Streams) GetEntityData() *types.CommonEntityData {
    streams.EntityData.YFilter = streams.YFilter
    streams.EntityData.YangName = "streams"
    streams.EntityData.BundleName = "ietf"
    streams.EntityData.ParentYangName = "ietf-event-notifications"
    streams.EntityData.SegmentPath = "ietf-event-notifications:streams"
    streams.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    streams.EntityData.NamespaceTable = ietf.GetNamespaces()
    streams.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    streams.EntityData.Children = types.NewOrderedMap()
    streams.EntityData.Leafs = types.NewOrderedMap()
    streams.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", streams.Stream})

    streams.EntityData.YListKeys = []string {}

    return &(streams.EntityData)
}

// Filters
// This container contains a list of configurable filters
// that can be applied to subscriptions.  This facilitates
// the reuse of complex filters once defined.
type Filters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of configurable filters that can be applied to subscriptions. The
    // type is slice of Filters_Filter.
    Filter []*Filters_Filter
}

func (filters *Filters) GetEntityData() *types.CommonEntityData {
    filters.EntityData.YFilter = filters.YFilter
    filters.EntityData.YangName = "filters"
    filters.EntityData.BundleName = "ietf"
    filters.EntityData.ParentYangName = "ietf-event-notifications"
    filters.EntityData.SegmentPath = "ietf-event-notifications:filters"
    filters.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    filters.EntityData.NamespaceTable = ietf.GetNamespaces()
    filters.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    filters.EntityData.Children = types.NewOrderedMap()
    filters.EntityData.Children.Append("filter", types.YChild{"Filter", nil})
    for i := range filters.Filter {
        filters.EntityData.Children.Append(types.GetSegmentPath(filters.Filter[i]), types.YChild{"Filter", filters.Filter[i]})
    }
    filters.EntityData.Leafs = types.NewOrderedMap()

    filters.EntityData.YListKeys = []string {}

    return &(filters.EntityData)
}

// Filters_Filter
// A list of configurable filters that can be applied to
// subscriptions.
type Filters_Filter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An identifier to differentiate between filters.
    // The type is interface{} with range: 0..4294967295.
    FilterId interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}
}

func (filter *Filters_Filter) GetEntityData() *types.CommonEntityData {
    filter.EntityData.YFilter = filter.YFilter
    filter.EntityData.YangName = "filter"
    filter.EntityData.BundleName = "ietf"
    filter.EntityData.ParentYangName = "filters"
    filter.EntityData.SegmentPath = "filter" + types.AddKeyToken(filter.FilterId, "filter-id")
    filter.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    filter.EntityData.NamespaceTable = ietf.GetNamespaces()
    filter.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    filter.EntityData.Children = types.NewOrderedMap()
    filter.EntityData.Leafs = types.NewOrderedMap()
    filter.EntityData.Leafs.Append("filter-id", types.YLeaf{"FilterId", filter.FilterId})
    filter.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", filter.Filter})
    filter.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", filter.SubtreeFilter})
    filter.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", filter.XpathFilter})

    filter.EntityData.YListKeys = []string {"FilterId"}

    return &(filter.EntityData)
}

// SubscriptionConfig
// Contains the list of subscriptions that are configured,
// as opposed to established via RPC or other means.
type SubscriptionConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Content of a subscription. The type is slice of
    // SubscriptionConfig_Subscription.
    Subscription []*SubscriptionConfig_Subscription
}

func (subscriptionConfig *SubscriptionConfig) GetEntityData() *types.CommonEntityData {
    subscriptionConfig.EntityData.YFilter = subscriptionConfig.YFilter
    subscriptionConfig.EntityData.YangName = "subscription-config"
    subscriptionConfig.EntityData.BundleName = "ietf"
    subscriptionConfig.EntityData.ParentYangName = "ietf-event-notifications"
    subscriptionConfig.EntityData.SegmentPath = "ietf-event-notifications:subscription-config"
    subscriptionConfig.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    subscriptionConfig.EntityData.NamespaceTable = ietf.GetNamespaces()
    subscriptionConfig.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    subscriptionConfig.EntityData.Children = types.NewOrderedMap()
    subscriptionConfig.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range subscriptionConfig.Subscription {
        subscriptionConfig.EntityData.Children.Append(types.GetSegmentPath(subscriptionConfig.Subscription[i]), types.YChild{"Subscription", subscriptionConfig.Subscription[i]})
    }
    subscriptionConfig.EntityData.Leafs = types.NewOrderedMap()

    subscriptionConfig.EntityData.YListKeys = []string {}

    return &(subscriptionConfig.EntityData)
}

// SubscriptionConfig_Subscription
// Content of a subscription.
type SubscriptionConfig_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Identifier to use for this subscription. The type
    // is interface{} with range: 0..4294967295.
    SubscriptionId interface{}

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // References the interface for notifications. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    SourceInterface interface{}

    // Label of the vrf. The type is interface{} with range: 16..1048574.
    SourceVrf interface{}

    // The source address for the notifications. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    SourceAddress interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}

    // The push update's IP packet transport priority. This is made visible across
    // network hops to receiver. The transport priority is shared for all
    // receivers of a given subscription. The type is interface{} with range:
    // 0..63. The default value is 0.
    Dscp interface{}

    // Relative priority for a subscription.   Allows an underlying transport
    // layer perform informed load balance allocations between various
    // subscriptions. The type is interface{} with range: 0..255.
    SubscriptionPriority interface{}

    // Provides the Subscription ID of a parent subscription without which this
    // subscription should not exist. In other words, there is no reason to stream
    // these objects if another subscription is missing. The type is string.
    SubscriptionDependency interface{}

    // Set of receivers in a subscription.
    Receivers SubscriptionConfig_Subscription_Receivers
}

func (subscription *SubscriptionConfig_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "ietf"
    subscription.EntityData.ParentYangName = "subscription-config"
    subscription.EntityData.SegmentPath = "subscription" + types.AddKeyToken(subscription.SubscriptionId, "subscription-id")
    subscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    subscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("receivers", types.YChild{"Receivers", &subscription.Receivers})
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", subscription.SubscriptionId})
    subscription.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", subscription.Stream})
    subscription.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", subscription.Encoding})
    subscription.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", subscription.Filter})
    subscription.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", subscription.FilterRef})
    subscription.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", subscription.SubtreeFilter})
    subscription.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", subscription.XpathFilter})
    subscription.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", subscription.StartTime})
    subscription.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", subscription.StopTime})
    subscription.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", subscription.SourceInterface})
    subscription.EntityData.Leafs.Append("source-vrf", types.YLeaf{"SourceVrf", subscription.SourceVrf})
    subscription.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", subscription.SourceAddress})
    subscription.EntityData.Leafs.Append("period", types.YLeaf{"Period", subscription.Period})
    subscription.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", subscription.AnchorTime})
    subscription.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", subscription.NoSynchOnStart})
    subscription.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", subscription.DampeningPeriod})
    subscription.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", subscription.ExcludedChange})
    subscription.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", subscription.Dscp})
    subscription.EntityData.Leafs.Append("subscription-priority", types.YLeaf{"SubscriptionPriority", subscription.SubscriptionPriority})
    subscription.EntityData.Leafs.Append("subscription-dependency", types.YLeaf{"SubscriptionDependency", subscription.SubscriptionDependency})

    subscription.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(subscription.EntityData)
}

// SubscriptionConfig_Subscription_Receivers
// Set of receivers in a subscription.
type SubscriptionConfig_Subscription_Receivers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host or multipoint address intended as a target for the
    // notifications for a subscription. The type is slice of
    // SubscriptionConfig_Subscription_Receivers_Receiver.
    Receiver []*SubscriptionConfig_Subscription_Receivers_Receiver
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetEntityData() *types.CommonEntityData {
    receivers.EntityData.YFilter = receivers.YFilter
    receivers.EntityData.YangName = "receivers"
    receivers.EntityData.BundleName = "ietf"
    receivers.EntityData.ParentYangName = "subscription"
    receivers.EntityData.SegmentPath = "receivers"
    receivers.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    receivers.EntityData.NamespaceTable = ietf.GetNamespaces()
    receivers.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    receivers.EntityData.Children = types.NewOrderedMap()
    receivers.EntityData.Children.Append("receiver", types.YChild{"Receiver", nil})
    for i := range receivers.Receiver {
        receivers.EntityData.Children.Append(types.GetSegmentPath(receivers.Receiver[i]), types.YChild{"Receiver", receivers.Receiver[i]})
    }
    receivers.EntityData.Leafs = types.NewOrderedMap()

    receivers.EntityData.YListKeys = []string {}

    return &(receivers.EntityData)
}

// SubscriptionConfig_Subscription_Receivers_Receiver
// A single host or multipoint address intended as a target
// for the notifications for a subscription.
type SubscriptionConfig_Subscription_Receivers_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the address for the traffic to reach a
    // remote host. One of the following must be specified: an ipv4 address, an
    // ipv6 address, or a host name. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory.., or string with pattern:
    // ((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.
    // This attribute is mandatory..
    Address interface{}

    // This leaf specifies the port number to use for messages destined for a
    // receiver. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Port interface{}

    // This leaf specifies the transport protocol used to deliver messages
    // destined for the receiver. The type is one of the following: NetconfHttp2.
    // The default value is netconf.
    Protocol interface{}
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "ietf"
    receiver.EntityData.ParentYangName = "receivers"
    receiver.EntityData.SegmentPath = "receiver" + types.AddKeyToken(receiver.Address, "address")
    receiver.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    receiver.EntityData.NamespaceTable = ietf.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("address", types.YLeaf{"Address", receiver.Address})
    receiver.EntityData.Leafs.Append("port", types.YLeaf{"Port", receiver.Port})
    receiver.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", receiver.Protocol})

    receiver.EntityData.YListKeys = []string {"Address"}

    return &(receiver.EntityData)
}

// Subscriptions
// Contains the list of currently active subscriptions,
// i.e. subscriptions that are currently in effect,
// used for subscription management and monitoring purposes.
// This includes subscriptions that have been setup via RPC
// primitives, e.g. create-subscription, establish-
// subscription, and modify-subscription, as well as
// subscriptions that have been established via
//     configuration.
type Subscriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Content of a subscription. Subscriptions can be created using a control
    // channel or RPC, or be established through configuration. The type is slice
    // of Subscriptions_Subscription.
    Subscription []*Subscriptions_Subscription
}

func (subscriptions *Subscriptions) GetEntityData() *types.CommonEntityData {
    subscriptions.EntityData.YFilter = subscriptions.YFilter
    subscriptions.EntityData.YangName = "subscriptions"
    subscriptions.EntityData.BundleName = "ietf"
    subscriptions.EntityData.ParentYangName = "ietf-event-notifications"
    subscriptions.EntityData.SegmentPath = "ietf-event-notifications:subscriptions"
    subscriptions.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    subscriptions.EntityData.NamespaceTable = ietf.GetNamespaces()
    subscriptions.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    subscriptions.EntityData.Children = types.NewOrderedMap()
    subscriptions.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range subscriptions.Subscription {
        subscriptions.EntityData.Children.Append(types.GetSegmentPath(subscriptions.Subscription[i]), types.YChild{"Subscription", subscriptions.Subscription[i]})
    }
    subscriptions.EntityData.Leafs = types.NewOrderedMap()

    subscriptions.EntityData.YListKeys = []string {}

    return &(subscriptions.EntityData)
}

// Subscriptions_Subscription
// Content of a subscription.
// Subscriptions can be created using a control channel
// or RPC, or be established through configuration.
type Subscriptions_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Identifier of this subscription. The type is
    // interface{} with range: 0..4294967295.
    SubscriptionId interface{}

    // The presence of this leaf indicates that the subscription originated from
    // configuration, not through a control channel or RPC. The type is
    // interface{}.
    ConfiguredSubscription interface{}

    // The status of the subscription. The type is one of the following:
    // InactiveSuspendedInErrorActive.
    SubscriptionStatus interface{}

    // Indicates which stream of events is of interest. If not present, events in
    // the default NETCONF stream will be sent. The type is one of the following:
    // NETCONFCustomStreamYangPush.
    Stream interface{}

    // The type of encoding for the subscribed data. Default is XML. The type is
    // one of the following: EncodeJsonEncodeXml. The default value is encode-xml.
    Encoding interface{}

    // Filter per RFC 5277. Notification filter. If a filter element is specified
    // to look for data of a particular value, and the data item is not present
    // within a particular event notification for its value to be checked against,
    // the notification will be filtered out. For example, if one were to check
    // for 'severity=critical' in a configuration event notification where this
    // field was not supported, then the notification would be filtered out. For
    // subtree filtering, a non-empty node set means that the filter matches.  For
    // XPath filtering, the mechanisms defined in [XPATH] should be used to
    // convert the returned value to boolean. The type is string.
    Filter interface{}

    // References filter which is associated with the subscription. The type is
    // string with range: 0..4294967295. Refers to
    // event_notifications.Filters_Filter_FilterId
    FilterRef interface{}

    // Subtree-filter used to specify the data nodes targeted for subscription
    // within a subtree, or subtrees, of a conceptual YANG datastore.  Objects
    // matching the filter criteria will traverse the filter. The syntax follows
    // the subtree filter syntax specified in RFC 6241, section 6. The type is
    // string.
    SubtreeFilter interface{}

    // Xpath defining the data items of interest. The type is string.
    XpathFilter interface{}

    // Used to trigger the replay feature and indicate that the replay should
    // start at the time specified.  If <startTime> is not present, this is not a
    // replay subscription.  It is not valid to specify start times that are later
    // than the current time.  If the <startTime> specified is earlier than the
    // log can support, the replay will begin with the earliest available
    // notification.  This parameter is of type dateTime and compliant to
    // [RFC3339].  Implementations must support time zones. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StopTime interface{}

    // References the interface for notifications. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    SourceInterface interface{}

    // Label of the vrf. The type is interface{} with range: 16..1048574.
    SourceVrf interface{}

    // The source address for the notifications. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    SourceAddress interface{}

    // Duration of time which should occur between periodic push updates.  Where
    // the anchor of a start-time is available, the push will include the objects
    // and their values which exist at an exact multiple of timeticks aligning to
    // this start-time anchor. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    Period interface{}

    // Designates a timestamp from which the series of periodic push updates are
    // computed. The next update will take place at the next period interval from
    // the anchor time.  For example, for an anchor time at the top of a minute
    // and a period interval of a minute, the next update will be sent at the top
    // of the next minute. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    AnchorTime interface{}

    // This leaf acts as a flag that determines behavior at the start of the
    // subscription.  When present, synchronization of state at the beginning of
    // the subscription is outside the scope of the subscription. Only updates
    // about changes that are observed from the start time, i.e. only
    // push-change-update notifications are sent. When absent (default behavior),
    // in order to facilitate a receiver's synchronization, a full update is sent
    // when the subscription starts using a push-update notification, just like in
    // the case of a periodic subscription.  After that, push-change-update
    // notifications only are sent unless the Publisher chooses to resynch the
    // subscription again. The type is interface{}.
    NoSynchOnStart interface{}

    // Minimum amount of time that needs to have passed since the last time an
    // update was provided. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    DampeningPeriod interface{}

    // Use to restrict which changes trigger an update. For example, if modify is
    // excluded, only creation and deletion of objects is reported. The type is
    // slice of ChangeType.
    ExcludedChange []interface{}

    // The push update's IP packet transport priority. This is made visible across
    // network hops to receiver. The transport priority is shared for all
    // receivers of a given subscription. The type is interface{} with range:
    // 0..63. The default value is 0.
    Dscp interface{}

    // Relative priority for a subscription.   Allows an underlying transport
    // layer perform informed load balance allocations between various
    // subscriptions. The type is interface{} with range: 0..255.
    SubscriptionPriority interface{}

    // Provides the Subscription ID of a parent subscription without which this
    // subscription should not exist. In other words, there is no reason to stream
    // these objects if another subscription is missing. The type is string.
    SubscriptionDependency interface{}

    // Set of receivers in a subscription.
    Receivers Subscriptions_Subscription_Receivers
}

func (subscription *Subscriptions_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "ietf"
    subscription.EntityData.ParentYangName = "subscriptions"
    subscription.EntityData.SegmentPath = "subscription" + types.AddKeyToken(subscription.SubscriptionId, "subscription-id")
    subscription.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    subscription.EntityData.NamespaceTable = ietf.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("receivers", types.YChild{"Receivers", &subscription.Receivers})
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", subscription.SubscriptionId})
    subscription.EntityData.Leafs.Append("configured-subscription", types.YLeaf{"ConfiguredSubscription", subscription.ConfiguredSubscription})
    subscription.EntityData.Leafs.Append("subscription-status", types.YLeaf{"SubscriptionStatus", subscription.SubscriptionStatus})
    subscription.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", subscription.Stream})
    subscription.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", subscription.Encoding})
    subscription.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", subscription.Filter})
    subscription.EntityData.Leafs.Append("filter-ref", types.YLeaf{"FilterRef", subscription.FilterRef})
    subscription.EntityData.Leafs.Append("subtree-filter", types.YLeaf{"SubtreeFilter", subscription.SubtreeFilter})
    subscription.EntityData.Leafs.Append("xpath-filter", types.YLeaf{"XpathFilter", subscription.XpathFilter})
    subscription.EntityData.Leafs.Append("startTime", types.YLeaf{"StartTime", subscription.StartTime})
    subscription.EntityData.Leafs.Append("stopTime", types.YLeaf{"StopTime", subscription.StopTime})
    subscription.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", subscription.SourceInterface})
    subscription.EntityData.Leafs.Append("source-vrf", types.YLeaf{"SourceVrf", subscription.SourceVrf})
    subscription.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", subscription.SourceAddress})
    subscription.EntityData.Leafs.Append("period", types.YLeaf{"Period", subscription.Period})
    subscription.EntityData.Leafs.Append("anchor-time", types.YLeaf{"AnchorTime", subscription.AnchorTime})
    subscription.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", subscription.NoSynchOnStart})
    subscription.EntityData.Leafs.Append("dampening-period", types.YLeaf{"DampeningPeriod", subscription.DampeningPeriod})
    subscription.EntityData.Leafs.Append("excluded-change", types.YLeaf{"ExcludedChange", subscription.ExcludedChange})
    subscription.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", subscription.Dscp})
    subscription.EntityData.Leafs.Append("subscription-priority", types.YLeaf{"SubscriptionPriority", subscription.SubscriptionPriority})
    subscription.EntityData.Leafs.Append("subscription-dependency", types.YLeaf{"SubscriptionDependency", subscription.SubscriptionDependency})

    subscription.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(subscription.EntityData)
}

// Subscriptions_Subscription_Receivers
// Set of receivers in a subscription.
type Subscriptions_Subscription_Receivers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host or multipoint address intended as a target for the
    // notifications for a subscription. The type is slice of
    // Subscriptions_Subscription_Receivers_Receiver.
    Receiver []*Subscriptions_Subscription_Receivers_Receiver
}

func (receivers *Subscriptions_Subscription_Receivers) GetEntityData() *types.CommonEntityData {
    receivers.EntityData.YFilter = receivers.YFilter
    receivers.EntityData.YangName = "receivers"
    receivers.EntityData.BundleName = "ietf"
    receivers.EntityData.ParentYangName = "subscription"
    receivers.EntityData.SegmentPath = "receivers"
    receivers.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    receivers.EntityData.NamespaceTable = ietf.GetNamespaces()
    receivers.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    receivers.EntityData.Children = types.NewOrderedMap()
    receivers.EntityData.Children.Append("receiver", types.YChild{"Receiver", nil})
    for i := range receivers.Receiver {
        receivers.EntityData.Children.Append(types.GetSegmentPath(receivers.Receiver[i]), types.YChild{"Receiver", receivers.Receiver[i]})
    }
    receivers.EntityData.Leafs = types.NewOrderedMap()

    receivers.EntityData.YListKeys = []string {}

    return &(receivers.EntityData)
}

// Subscriptions_Subscription_Receivers_Receiver
// A single host or multipoint address intended as a target
// for the notifications for a subscription.
type Subscriptions_Subscription_Receivers_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the address for the traffic to reach a
    // remote host. One of the following must be specified: an ipv4 address, an
    // ipv6 address, or a host name. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory.., or string with pattern:
    // ((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.
    // This attribute is mandatory..
    Address interface{}

    // This leaf specifies the port number to use for messages destined for a
    // receiver. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Port interface{}

    // This leaf specifies the transport protocol used to deliver messages
    // destined for the receiver. The type is one of the following: NetconfHttp2.
    // The default value is netconf.
    Protocol interface{}
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "ietf"
    receiver.EntityData.ParentYangName = "receivers"
    receiver.EntityData.SegmentPath = "receiver" + types.AddKeyToken(receiver.Address, "address")
    receiver.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    receiver.EntityData.NamespaceTable = ietf.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("address", types.YLeaf{"Address", receiver.Address})
    receiver.EntityData.Leafs.Append("port", types.YLeaf{"Port", receiver.Port})
    receiver.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", receiver.Protocol})

    receiver.EntityData.YListKeys = []string {"Address"}

    return &(receiver.EntityData)
}

