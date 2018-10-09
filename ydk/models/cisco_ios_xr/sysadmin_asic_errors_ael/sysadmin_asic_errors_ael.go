package sysadmin_asic_errors_ael

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_asic_errors_ael"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-asic-errors-ael asic-errors}", reflect.TypeOf(AsicErrors{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-asic-errors-ael:asic-errors", reflect.TypeOf(AsicErrors{}))
}

// AsicErrors
type AsicErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    DeviceName interface{}

    // The type is slice of AsicErrors_Instance.
    Instance []*AsicErrors_Instance

    
    ShowAllInstances AsicErrors_ShowAllInstances
}

func (asicErrors *AsicErrors) GetEntityData() *types.CommonEntityData {
    asicErrors.EntityData.YFilter = asicErrors.YFilter
    asicErrors.EntityData.YangName = "asic-errors"
    asicErrors.EntityData.BundleName = "cisco_ios_xr"
    asicErrors.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-asic-errors-ael"
    asicErrors.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-asic-errors-ael:asic-errors" + types.AddKeyToken(asicErrors.DeviceName, "device-name")
    asicErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrors.EntityData.Children = types.NewOrderedMap()
    asicErrors.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range asicErrors.Instance {
        asicErrors.EntityData.Children.Append(types.GetSegmentPath(asicErrors.Instance[i]), types.YChild{"Instance", asicErrors.Instance[i]})
    }
    asicErrors.EntityData.Children.Append("show-all-instances", types.YChild{"ShowAllInstances", &asicErrors.ShowAllInstances})
    asicErrors.EntityData.Leafs = types.NewOrderedMap()
    asicErrors.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", asicErrors.DeviceName})

    asicErrors.EntityData.YListKeys = []string {"DeviceName"}

    return &(asicErrors.EntityData)
}

// AsicErrors_Instance
type AsicErrors_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    InstanceNum interface{}

    
    Sbe AsicErrors_Instance_Sbe

    
    Mbe AsicErrors_Instance_Mbe

    
    Parity AsicErrors_Instance_Parity

    
    Generic AsicErrors_Instance_Generic

    
    Crc AsicErrors_Instance_Crc

    
    Reset AsicErrors_Instance_Reset

    
    Barrier AsicErrors_Instance_Barrier

    
    Unexpected AsicErrors_Instance_Unexpected

    
    Link AsicErrors_Instance_Link

    
    OorThresh AsicErrors_Instance_OorThresh

    
    Bp AsicErrors_Instance_Bp

    
    Io AsicErrors_Instance_Io

    
    Ucode AsicErrors_Instance_Ucode

    
    Config AsicErrors_Instance_Config

    
    Indirect AsicErrors_Instance_Indirect

    
    Nonerr AsicErrors_Instance_Nonerr

    
    Summary AsicErrors_Instance_Summary

    
    All AsicErrors_Instance_All
}

func (instance *AsicErrors_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "asic-errors"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceNum, "instance-num")
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("sbe", types.YChild{"Sbe", &instance.Sbe})
    instance.EntityData.Children.Append("mbe", types.YChild{"Mbe", &instance.Mbe})
    instance.EntityData.Children.Append("parity", types.YChild{"Parity", &instance.Parity})
    instance.EntityData.Children.Append("generic", types.YChild{"Generic", &instance.Generic})
    instance.EntityData.Children.Append("crc", types.YChild{"Crc", &instance.Crc})
    instance.EntityData.Children.Append("reset", types.YChild{"Reset", &instance.Reset})
    instance.EntityData.Children.Append("barrier", types.YChild{"Barrier", &instance.Barrier})
    instance.EntityData.Children.Append("unexpected", types.YChild{"Unexpected", &instance.Unexpected})
    instance.EntityData.Children.Append("link", types.YChild{"Link", &instance.Link})
    instance.EntityData.Children.Append("oor-thresh", types.YChild{"OorThresh", &instance.OorThresh})
    instance.EntityData.Children.Append("bp", types.YChild{"Bp", &instance.Bp})
    instance.EntityData.Children.Append("io", types.YChild{"Io", &instance.Io})
    instance.EntityData.Children.Append("ucode", types.YChild{"Ucode", &instance.Ucode})
    instance.EntityData.Children.Append("config", types.YChild{"Config", &instance.Config})
    instance.EntityData.Children.Append("indirect", types.YChild{"Indirect", &instance.Indirect})
    instance.EntityData.Children.Append("nonerr", types.YChild{"Nonerr", &instance.Nonerr})
    instance.EntityData.Children.Append("summary", types.YChild{"Summary", &instance.Summary})
    instance.EntityData.Children.Append("all", types.YChild{"All", &instance.All})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-num", types.YLeaf{"InstanceNum", instance.InstanceNum})

    instance.EntityData.YListKeys = []string {"InstanceNum"}

    return &(instance.EntityData)
}

// AsicErrors_Instance_Sbe
type AsicErrors_Instance_Sbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Sbe_Location.
    Location []*AsicErrors_Instance_Sbe_Location
}

func (sbe *AsicErrors_Instance_Sbe) GetEntityData() *types.CommonEntityData {
    sbe.EntityData.YFilter = sbe.YFilter
    sbe.EntityData.YangName = "sbe"
    sbe.EntityData.BundleName = "cisco_ios_xr"
    sbe.EntityData.ParentYangName = "instance"
    sbe.EntityData.SegmentPath = "sbe"
    sbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbe.EntityData.Children = types.NewOrderedMap()
    sbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range sbe.Location {
        sbe.EntityData.Children.Append(types.GetSegmentPath(sbe.Location[i]), types.YChild{"Location", sbe.Location[i]})
    }
    sbe.EntityData.Leafs = types.NewOrderedMap()

    sbe.EntityData.YListKeys = []string {}

    return &(sbe.EntityData)
}

// AsicErrors_Instance_Sbe_Location
type AsicErrors_Instance_Sbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Sbe_Location_LogLst.
    LogLst []*AsicErrors_Instance_Sbe_Location_LogLst
}

func (location *AsicErrors_Instance_Sbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Sbe_Location_LogLst
type AsicErrors_Instance_Sbe_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Sbe_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Mbe
type AsicErrors_Instance_Mbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Mbe_Location.
    Location []*AsicErrors_Instance_Mbe_Location
}

func (mbe *AsicErrors_Instance_Mbe) GetEntityData() *types.CommonEntityData {
    mbe.EntityData.YFilter = mbe.YFilter
    mbe.EntityData.YangName = "mbe"
    mbe.EntityData.BundleName = "cisco_ios_xr"
    mbe.EntityData.ParentYangName = "instance"
    mbe.EntityData.SegmentPath = "mbe"
    mbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mbe.EntityData.Children = types.NewOrderedMap()
    mbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range mbe.Location {
        mbe.EntityData.Children.Append(types.GetSegmentPath(mbe.Location[i]), types.YChild{"Location", mbe.Location[i]})
    }
    mbe.EntityData.Leafs = types.NewOrderedMap()

    mbe.EntityData.YListKeys = []string {}

    return &(mbe.EntityData)
}

// AsicErrors_Instance_Mbe_Location
type AsicErrors_Instance_Mbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Mbe_Location_LogLst.
    LogLst []*AsicErrors_Instance_Mbe_Location_LogLst
}

func (location *AsicErrors_Instance_Mbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "mbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Mbe_Location_LogLst
type AsicErrors_Instance_Mbe_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Mbe_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Parity
type AsicErrors_Instance_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Parity_Location.
    Location []*AsicErrors_Instance_Parity_Location
}

func (parity *AsicErrors_Instance_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "instance"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = types.NewOrderedMap()
    parity.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range parity.Location {
        parity.EntityData.Children.Append(types.GetSegmentPath(parity.Location[i]), types.YChild{"Location", parity.Location[i]})
    }
    parity.EntityData.Leafs = types.NewOrderedMap()

    parity.EntityData.YListKeys = []string {}

    return &(parity.EntityData)
}

// AsicErrors_Instance_Parity_Location
type AsicErrors_Instance_Parity_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Parity_Location_LogLst.
    LogLst []*AsicErrors_Instance_Parity_Location_LogLst
}

func (location *AsicErrors_Instance_Parity_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "parity"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Parity_Location_LogLst
type AsicErrors_Instance_Parity_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Parity_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Generic
type AsicErrors_Instance_Generic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Generic_Location.
    Location []*AsicErrors_Instance_Generic_Location
}

func (generic *AsicErrors_Instance_Generic) GetEntityData() *types.CommonEntityData {
    generic.EntityData.YFilter = generic.YFilter
    generic.EntityData.YangName = "generic"
    generic.EntityData.BundleName = "cisco_ios_xr"
    generic.EntityData.ParentYangName = "instance"
    generic.EntityData.SegmentPath = "generic"
    generic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generic.EntityData.Children = types.NewOrderedMap()
    generic.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range generic.Location {
        generic.EntityData.Children.Append(types.GetSegmentPath(generic.Location[i]), types.YChild{"Location", generic.Location[i]})
    }
    generic.EntityData.Leafs = types.NewOrderedMap()

    generic.EntityData.YListKeys = []string {}

    return &(generic.EntityData)
}

// AsicErrors_Instance_Generic_Location
type AsicErrors_Instance_Generic_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Generic_Location_LogLst.
    LogLst []*AsicErrors_Instance_Generic_Location_LogLst
}

func (location *AsicErrors_Instance_Generic_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "generic"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Generic_Location_LogLst
type AsicErrors_Instance_Generic_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Generic_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Crc
type AsicErrors_Instance_Crc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Crc_Location.
    Location []*AsicErrors_Instance_Crc_Location
}

func (crc *AsicErrors_Instance_Crc) GetEntityData() *types.CommonEntityData {
    crc.EntityData.YFilter = crc.YFilter
    crc.EntityData.YangName = "crc"
    crc.EntityData.BundleName = "cisco_ios_xr"
    crc.EntityData.ParentYangName = "instance"
    crc.EntityData.SegmentPath = "crc"
    crc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crc.EntityData.Children = types.NewOrderedMap()
    crc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range crc.Location {
        crc.EntityData.Children.Append(types.GetSegmentPath(crc.Location[i]), types.YChild{"Location", crc.Location[i]})
    }
    crc.EntityData.Leafs = types.NewOrderedMap()

    crc.EntityData.YListKeys = []string {}

    return &(crc.EntityData)
}

// AsicErrors_Instance_Crc_Location
type AsicErrors_Instance_Crc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Crc_Location_LogLst.
    LogLst []*AsicErrors_Instance_Crc_Location_LogLst
}

func (location *AsicErrors_Instance_Crc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "crc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Crc_Location_LogLst
type AsicErrors_Instance_Crc_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Crc_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Reset
type AsicErrors_Instance_Reset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Reset_Location.
    Location []*AsicErrors_Instance_Reset_Location
}

func (reset *AsicErrors_Instance_Reset) GetEntityData() *types.CommonEntityData {
    reset.EntityData.YFilter = reset.YFilter
    reset.EntityData.YangName = "reset"
    reset.EntityData.BundleName = "cisco_ios_xr"
    reset.EntityData.ParentYangName = "instance"
    reset.EntityData.SegmentPath = "reset"
    reset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reset.EntityData.Children = types.NewOrderedMap()
    reset.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reset.Location {
        reset.EntityData.Children.Append(types.GetSegmentPath(reset.Location[i]), types.YChild{"Location", reset.Location[i]})
    }
    reset.EntityData.Leafs = types.NewOrderedMap()

    reset.EntityData.YListKeys = []string {}

    return &(reset.EntityData)
}

// AsicErrors_Instance_Reset_Location
type AsicErrors_Instance_Reset_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Reset_Location_LogLst.
    LogLst []*AsicErrors_Instance_Reset_Location_LogLst
}

func (location *AsicErrors_Instance_Reset_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reset"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Reset_Location_LogLst
type AsicErrors_Instance_Reset_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Reset_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Barrier
type AsicErrors_Instance_Barrier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Barrier_Location.
    Location []*AsicErrors_Instance_Barrier_Location
}

func (barrier *AsicErrors_Instance_Barrier) GetEntityData() *types.CommonEntityData {
    barrier.EntityData.YFilter = barrier.YFilter
    barrier.EntityData.YangName = "barrier"
    barrier.EntityData.BundleName = "cisco_ios_xr"
    barrier.EntityData.ParentYangName = "instance"
    barrier.EntityData.SegmentPath = "barrier"
    barrier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrier.EntityData.Children = types.NewOrderedMap()
    barrier.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range barrier.Location {
        barrier.EntityData.Children.Append(types.GetSegmentPath(barrier.Location[i]), types.YChild{"Location", barrier.Location[i]})
    }
    barrier.EntityData.Leafs = types.NewOrderedMap()

    barrier.EntityData.YListKeys = []string {}

    return &(barrier.EntityData)
}

// AsicErrors_Instance_Barrier_Location
type AsicErrors_Instance_Barrier_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Barrier_Location_LogLst.
    LogLst []*AsicErrors_Instance_Barrier_Location_LogLst
}

func (location *AsicErrors_Instance_Barrier_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "barrier"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Barrier_Location_LogLst
type AsicErrors_Instance_Barrier_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Barrier_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Unexpected
type AsicErrors_Instance_Unexpected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Unexpected_Location.
    Location []*AsicErrors_Instance_Unexpected_Location
}

func (unexpected *AsicErrors_Instance_Unexpected) GetEntityData() *types.CommonEntityData {
    unexpected.EntityData.YFilter = unexpected.YFilter
    unexpected.EntityData.YangName = "unexpected"
    unexpected.EntityData.BundleName = "cisco_ios_xr"
    unexpected.EntityData.ParentYangName = "instance"
    unexpected.EntityData.SegmentPath = "unexpected"
    unexpected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpected.EntityData.Children = types.NewOrderedMap()
    unexpected.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range unexpected.Location {
        unexpected.EntityData.Children.Append(types.GetSegmentPath(unexpected.Location[i]), types.YChild{"Location", unexpected.Location[i]})
    }
    unexpected.EntityData.Leafs = types.NewOrderedMap()

    unexpected.EntityData.YListKeys = []string {}

    return &(unexpected.EntityData)
}

// AsicErrors_Instance_Unexpected_Location
type AsicErrors_Instance_Unexpected_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Unexpected_Location_LogLst.
    LogLst []*AsicErrors_Instance_Unexpected_Location_LogLst
}

func (location *AsicErrors_Instance_Unexpected_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "unexpected"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Unexpected_Location_LogLst
type AsicErrors_Instance_Unexpected_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Unexpected_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Link
type AsicErrors_Instance_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Link_Location.
    Location []*AsicErrors_Instance_Link_Location
}

func (link *AsicErrors_Instance_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "cisco_ios_xr"
    link.EntityData.ParentYangName = "instance"
    link.EntityData.SegmentPath = "link"
    link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    link.EntityData.Children = types.NewOrderedMap()
    link.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range link.Location {
        link.EntityData.Children.Append(types.GetSegmentPath(link.Location[i]), types.YChild{"Location", link.Location[i]})
    }
    link.EntityData.Leafs = types.NewOrderedMap()

    link.EntityData.YListKeys = []string {}

    return &(link.EntityData)
}

// AsicErrors_Instance_Link_Location
type AsicErrors_Instance_Link_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Link_Location_LogLst.
    LogLst []*AsicErrors_Instance_Link_Location_LogLst
}

func (location *AsicErrors_Instance_Link_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "link"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Link_Location_LogLst
type AsicErrors_Instance_Link_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Link_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_OorThresh
type AsicErrors_Instance_OorThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_OorThresh_Location.
    Location []*AsicErrors_Instance_OorThresh_Location
}

func (oorThresh *AsicErrors_Instance_OorThresh) GetEntityData() *types.CommonEntityData {
    oorThresh.EntityData.YFilter = oorThresh.YFilter
    oorThresh.EntityData.YangName = "oor-thresh"
    oorThresh.EntityData.BundleName = "cisco_ios_xr"
    oorThresh.EntityData.ParentYangName = "instance"
    oorThresh.EntityData.SegmentPath = "oor-thresh"
    oorThresh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorThresh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorThresh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorThresh.EntityData.Children = types.NewOrderedMap()
    oorThresh.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range oorThresh.Location {
        oorThresh.EntityData.Children.Append(types.GetSegmentPath(oorThresh.Location[i]), types.YChild{"Location", oorThresh.Location[i]})
    }
    oorThresh.EntityData.Leafs = types.NewOrderedMap()

    oorThresh.EntityData.YListKeys = []string {}

    return &(oorThresh.EntityData)
}

// AsicErrors_Instance_OorThresh_Location
type AsicErrors_Instance_OorThresh_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_OorThresh_Location_LogLst.
    LogLst []*AsicErrors_Instance_OorThresh_Location_LogLst
}

func (location *AsicErrors_Instance_OorThresh_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "oor-thresh"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_OorThresh_Location_LogLst
type AsicErrors_Instance_OorThresh_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_OorThresh_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Bp
type AsicErrors_Instance_Bp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Bp_Location.
    Location []*AsicErrors_Instance_Bp_Location
}

func (bp *AsicErrors_Instance_Bp) GetEntityData() *types.CommonEntityData {
    bp.EntityData.YFilter = bp.YFilter
    bp.EntityData.YangName = "bp"
    bp.EntityData.BundleName = "cisco_ios_xr"
    bp.EntityData.ParentYangName = "instance"
    bp.EntityData.SegmentPath = "bp"
    bp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bp.EntityData.Children = types.NewOrderedMap()
    bp.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range bp.Location {
        bp.EntityData.Children.Append(types.GetSegmentPath(bp.Location[i]), types.YChild{"Location", bp.Location[i]})
    }
    bp.EntityData.Leafs = types.NewOrderedMap()

    bp.EntityData.YListKeys = []string {}

    return &(bp.EntityData)
}

// AsicErrors_Instance_Bp_Location
type AsicErrors_Instance_Bp_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Bp_Location_LogLst.
    LogLst []*AsicErrors_Instance_Bp_Location_LogLst
}

func (location *AsicErrors_Instance_Bp_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "bp"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Bp_Location_LogLst
type AsicErrors_Instance_Bp_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Bp_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Io
type AsicErrors_Instance_Io struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Io_Location.
    Location []*AsicErrors_Instance_Io_Location
}

func (io *AsicErrors_Instance_Io) GetEntityData() *types.CommonEntityData {
    io.EntityData.YFilter = io.YFilter
    io.EntityData.YangName = "io"
    io.EntityData.BundleName = "cisco_ios_xr"
    io.EntityData.ParentYangName = "instance"
    io.EntityData.SegmentPath = "io"
    io.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    io.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    io.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    io.EntityData.Children = types.NewOrderedMap()
    io.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range io.Location {
        io.EntityData.Children.Append(types.GetSegmentPath(io.Location[i]), types.YChild{"Location", io.Location[i]})
    }
    io.EntityData.Leafs = types.NewOrderedMap()

    io.EntityData.YListKeys = []string {}

    return &(io.EntityData)
}

// AsicErrors_Instance_Io_Location
type AsicErrors_Instance_Io_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Io_Location_LogLst.
    LogLst []*AsicErrors_Instance_Io_Location_LogLst
}

func (location *AsicErrors_Instance_Io_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "io"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Io_Location_LogLst
type AsicErrors_Instance_Io_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Io_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Ucode
type AsicErrors_Instance_Ucode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Ucode_Location.
    Location []*AsicErrors_Instance_Ucode_Location
}

func (ucode *AsicErrors_Instance_Ucode) GetEntityData() *types.CommonEntityData {
    ucode.EntityData.YFilter = ucode.YFilter
    ucode.EntityData.YangName = "ucode"
    ucode.EntityData.BundleName = "cisco_ios_xr"
    ucode.EntityData.ParentYangName = "instance"
    ucode.EntityData.SegmentPath = "ucode"
    ucode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucode.EntityData.Children = types.NewOrderedMap()
    ucode.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range ucode.Location {
        ucode.EntityData.Children.Append(types.GetSegmentPath(ucode.Location[i]), types.YChild{"Location", ucode.Location[i]})
    }
    ucode.EntityData.Leafs = types.NewOrderedMap()

    ucode.EntityData.YListKeys = []string {}

    return &(ucode.EntityData)
}

// AsicErrors_Instance_Ucode_Location
type AsicErrors_Instance_Ucode_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Ucode_Location_LogLst.
    LogLst []*AsicErrors_Instance_Ucode_Location_LogLst
}

func (location *AsicErrors_Instance_Ucode_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "ucode"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Ucode_Location_LogLst
type AsicErrors_Instance_Ucode_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Ucode_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Config
type AsicErrors_Instance_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Config_Location.
    Location []*AsicErrors_Instance_Config_Location
}

func (config *AsicErrors_Instance_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "instance"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range config.Location {
        config.EntityData.Children.Append(types.GetSegmentPath(config.Location[i]), types.YChild{"Location", config.Location[i]})
    }
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// AsicErrors_Instance_Config_Location
type AsicErrors_Instance_Config_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Config_Location_LogLst.
    LogLst []*AsicErrors_Instance_Config_Location_LogLst
}

func (location *AsicErrors_Instance_Config_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "config"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Config_Location_LogLst
type AsicErrors_Instance_Config_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Config_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Indirect
type AsicErrors_Instance_Indirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Indirect_Location.
    Location []*AsicErrors_Instance_Indirect_Location
}

func (indirect *AsicErrors_Instance_Indirect) GetEntityData() *types.CommonEntityData {
    indirect.EntityData.YFilter = indirect.YFilter
    indirect.EntityData.YangName = "indirect"
    indirect.EntityData.BundleName = "cisco_ios_xr"
    indirect.EntityData.ParentYangName = "instance"
    indirect.EntityData.SegmentPath = "indirect"
    indirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirect.EntityData.Children = types.NewOrderedMap()
    indirect.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range indirect.Location {
        indirect.EntityData.Children.Append(types.GetSegmentPath(indirect.Location[i]), types.YChild{"Location", indirect.Location[i]})
    }
    indirect.EntityData.Leafs = types.NewOrderedMap()

    indirect.EntityData.YListKeys = []string {}

    return &(indirect.EntityData)
}

// AsicErrors_Instance_Indirect_Location
type AsicErrors_Instance_Indirect_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Indirect_Location_LogLst.
    LogLst []*AsicErrors_Instance_Indirect_Location_LogLst
}

func (location *AsicErrors_Instance_Indirect_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "indirect"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Indirect_Location_LogLst
type AsicErrors_Instance_Indirect_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Indirect_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Nonerr
type AsicErrors_Instance_Nonerr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Nonerr_Location.
    Location []*AsicErrors_Instance_Nonerr_Location
}

func (nonerr *AsicErrors_Instance_Nonerr) GetEntityData() *types.CommonEntityData {
    nonerr.EntityData.YFilter = nonerr.YFilter
    nonerr.EntityData.YangName = "nonerr"
    nonerr.EntityData.BundleName = "cisco_ios_xr"
    nonerr.EntityData.ParentYangName = "instance"
    nonerr.EntityData.SegmentPath = "nonerr"
    nonerr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonerr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonerr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonerr.EntityData.Children = types.NewOrderedMap()
    nonerr.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nonerr.Location {
        nonerr.EntityData.Children.Append(types.GetSegmentPath(nonerr.Location[i]), types.YChild{"Location", nonerr.Location[i]})
    }
    nonerr.EntityData.Leafs = types.NewOrderedMap()

    nonerr.EntityData.YListKeys = []string {}

    return &(nonerr.EntityData)
}

// AsicErrors_Instance_Nonerr_Location
type AsicErrors_Instance_Nonerr_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Nonerr_Location_LogLst.
    LogLst []*AsicErrors_Instance_Nonerr_Location_LogLst
}

func (location *AsicErrors_Instance_Nonerr_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nonerr"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Nonerr_Location_LogLst
type AsicErrors_Instance_Nonerr_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Nonerr_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_Summary
type AsicErrors_Instance_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_Summary_Location.
    Location []*AsicErrors_Instance_Summary_Location
}

func (summary *AsicErrors_Instance_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "instance"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range summary.Location {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Location[i]), types.YChild{"Location", summary.Location[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// AsicErrors_Instance_Summary_Location
type AsicErrors_Instance_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_Summary_Location_LogLst.
    LogLst []*AsicErrors_Instance_Summary_Location_LogLst
}

func (location *AsicErrors_Instance_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_Summary_Location_LogLst
type AsicErrors_Instance_Summary_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_Summary_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_All
type AsicErrors_Instance_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History AsicErrors_Instance_All_History

    // The type is slice of AsicErrors_Instance_All_Location.
    Location []*AsicErrors_Instance_All_Location
}

func (all *AsicErrors_Instance_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "instance"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("history", types.YChild{"History", &all.History})
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// AsicErrors_Instance_All_History
type AsicErrors_Instance_All_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_Instance_All_History_Location.
    Location []*AsicErrors_Instance_All_History_Location
}

func (history *AsicErrors_Instance_All_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "all"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range history.Location {
        history.EntityData.Children.Append(types.GetSegmentPath(history.Location[i]), types.YChild{"Location", history.Location[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// AsicErrors_Instance_All_History_Location
type AsicErrors_Instance_All_History_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_All_History_Location_LogLst.
    LogLst []*AsicErrors_Instance_All_History_Location_LogLst
}

func (location *AsicErrors_Instance_All_History_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "history"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_All_History_Location_LogLst
type AsicErrors_Instance_All_History_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_All_History_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_Instance_All_Location
type AsicErrors_Instance_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_Instance_All_Location_LogLst.
    LogLst []*AsicErrors_Instance_All_Location_LogLst
}

func (location *AsicErrors_Instance_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_Instance_All_Location_LogLst
type AsicErrors_Instance_All_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_Instance_All_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances
type AsicErrors_ShowAllInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Sbe AsicErrors_ShowAllInstances_Sbe

    
    Mbe AsicErrors_ShowAllInstances_Mbe

    
    Parity AsicErrors_ShowAllInstances_Parity

    
    Generic AsicErrors_ShowAllInstances_Generic

    
    Crc AsicErrors_ShowAllInstances_Crc

    
    Reset AsicErrors_ShowAllInstances_Reset

    
    Barrier AsicErrors_ShowAllInstances_Barrier

    
    Unexpected AsicErrors_ShowAllInstances_Unexpected

    
    Link AsicErrors_ShowAllInstances_Link

    
    OorThresh AsicErrors_ShowAllInstances_OorThresh

    
    Bp AsicErrors_ShowAllInstances_Bp

    
    Io AsicErrors_ShowAllInstances_Io

    
    Ucode AsicErrors_ShowAllInstances_Ucode

    
    Config AsicErrors_ShowAllInstances_Config

    
    Indirect AsicErrors_ShowAllInstances_Indirect

    
    Nonerr AsicErrors_ShowAllInstances_Nonerr

    
    Summary AsicErrors_ShowAllInstances_Summary

    
    All AsicErrors_ShowAllInstances_All
}

func (showAllInstances *AsicErrors_ShowAllInstances) GetEntityData() *types.CommonEntityData {
    showAllInstances.EntityData.YFilter = showAllInstances.YFilter
    showAllInstances.EntityData.YangName = "show-all-instances"
    showAllInstances.EntityData.BundleName = "cisco_ios_xr"
    showAllInstances.EntityData.ParentYangName = "asic-errors"
    showAllInstances.EntityData.SegmentPath = "show-all-instances"
    showAllInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    showAllInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    showAllInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    showAllInstances.EntityData.Children = types.NewOrderedMap()
    showAllInstances.EntityData.Children.Append("sbe", types.YChild{"Sbe", &showAllInstances.Sbe})
    showAllInstances.EntityData.Children.Append("mbe", types.YChild{"Mbe", &showAllInstances.Mbe})
    showAllInstances.EntityData.Children.Append("parity", types.YChild{"Parity", &showAllInstances.Parity})
    showAllInstances.EntityData.Children.Append("generic", types.YChild{"Generic", &showAllInstances.Generic})
    showAllInstances.EntityData.Children.Append("crc", types.YChild{"Crc", &showAllInstances.Crc})
    showAllInstances.EntityData.Children.Append("reset", types.YChild{"Reset", &showAllInstances.Reset})
    showAllInstances.EntityData.Children.Append("barrier", types.YChild{"Barrier", &showAllInstances.Barrier})
    showAllInstances.EntityData.Children.Append("unexpected", types.YChild{"Unexpected", &showAllInstances.Unexpected})
    showAllInstances.EntityData.Children.Append("link", types.YChild{"Link", &showAllInstances.Link})
    showAllInstances.EntityData.Children.Append("oor-thresh", types.YChild{"OorThresh", &showAllInstances.OorThresh})
    showAllInstances.EntityData.Children.Append("bp", types.YChild{"Bp", &showAllInstances.Bp})
    showAllInstances.EntityData.Children.Append("io", types.YChild{"Io", &showAllInstances.Io})
    showAllInstances.EntityData.Children.Append("ucode", types.YChild{"Ucode", &showAllInstances.Ucode})
    showAllInstances.EntityData.Children.Append("config", types.YChild{"Config", &showAllInstances.Config})
    showAllInstances.EntityData.Children.Append("indirect", types.YChild{"Indirect", &showAllInstances.Indirect})
    showAllInstances.EntityData.Children.Append("nonerr", types.YChild{"Nonerr", &showAllInstances.Nonerr})
    showAllInstances.EntityData.Children.Append("summary", types.YChild{"Summary", &showAllInstances.Summary})
    showAllInstances.EntityData.Children.Append("all", types.YChild{"All", &showAllInstances.All})
    showAllInstances.EntityData.Leafs = types.NewOrderedMap()

    showAllInstances.EntityData.YListKeys = []string {}

    return &(showAllInstances.EntityData)
}

// AsicErrors_ShowAllInstances_Sbe
type AsicErrors_ShowAllInstances_Sbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Sbe_Location.
    Location []*AsicErrors_ShowAllInstances_Sbe_Location
}

func (sbe *AsicErrors_ShowAllInstances_Sbe) GetEntityData() *types.CommonEntityData {
    sbe.EntityData.YFilter = sbe.YFilter
    sbe.EntityData.YangName = "sbe"
    sbe.EntityData.BundleName = "cisco_ios_xr"
    sbe.EntityData.ParentYangName = "show-all-instances"
    sbe.EntityData.SegmentPath = "sbe"
    sbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbe.EntityData.Children = types.NewOrderedMap()
    sbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range sbe.Location {
        sbe.EntityData.Children.Append(types.GetSegmentPath(sbe.Location[i]), types.YChild{"Location", sbe.Location[i]})
    }
    sbe.EntityData.Leafs = types.NewOrderedMap()

    sbe.EntityData.YListKeys = []string {}

    return &(sbe.EntityData)
}

// AsicErrors_ShowAllInstances_Sbe_Location
type AsicErrors_ShowAllInstances_Sbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Sbe_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Sbe_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Sbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Sbe_Location_LogLst
type AsicErrors_ShowAllInstances_Sbe_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Sbe_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Mbe
type AsicErrors_ShowAllInstances_Mbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Mbe_Location.
    Location []*AsicErrors_ShowAllInstances_Mbe_Location
}

func (mbe *AsicErrors_ShowAllInstances_Mbe) GetEntityData() *types.CommonEntityData {
    mbe.EntityData.YFilter = mbe.YFilter
    mbe.EntityData.YangName = "mbe"
    mbe.EntityData.BundleName = "cisco_ios_xr"
    mbe.EntityData.ParentYangName = "show-all-instances"
    mbe.EntityData.SegmentPath = "mbe"
    mbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mbe.EntityData.Children = types.NewOrderedMap()
    mbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range mbe.Location {
        mbe.EntityData.Children.Append(types.GetSegmentPath(mbe.Location[i]), types.YChild{"Location", mbe.Location[i]})
    }
    mbe.EntityData.Leafs = types.NewOrderedMap()

    mbe.EntityData.YListKeys = []string {}

    return &(mbe.EntityData)
}

// AsicErrors_ShowAllInstances_Mbe_Location
type AsicErrors_ShowAllInstances_Mbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Mbe_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Mbe_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Mbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "mbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Mbe_Location_LogLst
type AsicErrors_ShowAllInstances_Mbe_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Mbe_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Parity
type AsicErrors_ShowAllInstances_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Parity_Location.
    Location []*AsicErrors_ShowAllInstances_Parity_Location
}

func (parity *AsicErrors_ShowAllInstances_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "show-all-instances"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = types.NewOrderedMap()
    parity.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range parity.Location {
        parity.EntityData.Children.Append(types.GetSegmentPath(parity.Location[i]), types.YChild{"Location", parity.Location[i]})
    }
    parity.EntityData.Leafs = types.NewOrderedMap()

    parity.EntityData.YListKeys = []string {}

    return &(parity.EntityData)
}

// AsicErrors_ShowAllInstances_Parity_Location
type AsicErrors_ShowAllInstances_Parity_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Parity_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Parity_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Parity_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "parity"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Parity_Location_LogLst
type AsicErrors_ShowAllInstances_Parity_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Parity_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Generic
type AsicErrors_ShowAllInstances_Generic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Generic_Location.
    Location []*AsicErrors_ShowAllInstances_Generic_Location
}

func (generic *AsicErrors_ShowAllInstances_Generic) GetEntityData() *types.CommonEntityData {
    generic.EntityData.YFilter = generic.YFilter
    generic.EntityData.YangName = "generic"
    generic.EntityData.BundleName = "cisco_ios_xr"
    generic.EntityData.ParentYangName = "show-all-instances"
    generic.EntityData.SegmentPath = "generic"
    generic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generic.EntityData.Children = types.NewOrderedMap()
    generic.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range generic.Location {
        generic.EntityData.Children.Append(types.GetSegmentPath(generic.Location[i]), types.YChild{"Location", generic.Location[i]})
    }
    generic.EntityData.Leafs = types.NewOrderedMap()

    generic.EntityData.YListKeys = []string {}

    return &(generic.EntityData)
}

// AsicErrors_ShowAllInstances_Generic_Location
type AsicErrors_ShowAllInstances_Generic_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Generic_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Generic_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Generic_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "generic"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Generic_Location_LogLst
type AsicErrors_ShowAllInstances_Generic_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Generic_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Crc
type AsicErrors_ShowAllInstances_Crc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Crc_Location.
    Location []*AsicErrors_ShowAllInstances_Crc_Location
}

func (crc *AsicErrors_ShowAllInstances_Crc) GetEntityData() *types.CommonEntityData {
    crc.EntityData.YFilter = crc.YFilter
    crc.EntityData.YangName = "crc"
    crc.EntityData.BundleName = "cisco_ios_xr"
    crc.EntityData.ParentYangName = "show-all-instances"
    crc.EntityData.SegmentPath = "crc"
    crc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crc.EntityData.Children = types.NewOrderedMap()
    crc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range crc.Location {
        crc.EntityData.Children.Append(types.GetSegmentPath(crc.Location[i]), types.YChild{"Location", crc.Location[i]})
    }
    crc.EntityData.Leafs = types.NewOrderedMap()

    crc.EntityData.YListKeys = []string {}

    return &(crc.EntityData)
}

// AsicErrors_ShowAllInstances_Crc_Location
type AsicErrors_ShowAllInstances_Crc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Crc_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Crc_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Crc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "crc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Crc_Location_LogLst
type AsicErrors_ShowAllInstances_Crc_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Crc_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Reset
type AsicErrors_ShowAllInstances_Reset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Reset_Location.
    Location []*AsicErrors_ShowAllInstances_Reset_Location
}

func (reset *AsicErrors_ShowAllInstances_Reset) GetEntityData() *types.CommonEntityData {
    reset.EntityData.YFilter = reset.YFilter
    reset.EntityData.YangName = "reset"
    reset.EntityData.BundleName = "cisco_ios_xr"
    reset.EntityData.ParentYangName = "show-all-instances"
    reset.EntityData.SegmentPath = "reset"
    reset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reset.EntityData.Children = types.NewOrderedMap()
    reset.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reset.Location {
        reset.EntityData.Children.Append(types.GetSegmentPath(reset.Location[i]), types.YChild{"Location", reset.Location[i]})
    }
    reset.EntityData.Leafs = types.NewOrderedMap()

    reset.EntityData.YListKeys = []string {}

    return &(reset.EntityData)
}

// AsicErrors_ShowAllInstances_Reset_Location
type AsicErrors_ShowAllInstances_Reset_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Reset_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Reset_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Reset_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reset"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Reset_Location_LogLst
type AsicErrors_ShowAllInstances_Reset_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Reset_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Barrier
type AsicErrors_ShowAllInstances_Barrier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Barrier_Location.
    Location []*AsicErrors_ShowAllInstances_Barrier_Location
}

func (barrier *AsicErrors_ShowAllInstances_Barrier) GetEntityData() *types.CommonEntityData {
    barrier.EntityData.YFilter = barrier.YFilter
    barrier.EntityData.YangName = "barrier"
    barrier.EntityData.BundleName = "cisco_ios_xr"
    barrier.EntityData.ParentYangName = "show-all-instances"
    barrier.EntityData.SegmentPath = "barrier"
    barrier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrier.EntityData.Children = types.NewOrderedMap()
    barrier.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range barrier.Location {
        barrier.EntityData.Children.Append(types.GetSegmentPath(barrier.Location[i]), types.YChild{"Location", barrier.Location[i]})
    }
    barrier.EntityData.Leafs = types.NewOrderedMap()

    barrier.EntityData.YListKeys = []string {}

    return &(barrier.EntityData)
}

// AsicErrors_ShowAllInstances_Barrier_Location
type AsicErrors_ShowAllInstances_Barrier_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Barrier_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Barrier_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Barrier_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "barrier"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Barrier_Location_LogLst
type AsicErrors_ShowAllInstances_Barrier_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Barrier_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Unexpected
type AsicErrors_ShowAllInstances_Unexpected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Unexpected_Location.
    Location []*AsicErrors_ShowAllInstances_Unexpected_Location
}

func (unexpected *AsicErrors_ShowAllInstances_Unexpected) GetEntityData() *types.CommonEntityData {
    unexpected.EntityData.YFilter = unexpected.YFilter
    unexpected.EntityData.YangName = "unexpected"
    unexpected.EntityData.BundleName = "cisco_ios_xr"
    unexpected.EntityData.ParentYangName = "show-all-instances"
    unexpected.EntityData.SegmentPath = "unexpected"
    unexpected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpected.EntityData.Children = types.NewOrderedMap()
    unexpected.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range unexpected.Location {
        unexpected.EntityData.Children.Append(types.GetSegmentPath(unexpected.Location[i]), types.YChild{"Location", unexpected.Location[i]})
    }
    unexpected.EntityData.Leafs = types.NewOrderedMap()

    unexpected.EntityData.YListKeys = []string {}

    return &(unexpected.EntityData)
}

// AsicErrors_ShowAllInstances_Unexpected_Location
type AsicErrors_ShowAllInstances_Unexpected_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of
    // AsicErrors_ShowAllInstances_Unexpected_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Unexpected_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Unexpected_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "unexpected"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Unexpected_Location_LogLst
type AsicErrors_ShowAllInstances_Unexpected_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Unexpected_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Link
type AsicErrors_ShowAllInstances_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Link_Location.
    Location []*AsicErrors_ShowAllInstances_Link_Location
}

func (link *AsicErrors_ShowAllInstances_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "cisco_ios_xr"
    link.EntityData.ParentYangName = "show-all-instances"
    link.EntityData.SegmentPath = "link"
    link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    link.EntityData.Children = types.NewOrderedMap()
    link.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range link.Location {
        link.EntityData.Children.Append(types.GetSegmentPath(link.Location[i]), types.YChild{"Location", link.Location[i]})
    }
    link.EntityData.Leafs = types.NewOrderedMap()

    link.EntityData.YListKeys = []string {}

    return &(link.EntityData)
}

// AsicErrors_ShowAllInstances_Link_Location
type AsicErrors_ShowAllInstances_Link_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Link_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Link_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Link_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "link"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Link_Location_LogLst
type AsicErrors_ShowAllInstances_Link_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Link_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_OorThresh
type AsicErrors_ShowAllInstances_OorThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_OorThresh_Location.
    Location []*AsicErrors_ShowAllInstances_OorThresh_Location
}

func (oorThresh *AsicErrors_ShowAllInstances_OorThresh) GetEntityData() *types.CommonEntityData {
    oorThresh.EntityData.YFilter = oorThresh.YFilter
    oorThresh.EntityData.YangName = "oor-thresh"
    oorThresh.EntityData.BundleName = "cisco_ios_xr"
    oorThresh.EntityData.ParentYangName = "show-all-instances"
    oorThresh.EntityData.SegmentPath = "oor-thresh"
    oorThresh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorThresh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorThresh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorThresh.EntityData.Children = types.NewOrderedMap()
    oorThresh.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range oorThresh.Location {
        oorThresh.EntityData.Children.Append(types.GetSegmentPath(oorThresh.Location[i]), types.YChild{"Location", oorThresh.Location[i]})
    }
    oorThresh.EntityData.Leafs = types.NewOrderedMap()

    oorThresh.EntityData.YListKeys = []string {}

    return &(oorThresh.EntityData)
}

// AsicErrors_ShowAllInstances_OorThresh_Location
type AsicErrors_ShowAllInstances_OorThresh_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_OorThresh_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_OorThresh_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_OorThresh_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "oor-thresh"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_OorThresh_Location_LogLst
type AsicErrors_ShowAllInstances_OorThresh_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_OorThresh_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Bp
type AsicErrors_ShowAllInstances_Bp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Bp_Location.
    Location []*AsicErrors_ShowAllInstances_Bp_Location
}

func (bp *AsicErrors_ShowAllInstances_Bp) GetEntityData() *types.CommonEntityData {
    bp.EntityData.YFilter = bp.YFilter
    bp.EntityData.YangName = "bp"
    bp.EntityData.BundleName = "cisco_ios_xr"
    bp.EntityData.ParentYangName = "show-all-instances"
    bp.EntityData.SegmentPath = "bp"
    bp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bp.EntityData.Children = types.NewOrderedMap()
    bp.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range bp.Location {
        bp.EntityData.Children.Append(types.GetSegmentPath(bp.Location[i]), types.YChild{"Location", bp.Location[i]})
    }
    bp.EntityData.Leafs = types.NewOrderedMap()

    bp.EntityData.YListKeys = []string {}

    return &(bp.EntityData)
}

// AsicErrors_ShowAllInstances_Bp_Location
type AsicErrors_ShowAllInstances_Bp_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Bp_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Bp_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Bp_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "bp"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Bp_Location_LogLst
type AsicErrors_ShowAllInstances_Bp_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Bp_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Io
type AsicErrors_ShowAllInstances_Io struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Io_Location.
    Location []*AsicErrors_ShowAllInstances_Io_Location
}

func (io *AsicErrors_ShowAllInstances_Io) GetEntityData() *types.CommonEntityData {
    io.EntityData.YFilter = io.YFilter
    io.EntityData.YangName = "io"
    io.EntityData.BundleName = "cisco_ios_xr"
    io.EntityData.ParentYangName = "show-all-instances"
    io.EntityData.SegmentPath = "io"
    io.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    io.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    io.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    io.EntityData.Children = types.NewOrderedMap()
    io.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range io.Location {
        io.EntityData.Children.Append(types.GetSegmentPath(io.Location[i]), types.YChild{"Location", io.Location[i]})
    }
    io.EntityData.Leafs = types.NewOrderedMap()

    io.EntityData.YListKeys = []string {}

    return &(io.EntityData)
}

// AsicErrors_ShowAllInstances_Io_Location
type AsicErrors_ShowAllInstances_Io_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Io_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Io_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Io_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "io"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Io_Location_LogLst
type AsicErrors_ShowAllInstances_Io_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Io_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Ucode
type AsicErrors_ShowAllInstances_Ucode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Ucode_Location.
    Location []*AsicErrors_ShowAllInstances_Ucode_Location
}

func (ucode *AsicErrors_ShowAllInstances_Ucode) GetEntityData() *types.CommonEntityData {
    ucode.EntityData.YFilter = ucode.YFilter
    ucode.EntityData.YangName = "ucode"
    ucode.EntityData.BundleName = "cisco_ios_xr"
    ucode.EntityData.ParentYangName = "show-all-instances"
    ucode.EntityData.SegmentPath = "ucode"
    ucode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucode.EntityData.Children = types.NewOrderedMap()
    ucode.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range ucode.Location {
        ucode.EntityData.Children.Append(types.GetSegmentPath(ucode.Location[i]), types.YChild{"Location", ucode.Location[i]})
    }
    ucode.EntityData.Leafs = types.NewOrderedMap()

    ucode.EntityData.YListKeys = []string {}

    return &(ucode.EntityData)
}

// AsicErrors_ShowAllInstances_Ucode_Location
type AsicErrors_ShowAllInstances_Ucode_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Ucode_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Ucode_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Ucode_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "ucode"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Ucode_Location_LogLst
type AsicErrors_ShowAllInstances_Ucode_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Ucode_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Config
type AsicErrors_ShowAllInstances_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Config_Location.
    Location []*AsicErrors_ShowAllInstances_Config_Location
}

func (config *AsicErrors_ShowAllInstances_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "show-all-instances"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range config.Location {
        config.EntityData.Children.Append(types.GetSegmentPath(config.Location[i]), types.YChild{"Location", config.Location[i]})
    }
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// AsicErrors_ShowAllInstances_Config_Location
type AsicErrors_ShowAllInstances_Config_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Config_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Config_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Config_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "config"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Config_Location_LogLst
type AsicErrors_ShowAllInstances_Config_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Config_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Indirect
type AsicErrors_ShowAllInstances_Indirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Indirect_Location.
    Location []*AsicErrors_ShowAllInstances_Indirect_Location
}

func (indirect *AsicErrors_ShowAllInstances_Indirect) GetEntityData() *types.CommonEntityData {
    indirect.EntityData.YFilter = indirect.YFilter
    indirect.EntityData.YangName = "indirect"
    indirect.EntityData.BundleName = "cisco_ios_xr"
    indirect.EntityData.ParentYangName = "show-all-instances"
    indirect.EntityData.SegmentPath = "indirect"
    indirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirect.EntityData.Children = types.NewOrderedMap()
    indirect.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range indirect.Location {
        indirect.EntityData.Children.Append(types.GetSegmentPath(indirect.Location[i]), types.YChild{"Location", indirect.Location[i]})
    }
    indirect.EntityData.Leafs = types.NewOrderedMap()

    indirect.EntityData.YListKeys = []string {}

    return &(indirect.EntityData)
}

// AsicErrors_ShowAllInstances_Indirect_Location
type AsicErrors_ShowAllInstances_Indirect_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Indirect_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Indirect_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Indirect_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "indirect"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Indirect_Location_LogLst
type AsicErrors_ShowAllInstances_Indirect_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Indirect_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Nonerr
type AsicErrors_ShowAllInstances_Nonerr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Nonerr_Location.
    Location []*AsicErrors_ShowAllInstances_Nonerr_Location
}

func (nonerr *AsicErrors_ShowAllInstances_Nonerr) GetEntityData() *types.CommonEntityData {
    nonerr.EntityData.YFilter = nonerr.YFilter
    nonerr.EntityData.YangName = "nonerr"
    nonerr.EntityData.BundleName = "cisco_ios_xr"
    nonerr.EntityData.ParentYangName = "show-all-instances"
    nonerr.EntityData.SegmentPath = "nonerr"
    nonerr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonerr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonerr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonerr.EntityData.Children = types.NewOrderedMap()
    nonerr.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nonerr.Location {
        nonerr.EntityData.Children.Append(types.GetSegmentPath(nonerr.Location[i]), types.YChild{"Location", nonerr.Location[i]})
    }
    nonerr.EntityData.Leafs = types.NewOrderedMap()

    nonerr.EntityData.YListKeys = []string {}

    return &(nonerr.EntityData)
}

// AsicErrors_ShowAllInstances_Nonerr_Location
type AsicErrors_ShowAllInstances_Nonerr_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Nonerr_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Nonerr_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Nonerr_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nonerr"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Nonerr_Location_LogLst
type AsicErrors_ShowAllInstances_Nonerr_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Nonerr_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_Summary
type AsicErrors_ShowAllInstances_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_Summary_Location.
    Location []*AsicErrors_ShowAllInstances_Summary_Location
}

func (summary *AsicErrors_ShowAllInstances_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "show-all-instances"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range summary.Location {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Location[i]), types.YChild{"Location", summary.Location[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// AsicErrors_ShowAllInstances_Summary_Location
type AsicErrors_ShowAllInstances_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_Summary_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_Summary_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_Summary_Location_LogLst
type AsicErrors_ShowAllInstances_Summary_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_Summary_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

// AsicErrors_ShowAllInstances_All
type AsicErrors_ShowAllInstances_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AsicErrors_ShowAllInstances_All_Location.
    Location []*AsicErrors_ShowAllInstances_All_Location
}

func (all *AsicErrors_ShowAllInstances_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "show-all-instances"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// AsicErrors_ShowAllInstances_All_Location
type AsicErrors_ShowAllInstances_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}

    // The type is slice of AsicErrors_ShowAllInstances_All_Location_LogLst.
    LogLst []*AsicErrors_ShowAllInstances_All_Location_LogLst
}

func (location *AsicErrors_ShowAllInstances_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("log-lst", types.YChild{"LogLst", nil})
    for i := range location.LogLst {
        location.EntityData.Children.Append(types.GetSegmentPath(location.LogLst[i]), types.YChild{"LogLst", location.LogLst[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// AsicErrors_ShowAllInstances_All_Location_LogLst
type AsicErrors_ShowAllInstances_All_Location_LogLst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogLine interface{}
}

func (logLst *AsicErrors_ShowAllInstances_All_Location_LogLst) GetEntityData() *types.CommonEntityData {
    logLst.EntityData.YFilter = logLst.YFilter
    logLst.EntityData.YangName = "log-lst"
    logLst.EntityData.BundleName = "cisco_ios_xr"
    logLst.EntityData.ParentYangName = "location"
    logLst.EntityData.SegmentPath = "log-lst"
    logLst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logLst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logLst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logLst.EntityData.Children = types.NewOrderedMap()
    logLst.EntityData.Leafs = types.NewOrderedMap()
    logLst.EntityData.Leafs.Append("log-line", types.YLeaf{"LogLine", logLst.LogLine})

    logLst.EntityData.YListKeys = []string {}

    return &(logLst.EntityData)
}

