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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EstablishSubscription_Input

    
    Output EstablishSubscription_Output
}

func (establishSubscription *EstablishSubscription) GetFilter() yfilter.YFilter { return establishSubscription.YFilter }

func (establishSubscription *EstablishSubscription) SetFilter(yf yfilter.YFilter) { establishSubscription.YFilter = yf }

func (establishSubscription *EstablishSubscription) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (establishSubscription *EstablishSubscription) GetSegmentPath() string {
    return "ietf-event-notifications:establish-subscription"
}

func (establishSubscription *EstablishSubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &establishSubscription.Input
    }
    if childYangName == "output" {
        return &establishSubscription.Output
    }
    return nil
}

func (establishSubscription *EstablishSubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &establishSubscription.Input
    children["output"] = &establishSubscription.Output
    return children
}

func (establishSubscription *EstablishSubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (establishSubscription *EstablishSubscription) GetBundleName() string { return "ietf" }

func (establishSubscription *EstablishSubscription) GetYangName() string { return "establish-subscription" }

func (establishSubscription *EstablishSubscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (establishSubscription *EstablishSubscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (establishSubscription *EstablishSubscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (establishSubscription *EstablishSubscription) SetParent(parent types.Entity) { establishSubscription.parent = parent }

func (establishSubscription *EstablishSubscription) GetParent() types.Entity { return establishSubscription.parent }

func (establishSubscription *EstablishSubscription) GetParentYangName() string { return "ietf-event-notifications" }

// EstablishSubscription_Input
type EstablishSubscription_Input struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (input *EstablishSubscription_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EstablishSubscription_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EstablishSubscription_Input) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    if yname == "dscp" { return "Dscp" }
    if yname == "subscription-priority" { return "SubscriptionPriority" }
    if yname == "subscription-dependency" { return "SubscriptionDependency" }
    return ""
}

func (input *EstablishSubscription_Input) GetSegmentPath() string {
    return "input"
}

func (input *EstablishSubscription_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EstablishSubscription_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EstablishSubscription_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream"] = input.Stream
    leafs["encoding"] = input.Encoding
    leafs["filter"] = input.Filter
    leafs["filter-ref"] = input.FilterRef
    leafs["subtree-filter"] = input.SubtreeFilter
    leafs["xpath-filter"] = input.XpathFilter
    leafs["startTime"] = input.Starttime
    leafs["stopTime"] = input.Stoptime
    leafs["period"] = input.Period
    leafs["anchor-time"] = input.AnchorTime
    leafs["no-synch-on-start"] = input.NoSynchOnStart
    leafs["dampening-period"] = input.DampeningPeriod
    leafs["excluded-change"] = input.ExcludedChange
    leafs["dscp"] = input.Dscp
    leafs["subscription-priority"] = input.SubscriptionPriority
    leafs["subscription-dependency"] = input.SubscriptionDependency
    return leafs
}

func (input *EstablishSubscription_Input) GetBundleName() string { return "ietf" }

func (input *EstablishSubscription_Input) GetYangName() string { return "input" }

func (input *EstablishSubscription_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *EstablishSubscription_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *EstablishSubscription_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *EstablishSubscription_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EstablishSubscription_Input) GetParent() types.Entity { return input.parent }

func (input *EstablishSubscription_Input) GetParentYangName() string { return "establish-subscription" }

// EstablishSubscription_Output
type EstablishSubscription_Output struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (output *EstablishSubscription_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *EstablishSubscription_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *EstablishSubscription_Output) GetGoName(yname string) string {
    if yname == "subscription-result" { return "SubscriptionResult" }
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    if yname == "dscp" { return "Dscp" }
    if yname == "subscription-priority" { return "SubscriptionPriority" }
    if yname == "subscription-dependency" { return "SubscriptionDependency" }
    return ""
}

func (output *EstablishSubscription_Output) GetSegmentPath() string {
    return "output"
}

func (output *EstablishSubscription_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *EstablishSubscription_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *EstablishSubscription_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-result"] = output.SubscriptionResult
    leafs["subscription-id"] = output.SubscriptionId
    leafs["stream"] = output.Stream
    leafs["encoding"] = output.Encoding
    leafs["filter"] = output.Filter
    leafs["filter-ref"] = output.FilterRef
    leafs["subtree-filter"] = output.SubtreeFilter
    leafs["xpath-filter"] = output.XpathFilter
    leafs["startTime"] = output.Starttime
    leafs["stopTime"] = output.Stoptime
    leafs["period"] = output.Period
    leafs["anchor-time"] = output.AnchorTime
    leafs["no-synch-on-start"] = output.NoSynchOnStart
    leafs["dampening-period"] = output.DampeningPeriod
    leafs["excluded-change"] = output.ExcludedChange
    leafs["dscp"] = output.Dscp
    leafs["subscription-priority"] = output.SubscriptionPriority
    leafs["subscription-dependency"] = output.SubscriptionDependency
    return leafs
}

func (output *EstablishSubscription_Output) GetBundleName() string { return "ietf" }

func (output *EstablishSubscription_Output) GetYangName() string { return "output" }

func (output *EstablishSubscription_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *EstablishSubscription_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *EstablishSubscription_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *EstablishSubscription_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *EstablishSubscription_Output) GetParent() types.Entity { return output.parent }

func (output *EstablishSubscription_Output) GetParentYangName() string { return "establish-subscription" }

// CreateSubscription
// This operation initiates an event notification subscription
// that will send asynchronous event notifications to the
// initiator of the command until the association terminates.
// It is not possible to modify or delete a subscription
// that was created using this operation.  It is included for
// reasons of backward compatibility with RFC 5277
// implementations.
type CreateSubscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CreateSubscription_Input
}

func (createSubscription *CreateSubscription) GetFilter() yfilter.YFilter { return createSubscription.YFilter }

func (createSubscription *CreateSubscription) SetFilter(yf yfilter.YFilter) { createSubscription.YFilter = yf }

func (createSubscription *CreateSubscription) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (createSubscription *CreateSubscription) GetSegmentPath() string {
    return "ietf-event-notifications:create-subscription"
}

func (createSubscription *CreateSubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &createSubscription.Input
    }
    return nil
}

func (createSubscription *CreateSubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &createSubscription.Input
    return children
}

func (createSubscription *CreateSubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (createSubscription *CreateSubscription) GetBundleName() string { return "ietf" }

func (createSubscription *CreateSubscription) GetYangName() string { return "create-subscription" }

func (createSubscription *CreateSubscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (createSubscription *CreateSubscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (createSubscription *CreateSubscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (createSubscription *CreateSubscription) SetParent(parent types.Entity) { createSubscription.parent = parent }

func (createSubscription *CreateSubscription) GetParent() types.Entity { return createSubscription.parent }

func (createSubscription *CreateSubscription) GetParentYangName() string { return "ietf-event-notifications" }

// CreateSubscription_Input
type CreateSubscription_Input struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}
}

func (input *CreateSubscription_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CreateSubscription_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CreateSubscription_Input) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    return ""
}

func (input *CreateSubscription_Input) GetSegmentPath() string {
    return "input"
}

func (input *CreateSubscription_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CreateSubscription_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CreateSubscription_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream"] = input.Stream
    leafs["encoding"] = input.Encoding
    leafs["filter"] = input.Filter
    leafs["startTime"] = input.Starttime
    leafs["stopTime"] = input.Stoptime
    return leafs
}

func (input *CreateSubscription_Input) GetBundleName() string { return "ietf" }

func (input *CreateSubscription_Input) GetYangName() string { return "input" }

func (input *CreateSubscription_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *CreateSubscription_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *CreateSubscription_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *CreateSubscription_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CreateSubscription_Input) GetParent() types.Entity { return input.parent }

func (input *CreateSubscription_Input) GetParentYangName() string { return "create-subscription" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ModifySubscription_Input

    
    Output ModifySubscription_Output
}

func (modifySubscription *ModifySubscription) GetFilter() yfilter.YFilter { return modifySubscription.YFilter }

func (modifySubscription *ModifySubscription) SetFilter(yf yfilter.YFilter) { modifySubscription.YFilter = yf }

func (modifySubscription *ModifySubscription) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (modifySubscription *ModifySubscription) GetSegmentPath() string {
    return "ietf-event-notifications:modify-subscription"
}

func (modifySubscription *ModifySubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &modifySubscription.Input
    }
    if childYangName == "output" {
        return &modifySubscription.Output
    }
    return nil
}

func (modifySubscription *ModifySubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &modifySubscription.Input
    children["output"] = &modifySubscription.Output
    return children
}

func (modifySubscription *ModifySubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (modifySubscription *ModifySubscription) GetBundleName() string { return "ietf" }

func (modifySubscription *ModifySubscription) GetYangName() string { return "modify-subscription" }

func (modifySubscription *ModifySubscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (modifySubscription *ModifySubscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (modifySubscription *ModifySubscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (modifySubscription *ModifySubscription) SetParent(parent types.Entity) { modifySubscription.parent = parent }

func (modifySubscription *ModifySubscription) GetParent() types.Entity { return modifySubscription.parent }

func (modifySubscription *ModifySubscription) GetParentYangName() string { return "ietf-event-notifications" }

// ModifySubscription_Input
type ModifySubscription_Input struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (input *ModifySubscription_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ModifySubscription_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ModifySubscription_Input) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    return ""
}

func (input *ModifySubscription_Input) GetSegmentPath() string {
    return "input"
}

func (input *ModifySubscription_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *ModifySubscription_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *ModifySubscription_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = input.SubscriptionId
    leafs["filter"] = input.Filter
    leafs["filter-ref"] = input.FilterRef
    leafs["subtree-filter"] = input.SubtreeFilter
    leafs["xpath-filter"] = input.XpathFilter
    leafs["startTime"] = input.Starttime
    leafs["stopTime"] = input.Stoptime
    leafs["period"] = input.Period
    leafs["anchor-time"] = input.AnchorTime
    leafs["no-synch-on-start"] = input.NoSynchOnStart
    leafs["dampening-period"] = input.DampeningPeriod
    leafs["excluded-change"] = input.ExcludedChange
    return leafs
}

func (input *ModifySubscription_Input) GetBundleName() string { return "ietf" }

func (input *ModifySubscription_Input) GetYangName() string { return "input" }

func (input *ModifySubscription_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *ModifySubscription_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *ModifySubscription_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *ModifySubscription_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ModifySubscription_Input) GetParent() types.Entity { return input.parent }

func (input *ModifySubscription_Input) GetParentYangName() string { return "modify-subscription" }

// ModifySubscription_Output
type ModifySubscription_Output struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (output *ModifySubscription_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *ModifySubscription_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *ModifySubscription_Output) GetGoName(yname string) string {
    if yname == "subscription-result" { return "SubscriptionResult" }
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    if yname == "dscp" { return "Dscp" }
    if yname == "subscription-priority" { return "SubscriptionPriority" }
    if yname == "subscription-dependency" { return "SubscriptionDependency" }
    return ""
}

func (output *ModifySubscription_Output) GetSegmentPath() string {
    return "output"
}

func (output *ModifySubscription_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *ModifySubscription_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *ModifySubscription_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-result"] = output.SubscriptionResult
    leafs["subscription-id"] = output.SubscriptionId
    leafs["stream"] = output.Stream
    leafs["encoding"] = output.Encoding
    leafs["filter"] = output.Filter
    leafs["filter-ref"] = output.FilterRef
    leafs["subtree-filter"] = output.SubtreeFilter
    leafs["xpath-filter"] = output.XpathFilter
    leafs["startTime"] = output.Starttime
    leafs["stopTime"] = output.Stoptime
    leafs["period"] = output.Period
    leafs["anchor-time"] = output.AnchorTime
    leafs["no-synch-on-start"] = output.NoSynchOnStart
    leafs["dampening-period"] = output.DampeningPeriod
    leafs["excluded-change"] = output.ExcludedChange
    leafs["dscp"] = output.Dscp
    leafs["subscription-priority"] = output.SubscriptionPriority
    leafs["subscription-dependency"] = output.SubscriptionDependency
    return leafs
}

func (output *ModifySubscription_Output) GetBundleName() string { return "ietf" }

func (output *ModifySubscription_Output) GetYangName() string { return "output" }

func (output *ModifySubscription_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *ModifySubscription_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *ModifySubscription_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *ModifySubscription_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *ModifySubscription_Output) GetParent() types.Entity { return output.parent }

func (output *ModifySubscription_Output) GetParentYangName() string { return "modify-subscription" }

// DeleteSubscription
// This RPC allows a subscriber to delete a subscription that
// was previously created using establish-subscription.
type DeleteSubscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input DeleteSubscription_Input

    
    Output DeleteSubscription_Output
}

func (deleteSubscription *DeleteSubscription) GetFilter() yfilter.YFilter { return deleteSubscription.YFilter }

func (deleteSubscription *DeleteSubscription) SetFilter(yf yfilter.YFilter) { deleteSubscription.YFilter = yf }

func (deleteSubscription *DeleteSubscription) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (deleteSubscription *DeleteSubscription) GetSegmentPath() string {
    return "ietf-event-notifications:delete-subscription"
}

func (deleteSubscription *DeleteSubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &deleteSubscription.Input
    }
    if childYangName == "output" {
        return &deleteSubscription.Output
    }
    return nil
}

func (deleteSubscription *DeleteSubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &deleteSubscription.Input
    children["output"] = &deleteSubscription.Output
    return children
}

func (deleteSubscription *DeleteSubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (deleteSubscription *DeleteSubscription) GetBundleName() string { return "ietf" }

func (deleteSubscription *DeleteSubscription) GetYangName() string { return "delete-subscription" }

func (deleteSubscription *DeleteSubscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (deleteSubscription *DeleteSubscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (deleteSubscription *DeleteSubscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (deleteSubscription *DeleteSubscription) SetParent(parent types.Entity) { deleteSubscription.parent = parent }

func (deleteSubscription *DeleteSubscription) GetParent() types.Entity { return deleteSubscription.parent }

func (deleteSubscription *DeleteSubscription) GetParentYangName() string { return "ietf-event-notifications" }

// DeleteSubscription_Input
type DeleteSubscription_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Identifier of the subscription that is to be deleted. Only subscriptions
    // that were created using establish-subscription can be deleted via this RPC.
    // The type is interface{} with range: 0..4294967295. This attribute is
    // mandatory.
    SubscriptionId interface{}
}

func (input *DeleteSubscription_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *DeleteSubscription_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *DeleteSubscription_Input) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    return ""
}

func (input *DeleteSubscription_Input) GetSegmentPath() string {
    return "input"
}

func (input *DeleteSubscription_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *DeleteSubscription_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *DeleteSubscription_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = input.SubscriptionId
    return leafs
}

func (input *DeleteSubscription_Input) GetBundleName() string { return "ietf" }

func (input *DeleteSubscription_Input) GetYangName() string { return "input" }

func (input *DeleteSubscription_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *DeleteSubscription_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *DeleteSubscription_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *DeleteSubscription_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *DeleteSubscription_Input) GetParent() types.Entity { return input.parent }

func (input *DeleteSubscription_Input) GetParentYangName() string { return "delete-subscription" }

// DeleteSubscription_Output
type DeleteSubscription_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates whether subscription is operational, or if a problem was
    // encountered. The type is one of the following:
    // ErrorErrorNoSuchOptionErrorNoSuchSubscriptionErrorOtherErrorInsufficientResourcesErrorConfiguredSubscriptionErrorDataNotAuthorizedOk.
    // This attribute is mandatory.
    SubscriptionResult interface{}
}

func (output *DeleteSubscription_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DeleteSubscription_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DeleteSubscription_Output) GetGoName(yname string) string {
    if yname == "subscription-result" { return "SubscriptionResult" }
    return ""
}

func (output *DeleteSubscription_Output) GetSegmentPath() string {
    return "output"
}

func (output *DeleteSubscription_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DeleteSubscription_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DeleteSubscription_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-result"] = output.SubscriptionResult
    return leafs
}

func (output *DeleteSubscription_Output) GetBundleName() string { return "ietf" }

func (output *DeleteSubscription_Output) GetYangName() string { return "output" }

func (output *DeleteSubscription_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *DeleteSubscription_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *DeleteSubscription_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *DeleteSubscription_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DeleteSubscription_Output) GetParent() types.Entity { return output.parent }

func (output *DeleteSubscription_Output) GetParentYangName() string { return "delete-subscription" }

// Streams
// This container contains a leaf list of built-in
// streams that are provided by the system.
type Streams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Identifies the built-in streams that are supported by the system.  Built-in
    // streams are associated with their own identities, each of which carries a
    // special semantics. In case configurable custom streams are supported, as
    // indicated by the custom-stream identity, the configuration of those custom
    // streams is provided         separately. The type is slice of ['NETCONF',
    // 'CustomStream', 'YangPush'].
    Stream []interface{}
}

func (streams *Streams) GetFilter() yfilter.YFilter { return streams.YFilter }

func (streams *Streams) SetFilter(yf yfilter.YFilter) { streams.YFilter = yf }

func (streams *Streams) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    return ""
}

func (streams *Streams) GetSegmentPath() string {
    return "ietf-event-notifications:streams"
}

func (streams *Streams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (streams *Streams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (streams *Streams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream"] = streams.Stream
    return leafs
}

func (streams *Streams) GetBundleName() string { return "ietf" }

func (streams *Streams) GetYangName() string { return "streams" }

func (streams *Streams) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (streams *Streams) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (streams *Streams) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (streams *Streams) SetParent(parent types.Entity) { streams.parent = parent }

func (streams *Streams) GetParent() types.Entity { return streams.parent }

func (streams *Streams) GetParentYangName() string { return "ietf-event-notifications" }

// Filters
// This container contains a list of configurable filters
// that can be applied to subscriptions.  This facilitates
// the reuse of complex filters once defined.
type Filters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of configurable filters that can be applied to subscriptions. The
    // type is slice of Filters_Filter.
    Filter []Filters_Filter
}

func (filters *Filters) GetFilter() yfilter.YFilter { return filters.YFilter }

func (filters *Filters) SetFilter(yf yfilter.YFilter) { filters.YFilter = yf }

func (filters *Filters) GetGoName(yname string) string {
    if yname == "filter" { return "Filter" }
    return ""
}

func (filters *Filters) GetSegmentPath() string {
    return "ietf-event-notifications:filters"
}

func (filters *Filters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "filter" {
        for _, c := range filters.Filter {
            if filters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Filters_Filter{}
        filters.Filter = append(filters.Filter, child)
        return &filters.Filter[len(filters.Filter)-1]
    }
    return nil
}

func (filters *Filters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range filters.Filter {
        children[filters.Filter[i].GetSegmentPath()] = &filters.Filter[i]
    }
    return children
}

func (filters *Filters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (filters *Filters) GetBundleName() string { return "ietf" }

func (filters *Filters) GetYangName() string { return "filters" }

func (filters *Filters) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (filters *Filters) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (filters *Filters) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (filters *Filters) SetParent(parent types.Entity) { filters.parent = parent }

func (filters *Filters) GetParent() types.Entity { return filters.parent }

func (filters *Filters) GetParentYangName() string { return "ietf-event-notifications" }

// Filters_Filter
// A list of configurable filters that can be applied to
// subscriptions.
type Filters_Filter struct {
    parent types.Entity
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

func (filter *Filters_Filter) GetFilter() yfilter.YFilter { return filter.YFilter }

func (filter *Filters_Filter) SetFilter(yf yfilter.YFilter) { filter.YFilter = yf }

func (filter *Filters_Filter) GetGoName(yname string) string {
    if yname == "filter-id" { return "FilterId" }
    if yname == "filter" { return "Filter" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    return ""
}

func (filter *Filters_Filter) GetSegmentPath() string {
    return "filter" + "[filter-id='" + fmt.Sprintf("%v", filter.FilterId) + "']"
}

func (filter *Filters_Filter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (filter *Filters_Filter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (filter *Filters_Filter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter-id"] = filter.FilterId
    leafs["filter"] = filter.Filter
    leafs["subtree-filter"] = filter.SubtreeFilter
    leafs["xpath-filter"] = filter.XpathFilter
    return leafs
}

func (filter *Filters_Filter) GetBundleName() string { return "ietf" }

func (filter *Filters_Filter) GetYangName() string { return "filter" }

func (filter *Filters_Filter) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (filter *Filters_Filter) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (filter *Filters_Filter) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (filter *Filters_Filter) SetParent(parent types.Entity) { filter.parent = parent }

func (filter *Filters_Filter) GetParent() types.Entity { return filter.parent }

func (filter *Filters_Filter) GetParentYangName() string { return "filters" }

// SubscriptionConfig
// Contains the list of subscriptions that are configured,
// as opposed to established via RPC or other means.
type SubscriptionConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Content of a subscription. The type is slice of
    // SubscriptionConfig_Subscription.
    Subscription []SubscriptionConfig_Subscription
}

func (subscriptionConfig *SubscriptionConfig) GetFilter() yfilter.YFilter { return subscriptionConfig.YFilter }

func (subscriptionConfig *SubscriptionConfig) SetFilter(yf yfilter.YFilter) { subscriptionConfig.YFilter = yf }

func (subscriptionConfig *SubscriptionConfig) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (subscriptionConfig *SubscriptionConfig) GetSegmentPath() string {
    return "ietf-event-notifications:subscription-config"
}

func (subscriptionConfig *SubscriptionConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range subscriptionConfig.Subscription {
            if subscriptionConfig.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriptionConfig_Subscription{}
        subscriptionConfig.Subscription = append(subscriptionConfig.Subscription, child)
        return &subscriptionConfig.Subscription[len(subscriptionConfig.Subscription)-1]
    }
    return nil
}

func (subscriptionConfig *SubscriptionConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriptionConfig.Subscription {
        children[subscriptionConfig.Subscription[i].GetSegmentPath()] = &subscriptionConfig.Subscription[i]
    }
    return children
}

func (subscriptionConfig *SubscriptionConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriptionConfig *SubscriptionConfig) GetBundleName() string { return "ietf" }

func (subscriptionConfig *SubscriptionConfig) GetYangName() string { return "subscription-config" }

func (subscriptionConfig *SubscriptionConfig) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (subscriptionConfig *SubscriptionConfig) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (subscriptionConfig *SubscriptionConfig) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (subscriptionConfig *SubscriptionConfig) SetParent(parent types.Entity) { subscriptionConfig.parent = parent }

func (subscriptionConfig *SubscriptionConfig) GetParent() types.Entity { return subscriptionConfig.parent }

func (subscriptionConfig *SubscriptionConfig) GetParentYangName() string { return "ietf-event-notifications" }

// SubscriptionConfig_Subscription
// Content of a subscription.
type SubscriptionConfig_Subscription struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (subscription *SubscriptionConfig_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *SubscriptionConfig_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *SubscriptionConfig_Subscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "source-vrf" { return "SourceVrf" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    if yname == "dscp" { return "Dscp" }
    if yname == "subscription-priority" { return "SubscriptionPriority" }
    if yname == "subscription-dependency" { return "SubscriptionDependency" }
    if yname == "receivers" { return "Receivers" }
    return ""
}

func (subscription *SubscriptionConfig_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
}

func (subscription *SubscriptionConfig_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "receivers" {
        return &subscription.Receivers
    }
    return nil
}

func (subscription *SubscriptionConfig_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["receivers"] = &subscription.Receivers
    return children
}

func (subscription *SubscriptionConfig_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = subscription.SubscriptionId
    leafs["stream"] = subscription.Stream
    leafs["encoding"] = subscription.Encoding
    leafs["filter"] = subscription.Filter
    leafs["filter-ref"] = subscription.FilterRef
    leafs["subtree-filter"] = subscription.SubtreeFilter
    leafs["xpath-filter"] = subscription.XpathFilter
    leafs["startTime"] = subscription.Starttime
    leafs["stopTime"] = subscription.Stoptime
    leafs["source-interface"] = subscription.SourceInterface
    leafs["source-vrf"] = subscription.SourceVrf
    leafs["source-address"] = subscription.SourceAddress
    leafs["period"] = subscription.Period
    leafs["anchor-time"] = subscription.AnchorTime
    leafs["no-synch-on-start"] = subscription.NoSynchOnStart
    leafs["dampening-period"] = subscription.DampeningPeriod
    leafs["excluded-change"] = subscription.ExcludedChange
    leafs["dscp"] = subscription.Dscp
    leafs["subscription-priority"] = subscription.SubscriptionPriority
    leafs["subscription-dependency"] = subscription.SubscriptionDependency
    return leafs
}

func (subscription *SubscriptionConfig_Subscription) GetBundleName() string { return "ietf" }

func (subscription *SubscriptionConfig_Subscription) GetYangName() string { return "subscription" }

func (subscription *SubscriptionConfig_Subscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (subscription *SubscriptionConfig_Subscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (subscription *SubscriptionConfig_Subscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (subscription *SubscriptionConfig_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *SubscriptionConfig_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *SubscriptionConfig_Subscription) GetParentYangName() string { return "subscription-config" }

// SubscriptionConfig_Subscription_Receivers
// Set of receivers in a subscription.
type SubscriptionConfig_Subscription_Receivers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single host or multipoint address intended as a target for the
    // notifications for a subscription. The type is slice of
    // SubscriptionConfig_Subscription_Receivers_Receiver.
    Receiver []SubscriptionConfig_Subscription_Receivers_Receiver
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetFilter() yfilter.YFilter { return receivers.YFilter }

func (receivers *SubscriptionConfig_Subscription_Receivers) SetFilter(yf yfilter.YFilter) { receivers.YFilter = yf }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetGoName(yname string) string {
    if yname == "receiver" { return "Receiver" }
    return ""
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetSegmentPath() string {
    return "receivers"
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "receiver" {
        for _, c := range receivers.Receiver {
            if receivers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriptionConfig_Subscription_Receivers_Receiver{}
        receivers.Receiver = append(receivers.Receiver, child)
        return &receivers.Receiver[len(receivers.Receiver)-1]
    }
    return nil
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range receivers.Receiver {
        children[receivers.Receiver[i].GetSegmentPath()] = &receivers.Receiver[i]
    }
    return children
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (receivers *SubscriptionConfig_Subscription_Receivers) GetBundleName() string { return "ietf" }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetYangName() string { return "receivers" }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (receivers *SubscriptionConfig_Subscription_Receivers) SetParent(parent types.Entity) { receivers.parent = parent }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetParent() types.Entity { return receivers.parent }

func (receivers *SubscriptionConfig_Subscription_Receivers) GetParentYangName() string { return "subscription" }

// SubscriptionConfig_Subscription_Receivers_Receiver
// A single host or multipoint address intended as a target
// for the notifications for a subscription.
type SubscriptionConfig_Subscription_Receivers_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the address for the traffic to reach a
    // remote host. One of the following must be specified: an ipv4 address, an
    // ipv6 address, or a host name. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory.., or string with length: 1..253 This attribute
    // is mandatory..
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

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetSegmentPath() string {
    return "receiver" + "[address='" + fmt.Sprintf("%v", receiver.Address) + "']"
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = receiver.Address
    leafs["port"] = receiver.Port
    leafs["protocol"] = receiver.Protocol
    return leafs
}

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetBundleName() string { return "ietf" }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetYangName() string { return "receiver" }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *SubscriptionConfig_Subscription_Receivers_Receiver) GetParentYangName() string { return "receivers" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Content of a subscription. Subscriptions can be created using a control
    // channel or RPC, or be established through configuration. The type is slice
    // of Subscriptions_Subscription.
    Subscription []Subscriptions_Subscription
}

func (subscriptions *Subscriptions) GetFilter() yfilter.YFilter { return subscriptions.YFilter }

func (subscriptions *Subscriptions) SetFilter(yf yfilter.YFilter) { subscriptions.YFilter = yf }

func (subscriptions *Subscriptions) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (subscriptions *Subscriptions) GetSegmentPath() string {
    return "ietf-event-notifications:subscriptions"
}

func (subscriptions *Subscriptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range subscriptions.Subscription {
            if subscriptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Subscriptions_Subscription{}
        subscriptions.Subscription = append(subscriptions.Subscription, child)
        return &subscriptions.Subscription[len(subscriptions.Subscription)-1]
    }
    return nil
}

func (subscriptions *Subscriptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriptions.Subscription {
        children[subscriptions.Subscription[i].GetSegmentPath()] = &subscriptions.Subscription[i]
    }
    return children
}

func (subscriptions *Subscriptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriptions *Subscriptions) GetBundleName() string { return "ietf" }

func (subscriptions *Subscriptions) GetYangName() string { return "subscriptions" }

func (subscriptions *Subscriptions) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (subscriptions *Subscriptions) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (subscriptions *Subscriptions) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (subscriptions *Subscriptions) SetParent(parent types.Entity) { subscriptions.parent = parent }

func (subscriptions *Subscriptions) GetParent() types.Entity { return subscriptions.parent }

func (subscriptions *Subscriptions) GetParentYangName() string { return "ietf-event-notifications" }

// Subscriptions_Subscription
// Content of a subscription.
// Subscriptions can be created using a control channel
// or RPC, or be established through configuration.
type Subscriptions_Subscription struct {
    parent types.Entity
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
    Starttime interface{}

    // Used with the optional replay feature to indicate the newest notifications
    // of interest.  If <stopTime> is not present, the notifications will continue
    // until the subscription is terminated.  Must be used with and be later than
    // <startTime>.  Values of <stopTime> in the future are valid.  This parameter
    // is of type dateTime and compliant to [RFC3339].  Implementations must
    // support time zones. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Stoptime interface{}

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

func (subscription *Subscriptions_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *Subscriptions_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *Subscriptions_Subscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "configured-subscription" { return "ConfiguredSubscription" }
    if yname == "subscription-status" { return "SubscriptionStatus" }
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "filter" { return "Filter" }
    if yname == "filter-ref" { return "FilterRef" }
    if yname == "subtree-filter" { return "SubtreeFilter" }
    if yname == "xpath-filter" { return "XpathFilter" }
    if yname == "startTime" { return "Starttime" }
    if yname == "stopTime" { return "Stoptime" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "source-vrf" { return "SourceVrf" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "period" { return "Period" }
    if yname == "anchor-time" { return "AnchorTime" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "dampening-period" { return "DampeningPeriod" }
    if yname == "excluded-change" { return "ExcludedChange" }
    if yname == "dscp" { return "Dscp" }
    if yname == "subscription-priority" { return "SubscriptionPriority" }
    if yname == "subscription-dependency" { return "SubscriptionDependency" }
    if yname == "receivers" { return "Receivers" }
    return ""
}

func (subscription *Subscriptions_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
}

func (subscription *Subscriptions_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "receivers" {
        return &subscription.Receivers
    }
    return nil
}

func (subscription *Subscriptions_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["receivers"] = &subscription.Receivers
    return children
}

func (subscription *Subscriptions_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = subscription.SubscriptionId
    leafs["configured-subscription"] = subscription.ConfiguredSubscription
    leafs["subscription-status"] = subscription.SubscriptionStatus
    leafs["stream"] = subscription.Stream
    leafs["encoding"] = subscription.Encoding
    leafs["filter"] = subscription.Filter
    leafs["filter-ref"] = subscription.FilterRef
    leafs["subtree-filter"] = subscription.SubtreeFilter
    leafs["xpath-filter"] = subscription.XpathFilter
    leafs["startTime"] = subscription.Starttime
    leafs["stopTime"] = subscription.Stoptime
    leafs["source-interface"] = subscription.SourceInterface
    leafs["source-vrf"] = subscription.SourceVrf
    leafs["source-address"] = subscription.SourceAddress
    leafs["period"] = subscription.Period
    leafs["anchor-time"] = subscription.AnchorTime
    leafs["no-synch-on-start"] = subscription.NoSynchOnStart
    leafs["dampening-period"] = subscription.DampeningPeriod
    leafs["excluded-change"] = subscription.ExcludedChange
    leafs["dscp"] = subscription.Dscp
    leafs["subscription-priority"] = subscription.SubscriptionPriority
    leafs["subscription-dependency"] = subscription.SubscriptionDependency
    return leafs
}

func (subscription *Subscriptions_Subscription) GetBundleName() string { return "ietf" }

func (subscription *Subscriptions_Subscription) GetYangName() string { return "subscription" }

func (subscription *Subscriptions_Subscription) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (subscription *Subscriptions_Subscription) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (subscription *Subscriptions_Subscription) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (subscription *Subscriptions_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *Subscriptions_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *Subscriptions_Subscription) GetParentYangName() string { return "subscriptions" }

// Subscriptions_Subscription_Receivers
// Set of receivers in a subscription.
type Subscriptions_Subscription_Receivers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single host or multipoint address intended as a target for the
    // notifications for a subscription. The type is slice of
    // Subscriptions_Subscription_Receivers_Receiver.
    Receiver []Subscriptions_Subscription_Receivers_Receiver
}

func (receivers *Subscriptions_Subscription_Receivers) GetFilter() yfilter.YFilter { return receivers.YFilter }

func (receivers *Subscriptions_Subscription_Receivers) SetFilter(yf yfilter.YFilter) { receivers.YFilter = yf }

func (receivers *Subscriptions_Subscription_Receivers) GetGoName(yname string) string {
    if yname == "receiver" { return "Receiver" }
    return ""
}

func (receivers *Subscriptions_Subscription_Receivers) GetSegmentPath() string {
    return "receivers"
}

func (receivers *Subscriptions_Subscription_Receivers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "receiver" {
        for _, c := range receivers.Receiver {
            if receivers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Subscriptions_Subscription_Receivers_Receiver{}
        receivers.Receiver = append(receivers.Receiver, child)
        return &receivers.Receiver[len(receivers.Receiver)-1]
    }
    return nil
}

func (receivers *Subscriptions_Subscription_Receivers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range receivers.Receiver {
        children[receivers.Receiver[i].GetSegmentPath()] = &receivers.Receiver[i]
    }
    return children
}

func (receivers *Subscriptions_Subscription_Receivers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (receivers *Subscriptions_Subscription_Receivers) GetBundleName() string { return "ietf" }

func (receivers *Subscriptions_Subscription_Receivers) GetYangName() string { return "receivers" }

func (receivers *Subscriptions_Subscription_Receivers) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (receivers *Subscriptions_Subscription_Receivers) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (receivers *Subscriptions_Subscription_Receivers) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (receivers *Subscriptions_Subscription_Receivers) SetParent(parent types.Entity) { receivers.parent = parent }

func (receivers *Subscriptions_Subscription_Receivers) GetParent() types.Entity { return receivers.parent }

func (receivers *Subscriptions_Subscription_Receivers) GetParentYangName() string { return "subscription" }

// Subscriptions_Subscription_Receivers_Receiver
// A single host or multipoint address intended as a target
// for the notifications for a subscription.
type Subscriptions_Subscription_Receivers_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the address for the traffic to reach a
    // remote host. One of the following must be specified: an ipv4 address, an
    // ipv6 address, or a host name. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory.., or string with length: 1..253 This attribute
    // is mandatory..
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

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetSegmentPath() string {
    return "receiver" + "[address='" + fmt.Sprintf("%v", receiver.Address) + "']"
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = receiver.Address
    leafs["port"] = receiver.Port
    leafs["protocol"] = receiver.Protocol
    return leafs
}

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetBundleName() string { return "ietf" }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetYangName() string { return "receiver" }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Subscriptions_Subscription_Receivers_Receiver) GetParentYangName() string { return "receivers" }

