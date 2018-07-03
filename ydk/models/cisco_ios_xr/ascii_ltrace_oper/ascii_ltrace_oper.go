// This module contains a collection of YANG definitions
// for Cisco IOS-XR ascii-ltrace package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ltrace: ASCII ltrace data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ascii_ltrace_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ascii_ltrace_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ascii-ltrace-oper ltrace}", reflect.TypeOf(Ltrace{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ascii-ltrace-oper:ltrace", reflect.TypeOf(Ltrace{}))
}

// Ltrace
// ASCII ltrace data
type Ltrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // feature.
    Features Ltrace_Features
}

func (ltrace *Ltrace) GetEntityData() *types.CommonEntityData {
    ltrace.EntityData.YFilter = ltrace.YFilter
    ltrace.EntityData.YangName = "ltrace"
    ltrace.EntityData.BundleName = "cisco_ios_xr"
    ltrace.EntityData.ParentYangName = "Cisco-IOS-XR-ascii-ltrace-oper"
    ltrace.EntityData.SegmentPath = "Cisco-IOS-XR-ascii-ltrace-oper:ltrace"
    ltrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ltrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ltrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ltrace.EntityData.Children = types.NewOrderedMap()
    ltrace.EntityData.Children.Append("features", types.YChild{"Features", &ltrace.Features})
    ltrace.EntityData.Leafs = types.NewOrderedMap()

    ltrace.EntityData.YListKeys = []string {}

    return &(ltrace.EntityData)
}

// Ltrace_Features
// feature
type Ltrace_Features struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // feature. The type is slice of Ltrace_Features_Feature.
    Feature []*Ltrace_Features_Feature
}

func (features *Ltrace_Features) GetEntityData() *types.CommonEntityData {
    features.EntityData.YFilter = features.YFilter
    features.EntityData.YangName = "features"
    features.EntityData.BundleName = "cisco_ios_xr"
    features.EntityData.ParentYangName = "ltrace"
    features.EntityData.SegmentPath = "features"
    features.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    features.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    features.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    features.EntityData.Children = types.NewOrderedMap()
    features.EntityData.Children.Append("feature", types.YChild{"Feature", nil})
    for i := range features.Feature {
        features.EntityData.Children.Append(types.GetSegmentPath(features.Feature[i]), types.YChild{"Feature", features.Feature[i]})
    }
    features.EntityData.Leafs = types.NewOrderedMap()

    features.EntityData.YListKeys = []string {}

    return &(features.EntityData)
}

// Ltrace_Features_Feature
// feature
type Ltrace_Features_Feature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // feature name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    FeatureName interface{}

    // trace buffer name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    TraceBuf interface{}

    // trace.
    Traces Ltrace_Features_Feature_Traces
}

func (feature *Ltrace_Features_Feature) GetEntityData() *types.CommonEntityData {
    feature.EntityData.YFilter = feature.YFilter
    feature.EntityData.YangName = "feature"
    feature.EntityData.BundleName = "cisco_ios_xr"
    feature.EntityData.ParentYangName = "features"
    feature.EntityData.SegmentPath = "feature"
    feature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    feature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    feature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    feature.EntityData.Children = types.NewOrderedMap()
    feature.EntityData.Children.Append("traces", types.YChild{"Traces", &feature.Traces})
    feature.EntityData.Leafs = types.NewOrderedMap()
    feature.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", feature.FeatureName})
    feature.EntityData.Leafs.Append("trace-buf", types.YLeaf{"TraceBuf", feature.TraceBuf})

    feature.EntityData.YListKeys = []string {}

    return &(feature.EntityData)
}

// Ltrace_Features_Feature_Traces
// trace
type Ltrace_Features_Feature_Traces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // trace. The type is slice of Ltrace_Features_Feature_Traces_Trace.
    Trace []*Ltrace_Features_Feature_Traces_Trace
}

func (traces *Ltrace_Features_Feature_Traces) GetEntityData() *types.CommonEntityData {
    traces.EntityData.YFilter = traces.YFilter
    traces.EntityData.YangName = "traces"
    traces.EntityData.BundleName = "cisco_ios_xr"
    traces.EntityData.ParentYangName = "feature"
    traces.EntityData.SegmentPath = "traces"
    traces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traces.EntityData.Children = types.NewOrderedMap()
    traces.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range traces.Trace {
        traces.EntityData.Children.Append(types.GetSegmentPath(traces.Trace[i]), types.YChild{"Trace", traces.Trace[i]})
    }
    traces.EntityData.Leafs = types.NewOrderedMap()

    traces.EntityData.YListKeys = []string {}

    return &(traces.EntityData)
}

// Ltrace_Features_Feature_Traces_Trace
// trace
type Ltrace_Features_Feature_Traces_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ltrace ID of ltrace. The type is interface{} with
    // range: 0..4294967295.
    LtraceId interface{}

    // timestamp. The type is string.
    Timestamp interface{}

    // a single line of a trace point. The type is string.
    Line interface{}
}

func (trace *Ltrace_Features_Feature_Traces_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "traces"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.LtraceId, "ltrace-id")
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("ltrace-id", types.YLeaf{"LtraceId", trace.LtraceId})
    trace.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", trace.Timestamp})
    trace.EntityData.Leafs.Append("line", types.YLeaf{"Line", trace.Line})

    trace.EntityData.YListKeys = []string {"LtraceId"}

    return &(trace.EntityData)
}

