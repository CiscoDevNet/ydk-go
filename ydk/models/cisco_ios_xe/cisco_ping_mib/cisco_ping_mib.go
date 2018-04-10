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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table of ping request entries.
    Ciscopingtable CISCOPINGMIB_Ciscopingtable
}

func (cISCOPINGMIB *CISCOPINGMIB) GetEntityData() *types.CommonEntityData {
    cISCOPINGMIB.EntityData.YFilter = cISCOPINGMIB.YFilter
    cISCOPINGMIB.EntityData.YangName = "CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPINGMIB.EntityData.ParentYangName = "CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.SegmentPath = "CISCO-PING-MIB:CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPINGMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOPINGMIB.EntityData.Children["ciscoPingTable"] = types.YChild{"Ciscopingtable", &cISCOPINGMIB.Ciscopingtable}
    cISCOPINGMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOPINGMIB.EntityData)
}

// CISCOPINGMIB_Ciscopingtable
// A table of ping request entries.
type CISCOPINGMIB_Ciscopingtable struct {
    EntityData types.CommonEntityData
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

func (ciscopingtable *CISCOPINGMIB_Ciscopingtable) GetEntityData() *types.CommonEntityData {
    ciscopingtable.EntityData.YFilter = ciscopingtable.YFilter
    ciscopingtable.EntityData.YangName = "ciscoPingTable"
    ciscopingtable.EntityData.BundleName = "cisco_ios_xe"
    ciscopingtable.EntityData.ParentYangName = "CISCO-PING-MIB"
    ciscopingtable.EntityData.SegmentPath = "ciscoPingTable"
    ciscopingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscopingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscopingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscopingtable.EntityData.Children = make(map[string]types.YChild)
    ciscopingtable.EntityData.Children["ciscoPingEntry"] = types.YChild{"Ciscopingentry", nil}
    for i := range ciscopingtable.Ciscopingentry {
        ciscopingtable.EntityData.Children[types.GetSegmentPath(&ciscopingtable.Ciscopingentry[i])] = types.YChild{"Ciscopingentry", &ciscopingtable.Ciscopingentry[i]}
    }
    ciscopingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscopingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ciscopingentry *CISCOPINGMIB_Ciscopingtable_Ciscopingentry) GetEntityData() *types.CommonEntityData {
    ciscopingentry.EntityData.YFilter = ciscopingentry.YFilter
    ciscopingentry.EntityData.YangName = "ciscoPingEntry"
    ciscopingentry.EntityData.BundleName = "cisco_ios_xe"
    ciscopingentry.EntityData.ParentYangName = "ciscoPingTable"
    ciscopingentry.EntityData.SegmentPath = "ciscoPingEntry" + "[ciscoPingSerialNumber='" + fmt.Sprintf("%v", ciscopingentry.Ciscopingserialnumber) + "']"
    ciscopingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscopingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscopingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscopingentry.EntityData.Children = make(map[string]types.YChild)
    ciscopingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscopingentry.EntityData.Leafs["ciscoPingSerialNumber"] = types.YLeaf{"Ciscopingserialnumber", ciscopingentry.Ciscopingserialnumber}
    ciscopingentry.EntityData.Leafs["ciscoPingProtocol"] = types.YLeaf{"Ciscopingprotocol", ciscopingentry.Ciscopingprotocol}
    ciscopingentry.EntityData.Leafs["ciscoPingAddress"] = types.YLeaf{"Ciscopingaddress", ciscopingentry.Ciscopingaddress}
    ciscopingentry.EntityData.Leafs["ciscoPingPacketCount"] = types.YLeaf{"Ciscopingpacketcount", ciscopingentry.Ciscopingpacketcount}
    ciscopingentry.EntityData.Leafs["ciscoPingPacketSize"] = types.YLeaf{"Ciscopingpacketsize", ciscopingentry.Ciscopingpacketsize}
    ciscopingentry.EntityData.Leafs["ciscoPingPacketTimeout"] = types.YLeaf{"Ciscopingpackettimeout", ciscopingentry.Ciscopingpackettimeout}
    ciscopingentry.EntityData.Leafs["ciscoPingDelay"] = types.YLeaf{"Ciscopingdelay", ciscopingentry.Ciscopingdelay}
    ciscopingentry.EntityData.Leafs["ciscoPingTrapOnCompletion"] = types.YLeaf{"Ciscopingtraponcompletion", ciscopingentry.Ciscopingtraponcompletion}
    ciscopingentry.EntityData.Leafs["ciscoPingSentPackets"] = types.YLeaf{"Ciscopingsentpackets", ciscopingentry.Ciscopingsentpackets}
    ciscopingentry.EntityData.Leafs["ciscoPingReceivedPackets"] = types.YLeaf{"Ciscopingreceivedpackets", ciscopingentry.Ciscopingreceivedpackets}
    ciscopingentry.EntityData.Leafs["ciscoPingMinRtt"] = types.YLeaf{"Ciscopingminrtt", ciscopingentry.Ciscopingminrtt}
    ciscopingentry.EntityData.Leafs["ciscoPingAvgRtt"] = types.YLeaf{"Ciscopingavgrtt", ciscopingentry.Ciscopingavgrtt}
    ciscopingentry.EntityData.Leafs["ciscoPingMaxRtt"] = types.YLeaf{"Ciscopingmaxrtt", ciscopingentry.Ciscopingmaxrtt}
    ciscopingentry.EntityData.Leafs["ciscoPingCompleted"] = types.YLeaf{"Ciscopingcompleted", ciscopingentry.Ciscopingcompleted}
    ciscopingentry.EntityData.Leafs["ciscoPingEntryOwner"] = types.YLeaf{"Ciscopingentryowner", ciscopingentry.Ciscopingentryowner}
    ciscopingentry.EntityData.Leafs["ciscoPingEntryStatus"] = types.YLeaf{"Ciscopingentrystatus", ciscopingentry.Ciscopingentrystatus}
    ciscopingentry.EntityData.Leafs["ciscoPingVrfName"] = types.YLeaf{"Ciscopingvrfname", ciscopingentry.Ciscopingvrfname}
    return &(ciscopingentry.EntityData)
}

