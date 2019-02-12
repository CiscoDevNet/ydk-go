// This module contains a collection of YANG definitions for
// monitoring transceivers in a Network Element.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package transceiver_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package transceiver_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-transceiver-oper transceiver-oper-data}", reflect.TypeOf(TransceiverOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-transceiver-oper:transceiver-oper-data", reflect.TypeOf(TransceiverOperData{}))
}

// XcvrSonetCode represents SONET/SDH application code supported by the port
type XcvrSonetCode string

const (
    XcvrSonetCode_oc_48_short_reach XcvrSonetCode = "oc-48-short-reach"

    XcvrSonetCode_oc_48_intermediate_reach XcvrSonetCode = "oc-48-intermediate-reach"

    XcvrSonetCode_oc_48_long_reach XcvrSonetCode = "oc-48-long-reach"

    XcvrSonetCode_sonet_sr_compliant XcvrSonetCode = "sonet-sr-compliant"

    XcvrSonetCode_sonet_lr_compliant XcvrSonetCode = "sonet-lr-compliant"

    XcvrSonetCode_oc_192_short_reach XcvrSonetCode = "oc-192-short-reach"

    XcvrSonetCode_escon_smf_1310_laser XcvrSonetCode = "escon-smf-1310-laser"

    XcvrSonetCode_escon_mmf_1310_led XcvrSonetCode = "escon-mmf-1310-led"

    XcvrSonetCode_unknown XcvrSonetCode = "unknown"
)

// XcvrOtnCode represents OTN application code supported by the port
type XcvrOtnCode string

const (
    XcvrOtnCode_p1l1_2d1 XcvrOtnCode = "p1l1-2d1"

    XcvrOtnCode_p1s1_2d2 XcvrOtnCode = "p1s1-2d2"

    XcvrOtnCode_p1l1_2d2 XcvrOtnCode = "p1l1-2d2"

    XcvrOtnCode_otn_undefined XcvrOtnCode = "otn-undefined"
)

// TransceiverOperData
// Top-level container for transceiver operational data
type TransceiverOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of transceiver instance, keyed by name. The type is slice of
    // TransceiverOperData_Transceiver.
    Transceiver []*TransceiverOperData_Transceiver
}

func (transceiverOperData *TransceiverOperData) GetEntityData() *types.CommonEntityData {
    transceiverOperData.EntityData.YFilter = transceiverOperData.YFilter
    transceiverOperData.EntityData.YangName = "transceiver-oper-data"
    transceiverOperData.EntityData.BundleName = "cisco_ios_xe"
    transceiverOperData.EntityData.ParentYangName = "Cisco-IOS-XE-transceiver-oper"
    transceiverOperData.EntityData.SegmentPath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data"
    transceiverOperData.EntityData.AbsolutePath = transceiverOperData.EntityData.SegmentPath
    transceiverOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transceiverOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transceiverOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transceiverOperData.EntityData.Children = types.NewOrderedMap()
    transceiverOperData.EntityData.Children.Append("transceiver", types.YChild{"Transceiver", nil})
    for i := range transceiverOperData.Transceiver {
        transceiverOperData.EntityData.Children.Append(types.GetSegmentPath(transceiverOperData.Transceiver[i]), types.YChild{"Transceiver", transceiverOperData.Transceiver[i]})
    }
    transceiverOperData.EntityData.Leafs = types.NewOrderedMap()

    transceiverOperData.EntityData.YListKeys = []string {}

    return &(transceiverOperData.EntityData)
}

// TransceiverOperData_Transceiver
// List of transceiver instance, keyed by name
type TransceiverOperData_Transceiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string.
    Name interface{}

    // Turns power on/off to the transceiver. The type is bool.
    Enabled interface{}

    // Transceiver is present on the port. The type is bool.
    Present interface{}

    // Indicates the type of optical transceiver used on this port. The type is
    // string.
    Identifier interface{}

    // Connector type used on this port. The type is string.
    Connector interface{}

    // Ethernet PMD (physical medium dependent sublayer) that the transceiver
    // supports. The SFF/QSFP MSAs have registers for this and CFP MSA has
    // similar. The type is string.
    EthernetPmd interface{}

    // Full name of transceiver vendor. The type is string.
    Vendor interface{}

    // Transceiver vendor's part number. The type is string.
    VendorPart interface{}

    // Transceiver vendor's revision number. The type is string.
    VendorRev interface{}

    // Transceiver serial number. The type is string.
    SerialNo interface{}

    // Indicates if a fault condition exists in the transceiver. The type is bool.
    FaultCondition interface{}

    // Representation of the transceiver date. The type is string.
    Date interface{}

    // SONET/SDH application code supported by the port. The type is
    // XcvrSonetCode.
    Sonet interface{}

    // OTN application code supported by the port. The type is XcvrOtnCode.
    Otn interface{}

    // Internally measured temperature in degrees Celsius. The type is string with
    // range: -92233720368547758.08..92233720368547758.07.
    InternalTemp interface{}

    // The output optical power of overall transceiver(dBm).
    OutputPower TransceiverOperData_Transceiver_OutputPower

    // The input optical power of overall transceiver(dBm).
    InputPower TransceiverOperData_Transceiver_InputPower

    // The current applied by the system to the transmit laser to achieve the
    // output power(mA).
    LaserBiasCurrent TransceiverOperData_Transceiver_LaserBiasCurrent

    // List of physical channel for transceiver. The type is slice of
    // TransceiverOperData_Transceiver_XcvrPhysicalChannel.
    XcvrPhysicalChannel []*TransceiverOperData_Transceiver_XcvrPhysicalChannel
}

func (transceiver *TransceiverOperData_Transceiver) GetEntityData() *types.CommonEntityData {
    transceiver.EntityData.YFilter = transceiver.YFilter
    transceiver.EntityData.YangName = "transceiver"
    transceiver.EntityData.BundleName = "cisco_ios_xe"
    transceiver.EntityData.ParentYangName = "transceiver-oper-data"
    transceiver.EntityData.SegmentPath = "transceiver" + types.AddKeyToken(transceiver.Name, "name")
    transceiver.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/" + transceiver.EntityData.SegmentPath
    transceiver.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transceiver.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transceiver.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transceiver.EntityData.Children = types.NewOrderedMap()
    transceiver.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &transceiver.OutputPower})
    transceiver.EntityData.Children.Append("input-power", types.YChild{"InputPower", &transceiver.InputPower})
    transceiver.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &transceiver.LaserBiasCurrent})
    transceiver.EntityData.Children.Append("xcvr-physical-channel", types.YChild{"XcvrPhysicalChannel", nil})
    for i := range transceiver.XcvrPhysicalChannel {
        transceiver.EntityData.Children.Append(types.GetSegmentPath(transceiver.XcvrPhysicalChannel[i]), types.YChild{"XcvrPhysicalChannel", transceiver.XcvrPhysicalChannel[i]})
    }
    transceiver.EntityData.Leafs = types.NewOrderedMap()
    transceiver.EntityData.Leafs.Append("name", types.YLeaf{"Name", transceiver.Name})
    transceiver.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", transceiver.Enabled})
    transceiver.EntityData.Leafs.Append("present", types.YLeaf{"Present", transceiver.Present})
    transceiver.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", transceiver.Identifier})
    transceiver.EntityData.Leafs.Append("connector", types.YLeaf{"Connector", transceiver.Connector})
    transceiver.EntityData.Leafs.Append("ethernet-pmd", types.YLeaf{"EthernetPmd", transceiver.EthernetPmd})
    transceiver.EntityData.Leafs.Append("vendor", types.YLeaf{"Vendor", transceiver.Vendor})
    transceiver.EntityData.Leafs.Append("vendor-part", types.YLeaf{"VendorPart", transceiver.VendorPart})
    transceiver.EntityData.Leafs.Append("vendor-rev", types.YLeaf{"VendorRev", transceiver.VendorRev})
    transceiver.EntityData.Leafs.Append("serial-no", types.YLeaf{"SerialNo", transceiver.SerialNo})
    transceiver.EntityData.Leafs.Append("fault-condition", types.YLeaf{"FaultCondition", transceiver.FaultCondition})
    transceiver.EntityData.Leafs.Append("date", types.YLeaf{"Date", transceiver.Date})
    transceiver.EntityData.Leafs.Append("sonet", types.YLeaf{"Sonet", transceiver.Sonet})
    transceiver.EntityData.Leafs.Append("otn", types.YLeaf{"Otn", transceiver.Otn})
    transceiver.EntityData.Leafs.Append("internal-temp", types.YLeaf{"InternalTemp", transceiver.InternalTemp})

    transceiver.EntityData.YListKeys = []string {"Name"}

    return &(transceiver.EntityData)
}

// TransceiverOperData_Transceiver_OutputPower
// The output optical power of overall transceiver(dBm)
type TransceiverOperData_Transceiver_OutputPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (outputPower *TransceiverOperData_Transceiver_OutputPower) GetEntityData() *types.CommonEntityData {
    outputPower.EntityData.YFilter = outputPower.YFilter
    outputPower.EntityData.YangName = "output-power"
    outputPower.EntityData.BundleName = "cisco_ios_xe"
    outputPower.EntityData.ParentYangName = "transceiver"
    outputPower.EntityData.SegmentPath = "output-power"
    outputPower.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/" + outputPower.EntityData.SegmentPath
    outputPower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    outputPower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    outputPower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    outputPower.EntityData.Children = types.NewOrderedMap()
    outputPower.EntityData.Leafs = types.NewOrderedMap()
    outputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPower.Instant})
    outputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPower.Avg})
    outputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPower.Max})
    outputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPower.Min})

    outputPower.EntityData.YListKeys = []string {}

    return &(outputPower.EntityData)
}

// TransceiverOperData_Transceiver_InputPower
// The input optical power of overall transceiver(dBm)
type TransceiverOperData_Transceiver_InputPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (inputPower *TransceiverOperData_Transceiver_InputPower) GetEntityData() *types.CommonEntityData {
    inputPower.EntityData.YFilter = inputPower.YFilter
    inputPower.EntityData.YangName = "input-power"
    inputPower.EntityData.BundleName = "cisco_ios_xe"
    inputPower.EntityData.ParentYangName = "transceiver"
    inputPower.EntityData.SegmentPath = "input-power"
    inputPower.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/" + inputPower.EntityData.SegmentPath
    inputPower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inputPower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inputPower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inputPower.EntityData.Children = types.NewOrderedMap()
    inputPower.EntityData.Leafs = types.NewOrderedMap()
    inputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPower.Instant})
    inputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPower.Avg})
    inputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPower.Max})
    inputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPower.Min})

    inputPower.EntityData.YListKeys = []string {}

    return &(inputPower.EntityData)
}

// TransceiverOperData_Transceiver_LaserBiasCurrent
// The current applied by the system to the transmit laser to achieve the output power(mA)
type TransceiverOperData_Transceiver_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (laserBiasCurrent *TransceiverOperData_Transceiver_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
    laserBiasCurrent.EntityData.YFilter = laserBiasCurrent.YFilter
    laserBiasCurrent.EntityData.YangName = "laser-bias-current"
    laserBiasCurrent.EntityData.BundleName = "cisco_ios_xe"
    laserBiasCurrent.EntityData.ParentYangName = "transceiver"
    laserBiasCurrent.EntityData.SegmentPath = "laser-bias-current"
    laserBiasCurrent.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/" + laserBiasCurrent.EntityData.SegmentPath
    laserBiasCurrent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    laserBiasCurrent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    laserBiasCurrent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    laserBiasCurrent.EntityData.Children = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", laserBiasCurrent.Instant})
    laserBiasCurrent.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", laserBiasCurrent.Avg})
    laserBiasCurrent.EntityData.Leafs.Append("max", types.YLeaf{"Max", laserBiasCurrent.Max})
    laserBiasCurrent.EntityData.Leafs.Append("min", types.YLeaf{"Min", laserBiasCurrent.Min})

    laserBiasCurrent.EntityData.YListKeys = []string {}

    return &(laserBiasCurrent.EntityData)
}

// TransceiverOperData_Transceiver_XcvrPhysicalChannel
// List of physical channel for transceiver
type TransceiverOperData_Transceiver_XcvrPhysicalChannel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the physical channnel. The type is
    // interface{} with range: 0..65535.
    Index interface{}

    // Text description for the client physical channel. The type is string.
    Description interface{}

    // Enable (true) or disable (false) the transmit label for the channel. The
    // type is bool.
    TxLaser interface{}

    // Target output optical power level of the optical channel(dBm). The type is
    // string with range: -92233720368547758.08..92233720368547758.07.
    TargetOutputPower interface{}

    // The frequency in MHz of the individual physical channel. The type is
    // interface{} with range: 0..18446744073709551615.
    OutputFrequency interface{}

    // The output optical power of a physical channel(dBm).
    OutputPower TransceiverOperData_Transceiver_XcvrPhysicalChannel_OutputPower

    // The input optical power of a physical channel(dBm).
    InputPower TransceiverOperData_Transceiver_XcvrPhysicalChannel_InputPower

    // The current applied by the system to the transmit laser to achieve the
    // output power(mA).
    LaserBiasCurrent TransceiverOperData_Transceiver_XcvrPhysicalChannel_LaserBiasCurrent
}

func (xcvrPhysicalChannel *TransceiverOperData_Transceiver_XcvrPhysicalChannel) GetEntityData() *types.CommonEntityData {
    xcvrPhysicalChannel.EntityData.YFilter = xcvrPhysicalChannel.YFilter
    xcvrPhysicalChannel.EntityData.YangName = "xcvr-physical-channel"
    xcvrPhysicalChannel.EntityData.BundleName = "cisco_ios_xe"
    xcvrPhysicalChannel.EntityData.ParentYangName = "transceiver"
    xcvrPhysicalChannel.EntityData.SegmentPath = "xcvr-physical-channel" + types.AddKeyToken(xcvrPhysicalChannel.Index, "index")
    xcvrPhysicalChannel.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/" + xcvrPhysicalChannel.EntityData.SegmentPath
    xcvrPhysicalChannel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    xcvrPhysicalChannel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    xcvrPhysicalChannel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    xcvrPhysicalChannel.EntityData.Children = types.NewOrderedMap()
    xcvrPhysicalChannel.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &xcvrPhysicalChannel.OutputPower})
    xcvrPhysicalChannel.EntityData.Children.Append("input-power", types.YChild{"InputPower", &xcvrPhysicalChannel.InputPower})
    xcvrPhysicalChannel.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &xcvrPhysicalChannel.LaserBiasCurrent})
    xcvrPhysicalChannel.EntityData.Leafs = types.NewOrderedMap()
    xcvrPhysicalChannel.EntityData.Leafs.Append("index", types.YLeaf{"Index", xcvrPhysicalChannel.Index})
    xcvrPhysicalChannel.EntityData.Leafs.Append("description", types.YLeaf{"Description", xcvrPhysicalChannel.Description})
    xcvrPhysicalChannel.EntityData.Leafs.Append("tx-laser", types.YLeaf{"TxLaser", xcvrPhysicalChannel.TxLaser})
    xcvrPhysicalChannel.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", xcvrPhysicalChannel.TargetOutputPower})
    xcvrPhysicalChannel.EntityData.Leafs.Append("output-frequency", types.YLeaf{"OutputFrequency", xcvrPhysicalChannel.OutputFrequency})

    xcvrPhysicalChannel.EntityData.YListKeys = []string {"Index"}

    return &(xcvrPhysicalChannel.EntityData)
}

// TransceiverOperData_Transceiver_XcvrPhysicalChannel_OutputPower
// The output optical power of a physical channel(dBm)
type TransceiverOperData_Transceiver_XcvrPhysicalChannel_OutputPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (outputPower *TransceiverOperData_Transceiver_XcvrPhysicalChannel_OutputPower) GetEntityData() *types.CommonEntityData {
    outputPower.EntityData.YFilter = outputPower.YFilter
    outputPower.EntityData.YangName = "output-power"
    outputPower.EntityData.BundleName = "cisco_ios_xe"
    outputPower.EntityData.ParentYangName = "xcvr-physical-channel"
    outputPower.EntityData.SegmentPath = "output-power"
    outputPower.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/xcvr-physical-channel/" + outputPower.EntityData.SegmentPath
    outputPower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    outputPower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    outputPower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    outputPower.EntityData.Children = types.NewOrderedMap()
    outputPower.EntityData.Leafs = types.NewOrderedMap()
    outputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPower.Instant})
    outputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPower.Avg})
    outputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPower.Max})
    outputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPower.Min})

    outputPower.EntityData.YListKeys = []string {}

    return &(outputPower.EntityData)
}

// TransceiverOperData_Transceiver_XcvrPhysicalChannel_InputPower
// The input optical power of a physical channel(dBm)
type TransceiverOperData_Transceiver_XcvrPhysicalChannel_InputPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (inputPower *TransceiverOperData_Transceiver_XcvrPhysicalChannel_InputPower) GetEntityData() *types.CommonEntityData {
    inputPower.EntityData.YFilter = inputPower.YFilter
    inputPower.EntityData.YangName = "input-power"
    inputPower.EntityData.BundleName = "cisco_ios_xe"
    inputPower.EntityData.ParentYangName = "xcvr-physical-channel"
    inputPower.EntityData.SegmentPath = "input-power"
    inputPower.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/xcvr-physical-channel/" + inputPower.EntityData.SegmentPath
    inputPower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inputPower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inputPower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inputPower.EntityData.Children = types.NewOrderedMap()
    inputPower.EntityData.Leafs = types.NewOrderedMap()
    inputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPower.Instant})
    inputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPower.Avg})
    inputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPower.Max})
    inputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPower.Min})

    inputPower.EntityData.YListKeys = []string {}

    return &(inputPower.EntityData)
}

// TransceiverOperData_Transceiver_XcvrPhysicalChannel_LaserBiasCurrent
// The current applied by the system to the transmit laser to achieve the output power(mA)
type TransceiverOperData_Transceiver_XcvrPhysicalChannel_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instant value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Instant interface{}

    // Average value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Avg interface{}

    // Maximum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Max interface{}

    // Minimum value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Min interface{}
}

func (laserBiasCurrent *TransceiverOperData_Transceiver_XcvrPhysicalChannel_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
    laserBiasCurrent.EntityData.YFilter = laserBiasCurrent.YFilter
    laserBiasCurrent.EntityData.YangName = "laser-bias-current"
    laserBiasCurrent.EntityData.BundleName = "cisco_ios_xe"
    laserBiasCurrent.EntityData.ParentYangName = "xcvr-physical-channel"
    laserBiasCurrent.EntityData.SegmentPath = "laser-bias-current"
    laserBiasCurrent.EntityData.AbsolutePath = "Cisco-IOS-XE-transceiver-oper:transceiver-oper-data/transceiver/xcvr-physical-channel/" + laserBiasCurrent.EntityData.SegmentPath
    laserBiasCurrent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    laserBiasCurrent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    laserBiasCurrent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    laserBiasCurrent.EntityData.Children = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", laserBiasCurrent.Instant})
    laserBiasCurrent.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", laserBiasCurrent.Avg})
    laserBiasCurrent.EntityData.Leafs.Append("max", types.YLeaf{"Max", laserBiasCurrent.Max})
    laserBiasCurrent.EntityData.Leafs.Append("min", types.YLeaf{"Min", laserBiasCurrent.Min})

    laserBiasCurrent.EntityData.YListKeys = []string {}

    return &(laserBiasCurrent.EntityData)
}

