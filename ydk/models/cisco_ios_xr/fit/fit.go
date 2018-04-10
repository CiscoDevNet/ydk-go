// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
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
    Asic []Set_Asic
}

func (set *Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "fit"
    set.EntityData.SegmentPath = "fit:set"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["asic"] = types.YChild{"Asic", nil}
    for i := range set.Asic {
        set.EntityData.Children[types.GetSegmentPath(&set.Asic[i])] = types.YChild{"Asic", &set.Asic[i]}
    }
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(set.EntityData)
}

// Set_Asic
type Set_Asic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    AsicName interface{}

    // The type is slice of Set_Asic_Instance.
    Instance []Set_Asic_Instance
}

func (asic *Set_Asic) GetEntityData() *types.CommonEntityData {
    asic.EntityData.YFilter = asic.YFilter
    asic.EntityData.YangName = "asic"
    asic.EntityData.BundleName = "cisco_ios_xr"
    asic.EntityData.ParentYangName = "set"
    asic.EntityData.SegmentPath = "asic" + "[asic-name='" + fmt.Sprintf("%v", asic.AsicName) + "']"
    asic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asic.EntityData.Children = make(map[string]types.YChild)
    asic.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range asic.Instance {
        asic.EntityData.Children[types.GetSegmentPath(&asic.Instance[i])] = types.YChild{"Instance", &asic.Instance[i]}
    }
    asic.EntityData.Leafs = make(map[string]types.YLeaf)
    asic.EntityData.Leafs["asic-name"] = types.YLeaf{"AsicName", asic.AsicName}
    return &(asic.EntityData)
}

// Set_Asic_Instance
type Set_Asic_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    InstanceIds interface{}

    
    FaultInjection Set_Asic_Instance_FaultInjection
}

func (instance *Set_Asic_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "asic"
    instance.EntityData.SegmentPath = "instance" + "[instance-ids='" + fmt.Sprintf("%v", instance.InstanceIds) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["fault-injection"] = types.YChild{"FaultInjection", &instance.FaultInjection}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-ids"] = types.YLeaf{"InstanceIds", instance.InstanceIds}
    return &(instance.EntityData)
}

// Set_Asic_Instance_FaultInjection
type Set_Asic_Instance_FaultInjection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Set_Asic_Instance_FaultInjection_Module.
    Module []Set_Asic_Instance_FaultInjection_Module
}

func (faultInjection *Set_Asic_Instance_FaultInjection) GetEntityData() *types.CommonEntityData {
    faultInjection.EntityData.YFilter = faultInjection.YFilter
    faultInjection.EntityData.YangName = "fault-injection"
    faultInjection.EntityData.BundleName = "cisco_ios_xr"
    faultInjection.EntityData.ParentYangName = "instance"
    faultInjection.EntityData.SegmentPath = "fault-injection"
    faultInjection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultInjection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultInjection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultInjection.EntityData.Children = make(map[string]types.YChild)
    faultInjection.EntityData.Children["module"] = types.YChild{"Module", nil}
    for i := range faultInjection.Module {
        faultInjection.EntityData.Children[types.GetSegmentPath(&faultInjection.Module[i])] = types.YChild{"Module", &faultInjection.Module[i]}
    }
    faultInjection.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faultInjection.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module
type Set_Asic_Instance_FaultInjection_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    ModuleName interface{}

    
    FaultType Set_Asic_Instance_FaultInjection_Module_FaultType
}

func (module *Set_Asic_Instance_FaultInjection_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "fault-injection"
    module.EntityData.SegmentPath = "module" + "[module-name='" + fmt.Sprintf("%v", module.ModuleName) + "']"
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = make(map[string]types.YChild)
    module.EntityData.Children["fault-type"] = types.YChild{"FaultType", &module.FaultType}
    module.EntityData.Leafs = make(map[string]types.YLeaf)
    module.EntityData.Leafs["module-name"] = types.YLeaf{"ModuleName", module.ModuleName}
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
    faultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType.EntityData.Children = make(map[string]types.YChild)
    faultType.EntityData.Children["ecc"] = types.YChild{"Ecc", &faultType.Ecc}
    faultType.EntityData.Children["parity"] = types.YChild{"Parity", &faultType.Parity}
    faultType.EntityData.Children["other"] = types.YChild{"Other", &faultType.Other}
    faultType.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faultType.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst.
    BlockNameLst []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst
}

func (ecc *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc) GetEntityData() *types.CommonEntityData {
    ecc.EntityData.YFilter = ecc.YFilter
    ecc.EntityData.YangName = "ecc"
    ecc.EntityData.BundleName = "cisco_ios_xr"
    ecc.EntityData.ParentYangName = "fault-type"
    ecc.EntityData.SegmentPath = "ecc"
    ecc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ecc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ecc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ecc.EntityData.Children = make(map[string]types.YChild)
    ecc.EntityData.Children["all"] = types.YChild{"All", &ecc.All}
    ecc.EntityData.Children["block-name-lst"] = types.YChild{"BlockNameLst", nil}
    for i := range ecc.BlockNameLst {
        ecc.EntityData.Children[types.GetSegmentPath(&ecc.BlockNameLst[i])] = types.YChild{"BlockNameLst", &ecc.BlockNameLst[i]}
    }
    ecc.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ecc.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold.
    Threshold []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "ecc"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["threshold"] = types.YChild{"Threshold", nil}
    for i := range all.Threshold {
        all.EntityData.Children[types.GetSegmentPath(&all.Threshold[i])] = types.YChild{"Threshold", &all.Threshold[i]}
    }
    all.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range all.Location {
        all.EntityData.Children[types.GetSegmentPath(&all.Location[i])] = types.YChild{"Location", &all.Location[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + "[num-seconds='" + fmt.Sprintf("%v", threshold.NumSeconds) + "']"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range threshold.Location {
        threshold.EntityData.Children[types.GetSegmentPath(&threshold.Location[i])] = types.YChild{"Location", &threshold.Location[i]}
    }
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", threshold.NumSeconds}
    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + "[block-name='" + fmt.Sprintf("%v", blockNameLst.BlockName) + "']"
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = make(map[string]types.YChild)
    blockNameLst.EntityData.Children["one"] = types.YChild{"One", &blockNameLst.One}
    blockNameLst.EntityData.Children["continuous"] = types.YChild{"Continuous", &blockNameLst.Continuous}
    blockNameLst.EntityData.Children["stop"] = types.YChild{"Stop", &blockNameLst.Stop}
    blockNameLst.EntityData.Leafs = make(map[string]types.YLeaf)
    blockNameLst.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockNameLst.BlockName}
    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = make(map[string]types.YChild)
    one.EntityData.Children["rate"] = types.YChild{"Rate", &one.Rate}
    one.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range one.Location {
        one.EntityData.Children[types.GetSegmentPath(&one.Location[i])] = types.YChild{"Location", &one.Location[i]}
    }
    one.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = make(map[string]types.YChild)
    continuous.EntityData.Children["rate"] = types.YChild{"Rate", &continuous.Rate}
    continuous.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range continuous.Location {
        continuous.EntityData.Children[types.GetSegmentPath(&continuous.Location[i])] = types.YChild{"Location", &continuous.Location[i]}
    }
    continuous.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range stop.Location {
        stop.EntityData.Children[types.GetSegmentPath(&stop.Location[i])] = types.YChild{"Location", &stop.Location[i]}
    }
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Ecc_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst.
    BlockNameLst []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst
}

func (parity *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "fault-type"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = make(map[string]types.YChild)
    parity.EntityData.Children["all"] = types.YChild{"All", &parity.All}
    parity.EntityData.Children["block-name-lst"] = types.YChild{"BlockNameLst", nil}
    for i := range parity.BlockNameLst {
        parity.EntityData.Children[types.GetSegmentPath(&parity.BlockNameLst[i])] = types.YChild{"BlockNameLst", &parity.BlockNameLst[i]}
    }
    parity.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(parity.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold.
    Threshold []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "parity"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["threshold"] = types.YChild{"Threshold", nil}
    for i := range all.Threshold {
        all.EntityData.Children[types.GetSegmentPath(&all.Threshold[i])] = types.YChild{"Threshold", &all.Threshold[i]}
    }
    all.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range all.Location {
        all.EntityData.Children[types.GetSegmentPath(&all.Location[i])] = types.YChild{"Location", &all.Location[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + "[num-seconds='" + fmt.Sprintf("%v", threshold.NumSeconds) + "']"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range threshold.Location {
        threshold.EntityData.Children[types.GetSegmentPath(&threshold.Location[i])] = types.YChild{"Location", &threshold.Location[i]}
    }
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", threshold.NumSeconds}
    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + "[block-name='" + fmt.Sprintf("%v", blockNameLst.BlockName) + "']"
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = make(map[string]types.YChild)
    blockNameLst.EntityData.Children["one"] = types.YChild{"One", &blockNameLst.One}
    blockNameLst.EntityData.Children["continuous"] = types.YChild{"Continuous", &blockNameLst.Continuous}
    blockNameLst.EntityData.Children["stop"] = types.YChild{"Stop", &blockNameLst.Stop}
    blockNameLst.EntityData.Leafs = make(map[string]types.YLeaf)
    blockNameLst.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockNameLst.BlockName}
    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = make(map[string]types.YChild)
    one.EntityData.Children["rate"] = types.YChild{"Rate", &one.Rate}
    one.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range one.Location {
        one.EntityData.Children[types.GetSegmentPath(&one.Location[i])] = types.YChild{"Location", &one.Location[i]}
    }
    one.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = make(map[string]types.YChild)
    continuous.EntityData.Children["rate"] = types.YChild{"Rate", &continuous.Rate}
    continuous.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range continuous.Location {
        continuous.EntityData.Children[types.GetSegmentPath(&continuous.Location[i])] = types.YChild{"Location", &continuous.Location[i]}
    }
    continuous.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range stop.Location {
        stop.EntityData.Children[types.GetSegmentPath(&stop.Location[i])] = types.YChild{"Location", &stop.Location[i]}
    }
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Parity_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst.
    BlockNameLst []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst
}

func (other *Set_Asic_Instance_FaultInjection_Module_FaultType_Other) GetEntityData() *types.CommonEntityData {
    other.EntityData.YFilter = other.YFilter
    other.EntityData.YangName = "other"
    other.EntityData.BundleName = "cisco_ios_xr"
    other.EntityData.ParentYangName = "fault-type"
    other.EntityData.SegmentPath = "other"
    other.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    other.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    other.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    other.EntityData.Children = make(map[string]types.YChild)
    other.EntityData.Children["all"] = types.YChild{"All", &other.All}
    other.EntityData.Children["block-name-lst"] = types.YChild{"BlockNameLst", nil}
    for i := range other.BlockNameLst {
        other.EntityData.Children[types.GetSegmentPath(&other.BlockNameLst[i])] = types.YChild{"BlockNameLst", &other.BlockNameLst[i]}
    }
    other.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(other.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold.
    Threshold []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location
}

func (all *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "other"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["threshold"] = types.YChild{"Threshold", nil}
    for i := range all.Threshold {
        all.EntityData.Children[types.GetSegmentPath(&all.Threshold[i])] = types.YChild{"Threshold", &all.Threshold[i]}
    }
    all.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range all.Location {
        all.EntityData.Children[types.GetSegmentPath(&all.Location[i])] = types.YChild{"Location", &all.Location[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location
}

func (threshold *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "all"
    threshold.EntityData.SegmentPath = "threshold" + "[num-seconds='" + fmt.Sprintf("%v", threshold.NumSeconds) + "']"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range threshold.Location {
        threshold.EntityData.Children[types.GetSegmentPath(&threshold.Location[i])] = types.YChild{"Location", &threshold.Location[i]}
    }
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", threshold.NumSeconds}
    return &(threshold.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    blockNameLst.EntityData.SegmentPath = "block-name-lst" + "[block-name='" + fmt.Sprintf("%v", blockNameLst.BlockName) + "']"
    blockNameLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockNameLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockNameLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockNameLst.EntityData.Children = make(map[string]types.YChild)
    blockNameLst.EntityData.Children["one"] = types.YChild{"One", &blockNameLst.One}
    blockNameLst.EntityData.Children["continuous"] = types.YChild{"Continuous", &blockNameLst.Continuous}
    blockNameLst.EntityData.Children["stop"] = types.YChild{"Stop", &blockNameLst.Stop}
    blockNameLst.EntityData.Leafs = make(map[string]types.YLeaf)
    blockNameLst.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockNameLst.BlockName}
    return &(blockNameLst.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location
}

func (one *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One) GetEntityData() *types.CommonEntityData {
    one.EntityData.YFilter = one.YFilter
    one.EntityData.YangName = "one"
    one.EntityData.BundleName = "cisco_ios_xr"
    one.EntityData.ParentYangName = "block-name-lst"
    one.EntityData.SegmentPath = "one"
    one.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    one.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    one.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    one.EntityData.Children = make(map[string]types.YChild)
    one.EntityData.Children["rate"] = types.YChild{"Rate", &one.Rate}
    one.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range one.Location {
        one.EntityData.Children[types.GetSegmentPath(&one.Location[i])] = types.YChild{"Location", &one.Location[i]}
    }
    one.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(one.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "one"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_One_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "one"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rate Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location
}

func (continuous *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous) GetEntityData() *types.CommonEntityData {
    continuous.EntityData.YFilter = continuous.YFilter
    continuous.EntityData.YangName = "continuous"
    continuous.EntityData.BundleName = "cisco_ios_xr"
    continuous.EntityData.ParentYangName = "block-name-lst"
    continuous.EntityData.SegmentPath = "continuous"
    continuous.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuous.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuous.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuous.EntityData.Children = make(map[string]types.YChild)
    continuous.EntityData.Children["rate"] = types.YChild{"Rate", &continuous.Rate}
    continuous.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range continuous.Location {
        continuous.EntityData.Children[types.GetSegmentPath(&continuous.Location[i])] = types.YChild{"Location", &continuous.Location[i]}
    }
    continuous.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(continuous.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber.
    ErrorNumber []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber
}

func (rate *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "continuous"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = make(map[string]types.YChild)
    rate.EntityData.Children["error-number"] = types.YChild{"ErrorNumber", nil}
    for i := range rate.ErrorNumber {
        rate.EntityData.Children[types.GetSegmentPath(&rate.ErrorNumber[i])] = types.YChild{"ErrorNumber", &rate.ErrorNumber[i]}
    }
    rate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rate.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumErrs interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration.
    Duration []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
}

func (errorNumber *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber) GetEntityData() *types.CommonEntityData {
    errorNumber.EntityData.YFilter = errorNumber.YFilter
    errorNumber.EntityData.YangName = "error-number"
    errorNumber.EntityData.BundleName = "cisco_ios_xr"
    errorNumber.EntityData.ParentYangName = "rate"
    errorNumber.EntityData.SegmentPath = "error-number" + "[num-errs='" + fmt.Sprintf("%v", errorNumber.NumErrs) + "']"
    errorNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorNumber.EntityData.Children = make(map[string]types.YChild)
    errorNumber.EntityData.Children["duration"] = types.YChild{"Duration", nil}
    for i := range errorNumber.Duration {
        errorNumber.EntityData.Children[types.GetSegmentPath(&errorNumber.Duration[i])] = types.YChild{"Duration", &errorNumber.Duration[i]}
    }
    errorNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    errorNumber.EntityData.Leafs["num-errs"] = types.YLeaf{"NumErrs", errorNumber.NumErrs}
    return &(errorNumber.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NumSeconds interface{}

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
}

func (duration *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration) GetEntityData() *types.CommonEntityData {
    duration.EntityData.YFilter = duration.YFilter
    duration.EntityData.YangName = "duration"
    duration.EntityData.BundleName = "cisco_ios_xr"
    duration.EntityData.ParentYangName = "error-number"
    duration.EntityData.SegmentPath = "duration" + "[num-seconds='" + fmt.Sprintf("%v", duration.NumSeconds) + "']"
    duration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duration.EntityData.Children = make(map[string]types.YChild)
    duration.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range duration.Location {
        duration.EntityData.Children[types.GetSegmentPath(&duration.Location[i])] = types.YChild{"Location", &duration.Location[i]}
    }
    duration.EntityData.Leafs = make(map[string]types.YLeaf)
    duration.EntityData.Leafs["num-seconds"] = types.YLeaf{"NumSeconds", duration.NumSeconds}
    return &(duration.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Rate_ErrorNumber_Duration_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "duration"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Continuous_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "continuous"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location.
    Location []Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location
}

func (stop *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "block-name-lst"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range stop.Location {
        stop.EntityData.Children[types.GetSegmentPath(&stop.Location[i])] = types.YChild{"Location", &stop.Location[i]}
    }
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stop.EntityData)
}

// Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location
type Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    FitLocationName interface{}
}

func (location *Set_Asic_Instance_FaultInjection_Module_FaultType_Other_BlockNameLst_Stop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "stop"
    location.EntityData.SegmentPath = "location" + "[fit-location-name='" + fmt.Sprintf("%v", location.FitLocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["fit-location-name"] = types.YLeaf{"FitLocationName", location.FitLocationName}
    return &(location.EntityData)
}

