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
    parent types.Entity
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

func (eTHERWIS *ETHERWIS) GetFilter() yfilter.YFilter { return eTHERWIS.YFilter }

func (eTHERWIS *ETHERWIS) SetFilter(yf yfilter.YFilter) { eTHERWIS.YFilter = yf }

func (eTHERWIS *ETHERWIS) GetGoName(yname string) string {
    if yname == "etherWisDeviceTable" { return "Etherwisdevicetable" }
    if yname == "etherWisSectionCurrentTable" { return "Etherwissectioncurrenttable" }
    if yname == "etherWisPathCurrentTable" { return "Etherwispathcurrenttable" }
    if yname == "etherWisFarEndPathCurrentTable" { return "Etherwisfarendpathcurrenttable" }
    return ""
}

func (eTHERWIS *ETHERWIS) GetSegmentPath() string {
    return "ETHER-WIS:ETHER-WIS"
}

func (eTHERWIS *ETHERWIS) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherWisDeviceTable" {
        return &eTHERWIS.Etherwisdevicetable
    }
    if childYangName == "etherWisSectionCurrentTable" {
        return &eTHERWIS.Etherwissectioncurrenttable
    }
    if childYangName == "etherWisPathCurrentTable" {
        return &eTHERWIS.Etherwispathcurrenttable
    }
    if childYangName == "etherWisFarEndPathCurrentTable" {
        return &eTHERWIS.Etherwisfarendpathcurrenttable
    }
    return nil
}

func (eTHERWIS *ETHERWIS) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["etherWisDeviceTable"] = &eTHERWIS.Etherwisdevicetable
    children["etherWisSectionCurrentTable"] = &eTHERWIS.Etherwissectioncurrenttable
    children["etherWisPathCurrentTable"] = &eTHERWIS.Etherwispathcurrenttable
    children["etherWisFarEndPathCurrentTable"] = &eTHERWIS.Etherwisfarendpathcurrenttable
    return children
}

func (eTHERWIS *ETHERWIS) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eTHERWIS *ETHERWIS) GetBundleName() string { return "cisco_ios_xe" }

func (eTHERWIS *ETHERWIS) GetYangName() string { return "ETHER-WIS" }

func (eTHERWIS *ETHERWIS) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eTHERWIS *ETHERWIS) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eTHERWIS *ETHERWIS) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eTHERWIS *ETHERWIS) SetParent(parent types.Entity) { eTHERWIS.parent = parent }

func (eTHERWIS *ETHERWIS) GetParent() types.Entity { return eTHERWIS.parent }

func (eTHERWIS *ETHERWIS) GetParentYangName() string { return "ETHER-WIS" }

// ETHERWIS_Etherwisdevicetable
// The table for Ethernet WIS devices
type ETHERWIS_Etherwisdevicetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Ethernet WIS device table.  For each instance of this
    // object there MUST be a corresponding instance of sonetMediumEntry. The type
    // is slice of ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry.
    Etherwisdeviceentry []ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetFilter() yfilter.YFilter { return etherwisdevicetable.YFilter }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) SetFilter(yf yfilter.YFilter) { etherwisdevicetable.YFilter = yf }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetGoName(yname string) string {
    if yname == "etherWisDeviceEntry" { return "Etherwisdeviceentry" }
    return ""
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetSegmentPath() string {
    return "etherWisDeviceTable"
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherWisDeviceEntry" {
        for _, c := range etherwisdevicetable.Etherwisdeviceentry {
            if etherwisdevicetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry{}
        etherwisdevicetable.Etherwisdeviceentry = append(etherwisdevicetable.Etherwisdeviceentry, child)
        return &etherwisdevicetable.Etherwisdeviceentry[len(etherwisdevicetable.Etherwisdeviceentry)-1]
    }
    return nil
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherwisdevicetable.Etherwisdeviceentry {
        children[etherwisdevicetable.Etherwisdeviceentry[i].GetSegmentPath()] = &etherwisdevicetable.Etherwisdeviceentry[i]
    }
    return children
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetBundleName() string { return "cisco_ios_xe" }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetYangName() string { return "etherWisDeviceTable" }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) SetParent(parent types.Entity) { etherwisdevicetable.parent = parent }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetParent() types.Entity { return etherwisdevicetable.parent }

func (etherwisdevicetable *ETHERWIS_Etherwisdevicetable) GetParentYangName() string { return "ETHER-WIS" }

// ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry
// An entry in the Ethernet WIS device table.  For each
// instance of this object there MUST be a corresponding
// instance of sonetMediumEntry.
type ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry struct {
    parent types.Entity
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

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetFilter() yfilter.YFilter { return etherwisdeviceentry.YFilter }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) SetFilter(yf yfilter.YFilter) { etherwisdeviceentry.YFilter = yf }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "etherWisDeviceTxTestPatternMode" { return "Etherwisdevicetxtestpatternmode" }
    if yname == "etherWisDeviceRxTestPatternMode" { return "Etherwisdevicerxtestpatternmode" }
    if yname == "etherWisDeviceRxTestPatternErrors" { return "Etherwisdevicerxtestpatternerrors" }
    return ""
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetSegmentPath() string {
    return "etherWisDeviceEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwisdeviceentry.Ifindex) + "']"
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = etherwisdeviceentry.Ifindex
    leafs["etherWisDeviceTxTestPatternMode"] = etherwisdeviceentry.Etherwisdevicetxtestpatternmode
    leafs["etherWisDeviceRxTestPatternMode"] = etherwisdeviceentry.Etherwisdevicerxtestpatternmode
    leafs["etherWisDeviceRxTestPatternErrors"] = etherwisdeviceentry.Etherwisdevicerxtestpatternerrors
    return leafs
}

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetBundleName() string { return "cisco_ios_xe" }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetYangName() string { return "etherWisDeviceEntry" }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) SetParent(parent types.Entity) { etherwisdeviceentry.parent = parent }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetParent() types.Entity { return etherwisdeviceentry.parent }

func (etherwisdeviceentry *ETHERWIS_Etherwisdevicetable_Etherwisdeviceentry) GetParentYangName() string { return "etherWisDeviceTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the etherWisSectionCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of sonetSectionCurrentEntry.
    // The type is slice of
    // ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry.
    Etherwissectioncurrententry []ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetFilter() yfilter.YFilter { return etherwissectioncurrenttable.YFilter }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) SetFilter(yf yfilter.YFilter) { etherwissectioncurrenttable.YFilter = yf }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetGoName(yname string) string {
    if yname == "etherWisSectionCurrentEntry" { return "Etherwissectioncurrententry" }
    return ""
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetSegmentPath() string {
    return "etherWisSectionCurrentTable"
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherWisSectionCurrentEntry" {
        for _, c := range etherwissectioncurrenttable.Etherwissectioncurrententry {
            if etherwissectioncurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry{}
        etherwissectioncurrenttable.Etherwissectioncurrententry = append(etherwissectioncurrenttable.Etherwissectioncurrententry, child)
        return &etherwissectioncurrenttable.Etherwissectioncurrententry[len(etherwissectioncurrenttable.Etherwissectioncurrententry)-1]
    }
    return nil
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherwissectioncurrenttable.Etherwissectioncurrententry {
        children[etherwissectioncurrenttable.Etherwissectioncurrententry[i].GetSegmentPath()] = &etherwissectioncurrenttable.Etherwissectioncurrententry[i]
    }
    return children
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetYangName() string { return "etherWisSectionCurrentTable" }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) SetParent(parent types.Entity) { etherwissectioncurrenttable.parent = parent }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetParent() types.Entity { return etherwissectioncurrenttable.parent }

func (etherwissectioncurrenttable *ETHERWIS_Etherwissectioncurrenttable) GetParentYangName() string { return "ETHER-WIS" }

// ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry
// An entry in the etherWisSectionCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetSectionCurrentEntry.
type ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry struct {
    parent types.Entity
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

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetFilter() yfilter.YFilter { return etherwissectioncurrententry.YFilter }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) SetFilter(yf yfilter.YFilter) { etherwissectioncurrententry.YFilter = yf }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "etherWisSectionCurrentJ0Transmitted" { return "Etherwissectioncurrentj0Transmitted" }
    if yname == "etherWisSectionCurrentJ0Received" { return "Etherwissectioncurrentj0Received" }
    return ""
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetSegmentPath() string {
    return "etherWisSectionCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwissectioncurrententry.Ifindex) + "']"
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = etherwissectioncurrententry.Ifindex
    leafs["etherWisSectionCurrentJ0Transmitted"] = etherwissectioncurrententry.Etherwissectioncurrentj0Transmitted
    leafs["etherWisSectionCurrentJ0Received"] = etherwissectioncurrententry.Etherwissectioncurrentj0Received
    return leafs
}

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetYangName() string { return "etherWisSectionCurrentEntry" }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) SetParent(parent types.Entity) { etherwissectioncurrententry.parent = parent }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetParent() types.Entity { return etherwissectioncurrententry.parent }

func (etherwissectioncurrententry *ETHERWIS_Etherwissectioncurrenttable_Etherwissectioncurrententry) GetParentYangName() string { return "etherWisSectionCurrentTable" }

// ETHERWIS_Etherwispathcurrenttable
// The table for the current state of Ethernet WIS paths.
type ETHERWIS_Etherwispathcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the etherWisPathCurrentTable.  For each instance of this object
    // there MUST be a corresponding instance of sonetPathCurrentEntry. The type
    // is slice of ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry.
    Etherwispathcurrententry []ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetFilter() yfilter.YFilter { return etherwispathcurrenttable.YFilter }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) SetFilter(yf yfilter.YFilter) { etherwispathcurrenttable.YFilter = yf }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetGoName(yname string) string {
    if yname == "etherWisPathCurrentEntry" { return "Etherwispathcurrententry" }
    return ""
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetSegmentPath() string {
    return "etherWisPathCurrentTable"
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherWisPathCurrentEntry" {
        for _, c := range etherwispathcurrenttable.Etherwispathcurrententry {
            if etherwispathcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry{}
        etherwispathcurrenttable.Etherwispathcurrententry = append(etherwispathcurrenttable.Etherwispathcurrententry, child)
        return &etherwispathcurrenttable.Etherwispathcurrententry[len(etherwispathcurrenttable.Etherwispathcurrententry)-1]
    }
    return nil
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherwispathcurrenttable.Etherwispathcurrententry {
        children[etherwispathcurrenttable.Etherwispathcurrententry[i].GetSegmentPath()] = &etherwispathcurrenttable.Etherwispathcurrententry[i]
    }
    return children
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetYangName() string { return "etherWisPathCurrentTable" }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) SetParent(parent types.Entity) { etherwispathcurrenttable.parent = parent }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetParent() types.Entity { return etherwispathcurrenttable.parent }

func (etherwispathcurrenttable *ETHERWIS_Etherwispathcurrenttable) GetParentYangName() string { return "ETHER-WIS" }

// ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry
// An entry in the etherWisPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetPathCurrentEntry.
type ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry struct {
    parent types.Entity
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

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetFilter() yfilter.YFilter { return etherwispathcurrententry.YFilter }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) SetFilter(yf yfilter.YFilter) { etherwispathcurrententry.YFilter = yf }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "etherWisPathCurrentStatus" { return "Etherwispathcurrentstatus" }
    if yname == "etherWisPathCurrentJ1Transmitted" { return "Etherwispathcurrentj1Transmitted" }
    if yname == "etherWisPathCurrentJ1Received" { return "Etherwispathcurrentj1Received" }
    return ""
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetSegmentPath() string {
    return "etherWisPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwispathcurrententry.Ifindex) + "']"
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = etherwispathcurrententry.Ifindex
    leafs["etherWisPathCurrentStatus"] = etherwispathcurrententry.Etherwispathcurrentstatus
    leafs["etherWisPathCurrentJ1Transmitted"] = etherwispathcurrententry.Etherwispathcurrentj1Transmitted
    leafs["etherWisPathCurrentJ1Received"] = etherwispathcurrententry.Etherwispathcurrentj1Received
    return leafs
}

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetYangName() string { return "etherWisPathCurrentEntry" }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) SetParent(parent types.Entity) { etherwispathcurrententry.parent = parent }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetParent() types.Entity { return etherwispathcurrententry.parent }

func (etherwispathcurrententry *ETHERWIS_Etherwispathcurrenttable_Etherwispathcurrententry) GetParentYangName() string { return "etherWisPathCurrentTable" }

// ETHERWIS_Etherwisfarendpathcurrenttable
// The table for the current far-end state of Ethernet WIS
// paths.
type ETHERWIS_Etherwisfarendpathcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the etherWisFarEndPathCurrentTable.  For each instance of this
    // object there MUST be a corresponding instance of
    // sonetFarEndPathCurrentEntry. The type is slice of
    // ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry.
    Etherwisfarendpathcurrententry []ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetFilter() yfilter.YFilter { return etherwisfarendpathcurrenttable.YFilter }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) SetFilter(yf yfilter.YFilter) { etherwisfarendpathcurrenttable.YFilter = yf }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetGoName(yname string) string {
    if yname == "etherWisFarEndPathCurrentEntry" { return "Etherwisfarendpathcurrententry" }
    return ""
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetSegmentPath() string {
    return "etherWisFarEndPathCurrentTable"
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherWisFarEndPathCurrentEntry" {
        for _, c := range etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry {
            if etherwisfarendpathcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry{}
        etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry = append(etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry, child)
        return &etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry[len(etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry)-1]
    }
    return nil
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry {
        children[etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry[i].GetSegmentPath()] = &etherwisfarendpathcurrenttable.Etherwisfarendpathcurrententry[i]
    }
    return children
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetYangName() string { return "etherWisFarEndPathCurrentTable" }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) SetParent(parent types.Entity) { etherwisfarendpathcurrenttable.parent = parent }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetParent() types.Entity { return etherwisfarendpathcurrenttable.parent }

func (etherwisfarendpathcurrenttable *ETHERWIS_Etherwisfarendpathcurrenttable) GetParentYangName() string { return "ETHER-WIS" }

// ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry
// An entry in the etherWisFarEndPathCurrentTable.  For each
// instance of this object there MUST be a corresponding
// instance of sonetFarEndPathCurrentEntry.
type ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry struct {
    parent types.Entity
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

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetFilter() yfilter.YFilter { return etherwisfarendpathcurrententry.YFilter }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) SetFilter(yf yfilter.YFilter) { etherwisfarendpathcurrententry.YFilter = yf }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "etherWisFarEndPathCurrentStatus" { return "Etherwisfarendpathcurrentstatus" }
    return ""
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetSegmentPath() string {
    return "etherWisFarEndPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", etherwisfarendpathcurrententry.Ifindex) + "']"
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = etherwisfarendpathcurrententry.Ifindex
    leafs["etherWisFarEndPathCurrentStatus"] = etherwisfarendpathcurrententry.Etherwisfarendpathcurrentstatus
    return leafs
}

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetYangName() string { return "etherWisFarEndPathCurrentEntry" }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) SetParent(parent types.Entity) { etherwisfarendpathcurrententry.parent = parent }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetParent() types.Entity { return etherwisfarendpathcurrententry.parent }

func (etherwisfarendpathcurrententry *ETHERWIS_Etherwisfarendpathcurrenttable_Etherwisfarendpathcurrententry) GetParentYangName() string { return "etherWisFarEndPathCurrentTable" }

