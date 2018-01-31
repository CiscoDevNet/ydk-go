// The MIB module for management of the Cisco Discovery
// Protocol in Cisco devices.
package cisco_cdp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cdp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CDP-MIB CISCO-CDP-MIB}", reflect.TypeOf(CISCOCDPMIB{}))
    ydk.RegisterEntity("CISCO-CDP-MIB:CISCO-CDP-MIB", reflect.TypeOf(CISCOCDPMIB{}))
}

// CISCOCDPMIB
type CISCOCDPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cdpglobal CISCOCDPMIB_Cdpglobal

    // The (conceptual) table containing the status of CDP on the device's
    // interfaces.
    Cdpinterfacetable CISCOCDPMIB_Cdpinterfacetable

    // This table contains the additional CDP configuration on the device's
    // interfaces.
    Cdpinterfaceexttable CISCOCDPMIB_Cdpinterfaceexttable

    // The (conceptual) table containing the cached information obtained via
    // receiving CDP messages.
    Cdpcachetable CISCOCDPMIB_Cdpcachetable

    // The (conceptual) table containing the list of  network-layer addresses of a
    // neighbor interface, as reported in the Address TLV of the most recently
    // received CDP message. The first address included in the Address TLV is
    // saved in cdpCacheAddress.  This table contains the remainder of the
    // addresses in the Address TLV.
    Cdpctaddresstable CISCOCDPMIB_Cdpctaddresstable
}

func (cISCOCDPMIB *CISCOCDPMIB) GetFilter() yfilter.YFilter { return cISCOCDPMIB.YFilter }

func (cISCOCDPMIB *CISCOCDPMIB) SetFilter(yf yfilter.YFilter) { cISCOCDPMIB.YFilter = yf }

func (cISCOCDPMIB *CISCOCDPMIB) GetGoName(yname string) string {
    if yname == "cdpGlobal" { return "Cdpglobal" }
    if yname == "cdpInterfaceTable" { return "Cdpinterfacetable" }
    if yname == "cdpInterfaceExtTable" { return "Cdpinterfaceexttable" }
    if yname == "cdpCacheTable" { return "Cdpcachetable" }
    if yname == "cdpCtAddressTable" { return "Cdpctaddresstable" }
    return ""
}

func (cISCOCDPMIB *CISCOCDPMIB) GetSegmentPath() string {
    return "CISCO-CDP-MIB:CISCO-CDP-MIB"
}

func (cISCOCDPMIB *CISCOCDPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdpGlobal" {
        return &cISCOCDPMIB.Cdpglobal
    }
    if childYangName == "cdpInterfaceTable" {
        return &cISCOCDPMIB.Cdpinterfacetable
    }
    if childYangName == "cdpInterfaceExtTable" {
        return &cISCOCDPMIB.Cdpinterfaceexttable
    }
    if childYangName == "cdpCacheTable" {
        return &cISCOCDPMIB.Cdpcachetable
    }
    if childYangName == "cdpCtAddressTable" {
        return &cISCOCDPMIB.Cdpctaddresstable
    }
    return nil
}

func (cISCOCDPMIB *CISCOCDPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cdpGlobal"] = &cISCOCDPMIB.Cdpglobal
    children["cdpInterfaceTable"] = &cISCOCDPMIB.Cdpinterfacetable
    children["cdpInterfaceExtTable"] = &cISCOCDPMIB.Cdpinterfaceexttable
    children["cdpCacheTable"] = &cISCOCDPMIB.Cdpcachetable
    children["cdpCtAddressTable"] = &cISCOCDPMIB.Cdpctaddresstable
    return children
}

func (cISCOCDPMIB *CISCOCDPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCDPMIB *CISCOCDPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCDPMIB *CISCOCDPMIB) GetYangName() string { return "CISCO-CDP-MIB" }

func (cISCOCDPMIB *CISCOCDPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCDPMIB *CISCOCDPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCDPMIB *CISCOCDPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCDPMIB *CISCOCDPMIB) SetParent(parent types.Entity) { cISCOCDPMIB.parent = parent }

func (cISCOCDPMIB *CISCOCDPMIB) GetParent() types.Entity { return cISCOCDPMIB.parent }

func (cISCOCDPMIB *CISCOCDPMIB) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpglobal
type CISCOCDPMIB_Cdpglobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An indication of whether the Cisco Discovery Protocol is currently running.
    // Entries in cdpCacheTable are deleted when CDP is disabled. The type is
    // bool.
    Cdpglobalrun interface{}

    // The interval at which CDP messages are to be generated. The default value
    // is 60 seconds. The type is interface{} with range: 5..254. Units are
    // seconds.
    Cdpglobalmessageinterval interface{}

    // The time for the receiving device holds CDP message. The default value is
    // 180 seconds. The type is interface{} with range: 10..255. Units are
    // seconds.
    Cdpglobalholdtime interface{}

    // The device ID advertised by this device. The format of this device id is
    // characterized by the value of  cdpGlobalDeviceIdFormat object. The type is
    // string.
    Cdpglobaldeviceid interface{}

    // Indicates the time when the cache table was last changed. It is the most
    // recent time at which any row was last created, modified or deleted. The
    // type is interface{} with range: 0..4294967295.
    Cdpgloballastchange interface{}

    // Indicate the Device-Id format capability of the device.  serialNumber(0)
    // indicates that the device supports using serial number as the format for
    // its DeviceId.  macAddress(1) indicates that the device supports using layer
    // 2 MAC address as the format for its DeviceId.  other(2) indicates that the
    // device supports using its platform specific format as the format for its
    // DeviceId. The type is map[string]bool.
    Cdpglobaldeviceidformatcpb interface{}

    // An indication of the format of Device-Id contained in the corresponding
    // instance of cdpGlobalDeviceId. User can only specify the formats that the
    // device is capable of as denoted in cdpGlobalDeviceIdFormatCpb object. 
    // serialNumber(1) indicates that the value of cdpGlobalDeviceId  object is in
    // the form of an ASCII string contain the device serial number.  
    // macAddress(2) indicates that the value of cdpGlobalDeviceId  object is in
    // the form of Layer 2 MAC address.  other(3) indicates that the value of
    // cdpGlobalDeviceId object is in the form of a platform specific ASCII string
    // contain info that identifies the device. For example: ASCII string contains
    // serialNumber appended/prepened with system name. The type is
    // Cdpglobaldeviceidformat.
    Cdpglobaldeviceidformat interface{}
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetFilter() yfilter.YFilter { return cdpglobal.YFilter }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) SetFilter(yf yfilter.YFilter) { cdpglobal.YFilter = yf }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetGoName(yname string) string {
    if yname == "cdpGlobalRun" { return "Cdpglobalrun" }
    if yname == "cdpGlobalMessageInterval" { return "Cdpglobalmessageinterval" }
    if yname == "cdpGlobalHoldTime" { return "Cdpglobalholdtime" }
    if yname == "cdpGlobalDeviceId" { return "Cdpglobaldeviceid" }
    if yname == "cdpGlobalLastChange" { return "Cdpgloballastchange" }
    if yname == "cdpGlobalDeviceIdFormatCpb" { return "Cdpglobaldeviceidformatcpb" }
    if yname == "cdpGlobalDeviceIdFormat" { return "Cdpglobaldeviceidformat" }
    return ""
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetSegmentPath() string {
    return "cdpGlobal"
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdpGlobalRun"] = cdpglobal.Cdpglobalrun
    leafs["cdpGlobalMessageInterval"] = cdpglobal.Cdpglobalmessageinterval
    leafs["cdpGlobalHoldTime"] = cdpglobal.Cdpglobalholdtime
    leafs["cdpGlobalDeviceId"] = cdpglobal.Cdpglobaldeviceid
    leafs["cdpGlobalLastChange"] = cdpglobal.Cdpgloballastchange
    leafs["cdpGlobalDeviceIdFormatCpb"] = cdpglobal.Cdpglobaldeviceidformatcpb
    leafs["cdpGlobalDeviceIdFormat"] = cdpglobal.Cdpglobaldeviceidformat
    return leafs
}

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetBundleName() string { return "cisco_ios_xe" }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetYangName() string { return "cdpGlobal" }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) SetParent(parent types.Entity) { cdpglobal.parent = parent }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetParent() types.Entity { return cdpglobal.parent }

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat represents contains serialNumber appended/prepened with system name.
type CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat string

const (
    CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat_serialNumber CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat = "serialNumber"

    CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat_macAddress CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat = "macAddress"

    CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat_other CISCOCDPMIB_Cdpglobal_Cdpglobaldeviceidformat = "other"
)

// CISCOCDPMIB_Cdpinterfacetable
// The (conceptual) table containing the status of CDP on
// the device's interfaces.
type CISCOCDPMIB_Cdpinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpInterfaceTable, containing the status
    // of CDP on an interface. The type is slice of
    // CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry.
    Cdpinterfaceentry []CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetFilter() yfilter.YFilter { return cdpinterfacetable.YFilter }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) SetFilter(yf yfilter.YFilter) { cdpinterfacetable.YFilter = yf }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetGoName(yname string) string {
    if yname == "cdpInterfaceEntry" { return "Cdpinterfaceentry" }
    return ""
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetSegmentPath() string {
    return "cdpInterfaceTable"
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdpInterfaceEntry" {
        for _, c := range cdpinterfacetable.Cdpinterfaceentry {
            if cdpinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry{}
        cdpinterfacetable.Cdpinterfaceentry = append(cdpinterfacetable.Cdpinterfaceentry, child)
        return &cdpinterfacetable.Cdpinterfaceentry[len(cdpinterfacetable.Cdpinterfaceentry)-1]
    }
    return nil
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdpinterfacetable.Cdpinterfaceentry {
        children[cdpinterfacetable.Cdpinterfaceentry[i].GetSegmentPath()] = &cdpinterfacetable.Cdpinterfaceentry[i]
    }
    return children
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetYangName() string { return "cdpInterfaceTable" }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) SetParent(parent types.Entity) { cdpinterfacetable.parent = parent }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetParent() types.Entity { return cdpinterfacetable.parent }

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry
// An entry (conceptual row) in the cdpInterfaceTable,
// containing the status of CDP on an interface.
type CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the local interface.  For
    // 802.3 Repeaters on which the repeater ports do not have ifIndex values
    // assigned, this value is a unique value for the port, and greater than any
    // ifIndex value supported by the repeater; in this case, the specific port is
    // indicated by corresponding values of cdpInterfaceGroup and
    // cdpInterfacePort, where these values correspond to the group number and
    // port number values of RFC 1516. The type is interface{} with range:
    // 0..2147483647.
    Cdpinterfaceifindex interface{}

    // An indication of whether the Cisco Discovery Protocol is currently running
    // on this interface.  This variable has no effect when CDP is disabled
    // (cdpGlobalRun = FALSE). The type is bool.
    Cdpinterfaceenable interface{}

    // The interval at which CDP messages are to be generated on this interface. 
    // The default value is 60 seconds. The type is interface{} with range:
    // 0..254. Units are seconds.
    Cdpinterfacemessageinterval interface{}

    // This object is only relevant to interfaces which are repeater ports on
    // 802.3 repeaters.  In this situation, it indicates the RFC1516 group number
    // of the repeater port which corresponds to this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    Cdpinterfacegroup interface{}

    // This object is only relevant to interfaces which are repeater ports on
    // 802.3 repeaters.  In this situation, it indicates the RFC1516 port number
    // of the repeater port which corresponds to this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    Cdpinterfaceport interface{}

    // The name of the local interface as advertised by CDP in the Port-ID TLV.
    // The type is string.
    Cdpinterfacename interface{}
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetFilter() yfilter.YFilter { return cdpinterfaceentry.YFilter }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) SetFilter(yf yfilter.YFilter) { cdpinterfaceentry.YFilter = yf }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetGoName(yname string) string {
    if yname == "cdpInterfaceIfIndex" { return "Cdpinterfaceifindex" }
    if yname == "cdpInterfaceEnable" { return "Cdpinterfaceenable" }
    if yname == "cdpInterfaceMessageInterval" { return "Cdpinterfacemessageinterval" }
    if yname == "cdpInterfaceGroup" { return "Cdpinterfacegroup" }
    if yname == "cdpInterfacePort" { return "Cdpinterfaceport" }
    if yname == "cdpInterfaceName" { return "Cdpinterfacename" }
    return ""
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetSegmentPath() string {
    return "cdpInterfaceEntry" + "[cdpInterfaceIfIndex='" + fmt.Sprintf("%v", cdpinterfaceentry.Cdpinterfaceifindex) + "']"
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdpInterfaceIfIndex"] = cdpinterfaceentry.Cdpinterfaceifindex
    leafs["cdpInterfaceEnable"] = cdpinterfaceentry.Cdpinterfaceenable
    leafs["cdpInterfaceMessageInterval"] = cdpinterfaceentry.Cdpinterfacemessageinterval
    leafs["cdpInterfaceGroup"] = cdpinterfaceentry.Cdpinterfacegroup
    leafs["cdpInterfacePort"] = cdpinterfaceentry.Cdpinterfaceport
    leafs["cdpInterfaceName"] = cdpinterfaceentry.Cdpinterfacename
    return leafs
}

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetYangName() string { return "cdpInterfaceEntry" }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) SetParent(parent types.Entity) { cdpinterfaceentry.parent = parent }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetParent() types.Entity { return cdpinterfaceentry.parent }

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetParentYangName() string { return "cdpInterfaceTable" }

// CISCOCDPMIB_Cdpinterfaceexttable
// This table contains the additional CDP configuration on
// the device's interfaces.
type CISCOCDPMIB_Cdpinterfaceexttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the cdpInterfaceExtTable contains the values configured for
    // Extented Trust TLV and COS (Class of Service) for Untrusted Ports TLV on an
    // interface which supports the sending of these TLVs. The type is slice of
    // CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry.
    Cdpinterfaceextentry []CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetFilter() yfilter.YFilter { return cdpinterfaceexttable.YFilter }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) SetFilter(yf yfilter.YFilter) { cdpinterfaceexttable.YFilter = yf }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetGoName(yname string) string {
    if yname == "cdpInterfaceExtEntry" { return "Cdpinterfaceextentry" }
    return ""
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetSegmentPath() string {
    return "cdpInterfaceExtTable"
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdpInterfaceExtEntry" {
        for _, c := range cdpinterfaceexttable.Cdpinterfaceextentry {
            if cdpinterfaceexttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry{}
        cdpinterfaceexttable.Cdpinterfaceextentry = append(cdpinterfaceexttable.Cdpinterfaceextentry, child)
        return &cdpinterfaceexttable.Cdpinterfaceextentry[len(cdpinterfaceexttable.Cdpinterfaceextentry)-1]
    }
    return nil
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdpinterfaceexttable.Cdpinterfaceextentry {
        children[cdpinterfaceexttable.Cdpinterfaceextentry[i].GetSegmentPath()] = &cdpinterfaceexttable.Cdpinterfaceextentry[i]
    }
    return children
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetBundleName() string { return "cisco_ios_xe" }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetYangName() string { return "cdpInterfaceExtTable" }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) SetParent(parent types.Entity) { cdpinterfaceexttable.parent = parent }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetParent() types.Entity { return cdpinterfaceexttable.parent }

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry
// An entry in the cdpInterfaceExtTable contains the values
// configured for Extented Trust TLV and COS (Class of Service)
// for Untrusted Ports TLV on an interface which supports the
// sending of these TLVs.
type CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // Indicates the value to be sent by Extended Trust TLV.  If trusted(1) is
    // configured, the value of Extended Trust TLV is one byte in length with its
    // least significant bit equal to 1 to indicate extended trust. All other bits
    // are 0.  If noTrust(2) is configured, the value of Extended Trust TLV is one
    // byte in length with its least significant bit equal to 0 to indicate no
    // extended trust. All other bits are 0. The type is
    // Cdpinterfaceextendedtrust.
    Cdpinterfaceextendedtrust interface{}

    // Indicates the value to be sent by COS for Untrusted Ports TLV. The type is
    // interface{} with range: 0..7.
    Cdpinterfacecosforuntrustedport interface{}
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetFilter() yfilter.YFilter { return cdpinterfaceextentry.YFilter }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) SetFilter(yf yfilter.YFilter) { cdpinterfaceextentry.YFilter = yf }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cdpInterfaceExtendedTrust" { return "Cdpinterfaceextendedtrust" }
    if yname == "cdpInterfaceCosForUntrustedPort" { return "Cdpinterfacecosforuntrustedport" }
    return ""
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetSegmentPath() string {
    return "cdpInterfaceExtEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdpinterfaceextentry.Ifindex) + "']"
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cdpinterfaceextentry.Ifindex
    leafs["cdpInterfaceExtendedTrust"] = cdpinterfaceextentry.Cdpinterfaceextendedtrust
    leafs["cdpInterfaceCosForUntrustedPort"] = cdpinterfaceextentry.Cdpinterfacecosforuntrustedport
    return leafs
}

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetYangName() string { return "cdpInterfaceExtEntry" }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) SetParent(parent types.Entity) { cdpinterfaceextentry.parent = parent }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetParent() types.Entity { return cdpinterfaceextentry.parent }

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetParentYangName() string { return "cdpInterfaceExtTable" }

// CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust represents 0 to indicate no extended trust. All other bits are 0.
type CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust string

const (
    CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust_trusted CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust = "trusted"

    CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust_noTrust CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry_Cdpinterfaceextendedtrust = "noTrust"
)

// CISCOCDPMIB_Cdpcachetable
// The (conceptual) table containing the cached
// information obtained via receiving CDP messages.
type CISCOCDPMIB_Cdpcachetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpCacheTable, containing the information
    // received via CDP on one interface from one device.  Entries appear when a
    // CDP advertisement is received from a neighbor device.  Entries disappear
    // when CDP is disabled on the interface, or globally. The type is slice of
    // CISCOCDPMIB_Cdpcachetable_Cdpcacheentry.
    Cdpcacheentry []CISCOCDPMIB_Cdpcachetable_Cdpcacheentry
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetFilter() yfilter.YFilter { return cdpcachetable.YFilter }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) SetFilter(yf yfilter.YFilter) { cdpcachetable.YFilter = yf }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetGoName(yname string) string {
    if yname == "cdpCacheEntry" { return "Cdpcacheentry" }
    return ""
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetSegmentPath() string {
    return "cdpCacheTable"
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdpCacheEntry" {
        for _, c := range cdpcachetable.Cdpcacheentry {
            if cdpcachetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCDPMIB_Cdpcachetable_Cdpcacheentry{}
        cdpcachetable.Cdpcacheentry = append(cdpcachetable.Cdpcacheentry, child)
        return &cdpcachetable.Cdpcacheentry[len(cdpcachetable.Cdpcacheentry)-1]
    }
    return nil
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdpcachetable.Cdpcacheentry {
        children[cdpcachetable.Cdpcacheentry[i].GetSegmentPath()] = &cdpcachetable.Cdpcacheentry[i]
    }
    return children
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetYangName() string { return "cdpCacheTable" }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) SetParent(parent types.Entity) { cdpcachetable.parent = parent }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetParent() types.Entity { return cdpcachetable.parent }

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpcachetable_Cdpcacheentry
// An entry (conceptual row) in the cdpCacheTable,
// containing the information received via CDP on one
// interface from one device.  Entries appear when
// a CDP advertisement is received from a neighbor
// device.  Entries disappear when CDP is disabled
// on the interface, or globally.
type CISCOCDPMIB_Cdpcachetable_Cdpcacheentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Normally, the ifIndex value of the local
    // interface. For 802.3 Repeaters for which the repeater ports do not have
    // ifIndex values assigned, this value is a unique value for the port, and
    // greater than any ifIndex value supported by the repeater; the specific port
    // number in this case, is given by the corresponding value of
    // cdpInterfacePort. The type is interface{} with range: 0..2147483647.
    Cdpcacheifindex interface{}

    // This attribute is a key. A unique value for each device from which CDP
    // messages are being received. The type is interface{} with range:
    // 0..2147483647.
    Cdpcachedeviceindex interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCacheAddress. The type is CiscoNetworkProtocol.
    Cdpcacheaddresstype interface{}

    // The (first) network-layer address of the device as reported in the Address
    // TLV of the most recently received CDP message.  For example, if the
    // corresponding instance of cacheAddressType had the value 'ip(1)', then this
    // object  would be an IPv4-address.  If the neighbor device is 
    // SNMP-manageable, it is supposed to generate its CDP messages such that this
    // address is one at which it will receive SNMP messages. Use
    // cdpCtAddressTable to extract the remaining addresses from the Address TLV
    // received most recently. The type is string.
    Cdpcacheaddress interface{}

    // The Version string as reported in the most recent CDP message.  The
    // zero-length string indicates no Version field (TLV) was reported in the
    // most recent CDP message. The type is string.
    Cdpcacheversion interface{}

    // The Device-ID string as reported in the most recent CDP message.  The
    // zero-length string indicates no Device-ID field (TLV) was reported in the
    // most recent CDP message. The type is string.
    Cdpcachedeviceid interface{}

    // The Port-ID string as reported in the most recent CDP message.  This will
    // typically be the value of the ifName object (e.g., 'Ethernet0').  The
    // zero-length string indicates no Port-ID field (TLV) was reported in the
    // most recent CDP message. The type is string.
    Cdpcachedeviceport interface{}

    // The Device's Hardware Platform as reported in the most recent CDP message. 
    // The zero-length string indicates that no Platform field (TLV) was reported
    // in the most recent CDP message. The type is string.
    Cdpcacheplatform interface{}

    // The Device's Functional Capabilities as reported in the most recent CDP
    // message.  For latest set of specific values, see the latest version of the
    // CDP specification. The zero-length string indicates no Capabilities field
    // (TLV) was reported in the most recent CDP message. The type is string with
    // length: 0..4.
    Cdpcachecapabilities interface{}

    // The VTP Management Domain for the remote device's interface,  as reported
    // in the most recently received CDP message. This object is not instantiated
    // if no VTP Management Domain field (TLV) was reported in the most recently
    // received CDP message. The type is string with length: 0..32.
    Cdpcachevtpmgmtdomain interface{}

    // The remote device's interface's native VLAN, as reported in the  most
    // recent CDP message.  The value 0 indicates no native VLAN field (TLV) was
    // reported in the most recent CDP message. The type is interface{} with
    // range: 0..4095.
    Cdpcachenativevlan interface{}

    // The remote device's interface's duplex mode, as reported in the  most
    // recent CDP message.  The value unknown(1) indicates no duplex mode field
    // (TLV) was reported in the most recent CDP message. The type is
    // Cdpcacheduplex.
    Cdpcacheduplex interface{}

    // The remote device's Appliance ID, as reported in the  most recent CDP
    // message. This object is not instantiated if no Appliance VLAN-ID field
    // (TLV) was reported in the most recently received CDP message. The type is
    // interface{} with range: 0..255.
    Cdpcacheapplianceid interface{}

    // The remote device's VoIP VLAN ID, as reported in the  most recent CDP
    // message. This object is not instantiated if no Appliance VLAN-ID field
    // (TLV) was reported in the most recently received CDP message. The type is
    // interface{} with range: 0..4095.
    Cdpcachevlanid interface{}

    // The amount of power consumed by remote device, as reported in the most
    // recent CDP message. This object is not instantiated if no Power Consumption
    // field (TLV) was reported in the most recently received CDP message. The
    // type is interface{} with range: 0..4294967295. Units are milliwatts.
    Cdpcachepowerconsumption interface{}

    // Indicates the size of the largest datagram that can be sent/received by
    // remote device, as reported in the most recent CDP message. This object is
    // not instantiated if no MTU field (TLV) was reported in the most recently
    // received CDP message. The type is interface{} with range: 0..4294967295.
    Cdpcachemtu interface{}

    // Indicates the value of the remote device's sysName MIB object. By
    // convention, it is the device's fully qualified domain name. This object is
    // not instantiated if no sysName field (TLV) was reported in the most
    // recently received CDP message. The type is string with length: 0..255.
    Cdpcachesysname interface{}

    // Indicates the value of the remote device's sysObjectID MIB object. This
    // object is not instantiated if no sysObjectID field (TLV) was reported in
    // the most recently received CDP message. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Cdpcachesysobjectid interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCachePrimaryMgmtAddress. The type is CiscoNetworkProtocol.
    Cdpcacheprimarymgmtaddrtype interface{}

    // This object indicates the (first) network layer address at which the device
    // will accept SNMP messages as reported in the first address in the 
    // Management-Address TLV of the most recently received CDP message.  If the
    // corresponding instance of  cdpCachePrimaryMgmtAddrType has the value
    // 'ip(1)', then this object would be an IP-address. If the  remote device is
    // not currently manageable via any  network protocol, then it reports the
    // special value  of the IPv4 address 0.0.0.0, and that address is  recorded
    // in this object.  If the most recently received CDP message did not contain
    // the Management-Address TLV, then this object is not instanstiated. The type
    // is string.
    Cdpcacheprimarymgmtaddr interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCacheSecondaryMgmtAddress. The type is CiscoNetworkProtocol.
    Cdpcachesecondarymgmtaddrtype interface{}

    // This object indicates the alternate network layer address at which the
    // device will accept SNMP messages as reported in the second address in the 
    // Management-Address TLV of the most recently received CDP message.  If the
    // corresponding instance of cdpCacheSecondaryMgmtAddrType has the value
    // 'ip(1)', then this object would be an IP-address. If the  remote device
    // reports the special value of the  IPv4 address 0.0.0.0, that address is
    // recorded in  this object.  If the most recently received CDP message did
    // not contain the Management-Address TLV, or if that TLV contained only one
    // address, then this object is not instanstiated. The type is string.
    Cdpcachesecondarymgmtaddr interface{}

    // Indicates the physical location, as reported by the most recent CDP
    // message, of a connector which is on, or physically connected to, the remote
    // device's interface over which the CDP packet is sent. This object is not
    // instantiated if no Physical Location field (TLV) was reported by the most
    // recently received CDP message. The type is string.
    Cdpcachephyslocation interface{}

    // Indicates the time when this cache entry was last changed. This object is
    // initialised to the current time when the entry gets created and updated to
    // the current time whenever the value of any (other) object instance in the
    // corresponding row is modified. The type is interface{} with range:
    // 0..4294967295.
    Cdpcachelastchange interface{}
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetFilter() yfilter.YFilter { return cdpcacheentry.YFilter }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) SetFilter(yf yfilter.YFilter) { cdpcacheentry.YFilter = yf }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetGoName(yname string) string {
    if yname == "cdpCacheIfIndex" { return "Cdpcacheifindex" }
    if yname == "cdpCacheDeviceIndex" { return "Cdpcachedeviceindex" }
    if yname == "cdpCacheAddressType" { return "Cdpcacheaddresstype" }
    if yname == "cdpCacheAddress" { return "Cdpcacheaddress" }
    if yname == "cdpCacheVersion" { return "Cdpcacheversion" }
    if yname == "cdpCacheDeviceId" { return "Cdpcachedeviceid" }
    if yname == "cdpCacheDevicePort" { return "Cdpcachedeviceport" }
    if yname == "cdpCachePlatform" { return "Cdpcacheplatform" }
    if yname == "cdpCacheCapabilities" { return "Cdpcachecapabilities" }
    if yname == "cdpCacheVTPMgmtDomain" { return "Cdpcachevtpmgmtdomain" }
    if yname == "cdpCacheNativeVLAN" { return "Cdpcachenativevlan" }
    if yname == "cdpCacheDuplex" { return "Cdpcacheduplex" }
    if yname == "cdpCacheApplianceID" { return "Cdpcacheapplianceid" }
    if yname == "cdpCacheVlanID" { return "Cdpcachevlanid" }
    if yname == "cdpCachePowerConsumption" { return "Cdpcachepowerconsumption" }
    if yname == "cdpCacheMTU" { return "Cdpcachemtu" }
    if yname == "cdpCacheSysName" { return "Cdpcachesysname" }
    if yname == "cdpCacheSysObjectID" { return "Cdpcachesysobjectid" }
    if yname == "cdpCachePrimaryMgmtAddrType" { return "Cdpcacheprimarymgmtaddrtype" }
    if yname == "cdpCachePrimaryMgmtAddr" { return "Cdpcacheprimarymgmtaddr" }
    if yname == "cdpCacheSecondaryMgmtAddrType" { return "Cdpcachesecondarymgmtaddrtype" }
    if yname == "cdpCacheSecondaryMgmtAddr" { return "Cdpcachesecondarymgmtaddr" }
    if yname == "cdpCachePhysLocation" { return "Cdpcachephyslocation" }
    if yname == "cdpCacheLastChange" { return "Cdpcachelastchange" }
    return ""
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetSegmentPath() string {
    return "cdpCacheEntry" + "[cdpCacheIfIndex='" + fmt.Sprintf("%v", cdpcacheentry.Cdpcacheifindex) + "']" + "[cdpCacheDeviceIndex='" + fmt.Sprintf("%v", cdpcacheentry.Cdpcachedeviceindex) + "']"
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdpCacheIfIndex"] = cdpcacheentry.Cdpcacheifindex
    leafs["cdpCacheDeviceIndex"] = cdpcacheentry.Cdpcachedeviceindex
    leafs["cdpCacheAddressType"] = cdpcacheentry.Cdpcacheaddresstype
    leafs["cdpCacheAddress"] = cdpcacheentry.Cdpcacheaddress
    leafs["cdpCacheVersion"] = cdpcacheentry.Cdpcacheversion
    leafs["cdpCacheDeviceId"] = cdpcacheentry.Cdpcachedeviceid
    leafs["cdpCacheDevicePort"] = cdpcacheentry.Cdpcachedeviceport
    leafs["cdpCachePlatform"] = cdpcacheentry.Cdpcacheplatform
    leafs["cdpCacheCapabilities"] = cdpcacheentry.Cdpcachecapabilities
    leafs["cdpCacheVTPMgmtDomain"] = cdpcacheentry.Cdpcachevtpmgmtdomain
    leafs["cdpCacheNativeVLAN"] = cdpcacheentry.Cdpcachenativevlan
    leafs["cdpCacheDuplex"] = cdpcacheentry.Cdpcacheduplex
    leafs["cdpCacheApplianceID"] = cdpcacheentry.Cdpcacheapplianceid
    leafs["cdpCacheVlanID"] = cdpcacheentry.Cdpcachevlanid
    leafs["cdpCachePowerConsumption"] = cdpcacheentry.Cdpcachepowerconsumption
    leafs["cdpCacheMTU"] = cdpcacheentry.Cdpcachemtu
    leafs["cdpCacheSysName"] = cdpcacheentry.Cdpcachesysname
    leafs["cdpCacheSysObjectID"] = cdpcacheentry.Cdpcachesysobjectid
    leafs["cdpCachePrimaryMgmtAddrType"] = cdpcacheentry.Cdpcacheprimarymgmtaddrtype
    leafs["cdpCachePrimaryMgmtAddr"] = cdpcacheentry.Cdpcacheprimarymgmtaddr
    leafs["cdpCacheSecondaryMgmtAddrType"] = cdpcacheentry.Cdpcachesecondarymgmtaddrtype
    leafs["cdpCacheSecondaryMgmtAddr"] = cdpcacheentry.Cdpcachesecondarymgmtaddr
    leafs["cdpCachePhysLocation"] = cdpcacheentry.Cdpcachephyslocation
    leafs["cdpCacheLastChange"] = cdpcacheentry.Cdpcachelastchange
    return leafs
}

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetYangName() string { return "cdpCacheEntry" }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) SetParent(parent types.Entity) { cdpcacheentry.parent = parent }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetParent() types.Entity { return cdpcacheentry.parent }

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetParentYangName() string { return "cdpCacheTable" }

// CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex represents recent CDP message.
type CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex string

const (
    CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex_unknown CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex = "unknown"

    CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex_halfduplex CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex = "halfduplex"

    CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex_fullduplex CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheduplex = "fullduplex"
)

// CISCOCDPMIB_Cdpctaddresstable
// The (conceptual) table containing the list of 
// network-layer addresses of a neighbor interface,
// as reported in the Address TLV of the most recently
// received CDP message. The first address included in
// the Address TLV is saved in cdpCacheAddress.  This
// table contains the remainder of the addresses in the
// Address TLV.
type CISCOCDPMIB_Cdpctaddresstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpCtAddressTable, containing the
    // information on one address received via CDP  on one interface from one
    // device.  Entries appear  when a CDP advertisement is received from a
    // neighbor device, with an Address TLV.  Entries disappear when CDP is
    // disabled on the interface, or globally. An entry  or entries would also
    // disappear if the most recently received CDP packet contain fewer address
    // entries in the Address TLV, than are currently present in the CDP cache.
    // The type is slice of CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry.
    Cdpctaddressentry []CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetFilter() yfilter.YFilter { return cdpctaddresstable.YFilter }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) SetFilter(yf yfilter.YFilter) { cdpctaddresstable.YFilter = yf }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetGoName(yname string) string {
    if yname == "cdpCtAddressEntry" { return "Cdpctaddressentry" }
    return ""
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetSegmentPath() string {
    return "cdpCtAddressTable"
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdpCtAddressEntry" {
        for _, c := range cdpctaddresstable.Cdpctaddressentry {
            if cdpctaddresstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry{}
        cdpctaddresstable.Cdpctaddressentry = append(cdpctaddresstable.Cdpctaddressentry, child)
        return &cdpctaddresstable.Cdpctaddressentry[len(cdpctaddresstable.Cdpctaddressentry)-1]
    }
    return nil
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdpctaddresstable.Cdpctaddressentry {
        children[cdpctaddresstable.Cdpctaddressentry[i].GetSegmentPath()] = &cdpctaddresstable.Cdpctaddressentry[i]
    }
    return children
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetBundleName() string { return "cisco_ios_xe" }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetYangName() string { return "cdpCtAddressTable" }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) SetParent(parent types.Entity) { cdpctaddresstable.parent = parent }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetParent() types.Entity { return cdpctaddresstable.parent }

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetParentYangName() string { return "CISCO-CDP-MIB" }

// CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry
// An entry (conceptual row) in the cdpCtAddressTable,
// containing the information on one address received via CDP 
// on one interface from one device.  Entries appear 
// when a CDP advertisement is received from a neighbor
// device, with an Address TLV.  Entries disappear when
// CDP is disabled on the interface, or globally. An entry 
// or entries would also disappear if the most recently
// received CDP packet contain fewer address entries in the
// Address TLV, than are currently present in the CDP cache.
type CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_cdp_mib.CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcacheifindex
    Cdpcacheifindex interface{}

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_cdp_mib.CISCOCDPMIB_Cdpcachetable_Cdpcacheentry_Cdpcachedeviceindex
    Cdpcachedeviceindex interface{}

    // This attribute is a key. The index of the address entry for a given 
    // cdpCacheIfIndex,cdpCacheDeviceIndex pair. It has the value N-1 for the N-th
    // address in the Address TLV. The type is interface{} with range:
    // 1..2147483647.
    Cdpctaddressindex interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCtAddress. The type is CiscoNetworkProtocol.
    Cdpctaddresstype interface{}

    // The N-th network-layer address of the device as reported in the most recent
    // CDP message's Address TLV, where N-1 is given by the value of
    // cdpCtAddressIndex. For example, if the the corresponding instance of
    // cdpCtAddressType had the value 'ip(1)', then this object would be an
    // IPv4-address. NOTE - The 1st address received in the Address TLV is       
    // available using cdpCacheAddress. The type is string.
    Cdpctaddress interface{}
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetFilter() yfilter.YFilter { return cdpctaddressentry.YFilter }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) SetFilter(yf yfilter.YFilter) { cdpctaddressentry.YFilter = yf }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetGoName(yname string) string {
    if yname == "cdpCacheIfIndex" { return "Cdpcacheifindex" }
    if yname == "cdpCacheDeviceIndex" { return "Cdpcachedeviceindex" }
    if yname == "cdpCtAddressIndex" { return "Cdpctaddressindex" }
    if yname == "cdpCtAddressType" { return "Cdpctaddresstype" }
    if yname == "cdpCtAddress" { return "Cdpctaddress" }
    return ""
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetSegmentPath() string {
    return "cdpCtAddressEntry" + "[cdpCacheIfIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpcacheifindex) + "']" + "[cdpCacheDeviceIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpcachedeviceindex) + "']" + "[cdpCtAddressIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpctaddressindex) + "']"
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdpCacheIfIndex"] = cdpctaddressentry.Cdpcacheifindex
    leafs["cdpCacheDeviceIndex"] = cdpctaddressentry.Cdpcachedeviceindex
    leafs["cdpCtAddressIndex"] = cdpctaddressentry.Cdpctaddressindex
    leafs["cdpCtAddressType"] = cdpctaddressentry.Cdpctaddresstype
    leafs["cdpCtAddress"] = cdpctaddressentry.Cdpctaddress
    return leafs
}

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetYangName() string { return "cdpCtAddressEntry" }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) SetParent(parent types.Entity) { cdpctaddressentry.parent = parent }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetParent() types.Entity { return cdpctaddressentry.parent }

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetParentYangName() string { return "cdpCtAddressTable" }

