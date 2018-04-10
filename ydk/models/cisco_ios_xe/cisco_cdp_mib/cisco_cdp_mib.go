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

func (cISCOCDPMIB *CISCOCDPMIB) GetEntityData() *types.CommonEntityData {
    cISCOCDPMIB.EntityData.YFilter = cISCOCDPMIB.YFilter
    cISCOCDPMIB.EntityData.YangName = "CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCDPMIB.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.SegmentPath = "CISCO-CDP-MIB:CISCO-CDP-MIB"
    cISCOCDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCDPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOCDPMIB.EntityData.Children["cdpGlobal"] = types.YChild{"Cdpglobal", &cISCOCDPMIB.Cdpglobal}
    cISCOCDPMIB.EntityData.Children["cdpInterfaceTable"] = types.YChild{"Cdpinterfacetable", &cISCOCDPMIB.Cdpinterfacetable}
    cISCOCDPMIB.EntityData.Children["cdpInterfaceExtTable"] = types.YChild{"Cdpinterfaceexttable", &cISCOCDPMIB.Cdpinterfaceexttable}
    cISCOCDPMIB.EntityData.Children["cdpCacheTable"] = types.YChild{"Cdpcachetable", &cISCOCDPMIB.Cdpcachetable}
    cISCOCDPMIB.EntityData.Children["cdpCtAddressTable"] = types.YChild{"Cdpctaddresstable", &cISCOCDPMIB.Cdpctaddresstable}
    cISCOCDPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOCDPMIB.EntityData)
}

// CISCOCDPMIB_Cdpglobal
type CISCOCDPMIB_Cdpglobal struct {
    EntityData types.CommonEntityData
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

func (cdpglobal *CISCOCDPMIB_Cdpglobal) GetEntityData() *types.CommonEntityData {
    cdpglobal.EntityData.YFilter = cdpglobal.YFilter
    cdpglobal.EntityData.YangName = "cdpGlobal"
    cdpglobal.EntityData.BundleName = "cisco_ios_xe"
    cdpglobal.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpglobal.EntityData.SegmentPath = "cdpGlobal"
    cdpglobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpglobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpglobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpglobal.EntityData.Children = make(map[string]types.YChild)
    cdpglobal.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpglobal.EntityData.Leafs["cdpGlobalRun"] = types.YLeaf{"Cdpglobalrun", cdpglobal.Cdpglobalrun}
    cdpglobal.EntityData.Leafs["cdpGlobalMessageInterval"] = types.YLeaf{"Cdpglobalmessageinterval", cdpglobal.Cdpglobalmessageinterval}
    cdpglobal.EntityData.Leafs["cdpGlobalHoldTime"] = types.YLeaf{"Cdpglobalholdtime", cdpglobal.Cdpglobalholdtime}
    cdpglobal.EntityData.Leafs["cdpGlobalDeviceId"] = types.YLeaf{"Cdpglobaldeviceid", cdpglobal.Cdpglobaldeviceid}
    cdpglobal.EntityData.Leafs["cdpGlobalLastChange"] = types.YLeaf{"Cdpgloballastchange", cdpglobal.Cdpgloballastchange}
    cdpglobal.EntityData.Leafs["cdpGlobalDeviceIdFormatCpb"] = types.YLeaf{"Cdpglobaldeviceidformatcpb", cdpglobal.Cdpglobaldeviceidformatcpb}
    cdpglobal.EntityData.Leafs["cdpGlobalDeviceIdFormat"] = types.YLeaf{"Cdpglobaldeviceidformat", cdpglobal.Cdpglobaldeviceidformat}
    return &(cdpglobal.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpInterfaceTable, containing the status
    // of CDP on an interface. The type is slice of
    // CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry.
    Cdpinterfaceentry []CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry
}

func (cdpinterfacetable *CISCOCDPMIB_Cdpinterfacetable) GetEntityData() *types.CommonEntityData {
    cdpinterfacetable.EntityData.YFilter = cdpinterfacetable.YFilter
    cdpinterfacetable.EntityData.YangName = "cdpInterfaceTable"
    cdpinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    cdpinterfacetable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpinterfacetable.EntityData.SegmentPath = "cdpInterfaceTable"
    cdpinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpinterfacetable.EntityData.Children = make(map[string]types.YChild)
    cdpinterfacetable.EntityData.Children["cdpInterfaceEntry"] = types.YChild{"Cdpinterfaceentry", nil}
    for i := range cdpinterfacetable.Cdpinterfaceentry {
        cdpinterfacetable.EntityData.Children[types.GetSegmentPath(&cdpinterfacetable.Cdpinterfaceentry[i])] = types.YChild{"Cdpinterfaceentry", &cdpinterfacetable.Cdpinterfaceentry[i]}
    }
    cdpinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpinterfacetable.EntityData)
}

// CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry
// An entry (conceptual row) in the cdpInterfaceTable,
// containing the status of CDP on an interface.
type CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry struct {
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

func (cdpinterfaceentry *CISCOCDPMIB_Cdpinterfacetable_Cdpinterfaceentry) GetEntityData() *types.CommonEntityData {
    cdpinterfaceentry.EntityData.YFilter = cdpinterfaceentry.YFilter
    cdpinterfaceentry.EntityData.YangName = "cdpInterfaceEntry"
    cdpinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    cdpinterfaceentry.EntityData.ParentYangName = "cdpInterfaceTable"
    cdpinterfaceentry.EntityData.SegmentPath = "cdpInterfaceEntry" + "[cdpInterfaceIfIndex='" + fmt.Sprintf("%v", cdpinterfaceentry.Cdpinterfaceifindex) + "']"
    cdpinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    cdpinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpinterfaceentry.EntityData.Leafs["cdpInterfaceIfIndex"] = types.YLeaf{"Cdpinterfaceifindex", cdpinterfaceentry.Cdpinterfaceifindex}
    cdpinterfaceentry.EntityData.Leafs["cdpInterfaceEnable"] = types.YLeaf{"Cdpinterfaceenable", cdpinterfaceentry.Cdpinterfaceenable}
    cdpinterfaceentry.EntityData.Leafs["cdpInterfaceMessageInterval"] = types.YLeaf{"Cdpinterfacemessageinterval", cdpinterfaceentry.Cdpinterfacemessageinterval}
    cdpinterfaceentry.EntityData.Leafs["cdpInterfaceGroup"] = types.YLeaf{"Cdpinterfacegroup", cdpinterfaceentry.Cdpinterfacegroup}
    cdpinterfaceentry.EntityData.Leafs["cdpInterfacePort"] = types.YLeaf{"Cdpinterfaceport", cdpinterfaceentry.Cdpinterfaceport}
    cdpinterfaceentry.EntityData.Leafs["cdpInterfaceName"] = types.YLeaf{"Cdpinterfacename", cdpinterfaceentry.Cdpinterfacename}
    return &(cdpinterfaceentry.EntityData)
}

// CISCOCDPMIB_Cdpinterfaceexttable
// This table contains the additional CDP configuration on
// the device's interfaces.
type CISCOCDPMIB_Cdpinterfaceexttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cdpInterfaceExtTable contains the values configured for
    // Extented Trust TLV and COS (Class of Service) for Untrusted Ports TLV on an
    // interface which supports the sending of these TLVs. The type is slice of
    // CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry.
    Cdpinterfaceextentry []CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry
}

func (cdpinterfaceexttable *CISCOCDPMIB_Cdpinterfaceexttable) GetEntityData() *types.CommonEntityData {
    cdpinterfaceexttable.EntityData.YFilter = cdpinterfaceexttable.YFilter
    cdpinterfaceexttable.EntityData.YangName = "cdpInterfaceExtTable"
    cdpinterfaceexttable.EntityData.BundleName = "cisco_ios_xe"
    cdpinterfaceexttable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpinterfaceexttable.EntityData.SegmentPath = "cdpInterfaceExtTable"
    cdpinterfaceexttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpinterfaceexttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpinterfaceexttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpinterfaceexttable.EntityData.Children = make(map[string]types.YChild)
    cdpinterfaceexttable.EntityData.Children["cdpInterfaceExtEntry"] = types.YChild{"Cdpinterfaceextentry", nil}
    for i := range cdpinterfaceexttable.Cdpinterfaceextentry {
        cdpinterfaceexttable.EntityData.Children[types.GetSegmentPath(&cdpinterfaceexttable.Cdpinterfaceextentry[i])] = types.YChild{"Cdpinterfaceextentry", &cdpinterfaceexttable.Cdpinterfaceextentry[i]}
    }
    cdpinterfaceexttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpinterfaceexttable.EntityData)
}

// CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry
// An entry in the cdpInterfaceExtTable contains the values
// configured for Extented Trust TLV and COS (Class of Service)
// for Untrusted Ports TLV on an interface which supports the
// sending of these TLVs.
type CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry struct {
    EntityData types.CommonEntityData
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

func (cdpinterfaceextentry *CISCOCDPMIB_Cdpinterfaceexttable_Cdpinterfaceextentry) GetEntityData() *types.CommonEntityData {
    cdpinterfaceextentry.EntityData.YFilter = cdpinterfaceextentry.YFilter
    cdpinterfaceextentry.EntityData.YangName = "cdpInterfaceExtEntry"
    cdpinterfaceextentry.EntityData.BundleName = "cisco_ios_xe"
    cdpinterfaceextentry.EntityData.ParentYangName = "cdpInterfaceExtTable"
    cdpinterfaceextentry.EntityData.SegmentPath = "cdpInterfaceExtEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdpinterfaceextentry.Ifindex) + "']"
    cdpinterfaceextentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpinterfaceextentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpinterfaceextentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpinterfaceextentry.EntityData.Children = make(map[string]types.YChild)
    cdpinterfaceextentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpinterfaceextentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdpinterfaceextentry.Ifindex}
    cdpinterfaceextentry.EntityData.Leafs["cdpInterfaceExtendedTrust"] = types.YLeaf{"Cdpinterfaceextendedtrust", cdpinterfaceextentry.Cdpinterfaceextendedtrust}
    cdpinterfaceextentry.EntityData.Leafs["cdpInterfaceCosForUntrustedPort"] = types.YLeaf{"Cdpinterfacecosforuntrustedport", cdpinterfaceextentry.Cdpinterfacecosforuntrustedport}
    return &(cdpinterfaceextentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cdpCacheTable, containing the information
    // received via CDP on one interface from one device.  Entries appear when a
    // CDP advertisement is received from a neighbor device.  Entries disappear
    // when CDP is disabled on the interface, or globally. The type is slice of
    // CISCOCDPMIB_Cdpcachetable_Cdpcacheentry.
    Cdpcacheentry []CISCOCDPMIB_Cdpcachetable_Cdpcacheentry
}

func (cdpcachetable *CISCOCDPMIB_Cdpcachetable) GetEntityData() *types.CommonEntityData {
    cdpcachetable.EntityData.YFilter = cdpcachetable.YFilter
    cdpcachetable.EntityData.YangName = "cdpCacheTable"
    cdpcachetable.EntityData.BundleName = "cisco_ios_xe"
    cdpcachetable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpcachetable.EntityData.SegmentPath = "cdpCacheTable"
    cdpcachetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpcachetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpcachetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpcachetable.EntityData.Children = make(map[string]types.YChild)
    cdpcachetable.EntityData.Children["cdpCacheEntry"] = types.YChild{"Cdpcacheentry", nil}
    for i := range cdpcachetable.Cdpcacheentry {
        cdpcachetable.EntityData.Children[types.GetSegmentPath(&cdpcachetable.Cdpcacheentry[i])] = types.YChild{"Cdpcacheentry", &cdpcachetable.Cdpcacheentry[i]}
    }
    cdpcachetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpcachetable.EntityData)
}

// CISCOCDPMIB_Cdpcachetable_Cdpcacheentry
// An entry (conceptual row) in the cdpCacheTable,
// containing the information received via CDP on one
// interface from one device.  Entries appear when
// a CDP advertisement is received from a neighbor
// device.  Entries disappear when CDP is disabled
// on the interface, or globally.
type CISCOCDPMIB_Cdpcachetable_Cdpcacheentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (cdpcacheentry *CISCOCDPMIB_Cdpcachetable_Cdpcacheentry) GetEntityData() *types.CommonEntityData {
    cdpcacheentry.EntityData.YFilter = cdpcacheentry.YFilter
    cdpcacheentry.EntityData.YangName = "cdpCacheEntry"
    cdpcacheentry.EntityData.BundleName = "cisco_ios_xe"
    cdpcacheentry.EntityData.ParentYangName = "cdpCacheTable"
    cdpcacheentry.EntityData.SegmentPath = "cdpCacheEntry" + "[cdpCacheIfIndex='" + fmt.Sprintf("%v", cdpcacheentry.Cdpcacheifindex) + "']" + "[cdpCacheDeviceIndex='" + fmt.Sprintf("%v", cdpcacheentry.Cdpcachedeviceindex) + "']"
    cdpcacheentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpcacheentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpcacheentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpcacheentry.EntityData.Children = make(map[string]types.YChild)
    cdpcacheentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpcacheentry.EntityData.Leafs["cdpCacheIfIndex"] = types.YLeaf{"Cdpcacheifindex", cdpcacheentry.Cdpcacheifindex}
    cdpcacheentry.EntityData.Leafs["cdpCacheDeviceIndex"] = types.YLeaf{"Cdpcachedeviceindex", cdpcacheentry.Cdpcachedeviceindex}
    cdpcacheentry.EntityData.Leafs["cdpCacheAddressType"] = types.YLeaf{"Cdpcacheaddresstype", cdpcacheentry.Cdpcacheaddresstype}
    cdpcacheentry.EntityData.Leafs["cdpCacheAddress"] = types.YLeaf{"Cdpcacheaddress", cdpcacheentry.Cdpcacheaddress}
    cdpcacheentry.EntityData.Leafs["cdpCacheVersion"] = types.YLeaf{"Cdpcacheversion", cdpcacheentry.Cdpcacheversion}
    cdpcacheentry.EntityData.Leafs["cdpCacheDeviceId"] = types.YLeaf{"Cdpcachedeviceid", cdpcacheentry.Cdpcachedeviceid}
    cdpcacheentry.EntityData.Leafs["cdpCacheDevicePort"] = types.YLeaf{"Cdpcachedeviceport", cdpcacheentry.Cdpcachedeviceport}
    cdpcacheentry.EntityData.Leafs["cdpCachePlatform"] = types.YLeaf{"Cdpcacheplatform", cdpcacheentry.Cdpcacheplatform}
    cdpcacheentry.EntityData.Leafs["cdpCacheCapabilities"] = types.YLeaf{"Cdpcachecapabilities", cdpcacheentry.Cdpcachecapabilities}
    cdpcacheentry.EntityData.Leafs["cdpCacheVTPMgmtDomain"] = types.YLeaf{"Cdpcachevtpmgmtdomain", cdpcacheentry.Cdpcachevtpmgmtdomain}
    cdpcacheentry.EntityData.Leafs["cdpCacheNativeVLAN"] = types.YLeaf{"Cdpcachenativevlan", cdpcacheentry.Cdpcachenativevlan}
    cdpcacheentry.EntityData.Leafs["cdpCacheDuplex"] = types.YLeaf{"Cdpcacheduplex", cdpcacheentry.Cdpcacheduplex}
    cdpcacheentry.EntityData.Leafs["cdpCacheApplianceID"] = types.YLeaf{"Cdpcacheapplianceid", cdpcacheentry.Cdpcacheapplianceid}
    cdpcacheentry.EntityData.Leafs["cdpCacheVlanID"] = types.YLeaf{"Cdpcachevlanid", cdpcacheentry.Cdpcachevlanid}
    cdpcacheentry.EntityData.Leafs["cdpCachePowerConsumption"] = types.YLeaf{"Cdpcachepowerconsumption", cdpcacheentry.Cdpcachepowerconsumption}
    cdpcacheentry.EntityData.Leafs["cdpCacheMTU"] = types.YLeaf{"Cdpcachemtu", cdpcacheentry.Cdpcachemtu}
    cdpcacheentry.EntityData.Leafs["cdpCacheSysName"] = types.YLeaf{"Cdpcachesysname", cdpcacheentry.Cdpcachesysname}
    cdpcacheentry.EntityData.Leafs["cdpCacheSysObjectID"] = types.YLeaf{"Cdpcachesysobjectid", cdpcacheentry.Cdpcachesysobjectid}
    cdpcacheentry.EntityData.Leafs["cdpCachePrimaryMgmtAddrType"] = types.YLeaf{"Cdpcacheprimarymgmtaddrtype", cdpcacheentry.Cdpcacheprimarymgmtaddrtype}
    cdpcacheentry.EntityData.Leafs["cdpCachePrimaryMgmtAddr"] = types.YLeaf{"Cdpcacheprimarymgmtaddr", cdpcacheentry.Cdpcacheprimarymgmtaddr}
    cdpcacheentry.EntityData.Leafs["cdpCacheSecondaryMgmtAddrType"] = types.YLeaf{"Cdpcachesecondarymgmtaddrtype", cdpcacheentry.Cdpcachesecondarymgmtaddrtype}
    cdpcacheentry.EntityData.Leafs["cdpCacheSecondaryMgmtAddr"] = types.YLeaf{"Cdpcachesecondarymgmtaddr", cdpcacheentry.Cdpcachesecondarymgmtaddr}
    cdpcacheentry.EntityData.Leafs["cdpCachePhysLocation"] = types.YLeaf{"Cdpcachephyslocation", cdpcacheentry.Cdpcachephyslocation}
    cdpcacheentry.EntityData.Leafs["cdpCacheLastChange"] = types.YLeaf{"Cdpcachelastchange", cdpcacheentry.Cdpcachelastchange}
    return &(cdpcacheentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdpctaddresstable *CISCOCDPMIB_Cdpctaddresstable) GetEntityData() *types.CommonEntityData {
    cdpctaddresstable.EntityData.YFilter = cdpctaddresstable.YFilter
    cdpctaddresstable.EntityData.YangName = "cdpCtAddressTable"
    cdpctaddresstable.EntityData.BundleName = "cisco_ios_xe"
    cdpctaddresstable.EntityData.ParentYangName = "CISCO-CDP-MIB"
    cdpctaddresstable.EntityData.SegmentPath = "cdpCtAddressTable"
    cdpctaddresstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpctaddresstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpctaddresstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpctaddresstable.EntityData.Children = make(map[string]types.YChild)
    cdpctaddresstable.EntityData.Children["cdpCtAddressEntry"] = types.YChild{"Cdpctaddressentry", nil}
    for i := range cdpctaddresstable.Cdpctaddressentry {
        cdpctaddresstable.EntityData.Children[types.GetSegmentPath(&cdpctaddresstable.Cdpctaddressentry[i])] = types.YChild{"Cdpctaddressentry", &cdpctaddresstable.Cdpctaddressentry[i]}
    }
    cdpctaddresstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpctaddresstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdpctaddressentry *CISCOCDPMIB_Cdpctaddresstable_Cdpctaddressentry) GetEntityData() *types.CommonEntityData {
    cdpctaddressentry.EntityData.YFilter = cdpctaddressentry.YFilter
    cdpctaddressentry.EntityData.YangName = "cdpCtAddressEntry"
    cdpctaddressentry.EntityData.BundleName = "cisco_ios_xe"
    cdpctaddressentry.EntityData.ParentYangName = "cdpCtAddressTable"
    cdpctaddressentry.EntityData.SegmentPath = "cdpCtAddressEntry" + "[cdpCacheIfIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpcacheifindex) + "']" + "[cdpCacheDeviceIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpcachedeviceindex) + "']" + "[cdpCtAddressIndex='" + fmt.Sprintf("%v", cdpctaddressentry.Cdpctaddressindex) + "']"
    cdpctaddressentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpctaddressentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpctaddressentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpctaddressentry.EntityData.Children = make(map[string]types.YChild)
    cdpctaddressentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpctaddressentry.EntityData.Leafs["cdpCacheIfIndex"] = types.YLeaf{"Cdpcacheifindex", cdpctaddressentry.Cdpcacheifindex}
    cdpctaddressentry.EntityData.Leafs["cdpCacheDeviceIndex"] = types.YLeaf{"Cdpcachedeviceindex", cdpctaddressentry.Cdpcachedeviceindex}
    cdpctaddressentry.EntityData.Leafs["cdpCtAddressIndex"] = types.YLeaf{"Cdpctaddressindex", cdpctaddressentry.Cdpctaddressindex}
    cdpctaddressentry.EntityData.Leafs["cdpCtAddressType"] = types.YLeaf{"Cdpctaddresstype", cdpctaddressentry.Cdpctaddresstype}
    cdpctaddressentry.EntityData.Leafs["cdpCtAddress"] = types.YLeaf{"Cdpctaddress", cdpctaddressentry.Cdpctaddress}
    return &(cdpctaddressentry.EntityData)
}

