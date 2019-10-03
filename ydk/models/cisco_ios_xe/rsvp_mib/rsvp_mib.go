// The MIB module to describe the RSVP Protocol
package rsvp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rsvp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:RSVP-MIB RSVP-MIB}", reflect.TypeOf(RSVPMIB{}))
    ydk.RegisterEntity("RSVP-MIB:RSVP-MIB", reflect.TypeOf(RSVPMIB{}))
}

// RsvpEncapsulation represents Neighbor is perceived to be using.
type RsvpEncapsulation string

const (
    RsvpEncapsulation_ip RsvpEncapsulation = "ip"

    RsvpEncapsulation_udp RsvpEncapsulation = "udp"

    RsvpEncapsulation_both RsvpEncapsulation = "both"
)

// RSVPMIB
type RSVPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RsvpGenObjects RSVPMIB_RsvpGenObjects

    // A table  of	 all  sessions	seen  by  a  given system.
    RsvpSessionTable RSVPMIB_RsvpSessionTable

    // Information	describing the	state  information displayed by	senders	in
    // PATH	messages.
    RsvpSenderTable RSVPMIB_RsvpSenderTable

    // List of outgoing interfaces	that PATH messages use.	 The  ifIndex  is the
    // ifIndex value of the egress interface.
    RsvpSenderOutInterfaceTable RSVPMIB_RsvpSenderOutInterfaceTable

    // Information	describing the	state  information displayed by	receivers in
    // RESV messages.
    RsvpResvTable RSVPMIB_RsvpResvTable

    // Information	describing the	state  information displayed upstream in RESV
    // messages.
    RsvpResvFwdTable RSVPMIB_RsvpResvFwdTable

    // The	RSVP-specific attributes of  the  system's interfaces.
    RsvpIfTable RSVPMIB_RsvpIfTable

    // Information	describing  the	 Neighbors  of	an RSVP	system.
    RsvpNbrTable RSVPMIB_RsvpNbrTable
}

func (rSVPMIB *RSVPMIB) GetEntityData() *types.CommonEntityData {
    rSVPMIB.EntityData.YFilter = rSVPMIB.YFilter
    rSVPMIB.EntityData.YangName = "RSVP-MIB"
    rSVPMIB.EntityData.BundleName = "cisco_ios_xe"
    rSVPMIB.EntityData.ParentYangName = "RSVP-MIB"
    rSVPMIB.EntityData.SegmentPath = "RSVP-MIB:RSVP-MIB"
    rSVPMIB.EntityData.AbsolutePath = rSVPMIB.EntityData.SegmentPath
    rSVPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rSVPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rSVPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rSVPMIB.EntityData.Children = types.NewOrderedMap()
    rSVPMIB.EntityData.Children.Append("rsvpGenObjects", types.YChild{"RsvpGenObjects", &rSVPMIB.RsvpGenObjects})
    rSVPMIB.EntityData.Children.Append("rsvpSessionTable", types.YChild{"RsvpSessionTable", &rSVPMIB.RsvpSessionTable})
    rSVPMIB.EntityData.Children.Append("rsvpSenderTable", types.YChild{"RsvpSenderTable", &rSVPMIB.RsvpSenderTable})
    rSVPMIB.EntityData.Children.Append("rsvpSenderOutInterfaceTable", types.YChild{"RsvpSenderOutInterfaceTable", &rSVPMIB.RsvpSenderOutInterfaceTable})
    rSVPMIB.EntityData.Children.Append("rsvpResvTable", types.YChild{"RsvpResvTable", &rSVPMIB.RsvpResvTable})
    rSVPMIB.EntityData.Children.Append("rsvpResvFwdTable", types.YChild{"RsvpResvFwdTable", &rSVPMIB.RsvpResvFwdTable})
    rSVPMIB.EntityData.Children.Append("rsvpIfTable", types.YChild{"RsvpIfTable", &rSVPMIB.RsvpIfTable})
    rSVPMIB.EntityData.Children.Append("rsvpNbrTable", types.YChild{"RsvpNbrTable", &rSVPMIB.RsvpNbrTable})
    rSVPMIB.EntityData.Leafs = types.NewOrderedMap()

    rSVPMIB.EntityData.YListKeys = []string {}

    return &(rSVPMIB.EntityData)
}

// RSVPMIB_RsvpGenObjects
type RSVPMIB_RsvpGenObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object	keeps a	count of the number of bad RSVP	packets	received. The
    // type is interface{} with range: 0..4294967295.
    RsvpBadPackets interface{}

    // This  object  is  used  to	assign	values	to rsvpSenderNumber   as  
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpSenderEntry.   If  the  SET  fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    RsvpSenderNewIndex interface{}

    // This  object  is  used  to	assign	values	to rsvpResvNumber   as	 described 
    // in	  'Textual Conventions for SNMPv2'.  The  network  manager reads  the 
    // object,	and  then writes the value back	in the SET that	creates	a new
    // instance	of rsvpResvEntry.   If the SET fails with the code
    // 'inconsistentValue',	then the process  must	be repeated;  If the SET
    // succeeds, then	the object is incremented, and the new instance	is created
    // according to	the manager's directions. The type is interface{} with range:
    // 0..2147483647.
    RsvpResvNewIndex interface{}

    // This  object  is  used  to	assign	values	to rsvpResvFwdNumber   as 
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpResvFwdEntry.   If  the	SET fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    RsvpResvFwdNewIndex interface{}

    // This  object  is  used  to	assign	values	to rsvpSessionNumber   as 
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpSessionEntry.   If  the	SET fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    RsvpSessionNewIndex interface{}
}

func (rsvpGenObjects *RSVPMIB_RsvpGenObjects) GetEntityData() *types.CommonEntityData {
    rsvpGenObjects.EntityData.YFilter = rsvpGenObjects.YFilter
    rsvpGenObjects.EntityData.YangName = "rsvpGenObjects"
    rsvpGenObjects.EntityData.BundleName = "cisco_ios_xe"
    rsvpGenObjects.EntityData.ParentYangName = "RSVP-MIB"
    rsvpGenObjects.EntityData.SegmentPath = "rsvpGenObjects"
    rsvpGenObjects.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpGenObjects.EntityData.SegmentPath
    rsvpGenObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpGenObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpGenObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpGenObjects.EntityData.Children = types.NewOrderedMap()
    rsvpGenObjects.EntityData.Leafs = types.NewOrderedMap()
    rsvpGenObjects.EntityData.Leafs.Append("rsvpBadPackets", types.YLeaf{"RsvpBadPackets", rsvpGenObjects.RsvpBadPackets})
    rsvpGenObjects.EntityData.Leafs.Append("rsvpSenderNewIndex", types.YLeaf{"RsvpSenderNewIndex", rsvpGenObjects.RsvpSenderNewIndex})
    rsvpGenObjects.EntityData.Leafs.Append("rsvpResvNewIndex", types.YLeaf{"RsvpResvNewIndex", rsvpGenObjects.RsvpResvNewIndex})
    rsvpGenObjects.EntityData.Leafs.Append("rsvpResvFwdNewIndex", types.YLeaf{"RsvpResvFwdNewIndex", rsvpGenObjects.RsvpResvFwdNewIndex})
    rsvpGenObjects.EntityData.Leafs.Append("rsvpSessionNewIndex", types.YLeaf{"RsvpSessionNewIndex", rsvpGenObjects.RsvpSessionNewIndex})

    rsvpGenObjects.EntityData.YListKeys = []string {}

    return &(rsvpGenObjects.EntityData)
}

// RSVPMIB_RsvpSessionTable
// A table  of	 all  sessions	seen  by  a  given
// system.
type RSVPMIB_RsvpSessionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single session seen by a given system. The type is slice of
    // RSVPMIB_RsvpSessionTable_RsvpSessionEntry.
    RsvpSessionEntry []*RSVPMIB_RsvpSessionTable_RsvpSessionEntry
}

func (rsvpSessionTable *RSVPMIB_RsvpSessionTable) GetEntityData() *types.CommonEntityData {
    rsvpSessionTable.EntityData.YFilter = rsvpSessionTable.YFilter
    rsvpSessionTable.EntityData.YangName = "rsvpSessionTable"
    rsvpSessionTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpSessionTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpSessionTable.EntityData.SegmentPath = "rsvpSessionTable"
    rsvpSessionTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpSessionTable.EntityData.SegmentPath
    rsvpSessionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSessionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSessionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSessionTable.EntityData.Children = types.NewOrderedMap()
    rsvpSessionTable.EntityData.Children.Append("rsvpSessionEntry", types.YChild{"RsvpSessionEntry", nil})
    for i := range rsvpSessionTable.RsvpSessionEntry {
        rsvpSessionTable.EntityData.Children.Append(types.GetSegmentPath(rsvpSessionTable.RsvpSessionEntry[i]), types.YChild{"RsvpSessionEntry", rsvpSessionTable.RsvpSessionEntry[i]})
    }
    rsvpSessionTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpSessionTable.EntityData.YListKeys = []string {}

    return &(rsvpSessionTable.EntityData)
}

// RSVPMIB_RsvpSessionTable_RsvpSessionEntry
// A single session seen by a given system.
type RSVPMIB_RsvpSessionTable_RsvpSessionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The	number of this session.	 This is for  SNMP
    // Indexing  purposes  only and	has no relation	to any protocol	value. The
    // type is interface{} with range: 0..2147483647.
    RsvpSessionNumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    RsvpSessionType interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    RsvpSessionDestAddr interface{}

    // The	CIDR prefix length of the session address, which   is	32  for	 IP4 
    // host  and	 multicast addresses, and 128  for  IP6	 addresses.   This object
    // may not be changed when the value of the RowStatus object is 'active'. The
    // type is interface{} with range: 0..128.
    RsvpSessionDestAddrLength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    RsvpSessionProtocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpSenderProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    RsvpSessionPort interface{}

    // The	number of distinct senders currently known to be part of this session.
    // The type is interface{} with range: 0..4294967295.
    RsvpSessionSenders interface{}

    // The	number of reservations being requested	of this	system for this	session.
    // The type is interface{} with range: 0..4294967295.
    RsvpSessionReceivers interface{}

    // The	number of reservation requests this system is sending upstream for
    // this	session. The type is interface{} with range: 0..4294967295.
    RsvpSessionRequests interface{}
}

func (rsvpSessionEntry *RSVPMIB_RsvpSessionTable_RsvpSessionEntry) GetEntityData() *types.CommonEntityData {
    rsvpSessionEntry.EntityData.YFilter = rsvpSessionEntry.YFilter
    rsvpSessionEntry.EntityData.YangName = "rsvpSessionEntry"
    rsvpSessionEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpSessionEntry.EntityData.ParentYangName = "rsvpSessionTable"
    rsvpSessionEntry.EntityData.SegmentPath = "rsvpSessionEntry" + types.AddKeyToken(rsvpSessionEntry.RsvpSessionNumber, "rsvpSessionNumber")
    rsvpSessionEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpSessionTable/" + rsvpSessionEntry.EntityData.SegmentPath
    rsvpSessionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSessionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSessionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSessionEntry.EntityData.Children = types.NewOrderedMap()
    rsvpSessionEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionNumber", types.YLeaf{"RsvpSessionNumber", rsvpSessionEntry.RsvpSessionNumber})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionType", types.YLeaf{"RsvpSessionType", rsvpSessionEntry.RsvpSessionType})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionDestAddr", types.YLeaf{"RsvpSessionDestAddr", rsvpSessionEntry.RsvpSessionDestAddr})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionDestAddrLength", types.YLeaf{"RsvpSessionDestAddrLength", rsvpSessionEntry.RsvpSessionDestAddrLength})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionProtocol", types.YLeaf{"RsvpSessionProtocol", rsvpSessionEntry.RsvpSessionProtocol})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionPort", types.YLeaf{"RsvpSessionPort", rsvpSessionEntry.RsvpSessionPort})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionSenders", types.YLeaf{"RsvpSessionSenders", rsvpSessionEntry.RsvpSessionSenders})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionReceivers", types.YLeaf{"RsvpSessionReceivers", rsvpSessionEntry.RsvpSessionReceivers})
    rsvpSessionEntry.EntityData.Leafs.Append("rsvpSessionRequests", types.YLeaf{"RsvpSessionRequests", rsvpSessionEntry.RsvpSessionRequests})

    rsvpSessionEntry.EntityData.YListKeys = []string {"RsvpSessionNumber"}

    return &(rsvpSessionEntry.EntityData)
}

// RSVPMIB_RsvpSenderTable
// Information	describing the	state  information
// displayed by	senders	in PATH	messages.
type RSVPMIB_RsvpSenderTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed by	a single
    // sender's PATH message. The type is slice of
    // RSVPMIB_RsvpSenderTable_RsvpSenderEntry.
    RsvpSenderEntry []*RSVPMIB_RsvpSenderTable_RsvpSenderEntry
}

func (rsvpSenderTable *RSVPMIB_RsvpSenderTable) GetEntityData() *types.CommonEntityData {
    rsvpSenderTable.EntityData.YFilter = rsvpSenderTable.YFilter
    rsvpSenderTable.EntityData.YangName = "rsvpSenderTable"
    rsvpSenderTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpSenderTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpSenderTable.EntityData.SegmentPath = "rsvpSenderTable"
    rsvpSenderTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpSenderTable.EntityData.SegmentPath
    rsvpSenderTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSenderTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSenderTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSenderTable.EntityData.Children = types.NewOrderedMap()
    rsvpSenderTable.EntityData.Children.Append("rsvpSenderEntry", types.YChild{"RsvpSenderEntry", nil})
    for i := range rsvpSenderTable.RsvpSenderEntry {
        rsvpSenderTable.EntityData.Children.Append(types.GetSegmentPath(rsvpSenderTable.RsvpSenderEntry[i]), types.YChild{"RsvpSenderEntry", rsvpSenderTable.RsvpSenderEntry[i]})
    }
    rsvpSenderTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpSenderTable.EntityData.YListKeys = []string {}

    return &(rsvpSenderTable.EntityData)
}

// RSVPMIB_RsvpSenderTable_RsvpSenderEntry
// Information	describing the	state  information
// displayed by	a single sender's PATH message.
type RSVPMIB_RsvpSenderTable_RsvpSenderEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_RsvpSessionTable_RsvpSessionEntry_RsvpSessionNumber
    RsvpSessionNumber interface{}

    // This attribute is a key. The	number of this sender.	This is	 for  SNMP
    // Indexing  purposes  only and	has no relation	to any protocol	value. The
    // type is interface{} with range: 0..2147483647.
    RsvpSenderNumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    RsvpSenderType interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    RsvpSenderDestAddr interface{}

    // The	source address used by this sender in this session.   This  object may
    // not be changed when the value of	the RowStatus object is	'active'. The type
    // is string with length: 4..16.
    RsvpSenderAddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpSenderDestAddrLength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpSenderAddrLength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    RsvpSenderProtocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpSenderProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    RsvpSenderDestPort interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpSenderProtocol is 50 (ESP) or 51	(AH), this represents a	generalized
    // port identifier (GPI). A  value of zero indicates that the IP protocol in
    // use does not have	ports.	 This  object  may not	be changed when	the value
    // of the RowStatus object is 'active'. The type is string with length: 2..4.
    RsvpSenderPort interface{}

    // The	flow ID	that  this  sender  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    RsvpSenderFlowId interface{}

    // The	address	used  by  the  previous	 RSVP  hop (which may be the original
    // sender). The type is string with length: 4..16.
    RsvpSenderHopAddr interface{}

    // The	 Logical  Interface  Handle  used  by  the previous  RSVP  hop	(which
    // may be the original sender). The type is interface{} with range:
    // -2147483648..2147483647.
    RsvpSenderHopLih interface{}

    // The	ifIndex	value of the  interface	 on  which this	PATH message was most
    // recently received. The type is interface{} with range: 1..2147483647.
    RsvpSenderInterface interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpSenderTSpecPeakRate  (if	 supported  by the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpSenderTSpecRate.  Note	that this is a prediction, often based	on
    // the	general	 capability  of	a type of codec	or particular encoding;	the
    // measured average  rate may be significantly	lower. The type is interface{}
    // with range: 0..2147483647. Units are bits per second.
    RsvpSenderTSpecRate interface{}

    // The	Peak Bit Rate of the sender's data stream. Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    RsvpSenderTSpecPeakRate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time. The
    // type is interface{} with range: 0..2147483647. Units are bytes.
    RsvpSenderTSpecBurst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    RsvpSenderTSpecMinTU interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    RsvpSenderTSpecMaxTU interface{}

    // The	 interval  between  refresh  messages	as advertised by the Previous
    // Hop. The type is interface{} with range: 0..2147483647.
    RsvpSenderInterval interface{}

    // If TRUE, the node believes that  the  previous IP  hop  is	an  RSVP 
    // hop.	If FALSE, the node believes that the previous IP hop may not be	an
    // RSVP	hop. The type is bool.
    RsvpSenderRSVPHop interface{}

    // The	time of	 the  last  change  in	this  PATH message;  This  is either the
    // first time it was received or the time	of the most recent  change in
    // parameters. The type is interface{} with range: 0..4294967295.
    RsvpSenderLastChange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 4..65536.
    RsvpSenderPolicy interface{}

    // The	global break bit general  characterization parameter  from 
    // the	ADSPEC.	 If TRUE, at least one non-IS hop was detected in  the	path.	If
    // FALSE, no non-IS hops were detected. The type is bool.
    RsvpSenderAdspecBreak interface{}

    // The	  hop	count	general	  characterization parameter from the ADSPEC.  A
    // return	of zero	or noSuchValue	indicates  one	of  the	 following conditions:
    // the invalid bit was set    the parameter was	not present. The type is
    // interface{} with range: 0..65535.
    RsvpSenderAdspecHopCount interface{}

    // The	  path	  bandwidth    estimate	   general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: 0..2147483647. Units are bits
    // per second.
    RsvpSenderAdspecPathBw interface{}

    // The	   minimum    path     latency	   general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: -2147483648..2147483647. Units
    // are microseconds.
    RsvpSenderAdspecMinLatency interface{}

    // The	composed Maximum Transmission Unit general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: 0..65535. Units are bytes.
    RsvpSenderAdspecMtu interface{}

    // If TRUE,  the  ADSPEC  contains  a	Guaranteed Service  fragment.	If  FALSE,
    // the ADSPEC does not contain a Guaranteed Service fragment. The type is
    // bool.
    RsvpSenderAdspecGuaranteedSvc interface{}

    // If TRUE, the Guaranteed Service  fragment  has its	'break'	 bit  set, 
    // indicating that one	or more	nodes along the	path do	 not  support  the
    // guaranteed	  service.     If    FALSE,    and rsvpSenderAdspecGuaranteedSvc
    // is   TRUE,   the 'break' bit is not set.  If rsvpSenderAdspecGuaranteedSvc
    // is FALSE, this returns FALSE or noSuchValue. The type is bool.
    RsvpSenderAdspecGuaranteedBreak interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  end-to-end	 composed
    // value  for  the guaranteed service 'C' parameter.  A	return	of zero	  or 
    // noSuchValue  indicates  one  of  the following conditions:     the invalid
    // bit was set    the parameter was	not present  If
    // rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or noSuchValue.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // bytes.
    RsvpSenderAdspecGuaranteedCtot interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  end-to-end	 composed
    // value  for  the guaranteed service 'D' parameter.  A	return	of zero	  or 
    // noSuchValue  indicates  one  of  the following conditions:     the invalid
    // bit was set    the parameter was	not present  If
    // rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or noSuchValue.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // microseconds.
    RsvpSenderAdspecGuaranteedDtot interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  composed  value  for
    // the	guaranteed service 'C' parameter since the last	 reshaping point.    A	
    // return	 of  zero  or  noSuchValue indicates one of the	following
    // conditions:     the invalid bit was set    the parameter was	not present 
    // If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are bytes.
    RsvpSenderAdspecGuaranteedCsum interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  composed  value  for
    // the	guaranteed service 'D' parameter since the last	 reshaping point.    A	
    // return	 of  zero  or  noSuchValue indicates one of the	following
    // conditions:     the invalid bit was set    the parameter was	not present 
    // If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    RsvpSenderAdspecGuaranteedDsum interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is  the  service-specific 
    // override	of the hop count general characterization  parameter  from
    // the	ADSPEC.	  A  return of zero or noSuchValue indicates one of
    // the	following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns
    // zero	or noSuchValue. The type is interface{} with range: 0..65535.
    RsvpSenderAdspecGuaranteedHopCount interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is  the  service-specific 
    // override of the path bandwidth  estimate	general	  characterization
    // parameter from the ADSPEC.  A return	of zero	or noSuchValue	indicates 
    // one	of  the	 following conditions:     the invalid bit was set    the
    // parameter was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this
    // returns zero	or noSuchValue. The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    RsvpSenderAdspecGuaranteedPathBw interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is the service-specific
    // override of the minimum path	latency	general	characterization parameter
    // from	  the	ADSPEC.	   A  return  of  zero	or noSuchValue	indicates  one	of 
    // the	 following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns
    // zero	or noSuchValue. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    RsvpSenderAdspecGuaranteedMinLatency interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the   service-specific	
    // override  of  the composed  Maximum  Transmission  Unit   general
    // characterization  parameter from the	ADSPEC.	 A return of zero or
    // noSuchValue indicates one	of the following conditions:     the invalid bit
    // was set    the parameter was	not present  If rsvpSenderAdspecGuaranteedSvc
    // is FALSE, this returns zero	or noSuchValue. The type is interface{} with
    // range: 0..65535. Units are bytes.
    RsvpSenderAdspecGuaranteedMtu interface{}

    // If TRUE, the ADSPEC	contains a Controlled Load Service  fragment.	If 
    // FALSE, the ADSPEC does not	 contain   a   Controlled   Load   Service
    // fragment. The type is bool.
    RsvpSenderAdspecCtrlLoadSvc interface{}

    // If TRUE, the Controlled Load Service  fragment has its 'break' bit set,
    // indicating that one	or more	nodes along the	path do	 not  support  the
    // controlled	load   service.	   If  FALSE,  and rsvpSenderAdspecCtrlLoadSvc	
    // is   TRUE,    the 'break' bit is not set.  If rsvpSenderAdspecCtrlLoadSvc
    // is  FALSE,  this returns FALSE or noSuchValue. The type is bool.
    RsvpSenderAdspecCtrlLoadBreak interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is  the  service-specific 
    // override	of the hop count general characterization  parameter  from
    // the	ADSPEC.	  A  return of zero or noSuchValue indicates one of
    // the	following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this returns
    // zero	or noSuchValue. The type is interface{} with range: 0..65535.
    RsvpSenderAdspecCtrlLoadHopCount interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is  the  service-specific 
    // override of the path bandwidth  estimate	general	  characterization
    // parameter from the ADSPEC.  A return	of zero	or noSuchValue	indicates 
    // one	of  the	 following conditions:     the invalid bit was set    the
    // parameter was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this
    // returns zero	or noSuchValue. The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    RsvpSenderAdspecCtrlLoadPathBw interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is the service-specific
    // override of the minimum path	latency	general	characterization parameter
    // from	  the	ADSPEC.	   A  return  of  zero	or noSuchValue	indicates  one	of 
    // the	 following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this returns
    // zero	or noSuchValue. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    RsvpSenderAdspecCtrlLoadMinLatency interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is	the   service-specific	
    // override  of  the composed  Maximum  Transmission  Unit   general
    // characterization  parameter from the	ADSPEC.	 A return of zero or
    // noSuchValue indicates one	of the following conditions:     the invalid bit
    // was set    the parameter was	not present  If rsvpSenderAdspecCtrlLoadSvc is
    // FALSE,  this returns zero	or noSuchValue. The type is interface{} with
    // range: 0..65535. Units are bytes.
    RsvpSenderAdspecCtrlLoadMtu interface{}

    // 'active' for all active PATH  messages.   This object  may	be  used  to 
    // install  static PATH information or delete PATH information. The type is
    // RowStatus.
    RsvpSenderStatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    RsvpSenderTTL interface{}
}

func (rsvpSenderEntry *RSVPMIB_RsvpSenderTable_RsvpSenderEntry) GetEntityData() *types.CommonEntityData {
    rsvpSenderEntry.EntityData.YFilter = rsvpSenderEntry.YFilter
    rsvpSenderEntry.EntityData.YangName = "rsvpSenderEntry"
    rsvpSenderEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpSenderEntry.EntityData.ParentYangName = "rsvpSenderTable"
    rsvpSenderEntry.EntityData.SegmentPath = "rsvpSenderEntry" + types.AddKeyToken(rsvpSenderEntry.RsvpSessionNumber, "rsvpSessionNumber") + types.AddKeyToken(rsvpSenderEntry.RsvpSenderNumber, "rsvpSenderNumber")
    rsvpSenderEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpSenderTable/" + rsvpSenderEntry.EntityData.SegmentPath
    rsvpSenderEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSenderEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSenderEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSenderEntry.EntityData.Children = types.NewOrderedMap()
    rsvpSenderEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSessionNumber", types.YLeaf{"RsvpSessionNumber", rsvpSenderEntry.RsvpSessionNumber})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderNumber", types.YLeaf{"RsvpSenderNumber", rsvpSenderEntry.RsvpSenderNumber})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderType", types.YLeaf{"RsvpSenderType", rsvpSenderEntry.RsvpSenderType})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderDestAddr", types.YLeaf{"RsvpSenderDestAddr", rsvpSenderEntry.RsvpSenderDestAddr})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAddr", types.YLeaf{"RsvpSenderAddr", rsvpSenderEntry.RsvpSenderAddr})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderDestAddrLength", types.YLeaf{"RsvpSenderDestAddrLength", rsvpSenderEntry.RsvpSenderDestAddrLength})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAddrLength", types.YLeaf{"RsvpSenderAddrLength", rsvpSenderEntry.RsvpSenderAddrLength})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderProtocol", types.YLeaf{"RsvpSenderProtocol", rsvpSenderEntry.RsvpSenderProtocol})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderDestPort", types.YLeaf{"RsvpSenderDestPort", rsvpSenderEntry.RsvpSenderDestPort})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderPort", types.YLeaf{"RsvpSenderPort", rsvpSenderEntry.RsvpSenderPort})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderFlowId", types.YLeaf{"RsvpSenderFlowId", rsvpSenderEntry.RsvpSenderFlowId})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderHopAddr", types.YLeaf{"RsvpSenderHopAddr", rsvpSenderEntry.RsvpSenderHopAddr})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderHopLih", types.YLeaf{"RsvpSenderHopLih", rsvpSenderEntry.RsvpSenderHopLih})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderInterface", types.YLeaf{"RsvpSenderInterface", rsvpSenderEntry.RsvpSenderInterface})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTSpecRate", types.YLeaf{"RsvpSenderTSpecRate", rsvpSenderEntry.RsvpSenderTSpecRate})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTSpecPeakRate", types.YLeaf{"RsvpSenderTSpecPeakRate", rsvpSenderEntry.RsvpSenderTSpecPeakRate})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTSpecBurst", types.YLeaf{"RsvpSenderTSpecBurst", rsvpSenderEntry.RsvpSenderTSpecBurst})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTSpecMinTU", types.YLeaf{"RsvpSenderTSpecMinTU", rsvpSenderEntry.RsvpSenderTSpecMinTU})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTSpecMaxTU", types.YLeaf{"RsvpSenderTSpecMaxTU", rsvpSenderEntry.RsvpSenderTSpecMaxTU})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderInterval", types.YLeaf{"RsvpSenderInterval", rsvpSenderEntry.RsvpSenderInterval})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderRSVPHop", types.YLeaf{"RsvpSenderRSVPHop", rsvpSenderEntry.RsvpSenderRSVPHop})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderLastChange", types.YLeaf{"RsvpSenderLastChange", rsvpSenderEntry.RsvpSenderLastChange})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderPolicy", types.YLeaf{"RsvpSenderPolicy", rsvpSenderEntry.RsvpSenderPolicy})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecBreak", types.YLeaf{"RsvpSenderAdspecBreak", rsvpSenderEntry.RsvpSenderAdspecBreak})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecHopCount", types.YLeaf{"RsvpSenderAdspecHopCount", rsvpSenderEntry.RsvpSenderAdspecHopCount})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecPathBw", types.YLeaf{"RsvpSenderAdspecPathBw", rsvpSenderEntry.RsvpSenderAdspecPathBw})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecMinLatency", types.YLeaf{"RsvpSenderAdspecMinLatency", rsvpSenderEntry.RsvpSenderAdspecMinLatency})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecMtu", types.YLeaf{"RsvpSenderAdspecMtu", rsvpSenderEntry.RsvpSenderAdspecMtu})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedSvc", types.YLeaf{"RsvpSenderAdspecGuaranteedSvc", rsvpSenderEntry.RsvpSenderAdspecGuaranteedSvc})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedBreak", types.YLeaf{"RsvpSenderAdspecGuaranteedBreak", rsvpSenderEntry.RsvpSenderAdspecGuaranteedBreak})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedCtot", types.YLeaf{"RsvpSenderAdspecGuaranteedCtot", rsvpSenderEntry.RsvpSenderAdspecGuaranteedCtot})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedDtot", types.YLeaf{"RsvpSenderAdspecGuaranteedDtot", rsvpSenderEntry.RsvpSenderAdspecGuaranteedDtot})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedCsum", types.YLeaf{"RsvpSenderAdspecGuaranteedCsum", rsvpSenderEntry.RsvpSenderAdspecGuaranteedCsum})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedDsum", types.YLeaf{"RsvpSenderAdspecGuaranteedDsum", rsvpSenderEntry.RsvpSenderAdspecGuaranteedDsum})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedHopCount", types.YLeaf{"RsvpSenderAdspecGuaranteedHopCount", rsvpSenderEntry.RsvpSenderAdspecGuaranteedHopCount})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedPathBw", types.YLeaf{"RsvpSenderAdspecGuaranteedPathBw", rsvpSenderEntry.RsvpSenderAdspecGuaranteedPathBw})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedMinLatency", types.YLeaf{"RsvpSenderAdspecGuaranteedMinLatency", rsvpSenderEntry.RsvpSenderAdspecGuaranteedMinLatency})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecGuaranteedMtu", types.YLeaf{"RsvpSenderAdspecGuaranteedMtu", rsvpSenderEntry.RsvpSenderAdspecGuaranteedMtu})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadSvc", types.YLeaf{"RsvpSenderAdspecCtrlLoadSvc", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadSvc})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadBreak", types.YLeaf{"RsvpSenderAdspecCtrlLoadBreak", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadBreak})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadHopCount", types.YLeaf{"RsvpSenderAdspecCtrlLoadHopCount", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadHopCount})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadPathBw", types.YLeaf{"RsvpSenderAdspecCtrlLoadPathBw", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadPathBw})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadMinLatency", types.YLeaf{"RsvpSenderAdspecCtrlLoadMinLatency", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadMinLatency})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderAdspecCtrlLoadMtu", types.YLeaf{"RsvpSenderAdspecCtrlLoadMtu", rsvpSenderEntry.RsvpSenderAdspecCtrlLoadMtu})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderStatus", types.YLeaf{"RsvpSenderStatus", rsvpSenderEntry.RsvpSenderStatus})
    rsvpSenderEntry.EntityData.Leafs.Append("rsvpSenderTTL", types.YLeaf{"RsvpSenderTTL", rsvpSenderEntry.RsvpSenderTTL})

    rsvpSenderEntry.EntityData.YListKeys = []string {"RsvpSessionNumber", "RsvpSenderNumber"}

    return &(rsvpSenderEntry.EntityData)
}

// RSVPMIB_RsvpSenderOutInterfaceTable
// List of outgoing interfaces	that PATH messages
// use.	 The  ifIndex  is the ifIndex value of the
// egress interface.
type RSVPMIB_RsvpSenderOutInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of outgoing interfaces	that a	particular PATH	message	has. The type is
    // slice of RSVPMIB_RsvpSenderOutInterfaceTable_RsvpSenderOutInterfaceEntry.
    RsvpSenderOutInterfaceEntry []*RSVPMIB_RsvpSenderOutInterfaceTable_RsvpSenderOutInterfaceEntry
}

func (rsvpSenderOutInterfaceTable *RSVPMIB_RsvpSenderOutInterfaceTable) GetEntityData() *types.CommonEntityData {
    rsvpSenderOutInterfaceTable.EntityData.YFilter = rsvpSenderOutInterfaceTable.YFilter
    rsvpSenderOutInterfaceTable.EntityData.YangName = "rsvpSenderOutInterfaceTable"
    rsvpSenderOutInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpSenderOutInterfaceTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpSenderOutInterfaceTable.EntityData.SegmentPath = "rsvpSenderOutInterfaceTable"
    rsvpSenderOutInterfaceTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpSenderOutInterfaceTable.EntityData.SegmentPath
    rsvpSenderOutInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSenderOutInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSenderOutInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSenderOutInterfaceTable.EntityData.Children = types.NewOrderedMap()
    rsvpSenderOutInterfaceTable.EntityData.Children.Append("rsvpSenderOutInterfaceEntry", types.YChild{"RsvpSenderOutInterfaceEntry", nil})
    for i := range rsvpSenderOutInterfaceTable.RsvpSenderOutInterfaceEntry {
        rsvpSenderOutInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(rsvpSenderOutInterfaceTable.RsvpSenderOutInterfaceEntry[i]), types.YChild{"RsvpSenderOutInterfaceEntry", rsvpSenderOutInterfaceTable.RsvpSenderOutInterfaceEntry[i]})
    }
    rsvpSenderOutInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpSenderOutInterfaceTable.EntityData.YListKeys = []string {}

    return &(rsvpSenderOutInterfaceTable.EntityData)
}

// RSVPMIB_RsvpSenderOutInterfaceTable_RsvpSenderOutInterfaceEntry
// List of outgoing interfaces	that a	particular
// PATH	message	has.
type RSVPMIB_RsvpSenderOutInterfaceTable_RsvpSenderOutInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_RsvpSessionTable_RsvpSessionEntry_RsvpSessionNumber
    RsvpSessionNumber interface{}

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to rsvp_mib.RSVPMIB_RsvpSenderTable_RsvpSenderEntry_RsvpSenderNumber
    RsvpSenderNumber interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // 'active' for all active PATH messages. The type is RowStatus.
    RsvpSenderOutInterfaceStatus interface{}
}

func (rsvpSenderOutInterfaceEntry *RSVPMIB_RsvpSenderOutInterfaceTable_RsvpSenderOutInterfaceEntry) GetEntityData() *types.CommonEntityData {
    rsvpSenderOutInterfaceEntry.EntityData.YFilter = rsvpSenderOutInterfaceEntry.YFilter
    rsvpSenderOutInterfaceEntry.EntityData.YangName = "rsvpSenderOutInterfaceEntry"
    rsvpSenderOutInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpSenderOutInterfaceEntry.EntityData.ParentYangName = "rsvpSenderOutInterfaceTable"
    rsvpSenderOutInterfaceEntry.EntityData.SegmentPath = "rsvpSenderOutInterfaceEntry" + types.AddKeyToken(rsvpSenderOutInterfaceEntry.RsvpSessionNumber, "rsvpSessionNumber") + types.AddKeyToken(rsvpSenderOutInterfaceEntry.RsvpSenderNumber, "rsvpSenderNumber") + types.AddKeyToken(rsvpSenderOutInterfaceEntry.IfIndex, "ifIndex")
    rsvpSenderOutInterfaceEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpSenderOutInterfaceTable/" + rsvpSenderOutInterfaceEntry.EntityData.SegmentPath
    rsvpSenderOutInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpSenderOutInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpSenderOutInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpSenderOutInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    rsvpSenderOutInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpSenderOutInterfaceEntry.EntityData.Leafs.Append("rsvpSessionNumber", types.YLeaf{"RsvpSessionNumber", rsvpSenderOutInterfaceEntry.RsvpSessionNumber})
    rsvpSenderOutInterfaceEntry.EntityData.Leafs.Append("rsvpSenderNumber", types.YLeaf{"RsvpSenderNumber", rsvpSenderOutInterfaceEntry.RsvpSenderNumber})
    rsvpSenderOutInterfaceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", rsvpSenderOutInterfaceEntry.IfIndex})
    rsvpSenderOutInterfaceEntry.EntityData.Leafs.Append("rsvpSenderOutInterfaceStatus", types.YLeaf{"RsvpSenderOutInterfaceStatus", rsvpSenderOutInterfaceEntry.RsvpSenderOutInterfaceStatus})

    rsvpSenderOutInterfaceEntry.EntityData.YListKeys = []string {"RsvpSessionNumber", "RsvpSenderNumber", "IfIndex"}

    return &(rsvpSenderOutInterfaceEntry.EntityData)
}

// RSVPMIB_RsvpResvTable
// Information	describing the	state  information
// displayed by	receivers in RESV messages.
type RSVPMIB_RsvpResvTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed  by  a single
    // receiver's RESV message concerning a	single sender. The type is slice of
    // RSVPMIB_RsvpResvTable_RsvpResvEntry.
    RsvpResvEntry []*RSVPMIB_RsvpResvTable_RsvpResvEntry
}

func (rsvpResvTable *RSVPMIB_RsvpResvTable) GetEntityData() *types.CommonEntityData {
    rsvpResvTable.EntityData.YFilter = rsvpResvTable.YFilter
    rsvpResvTable.EntityData.YangName = "rsvpResvTable"
    rsvpResvTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpResvTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpResvTable.EntityData.SegmentPath = "rsvpResvTable"
    rsvpResvTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpResvTable.EntityData.SegmentPath
    rsvpResvTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpResvTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpResvTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpResvTable.EntityData.Children = types.NewOrderedMap()
    rsvpResvTable.EntityData.Children.Append("rsvpResvEntry", types.YChild{"RsvpResvEntry", nil})
    for i := range rsvpResvTable.RsvpResvEntry {
        rsvpResvTable.EntityData.Children.Append(types.GetSegmentPath(rsvpResvTable.RsvpResvEntry[i]), types.YChild{"RsvpResvEntry", rsvpResvTable.RsvpResvEntry[i]})
    }
    rsvpResvTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpResvTable.EntityData.YListKeys = []string {}

    return &(rsvpResvTable.EntityData)
}

// RSVPMIB_RsvpResvTable_RsvpResvEntry
// Information	describing the	state  information
// displayed  by  a single receiver's RESV message
// concerning a	single sender.
type RSVPMIB_RsvpResvTable_RsvpResvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_RsvpSessionTable_RsvpSessionEntry_RsvpSessionNumber
    RsvpSessionNumber interface{}

    // This attribute is a key. The	number of this reservation request.   This is 
    // for  SNMP Indexing purposes only	and has	no relation to any protocol value.
    // The type is interface{} with range: 0..2147483647.
    RsvpResvNumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    RsvpResvType interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    RsvpResvDestAddr interface{}

    // The	source address of the sender  selected	by this	 reservation.	The  value
    // of	all zeroes indicates 'all senders'.  This object  may  not be 
    // changed	when  the  value  of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    RsvpResvSenderAddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpResvDestAddrLength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpResvSenderAddrLength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    RsvpResvProtocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpResvProtocol,  is  50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    RsvpResvDestPort interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpResvProtocol  is	 50 (ESP) or 51	(AH), this represents a	generalized
    // port identifier (GPI). A  value of zero indicates that the IP protocol in
    // use does not have	ports.	 This  object  may not	be changed when	the value
    // of the RowStatus object is 'active'. The type is string with length: 2..4.
    RsvpResvPort interface{}

    // The	address	used by	the next RSVP  hop  (which may be the ultimate
    // receiver). The type is string with length: 4..16.
    RsvpResvHopAddr interface{}

    // The	Logical	Interface Handle received from the previous  RSVP  hop	(which
    // may be the ultimate receiver). The type is interface{} with range:
    // -2147483648..2147483647.
    RsvpResvHopLih interface{}

    // The	ifIndex	value of the  interface	 on  which this	RESV message was most
    // recently received. The type is interface{} with range: 1..2147483647.
    RsvpResvInterface interface{}

    // The	QoS Service  classification  requested	by the receiver. The type is
    // QosService.
    RsvpResvService interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpResvTSpecPeakRate   (if	supported  by  the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpResvTSpecRate.  Note	that this is a prediction, often based	on
    // the	general	 capability  of	a type of codec	or particular encoding;	the
    // measured average  rate may be significantly	lower. The type is interface{}
    // with range: 0..2147483647. Units are bits per second.
    RsvpResvTSpecRate interface{}

    // The	Peak Bit Rate of the sender's data stream. Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    RsvpResvTSpecPeakRate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time.  If
    // this is less than	 the  sender's	advertised burst  size,	the receiver
    // is	asking the network to provide flow pacing  beyond  what	 would	be
    // provided   under   normal  circumstances.  Such pacing is at	the network's
    // option. The type is interface{} with range: 0..2147483647. Units are bytes.
    RsvpResvTSpecBurst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    RsvpResvTSpecMinTU interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    RsvpResvTSpecMaxTU interface{}

    // If the requested  service  is  Guaranteed,	as specified   by 
    // rsvpResvService,  this  is  the clearing  rate   that   is	being	requested.
    // Otherwise,  it is zero, or the agent	may return noSuchValue. The type is
    // interface{} with range: 0..2147483647. Units are bits per second.
    RsvpResvRSpecRate interface{}

    // If the requested  service  is  Guaranteed,	as specified by	rsvpResvService,
    // this is the delay slack.  Otherwise, it is zero, or the agent may return
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    RsvpResvRSpecSlack interface{}

    // The	 interval  between  refresh  messages	as advertised by the Next Hop.
    // The type is interface{} with range: 0..2147483647.
    RsvpResvInterval interface{}

    // The	contents of the	scope object, displayed	as an  uninterpreted  string 
    // of octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length.  If the length  is  non-zero,	 this 
    // contains	 a series of IP4 or IP6	addresses. The type is string with length:
    // 0..65536.
    RsvpResvScope interface{}

    // If TRUE, a reservation shared among	senders	is requested.  If FALSE, a
    // reservation specific	to this	sender is requested. The type is bool.
    RsvpResvShared interface{}

    // If TRUE, individual	senders	are  listed  using Filter  Specifications.  
    // If	FALSE, all senders are implicitly selected.  The Scope Object will
    // contain  a list of senders that need	to receive this	reservation request 
    // for  the  purpose	of routing the RESV message. The type is bool.
    RsvpResvExplicit interface{}

    // If TRUE, the node believes that  the  previous IP  hop  is	an  RSVP 
    // hop.	If FALSE, the node believes that the previous IP hop may not be	an
    // RSVP	hop. The type is bool.
    RsvpResvRSVPHop interface{}

    // The	 time  of  the	 last	change	 in   this reservation	request;  This is
    // either the first time	it was received	or the time  of	 the  most recent
    // change in parameters. The type is interface{} with range: 0..4294967295.
    RsvpResvLastChange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    RsvpResvPolicy interface{}

    // 'active' for all active RESV  messages.   This object  may	be  used  to 
    // install  static RESV information or delete RESV information. The type is
    // RowStatus.
    RsvpResvStatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    RsvpResvTTL interface{}

    // The	flow ID	that this receiver  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    RsvpResvFlowId interface{}
}

func (rsvpResvEntry *RSVPMIB_RsvpResvTable_RsvpResvEntry) GetEntityData() *types.CommonEntityData {
    rsvpResvEntry.EntityData.YFilter = rsvpResvEntry.YFilter
    rsvpResvEntry.EntityData.YangName = "rsvpResvEntry"
    rsvpResvEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpResvEntry.EntityData.ParentYangName = "rsvpResvTable"
    rsvpResvEntry.EntityData.SegmentPath = "rsvpResvEntry" + types.AddKeyToken(rsvpResvEntry.RsvpSessionNumber, "rsvpSessionNumber") + types.AddKeyToken(rsvpResvEntry.RsvpResvNumber, "rsvpResvNumber")
    rsvpResvEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpResvTable/" + rsvpResvEntry.EntityData.SegmentPath
    rsvpResvEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpResvEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpResvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpResvEntry.EntityData.Children = types.NewOrderedMap()
    rsvpResvEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpResvEntry.EntityData.Leafs.Append("rsvpSessionNumber", types.YLeaf{"RsvpSessionNumber", rsvpResvEntry.RsvpSessionNumber})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvNumber", types.YLeaf{"RsvpResvNumber", rsvpResvEntry.RsvpResvNumber})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvType", types.YLeaf{"RsvpResvType", rsvpResvEntry.RsvpResvType})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvDestAddr", types.YLeaf{"RsvpResvDestAddr", rsvpResvEntry.RsvpResvDestAddr})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvSenderAddr", types.YLeaf{"RsvpResvSenderAddr", rsvpResvEntry.RsvpResvSenderAddr})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvDestAddrLength", types.YLeaf{"RsvpResvDestAddrLength", rsvpResvEntry.RsvpResvDestAddrLength})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvSenderAddrLength", types.YLeaf{"RsvpResvSenderAddrLength", rsvpResvEntry.RsvpResvSenderAddrLength})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvProtocol", types.YLeaf{"RsvpResvProtocol", rsvpResvEntry.RsvpResvProtocol})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvDestPort", types.YLeaf{"RsvpResvDestPort", rsvpResvEntry.RsvpResvDestPort})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvPort", types.YLeaf{"RsvpResvPort", rsvpResvEntry.RsvpResvPort})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvHopAddr", types.YLeaf{"RsvpResvHopAddr", rsvpResvEntry.RsvpResvHopAddr})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvHopLih", types.YLeaf{"RsvpResvHopLih", rsvpResvEntry.RsvpResvHopLih})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvInterface", types.YLeaf{"RsvpResvInterface", rsvpResvEntry.RsvpResvInterface})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvService", types.YLeaf{"RsvpResvService", rsvpResvEntry.RsvpResvService})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTSpecRate", types.YLeaf{"RsvpResvTSpecRate", rsvpResvEntry.RsvpResvTSpecRate})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTSpecPeakRate", types.YLeaf{"RsvpResvTSpecPeakRate", rsvpResvEntry.RsvpResvTSpecPeakRate})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTSpecBurst", types.YLeaf{"RsvpResvTSpecBurst", rsvpResvEntry.RsvpResvTSpecBurst})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTSpecMinTU", types.YLeaf{"RsvpResvTSpecMinTU", rsvpResvEntry.RsvpResvTSpecMinTU})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTSpecMaxTU", types.YLeaf{"RsvpResvTSpecMaxTU", rsvpResvEntry.RsvpResvTSpecMaxTU})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvRSpecRate", types.YLeaf{"RsvpResvRSpecRate", rsvpResvEntry.RsvpResvRSpecRate})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvRSpecSlack", types.YLeaf{"RsvpResvRSpecSlack", rsvpResvEntry.RsvpResvRSpecSlack})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvInterval", types.YLeaf{"RsvpResvInterval", rsvpResvEntry.RsvpResvInterval})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvScope", types.YLeaf{"RsvpResvScope", rsvpResvEntry.RsvpResvScope})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvShared", types.YLeaf{"RsvpResvShared", rsvpResvEntry.RsvpResvShared})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvExplicit", types.YLeaf{"RsvpResvExplicit", rsvpResvEntry.RsvpResvExplicit})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvRSVPHop", types.YLeaf{"RsvpResvRSVPHop", rsvpResvEntry.RsvpResvRSVPHop})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvLastChange", types.YLeaf{"RsvpResvLastChange", rsvpResvEntry.RsvpResvLastChange})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvPolicy", types.YLeaf{"RsvpResvPolicy", rsvpResvEntry.RsvpResvPolicy})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvStatus", types.YLeaf{"RsvpResvStatus", rsvpResvEntry.RsvpResvStatus})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvTTL", types.YLeaf{"RsvpResvTTL", rsvpResvEntry.RsvpResvTTL})
    rsvpResvEntry.EntityData.Leafs.Append("rsvpResvFlowId", types.YLeaf{"RsvpResvFlowId", rsvpResvEntry.RsvpResvFlowId})

    rsvpResvEntry.EntityData.YListKeys = []string {"RsvpSessionNumber", "RsvpResvNumber"}

    return &(rsvpResvEntry.EntityData)
}

// RSVPMIB_RsvpResvFwdTable
// Information	describing the	state  information
// displayed upstream in RESV messages.
type RSVPMIB_RsvpResvFwdTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed   upstream	  in  
    // an   RESV   message concerning a	single sender. The type is slice of
    // RSVPMIB_RsvpResvFwdTable_RsvpResvFwdEntry.
    RsvpResvFwdEntry []*RSVPMIB_RsvpResvFwdTable_RsvpResvFwdEntry
}

func (rsvpResvFwdTable *RSVPMIB_RsvpResvFwdTable) GetEntityData() *types.CommonEntityData {
    rsvpResvFwdTable.EntityData.YFilter = rsvpResvFwdTable.YFilter
    rsvpResvFwdTable.EntityData.YangName = "rsvpResvFwdTable"
    rsvpResvFwdTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpResvFwdTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpResvFwdTable.EntityData.SegmentPath = "rsvpResvFwdTable"
    rsvpResvFwdTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpResvFwdTable.EntityData.SegmentPath
    rsvpResvFwdTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpResvFwdTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpResvFwdTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpResvFwdTable.EntityData.Children = types.NewOrderedMap()
    rsvpResvFwdTable.EntityData.Children.Append("rsvpResvFwdEntry", types.YChild{"RsvpResvFwdEntry", nil})
    for i := range rsvpResvFwdTable.RsvpResvFwdEntry {
        rsvpResvFwdTable.EntityData.Children.Append(types.GetSegmentPath(rsvpResvFwdTable.RsvpResvFwdEntry[i]), types.YChild{"RsvpResvFwdEntry", rsvpResvFwdTable.RsvpResvFwdEntry[i]})
    }
    rsvpResvFwdTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpResvFwdTable.EntityData.YListKeys = []string {}

    return &(rsvpResvFwdTable.EntityData)
}

// RSVPMIB_RsvpResvFwdTable_RsvpResvFwdEntry
// Information	describing the	state  information
// displayed   upstream	  in   an   RESV   message
// concerning a	single sender.
type RSVPMIB_RsvpResvFwdTable_RsvpResvFwdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_RsvpSessionTable_RsvpSessionEntry_RsvpSessionNumber
    RsvpSessionNumber interface{}

    // This attribute is a key. The	number of this reservation request.   This is 
    // for  SNMP Indexing purposes only	and has	no relation to any protocol value.
    // The type is interface{} with range: 0..2147483647.
    RsvpResvFwdNumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    RsvpResvFwdType interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    RsvpResvFwdDestAddr interface{}

    // The	source address of the sender  selected	by this	 reservation.	The  value
    // of	all zeroes indicates 'all senders'.  This object  may  not be 
    // changed	when  the  value  of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    RsvpResvFwdSenderAddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpResvFwdDestAddrLength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    RsvpResvFwdSenderAddrLength interface{}

    // The	IP Protocol used by a session. for  secure sessions,  this  indicates 
    // IP  Security.  This object may not be changed when the value of the
    // RowStatus object is 'active'. The type is interface{} with range: 1..255.
    RsvpResvFwdProtocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by
    // rsvpResvFwdProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    RsvpResvFwdDestPort interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpResvFwdProtocol	is  50	(ESP)  or 51 (AH), this	represents a generalized
    // port	identifier (GPI).   A  value of	zero indicates that the	IP protocol in
    // use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    RsvpResvFwdPort interface{}

    // The	address	of the (previous) RSVP	that  will receive this	message. The
    // type is string with length: 4..16.
    RsvpResvFwdHopAddr interface{}

    // The	 Logical  Interface  Handle  sent  to  the (previous)	RSVP   that  
    // will   receive  this message. The type is interface{} with range:
    // -2147483648..2147483647.
    RsvpResvFwdHopLih interface{}

    // The	ifIndex	value of the  interface	 on  which this	RESV message was most
    // recently sent. The type is interface{} with range: 1..2147483647.
    RsvpResvFwdInterface interface{}

    // The	QoS Service classification requested. The type is QosService.
    RsvpResvFwdService interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpResvFwdTSpecPeakRate  (if  supported by the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpResvFwdTSpecRate.  Note	that this is a prediction, often
    // based	on the	general	 capability  of	a type of codec	or particular
    // encoding;	the measured average  rate may be significantly	lower. The type
    // is interface{} with range: 0..2147483647. Units are bits per second.
    RsvpResvFwdTSpecRate interface{}

    // The	Peak Bit Rate of the sender's data  stream Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    RsvpResvFwdTSpecPeakRate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time.  If
    // this is less than	 the  sender's	advertised burst  size,	the receiver
    // is	asking the network to provide flow pacing  beyond  what	 would	be
    // provided   under   normal  circumstances.  Such pacing is at	the network's
    // option. The type is interface{} with range: 0..2147483647. Units are bytes.
    RsvpResvFwdTSpecBurst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    RsvpResvFwdTSpecMinTU interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    RsvpResvFwdTSpecMaxTU interface{}

    // If the requested  service  is  Guaranteed,	as specified   by 
    // rsvpResvService,  this  is  the clearing  rate   that   is	being	requested.
    // Otherwise,  it is zero, or the agent	may return noSuchValue. The type is
    // interface{} with range: 0..2147483647. Units are bytes per second.
    RsvpResvFwdRSpecRate interface{}

    // If the requested  service  is  Guaranteed,	as specified by	rsvpResvService,
    // this is the delay slack.  Otherwise, it is zero, or the agent may return
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    RsvpResvFwdRSpecSlack interface{}

    // The	  interval   between   refresh	  messages advertised to the Previous
    // Hop. The type is interface{} with range: 0..2147483647.
    RsvpResvFwdInterval interface{}

    // The	contents of the	scope object, displayed	as an  uninterpreted  string 
    // of octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    RsvpResvFwdScope interface{}

    // If TRUE, a reservation shared among	senders	is requested.  If FALSE, a
    // reservation specific	to this	sender is requested. The type is bool.
    RsvpResvFwdShared interface{}

    // If TRUE, individual	senders	are  listed  using Filter  Specifications.  
    // If	FALSE, all senders are implicitly selected.  The Scope Object will
    // contain  a list of senders that need	to receive this	reservation request 
    // for  the  purpose	of routing the RESV message. The type is bool.
    RsvpResvFwdExplicit interface{}

    // If TRUE, the node believes that  the  next	IP hop	is  an	RSVP  hop.   If	
    // FALSE,	 the  node believes that the next IP hop  may  not  be	an RSVP	hop.
    // The type is bool.
    RsvpResvFwdRSVPHop interface{}

    // The	time of	the last change	in  this  request; This	 is  either  the first
    // time it was sent	or the	time  of  the  most   recent   change	in
    // parameters. The type is interface{} with range: 0..4294967295.
    RsvpResvFwdLastChange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    RsvpResvFwdPolicy interface{}

    // 'active' for all active RESV  messages.   This object may be used to
    // delete	RESV information. The type is RowStatus.
    RsvpResvFwdStatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    RsvpResvFwdTTL interface{}

    // The	flow ID	that this receiver  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    RsvpResvFwdFlowId interface{}
}

func (rsvpResvFwdEntry *RSVPMIB_RsvpResvFwdTable_RsvpResvFwdEntry) GetEntityData() *types.CommonEntityData {
    rsvpResvFwdEntry.EntityData.YFilter = rsvpResvFwdEntry.YFilter
    rsvpResvFwdEntry.EntityData.YangName = "rsvpResvFwdEntry"
    rsvpResvFwdEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpResvFwdEntry.EntityData.ParentYangName = "rsvpResvFwdTable"
    rsvpResvFwdEntry.EntityData.SegmentPath = "rsvpResvFwdEntry" + types.AddKeyToken(rsvpResvFwdEntry.RsvpSessionNumber, "rsvpSessionNumber") + types.AddKeyToken(rsvpResvFwdEntry.RsvpResvFwdNumber, "rsvpResvFwdNumber")
    rsvpResvFwdEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpResvFwdTable/" + rsvpResvFwdEntry.EntityData.SegmentPath
    rsvpResvFwdEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpResvFwdEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpResvFwdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpResvFwdEntry.EntityData.Children = types.NewOrderedMap()
    rsvpResvFwdEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpSessionNumber", types.YLeaf{"RsvpSessionNumber", rsvpResvFwdEntry.RsvpSessionNumber})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdNumber", types.YLeaf{"RsvpResvFwdNumber", rsvpResvFwdEntry.RsvpResvFwdNumber})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdType", types.YLeaf{"RsvpResvFwdType", rsvpResvFwdEntry.RsvpResvFwdType})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdDestAddr", types.YLeaf{"RsvpResvFwdDestAddr", rsvpResvFwdEntry.RsvpResvFwdDestAddr})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdSenderAddr", types.YLeaf{"RsvpResvFwdSenderAddr", rsvpResvFwdEntry.RsvpResvFwdSenderAddr})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdDestAddrLength", types.YLeaf{"RsvpResvFwdDestAddrLength", rsvpResvFwdEntry.RsvpResvFwdDestAddrLength})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdSenderAddrLength", types.YLeaf{"RsvpResvFwdSenderAddrLength", rsvpResvFwdEntry.RsvpResvFwdSenderAddrLength})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdProtocol", types.YLeaf{"RsvpResvFwdProtocol", rsvpResvFwdEntry.RsvpResvFwdProtocol})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdDestPort", types.YLeaf{"RsvpResvFwdDestPort", rsvpResvFwdEntry.RsvpResvFwdDestPort})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdPort", types.YLeaf{"RsvpResvFwdPort", rsvpResvFwdEntry.RsvpResvFwdPort})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdHopAddr", types.YLeaf{"RsvpResvFwdHopAddr", rsvpResvFwdEntry.RsvpResvFwdHopAddr})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdHopLih", types.YLeaf{"RsvpResvFwdHopLih", rsvpResvFwdEntry.RsvpResvFwdHopLih})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdInterface", types.YLeaf{"RsvpResvFwdInterface", rsvpResvFwdEntry.RsvpResvFwdInterface})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdService", types.YLeaf{"RsvpResvFwdService", rsvpResvFwdEntry.RsvpResvFwdService})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTSpecRate", types.YLeaf{"RsvpResvFwdTSpecRate", rsvpResvFwdEntry.RsvpResvFwdTSpecRate})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTSpecPeakRate", types.YLeaf{"RsvpResvFwdTSpecPeakRate", rsvpResvFwdEntry.RsvpResvFwdTSpecPeakRate})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTSpecBurst", types.YLeaf{"RsvpResvFwdTSpecBurst", rsvpResvFwdEntry.RsvpResvFwdTSpecBurst})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTSpecMinTU", types.YLeaf{"RsvpResvFwdTSpecMinTU", rsvpResvFwdEntry.RsvpResvFwdTSpecMinTU})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTSpecMaxTU", types.YLeaf{"RsvpResvFwdTSpecMaxTU", rsvpResvFwdEntry.RsvpResvFwdTSpecMaxTU})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdRSpecRate", types.YLeaf{"RsvpResvFwdRSpecRate", rsvpResvFwdEntry.RsvpResvFwdRSpecRate})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdRSpecSlack", types.YLeaf{"RsvpResvFwdRSpecSlack", rsvpResvFwdEntry.RsvpResvFwdRSpecSlack})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdInterval", types.YLeaf{"RsvpResvFwdInterval", rsvpResvFwdEntry.RsvpResvFwdInterval})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdScope", types.YLeaf{"RsvpResvFwdScope", rsvpResvFwdEntry.RsvpResvFwdScope})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdShared", types.YLeaf{"RsvpResvFwdShared", rsvpResvFwdEntry.RsvpResvFwdShared})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdExplicit", types.YLeaf{"RsvpResvFwdExplicit", rsvpResvFwdEntry.RsvpResvFwdExplicit})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdRSVPHop", types.YLeaf{"RsvpResvFwdRSVPHop", rsvpResvFwdEntry.RsvpResvFwdRSVPHop})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdLastChange", types.YLeaf{"RsvpResvFwdLastChange", rsvpResvFwdEntry.RsvpResvFwdLastChange})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdPolicy", types.YLeaf{"RsvpResvFwdPolicy", rsvpResvFwdEntry.RsvpResvFwdPolicy})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdStatus", types.YLeaf{"RsvpResvFwdStatus", rsvpResvFwdEntry.RsvpResvFwdStatus})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdTTL", types.YLeaf{"RsvpResvFwdTTL", rsvpResvFwdEntry.RsvpResvFwdTTL})
    rsvpResvFwdEntry.EntityData.Leafs.Append("rsvpResvFwdFlowId", types.YLeaf{"RsvpResvFwdFlowId", rsvpResvFwdEntry.RsvpResvFwdFlowId})

    rsvpResvFwdEntry.EntityData.YListKeys = []string {"RsvpSessionNumber", "RsvpResvFwdNumber"}

    return &(rsvpResvFwdEntry.EntityData)
}

// RSVPMIB_RsvpIfTable
// The	RSVP-specific attributes of  the  system's
// interfaces.
type RSVPMIB_RsvpIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The	RSVP-specific attributes of  the  a  given interface. The type is slice
    // of RSVPMIB_RsvpIfTable_RsvpIfEntry.
    RsvpIfEntry []*RSVPMIB_RsvpIfTable_RsvpIfEntry
}

func (rsvpIfTable *RSVPMIB_RsvpIfTable) GetEntityData() *types.CommonEntityData {
    rsvpIfTable.EntityData.YFilter = rsvpIfTable.YFilter
    rsvpIfTable.EntityData.YangName = "rsvpIfTable"
    rsvpIfTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpIfTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpIfTable.EntityData.SegmentPath = "rsvpIfTable"
    rsvpIfTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpIfTable.EntityData.SegmentPath
    rsvpIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpIfTable.EntityData.Children = types.NewOrderedMap()
    rsvpIfTable.EntityData.Children.Append("rsvpIfEntry", types.YChild{"RsvpIfEntry", nil})
    for i := range rsvpIfTable.RsvpIfEntry {
        rsvpIfTable.EntityData.Children.Append(types.GetSegmentPath(rsvpIfTable.RsvpIfEntry[i]), types.YChild{"RsvpIfEntry", rsvpIfTable.RsvpIfEntry[i]})
    }
    rsvpIfTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpIfTable.EntityData.YListKeys = []string {}

    return &(rsvpIfTable.EntityData)
}

// RSVPMIB_RsvpIfTable_RsvpIfEntry
// The	RSVP-specific attributes of  the  a  given
// interface.
type RSVPMIB_RsvpIfTable_RsvpIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The	number of neighbors perceived to be  using only	the RSVP UDP
    // Encapsulation. The type is interface{} with range: 0..4294967295.
    RsvpIfUdpNbrs interface{}

    // The	number of neighbors perceived to be  using only	the RSVP IP
    // Encapsulation. The type is interface{} with range: 0..4294967295.
    RsvpIfIpNbrs interface{}

    // The	number of neighbors  currently	perceived; this	 will  exceed
    // rsvpIfIpNbrs + rsvpIfUdpNbrs by  the  number   of	  neighbors   using  
    // both encapsulations. The type is interface{} with range: 0..4294967295.
    RsvpIfNbrs interface{}

    // The	value of the RSVP value	'Kb', Which is the minimum   number   of 
    // refresh  intervals  that blockade state will last once entered. The type is
    // interface{} with range: 1..65536.
    RsvpIfRefreshBlockadeMultiple interface{}

    // The	value of the RSVP value	'K', which is  the number  of  refresh
    // intervals which must elapse (minimum) before a PATH or RESV  message  which
    // is not being	refreshed will be aged out. The type is interface{} with
    // range: 1..65536.
    RsvpIfRefreshMultiple interface{}

    // The	value of SEND_TTL used on  this	 interface for	messages  this node
    // originates.	 If set	to zero, the node determines  the  TTL	via  other
    // means. The type is interface{} with range: 0..255.
    RsvpIfTTL interface{}

    // The	value of the RSVP value	'R', which is  the minimum period between
    // refresh transmissions	of a given PATH	or RESV	message	on an interface. The
    // type is interface{} with range: 0..2147483647. Units are milliseconds.
    RsvpIfRefreshInterval interface{}

    // The	approximate period from	the time  a  route is  changed	to  the	 time  a
    // resulting message appears on the interface. The type is interface{} with
    // range: 0..2147483647. Units are hundredths	of a second.
    RsvpIfRouteDelay interface{}

    // If TRUE, RSVP is enabled  on  this	Interface. If	FALSE,	 RSVP	is  not	
    // enabled  on  this interface. The type is bool.
    RsvpIfEnabled interface{}

    // If TRUE, manual configuration forces  the  use of  UDP  encapsulation  on 
    // the  interface.	If FALSE,  UDP	encapsulation  is  only	 used	if
    // rsvpIfUdpNbrs is not	zero. The type is bool.
    RsvpIfUdpRequired interface{}

    // 'active' on	interfaces that	are configured for RSVP. The type is RowStatus.
    RsvpIfStatus interface{}
}

func (rsvpIfEntry *RSVPMIB_RsvpIfTable_RsvpIfEntry) GetEntityData() *types.CommonEntityData {
    rsvpIfEntry.EntityData.YFilter = rsvpIfEntry.YFilter
    rsvpIfEntry.EntityData.YangName = "rsvpIfEntry"
    rsvpIfEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpIfEntry.EntityData.ParentYangName = "rsvpIfTable"
    rsvpIfEntry.EntityData.SegmentPath = "rsvpIfEntry" + types.AddKeyToken(rsvpIfEntry.IfIndex, "ifIndex")
    rsvpIfEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpIfTable/" + rsvpIfEntry.EntityData.SegmentPath
    rsvpIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpIfEntry.EntityData.Children = types.NewOrderedMap()
    rsvpIfEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", rsvpIfEntry.IfIndex})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfUdpNbrs", types.YLeaf{"RsvpIfUdpNbrs", rsvpIfEntry.RsvpIfUdpNbrs})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfIpNbrs", types.YLeaf{"RsvpIfIpNbrs", rsvpIfEntry.RsvpIfIpNbrs})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfNbrs", types.YLeaf{"RsvpIfNbrs", rsvpIfEntry.RsvpIfNbrs})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfRefreshBlockadeMultiple", types.YLeaf{"RsvpIfRefreshBlockadeMultiple", rsvpIfEntry.RsvpIfRefreshBlockadeMultiple})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfRefreshMultiple", types.YLeaf{"RsvpIfRefreshMultiple", rsvpIfEntry.RsvpIfRefreshMultiple})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfTTL", types.YLeaf{"RsvpIfTTL", rsvpIfEntry.RsvpIfTTL})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfRefreshInterval", types.YLeaf{"RsvpIfRefreshInterval", rsvpIfEntry.RsvpIfRefreshInterval})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfRouteDelay", types.YLeaf{"RsvpIfRouteDelay", rsvpIfEntry.RsvpIfRouteDelay})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfEnabled", types.YLeaf{"RsvpIfEnabled", rsvpIfEntry.RsvpIfEnabled})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfUdpRequired", types.YLeaf{"RsvpIfUdpRequired", rsvpIfEntry.RsvpIfUdpRequired})
    rsvpIfEntry.EntityData.Leafs.Append("rsvpIfStatus", types.YLeaf{"RsvpIfStatus", rsvpIfEntry.RsvpIfStatus})

    rsvpIfEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(rsvpIfEntry.EntityData)
}

// RSVPMIB_RsvpNbrTable
// Information	describing  the	 Neighbors  of	an
// RSVP	system.
type RSVPMIB_RsvpNbrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	  describing   a    single    RSVP Neighbor. The type is slice
    // of RSVPMIB_RsvpNbrTable_RsvpNbrEntry.
    RsvpNbrEntry []*RSVPMIB_RsvpNbrTable_RsvpNbrEntry
}

func (rsvpNbrTable *RSVPMIB_RsvpNbrTable) GetEntityData() *types.CommonEntityData {
    rsvpNbrTable.EntityData.YFilter = rsvpNbrTable.YFilter
    rsvpNbrTable.EntityData.YangName = "rsvpNbrTable"
    rsvpNbrTable.EntityData.BundleName = "cisco_ios_xe"
    rsvpNbrTable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpNbrTable.EntityData.SegmentPath = "rsvpNbrTable"
    rsvpNbrTable.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/" + rsvpNbrTable.EntityData.SegmentPath
    rsvpNbrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpNbrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpNbrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpNbrTable.EntityData.Children = types.NewOrderedMap()
    rsvpNbrTable.EntityData.Children.Append("rsvpNbrEntry", types.YChild{"RsvpNbrEntry", nil})
    for i := range rsvpNbrTable.RsvpNbrEntry {
        rsvpNbrTable.EntityData.Children.Append(types.GetSegmentPath(rsvpNbrTable.RsvpNbrEntry[i]), types.YChild{"RsvpNbrEntry", rsvpNbrTable.RsvpNbrEntry[i]})
    }
    rsvpNbrTable.EntityData.Leafs = types.NewOrderedMap()

    rsvpNbrTable.EntityData.YListKeys = []string {}

    return &(rsvpNbrTable.EntityData)
}

// RSVPMIB_RsvpNbrTable_RsvpNbrEntry
// Information	  describing   a    single    RSVP
// Neighbor.
type RSVPMIB_RsvpNbrTable_RsvpNbrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The	IP4 or IP6 Address used	by this	 neighbor.
    // This	 object	 may not be changed when the value of the RowStatus object is
    // 'active'. The type is string with length: 4..16.
    RsvpNbrAddress interface{}

    // The	  encapsulation	  being	  used	 by   this neighbor. The type is
    // RsvpEncapsulation.
    RsvpNbrProtocol interface{}

    // 'active' for all neighbors.	 This  object  may be	used   to  configure 
    // neighbors.   In  the presence   of   configured	 neighbors,    the
    // implementation  may	(but  is  not required to) limit the  set  of  valid 
    // neighbors	 to  those configured. The type is RowStatus.
    RsvpNbrStatus interface{}
}

func (rsvpNbrEntry *RSVPMIB_RsvpNbrTable_RsvpNbrEntry) GetEntityData() *types.CommonEntityData {
    rsvpNbrEntry.EntityData.YFilter = rsvpNbrEntry.YFilter
    rsvpNbrEntry.EntityData.YangName = "rsvpNbrEntry"
    rsvpNbrEntry.EntityData.BundleName = "cisco_ios_xe"
    rsvpNbrEntry.EntityData.ParentYangName = "rsvpNbrTable"
    rsvpNbrEntry.EntityData.SegmentPath = "rsvpNbrEntry" + types.AddKeyToken(rsvpNbrEntry.IfIndex, "ifIndex") + types.AddKeyToken(rsvpNbrEntry.RsvpNbrAddress, "rsvpNbrAddress")
    rsvpNbrEntry.EntityData.AbsolutePath = "RSVP-MIB:RSVP-MIB/rsvpNbrTable/" + rsvpNbrEntry.EntityData.SegmentPath
    rsvpNbrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpNbrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpNbrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpNbrEntry.EntityData.Children = types.NewOrderedMap()
    rsvpNbrEntry.EntityData.Leafs = types.NewOrderedMap()
    rsvpNbrEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", rsvpNbrEntry.IfIndex})
    rsvpNbrEntry.EntityData.Leafs.Append("rsvpNbrAddress", types.YLeaf{"RsvpNbrAddress", rsvpNbrEntry.RsvpNbrAddress})
    rsvpNbrEntry.EntityData.Leafs.Append("rsvpNbrProtocol", types.YLeaf{"RsvpNbrProtocol", rsvpNbrEntry.RsvpNbrProtocol})
    rsvpNbrEntry.EntityData.Leafs.Append("rsvpNbrStatus", types.YLeaf{"RsvpNbrStatus", rsvpNbrEntry.RsvpNbrStatus})

    rsvpNbrEntry.EntityData.YListKeys = []string {"IfIndex", "RsvpNbrAddress"}

    return &(rsvpNbrEntry.EntityData)
}

