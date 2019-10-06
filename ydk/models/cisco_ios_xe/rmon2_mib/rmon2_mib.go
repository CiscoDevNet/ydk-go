// The MIB module for managing remote monitoring
// device implementations. This MIB module
// augments the original RMON MIB as specified in
// RFC 1757.
package rmon2_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rmon2_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:RMON2-MIB RMON2-MIB}", reflect.TypeOf(RMON2MIB{}))
    ydk.RegisterEntity("RMON2-MIB:RMON2-MIB", reflect.TypeOf(RMON2MIB{}))
}

// RMON2MIB
type RMON2MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ProtocolDir RMON2MIB_ProtocolDir

    
    AddressMap RMON2MIB_AddressMap

    
    ProbeConfig RMON2MIB_ProbeConfig

    // This table lists the protocols that this agent has the capability to decode
    // and count.  There is one entry in this table for each such protocol.  These
    // protocols represent different network layer, transport layer, and
    // higher-layer protocols.  The agent should boot up with this table
    // preconfigured with those protocols that it knows about and wishes to
    // monitor.  Implementations are strongly encouraged to support protocols
    // higher than the network layer (at least for the protocol distribution
    // group), even for implementations that don't support the application layer
    // groups.
    ProtocolDirTable RMON2MIB_ProtocolDirTable

    // Controls the setup of protocol type distribution statistics tables. 
    // Implementations are encouraged to add an entry per monitored interface upon
    // initialization so that a default collection of protocol statistics is
    // available.  Rationale: This table controls collection of very basic
    // statistics for any or all of the protocols detected on a given interface.
    // An NMS can use this table to quickly determine bandwidth allocation
    // utilized by different protocols.  A media-specific statistics collection
    // could also be configured (e.g. etherStats, trPStats) to easily obtain total
    // frame, octet, and droppedEvents for the same interface.
    ProtocolDistControlTable RMON2MIB_ProtocolDistControlTable

    // An entry is made in this table for every protocol in the     
    // protocolDirTable which has been seen in at least one packet. Counters are
    // updated in this table for every protocol type that is encountered when
    // parsing a packet, but no counters are updated for packets with MAC-layer
    // errors.  Note that if a protocolDirEntry is deleted, all associated entries
    // in this table are removed.
    ProtocolDistStatsTable RMON2MIB_ProtocolDistStatsTable

    // A table to control the collection of network layer address to physical
    // address to interface mappings.  Note that this is not like the typical RMON
    // controlTable and dataTable in which each entry creates its own data table. 
    // Each entry in this table enables the discovery of addresses on a new
    // interface and the placement of address mappings into the central
    // addressMapTable.  Implementations are encouraged to add an entry per
    // monitored interface upon initialization so that a default collection of
    // address mappings is available.
    AddressMapControlTable RMON2MIB_AddressMapControlTable

    // A table of network layer address to physical address to interface mappings.
    // The probe will add entries to this table based on the source MAC and
    // network addresses seen in packets without MAC-level errors. The probe will
    // populate this table for all protocols in the protocol directory table whose
    // value of protocolDirAddressMapConfig is equal to supportedOn(3), and will
    // delete any entries whose protocolDirEntry is deleted or has a
    // protocolDirAddressMapConfig value of supportedOff(2).
    AddressMapTable RMON2MIB_AddressMapTable

    // A list of higher layer (i.e. non-MAC) host table control entries.  These
    // entries will enable the collection of the network and application level
    // host tables indexed by network addresses. Both the network and application
    // level host tables are controlled by this table is so that they will both be
    // created and deleted at the same time, further increasing the ease with
    // which they can be implemented as a single datastore (note that if an
    // implementation stores application layer host records in memory, it can
    // derive network layer host records from them).  Entries in the nlHostTable
    // will be created on behalf of each entry in this table. Additionally, if
    // this probe implements the alHostTable, entries in the alHostTable will be
    // created on behalf of each entry in this table.  Implementations are
    // encouraged to add an entry per monitored interface upon initialization so
    // that a default collection of host statistics is available.
    HlHostControlTable RMON2MIB_HlHostControlTable

    // A collection of statistics for a particular network layer address that has
    // been discovered on an interface of this device.  The probe will populate
    // this table for all network layer protocols in the protocol directory table
    // whose value of protocolDirHostConfig is equal to supportedOn(3), and will
    // delete any entries whose protocolDirEntry is deleted or has a
    // protocolDirHostConfig value of supportedOff(2).  The probe will add to this
    // table all addresses seen as the source or destination address in all
    // packets with no MAC errors, and will increment octet and packet counts in
    // the table for all packets with no MAC errors.
    NlHostTable RMON2MIB_NlHostTable

    // A list of higher layer (i.e. non-MAC) matrix control entries.  These
    // entries will enable the collection of the network and application level
    // matrix tables containing conversation statistics indexed by pairs of
    // network addresses. Both the network and application level matrix tables are
    // controlled by this table is so that they will both be created and deleted
    // at the same time, further increasing the ease with which they can be
    // implemented as a single datastore (note that if an implementation stores
    // application layer matrix records in memory, it can derive network layer
    // matrix records from them).  Entries in the nlMatrixSDTable and
    // nlMatrixDSTable will be created on behalf of each entry in this table. 
    // Additionally, if this probe implements the alMatrix tables, entries in the
    // alMatrix tables will be created on behalf of each entry in this table.
    HlMatrixControlTable RMON2MIB_HlMatrixControlTable

    // A list of traffic matrix entries which collect statistics for conversations
    // between two network-level addresses.  This table is indexed first by the
    // source address and then by the destination address to make it convenient to
    // collect all conversations from a particular address.  The probe will
    // populate this table for all network layer protocols in the protocol
    // directory table whose value of protocolDirMatrixConfig is equal to
    // supportedOn(3), and will delete any entries whose protocolDirEntry is
    // deleted or has a protocolDirMatrixConfig value of supportedOff(2).      The
    // probe will add to this table all pairs of addresses seen in all packets
    // with no MAC errors, and will increment octet and packet counts in the table
    // for all packets with no MAC errors.  Further, this table will only contain
    // entries that have a corresponding entry in the nlMatrixDSTable with the
    // same source address and destination address.
    NlMatrixSDTable RMON2MIB_NlMatrixSDTable

    // A list of traffic matrix entries which collect statistics for conversations
    // between two network-level addresses.  This table is indexed first by the
    // destination address and then by the source address to make it convenient to
    // collect all conversations to a particular address.  The probe will populate
    // this table for all network layer protocols in the protocol directory table
    // whose value of protocolDirMatrixConfig is equal to supportedOn(3), and will
    // delete any entries whose protocolDirEntry is deleted or has a
    // protocolDirMatrixConfig value of supportedOff(2).  The probe will add to
    // this table all pairs of addresses seen in all packets with no MAC errors,
    // and will increment      octet and packet counts in the table for all
    // packets with no MAC errors.  Further, this table will only contain entries
    // that have a corresponding entry in the nlMatrixSDTable with the same source
    // address and destination address.
    NlMatrixDSTable RMON2MIB_NlMatrixDSTable

    // A set of parameters that control the creation of a report of the top N
    // matrix entries according to a selected metric.
    NlMatrixTopNControlTable RMON2MIB_NlMatrixTopNControlTable

    // A set of statistics for those network layer matrix entries that have
    // counted the highest number of octets or packets.
    NlMatrixTopNTable RMON2MIB_NlMatrixTopNTable

    // A collection of statistics for a particular protocol from a particular
    // network address that has been discovered on an interface of this device. 
    // The probe will populate this table for all protocols in the protocol
    // directory table whose value of protocolDirHostConfig is equal to
    // supportedOn(3), and will delete any entries whose protocolDirEntry is
    // deleted or has a protocolDirHostConfig value of supportedOff(2).  The probe
    // will add to this table all addresses seen as the source or destination
    // address in all packets with no MAC errors, and will increment octet and
    // packet counts in the table for all packets with no MAC errors.  Further,
    // entries will only be added to this table if their address exists in the
    // nlHostTable and will be deleted from this table if their address is deleted
    // from the nlHostTable.
    AlHostTable RMON2MIB_AlHostTable

    // A list of application traffic matrix entries which collect      statistics
    // for conversations of a particular protocol between two network-level
    // addresses.  This table is indexed first by the source address and then by
    // the destination address to make it convenient to collect all statistics
    // from a particular address.  The probe will populate this table for all
    // protocols in the protocol directory table whose value of
    // protocolDirMatrixConfig is equal to supportedOn(3), and will delete any
    // entries whose protocolDirEntry is deleted or has a protocolDirMatrixConfig
    // value of supportedOff(2).  The probe will add to this table all pairs of
    // addresses for all protocols seen in all packets with no MAC errors, and
    // will increment octet and packet counts in the table for all packets with no
    // MAC errors.  Further, entries will only be added to this table if their
    // address pair exists in the nlMatrixSDTable and will be deleted from this
    // table if the address pair is deleted from the nlMatrixSDTable.
    AlMatrixSDTable RMON2MIB_AlMatrixSDTable

    // A list of application traffic matrix entries which collect statistics for
    // conversations of a particular protocol between two network-level addresses.
    // This table is indexed first by the destination address and then by the
    // source address to make it convenient to collect all statistics to a
    // particular address.  The probe will populate this table for all protocols
    // in the protocol directory table whose value of protocolDirMatrixConfig is
    // equal to supportedOn(3), and will delete any entries whose protocolDirEntry
    // is deleted or has a protocolDirMatrixConfig value of supportedOff(2).  The
    // probe will add to this table all pairs of addresses for all protocols seen
    // in all packets with no MAC errors, and will increment octet and packet
    // counts in the table for all packets with no MAC errors.  Further, entries
    // will only be added to this table if their address pair exists in the
    // nlMatrixDSTable and will be deleted from this table if the address pair is
    // deleted from the nlMatrixDSTable.
    AlMatrixDSTable RMON2MIB_AlMatrixDSTable

    // A set of parameters that control the creation of a report of the top N
    // matrix entries according to a selected metric.
    AlMatrixTopNControlTable RMON2MIB_AlMatrixTopNControlTable

    // A set of statistics for those application layer matrix entries that have
    // counted the highest number of octets or packets.
    AlMatrixTopNTable RMON2MIB_AlMatrixTopNTable

    // A list of data-collection configuration entries.
    UsrHistoryControlTable RMON2MIB_UsrHistoryControlTable

    // A list of data-collection configuration entries.
    UsrHistoryObjectTable RMON2MIB_UsrHistoryObjectTable

    // A list of user defined history entries.
    UsrHistoryTable RMON2MIB_UsrHistoryTable

    // A table of serial interface configuration entries.  This data will be
    // stored in non-volatile memory and preserved across probe resets or power
    // loss.
    SerialConfigTable RMON2MIB_SerialConfigTable

    // A table of netConfigEntries.
    NetConfigTable RMON2MIB_NetConfigTable

    // A list of trap destination entries.
    TrapDestTable RMON2MIB_TrapDestTable

    // A list of serialConnectionEntries.
    SerialConnectionTable RMON2MIB_SerialConnectionTable
}

func (rMON2MIB *RMON2MIB) GetEntityData() *types.CommonEntityData {
    rMON2MIB.EntityData.YFilter = rMON2MIB.YFilter
    rMON2MIB.EntityData.YangName = "RMON2-MIB"
    rMON2MIB.EntityData.BundleName = "cisco_ios_xe"
    rMON2MIB.EntityData.ParentYangName = "RMON2-MIB"
    rMON2MIB.EntityData.SegmentPath = "RMON2-MIB:RMON2-MIB"
    rMON2MIB.EntityData.AbsolutePath = rMON2MIB.EntityData.SegmentPath
    rMON2MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rMON2MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rMON2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rMON2MIB.EntityData.Children = types.NewOrderedMap()
    rMON2MIB.EntityData.Children.Append("protocolDir", types.YChild{"ProtocolDir", &rMON2MIB.ProtocolDir})
    rMON2MIB.EntityData.Children.Append("addressMap", types.YChild{"AddressMap", &rMON2MIB.AddressMap})
    rMON2MIB.EntityData.Children.Append("probeConfig", types.YChild{"ProbeConfig", &rMON2MIB.ProbeConfig})
    rMON2MIB.EntityData.Children.Append("protocolDirTable", types.YChild{"ProtocolDirTable", &rMON2MIB.ProtocolDirTable})
    rMON2MIB.EntityData.Children.Append("protocolDistControlTable", types.YChild{"ProtocolDistControlTable", &rMON2MIB.ProtocolDistControlTable})
    rMON2MIB.EntityData.Children.Append("protocolDistStatsTable", types.YChild{"ProtocolDistStatsTable", &rMON2MIB.ProtocolDistStatsTable})
    rMON2MIB.EntityData.Children.Append("addressMapControlTable", types.YChild{"AddressMapControlTable", &rMON2MIB.AddressMapControlTable})
    rMON2MIB.EntityData.Children.Append("addressMapTable", types.YChild{"AddressMapTable", &rMON2MIB.AddressMapTable})
    rMON2MIB.EntityData.Children.Append("hlHostControlTable", types.YChild{"HlHostControlTable", &rMON2MIB.HlHostControlTable})
    rMON2MIB.EntityData.Children.Append("nlHostTable", types.YChild{"NlHostTable", &rMON2MIB.NlHostTable})
    rMON2MIB.EntityData.Children.Append("hlMatrixControlTable", types.YChild{"HlMatrixControlTable", &rMON2MIB.HlMatrixControlTable})
    rMON2MIB.EntityData.Children.Append("nlMatrixSDTable", types.YChild{"NlMatrixSDTable", &rMON2MIB.NlMatrixSDTable})
    rMON2MIB.EntityData.Children.Append("nlMatrixDSTable", types.YChild{"NlMatrixDSTable", &rMON2MIB.NlMatrixDSTable})
    rMON2MIB.EntityData.Children.Append("nlMatrixTopNControlTable", types.YChild{"NlMatrixTopNControlTable", &rMON2MIB.NlMatrixTopNControlTable})
    rMON2MIB.EntityData.Children.Append("nlMatrixTopNTable", types.YChild{"NlMatrixTopNTable", &rMON2MIB.NlMatrixTopNTable})
    rMON2MIB.EntityData.Children.Append("alHostTable", types.YChild{"AlHostTable", &rMON2MIB.AlHostTable})
    rMON2MIB.EntityData.Children.Append("alMatrixSDTable", types.YChild{"AlMatrixSDTable", &rMON2MIB.AlMatrixSDTable})
    rMON2MIB.EntityData.Children.Append("alMatrixDSTable", types.YChild{"AlMatrixDSTable", &rMON2MIB.AlMatrixDSTable})
    rMON2MIB.EntityData.Children.Append("alMatrixTopNControlTable", types.YChild{"AlMatrixTopNControlTable", &rMON2MIB.AlMatrixTopNControlTable})
    rMON2MIB.EntityData.Children.Append("alMatrixTopNTable", types.YChild{"AlMatrixTopNTable", &rMON2MIB.AlMatrixTopNTable})
    rMON2MIB.EntityData.Children.Append("usrHistoryControlTable", types.YChild{"UsrHistoryControlTable", &rMON2MIB.UsrHistoryControlTable})
    rMON2MIB.EntityData.Children.Append("usrHistoryObjectTable", types.YChild{"UsrHistoryObjectTable", &rMON2MIB.UsrHistoryObjectTable})
    rMON2MIB.EntityData.Children.Append("usrHistoryTable", types.YChild{"UsrHistoryTable", &rMON2MIB.UsrHistoryTable})
    rMON2MIB.EntityData.Children.Append("serialConfigTable", types.YChild{"SerialConfigTable", &rMON2MIB.SerialConfigTable})
    rMON2MIB.EntityData.Children.Append("netConfigTable", types.YChild{"NetConfigTable", &rMON2MIB.NetConfigTable})
    rMON2MIB.EntityData.Children.Append("trapDestTable", types.YChild{"TrapDestTable", &rMON2MIB.TrapDestTable})
    rMON2MIB.EntityData.Children.Append("serialConnectionTable", types.YChild{"SerialConnectionTable", &rMON2MIB.SerialConnectionTable})
    rMON2MIB.EntityData.Leafs = types.NewOrderedMap()

    rMON2MIB.EntityData.YListKeys = []string {}

    return &(rMON2MIB.EntityData)
}

// RMON2MIB_ProtocolDir
type RMON2MIB_ProtocolDir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time the protocol directory was last
    // modified, either through insertions or deletions, or through modifications
    // of either the protocolDirAddressMapConfig, protocolDirHostConfig, or
    // protocolDirMatrixConfig. The type is interface{} with range: 0..4294967295.
    ProtocolDirLastChange interface{}
}

func (protocolDir *RMON2MIB_ProtocolDir) GetEntityData() *types.CommonEntityData {
    protocolDir.EntityData.YFilter = protocolDir.YFilter
    protocolDir.EntityData.YangName = "protocolDir"
    protocolDir.EntityData.BundleName = "cisco_ios_xe"
    protocolDir.EntityData.ParentYangName = "RMON2-MIB"
    protocolDir.EntityData.SegmentPath = "protocolDir"
    protocolDir.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + protocolDir.EntityData.SegmentPath
    protocolDir.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDir.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDir.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDir.EntityData.Children = types.NewOrderedMap()
    protocolDir.EntityData.Leafs = types.NewOrderedMap()
    protocolDir.EntityData.Leafs.Append("protocolDirLastChange", types.YLeaf{"ProtocolDirLastChange", protocolDir.ProtocolDirLastChange})

    protocolDir.EntityData.YListKeys = []string {}

    return &(protocolDir.EntityData)
}

// RMON2MIB_AddressMap
type RMON2MIB_AddressMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an address mapping entry has been inserted into the
    // addressMapTable.  If an entry is inserted, then deleted, and then inserted,
    // this counter will be incremented by 2.  Note that the table size can be
    // determined by subtracting addressMapDeletes from addressMapInserts. The
    // type is interface{} with range: 0..4294967295.
    AddressMapInserts interface{}

    // The number of times an address mapping entry has been deleted from the
    // addressMapTable (for any reason).  If an entry is deleted, then inserted,
    // and then deleted, this counter will be incremented by 2.  Note that the
    // table size can be determined by subtracting addressMapDeletes from
    // addressMapInserts. The type is interface{} with range: 0..4294967295.
    AddressMapDeletes interface{}

    // The maximum number of entries that are desired in the addressMapTable. The
    // probe will not create more than this number of entries in the table, but
    // may choose to create fewer entries in this table for any reason including
    // the lack of resources.  If this object is set to a value less than the
    // current number of entries, enough entries are chosen in an
    // implementation-dependent manner and deleted so that the number of entries
    // in the table equals the value of this object.  If this value is set to -1,
    // the probe may create any number of entries in this table.  This object may
    // be used to control how resources are allocated on the probe for the various
    // RMON functions. The type is interface{} with range: -1..2147483647.
    AddressMapMaxDesiredEntries interface{}
}

func (addressMap *RMON2MIB_AddressMap) GetEntityData() *types.CommonEntityData {
    addressMap.EntityData.YFilter = addressMap.YFilter
    addressMap.EntityData.YangName = "addressMap"
    addressMap.EntityData.BundleName = "cisco_ios_xe"
    addressMap.EntityData.ParentYangName = "RMON2-MIB"
    addressMap.EntityData.SegmentPath = "addressMap"
    addressMap.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + addressMap.EntityData.SegmentPath
    addressMap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressMap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressMap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressMap.EntityData.Children = types.NewOrderedMap()
    addressMap.EntityData.Leafs = types.NewOrderedMap()
    addressMap.EntityData.Leafs.Append("addressMapInserts", types.YLeaf{"AddressMapInserts", addressMap.AddressMapInserts})
    addressMap.EntityData.Leafs.Append("addressMapDeletes", types.YLeaf{"AddressMapDeletes", addressMap.AddressMapDeletes})
    addressMap.EntityData.Leafs.Append("addressMapMaxDesiredEntries", types.YLeaf{"AddressMapMaxDesiredEntries", addressMap.AddressMapMaxDesiredEntries})

    addressMap.EntityData.YListKeys = []string {}

    return &(addressMap.EntityData)
}

// RMON2MIB_ProbeConfig
type RMON2MIB_ProbeConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of the RMON MIB groups supported on at least one interface by
    // this probe. The type is string with length: 1..1.
    ProbeCapabilities interface{}

    // The software revision of this device.  This string will have a zero length
    // if the revision is unknown. The type is string with length: 0..31.
    ProbeSoftwareRev interface{}

    // The hardware revision of this device.  This string will have a zero length
    // if the revision is unknown. The type is string with length: 0..31.
    ProbeHardwareRev interface{}

    // Probe's current date and time.  field  octets  contents                 
    // range -----  ------  --------                  -----   1      1-2   year   
    // 0..65536   2       3    month                     1..12   3       4    day 
    // 1..31   4       5    hour                      0..23   5       6    minutes
    // 0..59   6       7    seconds                   0..60                 (use
    // 60 for leap-second)   7       8    deci-seconds              0..9   8      
    // 9    direction from UTC        '+' / '-'   9      10    hours from UTC     
    // 0..11  10      11    minutes from UTC          0..59  For example, Tuesday
    // May 26, 1992 at 1:30:15 PM EDT would be displayed as:             
    // 1992-5-26,13:30:15.0,-4:0  Note that if only local time is known, then
    // timezone information (fields 8-10) is not present, and if no time
    // information is known, the null string is returned. The type is string with
    // length: 0..0 | 8..8 | 11..11.
    ProbeDateTime interface{}

    // Setting this object to warmBoot(2) causes the device to restart the
    // application software with current configuration parameters saved in
    // non-volatile memory.  Setting this object to coldBoot(3) causes the device
    // to reinitialize configuration parameters in non-volatile memory to default
    // values and restart the application software.  When the device is running
    // normally, this variable has a value of running(1). The type is
    // ProbeResetControl.
    ProbeResetControl interface{}

    // The file name to be downloaded from the TFTP server when a download is next
    // requested via this MIB.  This value is set to the zero length string when
    // no file name has been specified. The type is string with length: 0..127.
    ProbeDownloadFile interface{}

    // The IP address of the TFTP server that contains the boot image to load when
    // a download is next requested via this MIB. This value is set to `0.0.0.0'
    // when no IP address has been specified. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ProbeDownloadTFTPServer interface{}

    // When this object is set to downloadToRAM(2) or downloadToPROM(3), the
    // device will discontinue its normal operation and begin download of the
    // image specified by probeDownloadFile from the server specified by
    // probeDownloadTFTPServer using the TFTP protocol.  If downloadToRAM(2) is
    // specified, the new image is copied to RAM only (the old image remains
    // unaltered in the flash EPROM).  If downloadToPROM(3) is specified the new
    // image is written to the flash EPROM memory after its checksum has been
    // verified to be correct. When the download process is completed, the device
    // will      warm boot to restart the newly loaded application. When the
    // device is not downloading, this object will have a value of
    // notDownloading(1). The type is ProbeDownloadAction.
    ProbeDownloadAction interface{}

    // The status of the last download procedure, if any.  This object will have a
    // value of downloadStatusUnknown(2) if no download process has been
    // performed. The type is ProbeDownloadStatus.
    ProbeDownloadStatus interface{}

    // The IP Address of the default gateway.  If this value is undefined or
    // unknown, it shall have the value 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetDefaultGateway interface{}
}

func (probeConfig *RMON2MIB_ProbeConfig) GetEntityData() *types.CommonEntityData {
    probeConfig.EntityData.YFilter = probeConfig.YFilter
    probeConfig.EntityData.YangName = "probeConfig"
    probeConfig.EntityData.BundleName = "cisco_ios_xe"
    probeConfig.EntityData.ParentYangName = "RMON2-MIB"
    probeConfig.EntityData.SegmentPath = "probeConfig"
    probeConfig.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + probeConfig.EntityData.SegmentPath
    probeConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    probeConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    probeConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    probeConfig.EntityData.Children = types.NewOrderedMap()
    probeConfig.EntityData.Leafs = types.NewOrderedMap()
    probeConfig.EntityData.Leafs.Append("probeCapabilities", types.YLeaf{"ProbeCapabilities", probeConfig.ProbeCapabilities})
    probeConfig.EntityData.Leafs.Append("probeSoftwareRev", types.YLeaf{"ProbeSoftwareRev", probeConfig.ProbeSoftwareRev})
    probeConfig.EntityData.Leafs.Append("probeHardwareRev", types.YLeaf{"ProbeHardwareRev", probeConfig.ProbeHardwareRev})
    probeConfig.EntityData.Leafs.Append("probeDateTime", types.YLeaf{"ProbeDateTime", probeConfig.ProbeDateTime})
    probeConfig.EntityData.Leafs.Append("probeResetControl", types.YLeaf{"ProbeResetControl", probeConfig.ProbeResetControl})
    probeConfig.EntityData.Leafs.Append("probeDownloadFile", types.YLeaf{"ProbeDownloadFile", probeConfig.ProbeDownloadFile})
    probeConfig.EntityData.Leafs.Append("probeDownloadTFTPServer", types.YLeaf{"ProbeDownloadTFTPServer", probeConfig.ProbeDownloadTFTPServer})
    probeConfig.EntityData.Leafs.Append("probeDownloadAction", types.YLeaf{"ProbeDownloadAction", probeConfig.ProbeDownloadAction})
    probeConfig.EntityData.Leafs.Append("probeDownloadStatus", types.YLeaf{"ProbeDownloadStatus", probeConfig.ProbeDownloadStatus})
    probeConfig.EntityData.Leafs.Append("netDefaultGateway", types.YLeaf{"NetDefaultGateway", probeConfig.NetDefaultGateway})

    probeConfig.EntityData.YListKeys = []string {}

    return &(probeConfig.EntityData)
}

// RMON2MIB_ProbeConfig_ProbeDownloadAction represents a value of notDownloading(1).
type RMON2MIB_ProbeConfig_ProbeDownloadAction string

const (
    RMON2MIB_ProbeConfig_ProbeDownloadAction_notDownloading RMON2MIB_ProbeConfig_ProbeDownloadAction = "notDownloading"

    RMON2MIB_ProbeConfig_ProbeDownloadAction_downloadToPROM RMON2MIB_ProbeConfig_ProbeDownloadAction = "downloadToPROM"

    RMON2MIB_ProbeConfig_ProbeDownloadAction_downloadToRAM RMON2MIB_ProbeConfig_ProbeDownloadAction = "downloadToRAM"
)

// RMON2MIB_ProbeConfig_ProbeDownloadStatus represents download process has been performed.
type RMON2MIB_ProbeConfig_ProbeDownloadStatus string

const (
    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadSuccess RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadSuccess"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadStatusUnknown RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadStatusUnknown"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadGeneralError RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadGeneralError"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadNoResponseFromServer RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadNoResponseFromServer"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadChecksumError RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadChecksumError"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadIncompatibleImage RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadIncompatibleImage"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadTftpFileNotFound RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadTftpFileNotFound"

    RMON2MIB_ProbeConfig_ProbeDownloadStatus_downloadTftpAccessViolation RMON2MIB_ProbeConfig_ProbeDownloadStatus = "downloadTftpAccessViolation"
)

// RMON2MIB_ProbeConfig_ProbeResetControl represents running(1).
type RMON2MIB_ProbeConfig_ProbeResetControl string

const (
    RMON2MIB_ProbeConfig_ProbeResetControl_running RMON2MIB_ProbeConfig_ProbeResetControl = "running"

    RMON2MIB_ProbeConfig_ProbeResetControl_warmBoot RMON2MIB_ProbeConfig_ProbeResetControl = "warmBoot"

    RMON2MIB_ProbeConfig_ProbeResetControl_coldBoot RMON2MIB_ProbeConfig_ProbeResetControl = "coldBoot"
)

// RMON2MIB_ProtocolDirTable
// This table lists the protocols that this agent has the
// capability to decode and count.  There is one entry in this
// table for each such protocol.  These protocols represent
// different network layer, transport layer, and higher-layer
// protocols.  The agent should boot up with this table
// preconfigured with those protocols that it knows about and
// wishes to monitor.  Implementations are strongly encouraged to
// support protocols higher than the network layer (at least for
// the protocol distribution group), even for implementations
// that don't support the application layer groups.
type RMON2MIB_ProtocolDirTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDirTable.  An example of the indexing of
    // this entry is protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
    // encoding of a length of 8, followed by 8 subids encoding the protocolDirID
    // of 1.2048, followed by a length of 2 and the 2 subids encoding zero-valued
    // parameters. The type is slice of
    // RMON2MIB_ProtocolDirTable_ProtocolDirEntry.
    ProtocolDirEntry []*RMON2MIB_ProtocolDirTable_ProtocolDirEntry
}

func (protocolDirTable *RMON2MIB_ProtocolDirTable) GetEntityData() *types.CommonEntityData {
    protocolDirTable.EntityData.YFilter = protocolDirTable.YFilter
    protocolDirTable.EntityData.YangName = "protocolDirTable"
    protocolDirTable.EntityData.BundleName = "cisco_ios_xe"
    protocolDirTable.EntityData.ParentYangName = "RMON2-MIB"
    protocolDirTable.EntityData.SegmentPath = "protocolDirTable"
    protocolDirTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + protocolDirTable.EntityData.SegmentPath
    protocolDirTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDirTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDirTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDirTable.EntityData.Children = types.NewOrderedMap()
    protocolDirTable.EntityData.Children.Append("protocolDirEntry", types.YChild{"ProtocolDirEntry", nil})
    for i := range protocolDirTable.ProtocolDirEntry {
        protocolDirTable.EntityData.Children.Append(types.GetSegmentPath(protocolDirTable.ProtocolDirEntry[i]), types.YChild{"ProtocolDirEntry", protocolDirTable.ProtocolDirEntry[i]})
    }
    protocolDirTable.EntityData.Leafs = types.NewOrderedMap()

    protocolDirTable.EntityData.YListKeys = []string {}

    return &(protocolDirTable.EntityData)
}

// RMON2MIB_ProtocolDirTable_ProtocolDirEntry
// A conceptual row in the protocolDirTable.
// 
// An example of the indexing of this entry is
// protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
// encoding of a length of 8, followed by 8 subids encoding the
// protocolDirID of 1.2048, followed by a length of 2 and the
// 2 subids encoding zero-valued parameters.
type RMON2MIB_ProtocolDirTable_ProtocolDirEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier for a particular protocol. 
    // Standard identifiers will be defined in a manner such that they can often
    // be used as specifications for new protocols - i.e. a tree-structured
    // assignment mechanism that matches the protocol encapsulation `tree' and
    // which has algorithmic assignment mechanisms for certain subtrees. See RFC
    // XXX for more details.  Despite the algorithmic mechanism, the probe will
    // only place entries in here for those protocols it chooses to collect.  In
    // other words, it need not populate this table with all of the possible
    // ethernet protocol types, nor need it create them on the fly when it sees
    // them.  Whether or not it does these things is a matter of product
    // definition (cost/benefit, usability), and is up to the designer of the
    // product.  If an entry is written to this table with a protocolDirID that
    // the agent doesn't understand, either directly or algorithmically, the SET
    // request will be rejected with an inconsistentName or badValue (for SNMPv1)
    // error. The type is string.
    ProtocolDirID interface{}

    // This attribute is a key. A set of parameters for the associated
    // protocolDirID. See the associated RMON2 Protocol Identifiers document for a
    // description of the possible parameters. There will be one octet in this
    // string for each sub-identifier in the protocolDirID, and the parameters
    // will appear here in the same order as the associated sub-identifiers appear
    // in the protocolDirID.  Every node in the protocolDirID tree has a
    // different, optional set of parameters defined (that is, the definition of
    // parameters for a node is optional).  The proper parameter value for each
    // node is included in this string.  Note that the inclusion of a parameter
    // value in this string for each node is not optional - what is optional is
    // that a node may have no parameters defined, in which case the parameter
    // field for that node will be zero. The type is string.
    ProtocolDirParameters interface{}

    // The locally arbitrary, but unique identifier associated with this
    // protocolDir entry.  The value for each supported protocol must remain
    // constant at least from one re-initialization of the entity's network
    // management system to the next re-initialization, except that if a protocol
    // is deleted and re-created, it must be re-created with a new value that has
    // not been used since the last re-initialization.  The specific value is
    // meaningful only within a given SNMP entity. A protocolDirLocalIndex must
    // not be re-used until the next agent-restart in the event the protocol
    // directory entry is deleted. The type is interface{} with range:
    // 1..2147483647.
    ProtocolDirLocalIndex interface{}

    // A textual description of the protocol encapsulation. A probe may choose to
    // describe only a subset of the entire encapsulation (e.g. only the highest
    // layer).  This object is intended for human consumption only.  This object
    // may not be modified if the associated protocolDirStatus object is equal to
    // active(1). The type is string with length: 1..64.
    ProtocolDirDescr interface{}

    // This object describes 2 attributes of this protocol directory entry.  The
    // presence or absence of the `extensible' bit describes whether or not this
    // protocol directory entry can be extended      by the user by creating
    // protocol directory entries which are children of this protocol.  An example
    // of an entry that will often allow extensibility is `ip.udp'.  The probe may
    // automatically populate some children of this node such as `ip.udp.snmp' and
    // `ip.udp.dns'. A probe administrator or user may also populate additional
    // children via remote SNMP requests that create entries in this table.  When
    // a child node is added for a protocol for which the probe has no built in
    // support, extending a parent node (for which the probe does have built in
    // support), that child node is not extendible.  This is termed `limited
    // extensibility'.  When a child node is added through this extensibility
    // mechanism, the values of protocolDirLocalIndex and protocolDirType shall be
    // assigned by the agent.  The other objects in the entry will be assigned by
    // the manager who is creating the new entry.  This object also describes
    // whether or not this agent can recognize addresses for this protocol, should
    // it be a network level protocol.  That is, while a probe may be able to
    // recognize packets of a particular network layer protocol and count them, it
    // takes additional logic to be able to recognize the addresses in this
    // protocol and to populate network layer or application layer tables with the
    // addresses in this protocol.  If this bit is set, the agent will recognize
    // network layer addresses for this protoocl and populate the network and
    // application layer host and matrix tables with these protocols.  Note that
    // when an entry is created, the agent will supply values for the bits that
    // match the capabilities of the agent with respect to this protocol.  Note
    // that since row creations usually exercise the limited extensibility
    // feature, these bits will usually be set to zero. The type is string with
    // length: 1..1.
    ProtocolDirType interface{}

    // This object describes and configures the probe's support for address
    // mapping for this protocol.  When the probe creates entries in this table
    // for all protocols that it understands, it will set the entry to
    // notSupported(1) if it doesn't have the capability to perform address
    // mapping for the protocol or if this protocol is not a network-layer
    // protocol.  When an entry is created in this table by a management operation
    // as part of the limited extensibility feature, the probe must set this value
    // to notSupported(1), because limited extensibility of the protocolDirTable
    // does not extend to interpreting addresses of the extended protocols.  If
    // the value of this object is notSupported(1), the probe will not perform
    // address mapping for this protocol and shall not allow this object to be
    // changed to any other value. If the value of this object is supportedOn(3),
    // the probe supports address mapping for this protocol and is configured to
    // perform address mapping for this protocol for all
    // addressMappingControlEntries and all interfaces. If the value of this
    // object is supportedOff(2), the probe supports address mapping for this
    // protocol but is configured to not perform address mapping for this protocol
    // for any addressMappingControlEntries and all interfaces. Whenever this
    // value changes from supportedOn(3) to supportedOff(2), the probe shall
    // delete all related entries in the addressMappingTable. The type is
    // ProtocolDirAddressMapConfig.
    ProtocolDirAddressMapConfig interface{}

    // This object describes and configures the probe's support for the network
    // layer and application layer host tables for this protocol.  When the probe
    // creates entries in this table for all protocols that it understands, it
    // will set the entry to notSupported(1) if it doesn't have the capability to
    // track the nlHostTable for this protocol or if the alHostTable is
    // implemented but doesn't have the capability to track this protocol.  Note
    // that if the alHostTable is implemented, the probe may only support a
    // protocol if it is supported in both the nlHostTable and the alHostTable.   
    // If the associated protocolDirType object has the addressRecognitionCapable
    // bit set, then this is a network layer protocol for which the probe
    // recognizes addresses, and thus the probe will populate the nlHostTable and
    // alHostTable with addresses it discovers for this protocol.  If the value of
    // this object is notSupported(1), the probe will not track the nlHostTable or
    // alHostTable for this protocol and shall not allow this object to be changed
    // to any other value. If the value of this object is supportedOn(3), the
    // probe supports tracking of the nlHostTable and alHostTable for this
    // protocol and is configured to track both tables for this protocol for all
    // control entries and all interfaces. If the value of this object is
    // supportedOff(2), the probe supports tracking of the nlHostTable and
    // alHostTable for this protocol but is configured to not track these tables
    // for any control entries or interfaces. Whenever this value changes from
    // supportedOn(3) to supportedOff(2), the probe shall delete all related
    // entries in the nlHostTable and alHostTable.  Note that since each
    // alHostEntry references 2 protocol directory entries, one for the network
    // address and one for the type of the highest protocol recognized, that an
    // entry will only be created in that table if this value is supportedOn(3)
    // for both protocols. The type is ProtocolDirHostConfig.
    ProtocolDirHostConfig interface{}

    // This object describes and configures the probe's support for the network
    // layer and application layer matrix tables for this protocol.  When the
    // probe creates entries in this table for all protocols that it understands,
    // it will set the entry to notSupported(1) if it doesn't have the capability
    // to track the nlMatrixTables for this protocol or if the alMatrixTables are
    // implemented but don't have the capability to track this protocol.  Note
    // that if the alMatrix tables are implemented, the probe may only support a
    // protocol if it is supported in the the both of the nlMatrixTables and both
    // of the alMatrixTables.      If the associated protocolDirType object has
    // the addressRecognitionCapable bit set, then this is a network layer
    // protocol for which the probe recognizes addresses, and thus the probe will
    // populate both of the nlMatrixTables and both of the alMatrixTables with
    // addresses it discovers for this protocol.  If the value of this object is
    // notSupported(1), the probe will not track either of the nlMatrixTables or
    // the alMatrixTables for this protocol and shall not allow this object to be
    // changed to any other value. If the value of this object is supportedOn(3),
    // the probe supports tracking of both of the nlMatrixTables and (if
    // implemented) both of the alMatrixTables for this protocol and is configured
    // to track these tables for this protocol for all control entries and all
    // interfaces. If the value of this object is supportedOff(2), the probe
    // supports tracking of both of the nlMatrixTables and (if implemented) both
    // of the alMatrixTables for this protocol but is configured to not track
    // these tables for this protocol for any control entries or interfaces.
    // Whenever this value changes from supportedOn(3) to supportedOff(2), the
    // probe shall delete all related entries in the nlMatrixTables and the
    // alMatrixTables.  Note that since each alMatrixEntry references 2 protocol
    // directory entries, one for the network address and one for the type of the
    // highest protocol recognized, that an entry will only be created in that
    // table if this value is supportedOn(3) for both protocols. The type is
    // ProtocolDirMatrixConfig.
    ProtocolDirMatrixConfig interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    ProtocolDirOwner interface{}

    // The status of this protocol directory entry.  An entry may not exist in the
    // active state unless all      objects in the entry have an appropriate
    // value.  If this object is not equal to active(1), all associated entries in
    // the nlHostTable, nlMatrixSDTable, nlMatrixDSTable, alHostTable,
    // alMatrixSDTable, and alMatrixDSTable shall be deleted. The type is
    // RowStatus.
    ProtocolDirStatus interface{}
}

func (protocolDirEntry *RMON2MIB_ProtocolDirTable_ProtocolDirEntry) GetEntityData() *types.CommonEntityData {
    protocolDirEntry.EntityData.YFilter = protocolDirEntry.YFilter
    protocolDirEntry.EntityData.YangName = "protocolDirEntry"
    protocolDirEntry.EntityData.BundleName = "cisco_ios_xe"
    protocolDirEntry.EntityData.ParentYangName = "protocolDirTable"
    protocolDirEntry.EntityData.SegmentPath = "protocolDirEntry" + types.AddKeyToken(protocolDirEntry.ProtocolDirID, "protocolDirID") + types.AddKeyToken(protocolDirEntry.ProtocolDirParameters, "protocolDirParameters")
    protocolDirEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/protocolDirTable/" + protocolDirEntry.EntityData.SegmentPath
    protocolDirEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDirEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDirEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDirEntry.EntityData.Children = types.NewOrderedMap()
    protocolDirEntry.EntityData.Leafs = types.NewOrderedMap()
    protocolDirEntry.EntityData.Leafs.Append("protocolDirID", types.YLeaf{"ProtocolDirID", protocolDirEntry.ProtocolDirID})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirParameters", types.YLeaf{"ProtocolDirParameters", protocolDirEntry.ProtocolDirParameters})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", protocolDirEntry.ProtocolDirLocalIndex})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirDescr", types.YLeaf{"ProtocolDirDescr", protocolDirEntry.ProtocolDirDescr})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirType", types.YLeaf{"ProtocolDirType", protocolDirEntry.ProtocolDirType})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirAddressMapConfig", types.YLeaf{"ProtocolDirAddressMapConfig", protocolDirEntry.ProtocolDirAddressMapConfig})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirHostConfig", types.YLeaf{"ProtocolDirHostConfig", protocolDirEntry.ProtocolDirHostConfig})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirMatrixConfig", types.YLeaf{"ProtocolDirMatrixConfig", protocolDirEntry.ProtocolDirMatrixConfig})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirOwner", types.YLeaf{"ProtocolDirOwner", protocolDirEntry.ProtocolDirOwner})
    protocolDirEntry.EntityData.Leafs.Append("protocolDirStatus", types.YLeaf{"ProtocolDirStatus", protocolDirEntry.ProtocolDirStatus})

    protocolDirEntry.EntityData.YListKeys = []string {"ProtocolDirID", "ProtocolDirParameters"}

    return &(protocolDirEntry.EntityData)
}

// RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig represents the addressMappingTable.
type RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig string

const (
    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig_notSupported RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig = "notSupported"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig_supportedOff RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig = "supportedOff"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig_supportedOn RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirAddressMapConfig = "supportedOn"
)

// RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig represents for both protocols.
type RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig string

const (
    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig_notSupported RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig = "notSupported"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig_supportedOff RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig = "supportedOff"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig_supportedOn RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirHostConfig = "supportedOn"
)

// RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig represents for both protocols.
type RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig string

const (
    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig_notSupported RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig = "notSupported"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig_supportedOff RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig = "supportedOff"

    RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig_supportedOn RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirMatrixConfig = "supportedOn"
)

// RMON2MIB_ProtocolDistControlTable
// Controls the setup of protocol type distribution statistics
// tables.
// 
// Implementations are encouraged to add an entry per monitored
// interface upon initialization so that a default collection
// of protocol statistics is available.
// 
// Rationale:
// This table controls collection of very basic statistics
// for any or all of the protocols detected on a given interface.
// An NMS can use this table to quickly determine bandwidth
// allocation utilized by different protocols.
// 
// A media-specific statistics collection could also
// be configured (e.g. etherStats, trPStats) to easily obtain
// total frame, octet, and droppedEvents for the same
// interface.
type RMON2MIB_ProtocolDistControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDistControlTable.  An example of the
    // indexing of this entry is      protocolDistControlDroppedFrames.7. The type
    // is slice of RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry.
    ProtocolDistControlEntry []*RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry
}

func (protocolDistControlTable *RMON2MIB_ProtocolDistControlTable) GetEntityData() *types.CommonEntityData {
    protocolDistControlTable.EntityData.YFilter = protocolDistControlTable.YFilter
    protocolDistControlTable.EntityData.YangName = "protocolDistControlTable"
    protocolDistControlTable.EntityData.BundleName = "cisco_ios_xe"
    protocolDistControlTable.EntityData.ParentYangName = "RMON2-MIB"
    protocolDistControlTable.EntityData.SegmentPath = "protocolDistControlTable"
    protocolDistControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + protocolDistControlTable.EntityData.SegmentPath
    protocolDistControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDistControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDistControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDistControlTable.EntityData.Children = types.NewOrderedMap()
    protocolDistControlTable.EntityData.Children.Append("protocolDistControlEntry", types.YChild{"ProtocolDistControlEntry", nil})
    for i := range protocolDistControlTable.ProtocolDistControlEntry {
        protocolDistControlTable.EntityData.Children.Append(types.GetSegmentPath(protocolDistControlTable.ProtocolDistControlEntry[i]), types.YChild{"ProtocolDistControlEntry", protocolDistControlTable.ProtocolDistControlEntry[i]})
    }
    protocolDistControlTable.EntityData.Leafs = types.NewOrderedMap()

    protocolDistControlTable.EntityData.YListKeys = []string {}

    return &(protocolDistControlTable.EntityData)
}

// RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry
// A conceptual row in the protocolDistControlTable.
// 
// An example of the indexing of this entry is
// 
// 
// 
// 
// 
// protocolDistControlDroppedFrames.7
type RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique index for this protocolDistControlEntry.
    // The type is interface{} with range: 1..65535.
    ProtocolDistControlIndex interface{}

    // The source of data for the this protocol distribution.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // protocolDistControlStatus object is equal to active(1). The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    ProtocolDistControlDataSource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.       This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    ProtocolDistControlDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    ProtocolDistControlCreateTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    ProtocolDistControlOwner interface{}

    // The status of this row.  An entry may not exist in the active state unless
    // all objects in the entry have an appropriate value.  If this object is not
    // equal to active(1), all associated entries in the protocolDistStatsTable
    // shall be deleted. The type is RowStatus.
    ProtocolDistControlStatus interface{}
}

func (protocolDistControlEntry *RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry) GetEntityData() *types.CommonEntityData {
    protocolDistControlEntry.EntityData.YFilter = protocolDistControlEntry.YFilter
    protocolDistControlEntry.EntityData.YangName = "protocolDistControlEntry"
    protocolDistControlEntry.EntityData.BundleName = "cisco_ios_xe"
    protocolDistControlEntry.EntityData.ParentYangName = "protocolDistControlTable"
    protocolDistControlEntry.EntityData.SegmentPath = "protocolDistControlEntry" + types.AddKeyToken(protocolDistControlEntry.ProtocolDistControlIndex, "protocolDistControlIndex")
    protocolDistControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/protocolDistControlTable/" + protocolDistControlEntry.EntityData.SegmentPath
    protocolDistControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDistControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDistControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDistControlEntry.EntityData.Children = types.NewOrderedMap()
    protocolDistControlEntry.EntityData.Leafs = types.NewOrderedMap()
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlIndex", types.YLeaf{"ProtocolDistControlIndex", protocolDistControlEntry.ProtocolDistControlIndex})
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlDataSource", types.YLeaf{"ProtocolDistControlDataSource", protocolDistControlEntry.ProtocolDistControlDataSource})
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlDroppedFrames", types.YLeaf{"ProtocolDistControlDroppedFrames", protocolDistControlEntry.ProtocolDistControlDroppedFrames})
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlCreateTime", types.YLeaf{"ProtocolDistControlCreateTime", protocolDistControlEntry.ProtocolDistControlCreateTime})
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlOwner", types.YLeaf{"ProtocolDistControlOwner", protocolDistControlEntry.ProtocolDistControlOwner})
    protocolDistControlEntry.EntityData.Leafs.Append("protocolDistControlStatus", types.YLeaf{"ProtocolDistControlStatus", protocolDistControlEntry.ProtocolDistControlStatus})

    protocolDistControlEntry.EntityData.YListKeys = []string {"ProtocolDistControlIndex"}

    return &(protocolDistControlEntry.EntityData)
}

// RMON2MIB_ProtocolDistStatsTable
// An entry is made in this table for every protocol in the
// 
// 
// 
// 
// 
// protocolDirTable which has been seen in at least one packet.
// Counters are updated in this table for every protocol type
// that is encountered when parsing a packet, but no counters are
// updated for packets with MAC-layer errors.
// 
// Note that if a protocolDirEntry is deleted, all associated
// entries in this table are removed.
type RMON2MIB_ProtocolDistStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDistStatsTable.  The index is composed of
    // the protocolDistControlIndex of the associated protocolDistControlEntry
    // followed by the protocolDirLocalIndex of the associated protocol that this
    // entry represents.  In other words, the index identifies the protocol
    // distribution an entry is a part of as well as the particular protocol that
    // it represents.  An example of the indexing of this entry is
    // protocolDistStatsPkts.1.18. The type is slice of
    // RMON2MIB_ProtocolDistStatsTable_ProtocolDistStatsEntry.
    ProtocolDistStatsEntry []*RMON2MIB_ProtocolDistStatsTable_ProtocolDistStatsEntry
}

func (protocolDistStatsTable *RMON2MIB_ProtocolDistStatsTable) GetEntityData() *types.CommonEntityData {
    protocolDistStatsTable.EntityData.YFilter = protocolDistStatsTable.YFilter
    protocolDistStatsTable.EntityData.YangName = "protocolDistStatsTable"
    protocolDistStatsTable.EntityData.BundleName = "cisco_ios_xe"
    protocolDistStatsTable.EntityData.ParentYangName = "RMON2-MIB"
    protocolDistStatsTable.EntityData.SegmentPath = "protocolDistStatsTable"
    protocolDistStatsTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + protocolDistStatsTable.EntityData.SegmentPath
    protocolDistStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDistStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDistStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDistStatsTable.EntityData.Children = types.NewOrderedMap()
    protocolDistStatsTable.EntityData.Children.Append("protocolDistStatsEntry", types.YChild{"ProtocolDistStatsEntry", nil})
    for i := range protocolDistStatsTable.ProtocolDistStatsEntry {
        protocolDistStatsTable.EntityData.Children.Append(types.GetSegmentPath(protocolDistStatsTable.ProtocolDistStatsEntry[i]), types.YChild{"ProtocolDistStatsEntry", protocolDistStatsTable.ProtocolDistStatsEntry[i]})
    }
    protocolDistStatsTable.EntityData.Leafs = types.NewOrderedMap()

    protocolDistStatsTable.EntityData.YListKeys = []string {}

    return &(protocolDistStatsTable.EntityData)
}

// RMON2MIB_ProtocolDistStatsTable_ProtocolDistStatsEntry
// A conceptual row in the protocolDistStatsTable.
// 
// The index is composed of the protocolDistControlIndex of the
// associated protocolDistControlEntry followed by the
// protocolDirLocalIndex of the associated protocol that this
// entry represents.  In other words, the index identifies the
// protocol distribution an entry is a part of as well as the
// particular protocol that it represents.
// 
// An example of the indexing of this entry is
// protocolDistStatsPkts.1.18
type RMON2MIB_ProtocolDistStatsTable_ProtocolDistStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_ProtocolDistControlTable_ProtocolDistControlEntry_ProtocolDistControlIndex
    ProtocolDistControlIndex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // The number of packets without errors received of this protocol type.  Note
    // that this is the number of link-layer packets, so if a single network-layer
    // packet is fragmented into several link-layer frames, this counter is
    // incremented several times. The type is interface{} with range:
    // 0..4294967295.
    ProtocolDistStatsPkts interface{}

    // The number of octets in packets received of this protocol type since it was
    // added to the protocolDistStatsTable (excluding framing bits but including
    // FCS octets), except for those octets in packets that contained errors. 
    // Note this doesn't count just those octets in the particular protocol
    // frames, but includes the entire packet that contained the protocol. The
    // type is interface{} with range: 0..4294967295.
    ProtocolDistStatsOctets interface{}
}

func (protocolDistStatsEntry *RMON2MIB_ProtocolDistStatsTable_ProtocolDistStatsEntry) GetEntityData() *types.CommonEntityData {
    protocolDistStatsEntry.EntityData.YFilter = protocolDistStatsEntry.YFilter
    protocolDistStatsEntry.EntityData.YangName = "protocolDistStatsEntry"
    protocolDistStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    protocolDistStatsEntry.EntityData.ParentYangName = "protocolDistStatsTable"
    protocolDistStatsEntry.EntityData.SegmentPath = "protocolDistStatsEntry" + types.AddKeyToken(protocolDistStatsEntry.ProtocolDistControlIndex, "protocolDistControlIndex") + types.AddKeyToken(protocolDistStatsEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex")
    protocolDistStatsEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/protocolDistStatsTable/" + protocolDistStatsEntry.EntityData.SegmentPath
    protocolDistStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolDistStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolDistStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolDistStatsEntry.EntityData.Children = types.NewOrderedMap()
    protocolDistStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    protocolDistStatsEntry.EntityData.Leafs.Append("protocolDistControlIndex", types.YLeaf{"ProtocolDistControlIndex", protocolDistStatsEntry.ProtocolDistControlIndex})
    protocolDistStatsEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", protocolDistStatsEntry.ProtocolDirLocalIndex})
    protocolDistStatsEntry.EntityData.Leafs.Append("protocolDistStatsPkts", types.YLeaf{"ProtocolDistStatsPkts", protocolDistStatsEntry.ProtocolDistStatsPkts})
    protocolDistStatsEntry.EntityData.Leafs.Append("protocolDistStatsOctets", types.YLeaf{"ProtocolDistStatsOctets", protocolDistStatsEntry.ProtocolDistStatsOctets})

    protocolDistStatsEntry.EntityData.YListKeys = []string {"ProtocolDistControlIndex", "ProtocolDirLocalIndex"}

    return &(protocolDistStatsEntry.EntityData)
}

// RMON2MIB_AddressMapControlTable
// A table to control the collection of network layer address to
// physical address to interface mappings.
// 
// Note that this is not like the typical RMON
// controlTable and dataTable in which each entry creates
// its own data table.  Each entry in this table enables the
// discovery of addresses on a new interface and the placement
// of address mappings into the central addressMapTable.
// 
// Implementations are encouraged to add an entry per monitored
// interface upon initialization so that a default collection
// of address mappings is available.
type RMON2MIB_AddressMapControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapControlTable.      An example of the
    // indexing of this entry is addressMapControlDroppedFrames.1. The type is
    // slice of RMON2MIB_AddressMapControlTable_AddressMapControlEntry.
    AddressMapControlEntry []*RMON2MIB_AddressMapControlTable_AddressMapControlEntry
}

func (addressMapControlTable *RMON2MIB_AddressMapControlTable) GetEntityData() *types.CommonEntityData {
    addressMapControlTable.EntityData.YFilter = addressMapControlTable.YFilter
    addressMapControlTable.EntityData.YangName = "addressMapControlTable"
    addressMapControlTable.EntityData.BundleName = "cisco_ios_xe"
    addressMapControlTable.EntityData.ParentYangName = "RMON2-MIB"
    addressMapControlTable.EntityData.SegmentPath = "addressMapControlTable"
    addressMapControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + addressMapControlTable.EntityData.SegmentPath
    addressMapControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressMapControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressMapControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressMapControlTable.EntityData.Children = types.NewOrderedMap()
    addressMapControlTable.EntityData.Children.Append("addressMapControlEntry", types.YChild{"AddressMapControlEntry", nil})
    for i := range addressMapControlTable.AddressMapControlEntry {
        addressMapControlTable.EntityData.Children.Append(types.GetSegmentPath(addressMapControlTable.AddressMapControlEntry[i]), types.YChild{"AddressMapControlEntry", addressMapControlTable.AddressMapControlEntry[i]})
    }
    addressMapControlTable.EntityData.Leafs = types.NewOrderedMap()

    addressMapControlTable.EntityData.YListKeys = []string {}

    return &(addressMapControlTable.EntityData)
}

// RMON2MIB_AddressMapControlTable_AddressMapControlEntry
// A conceptual row in the addressMapControlTable.
// 
// 
// 
// 
// 
// An example of the indexing of this entry is
// addressMapControlDroppedFrames.1
type RMON2MIB_AddressMapControlTable_AddressMapControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique index for this entry in the
    // addressMapControlTable. The type is interface{} with range: 1..65535.
    AddressMapControlIndex interface{}

    // The source of data for this addressMapControlEntry. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    AddressMapControlDataSource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    AddressMapControlDroppedFrames interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    AddressMapControlOwner interface{}

    // The status of this addressMap control entry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // addressMapTable shall be deleted. The type is RowStatus.
    AddressMapControlStatus interface{}
}

func (addressMapControlEntry *RMON2MIB_AddressMapControlTable_AddressMapControlEntry) GetEntityData() *types.CommonEntityData {
    addressMapControlEntry.EntityData.YFilter = addressMapControlEntry.YFilter
    addressMapControlEntry.EntityData.YangName = "addressMapControlEntry"
    addressMapControlEntry.EntityData.BundleName = "cisco_ios_xe"
    addressMapControlEntry.EntityData.ParentYangName = "addressMapControlTable"
    addressMapControlEntry.EntityData.SegmentPath = "addressMapControlEntry" + types.AddKeyToken(addressMapControlEntry.AddressMapControlIndex, "addressMapControlIndex")
    addressMapControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/addressMapControlTable/" + addressMapControlEntry.EntityData.SegmentPath
    addressMapControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressMapControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressMapControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressMapControlEntry.EntityData.Children = types.NewOrderedMap()
    addressMapControlEntry.EntityData.Leafs = types.NewOrderedMap()
    addressMapControlEntry.EntityData.Leafs.Append("addressMapControlIndex", types.YLeaf{"AddressMapControlIndex", addressMapControlEntry.AddressMapControlIndex})
    addressMapControlEntry.EntityData.Leafs.Append("addressMapControlDataSource", types.YLeaf{"AddressMapControlDataSource", addressMapControlEntry.AddressMapControlDataSource})
    addressMapControlEntry.EntityData.Leafs.Append("addressMapControlDroppedFrames", types.YLeaf{"AddressMapControlDroppedFrames", addressMapControlEntry.AddressMapControlDroppedFrames})
    addressMapControlEntry.EntityData.Leafs.Append("addressMapControlOwner", types.YLeaf{"AddressMapControlOwner", addressMapControlEntry.AddressMapControlOwner})
    addressMapControlEntry.EntityData.Leafs.Append("addressMapControlStatus", types.YLeaf{"AddressMapControlStatus", addressMapControlEntry.AddressMapControlStatus})

    addressMapControlEntry.EntityData.YListKeys = []string {"AddressMapControlIndex"}

    return &(addressMapControlEntry.EntityData)
}

// RMON2MIB_AddressMapTable
// A table of network layer address to physical address to
// interface mappings.
// 
// The probe will add entries to this table based on the source
// MAC and network addresses seen in packets without MAC-level
// errors. The probe will populate this table for all protocols
// in the protocol directory table whose value of
// protocolDirAddressMapConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirAddressMapConfig value of supportedOff(2).
type RMON2MIB_AddressMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapTable. The protocolDirLocalIndex in the
    // index identifies the network layer protocol of the
    // addressMapNetworkAddress.      An example of the indexing of this entry is
    // addressMapSource.783495.18.4.128.2.6.6.11.1.3.6.1.2.1.2.2.1.1.1. The type
    // is slice of RMON2MIB_AddressMapTable_AddressMapEntry.
    AddressMapEntry []*RMON2MIB_AddressMapTable_AddressMapEntry
}

func (addressMapTable *RMON2MIB_AddressMapTable) GetEntityData() *types.CommonEntityData {
    addressMapTable.EntityData.YFilter = addressMapTable.YFilter
    addressMapTable.EntityData.YangName = "addressMapTable"
    addressMapTable.EntityData.BundleName = "cisco_ios_xe"
    addressMapTable.EntityData.ParentYangName = "RMON2-MIB"
    addressMapTable.EntityData.SegmentPath = "addressMapTable"
    addressMapTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + addressMapTable.EntityData.SegmentPath
    addressMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressMapTable.EntityData.Children = types.NewOrderedMap()
    addressMapTable.EntityData.Children.Append("addressMapEntry", types.YChild{"AddressMapEntry", nil})
    for i := range addressMapTable.AddressMapEntry {
        addressMapTable.EntityData.Children.Append(types.GetSegmentPath(addressMapTable.AddressMapEntry[i]), types.YChild{"AddressMapEntry", addressMapTable.AddressMapEntry[i]})
    }
    addressMapTable.EntityData.Leafs = types.NewOrderedMap()

    addressMapTable.EntityData.YListKeys = []string {}

    return &(addressMapTable.EntityData)
}

// RMON2MIB_AddressMapTable_AddressMapEntry
// A conceptual row in the addressMapTable.
// The protocolDirLocalIndex in the index identifies the network
// layer protocol of the addressMapNetworkAddress.
// 
// 
// 
// 
// 
// An example of the indexing of this entry is
// addressMapSource.783495.18.4.128.2.6.6.11.1.3.6.1.2.1.2.2.1.1.1
type RMON2MIB_AddressMapTable_AddressMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    AddressMapTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The network address for this relation.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the protocolDirLocalIndex component of the index.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    AddressMapNetworkAddress interface{}

    // This attribute is a key. The interface or port on which the associated
    // network address was most recently seen.      If this address mapping was
    // discovered on an interface, this object shall identify the instance of the
    // ifIndex object, defined in [3,5], for the desired interface. For example,
    // if an entry were to receive data from interface #1, this object would be
    // set to ifIndex.1.  If this address mapping was discovered on a port, this
    // object shall identify the instance of the rptrGroupPortIndex object,
    // defined in [RFC1516], for the desired port. For example, if an entry were
    // to receive data from group #1, port #1, this object would be set to
    // rptrGroupPortIndex.1.1.  Note that while the dataSource associated with
    // this entry may only point to index objects, this object may at times point
    // to repeater port objects. This situation occurs when the dataSource points
    // to an interface which is a locally attached repeater and the agent has
    // additional information about the source port of traffic seen on that
    // repeater. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    AddressMapSource interface{}

    // The last source physical address on which the associated network address
    // was seen.  If the protocol of the associated network address was
    // encapsulated inside of a network-level or higher protocol, this will be the
    // address of the next-lower protocol with the addressRecognitionCapable bit
    // enabled and will be formatted as specified for that protocol. The type is
    // string.
    AddressMapPhysicalAddress interface{}

    // The value of sysUpTime at the time this entry was last created or the
    // values of the physical address changed.  This can be used to help detect
    // duplicate address problems, in which case this object will be updated
    // frequently. The type is interface{} with range: 0..4294967295.
    AddressMapLastChange interface{}
}

func (addressMapEntry *RMON2MIB_AddressMapTable_AddressMapEntry) GetEntityData() *types.CommonEntityData {
    addressMapEntry.EntityData.YFilter = addressMapEntry.YFilter
    addressMapEntry.EntityData.YangName = "addressMapEntry"
    addressMapEntry.EntityData.BundleName = "cisco_ios_xe"
    addressMapEntry.EntityData.ParentYangName = "addressMapTable"
    addressMapEntry.EntityData.SegmentPath = "addressMapEntry" + types.AddKeyToken(addressMapEntry.AddressMapTimeMark, "addressMapTimeMark") + types.AddKeyToken(addressMapEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(addressMapEntry.AddressMapNetworkAddress, "addressMapNetworkAddress") + types.AddKeyToken(addressMapEntry.AddressMapSource, "addressMapSource")
    addressMapEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/addressMapTable/" + addressMapEntry.EntityData.SegmentPath
    addressMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressMapEntry.EntityData.Children = types.NewOrderedMap()
    addressMapEntry.EntityData.Leafs = types.NewOrderedMap()
    addressMapEntry.EntityData.Leafs.Append("addressMapTimeMark", types.YLeaf{"AddressMapTimeMark", addressMapEntry.AddressMapTimeMark})
    addressMapEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", addressMapEntry.ProtocolDirLocalIndex})
    addressMapEntry.EntityData.Leafs.Append("addressMapNetworkAddress", types.YLeaf{"AddressMapNetworkAddress", addressMapEntry.AddressMapNetworkAddress})
    addressMapEntry.EntityData.Leafs.Append("addressMapSource", types.YLeaf{"AddressMapSource", addressMapEntry.AddressMapSource})
    addressMapEntry.EntityData.Leafs.Append("addressMapPhysicalAddress", types.YLeaf{"AddressMapPhysicalAddress", addressMapEntry.AddressMapPhysicalAddress})
    addressMapEntry.EntityData.Leafs.Append("addressMapLastChange", types.YLeaf{"AddressMapLastChange", addressMapEntry.AddressMapLastChange})

    addressMapEntry.EntityData.YListKeys = []string {"AddressMapTimeMark", "ProtocolDirLocalIndex", "AddressMapNetworkAddress", "AddressMapSource"}

    return &(addressMapEntry.EntityData)
}

// RMON2MIB_HlHostControlTable
// A list of higher layer (i.e. non-MAC) host table control entries.
// 
// These entries will enable the collection of the network and
// application level host tables indexed by network addresses.
// Both the network and application level host tables are
// controlled by this table is so that they will both be created
// and deleted at the same time, further increasing the ease with
// which they can be implemented as a single datastore (note that
// if an implementation stores application layer host records in
// memory, it can derive network layer host records from them).
// 
// Entries in the nlHostTable will be created on behalf of each
// entry in this table. Additionally, if this probe implements
// the alHostTable, entries in the alHostTable will be created on
// behalf of each entry in this table.
// 
// Implementations are encouraged to add an entry per monitored
// interface upon initialization so that a default collection
// of host statistics is available.
type RMON2MIB_HlHostControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the hlHostControlTable.  An example of the indexing of
    // this entry is hlHostControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_HlHostControlTable_HlHostControlEntry.
    HlHostControlEntry []*RMON2MIB_HlHostControlTable_HlHostControlEntry
}

func (hlHostControlTable *RMON2MIB_HlHostControlTable) GetEntityData() *types.CommonEntityData {
    hlHostControlTable.EntityData.YFilter = hlHostControlTable.YFilter
    hlHostControlTable.EntityData.YangName = "hlHostControlTable"
    hlHostControlTable.EntityData.BundleName = "cisco_ios_xe"
    hlHostControlTable.EntityData.ParentYangName = "RMON2-MIB"
    hlHostControlTable.EntityData.SegmentPath = "hlHostControlTable"
    hlHostControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + hlHostControlTable.EntityData.SegmentPath
    hlHostControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlHostControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlHostControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlHostControlTable.EntityData.Children = types.NewOrderedMap()
    hlHostControlTable.EntityData.Children.Append("hlHostControlEntry", types.YChild{"HlHostControlEntry", nil})
    for i := range hlHostControlTable.HlHostControlEntry {
        hlHostControlTable.EntityData.Children.Append(types.GetSegmentPath(hlHostControlTable.HlHostControlEntry[i]), types.YChild{"HlHostControlEntry", hlHostControlTable.HlHostControlEntry[i]})
    }
    hlHostControlTable.EntityData.Leafs = types.NewOrderedMap()

    hlHostControlTable.EntityData.YListKeys = []string {}

    return &(hlHostControlTable.EntityData)
}

// RMON2MIB_HlHostControlTable_HlHostControlEntry
// A conceptual row in the hlHostControlTable.
// 
// An example of the indexing of this entry is
// hlHostControlNlDroppedFrames.1
type RMON2MIB_HlHostControlTable_HlHostControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hlHostControlTable.  Each such entry defines a function that discovers
    // hosts on a particular interface and places statistics about them in the
    // nlHostTable, and optionally in the alHostTable, on behalf of this
    // hlHostControlEntry. The type is interface{} with range: 1..65535.
    HlHostControlIndex interface{}

    // The source of data for the associated host tables.  The statistics in this
    // group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // hlHostControlStatus object is equal to active(1). The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    HlHostControlDataSource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for the associated      nlHost entries for whatever reason. 
    // Most often, this event occurs when the probe is out of some resources and
    // decides to shed load from this collection.  This count does not include
    // packets that were not counted because they had MAC-layer errors.  Note that
    // if the nlHostTable is inactive because no protocols are enabled in the
    // protocol directory, this value should be 0.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    HlHostControlNlDroppedFrames interface{}

    // The number of times an nlHost entry has been inserted into the nlHost
    // table.  If an entry is inserted, then deleted, and then inserted, this
    // counter will be incremented by 2.  To allow for efficient implementation
    // strategies, agents may delay updating this object for short periods of
    // time.  For example, an implementation strategy may allow internal data
    // structures to differ from those visible via SNMP for short periods of time.
    // This counter may reflect the internal data structures for those short
    // periods of time.  Note that the table size can be determined by subtracting
    // hlHostControlNlDeletes from hlHostControlNlInserts. The type is interface{}
    // with range: 0..4294967295.
    HlHostControlNlInserts interface{}

    // The number of times an nlHost entry has been deleted from the nlHost table
    // (for any reason).  If an entry is deleted, then inserted, and then deleted,
    // this counter will be incremented by 2.  To allow for efficient
    // implementation strategies, agents may delay updating this object for short
    // periods of time.  For example, an implementation strategy may allow
    // internal      data structures to differ from those visible via SNMP for
    // short periods of time.  This counter may reflect the internal data
    // structures for those short periods of time.  Note that the table size can
    // be determined by subtracting hlHostControlNlDeletes from
    // hlHostControlNlInserts. The type is interface{} with range: 0..4294967295.
    HlHostControlNlDeletes interface{}

    // The maximum number of entries that are desired in the nlHostTable on behalf
    // of this control entry. The probe will not create more than this number of
    // associated entries in the table, but may choose to create fewer entries in
    // this table for any reason including the lack of resources.  If this object
    // is set to a value less than the current number of entries, enough entries
    // are chosen in an implementation-dependent manner and deleted so that the
    // number of entries in the table equals the value of this object.  If this
    // value is set to -1, the probe may create any number of entries in this
    // table.  If the associated hlHostControlStatus object is equal to `active',
    // this object may not be modified.  This object may be used to control how
    // resources are allocated on the probe for the various RMON functions. The
    // type is interface{} with range: -1..2147483647.
    HlHostControlNlMaxDesiredEntries interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for the associated alHost entries for whatever reason.  Most
    // often, this event occurs when the probe is out of some resources and
    // decides to shed load from this collection.  This count does not include
    // packets that were not counted because they had MAC-layer errors.       Note
    // that if the alHostTable is not implemented or is inactive because no
    // protocols are enabled in the protocol directory, this value should be 0. 
    // Note that, unlike the dropEvents counter, this number is the exact number
    // of frames dropped. The type is interface{} with range: 0..4294967295.
    HlHostControlAlDroppedFrames interface{}

    // The number of times an alHost entry has been inserted into the alHost
    // table.  If an entry is inserted, then deleted, and then inserted, this
    // counter will be incremented by 2.  To allow for efficient implementation
    // strategies, agents may delay updating this object for short periods of
    // time.  For example, an implementation strategy may allow internal data
    // structures to differ from those visible via SNMP for short periods of time.
    // This counter may reflect the internal data structures for those short
    // periods of time.  Note that the table size can be determined by subtracting
    // hlHostControlAlDeletes from hlHostControlAlInserts. The type is interface{}
    // with range: 0..4294967295.
    HlHostControlAlInserts interface{}

    // The number of times an alHost entry has been deleted from the alHost table
    // (for any reason).  If an entry is deleted, then inserted, and then deleted,
    // this counter will be incremented by 2.  To allow for efficient
    // implementation strategies, agents may delay updating this object for short
    // periods of time.  For example, an implementation strategy may allow
    // internal data structures to differ from those visible via SNMP for short
    // periods of time.  This counter may reflect the internal data structures for
    // those short periods of time.  Note that the table size can be determined by
    // subtracting hlHostControlAlDeletes from hlHostControlAlInserts. The type is
    // interface{} with range: 0..4294967295.
    HlHostControlAlDeletes interface{}

    // The maximum number of entries that are desired in the alHost table on
    // behalf of this control entry. The probe will not create more than this
    // number of associated entries in the table, but may choose to create fewer
    // entries in this table for any reason including the lack of resources.  If
    // this object is set to a value less than the current number of entries,
    // enough entries are chosen in an implementation-dependent manner and deleted
    // so that the number of entries in the table equals the value of this object.
    // If this value is set to -1, the probe may create any number of entries in
    // this table.  If the associated hlHostControlStatus object is equal to
    // `active', this object may not be modified.  This object may be used to
    // control how resources are allocated on the probe for the various RMON
    // functions. The type is interface{} with range: -1..2147483647.
    HlHostControlAlMaxDesiredEntries interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    HlHostControlOwner interface{}

    // The status of this hlHostControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // nlHostTable and alHostTable shall be deleted. The type is RowStatus.
    HlHostControlStatus interface{}
}

func (hlHostControlEntry *RMON2MIB_HlHostControlTable_HlHostControlEntry) GetEntityData() *types.CommonEntityData {
    hlHostControlEntry.EntityData.YFilter = hlHostControlEntry.YFilter
    hlHostControlEntry.EntityData.YangName = "hlHostControlEntry"
    hlHostControlEntry.EntityData.BundleName = "cisco_ios_xe"
    hlHostControlEntry.EntityData.ParentYangName = "hlHostControlTable"
    hlHostControlEntry.EntityData.SegmentPath = "hlHostControlEntry" + types.AddKeyToken(hlHostControlEntry.HlHostControlIndex, "hlHostControlIndex")
    hlHostControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/hlHostControlTable/" + hlHostControlEntry.EntityData.SegmentPath
    hlHostControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlHostControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlHostControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlHostControlEntry.EntityData.Children = types.NewOrderedMap()
    hlHostControlEntry.EntityData.Leafs = types.NewOrderedMap()
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlIndex", types.YLeaf{"HlHostControlIndex", hlHostControlEntry.HlHostControlIndex})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlDataSource", types.YLeaf{"HlHostControlDataSource", hlHostControlEntry.HlHostControlDataSource})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlNlDroppedFrames", types.YLeaf{"HlHostControlNlDroppedFrames", hlHostControlEntry.HlHostControlNlDroppedFrames})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlNlInserts", types.YLeaf{"HlHostControlNlInserts", hlHostControlEntry.HlHostControlNlInserts})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlNlDeletes", types.YLeaf{"HlHostControlNlDeletes", hlHostControlEntry.HlHostControlNlDeletes})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlNlMaxDesiredEntries", types.YLeaf{"HlHostControlNlMaxDesiredEntries", hlHostControlEntry.HlHostControlNlMaxDesiredEntries})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlAlDroppedFrames", types.YLeaf{"HlHostControlAlDroppedFrames", hlHostControlEntry.HlHostControlAlDroppedFrames})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlAlInserts", types.YLeaf{"HlHostControlAlInserts", hlHostControlEntry.HlHostControlAlInserts})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlAlDeletes", types.YLeaf{"HlHostControlAlDeletes", hlHostControlEntry.HlHostControlAlDeletes})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlAlMaxDesiredEntries", types.YLeaf{"HlHostControlAlMaxDesiredEntries", hlHostControlEntry.HlHostControlAlMaxDesiredEntries})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlOwner", types.YLeaf{"HlHostControlOwner", hlHostControlEntry.HlHostControlOwner})
    hlHostControlEntry.EntityData.Leafs.Append("hlHostControlStatus", types.YLeaf{"HlHostControlStatus", hlHostControlEntry.HlHostControlStatus})

    hlHostControlEntry.EntityData.YListKeys = []string {"HlHostControlIndex"}

    return &(hlHostControlEntry.EntityData)
}

// RMON2MIB_NlHostTable
// A collection of statistics for a particular network layer
// address that has been discovered on an interface of this
// device.
// 
// The probe will populate this table for all network layer
// protocols in the protocol directory table whose value of
// protocolDirHostConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirHostConfig value of supportedOff(2).
// 
// The probe will add to this table all addresses seen
// as the source or destination address in all packets with no
// MAC errors, and will increment octet and packet counts in the
// table for all packets with no MAC errors.
type RMON2MIB_NlHostTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlHostTable.  The hlHostControlIndex value in the
    // index identifies the hlHostControlEntry on whose behalf this entry was
    // created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlHostAddress.  An example of the indexing of
    // this entry is nlHostOutPkts.1.783495.18.4.128.2.6.6. The type is slice of
    // RMON2MIB_NlHostTable_NlHostEntry.
    NlHostEntry []*RMON2MIB_NlHostTable_NlHostEntry
}

func (nlHostTable *RMON2MIB_NlHostTable) GetEntityData() *types.CommonEntityData {
    nlHostTable.EntityData.YFilter = nlHostTable.YFilter
    nlHostTable.EntityData.YangName = "nlHostTable"
    nlHostTable.EntityData.BundleName = "cisco_ios_xe"
    nlHostTable.EntityData.ParentYangName = "RMON2-MIB"
    nlHostTable.EntityData.SegmentPath = "nlHostTable"
    nlHostTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + nlHostTable.EntityData.SegmentPath
    nlHostTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlHostTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlHostTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlHostTable.EntityData.Children = types.NewOrderedMap()
    nlHostTable.EntityData.Children.Append("nlHostEntry", types.YChild{"NlHostEntry", nil})
    for i := range nlHostTable.NlHostEntry {
        nlHostTable.EntityData.Children.Append(types.GetSegmentPath(nlHostTable.NlHostEntry[i]), types.YChild{"NlHostEntry", nlHostTable.NlHostEntry[i]})
    }
    nlHostTable.EntityData.Leafs = types.NewOrderedMap()

    nlHostTable.EntityData.YListKeys = []string {}

    return &(nlHostTable.EntityData)
}

// RMON2MIB_NlHostTable_NlHostEntry
// A conceptual row in the nlHostTable.
// 
// The hlHostControlIndex value in the index identifies the
// hlHostControlEntry on whose behalf this entry was created.
// The protocolDirLocalIndex value in the index identifies the
// network layer protocol of the nlHostAddress.
// 
// An example of the indexing of this entry is
// nlHostOutPkts.1.783495.18.4.128.2.6.6.
type RMON2MIB_NlHostTable_NlHostEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlHostControlTable_HlHostControlEntry_HlHostControlIndex
    HlHostControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    NlHostTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The network address for this nlHostEntry.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the protocolDirLocalIndex component of the index.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    NlHostAddress interface{}

    // The number of packets without errors transmitted to this address since it
    // was added to the nlHostTable.  Note that this is the number of link-layer
    // packets, so if a single network-layer packet is fragmented into several
    // link-layer frames, this counter is incremented several times. The type is
    // interface{} with range: 0..4294967295.
    NlHostInPkts interface{}

    // The number of packets without errors transmitted by      this address since
    // it was added to the nlHostTable.  Note that this is the number of
    // link-layer packets, so if a single network-layer packet is fragmented into
    // several link-layer frames, this counter is incremented several times. The
    // type is interface{} with range: 0..4294967295.
    NlHostOutPkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // nlHostTable (excluding framing bits but including FCS octets), excluding
    // those octets in packets that contained errors.  Note this doesn't count
    // just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    NlHostInOctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // nlHostTable (excluding framing bits but including FCS octets), excluding
    // those octets in packets that contained errors.  Note this doesn't count
    // just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    NlHostOutOctets interface{}

    // The number of packets without errors transmitted by this address that were
    // directed to any MAC broadcast addresses or to any MAC multicast addresses
    // since this host was added to the nlHostTable. Note that this is the number
    // of link-layer packets, so if a single network-layer packet is fragmented
    // into several link-layer frames, this counter is incremented several times.
    // The type is interface{} with range: 0..4294967295.
    NlHostOutMacNonUnicastPkts interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    NlHostCreateTime interface{}
}

func (nlHostEntry *RMON2MIB_NlHostTable_NlHostEntry) GetEntityData() *types.CommonEntityData {
    nlHostEntry.EntityData.YFilter = nlHostEntry.YFilter
    nlHostEntry.EntityData.YangName = "nlHostEntry"
    nlHostEntry.EntityData.BundleName = "cisco_ios_xe"
    nlHostEntry.EntityData.ParentYangName = "nlHostTable"
    nlHostEntry.EntityData.SegmentPath = "nlHostEntry" + types.AddKeyToken(nlHostEntry.HlHostControlIndex, "hlHostControlIndex") + types.AddKeyToken(nlHostEntry.NlHostTimeMark, "nlHostTimeMark") + types.AddKeyToken(nlHostEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(nlHostEntry.NlHostAddress, "nlHostAddress")
    nlHostEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/nlHostTable/" + nlHostEntry.EntityData.SegmentPath
    nlHostEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlHostEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlHostEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlHostEntry.EntityData.Children = types.NewOrderedMap()
    nlHostEntry.EntityData.Leafs = types.NewOrderedMap()
    nlHostEntry.EntityData.Leafs.Append("hlHostControlIndex", types.YLeaf{"HlHostControlIndex", nlHostEntry.HlHostControlIndex})
    nlHostEntry.EntityData.Leafs.Append("nlHostTimeMark", types.YLeaf{"NlHostTimeMark", nlHostEntry.NlHostTimeMark})
    nlHostEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", nlHostEntry.ProtocolDirLocalIndex})
    nlHostEntry.EntityData.Leafs.Append("nlHostAddress", types.YLeaf{"NlHostAddress", nlHostEntry.NlHostAddress})
    nlHostEntry.EntityData.Leafs.Append("nlHostInPkts", types.YLeaf{"NlHostInPkts", nlHostEntry.NlHostInPkts})
    nlHostEntry.EntityData.Leafs.Append("nlHostOutPkts", types.YLeaf{"NlHostOutPkts", nlHostEntry.NlHostOutPkts})
    nlHostEntry.EntityData.Leafs.Append("nlHostInOctets", types.YLeaf{"NlHostInOctets", nlHostEntry.NlHostInOctets})
    nlHostEntry.EntityData.Leafs.Append("nlHostOutOctets", types.YLeaf{"NlHostOutOctets", nlHostEntry.NlHostOutOctets})
    nlHostEntry.EntityData.Leafs.Append("nlHostOutMacNonUnicastPkts", types.YLeaf{"NlHostOutMacNonUnicastPkts", nlHostEntry.NlHostOutMacNonUnicastPkts})
    nlHostEntry.EntityData.Leafs.Append("nlHostCreateTime", types.YLeaf{"NlHostCreateTime", nlHostEntry.NlHostCreateTime})

    nlHostEntry.EntityData.YListKeys = []string {"HlHostControlIndex", "NlHostTimeMark", "ProtocolDirLocalIndex", "NlHostAddress"}

    return &(nlHostEntry.EntityData)
}

// RMON2MIB_HlMatrixControlTable
// A list of higher layer (i.e. non-MAC) matrix control entries.
// 
// These entries will enable the collection of the network and
// application level matrix tables containing conversation
// statistics indexed by pairs of network addresses.
// Both the network and application level matrix tables are
// controlled by this table is so that they will both be created
// and deleted at the same time, further increasing the ease with
// which they can be implemented as a single datastore (note that
// if an implementation stores application layer matrix records
// in memory, it can derive network layer matrix records from
// them).
// 
// Entries in the nlMatrixSDTable and nlMatrixDSTable will be
// created on behalf of each entry in this table.  Additionally,
// if this probe implements the alMatrix tables, entries in the
// alMatrix tables will be created on behalf of each entry in
// this table.
type RMON2MIB_HlMatrixControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the hlMatrixControlTable.  An example of indexing of
    // this entry is hlMatrixControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry.
    HlMatrixControlEntry []*RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry
}

func (hlMatrixControlTable *RMON2MIB_HlMatrixControlTable) GetEntityData() *types.CommonEntityData {
    hlMatrixControlTable.EntityData.YFilter = hlMatrixControlTable.YFilter
    hlMatrixControlTable.EntityData.YangName = "hlMatrixControlTable"
    hlMatrixControlTable.EntityData.BundleName = "cisco_ios_xe"
    hlMatrixControlTable.EntityData.ParentYangName = "RMON2-MIB"
    hlMatrixControlTable.EntityData.SegmentPath = "hlMatrixControlTable"
    hlMatrixControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + hlMatrixControlTable.EntityData.SegmentPath
    hlMatrixControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlMatrixControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlMatrixControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlMatrixControlTable.EntityData.Children = types.NewOrderedMap()
    hlMatrixControlTable.EntityData.Children.Append("hlMatrixControlEntry", types.YChild{"HlMatrixControlEntry", nil})
    for i := range hlMatrixControlTable.HlMatrixControlEntry {
        hlMatrixControlTable.EntityData.Children.Append(types.GetSegmentPath(hlMatrixControlTable.HlMatrixControlEntry[i]), types.YChild{"HlMatrixControlEntry", hlMatrixControlTable.HlMatrixControlEntry[i]})
    }
    hlMatrixControlTable.EntityData.Leafs = types.NewOrderedMap()

    hlMatrixControlTable.EntityData.YListKeys = []string {}

    return &(hlMatrixControlTable.EntityData)
}

// RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry
// A conceptual row in the hlMatrixControlTable.
// 
// An example of indexing of this entry is
// hlMatrixControlNlDroppedFrames.1
type RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hlMatrixControlTable.  Each such entry defines a function that discovers
    // conversations on a particular interface and places statistics about them in
    // the nlMatrixSDTable and the nlMatrixDSTable, and optionally the
    // alMatrixSDTable and alMatrixDSTable, on behalf of this
    // hlMatrixControlEntry. The type is interface{} with range: 1..65535.
    HlMatrixControlIndex interface{}

    // The source of the data for the associated matrix tables.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // hlMatrixControlStatus object is equal to active(1). The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    HlMatrixControlDataSource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that if the nlMatrixTables
    // are inactive because no protocols are enabled in the protocol directory,
    // this value should be 0.  Note that, unlike the dropEvents counter, this
    // number is the exact number of frames dropped. The type is interface{} with
    // range: 0..4294967295.
    HlMatrixControlNlDroppedFrames interface{}

    // The number of times an nlMatrix entry has been inserted into the nlMatrix
    // tables.  If an entry is inserted, then deleted, and then inserted, this
    // counter will be incremented by 2.  The addition of a conversation into both
    // the nlMatrixSDTable and nlMatrixDSTable shall be counted as two insertions
    // (even though every addition into one table must be accompanied by an
    // insertion into the other).  To allow for efficient implementation
    // strategies, agents may delay updating this object for short periods of
    // time.  For example, an implementation strategy may allow internal data
    // structures to differ from those visible via SNMP for short periods of time.
    // This counter may reflect the internal data structures for those short
    // periods of time.      Note that the sum of then nlMatrixSDTable and
    // nlMatrixDSTable sizes can be determined by subtracting
    // hlMatrixControlNlDeletes from hlMatrixControlNlInserts. The type is
    // interface{} with range: 0..4294967295.
    HlMatrixControlNlInserts interface{}

    // The number of times an nlMatrix entry has been deleted from the nlMatrix
    // tables (for any reason).  If an entry is deleted, then inserted, and then
    // deleted, this counter will be incremented by 2.  The deletion of a
    // conversation from both the nlMatrixSDTable and nlMatrixDSTable shall be
    // counted as two deletions (even though every deletion from one table must be
    // accompanied by a deletion from the other).  To allow for efficient
    // implementation strategies, agents may delay updating this object for short
    // periods of time.  For example, an implementation strategy may allow
    // internal data structures to differ from those visible via SNMP for short
    // periods of time.  This counter may reflect the internal data structures for
    // those short periods of time.  Note that the table size can be determined by
    // subtracting hlMatrixControlNlDeletes from hlMatrixControlNlInserts. The
    // type is interface{} with range: 0..4294967295.
    HlMatrixControlNlDeletes interface{}

    // The maximum number of entries that are desired in the nlMatrix tables on
    // behalf of this control entry. The probe will not create more than this
    // number of associated entries in the table, but may choose to create fewer
    // entries in this table for any reason including the lack of resources.  If
    // this object is set to a value less than the current number of entries,
    // enough entries are chosen in an implementation-dependent manner and deleted
    // so that the number of entries in the table equals the value of this object.
    // If this value is set to -1, the probe may create any number of entries in
    // this table.  If the associated      hlMatrixControlStatus object is equal
    // to `active', this object may not be modified.  This object may be used to
    // control how resources are allocated on the probe for the various RMON
    // functions. The type is interface{} with range: -1..2147483647.
    HlMatrixControlNlMaxDesiredEntries interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that if the alMatrixTables
    // are not implemented or are inactive because no protocols are enabled in the
    // protocol directory, this value should be 0.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    HlMatrixControlAlDroppedFrames interface{}

    // The number of times an alMatrix entry has been inserted into the alMatrix
    // tables.  If an entry is inserted, then deleted, and then inserted, this
    // counter will be incremented by 2.  The addition of a conversation into both
    // the alMatrixSDTable and alMatrixDSTable shall be counted as two insertions
    // (even though every addition into one table must be accompanied by an
    // insertion into the other).  To allow for efficient implementation
    // strategies, agents may delay updating this object for short periods of
    // time.  For example, an implementation strategy may allow internal data
    // structures to differ from those visible via SNMP for short periods of time.
    // This counter may reflect the internal      data structures for those short
    // periods of time.  Note that the table size can be determined by subtracting
    // hlMatrixControlAlDeletes from hlMatrixControlAlInserts. The type is
    // interface{} with range: 0..4294967295.
    HlMatrixControlAlInserts interface{}

    // The number of times an alMatrix entry has been deleted from the alMatrix
    // tables.  If an entry is deleted, then inserted, and then deleted, this
    // counter will be incremented by 2.  The deletion of a conversation from both
    // the alMatrixSDTable and alMatrixDSTable shall be counted as two deletions
    // (even though every deletion from one table must be accompanied by a
    // deletion from the other).  To allow for efficient implementation
    // strategies, agents may delay updating this object for short periods of
    // time.  For example, an implementation strategy may allow internal data
    // structures to differ from those visible via SNMP for short periods of time.
    // This counter may reflect the internal data structures for those short
    // periods of time.  Note that the table size can be determined by subtracting
    // hlMatrixControlAlDeletes from hlMatrixControlAlInserts. The type is
    // interface{} with range: 0..4294967295.
    HlMatrixControlAlDeletes interface{}

    // The maximum number of entries that are desired in the alMatrix tables on
    // behalf of this control entry. The probe will not create more than this
    // number of associated entries in the table, but may choose to create fewer
    // entries in this table for any reason including the lack of resources.  If
    // this object is set to a value less than the current number of entries,
    // enough entries are chosen in an implementation-dependent manner and deleted
    // so that the number of entries in the table equals the value of this object.
    // If this value is set to -1, the probe may create any number of entries in
    // this table.  If the associated      hlMatrixControlStatus object is equal
    // to `active', this object may not be modified.  This object may be used to
    // control how resources are allocated on the probe for the various RMON
    // functions. The type is interface{} with range: -1..2147483647.
    HlMatrixControlAlMaxDesiredEntries interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    HlMatrixControlOwner interface{}

    // The status of this hlMatrixControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // nlMatrixSDTable, nlMatrixDSTable, alMatrixSDTable, and the alMatrixDSTable
    // shall be deleted by the agent. The type is RowStatus.
    HlMatrixControlStatus interface{}
}

func (hlMatrixControlEntry *RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry) GetEntityData() *types.CommonEntityData {
    hlMatrixControlEntry.EntityData.YFilter = hlMatrixControlEntry.YFilter
    hlMatrixControlEntry.EntityData.YangName = "hlMatrixControlEntry"
    hlMatrixControlEntry.EntityData.BundleName = "cisco_ios_xe"
    hlMatrixControlEntry.EntityData.ParentYangName = "hlMatrixControlTable"
    hlMatrixControlEntry.EntityData.SegmentPath = "hlMatrixControlEntry" + types.AddKeyToken(hlMatrixControlEntry.HlMatrixControlIndex, "hlMatrixControlIndex")
    hlMatrixControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/hlMatrixControlTable/" + hlMatrixControlEntry.EntityData.SegmentPath
    hlMatrixControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlMatrixControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlMatrixControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlMatrixControlEntry.EntityData.Children = types.NewOrderedMap()
    hlMatrixControlEntry.EntityData.Leafs = types.NewOrderedMap()
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlIndex", types.YLeaf{"HlMatrixControlIndex", hlMatrixControlEntry.HlMatrixControlIndex})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlDataSource", types.YLeaf{"HlMatrixControlDataSource", hlMatrixControlEntry.HlMatrixControlDataSource})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlNlDroppedFrames", types.YLeaf{"HlMatrixControlNlDroppedFrames", hlMatrixControlEntry.HlMatrixControlNlDroppedFrames})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlNlInserts", types.YLeaf{"HlMatrixControlNlInserts", hlMatrixControlEntry.HlMatrixControlNlInserts})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlNlDeletes", types.YLeaf{"HlMatrixControlNlDeletes", hlMatrixControlEntry.HlMatrixControlNlDeletes})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlNlMaxDesiredEntries", types.YLeaf{"HlMatrixControlNlMaxDesiredEntries", hlMatrixControlEntry.HlMatrixControlNlMaxDesiredEntries})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlAlDroppedFrames", types.YLeaf{"HlMatrixControlAlDroppedFrames", hlMatrixControlEntry.HlMatrixControlAlDroppedFrames})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlAlInserts", types.YLeaf{"HlMatrixControlAlInserts", hlMatrixControlEntry.HlMatrixControlAlInserts})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlAlDeletes", types.YLeaf{"HlMatrixControlAlDeletes", hlMatrixControlEntry.HlMatrixControlAlDeletes})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlAlMaxDesiredEntries", types.YLeaf{"HlMatrixControlAlMaxDesiredEntries", hlMatrixControlEntry.HlMatrixControlAlMaxDesiredEntries})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlOwner", types.YLeaf{"HlMatrixControlOwner", hlMatrixControlEntry.HlMatrixControlOwner})
    hlMatrixControlEntry.EntityData.Leafs.Append("hlMatrixControlStatus", types.YLeaf{"HlMatrixControlStatus", hlMatrixControlEntry.HlMatrixControlStatus})

    hlMatrixControlEntry.EntityData.YListKeys = []string {"HlMatrixControlIndex"}

    return &(hlMatrixControlEntry.EntityData)
}

// RMON2MIB_NlMatrixSDTable
// A list of traffic matrix entries which collect statistics for
// conversations between two network-level addresses.  This table
// is indexed first by the source address and then by the
// destination address to make it convenient to collect all
// conversations from a particular address.
// 
// The probe will populate this table for all network layer
// protocols in the protocol directory table whose value of
// protocolDirMatrixConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirMatrixConfig value of supportedOff(2).
// 
// 
// 
// 
// 
// The probe will add to this table all pairs of addresses
// seen in all packets with no MAC errors, and will increment
// octet and packet counts in the table for all packets with no
// MAC errors.
// 
// Further, this table will only contain entries that have a
// corresponding entry in the nlMatrixDSTable with the same
// source address and destination address.
type RMON2MIB_NlMatrixSDTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixSDTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlMatrixSDSourceAddress and
    // nlMatrixSDDestAddress.  An example of the indexing of this table is
    // nlMatrixSDPkts.1.783495.18.4.128.2.6.6.4.128.2.6.7. The type is slice of
    // RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry.
    NlMatrixSDEntry []*RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry
}

func (nlMatrixSDTable *RMON2MIB_NlMatrixSDTable) GetEntityData() *types.CommonEntityData {
    nlMatrixSDTable.EntityData.YFilter = nlMatrixSDTable.YFilter
    nlMatrixSDTable.EntityData.YangName = "nlMatrixSDTable"
    nlMatrixSDTable.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixSDTable.EntityData.ParentYangName = "RMON2-MIB"
    nlMatrixSDTable.EntityData.SegmentPath = "nlMatrixSDTable"
    nlMatrixSDTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + nlMatrixSDTable.EntityData.SegmentPath
    nlMatrixSDTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixSDTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixSDTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixSDTable.EntityData.Children = types.NewOrderedMap()
    nlMatrixSDTable.EntityData.Children.Append("nlMatrixSDEntry", types.YChild{"NlMatrixSDEntry", nil})
    for i := range nlMatrixSDTable.NlMatrixSDEntry {
        nlMatrixSDTable.EntityData.Children.Append(types.GetSegmentPath(nlMatrixSDTable.NlMatrixSDEntry[i]), types.YChild{"NlMatrixSDEntry", nlMatrixSDTable.NlMatrixSDEntry[i]})
    }
    nlMatrixSDTable.EntityData.Leafs = types.NewOrderedMap()

    nlMatrixSDTable.EntityData.YListKeys = []string {}

    return &(nlMatrixSDTable.EntityData)
}

// RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry
// A conceptual row in the nlMatrixSDTable.
// 
// The hlMatrixControlIndex value in the index identifies the
// hlMatrixControlEntry on whose behalf this entry was created.
// The protocolDirLocalIndex value in the index identifies the
// network layer protocol of the nlMatrixSDSourceAddress and
// nlMatrixSDDestAddress.
// 
// An example of the indexing of this table is
// nlMatrixSDPkts.1.783495.18.4.128.2.6.6.4.128.2.6.7
type RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry_HlMatrixControlIndex
    HlMatrixControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    NlMatrixSDTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The network source address for this
    // nlMatrixSDEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    NlMatrixSDSourceAddress interface{}

    // This attribute is a key. The network destination address for this
    // nlMatrixSDEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    NlMatrixSDDestAddress interface{}

    // The number of packets without errors transmitted from the source address to
    // the destination address since this entry was added to the nlMatrixSDTable. 
    // Note that this is the number of link-layer packets, so if a single
    // network-layer packet is fragmented into several link-layer frames, this
    // counter is incremented several times. The type is interface{} with range:
    // 0..4294967295.
    NlMatrixSDPkts interface{}

    // The number of octets transmitted from the source address to the destination
    // address since this entry was added to the nlMatrixSDTable (excluding
    // framing bits but including FCS octets), excluding those octets in packets
    // that contained errors.  Note this doesn't count just those octets in the
    // particular protocol frames, but includes the entire packet that contained
    // the protocol. The type is interface{} with range: 0..4294967295.
    NlMatrixSDOctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    NlMatrixSDCreateTime interface{}
}

func (nlMatrixSDEntry *RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry) GetEntityData() *types.CommonEntityData {
    nlMatrixSDEntry.EntityData.YFilter = nlMatrixSDEntry.YFilter
    nlMatrixSDEntry.EntityData.YangName = "nlMatrixSDEntry"
    nlMatrixSDEntry.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixSDEntry.EntityData.ParentYangName = "nlMatrixSDTable"
    nlMatrixSDEntry.EntityData.SegmentPath = "nlMatrixSDEntry" + types.AddKeyToken(nlMatrixSDEntry.HlMatrixControlIndex, "hlMatrixControlIndex") + types.AddKeyToken(nlMatrixSDEntry.NlMatrixSDTimeMark, "nlMatrixSDTimeMark") + types.AddKeyToken(nlMatrixSDEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(nlMatrixSDEntry.NlMatrixSDSourceAddress, "nlMatrixSDSourceAddress") + types.AddKeyToken(nlMatrixSDEntry.NlMatrixSDDestAddress, "nlMatrixSDDestAddress")
    nlMatrixSDEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/nlMatrixSDTable/" + nlMatrixSDEntry.EntityData.SegmentPath
    nlMatrixSDEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixSDEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixSDEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixSDEntry.EntityData.Children = types.NewOrderedMap()
    nlMatrixSDEntry.EntityData.Leafs = types.NewOrderedMap()
    nlMatrixSDEntry.EntityData.Leafs.Append("hlMatrixControlIndex", types.YLeaf{"HlMatrixControlIndex", nlMatrixSDEntry.HlMatrixControlIndex})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDTimeMark", types.YLeaf{"NlMatrixSDTimeMark", nlMatrixSDEntry.NlMatrixSDTimeMark})
    nlMatrixSDEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", nlMatrixSDEntry.ProtocolDirLocalIndex})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDSourceAddress", types.YLeaf{"NlMatrixSDSourceAddress", nlMatrixSDEntry.NlMatrixSDSourceAddress})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDDestAddress", types.YLeaf{"NlMatrixSDDestAddress", nlMatrixSDEntry.NlMatrixSDDestAddress})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDPkts", types.YLeaf{"NlMatrixSDPkts", nlMatrixSDEntry.NlMatrixSDPkts})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDOctets", types.YLeaf{"NlMatrixSDOctets", nlMatrixSDEntry.NlMatrixSDOctets})
    nlMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDCreateTime", types.YLeaf{"NlMatrixSDCreateTime", nlMatrixSDEntry.NlMatrixSDCreateTime})

    nlMatrixSDEntry.EntityData.YListKeys = []string {"HlMatrixControlIndex", "NlMatrixSDTimeMark", "ProtocolDirLocalIndex", "NlMatrixSDSourceAddress", "NlMatrixSDDestAddress"}

    return &(nlMatrixSDEntry.EntityData)
}

// RMON2MIB_NlMatrixDSTable
// A list of traffic matrix entries which collect statistics for
// conversations between two network-level addresses.  This table
// is indexed first by the destination address and then by the
// source address to make it convenient to collect all
// conversations to a particular address.
// 
// The probe will populate this table for all network layer
// protocols in the protocol directory table whose value of
// protocolDirMatrixConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirMatrixConfig value of supportedOff(2).
// 
// The probe will add to this table all pairs of addresses
// seen in all packets with no MAC errors, and will increment
// 
// 
// 
// 
// 
// octet and packet counts in the table for all packets with no
// MAC errors.
// 
// Further, this table will only contain entries that have a
// corresponding entry in the nlMatrixSDTable with the same
// source address and destination address.
type RMON2MIB_NlMatrixDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixDSTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlMatrixDSSourceAddress and
    // nlMatrixDSDestAddress.  An example of the indexing of this table is
    // nlMatrixDSPkts.1.783495.18.4.128.2.6.7.4.128.2.6.6. The type is slice of
    // RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry.
    NlMatrixDSEntry []*RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry
}

func (nlMatrixDSTable *RMON2MIB_NlMatrixDSTable) GetEntityData() *types.CommonEntityData {
    nlMatrixDSTable.EntityData.YFilter = nlMatrixDSTable.YFilter
    nlMatrixDSTable.EntityData.YangName = "nlMatrixDSTable"
    nlMatrixDSTable.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixDSTable.EntityData.ParentYangName = "RMON2-MIB"
    nlMatrixDSTable.EntityData.SegmentPath = "nlMatrixDSTable"
    nlMatrixDSTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + nlMatrixDSTable.EntityData.SegmentPath
    nlMatrixDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixDSTable.EntityData.Children = types.NewOrderedMap()
    nlMatrixDSTable.EntityData.Children.Append("nlMatrixDSEntry", types.YChild{"NlMatrixDSEntry", nil})
    for i := range nlMatrixDSTable.NlMatrixDSEntry {
        nlMatrixDSTable.EntityData.Children.Append(types.GetSegmentPath(nlMatrixDSTable.NlMatrixDSEntry[i]), types.YChild{"NlMatrixDSEntry", nlMatrixDSTable.NlMatrixDSEntry[i]})
    }
    nlMatrixDSTable.EntityData.Leafs = types.NewOrderedMap()

    nlMatrixDSTable.EntityData.YListKeys = []string {}

    return &(nlMatrixDSTable.EntityData)
}

// RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry
// A conceptual row in the nlMatrixDSTable.
// 
// The hlMatrixControlIndex value in the index identifies the
// hlMatrixControlEntry on whose behalf this entry was created.
// The protocolDirLocalIndex value in the index identifies the
// network layer protocol of the nlMatrixDSSourceAddress and
// nlMatrixDSDestAddress.
// 
// An example of the indexing of this table is
// nlMatrixDSPkts.1.783495.18.4.128.2.6.7.4.128.2.6.6
type RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry_HlMatrixControlIndex
    HlMatrixControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    NlMatrixDSTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The network destination address for this
    // nlMatrixDSEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    NlMatrixDSDestAddress interface{}

    // This attribute is a key. The network source address for this
    // nlMatrixDSEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    NlMatrixDSSourceAddress interface{}

    // The number of packets without errors transmitted from the source address to
    // the destination address since this entry was added to the nlMatrixDSTable. 
    // Note that this is the number of link-layer packets, so if a single
    // network-layer packet is fragmented into several link-layer frames, this
    // counter is incremented several times. The type is interface{} with range:
    // 0..4294967295.
    NlMatrixDSPkts interface{}

    // The number of octets transmitted from the source address to the destination
    // address since this entry was added to the nlMatrixDSTable (excluding
    // framing bits but including FCS octets), excluding those octets in packets
    // that contained errors.  Note this doesn't count just those octets in the
    // particular protocol frames, but includes the entire packet that contained
    // the protocol. The type is interface{} with range: 0..4294967295.
    NlMatrixDSOctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    NlMatrixDSCreateTime interface{}
}

func (nlMatrixDSEntry *RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry) GetEntityData() *types.CommonEntityData {
    nlMatrixDSEntry.EntityData.YFilter = nlMatrixDSEntry.YFilter
    nlMatrixDSEntry.EntityData.YangName = "nlMatrixDSEntry"
    nlMatrixDSEntry.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixDSEntry.EntityData.ParentYangName = "nlMatrixDSTable"
    nlMatrixDSEntry.EntityData.SegmentPath = "nlMatrixDSEntry" + types.AddKeyToken(nlMatrixDSEntry.HlMatrixControlIndex, "hlMatrixControlIndex") + types.AddKeyToken(nlMatrixDSEntry.NlMatrixDSTimeMark, "nlMatrixDSTimeMark") + types.AddKeyToken(nlMatrixDSEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(nlMatrixDSEntry.NlMatrixDSDestAddress, "nlMatrixDSDestAddress") + types.AddKeyToken(nlMatrixDSEntry.NlMatrixDSSourceAddress, "nlMatrixDSSourceAddress")
    nlMatrixDSEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/nlMatrixDSTable/" + nlMatrixDSEntry.EntityData.SegmentPath
    nlMatrixDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixDSEntry.EntityData.Children = types.NewOrderedMap()
    nlMatrixDSEntry.EntityData.Leafs = types.NewOrderedMap()
    nlMatrixDSEntry.EntityData.Leafs.Append("hlMatrixControlIndex", types.YLeaf{"HlMatrixControlIndex", nlMatrixDSEntry.HlMatrixControlIndex})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSTimeMark", types.YLeaf{"NlMatrixDSTimeMark", nlMatrixDSEntry.NlMatrixDSTimeMark})
    nlMatrixDSEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", nlMatrixDSEntry.ProtocolDirLocalIndex})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSDestAddress", types.YLeaf{"NlMatrixDSDestAddress", nlMatrixDSEntry.NlMatrixDSDestAddress})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSSourceAddress", types.YLeaf{"NlMatrixDSSourceAddress", nlMatrixDSEntry.NlMatrixDSSourceAddress})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSPkts", types.YLeaf{"NlMatrixDSPkts", nlMatrixDSEntry.NlMatrixDSPkts})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSOctets", types.YLeaf{"NlMatrixDSOctets", nlMatrixDSEntry.NlMatrixDSOctets})
    nlMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSCreateTime", types.YLeaf{"NlMatrixDSCreateTime", nlMatrixDSEntry.NlMatrixDSCreateTime})

    nlMatrixDSEntry.EntityData.YListKeys = []string {"HlMatrixControlIndex", "NlMatrixDSTimeMark", "ProtocolDirLocalIndex", "NlMatrixDSDestAddress", "NlMatrixDSSourceAddress"}

    return &(nlMatrixDSEntry.EntityData)
}

// RMON2MIB_NlMatrixTopNControlTable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_NlMatrixTopNControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNControlTable.  An example of the
    // indexing of this table is nlMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry.
    NlMatrixTopNControlEntry []*RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry
}

func (nlMatrixTopNControlTable *RMON2MIB_NlMatrixTopNControlTable) GetEntityData() *types.CommonEntityData {
    nlMatrixTopNControlTable.EntityData.YFilter = nlMatrixTopNControlTable.YFilter
    nlMatrixTopNControlTable.EntityData.YangName = "nlMatrixTopNControlTable"
    nlMatrixTopNControlTable.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixTopNControlTable.EntityData.ParentYangName = "RMON2-MIB"
    nlMatrixTopNControlTable.EntityData.SegmentPath = "nlMatrixTopNControlTable"
    nlMatrixTopNControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + nlMatrixTopNControlTable.EntityData.SegmentPath
    nlMatrixTopNControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixTopNControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixTopNControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixTopNControlTable.EntityData.Children = types.NewOrderedMap()
    nlMatrixTopNControlTable.EntityData.Children.Append("nlMatrixTopNControlEntry", types.YChild{"NlMatrixTopNControlEntry", nil})
    for i := range nlMatrixTopNControlTable.NlMatrixTopNControlEntry {
        nlMatrixTopNControlTable.EntityData.Children.Append(types.GetSegmentPath(nlMatrixTopNControlTable.NlMatrixTopNControlEntry[i]), types.YChild{"NlMatrixTopNControlEntry", nlMatrixTopNControlTable.NlMatrixTopNControlEntry[i]})
    }
    nlMatrixTopNControlTable.EntityData.Leafs = types.NewOrderedMap()

    nlMatrixTopNControlTable.EntityData.YListKeys = []string {}

    return &(nlMatrixTopNControlTable.EntityData)
}

// RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry
// A conceptual row in the nlMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// nlMatrixTopNControlDuration.3
type RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely identifies an entry in the
    // nlMatrixTopNControlTable.  Each such entry defines one top N report
    // prepared for one interface. The type is interface{} with range: 1..65535.
    NlMatrixTopNControlIndex interface{}

    // The nlMatrix[SD/DS] table for which a top N report will be prepared on
    // behalf of this entry.  The nlMatrix[SD/DS] table is identified by the value
    // of the hlMatrixControlIndex for that table - that value is used here to
    // identify the particular table.  This object may not be modified if the
    // associated nlMatrixTopNControlStatus object is equal to active(1). The type
    // is interface{} with range: 1..65535.
    NlMatrixTopNControlMatrixIndex interface{}

    // The variable for each nlMatrix[SD/DS] entry that the nlMatrixTopNEntries
    // are sorted by.      This object may not be modified if the associated
    // nlMatrixTopNControlStatus object is equal to active(1). The type is
    // NlMatrixTopNControlRateBase.
    NlMatrixTopNControlRateBase interface{}

    // The number of seconds left in the report currently being collected.  When
    // this object is modified by the management station, a new collection is
    // started, possibly aborting a currently running report.  The new value is
    // used as the requested duration of this report, and is immediately loaded
    // into the associated nlMatrixTopNControlDuration object. When the report
    // finishes, the probe will automatically start another collection with the
    // same initial value of nlMatrixTopNControlTimeRemaining.  Thus the
    // management station may simply read the resulting reports repeatedly,
    // checking the startTime and duration each time to ensure that a report was
    // not missed or that the report parameters were not changed.  While the value
    // of this object is non-zero, it decrements by one per second until it
    // reaches zero.  At the time that this object decrements to zero, the report
    // is made accessible in the nlMatrixTopNTable, overwriting any report that
    // may be there.  When this object is modified by the management station, any
    // associated entries in the nlMatrixTopNTable shall be deleted.  (Note that
    // this is a different algorithm than the one used in the hostTopNTable). The
    // type is interface{} with range: 0..2147483647.
    NlMatrixTopNControlTimeRemaining interface{}

    // The number of reports that have been generated by this entry. The type is
    // interface{} with range: 0..4294967295.
    NlMatrixTopNControlGeneratedReports interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval.  When the associated nlMatrixTopNControlTimeRemaining
    // object is set, this object shall be set by the probe to the same value and
    // shall not be modified until the next time the
    // nlMatrixTopNControlTimeRemaining is set. This value shall be zero if no
    // reports have been requested for this nlMatrixTopNControlEntry. The type is
    // interface{} with range: -2147483648..2147483647.
    NlMatrixTopNControlDuration interface{}

    // The maximum number of matrix entries requested for this report.  When this
    // object is created or modified, the probe should set
    // nlMatrixTopNControlGrantedSize as closely to this object as is possible for
    // the particular probe implementation and available resources. The type is
    // interface{} with range: 0..2147483647.
    NlMatrixTopNControlRequestedSize interface{}

    // The maximum number of matrix entries in this report.  When the associated
    // nlMatrixTopNControlRequestedSize object is created or modified, the probe
    // should set this object as closely to the requested value as is possible for
    // the particular implementation and available resources. The probe must not
    // lower this value except as a result of a set to the associated
    // nlMatrixTopNControlRequestedSize object.  If the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNPkts, when the next
    // topN report is generated, matrix entries with the highest value of
    // nlMatrixTopNPktRate shall be placed in this table in decreasing order of
    // this rate until there is no more room or until there are no more     
    // matrix entries.  If the value of nlMatrixTopNControlRateBase is equal to
    // nlMatrixTopNOctets, when the next topN report is generated, matrix entries
    // with the highest value of nlMatrixTopNOctetRate shall be placed in this
    // table in decreasing order of this rate until there is no more room or until
    // there are no more matrix entries.  It is an implementation-specific matter
    // how entries with the same value of nlMatrixTopNPktRate or
    // nlMatrixTopNOctetRate are sorted.  It is also an implementation-specific
    // matter as to whether or not zero-valued entries are available. The type is
    // interface{} with range: 0..2147483647.
    NlMatrixTopNControlGrantedSize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated
    // nlMatrixTopNControlTimeRemaining object was modified to start the requested
    // report or the time the report was last automatically (re)started.  This
    // object may be used by the management station to determine if a report was
    // missed or not. The type is interface{} with range: 0..4294967295.
    NlMatrixTopNControlStartTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    NlMatrixTopNControlOwner interface{}

    // The status of this nlMatrixTopNControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.    
    // If this object is not equal to active(1), all associated entries in the
    // nlMatrixTopNTable shall be deleted by the agent. The type is RowStatus.
    NlMatrixTopNControlStatus interface{}
}

func (nlMatrixTopNControlEntry *RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry) GetEntityData() *types.CommonEntityData {
    nlMatrixTopNControlEntry.EntityData.YFilter = nlMatrixTopNControlEntry.YFilter
    nlMatrixTopNControlEntry.EntityData.YangName = "nlMatrixTopNControlEntry"
    nlMatrixTopNControlEntry.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixTopNControlEntry.EntityData.ParentYangName = "nlMatrixTopNControlTable"
    nlMatrixTopNControlEntry.EntityData.SegmentPath = "nlMatrixTopNControlEntry" + types.AddKeyToken(nlMatrixTopNControlEntry.NlMatrixTopNControlIndex, "nlMatrixTopNControlIndex")
    nlMatrixTopNControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/nlMatrixTopNControlTable/" + nlMatrixTopNControlEntry.EntityData.SegmentPath
    nlMatrixTopNControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixTopNControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixTopNControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixTopNControlEntry.EntityData.Children = types.NewOrderedMap()
    nlMatrixTopNControlEntry.EntityData.Leafs = types.NewOrderedMap()
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlIndex", types.YLeaf{"NlMatrixTopNControlIndex", nlMatrixTopNControlEntry.NlMatrixTopNControlIndex})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlMatrixIndex", types.YLeaf{"NlMatrixTopNControlMatrixIndex", nlMatrixTopNControlEntry.NlMatrixTopNControlMatrixIndex})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlRateBase", types.YLeaf{"NlMatrixTopNControlRateBase", nlMatrixTopNControlEntry.NlMatrixTopNControlRateBase})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlTimeRemaining", types.YLeaf{"NlMatrixTopNControlTimeRemaining", nlMatrixTopNControlEntry.NlMatrixTopNControlTimeRemaining})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlGeneratedReports", types.YLeaf{"NlMatrixTopNControlGeneratedReports", nlMatrixTopNControlEntry.NlMatrixTopNControlGeneratedReports})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlDuration", types.YLeaf{"NlMatrixTopNControlDuration", nlMatrixTopNControlEntry.NlMatrixTopNControlDuration})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlRequestedSize", types.YLeaf{"NlMatrixTopNControlRequestedSize", nlMatrixTopNControlEntry.NlMatrixTopNControlRequestedSize})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlGrantedSize", types.YLeaf{"NlMatrixTopNControlGrantedSize", nlMatrixTopNControlEntry.NlMatrixTopNControlGrantedSize})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlStartTime", types.YLeaf{"NlMatrixTopNControlStartTime", nlMatrixTopNControlEntry.NlMatrixTopNControlStartTime})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlOwner", types.YLeaf{"NlMatrixTopNControlOwner", nlMatrixTopNControlEntry.NlMatrixTopNControlOwner})
    nlMatrixTopNControlEntry.EntityData.Leafs.Append("nlMatrixTopNControlStatus", types.YLeaf{"NlMatrixTopNControlStatus", nlMatrixTopNControlEntry.NlMatrixTopNControlStatus})

    nlMatrixTopNControlEntry.EntityData.YListKeys = []string {"NlMatrixTopNControlIndex"}

    return &(nlMatrixTopNControlEntry.EntityData)
}

// RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase represents nlMatrixTopNControlStatus object is equal to active(1).
type RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase string

const (
    RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase_nlMatrixTopNPkts RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase = "nlMatrixTopNPkts"

    RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase_nlMatrixTopNOctets RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlRateBase = "nlMatrixTopNOctets"
)

// RMON2MIB_NlMatrixTopNTable
// A set of statistics for those network layer matrix entries
// that have counted the highest number of octets or packets.
type RMON2MIB_NlMatrixTopNTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNTable.  The nlMatrixTopNControlIndex
    // value in the index identifies the nlMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // nlMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_NlMatrixTopNTable_NlMatrixTopNEntry.
    NlMatrixTopNEntry []*RMON2MIB_NlMatrixTopNTable_NlMatrixTopNEntry
}

func (nlMatrixTopNTable *RMON2MIB_NlMatrixTopNTable) GetEntityData() *types.CommonEntityData {
    nlMatrixTopNTable.EntityData.YFilter = nlMatrixTopNTable.YFilter
    nlMatrixTopNTable.EntityData.YangName = "nlMatrixTopNTable"
    nlMatrixTopNTable.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixTopNTable.EntityData.ParentYangName = "RMON2-MIB"
    nlMatrixTopNTable.EntityData.SegmentPath = "nlMatrixTopNTable"
    nlMatrixTopNTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + nlMatrixTopNTable.EntityData.SegmentPath
    nlMatrixTopNTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixTopNTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixTopNTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixTopNTable.EntityData.Children = types.NewOrderedMap()
    nlMatrixTopNTable.EntityData.Children.Append("nlMatrixTopNEntry", types.YChild{"NlMatrixTopNEntry", nil})
    for i := range nlMatrixTopNTable.NlMatrixTopNEntry {
        nlMatrixTopNTable.EntityData.Children.Append(types.GetSegmentPath(nlMatrixTopNTable.NlMatrixTopNEntry[i]), types.YChild{"NlMatrixTopNEntry", nlMatrixTopNTable.NlMatrixTopNEntry[i]})
    }
    nlMatrixTopNTable.EntityData.Leafs = types.NewOrderedMap()

    nlMatrixTopNTable.EntityData.YListKeys = []string {}

    return &(nlMatrixTopNTable.EntityData)
}

// RMON2MIB_NlMatrixTopNTable_NlMatrixTopNEntry
// A conceptual row in the nlMatrixTopNTable.
// 
// The nlMatrixTopNControlIndex value in the index identifies the
// nlMatrixTopNControlEntry on whose behalf this entry was
// created.
// 
// An example of the indexing of this table is
// nlMatrixTopNPktRate.3.10
type RMON2MIB_NlMatrixTopNTable_NlMatrixTopNEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_NlMatrixTopNControlTable_NlMatrixTopNControlEntry_NlMatrixTopNControlIndex
    NlMatrixTopNControlIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // nlMatrixTopNTable among those in the same report.      This index is
    // between 1 and N, where N is the number of entries in this report.  If the
    // value of nlMatrixTopNControlRateBase is equal to nlMatrixTopNPkts,
    // increasing values of nlMatrixTopNIndex shall be assigned to entries with
    // decreasing values of nlMatrixTopNPktRate until index N is assigned or there
    // are no more nlMatrixTopNEntries.  If the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNOctets, increasing
    // values of nlMatrixTopNIndex shall be assigned to entries with decreasing
    // values of nlMatrixTopNOctetRate until index N is assigned or there are no
    // more nlMatrixTopNEntries. The type is interface{} with range: 1..65535.
    NlMatrixTopNIndex interface{}

    // The protocolDirLocalIndex of the network layer protocol of this entry's
    // network address. The type is interface{} with range: 1..2147483647.
    NlMatrixTopNProtocolDirLocalIndex interface{}

    // The network layer address of the source host in this conversation.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the associated nlMatrixTopNProtocolDirLocalIndex.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    NlMatrixTopNSourceAddress interface{}

    // The network layer address of the destination host in this conversation. 
    // This is represented as an octet string with specific semantics and length
    // as identified by the associated nlMatrixTopNProtocolDirLocalIndex.  For
    // example, if the nlMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    NlMatrixTopNDestAddress interface{}

    // The number of packets seen from the source host to the destination host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDPkts object. If the value of nlMatrixTopNControlRateBase is
    // nlMatrixTopNPkts, this variable will be used to sort this report. The type
    // is interface{} with range: 0..4294967295.
    NlMatrixTopNPktRate interface{}

    // The number of packets seen from the destination host to the source host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDPkts object (note that the corresponding nlMatrixSDPkts object
    // selected is the one whose source address is equal to
    // nlMatrixTopNDestAddress and whose destination address is equal to
    // nlMatrixTopNSourceAddress.)  Note that if the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNPkts, the sort of topN
    // entries is based entirely on nlMatrixTopNPktRate, and not on the value of
    // this object. The type is interface{} with range: 0..4294967295.
    NlMatrixTopNReversePktRate interface{}

    // The number of octets seen from the source host to the destination host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDOctets object.  If the value of nlMatrixTopNControlRateBase is
    // nlMatrixTopNOctets, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    NlMatrixTopNOctetRate interface{}

    // The number of octets seen from the destination host to the source host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixDSOctets object (note that the corresponding nlMatrixSDOctets
    // object selected is the one whose source address is equal to
    // nlMatrixTopNDestAddress and whose destination address is equal to
    // nlMatrixTopNSourceAddress.)  Note that if the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNOctets, the sort of
    // topN entries is based entirely on nlMatrixTopNOctetRate, and not on the
    // value of this object. The type is interface{} with range: 0..4294967295.
    NlMatrixTopNReverseOctetRate interface{}
}

func (nlMatrixTopNEntry *RMON2MIB_NlMatrixTopNTable_NlMatrixTopNEntry) GetEntityData() *types.CommonEntityData {
    nlMatrixTopNEntry.EntityData.YFilter = nlMatrixTopNEntry.YFilter
    nlMatrixTopNEntry.EntityData.YangName = "nlMatrixTopNEntry"
    nlMatrixTopNEntry.EntityData.BundleName = "cisco_ios_xe"
    nlMatrixTopNEntry.EntityData.ParentYangName = "nlMatrixTopNTable"
    nlMatrixTopNEntry.EntityData.SegmentPath = "nlMatrixTopNEntry" + types.AddKeyToken(nlMatrixTopNEntry.NlMatrixTopNControlIndex, "nlMatrixTopNControlIndex") + types.AddKeyToken(nlMatrixTopNEntry.NlMatrixTopNIndex, "nlMatrixTopNIndex")
    nlMatrixTopNEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/nlMatrixTopNTable/" + nlMatrixTopNEntry.EntityData.SegmentPath
    nlMatrixTopNEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlMatrixTopNEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlMatrixTopNEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlMatrixTopNEntry.EntityData.Children = types.NewOrderedMap()
    nlMatrixTopNEntry.EntityData.Leafs = types.NewOrderedMap()
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNControlIndex", types.YLeaf{"NlMatrixTopNControlIndex", nlMatrixTopNEntry.NlMatrixTopNControlIndex})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNIndex", types.YLeaf{"NlMatrixTopNIndex", nlMatrixTopNEntry.NlMatrixTopNIndex})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNProtocolDirLocalIndex", types.YLeaf{"NlMatrixTopNProtocolDirLocalIndex", nlMatrixTopNEntry.NlMatrixTopNProtocolDirLocalIndex})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNSourceAddress", types.YLeaf{"NlMatrixTopNSourceAddress", nlMatrixTopNEntry.NlMatrixTopNSourceAddress})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNDestAddress", types.YLeaf{"NlMatrixTopNDestAddress", nlMatrixTopNEntry.NlMatrixTopNDestAddress})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNPktRate", types.YLeaf{"NlMatrixTopNPktRate", nlMatrixTopNEntry.NlMatrixTopNPktRate})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNReversePktRate", types.YLeaf{"NlMatrixTopNReversePktRate", nlMatrixTopNEntry.NlMatrixTopNReversePktRate})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNOctetRate", types.YLeaf{"NlMatrixTopNOctetRate", nlMatrixTopNEntry.NlMatrixTopNOctetRate})
    nlMatrixTopNEntry.EntityData.Leafs.Append("nlMatrixTopNReverseOctetRate", types.YLeaf{"NlMatrixTopNReverseOctetRate", nlMatrixTopNEntry.NlMatrixTopNReverseOctetRate})

    nlMatrixTopNEntry.EntityData.YListKeys = []string {"NlMatrixTopNControlIndex", "NlMatrixTopNIndex"}

    return &(nlMatrixTopNEntry.EntityData)
}

// RMON2MIB_AlHostTable
// A collection of statistics for a particular protocol from a
// particular network address that has been discovered on an
// interface of this device.
// 
// The probe will populate this table for all protocols in the
// protocol directory table whose value of
// protocolDirHostConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirHostConfig value of supportedOff(2).
// 
// The probe will add to this table all addresses
// seen as the source or destination address in all packets with
// no MAC errors, and will increment octet and packet counts in
// the table for all packets with no MAC errors.  Further,
// entries will only be added to this table if their address
// exists in the nlHostTable and will be deleted from this table
// if their address is deleted from the nlHostTable.
type RMON2MIB_AlHostTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alHostTable.  The hlHostControlIndex value in the
    // index identifies the hlHostControlEntry on whose behalf this entry was
    // created. The first protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the address. The nlHostAddress value in the index
    // identifies the network layer address of this entry. The second
    // protocolDirLocalIndex value in the index identifies the protocol that is
    // counted by this entry.  An example of the indexing in this entry is
    // alHostOutPkts.1.783495.18.4.128.2.6.6.34. The type is slice of
    // RMON2MIB_AlHostTable_AlHostEntry.
    AlHostEntry []*RMON2MIB_AlHostTable_AlHostEntry
}

func (alHostTable *RMON2MIB_AlHostTable) GetEntityData() *types.CommonEntityData {
    alHostTable.EntityData.YFilter = alHostTable.YFilter
    alHostTable.EntityData.YangName = "alHostTable"
    alHostTable.EntityData.BundleName = "cisco_ios_xe"
    alHostTable.EntityData.ParentYangName = "RMON2-MIB"
    alHostTable.EntityData.SegmentPath = "alHostTable"
    alHostTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + alHostTable.EntityData.SegmentPath
    alHostTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alHostTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alHostTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alHostTable.EntityData.Children = types.NewOrderedMap()
    alHostTable.EntityData.Children.Append("alHostEntry", types.YChild{"AlHostEntry", nil})
    for i := range alHostTable.AlHostEntry {
        alHostTable.EntityData.Children.Append(types.GetSegmentPath(alHostTable.AlHostEntry[i]), types.YChild{"AlHostEntry", alHostTable.AlHostEntry[i]})
    }
    alHostTable.EntityData.Leafs = types.NewOrderedMap()

    alHostTable.EntityData.YListKeys = []string {}

    return &(alHostTable.EntityData)
}

// RMON2MIB_AlHostTable_AlHostEntry
// A conceptual row in the alHostTable.
// 
// The hlHostControlIndex value in the index identifies the
// hlHostControlEntry on whose behalf this entry was created.
// The first protocolDirLocalIndex value in the index identifies
// the network layer protocol of the address.
// The nlHostAddress value in the index identifies the network
// layer address of this entry.
// The second protocolDirLocalIndex value in the index identifies
// the protocol that is counted by this entry.
// 
// An example of the indexing in this entry is
// alHostOutPkts.1.783495.18.4.128.2.6.6.34
type RMON2MIB_AlHostTable_AlHostEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlHostControlTable_HlHostControlEntry_HlHostControlIndex
    HlHostControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    AlHostTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_NlHostTable_NlHostEntry_NlHostAddress
    NlHostAddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex2 interface{}

    // The number of packets of this protocol type without errors transmitted to
    // this address since it was added to the alHostTable.  Note that this is the
    // number of link-layer packets, so if a single network-layer packet is
    // fragmented into several link-layer frames, this counter is incremented
    // several times. The type is interface{} with range: 0..4294967295.
    AlHostInPkts interface{}

    // The number of packets of this protocol type without errors transmitted by
    // this address since it was added to the alHostTable.  Note that this is the
    // number of link-layer packets, so if a single network-layer packet is
    // fragmented into several link-layer frames, this counter is incremented
    // several times. The type is interface{} with range: 0..4294967295.
    AlHostOutPkts interface{}

    // The number of octets transmitted to this address of this protocol type
    // since it was added to the alHostTable (excluding framing bits but including
    // FCS octets), excluding those octets in packets that contained errors.  Note
    // this doesn't count just those octets in the particular protocol frames, but
    // includes the entire packet that contained the protocol. The type is
    // interface{} with range: 0..4294967295.
    AlHostInOctets interface{}

    // The number of octets transmitted by this address of this protocol type
    // since it was added to the alHostTable (excluding framing bits but including
    // FCS octets), excluding those octets in packets that contained errors.  Note
    // this doesn't count just those octets in the particular protocol frames, but
    // includes the entire packet that contained the protocol. The type is
    // interface{} with range: 0..4294967295.
    AlHostOutOctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    AlHostCreateTime interface{}
}

func (alHostEntry *RMON2MIB_AlHostTable_AlHostEntry) GetEntityData() *types.CommonEntityData {
    alHostEntry.EntityData.YFilter = alHostEntry.YFilter
    alHostEntry.EntityData.YangName = "alHostEntry"
    alHostEntry.EntityData.BundleName = "cisco_ios_xe"
    alHostEntry.EntityData.ParentYangName = "alHostTable"
    alHostEntry.EntityData.SegmentPath = "alHostEntry" + types.AddKeyToken(alHostEntry.HlHostControlIndex, "hlHostControlIndex") + types.AddKeyToken(alHostEntry.AlHostTimeMark, "alHostTimeMark") + types.AddKeyToken(alHostEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(alHostEntry.NlHostAddress, "nlHostAddress") + types.AddKeyToken(alHostEntry.ProtocolDirLocalIndex2, "protocolDirLocalIndex_2")
    alHostEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/alHostTable/" + alHostEntry.EntityData.SegmentPath
    alHostEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alHostEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alHostEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alHostEntry.EntityData.Children = types.NewOrderedMap()
    alHostEntry.EntityData.Leafs = types.NewOrderedMap()
    alHostEntry.EntityData.Leafs.Append("hlHostControlIndex", types.YLeaf{"HlHostControlIndex", alHostEntry.HlHostControlIndex})
    alHostEntry.EntityData.Leafs.Append("alHostTimeMark", types.YLeaf{"AlHostTimeMark", alHostEntry.AlHostTimeMark})
    alHostEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", alHostEntry.ProtocolDirLocalIndex})
    alHostEntry.EntityData.Leafs.Append("nlHostAddress", types.YLeaf{"NlHostAddress", alHostEntry.NlHostAddress})
    alHostEntry.EntityData.Leafs.Append("protocolDirLocalIndex_2", types.YLeaf{"ProtocolDirLocalIndex2", alHostEntry.ProtocolDirLocalIndex2})
    alHostEntry.EntityData.Leafs.Append("alHostInPkts", types.YLeaf{"AlHostInPkts", alHostEntry.AlHostInPkts})
    alHostEntry.EntityData.Leafs.Append("alHostOutPkts", types.YLeaf{"AlHostOutPkts", alHostEntry.AlHostOutPkts})
    alHostEntry.EntityData.Leafs.Append("alHostInOctets", types.YLeaf{"AlHostInOctets", alHostEntry.AlHostInOctets})
    alHostEntry.EntityData.Leafs.Append("alHostOutOctets", types.YLeaf{"AlHostOutOctets", alHostEntry.AlHostOutOctets})
    alHostEntry.EntityData.Leafs.Append("alHostCreateTime", types.YLeaf{"AlHostCreateTime", alHostEntry.AlHostCreateTime})

    alHostEntry.EntityData.YListKeys = []string {"HlHostControlIndex", "AlHostTimeMark", "ProtocolDirLocalIndex", "NlHostAddress", "ProtocolDirLocalIndex2"}

    return &(alHostEntry.EntityData)
}

// RMON2MIB_AlMatrixSDTable
// A list of application traffic matrix entries which collect
// 
// 
// 
// 
// 
// statistics for conversations of a particular protocol between
// two network-level addresses.  This table is indexed first by
// the source address and then by the destination address to make
// it convenient to collect all statistics from a particular
// address.
// 
// The probe will populate this table for all protocols in the
// protocol directory table whose value of
// protocolDirMatrixConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirMatrixConfig value of supportedOff(2).
// 
// The probe will add to this table all pairs of addresses for
// all protocols seen in all packets with no MAC errors, and will
// increment octet and packet counts in the table for all packets
// with no MAC errors.  Further, entries will only be added to
// this table if their address pair exists in the nlMatrixSDTable
// and will be deleted from this table if the address pair is
// deleted from the nlMatrixSDTable.
type RMON2MIB_AlMatrixSDTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixSDTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The first protocolDirLocalIndex value in the index identifies
    // the network layer protocol of the nlMatrixSDSourceAddress and
    // nlMatrixSDDestAddress. The nlMatrixSDSourceAddress value in the index
    // identifies the network layer address of the source host in this
    // conversation. The nlMatrixSDDestAddress value in the index identifies the
    // network layer address of the destination host in this conversation. The
    // second protocolDirLocalIndex value in the index identifies the protocol
    // that is counted by this entry.  An example of the indexing of this entry is
    // alMatrixSDPkts.1.783495.18.4.128.2.6.6.4.128.2.6.7.34. The type is slice of
    // RMON2MIB_AlMatrixSDTable_AlMatrixSDEntry.
    AlMatrixSDEntry []*RMON2MIB_AlMatrixSDTable_AlMatrixSDEntry
}

func (alMatrixSDTable *RMON2MIB_AlMatrixSDTable) GetEntityData() *types.CommonEntityData {
    alMatrixSDTable.EntityData.YFilter = alMatrixSDTable.YFilter
    alMatrixSDTable.EntityData.YangName = "alMatrixSDTable"
    alMatrixSDTable.EntityData.BundleName = "cisco_ios_xe"
    alMatrixSDTable.EntityData.ParentYangName = "RMON2-MIB"
    alMatrixSDTable.EntityData.SegmentPath = "alMatrixSDTable"
    alMatrixSDTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + alMatrixSDTable.EntityData.SegmentPath
    alMatrixSDTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixSDTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixSDTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixSDTable.EntityData.Children = types.NewOrderedMap()
    alMatrixSDTable.EntityData.Children.Append("alMatrixSDEntry", types.YChild{"AlMatrixSDEntry", nil})
    for i := range alMatrixSDTable.AlMatrixSDEntry {
        alMatrixSDTable.EntityData.Children.Append(types.GetSegmentPath(alMatrixSDTable.AlMatrixSDEntry[i]), types.YChild{"AlMatrixSDEntry", alMatrixSDTable.AlMatrixSDEntry[i]})
    }
    alMatrixSDTable.EntityData.Leafs = types.NewOrderedMap()

    alMatrixSDTable.EntityData.YListKeys = []string {}

    return &(alMatrixSDTable.EntityData)
}

// RMON2MIB_AlMatrixSDTable_AlMatrixSDEntry
// A conceptual row in the alMatrixSDTable.
// 
// The hlMatrixControlIndex value in the index identifies the
// hlMatrixControlEntry on whose behalf this entry was created.
// The first protocolDirLocalIndex value in the index identifies
// the network layer protocol of the nlMatrixSDSourceAddress and
// nlMatrixSDDestAddress.
// The nlMatrixSDSourceAddress value in the index identifies the
// network layer address of the source host in this conversation.
// The nlMatrixSDDestAddress value in the index identifies the
// network layer address of the destination host in this
// conversation.
// The second protocolDirLocalIndex value in the index identifies
// the protocol that is counted by this entry.
// 
// An example of the indexing of this entry is
// alMatrixSDPkts.1.783495.18.4.128.2.6.6.4.128.2.6.7.34
type RMON2MIB_AlMatrixSDTable_AlMatrixSDEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry_HlMatrixControlIndex
    HlMatrixControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    AlMatrixSDTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry_NlMatrixSDSourceAddress
    NlMatrixSDSourceAddress interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_NlMatrixSDTable_NlMatrixSDEntry_NlMatrixSDDestAddress
    NlMatrixSDDestAddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex2 interface{}

    // The number of packets of this protocol type without errors transmitted from
    // the source address to the destination address since this entry was added to
    // the alMatrixSDTable.  Note that this is the number of link-layer packets,
    // so if a single network-layer packet is fragmented into several link-layer
    // frames, this counter is incremented several times. The type is interface{}
    // with range: 0..4294967295.
    AlMatrixSDPkts interface{}

    // The number of octets in packets of this protocol type transmitted from the
    // source address to the destination address since this entry was added to the
    // alMatrixSDTable (excluding framing bits but including FCS octets),
    // excluding those octets in packets that contained errors.  Note this doesn't
    // count just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    AlMatrixSDOctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    AlMatrixSDCreateTime interface{}
}

func (alMatrixSDEntry *RMON2MIB_AlMatrixSDTable_AlMatrixSDEntry) GetEntityData() *types.CommonEntityData {
    alMatrixSDEntry.EntityData.YFilter = alMatrixSDEntry.YFilter
    alMatrixSDEntry.EntityData.YangName = "alMatrixSDEntry"
    alMatrixSDEntry.EntityData.BundleName = "cisco_ios_xe"
    alMatrixSDEntry.EntityData.ParentYangName = "alMatrixSDTable"
    alMatrixSDEntry.EntityData.SegmentPath = "alMatrixSDEntry" + types.AddKeyToken(alMatrixSDEntry.HlMatrixControlIndex, "hlMatrixControlIndex") + types.AddKeyToken(alMatrixSDEntry.AlMatrixSDTimeMark, "alMatrixSDTimeMark") + types.AddKeyToken(alMatrixSDEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(alMatrixSDEntry.NlMatrixSDSourceAddress, "nlMatrixSDSourceAddress") + types.AddKeyToken(alMatrixSDEntry.NlMatrixSDDestAddress, "nlMatrixSDDestAddress") + types.AddKeyToken(alMatrixSDEntry.ProtocolDirLocalIndex2, "protocolDirLocalIndex_2")
    alMatrixSDEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/alMatrixSDTable/" + alMatrixSDEntry.EntityData.SegmentPath
    alMatrixSDEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixSDEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixSDEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixSDEntry.EntityData.Children = types.NewOrderedMap()
    alMatrixSDEntry.EntityData.Leafs = types.NewOrderedMap()
    alMatrixSDEntry.EntityData.Leafs.Append("hlMatrixControlIndex", types.YLeaf{"HlMatrixControlIndex", alMatrixSDEntry.HlMatrixControlIndex})
    alMatrixSDEntry.EntityData.Leafs.Append("alMatrixSDTimeMark", types.YLeaf{"AlMatrixSDTimeMark", alMatrixSDEntry.AlMatrixSDTimeMark})
    alMatrixSDEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", alMatrixSDEntry.ProtocolDirLocalIndex})
    alMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDSourceAddress", types.YLeaf{"NlMatrixSDSourceAddress", alMatrixSDEntry.NlMatrixSDSourceAddress})
    alMatrixSDEntry.EntityData.Leafs.Append("nlMatrixSDDestAddress", types.YLeaf{"NlMatrixSDDestAddress", alMatrixSDEntry.NlMatrixSDDestAddress})
    alMatrixSDEntry.EntityData.Leafs.Append("protocolDirLocalIndex_2", types.YLeaf{"ProtocolDirLocalIndex2", alMatrixSDEntry.ProtocolDirLocalIndex2})
    alMatrixSDEntry.EntityData.Leafs.Append("alMatrixSDPkts", types.YLeaf{"AlMatrixSDPkts", alMatrixSDEntry.AlMatrixSDPkts})
    alMatrixSDEntry.EntityData.Leafs.Append("alMatrixSDOctets", types.YLeaf{"AlMatrixSDOctets", alMatrixSDEntry.AlMatrixSDOctets})
    alMatrixSDEntry.EntityData.Leafs.Append("alMatrixSDCreateTime", types.YLeaf{"AlMatrixSDCreateTime", alMatrixSDEntry.AlMatrixSDCreateTime})

    alMatrixSDEntry.EntityData.YListKeys = []string {"HlMatrixControlIndex", "AlMatrixSDTimeMark", "ProtocolDirLocalIndex", "NlMatrixSDSourceAddress", "NlMatrixSDDestAddress", "ProtocolDirLocalIndex2"}

    return &(alMatrixSDEntry.EntityData)
}

// RMON2MIB_AlMatrixDSTable
// A list of application traffic matrix entries which collect
// statistics for conversations of a particular protocol between
// two network-level addresses.  This table is indexed first by
// the destination address and then by the source address to make
// it convenient to collect all statistics to a particular
// address.
// 
// The probe will populate this table for all protocols in the
// protocol directory table whose value of
// protocolDirMatrixConfig is equal to supportedOn(3), and
// will delete any entries whose protocolDirEntry is deleted or
// has a protocolDirMatrixConfig value of supportedOff(2).
// 
// The probe will add to this table all pairs of addresses for
// all protocols seen in all packets with no MAC errors, and will
// increment octet and packet counts in the table for all packets
// with no MAC errors.  Further, entries will only be added to
// this table if their address pair exists in the nlMatrixDSTable
// and will be deleted from this table if the address pair is
// deleted from the nlMatrixDSTable.
type RMON2MIB_AlMatrixDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixDSTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The first protocolDirLocalIndex value in the index identifies
    // the network layer protocol of the alMatrixDSSourceAddress and
    // alMatrixDSDestAddress.      The nlMatrixDSDestAddress value in the index
    // identifies the network layer address of the destination host in this
    // conversation. The nlMatrixDSSourceAddress value in the index identifies the
    // network layer address of the source host in this conversation. The second
    // protocolDirLocalIndex value in the index identifies the protocol that is
    // counted by this entry.  An example of the indexing of this entry is
    // alMatrixDSPkts.1.783495.18.4.128.2.6.7.4.128.2.6.6.34. The type is slice of
    // RMON2MIB_AlMatrixDSTable_AlMatrixDSEntry.
    AlMatrixDSEntry []*RMON2MIB_AlMatrixDSTable_AlMatrixDSEntry
}

func (alMatrixDSTable *RMON2MIB_AlMatrixDSTable) GetEntityData() *types.CommonEntityData {
    alMatrixDSTable.EntityData.YFilter = alMatrixDSTable.YFilter
    alMatrixDSTable.EntityData.YangName = "alMatrixDSTable"
    alMatrixDSTable.EntityData.BundleName = "cisco_ios_xe"
    alMatrixDSTable.EntityData.ParentYangName = "RMON2-MIB"
    alMatrixDSTable.EntityData.SegmentPath = "alMatrixDSTable"
    alMatrixDSTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + alMatrixDSTable.EntityData.SegmentPath
    alMatrixDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixDSTable.EntityData.Children = types.NewOrderedMap()
    alMatrixDSTable.EntityData.Children.Append("alMatrixDSEntry", types.YChild{"AlMatrixDSEntry", nil})
    for i := range alMatrixDSTable.AlMatrixDSEntry {
        alMatrixDSTable.EntityData.Children.Append(types.GetSegmentPath(alMatrixDSTable.AlMatrixDSEntry[i]), types.YChild{"AlMatrixDSEntry", alMatrixDSTable.AlMatrixDSEntry[i]})
    }
    alMatrixDSTable.EntityData.Leafs = types.NewOrderedMap()

    alMatrixDSTable.EntityData.YListKeys = []string {}

    return &(alMatrixDSTable.EntityData)
}

// RMON2MIB_AlMatrixDSTable_AlMatrixDSEntry
// A conceptual row in the alMatrixDSTable.
// 
// The hlMatrixControlIndex value in the index identifies the
// hlMatrixControlEntry on whose behalf this entry was created.
// The first protocolDirLocalIndex value in the index identifies
// the network layer protocol of the alMatrixDSSourceAddress and
// alMatrixDSDestAddress.
// 
// 
// 
// 
// 
// The nlMatrixDSDestAddress value in the index identifies the
// network layer address of the destination host in this
// conversation.
// The nlMatrixDSSourceAddress value in the index identifies the
// network layer address of the source host in this conversation.
// The second protocolDirLocalIndex value in the index identifies
// the protocol that is counted by this entry.
// 
// An example of the indexing of this entry is
// alMatrixDSPkts.1.783495.18.4.128.2.6.7.4.128.2.6.6.34
type RMON2MIB_AlMatrixDSTable_AlMatrixDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_HlMatrixControlTable_HlMatrixControlEntry_HlMatrixControlIndex
    HlMatrixControlIndex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    AlMatrixDSTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry_NlMatrixDSDestAddress
    NlMatrixDSDestAddress interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_NlMatrixDSTable_NlMatrixDSEntry_NlMatrixDSSourceAddress
    NlMatrixDSSourceAddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_ProtocolDirTable_ProtocolDirEntry_ProtocolDirLocalIndex
    ProtocolDirLocalIndex2 interface{}

    // The number of packets of this protocol type without errors transmitted from
    // the source address to the destination address since this entry was added to
    // the alMatrixDSTable.  Note that this is the number of link-layer packets,
    // so if a single network-layer packet is fragmented into several link-layer
    // frames, this counter is incremented several times. The type is interface{}
    // with range: 0..4294967295.
    AlMatrixDSPkts interface{}

    // The number of octets in packets of this protocol type transmitted from the
    // source address to the destination address since this entry was added to the
    // alMatrixDSTable (excluding framing bits but including FCS octets),
    // excluding those octets in packets that contained errors.  Note this doesn't
    // count just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    AlMatrixDSOctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    AlMatrixDSCreateTime interface{}
}

func (alMatrixDSEntry *RMON2MIB_AlMatrixDSTable_AlMatrixDSEntry) GetEntityData() *types.CommonEntityData {
    alMatrixDSEntry.EntityData.YFilter = alMatrixDSEntry.YFilter
    alMatrixDSEntry.EntityData.YangName = "alMatrixDSEntry"
    alMatrixDSEntry.EntityData.BundleName = "cisco_ios_xe"
    alMatrixDSEntry.EntityData.ParentYangName = "alMatrixDSTable"
    alMatrixDSEntry.EntityData.SegmentPath = "alMatrixDSEntry" + types.AddKeyToken(alMatrixDSEntry.HlMatrixControlIndex, "hlMatrixControlIndex") + types.AddKeyToken(alMatrixDSEntry.AlMatrixDSTimeMark, "alMatrixDSTimeMark") + types.AddKeyToken(alMatrixDSEntry.ProtocolDirLocalIndex, "protocolDirLocalIndex") + types.AddKeyToken(alMatrixDSEntry.NlMatrixDSDestAddress, "nlMatrixDSDestAddress") + types.AddKeyToken(alMatrixDSEntry.NlMatrixDSSourceAddress, "nlMatrixDSSourceAddress") + types.AddKeyToken(alMatrixDSEntry.ProtocolDirLocalIndex2, "protocolDirLocalIndex_2")
    alMatrixDSEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/alMatrixDSTable/" + alMatrixDSEntry.EntityData.SegmentPath
    alMatrixDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixDSEntry.EntityData.Children = types.NewOrderedMap()
    alMatrixDSEntry.EntityData.Leafs = types.NewOrderedMap()
    alMatrixDSEntry.EntityData.Leafs.Append("hlMatrixControlIndex", types.YLeaf{"HlMatrixControlIndex", alMatrixDSEntry.HlMatrixControlIndex})
    alMatrixDSEntry.EntityData.Leafs.Append("alMatrixDSTimeMark", types.YLeaf{"AlMatrixDSTimeMark", alMatrixDSEntry.AlMatrixDSTimeMark})
    alMatrixDSEntry.EntityData.Leafs.Append("protocolDirLocalIndex", types.YLeaf{"ProtocolDirLocalIndex", alMatrixDSEntry.ProtocolDirLocalIndex})
    alMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSDestAddress", types.YLeaf{"NlMatrixDSDestAddress", alMatrixDSEntry.NlMatrixDSDestAddress})
    alMatrixDSEntry.EntityData.Leafs.Append("nlMatrixDSSourceAddress", types.YLeaf{"NlMatrixDSSourceAddress", alMatrixDSEntry.NlMatrixDSSourceAddress})
    alMatrixDSEntry.EntityData.Leafs.Append("protocolDirLocalIndex_2", types.YLeaf{"ProtocolDirLocalIndex2", alMatrixDSEntry.ProtocolDirLocalIndex2})
    alMatrixDSEntry.EntityData.Leafs.Append("alMatrixDSPkts", types.YLeaf{"AlMatrixDSPkts", alMatrixDSEntry.AlMatrixDSPkts})
    alMatrixDSEntry.EntityData.Leafs.Append("alMatrixDSOctets", types.YLeaf{"AlMatrixDSOctets", alMatrixDSEntry.AlMatrixDSOctets})
    alMatrixDSEntry.EntityData.Leafs.Append("alMatrixDSCreateTime", types.YLeaf{"AlMatrixDSCreateTime", alMatrixDSEntry.AlMatrixDSCreateTime})

    alMatrixDSEntry.EntityData.YListKeys = []string {"HlMatrixControlIndex", "AlMatrixDSTimeMark", "ProtocolDirLocalIndex", "NlMatrixDSDestAddress", "NlMatrixDSSourceAddress", "ProtocolDirLocalIndex2"}

    return &(alMatrixDSEntry.EntityData)
}

// RMON2MIB_AlMatrixTopNControlTable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_AlMatrixTopNControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNControlTable.  An example of the
    // indexing of this table is alMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry.
    AlMatrixTopNControlEntry []*RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry
}

func (alMatrixTopNControlTable *RMON2MIB_AlMatrixTopNControlTable) GetEntityData() *types.CommonEntityData {
    alMatrixTopNControlTable.EntityData.YFilter = alMatrixTopNControlTable.YFilter
    alMatrixTopNControlTable.EntityData.YangName = "alMatrixTopNControlTable"
    alMatrixTopNControlTable.EntityData.BundleName = "cisco_ios_xe"
    alMatrixTopNControlTable.EntityData.ParentYangName = "RMON2-MIB"
    alMatrixTopNControlTable.EntityData.SegmentPath = "alMatrixTopNControlTable"
    alMatrixTopNControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + alMatrixTopNControlTable.EntityData.SegmentPath
    alMatrixTopNControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixTopNControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixTopNControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixTopNControlTable.EntityData.Children = types.NewOrderedMap()
    alMatrixTopNControlTable.EntityData.Children.Append("alMatrixTopNControlEntry", types.YChild{"AlMatrixTopNControlEntry", nil})
    for i := range alMatrixTopNControlTable.AlMatrixTopNControlEntry {
        alMatrixTopNControlTable.EntityData.Children.Append(types.GetSegmentPath(alMatrixTopNControlTable.AlMatrixTopNControlEntry[i]), types.YChild{"AlMatrixTopNControlEntry", alMatrixTopNControlTable.AlMatrixTopNControlEntry[i]})
    }
    alMatrixTopNControlTable.EntityData.Leafs = types.NewOrderedMap()

    alMatrixTopNControlTable.EntityData.YListKeys = []string {}

    return &(alMatrixTopNControlTable.EntityData)
}

// RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry
// A conceptual row in the alMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// alMatrixTopNControlDuration.3
type RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely identifies an entry in the
    // alMatrixTopNControlTable.  Each such entry defines one top N report
    // prepared for one interface. The type is interface{} with range: 1..65535.
    AlMatrixTopNControlIndex interface{}

    // The alMatrix[SD/DS] table for which a top N report will be prepared on
    // behalf of this entry.  The alMatrix[SD/DS] table is identified by the value
    // of the hlMatrixControlIndex for that table - that value is used here to
    // identify the particular table.  This object may not be modified if the
    // associated alMatrixTopNControlStatus object is equal to active(1). The type
    // is interface{} with range: 1..65535.
    AlMatrixTopNControlMatrixIndex interface{}

    // The variable for each alMatrix[SD/DS] entry that the     
    // alMatrixTopNEntries are sorted by, as well as the selector of the view of
    // the matrix table that will be used.  The values alMatrixTopNTerminalsPkts
    // and alMatrixTopNTerminalsOctets cause collection only from protocols that
    // have no child protocols that are counted.  The values alMatrixTopNAllPkts
    // and alMatrixTopNAllOctets cause collection from all alMatrix entries.  This
    // object may not be modified if the associated alMatrixTopNControlStatus
    // object is equal to active(1). The type is AlMatrixTopNControlRateBase.
    AlMatrixTopNControlRateBase interface{}

    // The number of seconds left in the report currently being collected.  When
    // this object is modified by the management station, a new collection is
    // started, possibly aborting a currently running report.  The new value is
    // used as the requested duration of this report, and is immediately loaded
    // into the associated alMatrixTopNControlDuration object. When the report
    // finishes, the probe will automatically start another collection with the
    // same initial value of alMatrixTopNControlTimeRemaining.  Thus the
    // management station may simply read the resulting reports repeatedly,
    // checking the startTime and duration each time to ensure that a report was
    // not missed or that the report parameters were not changed.  While the value
    // of this object is non-zero, it decrements by one per second until it
    // reaches zero.  At the time that this object decrements to zero, the report
    // is made accessible in the alMatrixTopNTable, overwriting any report that
    // may be there.  When this object is modified by the management station, any
    // associated entries in the alMatrixTopNTable shall be deleted.  (Note that
    // this is a different algorithm than the one used in the hostTopNTable). The
    // type is interface{} with range: 0..2147483647.
    AlMatrixTopNControlTimeRemaining interface{}

    // The number of reports that have been generated by this entry. The type is
    // interface{} with range: 0..4294967295.
    AlMatrixTopNControlGeneratedReports interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval.  When the associated alMatrixTopNControlTimeRemaining
    // object is set, this object shall be set by the probe to the same value and
    // shall not be modified until the next time the
    // alMatrixTopNControlTimeRemaining is set.  This value shall be zero if no
    // reports have been requested for this alMatrixTopNControlEntry. The type is
    // interface{} with range: -2147483648..2147483647.
    AlMatrixTopNControlDuration interface{}

    // The maximum number of matrix entries requested for this report.  When this
    // object is created or modified, the probe should set
    // alMatrixTopNControlGrantedSize as closely to this object as is possible for
    // the particular probe implementation and available resources. The type is
    // interface{} with range: 0..2147483647.
    AlMatrixTopNControlRequestedSize interface{}

    // The maximum number of matrix entries in this report.  When the associated
    // alMatrixTopNControlRequestedSize object is created or modified, the probe
    // should set this      object as closely to the requested value as is
    // possible for the particular implementation and available resources. The
    // probe must not lower this value except as a result of a set to the
    // associated alMatrixTopNControlRequestedSize object.  If the value of
    // alMatrixTopNControlRateBase is equal to alMatrixTopNTerminalsPkts or
    // alMatrixTopNAllPkts, when the next topN report is generated, matrix entries
    // with the highest value of alMatrixTopNPktRate shall be placed in this table
    // in decreasing order of this rate until there is no more room or until there
    // are no more matrix entries.  If the value of alMatrixTopNControlRateBase is
    // equal to alMatrixTopNTerminalsOctets or alMatrixTopNAllOctets, when the
    // next topN report is generated, matrix entries with the highest value of
    // alMatrixTopNOctetRate shall be placed in this table in decreasing order of
    // this rate until there is no more room or until there are no more matrix
    // entries.  It is an implementation-specific matter how entries with the same
    // value of alMatrixTopNPktRate or alMatrixTopNOctetRate are sorted.  It is
    // also an implementation-specific matter as to whether or not zero-valued
    // entries are available. The type is interface{} with range: 0..2147483647.
    AlMatrixTopNControlGrantedSize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated
    // alMatrixTopNControlTimeRemaining object was modified to start the requested
    // report or the time the report was last automatically (re)started.  This
    // object may be used by the management station to determine if a report was
    // missed or not. The type is interface{} with range: 0..4294967295.
    AlMatrixTopNControlStartTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    AlMatrixTopNControlOwner interface{}

    // The status of this alMatrixTopNControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // alMatrixTopNTable shall be deleted by the agent. The type is RowStatus.
    AlMatrixTopNControlStatus interface{}
}

func (alMatrixTopNControlEntry *RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry) GetEntityData() *types.CommonEntityData {
    alMatrixTopNControlEntry.EntityData.YFilter = alMatrixTopNControlEntry.YFilter
    alMatrixTopNControlEntry.EntityData.YangName = "alMatrixTopNControlEntry"
    alMatrixTopNControlEntry.EntityData.BundleName = "cisco_ios_xe"
    alMatrixTopNControlEntry.EntityData.ParentYangName = "alMatrixTopNControlTable"
    alMatrixTopNControlEntry.EntityData.SegmentPath = "alMatrixTopNControlEntry" + types.AddKeyToken(alMatrixTopNControlEntry.AlMatrixTopNControlIndex, "alMatrixTopNControlIndex")
    alMatrixTopNControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/alMatrixTopNControlTable/" + alMatrixTopNControlEntry.EntityData.SegmentPath
    alMatrixTopNControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixTopNControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixTopNControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixTopNControlEntry.EntityData.Children = types.NewOrderedMap()
    alMatrixTopNControlEntry.EntityData.Leafs = types.NewOrderedMap()
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlIndex", types.YLeaf{"AlMatrixTopNControlIndex", alMatrixTopNControlEntry.AlMatrixTopNControlIndex})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlMatrixIndex", types.YLeaf{"AlMatrixTopNControlMatrixIndex", alMatrixTopNControlEntry.AlMatrixTopNControlMatrixIndex})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlRateBase", types.YLeaf{"AlMatrixTopNControlRateBase", alMatrixTopNControlEntry.AlMatrixTopNControlRateBase})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlTimeRemaining", types.YLeaf{"AlMatrixTopNControlTimeRemaining", alMatrixTopNControlEntry.AlMatrixTopNControlTimeRemaining})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlGeneratedReports", types.YLeaf{"AlMatrixTopNControlGeneratedReports", alMatrixTopNControlEntry.AlMatrixTopNControlGeneratedReports})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlDuration", types.YLeaf{"AlMatrixTopNControlDuration", alMatrixTopNControlEntry.AlMatrixTopNControlDuration})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlRequestedSize", types.YLeaf{"AlMatrixTopNControlRequestedSize", alMatrixTopNControlEntry.AlMatrixTopNControlRequestedSize})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlGrantedSize", types.YLeaf{"AlMatrixTopNControlGrantedSize", alMatrixTopNControlEntry.AlMatrixTopNControlGrantedSize})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlStartTime", types.YLeaf{"AlMatrixTopNControlStartTime", alMatrixTopNControlEntry.AlMatrixTopNControlStartTime})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlOwner", types.YLeaf{"AlMatrixTopNControlOwner", alMatrixTopNControlEntry.AlMatrixTopNControlOwner})
    alMatrixTopNControlEntry.EntityData.Leafs.Append("alMatrixTopNControlStatus", types.YLeaf{"AlMatrixTopNControlStatus", alMatrixTopNControlEntry.AlMatrixTopNControlStatus})

    alMatrixTopNControlEntry.EntityData.YListKeys = []string {"AlMatrixTopNControlIndex"}

    return &(alMatrixTopNControlEntry.EntityData)
}

// RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase represents alMatrixTopNControlStatus object is equal to active(1).
type RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase string

const (
    RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase_alMatrixTopNTerminalsPkts RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase = "alMatrixTopNTerminalsPkts"

    RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase_alMatrixTopNTerminalsOctets RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase = "alMatrixTopNTerminalsOctets"

    RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase_alMatrixTopNAllPkts RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase = "alMatrixTopNAllPkts"

    RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase_alMatrixTopNAllOctets RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlRateBase = "alMatrixTopNAllOctets"
)

// RMON2MIB_AlMatrixTopNTable
// A set of statistics for those application layer matrix
// entries that have counted the highest number of octets or
// packets.
type RMON2MIB_AlMatrixTopNTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNTable.  The alMatrixTopNControlIndex
    // value in the index identifies the alMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // alMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_AlMatrixTopNTable_AlMatrixTopNEntry.
    AlMatrixTopNEntry []*RMON2MIB_AlMatrixTopNTable_AlMatrixTopNEntry
}

func (alMatrixTopNTable *RMON2MIB_AlMatrixTopNTable) GetEntityData() *types.CommonEntityData {
    alMatrixTopNTable.EntityData.YFilter = alMatrixTopNTable.YFilter
    alMatrixTopNTable.EntityData.YangName = "alMatrixTopNTable"
    alMatrixTopNTable.EntityData.BundleName = "cisco_ios_xe"
    alMatrixTopNTable.EntityData.ParentYangName = "RMON2-MIB"
    alMatrixTopNTable.EntityData.SegmentPath = "alMatrixTopNTable"
    alMatrixTopNTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + alMatrixTopNTable.EntityData.SegmentPath
    alMatrixTopNTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixTopNTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixTopNTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixTopNTable.EntityData.Children = types.NewOrderedMap()
    alMatrixTopNTable.EntityData.Children.Append("alMatrixTopNEntry", types.YChild{"AlMatrixTopNEntry", nil})
    for i := range alMatrixTopNTable.AlMatrixTopNEntry {
        alMatrixTopNTable.EntityData.Children.Append(types.GetSegmentPath(alMatrixTopNTable.AlMatrixTopNEntry[i]), types.YChild{"AlMatrixTopNEntry", alMatrixTopNTable.AlMatrixTopNEntry[i]})
    }
    alMatrixTopNTable.EntityData.Leafs = types.NewOrderedMap()

    alMatrixTopNTable.EntityData.YListKeys = []string {}

    return &(alMatrixTopNTable.EntityData)
}

// RMON2MIB_AlMatrixTopNTable_AlMatrixTopNEntry
// A conceptual row in the alMatrixTopNTable.
// 
// The alMatrixTopNControlIndex value in the index identifies
// the alMatrixTopNControlEntry on whose behalf this entry was
// created.
// 
// An example of the indexing of this table is
// alMatrixTopNPktRate.3.10
type RMON2MIB_AlMatrixTopNTable_AlMatrixTopNEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_AlMatrixTopNControlTable_AlMatrixTopNControlEntry_AlMatrixTopNControlIndex
    AlMatrixTopNControlIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // alMatrixTopNTable among those in the same report. This index is between 1
    // and N, where N is the number of entries in this report.  If the value of
    // alMatrixTopNControlRateBase is equal to alMatrixTopNTerminalsPkts or
    // alMatrixTopNAllPkts, increasing values of alMatrixTopNIndex shall be
    // assigned to entries with decreasing values of alMatrixTopNPktRate until
    // index N is assigned or there are no more alMatrixTopNEntries.  If the value
    // of alMatrixTopNControlRateBase is equal to alMatrixTopNTerminalsOctets or
    // alMatrixTopNAllOctets, increasing values of alMatrixTopNIndex shall be
    // assigned to entries with decreasing values of alMatrixTopNOctetRate until
    // index N is assigned or there are no more alMatrixTopNEntries. The type is
    // interface{} with range: 1..65535.
    AlMatrixTopNIndex interface{}

    // The protocolDirLocalIndex of the network layer protocol of this entry's
    // network address. The type is interface{} with range: 1..2147483647.
    AlMatrixTopNProtocolDirLocalIndex interface{}

    // The network layer address of the source host in this conversation. This is
    // represented as an octet string with specific semantics and length as
    // identified      by the associated alMatrixTopNProtocolDirLocalIndex.  For
    // example, if the alMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    AlMatrixTopNSourceAddress interface{}

    // The network layer address of the destination host in this conversation. 
    // This is represented as an octet string with specific semantics and length
    // as identified by the associated alMatrixTopNProtocolDirLocalIndex.  For
    // example, if the alMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    AlMatrixTopNDestAddress interface{}

    // The type of the protocol counted by this matrix entry. The type is
    // interface{} with range: 1..2147483647.
    AlMatrixTopNAppProtocolDirLocalIndex interface{}

    // The number of packets seen of this protocol from the source host to the
    // destination host during this sampling interval, counted using the rules for
    // counting the alMatrixSDPkts object.  If the value of
    // alMatrixTopNControlRateBase is alMatrixTopNTerminalsPkts or
    // alMatrixTopNAllPkts, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    AlMatrixTopNPktRate interface{}

    // The number of packets seen of this protocol from the destination host to
    // the source host during this sampling interval, counted using the rules for
    // counting the alMatrixDSPkts object  (note that the corresponding
    // alMatrixSDPkts object selected is the one whose source address is equal to
    // alMatrixTopNDestAddress and whose destination address is equal to
    // alMatrixTopNSourceAddress.)  Note that if the value of
    // alMatrixTopNControlRateBase is equal to alMatrixTopNTerminalsPkts or
    // alMatrixTopNAllPkts, the sort of topN entries is based entirely on
    // alMatrixTopNPktRate, and not on the value of this object. The type is
    // interface{} with range: 0..4294967295.
    AlMatrixTopNReversePktRate interface{}

    // The number of octets seen of this protocol from the source host to the
    // destination host during this sampling interval, counted using the rules for
    // counting the alMatrixSDOctets object.  If the value of
    // alMatrixTopNControlRateBase is alMatrixTopNTerminalsOctets or
    // alMatrixTopNAllOctets, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    AlMatrixTopNOctetRate interface{}

    // The number of octets seen of this protocol from the destination host to the
    // source host during this sampling interval, counted using the rules for
    // counting the alMatrixDSOctets object  (note that the corresponding
    // alMatrixSDOctets object selected is the one whose source address is equal
    // to alMatrixTopNDestAddress and whose destination address is equal to
    // alMatrixTopNSourceAddress.)  Note that if the value of
    // alMatrixTopNControlRateBase is equal      to alMatrixTopNTerminalsOctets or
    // alMatrixTopNAllOctets, the sort of topN entries is based entirely on
    // alMatrixTopNOctetRate, and not on the value of this object. The type is
    // interface{} with range: 0..4294967295.
    AlMatrixTopNReverseOctetRate interface{}
}

func (alMatrixTopNEntry *RMON2MIB_AlMatrixTopNTable_AlMatrixTopNEntry) GetEntityData() *types.CommonEntityData {
    alMatrixTopNEntry.EntityData.YFilter = alMatrixTopNEntry.YFilter
    alMatrixTopNEntry.EntityData.YangName = "alMatrixTopNEntry"
    alMatrixTopNEntry.EntityData.BundleName = "cisco_ios_xe"
    alMatrixTopNEntry.EntityData.ParentYangName = "alMatrixTopNTable"
    alMatrixTopNEntry.EntityData.SegmentPath = "alMatrixTopNEntry" + types.AddKeyToken(alMatrixTopNEntry.AlMatrixTopNControlIndex, "alMatrixTopNControlIndex") + types.AddKeyToken(alMatrixTopNEntry.AlMatrixTopNIndex, "alMatrixTopNIndex")
    alMatrixTopNEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/alMatrixTopNTable/" + alMatrixTopNEntry.EntityData.SegmentPath
    alMatrixTopNEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alMatrixTopNEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alMatrixTopNEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alMatrixTopNEntry.EntityData.Children = types.NewOrderedMap()
    alMatrixTopNEntry.EntityData.Leafs = types.NewOrderedMap()
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNControlIndex", types.YLeaf{"AlMatrixTopNControlIndex", alMatrixTopNEntry.AlMatrixTopNControlIndex})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNIndex", types.YLeaf{"AlMatrixTopNIndex", alMatrixTopNEntry.AlMatrixTopNIndex})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNProtocolDirLocalIndex", types.YLeaf{"AlMatrixTopNProtocolDirLocalIndex", alMatrixTopNEntry.AlMatrixTopNProtocolDirLocalIndex})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNSourceAddress", types.YLeaf{"AlMatrixTopNSourceAddress", alMatrixTopNEntry.AlMatrixTopNSourceAddress})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNDestAddress", types.YLeaf{"AlMatrixTopNDestAddress", alMatrixTopNEntry.AlMatrixTopNDestAddress})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNAppProtocolDirLocalIndex", types.YLeaf{"AlMatrixTopNAppProtocolDirLocalIndex", alMatrixTopNEntry.AlMatrixTopNAppProtocolDirLocalIndex})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNPktRate", types.YLeaf{"AlMatrixTopNPktRate", alMatrixTopNEntry.AlMatrixTopNPktRate})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNReversePktRate", types.YLeaf{"AlMatrixTopNReversePktRate", alMatrixTopNEntry.AlMatrixTopNReversePktRate})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNOctetRate", types.YLeaf{"AlMatrixTopNOctetRate", alMatrixTopNEntry.AlMatrixTopNOctetRate})
    alMatrixTopNEntry.EntityData.Leafs.Append("alMatrixTopNReverseOctetRate", types.YLeaf{"AlMatrixTopNReverseOctetRate", alMatrixTopNEntry.AlMatrixTopNReverseOctetRate})

    alMatrixTopNEntry.EntityData.YListKeys = []string {"AlMatrixTopNControlIndex", "AlMatrixTopNIndex"}

    return &(alMatrixTopNEntry.EntityData)
}

// RMON2MIB_UsrHistoryControlTable
// A list of data-collection configuration entries.
type RMON2MIB_UsrHistoryControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a group of user-defined MIB objects to be
    // sampled periodically (called a bucket-group).  For example, an instance of
    // usrHistoryControlInterval might be named usrHistoryControlInterval.1. The
    // type is slice of RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry.
    UsrHistoryControlEntry []*RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry
}

func (usrHistoryControlTable *RMON2MIB_UsrHistoryControlTable) GetEntityData() *types.CommonEntityData {
    usrHistoryControlTable.EntityData.YFilter = usrHistoryControlTable.YFilter
    usrHistoryControlTable.EntityData.YangName = "usrHistoryControlTable"
    usrHistoryControlTable.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryControlTable.EntityData.ParentYangName = "RMON2-MIB"
    usrHistoryControlTable.EntityData.SegmentPath = "usrHistoryControlTable"
    usrHistoryControlTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + usrHistoryControlTable.EntityData.SegmentPath
    usrHistoryControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryControlTable.EntityData.Children = types.NewOrderedMap()
    usrHistoryControlTable.EntityData.Children.Append("usrHistoryControlEntry", types.YChild{"UsrHistoryControlEntry", nil})
    for i := range usrHistoryControlTable.UsrHistoryControlEntry {
        usrHistoryControlTable.EntityData.Children.Append(types.GetSegmentPath(usrHistoryControlTable.UsrHistoryControlEntry[i]), types.YChild{"UsrHistoryControlEntry", usrHistoryControlTable.UsrHistoryControlEntry[i]})
    }
    usrHistoryControlTable.EntityData.Leafs = types.NewOrderedMap()

    usrHistoryControlTable.EntityData.YListKeys = []string {}

    return &(usrHistoryControlTable.EntityData)
}

// RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry
// A list of parameters that set up a group of user-defined
// MIB objects to be sampled periodically (called a
// bucket-group).
// 
// For example, an instance of usrHistoryControlInterval
// might be named usrHistoryControlInterval.1
type RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely identifies an entry in the
    // usrHistoryControlTable.  Each such entry defines a set of samples at a
    // particular interval for a specified set of MIB instances available from the
    // managed system. The type is interface{} with range: 1..65535.
    UsrHistoryControlIndex interface{}

    // The number of MIB objects to be collected in the portion of usrHistoryTable
    // associated with this usrHistoryControlEntry.  This object may not be
    // modified if the associated instance of usrHistoryControlStatus is equal to
    // active(1). The type is interface{} with range: 1..65535.
    UsrHistoryControlObjects interface{}

    // The requested number of discrete time intervals over which data is to be
    // saved in the part of the usrHistoryTable associated with this
    // usrHistoryControlEntry.  When this object is created or modified, the probe
    // should set usrHistoryControlBucketsGranted as closely to this object as is
    // possible for the particular probe implementation and available resources.
    // The type is interface{} with range: 1..65535.
    UsrHistoryControlBucketsRequested interface{}

    // The number of discrete sampling intervals over which data shall be saved in
    // the part of the usrHistoryTable associated with this
    // usrHistoryControlEntry.  When the associated
    // usrHistoryControlBucketsRequested      object is created or modified, the
    // probe should set this object as closely to the requested value as is
    // possible for the particular  probe implementation and available resources. 
    // The probe must not lower this value except as a result of a modification to
    // the associated usrHistoryControlBucketsRequested object.  The associated
    // usrHistoryControlBucketsRequested object should be set before or at the
    // same time as this object to allow the probe to accurately estimate the
    // resources required for this usrHistoryControlEntry.  There will be times
    // when the actual number of buckets associated with this entry is less than
    // the value of this object.  In this case, at the end of each sampling
    // interval, a new bucket will be added to the usrHistoryTable.  When the
    // number of buckets reaches the value of this object and a new bucket is to
    // be added to the usrHistoryTable, the oldest bucket associated with this
    // usrHistoryControlEntry shall be deleted by the agent so that the new bucket
    // can be added.  When the value of this object changes to a value less than
    // the current value, entries are deleted from the usrHistoryTable associated
    // with this usrHistoryControlEntry. Enough of the oldest of these entries
    // shall be deleted by the agent so that their number remains less than or
    // equal to the new value of this object.  When the value of this object
    // changes to a value greater than the current value, the number of associated
    // usrHistory entries may be allowed to grow. The type is interface{} with
    // range: 1..65535.
    UsrHistoryControlBucketsGranted interface{}

    // The interval in seconds over which the data is sampled for each bucket in
    // the part of the usrHistory table associated with this
    // usrHistoryControlEntry.  Because the counters in a bucket may overflow at
    // their maximum value with no indication, a prudent manager will take into
    // account the possibility of overflow in any of      the associated counters.
    // It is important to consider the minimum time in which any counter could
    // overflow on a particular media type and set the usrHistoryControlInterval
    // object to a value less than this interval.  This object may not be modified
    // if the associated usrHistoryControlStatus object is equal to active(1). The
    // type is interface{} with range: 1..2147483647.
    UsrHistoryControlInterval interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    UsrHistoryControlOwner interface{}

    // The status of this variable history control entry.  An entry may not exist
    // in the active state unless all objects in the entry have an appropriate
    // value.  If this object is not equal to active(1), all associated entries in
    // the usrHistoryTable shall be deleted. The type is RowStatus.
    UsrHistoryControlStatus interface{}
}

func (usrHistoryControlEntry *RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry) GetEntityData() *types.CommonEntityData {
    usrHistoryControlEntry.EntityData.YFilter = usrHistoryControlEntry.YFilter
    usrHistoryControlEntry.EntityData.YangName = "usrHistoryControlEntry"
    usrHistoryControlEntry.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryControlEntry.EntityData.ParentYangName = "usrHistoryControlTable"
    usrHistoryControlEntry.EntityData.SegmentPath = "usrHistoryControlEntry" + types.AddKeyToken(usrHistoryControlEntry.UsrHistoryControlIndex, "usrHistoryControlIndex")
    usrHistoryControlEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/usrHistoryControlTable/" + usrHistoryControlEntry.EntityData.SegmentPath
    usrHistoryControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryControlEntry.EntityData.Children = types.NewOrderedMap()
    usrHistoryControlEntry.EntityData.Leafs = types.NewOrderedMap()
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlIndex", types.YLeaf{"UsrHistoryControlIndex", usrHistoryControlEntry.UsrHistoryControlIndex})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlObjects", types.YLeaf{"UsrHistoryControlObjects", usrHistoryControlEntry.UsrHistoryControlObjects})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlBucketsRequested", types.YLeaf{"UsrHistoryControlBucketsRequested", usrHistoryControlEntry.UsrHistoryControlBucketsRequested})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlBucketsGranted", types.YLeaf{"UsrHistoryControlBucketsGranted", usrHistoryControlEntry.UsrHistoryControlBucketsGranted})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlInterval", types.YLeaf{"UsrHistoryControlInterval", usrHistoryControlEntry.UsrHistoryControlInterval})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlOwner", types.YLeaf{"UsrHistoryControlOwner", usrHistoryControlEntry.UsrHistoryControlOwner})
    usrHistoryControlEntry.EntityData.Leafs.Append("usrHistoryControlStatus", types.YLeaf{"UsrHistoryControlStatus", usrHistoryControlEntry.UsrHistoryControlStatus})

    usrHistoryControlEntry.EntityData.YListKeys = []string {"UsrHistoryControlIndex"}

    return &(usrHistoryControlEntry.EntityData)
}

// RMON2MIB_UsrHistoryObjectTable
// A list of data-collection configuration entries.
type RMON2MIB_UsrHistoryObjectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of MIB instances to be sampled periodically.  Entries in this table
    // are created when an associated usrHistoryControlObjects object is created. 
    // The usrHistoryControlIndex value in the index is that of the associated
    // usrHistoryControlEntry.  For example, an instance of
    // usrHistoryObjectVariable might be usrHistoryObjectVariable.1.3. The type is
    // slice of RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry.
    UsrHistoryObjectEntry []*RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry
}

func (usrHistoryObjectTable *RMON2MIB_UsrHistoryObjectTable) GetEntityData() *types.CommonEntityData {
    usrHistoryObjectTable.EntityData.YFilter = usrHistoryObjectTable.YFilter
    usrHistoryObjectTable.EntityData.YangName = "usrHistoryObjectTable"
    usrHistoryObjectTable.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryObjectTable.EntityData.ParentYangName = "RMON2-MIB"
    usrHistoryObjectTable.EntityData.SegmentPath = "usrHistoryObjectTable"
    usrHistoryObjectTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + usrHistoryObjectTable.EntityData.SegmentPath
    usrHistoryObjectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryObjectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryObjectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryObjectTable.EntityData.Children = types.NewOrderedMap()
    usrHistoryObjectTable.EntityData.Children.Append("usrHistoryObjectEntry", types.YChild{"UsrHistoryObjectEntry", nil})
    for i := range usrHistoryObjectTable.UsrHistoryObjectEntry {
        usrHistoryObjectTable.EntityData.Children.Append(types.GetSegmentPath(usrHistoryObjectTable.UsrHistoryObjectEntry[i]), types.YChild{"UsrHistoryObjectEntry", usrHistoryObjectTable.UsrHistoryObjectEntry[i]})
    }
    usrHistoryObjectTable.EntityData.Leafs = types.NewOrderedMap()

    usrHistoryObjectTable.EntityData.YListKeys = []string {}

    return &(usrHistoryObjectTable.EntityData)
}

// RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry
// A list of MIB instances to be sampled periodically.
// 
// Entries in this table are created when an associated
// usrHistoryControlObjects object is created.
// 
// The usrHistoryControlIndex value in the index is
// that of the associated usrHistoryControlEntry.
// 
// For example, an instance of usrHistoryObjectVariable might be
// usrHistoryObjectVariable.1.3
type RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry_UsrHistoryControlIndex
    UsrHistoryControlIndex interface{}

    // This attribute is a key. An index used to uniquely identify an entry in the
    // usrHistoryObject table.  Each such entry defines a MIB instance to be
    // collected periodically. The type is interface{} with range: 1..65535.
    UsrHistoryObjectIndex interface{}

    // The object identifier of the particular variable to be sampled.  Only
    // variables that resolve to an ASN.1 primitive type of Integer32 (Integer32,
    // Counter, Gauge, or TimeTicks) may be sampled.  Because SNMP access control
    // is articulated entirely in terms of the contents of MIB views, no access
    // control mechanism exists that can restrict the value of this object to
    // identify only those objects that exist in a particular MIB view. Because
    // there is thus no acceptable means of restricting the read access that could
    // be obtained through the user history      mechanism, the probe must only
    // grant write access to this object in those views that have read access to
    // all objects on the probe.  During a set operation, if the supplied variable
    // name is not available in the selected MIB view, a badValue error must be
    // returned.  This object may not be modified if the associated
    // usrHistoryControlStatus object is equal to active(1). The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    UsrHistoryObjectVariable interface{}

    // The method of sampling the selected variable for storage in the
    // usrHistoryTable.  If the value of this object is absoluteValue(1), the
    // value of the selected variable will be copied directly into the history
    // bucket.  If the value of this object is deltaValue(2), the value of the
    // selected variable at the last sample will be subtracted from the current
    // value, and the difference will be stored in the history bucket. If the
    // associated usrHistoryObjectVariable instance could not be obtained at the
    // previous sample interval, then a delta sample is not possible, and the
    // value of the associated usrHistoryValStatus object for this interval will
    // be valueNotAvailable(1).  This object may not be modified if the associated
    // usrHistoryControlStatus object is equal to active(1). The type is
    // UsrHistoryObjectSampleType.
    UsrHistoryObjectSampleType interface{}
}

func (usrHistoryObjectEntry *RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry) GetEntityData() *types.CommonEntityData {
    usrHistoryObjectEntry.EntityData.YFilter = usrHistoryObjectEntry.YFilter
    usrHistoryObjectEntry.EntityData.YangName = "usrHistoryObjectEntry"
    usrHistoryObjectEntry.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryObjectEntry.EntityData.ParentYangName = "usrHistoryObjectTable"
    usrHistoryObjectEntry.EntityData.SegmentPath = "usrHistoryObjectEntry" + types.AddKeyToken(usrHistoryObjectEntry.UsrHistoryControlIndex, "usrHistoryControlIndex") + types.AddKeyToken(usrHistoryObjectEntry.UsrHistoryObjectIndex, "usrHistoryObjectIndex")
    usrHistoryObjectEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/usrHistoryObjectTable/" + usrHistoryObjectEntry.EntityData.SegmentPath
    usrHistoryObjectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryObjectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryObjectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryObjectEntry.EntityData.Children = types.NewOrderedMap()
    usrHistoryObjectEntry.EntityData.Leafs = types.NewOrderedMap()
    usrHistoryObjectEntry.EntityData.Leafs.Append("usrHistoryControlIndex", types.YLeaf{"UsrHistoryControlIndex", usrHistoryObjectEntry.UsrHistoryControlIndex})
    usrHistoryObjectEntry.EntityData.Leafs.Append("usrHistoryObjectIndex", types.YLeaf{"UsrHistoryObjectIndex", usrHistoryObjectEntry.UsrHistoryObjectIndex})
    usrHistoryObjectEntry.EntityData.Leafs.Append("usrHistoryObjectVariable", types.YLeaf{"UsrHistoryObjectVariable", usrHistoryObjectEntry.UsrHistoryObjectVariable})
    usrHistoryObjectEntry.EntityData.Leafs.Append("usrHistoryObjectSampleType", types.YLeaf{"UsrHistoryObjectSampleType", usrHistoryObjectEntry.UsrHistoryObjectSampleType})

    usrHistoryObjectEntry.EntityData.YListKeys = []string {"UsrHistoryControlIndex", "UsrHistoryObjectIndex"}

    return &(usrHistoryObjectEntry.EntityData)
}

// RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType represents usrHistoryControlStatus object is equal to active(1).
type RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType string

const (
    RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType_absoluteValue RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType = "absoluteValue"

    RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType_deltaValue RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectSampleType = "deltaValue"
)

// RMON2MIB_UsrHistoryTable
// A list of user defined history entries.
type RMON2MIB_UsrHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A historical sample of user-defined variables.  This sample is associated
    // with the usrHistoryControlEntry which set up the parameters for a regular
    // collection of these samples.  The usrHistoryControlIndex value in the index
    // identifies the usrHistoryControlEntry on whose behalf this entry was
    // created.  The usrHistoryObjectIndex value in the index identifies the
    // usrHistoryObjectEntry on whose behalf this entry was created.  For example,
    // an instance of usrHistoryAbsValue, which represents the 14th sample of a
    // variable collected as specified by usrHistoryControlEntry.1 and
    // usrHistoryObjectEntry.1.5, would be named usrHistoryAbsValue.1.14.5. The
    // type is slice of RMON2MIB_UsrHistoryTable_UsrHistoryEntry.
    UsrHistoryEntry []*RMON2MIB_UsrHistoryTable_UsrHistoryEntry
}

func (usrHistoryTable *RMON2MIB_UsrHistoryTable) GetEntityData() *types.CommonEntityData {
    usrHistoryTable.EntityData.YFilter = usrHistoryTable.YFilter
    usrHistoryTable.EntityData.YangName = "usrHistoryTable"
    usrHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryTable.EntityData.ParentYangName = "RMON2-MIB"
    usrHistoryTable.EntityData.SegmentPath = "usrHistoryTable"
    usrHistoryTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + usrHistoryTable.EntityData.SegmentPath
    usrHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryTable.EntityData.Children = types.NewOrderedMap()
    usrHistoryTable.EntityData.Children.Append("usrHistoryEntry", types.YChild{"UsrHistoryEntry", nil})
    for i := range usrHistoryTable.UsrHistoryEntry {
        usrHistoryTable.EntityData.Children.Append(types.GetSegmentPath(usrHistoryTable.UsrHistoryEntry[i]), types.YChild{"UsrHistoryEntry", usrHistoryTable.UsrHistoryEntry[i]})
    }
    usrHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    usrHistoryTable.EntityData.YListKeys = []string {}

    return &(usrHistoryTable.EntityData)
}

// RMON2MIB_UsrHistoryTable_UsrHistoryEntry
// A historical sample of user-defined variables.  This sample
// is associated with the usrHistoryControlEntry which set up the
// parameters for a regular collection of these samples.
// 
// The usrHistoryControlIndex value in the index identifies the
// usrHistoryControlEntry on whose behalf this entry was created.
// 
// The usrHistoryObjectIndex value in the index identifies the
// usrHistoryObjectEntry on whose behalf this entry was created.
// 
// For example, an instance of usrHistoryAbsValue, which represents
// the 14th sample of a variable collected as specified by
// usrHistoryControlEntry.1 and usrHistoryObjectEntry.1.5,
// would be named usrHistoryAbsValue.1.14.5
type RMON2MIB_UsrHistoryTable_UsrHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_UsrHistoryControlTable_UsrHistoryControlEntry_UsrHistoryControlIndex
    UsrHistoryControlIndex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // usrHistoryControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range: 1..2147483647.
    UsrHistorySampleIndex interface{}

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_UsrHistoryObjectTable_UsrHistoryObjectEntry_UsrHistoryObjectIndex
    UsrHistoryObjectIndex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval. Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    UsrHistoryIntervalStart interface{}

    // The value of sysUpTime at the end of the interval over which this sample
    // was measured. The type is interface{} with range: 0..4294967295.
    UsrHistoryIntervalEnd interface{}

    // The absolute value (i.e. unsigned value) of the user-specified statistic
    // during the last sampling period. The value during the current sampling
    // period is not made available until the period is completed.  To obtain the
    // true value for this sampling interval, the associated instance of
    // usrHistoryValStatus must be checked, and usrHistoryAbsValue adjusted as
    // necessary.  If the MIB instance could not be accessed during the sampling
    // interval, then this object will have a value of zero and the associated
    // instance of usrHistoryValStatus will be set to 'valueNotAvailable(1)'. The
    // type is interface{} with range: 0..4294967295.
    UsrHistoryAbsValue interface{}

    // This object indicates the validity and sign of the data in the associated
    // instance of usrHistoryAbsValue.  If the MIB instance could not be accessed
    // during the sampling interval, then 'valueNotAvailable(1)' will be returned.
    // If the sample is valid and actual value of the sample is greater than or
    // equal to zero then 'valuePositive(2)' is returned.  If the sample is valid
    // and the actual value of the sample is less than zero, 'valueNegative(3)'
    // will be returned. The associated instance of usrHistoryAbsValue should be
    // multiplied by -1 to obtain the true sample value. The type is
    // UsrHistoryValStatus.
    UsrHistoryValStatus interface{}
}

func (usrHistoryEntry *RMON2MIB_UsrHistoryTable_UsrHistoryEntry) GetEntityData() *types.CommonEntityData {
    usrHistoryEntry.EntityData.YFilter = usrHistoryEntry.YFilter
    usrHistoryEntry.EntityData.YangName = "usrHistoryEntry"
    usrHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    usrHistoryEntry.EntityData.ParentYangName = "usrHistoryTable"
    usrHistoryEntry.EntityData.SegmentPath = "usrHistoryEntry" + types.AddKeyToken(usrHistoryEntry.UsrHistoryControlIndex, "usrHistoryControlIndex") + types.AddKeyToken(usrHistoryEntry.UsrHistorySampleIndex, "usrHistorySampleIndex") + types.AddKeyToken(usrHistoryEntry.UsrHistoryObjectIndex, "usrHistoryObjectIndex")
    usrHistoryEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/usrHistoryTable/" + usrHistoryEntry.EntityData.SegmentPath
    usrHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrHistoryEntry.EntityData.Children = types.NewOrderedMap()
    usrHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryControlIndex", types.YLeaf{"UsrHistoryControlIndex", usrHistoryEntry.UsrHistoryControlIndex})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistorySampleIndex", types.YLeaf{"UsrHistorySampleIndex", usrHistoryEntry.UsrHistorySampleIndex})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryObjectIndex", types.YLeaf{"UsrHistoryObjectIndex", usrHistoryEntry.UsrHistoryObjectIndex})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryIntervalStart", types.YLeaf{"UsrHistoryIntervalStart", usrHistoryEntry.UsrHistoryIntervalStart})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryIntervalEnd", types.YLeaf{"UsrHistoryIntervalEnd", usrHistoryEntry.UsrHistoryIntervalEnd})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryAbsValue", types.YLeaf{"UsrHistoryAbsValue", usrHistoryEntry.UsrHistoryAbsValue})
    usrHistoryEntry.EntityData.Leafs.Append("usrHistoryValStatus", types.YLeaf{"UsrHistoryValStatus", usrHistoryEntry.UsrHistoryValStatus})

    usrHistoryEntry.EntityData.YListKeys = []string {"UsrHistoryControlIndex", "UsrHistorySampleIndex", "UsrHistoryObjectIndex"}

    return &(usrHistoryEntry.EntityData)
}

// RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus represents by -1 to obtain the true sample value.
type RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus string

const (
    RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus_valueNotAvailable RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus = "valueNotAvailable"

    RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus_valuePositive RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus = "valuePositive"

    RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus_valueNegative RMON2MIB_UsrHistoryTable_UsrHistoryEntry_UsrHistoryValStatus = "valueNegative"
)

// RMON2MIB_SerialConfigTable
// A table of serial interface configuration entries.  This data
// will be stored in non-volatile memory and preserved across
// probe resets or power loss.
type RMON2MIB_SerialConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular serial interface on this
    // device. If the device has no serial interfaces, this table is empty.  The
    // index is composed of the ifIndex assigned to this serial line interface.
    // The type is slice of RMON2MIB_SerialConfigTable_SerialConfigEntry.
    SerialConfigEntry []*RMON2MIB_SerialConfigTable_SerialConfigEntry
}

func (serialConfigTable *RMON2MIB_SerialConfigTable) GetEntityData() *types.CommonEntityData {
    serialConfigTable.EntityData.YFilter = serialConfigTable.YFilter
    serialConfigTable.EntityData.YangName = "serialConfigTable"
    serialConfigTable.EntityData.BundleName = "cisco_ios_xe"
    serialConfigTable.EntityData.ParentYangName = "RMON2-MIB"
    serialConfigTable.EntityData.SegmentPath = "serialConfigTable"
    serialConfigTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + serialConfigTable.EntityData.SegmentPath
    serialConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialConfigTable.EntityData.Children = types.NewOrderedMap()
    serialConfigTable.EntityData.Children.Append("serialConfigEntry", types.YChild{"SerialConfigEntry", nil})
    for i := range serialConfigTable.SerialConfigEntry {
        serialConfigTable.EntityData.Children.Append(types.GetSegmentPath(serialConfigTable.SerialConfigEntry[i]), types.YChild{"SerialConfigEntry", serialConfigTable.SerialConfigEntry[i]})
    }
    serialConfigTable.EntityData.Leafs = types.NewOrderedMap()

    serialConfigTable.EntityData.YListKeys = []string {}

    return &(serialConfigTable.EntityData)
}

// RMON2MIB_SerialConfigTable_SerialConfigEntry
// A set of configuration parameters for a particular
// serial interface on this device. If the device has no serial
// interfaces, this table is empty.
// 
// The index is composed of the ifIndex assigned to this serial
// line interface.
type RMON2MIB_SerialConfigTable_SerialConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range:
    // -2147483648..2147483647. Refers to
    // rfc1213_mib.RFC1213MIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The type of incoming connection to expect on this serial interface. The
    // type is SerialMode.
    SerialMode interface{}

    // The type of data link encapsulation to be used on this serial interface.
    // The type is SerialProtocol.
    SerialProtocol interface{}

    // This timeout value is used when the Management Station has initiated the
    // conversation over the serial link. This variable represents the number of
    // seconds of inactivity allowed before terminating the connection on this
    // serial interface. Use the      serialDialoutTimeout in the case where the
    // probe has initiated the connection for the purpose of sending a trap. The
    // type is interface{} with range: 1..65535.
    SerialTimeout interface{}

    // A control string which controls how a modem attached to this serial
    // interface should be initialized.  The initialization is performed once
    // during startup and again after each connection is terminated if the
    // associated serialMode has the value of modem(2).  A control string that is
    // appropriate for a wide variety of modems is: '^s^MATE0Q0V1X4 S0=1 S2=43^M'.
    // The type is string with length: 0..255.
    SerialModemInitString interface{}

    // A control string which specifies how to disconnect a modem connection on
    // this serial interface.  This object is only meaningful if the associated
    // serialMode has the value of modem(2). A control string that is appropriate
    // for a wide variety of modems is: '^d2^s+++^d2^sATH0^M^d2'. The type is
    // string with length: 0..255.
    SerialModemHangUpString interface{}

    // An ASCII string containing substrings that describe the expected modem
    // connection response code and associated bps rate.  The substrings are
    // delimited by the first character in the string, for example:   
    // /CONNECT/300/CONNECT 1200/1200/CONNECT 2400/2400/    CONNECT
    // 4800/4800/CONNECT 9600/9600 will be interpreted as:     response code   
    // bps rate     CONNECT            300     CONNECT 1200      1200         
    // CONNECT 2400      2400     CONNECT 4800      4800     CONNECT 9600     
    // 9600 The agent will use the information in this string to adjust the bps
    // rate of this serial interface once a modem connection is established.  A
    // value that is appropriate for a wide variety of modems is:
    // '/CONNECT/300/CONNECT 1200/1200/CONNECT 2400/2400/  CONNECT
    // 4800/4800/CONNECT 9600/9600/CONNECT 14400/14400/ CONNECT
    // 19200/19200/CONNECT 38400/38400/'. The type is string with length: 0..255.
    SerialModemConnectResp interface{}

    // An ASCII string containing response codes that may be generated by a modem
    // to report the reason why a connection attempt has failed.  The response
    // codes are delimited by the first character in the string, for example:   
    // /NO CARRIER/BUSY/NO DIALTONE/NO ANSWER/ERROR/ If one of these response
    // codes is received via this serial interface while attempting to make a
    // modem connection, the agent will issue the hang up command as specified by
    // serialModemHangUpString.  A value that is appropriate for a wide variety of
    // modems is: '/NO CARRIER/BUSY/NO DIALTONE/NO ANSWER/ERROR/'. The type is
    // string with length: 0..255.
    SerialModemNoConnectResp interface{}

    // This timeout value is used when the probe initiates the serial connection
    // with the intention of contacting a management station. This variable
    // represents the number of seconds of inactivity allowed before terminating
    // the connection on this serial interface. The type is interface{} with
    // range: 1..65535.
    SerialDialoutTimeout interface{}

    // The status of this serialConfigEntry.  An entry may not exist in the active
    // state unless all objects in the entry have an appropriate value. The type
    // is RowStatus.
    SerialStatus interface{}
}

func (serialConfigEntry *RMON2MIB_SerialConfigTable_SerialConfigEntry) GetEntityData() *types.CommonEntityData {
    serialConfigEntry.EntityData.YFilter = serialConfigEntry.YFilter
    serialConfigEntry.EntityData.YangName = "serialConfigEntry"
    serialConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    serialConfigEntry.EntityData.ParentYangName = "serialConfigTable"
    serialConfigEntry.EntityData.SegmentPath = "serialConfigEntry" + types.AddKeyToken(serialConfigEntry.IfIndex, "ifIndex")
    serialConfigEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/serialConfigTable/" + serialConfigEntry.EntityData.SegmentPath
    serialConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialConfigEntry.EntityData.Children = types.NewOrderedMap()
    serialConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    serialConfigEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", serialConfigEntry.IfIndex})
    serialConfigEntry.EntityData.Leafs.Append("serialMode", types.YLeaf{"SerialMode", serialConfigEntry.SerialMode})
    serialConfigEntry.EntityData.Leafs.Append("serialProtocol", types.YLeaf{"SerialProtocol", serialConfigEntry.SerialProtocol})
    serialConfigEntry.EntityData.Leafs.Append("serialTimeout", types.YLeaf{"SerialTimeout", serialConfigEntry.SerialTimeout})
    serialConfigEntry.EntityData.Leafs.Append("serialModemInitString", types.YLeaf{"SerialModemInitString", serialConfigEntry.SerialModemInitString})
    serialConfigEntry.EntityData.Leafs.Append("serialModemHangUpString", types.YLeaf{"SerialModemHangUpString", serialConfigEntry.SerialModemHangUpString})
    serialConfigEntry.EntityData.Leafs.Append("serialModemConnectResp", types.YLeaf{"SerialModemConnectResp", serialConfigEntry.SerialModemConnectResp})
    serialConfigEntry.EntityData.Leafs.Append("serialModemNoConnectResp", types.YLeaf{"SerialModemNoConnectResp", serialConfigEntry.SerialModemNoConnectResp})
    serialConfigEntry.EntityData.Leafs.Append("serialDialoutTimeout", types.YLeaf{"SerialDialoutTimeout", serialConfigEntry.SerialDialoutTimeout})
    serialConfigEntry.EntityData.Leafs.Append("serialStatus", types.YLeaf{"SerialStatus", serialConfigEntry.SerialStatus})

    serialConfigEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(serialConfigEntry.EntityData)
}

// RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode represents interface.
type RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode string

const (
    RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode_direct RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode = "direct"

    RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode_modem RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialMode = "modem"
)

// RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol represents serial interface.
type RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol string

const (
    RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol_other RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol = "other"

    RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol_slip RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol = "slip"

    RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol_ppp RMON2MIB_SerialConfigTable_SerialConfigEntry_SerialProtocol = "ppp"
)

// RMON2MIB_NetConfigTable
// A table of netConfigEntries.
type RMON2MIB_NetConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular network interface on
    // this device. If the device has no network interface, this table is empty. 
    // The index is composed of the ifIndex assigned to the corresponding
    // interface. The type is slice of RMON2MIB_NetConfigTable_NetConfigEntry.
    NetConfigEntry []*RMON2MIB_NetConfigTable_NetConfigEntry
}

func (netConfigTable *RMON2MIB_NetConfigTable) GetEntityData() *types.CommonEntityData {
    netConfigTable.EntityData.YFilter = netConfigTable.YFilter
    netConfigTable.EntityData.YangName = "netConfigTable"
    netConfigTable.EntityData.BundleName = "cisco_ios_xe"
    netConfigTable.EntityData.ParentYangName = "RMON2-MIB"
    netConfigTable.EntityData.SegmentPath = "netConfigTable"
    netConfigTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + netConfigTable.EntityData.SegmentPath
    netConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netConfigTable.EntityData.Children = types.NewOrderedMap()
    netConfigTable.EntityData.Children.Append("netConfigEntry", types.YChild{"NetConfigEntry", nil})
    for i := range netConfigTable.NetConfigEntry {
        netConfigTable.EntityData.Children.Append(types.GetSegmentPath(netConfigTable.NetConfigEntry[i]), types.YChild{"NetConfigEntry", netConfigTable.NetConfigEntry[i]})
    }
    netConfigTable.EntityData.Leafs = types.NewOrderedMap()

    netConfigTable.EntityData.YListKeys = []string {}

    return &(netConfigTable.EntityData)
}

// RMON2MIB_NetConfigTable_NetConfigEntry
// A set of configuration parameters for a particular
// network interface on this device. If the device has no network
// interface, this table is empty.
// 
// The index is composed of the ifIndex assigned to the
// corresponding interface.
type RMON2MIB_NetConfigTable_NetConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range:
    // -2147483648..2147483647. Refers to
    // rfc1213_mib.RFC1213MIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The IP address of this Net interface.  The default value for this object is
    // 0.0.0.0.  If either the netConfigIPAddress or netConfigSubnetMask are
    // 0.0.0.0, then when the device boots, it may use BOOTP to try to figure out
    // what these values should be. If BOOTP fails, before the device can talk on
    // the network, this value must be configured (e.g., through a terminal
    // attached to the device). If BOOTP is      used, care should be taken to not
    // send BOOTP broadcasts too frequently and to eventually send very
    // infrequently if no replies are received. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetConfigIPAddress interface{}

    // The subnet mask of this Net interface.  The default value for this object
    // is 0.0.0.0.  If either the netConfigIPAddress or netConfigSubnetMask are
    // 0.0.0.0, then when the device boots, it may use BOOTP to try to figure out
    // what these values should be. If BOOTP fails, before the device can talk on
    // the network, this value must be configured (e.g., through a terminal
    // attached to the device). If BOOTP is used, care should be taken to not send
    // BOOTP broadcasts too frequently and to eventually send very infrequently if
    // no replies are received. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetConfigSubnetMask interface{}

    // The status of this netConfigEntry.  An entry may not exist in the active
    // state unless all objects in the entry have an appropriate value. The type
    // is RowStatus.
    NetConfigStatus interface{}
}

func (netConfigEntry *RMON2MIB_NetConfigTable_NetConfigEntry) GetEntityData() *types.CommonEntityData {
    netConfigEntry.EntityData.YFilter = netConfigEntry.YFilter
    netConfigEntry.EntityData.YangName = "netConfigEntry"
    netConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    netConfigEntry.EntityData.ParentYangName = "netConfigTable"
    netConfigEntry.EntityData.SegmentPath = "netConfigEntry" + types.AddKeyToken(netConfigEntry.IfIndex, "ifIndex")
    netConfigEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/netConfigTable/" + netConfigEntry.EntityData.SegmentPath
    netConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netConfigEntry.EntityData.Children = types.NewOrderedMap()
    netConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    netConfigEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", netConfigEntry.IfIndex})
    netConfigEntry.EntityData.Leafs.Append("netConfigIPAddress", types.YLeaf{"NetConfigIPAddress", netConfigEntry.NetConfigIPAddress})
    netConfigEntry.EntityData.Leafs.Append("netConfigSubnetMask", types.YLeaf{"NetConfigSubnetMask", netConfigEntry.NetConfigSubnetMask})
    netConfigEntry.EntityData.Leafs.Append("netConfigStatus", types.YLeaf{"NetConfigStatus", netConfigEntry.NetConfigStatus})

    netConfigEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(netConfigEntry.EntityData)
}

// RMON2MIB_TrapDestTable
// A list of trap destination entries.
type RMON2MIB_TrapDestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry includes a destination IP address to which to send traps for
    // this community. The type is slice of RMON2MIB_TrapDestTable_TrapDestEntry.
    TrapDestEntry []*RMON2MIB_TrapDestTable_TrapDestEntry
}

func (trapDestTable *RMON2MIB_TrapDestTable) GetEntityData() *types.CommonEntityData {
    trapDestTable.EntityData.YFilter = trapDestTable.YFilter
    trapDestTable.EntityData.YangName = "trapDestTable"
    trapDestTable.EntityData.BundleName = "cisco_ios_xe"
    trapDestTable.EntityData.ParentYangName = "RMON2-MIB"
    trapDestTable.EntityData.SegmentPath = "trapDestTable"
    trapDestTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + trapDestTable.EntityData.SegmentPath
    trapDestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trapDestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trapDestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trapDestTable.EntityData.Children = types.NewOrderedMap()
    trapDestTable.EntityData.Children.Append("trapDestEntry", types.YChild{"TrapDestEntry", nil})
    for i := range trapDestTable.TrapDestEntry {
        trapDestTable.EntityData.Children.Append(types.GetSegmentPath(trapDestTable.TrapDestEntry[i]), types.YChild{"TrapDestEntry", trapDestTable.TrapDestEntry[i]})
    }
    trapDestTable.EntityData.Leafs = types.NewOrderedMap()

    trapDestTable.EntityData.YListKeys = []string {}

    return &(trapDestTable.EntityData)
}

// RMON2MIB_TrapDestTable_TrapDestEntry
// This entry includes a destination IP address to which to send
// traps for this community.
type RMON2MIB_TrapDestTable_TrapDestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A value that uniquely identifies this
    // trapDestEntry. The type is interface{} with range: 1..65535.
    TrapDestIndex interface{}

    // A community to which this destination address belongs. This entry is
    // associated with any eventEntries in the RMON      MIB whose value of
    // eventCommunity is equal to the value of this object.  Every time an
    // associated event entry sends a trap due to an event, that trap will be sent
    // to each address in the trapDestTable with a trapDestCommunity equal to
    // eventCommunity.  This object may not be modified if the associated
    // trapDestStatus object is equal to active(1). The type is string with
    // length: 0..127.
    TrapDestCommunity interface{}

    // The protocol with which to send this trap. The type is TrapDestProtocol.
    TrapDestProtocol interface{}

    // The address to send traps on behalf of this entry.  If the associated
    // trapDestProtocol object is equal to ip(1), the encoding of this object is
    // the same as the snmpUDPAddress textual convention in [RFC1906]:   -- for a
    // SnmpUDPAddress of length 6:   --   -- octets   contents        encoding  
    // --  1-4     IP-address      network-byte order   --  5-6     UDP-port      
    // network-byte order  If the associated trapDestProtocol object is equal to
    // ipx(2), the encoding of this object is the same as the snmpIPXAddress
    // textual convention in [RFC1906]:   -- for a SnmpIPXAddress of length 12:  
    // --   -- octets   contents            encoding   --  1-4     network-number 
    // network-byte order   --  5-10    physical-address    network-byte order  
    // -- 11-12    socket-number       network-byte order  This object may not be
    // modified if the associated      trapDestStatus object is equal to
    // active(1). The type is string.
    TrapDestAddress interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    TrapDestOwner interface{}

    // The status of this trap destination entry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value. The
    // type is RowStatus.
    TrapDestStatus interface{}
}

func (trapDestEntry *RMON2MIB_TrapDestTable_TrapDestEntry) GetEntityData() *types.CommonEntityData {
    trapDestEntry.EntityData.YFilter = trapDestEntry.YFilter
    trapDestEntry.EntityData.YangName = "trapDestEntry"
    trapDestEntry.EntityData.BundleName = "cisco_ios_xe"
    trapDestEntry.EntityData.ParentYangName = "trapDestTable"
    trapDestEntry.EntityData.SegmentPath = "trapDestEntry" + types.AddKeyToken(trapDestEntry.TrapDestIndex, "trapDestIndex")
    trapDestEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/trapDestTable/" + trapDestEntry.EntityData.SegmentPath
    trapDestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trapDestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trapDestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trapDestEntry.EntityData.Children = types.NewOrderedMap()
    trapDestEntry.EntityData.Leafs = types.NewOrderedMap()
    trapDestEntry.EntityData.Leafs.Append("trapDestIndex", types.YLeaf{"TrapDestIndex", trapDestEntry.TrapDestIndex})
    trapDestEntry.EntityData.Leafs.Append("trapDestCommunity", types.YLeaf{"TrapDestCommunity", trapDestEntry.TrapDestCommunity})
    trapDestEntry.EntityData.Leafs.Append("trapDestProtocol", types.YLeaf{"TrapDestProtocol", trapDestEntry.TrapDestProtocol})
    trapDestEntry.EntityData.Leafs.Append("trapDestAddress", types.YLeaf{"TrapDestAddress", trapDestEntry.TrapDestAddress})
    trapDestEntry.EntityData.Leafs.Append("trapDestOwner", types.YLeaf{"TrapDestOwner", trapDestEntry.TrapDestOwner})
    trapDestEntry.EntityData.Leafs.Append("trapDestStatus", types.YLeaf{"TrapDestStatus", trapDestEntry.TrapDestStatus})

    trapDestEntry.EntityData.YListKeys = []string {"TrapDestIndex"}

    return &(trapDestEntry.EntityData)
}

// RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol represents The protocol with which to send this trap.
type RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol string

const (
    RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol_ip RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol = "ip"

    RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol_ipx RMON2MIB_TrapDestTable_TrapDestEntry_TrapDestProtocol = "ipx"
)

// RMON2MIB_SerialConnectionTable
// A list of serialConnectionEntries.
type RMON2MIB_SerialConnectionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a SLIP link over a serial line. The type is slice of
    // RMON2MIB_SerialConnectionTable_SerialConnectionEntry.
    SerialConnectionEntry []*RMON2MIB_SerialConnectionTable_SerialConnectionEntry
}

func (serialConnectionTable *RMON2MIB_SerialConnectionTable) GetEntityData() *types.CommonEntityData {
    serialConnectionTable.EntityData.YFilter = serialConnectionTable.YFilter
    serialConnectionTable.EntityData.YangName = "serialConnectionTable"
    serialConnectionTable.EntityData.BundleName = "cisco_ios_xe"
    serialConnectionTable.EntityData.ParentYangName = "RMON2-MIB"
    serialConnectionTable.EntityData.SegmentPath = "serialConnectionTable"
    serialConnectionTable.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/" + serialConnectionTable.EntityData.SegmentPath
    serialConnectionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialConnectionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialConnectionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialConnectionTable.EntityData.Children = types.NewOrderedMap()
    serialConnectionTable.EntityData.Children.Append("serialConnectionEntry", types.YChild{"SerialConnectionEntry", nil})
    for i := range serialConnectionTable.SerialConnectionEntry {
        serialConnectionTable.EntityData.Children.Append(types.GetSegmentPath(serialConnectionTable.SerialConnectionEntry[i]), types.YChild{"SerialConnectionEntry", serialConnectionTable.SerialConnectionEntry[i]})
    }
    serialConnectionTable.EntityData.Leafs = types.NewOrderedMap()

    serialConnectionTable.EntityData.YListKeys = []string {}

    return &(serialConnectionTable.EntityData)
}

// RMON2MIB_SerialConnectionTable_SerialConnectionEntry
// Configuration for a SLIP link over a serial line.
type RMON2MIB_SerialConnectionTable_SerialConnectionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A value that uniquely identifies this
    // serialConnection entry. The type is interface{} with range: 1..65535.
    SerialConnectIndex interface{}

    // The IP Address that can be reached at the other end of this serial
    // connection. This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SerialConnectDestIpAddress interface{}

    // The type of outgoing connection to make.  If this object has the value
    // direct(1), then a direct serial connection is assumed.  If this object has
    // the value modem(2), then serialConnectDialString will be used to make a
    // modem connection.  If this object has the value switch(3),      then
    // serialConnectSwitchConnectSeq will be used to establish the connection over
    // a serial data switch, and serialConnectSwitchDisconnectSeq will be used to
    // terminate the connection.  If this object has the value modem-switch(4),
    // then a modem connection will be made first followed by the switch
    // connection.  This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is
    // SerialConnectType.
    SerialConnectType interface{}

    // A control string which specifies how to dial the phone number in order to
    // establish a modem connection.  The string should include dialing prefix and
    // suffix.  For example: ``^s^MATD9,888-1234^M'' will instruct the Probe to
    // send a carriage return followed by the dialing prefix ``ATD'', the phone
    // number ``9,888-1234'', and a carriage return as the dialing suffix. This
    // object may not be modified if the associated serialConnectStatus object is
    // equal to active(1). The type is string with length: 0..255.
    SerialConnectDialString interface{}

    // A control string which specifies how to establish a data switch connection.
    // This object may not be modified if the associated serialConnectStatus
    // object is equal to active(1). The type is string with length: 0..255.
    SerialConnectSwitchConnectSeq interface{}

    // A control string which specifies how to terminate a data switch connection.
    // This object may not be modified if the associated      serialConnectStatus
    // object is equal to active(1). The type is string with length: 0..255.
    SerialConnectSwitchDisconnectSeq interface{}

    // A control string which specifies how to reset a data switch in the event of
    // a timeout. This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is string with
    // length: 0..255.
    SerialConnectSwitchResetSeq interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    SerialConnectOwner interface{}

    // The status of this serialConnectionEntry.  If the manager attempts to set
    // this object to active(1) when the serialConnectType is set to modem(2) or
    // modem-switch(4) and the serialConnectDialString is a zero-length string or
    // cannot be correctly parsed as a ConnectString, the set request will be
    // rejected with badValue(3).  If the manager attempts to set this object to
    // active(1) when the serialConnectType is set to switch(3) or modem-switch(4)
    // and the serialConnectSwitchConnectSeq, the
    // serialConnectSwitchDisconnectSeq, or the serialConnectSwitchResetSeq are
    // zero-length strings or cannot be correctly parsed as ConnectStrings, the
    // set request will be rejected with badValue(3).  An entry may not exist in
    // the active state unless all objects in the entry have an appropriate value.
    // The type is RowStatus.
    SerialConnectStatus interface{}
}

func (serialConnectionEntry *RMON2MIB_SerialConnectionTable_SerialConnectionEntry) GetEntityData() *types.CommonEntityData {
    serialConnectionEntry.EntityData.YFilter = serialConnectionEntry.YFilter
    serialConnectionEntry.EntityData.YangName = "serialConnectionEntry"
    serialConnectionEntry.EntityData.BundleName = "cisco_ios_xe"
    serialConnectionEntry.EntityData.ParentYangName = "serialConnectionTable"
    serialConnectionEntry.EntityData.SegmentPath = "serialConnectionEntry" + types.AddKeyToken(serialConnectionEntry.SerialConnectIndex, "serialConnectIndex")
    serialConnectionEntry.EntityData.AbsolutePath = "RMON2-MIB:RMON2-MIB/serialConnectionTable/" + serialConnectionEntry.EntityData.SegmentPath
    serialConnectionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialConnectionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialConnectionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialConnectionEntry.EntityData.Children = types.NewOrderedMap()
    serialConnectionEntry.EntityData.Leafs = types.NewOrderedMap()
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectIndex", types.YLeaf{"SerialConnectIndex", serialConnectionEntry.SerialConnectIndex})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectDestIpAddress", types.YLeaf{"SerialConnectDestIpAddress", serialConnectionEntry.SerialConnectDestIpAddress})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectType", types.YLeaf{"SerialConnectType", serialConnectionEntry.SerialConnectType})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectDialString", types.YLeaf{"SerialConnectDialString", serialConnectionEntry.SerialConnectDialString})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectSwitchConnectSeq", types.YLeaf{"SerialConnectSwitchConnectSeq", serialConnectionEntry.SerialConnectSwitchConnectSeq})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectSwitchDisconnectSeq", types.YLeaf{"SerialConnectSwitchDisconnectSeq", serialConnectionEntry.SerialConnectSwitchDisconnectSeq})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectSwitchResetSeq", types.YLeaf{"SerialConnectSwitchResetSeq", serialConnectionEntry.SerialConnectSwitchResetSeq})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectOwner", types.YLeaf{"SerialConnectOwner", serialConnectionEntry.SerialConnectOwner})
    serialConnectionEntry.EntityData.Leafs.Append("serialConnectStatus", types.YLeaf{"SerialConnectStatus", serialConnectionEntry.SerialConnectStatus})

    serialConnectionEntry.EntityData.YListKeys = []string {"SerialConnectIndex"}

    return &(serialConnectionEntry.EntityData)
}

// RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType represents serialConnectStatus object is equal to active(1).
type RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType string

const (
    RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType_direct RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType = "direct"

    RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType_modem RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType = "modem"

    RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType_switch_ RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType = "switch"

    RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType_modemSwitch RMON2MIB_SerialConnectionTable_SerialConnectionEntry_SerialConnectType = "modemSwitch"
)

