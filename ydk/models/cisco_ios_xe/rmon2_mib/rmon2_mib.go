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
    parent types.Entity
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

func (rMON2MIB *RMON2MIB) GetFilter() yfilter.YFilter { return rMON2MIB.YFilter }

func (rMON2MIB *RMON2MIB) SetFilter(yf yfilter.YFilter) { rMON2MIB.YFilter = yf }

func (rMON2MIB *RMON2MIB) GetGoName(yname string) string {
    if yname == "protocolDir" { return "Protocoldir" }
    if yname == "addressMap" { return "Addressmap" }
    if yname == "probeConfig" { return "Probeconfig" }
    if yname == "protocolDirTable" { return "Protocoldirtable" }
    if yname == "protocolDistControlTable" { return "Protocoldistcontroltable" }
    if yname == "protocolDistStatsTable" { return "Protocoldiststatstable" }
    if yname == "addressMapControlTable" { return "Addressmapcontroltable" }
    if yname == "addressMapTable" { return "Addressmaptable" }
    if yname == "hlHostControlTable" { return "Hlhostcontroltable" }
    if yname == "nlHostTable" { return "Nlhosttable" }
    if yname == "hlMatrixControlTable" { return "Hlmatrixcontroltable" }
    if yname == "nlMatrixSDTable" { return "Nlmatrixsdtable" }
    if yname == "nlMatrixDSTable" { return "Nlmatrixdstable" }
    if yname == "nlMatrixTopNControlTable" { return "Nlmatrixtopncontroltable" }
    if yname == "nlMatrixTopNTable" { return "Nlmatrixtopntable" }
    if yname == "alHostTable" { return "Alhosttable" }
    if yname == "alMatrixSDTable" { return "Almatrixsdtable" }
    if yname == "alMatrixDSTable" { return "Almatrixdstable" }
    if yname == "alMatrixTopNControlTable" { return "Almatrixtopncontroltable" }
    if yname == "alMatrixTopNTable" { return "Almatrixtopntable" }
    if yname == "usrHistoryControlTable" { return "Usrhistorycontroltable" }
    if yname == "usrHistoryObjectTable" { return "Usrhistoryobjecttable" }
    if yname == "usrHistoryTable" { return "Usrhistorytable" }
    if yname == "serialConfigTable" { return "Serialconfigtable" }
    if yname == "netConfigTable" { return "Netconfigtable" }
    if yname == "trapDestTable" { return "Trapdesttable" }
    if yname == "serialConnectionTable" { return "Serialconnectiontable" }
    return ""
}

func (rMON2MIB *RMON2MIB) GetSegmentPath() string {
    return "RMON2-MIB:RMON2-MIB"
}

func (rMON2MIB *RMON2MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocolDir" {
        return &rMON2MIB.Protocoldir
    }
    if childYangName == "addressMap" {
        return &rMON2MIB.Addressmap
    }
    if childYangName == "probeConfig" {
        return &rMON2MIB.Probeconfig
    }
    if childYangName == "protocolDirTable" {
        return &rMON2MIB.Protocoldirtable
    }
    if childYangName == "protocolDistControlTable" {
        return &rMON2MIB.Protocoldistcontroltable
    }
    if childYangName == "protocolDistStatsTable" {
        return &rMON2MIB.Protocoldiststatstable
    }
    if childYangName == "addressMapControlTable" {
        return &rMON2MIB.Addressmapcontroltable
    }
    if childYangName == "addressMapTable" {
        return &rMON2MIB.Addressmaptable
    }
    if childYangName == "hlHostControlTable" {
        return &rMON2MIB.Hlhostcontroltable
    }
    if childYangName == "nlHostTable" {
        return &rMON2MIB.Nlhosttable
    }
    if childYangName == "hlMatrixControlTable" {
        return &rMON2MIB.Hlmatrixcontroltable
    }
    if childYangName == "nlMatrixSDTable" {
        return &rMON2MIB.Nlmatrixsdtable
    }
    if childYangName == "nlMatrixDSTable" {
        return &rMON2MIB.Nlmatrixdstable
    }
    if childYangName == "nlMatrixTopNControlTable" {
        return &rMON2MIB.Nlmatrixtopncontroltable
    }
    if childYangName == "nlMatrixTopNTable" {
        return &rMON2MIB.Nlmatrixtopntable
    }
    if childYangName == "alHostTable" {
        return &rMON2MIB.Alhosttable
    }
    if childYangName == "alMatrixSDTable" {
        return &rMON2MIB.Almatrixsdtable
    }
    if childYangName == "alMatrixDSTable" {
        return &rMON2MIB.Almatrixdstable
    }
    if childYangName == "alMatrixTopNControlTable" {
        return &rMON2MIB.Almatrixtopncontroltable
    }
    if childYangName == "alMatrixTopNTable" {
        return &rMON2MIB.Almatrixtopntable
    }
    if childYangName == "usrHistoryControlTable" {
        return &rMON2MIB.Usrhistorycontroltable
    }
    if childYangName == "usrHistoryObjectTable" {
        return &rMON2MIB.Usrhistoryobjecttable
    }
    if childYangName == "usrHistoryTable" {
        return &rMON2MIB.Usrhistorytable
    }
    if childYangName == "serialConfigTable" {
        return &rMON2MIB.Serialconfigtable
    }
    if childYangName == "netConfigTable" {
        return &rMON2MIB.Netconfigtable
    }
    if childYangName == "trapDestTable" {
        return &rMON2MIB.Trapdesttable
    }
    if childYangName == "serialConnectionTable" {
        return &rMON2MIB.Serialconnectiontable
    }
    return nil
}

func (rMON2MIB *RMON2MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocolDir"] = &rMON2MIB.Protocoldir
    children["addressMap"] = &rMON2MIB.Addressmap
    children["probeConfig"] = &rMON2MIB.Probeconfig
    children["protocolDirTable"] = &rMON2MIB.Protocoldirtable
    children["protocolDistControlTable"] = &rMON2MIB.Protocoldistcontroltable
    children["protocolDistStatsTable"] = &rMON2MIB.Protocoldiststatstable
    children["addressMapControlTable"] = &rMON2MIB.Addressmapcontroltable
    children["addressMapTable"] = &rMON2MIB.Addressmaptable
    children["hlHostControlTable"] = &rMON2MIB.Hlhostcontroltable
    children["nlHostTable"] = &rMON2MIB.Nlhosttable
    children["hlMatrixControlTable"] = &rMON2MIB.Hlmatrixcontroltable
    children["nlMatrixSDTable"] = &rMON2MIB.Nlmatrixsdtable
    children["nlMatrixDSTable"] = &rMON2MIB.Nlmatrixdstable
    children["nlMatrixTopNControlTable"] = &rMON2MIB.Nlmatrixtopncontroltable
    children["nlMatrixTopNTable"] = &rMON2MIB.Nlmatrixtopntable
    children["alHostTable"] = &rMON2MIB.Alhosttable
    children["alMatrixSDTable"] = &rMON2MIB.Almatrixsdtable
    children["alMatrixDSTable"] = &rMON2MIB.Almatrixdstable
    children["alMatrixTopNControlTable"] = &rMON2MIB.Almatrixtopncontroltable
    children["alMatrixTopNTable"] = &rMON2MIB.Almatrixtopntable
    children["usrHistoryControlTable"] = &rMON2MIB.Usrhistorycontroltable
    children["usrHistoryObjectTable"] = &rMON2MIB.Usrhistoryobjecttable
    children["usrHistoryTable"] = &rMON2MIB.Usrhistorytable
    children["serialConfigTable"] = &rMON2MIB.Serialconfigtable
    children["netConfigTable"] = &rMON2MIB.Netconfigtable
    children["trapDestTable"] = &rMON2MIB.Trapdesttable
    children["serialConnectionTable"] = &rMON2MIB.Serialconnectiontable
    return children
}

func (rMON2MIB *RMON2MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rMON2MIB *RMON2MIB) GetBundleName() string { return "cisco_ios_xe" }

func (rMON2MIB *RMON2MIB) GetYangName() string { return "RMON2-MIB" }

func (rMON2MIB *RMON2MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rMON2MIB *RMON2MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rMON2MIB *RMON2MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rMON2MIB *RMON2MIB) SetParent(parent types.Entity) { rMON2MIB.parent = parent }

func (rMON2MIB *RMON2MIB) GetParent() types.Entity { return rMON2MIB.parent }

func (rMON2MIB *RMON2MIB) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Protocoldir
type RMON2MIB_Protocoldir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time the protocol directory was last
    // modified, either through insertions or deletions, or through modifications
    // of either the protocolDirAddressMapConfig, protocolDirHostConfig, or
    // protocolDirMatrixConfig. The type is interface{} with range: 0..4294967295.
    Protocoldirlastchange interface{}
}

func (protocoldir *RMON2MIB_Protocoldir) GetFilter() yfilter.YFilter { return protocoldir.YFilter }

func (protocoldir *RMON2MIB_Protocoldir) SetFilter(yf yfilter.YFilter) { protocoldir.YFilter = yf }

func (protocoldir *RMON2MIB_Protocoldir) GetGoName(yname string) string {
    if yname == "protocolDirLastChange" { return "Protocoldirlastchange" }
    return ""
}

func (protocoldir *RMON2MIB_Protocoldir) GetSegmentPath() string {
    return "protocolDir"
}

func (protocoldir *RMON2MIB_Protocoldir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocoldir *RMON2MIB_Protocoldir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocoldir *RMON2MIB_Protocoldir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocolDirLastChange"] = protocoldir.Protocoldirlastchange
    return leafs
}

func (protocoldir *RMON2MIB_Protocoldir) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldir *RMON2MIB_Protocoldir) GetYangName() string { return "protocolDir" }

func (protocoldir *RMON2MIB_Protocoldir) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldir *RMON2MIB_Protocoldir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldir *RMON2MIB_Protocoldir) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldir *RMON2MIB_Protocoldir) SetParent(parent types.Entity) { protocoldir.parent = parent }

func (protocoldir *RMON2MIB_Protocoldir) GetParent() types.Entity { return protocoldir.parent }

func (protocoldir *RMON2MIB_Protocoldir) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Addressmap
type RMON2MIB_Addressmap struct {
    parent types.Entity
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

func (addressmap *RMON2MIB_Addressmap) GetFilter() yfilter.YFilter { return addressmap.YFilter }

func (addressmap *RMON2MIB_Addressmap) SetFilter(yf yfilter.YFilter) { addressmap.YFilter = yf }

func (addressmap *RMON2MIB_Addressmap) GetGoName(yname string) string {
    if yname == "addressMapInserts" { return "Addressmapinserts" }
    if yname == "addressMapDeletes" { return "Addressmapdeletes" }
    if yname == "addressMapMaxDesiredEntries" { return "Addressmapmaxdesiredentries" }
    return ""
}

func (addressmap *RMON2MIB_Addressmap) GetSegmentPath() string {
    return "addressMap"
}

func (addressmap *RMON2MIB_Addressmap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressmap *RMON2MIB_Addressmap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressmap *RMON2MIB_Addressmap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addressMapInserts"] = addressmap.Addressmapinserts
    leafs["addressMapDeletes"] = addressmap.Addressmapdeletes
    leafs["addressMapMaxDesiredEntries"] = addressmap.Addressmapmaxdesiredentries
    return leafs
}

func (addressmap *RMON2MIB_Addressmap) GetBundleName() string { return "cisco_ios_xe" }

func (addressmap *RMON2MIB_Addressmap) GetYangName() string { return "addressMap" }

func (addressmap *RMON2MIB_Addressmap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressmap *RMON2MIB_Addressmap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressmap *RMON2MIB_Addressmap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressmap *RMON2MIB_Addressmap) SetParent(parent types.Entity) { addressmap.parent = parent }

func (addressmap *RMON2MIB_Addressmap) GetParent() types.Entity { return addressmap.parent }

func (addressmap *RMON2MIB_Addressmap) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Probeconfig
type RMON2MIB_Probeconfig struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netdefaultgateway interface{}
}

func (probeconfig *RMON2MIB_Probeconfig) GetFilter() yfilter.YFilter { return probeconfig.YFilter }

func (probeconfig *RMON2MIB_Probeconfig) SetFilter(yf yfilter.YFilter) { probeconfig.YFilter = yf }

func (probeconfig *RMON2MIB_Probeconfig) GetGoName(yname string) string {
    if yname == "probeCapabilities" { return "Probecapabilities" }
    if yname == "probeSoftwareRev" { return "Probesoftwarerev" }
    if yname == "probeHardwareRev" { return "Probehardwarerev" }
    if yname == "probeDateTime" { return "Probedatetime" }
    if yname == "probeResetControl" { return "Proberesetcontrol" }
    if yname == "probeDownloadFile" { return "Probedownloadfile" }
    if yname == "probeDownloadTFTPServer" { return "Probedownloadtftpserver" }
    if yname == "probeDownloadAction" { return "Probedownloadaction" }
    if yname == "probeDownloadStatus" { return "Probedownloadstatus" }
    if yname == "netDefaultGateway" { return "Netdefaultgateway" }
    return ""
}

func (probeconfig *RMON2MIB_Probeconfig) GetSegmentPath() string {
    return "probeConfig"
}

func (probeconfig *RMON2MIB_Probeconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (probeconfig *RMON2MIB_Probeconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (probeconfig *RMON2MIB_Probeconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probeCapabilities"] = probeconfig.Probecapabilities
    leafs["probeSoftwareRev"] = probeconfig.Probesoftwarerev
    leafs["probeHardwareRev"] = probeconfig.Probehardwarerev
    leafs["probeDateTime"] = probeconfig.Probedatetime
    leafs["probeResetControl"] = probeconfig.Proberesetcontrol
    leafs["probeDownloadFile"] = probeconfig.Probedownloadfile
    leafs["probeDownloadTFTPServer"] = probeconfig.Probedownloadtftpserver
    leafs["probeDownloadAction"] = probeconfig.Probedownloadaction
    leafs["probeDownloadStatus"] = probeconfig.Probedownloadstatus
    leafs["netDefaultGateway"] = probeconfig.Netdefaultgateway
    return leafs
}

func (probeconfig *RMON2MIB_Probeconfig) GetBundleName() string { return "cisco_ios_xe" }

func (probeconfig *RMON2MIB_Probeconfig) GetYangName() string { return "probeConfig" }

func (probeconfig *RMON2MIB_Probeconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (probeconfig *RMON2MIB_Probeconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (probeconfig *RMON2MIB_Probeconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (probeconfig *RMON2MIB_Probeconfig) SetParent(parent types.Entity) { probeconfig.parent = parent }

func (probeconfig *RMON2MIB_Probeconfig) GetParent() types.Entity { return probeconfig.parent }

func (probeconfig *RMON2MIB_Probeconfig) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDirTable.  An example of the indexing of
    // this entry is protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
    // encoding of a length of 8, followed by 8 subids encoding the protocolDirID
    // of 1.2048, followed by a length of 2 and the 2 subids encoding zero-valued
    // parameters. The type is slice of
    // RMON2MIB_Protocoldirtable_Protocoldirentry.
    Protocoldirentry []RMON2MIB_Protocoldirtable_Protocoldirentry
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetFilter() yfilter.YFilter { return protocoldirtable.YFilter }

func (protocoldirtable *RMON2MIB_Protocoldirtable) SetFilter(yf yfilter.YFilter) { protocoldirtable.YFilter = yf }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetGoName(yname string) string {
    if yname == "protocolDirEntry" { return "Protocoldirentry" }
    return ""
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetSegmentPath() string {
    return "protocolDirTable"
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocolDirEntry" {
        for _, c := range protocoldirtable.Protocoldirentry {
            if protocoldirtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Protocoldirtable_Protocoldirentry{}
        protocoldirtable.Protocoldirentry = append(protocoldirtable.Protocoldirentry, child)
        return &protocoldirtable.Protocoldirentry[len(protocoldirtable.Protocoldirentry)-1]
    }
    return nil
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocoldirtable.Protocoldirentry {
        children[protocoldirtable.Protocoldirentry[i].GetSegmentPath()] = &protocoldirtable.Protocoldirentry[i]
    }
    return children
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetYangName() string { return "protocolDirTable" }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldirtable *RMON2MIB_Protocoldirtable) SetParent(parent types.Entity) { protocoldirtable.parent = parent }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetParent() types.Entity { return protocoldirtable.parent }

func (protocoldirtable *RMON2MIB_Protocoldirtable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Protocoldirtable_Protocoldirentry
// A conceptual row in the protocolDirTable.
// 
// An example of the indexing of this entry is
// protocolDirLocalIndex.8.0.0.0.1.0.0.8.0.2.0.0, which is the
// encoding of a length of 8, followed by 8 subids encoding the
// protocolDirID of 1.2048, followed by a length of 2 and the
// 2 subids encoding zero-valued parameters.
type RMON2MIB_Protocoldirtable_Protocoldirentry struct {
    parent types.Entity
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

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetFilter() yfilter.YFilter { return protocoldirentry.YFilter }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) SetFilter(yf yfilter.YFilter) { protocoldirentry.YFilter = yf }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetGoName(yname string) string {
    if yname == "protocolDirID" { return "Protocoldirid" }
    if yname == "protocolDirParameters" { return "Protocoldirparameters" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "protocolDirDescr" { return "Protocoldirdescr" }
    if yname == "protocolDirType" { return "Protocoldirtype" }
    if yname == "protocolDirAddressMapConfig" { return "Protocoldiraddressmapconfig" }
    if yname == "protocolDirHostConfig" { return "Protocoldirhostconfig" }
    if yname == "protocolDirMatrixConfig" { return "Protocoldirmatrixconfig" }
    if yname == "protocolDirOwner" { return "Protocoldirowner" }
    if yname == "protocolDirStatus" { return "Protocoldirstatus" }
    return ""
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetSegmentPath() string {
    return "protocolDirEntry" + "[protocolDirID='" + fmt.Sprintf("%v", protocoldirentry.Protocoldirid) + "']" + "[protocolDirParameters='" + fmt.Sprintf("%v", protocoldirentry.Protocoldirparameters) + "']"
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocolDirID"] = protocoldirentry.Protocoldirid
    leafs["protocolDirParameters"] = protocoldirentry.Protocoldirparameters
    leafs["protocolDirLocalIndex"] = protocoldirentry.Protocoldirlocalindex
    leafs["protocolDirDescr"] = protocoldirentry.Protocoldirdescr
    leafs["protocolDirType"] = protocoldirentry.Protocoldirtype
    leafs["protocolDirAddressMapConfig"] = protocoldirentry.Protocoldiraddressmapconfig
    leafs["protocolDirHostConfig"] = protocoldirentry.Protocoldirhostconfig
    leafs["protocolDirMatrixConfig"] = protocoldirentry.Protocoldirmatrixconfig
    leafs["protocolDirOwner"] = protocoldirentry.Protocoldirowner
    leafs["protocolDirStatus"] = protocoldirentry.Protocoldirstatus
    return leafs
}

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetYangName() string { return "protocolDirEntry" }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) SetParent(parent types.Entity) { protocoldirentry.parent = parent }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetParent() types.Entity { return protocoldirentry.parent }

func (protocoldirentry *RMON2MIB_Protocoldirtable_Protocoldirentry) GetParentYangName() string { return "protocolDirTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the protocolDistControlTable.  An example of the
    // indexing of this entry is      protocolDistControlDroppedFrames.7. The type
    // is slice of RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry.
    Protocoldistcontrolentry []RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetFilter() yfilter.YFilter { return protocoldistcontroltable.YFilter }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) SetFilter(yf yfilter.YFilter) { protocoldistcontroltable.YFilter = yf }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetGoName(yname string) string {
    if yname == "protocolDistControlEntry" { return "Protocoldistcontrolentry" }
    return ""
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetSegmentPath() string {
    return "protocolDistControlTable"
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocolDistControlEntry" {
        for _, c := range protocoldistcontroltable.Protocoldistcontrolentry {
            if protocoldistcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry{}
        protocoldistcontroltable.Protocoldistcontrolentry = append(protocoldistcontroltable.Protocoldistcontrolentry, child)
        return &protocoldistcontroltable.Protocoldistcontrolentry[len(protocoldistcontroltable.Protocoldistcontrolentry)-1]
    }
    return nil
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocoldistcontroltable.Protocoldistcontrolentry {
        children[protocoldistcontroltable.Protocoldistcontrolentry[i].GetSegmentPath()] = &protocoldistcontroltable.Protocoldistcontrolentry[i]
    }
    return children
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetYangName() string { return "protocolDistControlTable" }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) SetParent(parent types.Entity) { protocoldistcontroltable.parent = parent }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetParent() types.Entity { return protocoldistcontroltable.parent }

func (protocoldistcontroltable *RMON2MIB_Protocoldistcontroltable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique index for this protocolDistControlEntry.
    // The type is interface{} with range: 1..65535.
    Protocoldistcontrolindex interface{}

    // The source of data for the this protocol distribution.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  This object may not be modified if the associated
    // protocolDistControlStatus object is equal to active(1). The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetFilter() yfilter.YFilter { return protocoldistcontrolentry.YFilter }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) SetFilter(yf yfilter.YFilter) { protocoldistcontrolentry.YFilter = yf }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetGoName(yname string) string {
    if yname == "protocolDistControlIndex" { return "Protocoldistcontrolindex" }
    if yname == "protocolDistControlDataSource" { return "Protocoldistcontroldatasource" }
    if yname == "protocolDistControlDroppedFrames" { return "Protocoldistcontroldroppedframes" }
    if yname == "protocolDistControlCreateTime" { return "Protocoldistcontrolcreatetime" }
    if yname == "protocolDistControlOwner" { return "Protocoldistcontrolowner" }
    if yname == "protocolDistControlStatus" { return "Protocoldistcontrolstatus" }
    return ""
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetSegmentPath() string {
    return "protocolDistControlEntry" + "[protocolDistControlIndex='" + fmt.Sprintf("%v", protocoldistcontrolentry.Protocoldistcontrolindex) + "']"
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocolDistControlIndex"] = protocoldistcontrolentry.Protocoldistcontrolindex
    leafs["protocolDistControlDataSource"] = protocoldistcontrolentry.Protocoldistcontroldatasource
    leafs["protocolDistControlDroppedFrames"] = protocoldistcontrolentry.Protocoldistcontroldroppedframes
    leafs["protocolDistControlCreateTime"] = protocoldistcontrolentry.Protocoldistcontrolcreatetime
    leafs["protocolDistControlOwner"] = protocoldistcontrolentry.Protocoldistcontrolowner
    leafs["protocolDistControlStatus"] = protocoldistcontrolentry.Protocoldistcontrolstatus
    return leafs
}

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetYangName() string { return "protocolDistControlEntry" }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) SetParent(parent types.Entity) { protocoldistcontrolentry.parent = parent }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetParent() types.Entity { return protocoldistcontrolentry.parent }

func (protocoldistcontrolentry *RMON2MIB_Protocoldistcontroltable_Protocoldistcontrolentry) GetParentYangName() string { return "protocolDistControlTable" }

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
    parent types.Entity
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

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetFilter() yfilter.YFilter { return protocoldiststatstable.YFilter }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) SetFilter(yf yfilter.YFilter) { protocoldiststatstable.YFilter = yf }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetGoName(yname string) string {
    if yname == "protocolDistStatsEntry" { return "Protocoldiststatsentry" }
    return ""
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetSegmentPath() string {
    return "protocolDistStatsTable"
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocolDistStatsEntry" {
        for _, c := range protocoldiststatstable.Protocoldiststatsentry {
            if protocoldiststatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry{}
        protocoldiststatstable.Protocoldiststatsentry = append(protocoldiststatstable.Protocoldiststatsentry, child)
        return &protocoldiststatstable.Protocoldiststatsentry[len(protocoldiststatstable.Protocoldiststatsentry)-1]
    }
    return nil
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocoldiststatstable.Protocoldiststatsentry {
        children[protocoldiststatstable.Protocoldiststatsentry[i].GetSegmentPath()] = &protocoldiststatstable.Protocoldiststatsentry[i]
    }
    return children
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetYangName() string { return "protocolDistStatsTable" }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) SetParent(parent types.Entity) { protocoldiststatstable.parent = parent }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetParent() types.Entity { return protocoldiststatstable.parent }

func (protocoldiststatstable *RMON2MIB_Protocoldiststatstable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetFilter() yfilter.YFilter { return protocoldiststatsentry.YFilter }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) SetFilter(yf yfilter.YFilter) { protocoldiststatsentry.YFilter = yf }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetGoName(yname string) string {
    if yname == "protocolDistControlIndex" { return "Protocoldistcontrolindex" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "protocolDistStatsPkts" { return "Protocoldiststatspkts" }
    if yname == "protocolDistStatsOctets" { return "Protocoldiststatsoctets" }
    return ""
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetSegmentPath() string {
    return "protocolDistStatsEntry" + "[protocolDistControlIndex='" + fmt.Sprintf("%v", protocoldiststatsentry.Protocoldistcontrolindex) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", protocoldiststatsentry.Protocoldirlocalindex) + "']"
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocolDistControlIndex"] = protocoldiststatsentry.Protocoldistcontrolindex
    leafs["protocolDirLocalIndex"] = protocoldiststatsentry.Protocoldirlocalindex
    leafs["protocolDistStatsPkts"] = protocoldiststatsentry.Protocoldiststatspkts
    leafs["protocolDistStatsOctets"] = protocoldiststatsentry.Protocoldiststatsoctets
    return leafs
}

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetYangName() string { return "protocolDistStatsEntry" }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) SetParent(parent types.Entity) { protocoldiststatsentry.parent = parent }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetParent() types.Entity { return protocoldiststatsentry.parent }

func (protocoldiststatsentry *RMON2MIB_Protocoldiststatstable_Protocoldiststatsentry) GetParentYangName() string { return "protocolDistStatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapControlTable.      An example of the
    // indexing of this entry is addressMapControlDroppedFrames.1. The type is
    // slice of RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry.
    Addressmapcontrolentry []RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetFilter() yfilter.YFilter { return addressmapcontroltable.YFilter }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) SetFilter(yf yfilter.YFilter) { addressmapcontroltable.YFilter = yf }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetGoName(yname string) string {
    if yname == "addressMapControlEntry" { return "Addressmapcontrolentry" }
    return ""
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetSegmentPath() string {
    return "addressMapControlTable"
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addressMapControlEntry" {
        for _, c := range addressmapcontroltable.Addressmapcontrolentry {
            if addressmapcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry{}
        addressmapcontroltable.Addressmapcontrolentry = append(addressmapcontroltable.Addressmapcontrolentry, child)
        return &addressmapcontroltable.Addressmapcontrolentry[len(addressmapcontroltable.Addressmapcontrolentry)-1]
    }
    return nil
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressmapcontroltable.Addressmapcontrolentry {
        children[addressmapcontroltable.Addressmapcontrolentry[i].GetSegmentPath()] = &addressmapcontroltable.Addressmapcontrolentry[i]
    }
    return children
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetYangName() string { return "addressMapControlTable" }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) SetParent(parent types.Entity) { addressmapcontroltable.parent = parent }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetParent() types.Entity { return addressmapcontroltable.parent }

func (addressmapcontroltable *RMON2MIB_Addressmapcontroltable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique index for this entry in the
    // addressMapControlTable. The type is interface{} with range: 1..65535.
    Addressmapcontrolindex interface{}

    // The source of data for this addressMapControlEntry. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetFilter() yfilter.YFilter { return addressmapcontrolentry.YFilter }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) SetFilter(yf yfilter.YFilter) { addressmapcontrolentry.YFilter = yf }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetGoName(yname string) string {
    if yname == "addressMapControlIndex" { return "Addressmapcontrolindex" }
    if yname == "addressMapControlDataSource" { return "Addressmapcontroldatasource" }
    if yname == "addressMapControlDroppedFrames" { return "Addressmapcontroldroppedframes" }
    if yname == "addressMapControlOwner" { return "Addressmapcontrolowner" }
    if yname == "addressMapControlStatus" { return "Addressmapcontrolstatus" }
    return ""
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetSegmentPath() string {
    return "addressMapControlEntry" + "[addressMapControlIndex='" + fmt.Sprintf("%v", addressmapcontrolentry.Addressmapcontrolindex) + "']"
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addressMapControlIndex"] = addressmapcontrolentry.Addressmapcontrolindex
    leafs["addressMapControlDataSource"] = addressmapcontrolentry.Addressmapcontroldatasource
    leafs["addressMapControlDroppedFrames"] = addressmapcontrolentry.Addressmapcontroldroppedframes
    leafs["addressMapControlOwner"] = addressmapcontrolentry.Addressmapcontrolowner
    leafs["addressMapControlStatus"] = addressmapcontrolentry.Addressmapcontrolstatus
    return leafs
}

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetYangName() string { return "addressMapControlEntry" }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) SetParent(parent types.Entity) { addressmapcontrolentry.parent = parent }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetParent() types.Entity { return addressmapcontrolentry.parent }

func (addressmapcontrolentry *RMON2MIB_Addressmapcontroltable_Addressmapcontrolentry) GetParentYangName() string { return "addressMapControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the addressMapTable. The protocolDirLocalIndex in the
    // index identifies the network layer protocol of the
    // addressMapNetworkAddress.      An example of the indexing of this entry is
    // addressMapSource.783495.18.4.128.2.6.6.11.1.3.6.1.2.1.2.2.1.1.1. The type
    // is slice of RMON2MIB_Addressmaptable_Addressmapentry.
    Addressmapentry []RMON2MIB_Addressmaptable_Addressmapentry
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetFilter() yfilter.YFilter { return addressmaptable.YFilter }

func (addressmaptable *RMON2MIB_Addressmaptable) SetFilter(yf yfilter.YFilter) { addressmaptable.YFilter = yf }

func (addressmaptable *RMON2MIB_Addressmaptable) GetGoName(yname string) string {
    if yname == "addressMapEntry" { return "Addressmapentry" }
    return ""
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetSegmentPath() string {
    return "addressMapTable"
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addressMapEntry" {
        for _, c := range addressmaptable.Addressmapentry {
            if addressmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Addressmaptable_Addressmapentry{}
        addressmaptable.Addressmapentry = append(addressmaptable.Addressmapentry, child)
        return &addressmaptable.Addressmapentry[len(addressmaptable.Addressmapentry)-1]
    }
    return nil
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressmaptable.Addressmapentry {
        children[addressmaptable.Addressmapentry[i].GetSegmentPath()] = &addressmaptable.Addressmapentry[i]
    }
    return children
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressmaptable *RMON2MIB_Addressmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (addressmaptable *RMON2MIB_Addressmaptable) GetYangName() string { return "addressMapTable" }

func (addressmaptable *RMON2MIB_Addressmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressmaptable *RMON2MIB_Addressmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressmaptable *RMON2MIB_Addressmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressmaptable *RMON2MIB_Addressmaptable) SetParent(parent types.Entity) { addressmaptable.parent = parent }

func (addressmaptable *RMON2MIB_Addressmaptable) GetParent() types.Entity { return addressmaptable.parent }

func (addressmaptable *RMON2MIB_Addressmaptable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetFilter() yfilter.YFilter { return addressmapentry.YFilter }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) SetFilter(yf yfilter.YFilter) { addressmapentry.YFilter = yf }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetGoName(yname string) string {
    if yname == "addressMapTimeMark" { return "Addressmaptimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "addressMapNetworkAddress" { return "Addressmapnetworkaddress" }
    if yname == "addressMapSource" { return "Addressmapsource" }
    if yname == "addressMapPhysicalAddress" { return "Addressmapphysicaladdress" }
    if yname == "addressMapLastChange" { return "Addressmaplastchange" }
    return ""
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetSegmentPath() string {
    return "addressMapEntry" + "[addressMapTimeMark='" + fmt.Sprintf("%v", addressmapentry.Addressmaptimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", addressmapentry.Protocoldirlocalindex) + "']" + "[addressMapNetworkAddress='" + fmt.Sprintf("%v", addressmapentry.Addressmapnetworkaddress) + "']" + "[addressMapSource='" + fmt.Sprintf("%v", addressmapentry.Addressmapsource) + "']"
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addressMapTimeMark"] = addressmapentry.Addressmaptimemark
    leafs["protocolDirLocalIndex"] = addressmapentry.Protocoldirlocalindex
    leafs["addressMapNetworkAddress"] = addressmapentry.Addressmapnetworkaddress
    leafs["addressMapSource"] = addressmapentry.Addressmapsource
    leafs["addressMapPhysicalAddress"] = addressmapentry.Addressmapphysicaladdress
    leafs["addressMapLastChange"] = addressmapentry.Addressmaplastchange
    return leafs
}

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetYangName() string { return "addressMapEntry" }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) SetParent(parent types.Entity) { addressmapentry.parent = parent }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetParent() types.Entity { return addressmapentry.parent }

func (addressmapentry *RMON2MIB_Addressmaptable_Addressmapentry) GetParentYangName() string { return "addressMapTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the hlHostControlTable.  An example of the indexing of
    // this entry is hlHostControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry.
    Hlhostcontrolentry []RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetFilter() yfilter.YFilter { return hlhostcontroltable.YFilter }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) SetFilter(yf yfilter.YFilter) { hlhostcontroltable.YFilter = yf }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetGoName(yname string) string {
    if yname == "hlHostControlEntry" { return "Hlhostcontrolentry" }
    return ""
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetSegmentPath() string {
    return "hlHostControlTable"
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hlHostControlEntry" {
        for _, c := range hlhostcontroltable.Hlhostcontrolentry {
            if hlhostcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry{}
        hlhostcontroltable.Hlhostcontrolentry = append(hlhostcontroltable.Hlhostcontrolentry, child)
        return &hlhostcontroltable.Hlhostcontrolentry[len(hlhostcontroltable.Hlhostcontrolentry)-1]
    }
    return nil
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hlhostcontroltable.Hlhostcontrolentry {
        children[hlhostcontroltable.Hlhostcontrolentry[i].GetSegmentPath()] = &hlhostcontroltable.Hlhostcontrolentry[i]
    }
    return children
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetYangName() string { return "hlHostControlTable" }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) SetParent(parent types.Entity) { hlhostcontroltable.parent = parent }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetParent() types.Entity { return hlhostcontroltable.parent }

func (hlhostcontroltable *RMON2MIB_Hlhostcontroltable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry
// A conceptual row in the hlHostControlTable.
// 
// An example of the indexing of this entry is
// hlHostControlNlDroppedFrames.1
type RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry struct {
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetFilter() yfilter.YFilter { return hlhostcontrolentry.YFilter }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) SetFilter(yf yfilter.YFilter) { hlhostcontrolentry.YFilter = yf }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetGoName(yname string) string {
    if yname == "hlHostControlIndex" { return "Hlhostcontrolindex" }
    if yname == "hlHostControlDataSource" { return "Hlhostcontroldatasource" }
    if yname == "hlHostControlNlDroppedFrames" { return "Hlhostcontrolnldroppedframes" }
    if yname == "hlHostControlNlInserts" { return "Hlhostcontrolnlinserts" }
    if yname == "hlHostControlNlDeletes" { return "Hlhostcontrolnldeletes" }
    if yname == "hlHostControlNlMaxDesiredEntries" { return "Hlhostcontrolnlmaxdesiredentries" }
    if yname == "hlHostControlAlDroppedFrames" { return "Hlhostcontrolaldroppedframes" }
    if yname == "hlHostControlAlInserts" { return "Hlhostcontrolalinserts" }
    if yname == "hlHostControlAlDeletes" { return "Hlhostcontrolaldeletes" }
    if yname == "hlHostControlAlMaxDesiredEntries" { return "Hlhostcontrolalmaxdesiredentries" }
    if yname == "hlHostControlOwner" { return "Hlhostcontrolowner" }
    if yname == "hlHostControlStatus" { return "Hlhostcontrolstatus" }
    return ""
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetSegmentPath() string {
    return "hlHostControlEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", hlhostcontrolentry.Hlhostcontrolindex) + "']"
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlHostControlIndex"] = hlhostcontrolentry.Hlhostcontrolindex
    leafs["hlHostControlDataSource"] = hlhostcontrolentry.Hlhostcontroldatasource
    leafs["hlHostControlNlDroppedFrames"] = hlhostcontrolentry.Hlhostcontrolnldroppedframes
    leafs["hlHostControlNlInserts"] = hlhostcontrolentry.Hlhostcontrolnlinserts
    leafs["hlHostControlNlDeletes"] = hlhostcontrolentry.Hlhostcontrolnldeletes
    leafs["hlHostControlNlMaxDesiredEntries"] = hlhostcontrolentry.Hlhostcontrolnlmaxdesiredentries
    leafs["hlHostControlAlDroppedFrames"] = hlhostcontrolentry.Hlhostcontrolaldroppedframes
    leafs["hlHostControlAlInserts"] = hlhostcontrolentry.Hlhostcontrolalinserts
    leafs["hlHostControlAlDeletes"] = hlhostcontrolentry.Hlhostcontrolaldeletes
    leafs["hlHostControlAlMaxDesiredEntries"] = hlhostcontrolentry.Hlhostcontrolalmaxdesiredentries
    leafs["hlHostControlOwner"] = hlhostcontrolentry.Hlhostcontrolowner
    leafs["hlHostControlStatus"] = hlhostcontrolentry.Hlhostcontrolstatus
    return leafs
}

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetYangName() string { return "hlHostControlEntry" }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) SetParent(parent types.Entity) { hlhostcontrolentry.parent = parent }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetParent() types.Entity { return hlhostcontrolentry.parent }

func (hlhostcontrolentry *RMON2MIB_Hlhostcontroltable_Hlhostcontrolentry) GetParentYangName() string { return "hlHostControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the nlHostTable.  The hlHostControlIndex value in the
    // index identifies the hlHostControlEntry on whose behalf this entry was
    // created. The protocolDirLocalIndex value in the index identifies the
    // network layer protocol of the nlHostAddress.  An example of the indexing of
    // this entry is nlHostOutPkts.1.783495.18.4.128.2.6.6. The type is slice of
    // RMON2MIB_Nlhosttable_Nlhostentry.
    Nlhostentry []RMON2MIB_Nlhosttable_Nlhostentry
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetFilter() yfilter.YFilter { return nlhosttable.YFilter }

func (nlhosttable *RMON2MIB_Nlhosttable) SetFilter(yf yfilter.YFilter) { nlhosttable.YFilter = yf }

func (nlhosttable *RMON2MIB_Nlhosttable) GetGoName(yname string) string {
    if yname == "nlHostEntry" { return "Nlhostentry" }
    return ""
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetSegmentPath() string {
    return "nlHostTable"
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlHostEntry" {
        for _, c := range nlhosttable.Nlhostentry {
            if nlhosttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Nlhosttable_Nlhostentry{}
        nlhosttable.Nlhostentry = append(nlhosttable.Nlhostentry, child)
        return &nlhosttable.Nlhostentry[len(nlhosttable.Nlhostentry)-1]
    }
    return nil
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlhosttable.Nlhostentry {
        children[nlhosttable.Nlhostentry[i].GetSegmentPath()] = &nlhosttable.Nlhostentry[i]
    }
    return children
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlhosttable *RMON2MIB_Nlhosttable) GetBundleName() string { return "cisco_ios_xe" }

func (nlhosttable *RMON2MIB_Nlhosttable) GetYangName() string { return "nlHostTable" }

func (nlhosttable *RMON2MIB_Nlhosttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlhosttable *RMON2MIB_Nlhosttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlhosttable *RMON2MIB_Nlhosttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlhosttable *RMON2MIB_Nlhosttable) SetParent(parent types.Entity) { nlhosttable.parent = parent }

func (nlhosttable *RMON2MIB_Nlhosttable) GetParent() types.Entity { return nlhosttable.parent }

func (nlhosttable *RMON2MIB_Nlhosttable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetFilter() yfilter.YFilter { return nlhostentry.YFilter }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) SetFilter(yf yfilter.YFilter) { nlhostentry.YFilter = yf }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetGoName(yname string) string {
    if yname == "hlHostControlIndex" { return "Hlhostcontrolindex" }
    if yname == "nlHostTimeMark" { return "Nlhosttimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlHostAddress" { return "Nlhostaddress" }
    if yname == "nlHostInPkts" { return "Nlhostinpkts" }
    if yname == "nlHostOutPkts" { return "Nlhostoutpkts" }
    if yname == "nlHostInOctets" { return "Nlhostinoctets" }
    if yname == "nlHostOutOctets" { return "Nlhostoutoctets" }
    if yname == "nlHostOutMacNonUnicastPkts" { return "Nlhostoutmacnonunicastpkts" }
    if yname == "nlHostCreateTime" { return "Nlhostcreatetime" }
    return ""
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetSegmentPath() string {
    return "nlHostEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", nlhostentry.Hlhostcontrolindex) + "']" + "[nlHostTimeMark='" + fmt.Sprintf("%v", nlhostentry.Nlhosttimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlhostentry.Protocoldirlocalindex) + "']" + "[nlHostAddress='" + fmt.Sprintf("%v", nlhostentry.Nlhostaddress) + "']"
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlHostControlIndex"] = nlhostentry.Hlhostcontrolindex
    leafs["nlHostTimeMark"] = nlhostentry.Nlhosttimemark
    leafs["protocolDirLocalIndex"] = nlhostentry.Protocoldirlocalindex
    leafs["nlHostAddress"] = nlhostentry.Nlhostaddress
    leafs["nlHostInPkts"] = nlhostentry.Nlhostinpkts
    leafs["nlHostOutPkts"] = nlhostentry.Nlhostoutpkts
    leafs["nlHostInOctets"] = nlhostentry.Nlhostinoctets
    leafs["nlHostOutOctets"] = nlhostentry.Nlhostoutoctets
    leafs["nlHostOutMacNonUnicastPkts"] = nlhostentry.Nlhostoutmacnonunicastpkts
    leafs["nlHostCreateTime"] = nlhostentry.Nlhostcreatetime
    return leafs
}

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetYangName() string { return "nlHostEntry" }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) SetParent(parent types.Entity) { nlhostentry.parent = parent }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetParent() types.Entity { return nlhostentry.parent }

func (nlhostentry *RMON2MIB_Nlhosttable_Nlhostentry) GetParentYangName() string { return "nlHostTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the hlMatrixControlTable.  An example of indexing of
    // this entry is hlMatrixControlNlDroppedFrames.1. The type is slice of
    // RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry.
    Hlmatrixcontrolentry []RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetFilter() yfilter.YFilter { return hlmatrixcontroltable.YFilter }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) SetFilter(yf yfilter.YFilter) { hlmatrixcontroltable.YFilter = yf }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetGoName(yname string) string {
    if yname == "hlMatrixControlEntry" { return "Hlmatrixcontrolentry" }
    return ""
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetSegmentPath() string {
    return "hlMatrixControlTable"
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hlMatrixControlEntry" {
        for _, c := range hlmatrixcontroltable.Hlmatrixcontrolentry {
            if hlmatrixcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry{}
        hlmatrixcontroltable.Hlmatrixcontrolentry = append(hlmatrixcontroltable.Hlmatrixcontrolentry, child)
        return &hlmatrixcontroltable.Hlmatrixcontrolentry[len(hlmatrixcontroltable.Hlmatrixcontrolentry)-1]
    }
    return nil
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hlmatrixcontroltable.Hlmatrixcontrolentry {
        children[hlmatrixcontroltable.Hlmatrixcontrolentry[i].GetSegmentPath()] = &hlmatrixcontroltable.Hlmatrixcontrolentry[i]
    }
    return children
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetYangName() string { return "hlMatrixControlTable" }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) SetParent(parent types.Entity) { hlmatrixcontroltable.parent = parent }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetParent() types.Entity { return hlmatrixcontroltable.parent }

func (hlmatrixcontroltable *RMON2MIB_Hlmatrixcontroltable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry
// A conceptual row in the hlMatrixControlTable.
// 
// An example of indexing of this entry is
// hlMatrixControlNlDroppedFrames.1
type RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry struct {
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetFilter() yfilter.YFilter { return hlmatrixcontrolentry.YFilter }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) SetFilter(yf yfilter.YFilter) { hlmatrixcontrolentry.YFilter = yf }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetGoName(yname string) string {
    if yname == "hlMatrixControlIndex" { return "Hlmatrixcontrolindex" }
    if yname == "hlMatrixControlDataSource" { return "Hlmatrixcontroldatasource" }
    if yname == "hlMatrixControlNlDroppedFrames" { return "Hlmatrixcontrolnldroppedframes" }
    if yname == "hlMatrixControlNlInserts" { return "Hlmatrixcontrolnlinserts" }
    if yname == "hlMatrixControlNlDeletes" { return "Hlmatrixcontrolnldeletes" }
    if yname == "hlMatrixControlNlMaxDesiredEntries" { return "Hlmatrixcontrolnlmaxdesiredentries" }
    if yname == "hlMatrixControlAlDroppedFrames" { return "Hlmatrixcontrolaldroppedframes" }
    if yname == "hlMatrixControlAlInserts" { return "Hlmatrixcontrolalinserts" }
    if yname == "hlMatrixControlAlDeletes" { return "Hlmatrixcontrolaldeletes" }
    if yname == "hlMatrixControlAlMaxDesiredEntries" { return "Hlmatrixcontrolalmaxdesiredentries" }
    if yname == "hlMatrixControlOwner" { return "Hlmatrixcontrolowner" }
    if yname == "hlMatrixControlStatus" { return "Hlmatrixcontrolstatus" }
    return ""
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetSegmentPath() string {
    return "hlMatrixControlEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", hlmatrixcontrolentry.Hlmatrixcontrolindex) + "']"
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlMatrixControlIndex"] = hlmatrixcontrolentry.Hlmatrixcontrolindex
    leafs["hlMatrixControlDataSource"] = hlmatrixcontrolentry.Hlmatrixcontroldatasource
    leafs["hlMatrixControlNlDroppedFrames"] = hlmatrixcontrolentry.Hlmatrixcontrolnldroppedframes
    leafs["hlMatrixControlNlInserts"] = hlmatrixcontrolentry.Hlmatrixcontrolnlinserts
    leafs["hlMatrixControlNlDeletes"] = hlmatrixcontrolentry.Hlmatrixcontrolnldeletes
    leafs["hlMatrixControlNlMaxDesiredEntries"] = hlmatrixcontrolentry.Hlmatrixcontrolnlmaxdesiredentries
    leafs["hlMatrixControlAlDroppedFrames"] = hlmatrixcontrolentry.Hlmatrixcontrolaldroppedframes
    leafs["hlMatrixControlAlInserts"] = hlmatrixcontrolentry.Hlmatrixcontrolalinserts
    leafs["hlMatrixControlAlDeletes"] = hlmatrixcontrolentry.Hlmatrixcontrolaldeletes
    leafs["hlMatrixControlAlMaxDesiredEntries"] = hlmatrixcontrolentry.Hlmatrixcontrolalmaxdesiredentries
    leafs["hlMatrixControlOwner"] = hlmatrixcontrolentry.Hlmatrixcontrolowner
    leafs["hlMatrixControlStatus"] = hlmatrixcontrolentry.Hlmatrixcontrolstatus
    return leafs
}

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetYangName() string { return "hlMatrixControlEntry" }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) SetParent(parent types.Entity) { hlmatrixcontrolentry.parent = parent }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetParent() types.Entity { return hlmatrixcontrolentry.parent }

func (hlmatrixcontrolentry *RMON2MIB_Hlmatrixcontroltable_Hlmatrixcontrolentry) GetParentYangName() string { return "hlMatrixControlTable" }

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
    parent types.Entity
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

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetFilter() yfilter.YFilter { return nlmatrixsdtable.YFilter }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) SetFilter(yf yfilter.YFilter) { nlmatrixsdtable.YFilter = yf }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetGoName(yname string) string {
    if yname == "nlMatrixSDEntry" { return "Nlmatrixsdentry" }
    return ""
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetSegmentPath() string {
    return "nlMatrixSDTable"
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlMatrixSDEntry" {
        for _, c := range nlmatrixsdtable.Nlmatrixsdentry {
            if nlmatrixsdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry{}
        nlmatrixsdtable.Nlmatrixsdentry = append(nlmatrixsdtable.Nlmatrixsdentry, child)
        return &nlmatrixsdtable.Nlmatrixsdentry[len(nlmatrixsdtable.Nlmatrixsdentry)-1]
    }
    return nil
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmatrixsdtable.Nlmatrixsdentry {
        children[nlmatrixsdtable.Nlmatrixsdentry[i].GetSegmentPath()] = &nlmatrixsdtable.Nlmatrixsdentry[i]
    }
    return children
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetYangName() string { return "nlMatrixSDTable" }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) SetParent(parent types.Entity) { nlmatrixsdtable.parent = parent }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetParent() types.Entity { return nlmatrixsdtable.parent }

func (nlmatrixsdtable *RMON2MIB_Nlmatrixsdtable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetFilter() yfilter.YFilter { return nlmatrixsdentry.YFilter }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) SetFilter(yf yfilter.YFilter) { nlmatrixsdentry.YFilter = yf }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetGoName(yname string) string {
    if yname == "hlMatrixControlIndex" { return "Hlmatrixcontrolindex" }
    if yname == "nlMatrixSDTimeMark" { return "Nlmatrixsdtimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlMatrixSDSourceAddress" { return "Nlmatrixsdsourceaddress" }
    if yname == "nlMatrixSDDestAddress" { return "Nlmatrixsddestaddress" }
    if yname == "nlMatrixSDPkts" { return "Nlmatrixsdpkts" }
    if yname == "nlMatrixSDOctets" { return "Nlmatrixsdoctets" }
    if yname == "nlMatrixSDCreateTime" { return "Nlmatrixsdcreatetime" }
    return ""
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetSegmentPath() string {
    return "nlMatrixSDEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", nlmatrixsdentry.Hlmatrixcontrolindex) + "']" + "[nlMatrixSDTimeMark='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsdtimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlmatrixsdentry.Protocoldirlocalindex) + "']" + "[nlMatrixSDSourceAddress='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsdsourceaddress) + "']" + "[nlMatrixSDDestAddress='" + fmt.Sprintf("%v", nlmatrixsdentry.Nlmatrixsddestaddress) + "']"
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlMatrixControlIndex"] = nlmatrixsdentry.Hlmatrixcontrolindex
    leafs["nlMatrixSDTimeMark"] = nlmatrixsdentry.Nlmatrixsdtimemark
    leafs["protocolDirLocalIndex"] = nlmatrixsdentry.Protocoldirlocalindex
    leafs["nlMatrixSDSourceAddress"] = nlmatrixsdentry.Nlmatrixsdsourceaddress
    leafs["nlMatrixSDDestAddress"] = nlmatrixsdentry.Nlmatrixsddestaddress
    leafs["nlMatrixSDPkts"] = nlmatrixsdentry.Nlmatrixsdpkts
    leafs["nlMatrixSDOctets"] = nlmatrixsdentry.Nlmatrixsdoctets
    leafs["nlMatrixSDCreateTime"] = nlmatrixsdentry.Nlmatrixsdcreatetime
    return leafs
}

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetYangName() string { return "nlMatrixSDEntry" }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) SetParent(parent types.Entity) { nlmatrixsdentry.parent = parent }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetParent() types.Entity { return nlmatrixsdentry.parent }

func (nlmatrixsdentry *RMON2MIB_Nlmatrixsdtable_Nlmatrixsdentry) GetParentYangName() string { return "nlMatrixSDTable" }

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
    parent types.Entity
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

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetFilter() yfilter.YFilter { return nlmatrixdstable.YFilter }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) SetFilter(yf yfilter.YFilter) { nlmatrixdstable.YFilter = yf }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetGoName(yname string) string {
    if yname == "nlMatrixDSEntry" { return "Nlmatrixdsentry" }
    return ""
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetSegmentPath() string {
    return "nlMatrixDSTable"
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlMatrixDSEntry" {
        for _, c := range nlmatrixdstable.Nlmatrixdsentry {
            if nlmatrixdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry{}
        nlmatrixdstable.Nlmatrixdsentry = append(nlmatrixdstable.Nlmatrixdsentry, child)
        return &nlmatrixdstable.Nlmatrixdsentry[len(nlmatrixdstable.Nlmatrixdsentry)-1]
    }
    return nil
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmatrixdstable.Nlmatrixdsentry {
        children[nlmatrixdstable.Nlmatrixdsentry[i].GetSegmentPath()] = &nlmatrixdstable.Nlmatrixdsentry[i]
    }
    return children
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetYangName() string { return "nlMatrixDSTable" }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) SetParent(parent types.Entity) { nlmatrixdstable.parent = parent }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetParent() types.Entity { return nlmatrixdstable.parent }

func (nlmatrixdstable *RMON2MIB_Nlmatrixdstable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetFilter() yfilter.YFilter { return nlmatrixdsentry.YFilter }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) SetFilter(yf yfilter.YFilter) { nlmatrixdsentry.YFilter = yf }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetGoName(yname string) string {
    if yname == "hlMatrixControlIndex" { return "Hlmatrixcontrolindex" }
    if yname == "nlMatrixDSTimeMark" { return "Nlmatrixdstimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlMatrixDSDestAddress" { return "Nlmatrixdsdestaddress" }
    if yname == "nlMatrixDSSourceAddress" { return "Nlmatrixdssourceaddress" }
    if yname == "nlMatrixDSPkts" { return "Nlmatrixdspkts" }
    if yname == "nlMatrixDSOctets" { return "Nlmatrixdsoctets" }
    if yname == "nlMatrixDSCreateTime" { return "Nlmatrixdscreatetime" }
    return ""
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetSegmentPath() string {
    return "nlMatrixDSEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", nlmatrixdsentry.Hlmatrixcontrolindex) + "']" + "[nlMatrixDSTimeMark='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdstimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", nlmatrixdsentry.Protocoldirlocalindex) + "']" + "[nlMatrixDSDestAddress='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdsdestaddress) + "']" + "[nlMatrixDSSourceAddress='" + fmt.Sprintf("%v", nlmatrixdsentry.Nlmatrixdssourceaddress) + "']"
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlMatrixControlIndex"] = nlmatrixdsentry.Hlmatrixcontrolindex
    leafs["nlMatrixDSTimeMark"] = nlmatrixdsentry.Nlmatrixdstimemark
    leafs["protocolDirLocalIndex"] = nlmatrixdsentry.Protocoldirlocalindex
    leafs["nlMatrixDSDestAddress"] = nlmatrixdsentry.Nlmatrixdsdestaddress
    leafs["nlMatrixDSSourceAddress"] = nlmatrixdsentry.Nlmatrixdssourceaddress
    leafs["nlMatrixDSPkts"] = nlmatrixdsentry.Nlmatrixdspkts
    leafs["nlMatrixDSOctets"] = nlmatrixdsentry.Nlmatrixdsoctets
    leafs["nlMatrixDSCreateTime"] = nlmatrixdsentry.Nlmatrixdscreatetime
    return leafs
}

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetYangName() string { return "nlMatrixDSEntry" }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) SetParent(parent types.Entity) { nlmatrixdsentry.parent = parent }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetParent() types.Entity { return nlmatrixdsentry.parent }

func (nlmatrixdsentry *RMON2MIB_Nlmatrixdstable_Nlmatrixdsentry) GetParentYangName() string { return "nlMatrixDSTable" }

// RMON2MIB_Nlmatrixtopncontroltable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_Nlmatrixtopncontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNControlTable.  An example of the
    // indexing of this table is nlMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry.
    Nlmatrixtopncontrolentry []RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetFilter() yfilter.YFilter { return nlmatrixtopncontroltable.YFilter }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) SetFilter(yf yfilter.YFilter) { nlmatrixtopncontroltable.YFilter = yf }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetGoName(yname string) string {
    if yname == "nlMatrixTopNControlEntry" { return "Nlmatrixtopncontrolentry" }
    return ""
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetSegmentPath() string {
    return "nlMatrixTopNControlTable"
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlMatrixTopNControlEntry" {
        for _, c := range nlmatrixtopncontroltable.Nlmatrixtopncontrolentry {
            if nlmatrixtopncontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry{}
        nlmatrixtopncontroltable.Nlmatrixtopncontrolentry = append(nlmatrixtopncontroltable.Nlmatrixtopncontrolentry, child)
        return &nlmatrixtopncontroltable.Nlmatrixtopncontrolentry[len(nlmatrixtopncontroltable.Nlmatrixtopncontrolentry)-1]
    }
    return nil
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmatrixtopncontroltable.Nlmatrixtopncontrolentry {
        children[nlmatrixtopncontroltable.Nlmatrixtopncontrolentry[i].GetSegmentPath()] = &nlmatrixtopncontroltable.Nlmatrixtopncontrolentry[i]
    }
    return children
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetYangName() string { return "nlMatrixTopNControlTable" }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) SetParent(parent types.Entity) { nlmatrixtopncontroltable.parent = parent }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetParent() types.Entity { return nlmatrixtopncontroltable.parent }

func (nlmatrixtopncontroltable *RMON2MIB_Nlmatrixtopncontroltable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry
// A conceptual row in the nlMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// nlMatrixTopNControlDuration.3
type RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry struct {
    parent types.Entity
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

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetFilter() yfilter.YFilter { return nlmatrixtopncontrolentry.YFilter }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) SetFilter(yf yfilter.YFilter) { nlmatrixtopncontrolentry.YFilter = yf }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetGoName(yname string) string {
    if yname == "nlMatrixTopNControlIndex" { return "Nlmatrixtopncontrolindex" }
    if yname == "nlMatrixTopNControlMatrixIndex" { return "Nlmatrixtopncontrolmatrixindex" }
    if yname == "nlMatrixTopNControlRateBase" { return "Nlmatrixtopncontrolratebase" }
    if yname == "nlMatrixTopNControlTimeRemaining" { return "Nlmatrixtopncontroltimeremaining" }
    if yname == "nlMatrixTopNControlGeneratedReports" { return "Nlmatrixtopncontrolgeneratedreports" }
    if yname == "nlMatrixTopNControlDuration" { return "Nlmatrixtopncontrolduration" }
    if yname == "nlMatrixTopNControlRequestedSize" { return "Nlmatrixtopncontrolrequestedsize" }
    if yname == "nlMatrixTopNControlGrantedSize" { return "Nlmatrixtopncontrolgrantedsize" }
    if yname == "nlMatrixTopNControlStartTime" { return "Nlmatrixtopncontrolstarttime" }
    if yname == "nlMatrixTopNControlOwner" { return "Nlmatrixtopncontrolowner" }
    if yname == "nlMatrixTopNControlStatus" { return "Nlmatrixtopncontrolstatus" }
    return ""
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetSegmentPath() string {
    return "nlMatrixTopNControlEntry" + "[nlMatrixTopNControlIndex='" + fmt.Sprintf("%v", nlmatrixtopncontrolentry.Nlmatrixtopncontrolindex) + "']"
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlMatrixTopNControlIndex"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolindex
    leafs["nlMatrixTopNControlMatrixIndex"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolmatrixindex
    leafs["nlMatrixTopNControlRateBase"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolratebase
    leafs["nlMatrixTopNControlTimeRemaining"] = nlmatrixtopncontrolentry.Nlmatrixtopncontroltimeremaining
    leafs["nlMatrixTopNControlGeneratedReports"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolgeneratedreports
    leafs["nlMatrixTopNControlDuration"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolduration
    leafs["nlMatrixTopNControlRequestedSize"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolrequestedsize
    leafs["nlMatrixTopNControlGrantedSize"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolgrantedsize
    leafs["nlMatrixTopNControlStartTime"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolstarttime
    leafs["nlMatrixTopNControlOwner"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolowner
    leafs["nlMatrixTopNControlStatus"] = nlmatrixtopncontrolentry.Nlmatrixtopncontrolstatus
    return leafs
}

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetYangName() string { return "nlMatrixTopNControlEntry" }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) SetParent(parent types.Entity) { nlmatrixtopncontrolentry.parent = parent }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetParent() types.Entity { return nlmatrixtopncontrolentry.parent }

func (nlmatrixtopncontrolentry *RMON2MIB_Nlmatrixtopncontroltable_Nlmatrixtopncontrolentry) GetParentYangName() string { return "nlMatrixTopNControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the nlMatrixTopNTable.  The nlMatrixTopNControlIndex
    // value in the index identifies the nlMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // nlMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry.
    Nlmatrixtopnentry []RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetFilter() yfilter.YFilter { return nlmatrixtopntable.YFilter }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) SetFilter(yf yfilter.YFilter) { nlmatrixtopntable.YFilter = yf }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetGoName(yname string) string {
    if yname == "nlMatrixTopNEntry" { return "Nlmatrixtopnentry" }
    return ""
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetSegmentPath() string {
    return "nlMatrixTopNTable"
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlMatrixTopNEntry" {
        for _, c := range nlmatrixtopntable.Nlmatrixtopnentry {
            if nlmatrixtopntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry{}
        nlmatrixtopntable.Nlmatrixtopnentry = append(nlmatrixtopntable.Nlmatrixtopnentry, child)
        return &nlmatrixtopntable.Nlmatrixtopnentry[len(nlmatrixtopntable.Nlmatrixtopnentry)-1]
    }
    return nil
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmatrixtopntable.Nlmatrixtopnentry {
        children[nlmatrixtopntable.Nlmatrixtopnentry[i].GetSegmentPath()] = &nlmatrixtopntable.Nlmatrixtopnentry[i]
    }
    return children
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetYangName() string { return "nlMatrixTopNTable" }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) SetParent(parent types.Entity) { nlmatrixtopntable.parent = parent }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetParent() types.Entity { return nlmatrixtopntable.parent }

func (nlmatrixtopntable *RMON2MIB_Nlmatrixtopntable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetFilter() yfilter.YFilter { return nlmatrixtopnentry.YFilter }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) SetFilter(yf yfilter.YFilter) { nlmatrixtopnentry.YFilter = yf }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetGoName(yname string) string {
    if yname == "nlMatrixTopNControlIndex" { return "Nlmatrixtopncontrolindex" }
    if yname == "nlMatrixTopNIndex" { return "Nlmatrixtopnindex" }
    if yname == "nlMatrixTopNProtocolDirLocalIndex" { return "Nlmatrixtopnprotocoldirlocalindex" }
    if yname == "nlMatrixTopNSourceAddress" { return "Nlmatrixtopnsourceaddress" }
    if yname == "nlMatrixTopNDestAddress" { return "Nlmatrixtopndestaddress" }
    if yname == "nlMatrixTopNPktRate" { return "Nlmatrixtopnpktrate" }
    if yname == "nlMatrixTopNReversePktRate" { return "Nlmatrixtopnreversepktrate" }
    if yname == "nlMatrixTopNOctetRate" { return "Nlmatrixtopnoctetrate" }
    if yname == "nlMatrixTopNReverseOctetRate" { return "Nlmatrixtopnreverseoctetrate" }
    return ""
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetSegmentPath() string {
    return "nlMatrixTopNEntry" + "[nlMatrixTopNControlIndex='" + fmt.Sprintf("%v", nlmatrixtopnentry.Nlmatrixtopncontrolindex) + "']" + "[nlMatrixTopNIndex='" + fmt.Sprintf("%v", nlmatrixtopnentry.Nlmatrixtopnindex) + "']"
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlMatrixTopNControlIndex"] = nlmatrixtopnentry.Nlmatrixtopncontrolindex
    leafs["nlMatrixTopNIndex"] = nlmatrixtopnentry.Nlmatrixtopnindex
    leafs["nlMatrixTopNProtocolDirLocalIndex"] = nlmatrixtopnentry.Nlmatrixtopnprotocoldirlocalindex
    leafs["nlMatrixTopNSourceAddress"] = nlmatrixtopnentry.Nlmatrixtopnsourceaddress
    leafs["nlMatrixTopNDestAddress"] = nlmatrixtopnentry.Nlmatrixtopndestaddress
    leafs["nlMatrixTopNPktRate"] = nlmatrixtopnentry.Nlmatrixtopnpktrate
    leafs["nlMatrixTopNReversePktRate"] = nlmatrixtopnentry.Nlmatrixtopnreversepktrate
    leafs["nlMatrixTopNOctetRate"] = nlmatrixtopnentry.Nlmatrixtopnoctetrate
    leafs["nlMatrixTopNReverseOctetRate"] = nlmatrixtopnentry.Nlmatrixtopnreverseoctetrate
    return leafs
}

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetYangName() string { return "nlMatrixTopNEntry" }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) SetParent(parent types.Entity) { nlmatrixtopnentry.parent = parent }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetParent() types.Entity { return nlmatrixtopnentry.parent }

func (nlmatrixtopnentry *RMON2MIB_Nlmatrixtopntable_Nlmatrixtopnentry) GetParentYangName() string { return "nlMatrixTopNTable" }

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
    parent types.Entity
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

func (alhosttable *RMON2MIB_Alhosttable) GetFilter() yfilter.YFilter { return alhosttable.YFilter }

func (alhosttable *RMON2MIB_Alhosttable) SetFilter(yf yfilter.YFilter) { alhosttable.YFilter = yf }

func (alhosttable *RMON2MIB_Alhosttable) GetGoName(yname string) string {
    if yname == "alHostEntry" { return "Alhostentry" }
    return ""
}

func (alhosttable *RMON2MIB_Alhosttable) GetSegmentPath() string {
    return "alHostTable"
}

func (alhosttable *RMON2MIB_Alhosttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alHostEntry" {
        for _, c := range alhosttable.Alhostentry {
            if alhosttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Alhosttable_Alhostentry{}
        alhosttable.Alhostentry = append(alhosttable.Alhostentry, child)
        return &alhosttable.Alhostentry[len(alhosttable.Alhostentry)-1]
    }
    return nil
}

func (alhosttable *RMON2MIB_Alhosttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alhosttable.Alhostentry {
        children[alhosttable.Alhostentry[i].GetSegmentPath()] = &alhosttable.Alhostentry[i]
    }
    return children
}

func (alhosttable *RMON2MIB_Alhosttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alhosttable *RMON2MIB_Alhosttable) GetBundleName() string { return "cisco_ios_xe" }

func (alhosttable *RMON2MIB_Alhosttable) GetYangName() string { return "alHostTable" }

func (alhosttable *RMON2MIB_Alhosttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (alhosttable *RMON2MIB_Alhosttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (alhosttable *RMON2MIB_Alhosttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (alhosttable *RMON2MIB_Alhosttable) SetParent(parent types.Entity) { alhosttable.parent = parent }

func (alhosttable *RMON2MIB_Alhosttable) GetParent() types.Entity { return alhosttable.parent }

func (alhosttable *RMON2MIB_Alhosttable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetFilter() yfilter.YFilter { return alhostentry.YFilter }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) SetFilter(yf yfilter.YFilter) { alhostentry.YFilter = yf }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetGoName(yname string) string {
    if yname == "hlHostControlIndex" { return "Hlhostcontrolindex" }
    if yname == "alHostTimeMark" { return "Alhosttimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlHostAddress" { return "Nlhostaddress" }
    if yname == "protocolDirLocalIndex_2" { return "Protocoldirlocalindex2" }
    if yname == "alHostInPkts" { return "Alhostinpkts" }
    if yname == "alHostOutPkts" { return "Alhostoutpkts" }
    if yname == "alHostInOctets" { return "Alhostinoctets" }
    if yname == "alHostOutOctets" { return "Alhostoutoctets" }
    if yname == "alHostCreateTime" { return "Alhostcreatetime" }
    return ""
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetSegmentPath() string {
    return "alHostEntry" + "[hlHostControlIndex='" + fmt.Sprintf("%v", alhostentry.Hlhostcontrolindex) + "']" + "[alHostTimeMark='" + fmt.Sprintf("%v", alhostentry.Alhosttimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", alhostentry.Protocoldirlocalindex) + "']" + "[nlHostAddress='" + fmt.Sprintf("%v", alhostentry.Nlhostaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", alhostentry.Protocoldirlocalindex2) + "']"
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlHostControlIndex"] = alhostentry.Hlhostcontrolindex
    leafs["alHostTimeMark"] = alhostentry.Alhosttimemark
    leafs["protocolDirLocalIndex"] = alhostentry.Protocoldirlocalindex
    leafs["nlHostAddress"] = alhostentry.Nlhostaddress
    leafs["protocolDirLocalIndex_2"] = alhostentry.Protocoldirlocalindex2
    leafs["alHostInPkts"] = alhostentry.Alhostinpkts
    leafs["alHostOutPkts"] = alhostentry.Alhostoutpkts
    leafs["alHostInOctets"] = alhostentry.Alhostinoctets
    leafs["alHostOutOctets"] = alhostentry.Alhostoutoctets
    leafs["alHostCreateTime"] = alhostentry.Alhostcreatetime
    return leafs
}

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetBundleName() string { return "cisco_ios_xe" }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetYangName() string { return "alHostEntry" }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) SetParent(parent types.Entity) { alhostentry.parent = parent }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetParent() types.Entity { return alhostentry.parent }

func (alhostentry *RMON2MIB_Alhosttable_Alhostentry) GetParentYangName() string { return "alHostTable" }

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
    parent types.Entity
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

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetFilter() yfilter.YFilter { return almatrixsdtable.YFilter }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) SetFilter(yf yfilter.YFilter) { almatrixsdtable.YFilter = yf }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetGoName(yname string) string {
    if yname == "alMatrixSDEntry" { return "Almatrixsdentry" }
    return ""
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetSegmentPath() string {
    return "alMatrixSDTable"
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alMatrixSDEntry" {
        for _, c := range almatrixsdtable.Almatrixsdentry {
            if almatrixsdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Almatrixsdtable_Almatrixsdentry{}
        almatrixsdtable.Almatrixsdentry = append(almatrixsdtable.Almatrixsdentry, child)
        return &almatrixsdtable.Almatrixsdentry[len(almatrixsdtable.Almatrixsdentry)-1]
    }
    return nil
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range almatrixsdtable.Almatrixsdentry {
        children[almatrixsdtable.Almatrixsdentry[i].GetSegmentPath()] = &almatrixsdtable.Almatrixsdentry[i]
    }
    return children
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetYangName() string { return "alMatrixSDTable" }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) SetParent(parent types.Entity) { almatrixsdtable.parent = parent }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetParent() types.Entity { return almatrixsdtable.parent }

func (almatrixsdtable *RMON2MIB_Almatrixsdtable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetFilter() yfilter.YFilter { return almatrixsdentry.YFilter }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) SetFilter(yf yfilter.YFilter) { almatrixsdentry.YFilter = yf }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetGoName(yname string) string {
    if yname == "hlMatrixControlIndex" { return "Hlmatrixcontrolindex" }
    if yname == "alMatrixSDTimeMark" { return "Almatrixsdtimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlMatrixSDSourceAddress" { return "Nlmatrixsdsourceaddress" }
    if yname == "nlMatrixSDDestAddress" { return "Nlmatrixsddestaddress" }
    if yname == "protocolDirLocalIndex_2" { return "Protocoldirlocalindex2" }
    if yname == "alMatrixSDPkts" { return "Almatrixsdpkts" }
    if yname == "alMatrixSDOctets" { return "Almatrixsdoctets" }
    if yname == "alMatrixSDCreateTime" { return "Almatrixsdcreatetime" }
    return ""
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetSegmentPath() string {
    return "alMatrixSDEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", almatrixsdentry.Hlmatrixcontrolindex) + "']" + "[alMatrixSDTimeMark='" + fmt.Sprintf("%v", almatrixsdentry.Almatrixsdtimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", almatrixsdentry.Protocoldirlocalindex) + "']" + "[nlMatrixSDSourceAddress='" + fmt.Sprintf("%v", almatrixsdentry.Nlmatrixsdsourceaddress) + "']" + "[nlMatrixSDDestAddress='" + fmt.Sprintf("%v", almatrixsdentry.Nlmatrixsddestaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", almatrixsdentry.Protocoldirlocalindex2) + "']"
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlMatrixControlIndex"] = almatrixsdentry.Hlmatrixcontrolindex
    leafs["alMatrixSDTimeMark"] = almatrixsdentry.Almatrixsdtimemark
    leafs["protocolDirLocalIndex"] = almatrixsdentry.Protocoldirlocalindex
    leafs["nlMatrixSDSourceAddress"] = almatrixsdentry.Nlmatrixsdsourceaddress
    leafs["nlMatrixSDDestAddress"] = almatrixsdentry.Nlmatrixsddestaddress
    leafs["protocolDirLocalIndex_2"] = almatrixsdentry.Protocoldirlocalindex2
    leafs["alMatrixSDPkts"] = almatrixsdentry.Almatrixsdpkts
    leafs["alMatrixSDOctets"] = almatrixsdentry.Almatrixsdoctets
    leafs["alMatrixSDCreateTime"] = almatrixsdentry.Almatrixsdcreatetime
    return leafs
}

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetYangName() string { return "alMatrixSDEntry" }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) SetParent(parent types.Entity) { almatrixsdentry.parent = parent }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetParent() types.Entity { return almatrixsdentry.parent }

func (almatrixsdentry *RMON2MIB_Almatrixsdtable_Almatrixsdentry) GetParentYangName() string { return "alMatrixSDTable" }

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
    parent types.Entity
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

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetFilter() yfilter.YFilter { return almatrixdstable.YFilter }

func (almatrixdstable *RMON2MIB_Almatrixdstable) SetFilter(yf yfilter.YFilter) { almatrixdstable.YFilter = yf }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetGoName(yname string) string {
    if yname == "alMatrixDSEntry" { return "Almatrixdsentry" }
    return ""
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetSegmentPath() string {
    return "alMatrixDSTable"
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alMatrixDSEntry" {
        for _, c := range almatrixdstable.Almatrixdsentry {
            if almatrixdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Almatrixdstable_Almatrixdsentry{}
        almatrixdstable.Almatrixdsentry = append(almatrixdstable.Almatrixdsentry, child)
        return &almatrixdstable.Almatrixdsentry[len(almatrixdstable.Almatrixdsentry)-1]
    }
    return nil
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range almatrixdstable.Almatrixdsentry {
        children[almatrixdstable.Almatrixdsentry[i].GetSegmentPath()] = &almatrixdstable.Almatrixdsentry[i]
    }
    return children
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetYangName() string { return "alMatrixDSTable" }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixdstable *RMON2MIB_Almatrixdstable) SetParent(parent types.Entity) { almatrixdstable.parent = parent }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetParent() types.Entity { return almatrixdstable.parent }

func (almatrixdstable *RMON2MIB_Almatrixdstable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetFilter() yfilter.YFilter { return almatrixdsentry.YFilter }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) SetFilter(yf yfilter.YFilter) { almatrixdsentry.YFilter = yf }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetGoName(yname string) string {
    if yname == "hlMatrixControlIndex" { return "Hlmatrixcontrolindex" }
    if yname == "alMatrixDSTimeMark" { return "Almatrixdstimemark" }
    if yname == "protocolDirLocalIndex" { return "Protocoldirlocalindex" }
    if yname == "nlMatrixDSDestAddress" { return "Nlmatrixdsdestaddress" }
    if yname == "nlMatrixDSSourceAddress" { return "Nlmatrixdssourceaddress" }
    if yname == "protocolDirLocalIndex_2" { return "Protocoldirlocalindex2" }
    if yname == "alMatrixDSPkts" { return "Almatrixdspkts" }
    if yname == "alMatrixDSOctets" { return "Almatrixdsoctets" }
    if yname == "alMatrixDSCreateTime" { return "Almatrixdscreatetime" }
    return ""
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetSegmentPath() string {
    return "alMatrixDSEntry" + "[hlMatrixControlIndex='" + fmt.Sprintf("%v", almatrixdsentry.Hlmatrixcontrolindex) + "']" + "[alMatrixDSTimeMark='" + fmt.Sprintf("%v", almatrixdsentry.Almatrixdstimemark) + "']" + "[protocolDirLocalIndex='" + fmt.Sprintf("%v", almatrixdsentry.Protocoldirlocalindex) + "']" + "[nlMatrixDSDestAddress='" + fmt.Sprintf("%v", almatrixdsentry.Nlmatrixdsdestaddress) + "']" + "[nlMatrixDSSourceAddress='" + fmt.Sprintf("%v", almatrixdsentry.Nlmatrixdssourceaddress) + "']" + "[protocolDirLocalIndex_2='" + fmt.Sprintf("%v", almatrixdsentry.Protocoldirlocalindex2) + "']"
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hlMatrixControlIndex"] = almatrixdsentry.Hlmatrixcontrolindex
    leafs["alMatrixDSTimeMark"] = almatrixdsentry.Almatrixdstimemark
    leafs["protocolDirLocalIndex"] = almatrixdsentry.Protocoldirlocalindex
    leafs["nlMatrixDSDestAddress"] = almatrixdsentry.Nlmatrixdsdestaddress
    leafs["nlMatrixDSSourceAddress"] = almatrixdsentry.Nlmatrixdssourceaddress
    leafs["protocolDirLocalIndex_2"] = almatrixdsentry.Protocoldirlocalindex2
    leafs["alMatrixDSPkts"] = almatrixdsentry.Almatrixdspkts
    leafs["alMatrixDSOctets"] = almatrixdsentry.Almatrixdsoctets
    leafs["alMatrixDSCreateTime"] = almatrixdsentry.Almatrixdscreatetime
    return leafs
}

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetYangName() string { return "alMatrixDSEntry" }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) SetParent(parent types.Entity) { almatrixdsentry.parent = parent }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetParent() types.Entity { return almatrixdsentry.parent }

func (almatrixdsentry *RMON2MIB_Almatrixdstable_Almatrixdsentry) GetParentYangName() string { return "alMatrixDSTable" }

// RMON2MIB_Almatrixtopncontroltable
// A set of parameters that control the creation of a
// report of the top N matrix entries according to
// a selected metric.
type RMON2MIB_Almatrixtopncontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNControlTable.  An example of the
    // indexing of this table is alMatrixTopNControlDuration.3. The type is slice
    // of RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry.
    Almatrixtopncontrolentry []RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetFilter() yfilter.YFilter { return almatrixtopncontroltable.YFilter }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) SetFilter(yf yfilter.YFilter) { almatrixtopncontroltable.YFilter = yf }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetGoName(yname string) string {
    if yname == "alMatrixTopNControlEntry" { return "Almatrixtopncontrolentry" }
    return ""
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetSegmentPath() string {
    return "alMatrixTopNControlTable"
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alMatrixTopNControlEntry" {
        for _, c := range almatrixtopncontroltable.Almatrixtopncontrolentry {
            if almatrixtopncontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry{}
        almatrixtopncontroltable.Almatrixtopncontrolentry = append(almatrixtopncontroltable.Almatrixtopncontrolentry, child)
        return &almatrixtopncontroltable.Almatrixtopncontrolentry[len(almatrixtopncontroltable.Almatrixtopncontrolentry)-1]
    }
    return nil
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range almatrixtopncontroltable.Almatrixtopncontrolentry {
        children[almatrixtopncontroltable.Almatrixtopncontrolentry[i].GetSegmentPath()] = &almatrixtopncontroltable.Almatrixtopncontrolentry[i]
    }
    return children
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetYangName() string { return "alMatrixTopNControlTable" }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) SetParent(parent types.Entity) { almatrixtopncontroltable.parent = parent }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetParent() types.Entity { return almatrixtopncontroltable.parent }

func (almatrixtopncontroltable *RMON2MIB_Almatrixtopncontroltable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry
// A conceptual row in the alMatrixTopNControlTable.
// 
// An example of the indexing of this table is
// alMatrixTopNControlDuration.3
type RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry struct {
    parent types.Entity
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

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetFilter() yfilter.YFilter { return almatrixtopncontrolentry.YFilter }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) SetFilter(yf yfilter.YFilter) { almatrixtopncontrolentry.YFilter = yf }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetGoName(yname string) string {
    if yname == "alMatrixTopNControlIndex" { return "Almatrixtopncontrolindex" }
    if yname == "alMatrixTopNControlMatrixIndex" { return "Almatrixtopncontrolmatrixindex" }
    if yname == "alMatrixTopNControlRateBase" { return "Almatrixtopncontrolratebase" }
    if yname == "alMatrixTopNControlTimeRemaining" { return "Almatrixtopncontroltimeremaining" }
    if yname == "alMatrixTopNControlGeneratedReports" { return "Almatrixtopncontrolgeneratedreports" }
    if yname == "alMatrixTopNControlDuration" { return "Almatrixtopncontrolduration" }
    if yname == "alMatrixTopNControlRequestedSize" { return "Almatrixtopncontrolrequestedsize" }
    if yname == "alMatrixTopNControlGrantedSize" { return "Almatrixtopncontrolgrantedsize" }
    if yname == "alMatrixTopNControlStartTime" { return "Almatrixtopncontrolstarttime" }
    if yname == "alMatrixTopNControlOwner" { return "Almatrixtopncontrolowner" }
    if yname == "alMatrixTopNControlStatus" { return "Almatrixtopncontrolstatus" }
    return ""
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetSegmentPath() string {
    return "alMatrixTopNControlEntry" + "[alMatrixTopNControlIndex='" + fmt.Sprintf("%v", almatrixtopncontrolentry.Almatrixtopncontrolindex) + "']"
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["alMatrixTopNControlIndex"] = almatrixtopncontrolentry.Almatrixtopncontrolindex
    leafs["alMatrixTopNControlMatrixIndex"] = almatrixtopncontrolentry.Almatrixtopncontrolmatrixindex
    leafs["alMatrixTopNControlRateBase"] = almatrixtopncontrolentry.Almatrixtopncontrolratebase
    leafs["alMatrixTopNControlTimeRemaining"] = almatrixtopncontrolentry.Almatrixtopncontroltimeremaining
    leafs["alMatrixTopNControlGeneratedReports"] = almatrixtopncontrolentry.Almatrixtopncontrolgeneratedreports
    leafs["alMatrixTopNControlDuration"] = almatrixtopncontrolentry.Almatrixtopncontrolduration
    leafs["alMatrixTopNControlRequestedSize"] = almatrixtopncontrolentry.Almatrixtopncontrolrequestedsize
    leafs["alMatrixTopNControlGrantedSize"] = almatrixtopncontrolentry.Almatrixtopncontrolgrantedsize
    leafs["alMatrixTopNControlStartTime"] = almatrixtopncontrolentry.Almatrixtopncontrolstarttime
    leafs["alMatrixTopNControlOwner"] = almatrixtopncontrolentry.Almatrixtopncontrolowner
    leafs["alMatrixTopNControlStatus"] = almatrixtopncontrolentry.Almatrixtopncontrolstatus
    return leafs
}

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetYangName() string { return "alMatrixTopNControlEntry" }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) SetParent(parent types.Entity) { almatrixtopncontrolentry.parent = parent }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetParent() types.Entity { return almatrixtopncontrolentry.parent }

func (almatrixtopncontrolentry *RMON2MIB_Almatrixtopncontroltable_Almatrixtopncontrolentry) GetParentYangName() string { return "alMatrixTopNControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the alMatrixTopNTable.  The alMatrixTopNControlIndex
    // value in the index identifies the alMatrixTopNControlEntry on whose behalf
    // this entry was created.  An example of the indexing of this table is
    // alMatrixTopNPktRate.3.10. The type is slice of
    // RMON2MIB_Almatrixtopntable_Almatrixtopnentry.
    Almatrixtopnentry []RMON2MIB_Almatrixtopntable_Almatrixtopnentry
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetFilter() yfilter.YFilter { return almatrixtopntable.YFilter }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) SetFilter(yf yfilter.YFilter) { almatrixtopntable.YFilter = yf }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetGoName(yname string) string {
    if yname == "alMatrixTopNEntry" { return "Almatrixtopnentry" }
    return ""
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetSegmentPath() string {
    return "alMatrixTopNTable"
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alMatrixTopNEntry" {
        for _, c := range almatrixtopntable.Almatrixtopnentry {
            if almatrixtopntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Almatrixtopntable_Almatrixtopnentry{}
        almatrixtopntable.Almatrixtopnentry = append(almatrixtopntable.Almatrixtopnentry, child)
        return &almatrixtopntable.Almatrixtopnentry[len(almatrixtopntable.Almatrixtopnentry)-1]
    }
    return nil
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range almatrixtopntable.Almatrixtopnentry {
        children[almatrixtopntable.Almatrixtopnentry[i].GetSegmentPath()] = &almatrixtopntable.Almatrixtopnentry[i]
    }
    return children
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetYangName() string { return "alMatrixTopNTable" }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) SetParent(parent types.Entity) { almatrixtopntable.parent = parent }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetParent() types.Entity { return almatrixtopntable.parent }

func (almatrixtopntable *RMON2MIB_Almatrixtopntable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetFilter() yfilter.YFilter { return almatrixtopnentry.YFilter }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) SetFilter(yf yfilter.YFilter) { almatrixtopnentry.YFilter = yf }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetGoName(yname string) string {
    if yname == "alMatrixTopNControlIndex" { return "Almatrixtopncontrolindex" }
    if yname == "alMatrixTopNIndex" { return "Almatrixtopnindex" }
    if yname == "alMatrixTopNProtocolDirLocalIndex" { return "Almatrixtopnprotocoldirlocalindex" }
    if yname == "alMatrixTopNSourceAddress" { return "Almatrixtopnsourceaddress" }
    if yname == "alMatrixTopNDestAddress" { return "Almatrixtopndestaddress" }
    if yname == "alMatrixTopNAppProtocolDirLocalIndex" { return "Almatrixtopnappprotocoldirlocalindex" }
    if yname == "alMatrixTopNPktRate" { return "Almatrixtopnpktrate" }
    if yname == "alMatrixTopNReversePktRate" { return "Almatrixtopnreversepktrate" }
    if yname == "alMatrixTopNOctetRate" { return "Almatrixtopnoctetrate" }
    if yname == "alMatrixTopNReverseOctetRate" { return "Almatrixtopnreverseoctetrate" }
    return ""
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetSegmentPath() string {
    return "alMatrixTopNEntry" + "[alMatrixTopNControlIndex='" + fmt.Sprintf("%v", almatrixtopnentry.Almatrixtopncontrolindex) + "']" + "[alMatrixTopNIndex='" + fmt.Sprintf("%v", almatrixtopnentry.Almatrixtopnindex) + "']"
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["alMatrixTopNControlIndex"] = almatrixtopnentry.Almatrixtopncontrolindex
    leafs["alMatrixTopNIndex"] = almatrixtopnentry.Almatrixtopnindex
    leafs["alMatrixTopNProtocolDirLocalIndex"] = almatrixtopnentry.Almatrixtopnprotocoldirlocalindex
    leafs["alMatrixTopNSourceAddress"] = almatrixtopnentry.Almatrixtopnsourceaddress
    leafs["alMatrixTopNDestAddress"] = almatrixtopnentry.Almatrixtopndestaddress
    leafs["alMatrixTopNAppProtocolDirLocalIndex"] = almatrixtopnentry.Almatrixtopnappprotocoldirlocalindex
    leafs["alMatrixTopNPktRate"] = almatrixtopnentry.Almatrixtopnpktrate
    leafs["alMatrixTopNReversePktRate"] = almatrixtopnentry.Almatrixtopnreversepktrate
    leafs["alMatrixTopNOctetRate"] = almatrixtopnentry.Almatrixtopnoctetrate
    leafs["alMatrixTopNReverseOctetRate"] = almatrixtopnentry.Almatrixtopnreverseoctetrate
    return leafs
}

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetBundleName() string { return "cisco_ios_xe" }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetYangName() string { return "alMatrixTopNEntry" }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) SetParent(parent types.Entity) { almatrixtopnentry.parent = parent }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetParent() types.Entity { return almatrixtopnentry.parent }

func (almatrixtopnentry *RMON2MIB_Almatrixtopntable_Almatrixtopnentry) GetParentYangName() string { return "alMatrixTopNTable" }

// RMON2MIB_Usrhistorycontroltable
// A list of data-collection configuration entries.
type RMON2MIB_Usrhistorycontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of parameters that set up a group of user-defined MIB objects to be
    // sampled periodically (called a bucket-group).  For example, an instance of
    // usrHistoryControlInterval might be named usrHistoryControlInterval.1. The
    // type is slice of RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry.
    Usrhistorycontrolentry []RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetFilter() yfilter.YFilter { return usrhistorycontroltable.YFilter }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) SetFilter(yf yfilter.YFilter) { usrhistorycontroltable.YFilter = yf }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetGoName(yname string) string {
    if yname == "usrHistoryControlEntry" { return "Usrhistorycontrolentry" }
    return ""
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetSegmentPath() string {
    return "usrHistoryControlTable"
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usrHistoryControlEntry" {
        for _, c := range usrhistorycontroltable.Usrhistorycontrolentry {
            if usrhistorycontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry{}
        usrhistorycontroltable.Usrhistorycontrolentry = append(usrhistorycontroltable.Usrhistorycontrolentry, child)
        return &usrhistorycontroltable.Usrhistorycontrolentry[len(usrhistorycontroltable.Usrhistorycontrolentry)-1]
    }
    return nil
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usrhistorycontroltable.Usrhistorycontrolentry {
        children[usrhistorycontroltable.Usrhistorycontrolentry[i].GetSegmentPath()] = &usrhistorycontroltable.Usrhistorycontrolentry[i]
    }
    return children
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetYangName() string { return "usrHistoryControlTable" }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) SetParent(parent types.Entity) { usrhistorycontroltable.parent = parent }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetParent() types.Entity { return usrhistorycontroltable.parent }

func (usrhistorycontroltable *RMON2MIB_Usrhistorycontroltable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry
// A list of parameters that set up a group of user-defined
// MIB objects to be sampled periodically (called a
// bucket-group).
// 
// For example, an instance of usrHistoryControlInterval
// might be named usrHistoryControlInterval.1
type RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry struct {
    parent types.Entity
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

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetFilter() yfilter.YFilter { return usrhistorycontrolentry.YFilter }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) SetFilter(yf yfilter.YFilter) { usrhistorycontrolentry.YFilter = yf }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetGoName(yname string) string {
    if yname == "usrHistoryControlIndex" { return "Usrhistorycontrolindex" }
    if yname == "usrHistoryControlObjects" { return "Usrhistorycontrolobjects" }
    if yname == "usrHistoryControlBucketsRequested" { return "Usrhistorycontrolbucketsrequested" }
    if yname == "usrHistoryControlBucketsGranted" { return "Usrhistorycontrolbucketsgranted" }
    if yname == "usrHistoryControlInterval" { return "Usrhistorycontrolinterval" }
    if yname == "usrHistoryControlOwner" { return "Usrhistorycontrolowner" }
    if yname == "usrHistoryControlStatus" { return "Usrhistorycontrolstatus" }
    return ""
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetSegmentPath() string {
    return "usrHistoryControlEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistorycontrolentry.Usrhistorycontrolindex) + "']"
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["usrHistoryControlIndex"] = usrhistorycontrolentry.Usrhistorycontrolindex
    leafs["usrHistoryControlObjects"] = usrhistorycontrolentry.Usrhistorycontrolobjects
    leafs["usrHistoryControlBucketsRequested"] = usrhistorycontrolentry.Usrhistorycontrolbucketsrequested
    leafs["usrHistoryControlBucketsGranted"] = usrhistorycontrolentry.Usrhistorycontrolbucketsgranted
    leafs["usrHistoryControlInterval"] = usrhistorycontrolentry.Usrhistorycontrolinterval
    leafs["usrHistoryControlOwner"] = usrhistorycontrolentry.Usrhistorycontrolowner
    leafs["usrHistoryControlStatus"] = usrhistorycontrolentry.Usrhistorycontrolstatus
    return leafs
}

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetYangName() string { return "usrHistoryControlEntry" }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) SetParent(parent types.Entity) { usrhistorycontrolentry.parent = parent }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetParent() types.Entity { return usrhistorycontrolentry.parent }

func (usrhistorycontrolentry *RMON2MIB_Usrhistorycontroltable_Usrhistorycontrolentry) GetParentYangName() string { return "usrHistoryControlTable" }

// RMON2MIB_Usrhistoryobjecttable
// A list of data-collection configuration entries.
type RMON2MIB_Usrhistoryobjecttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of MIB instances to be sampled periodically.  Entries in this table
    // are created when an associated usrHistoryControlObjects object is created. 
    // The usrHistoryControlIndex value in the index is that of the associated
    // usrHistoryControlEntry.  For example, an instance of
    // usrHistoryObjectVariable might be usrHistoryObjectVariable.1.3. The type is
    // slice of RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry.
    Usrhistoryobjectentry []RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetFilter() yfilter.YFilter { return usrhistoryobjecttable.YFilter }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) SetFilter(yf yfilter.YFilter) { usrhistoryobjecttable.YFilter = yf }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetGoName(yname string) string {
    if yname == "usrHistoryObjectEntry" { return "Usrhistoryobjectentry" }
    return ""
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetSegmentPath() string {
    return "usrHistoryObjectTable"
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usrHistoryObjectEntry" {
        for _, c := range usrhistoryobjecttable.Usrhistoryobjectentry {
            if usrhistoryobjecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry{}
        usrhistoryobjecttable.Usrhistoryobjectentry = append(usrhistoryobjecttable.Usrhistoryobjectentry, child)
        return &usrhistoryobjecttable.Usrhistoryobjectentry[len(usrhistoryobjecttable.Usrhistoryobjectentry)-1]
    }
    return nil
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usrhistoryobjecttable.Usrhistoryobjectentry {
        children[usrhistoryobjecttable.Usrhistoryobjectentry[i].GetSegmentPath()] = &usrhistoryobjecttable.Usrhistoryobjectentry[i]
    }
    return children
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetYangName() string { return "usrHistoryObjectTable" }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) SetParent(parent types.Entity) { usrhistoryobjecttable.parent = parent }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetParent() types.Entity { return usrhistoryobjecttable.parent }

func (usrhistoryobjecttable *RMON2MIB_Usrhistoryobjecttable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetFilter() yfilter.YFilter { return usrhistoryobjectentry.YFilter }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) SetFilter(yf yfilter.YFilter) { usrhistoryobjectentry.YFilter = yf }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetGoName(yname string) string {
    if yname == "usrHistoryControlIndex" { return "Usrhistorycontrolindex" }
    if yname == "usrHistoryObjectIndex" { return "Usrhistoryobjectindex" }
    if yname == "usrHistoryObjectVariable" { return "Usrhistoryobjectvariable" }
    if yname == "usrHistoryObjectSampleType" { return "Usrhistoryobjectsampletype" }
    return ""
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetSegmentPath() string {
    return "usrHistoryObjectEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistoryobjectentry.Usrhistorycontrolindex) + "']" + "[usrHistoryObjectIndex='" + fmt.Sprintf("%v", usrhistoryobjectentry.Usrhistoryobjectindex) + "']"
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["usrHistoryControlIndex"] = usrhistoryobjectentry.Usrhistorycontrolindex
    leafs["usrHistoryObjectIndex"] = usrhistoryobjectentry.Usrhistoryobjectindex
    leafs["usrHistoryObjectVariable"] = usrhistoryobjectentry.Usrhistoryobjectvariable
    leafs["usrHistoryObjectSampleType"] = usrhistoryobjectentry.Usrhistoryobjectsampletype
    return leafs
}

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetYangName() string { return "usrHistoryObjectEntry" }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) SetParent(parent types.Entity) { usrhistoryobjectentry.parent = parent }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetParent() types.Entity { return usrhistoryobjectentry.parent }

func (usrhistoryobjectentry *RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry) GetParentYangName() string { return "usrHistoryObjectTable" }

// RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype represents usrHistoryControlStatus object is equal to active(1).
type RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype string

const (
    RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype_absoluteValue RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype = "absoluteValue"

    RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype_deltaValue RMON2MIB_Usrhistoryobjecttable_Usrhistoryobjectentry_Usrhistoryobjectsampletype = "deltaValue"
)

// RMON2MIB_Usrhistorytable
// A list of user defined history entries.
type RMON2MIB_Usrhistorytable struct {
    parent types.Entity
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

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetFilter() yfilter.YFilter { return usrhistorytable.YFilter }

func (usrhistorytable *RMON2MIB_Usrhistorytable) SetFilter(yf yfilter.YFilter) { usrhistorytable.YFilter = yf }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetGoName(yname string) string {
    if yname == "usrHistoryEntry" { return "Usrhistoryentry" }
    return ""
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetSegmentPath() string {
    return "usrHistoryTable"
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usrHistoryEntry" {
        for _, c := range usrhistorytable.Usrhistoryentry {
            if usrhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Usrhistorytable_Usrhistoryentry{}
        usrhistorytable.Usrhistoryentry = append(usrhistorytable.Usrhistoryentry, child)
        return &usrhistorytable.Usrhistoryentry[len(usrhistorytable.Usrhistoryentry)-1]
    }
    return nil
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usrhistorytable.Usrhistoryentry {
        children[usrhistorytable.Usrhistoryentry[i].GetSegmentPath()] = &usrhistorytable.Usrhistoryentry[i]
    }
    return children
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetYangName() string { return "usrHistoryTable" }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistorytable *RMON2MIB_Usrhistorytable) SetParent(parent types.Entity) { usrhistorytable.parent = parent }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetParent() types.Entity { return usrhistorytable.parent }

func (usrhistorytable *RMON2MIB_Usrhistorytable) GetParentYangName() string { return "RMON2-MIB" }

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
    parent types.Entity
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

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetFilter() yfilter.YFilter { return usrhistoryentry.YFilter }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) SetFilter(yf yfilter.YFilter) { usrhistoryentry.YFilter = yf }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetGoName(yname string) string {
    if yname == "usrHistoryControlIndex" { return "Usrhistorycontrolindex" }
    if yname == "usrHistorySampleIndex" { return "Usrhistorysampleindex" }
    if yname == "usrHistoryObjectIndex" { return "Usrhistoryobjectindex" }
    if yname == "usrHistoryIntervalStart" { return "Usrhistoryintervalstart" }
    if yname == "usrHistoryIntervalEnd" { return "Usrhistoryintervalend" }
    if yname == "usrHistoryAbsValue" { return "Usrhistoryabsvalue" }
    if yname == "usrHistoryValStatus" { return "Usrhistoryvalstatus" }
    return ""
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetSegmentPath() string {
    return "usrHistoryEntry" + "[usrHistoryControlIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistorycontrolindex) + "']" + "[usrHistorySampleIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistorysampleindex) + "']" + "[usrHistoryObjectIndex='" + fmt.Sprintf("%v", usrhistoryentry.Usrhistoryobjectindex) + "']"
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["usrHistoryControlIndex"] = usrhistoryentry.Usrhistorycontrolindex
    leafs["usrHistorySampleIndex"] = usrhistoryentry.Usrhistorysampleindex
    leafs["usrHistoryObjectIndex"] = usrhistoryentry.Usrhistoryobjectindex
    leafs["usrHistoryIntervalStart"] = usrhistoryentry.Usrhistoryintervalstart
    leafs["usrHistoryIntervalEnd"] = usrhistoryentry.Usrhistoryintervalend
    leafs["usrHistoryAbsValue"] = usrhistoryentry.Usrhistoryabsvalue
    leafs["usrHistoryValStatus"] = usrhistoryentry.Usrhistoryvalstatus
    return leafs
}

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetYangName() string { return "usrHistoryEntry" }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) SetParent(parent types.Entity) { usrhistoryentry.parent = parent }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetParent() types.Entity { return usrhistoryentry.parent }

func (usrhistoryentry *RMON2MIB_Usrhistorytable_Usrhistoryentry) GetParentYangName() string { return "usrHistoryTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular serial interface on this
    // device. If the device has no serial interfaces, this table is empty.  The
    // index is composed of the ifIndex assigned to this serial line interface.
    // The type is slice of RMON2MIB_Serialconfigtable_Serialconfigentry.
    Serialconfigentry []RMON2MIB_Serialconfigtable_Serialconfigentry
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetFilter() yfilter.YFilter { return serialconfigtable.YFilter }

func (serialconfigtable *RMON2MIB_Serialconfigtable) SetFilter(yf yfilter.YFilter) { serialconfigtable.YFilter = yf }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetGoName(yname string) string {
    if yname == "serialConfigEntry" { return "Serialconfigentry" }
    return ""
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetSegmentPath() string {
    return "serialConfigTable"
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "serialConfigEntry" {
        for _, c := range serialconfigtable.Serialconfigentry {
            if serialconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Serialconfigtable_Serialconfigentry{}
        serialconfigtable.Serialconfigentry = append(serialconfigtable.Serialconfigentry, child)
        return &serialconfigtable.Serialconfigentry[len(serialconfigtable.Serialconfigentry)-1]
    }
    return nil
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serialconfigtable.Serialconfigentry {
        children[serialconfigtable.Serialconfigentry[i].GetSegmentPath()] = &serialconfigtable.Serialconfigentry[i]
    }
    return children
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetYangName() string { return "serialConfigTable" }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (serialconfigtable *RMON2MIB_Serialconfigtable) SetParent(parent types.Entity) { serialconfigtable.parent = parent }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetParent() types.Entity { return serialconfigtable.parent }

func (serialconfigtable *RMON2MIB_Serialconfigtable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Serialconfigtable_Serialconfigentry
// A set of configuration parameters for a particular
// serial interface on this device. If the device has no serial
// interfaces, this table is empty.
// 
// The index is composed of the ifIndex assigned to this serial
// line interface.
type RMON2MIB_Serialconfigtable_Serialconfigentry struct {
    parent types.Entity
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

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetFilter() yfilter.YFilter { return serialconfigentry.YFilter }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) SetFilter(yf yfilter.YFilter) { serialconfigentry.YFilter = yf }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "serialMode" { return "Serialmode" }
    if yname == "serialProtocol" { return "Serialprotocol" }
    if yname == "serialTimeout" { return "Serialtimeout" }
    if yname == "serialModemInitString" { return "Serialmodeminitstring" }
    if yname == "serialModemHangUpString" { return "Serialmodemhangupstring" }
    if yname == "serialModemConnectResp" { return "Serialmodemconnectresp" }
    if yname == "serialModemNoConnectResp" { return "Serialmodemnoconnectresp" }
    if yname == "serialDialoutTimeout" { return "Serialdialouttimeout" }
    if yname == "serialStatus" { return "Serialstatus" }
    return ""
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetSegmentPath() string {
    return "serialConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", serialconfigentry.Ifindex) + "']"
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = serialconfigentry.Ifindex
    leafs["serialMode"] = serialconfigentry.Serialmode
    leafs["serialProtocol"] = serialconfigentry.Serialprotocol
    leafs["serialTimeout"] = serialconfigentry.Serialtimeout
    leafs["serialModemInitString"] = serialconfigentry.Serialmodeminitstring
    leafs["serialModemHangUpString"] = serialconfigentry.Serialmodemhangupstring
    leafs["serialModemConnectResp"] = serialconfigentry.Serialmodemconnectresp
    leafs["serialModemNoConnectResp"] = serialconfigentry.Serialmodemnoconnectresp
    leafs["serialDialoutTimeout"] = serialconfigentry.Serialdialouttimeout
    leafs["serialStatus"] = serialconfigentry.Serialstatus
    return leafs
}

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetYangName() string { return "serialConfigEntry" }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) SetParent(parent types.Entity) { serialconfigentry.parent = parent }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetParent() types.Entity { return serialconfigentry.parent }

func (serialconfigentry *RMON2MIB_Serialconfigtable_Serialconfigentry) GetParentYangName() string { return "serialConfigTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of configuration parameters for a particular network interface on
    // this device. If the device has no network interface, this table is empty. 
    // The index is composed of the ifIndex assigned to the corresponding
    // interface. The type is slice of RMON2MIB_Netconfigtable_Netconfigentry.
    Netconfigentry []RMON2MIB_Netconfigtable_Netconfigentry
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetFilter() yfilter.YFilter { return netconfigtable.YFilter }

func (netconfigtable *RMON2MIB_Netconfigtable) SetFilter(yf yfilter.YFilter) { netconfigtable.YFilter = yf }

func (netconfigtable *RMON2MIB_Netconfigtable) GetGoName(yname string) string {
    if yname == "netConfigEntry" { return "Netconfigentry" }
    return ""
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetSegmentPath() string {
    return "netConfigTable"
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "netConfigEntry" {
        for _, c := range netconfigtable.Netconfigentry {
            if netconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Netconfigtable_Netconfigentry{}
        netconfigtable.Netconfigentry = append(netconfigtable.Netconfigentry, child)
        return &netconfigtable.Netconfigentry[len(netconfigtable.Netconfigentry)-1]
    }
    return nil
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range netconfigtable.Netconfigentry {
        children[netconfigtable.Netconfigentry[i].GetSegmentPath()] = &netconfigtable.Netconfigentry[i]
    }
    return children
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconfigtable *RMON2MIB_Netconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (netconfigtable *RMON2MIB_Netconfigtable) GetYangName() string { return "netConfigTable" }

func (netconfigtable *RMON2MIB_Netconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (netconfigtable *RMON2MIB_Netconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (netconfigtable *RMON2MIB_Netconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (netconfigtable *RMON2MIB_Netconfigtable) SetParent(parent types.Entity) { netconfigtable.parent = parent }

func (netconfigtable *RMON2MIB_Netconfigtable) GetParent() types.Entity { return netconfigtable.parent }

func (netconfigtable *RMON2MIB_Netconfigtable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Netconfigtable_Netconfigentry
// A set of configuration parameters for a particular
// network interface on this device. If the device has no network
// interface, this table is empty.
// 
// The index is composed of the ifIndex assigned to the
// corresponding interface.
type RMON2MIB_Netconfigtable_Netconfigentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netconfigipaddress interface{}

    // The subnet mask of this Net interface.  The default value for this object
    // is 0.0.0.0.  If either the netConfigIPAddress or netConfigSubnetMask are
    // 0.0.0.0, then when the device boots, it may use BOOTP to try to figure out
    // what these values should be. If BOOTP fails, before the device can talk on
    // the network, this value must be configured (e.g., through a terminal
    // attached to the device). If BOOTP is used, care should be taken to not send
    // BOOTP broadcasts too frequently and to eventually send very infrequently if
    // no replies are received. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netconfigsubnetmask interface{}

    // The status of this netConfigEntry.  An entry may not exist in the active
    // state unless all objects in the entry have an appropriate value. The type
    // is RowStatus.
    Netconfigstatus interface{}
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetFilter() yfilter.YFilter { return netconfigentry.YFilter }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) SetFilter(yf yfilter.YFilter) { netconfigentry.YFilter = yf }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "netConfigIPAddress" { return "Netconfigipaddress" }
    if yname == "netConfigSubnetMask" { return "Netconfigsubnetmask" }
    if yname == "netConfigStatus" { return "Netconfigstatus" }
    return ""
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetSegmentPath() string {
    return "netConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", netconfigentry.Ifindex) + "']"
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = netconfigentry.Ifindex
    leafs["netConfigIPAddress"] = netconfigentry.Netconfigipaddress
    leafs["netConfigSubnetMask"] = netconfigentry.Netconfigsubnetmask
    leafs["netConfigStatus"] = netconfigentry.Netconfigstatus
    return leafs
}

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetYangName() string { return "netConfigEntry" }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) SetParent(parent types.Entity) { netconfigentry.parent = parent }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetParent() types.Entity { return netconfigentry.parent }

func (netconfigentry *RMON2MIB_Netconfigtable_Netconfigentry) GetParentYangName() string { return "netConfigTable" }

// RMON2MIB_Trapdesttable
// A list of trap destination entries.
type RMON2MIB_Trapdesttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry includes a destination IP address to which to send traps for
    // this community. The type is slice of RMON2MIB_Trapdesttable_Trapdestentry.
    Trapdestentry []RMON2MIB_Trapdesttable_Trapdestentry
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetFilter() yfilter.YFilter { return trapdesttable.YFilter }

func (trapdesttable *RMON2MIB_Trapdesttable) SetFilter(yf yfilter.YFilter) { trapdesttable.YFilter = yf }

func (trapdesttable *RMON2MIB_Trapdesttable) GetGoName(yname string) string {
    if yname == "trapDestEntry" { return "Trapdestentry" }
    return ""
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetSegmentPath() string {
    return "trapDestTable"
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trapDestEntry" {
        for _, c := range trapdesttable.Trapdestentry {
            if trapdesttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Trapdesttable_Trapdestentry{}
        trapdesttable.Trapdestentry = append(trapdesttable.Trapdestentry, child)
        return &trapdesttable.Trapdestentry[len(trapdesttable.Trapdestentry)-1]
    }
    return nil
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapdesttable.Trapdestentry {
        children[trapdesttable.Trapdestentry[i].GetSegmentPath()] = &trapdesttable.Trapdestentry[i]
    }
    return children
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapdesttable *RMON2MIB_Trapdesttable) GetBundleName() string { return "cisco_ios_xe" }

func (trapdesttable *RMON2MIB_Trapdesttable) GetYangName() string { return "trapDestTable" }

func (trapdesttable *RMON2MIB_Trapdesttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (trapdesttable *RMON2MIB_Trapdesttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (trapdesttable *RMON2MIB_Trapdesttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (trapdesttable *RMON2MIB_Trapdesttable) SetParent(parent types.Entity) { trapdesttable.parent = parent }

func (trapdesttable *RMON2MIB_Trapdesttable) GetParent() types.Entity { return trapdesttable.parent }

func (trapdesttable *RMON2MIB_Trapdesttable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Trapdesttable_Trapdestentry
// This entry includes a destination IP address to which to send
// traps for this community.
type RMON2MIB_Trapdesttable_Trapdestentry struct {
    parent types.Entity
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

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetFilter() yfilter.YFilter { return trapdestentry.YFilter }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) SetFilter(yf yfilter.YFilter) { trapdestentry.YFilter = yf }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetGoName(yname string) string {
    if yname == "trapDestIndex" { return "Trapdestindex" }
    if yname == "trapDestCommunity" { return "Trapdestcommunity" }
    if yname == "trapDestProtocol" { return "Trapdestprotocol" }
    if yname == "trapDestAddress" { return "Trapdestaddress" }
    if yname == "trapDestOwner" { return "Trapdestowner" }
    if yname == "trapDestStatus" { return "Trapdeststatus" }
    return ""
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetSegmentPath() string {
    return "trapDestEntry" + "[trapDestIndex='" + fmt.Sprintf("%v", trapdestentry.Trapdestindex) + "']"
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trapDestIndex"] = trapdestentry.Trapdestindex
    leafs["trapDestCommunity"] = trapdestentry.Trapdestcommunity
    leafs["trapDestProtocol"] = trapdestentry.Trapdestprotocol
    leafs["trapDestAddress"] = trapdestentry.Trapdestaddress
    leafs["trapDestOwner"] = trapdestentry.Trapdestowner
    leafs["trapDestStatus"] = trapdestentry.Trapdeststatus
    return leafs
}

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetBundleName() string { return "cisco_ios_xe" }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetYangName() string { return "trapDestEntry" }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) SetParent(parent types.Entity) { trapdestentry.parent = parent }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetParent() types.Entity { return trapdestentry.parent }

func (trapdestentry *RMON2MIB_Trapdesttable_Trapdestentry) GetParentYangName() string { return "trapDestTable" }

// RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol represents The protocol with which to send this trap.
type RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol string

const (
    RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol_ip RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol = "ip"

    RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol_ipx RMON2MIB_Trapdesttable_Trapdestentry_Trapdestprotocol = "ipx"
)

// RMON2MIB_Serialconnectiontable
// A list of serialConnectionEntries.
type RMON2MIB_Serialconnectiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a SLIP link over a serial line. The type is slice of
    // RMON2MIB_Serialconnectiontable_Serialconnectionentry.
    Serialconnectionentry []RMON2MIB_Serialconnectiontable_Serialconnectionentry
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetFilter() yfilter.YFilter { return serialconnectiontable.YFilter }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) SetFilter(yf yfilter.YFilter) { serialconnectiontable.YFilter = yf }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetGoName(yname string) string {
    if yname == "serialConnectionEntry" { return "Serialconnectionentry" }
    return ""
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetSegmentPath() string {
    return "serialConnectionTable"
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "serialConnectionEntry" {
        for _, c := range serialconnectiontable.Serialconnectionentry {
            if serialconnectiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMON2MIB_Serialconnectiontable_Serialconnectionentry{}
        serialconnectiontable.Serialconnectionentry = append(serialconnectiontable.Serialconnectionentry, child)
        return &serialconnectiontable.Serialconnectionentry[len(serialconnectiontable.Serialconnectionentry)-1]
    }
    return nil
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serialconnectiontable.Serialconnectionentry {
        children[serialconnectiontable.Serialconnectionentry[i].GetSegmentPath()] = &serialconnectiontable.Serialconnectionentry[i]
    }
    return children
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetBundleName() string { return "cisco_ios_xe" }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetYangName() string { return "serialConnectionTable" }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) SetParent(parent types.Entity) { serialconnectiontable.parent = parent }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetParent() types.Entity { return serialconnectiontable.parent }

func (serialconnectiontable *RMON2MIB_Serialconnectiontable) GetParentYangName() string { return "RMON2-MIB" }

// RMON2MIB_Serialconnectiontable_Serialconnectionentry
// Configuration for a SLIP link over a serial line.
type RMON2MIB_Serialconnectiontable_Serialconnectionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A value that uniquely identifies this
    // serialConnection entry. The type is interface{} with range: 1..65535.
    Serialconnectindex interface{}

    // The IP Address that can be reached at the other end of this serial
    // connection. This object may not be modified if the associated
    // serialConnectStatus object is equal to active(1). The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetFilter() yfilter.YFilter { return serialconnectionentry.YFilter }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) SetFilter(yf yfilter.YFilter) { serialconnectionentry.YFilter = yf }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetGoName(yname string) string {
    if yname == "serialConnectIndex" { return "Serialconnectindex" }
    if yname == "serialConnectDestIpAddress" { return "Serialconnectdestipaddress" }
    if yname == "serialConnectType" { return "Serialconnecttype" }
    if yname == "serialConnectDialString" { return "Serialconnectdialstring" }
    if yname == "serialConnectSwitchConnectSeq" { return "Serialconnectswitchconnectseq" }
    if yname == "serialConnectSwitchDisconnectSeq" { return "Serialconnectswitchdisconnectseq" }
    if yname == "serialConnectSwitchResetSeq" { return "Serialconnectswitchresetseq" }
    if yname == "serialConnectOwner" { return "Serialconnectowner" }
    if yname == "serialConnectStatus" { return "Serialconnectstatus" }
    return ""
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetSegmentPath() string {
    return "serialConnectionEntry" + "[serialConnectIndex='" + fmt.Sprintf("%v", serialconnectionentry.Serialconnectindex) + "']"
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["serialConnectIndex"] = serialconnectionentry.Serialconnectindex
    leafs["serialConnectDestIpAddress"] = serialconnectionentry.Serialconnectdestipaddress
    leafs["serialConnectType"] = serialconnectionentry.Serialconnecttype
    leafs["serialConnectDialString"] = serialconnectionentry.Serialconnectdialstring
    leafs["serialConnectSwitchConnectSeq"] = serialconnectionentry.Serialconnectswitchconnectseq
    leafs["serialConnectSwitchDisconnectSeq"] = serialconnectionentry.Serialconnectswitchdisconnectseq
    leafs["serialConnectSwitchResetSeq"] = serialconnectionentry.Serialconnectswitchresetseq
    leafs["serialConnectOwner"] = serialconnectionentry.Serialconnectowner
    leafs["serialConnectStatus"] = serialconnectionentry.Serialconnectstatus
    return leafs
}

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetBundleName() string { return "cisco_ios_xe" }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetYangName() string { return "serialConnectionEntry" }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) SetParent(parent types.Entity) { serialconnectionentry.parent = parent }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetParent() types.Entity { return serialconnectionentry.parent }

func (serialconnectionentry *RMON2MIB_Serialconnectiontable_Serialconnectionentry) GetParentYangName() string { return "serialConnectionTable" }

// RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype represents serialConnectStatus object is equal to active(1).
type RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype string

const (
    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_direct RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "direct"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_modem RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "modem"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_switch_ RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "switch"

    RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype_modemSwitch RMON2MIB_Serialconnectiontable_Serialconnectionentry_Serialconnecttype = "modemSwitch"
)

