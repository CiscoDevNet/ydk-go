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
    EtherWisDeviceTable ETHERWIS_EtherWisDeviceTable

    // The table for the current state of Ethernet WIS sections.
    EtherWisSectionCurrentTable ETHERWIS_EtherWisSectionCurrentTable

    // The table for the current state of Ethernet WIS paths.
    EtherWisPathCurrentTable ETHERWIS_EtherWisPathCurrentTable

    // The table for the current far-end state of Ethernet WIS paths.
    EtherWisFarEndPathCurrentTable ETHERWIS_EtherWisFarEndPathCurrentTable
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

    eTHERWIS.EntityData.Children = types.NewOrderedMap()
    eTHERWIS.EntityData.Children.Append("etherWisDeviceTable", types.YChild{"EtherWisDeviceTable", &eTHERWIS.EtherWisDeviceTable})
    eTHERWIS.EntityData.Children.Append("etherWisSectionCurrentTable", types.YChild{"EtherWisSectionCurrentTable", &eTHERWIS.EtherWisSectionCurrentTable})
    eTHERWIS.EntityData.Children.Append("etherWisPathCurrentTable", types.YChild{"EtherWisPathCurrentTable", &eTHERWIS.EtherWisPathCurrentTable})
    eTHERWIS.EntityData.Children.Append("etherWisFarEndPathCurrentTable", types.YChild{"EtherWisFarEndPathCurrentTable", &eTHERWIS.EtherWisFarEndPathCurrentTable})
    eTHERWIS.EntityData.Leafs = types.NewOrderedMap()

    eTHERWIS.EntityData.YListKeys = []string {}

    return &(eTHERWIS.EntityData)
}

// ETHERWIS_EtherWisDeviceTable
// The table for Ethernet WIS devices
type ETHERWIS_EtherWisDeviceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Ethernet WIS device table.  For each instance of this
    // object there MUST be a corresponding instance of sonetMediumEntry. The type
    // is slice of ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry.
    EtherWisDeviceEntry []*ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry
}

func (etherWisDeviceTable *ETHERWIS_EtherWisDeviceTable) GetEntityData() *types.CommonEntityData {
    etherWisDeviceTable.EntityData.YFilter = etherWisDeviceTable.YFilter
    etherWisDeviceTable.EntityData.YangName = "etherWisDeviceTable"
    etherWisDeviceTable.EntityData.BundleName = "cisco_ios_xe"
    etherWisDeviceTable.EntityData.ParentYangName = "ETHER-WIS"
    etherWisDeviceTable.EntityData.SegmentPath = "etherWisDeviceTable"
    etherWisDeviceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisDeviceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisDeviceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisDeviceTable.EntityData.Children = types.NewOrderedMap()
    etherWisDeviceTable.EntityData.Children.Append("etherWisDeviceEntry", types.YChild{"EtherWisDeviceEntry", nil})
    for i := range etherWisDeviceTable.EtherWisDeviceEntry {
        etherWisDeviceTable.EntityData.Children.Append(types.GetSegmentPath(etherWisDeviceTable.EtherWisDeviceEntry[i]), types.YChild{"EtherWisDeviceEntry", etherWisDeviceTable.EtherWisDeviceEntry[i]})
    }
    etherWisDeviceTable.EntityData.Leafs = types.NewOrderedMap()

    etherWisDeviceTable.EntityData.YListKeys = []string {}

    return &(etherWisDeviceTable.EntityData)
}

// ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry
// An entry in the Ethernet WIS device table.  For each
// instance of this object there MUST be a corresponding
// instance of sonetMediumEntry.
type ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    // EtherWisDeviceTxTestPatternMode.
    EtherWisDeviceTxTestPatternMode interface{}

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
    // error inconsistentValue. The type is EtherWisDeviceRxTestPatternMode.
    EtherWisDeviceRxTestPatternMode interface{}

    // This object counts the number of errors detected when the WIS receive path
    // is operating in the PRBS31 test pattern mode.  It is reset to zero when the
    // WIS receive path initially enters that mode, and it increments each time
    // the PRBS pattern checker detects an error as described in [IEEE 802.3 Std.]
    // subclause 50.3.8.2 unless its value is 65535, in which case it remains
    // unchanged.  This object is writeable so that it may be reset upon explicit
    // request of a command generator application while the WIS receive path
    // continues to operate in PRBS31 test pattern mode. The type is interface{}
    // with range: 0..65535.
    EtherWisDeviceRxTestPatternErrors interface{}
}

func (etherWisDeviceEntry *ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry) GetEntityData() *types.CommonEntityData {
    etherWisDeviceEntry.EntityData.YFilter = etherWisDeviceEntry.YFilter
    etherWisDeviceEntry.EntityData.YangName = "etherWisDeviceEntry"
    etherWisDeviceEntry.EntityData.BundleName = "cisco_ios_xe"
    etherWisDeviceEntry.EntityData.ParentYangName = "etherWisDeviceTable"
    etherWisDeviceEntry.EntityData.SegmentPath = "etherWisDeviceEntry" + types.AddKeyToken(etherWisDeviceEntry.IfIndex, "ifIndex")
    etherWisDeviceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisDeviceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisDeviceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisDeviceEntry.EntityData.Children = types.NewOrderedMap()
    etherWisDeviceEntry.EntityData.Leafs = types.NewOrderedMap()
    etherWisDeviceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", etherWisDeviceEntry.IfIndex})
    etherWisDeviceEntry.EntityData.Leafs.Append("etherWisDeviceTxTestPatternMode", types.YLeaf{"EtherWisDeviceTxTestPatternMode", etherWisDeviceEntry.EtherWisDeviceTxTestPatternMode})
    etherWisDeviceEntry.EntityData.Leafs.Append("etherWisDeviceRxTestPatternMode", types.YLeaf{"EtherWisDeviceRxTestPatternMode", etherWisDeviceEntry.EtherWisDeviceRxTestPatternMode})
    etherWisDeviceEntry.EntityData.Leafs.Append("etherWisDeviceRxTestPatternErrors", types.YLeaf{"EtherWisDeviceRxTestPatternErrors", etherWisDeviceEntry.EtherWisDeviceRxTestPatternErrors})

    etherWisDeviceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(etherWisDeviceEntry.EntityData)
}

// ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode represents none(1) MUST be rejected with the error inconsistentValue.
type ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode string

const (
    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode_none ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode = "none"

    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode_prbs31 ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode = "prbs31"

    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode_mixedFrequency ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceRxTestPatternMode = "mixedFrequency"
)

// ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode represents MUST be rejected with the error inconsistentValue.
type ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode string

const (
    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode_none ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode = "none"

    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode_squareWave ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode = "squareWave"

    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode_prbs31 ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode = "prbs31"

    ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode_mixedFrequency ETHERWIS_EtherWisDeviceTable_EtherWisDeviceEntry_EtherWisDeviceTxTestPatternMode = "mixedFrequency"
)

// ETHERWIS_EtherWisSectionCurrentTable
// The table for the current state of Ethernet WIS sections.
type ETHERWIS_EtherWisSectionCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisSectionCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of sonetSectionCurrentEntry.
    // The type is slice of
    // ETHERWIS_EtherWisSectionCurrentTable_EtherWisSectionCurrentEntry.
    EtherWisSectionCurrentEntry []*ETHERWIS_EtherWisSectionCurrentTable_EtherWisSectionCurrentEntry
}

func (etherWisSectionCurrentTable *ETHERWIS_EtherWisSectionCurrentTable) GetEntityData() *types.CommonEntityData {
    etherWisSectionCurrentTable.EntityData.YFilter = etherWisSectionCurrentTable.YFilter
    etherWisSectionCurrentTable.EntityData.YangName = "etherWisSectionCurrentTable"
    etherWisSectionCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    etherWisSectionCurrentTable.EntityData.ParentYangName = "ETHER-WIS"
    etherWisSectionCurrentTable.EntityData.SegmentPath = "etherWisSectionCurrentTable"
    etherWisSectionCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisSectionCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisSectionCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisSectionCurrentTable.EntityData.Children = types.NewOrderedMap()
    etherWisSectionCurrentTable.EntityData.Children.Append("etherWisSectionCurrentEntry", types.YChild{"EtherWisSectionCurrentEntry", nil})
    for i := range etherWisSectionCurrentTable.EtherWisSectionCurrentEntry {
        etherWisSectionCurrentTable.EntityData.Children.Append(types.GetSegmentPath(etherWisSectionCurrentTable.EtherWisSectionCurrentEntry[i]), types.YChild{"EtherWisSectionCurrentEntry", etherWisSectionCurrentTable.EtherWisSectionCurrentEntry[i]})
    }
    etherWisSectionCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    etherWisSectionCurrentTable.EntityData.YListKeys = []string {}

    return &(etherWisSectionCurrentTable.EntityData)
}

// ETHERWIS_EtherWisSectionCurrentTable_EtherWisSectionCurrentEntry
// An entry in the etherWisSectionCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetSectionCurrentEntry.
type ETHERWIS_EtherWisSectionCurrentTable_EtherWisSectionCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This is the 16-octet section trace message that is transmitted in the J0
    // byte.  The value SHOULD be '89'h followed by fifteen octets of '00'h (or
    // some cyclic shift thereof) when the section trace function is not used, and
    // the implementation SHOULD use that value (or a cyclic shift thereof) as a
    // default if no other value has been set. The type is string with length: 16.
    EtherWisSectionCurrentJ0Transmitted interface{}

    // This is the 16-octet section trace message that was most recently received
    // in the J0 byte. The type is string with length: 16.
    EtherWisSectionCurrentJ0Received interface{}
}

func (etherWisSectionCurrentEntry *ETHERWIS_EtherWisSectionCurrentTable_EtherWisSectionCurrentEntry) GetEntityData() *types.CommonEntityData {
    etherWisSectionCurrentEntry.EntityData.YFilter = etherWisSectionCurrentEntry.YFilter
    etherWisSectionCurrentEntry.EntityData.YangName = "etherWisSectionCurrentEntry"
    etherWisSectionCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    etherWisSectionCurrentEntry.EntityData.ParentYangName = "etherWisSectionCurrentTable"
    etherWisSectionCurrentEntry.EntityData.SegmentPath = "etherWisSectionCurrentEntry" + types.AddKeyToken(etherWisSectionCurrentEntry.IfIndex, "ifIndex")
    etherWisSectionCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisSectionCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisSectionCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisSectionCurrentEntry.EntityData.Children = types.NewOrderedMap()
    etherWisSectionCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    etherWisSectionCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", etherWisSectionCurrentEntry.IfIndex})
    etherWisSectionCurrentEntry.EntityData.Leafs.Append("etherWisSectionCurrentJ0Transmitted", types.YLeaf{"EtherWisSectionCurrentJ0Transmitted", etherWisSectionCurrentEntry.EtherWisSectionCurrentJ0Transmitted})
    etherWisSectionCurrentEntry.EntityData.Leafs.Append("etherWisSectionCurrentJ0Received", types.YLeaf{"EtherWisSectionCurrentJ0Received", etherWisSectionCurrentEntry.EtherWisSectionCurrentJ0Received})

    etherWisSectionCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(etherWisSectionCurrentEntry.EntityData)
}

// ETHERWIS_EtherWisPathCurrentTable
// The table for the current state of Ethernet WIS paths.
type ETHERWIS_EtherWisPathCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisPathCurrentTable.  For each instance of this object
    // there MUST be a corresponding instance of sonetPathCurrentEntry. The type
    // is slice of ETHERWIS_EtherWisPathCurrentTable_EtherWisPathCurrentEntry.
    EtherWisPathCurrentEntry []*ETHERWIS_EtherWisPathCurrentTable_EtherWisPathCurrentEntry
}

func (etherWisPathCurrentTable *ETHERWIS_EtherWisPathCurrentTable) GetEntityData() *types.CommonEntityData {
    etherWisPathCurrentTable.EntityData.YFilter = etherWisPathCurrentTable.YFilter
    etherWisPathCurrentTable.EntityData.YangName = "etherWisPathCurrentTable"
    etherWisPathCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    etherWisPathCurrentTable.EntityData.ParentYangName = "ETHER-WIS"
    etherWisPathCurrentTable.EntityData.SegmentPath = "etherWisPathCurrentTable"
    etherWisPathCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisPathCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisPathCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisPathCurrentTable.EntityData.Children = types.NewOrderedMap()
    etherWisPathCurrentTable.EntityData.Children.Append("etherWisPathCurrentEntry", types.YChild{"EtherWisPathCurrentEntry", nil})
    for i := range etherWisPathCurrentTable.EtherWisPathCurrentEntry {
        etherWisPathCurrentTable.EntityData.Children.Append(types.GetSegmentPath(etherWisPathCurrentTable.EtherWisPathCurrentEntry[i]), types.YChild{"EtherWisPathCurrentEntry", etherWisPathCurrentTable.EtherWisPathCurrentEntry[i]})
    }
    etherWisPathCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    etherWisPathCurrentTable.EntityData.YListKeys = []string {}

    return &(etherWisPathCurrentTable.EntityData)
}

// ETHERWIS_EtherWisPathCurrentTable_EtherWisPathCurrentEntry
// An entry in the etherWisPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetPathCurrentEntry.
type ETHERWIS_EtherWisPathCurrentTable_EtherWisPathCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    EtherWisPathCurrentStatus interface{}

    // This is the 16-octet path trace message that is transmitted in the J1 byte.
    // The value SHOULD be '89'h followed by fifteen octets of '00'h (or some
    // cyclic shift thereof) when the path trace function is not used, and the
    // implementation SHOULD use that value (or a cyclic shift thereof) as a
    // default if no other value has been set. The type is string with length: 16.
    EtherWisPathCurrentJ1Transmitted interface{}

    // This is the 16-octet path trace message that was most recently received in
    // the J1 byte. The type is string with length: 16.
    EtherWisPathCurrentJ1Received interface{}
}

func (etherWisPathCurrentEntry *ETHERWIS_EtherWisPathCurrentTable_EtherWisPathCurrentEntry) GetEntityData() *types.CommonEntityData {
    etherWisPathCurrentEntry.EntityData.YFilter = etherWisPathCurrentEntry.YFilter
    etherWisPathCurrentEntry.EntityData.YangName = "etherWisPathCurrentEntry"
    etherWisPathCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    etherWisPathCurrentEntry.EntityData.ParentYangName = "etherWisPathCurrentTable"
    etherWisPathCurrentEntry.EntityData.SegmentPath = "etherWisPathCurrentEntry" + types.AddKeyToken(etherWisPathCurrentEntry.IfIndex, "ifIndex")
    etherWisPathCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisPathCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisPathCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisPathCurrentEntry.EntityData.Children = types.NewOrderedMap()
    etherWisPathCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    etherWisPathCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", etherWisPathCurrentEntry.IfIndex})
    etherWisPathCurrentEntry.EntityData.Leafs.Append("etherWisPathCurrentStatus", types.YLeaf{"EtherWisPathCurrentStatus", etherWisPathCurrentEntry.EtherWisPathCurrentStatus})
    etherWisPathCurrentEntry.EntityData.Leafs.Append("etherWisPathCurrentJ1Transmitted", types.YLeaf{"EtherWisPathCurrentJ1Transmitted", etherWisPathCurrentEntry.EtherWisPathCurrentJ1Transmitted})
    etherWisPathCurrentEntry.EntityData.Leafs.Append("etherWisPathCurrentJ1Received", types.YLeaf{"EtherWisPathCurrentJ1Received", etherWisPathCurrentEntry.EtherWisPathCurrentJ1Received})

    etherWisPathCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(etherWisPathCurrentEntry.EntityData)
}

// ETHERWIS_EtherWisFarEndPathCurrentTable
// The table for the current far-end state of Ethernet WIS
// paths.
type ETHERWIS_EtherWisFarEndPathCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the etherWisFarEndPathCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of
    // sonetFarEndPathCurrentEntry. The type is slice of
    // ETHERWIS_EtherWisFarEndPathCurrentTable_EtherWisFarEndPathCurrentEntry.
    EtherWisFarEndPathCurrentEntry []*ETHERWIS_EtherWisFarEndPathCurrentTable_EtherWisFarEndPathCurrentEntry
}

func (etherWisFarEndPathCurrentTable *ETHERWIS_EtherWisFarEndPathCurrentTable) GetEntityData() *types.CommonEntityData {
    etherWisFarEndPathCurrentTable.EntityData.YFilter = etherWisFarEndPathCurrentTable.YFilter
    etherWisFarEndPathCurrentTable.EntityData.YangName = "etherWisFarEndPathCurrentTable"
    etherWisFarEndPathCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    etherWisFarEndPathCurrentTable.EntityData.ParentYangName = "ETHER-WIS"
    etherWisFarEndPathCurrentTable.EntityData.SegmentPath = "etherWisFarEndPathCurrentTable"
    etherWisFarEndPathCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisFarEndPathCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisFarEndPathCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisFarEndPathCurrentTable.EntityData.Children = types.NewOrderedMap()
    etherWisFarEndPathCurrentTable.EntityData.Children.Append("etherWisFarEndPathCurrentEntry", types.YChild{"EtherWisFarEndPathCurrentEntry", nil})
    for i := range etherWisFarEndPathCurrentTable.EtherWisFarEndPathCurrentEntry {
        etherWisFarEndPathCurrentTable.EntityData.Children.Append(types.GetSegmentPath(etherWisFarEndPathCurrentTable.EtherWisFarEndPathCurrentEntry[i]), types.YChild{"EtherWisFarEndPathCurrentEntry", etherWisFarEndPathCurrentTable.EtherWisFarEndPathCurrentEntry[i]})
    }
    etherWisFarEndPathCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    etherWisFarEndPathCurrentTable.EntityData.YListKeys = []string {}

    return &(etherWisFarEndPathCurrentTable.EntityData)
}

// ETHERWIS_EtherWisFarEndPathCurrentTable_EtherWisFarEndPathCurrentEntry
// An entry in the etherWisFarEndPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetFarEndPathCurrentEntry.
type ETHERWIS_EtherWisFarEndPathCurrentTable_EtherWisFarEndPathCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This variable indicates the current status at the far end of the path using
    // a bit map that can indicate multiple defects at once.  The bit positions
    // are assigned as follows:  etherWisFarEndPayloadDefect(0)    A far end
    // payload defect (i.e., far end    PLM-P or LCD-P) is currently being
    // signaled    in G1 bits 5-7.  etherWisFarEndServerDefect(1)    A far end
    // server defect (i.e., far end    LOP-P or AIS-P) is currently being signaled
    // in G1 bits 5-7.  Note:  when this bit is set,    sonetPathSTSRDI MUST be
    // set in the corresponding    instance of sonetPathCurrentStatus. The type is
    // map[string]bool.
    EtherWisFarEndPathCurrentStatus interface{}
}

func (etherWisFarEndPathCurrentEntry *ETHERWIS_EtherWisFarEndPathCurrentTable_EtherWisFarEndPathCurrentEntry) GetEntityData() *types.CommonEntityData {
    etherWisFarEndPathCurrentEntry.EntityData.YFilter = etherWisFarEndPathCurrentEntry.YFilter
    etherWisFarEndPathCurrentEntry.EntityData.YangName = "etherWisFarEndPathCurrentEntry"
    etherWisFarEndPathCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    etherWisFarEndPathCurrentEntry.EntityData.ParentYangName = "etherWisFarEndPathCurrentTable"
    etherWisFarEndPathCurrentEntry.EntityData.SegmentPath = "etherWisFarEndPathCurrentEntry" + types.AddKeyToken(etherWisFarEndPathCurrentEntry.IfIndex, "ifIndex")
    etherWisFarEndPathCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherWisFarEndPathCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherWisFarEndPathCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherWisFarEndPathCurrentEntry.EntityData.Children = types.NewOrderedMap()
    etherWisFarEndPathCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    etherWisFarEndPathCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", etherWisFarEndPathCurrentEntry.IfIndex})
    etherWisFarEndPathCurrentEntry.EntityData.Leafs.Append("etherWisFarEndPathCurrentStatus", types.YLeaf{"EtherWisFarEndPathCurrentStatus", etherWisFarEndPathCurrentEntry.EtherWisFarEndPathCurrentStatus})

    etherWisFarEndPathCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(etherWisFarEndPathCurrentEntry.EntityData)
}

