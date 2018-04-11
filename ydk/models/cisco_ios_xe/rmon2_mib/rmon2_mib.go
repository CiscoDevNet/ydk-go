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

    
    Protocoldir RMON2MIB_Protocoldir

    
    Addressmap RMON2MIB_Addressmap

    
    Probeconfig RMON2MIB_Probeconfig

    // This table lists the protocols that this agent has the capability to decode
    // and count.  There is one entry in this table for each such protocol.  These
    // protocols represent different network layer, transport layer, and
    // higher-layer protocols.  The agent should boot up with this table
    // preconfigured with those protocols that it knows about and wishes to
    // monitor.  Implementations are strongly encouraged to support protocols
    // higher than the network layer (at least for the protocol distribution
    // group), even for implementations that don't support the application layer
    // groups.
    Protocoldirtable RMON2MIB_Protocoldirtable

    // Controls the setup of protocol type distribution statistics tables. 
    // Implementations are encouraged to add an entry per monitored interface upon
    // initialization so that a default collection of protocol statistics is
    // available.  Rationale: This table controls collection of very basic
    // statistics for any or all of the protocols detected on a given interface.
    // An NMS can use this table to quickly determine bandwidth allocation
    // utilized by different protocols.  A media-specific statistics collection
    // could also be configured (e.g. etherStats, trPStats) to easily obtain total
    // frame, octet, and droppedEvents for the same interface.
    Protocoldistcontroltable RMON2MIB_Protocoldistcontroltable

    // An entry is made in this table for every protocol in the     
    // protocolDirTable which has been seen in at least one packet. Counters are
    // updated in this table for every protocol type that is encountered when
    // parsing a packet, but no counters are updated for packets with MAC-layer
    // errors.  Note that if a protocolDirEntry is deleted, all associated entries
    // in this table are removed.
    Protocoldiststatstable RMON2MIB_Protocoldiststatstable

    // A table to control the collection of network layer address to physical
    // address to interface mappings.  Note that this is not like the typical RMON
    // controlTable and dataTable in which each entry creates its own data table. 
    // Each entry in this table enables the discovery of addresses on a new
    // interface and the placement of address mappings into the central
    // addressMapTable.  Implementations are encouraged to add an entry per
    // monitored interface upon initialization so that a default collection of
    // address mappings is available.
    Addressmapcontroltable RMON2MIB_Addressmapcontroltable

    // A table of network layer address to physical address to interface mappings.
    // The probe will add entries to this table based on the source MAC and
    // network addresses seen in packets without MAC-level errors. The probe will
    // populate this table for all protocols in the protocol directory table whose
    // value of protocolDirAddressMapConfig is equal to supportedOn(3), and will
    // delete any entries whose protocolDirEntry is deleted or has a
    // protocolDirAddressMapConfig value of supportedOff(2).
    Addressmaptable RMON2MIB_Addressmaptable

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
    Hlhostcontroltable RMON2MIB_Hlhostcontroltable

    // A collection of statistics for a particular network layer address that has
    // been discovered on an interface of this device.  The probe will populate
    // this table for all network layer protocols in the protocol directory table
    // whose value of protocolDirHostConfig is equal to supportedOn(3), and will
    // delete any entries whose protocolDirEntry is deleted or has a
    // protocolDirHostConfig value of supportedOff(2).  The probe will add to this
    // table all addresses seen as the source or destination address in all
    // packets with no MAC errors, and will increment octet and packet counts in
    // the table for all packets with no MAC errors.
    Nlhosttable RMON2MIB_Nlhosttable

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
    Hlmatrixcontroltable RMON2MIB_Hlmatrixcontroltable

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
    Nlmatrixsdtable RMON2MIB_Nlmatrixsdtable

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
    Nlmatrixdstable RMON2MIB_Nlmatrixdstable

    // A set of parameters that control the creation of a report of the top N
    // matrix entries according to a selected metric.
    Nlmatrixtopncontroltable RMON2MIB_Nlmatrixtopncontroltable

    // A set of statistics for those network layer matrix entries that have
    // counted the highest number of octets or packets.
    Nlmatrixtopntable RMON2MIB_Nlmatrixtopntable

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
    Alhosttable RMON2MIB_Alhosttable

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
    Almatrixsdtable RMON2MIB_Almatrixsdtable

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
    Almatrixdstable RMON2MIB_Almatrixdstable

    // A set of parameters that control the creation of a report of the top N
    // matrix entries according to a selected metric.
    Almatrixtopncontroltable RMON2MIB_Almatrixtopncontroltable

    // A set of statistics for those application layer matrix entries that have
    // counted the highest number of octets or packets.
    Almatrixtopntable RMON2MIB_Almatrixtopntable

    // A list of data-collection configuration entries.
    Usrhistorycontroltable RMON2MIB_Usrhistorycontroltable

    // A list of data-collection configuration entries.
    Usrhistoryobjecttable RMON2MIB_Usrhistoryobjecttable

    // A list of user defined history entries.
    Usrhistorytable RMON2MIB_Usrhistorytable

    // A table of serial interface configuration entries.  This data will be
    // stored in non-volatile memory and preserved across probe resets or power
    // loss.
    Serialconfigtable RMON2MIB_Serialconfigtable

    // A table of netConfigEntries.
    Netconfigtable RMON2MIB_Netconfigtable

    // A list of trap destination entries.
    Trapdesttable RMON2MIB_Trapdesttable

    // A list of serialConnectionEntries.
    Serialconnectiontable RMON2MIB_Serialconnectiontable
}

func (rMON2MIB *RMON2MIB) GetEntityData() *types.CommonEntityData {
    rMON2MIB.EntityData.YFilter = rMON2MIB.YFilter
    rMON2MIB.EntityData.YangName = "RMON2-MIB"
    rMON2MIB.EntityData.BundleName = "cisco_ios_xe"
    rMON2MIB.EntityData.ParentYangName = "RMON2-MIB"
    rMON2MIB.EntityData.SegmentPath = "RMON2-MIB:RMON2-MIB"
    rMON2MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rMON2MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rMON2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rMON2MIB.EntityData.Children = make(map[string]types.YChild)
    rMON2MIB.EntityData.Children["protocolDir"] = types.YChild{"Protocoldir", &rMON2MIB.Protocoldir}
    rMON2MIB.EntityData.Children["addressMap"] = types.YChild{"Addressmap", &rMON2MIB.Addressmap}
    rMON2MIB.EntityData.Children["probeConfig"] = types.YChild{"Probeconfig", &rMON2MIB.Probeconfig}
    rMON2MIB.EntityData.Children["protocolDirTable"] = types.YChild{"Protocoldirtable", &rMON2MIB.Protocoldirtable}
    rMON2MIB.EntityData.Children["protocolDistControlTable"] = types.YChild{"Protocoldistcontroltable", &rMON2MIB.Protocoldistcontroltable}
    rMON2MIB.EntityData.Children["protocolDistStatsTable"] = types.YChild{"Protocoldiststatstable", &rMON2MIB.Protocoldiststatstable}
    rMON2MIB.EntityData.Children["addressMapControlTable"] = types.YChild{"Addressmapcontroltable", &rMON2MIB.Addressmapcontroltable}
    rMON2MIB.EntityData.Children["addressMapTable"] = types.YChild{"Addressmaptable", &rMON2MIB.Addressmaptable}
    rMON2MIB.EntityData.Children["hlHostControlTable"] = types.YChild{"Hlhostcontroltable", &rMON2MIB.Hlhostcontroltable}
    rMON2MIB.EntityData.Children["nlHostTable"] = types.YChild{"Nlhosttable", &rMON2MIB.Nlhosttable}
    rMON2MIB.EntityData.Children["hlMatrixControlTable"] = types.YChild{"Hlmatrixcontroltable", &rMON2MIB.Hlmatrixcontroltable}
    rMON2MIB.EntityData.Children["nlMatrixSDTable"] = types.YChild{"Nlmatrixsdtable", &rMON2MIB.Nlmatrixsdtable}
    rMON2MIB.EntityData.Children["nlMatrixDSTable"] = types.YChild{"Nlmatrixdstable", &rMON2MIB.Nlmatrixdstable}
    rMON2MIB.EntityData.Children["nlMatrixTopNControlTable"] = types.YChild{"Nlmatrixtopncontroltable", &rMON2MIB.Nlmatrixtopncontroltable}
    rMON2MIB.EntityData.Children["nlMatrixTopNTable"] = types.YChild{"Nlmatrixtopntable", &rMON2MIB.Nlmatrixtopntable}
    rMON2MIB.EntityData.Children["alHostTable"] = types.YChild{"Alhosttable", &rMON2MIB.Alhosttable}
    rMON2MIB.EntityData.Children["alMatrixSDTable"] = types.YChild{"Almatrixsdtable", &rMON2MIB.Almatrixsdtable}
    rMON2MIB.EntityData.Children["alMatrixDSTable"] = types.YChild{"Almatrixdstable", &rMON2MIB.Almatrixdstable}
    rMON2MIB.EntityData.Children["alMatrixTopNControlTable"] = types.YChild{"Almatrixtopncontroltable", &rMON2MIB.Almatrixtopncontroltable}
    rMON2MIB.EntityData.Children["alMatrixTopNTable"] = types.YChild{"Almatrixtopntable", &rMON2MIB.Almatrixtopntable}
    rMON2MIB.EntityData.Children["usrHistoryControlTable"] = types.YChild{"Usrhistorycontroltable", &rMON2MIB.Usrhistorycontroltable}
    rMON2MIB.EntityData.Children["usrHistoryObjectTable"] = types.YChild{"Usrhistoryobjecttable", &rMON2MIB.Usrhistoryobjecttable}
    rMON2MIB.EntityData.Children["usrHistoryTable"] = types.YChild{"Usrhistorytable", &rMON2MIB.Usrhistorytable}
    rMON2MIB.EntityData.Children["serialConfigTable"] = types.YChild{"Serialconfigtable", &rMON2MIB.Serialconfigtable}
    rMON2MIB.EntityData.Children["netConfigTable"] = types.YChild{"Netconfigtable", &rMON2MIB.Netconfigtable}
    rMON2MIB.EntityData.Children["trapDestTable"] = types.YChild{"Trapdesttable", &rMON2MIB.Trapdesttable}
    rMON2MIB.EntityData.Children["serialConnectionTable"] = types.YChild{"Serialconnectiontable", &rMON2MIB.Serialconnectiontable}
    rMON2MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rMON2MIB.EntityData)
}

// RMON2MIB_Protocoldir
type RMON2MIB_Protocoldir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time the protocol directory was last
    // modified, either through insertions or deletions, or through modifications
    // of either the protocolDirAddressMapConfig, protocolDirHostConfig, or
    // protocolDirMatrixConfig. The type is interface{} with range: 0..4294967295.
    Protocoldirlastchange interface{}
}

func (protocoldir *RMON2MIB_Protocoldir) GetEntityData() *types.CommonEntityData {
    protocoldir.EntityData.YFilter = protocoldir.YFilter
    protocoldir.EntityData.YangName = "protocolDir"
    protocoldir.EntityData.BundleName = "cisco_ios_xe"
    protocoldir.EntityData.ParentYangName = "RMON2-MIB"
    protocoldir.EntityData.SegmentPath = "protocolDir"
    protocoldir.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldir.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldir.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldir.EntityData.Children = make(map[string]types.YChild)
    protocoldir.EntityData.Leafs = make(map[string]types.YLeaf)
    protocoldir.EntityData.Leafs["protocolDirLastChange"] = types.YLeaf{"Protocoldirlastchange", protocoldir.Protocoldirlastchange}
    return &(protocoldir.EntityData)
}

// RMON2MIB_Addressmap
type RMON2MIB_Addressmap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an address mapping entry has been inserted into the
    // addressMapTable.  If an entry is inserted, then deleted, and then inserted,
    // this counter will be incremented by 2.  Note that the table size can be
    // determined by subtracting addressMapDeletes from addressMapInserts. The
    // type is interface{} with range: 0..4294967295.
    Addressmapinserts interface{}

    // The number of times an address mapping entry has been deleted from the
    // addressMapTable (for any reason).  If an entry is deleted, then inserted,
    // and then deleted, this counter will be incremented by 2.  Note that the
    // table size can be determined by subtracting addressMapDeletes from
    // addressMapInserts. The type is interface{} with range: 0..4294967295.
    Addressmapdeletes interface{}

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
    Addressmapmaxdesiredentries interface{}
}

func (addressmap *RMON2MIB_Addressmap) GetEntityData() *types.CommonEntityData {
    addressmap.EntityData.YFilter = addressmap.YFilter
    addressmap.EntityData.YangName = "addressMap"
    addressmap.EntityData.BundleName = "cisco_ios_xe"
    addressmap.EntityData.ParentYangName = "RMON2-MIB"
    addressmap.EntityData.SegmentPath = "addressMap"
    addressmap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressmap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressmap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressmap.EntityData.Children = make(map[string]types.YChild)
    addressmap.EntityData.Leafs = make(map[string]types.YLeaf)
    addressmap.EntityData.Leafs["addressMapInserts"] = types.YLeaf{"Addressmapinserts", addressmap.Addressmapinserts}
    addressmap.EntityData.Leafs["addressMapDeletes"] = types.YLeaf{"Addressmapdeletes", addressmap.Addressmapdeletes}
    addressmap.EntityData.Leafs["addressMapMaxDesiredEntries"] = types.YLeaf{"Addressmapmaxdesiredentries", addressmap.Addressmapmaxdesiredentries}
    return &(addressmap.EntityData)
}

// RMON2MIB_Probeconfig
type RMON2MIB_Probeconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of the RMON MIB groups supported on at least one interface by
    // this probe. The type is string with length: 1.
    Probecapabilities interface{}

    // The software revision of this device.  This string will have a zero length
    // if the revision is unknown. The type is string with length: 0..31.
    Probesoftwarerev interface{}

    // The hardware revision of this device.  This string will have a zero length
    // if the revision is unknown. The type is string with length: 0..31.
    Probehardwarerev interface{}

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
    // length: 0 | 8 | 11.
    Probedatetime interface{}

    // Setting this object to warmBoot(2) causes the device to restart the
    // application software with current configuration parameters saved in
    // non-volatile memory.  Setting this object to coldBoot(3) causes the device
    // to reinitialize configuration parameters in non-volatile memory to default
    // values and restart the application software.  When the device is running
    // normally, this variable has a value of running(1). The type is
    // Proberesetcontrol.
    Proberesetcontrol interface{}

    // The file name to be downloaded from the TFTP server when a download is next
    // requested via this MIB.  This value is set to the zero length string when
    // no file name has been specified. The type is string with length: 0..127.
    Probedownloadfile interface{}

    // The IP address of the TFTP server that contains the boot image to load when
    // a download is next requested via this MIB. This value is set to `0.0.0.0'
    // when no IP address has been specified. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Probedownloadtftpserver interface{}

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
    // notDownloading(1). The type is Probedownloadaction.
    Probedownloadaction interface{}

    // The status of the last download procedure, if any.  This object will have a
    // value of downloadStatusUnknown(2) if no download process has been
    // performed. The type is Probedownloadstatus.
    Probedownloadstatus interface{}

    // The IP Address of the default gateway.  If this value is undefined or
    // unknown, it shall have the value 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Netdefaultgateway interface{}
}

func (probeconfig *RMON2MIB_Probeconfig) GetEntityData() *types.CommonEntityData {
    probeconfig.EntityData.YFilter = probeconfig.YFilter
    probeconfig.EntityData.YangName = "probeConfig"
    probeconfig.EntityData.BundleName = "cisco_ios_xe"
    probeconfig.EntityData.ParentYangName = "RMON2-MIB"
    probeconfig.EntityData.SegmentPath = "probeConfig"
    probeconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    probeconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    probeconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    probeconfig.EntityData.Children = make(map[string]types.YChild)
    probeconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    probeconfig.EntityData.Leafs["probeCapabilities"] = types.YLeaf{"Probecapabilities", probeconfig.Probecapabilities}
    probeconfig.EntityData.Leafs["probeSoftwareRev"] = types.YLeaf{"Probesoftwarerev", probeconfig.Probesoftwarerev}
    probeconfig.EntityData.Leafs["probeHardwareRev"] = types.YLeaf{"Probehardwarerev", probeconfig.Probehardwarerev}
    probeconfig.EntityData.Leafs["probeDateTime"] = types.YLeaf{"Probedatetime", probeconfig.Probedatetime}
    probeconfig.EntityData.Leafs["probeResetControl"] = types.YLeaf{"Proberesetcontrol", probeconfig.Proberesetcontrol}
    probeconfig.EntityData.Leafs["probeDownloadFile"] = types.YLeaf{"Probedownloadfile", probeconfig.Probedownloadfile}
    probeconfig.EntityData.Leafs["probeDownloadTFTPServer"] = types.YLeaf{"Probedownloadtftpserver", probeconfig.Probedownloadtftpserver}
    probeconfig.EntityData.Leafs["probeDownloadAction"] = types.YLeaf{"Probedownloadaction", probeconfig.Probedownloadaction}
    probeconfig.EntityData.Leafs["probeDownloadStatus"] = types.YLeaf{"Probedownloadstatus", probeconfig.Probedownloadstatus}
    probeconfig.EntityData.Leafs["netDefaultGateway"] = types.YLeaf{"Netdefaultgateway", probeconfig.Netdefaultgateway}
    return &(probeconfig.EntityData)
}

// RMON2MIB_Probeconfig_Probedownloadaction represents a value of notDownloading(1).
type RMON2MIB_Probeconfig_Probedownloadaction string

const (
    RMON2MIB_Probeconfig_Probedownloadaction_notDownloading RMON2MIB_Probeconfig_Probedownloadaction = "notDownloading"

    RMON2MIB_Probeconfig_Probedownloadaction_downloadToPROM RMON2MIB_Probeconfig_Probedownloadaction = "downloadToPROM"

    RMON2MIB_Probeconfig_Probedownloadaction_downloadToRAM RMON2MIB_Probeconfig_Probedownloadaction = "downloadToRAM"
)

// RMON2MIB_Probeconfig_Probedownloadstatus represents download process has been performed.
type RMON2MIB_Probeconfig_Probedownloadstatus string

const (
    RMON2MIB_Probeconfig_Probedownloadstatus_downloadSuccess RMON2MIB_Probeconfig_Probedownloadstatus = "downloadSuccess"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadStatusUnknown RMON2MIB_Probeconfig_Probedownloadstatus = "downloadStatusUnknown"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadGeneralError RMON2MIB_Probeconfig_Probedownloadstatus = "downloadGeneralError"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadNoResponseFromServer RMON2MIB_Probeconfig_Probedownloadstatus = "downloadNoResponseFromServer"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadChecksumError RMON2MIB_Probeconfig_Probedownloadstatus = "downloadChecksumError"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadIncompatibleImage RMON2MIB_Probeconfig_Probedownloadstatus = "downloadIncompatibleImage"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadTftpFileNotFound RMON2MIB_Probeconfig_Probedownloadstatus = "downloadTftpFileNotFound"

    RMON2MIB_Probeconfig_Probedownloadstatus_downloadTftpAccessViolation RMON2MIB_Probeconfig_Probedownloadstatus = "downloadTftpAccessViolation"
)

// RMON2MIB_Probeconfig_Proberesetcontrol represents running(1).
type RMON2MIB_Probeconfig_Proberesetcontrol string

const (
    RMON2MIB_Probeconfig_Proberesetcontrol_running RMON2MIB_Probeconfig_Proberesetcontrol = "running"

    RMON2MIB_Probeconfig_Proberesetcontrol_warmBoot RMON2MIB_Probeconfig_Proberesetcontrol = "warmBoot"

    RMON2MIB_Probeconfig_Proberesetcontrol_coldBoot RMON2MIB_Probeconfig_Proberesetcontrol = "coldBoot"
)

// RMON2MIB_Protocoldirtable
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
type RMON2MIB_Protocoldirtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDirTable.  An example of the indexing of
    // this entry is protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
    // encoding of a length of 8, followed by 8 subids encoding the protocolDirID
    // of 1.2048, followed by a length of 2 and the 2 subids encoding zero-valued
    // parameters. The type is slice of
    // RMON2MIB_Protocoldirtable_Protocoldirentry.
    Protocoldirentry []RMON2MIB_Protocoldirtable_Protocoldirentry
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetEntityData() *types.CommonEntityData {
    protocoldirtable.EntityData.YFilter = protocoldirtable.YFilter
    protocoldirtable.EntityData.YangName = "protocolDirTable"
    protocoldirtable.EntityData.BundleName = "cisco_ios_xe"
    protocoldirtable.EntityData.ParentYangName = "RMON2-MIB"
    protocoldirtable.EntityData.SegmentPath = "protocolDirTable"
    protocoldirtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldirtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldirtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldirtable.EntityData.Children = make(map[string]types.YChild)
    protocoldirtable.EntityData.Children["protocolDirEntry"] = types.YChild{"Protocoldirentry", nil}
    for i := range protocoldirtable.Protocoldirentry {
        protocoldirtable.EntityData.Children[types.GetSegmentPath(&protocoldirtable.Protocoldirentry[i])] = types.YChild{"Protocoldirentry", &protocoldirtable.Protocoldirentry[i]}
    }
    protocoldirtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocoldirtable.EntityData)
}

// RMON2MIB_Protocoldirtable_Protocoldirentry
// A conceptual row in the protocolDirTable.
// 
// An example of the indexing of this entry is
// protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
// encoding of a length of 8, followed by 8 subids encoding the
// protocolDirID of 1.2048, followed by a length of 2 and the
// 2 subids encoding zero-valued parameters.
type RMON2MIB_Protocoldirtable_Protocoldirentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Protocoldirid interface{}

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
    Protocoldirparameters interface{}

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
    Protocoldirlocalindex interface{}

    // A textual description of the protocol encapsulation. A probe may choose to
    // describe only a subset of the entire encapsulation (e.g. only the highest
    // layer).  This object is intended for human consumption only.  This object
    // may not be modified if the associated protocolDirStatus object is equal to
    // active(1). The type is string with length: 1..64.
    Protocoldirdescr interface{}

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
    // length: 1.
    Protocoldirtype interface{}

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
    // Protocoldiraddressmapconfig.
    Protocoldiraddressmapconfig interface{}

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
    // for both protocols. The type is Protocoldirhostconfig.
    Protocoldirhostconfig interface{}

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
    // Protocoldirmatrixconfig.
    Protocoldirmatrixconfig interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Protocoldirowner interface{}

    // The status of this protocol directory entry.  An entry may not exist in the
    // active state unless all      objects in the entry have an appropriate
    // value.  If this object is not equal to active(1), all associated entries in
    // the nlHostTable, nlMatrixSDTable, nlMatrixDSTable, alHostTable,
    // alMatrixSDTable, and alMatrixDSTable shall be deleted. The type is
    // RowStatus.
    Protocoldirstatus interface{}
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetEntityData() *types.CommonEntityData {
    protocoldirentry.EntityData.YFilter = protocoldirentry.YFilter
    protocoldirentry.EntityData.YangName = "protocolDirEntry"
    protocoldirentry.EntityData.BundleName = "cisco_ios_xe"
    protocoldirentry.EntityData.ParentYangName = "protocolDirTable"
    protocoldirentry.EntityData.SegmentPath = "protocolDirEntry" + "[protocolDirID='" + fmt.Sprintf("%v", protocoldirentry.Protocoldirid) + "']" + "[protocolDirParameters='" + fmt.Sprintf("%v", protocoldirentry.Protocoldirparameters) + "']"
    protocoldirentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldirentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldirentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldirentry.EntityData.Children = make(map[string]types.YChild)
    protocoldirentry.EntityData.Leafs = make(map[string]types.YLeaf)
    protocoldirentry.EntityData.Leafs["protocolDirID"] = types.YLeaf{"Protocoldirid", protocoldirentry.Protocoldirid}
    protocoldirentry.EntityData.Leafs["protocolDirParameters"] = types.YLeaf{"Protocoldirparameters", protocoldirentry.Protocoldirparameters}
    protocoldirentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", protocoldirentry.Protocoldirlocalindex}
    protocoldirentry.EntityData.Leafs["protocolDirDescr"] = types.YLeaf{"Protocoldirdescr", protocoldirentry.Protocoldirdescr}
    protocoldirentry.EntityData.Leafs["protocolDirType"] = types.YLeaf{"Protocoldirtype", protocoldirentry.Protocoldirtype}
    protocoldirentry.EntityData.Leafs["protocolDirAddressMapConfig"] = types.YLeaf{"Protocoldiraddressmapconfig", protocoldirentry.Protocoldiraddressmapconfig}
    protocoldirentry.EntityData.Leafs["protocolDirHostConfig"] = types.YLeaf{"Protocoldirhostconfig", protocoldirentry.Protocoldirhostconfig}
    protocoldirentry.EntityData.Leafs["protocolDirMatrixConfig"] = types.YLeaf{"Protocoldirmatrixconfig", protocoldirentry.Protocoldirmatrixconfig}
    protocoldirentry.EntityData.Leafs["protocolDirOwner"] = types.YLeaf{"Protocoldirowner", protocoldirentry.Protocoldirowner}
    protocoldirentry.EntityData.Leafs["protocolDirStatus"] = types.YLeaf{"Protocoldirstatus", protocoldirentry.Protocoldirstatus}
    return &(protocoldirentry.EntityData)
}

// RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig represents the addressMappingTable.
type RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig string

const (
    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig_notSupported RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig = "notSupported"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig_supportedOff RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig = "supportedOff"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig_supportedOn RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldiraddressmapconfig = "supportedOn"
)

// RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig represents for both protocols.
type RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig string

const (
    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig_notSupported RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig = "notSupported"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig_supportedOff RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig = "supportedOff"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig_supportedOn RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirhostconfig = "supportedOn"
)

// RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig represents for both protocols.
type RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig string

const (
    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig_notSupported RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig = "notSupported"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig_supportedOff RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig = "supportedOff"

    RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig_supportedOn RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirmatrixconfig = "supportedOn"
)

// RMON2MIB_Protocoldistcontroltable
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
type RMON2MIB_Protocoldistcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDistControlTable.  An example of the
    // indexing of this entry is      protocolDistControlDroppedFrames.7. The type
    // is slice of RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry.
    Protocoldistcontrolentry []RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetEntityData() *types.CommonEntityData {
    protocoldistcontroltable.EntityData.YFilter = protocoldistcontroltable.YFilter
    protocoldistcontroltable.EntityData.YangName = "protocolDistControlTable"
    protocoldistcontroltable.EntityData.BundleName = "cisco_ios_xe"
    protocoldistcontroltable.EntityData.ParentYangName = "RMON2-MIB"
    protocoldistcontroltable.EntityData.SegmentPath = "protocolDistControlTable"
    protocoldistcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldistcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldistcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldistcontroltable.EntityData.Children = make(map[string]types.YChild)
    protocoldistcontroltable.EntityData.Children["protocolDistControlEntry"] = types.YChild{"Protocoldistcontrolentry", nil}
    for i := range protocoldistcontroltable.Protocoldistcontrolentry {
        protocoldistcontroltable.EntityData.Children[types.GetSegmentPath(&protocoldistcontroltable.Protocoldistcontrolentry[i])] = types.YChild{"Protocoldistcontrolentry", &protocoldistcontroltable.Protocoldistcontrolentry[i]}
    }
    protocoldistcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocoldistcontroltable.EntityData)
}

// RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry
// A conceptual row in the protocolDistControlTable.
// 
// An example of the indexing of this entry is
// 
// 
// 
// 
// 
// protocolDistControlDroppedFrames.7
type RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique index for this protocolDistControlEntry.
    // The type is interface{} with range: 1..65535.
    Protocoldistcontrolindex interface{}

    // The source of data for the this protocol distribution.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // protocolDistControlStatus object is equal to active(1). The type is string
    // with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Protocoldistcontroldatasource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.       This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Protocoldistcontroldroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Protocoldistcontrolcreatetime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Protocoldistcontrolowner interface{}

    // The status of this row.  An entry may not exist in the active state unless
    // all objects in the entry have an appropriate value.  If this object is not
    // equal to active(1), all associated entries in the protocolDistStatsTable
    // shall be deleted. The type is RowStatus.
    Protocoldistcontrolstatus interface{}
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetEntityData() *types.CommonEntityData {
    protocoldistcontrolentry.EntityData.YFilter = protocoldistcontrolentry.YFilter
    protocoldistcontrolentry.EntityData.YangName = "protocolDistControlEntry"
    protocoldistcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    protocoldistcontrolentry.EntityData.ParentYangName = "protocolDistControlTable"
    protocoldistcontrolentry.EntityData.SegmentPath = "protocolDistControlEntry" + "[protocolDistControlIndex='" + fmt.Sprintf("%v", protocoldistcontrolentry.Protocoldistcontrolindex) + "']"
    protocoldistcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldistcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldistcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldistcontrolentry.EntityData.Children = make(map[string]types.YChild)
    protocoldistcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlIndex"] = types.YLeaf{"Protocoldistcontrolindex", protocoldistcontrolentry.Protocoldistcontrolindex}
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlDataSource"] = types.YLeaf{"Protocoldistcontroldatasource", protocoldistcontrolentry.Protocoldistcontroldatasource}
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlDroppedFrames"] = types.YLeaf{"Protocoldistcontroldroppedframes", protocoldistcontrolentry.Protocoldistcontroldroppedframes}
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlCreateTime"] = types.YLeaf{"Protocoldistcontrolcreatetime", protocoldistcontrolentry.Protocoldistcontrolcreatetime}
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlOwner"] = types.YLeaf{"Protocoldistcontrolowner", protocoldistcontrolentry.Protocoldistcontrolowner}
    protocoldistcontrolentry.EntityData.Leafs["protocolDistControlStatus"] = types.YLeaf{"Protocoldistcontrolstatus", protocoldistcontrolentry.Protocoldistcontrolstatus}
    return &(protocoldistcontrolentry.EntityData)
}

// RMON2MIB_Protocoldiststatstable
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
type RMON2MIB_Protocoldiststatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDistStatsTable.  The index is composed of
    // the protocolDistControlIndex of the associated protocolDistControlEntry
    // followed by the protocolDirLocalIndex of the associated protocol that this
    // entry represents.  In other words, the index identifies the protocol
    // distribution an entry is a part of as well as the particular protocol that
    // it represents.  An example of the indexing of this entry is
    // protocolDistStatsPkts.1.18. The type is slice of
    // RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry.
    Protocoldiststatsentry []RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetEntityData() *types.CommonEntityData {
    protocoldiststatstable.EntityData.YFilter = protocoldiststatstable.YFilter
    protocoldiststatstable.EntityData.YangName = "protocolDistStatsTable"
    protocoldiststatstable.EntityData.BundleName = "cisco_ios_xe"
    protocoldiststatstable.EntityData.ParentYangName = "RMON2-MIB"
    protocoldiststatstable.EntityData.SegmentPath = "protocolDistStatsTable"
    protocoldiststatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldiststatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldiststatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldiststatstable.EntityData.Children = make(map[string]types.YChild)
    protocoldiststatstable.EntityData.Children["protocolDistStatsEntry"] = types.YChild{"Protocoldiststatsentry", nil}
    for i := range protocoldiststatstable.Protocoldiststatsentry {
        protocoldiststatstable.EntityData.Children[types.GetSegmentPath(&protocoldiststatstable.Protocoldiststatsentry[i])] = types.YChild{"Protocoldiststatsentry", &protocoldiststatstable.Protocoldiststatsentry[i]}
    }
    protocoldiststatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocoldiststatstable.EntityData)
}

// RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry
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
type RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry_Protocoldistcontrolindex
    Protocoldistcontrolindex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // The number of packets without errors received of this protocol type.  Note
    // that this is the number of link-layer packets, so if a single network-layer
    // packet is fragmented into several link-layer frames, this counter is
    // incremented several times. The type is interface{} with range:
    // 0..4294967295.
    Protocoldiststatspkts interface{}

    // The number of octets in packets received of this protocol type since it was
    // added to the protocolDistStatsTable (excluding framing bits but including
    // FCS octets), except for those octets in packets that contained errors. 
    // Note this doesn't count just those octets in the particular protocol
    // frames, but includes the entire packet that contained the protocol. The
    // type is interface{} with range: 0..4294967295.
    Protocoldiststatsoctets interface{}
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetEntityData() *types.CommonEntityData {
    protocoldiststatsentry.EntityData.YFilter = protocoldiststatsentry.YFilter
    protocoldiststatsentry.EntityData.YangName = "protocolDistStatsEntry"
    protocoldiststatsentry.EntityData.BundleName = "cisco_ios_xe"
    protocoldiststatsentry.EntityData.ParentYangName = "protocolDistStatsTable"
    protocoldiststatsentry.EntityData.SegmentPath = "protocolDistStatsEntry" + "[protocolDistControlIndex='" + fmt.Sprintf("%v", protocoldiststatsentry.Protocoldistcontrolindex) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", protocoldiststatsentry.Protocoldirlocalindex) + "']"
    protocoldiststatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocoldiststatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocoldiststatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocoldiststatsentry.EntityData.Children = make(map[string]types.YChild)
    protocoldiststatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    protocoldiststatsentry.EntityData.Leafs["protocolDistControlIndex"] = types.YLeaf{"Protocoldistcontrolindex", protocoldiststatsentry.Protocoldistcontrolindex}
    protocoldiststatsentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", protocoldiststatsentry.Protocoldirlocalindex}
    protocoldiststatsentry.EntityData.Leafs["protocolDistStatsPkts"] = types.YLeaf{"Protocoldiststatspkts", protocoldiststatsentry.Protocoldiststatspkts}
    protocoldiststatsentry.EntityData.Leafs["protocolDistStatsOctets"] = types.YLeaf{"Protocoldiststatsoctets", protocoldiststatsentry.Protocoldiststatsoctets}
    return &(protocoldiststatsentry.EntityData)
}

// RMON2MIB_Addressmapcontroltable
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
type RMON2MIB_Addressmapcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapControlTable.      An example of the
    // indexing of this entry is addressMapControlDroppedFrames.1. The type is
    // slice of RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry.
    Addressmapcontrolentry []RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetEntityData() *types.CommonEntityData {
    addressmapcontroltable.EntityData.YFilter = addressmapcontroltable.YFilter
    addressmapcontroltable.EntityData.YangName = "addressMapControlTable"
    addressmapcontroltable.EntityData.BundleName = "cisco_ios_xe"
    addressmapcontroltable.EntityData.ParentYangName = "RMON2-MIB"
    addressmapcontroltable.EntityData.SegmentPath = "addressMapControlTable"
    addressmapcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressmapcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressmapcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressmapcontroltable.EntityData.Children = make(map[string]types.YChild)
    addressmapcontroltable.EntityData.Children["addressMapControlEntry"] = types.YChild{"Addressmapcontrolentry", nil}
    for i := range addressmapcontroltable.Addressmapcontrolentry {
        addressmapcontroltable.EntityData.Children[types.GetSegmentPath(&addressmapcontroltable.Addressmapcontrolentry[i])] = types.YChild{"Addressmapcontrolentry", &addressmapcontroltable.Addressmapcontrolentry[i]}
    }
    addressmapcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressmapcontroltable.EntityData)
}

// RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry
// A conceptual row in the addressMapControlTable.
// 
// 
// 
// 
// 
// An example of the indexing of this entry is
// addressMapControlDroppedFrames.1
type RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique index for this entry in the
    // addressMapControlTable. The type is interface{} with range: 1..65535.
    Addressmapcontrolindex interface{}

    // The source of data for this addressMapControlEntry. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Addressmapcontroldatasource interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Addressmapcontroldroppedframes interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Addressmapcontrolowner interface{}

    // The status of this addressMap control entry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // addressMapTable shall be deleted. The type is RowStatus.
    Addressmapcontrolstatus interface{}
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetEntityData() *types.CommonEntityData {
    addressmapcontrolentry.EntityData.YFilter = addressmapcontrolentry.YFilter
    addressmapcontrolentry.EntityData.YangName = "addressMapControlEntry"
    addressmapcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    addressmapcontrolentry.EntityData.ParentYangName = "addressMapControlTable"
    addressmapcontrolentry.EntityData.SegmentPath = "addressMapControlEntry" + "[addressMapControlIndex='" + fmt.Sprintf("%v", addressmapcontrolentry.Addressmapcontrolindex) + "']"
    addressmapcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressmapcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressmapcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressmapcontrolentry.EntityData.Children = make(map[string]types.YChild)
    addressmapcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    addressmapcontrolentry.EntityData.Leafs["addressMapControlIndex"] = types.YLeaf{"Addressmapcontrolindex", addressmapcontrolentry.Addressmapcontrolindex}
    addressmapcontrolentry.EntityData.Leafs["addressMapControlDataSource"] = types.YLeaf{"Addressmapcontroldatasource", addressmapcontrolentry.Addressmapcontroldatasource}
    addressmapcontrolentry.EntityData.Leafs["addressMapControlDroppedFrames"] = types.YLeaf{"Addressmapcontroldroppedframes", addressmapcontrolentry.Addressmapcontroldroppedframes}
    addressmapcontrolentry.EntityData.Leafs["addressMapControlOwner"] = types.YLeaf{"Addressmapcontrolowner", addressmapcontrolentry.Addressmapcontrolowner}
    addressmapcontrolentry.EntityData.Leafs["addressMapControlStatus"] = types.YLeaf{"Addressmapcontrolstatus", addressmapcontrolentry.Addressmapcontrolstatus}
    return &(addressmapcontrolentry.EntityData)
}

// RMON2MIB_Addressmaptable
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
type RMON2MIB_Addressmaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapTable. The protocolDirLocalIndex in the
    // index identifies the network layer protocol of the
    // addressMapNetworkAddress.      An example of the indexing of this entry is
    // addressMapSource.783495.18.4.128.2.6.6.11.1.3.6.1.2.1.2.2.1.1.1. The type
    // is slice of RMON2MIB_Addressmaptable_Addressmapentry.
    Addressmapentry []RMON2MIB_Addressmaptable_Addressmapentry
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetEntityData() *types.CommonEntityData {
    addressmaptable.EntityData.YFilter = addressmaptable.YFilter
    addressmaptable.EntityData.YangName = "addressMapTable"
    addressmaptable.EntityData.BundleName = "cisco_ios_xe"
    addressmaptable.EntityData.ParentYangName = "RMON2-MIB"
    addressmaptable.EntityData.SegmentPath = "addressMapTable"
    addressmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressmaptable.EntityData.Children = make(map[string]types.YChild)
    addressmaptable.EntityData.Children["addressMapEntry"] = types.YChild{"Addressmapentry", nil}
    for i := range addressmaptable.Addressmapentry {
        addressmaptable.EntityData.Children[types.GetSegmentPath(&addressmaptable.Addressmapentry[i])] = types.YChild{"Addressmapentry", &addressmaptable.Addressmapentry[i]}
    }
    addressmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressmaptable.EntityData)
}

// RMON2MIB_Addressmaptable_Addressmapentry
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
type RMON2MIB_Addressmaptable_Addressmapentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Addressmaptimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The network address for this relation.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the protocolDirLocalIndex component of the index.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    Addressmapnetworkaddress interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Addressmapsource interface{}

    // The last source physical address on which the associated network address
    // was seen.  If the protocol of the associated network address was
    // encapsulated inside of a network-level or higher protocol, this will be the
    // address of the next-lower protocol with the addressRecognitionCapable bit
    // enabled and will be formatted as specified for that protocol. The type is
    // string.
    Addressmapphysicaladdress interface{}

    // The value of sysUpTime at the time this entry was last created or the
    // values of the physical address changed.  This can be used to help detect
    // duplicate address problems, in which case this object will be updated
    // frequently. The type is interface{} with range: 0..4294967295.
    Addressmaplastchange interface{}
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetEntityData() *types.CommonEntityData {
    addressmapentry.EntityData.YFilter = addressmapentry.YFilter
    addressmapentry.EntityData.YangName = "addressMapEntry"
    addressmapentry.EntityData.BundleName = "cisco_ios_xe"
    addressmapentry.EntityData.ParentYangName = "addressMapTable"
    addressmapentry.EntityData.SegmentPath = "addressMapEntry" + "[addressMapTimeMark='" + fmt.Sprintf("%v", addressmapentry.Addressmaptimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", addressmapentry.Protocoldirlocalindex) + "']" + "[addressMapNetworkAddress='" + fmt.Sprintf("%v", addressmapentry.Addressmapnetworkaddress) + "']" + "[addressMapSource='" + fmt.Sprintf("%v", addressmapentry.Addressmapsource) + "']"
    addressmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressmapentry.EntityData.Children = make(map[string]types.YChild)
    addressmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    addressmapentry.EntityData.Leafs["addressMapTimeMark"] = types.YLeaf{"Addressmaptimemark", addressmapentry.Addressmaptimemark}
    addressmapentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", addressmapentry.Protocoldirlocalindex}
    addressmapentry.EntityData.Leafs["addressMapNetworkAddress"] = types.YLeaf{"Addressmapnetworkaddress", addressmapentry.Addressmapnetworkaddress}
    addressmapentry.EntityData.Leafs["addressMapSource"] = types.YLeaf{"Addressmapsource", addressmapentry.Addressmapsource}
    addressmapentry.EntityData.Leafs["addressMapPhysicalAddress"] = types.YLeaf{"Addressmapphysicaladdress", addressmapentry.Addressmapphysicaladdress}
    addressmapentry.EntityData.Leafs["addressMapLastChange"] = types.YLeaf{"Addressmaplastchange", addressmapentry.Addressmaplastchange}
    return &(addressmapentry.EntityData)
}

// RMON2MIB_Hlhostcontroltable
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
type RMON2MIB_Hlhostcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the hlHostControlTable.  An example of the indexing of
    // this entry is hlHostControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry.
    Hlhostcontrolentry []RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetEntityData() *types.CommonEntityData {
    hlhostcontroltable.EntityData.YFilter = hlhostcontroltable.YFilter
    hlhostcontroltable.EntityData.YangName = "hlHostControlTable"
    hlhostcontroltable.EntityData.BundleName = "cisco_ios_xe"
    hlhostcontroltable.EntityData.ParentYangName = "RMON2-MIB"
    hlhostcontroltable.EntityData.SegmentPath = "hlHostControlTable"
    hlhostcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlhostcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlhostcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlhostcontroltable.EntityData.Children = make(map[string]types.YChild)
    hlhostcontroltable.EntityData.Children["hlHostControlEntry"] = types.YChild{"Hlhostcontrolentry", nil}
    for i := range hlhostcontroltable.Hlhostcontrolentry {
        hlhostcontroltable.EntityData.Children[types.GetSegmentPath(&hlhostcontroltable.Hlhostcontrolentry[i])] = types.YChild{"Hlhostcontrolentry", &hlhostcontroltable.Hlhostcontrolentry[i]}
    }
    hlhostcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hlhostcontroltable.EntityData)
}

// RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry
// A conceptual row in the hlHostControlTable.
// 
// An example of the indexing of this entry is
// hlHostControlNlDroppedFrames.1
type RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hlHostControlTable.  Each such entry defines a function that discovers
    // hosts on a particular interface and places statistics about them in the
    // nlHostTable, and optionally in the alHostTable, on behalf of this
    // hlHostControlEntry. The type is interface{} with range: 1..65535.
    Hlhostcontrolindex interface{}

    // The source of data for the associated host tables.  The statistics in this
    // group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // hlHostControlStatus object is equal to active(1). The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Hlhostcontroldatasource interface{}

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
    Hlhostcontrolnldroppedframes interface{}

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
    Hlhostcontrolnlinserts interface{}

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
    Hlhostcontrolnldeletes interface{}

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
    Hlhostcontrolnlmaxdesiredentries interface{}

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
    Hlhostcontrolaldroppedframes interface{}

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
    Hlhostcontrolalinserts interface{}

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
    Hlhostcontrolaldeletes interface{}

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
    Hlhostcontrolalmaxdesiredentries interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Hlhostcontrolowner interface{}

    // The status of this hlHostControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // nlHostTable and alHostTable shall be deleted. The type is RowStatus.
    Hlhostcontrolstatus interface{}
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetEntityData() *types.CommonEntityData {
    hlhostcontrolentry.EntityData.YFilter = hlhostcontrolentry.YFilter
    hlhostcontrolentry.EntityData.YangName = "hlHostControlEntry"
    hlhostcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    hlhostcontrolentry.EntityData.ParentYangName = "hlHostControlTable"
    hlhostcontrolentry.EntityData.SegmentPath = "hlHostControlEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", hlhostcontrolentry.Hlhostcontrolindex) + "']"
    hlhostcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlhostcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlhostcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlhostcontrolentry.EntityData.Children = make(map[string]types.YChild)
    hlhostcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hlhostcontrolentry.EntityData.Leafs["hlHostControlIndex"] = types.YLeaf{"Hlhostcontrolindex", hlhostcontrolentry.Hlhostcontrolindex}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlDataSource"] = types.YLeaf{"Hlhostcontroldatasource", hlhostcontrolentry.Hlhostcontroldatasource}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlNlDroppedFrames"] = types.YLeaf{"Hlhostcontrolnldroppedframes", hlhostcontrolentry.Hlhostcontrolnldroppedframes}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlNlInserts"] = types.YLeaf{"Hlhostcontrolnlinserts", hlhostcontrolentry.Hlhostcontrolnlinserts}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlNlDeletes"] = types.YLeaf{"Hlhostcontrolnldeletes", hlhostcontrolentry.Hlhostcontrolnldeletes}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlNlMaxDesiredEntries"] = types.YLeaf{"Hlhostcontrolnlmaxdesiredentries", hlhostcontrolentry.Hlhostcontrolnlmaxdesiredentries}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlAlDroppedFrames"] = types.YLeaf{"Hlhostcontrolaldroppedframes", hlhostcontrolentry.Hlhostcontrolaldroppedframes}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlAlInserts"] = types.YLeaf{"Hlhostcontrolalinserts", hlhostcontrolentry.Hlhostcontrolalinserts}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlAlDeletes"] = types.YLeaf{"Hlhostcontrolaldeletes", hlhostcontrolentry.Hlhostcontrolaldeletes}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlAlMaxDesiredEntries"] = types.YLeaf{"Hlhostcontrolalmaxdesiredentries", hlhostcontrolentry.Hlhostcontrolalmaxdesiredentries}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlOwner"] = types.YLeaf{"Hlhostcontrolowner", hlhostcontrolentry.Hlhostcontrolowner}
    hlhostcontrolentry.EntityData.Leafs["hlHostControlStatus"] = types.YLeaf{"Hlhostcontrolstatus", hlhostcontrolentry.Hlhostcontrolstatus}
    return &(hlhostcontrolentry.EntityData)
}

// RMON2MIB_Nlhosttable
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
type RMON2MIB_Nlhosttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlHostTable.  The hlHostControlIndex value in the
    // index identifies the hlHostControlEntry on whose behalf this entry was
    // created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlHostAddress.  An example of the indexing of
    // this entry is nlHostOutPkts.1.783495.18.4.128.2.6.6. The type is slice of
    // RMON2MIB_Nlhosttable_Nlhostentry.
    Nlhostentry []RMON2MIB_Nlhosttable_Nlhostentry
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetEntityData() *types.CommonEntityData {
    nlhosttable.EntityData.YFilter = nlhosttable.YFilter
    nlhosttable.EntityData.YangName = "nlHostTable"
    nlhosttable.EntityData.BundleName = "cisco_ios_xe"
    nlhosttable.EntityData.ParentYangName = "RMON2-MIB"
    nlhosttable.EntityData.SegmentPath = "nlHostTable"
    nlhosttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlhosttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlhosttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlhosttable.EntityData.Children = make(map[string]types.YChild)
    nlhosttable.EntityData.Children["nlHostEntry"] = types.YChild{"Nlhostentry", nil}
    for i := range nlhosttable.Nlhostentry {
        nlhosttable.EntityData.Children[types.GetSegmentPath(&nlhosttable.Nlhostentry[i])] = types.YChild{"Nlhostentry", &nlhosttable.Nlhostentry[i]}
    }
    nlhosttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlhosttable.EntityData)
}

// RMON2MIB_Nlhosttable_Nlhostentry
// A conceptual row in the nlHostTable.
// 
// The hlHostControlIndex value in the index identifies the
// hlHostControlEntry on whose behalf this entry was created.
// The protocolDirLocalIndex value in the index identifies the
// network layer protocol of the nlHostAddress.
// 
// An example of the indexing of this entry is
// nlHostOutPkts.1.783495.18.4.128.2.6.6.
type RMON2MIB_Nlhosttable_Nlhostentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry_Hlhostcontrolindex
    Hlhostcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Nlhosttimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The network address for this nlHostEntry.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the protocolDirLocalIndex component of the index.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    Nlhostaddress interface{}

    // The number of packets without errors transmitted to this address since it
    // was added to the nlHostTable.  Note that this is the number of link-layer
    // packets, so if a single network-layer packet is fragmented into several
    // link-layer frames, this counter is incremented several times. The type is
    // interface{} with range: 0..4294967295.
    Nlhostinpkts interface{}

    // The number of packets without errors transmitted by      this address since
    // it was added to the nlHostTable.  Note that this is the number of
    // link-layer packets, so if a single network-layer packet is fragmented into
    // several link-layer frames, this counter is incremented several times. The
    // type is interface{} with range: 0..4294967295.
    Nlhostoutpkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // nlHostTable (excluding framing bits but including FCS octets), excluding
    // those octets in packets that contained errors.  Note this doesn't count
    // just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    Nlhostinoctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // nlHostTable (excluding framing bits but including FCS octets), excluding
    // those octets in packets that contained errors.  Note this doesn't count
    // just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    Nlhostoutoctets interface{}

    // The number of packets without errors transmitted by this address that were
    // directed to any MAC broadcast addresses or to any MAC multicast addresses
    // since this host was added to the nlHostTable. Note that this is the number
    // of link-layer packets, so if a single network-layer packet is fragmented
    // into several link-layer frames, this counter is incremented several times.
    // The type is interface{} with range: 0..4294967295.
    Nlhostoutmacnonunicastpkts interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Nlhostcreatetime interface{}
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetEntityData() *types.CommonEntityData {
    nlhostentry.EntityData.YFilter = nlhostentry.YFilter
    nlhostentry.EntityData.YangName = "nlHostEntry"
    nlhostentry.EntityData.BundleName = "cisco_ios_xe"
    nlhostentry.EntityData.ParentYangName = "nlHostTable"
    nlhostentry.EntityData.SegmentPath = "nlHostEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", nlhostentry.Hlhostcontrolindex) + "']" + "[nlHostTimeMark='" + fmt.Sprintf("%v", nlhostentry.Nlhosttimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlhostentry.Protocoldirlocalindex) + "']" + "[nlHostAddress='" + fmt.Sprintf("%v", nlhostentry.Nlhostaddress) + "']"
    nlhostentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlhostentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlhostentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlhostentry.EntityData.Children = make(map[string]types.YChild)
    nlhostentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlhostentry.EntityData.Leafs["hlHostControlIndex"] = types.YLeaf{"Hlhostcontrolindex", nlhostentry.Hlhostcontrolindex}
    nlhostentry.EntityData.Leafs["nlHostTimeMark"] = types.YLeaf{"Nlhosttimemark", nlhostentry.Nlhosttimemark}
    nlhostentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", nlhostentry.Protocoldirlocalindex}
    nlhostentry.EntityData.Leafs["nlHostAddress"] = types.YLeaf{"Nlhostaddress", nlhostentry.Nlhostaddress}
    nlhostentry.EntityData.Leafs["nlHostInPkts"] = types.YLeaf{"Nlhostinpkts", nlhostentry.Nlhostinpkts}
    nlhostentry.EntityData.Leafs["nlHostOutPkts"] = types.YLeaf{"Nlhostoutpkts", nlhostentry.Nlhostoutpkts}
    nlhostentry.EntityData.Leafs["nlHostInOctets"] = types.YLeaf{"Nlhostinoctets", nlhostentry.Nlhostinoctets}
    nlhostentry.EntityData.Leafs["nlHostOutOctets"] = types.YLeaf{"Nlhostoutoctets", nlhostentry.Nlhostoutoctets}
    nlhostentry.EntityData.Leafs["nlHostOutMacNonUnicastPkts"] = types.YLeaf{"Nlhostoutmacnonunicastpkts", nlhostentry.Nlhostoutmacnonunicastpkts}
    nlhostentry.EntityData.Leafs["nlHostCreateTime"] = types.YLeaf{"Nlhostcreatetime", nlhostentry.Nlhostcreatetime}
    return &(nlhostentry.EntityData)
}

// RMON2MIB_Hlmatrixcontroltable
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
type RMON2MIB_Hlmatrixcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the hlMatrixControlTable.  An example of indexing of
    // this entry is hlMatrixControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry.
    Hlmatrixcontrolentry []RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetEntityData() *types.CommonEntityData {
    hlmatrixcontroltable.EntityData.YFilter = hlmatrixcontroltable.YFilter
    hlmatrixcontroltable.EntityData.YangName = "hlMatrixControlTable"
    hlmatrixcontroltable.EntityData.BundleName = "cisco_ios_xe"
    hlmatrixcontroltable.EntityData.ParentYangName = "RMON2-MIB"
    hlmatrixcontroltable.EntityData.SegmentPath = "hlMatrixControlTable"
    hlmatrixcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlmatrixcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlmatrixcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlmatrixcontroltable.EntityData.Children = make(map[string]types.YChild)
    hlmatrixcontroltable.EntityData.Children["hlMatrixControlEntry"] = types.YChild{"Hlmatrixcontrolentry", nil}
    for i := range hlmatrixcontroltable.Hlmatrixcontrolentry {
        hlmatrixcontroltable.EntityData.Children[types.GetSegmentPath(&hlmatrixcontroltable.Hlmatrixcontrolentry[i])] = types.YChild{"Hlmatrixcontrolentry", &hlmatrixcontroltable.Hlmatrixcontrolentry[i]}
    }
    hlmatrixcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hlmatrixcontroltable.EntityData)
}

// RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry
// A conceptual row in the hlMatrixControlTable.
// 
// An example of indexing of this entry is
// hlMatrixControlNlDroppedFrames.1
type RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hlMatrixControlTable.  Each such entry defines a function that discovers
    // conversations on a particular interface and places statistics about them in
    // the nlMatrixSDTable and the nlMatrixDSTable, and optionally the
    // alMatrixSDTable and alMatrixDSTable, on behalf of this
    // hlMatrixControlEntry. The type is interface{} with range: 1..65535.
    Hlmatrixcontrolindex interface{}

    // The source of the data for the associated matrix tables.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // hlMatrixControlStatus object is equal to active(1). The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Hlmatrixcontroldatasource interface{}

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
    Hlmatrixcontrolnldroppedframes interface{}

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
    Hlmatrixcontrolnlinserts interface{}

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
    Hlmatrixcontrolnldeletes interface{}

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
    Hlmatrixcontrolnlmaxdesiredentries interface{}

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
    Hlmatrixcontrolaldroppedframes interface{}

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
    Hlmatrixcontrolalinserts interface{}

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
    Hlmatrixcontrolaldeletes interface{}

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
    Hlmatrixcontrolalmaxdesiredentries interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Hlmatrixcontrolowner interface{}

    // The status of this hlMatrixControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // nlMatrixSDTable, nlMatrixDSTable, alMatrixSDTable, and the alMatrixDSTable
    // shall be deleted by the agent. The type is RowStatus.
    Hlmatrixcontrolstatus interface{}
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetEntityData() *types.CommonEntityData {
    hlmatrixcontrolentry.EntityData.YFilter = hlmatrixcontrolentry.YFilter
    hlmatrixcontrolentry.EntityData.YangName = "hlMatrixControlEntry"
    hlmatrixcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    hlmatrixcontrolentry.EntityData.ParentYangName = "hlMatrixControlTable"
    hlmatrixcontrolentry.EntityData.SegmentPath = "hlMatrixControlEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", hlmatrixcontrolentry.Hlmatrixcontrolindex) + "']"
    hlmatrixcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hlmatrixcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hlmatrixcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hlmatrixcontrolentry.EntityData.Children = make(map[string]types.YChild)
    hlmatrixcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlIndex"] = types.YLeaf{"Hlmatrixcontrolindex", hlmatrixcontrolentry.Hlmatrixcontrolindex}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlDataSource"] = types.YLeaf{"Hlmatrixcontroldatasource", hlmatrixcontrolentry.Hlmatrixcontroldatasource}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlNlDroppedFrames"] = types.YLeaf{"Hlmatrixcontrolnldroppedframes", hlmatrixcontrolentry.Hlmatrixcontrolnldroppedframes}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlNlInserts"] = types.YLeaf{"Hlmatrixcontrolnlinserts", hlmatrixcontrolentry.Hlmatrixcontrolnlinserts}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlNlDeletes"] = types.YLeaf{"Hlmatrixcontrolnldeletes", hlmatrixcontrolentry.Hlmatrixcontrolnldeletes}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlNlMaxDesiredEntries"] = types.YLeaf{"Hlmatrixcontrolnlmaxdesiredentries", hlmatrixcontrolentry.Hlmatrixcontrolnlmaxdesiredentries}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlAlDroppedFrames"] = types.YLeaf{"Hlmatrixcontrolaldroppedframes", hlmatrixcontrolentry.Hlmatrixcontrolaldroppedframes}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlAlInserts"] = types.YLeaf{"Hlmatrixcontrolalinserts", hlmatrixcontrolentry.Hlmatrixcontrolalinserts}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlAlDeletes"] = types.YLeaf{"Hlmatrixcontrolaldeletes", hlmatrixcontrolentry.Hlmatrixcontrolaldeletes}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlAlMaxDesiredEntries"] = types.YLeaf{"Hlmatrixcontrolalmaxdesiredentries", hlmatrixcontrolentry.Hlmatrixcontrolalmaxdesiredentries}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlOwner"] = types.YLeaf{"Hlmatrixcontrolowner", hlmatrixcontrolentry.Hlmatrixcontrolowner}
    hlmatrixcontrolentry.EntityData.Leafs["hlMatrixControlStatus"] = types.YLeaf{"Hlmatrixcontrolstatus", hlmatrixcontrolentry.Hlmatrixcontrolstatus}
    return &(hlmatrixcontrolentry.EntityData)
}

// RMON2MIB_Nlmatrixsdtable
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
type RMON2MIB_Nlmatrixsdtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixSDTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlMatrixSDSourceAddress and
    // nlMatrixSDDestAddress.  An example of the indexing of this table is
    // nlMatrixSDPkts.1.783495.18.4.128.2.6.6.4.128.2.6.7. The type is slice of
    // RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry.
    Nlmatrixsdentry []RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetEntityData() *types.CommonEntityData {
    nlmatrixsdtable.EntityData.YFilter = nlmatrixsdtable.YFilter
    nlmatrixsdtable.EntityData.YangName = "nlMatrixSDTable"
    nlmatrixsdtable.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixsdtable.EntityData.ParentYangName = "RMON2-MIB"
    nlmatrixsdtable.EntityData.SegmentPath = "nlMatrixSDTable"
    nlmatrixsdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixsdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixsdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixsdtable.EntityData.Children = make(map[string]types.YChild)
    nlmatrixsdtable.EntityData.Children["nlMatrixSDEntry"] = types.YChild{"Nlmatrixsdentry", nil}
    for i := range nlmatrixsdtable.Nlmatrixsdentry {
        nlmatrixsdtable.EntityData.Children[types.GetSegmentPath(&nlmatrixsdtable.Nlmatrixsdentry[i])] = types.YChild{"Nlmatrixsdentry", &nlmatrixsdtable.Nlmatrixsdentry[i]}
    }
    nlmatrixsdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmatrixsdtable.EntityData)
}

// RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry
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
type RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry_Hlmatrixcontrolindex
    Hlmatrixcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Nlmatrixsdtimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The network source address for this
    // nlMatrixSDEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Nlmatrixsdsourceaddress interface{}

    // This attribute is a key. The network destination address for this
    // nlMatrixSDEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Nlmatrixsddestaddress interface{}

    // The number of packets without errors transmitted from the source address to
    // the destination address since this entry was added to the nlMatrixSDTable. 
    // Note that this is the number of link-layer packets, so if a single
    // network-layer packet is fragmented into several link-layer frames, this
    // counter is incremented several times. The type is interface{} with range:
    // 0..4294967295.
    Nlmatrixsdpkts interface{}

    // The number of octets transmitted from the source address to the destination
    // address since this entry was added to the nlMatrixSDTable (excluding
    // framing bits but including FCS octets), excluding those octets in packets
    // that contained errors.  Note this doesn't count just those octets in the
    // particular protocol frames, but includes the entire packet that contained
    // the protocol. The type is interface{} with range: 0..4294967295.
    Nlmatrixsdoctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Nlmatrixsdcreatetime interface{}
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetEntityData() *types.CommonEntityData {
    nlmatrixsdentry.EntityData.YFilter = nlmatrixsdentry.YFilter
    nlmatrixsdentry.EntityData.YangName = "nlMatrixSDEntry"
    nlmatrixsdentry.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixsdentry.EntityData.ParentYangName = "nlMatrixSDTable"
    nlmatrixsdentry.EntityData.SegmentPath = "nlMatrixSDEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", nlmatrixsdentry.Hlmatrixcontrolindex) + "']" + "[nlMatrixSDTimeMark='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsdtimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlmatrixsdentry.Protocoldirlocalindex) + "']" + "[nlMatrixSDSourceAddress='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsdsourceaddress) + "']" + "[nlMatrixSDDestAddress='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsddestaddress) + "']"
    nlmatrixsdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixsdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixsdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixsdentry.EntityData.Children = make(map[string]types.YChild)
    nlmatrixsdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmatrixsdentry.EntityData.Leafs["hlMatrixControlIndex"] = types.YLeaf{"Hlmatrixcontrolindex", nlmatrixsdentry.Hlmatrixcontrolindex}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDTimeMark"] = types.YLeaf{"Nlmatrixsdtimemark", nlmatrixsdentry.Nlmatrixsdtimemark}
    nlmatrixsdentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", nlmatrixsdentry.Protocoldirlocalindex}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDSourceAddress"] = types.YLeaf{"Nlmatrixsdsourceaddress", nlmatrixsdentry.Nlmatrixsdsourceaddress}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDDestAddress"] = types.YLeaf{"Nlmatrixsddestaddress", nlmatrixsdentry.Nlmatrixsddestaddress}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDPkts"] = types.YLeaf{"Nlmatrixsdpkts", nlmatrixsdentry.Nlmatrixsdpkts}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDOctets"] = types.YLeaf{"Nlmatrixsdoctets", nlmatrixsdentry.Nlmatrixsdoctets}
    nlmatrixsdentry.EntityData.Leafs["nlMatrixSDCreateTime"] = types.YLeaf{"Nlmatrixsdcreatetime", nlmatrixsdentry.Nlmatrixsdcreatetime}
    return &(nlmatrixsdentry.EntityData)
}

// RMON2MIB_Nlmatrixdstable
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
type RMON2MIB_Nlmatrixdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixDSTable.  The hlMatrixControlIndex value in
    // the index identifies the hlMatrixControlEntry on whose behalf this entry
    // was created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlMatrixDSSourceAddress and
    // nlMatrixDSDestAddress.  An example of the indexing of this table is
    // nlMatrixDSPkts.1.783495.18.4.128.2.6.7.4.128.2.6.6. The type is slice of
    // RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry.
    Nlmatrixdsentry []RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetEntityData() *types.CommonEntityData {
    nlmatrixdstable.EntityData.YFilter = nlmatrixdstable.YFilter
    nlmatrixdstable.EntityData.YangName = "nlMatrixDSTable"
    nlmatrixdstable.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixdstable.EntityData.ParentYangName = "RMON2-MIB"
    nlmatrixdstable.EntityData.SegmentPath = "nlMatrixDSTable"
    nlmatrixdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixdstable.EntityData.Children = make(map[string]types.YChild)
    nlmatrixdstable.EntityData.Children["nlMatrixDSEntry"] = types.YChild{"Nlmatrixdsentry", nil}
    for i := range nlmatrixdstable.Nlmatrixdsentry {
        nlmatrixdstable.EntityData.Children[types.GetSegmentPath(&nlmatrixdstable.Nlmatrixdsentry[i])] = types.YChild{"Nlmatrixdsentry", &nlmatrixdstable.Nlmatrixdsentry[i]}
    }
    nlmatrixdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmatrixdstable.EntityData)
}

// RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry
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
type RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry_Hlmatrixcontrolindex
    Hlmatrixcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Nlmatrixdstimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The network destination address for this
    // nlMatrixDSEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Nlmatrixdsdestaddress interface{}

    // This attribute is a key. The network source address for this
    // nlMatrixDSEntry.  This is represented as an octet string with specific
    // semantics and length as identified by the protocolDirLocalIndex component
    // of the index.  For example, if the protocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Nlmatrixdssourceaddress interface{}

    // The number of packets without errors transmitted from the source address to
    // the destination address since this entry was added to the nlMatrixDSTable. 
    // Note that this is the number of link-layer packets, so if a single
    // network-layer packet is fragmented into several link-layer frames, this
    // counter is incremented several times. The type is interface{} with range:
    // 0..4294967295.
    Nlmatrixdspkts interface{}

    // The number of octets transmitted from the source address to the destination
    // address since this entry was added to the nlMatrixDSTable (excluding
    // framing bits but including FCS octets), excluding those octets in packets
    // that contained errors.  Note this doesn't count just those octets in the
    // particular protocol frames, but includes the entire packet that contained
    // the protocol. The type is interface{} with range: 0..4294967295.
    Nlmatrixdsoctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Nlmatrixdscreatetime interface{}
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetEntityData() *types.CommonEntityData {
    nlmatrixdsentry.EntityData.YFilter = nlmatrixdsentry.YFilter
    nlmatrixdsentry.EntityData.YangName = "nlMatrixDSEntry"
    nlmatrixdsentry.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixdsentry.EntityData.ParentYangName = "nlMatrixDSTable"
    nlmatrixdsentry.EntityData.SegmentPath = "nlMatrixDSEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", nlmatrixdsentry.Hlmatrixcontrolindex) + "']" + "[nlMatrixDSTimeMark='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdstimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlmatrixdsentry.Protocoldirlocalindex) + "']" + "[nlMatrixDSDestAddress='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdsdestaddress) + "']" + "[nlMatrixDSSourceAddress='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdssourceaddress) + "']"
    nlmatrixdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixdsentry.EntityData.Children = make(map[string]types.YChild)
    nlmatrixdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmatrixdsentry.EntityData.Leafs["hlMatrixControlIndex"] = types.YLeaf{"Hlmatrixcontrolindex", nlmatrixdsentry.Hlmatrixcontrolindex}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSTimeMark"] = types.YLeaf{"Nlmatrixdstimemark", nlmatrixdsentry.Nlmatrixdstimemark}
    nlmatrixdsentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", nlmatrixdsentry.Protocoldirlocalindex}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSDestAddress"] = types.YLeaf{"Nlmatrixdsdestaddress", nlmatrixdsentry.Nlmatrixdsdestaddress}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSSourceAddress"] = types.YLeaf{"Nlmatrixdssourceaddress", nlmatrixdsentry.Nlmatrixdssourceaddress}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSPkts"] = types.YLeaf{"Nlmatrixdspkts", nlmatrixdsentry.Nlmatrixdspkts}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSOctets"] = types.YLeaf{"Nlmatrixdsoctets", nlmatrixdsentry.Nlmatrixdsoctets}
    nlmatrixdsentry.EntityData.Leafs["nlMatrixDSCreateTime"] = types.YLeaf{"Nlmatrixdscreatetime", nlmatrixdsentry.Nlmatrixdscreatetime}
    return &(nlmatrixdsentry.EntityData)
}

// RMON2MIB_Nlmatrixtopncontroltable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_Nlmatrixtopncontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNControlTable.  An example of the
    // indexing of this table is nlMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry.
    Nlmatrixtopncontrolentry []RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetEntityData() *types.CommonEntityData {
    nlmatrixtopncontroltable.EntityData.YFilter = nlmatrixtopncontroltable.YFilter
    nlmatrixtopncontroltable.EntityData.YangName = "nlMatrixTopNControlTable"
    nlmatrixtopncontroltable.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixtopncontroltable.EntityData.ParentYangName = "RMON2-MIB"
    nlmatrixtopncontroltable.EntityData.SegmentPath = "nlMatrixTopNControlTable"
    nlmatrixtopncontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixtopncontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixtopncontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixtopncontroltable.EntityData.Children = make(map[string]types.YChild)
    nlmatrixtopncontroltable.EntityData.Children["nlMatrixTopNControlEntry"] = types.YChild{"Nlmatrixtopncontrolentry", nil}
    for i := range nlmatrixtopncontroltable.Nlmatrixtopncontrolentry {
        nlmatrixtopncontroltable.EntityData.Children[types.GetSegmentPath(&nlmatrixtopncontroltable.Nlmatrixtopncontrolentry[i])] = types.YChild{"Nlmatrixtopncontrolentry", &nlmatrixtopncontroltable.Nlmatrixtopncontrolentry[i]}
    }
    nlmatrixtopncontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmatrixtopncontroltable.EntityData)
}

// RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry
// A conceptual row in the nlMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// nlMatrixTopNControlDuration.3
type RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // nlMatrixTopNControlTable.  Each such entry defines one top N report
    // prepared for one interface. The type is interface{} with range: 1..65535.
    Nlmatrixtopncontrolindex interface{}

    // The nlMatrix[SD/DS] table for which a top N report will be prepared on
    // behalf of this entry.  The nlMatrix[SD/DS] table is identified by the value
    // of the hlMatrixControlIndex for that table - that value is used here to
    // identify the particular table.  This object may not be modified if the
    // associated nlMatrixTopNControlStatus object is equal to active(1). The type
    // is interface{} with range: 1..65535.
    Nlmatrixtopncontrolmatrixindex interface{}

    // The variable for each nlMatrix[SD/DS] entry that the nlMatrixTopNEntries
    // are sorted by.      This object may not be modified if the associated
    // nlMatrixTopNControlStatus object is equal to active(1). The type is
    // Nlmatrixtopncontrolratebase.
    Nlmatrixtopncontrolratebase interface{}

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
    Nlmatrixtopncontroltimeremaining interface{}

    // The number of reports that have been generated by this entry. The type is
    // interface{} with range: 0..4294967295.
    Nlmatrixtopncontrolgeneratedreports interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval.  When the associated nlMatrixTopNControlTimeRemaining
    // object is set, this object shall be set by the probe to the same value and
    // shall not be modified until the next time the
    // nlMatrixTopNControlTimeRemaining is set. This value shall be zero if no
    // reports have been requested for this nlMatrixTopNControlEntry. The type is
    // interface{} with range: -2147483648..2147483647.
    Nlmatrixtopncontrolduration interface{}

    // The maximum number of matrix entries requested for this report.  When this
    // object is created or modified, the probe should set
    // nlMatrixTopNControlGrantedSize as closely to this object as is possible for
    // the particular probe implementation and available resources. The type is
    // interface{} with range: 0..2147483647.
    Nlmatrixtopncontrolrequestedsize interface{}

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
    Nlmatrixtopncontrolgrantedsize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated
    // nlMatrixTopNControlTimeRemaining object was modified to start the requested
    // report or the time the report was last automatically (re)started.  This
    // object may be used by the management station to determine if a report was
    // missed or not. The type is interface{} with range: 0..4294967295.
    Nlmatrixtopncontrolstarttime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Nlmatrixtopncontrolowner interface{}

    // The status of this nlMatrixTopNControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.    
    // If this object is not equal to active(1), all associated entries in the
    // nlMatrixTopNTable shall be deleted by the agent. The type is RowStatus.
    Nlmatrixtopncontrolstatus interface{}
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetEntityData() *types.CommonEntityData {
    nlmatrixtopncontrolentry.EntityData.YFilter = nlmatrixtopncontrolentry.YFilter
    nlmatrixtopncontrolentry.EntityData.YangName = "nlMatrixTopNControlEntry"
    nlmatrixtopncontrolentry.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixtopncontrolentry.EntityData.ParentYangName = "nlMatrixTopNControlTable"
    nlmatrixtopncontrolentry.EntityData.SegmentPath = "nlMatrixTopNControlEntry" + "[nlMatrixTopNControlIndex='" + fmt.Sprintf("%v", nlmatrixtopncontrolentry.Nlmatrixtopncontrolindex) + "']"
    nlmatrixtopncontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixtopncontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixtopncontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixtopncontrolentry.EntityData.Children = make(map[string]types.YChild)
    nlmatrixtopncontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlIndex"] = types.YLeaf{"Nlmatrixtopncontrolindex", nlmatrixtopncontrolentry.Nlmatrixtopncontrolindex}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlMatrixIndex"] = types.YLeaf{"Nlmatrixtopncontrolmatrixindex", nlmatrixtopncontrolentry.Nlmatrixtopncontrolmatrixindex}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlRateBase"] = types.YLeaf{"Nlmatrixtopncontrolratebase", nlmatrixtopncontrolentry.Nlmatrixtopncontrolratebase}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlTimeRemaining"] = types.YLeaf{"Nlmatrixtopncontroltimeremaining", nlmatrixtopncontrolentry.Nlmatrixtopncontroltimeremaining}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlGeneratedReports"] = types.YLeaf{"Nlmatrixtopncontrolgeneratedreports", nlmatrixtopncontrolentry.Nlmatrixtopncontrolgeneratedreports}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlDuration"] = types.YLeaf{"Nlmatrixtopncontrolduration", nlmatrixtopncontrolentry.Nlmatrixtopncontrolduration}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlRequestedSize"] = types.YLeaf{"Nlmatrixtopncontrolrequestedsize", nlmatrixtopncontrolentry.Nlmatrixtopncontrolrequestedsize}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlGrantedSize"] = types.YLeaf{"Nlmatrixtopncontrolgrantedsize", nlmatrixtopncontrolentry.Nlmatrixtopncontrolgrantedsize}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlStartTime"] = types.YLeaf{"Nlmatrixtopncontrolstarttime", nlmatrixtopncontrolentry.Nlmatrixtopncontrolstarttime}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlOwner"] = types.YLeaf{"Nlmatrixtopncontrolowner", nlmatrixtopncontrolentry.Nlmatrixtopncontrolowner}
    nlmatrixtopncontrolentry.EntityData.Leafs["nlMatrixTopNControlStatus"] = types.YLeaf{"Nlmatrixtopncontrolstatus", nlmatrixtopncontrolentry.Nlmatrixtopncontrolstatus}
    return &(nlmatrixtopncontrolentry.EntityData)
}

// RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase represents nlMatrixTopNControlStatus object is equal to active(1).
type RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase string

const (
    RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase_nlMatrixTopNPkts RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase = "nlMatrixTopNPkts"

    RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase_nlMatrixTopNOctets RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolratebase = "nlMatrixTopNOctets"
)

// RMON2MIB_Nlmatrixtopntable
// A set of statistics for those network layer matrix entries
// that have counted the highest number of octets or packets.
type RMON2MIB_Nlmatrixtopntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNTable.  The nlMatrixTopNControlIndex
    // value in the index identifies the nlMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // nlMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry.
    Nlmatrixtopnentry []RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetEntityData() *types.CommonEntityData {
    nlmatrixtopntable.EntityData.YFilter = nlmatrixtopntable.YFilter
    nlmatrixtopntable.EntityData.YangName = "nlMatrixTopNTable"
    nlmatrixtopntable.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixtopntable.EntityData.ParentYangName = "RMON2-MIB"
    nlmatrixtopntable.EntityData.SegmentPath = "nlMatrixTopNTable"
    nlmatrixtopntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixtopntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixtopntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixtopntable.EntityData.Children = make(map[string]types.YChild)
    nlmatrixtopntable.EntityData.Children["nlMatrixTopNEntry"] = types.YChild{"Nlmatrixtopnentry", nil}
    for i := range nlmatrixtopntable.Nlmatrixtopnentry {
        nlmatrixtopntable.EntityData.Children[types.GetSegmentPath(&nlmatrixtopntable.Nlmatrixtopnentry[i])] = types.YChild{"Nlmatrixtopnentry", &nlmatrixtopntable.Nlmatrixtopnentry[i]}
    }
    nlmatrixtopntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmatrixtopntable.EntityData)
}

// RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry
// A conceptual row in the nlMatrixTopNTable.
// 
// The nlMatrixTopNControlIndex value in the index identifies the
// nlMatrixTopNControlEntry on whose behalf this entry was
// created.
// 
// An example of the indexing of this table is
// nlMatrixTopNPktRate.3.10
type RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry_Nlmatrixtopncontrolindex
    Nlmatrixtopncontrolindex interface{}

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
    Nlmatrixtopnindex interface{}

    // The protocolDirLocalIndex of the network layer protocol of this entry's
    // network address. The type is interface{} with range: 1..2147483647.
    Nlmatrixtopnprotocoldirlocalindex interface{}

    // The network layer address of the source host in this conversation.  This is
    // represented as an octet string with specific semantics and length as
    // identified by the associated nlMatrixTopNProtocolDirLocalIndex.  For
    // example, if the protocolDirLocalIndex indicates an encapsulation of ip,
    // this object is encoded as a length octet of 4, followed by the 4 octets of
    // the ip address, in network byte order. The type is string.
    Nlmatrixtopnsourceaddress interface{}

    // The network layer address of the destination host in this conversation. 
    // This is represented as an octet string with specific semantics and length
    // as identified by the associated nlMatrixTopNProtocolDirLocalIndex.  For
    // example, if the nlMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Nlmatrixtopndestaddress interface{}

    // The number of packets seen from the source host to the destination host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDPkts object. If the value of nlMatrixTopNControlRateBase is
    // nlMatrixTopNPkts, this variable will be used to sort this report. The type
    // is interface{} with range: 0..4294967295.
    Nlmatrixtopnpktrate interface{}

    // The number of packets seen from the destination host to the source host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDPkts object (note that the corresponding nlMatrixSDPkts object
    // selected is the one whose source address is equal to
    // nlMatrixTopNDestAddress and whose destination address is equal to
    // nlMatrixTopNSourceAddress.)  Note that if the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNPkts, the sort of topN
    // entries is based entirely on nlMatrixTopNPktRate, and not on the value of
    // this object. The type is interface{} with range: 0..4294967295.
    Nlmatrixtopnreversepktrate interface{}

    // The number of octets seen from the source host to the destination host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixSDOctets object.  If the value of nlMatrixTopNControlRateBase is
    // nlMatrixTopNOctets, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    Nlmatrixtopnoctetrate interface{}

    // The number of octets seen from the destination host to the source host
    // during this sampling interval, counted using the rules for counting the
    // nlMatrixDSOctets object (note that the corresponding nlMatrixSDOctets
    // object selected is the one whose source address is equal to
    // nlMatrixTopNDestAddress and whose destination address is equal to
    // nlMatrixTopNSourceAddress.)  Note that if the value of
    // nlMatrixTopNControlRateBase is equal to nlMatrixTopNOctets, the sort of
    // topN entries is based entirely on nlMatrixTopNOctetRate, and not on the
    // value of this object. The type is interface{} with range: 0..4294967295.
    Nlmatrixtopnreverseoctetrate interface{}
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetEntityData() *types.CommonEntityData {
    nlmatrixtopnentry.EntityData.YFilter = nlmatrixtopnentry.YFilter
    nlmatrixtopnentry.EntityData.YangName = "nlMatrixTopNEntry"
    nlmatrixtopnentry.EntityData.BundleName = "cisco_ios_xe"
    nlmatrixtopnentry.EntityData.ParentYangName = "nlMatrixTopNTable"
    nlmatrixtopnentry.EntityData.SegmentPath = "nlMatrixTopNEntry" + "[nlMatrixTopNControlIndex='" + fmt.Sprintf("%v", nlmatrixtopnentry.Nlmatrixtopncontrolindex) + "']" + "[nlMatrixTopNIndex='" + fmt.Sprintf("%v", nlmatrixtopnentry.Nlmatrixtopnindex) + "']"
    nlmatrixtopnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmatrixtopnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmatrixtopnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmatrixtopnentry.EntityData.Children = make(map[string]types.YChild)
    nlmatrixtopnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNControlIndex"] = types.YLeaf{"Nlmatrixtopncontrolindex", nlmatrixtopnentry.Nlmatrixtopncontrolindex}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNIndex"] = types.YLeaf{"Nlmatrixtopnindex", nlmatrixtopnentry.Nlmatrixtopnindex}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNProtocolDirLocalIndex"] = types.YLeaf{"Nlmatrixtopnprotocoldirlocalindex", nlmatrixtopnentry.Nlmatrixtopnprotocoldirlocalindex}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNSourceAddress"] = types.YLeaf{"Nlmatrixtopnsourceaddress", nlmatrixtopnentry.Nlmatrixtopnsourceaddress}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNDestAddress"] = types.YLeaf{"Nlmatrixtopndestaddress", nlmatrixtopnentry.Nlmatrixtopndestaddress}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNPktRate"] = types.YLeaf{"Nlmatrixtopnpktrate", nlmatrixtopnentry.Nlmatrixtopnpktrate}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNReversePktRate"] = types.YLeaf{"Nlmatrixtopnreversepktrate", nlmatrixtopnentry.Nlmatrixtopnreversepktrate}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNOctetRate"] = types.YLeaf{"Nlmatrixtopnoctetrate", nlmatrixtopnentry.Nlmatrixtopnoctetrate}
    nlmatrixtopnentry.EntityData.Leafs["nlMatrixTopNReverseOctetRate"] = types.YLeaf{"Nlmatrixtopnreverseoctetrate", nlmatrixtopnentry.Nlmatrixtopnreverseoctetrate}
    return &(nlmatrixtopnentry.EntityData)
}

// RMON2MIB_Alhosttable
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
type RMON2MIB_Alhosttable struct {
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
    // RMON2MIB_Alhosttable_Alhostentry.
    Alhostentry []RMON2MIB_Alhosttable_Alhostentry
}

func (alhosttable *RMON2MIB_Alhosttable) GetEntityData() *types.CommonEntityData {
    alhosttable.EntityData.YFilter = alhosttable.YFilter
    alhosttable.EntityData.YangName = "alHostTable"
    alhosttable.EntityData.BundleName = "cisco_ios_xe"
    alhosttable.EntityData.ParentYangName = "RMON2-MIB"
    alhosttable.EntityData.SegmentPath = "alHostTable"
    alhosttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alhosttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alhosttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alhosttable.EntityData.Children = make(map[string]types.YChild)
    alhosttable.EntityData.Children["alHostEntry"] = types.YChild{"Alhostentry", nil}
    for i := range alhosttable.Alhostentry {
        alhosttable.EntityData.Children[types.GetSegmentPath(&alhosttable.Alhostentry[i])] = types.YChild{"Alhostentry", &alhosttable.Alhostentry[i]}
    }
    alhosttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alhosttable.EntityData)
}

// RMON2MIB_Alhosttable_Alhostentry
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
type RMON2MIB_Alhosttable_Alhostentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry_Hlhostcontrolindex
    Hlhostcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Alhosttimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_Nlhosttable_Nlhostentry_Nlhostaddress
    Nlhostaddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex2 interface{}

    // The number of packets of this protocol type without errors transmitted to
    // this address since it was added to the alHostTable.  Note that this is the
    // number of link-layer packets, so if a single network-layer packet is
    // fragmented into several link-layer frames, this counter is incremented
    // several times. The type is interface{} with range: 0..4294967295.
    Alhostinpkts interface{}

    // The number of packets of this protocol type without errors transmitted by
    // this address since it was added to the alHostTable.  Note that this is the
    // number of link-layer packets, so if a single network-layer packet is
    // fragmented into several link-layer frames, this counter is incremented
    // several times. The type is interface{} with range: 0..4294967295.
    Alhostoutpkts interface{}

    // The number of octets transmitted to this address of this protocol type
    // since it was added to the alHostTable (excluding framing bits but including
    // FCS octets), excluding those octets in packets that contained errors.  Note
    // this doesn't count just those octets in the particular protocol frames, but
    // includes the entire packet that contained the protocol. The type is
    // interface{} with range: 0..4294967295.
    Alhostinoctets interface{}

    // The number of octets transmitted by this address of this protocol type
    // since it was added to the alHostTable (excluding framing bits but including
    // FCS octets), excluding those octets in packets that contained errors.  Note
    // this doesn't count just those octets in the particular protocol frames, but
    // includes the entire packet that contained the protocol. The type is
    // interface{} with range: 0..4294967295.
    Alhostoutoctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Alhostcreatetime interface{}
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetEntityData() *types.CommonEntityData {
    alhostentry.EntityData.YFilter = alhostentry.YFilter
    alhostentry.EntityData.YangName = "alHostEntry"
    alhostentry.EntityData.BundleName = "cisco_ios_xe"
    alhostentry.EntityData.ParentYangName = "alHostTable"
    alhostentry.EntityData.SegmentPath = "alHostEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", alhostentry.Hlhostcontrolindex) + "']" + "[alHostTimeMark='" + fmt.Sprintf("%v", alhostentry.Alhosttimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", alhostentry.Protocoldirlocalindex) + "']" + "[nlHostAddress='" + fmt.Sprintf("%v", alhostentry.Nlhostaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", alhostentry.Protocoldirlocalindex2) + "']"
    alhostentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alhostentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alhostentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alhostentry.EntityData.Children = make(map[string]types.YChild)
    alhostentry.EntityData.Leafs = make(map[string]types.YLeaf)
    alhostentry.EntityData.Leafs["hlHostControlIndex"] = types.YLeaf{"Hlhostcontrolindex", alhostentry.Hlhostcontrolindex}
    alhostentry.EntityData.Leafs["alHostTimeMark"] = types.YLeaf{"Alhosttimemark", alhostentry.Alhosttimemark}
    alhostentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", alhostentry.Protocoldirlocalindex}
    alhostentry.EntityData.Leafs["nlHostAddress"] = types.YLeaf{"Nlhostaddress", alhostentry.Nlhostaddress}
    alhostentry.EntityData.Leafs["protocolDirLocalIndex_2"] = types.YLeaf{"Protocoldirlocalindex2", alhostentry.Protocoldirlocalindex2}
    alhostentry.EntityData.Leafs["alHostInPkts"] = types.YLeaf{"Alhostinpkts", alhostentry.Alhostinpkts}
    alhostentry.EntityData.Leafs["alHostOutPkts"] = types.YLeaf{"Alhostoutpkts", alhostentry.Alhostoutpkts}
    alhostentry.EntityData.Leafs["alHostInOctets"] = types.YLeaf{"Alhostinoctets", alhostentry.Alhostinoctets}
    alhostentry.EntityData.Leafs["alHostOutOctets"] = types.YLeaf{"Alhostoutoctets", alhostentry.Alhostoutoctets}
    alhostentry.EntityData.Leafs["alHostCreateTime"] = types.YLeaf{"Alhostcreatetime", alhostentry.Alhostcreatetime}
    return &(alhostentry.EntityData)
}

// RMON2MIB_Almatrixsdtable
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
type RMON2MIB_Almatrixsdtable struct {
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
    // RMON2MIB_Almatrixsdtable_Almatrixsdentry.
    Almatrixsdentry []RMON2MIB_Almatrixsdtable_Almatrixsdentry
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetEntityData() *types.CommonEntityData {
    almatrixsdtable.EntityData.YFilter = almatrixsdtable.YFilter
    almatrixsdtable.EntityData.YangName = "alMatrixSDTable"
    almatrixsdtable.EntityData.BundleName = "cisco_ios_xe"
    almatrixsdtable.EntityData.ParentYangName = "RMON2-MIB"
    almatrixsdtable.EntityData.SegmentPath = "alMatrixSDTable"
    almatrixsdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixsdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixsdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixsdtable.EntityData.Children = make(map[string]types.YChild)
    almatrixsdtable.EntityData.Children["alMatrixSDEntry"] = types.YChild{"Almatrixsdentry", nil}
    for i := range almatrixsdtable.Almatrixsdentry {
        almatrixsdtable.EntityData.Children[types.GetSegmentPath(&almatrixsdtable.Almatrixsdentry[i])] = types.YChild{"Almatrixsdentry", &almatrixsdtable.Almatrixsdentry[i]}
    }
    almatrixsdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(almatrixsdtable.EntityData)
}

// RMON2MIB_Almatrixsdtable_Almatrixsdentry
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
type RMON2MIB_Almatrixsdtable_Almatrixsdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry_Hlmatrixcontrolindex
    Hlmatrixcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Almatrixsdtimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry_Nlmatrixsdsourceaddress
    Nlmatrixsdsourceaddress interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry_Nlmatrixsddestaddress
    Nlmatrixsddestaddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex2 interface{}

    // The number of packets of this protocol type without errors transmitted from
    // the source address to the destination address since this entry was added to
    // the alMatrixSDTable.  Note that this is the number of link-layer packets,
    // so if a single network-layer packet is fragmented into several link-layer
    // frames, this counter is incremented several times. The type is interface{}
    // with range: 0..4294967295.
    Almatrixsdpkts interface{}

    // The number of octets in packets of this protocol type transmitted from the
    // source address to the destination address since this entry was added to the
    // alMatrixSDTable (excluding framing bits but including FCS octets),
    // excluding those octets in packets that contained errors.  Note this doesn't
    // count just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    Almatrixsdoctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Almatrixsdcreatetime interface{}
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetEntityData() *types.CommonEntityData {
    almatrixsdentry.EntityData.YFilter = almatrixsdentry.YFilter
    almatrixsdentry.EntityData.YangName = "alMatrixSDEntry"
    almatrixsdentry.EntityData.BundleName = "cisco_ios_xe"
    almatrixsdentry.EntityData.ParentYangName = "alMatrixSDTable"
    almatrixsdentry.EntityData.SegmentPath = "alMatrixSDEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", almatrixsdentry.Hlmatrixcontrolindex) + "']" + "[alMatrixSDTimeMark='" + fmt.Sprintf("%v", almatrixsdentry.Almatrixsdtimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", almatrixsdentry.Protocoldirlocalindex) + "']" + "[nlMatrixSDSourceAddress='" + fmt.Sprintf("%v", almatrixsdentry.Nlmatrixsdsourceaddress) + "']" + "[nlMatrixSDDestAddress='" + fmt.Sprintf("%v", almatrixsdentry.Nlmatrixsddestaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", almatrixsdentry.Protocoldirlocalindex2) + "']"
    almatrixsdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixsdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixsdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixsdentry.EntityData.Children = make(map[string]types.YChild)
    almatrixsdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    almatrixsdentry.EntityData.Leafs["hlMatrixControlIndex"] = types.YLeaf{"Hlmatrixcontrolindex", almatrixsdentry.Hlmatrixcontrolindex}
    almatrixsdentry.EntityData.Leafs["alMatrixSDTimeMark"] = types.YLeaf{"Almatrixsdtimemark", almatrixsdentry.Almatrixsdtimemark}
    almatrixsdentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", almatrixsdentry.Protocoldirlocalindex}
    almatrixsdentry.EntityData.Leafs["nlMatrixSDSourceAddress"] = types.YLeaf{"Nlmatrixsdsourceaddress", almatrixsdentry.Nlmatrixsdsourceaddress}
    almatrixsdentry.EntityData.Leafs["nlMatrixSDDestAddress"] = types.YLeaf{"Nlmatrixsddestaddress", almatrixsdentry.Nlmatrixsddestaddress}
    almatrixsdentry.EntityData.Leafs["protocolDirLocalIndex_2"] = types.YLeaf{"Protocoldirlocalindex2", almatrixsdentry.Protocoldirlocalindex2}
    almatrixsdentry.EntityData.Leafs["alMatrixSDPkts"] = types.YLeaf{"Almatrixsdpkts", almatrixsdentry.Almatrixsdpkts}
    almatrixsdentry.EntityData.Leafs["alMatrixSDOctets"] = types.YLeaf{"Almatrixsdoctets", almatrixsdentry.Almatrixsdoctets}
    almatrixsdentry.EntityData.Leafs["alMatrixSDCreateTime"] = types.YLeaf{"Almatrixsdcreatetime", almatrixsdentry.Almatrixsdcreatetime}
    return &(almatrixsdentry.EntityData)
}

// RMON2MIB_Almatrixdstable
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
type RMON2MIB_Almatrixdstable struct {
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
    // RMON2MIB_Almatrixdstable_Almatrixdsentry.
    Almatrixdsentry []RMON2MIB_Almatrixdstable_Almatrixdsentry
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetEntityData() *types.CommonEntityData {
    almatrixdstable.EntityData.YFilter = almatrixdstable.YFilter
    almatrixdstable.EntityData.YangName = "alMatrixDSTable"
    almatrixdstable.EntityData.BundleName = "cisco_ios_xe"
    almatrixdstable.EntityData.ParentYangName = "RMON2-MIB"
    almatrixdstable.EntityData.SegmentPath = "alMatrixDSTable"
    almatrixdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixdstable.EntityData.Children = make(map[string]types.YChild)
    almatrixdstable.EntityData.Children["alMatrixDSEntry"] = types.YChild{"Almatrixdsentry", nil}
    for i := range almatrixdstable.Almatrixdsentry {
        almatrixdstable.EntityData.Children[types.GetSegmentPath(&almatrixdstable.Almatrixdsentry[i])] = types.YChild{"Almatrixdsentry", &almatrixdstable.Almatrixdsentry[i]}
    }
    almatrixdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(almatrixdstable.EntityData)
}

// RMON2MIB_Almatrixdstable_Almatrixdsentry
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
type RMON2MIB_Almatrixdstable_Almatrixdsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry_Hlmatrixcontrolindex
    Hlmatrixcontrolindex interface{}

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Almatrixdstimemark interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry_Nlmatrixdsdestaddress
    Nlmatrixdsdestaddress interface{}

    // This attribute is a key. The type is string. Refers to
    // rmon2_mib.RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry_Nlmatrixdssourceaddress
    Nlmatrixdssourceaddress interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // rmon2_mib.RMON2MIB_Protocoldirtable_Protocoldirentry_Protocoldirlocalindex
    Protocoldirlocalindex2 interface{}

    // The number of packets of this protocol type without errors transmitted from
    // the source address to the destination address since this entry was added to
    // the alMatrixDSTable.  Note that this is the number of link-layer packets,
    // so if a single network-layer packet is fragmented into several link-layer
    // frames, this counter is incremented several times. The type is interface{}
    // with range: 0..4294967295.
    Almatrixdspkts interface{}

    // The number of octets in packets of this protocol type transmitted from the
    // source address to the destination address since this entry was added to the
    // alMatrixDSTable (excluding framing bits but including FCS octets),
    // excluding those octets in packets that contained errors.  Note this doesn't
    // count just those octets in the particular protocol frames, but includes the
    // entire packet that contained the protocol. The type is interface{} with
    // range: 0..4294967295.
    Almatrixdsoctets interface{}

    // The value of sysUpTime when this entry was last activated. This can be used
    // by the management station to ensure that the entry has not been deleted and
    // recreated between polls. The type is interface{} with range: 0..4294967295.
    Almatrixdscreatetime interface{}
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetEntityData() *types.CommonEntityData {
    almatrixdsentry.EntityData.YFilter = almatrixdsentry.YFilter
    almatrixdsentry.EntityData.YangName = "alMatrixDSEntry"
    almatrixdsentry.EntityData.BundleName = "cisco_ios_xe"
    almatrixdsentry.EntityData.ParentYangName = "alMatrixDSTable"
    almatrixdsentry.EntityData.SegmentPath = "alMatrixDSEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", almatrixdsentry.Hlmatrixcontrolindex) + "']" + "[alMatrixDSTimeMark='" + fmt.Sprintf("%v", almatrixdsentry.Almatrixdstimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", almatrixdsentry.Protocoldirlocalindex) + "']" + "[nlMatrixDSDestAddress='" + fmt.Sprintf("%v", almatrixdsentry.Nlmatrixdsdestaddress) + "']" + "[nlMatrixDSSourceAddress='" + fmt.Sprintf("%v", almatrixdsentry.Nlmatrixdssourceaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", almatrixdsentry.Protocoldirlocalindex2) + "']"
    almatrixdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixdsentry.EntityData.Children = make(map[string]types.YChild)
    almatrixdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    almatrixdsentry.EntityData.Leafs["hlMatrixControlIndex"] = types.YLeaf{"Hlmatrixcontrolindex", almatrixdsentry.Hlmatrixcontrolindex}
    almatrixdsentry.EntityData.Leafs["alMatrixDSTimeMark"] = types.YLeaf{"Almatrixdstimemark", almatrixdsentry.Almatrixdstimemark}
    almatrixdsentry.EntityData.Leafs["protocolDirLocalIndex"] = types.YLeaf{"Protocoldirlocalindex", almatrixdsentry.Protocoldirlocalindex}
    almatrixdsentry.EntityData.Leafs["nlMatrixDSDestAddress"] = types.YLeaf{"Nlmatrixdsdestaddress", almatrixdsentry.Nlmatrixdsdestaddress}
    almatrixdsentry.EntityData.Leafs["nlMatrixDSSourceAddress"] = types.YLeaf{"Nlmatrixdssourceaddress", almatrixdsentry.Nlmatrixdssourceaddress}
    almatrixdsentry.EntityData.Leafs["protocolDirLocalIndex_2"] = types.YLeaf{"Protocoldirlocalindex2", almatrixdsentry.Protocoldirlocalindex2}
    almatrixdsentry.EntityData.Leafs["alMatrixDSPkts"] = types.YLeaf{"Almatrixdspkts", almatrixdsentry.Almatrixdspkts}
    almatrixdsentry.EntityData.Leafs["alMatrixDSOctets"] = types.YLeaf{"Almatrixdsoctets", almatrixdsentry.Almatrixdsoctets}
    almatrixdsentry.EntityData.Leafs["alMatrixDSCreateTime"] = types.YLeaf{"Almatrixdscreatetime", almatrixdsentry.Almatrixdscreatetime}
    return &(almatrixdsentry.EntityData)
}

// RMON2MIB_Almatrixtopncontroltable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_Almatrixtopncontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNControlTable.  An example of the
    // indexing of this table is alMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry.
    Almatrixtopncontrolentry []RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetEntityData() *types.CommonEntityData {
    almatrixtopncontroltable.EntityData.YFilter = almatrixtopncontroltable.YFilter
    almatrixtopncontroltable.EntityData.YangName = "alMatrixTopNControlTable"
    almatrixtopncontroltable.EntityData.BundleName = "cisco_ios_xe"
    almatrixtopncontroltable.EntityData.ParentYangName = "RMON2-MIB"
    almatrixtopncontroltable.EntityData.SegmentPath = "alMatrixTopNControlTable"
    almatrixtopncontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixtopncontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixtopncontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixtopncontroltable.EntityData.Children = make(map[string]types.YChild)
    almatrixtopncontroltable.EntityData.Children["alMatrixTopNControlEntry"] = types.YChild{"Almatrixtopncontrolentry", nil}
    for i := range almatrixtopncontroltable.Almatrixtopncontrolentry {
        almatrixtopncontroltable.EntityData.Children[types.GetSegmentPath(&almatrixtopncontroltable.Almatrixtopncontrolentry[i])] = types.YChild{"Almatrixtopncontrolentry", &almatrixtopncontroltable.Almatrixtopncontrolentry[i]}
    }
    almatrixtopncontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(almatrixtopncontroltable.EntityData)
}

// RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry
// A conceptual row in the alMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// alMatrixTopNControlDuration.3
type RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // alMatrixTopNControlTable.  Each such entry defines one top N report
    // prepared for one interface. The type is interface{} with range: 1..65535.
    Almatrixtopncontrolindex interface{}

    // The alMatrix[SD/DS] table for which a top N report will be prepared on
    // behalf of this entry.  The alMatrix[SD/DS] table is identified by the value
    // of the hlMatrixControlIndex for that table - that value is used here to
    // identify the particular table.  This object may not be modified if the
    // associated alMatrixTopNControlStatus object is equal to active(1). The type
    // is interface{} with range: 1..65535.
    Almatrixtopncontrolmatrixindex interface{}

    // The variable for each alMatrix[SD/DS] entry that the     
    // alMatrixTopNEntries are sorted by, as well as the selector of the view of
    // the matrix table that will be used.  The values alMatrixTopNTerminalsPkts
    // and alMatrixTopNTerminalsOctets cause collection only from protocols that
    // have no child protocols that are counted.  The values alMatrixTopNAllPkts
    // and alMatrixTopNAllOctets cause collection from all alMatrix entries.  This
    // object may not be modified if the associated alMatrixTopNControlStatus
    // object is equal to active(1). The type is Almatrixtopncontrolratebase.
    Almatrixtopncontrolratebase interface{}

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
    Almatrixtopncontroltimeremaining interface{}

    // The number of reports that have been generated by this entry. The type is
    // interface{} with range: 0..4294967295.
    Almatrixtopncontrolgeneratedreports interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval.  When the associated alMatrixTopNControlTimeRemaining
    // object is set, this object shall be set by the probe to the same value and
    // shall not be modified until the next time the
    // alMatrixTopNControlTimeRemaining is set.  This value shall be zero if no
    // reports have been requested for this alMatrixTopNControlEntry. The type is
    // interface{} with range: -2147483648..2147483647.
    Almatrixtopncontrolduration interface{}

    // The maximum number of matrix entries requested for this report.  When this
    // object is created or modified, the probe should set
    // alMatrixTopNControlGrantedSize as closely to this object as is possible for
    // the particular probe implementation and available resources. The type is
    // interface{} with range: 0..2147483647.
    Almatrixtopncontrolrequestedsize interface{}

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
    Almatrixtopncontrolgrantedsize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated
    // alMatrixTopNControlTimeRemaining object was modified to start the requested
    // report or the time the report was last automatically (re)started.  This
    // object may be used by the management station to determine if a report was
    // missed or not. The type is interface{} with range: 0..4294967295.
    Almatrixtopncontrolstarttime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Almatrixtopncontrolowner interface{}

    // The status of this alMatrixTopNControlEntry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value.  If
    // this object is not equal to active(1), all associated entries in the
    // alMatrixTopNTable shall be deleted by the agent. The type is RowStatus.
    Almatrixtopncontrolstatus interface{}
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetEntityData() *types.CommonEntityData {
    almatrixtopncontrolentry.EntityData.YFilter = almatrixtopncontrolentry.YFilter
    almatrixtopncontrolentry.EntityData.YangName = "alMatrixTopNControlEntry"
    almatrixtopncontrolentry.EntityData.BundleName = "cisco_ios_xe"
    almatrixtopncontrolentry.EntityData.ParentYangName = "alMatrixTopNControlTable"
    almatrixtopncontrolentry.EntityData.SegmentPath = "alMatrixTopNControlEntry" + "[alMatrixTopNControlIndex='" + fmt.Sprintf("%v", almatrixtopncontrolentry.Almatrixtopncontrolindex) + "']"
    almatrixtopncontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixtopncontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixtopncontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixtopncontrolentry.EntityData.Children = make(map[string]types.YChild)
    almatrixtopncontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlIndex"] = types.YLeaf{"Almatrixtopncontrolindex", almatrixtopncontrolentry.Almatrixtopncontrolindex}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlMatrixIndex"] = types.YLeaf{"Almatrixtopncontrolmatrixindex", almatrixtopncontrolentry.Almatrixtopncontrolmatrixindex}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlRateBase"] = types.YLeaf{"Almatrixtopncontrolratebase", almatrixtopncontrolentry.Almatrixtopncontrolratebase}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlTimeRemaining"] = types.YLeaf{"Almatrixtopncontroltimeremaining", almatrixtopncontrolentry.Almatrixtopncontroltimeremaining}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlGeneratedReports"] = types.YLeaf{"Almatrixtopncontrolgeneratedreports", almatrixtopncontrolentry.Almatrixtopncontrolgeneratedreports}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlDuration"] = types.YLeaf{"Almatrixtopncontrolduration", almatrixtopncontrolentry.Almatrixtopncontrolduration}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlRequestedSize"] = types.YLeaf{"Almatrixtopncontrolrequestedsize", almatrixtopncontrolentry.Almatrixtopncontrolrequestedsize}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlGrantedSize"] = types.YLeaf{"Almatrixtopncontrolgrantedsize", almatrixtopncontrolentry.Almatrixtopncontrolgrantedsize}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlStartTime"] = types.YLeaf{"Almatrixtopncontrolstarttime", almatrixtopncontrolentry.Almatrixtopncontrolstarttime}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlOwner"] = types.YLeaf{"Almatrixtopncontrolowner", almatrixtopncontrolentry.Almatrixtopncontrolowner}
    almatrixtopncontrolentry.EntityData.Leafs["alMatrixTopNControlStatus"] = types.YLeaf{"Almatrixtopncontrolstatus", almatrixtopncontrolentry.Almatrixtopncontrolstatus}
    return &(almatrixtopncontrolentry.EntityData)
}

// RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase represents alMatrixTopNControlStatus object is equal to active(1).
type RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase string

const (
    RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase_alMatrixTopNTerminalsPkts RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase = "alMatrixTopNTerminalsPkts"

    RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase_alMatrixTopNTerminalsOctets RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase = "alMatrixTopNTerminalsOctets"

    RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase_alMatrixTopNAllPkts RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase = "alMatrixTopNAllPkts"

    RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase_alMatrixTopNAllOctets RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolratebase = "alMatrixTopNAllOctets"
)

// RMON2MIB_Almatrixtopntable
// A set of statistics for those application layer matrix
// entries that have counted the highest number of octets or
// packets.
type RMON2MIB_Almatrixtopntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNTable.  The alMatrixTopNControlIndex
    // value in the index identifies the alMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // alMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_Almatrixtopntable_Almatrixtopnentry.
    Almatrixtopnentry []RMON2MIB_Almatrixtopntable_Almatrixtopnentry
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetEntityData() *types.CommonEntityData {
    almatrixtopntable.EntityData.YFilter = almatrixtopntable.YFilter
    almatrixtopntable.EntityData.YangName = "alMatrixTopNTable"
    almatrixtopntable.EntityData.BundleName = "cisco_ios_xe"
    almatrixtopntable.EntityData.ParentYangName = "RMON2-MIB"
    almatrixtopntable.EntityData.SegmentPath = "alMatrixTopNTable"
    almatrixtopntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixtopntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixtopntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixtopntable.EntityData.Children = make(map[string]types.YChild)
    almatrixtopntable.EntityData.Children["alMatrixTopNEntry"] = types.YChild{"Almatrixtopnentry", nil}
    for i := range almatrixtopntable.Almatrixtopnentry {
        almatrixtopntable.EntityData.Children[types.GetSegmentPath(&almatrixtopntable.Almatrixtopnentry[i])] = types.YChild{"Almatrixtopnentry", &almatrixtopntable.Almatrixtopnentry[i]}
    }
    almatrixtopntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(almatrixtopntable.EntityData)
}

// RMON2MIB_Almatrixtopntable_Almatrixtopnentry
// A conceptual row in the alMatrixTopNTable.
// 
// The alMatrixTopNControlIndex value in the index identifies
// the alMatrixTopNControlEntry on whose behalf this entry was
// created.
// 
// An example of the indexing of this table is
// alMatrixTopNPktRate.3.10
type RMON2MIB_Almatrixtopntable_Almatrixtopnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry_Almatrixtopncontrolindex
    Almatrixtopncontrolindex interface{}

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
    Almatrixtopnindex interface{}

    // The protocolDirLocalIndex of the network layer protocol of this entry's
    // network address. The type is interface{} with range: 1..2147483647.
    Almatrixtopnprotocoldirlocalindex interface{}

    // The network layer address of the source host in this conversation. This is
    // represented as an octet string with specific semantics and length as
    // identified      by the associated alMatrixTopNProtocolDirLocalIndex.  For
    // example, if the alMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Almatrixtopnsourceaddress interface{}

    // The network layer address of the destination host in this conversation. 
    // This is represented as an octet string with specific semantics and length
    // as identified by the associated alMatrixTopNProtocolDirLocalIndex.  For
    // example, if the alMatrixTopNProtocolDirLocalIndex indicates an
    // encapsulation of ip, this object is encoded as a length octet of 4,
    // followed by the 4 octets of the ip address, in network byte order. The type
    // is string.
    Almatrixtopndestaddress interface{}

    // The type of the protocol counted by this matrix entry. The type is
    // interface{} with range: 1..2147483647.
    Almatrixtopnappprotocoldirlocalindex interface{}

    // The number of packets seen of this protocol from the source host to the
    // destination host during this sampling interval, counted using the rules for
    // counting the alMatrixSDPkts object.  If the value of
    // alMatrixTopNControlRateBase is alMatrixTopNTerminalsPkts or
    // alMatrixTopNAllPkts, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    Almatrixtopnpktrate interface{}

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
    Almatrixtopnreversepktrate interface{}

    // The number of octets seen of this protocol from the source host to the
    // destination host during this sampling interval, counted using the rules for
    // counting the alMatrixSDOctets object.  If the value of
    // alMatrixTopNControlRateBase is alMatrixTopNTerminalsOctets or
    // alMatrixTopNAllOctets, this variable will be used to sort this report. The
    // type is interface{} with range: 0..4294967295.
    Almatrixtopnoctetrate interface{}

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
    Almatrixtopnreverseoctetrate interface{}
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetEntityData() *types.CommonEntityData {
    almatrixtopnentry.EntityData.YFilter = almatrixtopnentry.YFilter
    almatrixtopnentry.EntityData.YangName = "alMatrixTopNEntry"
    almatrixtopnentry.EntityData.BundleName = "cisco_ios_xe"
    almatrixtopnentry.EntityData.ParentYangName = "alMatrixTopNTable"
    almatrixtopnentry.EntityData.SegmentPath = "alMatrixTopNEntry" + "[alMatrixTopNControlIndex='" + fmt.Sprintf("%v", almatrixtopnentry.Almatrixtopncontrolindex) + "']" + "[alMatrixTopNIndex='" + fmt.Sprintf("%v", almatrixtopnentry.Almatrixtopnindex) + "']"
    almatrixtopnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    almatrixtopnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    almatrixtopnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    almatrixtopnentry.EntityData.Children = make(map[string]types.YChild)
    almatrixtopnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNControlIndex"] = types.YLeaf{"Almatrixtopncontrolindex", almatrixtopnentry.Almatrixtopncontrolindex}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNIndex"] = types.YLeaf{"Almatrixtopnindex", almatrixtopnentry.Almatrixtopnindex}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNProtocolDirLocalIndex"] = types.YLeaf{"Almatrixtopnprotocoldirlocalindex", almatrixtopnentry.Almatrixtopnprotocoldirlocalindex}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNSourceAddress"] = types.YLeaf{"Almatrixtopnsourceaddress", almatrixtopnentry.Almatrixtopnsourceaddress}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNDestAddress"] = types.YLeaf{"Almatrixtopndestaddress", almatrixtopnentry.Almatrixtopndestaddress}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNAppProtocolDirLocalIndex"] = types.YLeaf{"Almatrixtopnappprotocoldirlocalindex", almatrixtopnentry.Almatrixtopnappprotocoldirlocalindex}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNPktRate"] = types.YLeaf{"Almatrixtopnpktrate", almatrixtopnentry.Almatrixtopnpktrate}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNReversePktRate"] = types.YLeaf{"Almatrixtopnreversepktrate", almatrixtopnentry.Almatrixtopnreversepktrate}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNOctetRate"] = types.YLeaf{"Almatrixtopnoctetrate", almatrixtopnentry.Almatrixtopnoctetrate}
    almatrixtopnentry.EntityData.Leafs["alMatrixTopNReverseOctetRate"] = types.YLeaf{"Almatrixtopnreverseoctetrate", almatrixtopnentry.Almatrixtopnreverseoctetrate}
    return &(almatrixtopnentry.EntityData)
}

// RMON2MIB_Usrhistorycontroltable
// A list of data-collection configuration entries.
type RMON2MIB_Usrhistorycontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a group of user-defined MIB objects to be
    // sampled periodically (called a bucket-group).  For example, an instance of
    // usrHistoryControlInterval might be named usrHistoryControlInterval.1. The
    // type is slice of RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry.
    Usrhistorycontrolentry []RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetEntityData() *types.CommonEntityData {
    usrhistorycontroltable.EntityData.YFilter = usrhistorycontroltable.YFilter
    usrhistorycontroltable.EntityData.YangName = "usrHistoryControlTable"
    usrhistorycontroltable.EntityData.BundleName = "cisco_ios_xe"
    usrhistorycontroltable.EntityData.ParentYangName = "RMON2-MIB"
    usrhistorycontroltable.EntityData.SegmentPath = "usrHistoryControlTable"
    usrhistorycontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistorycontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistorycontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistorycontroltable.EntityData.Children = make(map[string]types.YChild)
    usrhistorycontroltable.EntityData.Children["usrHistoryControlEntry"] = types.YChild{"Usrhistorycontrolentry", nil}
    for i := range usrhistorycontroltable.Usrhistorycontrolentry {
        usrhistorycontroltable.EntityData.Children[types.GetSegmentPath(&usrhistorycontroltable.Usrhistorycontrolentry[i])] = types.YChild{"Usrhistorycontrolentry", &usrhistorycontroltable.Usrhistorycontrolentry[i]}
    }
    usrhistorycontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usrhistorycontroltable.EntityData)
}

// RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry
// A list of parameters that set up a group of user-defined
// MIB objects to be sampled periodically (called a
// bucket-group).
// 
// For example, an instance of usrHistoryControlInterval
// might be named usrHistoryControlInterval.1
type RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // usrHistoryControlTable.  Each such entry defines a set of samples at a
    // particular interval for a specified set of MIB instances available from the
    // managed system. The type is interface{} with range: 1..65535.
    Usrhistorycontrolindex interface{}

    // The number of MIB objects to be collected in the portion of usrHistoryTable
    // associated with this usrHistoryControlEntry.  This object may not be
    // modified if the associated instance of usrHistoryControlStatus is equal to
    // active(1). The type is interface{} with range: 1..65535.
    Usrhistorycontrolobjects interface{}

    // The requested number of discrete time intervals over which data is to be
    // saved in the part of the usrHistoryTable associated with this
    // usrHistoryControlEntry.  When this object is created or modified, the probe
    // should set usrHistoryControlBucketsGranted as closely to this object as is
    // possible for the particular probe implementation and available resources.
    // The type is interface{} with range: 1..65535.
    Usrhistorycontrolbucketsrequested interface{}

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
    Usrhistorycontrolbucketsgranted interface{}

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
    Usrhistorycontrolinterval interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Usrhistorycontrolowner interface{}

    // The status of this variable history control entry.  An entry may not exist
    // in the active state unless all objects in the entry have an appropriate
    // value.  If this object is not equal to active(1), all associated entries in
    // the usrHistoryTable shall be deleted. The type is RowStatus.
    Usrhistorycontrolstatus interface{}
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetEntityData() *types.CommonEntityData {
    usrhistorycontrolentry.EntityData.YFilter = usrhistorycontrolentry.YFilter
    usrhistorycontrolentry.EntityData.YangName = "usrHistoryControlEntry"
    usrhistorycontrolentry.EntityData.BundleName = "cisco_ios_xe"
    usrhistorycontrolentry.EntityData.ParentYangName = "usrHistoryControlTable"
    usrhistorycontrolentry.EntityData.SegmentPath = "usrHistoryControlEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistorycontrolentry.Usrhistorycontrolindex) + "']"
    usrhistorycontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistorycontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistorycontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistorycontrolentry.EntityData.Children = make(map[string]types.YChild)
    usrhistorycontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlIndex"] = types.YLeaf{"Usrhistorycontrolindex", usrhistorycontrolentry.Usrhistorycontrolindex}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlObjects"] = types.YLeaf{"Usrhistorycontrolobjects", usrhistorycontrolentry.Usrhistorycontrolobjects}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlBucketsRequested"] = types.YLeaf{"Usrhistorycontrolbucketsrequested", usrhistorycontrolentry.Usrhistorycontrolbucketsrequested}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlBucketsGranted"] = types.YLeaf{"Usrhistorycontrolbucketsgranted", usrhistorycontrolentry.Usrhistorycontrolbucketsgranted}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlInterval"] = types.YLeaf{"Usrhistorycontrolinterval", usrhistorycontrolentry.Usrhistorycontrolinterval}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlOwner"] = types.YLeaf{"Usrhistorycontrolowner", usrhistorycontrolentry.Usrhistorycontrolowner}
    usrhistorycontrolentry.EntityData.Leafs["usrHistoryControlStatus"] = types.YLeaf{"Usrhistorycontrolstatus", usrhistorycontrolentry.Usrhistorycontrolstatus}
    return &(usrhistorycontrolentry.EntityData)
}

// RMON2MIB_Usrhistoryobjecttable
// A list of data-collection configuration entries.
type RMON2MIB_Usrhistoryobjecttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of MIB instances to be sampled periodically.  Entries in this table
    // are created when an associated usrHistoryControlObjects object is created. 
    // The usrHistoryControlIndex value in the index is that of the associated
    // usrHistoryControlEntry.  For example, an instance of
    // usrHistoryObjectVariable might be usrHistoryObjectVariable.1.3. The type is
    // slice of RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry.
    Usrhistoryobjectentry []RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetEntityData() *types.CommonEntityData {
    usrhistoryobjecttable.EntityData.YFilter = usrhistoryobjecttable.YFilter
    usrhistoryobjecttable.EntityData.YangName = "usrHistoryObjectTable"
    usrhistoryobjecttable.EntityData.BundleName = "cisco_ios_xe"
    usrhistoryobjecttable.EntityData.ParentYangName = "RMON2-MIB"
    usrhistoryobjecttable.EntityData.SegmentPath = "usrHistoryObjectTable"
    usrhistoryobjecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistoryobjecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistoryobjecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistoryobjecttable.EntityData.Children = make(map[string]types.YChild)
    usrhistoryobjecttable.EntityData.Children["usrHistoryObjectEntry"] = types.YChild{"Usrhistoryobjectentry", nil}
    for i := range usrhistoryobjecttable.Usrhistoryobjectentry {
        usrhistoryobjecttable.EntityData.Children[types.GetSegmentPath(&usrhistoryobjecttable.Usrhistoryobjectentry[i])] = types.YChild{"Usrhistoryobjectentry", &usrhistoryobjecttable.Usrhistoryobjectentry[i]}
    }
    usrhistoryobjecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usrhistoryobjecttable.EntityData)
}

// RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry
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
type RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry_Usrhistorycontrolindex
    Usrhistorycontrolindex interface{}

    // This attribute is a key. An index used to uniquely identify an entry in the
    // usrHistoryObject table.  Each such entry defines a MIB instance to be
    // collected periodically. The type is interface{} with range: 1..65535.
    Usrhistoryobjectindex interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Usrhistoryobjectvariable interface{}

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
    // Usrhistoryobjectsampletype.
    Usrhistoryobjectsampletype interface{}
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetEntityData() *types.CommonEntityData {
    usrhistoryobjectentry.EntityData.YFilter = usrhistoryobjectentry.YFilter
    usrhistoryobjectentry.EntityData.YangName = "usrHistoryObjectEntry"
    usrhistoryobjectentry.EntityData.BundleName = "cisco_ios_xe"
    usrhistoryobjectentry.EntityData.ParentYangName = "usrHistoryObjectTable"
    usrhistoryobjectentry.EntityData.SegmentPath = "usrHistoryObjectEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistoryobjectentry.Usrhistorycontrolindex) + "']" + "[usrHistoryObjectIndex='" + fmt.Sprintf("%v", usrhistoryobjectentry.Usrhistoryobjectindex) + "']"
    usrhistoryobjectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistoryobjectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistoryobjectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistoryobjectentry.EntityData.Children = make(map[string]types.YChild)
    usrhistoryobjectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    usrhistoryobjectentry.EntityData.Leafs["usrHistoryControlIndex"] = types.YLeaf{"Usrhistorycontrolindex", usrhistoryobjectentry.Usrhistorycontrolindex}
    usrhistoryobjectentry.EntityData.Leafs["usrHistoryObjectIndex"] = types.YLeaf{"Usrhistoryobjectindex", usrhistoryobjectentry.Usrhistoryobjectindex}
    usrhistoryobjectentry.EntityData.Leafs["usrHistoryObjectVariable"] = types.YLeaf{"Usrhistoryobjectvariable", usrhistoryobjectentry.Usrhistoryobjectvariable}
    usrhistoryobjectentry.EntityData.Leafs["usrHistoryObjectSampleType"] = types.YLeaf{"Usrhistoryobjectsampletype", usrhistoryobjectentry.Usrhistoryobjectsampletype}
    return &(usrhistoryobjectentry.EntityData)
}

// RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype represents usrHistoryControlStatus object is equal to active(1).
type RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype string

const (
    RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype_absoluteValue RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype = "absoluteValue"

    RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype_deltaValue RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype = "deltaValue"
)

// RMON2MIB_Usrhistorytable
// A list of user defined history entries.
type RMON2MIB_Usrhistorytable struct {
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
    // type is slice of RMON2MIB_Usrhistorytable_Usrhistoryentry.
    Usrhistoryentry []RMON2MIB_Usrhistorytable_Usrhistoryentry
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetEntityData() *types.CommonEntityData {
    usrhistorytable.EntityData.YFilter = usrhistorytable.YFilter
    usrhistorytable.EntityData.YangName = "usrHistoryTable"
    usrhistorytable.EntityData.BundleName = "cisco_ios_xe"
    usrhistorytable.EntityData.ParentYangName = "RMON2-MIB"
    usrhistorytable.EntityData.SegmentPath = "usrHistoryTable"
    usrhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistorytable.EntityData.Children = make(map[string]types.YChild)
    usrhistorytable.EntityData.Children["usrHistoryEntry"] = types.YChild{"Usrhistoryentry", nil}
    for i := range usrhistorytable.Usrhistoryentry {
        usrhistorytable.EntityData.Children[types.GetSegmentPath(&usrhistorytable.Usrhistoryentry[i])] = types.YChild{"Usrhistoryentry", &usrhistorytable.Usrhistoryentry[i]}
    }
    usrhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usrhistorytable.EntityData)
}

// RMON2MIB_Usrhistorytable_Usrhistoryentry
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
type RMON2MIB_Usrhistorytable_Usrhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry_Usrhistorycontrolindex
    Usrhistorycontrolindex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // usrHistoryControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range: 1..2147483647.
    Usrhistorysampleindex interface{}

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // rmon2_mib.RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectindex
    Usrhistoryobjectindex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval. Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    Usrhistoryintervalstart interface{}

    // The value of sysUpTime at the end of the interval over which this sample
    // was measured. The type is interface{} with range: 0..4294967295.
    Usrhistoryintervalend interface{}

    // The absolute value (i.e. unsigned value) of the user-specified statistic
    // during the last sampling period. The value during the current sampling
    // period is not made available until the period is completed.  To obtain the
    // true value for this sampling interval, the associated instance of
    // usrHistoryValStatus must be checked, and usrHistoryAbsValue adjusted as
    // necessary.  If the MIB instance could not be accessed during the sampling
    // interval, then this object will have a value of zero and the associated
    // instance of usrHistoryValStatus will be set to 'valueNotAvailable(1)'. The
    // type is interface{} with range: 0..4294967295.
    Usrhistoryabsvalue interface{}

    // This object indicates the validity and sign of the data in the associated
    // instance of usrHistoryAbsValue.  If the MIB instance could not be accessed
    // during the sampling interval, then 'valueNotAvailable(1)' will be returned.
    // If the sample is valid and actual value of the sample is greater than or
    // equal to zero then 'valuePositive(2)' is returned.  If the sample is valid
    // and the actual value of the sample is less than zero, 'valueNegative(3)'
    // will be returned. The associated instance of usrHistoryAbsValue should be
    // multiplied by -1 to obtain the true sample value. The type is
    // Usrhistoryvalstatus.
    Usrhistoryvalstatus interface{}
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetEntityData() *types.CommonEntityData {
    usrhistoryentry.EntityData.YFilter = usrhistoryentry.YFilter
    usrhistoryentry.EntityData.YangName = "usrHistoryEntry"
    usrhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    usrhistoryentry.EntityData.ParentYangName = "usrHistoryTable"
    usrhistoryentry.EntityData.SegmentPath = "usrHistoryEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistorycontrolindex) + "']" + "[usrHistorySampleIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistorysampleindex) + "']" + "[usrHistoryObjectIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistoryobjectindex) + "']"
    usrhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usrhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usrhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usrhistoryentry.EntityData.Children = make(map[string]types.YChild)
    usrhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    usrhistoryentry.EntityData.Leafs["usrHistoryControlIndex"] = types.YLeaf{"Usrhistorycontrolindex", usrhistoryentry.Usrhistorycontrolindex}
    usrhistoryentry.EntityData.Leafs["usrHistorySampleIndex"] = types.YLeaf{"Usrhistorysampleindex", usrhistoryentry.Usrhistorysampleindex}
    usrhistoryentry.EntityData.Leafs["usrHistoryObjectIndex"] = types.YLeaf{"Usrhistoryobjectindex", usrhistoryentry.Usrhistoryobjectindex}
    usrhistoryentry.EntityData.Leafs["usrHistoryIntervalStart"] = types.YLeaf{"Usrhistoryintervalstart", usrhistoryentry.Usrhistoryintervalstart}
    usrhistoryentry.EntityData.Leafs["usrHistoryIntervalEnd"] = types.YLeaf{"Usrhistoryintervalend", usrhistoryentry.Usrhistoryintervalend}
    usrhistoryentry.EntityData.Leafs["usrHistoryAbsValue"] = types.YLeaf{"Usrhistoryabsvalue", usrhistoryentry.Usrhistoryabsvalue}
    usrhistoryentry.EntityData.Leafs["usrHistoryValStatus"] = types.YLeaf{"Usrhistoryvalstatus", usrhistoryentry.Usrhistoryvalstatus}
    return &(usrhistoryentry.EntityData)
}

// RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus represents by -1 to obtain the true sample value.
type RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus string

const (
    RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus_valueNotAvailable RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus = "valueNotAvailable"

    RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus_valuePositive RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus = "valuePositive"

    RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus_valueNegative RMON2MIB_Usrhistorytable_Usrhistoryentry_Usrhistoryvalstatus = "valueNegative"
)

// RMON2MIB_Serialconfigtable
// A table of serial interface configuration entries.  This data
// will be stored in non-volatile memory and preserved across
// probe resets or power loss.
type RMON2MIB_Serialconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular serial interface on this
    // device. If the device has no serial interfaces, this table is empty.  The
    // index is composed of the ifIndex assigned to this serial line interface.
    // The type is slice of RMON2MIB_Serialconfigtable_Serialconfigentry.
    Serialconfigentry []RMON2MIB_Serialconfigtable_Serialconfigentry
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetEntityData() *types.CommonEntityData {
    serialconfigtable.EntityData.YFilter = serialconfigtable.YFilter
    serialconfigtable.EntityData.YangName = "serialConfigTable"
    serialconfigtable.EntityData.BundleName = "cisco_ios_xe"
    serialconfigtable.EntityData.ParentYangName = "RMON2-MIB"
    serialconfigtable.EntityData.SegmentPath = "serialConfigTable"
    serialconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialconfigtable.EntityData.Children = make(map[string]types.YChild)
    serialconfigtable.EntityData.Children["serialConfigEntry"] = types.YChild{"Serialconfigentry", nil}
    for i := range serialconfigtable.Serialconfigentry {
        serialconfigtable.EntityData.Children[types.GetSegmentPath(&serialconfigtable.Serialconfigentry[i])] = types.YChild{"Serialconfigentry", &serialconfigtable.Serialconfigentry[i]}
    }
    serialconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serialconfigtable.EntityData)
}

// RMON2MIB_Serialconfigtable_Serialconfigentry
// A set of configuration parameters for a particular
// serial interface on this device. If the device has no serial
// interfaces, this table is empty.
// 
// The index is composed of the ifIndex assigned to this serial
// line interface.
type RMON2MIB_Serialconfigtable_Serialconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range:
    // -2147483648..2147483647. Refers to
    // rfc1213_mib.RFC1213MIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The type of incoming connection to expect on this serial interface. The
    // type is Serialmode.
    Serialmode interface{}

    // The type of data link encapsulation to be used on this serial interface.
    // The type is Serialprotocol.
    Serialprotocol interface{}

    // This timeout value is used when the Management Station has initiated the
    // conversation over the serial link. This variable represents the number of
    // seconds of inactivity allowed before terminating the connection on this
    // serial interface. Use the      serialDialoutTimeout in the case where the
    // probe has initiated the connection for the purpose of sending a trap. The
    // type is interface{} with range: 1..65535.
    Serialtimeout interface{}

    // A control string which controls how a modem attached to this serial
    // interface should be initialized.  The initialization is performed once
    // during startup and again after each connection is terminated if the
    // associated serialMode has the value of modem(2).  A control string that is
    // appropriate for a wide variety of modems is: '^s^MATE0Q0V1X4 S0=1 S2=43^M'.
    // The type is string with length: 0..255.
    Serialmodeminitstring interface{}

    // A control string which specifies how to disconnect a modem connection on
    // this serial interface.  This object is only meaningful if the associated
    // serialMode has the value of modem(2). A control string that is appropriate
    // for a wide variety of modems is: '^d2^s+++^d2^sATH0^M^d2'. The type is
    // string with length: 0..255.
    Serialmodemhangupstring interface{}

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
    Serialmodemconnectresp interface{}

    // An ASCII string containing response codes that may be generated by a modem
    // to report the reason why a connection attempt has failed.  The response
    // codes are delimited by the first character in the string, for example:   
    // /NO CARRIER/BUSY/NO DIALTONE/NO ANSWER/ERROR/ If one of these response
    // codes is received via this serial interface while attempting to make a
    // modem connection, the agent will issue the hang up command as specified by
    // serialModemHangUpString.  A value that is appropriate for a wide variety of
    // modems is: '/NO CARRIER/BUSY/NO DIALTONE/NO ANSWER/ERROR/'. The type is
    // string with length: 0..255.
    Serialmodemnoconnectresp interface{}

    // This timeout value is used when the probe initiates the serial connection
    // with the intention of contacting a management station. This variable
    // represents the number of seconds of inactivity allowed before terminating
    // the connection on this serial interface. The type is interface{} with
    // range: 1..65535.
    Serialdialouttimeout interface{}

    // The status of this serialConfigEntry.  An entry may not exist in the active
    // state unless all objects in the entry have an appropriate value. The type
    // is RowStatus.
    Serialstatus interface{}
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetEntityData() *types.CommonEntityData {
    serialconfigentry.EntityData.YFilter = serialconfigentry.YFilter
    serialconfigentry.EntityData.YangName = "serialConfigEntry"
    serialconfigentry.EntityData.BundleName = "cisco_ios_xe"
    serialconfigentry.EntityData.ParentYangName = "serialConfigTable"
    serialconfigentry.EntityData.SegmentPath = "serialConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", serialconfigentry.Ifindex) + "']"
    serialconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialconfigentry.EntityData.Children = make(map[string]types.YChild)
    serialconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    serialconfigentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", serialconfigentry.Ifindex}
    serialconfigentry.EntityData.Leafs["serialMode"] = types.YLeaf{"Serialmode", serialconfigentry.Serialmode}
    serialconfigentry.EntityData.Leafs["serialProtocol"] = types.YLeaf{"Serialprotocol", serialconfigentry.Serialprotocol}
    serialconfigentry.EntityData.Leafs["serialTimeout"] = types.YLeaf{"Serialtimeout", serialconfigentry.Serialtimeout}
    serialconfigentry.EntityData.Leafs["serialModemInitString"] = types.YLeaf{"Serialmodeminitstring", serialconfigentry.Serialmodeminitstring}
    serialconfigentry.EntityData.Leafs["serialModemHangUpString"] = types.YLeaf{"Serialmodemhangupstring", serialconfigentry.Serialmodemhangupstring}
    serialconfigentry.EntityData.Leafs["serialModemConnectResp"] = types.YLeaf{"Serialmodemconnectresp", serialconfigentry.Serialmodemconnectresp}
    serialconfigentry.EntityData.Leafs["serialModemNoConnectResp"] = types.YLeaf{"Serialmodemnoconnectresp", serialconfigentry.Serialmodemnoconnectresp}
    serialconfigentry.EntityData.Leafs["serialDialoutTimeout"] = types.YLeaf{"Serialdialouttimeout", serialconfigentry.Serialdialouttimeout}
    serialconfigentry.EntityData.Leafs["serialStatus"] = types.YLeaf{"Serialstatus", serialconfigentry.Serialstatus}
    return &(serialconfigentry.EntityData)
}

// RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode represents interface.
type RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode string

const (
    RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode_direct RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode = "direct"

    RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode_modem RMON2MIB_Serialconfigtable_Serialconfigentry_Serialmode = "modem"
)

// RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol represents serial interface.
type RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol string

const (
    RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol_other RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol = "other"

    RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol_slip RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol = "slip"

    RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol_ppp RMON2MIB_Serialconfigtable_Serialconfigentry_Serialprotocol = "ppp"
)

// RMON2MIB_Netconfigtable
// A table of netConfigEntries.
type RMON2MIB_Netconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular network interface on
    // this device. If the device has no network interface, this table is empty. 
    // The index is composed of the ifIndex assigned to the corresponding
    // interface. The type is slice of RMON2MIB_Netconfigtable_Netconfigentry.
    Netconfigentry []RMON2MIB_Netconfigtable_Netconfigentry
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetEntityData() *types.CommonEntityData {
    netconfigtable.EntityData.YFilter = netconfigtable.YFilter
    netconfigtable.EntityData.YangName = "netConfigTable"
    netconfigtable.EntityData.BundleName = "cisco_ios_xe"
    netconfigtable.EntityData.ParentYangName = "RMON2-MIB"
    netconfigtable.EntityData.SegmentPath = "netConfigTable"
    netconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netconfigtable.EntityData.Children = make(map[string]types.YChild)
    netconfigtable.EntityData.Children["netConfigEntry"] = types.YChild{"Netconfigentry", nil}
    for i := range netconfigtable.Netconfigentry {
        netconfigtable.EntityData.Children[types.GetSegmentPath(&netconfigtable.Netconfigentry[i])] = types.YChild{"Netconfigentry", &netconfigtable.Netconfigentry[i]}
    }
    netconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netconfigtable.EntityData)
}

// RMON2MIB_Netconfigtable_Netconfigentry
// A set of configuration parameters for a particular
// network interface on this device. If the device has no network
// interface, this table is empty.
// 
// The index is composed of the ifIndex assigned to the
// corresponding interface.
type RMON2MIB_Netconfigtable_Netconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range:
    // -2147483648..2147483647. Refers to
    // rfc1213_mib.RFC1213MIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The IP address of this Net interface.  The default value for this object is
    // 0.0.0.0.  If either the netConfigIPAddress or netConfigSubnetMask are
    // 0.0.0.0, then when the device boots, it may use BOOTP to try to figure out
    // what these values should be. If BOOTP fails, before the device can talk on
    // the network, this value must be configured (e.g., through a terminal
    // attached to the device). If BOOTP is      used, care should be taken to not
    // send BOOTP broadcasts too frequently and to eventually send very
    // infrequently if no replies are received. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Netconfigipaddress interface{}

    // The subnet mask of this Net interface.  The default value for this object
    // is 0.0.0.0.  If either the netConfigIPAddress or netConfigSubnetMask are
    // 0.0.0.0, then when the device boots, it may use BOOTP to try to figure out
    // what these values should be. If BOOTP fails, before the device can talk on
    // the network, this value must be configured (e.g., through a terminal
    // attached to the device). If BOOTP is used, care should be taken to not send
    // BOOTP broadcasts too frequently and to eventually send very infrequently if
    // no replies are received. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Netconfigsubnetmask interface{}

    // The status of this netConfigEntry.  An entry may not exist in the active
    // state unless all objects in the entry have an appropriate value. The type
    // is RowStatus.
    Netconfigstatus interface{}
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetEntityData() *types.CommonEntityData {
    netconfigentry.EntityData.YFilter = netconfigentry.YFilter
    netconfigentry.EntityData.YangName = "netConfigEntry"
    netconfigentry.EntityData.BundleName = "cisco_ios_xe"
    netconfigentry.EntityData.ParentYangName = "netConfigTable"
    netconfigentry.EntityData.SegmentPath = "netConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", netconfigentry.Ifindex) + "']"
    netconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netconfigentry.EntityData.Children = make(map[string]types.YChild)
    netconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    netconfigentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", netconfigentry.Ifindex}
    netconfigentry.EntityData.Leafs["netConfigIPAddress"] = types.YLeaf{"Netconfigipaddress", netconfigentry.Netconfigipaddress}
    netconfigentry.EntityData.Leafs["netConfigSubnetMask"] = types.YLeaf{"Netconfigsubnetmask", netconfigentry.Netconfigsubnetmask}
    netconfigentry.EntityData.Leafs["netConfigStatus"] = types.YLeaf{"Netconfigstatus", netconfigentry.Netconfigstatus}
    return &(netconfigentry.EntityData)
}

// RMON2MIB_Trapdesttable
// A list of trap destination entries.
type RMON2MIB_Trapdesttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry includes a destination IP address to which to send traps for
    // this community. The type is slice of RMON2MIB_Trapdesttable_Trapdestentry.
    Trapdestentry []RMON2MIB_Trapdesttable_Trapdestentry
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetEntityData() *types.CommonEntityData {
    trapdesttable.EntityData.YFilter = trapdesttable.YFilter
    trapdesttable.EntityData.YangName = "trapDestTable"
    trapdesttable.EntityData.BundleName = "cisco_ios_xe"
    trapdesttable.EntityData.ParentYangName = "RMON2-MIB"
    trapdesttable.EntityData.SegmentPath = "trapDestTable"
    trapdesttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trapdesttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trapdesttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trapdesttable.EntityData.Children = make(map[string]types.YChild)
    trapdesttable.EntityData.Children["trapDestEntry"] = types.YChild{"Trapdestentry", nil}
    for i := range trapdesttable.Trapdestentry {
        trapdesttable.EntityData.Children[types.GetSegmentPath(&trapdesttable.Trapdestentry[i])] = types.YChild{"Trapdestentry", &trapdesttable.Trapdestentry[i]}
    }
    trapdesttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapdesttable.EntityData)
}

// RMON2MIB_Trapdesttable_Trapdestentry
// This entry includes a destination IP address to which to send
// traps for this community.
type RMON2MIB_Trapdesttable_Trapdestentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A value that uniquely identifies this
    // trapDestEntry. The type is interface{} with range: 1..65535.
    Trapdestindex interface{}

    // A community to which this destination address belongs. This entry is
    // associated with any eventEntries in the RMON      MIB whose value of
    // eventCommunity is equal to the value of this object.  Every time an
    // associated event entry sends a trap due to an event, that trap will be sent
    // to each address in the trapDestTable with a trapDestCommunity equal to
    // eventCommunity.  This object may not be modified if the associated
    // trapDestStatus object is equal to active(1). The type is string with
    // length: 0..127.
    Trapdestcommunity interface{}

    // The protocol with which to send this trap. The type is Trapdestprotocol.
    Trapdestprotocol interface{}

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
    Trapdestaddress interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Trapdestowner interface{}

    // The status of this trap destination entry.  An entry may not exist in the
    // active state unless all objects in the entry have an appropriate value. The
    // type is RowStatus.
    Trapdeststatus interface{}
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetEntityData() *types.CommonEntityData {
    trapdestentry.EntityData.YFilter = trapdestentry.YFilter
    trapdestentry.EntityData.YangName = "trapDestEntry"
    trapdestentry.EntityData.BundleName = "cisco_ios_xe"
    trapdestentry.EntityData.ParentYangName = "trapDestTable"
    trapdestentry.EntityData.SegmentPath = "trapDestEntry" + "[trapDestIndex='" + fmt.Sprintf("%v", trapdestentry.Trapdestindex) + "']"
    trapdestentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trapdestentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trapdestentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trapdestentry.EntityData.Children = make(map[string]types.YChild)
    trapdestentry.EntityData.Leafs = make(map[string]types.YLeaf)
    trapdestentry.EntityData.Leafs["trapDestIndex"] = types.YLeaf{"Trapdestindex", trapdestentry.Trapdestindex}
    trapdestentry.EntityData.Leafs["trapDestCommunity"] = types.YLeaf{"Trapdestcommunity", trapdestentry.Trapdestcommunity}
    trapdestentry.EntityData.Leafs["trapDestProtocol"] = types.YLeaf{"Trapdestprotocol", trapdestentry.Trapdestprotocol}
    trapdestentry.EntityData.Leafs["trapDestAddress"] = types.YLeaf{"Trapdestaddress", trapdestentry.Trapdestaddress}
    trapdestentry.EntityData.Leafs["trapDestOwner"] = types.YLeaf{"Trapdestowner", trapdestentry.Trapdestowner}
    trapdestentry.EntityData.Leafs["trapDestStatus"] = types.YLeaf{"Trapdeststatus", trapdestentry.Trapdeststatus}
    return &(trapdestentry.EntityData)
}

// RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol represents The protocol with which to send this trap.
type RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol string

const (
    RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol_ip RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol = "ip"

    RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol_ipx RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol = "ipx"
)

// RMON2MIB_Serialconnectiontable
// A list of serialConnectionEntries.
type RMON2MIB_Serialconnectiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a SLIP link over a serial line. The type is slice of
    // RMON2MIB_Serialconnectiontable_Serialconnectionentry.
    Serialconnectionentry []RMON2MIB_Serialconnectiontable_Serialconnectionentry
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetEntityData() *types.CommonEntityData {
    serialconnectiontable.EntityData.YFilter = serialconnectiontable.YFilter
    serialconnectiontable.EntityData.YangName = "serialConnectionTable"
    serialconnectiontable.EntityData.BundleName = "cisco_ios_xe"
    serialconnectiontable.EntityData.ParentYangName = "RMON2-MIB"
    serialconnectiontable.EntityData.SegmentPath = "serialConnectionTable"
    serialconnectiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialconnectiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialconnectiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialconnectiontable.EntityData.Children = make(map[string]types.YChild)
    serialconnectiontable.EntityData.Children["serialConnectionEntry"] = types.YChild{"Serialconnectionentry", nil}
    for i := range serialconnectiontable.Serialconnectionentry {
        serialconnectiontable.EntityData.Children[types.GetSegmentPath(&serialconnectiontable.Serialconnectionentry[i])] = types.YChild{"Serialconnectionentry", &serialconnectiontable.Serialconnectionentry[i]}
    }
    serialconnectiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serialconnectiontable.EntityData)
}

// RMON2MIB_Serialconnectiontable_Serialconnectionentry
// Configuration for a SLIP link over a serial line.
type RMON2MIB_Serialconnectiontable_Serialconnectionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A value that uniquely identifies this
    // serialConnection entry. The type is interface{} with range: 1..65535.
    Serialconnectindex interface{}

    // The IP Address that can be reached at the other end of this serial
    // connection. This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Serialconnectdestipaddress interface{}

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
    // Serialconnecttype.
    Serialconnecttype interface{}

    // A control string which specifies how to dial the phone number in order to
    // establish a modem connection.  The string should include dialing prefix and
    // suffix.  For example: ``^s^MATD9,888-1234^M'' will instruct the Probe to
    // send a carriage return followed by the dialing prefix ``ATD'', the phone
    // number ``9,888-1234'', and a carriage return as the dialing suffix. This
    // object may not be modified if the associated serialConnectStatus object is
    // equal to active(1). The type is string with length: 0..255.
    Serialconnectdialstring interface{}

    // A control string which specifies how to establish a data switch connection.
    // This object may not be modified if the associated serialConnectStatus
    // object is equal to active(1). The type is string with length: 0..255.
    Serialconnectswitchconnectseq interface{}

    // A control string which specifies how to terminate a data switch connection.
    // This object may not be modified if the associated      serialConnectStatus
    // object is equal to active(1). The type is string with length: 0..255.
    Serialconnectswitchdisconnectseq interface{}

    // A control string which specifies how to reset a data switch in the event of
    // a timeout. This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is string with
    // length: 0..255.
    Serialconnectswitchresetseq interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Serialconnectowner interface{}

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
    Serialconnectstatus interface{}
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetEntityData() *types.CommonEntityData {
    serialconnectionentry.EntityData.YFilter = serialconnectionentry.YFilter
    serialconnectionentry.EntityData.YangName = "serialConnectionEntry"
    serialconnectionentry.EntityData.BundleName = "cisco_ios_xe"
    serialconnectionentry.EntityData.ParentYangName = "serialConnectionTable"
    serialconnectionentry.EntityData.SegmentPath = "serialConnectionEntry" + "[serialConnectIndex='" + fmt.Sprintf("%v", serialconnectionentry.Serialconnectindex) + "']"
    serialconnectionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialconnectionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialconnectionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialconnectionentry.EntityData.Children = make(map[string]types.YChild)
    serialconnectionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    serialconnectionentry.EntityData.Leafs["serialConnectIndex"] = types.YLeaf{"Serialconnectindex", serialconnectionentry.Serialconnectindex}
    serialconnectionentry.EntityData.Leafs["serialConnectDestIpAddress"] = types.YLeaf{"Serialconnectdestipaddress", serialconnectionentry.Serialconnectdestipaddress}
    serialconnectionentry.EntityData.Leafs["serialConnectType"] = types.YLeaf{"Serialconnecttype", serialconnectionentry.Serialconnecttype}
    serialconnectionentry.EntityData.Leafs["serialConnectDialString"] = types.YLeaf{"Serialconnectdialstring", serialconnectionentry.Serialconnectdialstring}
    serialconnectionentry.EntityData.Leafs["serialConnectSwitchConnectSeq"] = types.YLeaf{"Serialconnectswitchconnectseq", serialconnectionentry.Serialconnectswitchconnectseq}
    serialconnectionentry.EntityData.Leafs["serialConnectSwitchDisconnectSeq"] = types.YLeaf{"Serialconnectswitchdisconnectseq", serialconnectionentry.Serialconnectswitchdisconnectseq}
    serialconnectionentry.EntityData.Leafs["serialConnectSwitchResetSeq"] = types.YLeaf{"Serialconnectswitchresetseq", serialconnectionentry.Serialconnectswitchresetseq}
    serialconnectionentry.EntityData.Leafs["serialConnectOwner"] = types.YLeaf{"Serialconnectowner", serialconnectionentry.Serialconnectowner}
    serialconnectionentry.EntityData.Leafs["serialConnectStatus"] = types.YLeaf{"Serialconnectstatus", serialconnectionentry.Serialconnectstatus}
    return &(serialconnectionentry.EntityData)
}

// RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype represents serialConnectStatus object is equal to active(1).
type RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype string

const (
    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_direct RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "direct"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_modem RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "modem"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_switch_ RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "switch"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_modemSwitch RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "modemSwitch"
)

