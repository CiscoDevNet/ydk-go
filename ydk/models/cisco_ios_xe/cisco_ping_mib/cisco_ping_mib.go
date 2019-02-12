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
    CiscoPingTable CISCOPINGMIB_CiscoPingTable
}

func (cISCOPINGMIB *CISCOPINGMIB) GetEntityData() *types.CommonEntityData {
    cISCOPINGMIB.EntityData.YFilter = cISCOPINGMIB.YFilter
    cISCOPINGMIB.EntityData.YangName = "CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPINGMIB.EntityData.ParentYangName = "CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.SegmentPath = "CISCO-PING-MIB:CISCO-PING-MIB"
    cISCOPINGMIB.EntityData.AbsolutePath = cISCOPINGMIB.EntityData.SegmentPath
    cISCOPINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPINGMIB.EntityData.Children = types.NewOrderedMap()
    cISCOPINGMIB.EntityData.Children.Append("ciscoPingTable", types.YChild{"CiscoPingTable", &cISCOPINGMIB.CiscoPingTable})
    cISCOPINGMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOPINGMIB.EntityData.YListKeys = []string {}

    return &(cISCOPINGMIB.EntityData)
}

// CISCOPINGMIB_CiscoPingTable
// A table of ping request entries.
type CISCOPINGMIB_CiscoPingTable struct {
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
    // CISCOPINGMIB_CiscoPingTable_CiscoPingEntry.
    CiscoPingEntry []*CISCOPINGMIB_CiscoPingTable_CiscoPingEntry
}

func (ciscoPingTable *CISCOPINGMIB_CiscoPingTable) GetEntityData() *types.CommonEntityData {
    ciscoPingTable.EntityData.YFilter = ciscoPingTable.YFilter
    ciscoPingTable.EntityData.YangName = "ciscoPingTable"
    ciscoPingTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoPingTable.EntityData.ParentYangName = "CISCO-PING-MIB"
    ciscoPingTable.EntityData.SegmentPath = "ciscoPingTable"
    ciscoPingTable.EntityData.AbsolutePath = "CISCO-PING-MIB:CISCO-PING-MIB/" + ciscoPingTable.EntityData.SegmentPath
    ciscoPingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoPingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoPingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoPingTable.EntityData.Children = types.NewOrderedMap()
    ciscoPingTable.EntityData.Children.Append("ciscoPingEntry", types.YChild{"CiscoPingEntry", nil})
    for i := range ciscoPingTable.CiscoPingEntry {
        ciscoPingTable.EntityData.Children.Append(types.GetSegmentPath(ciscoPingTable.CiscoPingEntry[i]), types.YChild{"CiscoPingEntry", ciscoPingTable.CiscoPingEntry[i]})
    }
    ciscoPingTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoPingTable.EntityData.YListKeys = []string {}

    return &(ciscoPingTable.EntityData)
}

// CISCOPINGMIB_CiscoPingTable_CiscoPingEntry
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
type CISCOPINGMIB_CiscoPingTable_CiscoPingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // ciscoPingTable.  A management station wishing to initiate a ping operation
    // should use a pseudo-random value for this object when creating or modifying
    // an instance of a ciscoPingEntry. The RowStatus semantics of the
    // ciscoPingEntryStatus object will prevent access conflicts. The type is
    // interface{} with range: 1..2147483647.
    CiscoPingSerialNumber interface{}

    // The protocol to use. Once an instance of this object is created, its value
    // can not be changed. The type is CiscoNetworkProtocol.
    CiscoPingProtocol interface{}

    // The address of the device to be pinged. An instance of this object cannot
    // be created until the associated instance of ciscoPingProtocol is created.
    // The type is string.
    CiscoPingAddress interface{}

    // Specifies the number of ping packets to send to the target in this
    // sequence. The type is interface{} with range: 1..2147483647.
    CiscoPingPacketCount interface{}

    // Specifies the size of ping packets to send to the target in this sequence. 
    // The lower and upper boundaries of this object are protocol-dependent. An
    // instance of this object cannot be modified unless the associated instance
    // of ciscoPingProtocol has been created (so as to allow protocol-specific
    // range checking on the new value). The type is interface{} with range:
    // -2147483648..2147483647.
    CiscoPingPacketSize interface{}

    // Specifies the amount of time to wait for a response to a transmitted packet
    // before declaring the packet 'dropped.'. The type is interface{} with range:
    // 0..3600000. Units are milliseconds.
    CiscoPingPacketTimeout interface{}

    // Specifies the minimum amount of time to wait before sending the next packet
    // in a sequence after receiving a response or declaring a timeout for a
    // previous packet.  The actual delay may be greater due to internal task
    // scheduling. The type is interface{} with range: 0..3600000. Units are
    // milliseconds.
    CiscoPingDelay interface{}

    // Specifies whether or not a ciscoPingCompletion trap should be issued on
    // completion of the sequence of pings.  If such a trap is desired, it is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the trap to be
    // delivered. The type is bool.
    CiscoPingTrapOnCompletion interface{}

    // The number of ping packets that have been sent to the target in this
    // sequence. The type is interface{} with range: 0..4294967295.
    CiscoPingSentPackets interface{}

    // The number of ping packets that have been received from the target in this
    // sequence. The type is interface{} with range: 0..4294967295.
    CiscoPingReceivedPackets interface{}

    // The minimum round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    CiscoPingMinRtt interface{}

    // The average round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    CiscoPingAvgRtt interface{}

    // The maximum round trip time of all the packets that have been sent in this
    // sequence.  This object will not be created until the first ping response in
    // a sequence is received. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    CiscoPingMaxRtt interface{}

    // Set to true when all the packets in this sequence have been either
    // responded to or timed out. The type is bool.
    CiscoPingCompleted interface{}

    // The entity that configured this entry. The type is string.
    CiscoPingEntryOwner interface{}

    // The status of this table entry.  Once the entry status is set to active,
    // the associate entry cannot be modified until the sequence completes
    // (ciscoPingCompleted is true). The type is RowStatus.
    CiscoPingEntryStatus interface{}

    // This field is used to specify the VPN name in  which the ping will be used.
    // For regular ping this field should not be configured. The agent will use
    // this field to identify the VPN routing Table for  this ping. This is the
    // same ascii string used in  the CLI to refer to this VPN. . The type is
    // string with length: 0..32.
    CiscoPingVrfName interface{}
}

func (ciscoPingEntry *CISCOPINGMIB_CiscoPingTable_CiscoPingEntry) GetEntityData() *types.CommonEntityData {
    ciscoPingEntry.EntityData.YFilter = ciscoPingEntry.YFilter
    ciscoPingEntry.EntityData.YangName = "ciscoPingEntry"
    ciscoPingEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoPingEntry.EntityData.ParentYangName = "ciscoPingTable"
    ciscoPingEntry.EntityData.SegmentPath = "ciscoPingEntry" + types.AddKeyToken(ciscoPingEntry.CiscoPingSerialNumber, "ciscoPingSerialNumber")
    ciscoPingEntry.EntityData.AbsolutePath = "CISCO-PING-MIB:CISCO-PING-MIB/ciscoPingTable/" + ciscoPingEntry.EntityData.SegmentPath
    ciscoPingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoPingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoPingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoPingEntry.EntityData.Children = types.NewOrderedMap()
    ciscoPingEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingSerialNumber", types.YLeaf{"CiscoPingSerialNumber", ciscoPingEntry.CiscoPingSerialNumber})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingProtocol", types.YLeaf{"CiscoPingProtocol", ciscoPingEntry.CiscoPingProtocol})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingAddress", types.YLeaf{"CiscoPingAddress", ciscoPingEntry.CiscoPingAddress})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingPacketCount", types.YLeaf{"CiscoPingPacketCount", ciscoPingEntry.CiscoPingPacketCount})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingPacketSize", types.YLeaf{"CiscoPingPacketSize", ciscoPingEntry.CiscoPingPacketSize})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingPacketTimeout", types.YLeaf{"CiscoPingPacketTimeout", ciscoPingEntry.CiscoPingPacketTimeout})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingDelay", types.YLeaf{"CiscoPingDelay", ciscoPingEntry.CiscoPingDelay})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingTrapOnCompletion", types.YLeaf{"CiscoPingTrapOnCompletion", ciscoPingEntry.CiscoPingTrapOnCompletion})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingSentPackets", types.YLeaf{"CiscoPingSentPackets", ciscoPingEntry.CiscoPingSentPackets})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingReceivedPackets", types.YLeaf{"CiscoPingReceivedPackets", ciscoPingEntry.CiscoPingReceivedPackets})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingMinRtt", types.YLeaf{"CiscoPingMinRtt", ciscoPingEntry.CiscoPingMinRtt})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingAvgRtt", types.YLeaf{"CiscoPingAvgRtt", ciscoPingEntry.CiscoPingAvgRtt})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingMaxRtt", types.YLeaf{"CiscoPingMaxRtt", ciscoPingEntry.CiscoPingMaxRtt})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingCompleted", types.YLeaf{"CiscoPingCompleted", ciscoPingEntry.CiscoPingCompleted})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingEntryOwner", types.YLeaf{"CiscoPingEntryOwner", ciscoPingEntry.CiscoPingEntryOwner})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingEntryStatus", types.YLeaf{"CiscoPingEntryStatus", ciscoPingEntry.CiscoPingEntryStatus})
    ciscoPingEntry.EntityData.Leafs.Append("ciscoPingVrfName", types.YLeaf{"CiscoPingVrfName", ciscoPingEntry.CiscoPingVrfName})

    ciscoPingEntry.EntityData.YListKeys = []string {"CiscoPingSerialNumber"}

    return &(ciscoPingEntry.EntityData)
}

