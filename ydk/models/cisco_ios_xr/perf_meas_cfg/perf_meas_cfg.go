// This module contains a collection of YANG definitions
// for Cisco IOS-XR perf-meas package configuration.
// 
// This module contains definitions
// for the following management objects:
//   performance-measurement: The root of performance-measurement
//     configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package perf_meas_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package perf_meas_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-perf-meas-cfg performance-measurement}", reflect.TypeOf(PerformanceMeasurement{}))
    ydk.RegisterEntity("Cisco-IOS-XR-perf-meas-cfg:performance-measurement", reflect.TypeOf(PerformanceMeasurement{}))
}

// PerformanceMeasurement
// The root of performance-measurement configuration
type PerformanceMeasurement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable the performance measurement feature. The type is interface{}.
    EnablePerformanceMeasurement interface{}

    // PM Delay Profile.
    DelayProfileInterface PerformanceMeasurement_DelayProfileInterface

    // Configure performance-measurement interfaces.
    Interfaces PerformanceMeasurement_Interfaces
}

func (performanceMeasurement *PerformanceMeasurement) GetEntityData() *types.CommonEntityData {
    performanceMeasurement.EntityData.YFilter = performanceMeasurement.YFilter
    performanceMeasurement.EntityData.YangName = "performance-measurement"
    performanceMeasurement.EntityData.BundleName = "cisco_ios_xr"
    performanceMeasurement.EntityData.ParentYangName = "Cisco-IOS-XR-perf-meas-cfg"
    performanceMeasurement.EntityData.SegmentPath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement"
    performanceMeasurement.EntityData.AbsolutePath = performanceMeasurement.EntityData.SegmentPath
    performanceMeasurement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    performanceMeasurement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    performanceMeasurement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    performanceMeasurement.EntityData.Children = types.NewOrderedMap()
    performanceMeasurement.EntityData.Children.Append("delay-profile-interface", types.YChild{"DelayProfileInterface", &performanceMeasurement.DelayProfileInterface})
    performanceMeasurement.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &performanceMeasurement.Interfaces})
    performanceMeasurement.EntityData.Leafs = types.NewOrderedMap()
    performanceMeasurement.EntityData.Leafs.Append("enable-performance-measurement", types.YLeaf{"EnablePerformanceMeasurement", performanceMeasurement.EnablePerformanceMeasurement})

    performanceMeasurement.EntityData.YListKeys = []string {}

    return &(performanceMeasurement.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface
// PM Delay Profile
type PerformanceMeasurement_DelayProfileInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Advertisement Profile.
    Advertisement PerformanceMeasurement_DelayProfileInterface_Advertisement

    // PM Delay Profile Probe.
    Probe PerformanceMeasurement_DelayProfileInterface_Probe
}

func (delayProfileInterface *PerformanceMeasurement_DelayProfileInterface) GetEntityData() *types.CommonEntityData {
    delayProfileInterface.EntityData.YFilter = delayProfileInterface.YFilter
    delayProfileInterface.EntityData.YangName = "delay-profile-interface"
    delayProfileInterface.EntityData.BundleName = "cisco_ios_xr"
    delayProfileInterface.EntityData.ParentYangName = "performance-measurement"
    delayProfileInterface.EntityData.SegmentPath = "delay-profile-interface"
    delayProfileInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/" + delayProfileInterface.EntityData.SegmentPath
    delayProfileInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayProfileInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayProfileInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayProfileInterface.EntityData.Children = types.NewOrderedMap()
    delayProfileInterface.EntityData.Children.Append("advertisement", types.YChild{"Advertisement", &delayProfileInterface.Advertisement})
    delayProfileInterface.EntityData.Children.Append("probe", types.YChild{"Probe", &delayProfileInterface.Probe})
    delayProfileInterface.EntityData.Leafs = types.NewOrderedMap()

    delayProfileInterface.EntityData.YListKeys = []string {}

    return &(delayProfileInterface.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface_Advertisement
// Advertisement Profile
type PerformanceMeasurement_DelayProfileInterface_Advertisement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accelerated Advertisement Profile.
    Accelerated PerformanceMeasurement_DelayProfileInterface_Advertisement_Accelerated

    // Periodic Advertisement Profile.
    Periodic PerformanceMeasurement_DelayProfileInterface_Advertisement_Periodic
}

func (advertisement *PerformanceMeasurement_DelayProfileInterface_Advertisement) GetEntityData() *types.CommonEntityData {
    advertisement.EntityData.YFilter = advertisement.YFilter
    advertisement.EntityData.YangName = "advertisement"
    advertisement.EntityData.BundleName = "cisco_ios_xr"
    advertisement.EntityData.ParentYangName = "delay-profile-interface"
    advertisement.EntityData.SegmentPath = "advertisement"
    advertisement.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/delay-profile-interface/" + advertisement.EntityData.SegmentPath
    advertisement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisement.EntityData.Children = types.NewOrderedMap()
    advertisement.EntityData.Children.Append("accelerated", types.YChild{"Accelerated", &advertisement.Accelerated})
    advertisement.EntityData.Children.Append("periodic", types.YChild{"Periodic", &advertisement.Periodic})
    advertisement.EntityData.Leafs = types.NewOrderedMap()

    advertisement.EntityData.YListKeys = []string {}

    return &(advertisement.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface_Advertisement_Accelerated
// Accelerated Advertisement Profile
type PerformanceMeasurement_DelayProfileInterface_Advertisement_Accelerated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accelerated advertisement threshold percentage. The type is interface{}
    // with range: 0..100. Units are percentage. The default value is 20.
    Threshold interface{}

    // Accelerated advertisement minimum value in uSec. The type is interface{}
    // with range: 0..100000. The default value is 500.
    MinimumChange interface{}

    // Enable Accelerated Advertisement. The type is interface{}.
    Enable interface{}
}

func (accelerated *PerformanceMeasurement_DelayProfileInterface_Advertisement_Accelerated) GetEntityData() *types.CommonEntityData {
    accelerated.EntityData.YFilter = accelerated.YFilter
    accelerated.EntityData.YangName = "accelerated"
    accelerated.EntityData.BundleName = "cisco_ios_xr"
    accelerated.EntityData.ParentYangName = "advertisement"
    accelerated.EntityData.SegmentPath = "accelerated"
    accelerated.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/delay-profile-interface/advertisement/" + accelerated.EntityData.SegmentPath
    accelerated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accelerated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accelerated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accelerated.EntityData.Children = types.NewOrderedMap()
    accelerated.EntityData.Leafs = types.NewOrderedMap()
    accelerated.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", accelerated.Threshold})
    accelerated.EntityData.Leafs.Append("minimum-change", types.YLeaf{"MinimumChange", accelerated.MinimumChange})
    accelerated.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", accelerated.Enable})

    accelerated.EntityData.YListKeys = []string {}

    return &(accelerated.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface_Advertisement_Periodic
// Periodic Advertisement Profile
type PerformanceMeasurement_DelayProfileInterface_Advertisement_Periodic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Periodic advertisement interval in seconds. The type is interface{} with
    // range: 30..3600. Units are second. The default value is 120.
    Interval interface{}

    // Periodic advertisement threshold percentage. The type is interface{} with
    // range: 0..100. Units are percentage. The default value is 10.
    Threshold interface{}

    // Periodic advertisement minimum value in uSec. The type is interface{} with
    // range: 0..100000. The default value is 500.
    MinimumChange interface{}

    // Enable Performance Measurement Periodic Advertisement. The type is
    // interface{}.
    Disable interface{}
}

func (periodic *PerformanceMeasurement_DelayProfileInterface_Advertisement_Periodic) GetEntityData() *types.CommonEntityData {
    periodic.EntityData.YFilter = periodic.YFilter
    periodic.EntityData.YangName = "periodic"
    periodic.EntityData.BundleName = "cisco_ios_xr"
    periodic.EntityData.ParentYangName = "advertisement"
    periodic.EntityData.SegmentPath = "periodic"
    periodic.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/delay-profile-interface/advertisement/" + periodic.EntityData.SegmentPath
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = types.NewOrderedMap()
    periodic.EntityData.Leafs = types.NewOrderedMap()
    periodic.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", periodic.Interval})
    periodic.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", periodic.Threshold})
    periodic.EntityData.Leafs.Append("minimum-change", types.YLeaf{"MinimumChange", periodic.MinimumChange})
    periodic.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", periodic.Disable})

    periodic.EntityData.YListKeys = []string {}

    return &(periodic.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface_Probe
// PM Delay Profile Probe
type PerformanceMeasurement_DelayProfileInterface_Probe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable one-way measurement. The type is interface{}.
    OneWayMeasurement interface{}

    // The value for delay profile probe interval in seconds. The type is
    // interface{} with range: 30..3600. Units are second. The default value is
    // 30.
    Interval interface{}

    // PM Delay Profile Probe Burst.
    Burst PerformanceMeasurement_DelayProfileInterface_Probe_Burst
}

func (probe *PerformanceMeasurement_DelayProfileInterface_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "delay-profile-interface"
    probe.EntityData.SegmentPath = "probe"
    probe.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/delay-profile-interface/" + probe.EntityData.SegmentPath
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = types.NewOrderedMap()
    probe.EntityData.Children.Append("burst", types.YChild{"Burst", &probe.Burst})
    probe.EntityData.Leafs = types.NewOrderedMap()
    probe.EntityData.Leafs.Append("one-way-measurement", types.YLeaf{"OneWayMeasurement", probe.OneWayMeasurement})
    probe.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", probe.Interval})

    probe.EntityData.YListKeys = []string {}

    return &(probe.EntityData)
}

// PerformanceMeasurement_DelayProfileInterface_Probe_Burst
// PM Delay Profile Probe Burst
type PerformanceMeasurement_DelayProfileInterface_Probe_Burst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value for delay profile probe burst count. The type is interface{} with
    // range: 1..30. The default value is 10.
    Count interface{}

    // The value for delay profile probe burst interval. The type is interface{}
    // with range: 30..15000. The default value is 3000.
    Interval interface{}
}

func (burst *PerformanceMeasurement_DelayProfileInterface_Probe_Burst) GetEntityData() *types.CommonEntityData {
    burst.EntityData.YFilter = burst.YFilter
    burst.EntityData.YangName = "burst"
    burst.EntityData.BundleName = "cisco_ios_xr"
    burst.EntityData.ParentYangName = "probe"
    burst.EntityData.SegmentPath = "burst"
    burst.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/delay-profile-interface/probe/" + burst.EntityData.SegmentPath
    burst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    burst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    burst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    burst.EntityData.Children = types.NewOrderedMap()
    burst.EntityData.Leafs = types.NewOrderedMap()
    burst.EntityData.Leafs.Append("count", types.YLeaf{"Count", burst.Count})
    burst.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", burst.Interval})

    burst.EntityData.YListKeys = []string {}

    return &(burst.EntityData)
}

// PerformanceMeasurement_Interfaces
// Configure performance-measurement interfaces
type PerformanceMeasurement_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a performance-measurement interface. The type is slice of
    // PerformanceMeasurement_Interfaces_Interface.
    Interface []*PerformanceMeasurement_Interfaces_Interface
}

func (interfaces *PerformanceMeasurement_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "performance-measurement"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerformanceMeasurement_Interfaces_Interface
// Configure a performance-measurement interface
type PerformanceMeasurement_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable interface submode. The type is interface{}.
    EnableInterface interface{}

    // Interface delay measurement.
    DelayMeasurement PerformanceMeasurement_Interfaces_Interface_DelayMeasurement
}

func (self *PerformanceMeasurement_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("delay-measurement", types.YChild{"DelayMeasurement", &self.DelayMeasurement})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable-interface", types.YLeaf{"EnableInterface", self.EnableInterface})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// PerformanceMeasurement_Interfaces_Interface_DelayMeasurement
// Interface delay measurement
type PerformanceMeasurement_Interfaces_Interface_DelayMeasurement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable interface delay measurement. The type is interface{}.
    EnableDelayMeasurement interface{}

    // The value for interface delay measurement advertisement delay in uSec. The
    // type is interface{} with range: 0..16777215.
    AdvertiseDelay interface{}
}

func (delayMeasurement *PerformanceMeasurement_Interfaces_Interface_DelayMeasurement) GetEntityData() *types.CommonEntityData {
    delayMeasurement.EntityData.YFilter = delayMeasurement.YFilter
    delayMeasurement.EntityData.YangName = "delay-measurement"
    delayMeasurement.EntityData.BundleName = "cisco_ios_xr"
    delayMeasurement.EntityData.ParentYangName = "interface"
    delayMeasurement.EntityData.SegmentPath = "delay-measurement"
    delayMeasurement.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-cfg:performance-measurement/interfaces/interface/" + delayMeasurement.EntityData.SegmentPath
    delayMeasurement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayMeasurement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayMeasurement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayMeasurement.EntityData.Children = types.NewOrderedMap()
    delayMeasurement.EntityData.Leafs = types.NewOrderedMap()
    delayMeasurement.EntityData.Leafs.Append("enable-delay-measurement", types.YLeaf{"EnableDelayMeasurement", delayMeasurement.EnableDelayMeasurement})
    delayMeasurement.EntityData.Leafs.Append("advertise-delay", types.YLeaf{"AdvertiseDelay", delayMeasurement.AdvertiseDelay})

    delayMeasurement.EntityData.YListKeys = []string {}

    return &(delayMeasurement.EntityData)
}

