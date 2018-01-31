// Modified description of ciscoPingAddress object.
package cisco_ping_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ping_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-PING-MIB CISCO-PING-MIB}", reflect.TypeOf(CISCOPINGMIB{}))
    ydk.RegisterEntity("CISCO-PING-MIB:CISCO-PING-MIB", reflect.TypeOf(CISCOPINGMIB{}))
}

// CISCOPINGMIB
type CISCOPINGMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table of ping request entries.
    Ciscopingtable CISCOPINGMIB_Ciscopingtable
}

func (cISCOPINGMIB *CISCOPINGMIB) GetFilter() yfilter.YFilter { return cISCOPINGMIB.YFilter }

func (cISCOPINGMIB *CISCOPINGMIB) SetFilter(yf yfilter.YFilter) { cISCOPINGMIB.YFilter = yf }

func (cISCOPINGMIB *CISCOPINGMIB) GetGoName(yname string) string {
    if yname == "ciscoPingTable" { return "Ciscopingtable" }
    return ""
}

func (cISCOPINGMIB *CISCOPINGMIB) GetSegmentPath() string {
    return "CISCO-PING-MIB:CISCO-PING-MIB"
}

func (cISCOPINGMIB *CISCOPINGMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoPingTable" {
        return &cISCOPINGMIB.Ciscopingtable
    }
    return nil
}

func (cISCOPINGMIB *CISCOPINGMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoPingTable"] = &cISCOPINGMIB.Ciscopingtable
    return children
}

func (cISCOPINGMIB *CISCOPINGMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOPINGMIB *CISCOPINGMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOPINGMIB *CISCOPINGMIB) GetYangName() string { return "CISCO-PING-MIB" }

func (cISCOPINGMIB *CISCOPINGMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOPINGMIB *CISCOPINGMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOPINGMIB *CISCOPINGMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOPINGMIB *CISCOPINGMIB) SetParent(parent types.Entity) { cISCOPINGMIB.parent = parent }

func (cISCOPINGMIB *CISCOPINGMIB) GetParent() types.Entity { return cISCOPINGMIB.parent }

func (cISCOPINGMIB *CISCOPINGMIB) GetParentYangName() string { return "CISCO-PING-MIB" }

// CISCOPINGMIB_Ciscopingtable
// A table of ping request entries.
type CISCOPINGMIB_Ciscopingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A ping request entry.  A management station wishing to create an entry
    // should first generate a pseudo-random serial number to be used as the index
    // to this sparse table.  The station should then create the associated
    // instance of the row status and row owner objects.  It must also, either in
    // the same or in successive PDUs, create the associated instance of the
    // protocol and address objects.  It should also modify the default values for
    // the other configuration objects if the defaults are not appropriate.  Once
    // the appropriate instance of all the configuration objects have been
    // created, either by an explicit SNMP set request or by default, the row
    // status should be set to active to initiate the request.  Note that this
    // entire procedure may be initiated via a single set request which specifies
    // a row status of createAndGo as well as specifies valid values for the
    // non-defaulted configuration objects.  Once the ping sequence has been
    // activated, it cannot be stopped -- it will run until the configured number
    // of packets have been sent.  Once the sequence completes, the management
    // station should retrieve the values of the status objects of interest, and
    // should then delete the entry.  In order to prevent old entries from
    // clogging the table, entries will be aged out, but an entry will never be
    // deleted within 5 minutes of completing. The type is slice of
    // CISCOPINGMIB_Ciscopingtable_Ciscopingentry.
    Ciscopingentry []CISCOPINGMIB_Ciscopingtable_Ciscopingentry
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetFilter() yfilter.YFilter { return ciscopingtable.YFilter }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) SetFilter(yf yfilter.YFilter) { ciscopingtable.YFilter = yf }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetGoName(yname string) string {
    if yname == "ciscoPingEntry" { return "Ciscopingentry" }
    return ""
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetSegmentPath() string {
    return "ciscoPingTable"
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoPingEntry" {
        for _, c := range ciscopingtable.Ciscopingentry {
            if ciscopingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPINGMIB_Ciscopingtable_Ciscopingentry{}
        ciscopingtable.Ciscopingentry = append(ciscopingtable.Ciscopingentry, child)
        return &ciscopingtable.Ciscopingentry[len(ciscopingtable.Ciscopingentry)-1]
    }
    return nil
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscopingtable.Ciscopingentry {
        children[ciscopingtable.Ciscopingentry[i].GetSegmentPath()] = &ciscopingtable.Ciscopingentry[i]
    }
    return children
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetYangName() string { return "ciscoPingTable" }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) SetParent(parent types.Entity) { ciscopingtable.parent = parent }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetParent() types.Entity { return ciscopingtable.parent }

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetParentYangName() string { return "CISCO-PING-MIB" }

// CISCOPINGMIB_Ciscopingtable_Ciscopingentry
// A ping request entry.
// 
// A management station wishing to create an entry should
// first generate a pseudo-random serial number to be used
// as the index to this sparse table.  The station should
// then create the associated instance of the row status
// and row owner objects.  It must also, either in the same
// or in successive PDUs, create the associated instance of
// the protocol and address objects.  It should also modify
// the default values for the other configuration objects
// if the defaults are not appropriate.
// 
// Once the appropriate instance of all the configuration
// objects have been created, either by an explicit SNMP
// set request or by default, the row status should be set
// to active to initiate the request.  Note that this entire
// procedure may be initiated via a single set request which
// specifies a row status of createAndGo as well as specifies
// valid values for the non-defaulted configuration objects.
// 
// Once the ping sequence has been activated, it cannot be
// stopped -- it will run until the configured number of
// packets have been sent.
// 
// Once the sequence completes, the management station should
// retrieve the values of the status objects of interest, and
// should then delete the entry.  In order to prevent old
// entries from clogging the table, entries will be aged out,
// but an entry will never be deleted within 5 minutes of
// completing.
type CISCOPINGMIB_Ciscopingtable_Ciscopingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object which specifies a unique entry in the
    // ciscoPingTable.  A management station wishing to initiate a ping operation
    // should use a pseudo-random value for this object when creating or modifying
    // an instance of a ciscoPingEntry. The RowStatus semantics of the
    // ciscoPingEntryStatus object will prevent access conflicts. The type is
    // interface{} with range: 1..2147483647.
    Ciscopingserialnumber interface{}

    // The protocol to use. Once an instance of this object is created, its value
    // can not be changed. The type is CiscoNetworkProtocol.
    Ciscopingprotocol interface{}

    // The address of the device to be pinged. An instance of this object cannot
    // be created until the associated instance of ciscoPingProtocol is created.
    // The type is string.
    Ciscopingaddress interface{}

    // Specifies the number of ping packets to send to the target in this
    // sequence. The type is interface{} with range: 1..2147483647.
    Ciscopingpacketcount interface{}

    // Specifies the size of ping packets to send to the target in this sequence. 
    // The lower and upper boundaries of this object are protocol-dependent. An
    // instance of this object cannot be modified unless the associated instance
    // of ciscoPingProtocol has been created (so as to allow protocol-specific
    // range checking on the new value). The type is interface{} with range:
    // -2147483648..2147483647.
    Ciscopingpacketsize interface{}

    // Specifies the amount of time to wait for a response to a transmitted packet
    // before declaring the packet 'dropped.'. The type is interface{} with range:
    // 0..3600000. Units are milliseconds.
    Ciscopingpackettimeout interface{}

    // Specifies the minimum amount of time to wait before sending the next packet
    // in a sequence after receiving a response or declaring a timeout for a
    // previous packet.  The actual delay may be greater due to internal task
    // scheduling. The type is interface{} with range: 0..3600000. Units are
    // milliseconds.
    Ciscopingdelay interface{}

    // Specifies whether or not a ciscoPingCompletion trap should be issued on
    // completion of the sequence of pings.  If such a trap is desired, it is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the trap to be
    // delivered. The type is bool.
    Ciscopingtraponcompletion interface{}

    // The number of ping packets that have been sent to the target in this
    // sequence. The type is interface{} with range: 0..4294967295.
    Ciscopingsentpackets interface{}

    // The number of ping packets that have been received from the target in this
    // sequence. The type is interface{} with range: 0..4294967295.
    Ciscopingreceivedpackets interface{}

    // The minimum round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    Ciscopingminrtt interface{}

    // The average round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    Ciscopingavgrtt interface{}

    // The maximum round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    Ciscopingmaxrtt interface{}

    // Set to true when all the packets in this sequence have been either
    // responded to or timed out. The type is bool.
    Ciscopingcompleted interface{}

    // The entity that configured this entry. The type is string.
    Ciscopingentryowner interface{}

    // The status of this table entry.  Once the entry status is set to active,
    // the associate entry cannot be modified until the sequence completes
    // (ciscoPingCompleted is true). The type is RowStatus.
    Ciscopingentrystatus interface{}

    // This field is used to specify the VPN name in  which the ping will be used.
    // For regular ping this field should not be configured. The agent will use
    // this field to identify the VPN routing Table for  this ping. This is the
    // same ascii string used in  the CLI to refer to this VPN. . The type is
    // string with length: 0..32.
    Ciscopingvrfname interface{}
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetFilter() yfilter.YFilter { return ciscopingentry.YFilter }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) SetFilter(yf yfilter.YFilter) { ciscopingentry.YFilter = yf }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetGoName(yname string) string {
    if yname == "ciscoPingSerialNumber" { return "Ciscopingserialnumber" }
    if yname == "ciscoPingProtocol" { return "Ciscopingprotocol" }
    if yname == "ciscoPingAddress" { return "Ciscopingaddress" }
    if yname == "ciscoPingPacketCount" { return "Ciscopingpacketcount" }
    if yname == "ciscoPingPacketSize" { return "Ciscopingpacketsize" }
    if yname == "ciscoPingPacketTimeout" { return "Ciscopingpackettimeout" }
    if yname == "ciscoPingDelay" { return "Ciscopingdelay" }
    if yname == "ciscoPingTrapOnCompletion" { return "Ciscopingtraponcompletion" }
    if yname == "ciscoPingSentPackets" { return "Ciscopingsentpackets" }
    if yname == "ciscoPingReceivedPackets" { return "Ciscopingreceivedpackets" }
    if yname == "ciscoPingMinRtt" { return "Ciscopingminrtt" }
    if yname == "ciscoPingAvgRtt" { return "Ciscopingavgrtt" }
    if yname == "ciscoPingMaxRtt" { return "Ciscopingmaxrtt" }
    if yname == "ciscoPingCompleted" { return "Ciscopingcompleted" }
    if yname == "ciscoPingEntryOwner" { return "Ciscopingentryowner" }
    if yname == "ciscoPingEntryStatus" { return "Ciscopingentrystatus" }
    if yname == "ciscoPingVrfName" { return "Ciscopingvrfname" }
    return ""
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetSegmentPath() string {
    return "ciscoPingEntry" + "[ciscoPingSerialNumber='" + fmt.Sprintf("%v", ciscopingentry.Ciscopingserialnumber) + "']"
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoPingSerialNumber"] = ciscopingentry.Ciscopingserialnumber
    leafs["ciscoPingProtocol"] = ciscopingentry.Ciscopingprotocol
    leafs["ciscoPingAddress"] = ciscopingentry.Ciscopingaddress
    leafs["ciscoPingPacketCount"] = ciscopingentry.Ciscopingpacketcount
    leafs["ciscoPingPacketSize"] = ciscopingentry.Ciscopingpacketsize
    leafs["ciscoPingPacketTimeout"] = ciscopingentry.Ciscopingpackettimeout
    leafs["ciscoPingDelay"] = ciscopingentry.Ciscopingdelay
    leafs["ciscoPingTrapOnCompletion"] = ciscopingentry.Ciscopingtraponcompletion
    leafs["ciscoPingSentPackets"] = ciscopingentry.Ciscopingsentpackets
    leafs["ciscoPingReceivedPackets"] = ciscopingentry.Ciscopingreceivedpackets
    leafs["ciscoPingMinRtt"] = ciscopingentry.Ciscopingminrtt
    leafs["ciscoPingAvgRtt"] = ciscopingentry.Ciscopingavgrtt
    leafs["ciscoPingMaxRtt"] = ciscopingentry.Ciscopingmaxrtt
    leafs["ciscoPingCompleted"] = ciscopingentry.Ciscopingcompleted
    leafs["ciscoPingEntryOwner"] = ciscopingentry.Ciscopingentryowner
    leafs["ciscoPingEntryStatus"] = ciscopingentry.Ciscopingentrystatus
    leafs["ciscoPingVrfName"] = ciscopingentry.Ciscopingvrfname
    return leafs
}

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetYangName() string { return "ciscoPingEntry" }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) SetParent(parent types.Entity) { ciscopingentry.parent = parent }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetParent() types.Entity { return ciscopingentry.parent }

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetParentYangName() string { return "ciscoPingTable" }

