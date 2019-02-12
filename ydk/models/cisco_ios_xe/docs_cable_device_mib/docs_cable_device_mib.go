// This is the MIB Module for DOCSIS-compliant cable modems
// and cable-modem termination systems.
// 
// Copyright (C) The IETF Trust (2006).  This version
// of this MIB module was published in RFC 4639; for full
// legal notices see the RFC itself.
package docs_cable_device_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_cable_device_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-CABLE-DEVICE-MIB DOCS-CABLE-DEVICE-MIB}", reflect.TypeOf(DOCSCABLEDEVICEMIB{}))
    ydk.RegisterEntity("DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB", reflect.TypeOf(DOCSCABLEDEVICEMIB{}))
}

// DOCSCABLEDEVICEMIB
type DOCSCABLEDEVICEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DocsDevBase DOCSCABLEDEVICEMIB_DocsDevBase

    
    DocsDevSoftware DOCSCABLEDEVICEMIB_DocsDevSoftware

    
    DocsDevServer DOCSCABLEDEVICEMIB_DocsDevServer

    
    DocsDevEvent DOCSCABLEDEVICEMIB_DocsDevEvent

    
    DocsDevFilter DOCSCABLEDEVICEMIB_DocsDevFilter

    
    DocsDevCpe DOCSCABLEDEVICEMIB_DocsDevCpe

    // This table controls access to SNMP objects by network management stations. 
    // If the table is empty, access to SNMP objects is unrestricted.  The objects
    // in this table MUST NOT persist across reboots.  The objects in this table
    // are only accessible from cable devices that are not capable of operating in
    // SNMP Coexistence mode (RFC 3584) or in SNMPv3 mode (RFC 3410). See the
    // conformance section for details.  Note that some devices are required by
    // other specifications (e.g., the DOCSIS OSSIv1.1 specification) to support
    // the legacy SNMPv1/v2c docsDevNmAccess mode for backward compatibility. 
    // This table is deprecated.  Instead, use the SNMP coexistence MIBs from RFC
    // 3584, the TARGET and NOTIFICATION MIBs from RFC 3413, and the View-Based
    // Access Control Model (VACM) MIBs for all SNMP protocol versions from RFC
    // 3415.
    DocsDevNmAccessTable DOCSCABLEDEVICEMIB_DocsDevNmAccessTable

    // This table allows control of the reporting of event classes.  For each
    // event priority, a combination of logging and reporting mechanisms may be
    // chosen.  The mapping of event types to priorities is vendor dependent. 
    // Vendors may also choose to allow the user to control that mapping through
    // proprietary means.  Table entries MUST persist across reboots for CMTS
    // devices and MUST NOT persist across reboots for CM devices.
    DocsDevEvControlTable DOCSCABLEDEVICEMIB_DocsDevEvControlTable

    // Contains a log of network and device events that may be of interest in
    // fault isolation and troubleshooting. If the local(0) bit is set in
    // docsDevEvReporting, entries in this table MUST persist across reboots.
    DocsDevEventTable DOCSCABLEDEVICEMIB_DocsDevEventTable

    // A list of filters to apply to (bridged) LLC traffic. The filters in this
    // table are applied to incoming traffic on the appropriate interface(s) 
    // prior to any further processing (e.g., before the packet is handed off for
    // level 3 processing, or for bridging). The specific action taken when no
    // filter is matched is controlled by docsDevFilterLLCUnmatchedAction.  Table
    // entries MUST NOT persist across reboots for any device.
    DocsDevFilterLLCTable DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable

    // An ordered list of filters or classifiers to apply to IP traffic. Filter
    // application is ordered by the filter index, rather than by a best match
    // algorithm (Note that this implies that the filter table may have gaps in
    // the index values).  Packets that match no filters will have policy 0 in the
    // docsDevFilterPolicyTable applied to them, if it exists.  Otherwise, Packets
    // that match no filters are discarded or forwarded according to the setting
    // of docsDevFilterIpDefault.  Any IP packet can theoretically match multiple
    // rows of this table.  When considering a packet, the table is scanned in row
    // index order (e.g., filter 10 is checked before filter 20).  If the packet
    // matches that filter (which means that it matches ALL criteria for that
    // row), actions appropriate to docsDevFilterIpControl and
    // docsDevFilterPolicyId are taken.  If the packet was discarded processing is
    // complete.  If docsDevFilterIpContinue is set to true, the filter comparison
    // continues with the next row in the table, looking for additional matches. 
    // If the packet matches no filter in the table, the packet is accepted or
    // dropped for further processing according to the setting of
    // docsDevFilterIpDefault. If the packet is accepted, the actions specified by
    // policy group 0 (e.g., the rows in docsDevFilterPolicyTable that have a
    // value of 0 for docsDevFilterPolicyId) are taken, if that policy group
    // exists.  Logically, this table is consulted twice during the processing of
    // any IP packet: once upon its acceptance from the L2 entity, and once upon
    // its transmission to the L2 entity.  In actuality, for cable modems, IP
    // filtering is generally the only IP processing done for transit traffic. 
    // This means that inbound and outbound filtering can generally be done at the
    // same time with one pass through the filter table.  The objects in this
    // table are only accessible from cable devices that are not operating in
    // DiffServ MIB mode (RFC 3289).  See the conformance section for details. 
    // Note that some devices are required by other specifications (e.g., the
    // DOCSIS OSSIv1.1 specification) to support the legacy SNMPv1/v2c
    // docsDevFilter mode for backward compatibility.  Table entries MUST NOT
    // persist across reboots for any device.  This table is deprecated.  Instead,
    // use the DiffServ MIB from RFC3289.
    DocsDevFilterIpTable DOCSCABLEDEVICEMIB_DocsDevFilterIpTable

    // A Table that maps between a policy group ID and a set of pointers to
    // policies to be applied.  All rows with the same docsDevFilterPolicyId are
    // part of the same group of policy pointers and are applied in the order in
    // this table.  docsDevFilterPolicyTable exists to allow multiple policy
    // actions (referenced by policy pointers) to be applied to any given
    // classified packet. The policy actions are applied in index order. For
    // example:  Index   ID    Type    Action  1      1      TOS     1  9      5  
    // TOS     1  12     1      IPSEC   3  This says that a packet that matches a
    // filter with policy id 1 first has TOS policy 1 applied (which might set the
    // TOS bits to enable a higher priority) and next has the IPSEC policy 3
    // applied (which may result in the packets being dumped into a secure VPN to
    // a remote encryptor).  Policy ID 0 is reserved for default actions and is
    // applied only to packets that match no filters in docsDevFilterIpTable. 
    // Table entries MUST NOT persist across reboots for any device.  This table
    // is deprecated.  Instead, use the DiffServ MIB from RFC3289.
    DocsDevFilterPolicyTable DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable

    // Table used to describe Type of Service (TOS) bits processing.  This table
    // is an adjunct to the docsDevFilterIpTable and the docsDevFilterPolicy
    // table.  Entries in the latter table can point to specific rows in this (and
    // other) tables and cause specific actions to be taken. This table permits
    // the manipulation of the value of the Type of Service bits in the IP header
    // of the matched packet as follows:  Set the tosBits of the packet to   
    // (tosBits & docsDevFilterTosAndMask) |    docsDevFilterTosOrMask  This
    // construct allows you to do a clear and set of all the TOS bits in a
    // flexible manner.  Table entries MUST NOT persist across reboots for any
    // device.  This table is deprecated.  Instead, use the DiffServ MIB from
    // RFC3289.
    DocsDevFilterTosTable DOCSCABLEDEVICEMIB_DocsDevFilterTosTable

    // This table lists the IPv4 addresses seen (or permitted) as source addresses
    // in packets originating from the customer interface on this device.  In
    // addition, this table can be provisioned with the specific addresses
    // permitted for the CPEs via the normal row creation mechanisms.  Table
    // entries MUST NOT persist across reboots for any device.  N.B.  Management
    // action can add entries in this table and in docsDevCpeIpTable past the
    // value of docsDevCpeIpMax.  docsDevCpeIpMax ONLY restricts the ability of
    // the CM to add learned addresses automatically.  This table is deprecated
    // and is replaced by docsDevCpeInetTable.
    DocsDevCpeTable DOCSCABLEDEVICEMIB_DocsDevCpeTable

    // This table lists the IP addresses seen (or permitted) as source addresses
    // in packets originating from the customer interface on this device.  In
    // addition, this table can be provisioned with the specific addresses
    // permitted for the CPEs via the normal row creation mechanisms.  N.B. 
    // Management action can add entries in this table and in docsDevCpeIpTable
    // past the value of docsDevCpeIpMax.  docsDevCpeIpMax ONLY restricts the
    // ability of the CM to add learned addresses automatically.  Table entries
    // MUST NOT persist across reboots for any device.  This table exactly mirrors
    // docsDevCpeTable and applies to IPv4 and IPv6 addresses.
    DocsDevCpeInetTable DOCSCABLEDEVICEMIB_DocsDevCpeInetTable
}

func (dOCSCABLEDEVICEMIB *DOCSCABLEDEVICEMIB) GetEntityData() *types.CommonEntityData {
    dOCSCABLEDEVICEMIB.EntityData.YFilter = dOCSCABLEDEVICEMIB.YFilter
    dOCSCABLEDEVICEMIB.EntityData.YangName = "DOCS-CABLE-DEVICE-MIB"
    dOCSCABLEDEVICEMIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSCABLEDEVICEMIB.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    dOCSCABLEDEVICEMIB.EntityData.SegmentPath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB"
    dOCSCABLEDEVICEMIB.EntityData.AbsolutePath = dOCSCABLEDEVICEMIB.EntityData.SegmentPath
    dOCSCABLEDEVICEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSCABLEDEVICEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSCABLEDEVICEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSCABLEDEVICEMIB.EntityData.Children = types.NewOrderedMap()
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevBase", types.YChild{"DocsDevBase", &dOCSCABLEDEVICEMIB.DocsDevBase})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevSoftware", types.YChild{"DocsDevSoftware", &dOCSCABLEDEVICEMIB.DocsDevSoftware})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevServer", types.YChild{"DocsDevServer", &dOCSCABLEDEVICEMIB.DocsDevServer})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevEvent", types.YChild{"DocsDevEvent", &dOCSCABLEDEVICEMIB.DocsDevEvent})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevFilter", types.YChild{"DocsDevFilter", &dOCSCABLEDEVICEMIB.DocsDevFilter})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevCpe", types.YChild{"DocsDevCpe", &dOCSCABLEDEVICEMIB.DocsDevCpe})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevNmAccessTable", types.YChild{"DocsDevNmAccessTable", &dOCSCABLEDEVICEMIB.DocsDevNmAccessTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevEvControlTable", types.YChild{"DocsDevEvControlTable", &dOCSCABLEDEVICEMIB.DocsDevEvControlTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevEventTable", types.YChild{"DocsDevEventTable", &dOCSCABLEDEVICEMIB.DocsDevEventTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevFilterLLCTable", types.YChild{"DocsDevFilterLLCTable", &dOCSCABLEDEVICEMIB.DocsDevFilterLLCTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevFilterIpTable", types.YChild{"DocsDevFilterIpTable", &dOCSCABLEDEVICEMIB.DocsDevFilterIpTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevFilterPolicyTable", types.YChild{"DocsDevFilterPolicyTable", &dOCSCABLEDEVICEMIB.DocsDevFilterPolicyTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevFilterTosTable", types.YChild{"DocsDevFilterTosTable", &dOCSCABLEDEVICEMIB.DocsDevFilterTosTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevCpeTable", types.YChild{"DocsDevCpeTable", &dOCSCABLEDEVICEMIB.DocsDevCpeTable})
    dOCSCABLEDEVICEMIB.EntityData.Children.Append("docsDevCpeInetTable", types.YChild{"DocsDevCpeInetTable", &dOCSCABLEDEVICEMIB.DocsDevCpeInetTable})
    dOCSCABLEDEVICEMIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSCABLEDEVICEMIB.EntityData.YListKeys = []string {}

    return &(dOCSCABLEDEVICEMIB.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevBase
type DOCSCABLEDEVICEMIB_DocsDevBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Defines the current role of this device.  cm(1) is a Cable Modem,
    // cmtsActive(2) is a Cable Modem Termination System that is controlling the
    // system of cable modems, and cmtsBackup(3) is a CMTS that is currently
    // connected but is not controlling the system (not currently used).  In
    // general, if this device is a 'cm', its role will not change during
    // operation or between reboots.  If the device is a 'cmts' it may change
    // between cmtsActive and cmtsBackup and back again during normal operation. 
    // NB: At this time, the DOCSIS standards do not support the concept of a
    // backup CMTS, but cmtsBackup is included for completeness. The type is
    // DocsDevRole.
    DocsDevRole interface{}

    // The current date and time, with time zone information (if known).  If the
    // real data and time cannot be determined, this shall represent elapsed time
    // from boot relative to the standard epoch '1970-1-1,0:0:0.0'.  In other
    // words, if this agent has been up for 3 minutes and not been able to
    // determine what the actual date and time are, this object will return the
    // value '1970-1-1,0:03:0.0'. The type is string.
    DocsDevDateTime interface{}

    // Setting this object to true(1) causes the device to reset.  Reading this
    // object always returns false(2). The type is bool.
    DocsDevResetNow interface{}

    // The manufacturer's serial number for this device. The type is string.
    DocsDevSerialNumber interface{}

    // This object controls operation of the spanning tree protocol (as
    // distinguished from transparent bridging).  If set to stEnabled(1), then the
    // spanning tree protocol is enabled, subject to bridging constraints. If
    // noStFilterBpdu(2), then spanning tree is not active, and Bridge PDUs
    // received are discarded.  If noStPassBpdu(3), then spanning tree is not
    // active, and Bridge PDUs are transparently forwarded.  Note that a device
    // need not implement all of these options, but that noStFilterBpdu(2) is
    // required. The type is DocsDevSTPControl.
    DocsDevSTPControl interface{}

    // This object controls the IGMP mode of operation for the CM or CMTS.  In
    // passive mode, the device forwards IGMP between interfaces as based on
    // knowledge of Multicast Session activity on the subscriber side interface
    // and the rules defined in the DOCSIS RFI specification.  In active mode, the
    // device terminates at and initiates IGMP through its interfaces as based on
    // the knowledge of Multicast Session activity on the subscriber side
    // interface. The type is DocsDevIgmpModeControl.
    DocsDevIgmpModeControl interface{}

    // The maximum number of CPEs that can be granted access through a CM during a
    // CM epoch.  This value can be obtained from the CM configuration file;
    // however, it may be adjusted by the CM according to hardware or software
    // limitations that have been imposed on the implementation. The type is
    // interface{} with range: 0..255. Units are CPEs.
    DocsDevMaxCpe interface{}
}

func (docsDevBase *DOCSCABLEDEVICEMIB_DocsDevBase) GetEntityData() *types.CommonEntityData {
    docsDevBase.EntityData.YFilter = docsDevBase.YFilter
    docsDevBase.EntityData.YangName = "docsDevBase"
    docsDevBase.EntityData.BundleName = "cisco_ios_xe"
    docsDevBase.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevBase.EntityData.SegmentPath = "docsDevBase"
    docsDevBase.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevBase.EntityData.SegmentPath
    docsDevBase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevBase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevBase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevBase.EntityData.Children = types.NewOrderedMap()
    docsDevBase.EntityData.Leafs = types.NewOrderedMap()
    docsDevBase.EntityData.Leafs.Append("docsDevRole", types.YLeaf{"DocsDevRole", docsDevBase.DocsDevRole})
    docsDevBase.EntityData.Leafs.Append("docsDevDateTime", types.YLeaf{"DocsDevDateTime", docsDevBase.DocsDevDateTime})
    docsDevBase.EntityData.Leafs.Append("docsDevResetNow", types.YLeaf{"DocsDevResetNow", docsDevBase.DocsDevResetNow})
    docsDevBase.EntityData.Leafs.Append("docsDevSerialNumber", types.YLeaf{"DocsDevSerialNumber", docsDevBase.DocsDevSerialNumber})
    docsDevBase.EntityData.Leafs.Append("docsDevSTPControl", types.YLeaf{"DocsDevSTPControl", docsDevBase.DocsDevSTPControl})
    docsDevBase.EntityData.Leafs.Append("docsDevIgmpModeControl", types.YLeaf{"DocsDevIgmpModeControl", docsDevBase.DocsDevIgmpModeControl})
    docsDevBase.EntityData.Leafs.Append("docsDevMaxCpe", types.YLeaf{"DocsDevMaxCpe", docsDevBase.DocsDevMaxCpe})

    docsDevBase.EntityData.YListKeys = []string {}

    return &(docsDevBase.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl represents subscriber side interface.
type DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl string

const (
    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl_passive DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl = "passive"

    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl_active DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevIgmpModeControl = "active"
)

// DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole represents completeness.
type DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole string

const (
    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole_cm DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole = "cm"

    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole_cmtsActive DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole = "cmtsActive"

    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole_cmtsBackup DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevRole = "cmtsBackup"
)

// DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl represents options, but that noStFilterBpdu(2) is required.
type DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl string

const (
    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl_stEnabled DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl = "stEnabled"

    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl_noStFilterBpdu DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl = "noStFilterBpdu"

    DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl_noStPassBpdu DOCSCABLEDEVICEMIB_DocsDevBase_DocsDevSTPControl = "noStPassBpdu"
)

// DOCSCABLEDEVICEMIB_DocsDevSoftware
type DOCSCABLEDEVICEMIB_DocsDevSoftware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The address of the TFTP server used for software upgrades.  If the TFTP
    // server is unknown or is a non-IPv4 address, return 0.0.0.0.  This object is
    // deprecated.  See docsDevSwServerAddress for its replacement.  This object
    // will have its value modified, given a valid SET to docsDevSwServerAddress.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevSwServer interface{}

    // The filename of the software image to be downloaded via TFTP, or the
    // abs_path (as defined in RFC2616) of the software image to be downloaded via
    // HTTP.  Unless set via SNMP, this is the filename or abs_path specified by
    // the provisioning server during the boot process  that corresponds to the
    // software version that is desired for this device.  If unknown, the value of
    // this object is the zero-length string. The type is string with length:
    // 0..64.
    DocsDevSwFilename interface{}

    // If set to upgradeFromMgt(1), the device will initiate a TFTP or HTTP
    // software image download.  After successfully receiving an image, the device
    // will set its state to ignoreProvisioningUpgrade(3) and reboot. If the
    // download process is interrupted (e.g., by a reset or power failure), the
    // device will load the previous image and, after re-initialization, continue
    // to attempt loading the image specified in docsDevSwFilename.  If set to
    // allowProvisioningUpgrade(2), the device will use the software version
    // information supplied by the provisioning server when next rebooting (this
    // does not cause a reboot).  When set to ignoreProvisioningUpgrade(3), the
    // device will disregard software image upgrade information from the
    // provisioning server.  Note that reading this object can return
    // upgradeFromMgt(1).  This indicates that a software download is currently in
    // progress, and that the device will reboot after successfully receiving an
    // image. The type is DocsDevSwAdminStatus.
    DocsDevSwAdminStatus interface{}

    // InProgress(1) indicates that a TFTP or HTTP download is underway, either as
    // a result of a version mismatch at provisioning or as a result of a
    // upgradeFromMgt request. No other docsDevSw* objects can be modified in this
    // state.  CompleteFromProvisioning(2) indicates that the last software
    // upgrade was a result of version mismatch at provisioning. 
    // CompleteFromMgt(3) indicates that the last software upgrade was a result of
    // setting docsDevSwAdminStatus to upgradeFromMgt.  Failed(4) indicates that
    // the last attempted download failed, ordinarily due to TFTP or HTTP timeout.
    // The type is DocsDevSwOperStatus.
    DocsDevSwOperStatus interface{}

    // The software version currently operating in this device. This string's
    // syntax is that used by the individual vendor to identify software versions.
    // For a CM, this string will describe the current software load.  For a CMTS,
    // this object SHOULD contain a human-readable representation either of the
    // vendor specific designation of the software for the chassis, or of the
    // software for the control processor.  If neither of these is applicable, the
    // value MUST be a zero-length string. The type is string.
    DocsDevSwCurrentVers interface{}

    // The type of address of the TFTP or HTTP server used for software upgrades. 
    // If docsDevSwServerTransportProtocol is currently set to tftp(1), attempting
    // to set this object to dns(16) MUST result in an error. The type is
    // InetAddressType.
    DocsDevSwServerAddressType interface{}

    // The address of the TFTP or HTTP server used for software upgrades.  If the
    // TFTP/HTTP server is unknown, return the zero- length address string (see
    // the TextualConvention).  If docsDevSwServer is also implemented in this
    // agent, this object is tied to it.  A set of this object to an IPv4 address
    // will result in also setting the value of docsDevSwServer to that address. 
    // If this object is set to an IPv6 address, docsDevSwServer is set to
    // 0.0.0.0. If docsDevSwServer is set, this object is also set to that value. 
    // Note that if both are set in the same action, the order of which one sets
    // the other is undefined. The type is string with length: 0..255.
    DocsDevSwServerAddress interface{}

    // This object specifies the transport protocol (TFTP or HTTP) to be used for
    // software upgrades.  If the value of this object is tftp(1), then the cable
    // device uses TFTP (RFC1350) read request packets to download the
    // docsDevSwFilename from the docsDevSwServerAddress in octet mode.  If the
    // value of this object is http(2), then the cable device uses HTTP 1.0
    // (RFC1945) or HTTP 1.1 (RFC2616) GET requests sent to host
    // docsDevSwServerAddress to download the software image from path
    // docsDevSwFilename.  If docsDevSwServerAddressType is currently set to
    // dns(16), attempting to set this object to tftp(1) MUST result in an error.
    // The type is DocsDevSwServerTransportProtocol.
    DocsDevSwServerTransportProtocol interface{}
}

func (docsDevSoftware *DOCSCABLEDEVICEMIB_DocsDevSoftware) GetEntityData() *types.CommonEntityData {
    docsDevSoftware.EntityData.YFilter = docsDevSoftware.YFilter
    docsDevSoftware.EntityData.YangName = "docsDevSoftware"
    docsDevSoftware.EntityData.BundleName = "cisco_ios_xe"
    docsDevSoftware.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevSoftware.EntityData.SegmentPath = "docsDevSoftware"
    docsDevSoftware.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevSoftware.EntityData.SegmentPath
    docsDevSoftware.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevSoftware.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevSoftware.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevSoftware.EntityData.Children = types.NewOrderedMap()
    docsDevSoftware.EntityData.Leafs = types.NewOrderedMap()
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwServer", types.YLeaf{"DocsDevSwServer", docsDevSoftware.DocsDevSwServer})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwFilename", types.YLeaf{"DocsDevSwFilename", docsDevSoftware.DocsDevSwFilename})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwAdminStatus", types.YLeaf{"DocsDevSwAdminStatus", docsDevSoftware.DocsDevSwAdminStatus})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwOperStatus", types.YLeaf{"DocsDevSwOperStatus", docsDevSoftware.DocsDevSwOperStatus})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwCurrentVers", types.YLeaf{"DocsDevSwCurrentVers", docsDevSoftware.DocsDevSwCurrentVers})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwServerAddressType", types.YLeaf{"DocsDevSwServerAddressType", docsDevSoftware.DocsDevSwServerAddressType})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwServerAddress", types.YLeaf{"DocsDevSwServerAddress", docsDevSoftware.DocsDevSwServerAddress})
    docsDevSoftware.EntityData.Leafs.Append("docsDevSwServerTransportProtocol", types.YLeaf{"DocsDevSwServerTransportProtocol", docsDevSoftware.DocsDevSwServerTransportProtocol})

    docsDevSoftware.EntityData.YListKeys = []string {}

    return &(docsDevSoftware.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus represents will reboot after successfully receiving an image.
type DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus string

const (
    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus_upgradeFromMgt DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus = "upgradeFromMgt"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus_allowProvisioningUpgrade DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus = "allowProvisioningUpgrade"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus_ignoreProvisioningUpgrade DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwAdminStatus = "ignoreProvisioningUpgrade"
)

// DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus represents failed, ordinarily due to TFTP or HTTP timeout.
type DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus string

const (
    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus_inProgress DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus = "inProgress"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus_completeFromProvisioning DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus = "completeFromProvisioning"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus_completeFromMgt DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus = "completeFromMgt"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus_failed DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus = "failed"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus_other DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwOperStatus = "other"
)

// DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol represents result in an error.
type DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol string

const (
    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol_tftp DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol = "tftp"

    DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol_http DOCSCABLEDEVICEMIB_DocsDevSoftware_DocsDevSwServerTransportProtocol = "http"
)

// DOCSCABLEDEVICEMIB_DocsDevServer
type DOCSCABLEDEVICEMIB_DocsDevServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If operational(1), the device has completed loading and processing of
    // configuration parameters, and the CMTS has completed the Registration
    // exchange.  If disabled(2), then the device was administratively disabled,
    // possibly by being refused network access in the configuration file.  If
    // waitingForDhcpOffer(3), then a Dynamic Host Configuration Protocol (DHCP)
    // Discover has been transmitted, and no offer has yet been received.  If
    // waitingForDhcpResponse(4), then a DHCP Request has been transmitted, and no
    // response has yet been received.  If waitingForTimeServer(5), then a Time
    // Request has been transmitted, and no response has yet been received. If
    // waitingForTftp(6), then a request to the TFTP parameter server has been
    // made, and no response received.  If refusedByCmts(7), then the Registration
    // Request/Response exchange with the CMTS failed.  If forwardingDenied(8),
    // then the registration process was completed, but the network access option
    // in the received configuration file prohibits forwarding.  If other(9), then
    // the registration process reached a point that does not fall into one of the
    // above categories.  If unknown(10), then the device has not yet begun the
    // registration process or is in some other indeterminate state. The type is
    // DocsDevServerBootState.
    DocsDevServerBootState interface{}

    // The IP address of the DHCP server that assigned an IP address to this
    // device. Returns 0.0.0.0 if DHCP is not used for IP address assignment, or
    // if this agent is not assigned an IPv4 address.  This object is deprecated
    // and is replaced by docsDevServerDhcpAddress. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevServerDhcp interface{}

    // The IP address of the Time server (RFC 0868).  Returns 0.0.0.0 if the time
    // server IP address is unknown, or if the time server is not an IPv4 server. 
    // This object is deprecated and is replaced by docsDevServerTimeAddress. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevServerTime interface{}

    // The IP address of the TFTP server responsible for downloading provisioning
    // and configuration parameters to this device. Returns 0.0.0.0 if the TFTP
    // server address is unknown or is not an IPv4 address.  This object is
    // deprecated and is replaced by docsDevServerConfigTftpAddress. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevServerTftp interface{}

    // The name of the device configuration file read from the TFTP server. 
    // Returns a zero-length string if the configuration file name is unknown. The
    // type is string.
    DocsDevServerConfigFile interface{}

    // The type of address of docsDevServerDhcpAddress.  If DHCP was not used,
    // this value should return unknown(0). The type is InetAddressType.
    DocsDevServerDhcpAddressType interface{}

    // The internet address of the DHCP server that assigned an IP address to this
    // device. Returns the zero length octet string if DHCP was not used for IP
    // address assignment. The type is string with length: 0..255.
    DocsDevServerDhcpAddress interface{}

    // The type of address of docsDevServerTimeAddress.  If no time server exists,
    // this value should return unknown(0). The type is InetAddressType.
    DocsDevServerTimeAddressType interface{}

    // The Internet address of the RFC 868 Time server, as provided by DHCP option
    // 4.  Note that if multiple values are provided to the CM in DHCP option 4,
    // the value of this MIB object MUST be the Time server address from which the
    // Time of Day reference was acquired as based on the DOCSIS RFI
    // specification.  During the period of time where the Time of Day have not
    // been acquired, the Time server address reported by the CM may report the
    // first address value in the DHCP option value or the last server address the
    // CM attempted to get the Time of day value.  Returns the zero-length octet
    // string if the time server IP address is not provisioned. The type is string
    // with length: 0..255.
    DocsDevServerTimeAddress interface{}

    // The type of address of docsDevServerConfigTftpAddress. If no TFTP server
    // exists, this value should return unknown(0). The type is InetAddressType.
    DocsDevServerConfigTftpAddressType interface{}

    // The internet address of the TFTP server responsible for downloading
    // provisioning and configuration parameters to this device.  Returns the
    // zero-length octet string if the config server address is unknown.  There
    // are certain security risks that are involved with using TFTP. The type is
    // string with length: 0..255.
    DocsDevServerConfigTftpAddress interface{}
}

func (docsDevServer *DOCSCABLEDEVICEMIB_DocsDevServer) GetEntityData() *types.CommonEntityData {
    docsDevServer.EntityData.YFilter = docsDevServer.YFilter
    docsDevServer.EntityData.YangName = "docsDevServer"
    docsDevServer.EntityData.BundleName = "cisco_ios_xe"
    docsDevServer.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevServer.EntityData.SegmentPath = "docsDevServer"
    docsDevServer.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevServer.EntityData.SegmentPath
    docsDevServer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevServer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevServer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevServer.EntityData.Children = types.NewOrderedMap()
    docsDevServer.EntityData.Leafs = types.NewOrderedMap()
    docsDevServer.EntityData.Leafs.Append("docsDevServerBootState", types.YLeaf{"DocsDevServerBootState", docsDevServer.DocsDevServerBootState})
    docsDevServer.EntityData.Leafs.Append("docsDevServerDhcp", types.YLeaf{"DocsDevServerDhcp", docsDevServer.DocsDevServerDhcp})
    docsDevServer.EntityData.Leafs.Append("docsDevServerTime", types.YLeaf{"DocsDevServerTime", docsDevServer.DocsDevServerTime})
    docsDevServer.EntityData.Leafs.Append("docsDevServerTftp", types.YLeaf{"DocsDevServerTftp", docsDevServer.DocsDevServerTftp})
    docsDevServer.EntityData.Leafs.Append("docsDevServerConfigFile", types.YLeaf{"DocsDevServerConfigFile", docsDevServer.DocsDevServerConfigFile})
    docsDevServer.EntityData.Leafs.Append("docsDevServerDhcpAddressType", types.YLeaf{"DocsDevServerDhcpAddressType", docsDevServer.DocsDevServerDhcpAddressType})
    docsDevServer.EntityData.Leafs.Append("docsDevServerDhcpAddress", types.YLeaf{"DocsDevServerDhcpAddress", docsDevServer.DocsDevServerDhcpAddress})
    docsDevServer.EntityData.Leafs.Append("docsDevServerTimeAddressType", types.YLeaf{"DocsDevServerTimeAddressType", docsDevServer.DocsDevServerTimeAddressType})
    docsDevServer.EntityData.Leafs.Append("docsDevServerTimeAddress", types.YLeaf{"DocsDevServerTimeAddress", docsDevServer.DocsDevServerTimeAddress})
    docsDevServer.EntityData.Leafs.Append("docsDevServerConfigTftpAddressType", types.YLeaf{"DocsDevServerConfigTftpAddressType", docsDevServer.DocsDevServerConfigTftpAddressType})
    docsDevServer.EntityData.Leafs.Append("docsDevServerConfigTftpAddress", types.YLeaf{"DocsDevServerConfigTftpAddress", docsDevServer.DocsDevServerConfigTftpAddress})

    docsDevServer.EntityData.YListKeys = []string {}

    return &(docsDevServer.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState represents state.
type DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState string

const (
    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_operational DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "operational"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_disabled DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "disabled"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_waitingForDhcpOffer DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "waitingForDhcpOffer"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_waitingForDhcpResponse DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "waitingForDhcpResponse"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_waitingForTimeServer DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "waitingForTimeServer"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_waitingForTftp DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "waitingForTftp"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_refusedByCmts DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "refusedByCmts"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_forwardingDenied DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "forwardingDenied"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_other DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "other"

    DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState_unknown DOCSCABLEDEVICEMIB_DocsDevServer_DocsDevServerBootState = "unknown"
)

// DOCSCABLEDEVICEMIB_DocsDevEvent
type DOCSCABLEDEVICEMIB_DocsDevEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Setting this object to resetLog(1) empties the event log.  All data is
    // deleted.  Setting it to useDefaultReporting(2) returns all event priorities
    // to their factory-default reporting.  Reading this object always returns
    // useDefaultReporting(2). The type is DocsDevEvControl.
    DocsDevEvControl interface{}

    // The IP address of the Syslog server. If 0.0.0.0, either syslog transmission
    // is inhibited, or the Syslog server address is not an IPv4 address.  This
    // object is deprecated and is replaced by docsDevEvSyslogAddress. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevEvSyslog interface{}

    // Controls the transmission of traps and syslog messages with respect to the
    // trap pacing threshold.  unconstrained(1) causes traps and syslog messages
    // to be transmitted without regard to the threshold settings. 
    // maintainBelowThreshold(2) causes trap transmission and syslog messages to
    // be suppressed if the number of traps would otherwise exceed the threshold. 
    // stopAtThreshold(3) causes trap transmission to cease at the threshold and
    // not to resume until directed to do so.  inhibited(4) causes all trap
    // transmission and syslog messages to be suppressed.  A single event is
    // always treated as a single event for threshold counting.  That is, an event
    // causing both a trap and a syslog message is still treated as a single
    // event.  Writing to this object resets the thresholding state. The type is
    // DocsDevEvThrottleAdminStatus.
    DocsDevEvThrottleAdminStatus interface{}

    // If true(1), trap and syslog transmission is currently inhibited due to
    // thresholds and/or the current setting of docsDevEvThrottleAdminStatus.  In
    // addition, this is true(1) when transmission is inhibited because no syslog
    // (docsDevEvSyslog) or trap (docsDevNmAccessEntry) destinations have been
    // set.  This object is deprecated and is replaced by
    // docsDevEvThrottleThresholdExceeded. The type is bool.
    DocsDevEvThrottleInhibited interface{}

    // Number of events per docsDevEvThrottleInterval permitted before throttling
    // is to occur.  A single event, whether the notification could result in
    // messages transmitted using syslog, SNMP, or both protocols, and regardless
    // of the number of destinations, (including zero) is always treated as a
    // single event for threshold counting.  For example, an event causing both a
    // trap and a syslog message is still treated as a single event.  All system
    // notifications that occur within the device should be taken into
    // consideration when calculating and monitoring the threshold. The type is
    // interface{} with range: 0..4294967295. Units are events.
    DocsDevEvThrottleThreshold interface{}

    // The interval over which docsDevEvThrottleThreshold applies. The type is
    // interface{} with range: 1..2147483647. Units are seconds.
    DocsDevEvThrottleInterval interface{}

    // The type of address of docsDevEvSyslogAddress.  If no syslog server exists,
    // this value should return unknown(0). The type is InetAddressType.
    DocsDevEvSyslogAddressType interface{}

    // The Internet address of the Syslog server, as provided by DHCP option 7 or
    // set via SNMP management.  If the address of the server is set to the
    // zero-length string, the 0.0.0.0 IPv4 address, or the 0: IPv6 address,
    // Syslog transmission is inhibited.  Note that if multiple values are
    // provided to the CM in DHCP option 7, the value of this MIB object MUST be
    // the first Syslog server address received.  By default at agent boot, this
    // object returns the zero length string. The type is string with length:
    // 0..255.
    DocsDevEvSyslogAddress interface{}

    // If true(1), trap and syslog transmission is currently inhibited due to
    // exceeding the trap/syslog event threshold in the current interval. The type
    // is bool.
    DocsDevEvThrottleThresholdExceeded interface{}
}

func (docsDevEvent *DOCSCABLEDEVICEMIB_DocsDevEvent) GetEntityData() *types.CommonEntityData {
    docsDevEvent.EntityData.YFilter = docsDevEvent.YFilter
    docsDevEvent.EntityData.YangName = "docsDevEvent"
    docsDevEvent.EntityData.BundleName = "cisco_ios_xe"
    docsDevEvent.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevEvent.EntityData.SegmentPath = "docsDevEvent"
    docsDevEvent.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevEvent.EntityData.SegmentPath
    docsDevEvent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevEvent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevEvent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevEvent.EntityData.Children = types.NewOrderedMap()
    docsDevEvent.EntityData.Leafs = types.NewOrderedMap()
    docsDevEvent.EntityData.Leafs.Append("docsDevEvControl", types.YLeaf{"DocsDevEvControl", docsDevEvent.DocsDevEvControl})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvSyslog", types.YLeaf{"DocsDevEvSyslog", docsDevEvent.DocsDevEvSyslog})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvThrottleAdminStatus", types.YLeaf{"DocsDevEvThrottleAdminStatus", docsDevEvent.DocsDevEvThrottleAdminStatus})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvThrottleInhibited", types.YLeaf{"DocsDevEvThrottleInhibited", docsDevEvent.DocsDevEvThrottleInhibited})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvThrottleThreshold", types.YLeaf{"DocsDevEvThrottleThreshold", docsDevEvent.DocsDevEvThrottleThreshold})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvThrottleInterval", types.YLeaf{"DocsDevEvThrottleInterval", docsDevEvent.DocsDevEvThrottleInterval})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvSyslogAddressType", types.YLeaf{"DocsDevEvSyslogAddressType", docsDevEvent.DocsDevEvSyslogAddressType})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvSyslogAddress", types.YLeaf{"DocsDevEvSyslogAddress", docsDevEvent.DocsDevEvSyslogAddress})
    docsDevEvent.EntityData.Leafs.Append("docsDevEvThrottleThresholdExceeded", types.YLeaf{"DocsDevEvThrottleThresholdExceeded", docsDevEvent.DocsDevEvThrottleThresholdExceeded})

    docsDevEvent.EntityData.YListKeys = []string {}

    return &(docsDevEvent.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl represents always returns useDefaultReporting(2).
type DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl string

const (
    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl_resetLog DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl = "resetLog"

    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl_useDefaultReporting DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvControl = "useDefaultReporting"
)

// DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus represents Writing to this object resets the thresholding state.
type DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus string

const (
    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus_unconstrained DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus = "unconstrained"

    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus_maintainBelowThreshold DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus = "maintainBelowThreshold"

    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus_stopAtThreshold DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus = "stopAtThreshold"

    DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus_inhibited DOCSCABLEDEVICEMIB_DocsDevEvent_DocsDevEvThrottleAdminStatus = "inhibited"
)

// DOCSCABLEDEVICEMIB_DocsDevFilter
type DOCSCABLEDEVICEMIB_DocsDevFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LLC (Link Level Control) filters can be defined on an inclusive or
    // exclusive basis: CMs can be configured to forward only packets matching a
    // set of layer three protocols, or to drop packets matching a set of layer
    // three protocols.  Typical use of these filters is to filter out possibly
    // harmful (given the context of a large metropolitan LAN) protocols.  If set
    // to discard(1), any L2 packet that does not match at least one filter in the
    // docsDevFilterLLCTable will be discarded.  If set to accept(2), any L2
    // packet that does not match at least one filter in the docsDevFilterLLCTable
    // will be accepted for further processing (e.g., bridging).  In other words,
    // if the packet does not match an entry in the table, it takes this action;
    // if it does match an entry in the table, it takes the opposite of this
    // action. The type is DocsDevFilterLLCUnmatchedAction.
    DocsDevFilterLLCUnmatchedAction interface{}

    // The default behavior for (bridged) packets that do not match IP filters (or
    // Internet filters, if implemented) is defined by docsDevFilterIpDefault.  If
    // set to discard(1), all packets not matching an IP filter in
    // docsDevFilterIpTable will be discarded.  If set to accept(2), all packets
    // not matching an IP filter or an Internet filter will be accepted for
    // further processing (e.g., bridging). The type is DocsDevFilterIpDefault.
    DocsDevFilterIpDefault interface{}
}

func (docsDevFilter *DOCSCABLEDEVICEMIB_DocsDevFilter) GetEntityData() *types.CommonEntityData {
    docsDevFilter.EntityData.YFilter = docsDevFilter.YFilter
    docsDevFilter.EntityData.YangName = "docsDevFilter"
    docsDevFilter.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilter.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevFilter.EntityData.SegmentPath = "docsDevFilter"
    docsDevFilter.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevFilter.EntityData.SegmentPath
    docsDevFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilter.EntityData.Children = types.NewOrderedMap()
    docsDevFilter.EntityData.Leafs = types.NewOrderedMap()
    docsDevFilter.EntityData.Leafs.Append("docsDevFilterLLCUnmatchedAction", types.YLeaf{"DocsDevFilterLLCUnmatchedAction", docsDevFilter.DocsDevFilterLLCUnmatchedAction})
    docsDevFilter.EntityData.Leafs.Append("docsDevFilterIpDefault", types.YLeaf{"DocsDevFilterIpDefault", docsDevFilter.DocsDevFilterIpDefault})

    docsDevFilter.EntityData.YListKeys = []string {}

    return &(docsDevFilter.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault represents processing (e.g., bridging).
type DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault string

const (
    DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault_discard DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault = "discard"

    DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault_accept DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterIpDefault = "accept"
)

// DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction represents takes the opposite of this action.
type DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction string

const (
    DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction_discard DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction = "discard"

    DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction_accept DOCSCABLEDEVICEMIB_DocsDevFilter_DocsDevFilterLLCUnmatchedAction = "accept"
)

// DOCSCABLEDEVICEMIB_DocsDevCpe
type DOCSCABLEDEVICEMIB_DocsDevCpe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls the population of docsDevFilterCpeTable. If set to
    // none, the filters must be set manually by a network management action
    // (either configuration or SNMP set). If set to any, the CM wiretaps the
    // packets originating from the ethernet and enrolls up to docsDevCpeIpMax
    // addresses as based on the source IPv4 or v6 addresses of those packets. The
    // type is DocsDevCpeEnroll.
    DocsDevCpeEnroll interface{}

    // This object controls the maximum number of CPEs allowed to be learned
    // behind this device. If set to zero, any number of CPEs may connect up to
    // the maximum permitted for the device. If set to -1, no filtering is done on
    // CPE source addresses, and no entries are made in the docsDevFilterCpeTable
    // via learning.  If an attempt is made to set this to a number greater than
    // that permitted for the device, it is set to that maximum. The type is
    // interface{} with range: -1..2147483647.
    DocsDevCpeIpMax interface{}
}

func (docsDevCpe *DOCSCABLEDEVICEMIB_DocsDevCpe) GetEntityData() *types.CommonEntityData {
    docsDevCpe.EntityData.YFilter = docsDevCpe.YFilter
    docsDevCpe.EntityData.YangName = "docsDevCpe"
    docsDevCpe.EntityData.BundleName = "cisco_ios_xe"
    docsDevCpe.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevCpe.EntityData.SegmentPath = "docsDevCpe"
    docsDevCpe.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevCpe.EntityData.SegmentPath
    docsDevCpe.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevCpe.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevCpe.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevCpe.EntityData.Children = types.NewOrderedMap()
    docsDevCpe.EntityData.Leafs = types.NewOrderedMap()
    docsDevCpe.EntityData.Leafs.Append("docsDevCpeEnroll", types.YLeaf{"DocsDevCpeEnroll", docsDevCpe.DocsDevCpeEnroll})
    docsDevCpe.EntityData.Leafs.Append("docsDevCpeIpMax", types.YLeaf{"DocsDevCpeIpMax", docsDevCpe.DocsDevCpeIpMax})

    docsDevCpe.EntityData.YListKeys = []string {}

    return &(docsDevCpe.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll represents those packets.
type DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll string

const (
    DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll_none DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll = "none"

    DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll_any DOCSCABLEDEVICEMIB_DocsDevCpe_DocsDevCpeEnroll = "any"
)

// DOCSCABLEDEVICEMIB_DocsDevNmAccessTable
// This table controls access to SNMP objects by network
// management stations.  If the table is empty, access to
// SNMP objects is unrestricted.  The objects in this table
// MUST NOT persist across reboots.  The objects in this
// table are only accessible from cable devices that are
// not capable of operating in SNMP Coexistence mode
// (RFC 3584) or in SNMPv3 mode (RFC 3410).
// See the conformance section for
// details.  Note that some devices are required by other
// specifications (e.g., the DOCSIS OSSIv1.1 specification)
// to support the legacy SNMPv1/v2c docsDevNmAccess mode
// for backward compatibility.
// 
// This table is deprecated.  Instead, use the SNMP
// coexistence MIBs from RFC 3584, the TARGET and
// NOTIFICATION MIBs from RFC 3413, and
// the View-Based Access Control Model (VACM) MIBs for
// all SNMP protocol versions from RFC 3415.
type DOCSCABLEDEVICEMIB_DocsDevNmAccessTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describing access to SNMP objects by a particular network
    // management station. An entry in this table is not readable unless the
    // management station has read-write permission (either implicit if the table
    // is empty, or explicit through an entry in this table). Entries are ordered
    // by docsDevNmAccessIndex.  The first matching entry (e.g., matching IP
    // address and community string) is used to derive access. The type is slice
    // of DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry.
    DocsDevNmAccessEntry []*DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry
}

func (docsDevNmAccessTable *DOCSCABLEDEVICEMIB_DocsDevNmAccessTable) GetEntityData() *types.CommonEntityData {
    docsDevNmAccessTable.EntityData.YFilter = docsDevNmAccessTable.YFilter
    docsDevNmAccessTable.EntityData.YangName = "docsDevNmAccessTable"
    docsDevNmAccessTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevNmAccessTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevNmAccessTable.EntityData.SegmentPath = "docsDevNmAccessTable"
    docsDevNmAccessTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevNmAccessTable.EntityData.SegmentPath
    docsDevNmAccessTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevNmAccessTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevNmAccessTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevNmAccessTable.EntityData.Children = types.NewOrderedMap()
    docsDevNmAccessTable.EntityData.Children.Append("docsDevNmAccessEntry", types.YChild{"DocsDevNmAccessEntry", nil})
    for i := range docsDevNmAccessTable.DocsDevNmAccessEntry {
        docsDevNmAccessTable.EntityData.Children.Append(types.GetSegmentPath(docsDevNmAccessTable.DocsDevNmAccessEntry[i]), types.YChild{"DocsDevNmAccessEntry", docsDevNmAccessTable.DocsDevNmAccessEntry[i]})
    }
    docsDevNmAccessTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevNmAccessTable.EntityData.YListKeys = []string {}

    return &(docsDevNmAccessTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry
// An entry describing access to SNMP objects by a
// particular network management station. An entry in
// this table is not readable unless the management station
// has read-write permission (either implicit if the table
// is empty, or explicit through an entry in this table).
// Entries are ordered by docsDevNmAccessIndex.  The first
// matching entry (e.g., matching IP address and community
// string) is used to derive access.
type DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index used to order the application of access
    // entries. The type is interface{} with range: 1..2147483647.
    DocsDevNmAccessIndex interface{}

    // The IP address (or subnet) of the network management station. The address
    // 0.0.0.0 is defined to mean any Network Management Station (NMS).  If traps
    // are enabled for this entry, then the value must be the address of a
    // specific device.  Implementations MAY recognize 255.255.255.255 as
    // equivalent to 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevNmAccessIp interface{}

    // The IP subnet mask of the network management stations. If traps are enabled
    // for this entry, then the value must be 0.0.0.0.  Implementations MAY
    // recognize 255.255.255.255 as equivalent to 0.0.0.0. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevNmAccessIpMask interface{}

    // The community string to be matched for access by this entry.  If set to a
    // zero-length string, then any community string will match.  When read, this
    // object SHOULD return a zero-length string. The type is string.
    DocsDevNmAccessCommunity interface{}

    // Specifies the type of access allowed to this NMS. Setting this object to
    // none(1) causes the table entry to be destroyed.  Read(2) allows access by
    // 'get' and 'get-next' PDUs.  ReadWrite(3) allows access by 'set' as well. 
    // RoWithtraps(4), rwWithTraps(5), and trapsOnly(6) control distribution of
    // Trap PDUs transmitted by this device. The type is DocsDevNmAccessControl.
    DocsDevNmAccessControl interface{}

    // Specifies the set of interfaces from which requests from this NMS will be
    // accepted.  Each octet within the value of this object specifies a set of
    // eight interfaces, the first octet specifying ports 1 through 8, the second
    // octet specifying interfaces 9 through 16, etc.  Within each octet, the most
    // significant bit represents the lowest numbered interface, and the least
    // significant bit represents the highest numbered interface.  Thus, each
    // interface is represented by a single bit within the value of this object. 
    // If that bit has a value of '1' then that interface is included in the set. 
    // Note that entries in this table apply only to link-layer interfaces (e.g.,
    // Ethernet and CATV MAC).  Bits representing upstream and downstream channel
    // interfaces MUST NOT be set to '1'. Note that if bits corresponding to
    // non-existing interfaces are set, the result is implementation specific. 
    // Note that according to the DOCSIS OSSIv1.1 specification, when ifIndex '1'
    // is included in the set, then this row applies to all CPE (customer-facing)
    // interfaces.  The size of this object is the minimum required to represent
    // all configured interfaces for this device. The type is string with length:
    // 1..32.
    DocsDevNmAccessInterfaces interface{}

    // Controls and reflects the status of rows in this table. Rows in this table
    // may be created by either the create-and-go or create-and-wait paradigm.  
    // There is no restriction on changing values in a row of this table while the
    // row is active.  The following objects MUST have valid values before this
    // object can be set to active: docsDevNmAccessIp, docsDevNmAccessStatus,
    // docsDevNmAccessIpMask, docsDevNmAccessCommunity, docsDevNmAccessControl,
    // and docsDevNmAccessInterfaces. The type is RowStatus.
    DocsDevNmAccessStatus interface{}

    // Specifies the TRAP version that is sent to this NMS. Setting this object to
    // disableSNMPv2trap (1) causes the trap in SNMPv1 format to be sent to a
    // particular NMS. Setting this object to enableSNMPv2trap (2) causes the trap
    // in SNMPv2 format be sent to a particular NMS. The type is
    // DocsDevNmAccessTrapVersion.
    DocsDevNmAccessTrapVersion interface{}
}

func (docsDevNmAccessEntry *DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry) GetEntityData() *types.CommonEntityData {
    docsDevNmAccessEntry.EntityData.YFilter = docsDevNmAccessEntry.YFilter
    docsDevNmAccessEntry.EntityData.YangName = "docsDevNmAccessEntry"
    docsDevNmAccessEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevNmAccessEntry.EntityData.ParentYangName = "docsDevNmAccessTable"
    docsDevNmAccessEntry.EntityData.SegmentPath = "docsDevNmAccessEntry" + types.AddKeyToken(docsDevNmAccessEntry.DocsDevNmAccessIndex, "docsDevNmAccessIndex")
    docsDevNmAccessEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevNmAccessTable/" + docsDevNmAccessEntry.EntityData.SegmentPath
    docsDevNmAccessEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevNmAccessEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevNmAccessEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevNmAccessEntry.EntityData.Children = types.NewOrderedMap()
    docsDevNmAccessEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessIndex", types.YLeaf{"DocsDevNmAccessIndex", docsDevNmAccessEntry.DocsDevNmAccessIndex})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessIp", types.YLeaf{"DocsDevNmAccessIp", docsDevNmAccessEntry.DocsDevNmAccessIp})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessIpMask", types.YLeaf{"DocsDevNmAccessIpMask", docsDevNmAccessEntry.DocsDevNmAccessIpMask})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessCommunity", types.YLeaf{"DocsDevNmAccessCommunity", docsDevNmAccessEntry.DocsDevNmAccessCommunity})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessControl", types.YLeaf{"DocsDevNmAccessControl", docsDevNmAccessEntry.DocsDevNmAccessControl})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessInterfaces", types.YLeaf{"DocsDevNmAccessInterfaces", docsDevNmAccessEntry.DocsDevNmAccessInterfaces})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessStatus", types.YLeaf{"DocsDevNmAccessStatus", docsDevNmAccessEntry.DocsDevNmAccessStatus})
    docsDevNmAccessEntry.EntityData.Leafs.Append("docsDevNmAccessTrapVersion", types.YLeaf{"DocsDevNmAccessTrapVersion", docsDevNmAccessEntry.DocsDevNmAccessTrapVersion})

    docsDevNmAccessEntry.EntityData.YListKeys = []string {"DocsDevNmAccessIndex"}

    return &(docsDevNmAccessEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl represents device.
type DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl string

const (
    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_none DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "none"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_read DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "read"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_readWrite DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "readWrite"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_roWithTraps DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "roWithTraps"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_rwWithTraps DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "rwWithTraps"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl_trapsOnly DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessControl = "trapsOnly"
)

// DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion represents trap in SNMPv2 format be sent to a particular NMS.
type DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion string

const (
    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion_disableSNMPv2trap DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion = "disableSNMPv2trap"

    DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion_enableSNMPv2trap DOCSCABLEDEVICEMIB_DocsDevNmAccessTable_DocsDevNmAccessEntry_DocsDevNmAccessTrapVersion = "enableSNMPv2trap"
)

// DOCSCABLEDEVICEMIB_DocsDevEvControlTable
// This table allows control of the reporting of event
// classes.  For each event priority, a combination of
// logging and reporting mechanisms may be chosen.  The
// mapping of event types to priorities is
// vendor dependent.  Vendors may also choose to allow
// the user to control that mapping through proprietary
// means.  Table entries MUST persist across reboots for
// CMTS devices and MUST NOT persist across reboots for CM
// devices.
type DOCSCABLEDEVICEMIB_DocsDevEvControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allows configuration of the reporting mechanisms for a particular event
    // priority. The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry.
    DocsDevEvControlEntry []*DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry
}

func (docsDevEvControlTable *DOCSCABLEDEVICEMIB_DocsDevEvControlTable) GetEntityData() *types.CommonEntityData {
    docsDevEvControlTable.EntityData.YFilter = docsDevEvControlTable.YFilter
    docsDevEvControlTable.EntityData.YangName = "docsDevEvControlTable"
    docsDevEvControlTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevEvControlTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevEvControlTable.EntityData.SegmentPath = "docsDevEvControlTable"
    docsDevEvControlTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevEvControlTable.EntityData.SegmentPath
    docsDevEvControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevEvControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevEvControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevEvControlTable.EntityData.Children = types.NewOrderedMap()
    docsDevEvControlTable.EntityData.Children.Append("docsDevEvControlEntry", types.YChild{"DocsDevEvControlEntry", nil})
    for i := range docsDevEvControlTable.DocsDevEvControlEntry {
        docsDevEvControlTable.EntityData.Children.Append(types.GetSegmentPath(docsDevEvControlTable.DocsDevEvControlEntry[i]), types.YChild{"DocsDevEvControlEntry", docsDevEvControlTable.DocsDevEvControlEntry[i]})
    }
    docsDevEvControlTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevEvControlTable.EntityData.YListKeys = []string {}

    return &(docsDevEvControlTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry
// Allows configuration of the reporting mechanisms for a
// particular event priority.
type DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The priority level that is controlled by this
    // entry. These are ordered from most (emergency) to least (debug) critical. 
    // Each event with a CM or CMTS has a particular priority level associated
    // with it (as defined by the vendor).  emergency(1) events indicate
    // vendor-specific fatal hardware or software errors that prevent normal
    // system operation.  alert(2) events indicate a serious failure that causes
    // the reporting system to reboot but is not caused by hardware or software
    // malfunctioning.  critical(3) events indicate a serious failure that
    // requires attention and prevents the device from transmitting data but that
    // could be recovered without rebooting the system.  error(4) and warning(5)
    // events indicate that a failure occurred that could interrupt the normal
    // data flow but that does not cause the device to re-register.  notice(6) and
    // information(7) events indicate a milestone or checkpoint in normal
    // operation that could be of particular importance for troubleshooting. 
    // debug(8) events are reserved for vendor-specific events.  During normal
    // operation, no event more critical than notice(6) should be generated. 
    // Events between warning and emergency should be generated at appropriate
    // levels of problems (e.g., emergency when the box is about to crash). The
    // type is DocsDevEvPriority.
    DocsDevEvPriority interface{}

    // Defines the action to be taken on occurrence of this event class. 
    // Implementations may not necessarily support all options for all event
    // classes but at minimum must allow traps and syslogging to be disabled.  If
    // the local(0) bit is set, then log to the internal log and update
    // non-volatile store, for backward compatibility with the original RFC 2669
    // definition. If the traps(1) bit is set, then generate an SNMP trap; if the
    // syslog(2) bit is set, then send a syslog message (assuming that the syslog
    // address is set).  If the localVolatile(8) bit is set, then log to the
    // internal log without updating non-volatile store.  If the stdInterface(9)
    // bit is set, then the agent ignores all other bits except the local(0),
    // syslog(2), and localVolatile(8) bits.  Setting the stdInterface(9) bit
    // indicates that RFC3413 and RFC3014 are being used to control event
    // reporting mechanisms. The type is string.
    DocsDevEvReporting interface{}
}

func (docsDevEvControlEntry *DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry) GetEntityData() *types.CommonEntityData {
    docsDevEvControlEntry.EntityData.YFilter = docsDevEvControlEntry.YFilter
    docsDevEvControlEntry.EntityData.YangName = "docsDevEvControlEntry"
    docsDevEvControlEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevEvControlEntry.EntityData.ParentYangName = "docsDevEvControlTable"
    docsDevEvControlEntry.EntityData.SegmentPath = "docsDevEvControlEntry" + types.AddKeyToken(docsDevEvControlEntry.DocsDevEvPriority, "docsDevEvPriority")
    docsDevEvControlEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevEvControlTable/" + docsDevEvControlEntry.EntityData.SegmentPath
    docsDevEvControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevEvControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevEvControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevEvControlEntry.EntityData.Children = types.NewOrderedMap()
    docsDevEvControlEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevEvControlEntry.EntityData.Leafs.Append("docsDevEvPriority", types.YLeaf{"DocsDevEvPriority", docsDevEvControlEntry.DocsDevEvPriority})
    docsDevEvControlEntry.EntityData.Leafs.Append("docsDevEvReporting", types.YLeaf{"DocsDevEvReporting", docsDevEvControlEntry.DocsDevEvReporting})

    docsDevEvControlEntry.EntityData.YListKeys = []string {"DocsDevEvPriority"}

    return &(docsDevEvControlEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority represents box is about to crash).
type DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority string

const (
    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_emergency DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "emergency"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_alert DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "alert"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_critical DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "critical"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_error_ DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "error"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_warning DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "warning"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_notice DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "notice"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_information DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "information"

    DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority_debug DOCSCABLEDEVICEMIB_DocsDevEvControlTable_DocsDevEvControlEntry_DocsDevEvPriority = "debug"
)

// DOCSCABLEDEVICEMIB_DocsDevEventTable
// Contains a log of network and device events that may be
// of interest in fault isolation and troubleshooting.
// If the local(0) bit is set in docsDevEvReporting,
// entries in this table MUST persist across reboots.
type DOCSCABLEDEVICEMIB_DocsDevEventTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a network or device event that may be of interest in fault
    // isolation and troubleshooting. Multiple sequential identical events are
    // represented by incrementing docsDevEvCounts and setting docsDevEvLastTime
    // to the current time rather than creating multiple rows.  Entries are
    // created with the first occurrence of an event.  docsDevEvControl can be
    // used to clear the table.  Individual events cannot be deleted. The type is
    // slice of DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry.
    DocsDevEventEntry []*DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry
}

func (docsDevEventTable *DOCSCABLEDEVICEMIB_DocsDevEventTable) GetEntityData() *types.CommonEntityData {
    docsDevEventTable.EntityData.YFilter = docsDevEventTable.YFilter
    docsDevEventTable.EntityData.YangName = "docsDevEventTable"
    docsDevEventTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevEventTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevEventTable.EntityData.SegmentPath = "docsDevEventTable"
    docsDevEventTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevEventTable.EntityData.SegmentPath
    docsDevEventTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevEventTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevEventTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevEventTable.EntityData.Children = types.NewOrderedMap()
    docsDevEventTable.EntityData.Children.Append("docsDevEventEntry", types.YChild{"DocsDevEventEntry", nil})
    for i := range docsDevEventTable.DocsDevEventEntry {
        docsDevEventTable.EntityData.Children.Append(types.GetSegmentPath(docsDevEventTable.DocsDevEventEntry[i]), types.YChild{"DocsDevEventEntry", docsDevEventTable.DocsDevEventEntry[i]})
    }
    docsDevEventTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevEventTable.EntityData.YListKeys = []string {}

    return &(docsDevEventTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry
// Describes a network or device event that may be of
// interest in fault isolation and troubleshooting.
// Multiple sequential identical events are represented by
// incrementing docsDevEvCounts and setting
// docsDevEvLastTime to the current time rather than
// creating multiple rows.
// 
// Entries are created with the first occurrence of an
// event.  docsDevEvControl can be used to clear the
// table.  Individual events cannot be deleted.
type DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Provides relative ordering of the objects in the
    // event log. This object will always increase except when (a) the log is
    // reset via docsDevEvControl, (b) the device reboots and does not implement
    // non-volatile storage for this log, or (c) it reaches the value 2^31.  The
    // next entry for all the above cases is 1. The type is interface{} with
    // range: 1..2147483647.
    DocsDevEvIndex interface{}

    // The value of docsDevDateTime at the time this entry was created. The type
    // is string.
    DocsDevEvFirstTime interface{}

    // When an entry reports only one event, this object will have the same value
    // as the corresponding instance of docsDevEvFirstTime.  When an entry reports
    // multiple events, this object will record the value that docsDevDateTime had
    // when the most recent event for this entry occurred. The type is string.
    DocsDevEvLastTime interface{}

    // The number of consecutive event instances reported by this entry.  This
    // starts at 1 with the creation of this row and increments by 1 for each
    // subsequent duplicate event. The type is interface{} with range:
    // 0..4294967295. Units are events.
    DocsDevEvCounts interface{}

    // The priority level of this event, as defined by the vendor.  These are
    // ordered from most serious (emergency) to least serious (debug). 
    // emergency(1) events indicate vendor-specific fatal hardware or software
    // errors that prevent normal system operation.  alert(2) events indicate a
    // serious failure that causes the reporting system to reboot but that is not
    // caused by hardware or software malfunctioning.  critical(3) events indicate
    // a serious failure that requires attention and prevents the device from
    // transmitting data but that could be recovered without rebooting the system.
    // error(4) and warning(5) events indicate that a failure occurred that could
    // interrupt the normal data flow but that does not cause the device to
    // re-register.  notice(6) and information(7) events indicate a milestone or
    // checkpoint in normal operation that could be of particular importance for
    // troubleshooting.  debug(8) events are reserved for vendor-specific events. 
    // During normal operation, no event more critical than notice(6) should be
    // generated.  Events between warning and emergency should be generated at
    // appropriate levels of problems (e.g., emergency when the box is about to
    // crash). The type is DocsDevEvLevel.
    DocsDevEvLevel interface{}

    // For this product, uniquely identifies the type of event that is reported by
    // this entry. The type is interface{} with range: 0..4294967295.
    DocsDevEvId interface{}

    // Provides a human-readable description of the event, including all relevant
    // context (interface numbers, etc.). The type is string.
    DocsDevEvText interface{}
}

func (docsDevEventEntry *DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry) GetEntityData() *types.CommonEntityData {
    docsDevEventEntry.EntityData.YFilter = docsDevEventEntry.YFilter
    docsDevEventEntry.EntityData.YangName = "docsDevEventEntry"
    docsDevEventEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevEventEntry.EntityData.ParentYangName = "docsDevEventTable"
    docsDevEventEntry.EntityData.SegmentPath = "docsDevEventEntry" + types.AddKeyToken(docsDevEventEntry.DocsDevEvIndex, "docsDevEvIndex")
    docsDevEventEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevEventTable/" + docsDevEventEntry.EntityData.SegmentPath
    docsDevEventEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevEventEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevEventEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevEventEntry.EntityData.Children = types.NewOrderedMap()
    docsDevEventEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvIndex", types.YLeaf{"DocsDevEvIndex", docsDevEventEntry.DocsDevEvIndex})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvFirstTime", types.YLeaf{"DocsDevEvFirstTime", docsDevEventEntry.DocsDevEvFirstTime})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvLastTime", types.YLeaf{"DocsDevEvLastTime", docsDevEventEntry.DocsDevEvLastTime})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvCounts", types.YLeaf{"DocsDevEvCounts", docsDevEventEntry.DocsDevEvCounts})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvLevel", types.YLeaf{"DocsDevEvLevel", docsDevEventEntry.DocsDevEvLevel})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvId", types.YLeaf{"DocsDevEvId", docsDevEventEntry.DocsDevEvId})
    docsDevEventEntry.EntityData.Leafs.Append("docsDevEvText", types.YLeaf{"DocsDevEvText", docsDevEventEntry.DocsDevEvText})

    docsDevEventEntry.EntityData.YListKeys = []string {"DocsDevEvIndex"}

    return &(docsDevEventEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel represents box is about to crash).
type DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel string

const (
    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_emergency DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "emergency"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_alert DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "alert"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_critical DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "critical"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_error_ DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "error"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_warning DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "warning"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_notice DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "notice"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_information DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "information"

    DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel_debug DOCSCABLEDEVICEMIB_DocsDevEventTable_DocsDevEventEntry_DocsDevEvLevel = "debug"
)

// DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable
// A list of filters to apply to (bridged) LLC
// traffic. The filters in this table are applied to
// incoming traffic on the appropriate interface(s)  prior
// to any further processing (e.g., before the packet
// is handed off for level 3 processing, or for bridging).
// The specific action taken when no filter is matched is
// controlled by docsDevFilterLLCUnmatchedAction.  Table
// entries MUST NOT persist across reboots for any device.
type DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a single filter to apply to (bridged) LLC traffic received on a
    // specified interface. . The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry.
    DocsDevFilterLLCEntry []*DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry
}

func (docsDevFilterLLCTable *DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable) GetEntityData() *types.CommonEntityData {
    docsDevFilterLLCTable.EntityData.YFilter = docsDevFilterLLCTable.YFilter
    docsDevFilterLLCTable.EntityData.YangName = "docsDevFilterLLCTable"
    docsDevFilterLLCTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterLLCTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevFilterLLCTable.EntityData.SegmentPath = "docsDevFilterLLCTable"
    docsDevFilterLLCTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevFilterLLCTable.EntityData.SegmentPath
    docsDevFilterLLCTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterLLCTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterLLCTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterLLCTable.EntityData.Children = types.NewOrderedMap()
    docsDevFilterLLCTable.EntityData.Children.Append("docsDevFilterLLCEntry", types.YChild{"DocsDevFilterLLCEntry", nil})
    for i := range docsDevFilterLLCTable.DocsDevFilterLLCEntry {
        docsDevFilterLLCTable.EntityData.Children.Append(types.GetSegmentPath(docsDevFilterLLCTable.DocsDevFilterLLCEntry[i]), types.YChild{"DocsDevFilterLLCEntry", docsDevFilterLLCTable.DocsDevFilterLLCEntry[i]})
    }
    docsDevFilterLLCTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevFilterLLCTable.EntityData.YListKeys = []string {}

    return &(docsDevFilterLLCTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry
// Describes a single filter to apply to (bridged) LLC
// traffic received on a specified interface. 
type DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index used for the identification of filters (note
    // that LLC filter order is irrelevant). The type is interface{} with range:
    // 1..2147483647.
    DocsDevFilterLLCIndex interface{}

    // Controls and reflects the status of rows in this table. There is no
    // restriction on changing any of the associated columns for this row while
    // this object is set to active.  Specifying only this object (with the
    // appropriate index) on a CM is sufficient to create a filter row that
    // matches all inbound packets on the ethernet interface and results in the
    // packets being discarded.  docsDevFilterLLCIfIndex (at least) must be
    // specified on a CMTS to create a row. The type is RowStatus.
    DocsDevFilterLLCStatus interface{}

    // The entry interface to which this filter applies.  The value corresponds to
    // ifIndex for either a CATV MAC or another network interface.  If the value
    // is zero, the filter applies to all interfaces.  In Cable Modems, the
    // default value is the customer side interface(s).  In CMTSs, this object has
    // to be specified to create a row in this table.  Note that according to the
    // DOCSIS OSSIv1.1 specification, ifIndex '1' in the CM means that this row
    // applies to all Cable Modem-to-CPE Interfaces (CMCI). The type is
    // interface{} with range: 0..2147483647.
    DocsDevFilterLLCIfIndex interface{}

    // The format of the value in docsDevFilterLLCProtocol: either a two-byte
    // Ethernet Ethertype, or a one-byte 802.2 Service Access Point (SAP) value. 
    // ethertype(1) also applies to Standard Network Access Protocol (SNAP)
    // encapsulated frames. The type is DocsDevFilterLLCProtocolType.
    DocsDevFilterLLCProtocolType interface{}

    // The layer-three protocol for which this filter applies. The protocol value
    // format depends on docsDevFilterLLCProtocolType.  Note that for SNAP frames,
    // ethertype filtering is performed rather than Destination Service Access
    // Point (DSAP) =0xAA. The type is interface{} with range: 0..65535.
    DocsDevFilterLLCProtocol interface{}

    // Counts the number of times this filter was matched. The type is interface{}
    // with range: 0..4294967295. Units are matches.
    DocsDevFilterLLCMatches interface{}
}

func (docsDevFilterLLCEntry *DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry) GetEntityData() *types.CommonEntityData {
    docsDevFilterLLCEntry.EntityData.YFilter = docsDevFilterLLCEntry.YFilter
    docsDevFilterLLCEntry.EntityData.YangName = "docsDevFilterLLCEntry"
    docsDevFilterLLCEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterLLCEntry.EntityData.ParentYangName = "docsDevFilterLLCTable"
    docsDevFilterLLCEntry.EntityData.SegmentPath = "docsDevFilterLLCEntry" + types.AddKeyToken(docsDevFilterLLCEntry.DocsDevFilterLLCIndex, "docsDevFilterLLCIndex")
    docsDevFilterLLCEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevFilterLLCTable/" + docsDevFilterLLCEntry.EntityData.SegmentPath
    docsDevFilterLLCEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterLLCEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterLLCEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterLLCEntry.EntityData.Children = types.NewOrderedMap()
    docsDevFilterLLCEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCIndex", types.YLeaf{"DocsDevFilterLLCIndex", docsDevFilterLLCEntry.DocsDevFilterLLCIndex})
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCStatus", types.YLeaf{"DocsDevFilterLLCStatus", docsDevFilterLLCEntry.DocsDevFilterLLCStatus})
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCIfIndex", types.YLeaf{"DocsDevFilterLLCIfIndex", docsDevFilterLLCEntry.DocsDevFilterLLCIfIndex})
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCProtocolType", types.YLeaf{"DocsDevFilterLLCProtocolType", docsDevFilterLLCEntry.DocsDevFilterLLCProtocolType})
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCProtocol", types.YLeaf{"DocsDevFilterLLCProtocol", docsDevFilterLLCEntry.DocsDevFilterLLCProtocol})
    docsDevFilterLLCEntry.EntityData.Leafs.Append("docsDevFilterLLCMatches", types.YLeaf{"DocsDevFilterLLCMatches", docsDevFilterLLCEntry.DocsDevFilterLLCMatches})

    docsDevFilterLLCEntry.EntityData.YListKeys = []string {"DocsDevFilterLLCIndex"}

    return &(docsDevFilterLLCEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType represents (SNAP) encapsulated frames.
type DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType string

const (
    DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType_ethertype DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType = "ethertype"

    DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType_dsap DOCSCABLEDEVICEMIB_DocsDevFilterLLCTable_DocsDevFilterLLCEntry_DocsDevFilterLLCProtocolType = "dsap"
)

// DOCSCABLEDEVICEMIB_DocsDevFilterIpTable
// An ordered list of filters or classifiers to apply to
// IP traffic. Filter application is ordered by the filter
// index, rather than by a best match algorithm (Note that
// this implies that the filter table may have gaps in the
// index values).  Packets that match no filters will have
// policy 0 in the docsDevFilterPolicyTable applied to
// them, if it exists.  Otherwise, Packets that match no
// filters are discarded or forwarded according to the
// setting of docsDevFilterIpDefault.
// 
// Any IP packet can theoretically match multiple rows of
// this table.  When considering a packet, the table is
// scanned in row index order (e.g., filter 10 is checked
// before filter 20).  If the packet matches that filter
// (which means that it matches ALL criteria for that row),
// actions appropriate to docsDevFilterIpControl and
// docsDevFilterPolicyId are taken.  If the packet was
// discarded processing is complete.  If
// docsDevFilterIpContinue is set to true, the filter
// comparison continues with the next row in the table,
// looking for additional matches.
// 
// If the packet matches no filter in the table, the packet
// is accepted or dropped for further processing
// according to the setting of docsDevFilterIpDefault.
// If the packet is accepted, the actions specified by
// policy group 0 (e.g., the rows in
// docsDevFilterPolicyTable that have a value of 0 for
// docsDevFilterPolicyId) are taken, if that policy
// group exists.
// 
// Logically, this table is consulted twice during the
// processing of any IP packet: once upon its acceptance
// from the L2 entity, and once upon its transmission to
// the L2 entity.  In actuality, for cable modems, IP
// filtering is generally the only IP processing done for
// transit traffic.  This means that inbound and outbound
// filtering can generally be done at the same time with
// one pass through the filter table.
// 
// The objects in this table are only accessible from cable
// devices that are not operating in DiffServ MIB mode
// (RFC 3289).  See the conformance section for details.
// 
// Note that some devices are required by other
// specifications (e.g., the DOCSIS OSSIv1.1 specification)
// to support the legacy SNMPv1/v2c docsDevFilter mode
// for backward compatibility.
// 
// Table entries MUST NOT persist across reboots for any
// device.
// 
// This table is deprecated.  Instead, use the DiffServ MIB
// from RFC3289.
type DOCSCABLEDEVICEMIB_DocsDevFilterIpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a filter to apply to IP traffic received on a specified
    // interface.  All identity objects in this table (e.g., source and
    // destination address/mask, protocol, source/dest port, TOS/mask, interface
    // and direction) must match their respective fields in the packet for any
    // given filter to match.  To create an entry in this table,
    // docsDevFilterIpIfIndex must be specified. The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry.
    DocsDevFilterIpEntry []*DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry
}

func (docsDevFilterIpTable *DOCSCABLEDEVICEMIB_DocsDevFilterIpTable) GetEntityData() *types.CommonEntityData {
    docsDevFilterIpTable.EntityData.YFilter = docsDevFilterIpTable.YFilter
    docsDevFilterIpTable.EntityData.YangName = "docsDevFilterIpTable"
    docsDevFilterIpTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterIpTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevFilterIpTable.EntityData.SegmentPath = "docsDevFilterIpTable"
    docsDevFilterIpTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevFilterIpTable.EntityData.SegmentPath
    docsDevFilterIpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterIpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterIpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterIpTable.EntityData.Children = types.NewOrderedMap()
    docsDevFilterIpTable.EntityData.Children.Append("docsDevFilterIpEntry", types.YChild{"DocsDevFilterIpEntry", nil})
    for i := range docsDevFilterIpTable.DocsDevFilterIpEntry {
        docsDevFilterIpTable.EntityData.Children.Append(types.GetSegmentPath(docsDevFilterIpTable.DocsDevFilterIpEntry[i]), types.YChild{"DocsDevFilterIpEntry", docsDevFilterIpTable.DocsDevFilterIpEntry[i]})
    }
    docsDevFilterIpTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevFilterIpTable.EntityData.YListKeys = []string {}

    return &(docsDevFilterIpTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry
// Describes a filter to apply to IP traffic received on a
// specified interface.  All identity objects in this table
// (e.g., source and destination address/mask, protocol,
// source/dest port, TOS/mask, interface and direction)
// must match their respective fields in the packet for
// any given filter to match.
// 
// To create an entry in this table, docsDevFilterIpIfIndex
// must be specified.
type DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index used to order the application of filters.
    // The filter with the lowest index is always applied first. The type is
    // interface{} with range: 1..2147483647.
    DocsDevFilterIpIndex interface{}

    // Controls and reflects the status of rows in this table.  Specifying only
    // this object (with the appropriate index) on a CM is sufficient to create a
    // filter row that matches all inbound packets on the ethernet interface and
    // results in the packets being discarded. docsDevFilterIpIfIndex (at least)
    // must be specified on a CMTS to create a row.  Creation of the rows may be
    // done via either create-and-wait or create-and-go, but the filter is not
    // applied until this object is set to (or changes to) active. There is no
    // restriction in changing any object in a row while this object is set to
    // active. The type is RowStatus.
    DocsDevFilterIpStatus interface{}

    // If set to discard(1), all packets matching this filter will be discarded,
    // and scanning of the remainder of the filter list will be aborted. If set to
    // accept(2), all packets matching this filter will be accepted for further
    // processing (e.g., bridging).  If docsDevFilterIpContinue is set to true,
    // see if there are other matches; otherwise, done.  If set to policy (3),
    // execute the policy entries matched by docsDevFilterIpPolicyId in
    // docsDevFilterPolicyTable.  If docsDevFilterIpContinue is set to true,
    // continue scanning the table for other matches; otherwise, done. The type is
    // DocsDevFilterIpControl.
    DocsDevFilterIpControl interface{}

    // The entry interface to which this filter applies. The value corresponds to
    // ifIndex for either a CATV MAC or another interface. If the value is zero,
    // the filter applies to all interfaces.  Default value in CMs is the index of
    // the customer-side (e.g., ethernet) interface(s).  In CMTSes, this object
    // MUST be specified to create a row in this table.  Note that according to
    // the DOCSIS OSSIv1.1 specification, ifIndex '1' in the Cable Modem means
    // that this row applies to all CMCI (customer-facing) interfaces. The type is
    // interface{} with range: 0..2147483647.
    DocsDevFilterIpIfIndex interface{}

    // Determines whether the filter is applied to inbound(1) traffic, outbound(2)
    // traffic, or traffic in both(3) directions. The type is
    // DocsDevFilterIpDirection.
    DocsDevFilterIpDirection interface{}

    // If set to true(1), the filter only applies to multicast and broadcast
    // traffic. If set to false(2), the filter applies to all traffic. The type is
    // bool.
    DocsDevFilterIpBroadcast interface{}

    // The source IP address, or portion thereof, that is to be matched for this
    // filter.  The source address is first masked (ANDed) against
    // docsDevFilterIpSmask before being compared to this value.  A value of 0 for
    // this object and 0 for the mask matches all IP addresses. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevFilterIpSaddr interface{}

    // A bit mask that is to be applied to the source address prior to matching. 
    // This mask is not necessarily the same as a subnet mask, but 1s bits must be
    // leftmost and contiguous. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevFilterIpSmask interface{}

    // The destination IP address, or portion thereof, that is to be matched for
    // this filter.  The destination address is first masked (ANDed) against
    // docsDevFilterIpDmask before being compared to this value.  A value of
    // 00000000 for this object and 00000000 for the mask matches all IP
    // addresses. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevFilterIpDaddr interface{}

    // A bit mask that is to be applied to the destination address prior to
    // matching. This mask is not necessarily the same as a subnet mask, but 1s
    // bits MUST be leftmost and contiguous. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevFilterIpDmask interface{}

    // The IP protocol value that is to be matched.  For example, icmp is 1, tcp
    // is 6, and udp is 17.  A value of 256 matches ANY protocol. The type is
    // interface{} with range: 0..256.
    DocsDevFilterIpProtocol interface{}

    // This is the inclusive lower bound of the transport-layer source port range
    // that is to be matched.  If the IP protocol of the packet is neither UDP nor
    // TCP, this object is ignored during matching. The type is interface{} with
    // range: 0..65535.
    DocsDevFilterIpSourcePortLow interface{}

    // This is the inclusive upper bound of the transport-layer source port range
    // that is to be matched.  If the IP protocol of the packet is neither UDP nor
    // TCP, this object is ignored during matching. The type is interface{} with
    // range: 0..65535.
    DocsDevFilterIpSourcePortHigh interface{}

    // This is the inclusive lower bound of the transport-layer destination port
    // range that is to be matched. If the IP protocol of the packet is neither
    // UDP nor TCP, this object is ignored during matching. The type is
    // interface{} with range: 0..65535.
    DocsDevFilterIpDestPortLow interface{}

    // This is the inclusive upper bound of the transport-layer destination port
    // range that is to be matched. If the IP protocol of the packet is neither
    // UDP nor TCP, this object is ignored during matching. The type is
    // interface{} with range: 0..65535.
    DocsDevFilterIpDestPortHigh interface{}

    // Counts the number of times this filter was matched. This object is
    // initialized to 0 at boot, or at row creation, and is reset only upon
    // reboot. The type is interface{} with range: 0..4294967295. Units are
    // matches.
    DocsDevFilterIpMatches interface{}

    // This is the value to be matched to the packet's TOS (Type of Service) value
    // (after the TOS value is ANDed with docsDevFilterIpTosMask).  A value for
    // this object of 0 and a mask of 0 matches all TOS values. The type is string
    // with length: 1.
    DocsDevFilterIpTos interface{}

    // The mask to be applied to the packet's TOS value before matching. The type
    // is string with length: 1.
    DocsDevFilterIpTosMask interface{}

    // If this value is set to true and docsDevFilterIpControl is anything but
    // discard (1), continue scanning and applying policies.  See Section 3.3.3
    // for more details. The type is bool.
    DocsDevFilterIpContinue interface{}

    // This object points to an entry in docsDevFilterPolicyTable.  If
    // docsDevFilterIpControl is set to policy (3), execute all matching policies
    // in docsDevFilterPolicyTable.  If no matching policy exists, treat as if
    // docsDevFilterIpControl were set to accept (1).  If this object is set to
    // the value of 0, there is no matching policy, and docsDevFilterPolicyTable
    // MUST NOT be consulted. The type is interface{} with range: 0..2147483647.
    DocsDevFilterIpPolicyId interface{}
}

func (docsDevFilterIpEntry *DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry) GetEntityData() *types.CommonEntityData {
    docsDevFilterIpEntry.EntityData.YFilter = docsDevFilterIpEntry.YFilter
    docsDevFilterIpEntry.EntityData.YangName = "docsDevFilterIpEntry"
    docsDevFilterIpEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterIpEntry.EntityData.ParentYangName = "docsDevFilterIpTable"
    docsDevFilterIpEntry.EntityData.SegmentPath = "docsDevFilterIpEntry" + types.AddKeyToken(docsDevFilterIpEntry.DocsDevFilterIpIndex, "docsDevFilterIpIndex")
    docsDevFilterIpEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevFilterIpTable/" + docsDevFilterIpEntry.EntityData.SegmentPath
    docsDevFilterIpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterIpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterIpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterIpEntry.EntityData.Children = types.NewOrderedMap()
    docsDevFilterIpEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpIndex", types.YLeaf{"DocsDevFilterIpIndex", docsDevFilterIpEntry.DocsDevFilterIpIndex})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpStatus", types.YLeaf{"DocsDevFilterIpStatus", docsDevFilterIpEntry.DocsDevFilterIpStatus})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpControl", types.YLeaf{"DocsDevFilterIpControl", docsDevFilterIpEntry.DocsDevFilterIpControl})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpIfIndex", types.YLeaf{"DocsDevFilterIpIfIndex", docsDevFilterIpEntry.DocsDevFilterIpIfIndex})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpDirection", types.YLeaf{"DocsDevFilterIpDirection", docsDevFilterIpEntry.DocsDevFilterIpDirection})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpBroadcast", types.YLeaf{"DocsDevFilterIpBroadcast", docsDevFilterIpEntry.DocsDevFilterIpBroadcast})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpSaddr", types.YLeaf{"DocsDevFilterIpSaddr", docsDevFilterIpEntry.DocsDevFilterIpSaddr})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpSmask", types.YLeaf{"DocsDevFilterIpSmask", docsDevFilterIpEntry.DocsDevFilterIpSmask})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpDaddr", types.YLeaf{"DocsDevFilterIpDaddr", docsDevFilterIpEntry.DocsDevFilterIpDaddr})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpDmask", types.YLeaf{"DocsDevFilterIpDmask", docsDevFilterIpEntry.DocsDevFilterIpDmask})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpProtocol", types.YLeaf{"DocsDevFilterIpProtocol", docsDevFilterIpEntry.DocsDevFilterIpProtocol})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpSourcePortLow", types.YLeaf{"DocsDevFilterIpSourcePortLow", docsDevFilterIpEntry.DocsDevFilterIpSourcePortLow})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpSourcePortHigh", types.YLeaf{"DocsDevFilterIpSourcePortHigh", docsDevFilterIpEntry.DocsDevFilterIpSourcePortHigh})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpDestPortLow", types.YLeaf{"DocsDevFilterIpDestPortLow", docsDevFilterIpEntry.DocsDevFilterIpDestPortLow})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpDestPortHigh", types.YLeaf{"DocsDevFilterIpDestPortHigh", docsDevFilterIpEntry.DocsDevFilterIpDestPortHigh})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpMatches", types.YLeaf{"DocsDevFilterIpMatches", docsDevFilterIpEntry.DocsDevFilterIpMatches})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpTos", types.YLeaf{"DocsDevFilterIpTos", docsDevFilterIpEntry.DocsDevFilterIpTos})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpTosMask", types.YLeaf{"DocsDevFilterIpTosMask", docsDevFilterIpEntry.DocsDevFilterIpTosMask})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpContinue", types.YLeaf{"DocsDevFilterIpContinue", docsDevFilterIpEntry.DocsDevFilterIpContinue})
    docsDevFilterIpEntry.EntityData.Leafs.Append("docsDevFilterIpPolicyId", types.YLeaf{"DocsDevFilterIpPolicyId", docsDevFilterIpEntry.DocsDevFilterIpPolicyId})

    docsDevFilterIpEntry.EntityData.YListKeys = []string {"DocsDevFilterIpIndex"}

    return &(docsDevFilterIpEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl represents scanning the table for other matches; otherwise, done.
type DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl string

const (
    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl_discard DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl = "discard"

    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl_accept DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl = "accept"

    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl_policy DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpControl = "policy"
)

// DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection represents directions.
type DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection string

const (
    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection_inbound DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection = "inbound"

    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection_outbound DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection = "outbound"

    DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection_both DOCSCABLEDEVICEMIB_DocsDevFilterIpTable_DocsDevFilterIpEntry_DocsDevFilterIpDirection = "both"
)

// DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable
// A Table that maps between a policy group ID and a set
// of pointers to policies to be applied.  All rows with
// the same docsDevFilterPolicyId are part of the same
// group of policy pointers and are applied in the order
// in this table.  docsDevFilterPolicyTable exists to
// allow multiple policy actions (referenced by policy
// pointers) to be applied to any given classified packet.
// The policy actions are applied in index order.
// For example:
// 
// Index   ID    Type    Action
//  1      1      TOS     1
//  9      5      TOS     1
//  12     1      IPSEC   3
// 
// This says that a packet that matches a filter with
// policy id 1 first has TOS policy 1 applied (which might
// set the TOS bits to enable a higher priority) and next
// has the IPSEC policy 3 applied (which may result in the
// packets being dumped into a secure VPN to a remote
// encryptor).
// 
// Policy ID 0 is reserved for default actions and is
// applied only to packets that match no filters in
// docsDevFilterIpTable.
// 
// Table entries MUST NOT persist across reboots for any
// device.
// 
// This table is deprecated.  Instead, use the DiffServ MIB
// from RFC3289.
type DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the docsDevFilterPolicyTable. Entries are created by Network
    // Management. To create an entry, docsDevFilterPolicyId MUST be specified.
    // The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable_DocsDevFilterPolicyEntry.
    DocsDevFilterPolicyEntry []*DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable_DocsDevFilterPolicyEntry
}

func (docsDevFilterPolicyTable *DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable) GetEntityData() *types.CommonEntityData {
    docsDevFilterPolicyTable.EntityData.YFilter = docsDevFilterPolicyTable.YFilter
    docsDevFilterPolicyTable.EntityData.YangName = "docsDevFilterPolicyTable"
    docsDevFilterPolicyTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterPolicyTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevFilterPolicyTable.EntityData.SegmentPath = "docsDevFilterPolicyTable"
    docsDevFilterPolicyTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevFilterPolicyTable.EntityData.SegmentPath
    docsDevFilterPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterPolicyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterPolicyTable.EntityData.Children = types.NewOrderedMap()
    docsDevFilterPolicyTable.EntityData.Children.Append("docsDevFilterPolicyEntry", types.YChild{"DocsDevFilterPolicyEntry", nil})
    for i := range docsDevFilterPolicyTable.DocsDevFilterPolicyEntry {
        docsDevFilterPolicyTable.EntityData.Children.Append(types.GetSegmentPath(docsDevFilterPolicyTable.DocsDevFilterPolicyEntry[i]), types.YChild{"DocsDevFilterPolicyEntry", docsDevFilterPolicyTable.DocsDevFilterPolicyEntry[i]})
    }
    docsDevFilterPolicyTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevFilterPolicyTable.EntityData.YListKeys = []string {}

    return &(docsDevFilterPolicyTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable_DocsDevFilterPolicyEntry
// An entry in the docsDevFilterPolicyTable. Entries are
// created by Network Management. To create an entry,
// docsDevFilterPolicyId MUST be specified.
type DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable_DocsDevFilterPolicyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index value for the table. The type is interface{}
    // with range: 1..2147483647.
    DocsDevFilterPolicyIndex interface{}

    // Policy ID for this entry.  If a policy ID can apply to multiple rows of
    // this table, all relevant policies are executed. Policy 0 (if populated) is
    // applied to all packets that do not match any of the filters.  N.B. If
    // docsDevFilterIpPolicyId is set to 0, it DOES NOT match policy 0 of this
    // table. The type is interface{} with range: 0..2147483647.
    DocsDevFilterPolicyId interface{}

    // Object used to create an entry in this table.  There is no restriction in
    // changing any object in a row while this object is set to active. The
    // following object MUST have a valid value before this object can be set to
    // active:  docsDevFilterPolicyPtr. The type is RowStatus.
    DocsDevFilterPolicyStatus interface{}

    // This object points to a row in an applicable filter policy table. 
    // Currently, the only standard policy table is docsDevFilterTosTable.  Per
    // the textual convention, this object points to the first accessible object
    // in the row; e.g., to point to a row in docsDevFilterTosTable with an index
    // of 21, the value of this object would be the object identifier
    // docsDevTosStatus.21.  Vendors are recommended to adhere to the same
    // convention when adding vendor-specific policy table extensions.  If this
    // pointer references an empty or non-existent row, then no policy action is
    // taken.  The default upon row creation is a null pointer that results in no
    // policy action being taken. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    DocsDevFilterPolicyPtr interface{}
}

func (docsDevFilterPolicyEntry *DOCSCABLEDEVICEMIB_DocsDevFilterPolicyTable_DocsDevFilterPolicyEntry) GetEntityData() *types.CommonEntityData {
    docsDevFilterPolicyEntry.EntityData.YFilter = docsDevFilterPolicyEntry.YFilter
    docsDevFilterPolicyEntry.EntityData.YangName = "docsDevFilterPolicyEntry"
    docsDevFilterPolicyEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterPolicyEntry.EntityData.ParentYangName = "docsDevFilterPolicyTable"
    docsDevFilterPolicyEntry.EntityData.SegmentPath = "docsDevFilterPolicyEntry" + types.AddKeyToken(docsDevFilterPolicyEntry.DocsDevFilterPolicyIndex, "docsDevFilterPolicyIndex")
    docsDevFilterPolicyEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevFilterPolicyTable/" + docsDevFilterPolicyEntry.EntityData.SegmentPath
    docsDevFilterPolicyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterPolicyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterPolicyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterPolicyEntry.EntityData.Children = types.NewOrderedMap()
    docsDevFilterPolicyEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevFilterPolicyEntry.EntityData.Leafs.Append("docsDevFilterPolicyIndex", types.YLeaf{"DocsDevFilterPolicyIndex", docsDevFilterPolicyEntry.DocsDevFilterPolicyIndex})
    docsDevFilterPolicyEntry.EntityData.Leafs.Append("docsDevFilterPolicyId", types.YLeaf{"DocsDevFilterPolicyId", docsDevFilterPolicyEntry.DocsDevFilterPolicyId})
    docsDevFilterPolicyEntry.EntityData.Leafs.Append("docsDevFilterPolicyStatus", types.YLeaf{"DocsDevFilterPolicyStatus", docsDevFilterPolicyEntry.DocsDevFilterPolicyStatus})
    docsDevFilterPolicyEntry.EntityData.Leafs.Append("docsDevFilterPolicyPtr", types.YLeaf{"DocsDevFilterPolicyPtr", docsDevFilterPolicyEntry.DocsDevFilterPolicyPtr})

    docsDevFilterPolicyEntry.EntityData.YListKeys = []string {"DocsDevFilterPolicyIndex"}

    return &(docsDevFilterPolicyEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterTosTable
// Table used to describe Type of Service (TOS) bits
// processing.
// 
// This table is an adjunct to the docsDevFilterIpTable
// and the docsDevFilterPolicy table.  Entries in the
// latter table can point to specific rows in this (and
// other) tables and cause specific actions to be taken.
// This table permits the manipulation of the value of the
// Type of Service bits in the IP header of the matched
// packet as follows:
// 
// Set the tosBits of the packet to
//    (tosBits & docsDevFilterTosAndMask) |
//    docsDevFilterTosOrMask
// 
// This construct allows you to do a clear and set of all
// the TOS bits in a flexible manner.
// 
// Table entries MUST NOT persist across reboots for any
// device.
// 
// This table is deprecated.  Instead, use the DiffServ MIB
// from RFC3289.
type DOCSCABLEDEVICEMIB_DocsDevFilterTosTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A TOS policy entry. The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevFilterTosTable_DocsDevFilterTosEntry.
    DocsDevFilterTosEntry []*DOCSCABLEDEVICEMIB_DocsDevFilterTosTable_DocsDevFilterTosEntry
}

func (docsDevFilterTosTable *DOCSCABLEDEVICEMIB_DocsDevFilterTosTable) GetEntityData() *types.CommonEntityData {
    docsDevFilterTosTable.EntityData.YFilter = docsDevFilterTosTable.YFilter
    docsDevFilterTosTable.EntityData.YangName = "docsDevFilterTosTable"
    docsDevFilterTosTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterTosTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevFilterTosTable.EntityData.SegmentPath = "docsDevFilterTosTable"
    docsDevFilterTosTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevFilterTosTable.EntityData.SegmentPath
    docsDevFilterTosTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterTosTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterTosTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterTosTable.EntityData.Children = types.NewOrderedMap()
    docsDevFilterTosTable.EntityData.Children.Append("docsDevFilterTosEntry", types.YChild{"DocsDevFilterTosEntry", nil})
    for i := range docsDevFilterTosTable.DocsDevFilterTosEntry {
        docsDevFilterTosTable.EntityData.Children.Append(types.GetSegmentPath(docsDevFilterTosTable.DocsDevFilterTosEntry[i]), types.YChild{"DocsDevFilterTosEntry", docsDevFilterTosTable.DocsDevFilterTosEntry[i]})
    }
    docsDevFilterTosTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevFilterTosTable.EntityData.YListKeys = []string {}

    return &(docsDevFilterTosTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevFilterTosTable_DocsDevFilterTosEntry
// A TOS policy entry.
type DOCSCABLEDEVICEMIB_DocsDevFilterTosTable_DocsDevFilterTosEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The unique index for this row.  There are no
    // ordering requirements for this table, and any valid index may be specified.
    // The type is interface{} with range: 1..2147483647.
    DocsDevFilterTosIndex interface{}

    // The object used to create and delete entries in this table. A row created
    // by specifying just this object results in a row that specifies no change to
    // the TOS bits.  A row may be created using either the create-and-go or
    // create-and-wait paradigms.  There is no restriction on the ability to
    // change values in this row while the row is active. The type is RowStatus.
    DocsDevFilterTosStatus interface{}

    // This value is bitwise ANDed with the matched packet's TOS bits. The type is
    // string with length: 1.
    DocsDevFilterTosAndMask interface{}

    // This value is bitwise ORed with the result from the AND procedure (tosBits
    // & docsDevFilterTosAndMask). The result then replaces the packet's TOS bits.
    // The type is string with length: 1.
    DocsDevFilterTosOrMask interface{}
}

func (docsDevFilterTosEntry *DOCSCABLEDEVICEMIB_DocsDevFilterTosTable_DocsDevFilterTosEntry) GetEntityData() *types.CommonEntityData {
    docsDevFilterTosEntry.EntityData.YFilter = docsDevFilterTosEntry.YFilter
    docsDevFilterTosEntry.EntityData.YangName = "docsDevFilterTosEntry"
    docsDevFilterTosEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevFilterTosEntry.EntityData.ParentYangName = "docsDevFilterTosTable"
    docsDevFilterTosEntry.EntityData.SegmentPath = "docsDevFilterTosEntry" + types.AddKeyToken(docsDevFilterTosEntry.DocsDevFilterTosIndex, "docsDevFilterTosIndex")
    docsDevFilterTosEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevFilterTosTable/" + docsDevFilterTosEntry.EntityData.SegmentPath
    docsDevFilterTosEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevFilterTosEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevFilterTosEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevFilterTosEntry.EntityData.Children = types.NewOrderedMap()
    docsDevFilterTosEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevFilterTosEntry.EntityData.Leafs.Append("docsDevFilterTosIndex", types.YLeaf{"DocsDevFilterTosIndex", docsDevFilterTosEntry.DocsDevFilterTosIndex})
    docsDevFilterTosEntry.EntityData.Leafs.Append("docsDevFilterTosStatus", types.YLeaf{"DocsDevFilterTosStatus", docsDevFilterTosEntry.DocsDevFilterTosStatus})
    docsDevFilterTosEntry.EntityData.Leafs.Append("docsDevFilterTosAndMask", types.YLeaf{"DocsDevFilterTosAndMask", docsDevFilterTosEntry.DocsDevFilterTosAndMask})
    docsDevFilterTosEntry.EntityData.Leafs.Append("docsDevFilterTosOrMask", types.YLeaf{"DocsDevFilterTosOrMask", docsDevFilterTosEntry.DocsDevFilterTosOrMask})

    docsDevFilterTosEntry.EntityData.YListKeys = []string {"DocsDevFilterTosIndex"}

    return &(docsDevFilterTosEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpeTable
// This table lists the IPv4 addresses seen (or permitted)
// as source addresses in packets originating from the
// customer interface on this device.  In addition, this
// table can be provisioned with the specific addresses
// permitted for the CPEs via the normal row creation
// mechanisms.  Table entries MUST NOT persist across
// reboots for any device.
// 
// N.B.  Management action can add entries in this table
// and in docsDevCpeIpTable past the value of
// docsDevCpeIpMax.  docsDevCpeIpMax ONLY restricts the
// ability of the CM to add learned addresses
// automatically.
// 
// This table is deprecated and is replaced by
// docsDevCpeInetTable.
type DOCSCABLEDEVICEMIB_DocsDevCpeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the docsDevFilterCpeTable.  There is one entry for each IPv4
    // CPE seen or provisioned.  If docsDevCpeIpMax is set to -1, this table is
    // ignored; otherwise, upon receipt of an IP packet from the customer
    // interface of the CM, the source IP address is checked against this table. 
    // If the address is in the table, packet processing continues.  If the
    // address is not in the table but docsDevCpeEnroll is set to any and the sum
    // of the table sizes of docsDevCpeTable and docsDevCpeInetTable is less than
    // docsDevCpeIpMax, the address is added to the table, and packet processing
    // continues.  Otherwise, the packet is dropped.  The filtering actions
    // specified by this table occur after any LLC filtering
    // (docsDevFilterLLCTable), but prior to any IP filtering
    // (docsDevFilterIpTable, docsDevNmAccessTable). The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry.
    DocsDevCpeEntry []*DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry
}

func (docsDevCpeTable *DOCSCABLEDEVICEMIB_DocsDevCpeTable) GetEntityData() *types.CommonEntityData {
    docsDevCpeTable.EntityData.YFilter = docsDevCpeTable.YFilter
    docsDevCpeTable.EntityData.YangName = "docsDevCpeTable"
    docsDevCpeTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevCpeTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevCpeTable.EntityData.SegmentPath = "docsDevCpeTable"
    docsDevCpeTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevCpeTable.EntityData.SegmentPath
    docsDevCpeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevCpeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevCpeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevCpeTable.EntityData.Children = types.NewOrderedMap()
    docsDevCpeTable.EntityData.Children.Append("docsDevCpeEntry", types.YChild{"DocsDevCpeEntry", nil})
    for i := range docsDevCpeTable.DocsDevCpeEntry {
        docsDevCpeTable.EntityData.Children.Append(types.GetSegmentPath(docsDevCpeTable.DocsDevCpeEntry[i]), types.YChild{"DocsDevCpeEntry", docsDevCpeTable.DocsDevCpeEntry[i]})
    }
    docsDevCpeTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevCpeTable.EntityData.YListKeys = []string {}

    return &(docsDevCpeTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry
// An entry in the docsDevFilterCpeTable.  There is one
// entry for each IPv4 CPE seen or provisioned.  If
// docsDevCpeIpMax is set to -1, this table is ignored;
// otherwise, upon receipt of an IP packet from the
// customer interface of the CM, the source IP address is
// checked against this table.  If the address is in the
// table, packet processing continues.  If the address is
// not in the table but docsDevCpeEnroll is set to any
// and the sum of the table sizes of docsDevCpeTable and
// docsDevCpeInetTable is less than docsDevCpeIpMax, the
// address is added to the table, and packet processing
// continues.  Otherwise, the packet is dropped.
// 
// The filtering actions specified by this table occur
// after any LLC filtering (docsDevFilterLLCTable), but
// prior to any IP filtering (docsDevFilterIpTable,
// docsDevNmAccessTable).
type DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IPv4 address to which this entry applies. 
    // N.B.  Attempts to set all zeros or all ones address values MUST be
    // rejected. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsDevCpeIp interface{}

    // This object describes how this entry was created.  If the value is
    // manual(2), this row was created by a network management action (either
    // configuration or SNMP set).  If set to learned(3), then it was found via
    // looking at the source IPv4 address of a received packet. The value other(1)
    // is used for any entries that do not meet manual(2) or learned(3) criteria.
    // The type is DocsDevCpeSource.
    DocsDevCpeSource interface{}

    // Standard object to manipulate rows.  To create a row in this table, one
    // only needs to specify this object. Management stations SHOULD use the
    // create-and-go mechanism for creating rows in this table. The type is
    // RowStatus.
    DocsDevCpeStatus interface{}
}

func (docsDevCpeEntry *DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry) GetEntityData() *types.CommonEntityData {
    docsDevCpeEntry.EntityData.YFilter = docsDevCpeEntry.YFilter
    docsDevCpeEntry.EntityData.YangName = "docsDevCpeEntry"
    docsDevCpeEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevCpeEntry.EntityData.ParentYangName = "docsDevCpeTable"
    docsDevCpeEntry.EntityData.SegmentPath = "docsDevCpeEntry" + types.AddKeyToken(docsDevCpeEntry.DocsDevCpeIp, "docsDevCpeIp")
    docsDevCpeEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevCpeTable/" + docsDevCpeEntry.EntityData.SegmentPath
    docsDevCpeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevCpeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevCpeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevCpeEntry.EntityData.Children = types.NewOrderedMap()
    docsDevCpeEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevCpeEntry.EntityData.Leafs.Append("docsDevCpeIp", types.YLeaf{"DocsDevCpeIp", docsDevCpeEntry.DocsDevCpeIp})
    docsDevCpeEntry.EntityData.Leafs.Append("docsDevCpeSource", types.YLeaf{"DocsDevCpeSource", docsDevCpeEntry.DocsDevCpeSource})
    docsDevCpeEntry.EntityData.Leafs.Append("docsDevCpeStatus", types.YLeaf{"DocsDevCpeStatus", docsDevCpeEntry.DocsDevCpeStatus})

    docsDevCpeEntry.EntityData.YListKeys = []string {"DocsDevCpeIp"}

    return &(docsDevCpeEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource represents meet manual(2) or learned(3) criteria.
type DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource string

const (
    DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource_other DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource = "other"

    DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource_manual DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource = "manual"

    DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource_learned DOCSCABLEDEVICEMIB_DocsDevCpeTable_DocsDevCpeEntry_DocsDevCpeSource = "learned"
)

// DOCSCABLEDEVICEMIB_DocsDevCpeInetTable
// This table lists the IP addresses seen (or permitted) as
// source addresses in packets originating from the
// customer interface on this device.  In addition, this
// table can be provisioned with the specific addresses
// permitted for the CPEs via the normal row creation
// mechanisms.
// 
// N.B.  Management action can add entries in this table
// and in docsDevCpeIpTable past the value of
// docsDevCpeIpMax.  docsDevCpeIpMax ONLY restricts the
// ability of the CM to add learned addresses
// automatically.
// 
// Table entries MUST NOT persist across reboots for any
// device.
// 
// This table exactly mirrors docsDevCpeTable and applies
// to IPv4 and IPv6 addresses.
type DOCSCABLEDEVICEMIB_DocsDevCpeInetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the docsDevFilterCpeInetTable. There is one entry for each IP
    // CPE seen or provisioned. If docsDevCpeIpMax is set to -1, this table is
    // ignored; otherwise, upon receipt of an IP packet from the customer
    // interface of the CM, the source IP address is checked against this table. 
    // If the address is in the table, packet processing continues.  If the
    // address is not in the table but docsDevCpeEnroll is set to any and the sum
    // of the table sizes for docsDevCpeTable and docsDevCpeInetTable is less than
    // docsDevCpeIpMax, the address is added to the table, and packet processing
    // continues.  Otherwise, the packet is dropped.  The filtering actions
    // specified by this table occur after any LLC filtering
    // (docsDevFilterLLCTable), but prior to any IP filtering
    // (docsDevFilterIpTable, docsDevNmAccessTable).  When an agent (cable modem)
    // restarts, then all dynamically created rows are lost. The type is slice of
    // DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry.
    DocsDevCpeInetEntry []*DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry
}

func (docsDevCpeInetTable *DOCSCABLEDEVICEMIB_DocsDevCpeInetTable) GetEntityData() *types.CommonEntityData {
    docsDevCpeInetTable.EntityData.YFilter = docsDevCpeInetTable.YFilter
    docsDevCpeInetTable.EntityData.YangName = "docsDevCpeInetTable"
    docsDevCpeInetTable.EntityData.BundleName = "cisco_ios_xe"
    docsDevCpeInetTable.EntityData.ParentYangName = "DOCS-CABLE-DEVICE-MIB"
    docsDevCpeInetTable.EntityData.SegmentPath = "docsDevCpeInetTable"
    docsDevCpeInetTable.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/" + docsDevCpeInetTable.EntityData.SegmentPath
    docsDevCpeInetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevCpeInetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevCpeInetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevCpeInetTable.EntityData.Children = types.NewOrderedMap()
    docsDevCpeInetTable.EntityData.Children.Append("docsDevCpeInetEntry", types.YChild{"DocsDevCpeInetEntry", nil})
    for i := range docsDevCpeInetTable.DocsDevCpeInetEntry {
        docsDevCpeInetTable.EntityData.Children.Append(types.GetSegmentPath(docsDevCpeInetTable.DocsDevCpeInetEntry[i]), types.YChild{"DocsDevCpeInetEntry", docsDevCpeInetTable.DocsDevCpeInetEntry[i]})
    }
    docsDevCpeInetTable.EntityData.Leafs = types.NewOrderedMap()

    docsDevCpeInetTable.EntityData.YListKeys = []string {}

    return &(docsDevCpeInetTable.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry
// An entry in the docsDevFilterCpeInetTable. There is one
// entry for each IP CPE seen or provisioned. If
// docsDevCpeIpMax is set to -1, this table is ignored;
// otherwise, upon receipt of an IP packet from the
// customer interface of the CM, the source IP address is
// checked against this table.  If the address is in the
// table, packet processing continues.  If the address is
// not in the table but docsDevCpeEnroll is set to any and
// the sum of the table sizes for docsDevCpeTable and
// docsDevCpeInetTable is less than docsDevCpeIpMax, the
// address is added to the table, and packet processing
// continues.  Otherwise, the packet is dropped.
// 
// The filtering actions specified by this table occur
// after any LLC filtering (docsDevFilterLLCTable), but
// prior to any IP filtering (docsDevFilterIpTable,
// docsDevNmAccessTable).
// 
// When an agent (cable modem) restarts, then all
// dynamically created rows are lost.
type DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type of internet address of
    // docsDevCpeInetAddr. The type is InetAddressType.
    DocsDevCpeInetType interface{}

    // This attribute is a key. The Internet address to which this entry applies. 
    // Implementors need to be aware that if the size of docsDevCpeInetAddr
    // exceeds 114 octets OIDs of instances of columns in this row will have more
    // than 128 sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or
    // SNMPv3.  Only unicast address are allowed for this object. The type is
    // string with length: 0..255.
    DocsDevCpeInetAddr interface{}

    // This object describes how this entry was created.  If the value is
    // manual(2), this row was created by a network management action (either
    // configuration or SNMP set).  If set to learned(3), then it was found via
    // looking at the source IP address of a received packet. The type is
    // DocsDevCpeInetSource.
    DocsDevCpeInetSource interface{}

    // Standard object to manipulate rows.  To create a row in this table, one
    // only needs to specify this object. Management stations SHOULD use the
    // create-and-go mechanism for creating rows in this table. The type is
    // RowStatus.
    DocsDevCpeInetRowStatus interface{}
}

func (docsDevCpeInetEntry *DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry) GetEntityData() *types.CommonEntityData {
    docsDevCpeInetEntry.EntityData.YFilter = docsDevCpeInetEntry.YFilter
    docsDevCpeInetEntry.EntityData.YangName = "docsDevCpeInetEntry"
    docsDevCpeInetEntry.EntityData.BundleName = "cisco_ios_xe"
    docsDevCpeInetEntry.EntityData.ParentYangName = "docsDevCpeInetTable"
    docsDevCpeInetEntry.EntityData.SegmentPath = "docsDevCpeInetEntry" + types.AddKeyToken(docsDevCpeInetEntry.DocsDevCpeInetType, "docsDevCpeInetType") + types.AddKeyToken(docsDevCpeInetEntry.DocsDevCpeInetAddr, "docsDevCpeInetAddr")
    docsDevCpeInetEntry.EntityData.AbsolutePath = "DOCS-CABLE-DEVICE-MIB:DOCS-CABLE-DEVICE-MIB/docsDevCpeInetTable/" + docsDevCpeInetEntry.EntityData.SegmentPath
    docsDevCpeInetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsDevCpeInetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsDevCpeInetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsDevCpeInetEntry.EntityData.Children = types.NewOrderedMap()
    docsDevCpeInetEntry.EntityData.Leafs = types.NewOrderedMap()
    docsDevCpeInetEntry.EntityData.Leafs.Append("docsDevCpeInetType", types.YLeaf{"DocsDevCpeInetType", docsDevCpeInetEntry.DocsDevCpeInetType})
    docsDevCpeInetEntry.EntityData.Leafs.Append("docsDevCpeInetAddr", types.YLeaf{"DocsDevCpeInetAddr", docsDevCpeInetEntry.DocsDevCpeInetAddr})
    docsDevCpeInetEntry.EntityData.Leafs.Append("docsDevCpeInetSource", types.YLeaf{"DocsDevCpeInetSource", docsDevCpeInetEntry.DocsDevCpeInetSource})
    docsDevCpeInetEntry.EntityData.Leafs.Append("docsDevCpeInetRowStatus", types.YLeaf{"DocsDevCpeInetRowStatus", docsDevCpeInetEntry.DocsDevCpeInetRowStatus})

    docsDevCpeInetEntry.EntityData.YListKeys = []string {"DocsDevCpeInetType", "DocsDevCpeInetAddr"}

    return &(docsDevCpeInetEntry.EntityData)
}

// DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource represents packet.
type DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource string

const (
    DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource_manual DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource = "manual"

    DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource_learned DOCSCABLEDEVICEMIB_DocsDevCpeInetTable_DocsDevCpeInetEntry_DocsDevCpeInetSource = "learned"
)

