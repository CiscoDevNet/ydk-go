// The objects in this MIB module are used in conjunction
// with objects in the SONET-MIB and the MAU-MIB to manage
// the Ethernet WAN Interface Sublayer (WIS).
// 
// The following reference is used throughout this MIB module:
// 
// [IEEE 802.3 Std] refers to:
//    IEEE Std 802.3, 2000 Edition: 'IEEE Standard for
//    Information technology - Telecommunications and
//    information exchange between systems - Local and
//    metropolitan area networks - Specific requirements -
//    Part 3: Carrier sense multiple access with collision
//    detection (CSMA/CD) access method and physical layer
//    specifications', as amended by IEEE Std 802.3ae-2002,
//    'IEEE Standard for Carrier Sense Multiple Access with
//    Collision Detection (CSMA/CD) Access Method and
//    Physical Layer Specifications - Media Access Control
//    (MAC) Parameters, Physical Layer and Management
//    Parameters for 10 Gb/s Operation', 30 August 2002.
// 
// Of particular interest are Clause 50, 'WAN Interface
// Sublayer (WIS), type 10GBASE-W', Clause 30, '10Mb/s,
// 100Mb/s, 1000Mb/s, and 10Gb/s MAC Control, and Link
// Aggregation Management', and Clause 45, 'Management
// Data Input/Output (MDIO) Interface'.
// 
// Copyright (C) The Internet Society (2003).  This version
// of this MIB module is part of RFC 3637;  see the RFC
// itself for full legal notices.
package ether_wis

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ether_wis"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ETHER-WIS ETHER-WIS}", reflect.TypeOf(ETHERWIS{}))
    ydk.RegisterEntity("ETHER-WIS:ETHER-WIS", reflect.TypeOf(ETHERWIS{}))
}

// ETHERWIS
type ETHERWIS struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The table for Ethernet WIS devices.
    Etherwisdevicetable ETHERWIS_Etherwisdevicetable

    // The table for the current state of Ethernet WIS sections.
    Etherwissectioncurrenttable ETHERWIS_Etherwissectioncurrenttable

    // The table for the current state of Ethernet WIS paths.
    Etherwispathcurrenttable ETHERWIS_Etherwispathcurrenttable

    // The table for the current far-end state of Ethernet WIS paths.
    Etherwisfarendpathcurrenttable ETHERWIS_Etherwisfarendpathcurrenttable
}

func (eTHERWIS *ETHERWIS) GetEntityData() *types.CommonEntityData {
    eTHERWIS.EntityData.YFilter = eTHERWIS.YFilter
    eTHERWIS.EntityData.YangName = "ETHER-WIS"
    eTHERWIS.EntityData.BundleName = "cisco_ios_xe"
    eTHERWIS.EntityData.ParentYangName = "ETHER-WIS"
    eTHERWIS.EntityData.SegmentPath = "ETHER-WIS:ETHER-WIS"
    eTHERWIS.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eTHERWIS.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eTHERWIS.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eTHERWIS.EntityData.Children = make(map[string]types.YChild)
    eTHERWIS.EntityData.Children["etherWisDeviceTable"] = types.YChild{"Etherwisdevicetable", &eTHERWIS.Etherwisdevicetable}
    eTHERWIS.EntityData.Children["etherWisSectionCurrentTable"] = types.YChild{"Etherwissectioncurrenttable", &eTHERWIS.Etherwissectioncurrenttable}
    eTHERWIS.EntityData.Children["etherWisPathCurrentTable"] = types.YChild{"Etherwispathcurrenttable", &eTHERWIS.Etherwispathcurrenttable}
    eTHERWIS.EntityData.Children["etherWisFarEndPathCurrentTable"] = types.YChild{"Etherwisfarendpathcurrenttable", &eTHERWIS.Etherwisfarendpathcurrenttable}
    eTHERWIS.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eTHERWIS.EntityData)
}

// ETHERWIS_Etherwisdevicetable
// The table for Ethernet WIS devices
type ETHERWIS_Etherwisdevicetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Ethernet WIS device table.  For each instance of this
    // object there MUST be a corresponding instance of sonetMediumEntry. The type
    // is slice of ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry.
    Etherwisdeviceentry []ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetEntityData() *types.CommonEntityData {
    etherwisdevicetable.EntityData.YFilter = etherwisdevicetable.YFilter
    etherwisdevicetable.EntityData.YangName = "etherWisDeviceTable"
    etherwisdevicetable.EntityData.BundleName = "cisco_ios_xe"
    etherwisdevicetable.EntityData.ParentYangName = "ETHER-WIS"
    etherwisdevicetable.EntityData.SegmentPath = "etherWisDeviceTable"
    etherwisdevicetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwisdevicetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwisdevicetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwisdevicetable.EntityData.Children = make(map[string]types.YChild)
    etherwisdevicetable.EntityData.Children["etherWisDeviceEntry"] = types.YChild{"Etherwisdeviceentry", nil}
    for i := range etherwisdevicetable.Etherwisdeviceentry {
        etherwisdevicetable.EntityData.Children[types.GetSegmentPath(&etherwisdevicetable.Etherwisdeviceentry[i])] = types.YChild{"Etherwisdeviceentry", &etherwisdevicetable.Etherwisdeviceentry[i]}
    }
    etherwisdevicetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherwisdevicetable.EntityData)
}

// ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry
// An entry in the Ethernet WIS device table.  For each
// instance of this object there MUST be a corresponding
// instance of sonetMediumEntry.
type ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable controls the transmit test pattern mode. The value none(1)
    // puts the the WIS transmit path into the normal operating mode.  The value
    // squareWave(2) puts the WIS transmit path into the square wave test pattern
    // mode described in [IEEE 802.3 Std.] subclause 50.3.8.1. The value prbs31(3)
    // puts the WIS transmit path into the PRBS31 test pattern mode described in
    // [IEEE 802.3 Std.] subclause 50.3.8.2.  The value mixedFrequency(4) puts the
    // WIS transmit path into the mixed frequency test pattern mode described in
    // [IEEE 802.3 Std.] subclause 50.3.8.3. Any attempt to set this object to a
    // value other than none(1) when the corresponding instance of ifAdminStatus
    // has the value up(1) MUST be rejected with the error inconsistentValue, and
    // any attempt to set the corresponding instance of ifAdminStatus to the value
    // up(1) when an instance of this object has a value other than none(1) MUST
    // be rejected with the error inconsistentValue. The type is
    // Etherwisdevicetxtestpatternmode.
    Etherwisdevicetxtestpatternmode interface{}

    // This variable controls the receive test pattern mode. The value none(1)
    // puts the the WIS receive path into the normal operating mode.  The value
    // prbs31(3) puts the WIS receive path into the PRBS31 test pattern mode
    // described in [IEEE 802.3 Std.] subclause 50.3.8.2.  The value
    // mixedFrequency(4) puts the WIS receive path into the mixed frequency test
    // pattern mode described in [IEEE 802.3 Std.] subclause 50.3.8.3.  Any
    // attempt to set this object to a value other than none(1) when the
    // corresponding instance of ifAdminStatus has the value up(1) MUST be
    // rejected with the error inconsistentValue, and any attempt to set the
    // corresponding instance of ifAdminStatus to the value up(1) when an instance
    // of this object has a value other than none(1) MUST be rejected with the
    // error inconsistentValue. The type is Etherwisdevicerxtestpatternmode.
    Etherwisdevicerxtestpatternmode interface{}

    // This object counts the number of errors detected when the WIS receive path
    // is operating in the PRBS31 test pattern mode.  It is reset to zero when the
    // WIS receive path initially enters that mode, and it increments each time
    // the PRBS pattern checker detects an error as described in [IEEE 802.3 Std.]
    // subclause 50.3.8.2 unless its value is 65535, in which case it remains
    // unchanged.  This object is writeable so that it may be reset upon explicit
    // request of a command generator application while the WIS receive path
    // continues to operate in PRBS31 test pattern mode. The type is interface{}
    // with range: 0..65535.
    Etherwisdevicerxtestpatternerrors interface{}
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetEntityData() *types.CommonEntityData {
    etherwisdeviceentry.EntityData.YFilter = etherwisdeviceentry.YFilter
    etherwisdeviceentry.EntityData.YangName = "etherWisDeviceEntry"
    etherwisdeviceentry.EntityData.BundleName = "cisco_ios_xe"
    etherwisdeviceentry.EntityData.ParentYangName = "etherWisDeviceTable"
    etherwisdeviceentry.EntityData.SegmentPath = "etherWisDeviceEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwisdeviceentry.Ifindex) + "']"
    etherwisdeviceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwisdeviceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwisdeviceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwisdeviceentry.EntityData.Children = make(map[string]types.YChild)
    etherwisdeviceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherwisdeviceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", etherwisdeviceentry.Ifindex}
    etherwisdeviceentry.EntityData.Leafs["etherWisDeviceTxTestPatternMode"] = types.YLeaf{"Etherwisdevicetxtestpatternmode", etherwisdeviceentry.Etherwisdevicetxtestpatternmode}
    etherwisdeviceentry.EntityData.Leafs["etherWisDeviceRxTestPatternMode"] = types.YLeaf{"Etherwisdevicerxtestpatternmode", etherwisdeviceentry.Etherwisdevicerxtestpatternmode}
    etherwisdeviceentry.EntityData.Leafs["etherWisDeviceRxTestPatternErrors"] = types.YLeaf{"Etherwisdevicerxtestpatternerrors", etherwisdeviceentry.Etherwisdevicerxtestpatternerrors}
    return &(etherwisdeviceentry.EntityData)
}

// ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode represents none(1) MUST be rejected with the error inconsistentValue.
type ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode string

const (
    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode_none ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode = "none"

    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode_prbs31 ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode = "prbs31"

    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode_mixedFrequency ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicerxtestpatternmode = "mixedFrequency"
)

// ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode represents MUST be rejected with the error inconsistentValue.
type ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode string

const (
    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode_none ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode = "none"

    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode_squareWave ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode = "squareWave"

    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode_prbs31 ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode = "prbs31"

    ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode_mixedFrequency ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry_Etherwisdevicetxtestpatternmode = "mixedFrequency"
)

// ETHERWIS_Etherwissectioncurrenttable
// The table for the current state of Ethernet WIS sections.
type ETHERWIS_Etherwissectioncurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisSectionCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of sonetSectionCurrentEntry.
    // The type is slice of
    // ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry.
    Etherwissectioncurrententry []ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetEntityData() *types.CommonEntityData {
    etherwissectioncurrenttable.EntityData.YFilter = etherwissectioncurrenttable.YFilter
    etherwissectioncurrenttable.EntityData.YangName = "etherWisSectionCurrentTable"
    etherwissectioncurrenttable.EntityData.BundleName = "cisco_ios_xe"
    etherwissectioncurrenttable.EntityData.ParentYangName = "ETHER-WIS"
    etherwissectioncurrenttable.EntityData.SegmentPath = "etherWisSectionCurrentTable"
    etherwissectioncurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwissectioncurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwissectioncurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwissectioncurrenttable.EntityData.Children = make(map[string]types.YChild)
    etherwissectioncurrenttable.EntityData.Children["etherWisSectionCurrentEntry"] = types.YChild{"Etherwissectioncurrententry", nil}
    for i := range etherwissectioncurrenttable.Etherwissectioncurrententry {
        etherwissectioncurrenttable.EntityData.Children[types.GetSegmentPath(&etherwissectioncurrenttable.Etherwissectioncurrententry[i])] = types.YChild{"Etherwissectioncurrententry", &etherwissectioncurrenttable.Etherwissectioncurrententry[i]}
    }
    etherwissectioncurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherwissectioncurrenttable.EntityData)
}

// ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry
// An entry in the etherWisSectionCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetSectionCurrentEntry.
type ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This is the 16-octet section trace message that is transmitted in the J0
    // byte.  The value SHOULD be '89'h followed by fifteen octets of '00'h (or
    // some cyclic shift thereof) when the section trace function is not used, and
    // the implementation SHOULD use that value (or a cyclic shift thereof) as a
    // default if no other value has been set. The type is string with length: 16.
    Etherwissectioncurrentj0Transmitted interface{}

    // This is the 16-octet section trace message that was most recently received
    // in the J0 byte. The type is string with length: 16.
    Etherwissectioncurrentj0Received interface{}
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetEntityData() *types.CommonEntityData {
    etherwissectioncurrententry.EntityData.YFilter = etherwissectioncurrententry.YFilter
    etherwissectioncurrententry.EntityData.YangName = "etherWisSectionCurrentEntry"
    etherwissectioncurrententry.EntityData.BundleName = "cisco_ios_xe"
    etherwissectioncurrententry.EntityData.ParentYangName = "etherWisSectionCurrentTable"
    etherwissectioncurrententry.EntityData.SegmentPath = "etherWisSectionCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwissectioncurrententry.Ifindex) + "']"
    etherwissectioncurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwissectioncurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwissectioncurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwissectioncurrententry.EntityData.Children = make(map[string]types.YChild)
    etherwissectioncurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherwissectioncurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", etherwissectioncurrententry.Ifindex}
    etherwissectioncurrententry.EntityData.Leafs["etherWisSectionCurrentJ0Transmitted"] = types.YLeaf{"Etherwissectioncurrentj0Transmitted", etherwissectioncurrententry.Etherwissectioncurrentj0Transmitted}
    etherwissectioncurrententry.EntityData.Leafs["etherWisSectionCurrentJ0Received"] = types.YLeaf{"Etherwissectioncurrentj0Received", etherwissectioncurrententry.Etherwissectioncurrentj0Received}
    return &(etherwissectioncurrententry.EntityData)
}

// ETHERWIS_Etherwispathcurrenttable
// The table for the current state of Ethernet WIS paths.
type ETHERWIS_Etherwispathcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisPathCurrentTable.  For each instance of this object
    // there MUST be a corresponding instance of sonetPathCurrentEntry. The type
    // is slice of ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry.
    Etherwispathcurrententry []ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetEntityData() *types.CommonEntityData {
    etherwispathcurrenttable.EntityData.YFilter = etherwispathcurrenttable.YFilter
    etherwispathcurrenttable.EntityData.YangName = "etherWisPathCurrentTable"
    etherwispathcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    etherwispathcurrenttable.EntityData.ParentYangName = "ETHER-WIS"
    etherwispathcurrenttable.EntityData.SegmentPath = "etherWisPathCurrentTable"
    etherwispathcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwispathcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwispathcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwispathcurrenttable.EntityData.Children = make(map[string]types.YChild)
    etherwispathcurrenttable.EntityData.Children["etherWisPathCurrentEntry"] = types.YChild{"Etherwispathcurrententry", nil}
    for i := range etherwispathcurrenttable.Etherwispathcurrententry {
        etherwispathcurrenttable.EntityData.Children[types.GetSegmentPath(&etherwispathcurrenttable.Etherwispathcurrententry[i])] = types.YChild{"Etherwispathcurrententry", &etherwispathcurrenttable.Etherwispathcurrententry[i]}
    }
    etherwispathcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherwispathcurrenttable.EntityData)
}

// ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry
// An entry in the etherWisPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetPathCurrentEntry.
type ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable indicates the current status of the path payload with a bit
    // map that can indicate multiple defects at once.  The bit positions are
    // assigned as follows:  etherWisPathLOP(0)    This bit is set to indicate
    // that an    LOP-P (Loss of Pointer - Path) defect    is being experienced. 
    // Note:  when this    bit is set, sonetPathSTSLOP MUST be set    in the
    // corresponding instance of    sonetPathCurrentStatus.  etherWisPathAIS(1)   
    // This bit is set to indicate that an    AIS-P (Alarm Indication Signal -
    // Path)    defect is being experienced.  Note:  when    this bit is set,
    // sonetPathSTSAIS MUST be    set in the corresponding instance of   
    // sonetPathCurrentStatus.  etherWisPathPLM(1)    This bit is set to indicate
    // that a    PLM-P (Payload Label Mismatch - Path)    defect is being
    // experienced.  Note:  when    this bit is set, sonetPathSignalLabelMismatch 
    // MUST be set in the corresponding instance of    sonetPathCurrentStatus. 
    // etherWisPathLCD(3)    This bit is set to indicate that an    LCD-P (Loss of
    // Codegroup Delination - Path)    defect is being experienced.  Since this   
    // defect is detected by the PCS and not by    the path layer itself, there is
    // no    corresponding bit in sonetPathCurrentStatus. The type is
    // map[string]bool.
    Etherwispathcurrentstatus interface{}

    // This is the 16-octet path trace message that is transmitted in the J1 byte.
    // The value SHOULD be '89'h followed by fifteen octets of '00'h (or some
    // cyclic shift thereof) when the path trace function is not used, and the
    // implementation SHOULD use that value (or a cyclic shift thereof) as a
    // default if no other value has been set. The type is string with length: 16.
    Etherwispathcurrentj1Transmitted interface{}

    // This is the 16-octet path trace message that was most recently received in
    // the J1 byte. The type is string with length: 16.
    Etherwispathcurrentj1Received interface{}
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetEntityData() *types.CommonEntityData {
    etherwispathcurrententry.EntityData.YFilter = etherwispathcurrententry.YFilter
    etherwispathcurrententry.EntityData.YangName = "etherWisPathCurrentEntry"
    etherwispathcurrententry.EntityData.BundleName = "cisco_ios_xe"
    etherwispathcurrententry.EntityData.ParentYangName = "etherWisPathCurrentTable"
    etherwispathcurrententry.EntityData.SegmentPath = "etherWisPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwispathcurrententry.Ifindex) + "']"
    etherwispathcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwispathcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwispathcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwispathcurrententry.EntityData.Children = make(map[string]types.YChild)
    etherwispathcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherwispathcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", etherwispathcurrententry.Ifindex}
    etherwispathcurrententry.EntityData.Leafs["etherWisPathCurrentStatus"] = types.YLeaf{"Etherwispathcurrentstatus", etherwispathcurrententry.Etherwispathcurrentstatus}
    etherwispathcurrententry.EntityData.Leafs["etherWisPathCurrentJ1Transmitted"] = types.YLeaf{"Etherwispathcurrentj1Transmitted", etherwispathcurrententry.Etherwispathcurrentj1Transmitted}
    etherwispathcurrententry.EntityData.Leafs["etherWisPathCurrentJ1Received"] = types.YLeaf{"Etherwispathcurrentj1Received", etherwispathcurrententry.Etherwispathcurrentj1Received}
    return &(etherwispathcurrententry.EntityData)
}

// ETHERWIS_Etherwisfarendpathcurrenttable
// The table for the current far-end state of Ethernet WIS
// paths.
type ETHERWIS_Etherwisfarendpathcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisFarEndPathCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of
    // sonetFarEndPathCurrentEntry. The type is slice of
    // ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry.
    Etherwisfarendpathcurrententry []ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetEntityData() *types.CommonEntityData {
    etherwisfarendpathcurrenttable.EntityData.YFilter = etherwisfarendpathcurrenttable.YFilter
    etherwisfarendpathcurrenttable.EntityData.YangName = "etherWisFarEndPathCurrentTable"
    etherwisfarendpathcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    etherwisfarendpathcurrenttable.EntityData.ParentYangName = "ETHER-WIS"
    etherwisfarendpathcurrenttable.EntityData.SegmentPath = "etherWisFarEndPathCurrentTable"
    etherwisfarendpathcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwisfarendpathcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwisfarendpathcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwisfarendpathcurrenttable.EntityData.Children = make(map[string]types.YChild)
    etherwisfarendpathcurrenttable.EntityData.Children["etherWisFarEndPathCurrentEntry"] = types.YChild{"Etherwisfarendpathcurrententry", nil}
    for i := range etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry {
        etherwisfarendpathcurrenttable.EntityData.Children[types.GetSegmentPath(&etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry[i])] = types.YChild{"Etherwisfarendpathcurrententry", &etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry[i]}
    }
    etherwisfarendpathcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherwisfarendpathcurrenttable.EntityData)
}

// ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry
// An entry in the etherWisFarEndPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetFarEndPathCurrentEntry.
type ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable indicates the current status at the far end of the path using
    // a bit map that can indicate multiple defects at once.  The bit positions
    // are assigned as follows:  etherWisFarEndPayloadDefect(0)    A far end
    // payload defect (i.e., far end    PLM-P or LCD-P) is currently being
    // signaled    in G1 bits 5-7.  etherWisFarEndServerDefect(1)    A far end
    // server defect (i.e., far end    LOP-P or AIS-P) is currently being signaled
    // in G1 bits 5-7.  Note:  when this bit is set,    sonetPathSTSRDI MUST be
    // set in the corresponding    instance of sonetPathCurrentStatus. The type is
    // map[string]bool.
    Etherwisfarendpathcurrentstatus interface{}
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetEntityData() *types.CommonEntityData {
    etherwisfarendpathcurrententry.EntityData.YFilter = etherwisfarendpathcurrententry.YFilter
    etherwisfarendpathcurrententry.EntityData.YangName = "etherWisFarEndPathCurrentEntry"
    etherwisfarendpathcurrententry.EntityData.BundleName = "cisco_ios_xe"
    etherwisfarendpathcurrententry.EntityData.ParentYangName = "etherWisFarEndPathCurrentTable"
    etherwisfarendpathcurrententry.EntityData.SegmentPath = "etherWisFarEndPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwisfarendpathcurrententry.Ifindex) + "']"
    etherwisfarendpathcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherwisfarendpathcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherwisfarendpathcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherwisfarendpathcurrententry.EntityData.Children = make(map[string]types.YChild)
    etherwisfarendpathcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherwisfarendpathcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", etherwisfarendpathcurrententry.Ifindex}
    etherwisfarendpathcurrententry.EntityData.Leafs["etherWisFarEndPathCurrentStatus"] = types.YLeaf{"Etherwisfarendpathcurrentstatus", etherwisfarendpathcurrententry.Etherwisfarendpathcurrentstatus}
    return &(etherwisfarendpathcurrententry.EntityData)
}

