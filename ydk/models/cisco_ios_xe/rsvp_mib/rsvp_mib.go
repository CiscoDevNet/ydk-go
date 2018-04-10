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

    
    Rsvpgenobjects RSVPMIB_Rsvpgenobjects

    // A table  of	 all  sessions	seen  by  a  given system.
    Rsvpsessiontable RSVPMIB_Rsvpsessiontable

    // Information	describing the	state  information displayed by	senders	in
    // PATH	messages.
    Rsvpsendertable RSVPMIB_Rsvpsendertable

    // List of outgoing interfaces	that PATH messages use.	 The  ifIndex  is the
    // ifIndex value of the egress interface.
    Rsvpsenderoutinterfacetable RSVPMIB_Rsvpsenderoutinterfacetable

    // Information	describing the	state  information displayed by	receivers in
    // RESV messages.
    Rsvpresvtable RSVPMIB_Rsvpresvtable

    // Information	describing the	state  information displayed upstream in RESV
    // messages.
    Rsvpresvfwdtable RSVPMIB_Rsvpresvfwdtable

    // The	RSVP-specific attributes of  the  system's interfaces.
    Rsvpiftable RSVPMIB_Rsvpiftable

    // Information	describing  the	 Neighbors  of	an RSVP	system.
    Rsvpnbrtable RSVPMIB_Rsvpnbrtable
}

func (rSVPMIB *RSVPMIB) GetEntityData() *types.CommonEntityData {
    rSVPMIB.EntityData.YFilter = rSVPMIB.YFilter
    rSVPMIB.EntityData.YangName = "RSVP-MIB"
    rSVPMIB.EntityData.BundleName = "cisco_ios_xe"
    rSVPMIB.EntityData.ParentYangName = "RSVP-MIB"
    rSVPMIB.EntityData.SegmentPath = "RSVP-MIB:RSVP-MIB"
    rSVPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rSVPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rSVPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rSVPMIB.EntityData.Children = make(map[string]types.YChild)
    rSVPMIB.EntityData.Children["rsvpGenObjects"] = types.YChild{"Rsvpgenobjects", &rSVPMIB.Rsvpgenobjects}
    rSVPMIB.EntityData.Children["rsvpSessionTable"] = types.YChild{"Rsvpsessiontable", &rSVPMIB.Rsvpsessiontable}
    rSVPMIB.EntityData.Children["rsvpSenderTable"] = types.YChild{"Rsvpsendertable", &rSVPMIB.Rsvpsendertable}
    rSVPMIB.EntityData.Children["rsvpSenderOutInterfaceTable"] = types.YChild{"Rsvpsenderoutinterfacetable", &rSVPMIB.Rsvpsenderoutinterfacetable}
    rSVPMIB.EntityData.Children["rsvpResvTable"] = types.YChild{"Rsvpresvtable", &rSVPMIB.Rsvpresvtable}
    rSVPMIB.EntityData.Children["rsvpResvFwdTable"] = types.YChild{"Rsvpresvfwdtable", &rSVPMIB.Rsvpresvfwdtable}
    rSVPMIB.EntityData.Children["rsvpIfTable"] = types.YChild{"Rsvpiftable", &rSVPMIB.Rsvpiftable}
    rSVPMIB.EntityData.Children["rsvpNbrTable"] = types.YChild{"Rsvpnbrtable", &rSVPMIB.Rsvpnbrtable}
    rSVPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rSVPMIB.EntityData)
}

// RSVPMIB_Rsvpgenobjects
type RSVPMIB_Rsvpgenobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object	keeps a	count of the number of bad RSVP	packets	received. The
    // type is interface{} with range: 0..4294967295.
    Rsvpbadpackets interface{}

    // This  object  is  used  to	assign	values	to rsvpSenderNumber   as  
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpSenderEntry.   If  the  SET  fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    Rsvpsendernewindex interface{}

    // This  object  is  used  to	assign	values	to rsvpResvNumber   as	 described 
    // in	  'Textual Conventions for SNMPv2'.  The  network  manager reads  the 
    // object,	and  then writes the value back	in the SET that	creates	a new
    // instance	of rsvpResvEntry.   If the SET fails with the code
    // 'inconsistentValue',	then the process  must	be repeated;  If the SET
    // succeeds, then	the object is incremented, and the new instance	is created
    // according to	the manager's directions. The type is interface{} with range:
    // 0..2147483647.
    Rsvpresvnewindex interface{}

    // This  object  is  used  to	assign	values	to rsvpResvFwdNumber   as 
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpResvFwdEntry.   If  the	SET fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    Rsvpresvfwdnewindex interface{}

    // This  object  is  used  to	assign	values	to rsvpSessionNumber   as 
    // described  in  'Textual Conventions for SNMPv2'.  The  network  manager
    // reads  the  object,	and  then writes the value back	in the SET
    // that	creates	a new instance	of rsvpSessionEntry.   If  the	SET fails with
    // the code	'inconsistentValue', then the process must be  repeated;  If  the 
    // SET	succeeds, then the object is incremented, and the new instance	is
    // created according to	the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    Rsvpsessionnewindex interface{}
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetEntityData() *types.CommonEntityData {
    rsvpgenobjects.EntityData.YFilter = rsvpgenobjects.YFilter
    rsvpgenobjects.EntityData.YangName = "rsvpGenObjects"
    rsvpgenobjects.EntityData.BundleName = "cisco_ios_xe"
    rsvpgenobjects.EntityData.ParentYangName = "RSVP-MIB"
    rsvpgenobjects.EntityData.SegmentPath = "rsvpGenObjects"
    rsvpgenobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpgenobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpgenobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpgenobjects.EntityData.Children = make(map[string]types.YChild)
    rsvpgenobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpgenobjects.EntityData.Leafs["rsvpBadPackets"] = types.YLeaf{"Rsvpbadpackets", rsvpgenobjects.Rsvpbadpackets}
    rsvpgenobjects.EntityData.Leafs["rsvpSenderNewIndex"] = types.YLeaf{"Rsvpsendernewindex", rsvpgenobjects.Rsvpsendernewindex}
    rsvpgenobjects.EntityData.Leafs["rsvpResvNewIndex"] = types.YLeaf{"Rsvpresvnewindex", rsvpgenobjects.Rsvpresvnewindex}
    rsvpgenobjects.EntityData.Leafs["rsvpResvFwdNewIndex"] = types.YLeaf{"Rsvpresvfwdnewindex", rsvpgenobjects.Rsvpresvfwdnewindex}
    rsvpgenobjects.EntityData.Leafs["rsvpSessionNewIndex"] = types.YLeaf{"Rsvpsessionnewindex", rsvpgenobjects.Rsvpsessionnewindex}
    return &(rsvpgenobjects.EntityData)
}

// RSVPMIB_Rsvpsessiontable
// A table  of	 all  sessions	seen  by  a  given
// system.
type RSVPMIB_Rsvpsessiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single session seen by a given system. The type is slice of
    // RSVPMIB_Rsvpsessiontable_Rsvpsessionentry.
    Rsvpsessionentry []RSVPMIB_Rsvpsessiontable_Rsvpsessionentry
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetEntityData() *types.CommonEntityData {
    rsvpsessiontable.EntityData.YFilter = rsvpsessiontable.YFilter
    rsvpsessiontable.EntityData.YangName = "rsvpSessionTable"
    rsvpsessiontable.EntityData.BundleName = "cisco_ios_xe"
    rsvpsessiontable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpsessiontable.EntityData.SegmentPath = "rsvpSessionTable"
    rsvpsessiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsessiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsessiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsessiontable.EntityData.Children = make(map[string]types.YChild)
    rsvpsessiontable.EntityData.Children["rsvpSessionEntry"] = types.YChild{"Rsvpsessionentry", nil}
    for i := range rsvpsessiontable.Rsvpsessionentry {
        rsvpsessiontable.EntityData.Children[types.GetSegmentPath(&rsvpsessiontable.Rsvpsessionentry[i])] = types.YChild{"Rsvpsessionentry", &rsvpsessiontable.Rsvpsessionentry[i]}
    }
    rsvpsessiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpsessiontable.EntityData)
}

// RSVPMIB_Rsvpsessiontable_Rsvpsessionentry
// A single session seen by a given system.
type RSVPMIB_Rsvpsessiontable_Rsvpsessionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The	number of this session.	 This is for  SNMP
    // Indexing  purposes  only and	has no relation	to any protocol	value. The
    // type is interface{} with range: 0..2147483647.
    Rsvpsessionnumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    Rsvpsessiontype interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    Rsvpsessiondestaddr interface{}

    // The	CIDR prefix length of the session address, which   is	32  for	 IP4 
    // host  and	 multicast addresses, and 128  for  IP6	 addresses.   This object
    // may not be changed when the value of the RowStatus object is 'active'. The
    // type is interface{} with range: 0..128.
    Rsvpsessiondestaddrlength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    Rsvpsessionprotocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpSenderProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    Rsvpsessionport interface{}

    // The	number of distinct senders currently known to be part of this session.
    // The type is interface{} with range: 0..4294967295.
    Rsvpsessionsenders interface{}

    // The	number of reservations being requested	of this	system for this	session.
    // The type is interface{} with range: 0..4294967295.
    Rsvpsessionreceivers interface{}

    // The	number of reservation requests this system is sending upstream for
    // this	session. The type is interface{} with range: 0..4294967295.
    Rsvpsessionrequests interface{}
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetEntityData() *types.CommonEntityData {
    rsvpsessionentry.EntityData.YFilter = rsvpsessionentry.YFilter
    rsvpsessionentry.EntityData.YangName = "rsvpSessionEntry"
    rsvpsessionentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpsessionentry.EntityData.ParentYangName = "rsvpSessionTable"
    rsvpsessionentry.EntityData.SegmentPath = "rsvpSessionEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsessionentry.Rsvpsessionnumber) + "']"
    rsvpsessionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsessionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsessionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsessionentry.EntityData.Children = make(map[string]types.YChild)
    rsvpsessionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpsessionentry.EntityData.Leafs["rsvpSessionNumber"] = types.YLeaf{"Rsvpsessionnumber", rsvpsessionentry.Rsvpsessionnumber}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionType"] = types.YLeaf{"Rsvpsessiontype", rsvpsessionentry.Rsvpsessiontype}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionDestAddr"] = types.YLeaf{"Rsvpsessiondestaddr", rsvpsessionentry.Rsvpsessiondestaddr}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionDestAddrLength"] = types.YLeaf{"Rsvpsessiondestaddrlength", rsvpsessionentry.Rsvpsessiondestaddrlength}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionProtocol"] = types.YLeaf{"Rsvpsessionprotocol", rsvpsessionentry.Rsvpsessionprotocol}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionPort"] = types.YLeaf{"Rsvpsessionport", rsvpsessionentry.Rsvpsessionport}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionSenders"] = types.YLeaf{"Rsvpsessionsenders", rsvpsessionentry.Rsvpsessionsenders}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionReceivers"] = types.YLeaf{"Rsvpsessionreceivers", rsvpsessionentry.Rsvpsessionreceivers}
    rsvpsessionentry.EntityData.Leafs["rsvpSessionRequests"] = types.YLeaf{"Rsvpsessionrequests", rsvpsessionentry.Rsvpsessionrequests}
    return &(rsvpsessionentry.EntityData)
}

// RSVPMIB_Rsvpsendertable
// Information	describing the	state  information
// displayed by	senders	in PATH	messages.
type RSVPMIB_Rsvpsendertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed by	a single
    // sender's PATH message. The type is slice of
    // RSVPMIB_Rsvpsendertable_Rsvpsenderentry.
    Rsvpsenderentry []RSVPMIB_Rsvpsendertable_Rsvpsenderentry
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetEntityData() *types.CommonEntityData {
    rsvpsendertable.EntityData.YFilter = rsvpsendertable.YFilter
    rsvpsendertable.EntityData.YangName = "rsvpSenderTable"
    rsvpsendertable.EntityData.BundleName = "cisco_ios_xe"
    rsvpsendertable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpsendertable.EntityData.SegmentPath = "rsvpSenderTable"
    rsvpsendertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsendertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsendertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsendertable.EntityData.Children = make(map[string]types.YChild)
    rsvpsendertable.EntityData.Children["rsvpSenderEntry"] = types.YChild{"Rsvpsenderentry", nil}
    for i := range rsvpsendertable.Rsvpsenderentry {
        rsvpsendertable.EntityData.Children[types.GetSegmentPath(&rsvpsendertable.Rsvpsenderentry[i])] = types.YChild{"Rsvpsenderentry", &rsvpsendertable.Rsvpsenderentry[i]}
    }
    rsvpsendertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpsendertable.EntityData)
}

// RSVPMIB_Rsvpsendertable_Rsvpsenderentry
// Information	describing the	state  information
// displayed by	a single sender's PATH message.
type RSVPMIB_Rsvpsendertable_Rsvpsenderentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_Rsvpsessiontable_Rsvpsessionentry_Rsvpsessionnumber
    Rsvpsessionnumber interface{}

    // This attribute is a key. The	number of this sender.	This is	 for  SNMP
    // Indexing  purposes  only and	has no relation	to any protocol	value. The
    // type is interface{} with range: 0..2147483647.
    Rsvpsendernumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    Rsvpsendertype interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    Rsvpsenderdestaddr interface{}

    // The	source address used by this sender in this session.   This  object may
    // not be changed when the value of	the RowStatus object is	'active'. The type
    // is string with length: 4..16.
    Rsvpsenderaddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpsenderdestaddrlength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpsenderaddrlength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    Rsvpsenderprotocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpSenderProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    Rsvpsenderdestport interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpSenderProtocol is 50 (ESP) or 51	(AH), this represents a	generalized
    // port identifier (GPI). A  value of zero indicates that the IP protocol in
    // use does not have	ports.	 This  object  may not	be changed when	the value
    // of the RowStatus object is 'active'. The type is string with length: 2..4.
    Rsvpsenderport interface{}

    // The	flow ID	that  this  sender  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    Rsvpsenderflowid interface{}

    // The	address	used  by  the  previous	 RSVP  hop (which may be the original
    // sender). The type is string with length: 4..16.
    Rsvpsenderhopaddr interface{}

    // The	 Logical  Interface  Handle  used  by  the previous  RSVP  hop	(which
    // may be the original sender). The type is interface{} with range:
    // -2147483648..2147483647.
    Rsvpsenderhoplih interface{}

    // The	ifIndex	value of the  interface	 on  which this	PATH message was most
    // recently received. The type is interface{} with range: 1..2147483647.
    Rsvpsenderinterface interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpSenderTSpecPeakRate  (if	 supported  by the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpSenderTSpecRate.  Note	that this is a prediction, often based	on
    // the	general	 capability  of	a type of codec	or particular encoding;	the
    // measured average  rate may be significantly	lower. The type is interface{}
    // with range: 0..2147483647. Units are bits per second.
    Rsvpsendertspecrate interface{}

    // The	Peak Bit Rate of the sender's data stream. Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    Rsvpsendertspecpeakrate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time. The
    // type is interface{} with range: 0..2147483647. Units are bytes.
    Rsvpsendertspecburst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    Rsvpsendertspecmintu interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    Rsvpsendertspecmaxtu interface{}

    // The	 interval  between  refresh  messages	as advertised by the Previous
    // Hop. The type is interface{} with range: 0..2147483647.
    Rsvpsenderinterval interface{}

    // If TRUE, the node believes that  the  previous IP  hop  is	an  RSVP 
    // hop.	If FALSE, the node believes that the previous IP hop may not be	an
    // RSVP	hop. The type is bool.
    Rsvpsenderrsvphop interface{}

    // The	time of	 the  last  change  in	this  PATH message;  This  is either the
    // first time it was received or the time	of the most recent  change in
    // parameters. The type is interface{} with range: 0..4294967295.
    Rsvpsenderlastchange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 4..65536.
    Rsvpsenderpolicy interface{}

    // The	global break bit general  characterization parameter  from 
    // the	ADSPEC.	 If TRUE, at least one non-IS hop was detected in  the	path.	If
    // FALSE, no non-IS hops were detected. The type is bool.
    Rsvpsenderadspecbreak interface{}

    // The	  hop	count	general	  characterization parameter from the ADSPEC.  A
    // return	of zero	or noSuchValue	indicates  one	of  the	 following conditions:
    // the invalid bit was set    the parameter was	not present. The type is
    // interface{} with range: 0..65535.
    Rsvpsenderadspechopcount interface{}

    // The	  path	  bandwidth    estimate	   general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: 0..2147483647. Units are bits
    // per second.
    Rsvpsenderadspecpathbw interface{}

    // The	   minimum    path     latency	   general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: -2147483648..2147483647. Units
    // are microseconds.
    Rsvpsenderadspecminlatency interface{}

    // The	composed Maximum Transmission Unit general characterization  parameter
    // from the	ADSPEC.	 A return of zero or noSuchValue indicates one	of the
    // following conditions:     the invalid bit was set    the parameter was	not
    // present. The type is interface{} with range: 0..65535. Units are bytes.
    Rsvpsenderadspecmtu interface{}

    // If TRUE,  the  ADSPEC  contains  a	Guaranteed Service  fragment.	If  FALSE,
    // the ADSPEC does not contain a Guaranteed Service fragment. The type is
    // bool.
    Rsvpsenderadspecguaranteedsvc interface{}

    // If TRUE, the Guaranteed Service  fragment  has its	'break'	 bit  set, 
    // indicating that one	or more	nodes along the	path do	 not  support  the
    // guaranteed	  service.     If    FALSE,    and rsvpSenderAdspecGuaranteedSvc
    // is   TRUE,   the 'break' bit is not set.  If rsvpSenderAdspecGuaranteedSvc
    // is FALSE, this returns FALSE or noSuchValue. The type is bool.
    Rsvpsenderadspecguaranteedbreak interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  end-to-end	 composed
    // value  for  the guaranteed service 'C' parameter.  A	return	of zero	  or 
    // noSuchValue  indicates  one  of  the following conditions:     the invalid
    // bit was set    the parameter was	not present  If
    // rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or noSuchValue.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // bytes.
    Rsvpsenderadspecguaranteedctot interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  end-to-end	 composed
    // value  for  the guaranteed service 'D' parameter.  A	return	of zero	  or 
    // noSuchValue  indicates  one  of  the following conditions:     the invalid
    // bit was set    the parameter was	not present  If
    // rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or noSuchValue.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // microseconds.
    Rsvpsenderadspecguaranteeddtot interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  composed  value  for
    // the	guaranteed service 'C' parameter since the last	 reshaping point.    A	
    // return	 of  zero  or  noSuchValue indicates one of the	following
    // conditions:     the invalid bit was set    the parameter was	not present 
    // If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are bytes.
    Rsvpsenderadspecguaranteedcsum interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the  composed  value  for
    // the	guaranteed service 'D' parameter since the last	 reshaping point.    A	
    // return	 of  zero  or  noSuchValue indicates one of the	following
    // conditions:     the invalid bit was set    the parameter was	not present 
    // If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns zero	or
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    Rsvpsenderadspecguaranteeddsum interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is  the  service-specific 
    // override	of the hop count general characterization  parameter  from
    // the	ADSPEC.	  A  return of zero or noSuchValue indicates one of
    // the	following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns
    // zero	or noSuchValue. The type is interface{} with range: 0..65535.
    Rsvpsenderadspecguaranteedhopcount interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is  the  service-specific 
    // override of the path bandwidth  estimate	general	  characterization
    // parameter from the ADSPEC.  A return	of zero	or noSuchValue	indicates 
    // one	of  the	 following conditions:     the invalid bit was set    the
    // parameter was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this
    // returns zero	or noSuchValue. The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    Rsvpsenderadspecguaranteedpathbw interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is the service-specific
    // override of the minimum path	latency	general	characterization parameter
    // from	  the	ADSPEC.	   A  return  of  zero	or noSuchValue	indicates  one	of 
    // the	 following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecGuaranteedSvc is FALSE, this returns
    // zero	or noSuchValue. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    Rsvpsenderadspecguaranteedminlatency interface{}

    // If rsvpSenderAdspecGuaranteedSvc is	TRUE, this is	the   service-specific	
    // override  of  the composed  Maximum  Transmission  Unit   general
    // characterization  parameter from the	ADSPEC.	 A return of zero or
    // noSuchValue indicates one	of the following conditions:     the invalid bit
    // was set    the parameter was	not present  If rsvpSenderAdspecGuaranteedSvc
    // is FALSE, this returns zero	or noSuchValue. The type is interface{} with
    // range: 0..65535. Units are bytes.
    Rsvpsenderadspecguaranteedmtu interface{}

    // If TRUE, the ADSPEC	contains a Controlled Load Service  fragment.	If 
    // FALSE, the ADSPEC does not	 contain   a   Controlled   Load   Service
    // fragment. The type is bool.
    Rsvpsenderadspecctrlloadsvc interface{}

    // If TRUE, the Controlled Load Service  fragment has its 'break' bit set,
    // indicating that one	or more	nodes along the	path do	 not  support  the
    // controlled	load   service.	   If  FALSE,  and rsvpSenderAdspecCtrlLoadSvc	
    // is   TRUE,    the 'break' bit is not set.  If rsvpSenderAdspecCtrlLoadSvc
    // is  FALSE,  this returns FALSE or noSuchValue. The type is bool.
    Rsvpsenderadspecctrlloadbreak interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is  the  service-specific 
    // override	of the hop count general characterization  parameter  from
    // the	ADSPEC.	  A  return of zero or noSuchValue indicates one of
    // the	following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this returns
    // zero	or noSuchValue. The type is interface{} with range: 0..65535.
    Rsvpsenderadspecctrlloadhopcount interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is  the  service-specific 
    // override of the path bandwidth  estimate	general	  characterization
    // parameter from the ADSPEC.  A return	of zero	or noSuchValue	indicates 
    // one	of  the	 following conditions:     the invalid bit was set    the
    // parameter was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this
    // returns zero	or noSuchValue. The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    Rsvpsenderadspecctrlloadpathbw interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is the service-specific
    // override of the minimum path	latency	general	characterization parameter
    // from	  the	ADSPEC.	   A  return  of  zero	or noSuchValue	indicates  one	of 
    // the	 following conditions:     the invalid bit was set    the parameter
    // was	not present  If rsvpSenderAdspecCtrlLoadSvc is  FALSE,  this returns
    // zero	or noSuchValue. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    Rsvpsenderadspecctrlloadminlatency interface{}

    // If rsvpSenderAdspecCtrlLoadSvc is  TRUE,  this is	the   service-specific	
    // override  of  the composed  Maximum  Transmission  Unit   general
    // characterization  parameter from the	ADSPEC.	 A return of zero or
    // noSuchValue indicates one	of the following conditions:     the invalid bit
    // was set    the parameter was	not present  If rsvpSenderAdspecCtrlLoadSvc is
    // FALSE,  this returns zero	or noSuchValue. The type is interface{} with
    // range: 0..65535. Units are bytes.
    Rsvpsenderadspecctrlloadmtu interface{}

    // 'active' for all active PATH  messages.   This object  may	be  used  to 
    // install  static PATH information or delete PATH information. The type is
    // RowStatus.
    Rsvpsenderstatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    Rsvpsenderttl interface{}
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetEntityData() *types.CommonEntityData {
    rsvpsenderentry.EntityData.YFilter = rsvpsenderentry.YFilter
    rsvpsenderentry.EntityData.YangName = "rsvpSenderEntry"
    rsvpsenderentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpsenderentry.EntityData.ParentYangName = "rsvpSenderTable"
    rsvpsenderentry.EntityData.SegmentPath = "rsvpSenderEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsenderentry.Rsvpsessionnumber) + "']" + "[rsvpSenderNumber='" + fmt.Sprintf("%v", rsvpsenderentry.Rsvpsendernumber) + "']"
    rsvpsenderentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsenderentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsenderentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsenderentry.EntityData.Children = make(map[string]types.YChild)
    rsvpsenderentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpsenderentry.EntityData.Leafs["rsvpSessionNumber"] = types.YLeaf{"Rsvpsessionnumber", rsvpsenderentry.Rsvpsessionnumber}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderNumber"] = types.YLeaf{"Rsvpsendernumber", rsvpsenderentry.Rsvpsendernumber}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderType"] = types.YLeaf{"Rsvpsendertype", rsvpsenderentry.Rsvpsendertype}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderDestAddr"] = types.YLeaf{"Rsvpsenderdestaddr", rsvpsenderentry.Rsvpsenderdestaddr}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAddr"] = types.YLeaf{"Rsvpsenderaddr", rsvpsenderentry.Rsvpsenderaddr}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderDestAddrLength"] = types.YLeaf{"Rsvpsenderdestaddrlength", rsvpsenderentry.Rsvpsenderdestaddrlength}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAddrLength"] = types.YLeaf{"Rsvpsenderaddrlength", rsvpsenderentry.Rsvpsenderaddrlength}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderProtocol"] = types.YLeaf{"Rsvpsenderprotocol", rsvpsenderentry.Rsvpsenderprotocol}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderDestPort"] = types.YLeaf{"Rsvpsenderdestport", rsvpsenderentry.Rsvpsenderdestport}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderPort"] = types.YLeaf{"Rsvpsenderport", rsvpsenderentry.Rsvpsenderport}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderFlowId"] = types.YLeaf{"Rsvpsenderflowid", rsvpsenderentry.Rsvpsenderflowid}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderHopAddr"] = types.YLeaf{"Rsvpsenderhopaddr", rsvpsenderentry.Rsvpsenderhopaddr}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderHopLih"] = types.YLeaf{"Rsvpsenderhoplih", rsvpsenderentry.Rsvpsenderhoplih}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderInterface"] = types.YLeaf{"Rsvpsenderinterface", rsvpsenderentry.Rsvpsenderinterface}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTSpecRate"] = types.YLeaf{"Rsvpsendertspecrate", rsvpsenderentry.Rsvpsendertspecrate}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTSpecPeakRate"] = types.YLeaf{"Rsvpsendertspecpeakrate", rsvpsenderentry.Rsvpsendertspecpeakrate}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTSpecBurst"] = types.YLeaf{"Rsvpsendertspecburst", rsvpsenderentry.Rsvpsendertspecburst}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTSpecMinTU"] = types.YLeaf{"Rsvpsendertspecmintu", rsvpsenderentry.Rsvpsendertspecmintu}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTSpecMaxTU"] = types.YLeaf{"Rsvpsendertspecmaxtu", rsvpsenderentry.Rsvpsendertspecmaxtu}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderInterval"] = types.YLeaf{"Rsvpsenderinterval", rsvpsenderentry.Rsvpsenderinterval}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderRSVPHop"] = types.YLeaf{"Rsvpsenderrsvphop", rsvpsenderentry.Rsvpsenderrsvphop}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderLastChange"] = types.YLeaf{"Rsvpsenderlastchange", rsvpsenderentry.Rsvpsenderlastchange}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderPolicy"] = types.YLeaf{"Rsvpsenderpolicy", rsvpsenderentry.Rsvpsenderpolicy}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecBreak"] = types.YLeaf{"Rsvpsenderadspecbreak", rsvpsenderentry.Rsvpsenderadspecbreak}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecHopCount"] = types.YLeaf{"Rsvpsenderadspechopcount", rsvpsenderentry.Rsvpsenderadspechopcount}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecPathBw"] = types.YLeaf{"Rsvpsenderadspecpathbw", rsvpsenderentry.Rsvpsenderadspecpathbw}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecMinLatency"] = types.YLeaf{"Rsvpsenderadspecminlatency", rsvpsenderentry.Rsvpsenderadspecminlatency}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecMtu"] = types.YLeaf{"Rsvpsenderadspecmtu", rsvpsenderentry.Rsvpsenderadspecmtu}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedSvc"] = types.YLeaf{"Rsvpsenderadspecguaranteedsvc", rsvpsenderentry.Rsvpsenderadspecguaranteedsvc}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedBreak"] = types.YLeaf{"Rsvpsenderadspecguaranteedbreak", rsvpsenderentry.Rsvpsenderadspecguaranteedbreak}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedCtot"] = types.YLeaf{"Rsvpsenderadspecguaranteedctot", rsvpsenderentry.Rsvpsenderadspecguaranteedctot}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedDtot"] = types.YLeaf{"Rsvpsenderadspecguaranteeddtot", rsvpsenderentry.Rsvpsenderadspecguaranteeddtot}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedCsum"] = types.YLeaf{"Rsvpsenderadspecguaranteedcsum", rsvpsenderentry.Rsvpsenderadspecguaranteedcsum}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedDsum"] = types.YLeaf{"Rsvpsenderadspecguaranteeddsum", rsvpsenderentry.Rsvpsenderadspecguaranteeddsum}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedHopCount"] = types.YLeaf{"Rsvpsenderadspecguaranteedhopcount", rsvpsenderentry.Rsvpsenderadspecguaranteedhopcount}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedPathBw"] = types.YLeaf{"Rsvpsenderadspecguaranteedpathbw", rsvpsenderentry.Rsvpsenderadspecguaranteedpathbw}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedMinLatency"] = types.YLeaf{"Rsvpsenderadspecguaranteedminlatency", rsvpsenderentry.Rsvpsenderadspecguaranteedminlatency}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecGuaranteedMtu"] = types.YLeaf{"Rsvpsenderadspecguaranteedmtu", rsvpsenderentry.Rsvpsenderadspecguaranteedmtu}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadSvc"] = types.YLeaf{"Rsvpsenderadspecctrlloadsvc", rsvpsenderentry.Rsvpsenderadspecctrlloadsvc}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadBreak"] = types.YLeaf{"Rsvpsenderadspecctrlloadbreak", rsvpsenderentry.Rsvpsenderadspecctrlloadbreak}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadHopCount"] = types.YLeaf{"Rsvpsenderadspecctrlloadhopcount", rsvpsenderentry.Rsvpsenderadspecctrlloadhopcount}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadPathBw"] = types.YLeaf{"Rsvpsenderadspecctrlloadpathbw", rsvpsenderentry.Rsvpsenderadspecctrlloadpathbw}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadMinLatency"] = types.YLeaf{"Rsvpsenderadspecctrlloadminlatency", rsvpsenderentry.Rsvpsenderadspecctrlloadminlatency}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderAdspecCtrlLoadMtu"] = types.YLeaf{"Rsvpsenderadspecctrlloadmtu", rsvpsenderentry.Rsvpsenderadspecctrlloadmtu}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderStatus"] = types.YLeaf{"Rsvpsenderstatus", rsvpsenderentry.Rsvpsenderstatus}
    rsvpsenderentry.EntityData.Leafs["rsvpSenderTTL"] = types.YLeaf{"Rsvpsenderttl", rsvpsenderentry.Rsvpsenderttl}
    return &(rsvpsenderentry.EntityData)
}

// RSVPMIB_Rsvpsenderoutinterfacetable
// List of outgoing interfaces	that PATH messages
// use.	 The  ifIndex  is the ifIndex value of the
// egress interface.
type RSVPMIB_Rsvpsenderoutinterfacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of outgoing interfaces	that a	particular PATH	message	has. The type is
    // slice of RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry.
    Rsvpsenderoutinterfaceentry []RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetEntityData() *types.CommonEntityData {
    rsvpsenderoutinterfacetable.EntityData.YFilter = rsvpsenderoutinterfacetable.YFilter
    rsvpsenderoutinterfacetable.EntityData.YangName = "rsvpSenderOutInterfaceTable"
    rsvpsenderoutinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    rsvpsenderoutinterfacetable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpsenderoutinterfacetable.EntityData.SegmentPath = "rsvpSenderOutInterfaceTable"
    rsvpsenderoutinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsenderoutinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsenderoutinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsenderoutinterfacetable.EntityData.Children = make(map[string]types.YChild)
    rsvpsenderoutinterfacetable.EntityData.Children["rsvpSenderOutInterfaceEntry"] = types.YChild{"Rsvpsenderoutinterfaceentry", nil}
    for i := range rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry {
        rsvpsenderoutinterfacetable.EntityData.Children[types.GetSegmentPath(&rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry[i])] = types.YChild{"Rsvpsenderoutinterfaceentry", &rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry[i]}
    }
    rsvpsenderoutinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpsenderoutinterfacetable.EntityData)
}

// RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry
// List of outgoing interfaces	that a	particular
// PATH	message	has.
type RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_Rsvpsessiontable_Rsvpsessionentry_Rsvpsessionnumber
    Rsvpsessionnumber interface{}

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to rsvp_mib.RSVPMIB_Rsvpsendertable_Rsvpsenderentry_Rsvpsendernumber
    Rsvpsendernumber interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // 'active' for all active PATH messages. The type is RowStatus.
    Rsvpsenderoutinterfacestatus interface{}
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetEntityData() *types.CommonEntityData {
    rsvpsenderoutinterfaceentry.EntityData.YFilter = rsvpsenderoutinterfaceentry.YFilter
    rsvpsenderoutinterfaceentry.EntityData.YangName = "rsvpSenderOutInterfaceEntry"
    rsvpsenderoutinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpsenderoutinterfaceentry.EntityData.ParentYangName = "rsvpSenderOutInterfaceTable"
    rsvpsenderoutinterfaceentry.EntityData.SegmentPath = "rsvpSenderOutInterfaceEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Rsvpsessionnumber) + "']" + "[rsvpSenderNumber='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Rsvpsendernumber) + "']" + "[ifIndex='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Ifindex) + "']"
    rsvpsenderoutinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpsenderoutinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpsenderoutinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpsenderoutinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    rsvpsenderoutinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpsenderoutinterfaceentry.EntityData.Leafs["rsvpSessionNumber"] = types.YLeaf{"Rsvpsessionnumber", rsvpsenderoutinterfaceentry.Rsvpsessionnumber}
    rsvpsenderoutinterfaceentry.EntityData.Leafs["rsvpSenderNumber"] = types.YLeaf{"Rsvpsendernumber", rsvpsenderoutinterfaceentry.Rsvpsendernumber}
    rsvpsenderoutinterfaceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", rsvpsenderoutinterfaceentry.Ifindex}
    rsvpsenderoutinterfaceentry.EntityData.Leafs["rsvpSenderOutInterfaceStatus"] = types.YLeaf{"Rsvpsenderoutinterfacestatus", rsvpsenderoutinterfaceentry.Rsvpsenderoutinterfacestatus}
    return &(rsvpsenderoutinterfaceentry.EntityData)
}

// RSVPMIB_Rsvpresvtable
// Information	describing the	state  information
// displayed by	receivers in RESV messages.
type RSVPMIB_Rsvpresvtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed  by  a single
    // receiver's RESV message concerning a	single sender. The type is slice of
    // RSVPMIB_Rsvpresvtable_Rsvpresventry.
    Rsvpresventry []RSVPMIB_Rsvpresvtable_Rsvpresventry
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetEntityData() *types.CommonEntityData {
    rsvpresvtable.EntityData.YFilter = rsvpresvtable.YFilter
    rsvpresvtable.EntityData.YangName = "rsvpResvTable"
    rsvpresvtable.EntityData.BundleName = "cisco_ios_xe"
    rsvpresvtable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpresvtable.EntityData.SegmentPath = "rsvpResvTable"
    rsvpresvtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpresvtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpresvtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpresvtable.EntityData.Children = make(map[string]types.YChild)
    rsvpresvtable.EntityData.Children["rsvpResvEntry"] = types.YChild{"Rsvpresventry", nil}
    for i := range rsvpresvtable.Rsvpresventry {
        rsvpresvtable.EntityData.Children[types.GetSegmentPath(&rsvpresvtable.Rsvpresventry[i])] = types.YChild{"Rsvpresventry", &rsvpresvtable.Rsvpresventry[i]}
    }
    rsvpresvtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpresvtable.EntityData)
}

// RSVPMIB_Rsvpresvtable_Rsvpresventry
// Information	describing the	state  information
// displayed  by  a single receiver's RESV message
// concerning a	single sender.
type RSVPMIB_Rsvpresvtable_Rsvpresventry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_Rsvpsessiontable_Rsvpsessionentry_Rsvpsessionnumber
    Rsvpsessionnumber interface{}

    // This attribute is a key. The	number of this reservation request.   This is 
    // for  SNMP Indexing purposes only	and has	no relation to any protocol value.
    // The type is interface{} with range: 0..2147483647.
    Rsvpresvnumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    Rsvpresvtype interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    Rsvpresvdestaddr interface{}

    // The	source address of the sender  selected	by this	 reservation.	The  value
    // of	all zeroes indicates 'all senders'.  This object  may  not be 
    // changed	when  the  value  of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    Rsvpresvsenderaddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpresvdestaddrlength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpresvsenderaddrlength interface{}

    // The	IP Protocol used by  this  session.   This object may not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    Rsvpresvprotocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by 
    // rsvpResvProtocol,  is  50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    Rsvpresvdestport interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpResvProtocol  is	 50 (ESP) or 51	(AH), this represents a	generalized
    // port identifier (GPI). A  value of zero indicates that the IP protocol in
    // use does not have	ports.	 This  object  may not	be changed when	the value
    // of the RowStatus object is 'active'. The type is string with length: 2..4.
    Rsvpresvport interface{}

    // The	address	used by	the next RSVP  hop  (which may be the ultimate
    // receiver). The type is string with length: 4..16.
    Rsvpresvhopaddr interface{}

    // The	Logical	Interface Handle received from the previous  RSVP  hop	(which
    // may be the ultimate receiver). The type is interface{} with range:
    // -2147483648..2147483647.
    Rsvpresvhoplih interface{}

    // The	ifIndex	value of the  interface	 on  which this	RESV message was most
    // recently received. The type is interface{} with range: 1..2147483647.
    Rsvpresvinterface interface{}

    // The	QoS Service  classification  requested	by the receiver. The type is
    // QosService.
    Rsvpresvservice interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpResvTSpecPeakRate   (if	supported  by  the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpResvTSpecRate.  Note	that this is a prediction, often based	on
    // the	general	 capability  of	a type of codec	or particular encoding;	the
    // measured average  rate may be significantly	lower. The type is interface{}
    // with range: 0..2147483647. Units are bits per second.
    Rsvpresvtspecrate interface{}

    // The	Peak Bit Rate of the sender's data stream. Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    Rsvpresvtspecpeakrate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time.  If
    // this is less than	 the  sender's	advertised burst  size,	the receiver
    // is	asking the network to provide flow pacing  beyond  what	 would	be
    // provided   under   normal  circumstances.  Such pacing is at	the network's
    // option. The type is interface{} with range: 0..2147483647. Units are bytes.
    Rsvpresvtspecburst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    Rsvpresvtspecmintu interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    Rsvpresvtspecmaxtu interface{}

    // If the requested  service  is  Guaranteed,	as specified   by 
    // rsvpResvService,  this  is  the clearing  rate   that   is	being	requested.
    // Otherwise,  it is zero, or the agent	may return noSuchValue. The type is
    // interface{} with range: 0..2147483647. Units are bits per second.
    Rsvpresvrspecrate interface{}

    // If the requested  service  is  Guaranteed,	as specified by	rsvpResvService,
    // this is the delay slack.  Otherwise, it is zero, or the agent may return
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    Rsvpresvrspecslack interface{}

    // The	 interval  between  refresh  messages	as advertised by the Next Hop.
    // The type is interface{} with range: 0..2147483647.
    Rsvpresvinterval interface{}

    // The	contents of the	scope object, displayed	as an  uninterpreted  string 
    // of octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length.  If the length  is  non-zero,	 this 
    // contains	 a series of IP4 or IP6	addresses. The type is string with length:
    // 0..65536.
    Rsvpresvscope interface{}

    // If TRUE, a reservation shared among	senders	is requested.  If FALSE, a
    // reservation specific	to this	sender is requested. The type is bool.
    Rsvpresvshared interface{}

    // If TRUE, individual	senders	are  listed  using Filter  Specifications.  
    // If	FALSE, all senders are implicitly selected.  The Scope Object will
    // contain  a list of senders that need	to receive this	reservation request 
    // for  the  purpose	of routing the RESV message. The type is bool.
    Rsvpresvexplicit interface{}

    // If TRUE, the node believes that  the  previous IP  hop  is	an  RSVP 
    // hop.	If FALSE, the node believes that the previous IP hop may not be	an
    // RSVP	hop. The type is bool.
    Rsvpresvrsvphop interface{}

    // The	 time  of  the	 last	change	 in   this reservation	request;  This is
    // either the first time	it was received	or the time  of	 the  most recent
    // change in parameters. The type is interface{} with range: 0..4294967295.
    Rsvpresvlastchange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    Rsvpresvpolicy interface{}

    // 'active' for all active RESV  messages.   This object  may	be  used  to 
    // install  static RESV information or delete RESV information. The type is
    // RowStatus.
    Rsvpresvstatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    Rsvpresvttl interface{}

    // The	flow ID	that this receiver  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    Rsvpresvflowid interface{}
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetEntityData() *types.CommonEntityData {
    rsvpresventry.EntityData.YFilter = rsvpresventry.YFilter
    rsvpresventry.EntityData.YangName = "rsvpResvEntry"
    rsvpresventry.EntityData.BundleName = "cisco_ios_xe"
    rsvpresventry.EntityData.ParentYangName = "rsvpResvTable"
    rsvpresventry.EntityData.SegmentPath = "rsvpResvEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpresventry.Rsvpsessionnumber) + "']" + "[rsvpResvNumber='" + fmt.Sprintf("%v", rsvpresventry.Rsvpresvnumber) + "']"
    rsvpresventry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpresventry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpresventry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpresventry.EntityData.Children = make(map[string]types.YChild)
    rsvpresventry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpresventry.EntityData.Leafs["rsvpSessionNumber"] = types.YLeaf{"Rsvpsessionnumber", rsvpresventry.Rsvpsessionnumber}
    rsvpresventry.EntityData.Leafs["rsvpResvNumber"] = types.YLeaf{"Rsvpresvnumber", rsvpresventry.Rsvpresvnumber}
    rsvpresventry.EntityData.Leafs["rsvpResvType"] = types.YLeaf{"Rsvpresvtype", rsvpresventry.Rsvpresvtype}
    rsvpresventry.EntityData.Leafs["rsvpResvDestAddr"] = types.YLeaf{"Rsvpresvdestaddr", rsvpresventry.Rsvpresvdestaddr}
    rsvpresventry.EntityData.Leafs["rsvpResvSenderAddr"] = types.YLeaf{"Rsvpresvsenderaddr", rsvpresventry.Rsvpresvsenderaddr}
    rsvpresventry.EntityData.Leafs["rsvpResvDestAddrLength"] = types.YLeaf{"Rsvpresvdestaddrlength", rsvpresventry.Rsvpresvdestaddrlength}
    rsvpresventry.EntityData.Leafs["rsvpResvSenderAddrLength"] = types.YLeaf{"Rsvpresvsenderaddrlength", rsvpresventry.Rsvpresvsenderaddrlength}
    rsvpresventry.EntityData.Leafs["rsvpResvProtocol"] = types.YLeaf{"Rsvpresvprotocol", rsvpresventry.Rsvpresvprotocol}
    rsvpresventry.EntityData.Leafs["rsvpResvDestPort"] = types.YLeaf{"Rsvpresvdestport", rsvpresventry.Rsvpresvdestport}
    rsvpresventry.EntityData.Leafs["rsvpResvPort"] = types.YLeaf{"Rsvpresvport", rsvpresventry.Rsvpresvport}
    rsvpresventry.EntityData.Leafs["rsvpResvHopAddr"] = types.YLeaf{"Rsvpresvhopaddr", rsvpresventry.Rsvpresvhopaddr}
    rsvpresventry.EntityData.Leafs["rsvpResvHopLih"] = types.YLeaf{"Rsvpresvhoplih", rsvpresventry.Rsvpresvhoplih}
    rsvpresventry.EntityData.Leafs["rsvpResvInterface"] = types.YLeaf{"Rsvpresvinterface", rsvpresventry.Rsvpresvinterface}
    rsvpresventry.EntityData.Leafs["rsvpResvService"] = types.YLeaf{"Rsvpresvservice", rsvpresventry.Rsvpresvservice}
    rsvpresventry.EntityData.Leafs["rsvpResvTSpecRate"] = types.YLeaf{"Rsvpresvtspecrate", rsvpresventry.Rsvpresvtspecrate}
    rsvpresventry.EntityData.Leafs["rsvpResvTSpecPeakRate"] = types.YLeaf{"Rsvpresvtspecpeakrate", rsvpresventry.Rsvpresvtspecpeakrate}
    rsvpresventry.EntityData.Leafs["rsvpResvTSpecBurst"] = types.YLeaf{"Rsvpresvtspecburst", rsvpresventry.Rsvpresvtspecburst}
    rsvpresventry.EntityData.Leafs["rsvpResvTSpecMinTU"] = types.YLeaf{"Rsvpresvtspecmintu", rsvpresventry.Rsvpresvtspecmintu}
    rsvpresventry.EntityData.Leafs["rsvpResvTSpecMaxTU"] = types.YLeaf{"Rsvpresvtspecmaxtu", rsvpresventry.Rsvpresvtspecmaxtu}
    rsvpresventry.EntityData.Leafs["rsvpResvRSpecRate"] = types.YLeaf{"Rsvpresvrspecrate", rsvpresventry.Rsvpresvrspecrate}
    rsvpresventry.EntityData.Leafs["rsvpResvRSpecSlack"] = types.YLeaf{"Rsvpresvrspecslack", rsvpresventry.Rsvpresvrspecslack}
    rsvpresventry.EntityData.Leafs["rsvpResvInterval"] = types.YLeaf{"Rsvpresvinterval", rsvpresventry.Rsvpresvinterval}
    rsvpresventry.EntityData.Leafs["rsvpResvScope"] = types.YLeaf{"Rsvpresvscope", rsvpresventry.Rsvpresvscope}
    rsvpresventry.EntityData.Leafs["rsvpResvShared"] = types.YLeaf{"Rsvpresvshared", rsvpresventry.Rsvpresvshared}
    rsvpresventry.EntityData.Leafs["rsvpResvExplicit"] = types.YLeaf{"Rsvpresvexplicit", rsvpresventry.Rsvpresvexplicit}
    rsvpresventry.EntityData.Leafs["rsvpResvRSVPHop"] = types.YLeaf{"Rsvpresvrsvphop", rsvpresventry.Rsvpresvrsvphop}
    rsvpresventry.EntityData.Leafs["rsvpResvLastChange"] = types.YLeaf{"Rsvpresvlastchange", rsvpresventry.Rsvpresvlastchange}
    rsvpresventry.EntityData.Leafs["rsvpResvPolicy"] = types.YLeaf{"Rsvpresvpolicy", rsvpresventry.Rsvpresvpolicy}
    rsvpresventry.EntityData.Leafs["rsvpResvStatus"] = types.YLeaf{"Rsvpresvstatus", rsvpresventry.Rsvpresvstatus}
    rsvpresventry.EntityData.Leafs["rsvpResvTTL"] = types.YLeaf{"Rsvpresvttl", rsvpresventry.Rsvpresvttl}
    rsvpresventry.EntityData.Leafs["rsvpResvFlowId"] = types.YLeaf{"Rsvpresvflowid", rsvpresventry.Rsvpresvflowid}
    return &(rsvpresventry.EntityData)
}

// RSVPMIB_Rsvpresvfwdtable
// Information	describing the	state  information
// displayed upstream in RESV messages.
type RSVPMIB_Rsvpresvfwdtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed   upstream	  in  
    // an   RESV   message concerning a	single sender. The type is slice of
    // RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry.
    Rsvpresvfwdentry []RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetEntityData() *types.CommonEntityData {
    rsvpresvfwdtable.EntityData.YFilter = rsvpresvfwdtable.YFilter
    rsvpresvfwdtable.EntityData.YangName = "rsvpResvFwdTable"
    rsvpresvfwdtable.EntityData.BundleName = "cisco_ios_xe"
    rsvpresvfwdtable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpresvfwdtable.EntityData.SegmentPath = "rsvpResvFwdTable"
    rsvpresvfwdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpresvfwdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpresvfwdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpresvfwdtable.EntityData.Children = make(map[string]types.YChild)
    rsvpresvfwdtable.EntityData.Children["rsvpResvFwdEntry"] = types.YChild{"Rsvpresvfwdentry", nil}
    for i := range rsvpresvfwdtable.Rsvpresvfwdentry {
        rsvpresvfwdtable.EntityData.Children[types.GetSegmentPath(&rsvpresvfwdtable.Rsvpresvfwdentry[i])] = types.YChild{"Rsvpresvfwdentry", &rsvpresvfwdtable.Rsvpresvfwdentry[i]}
    }
    rsvpresvfwdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpresvfwdtable.EntityData)
}

// RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry
// Information	describing the	state  information
// displayed   upstream	  in   an   RESV   message
// concerning a	single sender.
type RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // rsvp_mib.RSVPMIB_Rsvpsessiontable_Rsvpsessionentry_Rsvpsessionnumber
    Rsvpsessionnumber interface{}

    // This attribute is a key. The	number of this reservation request.   This is 
    // for  SNMP Indexing purposes only	and has	no relation to any protocol value.
    // The type is interface{} with range: 0..2147483647.
    Rsvpresvfwdnumber interface{}

    // The	type of	session	(IP4, IP6, IP6	with  flow information,	etc). The type
    // is interface{} with range: 1..255.
    Rsvpresvfwdtype interface{}

    // The	destination address used by all	senders	in this	 session.   This
    // object	may not	be changed when	the  value  of	the  RowStatus	object	is
    // 'active'. The type is string with length: 4..16.
    Rsvpresvfwddestaddr interface{}

    // The	source address of the sender  selected	by this	 reservation.	The  value
    // of	all zeroes indicates 'all senders'.  This object  may  not be 
    // changed	when  the  value  of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    Rsvpresvfwdsenderaddr interface{}

    // The	length of the destination address in bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpresvfwddestaddrlength interface{}

    // The	length of the sender's	address	 in  bits. This	 is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Rsvpresvfwdsenderaddrlength interface{}

    // The	IP Protocol used by a session. for  secure sessions,  this  indicates 
    // IP  Security.  This object may not be changed when the value of the
    // RowStatus object is 'active'. The type is interface{} with range: 1..255.
    Rsvpresvfwdprotocol interface{}

    // The	 UDP  or  TCP  port  number  used   as	 a destination	 port  for  all	
    // senders  in  this session.  If	the IP protocol	in use,	 specified by
    // rsvpResvFwdProtocol, is 50 (ESP) or 51 (AH), this	 represents  a	virtual	
    // destination  port number.   A value of	zero indicates that the	IP protocol
    // in use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    Rsvpresvfwddestport interface{}

    // The	UDP or TCP port	number used  as	 a  source port	 for  this sender in
    // this session.  If the IP	 protocol    in	   use,	   specified	by
    // rsvpResvFwdProtocol	is  50	(ESP)  or 51 (AH), this	represents a generalized
    // port	identifier (GPI).   A  value of	zero indicates that the	IP protocol in
    // use  does  not  have  ports.   This object may not be changed when the
    // value of the RowStatus object is 'active'. The type is string with length:
    // 2..4.
    Rsvpresvfwdport interface{}

    // The	address	of the (previous) RSVP	that  will receive this	message. The
    // type is string with length: 4..16.
    Rsvpresvfwdhopaddr interface{}

    // The	 Logical  Interface  Handle  sent  to  the (previous)	RSVP   that  
    // will   receive  this message. The type is interface{} with range:
    // -2147483648..2147483647.
    Rsvpresvfwdhoplih interface{}

    // The	ifIndex	value of the  interface	 on  which this	RESV message was most
    // recently sent. The type is interface{} with range: 1..2147483647.
    Rsvpresvfwdinterface interface{}

    // The	QoS Service classification requested. The type is QosService.
    Rsvpresvfwdservice interface{}

    // The	Average	Bit  Rate  of  the  sender's  data stream.    Within  a	
    // transmission  burst,  the arrival   rate    may    be	  as	fast	as
    // rsvpResvFwdTSpecPeakRate  (if  supported by the service model); however,
    // averaged across two	or more	 burst	intervals,  the	 rate  should  not
    // exceed rsvpResvFwdTSpecRate.  Note	that this is a prediction, often
    // based	on the	general	 capability  of	a type of codec	or particular
    // encoding;	the measured average  rate may be significantly	lower. The type
    // is interface{} with range: 0..2147483647. Units are bits per second.
    Rsvpresvfwdtspecrate interface{}

    // The	Peak Bit Rate of the sender's data  stream Traffic  arrival is not
    // expected to exceed this rate	at any time, apart  from  the  effects	of
    // jitter in the network.  If not specified in the TSpec, this returns zero or
    // noSuchValue. The type is interface{} with range: 0..2147483647. Units are
    // bits per second.
    Rsvpresvfwdtspecpeakrate interface{}

    // The	size of	the largest  burst  expected  from the sender at a time.  If
    // this is less than	 the  sender's	advertised burst  size,	the receiver
    // is	asking the network to provide flow pacing  beyond  what	 would	be
    // provided   under   normal  circumstances.  Such pacing is at	the network's
    // option. The type is interface{} with range: 0..2147483647. Units are bytes.
    Rsvpresvfwdtspecburst interface{}

    // The	minimum	message	size for  this	flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    Rsvpresvfwdtspecmintu interface{}

    // The	maximum	message	size for  this	flow.  The admission  algorithm	 will 
    // reject TSpecs whose Maximum Transmission	Unit, plus  the	 interface
    // headers, exceed the interface MTU. The type is interface{} with range:
    // 0..2147483647.
    Rsvpresvfwdtspecmaxtu interface{}

    // If the requested  service  is  Guaranteed,	as specified   by 
    // rsvpResvService,  this  is  the clearing  rate   that   is	being	requested.
    // Otherwise,  it is zero, or the agent	may return noSuchValue. The type is
    // interface{} with range: 0..2147483647. Units are bytes per second.
    Rsvpresvfwdrspecrate interface{}

    // If the requested  service  is  Guaranteed,	as specified by	rsvpResvService,
    // this is the delay slack.  Otherwise, it is zero, or the agent may return
    // noSuchValue. The type is interface{} with range: -2147483648..2147483647.
    // Units are microseconds.
    Rsvpresvfwdrspecslack interface{}

    // The	  interval   between   refresh	  messages advertised to the Previous
    // Hop. The type is interface{} with range: 0..2147483647.
    Rsvpresvfwdinterval interface{}

    // The	contents of the	scope object, displayed	as an  uninterpreted  string 
    // of octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    Rsvpresvfwdscope interface{}

    // If TRUE, a reservation shared among	senders	is requested.  If FALSE, a
    // reservation specific	to this	sender is requested. The type is bool.
    Rsvpresvfwdshared interface{}

    // If TRUE, individual	senders	are  listed  using Filter  Specifications.  
    // If	FALSE, all senders are implicitly selected.  The Scope Object will
    // contain  a list of senders that need	to receive this	reservation request 
    // for  the  purpose	of routing the RESV message. The type is bool.
    Rsvpresvfwdexplicit interface{}

    // If TRUE, the node believes that  the  next	IP hop	is  an	RSVP  hop.   If	
    // FALSE,	 the  node believes that the next IP hop  may  not  be	an RSVP	hop.
    // The type is bool.
    Rsvpresvfwdrsvphop interface{}

    // The	time of	the last change	in  this  request; This	 is  either  the first
    // time it was sent	or the	time  of  the  most   recent   change	in
    // parameters. The type is interface{} with range: 0..4294967295.
    Rsvpresvfwdlastchange interface{}

    // The	contents of the	policy	object,	 displayed as an uninterpreted string of
    // octets, including the object header.  In the absence of  such	an object,
    // this	should be of zero length. The type is string with length: 0..65536.
    Rsvpresvfwdpolicy interface{}

    // 'active' for all active RESV  messages.   This object may be used to
    // delete	RESV information. The type is RowStatus.
    Rsvpresvfwdstatus interface{}

    // The	TTL value in the RSVP header that was last received. The type is
    // interface{} with range: 0..255.
    Rsvpresvfwdttl interface{}

    // The	flow ID	that this receiver  is	using,	if this	 is  an	IPv6 session. The
    // type is interface{} with range: 0..16777215.
    Rsvpresvfwdflowid interface{}
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetEntityData() *types.CommonEntityData {
    rsvpresvfwdentry.EntityData.YFilter = rsvpresvfwdentry.YFilter
    rsvpresvfwdentry.EntityData.YangName = "rsvpResvFwdEntry"
    rsvpresvfwdentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpresvfwdentry.EntityData.ParentYangName = "rsvpResvFwdTable"
    rsvpresvfwdentry.EntityData.SegmentPath = "rsvpResvFwdEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpresvfwdentry.Rsvpsessionnumber) + "']" + "[rsvpResvFwdNumber='" + fmt.Sprintf("%v", rsvpresvfwdentry.Rsvpresvfwdnumber) + "']"
    rsvpresvfwdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpresvfwdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpresvfwdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpresvfwdentry.EntityData.Children = make(map[string]types.YChild)
    rsvpresvfwdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpresvfwdentry.EntityData.Leafs["rsvpSessionNumber"] = types.YLeaf{"Rsvpsessionnumber", rsvpresvfwdentry.Rsvpsessionnumber}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdNumber"] = types.YLeaf{"Rsvpresvfwdnumber", rsvpresvfwdentry.Rsvpresvfwdnumber}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdType"] = types.YLeaf{"Rsvpresvfwdtype", rsvpresvfwdentry.Rsvpresvfwdtype}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdDestAddr"] = types.YLeaf{"Rsvpresvfwddestaddr", rsvpresvfwdentry.Rsvpresvfwddestaddr}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdSenderAddr"] = types.YLeaf{"Rsvpresvfwdsenderaddr", rsvpresvfwdentry.Rsvpresvfwdsenderaddr}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdDestAddrLength"] = types.YLeaf{"Rsvpresvfwddestaddrlength", rsvpresvfwdentry.Rsvpresvfwddestaddrlength}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdSenderAddrLength"] = types.YLeaf{"Rsvpresvfwdsenderaddrlength", rsvpresvfwdentry.Rsvpresvfwdsenderaddrlength}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdProtocol"] = types.YLeaf{"Rsvpresvfwdprotocol", rsvpresvfwdentry.Rsvpresvfwdprotocol}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdDestPort"] = types.YLeaf{"Rsvpresvfwddestport", rsvpresvfwdentry.Rsvpresvfwddestport}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdPort"] = types.YLeaf{"Rsvpresvfwdport", rsvpresvfwdentry.Rsvpresvfwdport}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdHopAddr"] = types.YLeaf{"Rsvpresvfwdhopaddr", rsvpresvfwdentry.Rsvpresvfwdhopaddr}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdHopLih"] = types.YLeaf{"Rsvpresvfwdhoplih", rsvpresvfwdentry.Rsvpresvfwdhoplih}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdInterface"] = types.YLeaf{"Rsvpresvfwdinterface", rsvpresvfwdentry.Rsvpresvfwdinterface}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdService"] = types.YLeaf{"Rsvpresvfwdservice", rsvpresvfwdentry.Rsvpresvfwdservice}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTSpecRate"] = types.YLeaf{"Rsvpresvfwdtspecrate", rsvpresvfwdentry.Rsvpresvfwdtspecrate}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTSpecPeakRate"] = types.YLeaf{"Rsvpresvfwdtspecpeakrate", rsvpresvfwdentry.Rsvpresvfwdtspecpeakrate}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTSpecBurst"] = types.YLeaf{"Rsvpresvfwdtspecburst", rsvpresvfwdentry.Rsvpresvfwdtspecburst}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTSpecMinTU"] = types.YLeaf{"Rsvpresvfwdtspecmintu", rsvpresvfwdentry.Rsvpresvfwdtspecmintu}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTSpecMaxTU"] = types.YLeaf{"Rsvpresvfwdtspecmaxtu", rsvpresvfwdentry.Rsvpresvfwdtspecmaxtu}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdRSpecRate"] = types.YLeaf{"Rsvpresvfwdrspecrate", rsvpresvfwdentry.Rsvpresvfwdrspecrate}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdRSpecSlack"] = types.YLeaf{"Rsvpresvfwdrspecslack", rsvpresvfwdentry.Rsvpresvfwdrspecslack}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdInterval"] = types.YLeaf{"Rsvpresvfwdinterval", rsvpresvfwdentry.Rsvpresvfwdinterval}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdScope"] = types.YLeaf{"Rsvpresvfwdscope", rsvpresvfwdentry.Rsvpresvfwdscope}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdShared"] = types.YLeaf{"Rsvpresvfwdshared", rsvpresvfwdentry.Rsvpresvfwdshared}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdExplicit"] = types.YLeaf{"Rsvpresvfwdexplicit", rsvpresvfwdentry.Rsvpresvfwdexplicit}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdRSVPHop"] = types.YLeaf{"Rsvpresvfwdrsvphop", rsvpresvfwdentry.Rsvpresvfwdrsvphop}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdLastChange"] = types.YLeaf{"Rsvpresvfwdlastchange", rsvpresvfwdentry.Rsvpresvfwdlastchange}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdPolicy"] = types.YLeaf{"Rsvpresvfwdpolicy", rsvpresvfwdentry.Rsvpresvfwdpolicy}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdStatus"] = types.YLeaf{"Rsvpresvfwdstatus", rsvpresvfwdentry.Rsvpresvfwdstatus}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdTTL"] = types.YLeaf{"Rsvpresvfwdttl", rsvpresvfwdentry.Rsvpresvfwdttl}
    rsvpresvfwdentry.EntityData.Leafs["rsvpResvFwdFlowId"] = types.YLeaf{"Rsvpresvfwdflowid", rsvpresvfwdentry.Rsvpresvfwdflowid}
    return &(rsvpresvfwdentry.EntityData)
}

// RSVPMIB_Rsvpiftable
// The	RSVP-specific attributes of  the  system's
// interfaces.
type RSVPMIB_Rsvpiftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The	RSVP-specific attributes of  the  a  given interface. The type is slice
    // of RSVPMIB_Rsvpiftable_Rsvpifentry.
    Rsvpifentry []RSVPMIB_Rsvpiftable_Rsvpifentry
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetEntityData() *types.CommonEntityData {
    rsvpiftable.EntityData.YFilter = rsvpiftable.YFilter
    rsvpiftable.EntityData.YangName = "rsvpIfTable"
    rsvpiftable.EntityData.BundleName = "cisco_ios_xe"
    rsvpiftable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpiftable.EntityData.SegmentPath = "rsvpIfTable"
    rsvpiftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpiftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpiftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpiftable.EntityData.Children = make(map[string]types.YChild)
    rsvpiftable.EntityData.Children["rsvpIfEntry"] = types.YChild{"Rsvpifentry", nil}
    for i := range rsvpiftable.Rsvpifentry {
        rsvpiftable.EntityData.Children[types.GetSegmentPath(&rsvpiftable.Rsvpifentry[i])] = types.YChild{"Rsvpifentry", &rsvpiftable.Rsvpifentry[i]}
    }
    rsvpiftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpiftable.EntityData)
}

// RSVPMIB_Rsvpiftable_Rsvpifentry
// The	RSVP-specific attributes of  the  a  given
// interface.
type RSVPMIB_Rsvpiftable_Rsvpifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The	number of neighbors perceived to be  using only	the RSVP UDP
    // Encapsulation. The type is interface{} with range: 0..4294967295.
    Rsvpifudpnbrs interface{}

    // The	number of neighbors perceived to be  using only	the RSVP IP
    // Encapsulation. The type is interface{} with range: 0..4294967295.
    Rsvpifipnbrs interface{}

    // The	number of neighbors  currently	perceived; this	 will  exceed
    // rsvpIfIpNbrs + rsvpIfUdpNbrs by  the  number   of	  neighbors   using  
    // both encapsulations. The type is interface{} with range: 0..4294967295.
    Rsvpifnbrs interface{}

    // The	value of the RSVP value	'Kb', Which is the minimum   number   of 
    // refresh  intervals  that blockade state will last once entered. The type is
    // interface{} with range: 1..65536.
    Rsvpifrefreshblockademultiple interface{}

    // The	value of the RSVP value	'K', which is  the number  of  refresh
    // intervals which must elapse (minimum) before a PATH or RESV  message  which
    // is not being	refreshed will be aged out. The type is interface{} with
    // range: 1..65536.
    Rsvpifrefreshmultiple interface{}

    // The	value of SEND_TTL used on  this	 interface for	messages  this node
    // originates.	 If set	to zero, the node determines  the  TTL	via  other
    // means. The type is interface{} with range: 0..255.
    Rsvpifttl interface{}

    // The	value of the RSVP value	'R', which is  the minimum period between
    // refresh transmissions	of a given PATH	or RESV	message	on an interface. The
    // type is interface{} with range: 0..2147483647. Units are milliseconds.
    Rsvpifrefreshinterval interface{}

    // The	approximate period from	the time  a  route is  changed	to  the	 time  a
    // resulting message appears on the interface. The type is interface{} with
    // range: 0..2147483647. Units are hundredths	of a second.
    Rsvpifroutedelay interface{}

    // If TRUE, RSVP is enabled  on  this	Interface. If	FALSE,	 RSVP	is  not	
    // enabled  on  this interface. The type is bool.
    Rsvpifenabled interface{}

    // If TRUE, manual configuration forces  the  use of  UDP  encapsulation  on 
    // the  interface.	If FALSE,  UDP	encapsulation  is  only	 used	if
    // rsvpIfUdpNbrs is not	zero. The type is bool.
    Rsvpifudprequired interface{}

    // 'active' on	interfaces that	are configured for RSVP. The type is RowStatus.
    Rsvpifstatus interface{}
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetEntityData() *types.CommonEntityData {
    rsvpifentry.EntityData.YFilter = rsvpifentry.YFilter
    rsvpifentry.EntityData.YangName = "rsvpIfEntry"
    rsvpifentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpifentry.EntityData.ParentYangName = "rsvpIfTable"
    rsvpifentry.EntityData.SegmentPath = "rsvpIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", rsvpifentry.Ifindex) + "']"
    rsvpifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpifentry.EntityData.Children = make(map[string]types.YChild)
    rsvpifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", rsvpifentry.Ifindex}
    rsvpifentry.EntityData.Leafs["rsvpIfUdpNbrs"] = types.YLeaf{"Rsvpifudpnbrs", rsvpifentry.Rsvpifudpnbrs}
    rsvpifentry.EntityData.Leafs["rsvpIfIpNbrs"] = types.YLeaf{"Rsvpifipnbrs", rsvpifentry.Rsvpifipnbrs}
    rsvpifentry.EntityData.Leafs["rsvpIfNbrs"] = types.YLeaf{"Rsvpifnbrs", rsvpifentry.Rsvpifnbrs}
    rsvpifentry.EntityData.Leafs["rsvpIfRefreshBlockadeMultiple"] = types.YLeaf{"Rsvpifrefreshblockademultiple", rsvpifentry.Rsvpifrefreshblockademultiple}
    rsvpifentry.EntityData.Leafs["rsvpIfRefreshMultiple"] = types.YLeaf{"Rsvpifrefreshmultiple", rsvpifentry.Rsvpifrefreshmultiple}
    rsvpifentry.EntityData.Leafs["rsvpIfTTL"] = types.YLeaf{"Rsvpifttl", rsvpifentry.Rsvpifttl}
    rsvpifentry.EntityData.Leafs["rsvpIfRefreshInterval"] = types.YLeaf{"Rsvpifrefreshinterval", rsvpifentry.Rsvpifrefreshinterval}
    rsvpifentry.EntityData.Leafs["rsvpIfRouteDelay"] = types.YLeaf{"Rsvpifroutedelay", rsvpifentry.Rsvpifroutedelay}
    rsvpifentry.EntityData.Leafs["rsvpIfEnabled"] = types.YLeaf{"Rsvpifenabled", rsvpifentry.Rsvpifenabled}
    rsvpifentry.EntityData.Leafs["rsvpIfUdpRequired"] = types.YLeaf{"Rsvpifudprequired", rsvpifentry.Rsvpifudprequired}
    rsvpifentry.EntityData.Leafs["rsvpIfStatus"] = types.YLeaf{"Rsvpifstatus", rsvpifentry.Rsvpifstatus}
    return &(rsvpifentry.EntityData)
}

// RSVPMIB_Rsvpnbrtable
// Information	describing  the	 Neighbors  of	an
// RSVP	system.
type RSVPMIB_Rsvpnbrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information	  describing   a    single    RSVP Neighbor. The type is slice
    // of RSVPMIB_Rsvpnbrtable_Rsvpnbrentry.
    Rsvpnbrentry []RSVPMIB_Rsvpnbrtable_Rsvpnbrentry
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetEntityData() *types.CommonEntityData {
    rsvpnbrtable.EntityData.YFilter = rsvpnbrtable.YFilter
    rsvpnbrtable.EntityData.YangName = "rsvpNbrTable"
    rsvpnbrtable.EntityData.BundleName = "cisco_ios_xe"
    rsvpnbrtable.EntityData.ParentYangName = "RSVP-MIB"
    rsvpnbrtable.EntityData.SegmentPath = "rsvpNbrTable"
    rsvpnbrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpnbrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpnbrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpnbrtable.EntityData.Children = make(map[string]types.YChild)
    rsvpnbrtable.EntityData.Children["rsvpNbrEntry"] = types.YChild{"Rsvpnbrentry", nil}
    for i := range rsvpnbrtable.Rsvpnbrentry {
        rsvpnbrtable.EntityData.Children[types.GetSegmentPath(&rsvpnbrtable.Rsvpnbrentry[i])] = types.YChild{"Rsvpnbrentry", &rsvpnbrtable.Rsvpnbrentry[i]}
    }
    rsvpnbrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpnbrtable.EntityData)
}

// RSVPMIB_Rsvpnbrtable_Rsvpnbrentry
// Information	  describing   a    single    RSVP
// Neighbor.
type RSVPMIB_Rsvpnbrtable_Rsvpnbrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The	IP4 or IP6 Address used	by this	 neighbor.
    // This	 object	 may not be changed when the value of the RowStatus object is
    // 'active'. The type is string with length: 4..16.
    Rsvpnbraddress interface{}

    // The	  encapsulation	  being	  used	 by   this neighbor. The type is
    // RsvpEncapsulation.
    Rsvpnbrprotocol interface{}

    // 'active' for all neighbors.	 This  object  may be	used   to  configure 
    // neighbors.   In  the presence   of   configured	 neighbors,    the
    // implementation  may	(but  is  not required to) limit the  set  of  valid 
    // neighbors	 to  those configured. The type is RowStatus.
    Rsvpnbrstatus interface{}
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetEntityData() *types.CommonEntityData {
    rsvpnbrentry.EntityData.YFilter = rsvpnbrentry.YFilter
    rsvpnbrentry.EntityData.YangName = "rsvpNbrEntry"
    rsvpnbrentry.EntityData.BundleName = "cisco_ios_xe"
    rsvpnbrentry.EntityData.ParentYangName = "rsvpNbrTable"
    rsvpnbrentry.EntityData.SegmentPath = "rsvpNbrEntry" + "[ifIndex='" + fmt.Sprintf("%v", rsvpnbrentry.Ifindex) + "']" + "[rsvpNbrAddress='" + fmt.Sprintf("%v", rsvpnbrentry.Rsvpnbraddress) + "']"
    rsvpnbrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rsvpnbrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rsvpnbrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rsvpnbrentry.EntityData.Children = make(map[string]types.YChild)
    rsvpnbrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpnbrentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", rsvpnbrentry.Ifindex}
    rsvpnbrentry.EntityData.Leafs["rsvpNbrAddress"] = types.YLeaf{"Rsvpnbraddress", rsvpnbrentry.Rsvpnbraddress}
    rsvpnbrentry.EntityData.Leafs["rsvpNbrProtocol"] = types.YLeaf{"Rsvpnbrprotocol", rsvpnbrentry.Rsvpnbrprotocol}
    rsvpnbrentry.EntityData.Leafs["rsvpNbrStatus"] = types.YLeaf{"Rsvpnbrstatus", rsvpnbrentry.Rsvpnbrstatus}
    return &(rsvpnbrentry.EntityData)
}

