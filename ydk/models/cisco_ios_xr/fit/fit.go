// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package fit

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fit"))
    ydk.RegisterEntity("{http://cisco.com/panini/calvados/fit set}", reflect.TypeOf(Set{}))
    ydk.RegisterEntity("fit:set", reflect.TypeOf(Set{}))
}

// Set
type Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Set_Asic.
    Asic []*Set_Asic
}

func (set *Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "fit"
    set.EntityData.SegmentPath = "fit:set"
    set.EntityData.AbsolutePath = set.EntityData.SegmentPath
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("asic", types.YChild{"Asic", nil})
    for i := range set.Asic {
        set.EntityData.Children.Append(types.GetSegmentPath(set.Asic[i]), types.YChild{"Asic", set.Asic[i]})
    }
    set.EntityData.Leafs = types.NewOrderedMap()

    set.EntityData.YListKeys = []string {}

    return &(set.EntityData)
}

// Set_Asic
type Set_Asic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    AsicName interface{}

    // The type is slice of Set_Asic_Instance.
    Instance []*Set_Asic_Instance
}

func (asic *Set_Asic) GetEntityData() *types.CommonEntityData {
    asic.EntityData.YFilter = asic.YFilter
    asic.EntityData.YangName = "asic"
    asic.EntityData.BundleName = "cisco_ios_xr"
    asic.EntityData.ParentYangName = "set"
    asic.EntityData.SegmentPath = "asic" + types.AddKeyToken(asic.AsicName, "asic-name")
    asic.EntityData.AbsolutePath = "fit:set/" + asic.EntityData.SegmentPath
    asic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asic.EntityData.Children = types.NewOrderedMap()
    asic.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range asic.Instance {
        asic.EntityData.Children.Append(types.GetSegmentPath(asic.Instance[i]), types.YChild{"Instance", asic.Instance[i]})
    }
    asic.EntityData.Leafs = types.NewOrderedMap()
    asic.EntityData.Leafs.Append("asic-name", types.YLeaf{"AsicName", asic.AsicName})

    asic.EntityData.YListKeys = []string {"AsicName"}

    return &(asic.EntityData)
}

// Set_Asic_Instance
type Set_Asic_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    InstanceIds interface{}

    
    FaultInjection Set_Asic_Instance_FaultInjection
}

func (instance *Set_Asic_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "asic"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceIds, "instance-ids")
    instance.EntityData.AbsolutePath = "fit:set/asic/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("fault-injection", types.YChild{"FaultInjection", &instance.FaultInjection})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-ids", types.YLeaf{"InstanceIds", instance.InstanceIds})

    instance.EntityData.YListKeys = []string {"InstanceIds"}

    return &(instance.EntityData)
}

// Set_Asic_Instance_FaultInjection
type Set_Asic_Instance_FaultInjection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Set_Asic_Instance_FaultInjection_Module.
    Module []*Set_Asic_Instance_FaultInjection_Module
}

func (faultInjection *Set_Asic_Instance_FaultInjection) GetEntityData() *types.CommonEntityData {
    faultInjection.EntityData.YFilter = faultInjection.YFilter
    faultInjection.EntityData.YangName = "fault-injection"
    faultInjection.EntityData.BundleName = "cisco_ios_xr"
    faultInjection.EntityData.ParentYangName = "instance"
    faultInjection.EntityData.SegmentPath = "fault-injection"
    faultInjection.EntityData.AbsolutePath = "fit:set/asic/instance/" + faultInjection.EntityData.SegmentPath
    faultInjection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultInjection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultInjection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultInjection.EntityData.Children = types.NewOrderedMap()
    faultInjection.EntityData.Children.Append("module", types.YChild{"Module", nil})
    for i := range faultInjection.Module {
        faultInjection.EntityData.Children.Append(types.GetSegmentPath(faultInjection.Module[i]), types.YChild{"Module", faultInjection.Module[i]})
    }
    faultInjection.EntityData.Leafs = types.NewOrderedMap()

    faultInjection.EntityData.YListKeys = []string {}

    return &(faultInjection.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module
type Set_Asic_Instance_FaultInjection_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    ModuleName interface{}

    
    FaultType Set_Asic_Instance_FaultInjection_Module_FaultType
}

func (module *Set_Asic_Instance_FaultInjection_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "fault-injection"
    module.EntityData.SegmentPath = "module" + types.AddKeyToken(module.ModuleName, "module-name")
    module.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/" + module.EntityData.SegmentPath
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = types.NewOrderedMap()
    module.EntityData.Children.Append("fault-type", types.YChild{"FaultType", &module.FaultType})
    module.EntityData.Leafs = types.NewOrderedMap()
    module.EntityData.Leafs.Append("module-name", types.YLeaf{"ModuleName", module.ModuleName})

    module.EntityData.YListKeys = []string {"ModuleName"}

    return &(module.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType
type Set_Asic_Instance_FaultInjection_Module_FaultType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ecc Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc

    
    Parity Set_Asic_Instance_FaultInjection_Module_FaultType_Parity

    
    Other Set_Asic_Instance_FaultInjection_Module_FaultType_Other
}

func (faultType *Set_Asic_Instance_FaultInjection_Module_FaultType) GetEntityData() *types.CommonEntityData {
    faultType.EntityData.YFilter = faultType.YFilter
    faultType.EntityData.YangName = "fault-type"
    faultType.EntityData.BundleName = "cisco_ios_xr"
    faultType.EntityData.ParentYangName = "module"
    faultType.EntityData.SegmentPath = "fault-type"
    faultType.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/" + faultType.EntityData.SegmentPath
    faultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType.EntityData.Children = types.NewOrderedMap()
    faultType.EntityData.Children.Append("ecc", types.YChild{"Ecc", &faultType.Ecc})
    faultType.EntityData.Children.Append("parity", types.YChild{"Parity", &faultType.Parity})
    faultType.EntityData.Children.Append("other", types.YChild{"Other", &faultType.Other})
    faultType.EntityData.Leafs = types.NewOrderedMap()

    faultType.EntityData.YListKeys = []string {}

    return &(faultType.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst.
    BlockNameLst []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst
}

func (ecc *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc) GetEntityData() *types.CommonEntityData {
    ecc.EntityData.YFilter = ecc.YFilter
    ecc.EntityData.YangName = "ecc"
    ecc.EntityData.BundleName = "cisco_ios_xr"
    ecc.EntityData.ParentYangName = "fault-type"
    ecc.EntityData.SegmentPath = "ecc"
    ecc.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/" + ecc.EntityData.SegmentPath
    ecc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ecc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ecc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ecc.EntityData.Children = types.NewOrderedMap()
    ecc.EntityData.Children.Append("all", types.YChild{"All", &ecc.All})
    ecc.EntityData.Children.Append("block-name-lst", types.YChild{"BlockNameLst", nil})
    for i := range ecc.BlockNameLst {
        ecc.EntityData.Children.Append(types.GetSegmentPath(ecc.BlockNameLst[i]), types.YChild{"BlockNameLst", ecc.BlockNameLst[i]})
    }
    ecc.EntityData.Leafs = types.NewOrderedMap()

    ecc.EntityData.YListKeys = []string {}

    return &(ecc.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold.
    Threshold []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "ecc"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("threshold", types.YChild{"Threshold", nil})
    for i := range all.Threshold {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Threshold[i]), types.YChild{"Threshold", all.Threshold[i]})
    }
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + types.AddKeyToken(threshold.NumSeconds, "num-seconds")
    threshold.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/all/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range threshold.Location {
        threshold.EntityData.Children.Append(types.GetSegmentPath(threshold.Location[i]), types.YChild{"Location", threshold.Location[i]})
    }
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", threshold.NumSeconds})

    threshold.EntityData.YListKeys = []string {"NumSeconds"}

    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/all/threshold/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    BlockName interface{}

    
    One Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One

    
    Continuous Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous

    
    Stop Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop
}

func (blockNameLst *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst) GetEntityData() *types.CommonEntityData {
    blockNameLst.EntityData.YFilter = blockNameLst.YFilter
    blockNameLst.EntityData.YangName = "block-name-lst"
    blockNameLst.EntityData.BundleName = "cisco_ios_xr"
    blockNameLst.EntityData.ParentYangName = "ecc"
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + types.AddKeyToken(blockNameLst.BlockName, "block-name")
    blockNameLst.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/" + blockNameLst.EntityData.SegmentPath
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = types.NewOrderedMap()
    blockNameLst.EntityData.Children.Append("one", types.YChild{"One", &blockNameLst.One})
    blockNameLst.EntityData.Children.Append("continuous", types.YChild{"Continuous", &blockNameLst.Continuous})
    blockNameLst.EntityData.Children.Append("stop", types.YChild{"Stop", &blockNameLst.Stop})
    blockNameLst.EntityData.Leafs = types.NewOrderedMap()
    blockNameLst.EntityData.Leafs.Append("block-name", types.YLeaf{"BlockName", blockNameLst.BlockName})

    blockNameLst.EntityData.YListKeys = []string {"BlockName"}

    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/" + one.EntityData.SegmentPath
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = types.NewOrderedMap()
    one.EntityData.Children.Append("rate", types.YChild{"Rate", &one.Rate})
    one.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range one.Location {
        one.EntityData.Children.Append(types.GetSegmentPath(one.Location[i]), types.YChild{"Location", one.Location[i]})
    }
    one.EntityData.Leafs = types.NewOrderedMap()

    one.EntityData.YListKeys = []string {}

    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/one/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/one/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/one/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/one/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/one/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/" + continuous.EntityData.SegmentPath
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = types.NewOrderedMap()
    continuous.EntityData.Children.Append("rate", types.YChild{"Rate", &continuous.Rate})
    continuous.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range continuous.Location {
        continuous.EntityData.Children.Append(types.GetSegmentPath(continuous.Location[i]), types.YChild{"Location", continuous.Location[i]})
    }
    continuous.EntityData.Leafs = types.NewOrderedMap()

    continuous.EntityData.YListKeys = []string {}

    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/continuous/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/continuous/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/continuous/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/continuous/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/continuous/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range stop.Location {
        stop.EntityData.Children.Append(types.GetSegmentPath(stop.Location[i]), types.YChild{"Location", stop.Location[i]})
    }
    stop.EntityData.Leafs = types.NewOrderedMap()

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/ecc/block-name-lst/stop/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst.
    BlockNameLst []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst
}

func (parity *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "fault-type"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/" + parity.EntityData.SegmentPath
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = types.NewOrderedMap()
    parity.EntityData.Children.Append("all", types.YChild{"All", &parity.All})
    parity.EntityData.Children.Append("block-name-lst", types.YChild{"BlockNameLst", nil})
    for i := range parity.BlockNameLst {
        parity.EntityData.Children.Append(types.GetSegmentPath(parity.BlockNameLst[i]), types.YChild{"BlockNameLst", parity.BlockNameLst[i]})
    }
    parity.EntityData.Leafs = types.NewOrderedMap()

    parity.EntityData.YListKeys = []string {}

    return &(parity.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold.
    Threshold []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "parity"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("threshold", types.YChild{"Threshold", nil})
    for i := range all.Threshold {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Threshold[i]), types.YChild{"Threshold", all.Threshold[i]})
    }
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + types.AddKeyToken(threshold.NumSeconds, "num-seconds")
    threshold.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/all/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range threshold.Location {
        threshold.EntityData.Children.Append(types.GetSegmentPath(threshold.Location[i]), types.YChild{"Location", threshold.Location[i]})
    }
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", threshold.NumSeconds})

    threshold.EntityData.YListKeys = []string {"NumSeconds"}

    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/all/threshold/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    BlockName interface{}

    
    One Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One

    
    Continuous Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous

    
    Stop Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop
}

func (blockNameLst *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst) GetEntityData() *types.CommonEntityData {
    blockNameLst.EntityData.YFilter = blockNameLst.YFilter
    blockNameLst.EntityData.YangName = "block-name-lst"
    blockNameLst.EntityData.BundleName = "cisco_ios_xr"
    blockNameLst.EntityData.ParentYangName = "parity"
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + types.AddKeyToken(blockNameLst.BlockName, "block-name")
    blockNameLst.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/" + blockNameLst.EntityData.SegmentPath
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = types.NewOrderedMap()
    blockNameLst.EntityData.Children.Append("one", types.YChild{"One", &blockNameLst.One})
    blockNameLst.EntityData.Children.Append("continuous", types.YChild{"Continuous", &blockNameLst.Continuous})
    blockNameLst.EntityData.Children.Append("stop", types.YChild{"Stop", &blockNameLst.Stop})
    blockNameLst.EntityData.Leafs = types.NewOrderedMap()
    blockNameLst.EntityData.Leafs.Append("block-name", types.YLeaf{"BlockName", blockNameLst.BlockName})

    blockNameLst.EntityData.YListKeys = []string {"BlockName"}

    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/" + one.EntityData.SegmentPath
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = types.NewOrderedMap()
    one.EntityData.Children.Append("rate", types.YChild{"Rate", &one.Rate})
    one.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range one.Location {
        one.EntityData.Children.Append(types.GetSegmentPath(one.Location[i]), types.YChild{"Location", one.Location[i]})
    }
    one.EntityData.Leafs = types.NewOrderedMap()

    one.EntityData.YListKeys = []string {}

    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/one/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/one/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/one/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/one/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/one/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/" + continuous.EntityData.SegmentPath
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = types.NewOrderedMap()
    continuous.EntityData.Children.Append("rate", types.YChild{"Rate", &continuous.Rate})
    continuous.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range continuous.Location {
        continuous.EntityData.Children.Append(types.GetSegmentPath(continuous.Location[i]), types.YChild{"Location", continuous.Location[i]})
    }
    continuous.EntityData.Leafs = types.NewOrderedMap()

    continuous.EntityData.YListKeys = []string {}

    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/continuous/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/continuous/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/continuous/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/continuous/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/continuous/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range stop.Location {
        stop.EntityData.Children.Append(types.GetSegmentPath(stop.Location[i]), types.YChild{"Location", stop.Location[i]})
    }
    stop.EntityData.Leafs = types.NewOrderedMap()

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/parity/block-name-lst/stop/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst.
    BlockNameLst []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst
}

func (other *Set_Asic_Instance_FaultInjection_Module_FaultType_Other) GetEntityData() *types.CommonEntityData {
    other.EntityData.YFilter = other.YFilter
    other.EntityData.YangName = "other"
    other.EntityData.BundleName = "cisco_ios_xr"
    other.EntityData.ParentYangName = "fault-type"
    other.EntityData.SegmentPath = "other"
    other.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/" + other.EntityData.SegmentPath
    other.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    other.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    other.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    other.EntityData.Children = types.NewOrderedMap()
    other.EntityData.Children.Append("all", types.YChild{"All", &other.All})
    other.EntityData.Children.Append("block-name-lst", types.YChild{"BlockNameLst", nil})
    for i := range other.BlockNameLst {
        other.EntityData.Children.Append(types.GetSegmentPath(other.BlockNameLst[i]), types.YChild{"BlockNameLst", other.BlockNameLst[i]})
    }
    other.EntityData.Leafs = types.NewOrderedMap()

    other.EntityData.YListKeys = []string {}

    return &(other.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold.
    Threshold []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "other"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("threshold", types.YChild{"Threshold", nil})
    for i := range all.Threshold {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Threshold[i]), types.YChild{"Threshold", all.Threshold[i]})
    }
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + types.AddKeyToken(threshold.NumSeconds, "num-seconds")
    threshold.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/all/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range threshold.Location {
        threshold.EntityData.Children.Append(types.GetSegmentPath(threshold.Location[i]), types.YChild{"Location", threshold.Location[i]})
    }
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", threshold.NumSeconds})

    threshold.EntityData.YListKeys = []string {"NumSeconds"}

    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/all/threshold/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    BlockName interface{}

    
    One Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One

    
    Continuous Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous

    
    Stop Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop
}

func (blockNameLst *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst) GetEntityData() *types.CommonEntityData {
    blockNameLst.EntityData.YFilter = blockNameLst.YFilter
    blockNameLst.EntityData.YangName = "block-name-lst"
    blockNameLst.EntityData.BundleName = "cisco_ios_xr"
    blockNameLst.EntityData.ParentYangName = "other"
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + types.AddKeyToken(blockNameLst.BlockName, "block-name")
    blockNameLst.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/" + blockNameLst.EntityData.SegmentPath
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = types.NewOrderedMap()
    blockNameLst.EntityData.Children.Append("one", types.YChild{"One", &blockNameLst.One})
    blockNameLst.EntityData.Children.Append("continuous", types.YChild{"Continuous", &blockNameLst.Continuous})
    blockNameLst.EntityData.Children.Append("stop", types.YChild{"Stop", &blockNameLst.Stop})
    blockNameLst.EntityData.Leafs = types.NewOrderedMap()
    blockNameLst.EntityData.Leafs.Append("block-name", types.YLeaf{"BlockName", blockNameLst.BlockName})

    blockNameLst.EntityData.YListKeys = []string {"BlockName"}

    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/" + one.EntityData.SegmentPath
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = types.NewOrderedMap()
    one.EntityData.Children.Append("rate", types.YChild{"Rate", &one.Rate})
    one.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range one.Location {
        one.EntityData.Children.Append(types.GetSegmentPath(one.Location[i]), types.YChild{"Location", one.Location[i]})
    }
    one.EntityData.Leafs = types.NewOrderedMap()

    one.EntityData.YListKeys = []string {}

    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/one/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/one/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/one/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/one/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/one/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/" + continuous.EntityData.SegmentPath
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = types.NewOrderedMap()
    continuous.EntityData.Children.Append("rate", types.YChild{"Rate", &continuous.Rate})
    continuous.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range continuous.Location {
        continuous.EntityData.Children.Append(types.GetSegmentPath(continuous.Location[i]), types.YChild{"Location", continuous.Location[i]})
    }
    continuous.EntityData.Leafs = types.NewOrderedMap()

    continuous.EntityData.YListKeys = []string {}

    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/continuous/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Children.Append("error-number", types.YChild{"ErrorNumber", nil})
    for i := range rate.ErrorNumber {
        rate.EntityData.Children.Append(types.GetSegmentPath(rate.ErrorNumber[i]), types.YChild{"ErrorNumber", rate.ErrorNumber[i]})
    }
    rate.EntityData.Leafs = types.NewOrderedMap()

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + types.AddKeyToken(errorNumber.NumErrs, "num-errs")
    errorNumber.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/continuous/rate/" + errorNumber.EntityData.SegmentPath
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = types.NewOrderedMap()
    errorNumber.EntityData.Children.Append("duration", types.YChild{"Duration", nil})
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children.Append(types.GetSegmentPath(errorNumber.Duration[i]), types.YChild{"Duration", errorNumber.Duration[i]})
    }
    errorNumber.EntityData.Leafs = types.NewOrderedMap()
    errorNumber.EntityData.Leafs.Append("num-errs", types.YLeaf{"NumErrs", errorNumber.NumErrs})

    errorNumber.EntityData.YListKeys = []string {"NumErrs"}

    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + types.AddKeyToken(duration.NumSeconds, "num-seconds")
    duration.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/continuous/rate/error-number/" + duration.EntityData.SegmentPath
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = types.NewOrderedMap()
    duration.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range duration.Location {
        duration.EntityData.Children.Append(types.GetSegmentPath(duration.Location[i]), types.YChild{"Location", duration.Location[i]})
    }
    duration.EntityData.Leafs = types.NewOrderedMap()
    duration.EntityData.Leafs.Append("num-seconds", types.YLeaf{"NumSeconds", duration.NumSeconds})

    duration.EntityData.YListKeys = []string {"NumSeconds"}

    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/continuous/rate/error-number/duration/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/continuous/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location.
    Location []*Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range stop.Location {
        stop.EntityData.Children.Append(types.GetSegmentPath(stop.Location[i]), types.YChild{"Location", stop.Location[i]})
    }
    stop.EntityData.Leafs = types.NewOrderedMap()

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.FitLocationName, "fit-location-name")
    location.EntityData.AbsolutePath = "fit:set/asic/instance/fault-injection/module/fault-type/other/block-name-lst/stop/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("fit-location-name", types.YLeaf{"FitLocationName", location.FitLocationName})

    location.EntityData.YListKeys = []string {"FitLocationName"}

    return &(location.EntityData)
}

