// This module provides CLI for Status, ATTN, ALARM LED's.
package sysadmin_led_mgr_ui

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_led_mgr_ui"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-led-mgr-ui led}", reflect.TypeOf(Led{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-led-mgr-ui:led", reflect.TypeOf(Led{}))
}

// Led
// Calvados Led Manager Status, Attn, Alarm inventory
type Led struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Led_Location.
    Location []*Led_Location

    // show traceable processes. The type is slice of Led_Trace.
    Trace []*Led_Trace
}

func (led *Led) GetEntityData() *types.CommonEntityData {
    led.EntityData.YFilter = led.YFilter
    led.EntityData.YangName = "led"
    led.EntityData.BundleName = "cisco_ios_xr"
    led.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-led-mgr-ui"
    led.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-led-mgr-ui:led"
    led.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    led.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    led.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    led.EntityData.Children = types.NewOrderedMap()
    led.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range led.Location {
        led.EntityData.Children.Append(types.GetSegmentPath(led.Location[i]), types.YChild{"Location", led.Location[i]})
    }
    led.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range led.Trace {
        led.EntityData.Children.Append(types.GetSegmentPath(led.Trace[i]), types.YChild{"Trace", led.Trace[i]})
    }
    led.EntityData.Leafs = types.NewOrderedMap()

    led.EntityData.YListKeys = []string {}

    return &(led.EntityData)
}

// Led_Location
type Led_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Led_Location_LedAttributes.
    LedAttributes []*Led_Location_LedAttributes
}

func (location *Led_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "led"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("led_attributes", types.YChild{"LedAttributes", nil})
    for i := range location.LedAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LedAttributes[i]), types.YChild{"LedAttributes", location.LedAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Led_Location_LedAttributes
type Led_Location_LedAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LedName interface{}

    // The type is string.
    LedMode interface{}

    // The type is string.
    LedColor interface{}
}

func (ledAttributes *Led_Location_LedAttributes) GetEntityData() *types.CommonEntityData {
    ledAttributes.EntityData.YFilter = ledAttributes.YFilter
    ledAttributes.EntityData.YangName = "led_attributes"
    ledAttributes.EntityData.BundleName = "cisco_ios_xr"
    ledAttributes.EntityData.ParentYangName = "location"
    ledAttributes.EntityData.SegmentPath = "led_attributes" + types.AddKeyToken(ledAttributes.LedName, "led_name")
    ledAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ledAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ledAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ledAttributes.EntityData.Children = types.NewOrderedMap()
    ledAttributes.EntityData.Leafs = types.NewOrderedMap()
    ledAttributes.EntityData.Leafs.Append("led_name", types.YLeaf{"LedName", ledAttributes.LedName})
    ledAttributes.EntityData.Leafs.Append("led_mode", types.YLeaf{"LedMode", ledAttributes.LedMode})
    ledAttributes.EntityData.Leafs.Append("led_color", types.YLeaf{"LedColor", ledAttributes.LedColor})

    ledAttributes.EntityData.YListKeys = []string {"LedName"}

    return &(ledAttributes.EntityData)
}

// Led_Trace
// show traceable processes
type Led_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Led_Trace_Location.
    Location []*Led_Trace_Location
}

func (trace *Led_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "led"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
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

// Led_Trace_Location
type Led_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Led_Trace_Location_AllOptions.
    AllOptions []*Led_Trace_Location_AllOptions
}

func (location *Led_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
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

// Led_Trace_Location_AllOptions
type Led_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Led_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Led_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Led_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Led_Trace_Location_AllOptions_TraceBlocks
type Led_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Led_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

