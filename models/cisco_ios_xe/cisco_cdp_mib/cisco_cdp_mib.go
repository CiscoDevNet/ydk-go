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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CdpGlobal CISCOCDPMIB_CdpGlobal

    // The (conceptual) table containing the status of CDP on the device's
    // interfaces.
    CdpInterfaceTable CISCOCDPMIB_CdpInterfaceTable

    // This table contains the additional CDP configuration on the device's
    // interfaces.
    CdpInterfaceExtTable CISCOCDPMIB_CdpInterfaceExtTable

    // The (conceptual) table containing the cached information obtained via
    // receiving CDP messages.
    CdpCacheTable CISCOCDPMIB_CdpCacheTable

    // The (conceptual) table containing the list of  network-layer addresses of a
    // neighbor interface, as reported in the Address TLV of the most recently
    // received CDP message. The first address included in the Address TLV is
    // saved in cdpCacheAddress.  This table contains the remainder of the
    // addresses in the Address TLV.
    CdpCtAddressTable CISCOCDPMIB_CdpCtAddressTable
}

func (cISCOCDPMIB *CISCOCDPMIB) GetEntityData() *types.CommonEntityData {
    cISCOCDPMIB.EntityData.YFilter = cISCOCDPMIB.YFilter
    cISCOCDPMIB.EntityData.YangName = "CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCDPMIB.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.SegmentPath = "CISCO-CDP-MIB:CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCDPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCDPMIB.EntityData.Children.Append("cdpGlobal", types.YChild{"CdpGlobal", &cISCOCDPMIB.CdpGlobal})
    cISCOCDPMIB.EntityData.Children.Append("cdpInterfaceTable", types.YChild{"CdpInterfaceTable", &cISCOCDPMIB.CdpInterfaceTable})
    cISCOCDPMIB.EntityData.Children.Append("cdpInterfaceExtTable", types.YChild{"CdpInterfaceExtTable", &cISCOCDPMIB.CdpInterfaceExtTable})
    cISCOCDPMIB.EntityData.Children.Append("cdpCacheTable", types.YChild{"CdpCacheTable", &cISCOCDPMIB.CdpCacheTable})
    cISCOCDPMIB.EntityData.Children.Append("cdpCtAddressTable", types.YChild{"CdpCtAddressTable", &cISCOCDPMIB.CdpCtAddressTable})
    cISCOCDPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCDPMIB.EntityData.YListKeys = []string {}

    return &(cISCOCDPMIB.EntityData)
}

// CISCOCDPMIB_CdpGlobal
type CISCOCDPMIB_CdpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the Cisco Discovery Protocol is currently running.
    // Entries in cdpCacheTable are deleted when CDP is disabled. The type is
    // bool.
    CdpGlobalRun interface{}

    // The interval at which CDP messages are to be generated. The default value
    // is 60 seconds. The type is interface{} with range: 5..254. Units are
    // seconds.
    CdpGlobalMessageInterval interface{}

    // The time for the receiving device holds CDP message. The default value is
    // 180 seconds. The type is interface{} with range: 10..255. Units are
    // seconds.
    CdpGlobalHoldTime interface{}

    // The device ID advertised by this device. The format of this device id is
    // characterized by the value of  cdpGlobalDeviceIdFormat object. The type is
    // string.
    CdpGlobalDeviceId interface{}

    // Indicates the time when the cache table was last changed. It is the most
    // recent time at which any row was last created, modified or deleted. The
    // type is interface{} with range: 0..4294967295.
    CdpGlobalLastChange interface{}

    // Indicate the Device-Id format capability of the device.  serialNumber(0)
    // indicates that the device supports using serial number as the format for
    // its DeviceId.  macAddress(1) indicates that the device supports using layer
    // 2 MAC address as the format for its DeviceId.  other(2) indicates that the
    // device supports using its platform specific format as the format for its
    // DeviceId. The type is map[string]bool.
    CdpGlobalDeviceIdFormatCpb interface{}

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
    // CdpGlobalDeviceIdFormat.
    CdpGlobalDeviceIdFormat interface{}
}

func (cdpGlobal *CISCOCDPMIB_CdpGlobal) GetEntityData() *types.CommonEntityData {
    cdpGlobal.EntityData.YFilter = cdpGlobal.YFilter
    cdpGlobal.EntityData.YangName = "cdpGlobal"
    cdpGlobal.EntityData.BundleName = "cisco_ios_xe"
    cdpGlobal.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpGlobal.EntityData.SegmentPath = "cdpGlobal"
    cdpGlobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpGlobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpGlobal.EntityData.Children = types.NewOrderedMap()
    cdpGlobal.EntityData.Leafs = types.NewOrderedMap()
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalRun", types.YLeaf{"CdpGlobalRun", cdpGlobal.CdpGlobalRun})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalMessageInterval", types.YLeaf{"CdpGlobalMessageInterval", cdpGlobal.CdpGlobalMessageInterval})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalHoldTime", types.YLeaf{"CdpGlobalHoldTime", cdpGlobal.CdpGlobalHoldTime})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalDeviceId", types.YLeaf{"CdpGlobalDeviceId", cdpGlobal.CdpGlobalDeviceId})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalLastChange", types.YLeaf{"CdpGlobalLastChange", cdpGlobal.CdpGlobalLastChange})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalDeviceIdFormatCpb", types.YLeaf{"CdpGlobalDeviceIdFormatCpb", cdpGlobal.CdpGlobalDeviceIdFormatCpb})
    cdpGlobal.EntityData.Leafs.Append("cdpGlobalDeviceIdFormat", types.YLeaf{"CdpGlobalDeviceIdFormat", cdpGlobal.CdpGlobalDeviceIdFormat})

    cdpGlobal.EntityData.YListKeys = []string {}

    return &(cdpGlobal.EntityData)
}

// CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat represents contains serialNumber appended/prepened with system name.
type CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat string

const (
    CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat_serialNumber CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat = "serialNumber"

    CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat_macAddress CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat = "macAddress"

    CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat_other CISCOCDPMIB_CdpGlobal_CdpGlobalDeviceIdFormat = "other"
)

// CISCOCDPMIB_CdpInterfaceTable
// The (conceptual) table containing the status of CDP on
// the device's interfaces.
type CISCOCDPMIB_CdpInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpInterfaceTable, containing the status
    // of CDP on an interface. The type is slice of
    // CISCOCDPMIB_CdpInterfaceTable_CdpInterfaceEntry.
    CdpInterfaceEntry []*CISCOCDPMIB_CdpInterfaceTable_CdpInterfaceEntry
}

func (cdpInterfaceTable *CISCOCDPMIB_CdpInterfaceTable) GetEntityData() *types.CommonEntityData {
    cdpInterfaceTable.EntityData.YFilter = cdpInterfaceTable.YFilter
    cdpInterfaceTable.EntityData.YangName = "cdpInterfaceTable"
    cdpInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    cdpInterfaceTable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpInterfaceTable.EntityData.SegmentPath = "cdpInterfaceTable"
    cdpInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpInterfaceTable.EntityData.Children = types.NewOrderedMap()
    cdpInterfaceTable.EntityData.Children.Append("cdpInterfaceEntry", types.YChild{"CdpInterfaceEntry", nil})
    for i := range cdpInterfaceTable.CdpInterfaceEntry {
        cdpInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(cdpInterfaceTable.CdpInterfaceEntry[i]), types.YChild{"CdpInterfaceEntry", cdpInterfaceTable.CdpInterfaceEntry[i]})
    }
    cdpInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    cdpInterfaceTable.EntityData.YListKeys = []string {}

    return &(cdpInterfaceTable.EntityData)
}

// CISCOCDPMIB_CdpInterfaceTable_CdpInterfaceEntry
// An entry (conceptual row) in the cdpInterfaceTable,
// containing the status of CDP on an interface.
type CISCOCDPMIB_CdpInterfaceTable_CdpInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the local interface.  For
    // 802.3 Repeaters on which the repeater ports do not have ifIndex values
    // assigned, this value is a unique value for the port, and greater than any
    // ifIndex value supported by the repeater; in this case, the specific port is
    // indicated by corresponding values of cdpInterfaceGroup and
    // cdpInterfacePort, where these values correspond to the group number and
    // port number values of RFC 1516. The type is interface{} with range:
    // 0..2147483647.
    CdpInterfaceIfIndex interface{}

    // An indication of whether the Cisco Discovery Protocol is currently running
    // on this interface.  This variable has no effect when CDP is disabled
    // (cdpGlobalRun = FALSE). The type is bool.
    CdpInterfaceEnable interface{}

    // The interval at which CDP messages are to be generated on this interface. 
    // The default value is 60 seconds. The type is interface{} with range:
    // 0..254. Units are seconds.
    CdpInterfaceMessageInterval interface{}

    // This object is only relevant to interfaces which are repeater ports on
    // 802.3 repeaters.  In this situation, it indicates the RFC1516 group number
    // of the repeater port which corresponds to this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    CdpInterfaceGroup interface{}

    // This object is only relevant to interfaces which are repeater ports on
    // 802.3 repeaters.  In this situation, it indicates the RFC1516 port number
    // of the repeater port which corresponds to this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    CdpInterfacePort interface{}

    // The name of the local interface as advertised by CDP in the Port-ID TLV.
    // The type is string.
    CdpInterfaceName interface{}
}

func (cdpInterfaceEntry *CISCOCDPMIB_CdpInterfaceTable_CdpInterfaceEntry) GetEntityData() *types.CommonEntityData {
    cdpInterfaceEntry.EntityData.YFilter = cdpInterfaceEntry.YFilter
    cdpInterfaceEntry.EntityData.YangName = "cdpInterfaceEntry"
    cdpInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    cdpInterfaceEntry.EntityData.ParentYangName = "cdpInterfaceTable"
    cdpInterfaceEntry.EntityData.SegmentPath = "cdpInterfaceEntry" + types.AddKeyToken(cdpInterfaceEntry.CdpInterfaceIfIndex, "cdpInterfaceIfIndex")
    cdpInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    cdpInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfaceIfIndex", types.YLeaf{"CdpInterfaceIfIndex", cdpInterfaceEntry.CdpInterfaceIfIndex})
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfaceEnable", types.YLeaf{"CdpInterfaceEnable", cdpInterfaceEntry.CdpInterfaceEnable})
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfaceMessageInterval", types.YLeaf{"CdpInterfaceMessageInterval", cdpInterfaceEntry.CdpInterfaceMessageInterval})
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfaceGroup", types.YLeaf{"CdpInterfaceGroup", cdpInterfaceEntry.CdpInterfaceGroup})
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfacePort", types.YLeaf{"CdpInterfacePort", cdpInterfaceEntry.CdpInterfacePort})
    cdpInterfaceEntry.EntityData.Leafs.Append("cdpInterfaceName", types.YLeaf{"CdpInterfaceName", cdpInterfaceEntry.CdpInterfaceName})

    cdpInterfaceEntry.EntityData.YListKeys = []string {"CdpInterfaceIfIndex"}

    return &(cdpInterfaceEntry.EntityData)
}

// CISCOCDPMIB_CdpInterfaceExtTable
// This table contains the additional CDP configuration on
// the device's interfaces.
type CISCOCDPMIB_CdpInterfaceExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cdpInterfaceExtTable contains the values configured for
    // Extented Trust TLV and COS (Class of Service) for Untrusted Ports TLV on an
    // interface which supports the sending of these TLVs. The type is slice of
    // CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry.
    CdpInterfaceExtEntry []*CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry
}

func (cdpInterfaceExtTable *CISCOCDPMIB_CdpInterfaceExtTable) GetEntityData() *types.CommonEntityData {
    cdpInterfaceExtTable.EntityData.YFilter = cdpInterfaceExtTable.YFilter
    cdpInterfaceExtTable.EntityData.YangName = "cdpInterfaceExtTable"
    cdpInterfaceExtTable.EntityData.BundleName = "cisco_ios_xe"
    cdpInterfaceExtTable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpInterfaceExtTable.EntityData.SegmentPath = "cdpInterfaceExtTable"
    cdpInterfaceExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpInterfaceExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpInterfaceExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpInterfaceExtTable.EntityData.Children = types.NewOrderedMap()
    cdpInterfaceExtTable.EntityData.Children.Append("cdpInterfaceExtEntry", types.YChild{"CdpInterfaceExtEntry", nil})
    for i := range cdpInterfaceExtTable.CdpInterfaceExtEntry {
        cdpInterfaceExtTable.EntityData.Children.Append(types.GetSegmentPath(cdpInterfaceExtTable.CdpInterfaceExtEntry[i]), types.YChild{"CdpInterfaceExtEntry", cdpInterfaceExtTable.CdpInterfaceExtEntry[i]})
    }
    cdpInterfaceExtTable.EntityData.Leafs = types.NewOrderedMap()

    cdpInterfaceExtTable.EntityData.YListKeys = []string {}

    return &(cdpInterfaceExtTable.EntityData)
}

// CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry
// An entry in the cdpInterfaceExtTable contains the values
// configured for Extented Trust TLV and COS (Class of Service)
// for Untrusted Ports TLV on an interface which supports the
// sending of these TLVs.
type CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Indicates the value to be sent by Extended Trust TLV.  If trusted(1) is
    // configured, the value of Extended Trust TLV is one byte in length with its
    // least significant bit equal to 1 to indicate extended trust. All other bits
    // are 0.  If noTrust(2) is configured, the value of Extended Trust TLV is one
    // byte in length with its least significant bit equal to 0 to indicate no
    // extended trust. All other bits are 0. The type is
    // CdpInterfaceExtendedTrust.
    CdpInterfaceExtendedTrust interface{}

    // Indicates the value to be sent by COS for Untrusted Ports TLV. The type is
    // interface{} with range: 0..7.
    CdpInterfaceCosForUntrustedPort interface{}
}

func (cdpInterfaceExtEntry *CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry) GetEntityData() *types.CommonEntityData {
    cdpInterfaceExtEntry.EntityData.YFilter = cdpInterfaceExtEntry.YFilter
    cdpInterfaceExtEntry.EntityData.YangName = "cdpInterfaceExtEntry"
    cdpInterfaceExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdpInterfaceExtEntry.EntityData.ParentYangName = "cdpInterfaceExtTable"
    cdpInterfaceExtEntry.EntityData.SegmentPath = "cdpInterfaceExtEntry" + types.AddKeyToken(cdpInterfaceExtEntry.IfIndex, "ifIndex")
    cdpInterfaceExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpInterfaceExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpInterfaceExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpInterfaceExtEntry.EntityData.Children = types.NewOrderedMap()
    cdpInterfaceExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpInterfaceExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdpInterfaceExtEntry.IfIndex})
    cdpInterfaceExtEntry.EntityData.Leafs.Append("cdpInterfaceExtendedTrust", types.YLeaf{"CdpInterfaceExtendedTrust", cdpInterfaceExtEntry.CdpInterfaceExtendedTrust})
    cdpInterfaceExtEntry.EntityData.Leafs.Append("cdpInterfaceCosForUntrustedPort", types.YLeaf{"CdpInterfaceCosForUntrustedPort", cdpInterfaceExtEntry.CdpInterfaceCosForUntrustedPort})

    cdpInterfaceExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdpInterfaceExtEntry.EntityData)
}

// CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust represents 0 to indicate no extended trust. All other bits are 0.
type CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust string

const (
    CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust_trusted CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust = "trusted"

    CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust_noTrust CISCOCDPMIB_CdpInterfaceExtTable_CdpInterfaceExtEntry_CdpInterfaceExtendedTrust = "noTrust"
)

// CISCOCDPMIB_CdpCacheTable
// The (conceptual) table containing the cached
// information obtained via receiving CDP messages.
type CISCOCDPMIB_CdpCacheTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpCacheTable, containing the information
    // received via CDP on one interface from one device.  Entries appear when a
    // CDP advertisement is received from a neighbor device.  Entries disappear
    // when CDP is disabled on the interface, or globally. The type is slice of
    // CISCOCDPMIB_CdpCacheTable_CdpCacheEntry.
    CdpCacheEntry []*CISCOCDPMIB_CdpCacheTable_CdpCacheEntry
}

func (cdpCacheTable *CISCOCDPMIB_CdpCacheTable) GetEntityData() *types.CommonEntityData {
    cdpCacheTable.EntityData.YFilter = cdpCacheTable.YFilter
    cdpCacheTable.EntityData.YangName = "cdpCacheTable"
    cdpCacheTable.EntityData.BundleName = "cisco_ios_xe"
    cdpCacheTable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpCacheTable.EntityData.SegmentPath = "cdpCacheTable"
    cdpCacheTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpCacheTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpCacheTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpCacheTable.EntityData.Children = types.NewOrderedMap()
    cdpCacheTable.EntityData.Children.Append("cdpCacheEntry", types.YChild{"CdpCacheEntry", nil})
    for i := range cdpCacheTable.CdpCacheEntry {
        cdpCacheTable.EntityData.Children.Append(types.GetSegmentPath(cdpCacheTable.CdpCacheEntry[i]), types.YChild{"CdpCacheEntry", cdpCacheTable.CdpCacheEntry[i]})
    }
    cdpCacheTable.EntityData.Leafs = types.NewOrderedMap()

    cdpCacheTable.EntityData.YListKeys = []string {}

    return &(cdpCacheTable.EntityData)
}

// CISCOCDPMIB_CdpCacheTable_CdpCacheEntry
// An entry (conceptual row) in the cdpCacheTable,
// containing the information received via CDP on one
// interface from one device.  Entries appear when
// a CDP advertisement is received from a neighbor
// device.  Entries disappear when CDP is disabled
// on the interface, or globally.
type CISCOCDPMIB_CdpCacheTable_CdpCacheEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Normally, the ifIndex value of the local
    // interface. For 802.3 Repeaters for which the repeater ports do not have
    // ifIndex values assigned, this value is a unique value for the port, and
    // greater than any ifIndex value supported by the repeater; the specific port
    // number in this case, is given by the corresponding value of
    // cdpInterfacePort. The type is interface{} with range: 0..2147483647.
    CdpCacheIfIndex interface{}

    // This attribute is a key. A unique value for each device from which CDP
    // messages are being received. The type is interface{} with range:
    // 0..2147483647.
    CdpCacheDeviceIndex interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCacheAddress. The type is CiscoNetworkProtocol.
    CdpCacheAddressType interface{}

    // The (first) network-layer address of the device as reported in the Address
    // TLV of the most recently received CDP message.  For example, if the
    // corresponding instance of cacheAddressType had the value 'ip(1)', then this
    // object  would be an IPv4-address.  If the neighbor device is 
    // SNMP-manageable, it is supposed to generate its CDP messages such that this
    // address is one at which it will receive SNMP messages. Use
    // cdpCtAddressTable to extract the remaining addresses from the Address TLV
    // received most recently. The type is string.
    CdpCacheAddress interface{}

    // The Version string as reported in the most recent CDP message.  The
    // zero-length string indicates no Version field (TLV) was reported in the
    // most recent CDP message. The type is string.
    CdpCacheVersion interface{}

    // The Device-ID string as reported in the most recent CDP message.  The
    // zero-length string indicates no Device-ID field (TLV) was reported in the
    // most recent CDP message. The type is string.
    CdpCacheDeviceId interface{}

    // The Port-ID string as reported in the most recent CDP message.  This will
    // typically be the value of the ifName object (e.g., 'Ethernet0').  The
    // zero-length string indicates no Port-ID field (TLV) was reported in the
    // most recent CDP message. The type is string.
    CdpCacheDevicePort interface{}

    // The Device's Hardware Platform as reported in the most recent CDP message. 
    // The zero-length string indicates that no Platform field (TLV) was reported
    // in the most recent CDP message. The type is string.
    CdpCachePlatform interface{}

    // The Device's Functional Capabilities as reported in the most recent CDP
    // message.  For latest set of specific values, see the latest version of the
    // CDP specification. The zero-length string indicates no Capabilities field
    // (TLV) was reported in the most recent CDP message. The type is string with
    // length: 0..4.
    CdpCacheCapabilities interface{}

    // The VTP Management Domain for the remote device's interface,  as reported
    // in the most recently received CDP message. This object is not instantiated
    // if no VTP Management Domain field (TLV) was reported in the most recently
    // received CDP message. The type is string with length: 0..32.
    CdpCacheVTPMgmtDomain interface{}

    // The remote device's interface's native VLAN, as reported in the  most
    // recent CDP message.  The value 0 indicates no native VLAN field (TLV) was
    // reported in the most recent CDP message. The type is interface{} with
    // range: 0..4095.
    CdpCacheNativeVLAN interface{}

    // The remote device's interface's duplex mode, as reported in the  most
    // recent CDP message.  The value unknown(1) indicates no duplex mode field
    // (TLV) was reported in the most recent CDP message. The type is
    // CdpCacheDuplex.
    CdpCacheDuplex interface{}

    // The remote device's Appliance ID, as reported in the  most recent CDP
    // message. This object is not instantiated if no Appliance VLAN-ID field
    // (TLV) was reported in the most recently received CDP message. The type is
    // interface{} with range: 0..255.
    CdpCacheApplianceID interface{}

    // The remote device's VoIP VLAN ID, as reported in the  most recent CDP
    // message. This object is not instantiated if no Appliance VLAN-ID field
    // (TLV) was reported in the most recently received CDP message. The type is
    // interface{} with range: 0..4095.
    CdpCacheVlanID interface{}

    // The amount of power consumed by remote device, as reported in the most
    // recent CDP message. This object is not instantiated if no Power Consumption
    // field (TLV) was reported in the most recently received CDP message. The
    // type is interface{} with range: 0..4294967295. Units are milliwatts.
    CdpCachePowerConsumption interface{}

    // Indicates the size of the largest datagram that can be sent/received by
    // remote device, as reported in the most recent CDP message. This object is
    // not instantiated if no MTU field (TLV) was reported in the most recently
    // received CDP message. The type is interface{} with range: 0..4294967295.
    CdpCacheMTU interface{}

    // Indicates the value of the remote device's sysName MIB object. By
    // convention, it is the device's fully qualified domain name. This object is
    // not instantiated if no sysName field (TLV) was reported in the most
    // recently received CDP message. The type is string with length: 0..255.
    CdpCacheSysName interface{}

    // Indicates the value of the remote device's sysObjectID MIB object. This
    // object is not instantiated if no sysObjectID field (TLV) was reported in
    // the most recently received CDP message. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdpCacheSysObjectID interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCachePrimaryMgmtAddress. The type is CiscoNetworkProtocol.
    CdpCachePrimaryMgmtAddrType interface{}

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
    CdpCachePrimaryMgmtAddr interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCacheSecondaryMgmtAddress. The type is CiscoNetworkProtocol.
    CdpCacheSecondaryMgmtAddrType interface{}

    // This object indicates the alternate network layer address at which the
    // device will accept SNMP messages as reported in the second address in the 
    // Management-Address TLV of the most recently received CDP message.  If the
    // corresponding instance of cdpCacheSecondaryMgmtAddrType has the value
    // 'ip(1)', then this object would be an IP-address. If the  remote device
    // reports the special value of the  IPv4 address 0.0.0.0, that address is
    // recorded in  this object.  If the most recently received CDP message did
    // not contain the Management-Address TLV, or if that TLV contained only one
    // address, then this object is not instanstiated. The type is string.
    CdpCacheSecondaryMgmtAddr interface{}

    // Indicates the physical location, as reported by the most recent CDP
    // message, of a connector which is on, or physically connected to, the remote
    // device's interface over which the CDP packet is sent. This object is not
    // instantiated if no Physical Location field (TLV) was reported by the most
    // recently received CDP message. The type is string.
    CdpCachePhysLocation interface{}

    // Indicates the time when this cache entry was last changed. This object is
    // initialised to the current time when the entry gets created and updated to
    // the current time whenever the value of any (other) object instance in the
    // corresponding row is modified. The type is interface{} with range:
    // 0..4294967295.
    CdpCacheLastChange interface{}
}

func (cdpCacheEntry *CISCOCDPMIB_CdpCacheTable_CdpCacheEntry) GetEntityData() *types.CommonEntityData {
    cdpCacheEntry.EntityData.YFilter = cdpCacheEntry.YFilter
    cdpCacheEntry.EntityData.YangName = "cdpCacheEntry"
    cdpCacheEntry.EntityData.BundleName = "cisco_ios_xe"
    cdpCacheEntry.EntityData.ParentYangName = "cdpCacheTable"
    cdpCacheEntry.EntityData.SegmentPath = "cdpCacheEntry" + types.AddKeyToken(cdpCacheEntry.CdpCacheIfIndex, "cdpCacheIfIndex") + types.AddKeyToken(cdpCacheEntry.CdpCacheDeviceIndex, "cdpCacheDeviceIndex")
    cdpCacheEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpCacheEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpCacheEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpCacheEntry.EntityData.Children = types.NewOrderedMap()
    cdpCacheEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheIfIndex", types.YLeaf{"CdpCacheIfIndex", cdpCacheEntry.CdpCacheIfIndex})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheDeviceIndex", types.YLeaf{"CdpCacheDeviceIndex", cdpCacheEntry.CdpCacheDeviceIndex})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheAddressType", types.YLeaf{"CdpCacheAddressType", cdpCacheEntry.CdpCacheAddressType})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheAddress", types.YLeaf{"CdpCacheAddress", cdpCacheEntry.CdpCacheAddress})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheVersion", types.YLeaf{"CdpCacheVersion", cdpCacheEntry.CdpCacheVersion})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheDeviceId", types.YLeaf{"CdpCacheDeviceId", cdpCacheEntry.CdpCacheDeviceId})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheDevicePort", types.YLeaf{"CdpCacheDevicePort", cdpCacheEntry.CdpCacheDevicePort})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCachePlatform", types.YLeaf{"CdpCachePlatform", cdpCacheEntry.CdpCachePlatform})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheCapabilities", types.YLeaf{"CdpCacheCapabilities", cdpCacheEntry.CdpCacheCapabilities})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheVTPMgmtDomain", types.YLeaf{"CdpCacheVTPMgmtDomain", cdpCacheEntry.CdpCacheVTPMgmtDomain})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheNativeVLAN", types.YLeaf{"CdpCacheNativeVLAN", cdpCacheEntry.CdpCacheNativeVLAN})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheDuplex", types.YLeaf{"CdpCacheDuplex", cdpCacheEntry.CdpCacheDuplex})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheApplianceID", types.YLeaf{"CdpCacheApplianceID", cdpCacheEntry.CdpCacheApplianceID})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheVlanID", types.YLeaf{"CdpCacheVlanID", cdpCacheEntry.CdpCacheVlanID})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCachePowerConsumption", types.YLeaf{"CdpCachePowerConsumption", cdpCacheEntry.CdpCachePowerConsumption})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheMTU", types.YLeaf{"CdpCacheMTU", cdpCacheEntry.CdpCacheMTU})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheSysName", types.YLeaf{"CdpCacheSysName", cdpCacheEntry.CdpCacheSysName})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheSysObjectID", types.YLeaf{"CdpCacheSysObjectID", cdpCacheEntry.CdpCacheSysObjectID})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCachePrimaryMgmtAddrType", types.YLeaf{"CdpCachePrimaryMgmtAddrType", cdpCacheEntry.CdpCachePrimaryMgmtAddrType})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCachePrimaryMgmtAddr", types.YLeaf{"CdpCachePrimaryMgmtAddr", cdpCacheEntry.CdpCachePrimaryMgmtAddr})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheSecondaryMgmtAddrType", types.YLeaf{"CdpCacheSecondaryMgmtAddrType", cdpCacheEntry.CdpCacheSecondaryMgmtAddrType})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheSecondaryMgmtAddr", types.YLeaf{"CdpCacheSecondaryMgmtAddr", cdpCacheEntry.CdpCacheSecondaryMgmtAddr})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCachePhysLocation", types.YLeaf{"CdpCachePhysLocation", cdpCacheEntry.CdpCachePhysLocation})
    cdpCacheEntry.EntityData.Leafs.Append("cdpCacheLastChange", types.YLeaf{"CdpCacheLastChange", cdpCacheEntry.CdpCacheLastChange})

    cdpCacheEntry.EntityData.YListKeys = []string {"CdpCacheIfIndex", "CdpCacheDeviceIndex"}

    return &(cdpCacheEntry.EntityData)
}

// CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex represents recent CDP message.
type CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex string

const (
    CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex_unknown CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex = "unknown"

    CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex_halfduplex CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex = "halfduplex"

    CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex_fullduplex CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDuplex = "fullduplex"
)

// CISCOCDPMIB_CdpCtAddressTable
// The (conceptual) table containing the list of 
// network-layer addresses of a neighbor interface,
// as reported in the Address TLV of the most recently
// received CDP message. The first address included in
// the Address TLV is saved in cdpCacheAddress.  This
// table contains the remainder of the addresses in the
// Address TLV.
type CISCOCDPMIB_CdpCtAddressTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpCtAddressTable, containing the
    // information on one address received via CDP  on one interface from one
    // device.  Entries appear  when a CDP advertisement is received from a
    // neighbor device, with an Address TLV.  Entries disappear when CDP is
    // disabled on the interface, or globally. An entry  or entries would also
    // disappear if the most recently received CDP packet contain fewer address
    // entries in the Address TLV, than are currently present in the CDP cache.
    // The type is slice of CISCOCDPMIB_CdpCtAddressTable_CdpCtAddressEntry.
    CdpCtAddressEntry []*CISCOCDPMIB_CdpCtAddressTable_CdpCtAddressEntry
}

func (cdpCtAddressTable *CISCOCDPMIB_CdpCtAddressTable) GetEntityData() *types.CommonEntityData {
    cdpCtAddressTable.EntityData.YFilter = cdpCtAddressTable.YFilter
    cdpCtAddressTable.EntityData.YangName = "cdpCtAddressTable"
    cdpCtAddressTable.EntityData.BundleName = "cisco_ios_xe"
    cdpCtAddressTable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpCtAddressTable.EntityData.SegmentPath = "cdpCtAddressTable"
    cdpCtAddressTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpCtAddressTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpCtAddressTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpCtAddressTable.EntityData.Children = types.NewOrderedMap()
    cdpCtAddressTable.EntityData.Children.Append("cdpCtAddressEntry", types.YChild{"CdpCtAddressEntry", nil})
    for i := range cdpCtAddressTable.CdpCtAddressEntry {
        cdpCtAddressTable.EntityData.Children.Append(types.GetSegmentPath(cdpCtAddressTable.CdpCtAddressEntry[i]), types.YChild{"CdpCtAddressEntry", cdpCtAddressTable.CdpCtAddressEntry[i]})
    }
    cdpCtAddressTable.EntityData.Leafs = types.NewOrderedMap()

    cdpCtAddressTable.EntityData.YListKeys = []string {}

    return &(cdpCtAddressTable.EntityData)
}

// CISCOCDPMIB_CdpCtAddressTable_CdpCtAddressEntry
// An entry (conceptual row) in the cdpCtAddressTable,
// containing the information on one address received via CDP 
// on one interface from one device.  Entries appear 
// when a CDP advertisement is received from a neighbor
// device, with an Address TLV.  Entries disappear when
// CDP is disabled on the interface, or globally. An entry 
// or entries would also disappear if the most recently
// received CDP packet contain fewer address entries in the
// Address TLV, than are currently present in the CDP cache.
type CISCOCDPMIB_CdpCtAddressTable_CdpCtAddressEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_cdp_mib.CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheIfIndex
    CdpCacheIfIndex interface{}

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_cdp_mib.CISCOCDPMIB_CdpCacheTable_CdpCacheEntry_CdpCacheDeviceIndex
    CdpCacheDeviceIndex interface{}

    // This attribute is a key. The index of the address entry for a given 
    // cdpCacheIfIndex,cdpCacheDeviceIndex pair. It has the value N-1 for the N-th
    // address in the Address TLV. The type is interface{} with range:
    // 1..2147483647.
    CdpCtAddressIndex interface{}

    // An indication of the type of address contained in the corresponding
    // instance of cdpCtAddress. The type is CiscoNetworkProtocol.
    CdpCtAddressType interface{}

    // The N-th network-layer address of the device as reported in the most recent
    // CDP message's Address TLV, where N-1 is given by the value of
    // cdpCtAddressIndex. For example, if the the corresponding instance of
    // cdpCtAddressType had the value 'ip(1)', then this object would be an
    // IPv4-address. NOTE - The 1st address received in the Address TLV is       
    // available using cdpCacheAddress. The type is string.
    CdpCtAddress interface{}
}

func (cdpCtAddressEntry *CISCOCDPMIB_CdpCtAddressTable_CdpCtAddressEntry) GetEntityData() *types.CommonEntityData {
    cdpCtAddressEntry.EntityData.YFilter = cdpCtAddressEntry.YFilter
    cdpCtAddressEntry.EntityData.YangName = "cdpCtAddressEntry"
    cdpCtAddressEntry.EntityData.BundleName = "cisco_ios_xe"
    cdpCtAddressEntry.EntityData.ParentYangName = "cdpCtAddressTable"
    cdpCtAddressEntry.EntityData.SegmentPath = "cdpCtAddressEntry" + types.AddKeyToken(cdpCtAddressEntry.CdpCacheIfIndex, "cdpCacheIfIndex") + types.AddKeyToken(cdpCtAddressEntry.CdpCacheDeviceIndex, "cdpCacheDeviceIndex") + types.AddKeyToken(cdpCtAddressEntry.CdpCtAddressIndex, "cdpCtAddressIndex")
    cdpCtAddressEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpCtAddressEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpCtAddressEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpCtAddressEntry.EntityData.Children = types.NewOrderedMap()
    cdpCtAddressEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpCtAddressEntry.EntityData.Leafs.Append("cdpCacheIfIndex", types.YLeaf{"CdpCacheIfIndex", cdpCtAddressEntry.CdpCacheIfIndex})
    cdpCtAddressEntry.EntityData.Leafs.Append("cdpCacheDeviceIndex", types.YLeaf{"CdpCacheDeviceIndex", cdpCtAddressEntry.CdpCacheDeviceIndex})
    cdpCtAddressEntry.EntityData.Leafs.Append("cdpCtAddressIndex", types.YLeaf{"CdpCtAddressIndex", cdpCtAddressEntry.CdpCtAddressIndex})
    cdpCtAddressEntry.EntityData.Leafs.Append("cdpCtAddressType", types.YLeaf{"CdpCtAddressType", cdpCtAddressEntry.CdpCtAddressType})
    cdpCtAddressEntry.EntityData.Leafs.Append("cdpCtAddress", types.YLeaf{"CdpCtAddress", cdpCtAddressEntry.CdpCtAddress})

    cdpCtAddressEntry.EntityData.YListKeys = []string {"CdpCacheIfIndex", "CdpCacheDeviceIndex", "CdpCtAddressIndex"}

    return &(cdpCtAddressEntry.EntityData)
}

