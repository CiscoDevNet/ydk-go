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
    parent types.Entity
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

func (rSVPMIB *RSVPMIB) GetFilter() yfilter.YFilter { return rSVPMIB.YFilter }

func (rSVPMIB *RSVPMIB) SetFilter(yf yfilter.YFilter) { rSVPMIB.YFilter = yf }

func (rSVPMIB *RSVPMIB) GetGoName(yname string) string {
    if yname == "rsvpGenObjects" { return "Rsvpgenobjects" }
    if yname == "rsvpSessionTable" { return "Rsvpsessiontable" }
    if yname == "rsvpSenderTable" { return "Rsvpsendertable" }
    if yname == "rsvpSenderOutInterfaceTable" { return "Rsvpsenderoutinterfacetable" }
    if yname == "rsvpResvTable" { return "Rsvpresvtable" }
    if yname == "rsvpResvFwdTable" { return "Rsvpresvfwdtable" }
    if yname == "rsvpIfTable" { return "Rsvpiftable" }
    if yname == "rsvpNbrTable" { return "Rsvpnbrtable" }
    return ""
}

func (rSVPMIB *RSVPMIB) GetSegmentPath() string {
    return "RSVP-MIB:RSVP-MIB"
}

func (rSVPMIB *RSVPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpGenObjects" {
        return &rSVPMIB.Rsvpgenobjects
    }
    if childYangName == "rsvpSessionTable" {
        return &rSVPMIB.Rsvpsessiontable
    }
    if childYangName == "rsvpSenderTable" {
        return &rSVPMIB.Rsvpsendertable
    }
    if childYangName == "rsvpSenderOutInterfaceTable" {
        return &rSVPMIB.Rsvpsenderoutinterfacetable
    }
    if childYangName == "rsvpResvTable" {
        return &rSVPMIB.Rsvpresvtable
    }
    if childYangName == "rsvpResvFwdTable" {
        return &rSVPMIB.Rsvpresvfwdtable
    }
    if childYangName == "rsvpIfTable" {
        return &rSVPMIB.Rsvpiftable
    }
    if childYangName == "rsvpNbrTable" {
        return &rSVPMIB.Rsvpnbrtable
    }
    return nil
}

func (rSVPMIB *RSVPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rsvpGenObjects"] = &rSVPMIB.Rsvpgenobjects
    children["rsvpSessionTable"] = &rSVPMIB.Rsvpsessiontable
    children["rsvpSenderTable"] = &rSVPMIB.Rsvpsendertable
    children["rsvpSenderOutInterfaceTable"] = &rSVPMIB.Rsvpsenderoutinterfacetable
    children["rsvpResvTable"] = &rSVPMIB.Rsvpresvtable
    children["rsvpResvFwdTable"] = &rSVPMIB.Rsvpresvfwdtable
    children["rsvpIfTable"] = &rSVPMIB.Rsvpiftable
    children["rsvpNbrTable"] = &rSVPMIB.Rsvpnbrtable
    return children
}

func (rSVPMIB *RSVPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rSVPMIB *RSVPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (rSVPMIB *RSVPMIB) GetYangName() string { return "RSVP-MIB" }

func (rSVPMIB *RSVPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rSVPMIB *RSVPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rSVPMIB *RSVPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rSVPMIB *RSVPMIB) SetParent(parent types.Entity) { rSVPMIB.parent = parent }

func (rSVPMIB *RSVPMIB) GetParent() types.Entity { return rSVPMIB.parent }

func (rSVPMIB *RSVPMIB) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpgenobjects
type RSVPMIB_Rsvpgenobjects struct {
    parent types.Entity
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

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetFilter() yfilter.YFilter { return rsvpgenobjects.YFilter }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) SetFilter(yf yfilter.YFilter) { rsvpgenobjects.YFilter = yf }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetGoName(yname string) string {
    if yname == "rsvpBadPackets" { return "Rsvpbadpackets" }
    if yname == "rsvpSenderNewIndex" { return "Rsvpsendernewindex" }
    if yname == "rsvpResvNewIndex" { return "Rsvpresvnewindex" }
    if yname == "rsvpResvFwdNewIndex" { return "Rsvpresvfwdnewindex" }
    if yname == "rsvpSessionNewIndex" { return "Rsvpsessionnewindex" }
    return ""
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetSegmentPath() string {
    return "rsvpGenObjects"
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpBadPackets"] = rsvpgenobjects.Rsvpbadpackets
    leafs["rsvpSenderNewIndex"] = rsvpgenobjects.Rsvpsendernewindex
    leafs["rsvpResvNewIndex"] = rsvpgenobjects.Rsvpresvnewindex
    leafs["rsvpResvFwdNewIndex"] = rsvpgenobjects.Rsvpresvfwdnewindex
    leafs["rsvpSessionNewIndex"] = rsvpgenobjects.Rsvpsessionnewindex
    return leafs
}

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetYangName() string { return "rsvpGenObjects" }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) SetParent(parent types.Entity) { rsvpgenobjects.parent = parent }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetParent() types.Entity { return rsvpgenobjects.parent }

func (rsvpgenobjects *RSVPMIB_Rsvpgenobjects) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpsessiontable
// A table  of	 all  sessions	seen  by  a  given
// system.
type RSVPMIB_Rsvpsessiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single session seen by a given system. The type is slice of
    // RSVPMIB_Rsvpsessiontable_Rsvpsessionentry.
    Rsvpsessionentry []RSVPMIB_Rsvpsessiontable_Rsvpsessionentry
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetFilter() yfilter.YFilter { return rsvpsessiontable.YFilter }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) SetFilter(yf yfilter.YFilter) { rsvpsessiontable.YFilter = yf }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetGoName(yname string) string {
    if yname == "rsvpSessionEntry" { return "Rsvpsessionentry" }
    return ""
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetSegmentPath() string {
    return "rsvpSessionTable"
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpSessionEntry" {
        for _, c := range rsvpsessiontable.Rsvpsessionentry {
            if rsvpsessiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpsessiontable_Rsvpsessionentry{}
        rsvpsessiontable.Rsvpsessionentry = append(rsvpsessiontable.Rsvpsessionentry, child)
        return &rsvpsessiontable.Rsvpsessionentry[len(rsvpsessiontable.Rsvpsessionentry)-1]
    }
    return nil
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpsessiontable.Rsvpsessionentry {
        children[rsvpsessiontable.Rsvpsessionentry[i].GetSegmentPath()] = &rsvpsessiontable.Rsvpsessionentry[i]
    }
    return children
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetYangName() string { return "rsvpSessionTable" }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) SetParent(parent types.Entity) { rsvpsessiontable.parent = parent }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetParent() types.Entity { return rsvpsessiontable.parent }

func (rsvpsessiontable *RSVPMIB_Rsvpsessiontable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpsessiontable_Rsvpsessionentry
// A single session seen by a given system.
type RSVPMIB_Rsvpsessiontable_Rsvpsessionentry struct {
    parent types.Entity
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

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetFilter() yfilter.YFilter { return rsvpsessionentry.YFilter }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) SetFilter(yf yfilter.YFilter) { rsvpsessionentry.YFilter = yf }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetGoName(yname string) string {
    if yname == "rsvpSessionNumber" { return "Rsvpsessionnumber" }
    if yname == "rsvpSessionType" { return "Rsvpsessiontype" }
    if yname == "rsvpSessionDestAddr" { return "Rsvpsessiondestaddr" }
    if yname == "rsvpSessionDestAddrLength" { return "Rsvpsessiondestaddrlength" }
    if yname == "rsvpSessionProtocol" { return "Rsvpsessionprotocol" }
    if yname == "rsvpSessionPort" { return "Rsvpsessionport" }
    if yname == "rsvpSessionSenders" { return "Rsvpsessionsenders" }
    if yname == "rsvpSessionReceivers" { return "Rsvpsessionreceivers" }
    if yname == "rsvpSessionRequests" { return "Rsvpsessionrequests" }
    return ""
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetSegmentPath() string {
    return "rsvpSessionEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsessionentry.Rsvpsessionnumber) + "']"
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpSessionNumber"] = rsvpsessionentry.Rsvpsessionnumber
    leafs["rsvpSessionType"] = rsvpsessionentry.Rsvpsessiontype
    leafs["rsvpSessionDestAddr"] = rsvpsessionentry.Rsvpsessiondestaddr
    leafs["rsvpSessionDestAddrLength"] = rsvpsessionentry.Rsvpsessiondestaddrlength
    leafs["rsvpSessionProtocol"] = rsvpsessionentry.Rsvpsessionprotocol
    leafs["rsvpSessionPort"] = rsvpsessionentry.Rsvpsessionport
    leafs["rsvpSessionSenders"] = rsvpsessionentry.Rsvpsessionsenders
    leafs["rsvpSessionReceivers"] = rsvpsessionentry.Rsvpsessionreceivers
    leafs["rsvpSessionRequests"] = rsvpsessionentry.Rsvpsessionrequests
    return leafs
}

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetYangName() string { return "rsvpSessionEntry" }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) SetParent(parent types.Entity) { rsvpsessionentry.parent = parent }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetParent() types.Entity { return rsvpsessionentry.parent }

func (rsvpsessionentry *RSVPMIB_Rsvpsessiontable_Rsvpsessionentry) GetParentYangName() string { return "rsvpSessionTable" }

// RSVPMIB_Rsvpsendertable
// Information	describing the	state  information
// displayed by	senders	in PATH	messages.
type RSVPMIB_Rsvpsendertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed by	a single
    // sender's PATH message. The type is slice of
    // RSVPMIB_Rsvpsendertable_Rsvpsenderentry.
    Rsvpsenderentry []RSVPMIB_Rsvpsendertable_Rsvpsenderentry
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetFilter() yfilter.YFilter { return rsvpsendertable.YFilter }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) SetFilter(yf yfilter.YFilter) { rsvpsendertable.YFilter = yf }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetGoName(yname string) string {
    if yname == "rsvpSenderEntry" { return "Rsvpsenderentry" }
    return ""
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetSegmentPath() string {
    return "rsvpSenderTable"
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpSenderEntry" {
        for _, c := range rsvpsendertable.Rsvpsenderentry {
            if rsvpsendertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpsendertable_Rsvpsenderentry{}
        rsvpsendertable.Rsvpsenderentry = append(rsvpsendertable.Rsvpsenderentry, child)
        return &rsvpsendertable.Rsvpsenderentry[len(rsvpsendertable.Rsvpsenderentry)-1]
    }
    return nil
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpsendertable.Rsvpsenderentry {
        children[rsvpsendertable.Rsvpsenderentry[i].GetSegmentPath()] = &rsvpsendertable.Rsvpsenderentry[i]
    }
    return children
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetYangName() string { return "rsvpSenderTable" }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) SetParent(parent types.Entity) { rsvpsendertable.parent = parent }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetParent() types.Entity { return rsvpsendertable.parent }

func (rsvpsendertable *RSVPMIB_Rsvpsendertable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpsendertable_Rsvpsenderentry
// Information	describing the	state  information
// displayed by	a single sender's PATH message.
type RSVPMIB_Rsvpsendertable_Rsvpsenderentry struct {
    parent types.Entity
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

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetFilter() yfilter.YFilter { return rsvpsenderentry.YFilter }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) SetFilter(yf yfilter.YFilter) { rsvpsenderentry.YFilter = yf }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetGoName(yname string) string {
    if yname == "rsvpSessionNumber" { return "Rsvpsessionnumber" }
    if yname == "rsvpSenderNumber" { return "Rsvpsendernumber" }
    if yname == "rsvpSenderType" { return "Rsvpsendertype" }
    if yname == "rsvpSenderDestAddr" { return "Rsvpsenderdestaddr" }
    if yname == "rsvpSenderAddr" { return "Rsvpsenderaddr" }
    if yname == "rsvpSenderDestAddrLength" { return "Rsvpsenderdestaddrlength" }
    if yname == "rsvpSenderAddrLength" { return "Rsvpsenderaddrlength" }
    if yname == "rsvpSenderProtocol" { return "Rsvpsenderprotocol" }
    if yname == "rsvpSenderDestPort" { return "Rsvpsenderdestport" }
    if yname == "rsvpSenderPort" { return "Rsvpsenderport" }
    if yname == "rsvpSenderFlowId" { return "Rsvpsenderflowid" }
    if yname == "rsvpSenderHopAddr" { return "Rsvpsenderhopaddr" }
    if yname == "rsvpSenderHopLih" { return "Rsvpsenderhoplih" }
    if yname == "rsvpSenderInterface" { return "Rsvpsenderinterface" }
    if yname == "rsvpSenderTSpecRate" { return "Rsvpsendertspecrate" }
    if yname == "rsvpSenderTSpecPeakRate" { return "Rsvpsendertspecpeakrate" }
    if yname == "rsvpSenderTSpecBurst" { return "Rsvpsendertspecburst" }
    if yname == "rsvpSenderTSpecMinTU" { return "Rsvpsendertspecmintu" }
    if yname == "rsvpSenderTSpecMaxTU" { return "Rsvpsendertspecmaxtu" }
    if yname == "rsvpSenderInterval" { return "Rsvpsenderinterval" }
    if yname == "rsvpSenderRSVPHop" { return "Rsvpsenderrsvphop" }
    if yname == "rsvpSenderLastChange" { return "Rsvpsenderlastchange" }
    if yname == "rsvpSenderPolicy" { return "Rsvpsenderpolicy" }
    if yname == "rsvpSenderAdspecBreak" { return "Rsvpsenderadspecbreak" }
    if yname == "rsvpSenderAdspecHopCount" { return "Rsvpsenderadspechopcount" }
    if yname == "rsvpSenderAdspecPathBw" { return "Rsvpsenderadspecpathbw" }
    if yname == "rsvpSenderAdspecMinLatency" { return "Rsvpsenderadspecminlatency" }
    if yname == "rsvpSenderAdspecMtu" { return "Rsvpsenderadspecmtu" }
    if yname == "rsvpSenderAdspecGuaranteedSvc" { return "Rsvpsenderadspecguaranteedsvc" }
    if yname == "rsvpSenderAdspecGuaranteedBreak" { return "Rsvpsenderadspecguaranteedbreak" }
    if yname == "rsvpSenderAdspecGuaranteedCtot" { return "Rsvpsenderadspecguaranteedctot" }
    if yname == "rsvpSenderAdspecGuaranteedDtot" { return "Rsvpsenderadspecguaranteeddtot" }
    if yname == "rsvpSenderAdspecGuaranteedCsum" { return "Rsvpsenderadspecguaranteedcsum" }
    if yname == "rsvpSenderAdspecGuaranteedDsum" { return "Rsvpsenderadspecguaranteeddsum" }
    if yname == "rsvpSenderAdspecGuaranteedHopCount" { return "Rsvpsenderadspecguaranteedhopcount" }
    if yname == "rsvpSenderAdspecGuaranteedPathBw" { return "Rsvpsenderadspecguaranteedpathbw" }
    if yname == "rsvpSenderAdspecGuaranteedMinLatency" { return "Rsvpsenderadspecguaranteedminlatency" }
    if yname == "rsvpSenderAdspecGuaranteedMtu" { return "Rsvpsenderadspecguaranteedmtu" }
    if yname == "rsvpSenderAdspecCtrlLoadSvc" { return "Rsvpsenderadspecctrlloadsvc" }
    if yname == "rsvpSenderAdspecCtrlLoadBreak" { return "Rsvpsenderadspecctrlloadbreak" }
    if yname == "rsvpSenderAdspecCtrlLoadHopCount" { return "Rsvpsenderadspecctrlloadhopcount" }
    if yname == "rsvpSenderAdspecCtrlLoadPathBw" { return "Rsvpsenderadspecctrlloadpathbw" }
    if yname == "rsvpSenderAdspecCtrlLoadMinLatency" { return "Rsvpsenderadspecctrlloadminlatency" }
    if yname == "rsvpSenderAdspecCtrlLoadMtu" { return "Rsvpsenderadspecctrlloadmtu" }
    if yname == "rsvpSenderStatus" { return "Rsvpsenderstatus" }
    if yname == "rsvpSenderTTL" { return "Rsvpsenderttl" }
    return ""
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetSegmentPath() string {
    return "rsvpSenderEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsenderentry.Rsvpsessionnumber) + "']" + "[rsvpSenderNumber='" + fmt.Sprintf("%v", rsvpsenderentry.Rsvpsendernumber) + "']"
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpSessionNumber"] = rsvpsenderentry.Rsvpsessionnumber
    leafs["rsvpSenderNumber"] = rsvpsenderentry.Rsvpsendernumber
    leafs["rsvpSenderType"] = rsvpsenderentry.Rsvpsendertype
    leafs["rsvpSenderDestAddr"] = rsvpsenderentry.Rsvpsenderdestaddr
    leafs["rsvpSenderAddr"] = rsvpsenderentry.Rsvpsenderaddr
    leafs["rsvpSenderDestAddrLength"] = rsvpsenderentry.Rsvpsenderdestaddrlength
    leafs["rsvpSenderAddrLength"] = rsvpsenderentry.Rsvpsenderaddrlength
    leafs["rsvpSenderProtocol"] = rsvpsenderentry.Rsvpsenderprotocol
    leafs["rsvpSenderDestPort"] = rsvpsenderentry.Rsvpsenderdestport
    leafs["rsvpSenderPort"] = rsvpsenderentry.Rsvpsenderport
    leafs["rsvpSenderFlowId"] = rsvpsenderentry.Rsvpsenderflowid
    leafs["rsvpSenderHopAddr"] = rsvpsenderentry.Rsvpsenderhopaddr
    leafs["rsvpSenderHopLih"] = rsvpsenderentry.Rsvpsenderhoplih
    leafs["rsvpSenderInterface"] = rsvpsenderentry.Rsvpsenderinterface
    leafs["rsvpSenderTSpecRate"] = rsvpsenderentry.Rsvpsendertspecrate
    leafs["rsvpSenderTSpecPeakRate"] = rsvpsenderentry.Rsvpsendertspecpeakrate
    leafs["rsvpSenderTSpecBurst"] = rsvpsenderentry.Rsvpsendertspecburst
    leafs["rsvpSenderTSpecMinTU"] = rsvpsenderentry.Rsvpsendertspecmintu
    leafs["rsvpSenderTSpecMaxTU"] = rsvpsenderentry.Rsvpsendertspecmaxtu
    leafs["rsvpSenderInterval"] = rsvpsenderentry.Rsvpsenderinterval
    leafs["rsvpSenderRSVPHop"] = rsvpsenderentry.Rsvpsenderrsvphop
    leafs["rsvpSenderLastChange"] = rsvpsenderentry.Rsvpsenderlastchange
    leafs["rsvpSenderPolicy"] = rsvpsenderentry.Rsvpsenderpolicy
    leafs["rsvpSenderAdspecBreak"] = rsvpsenderentry.Rsvpsenderadspecbreak
    leafs["rsvpSenderAdspecHopCount"] = rsvpsenderentry.Rsvpsenderadspechopcount
    leafs["rsvpSenderAdspecPathBw"] = rsvpsenderentry.Rsvpsenderadspecpathbw
    leafs["rsvpSenderAdspecMinLatency"] = rsvpsenderentry.Rsvpsenderadspecminlatency
    leafs["rsvpSenderAdspecMtu"] = rsvpsenderentry.Rsvpsenderadspecmtu
    leafs["rsvpSenderAdspecGuaranteedSvc"] = rsvpsenderentry.Rsvpsenderadspecguaranteedsvc
    leafs["rsvpSenderAdspecGuaranteedBreak"] = rsvpsenderentry.Rsvpsenderadspecguaranteedbreak
    leafs["rsvpSenderAdspecGuaranteedCtot"] = rsvpsenderentry.Rsvpsenderadspecguaranteedctot
    leafs["rsvpSenderAdspecGuaranteedDtot"] = rsvpsenderentry.Rsvpsenderadspecguaranteeddtot
    leafs["rsvpSenderAdspecGuaranteedCsum"] = rsvpsenderentry.Rsvpsenderadspecguaranteedcsum
    leafs["rsvpSenderAdspecGuaranteedDsum"] = rsvpsenderentry.Rsvpsenderadspecguaranteeddsum
    leafs["rsvpSenderAdspecGuaranteedHopCount"] = rsvpsenderentry.Rsvpsenderadspecguaranteedhopcount
    leafs["rsvpSenderAdspecGuaranteedPathBw"] = rsvpsenderentry.Rsvpsenderadspecguaranteedpathbw
    leafs["rsvpSenderAdspecGuaranteedMinLatency"] = rsvpsenderentry.Rsvpsenderadspecguaranteedminlatency
    leafs["rsvpSenderAdspecGuaranteedMtu"] = rsvpsenderentry.Rsvpsenderadspecguaranteedmtu
    leafs["rsvpSenderAdspecCtrlLoadSvc"] = rsvpsenderentry.Rsvpsenderadspecctrlloadsvc
    leafs["rsvpSenderAdspecCtrlLoadBreak"] = rsvpsenderentry.Rsvpsenderadspecctrlloadbreak
    leafs["rsvpSenderAdspecCtrlLoadHopCount"] = rsvpsenderentry.Rsvpsenderadspecctrlloadhopcount
    leafs["rsvpSenderAdspecCtrlLoadPathBw"] = rsvpsenderentry.Rsvpsenderadspecctrlloadpathbw
    leafs["rsvpSenderAdspecCtrlLoadMinLatency"] = rsvpsenderentry.Rsvpsenderadspecctrlloadminlatency
    leafs["rsvpSenderAdspecCtrlLoadMtu"] = rsvpsenderentry.Rsvpsenderadspecctrlloadmtu
    leafs["rsvpSenderStatus"] = rsvpsenderentry.Rsvpsenderstatus
    leafs["rsvpSenderTTL"] = rsvpsenderentry.Rsvpsenderttl
    return leafs
}

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetYangName() string { return "rsvpSenderEntry" }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) SetParent(parent types.Entity) { rsvpsenderentry.parent = parent }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetParent() types.Entity { return rsvpsenderentry.parent }

func (rsvpsenderentry *RSVPMIB_Rsvpsendertable_Rsvpsenderentry) GetParentYangName() string { return "rsvpSenderTable" }

// RSVPMIB_Rsvpsenderoutinterfacetable
// List of outgoing interfaces	that PATH messages
// use.	 The  ifIndex  is the ifIndex value of the
// egress interface.
type RSVPMIB_Rsvpsenderoutinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of outgoing interfaces	that a	particular PATH	message	has. The type is
    // slice of RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry.
    Rsvpsenderoutinterfaceentry []RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetFilter() yfilter.YFilter { return rsvpsenderoutinterfacetable.YFilter }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) SetFilter(yf yfilter.YFilter) { rsvpsenderoutinterfacetable.YFilter = yf }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetGoName(yname string) string {
    if yname == "rsvpSenderOutInterfaceEntry" { return "Rsvpsenderoutinterfaceentry" }
    return ""
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetSegmentPath() string {
    return "rsvpSenderOutInterfaceTable"
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpSenderOutInterfaceEntry" {
        for _, c := range rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry {
            if rsvpsenderoutinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry{}
        rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry = append(rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry, child)
        return &rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry[len(rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry)-1]
    }
    return nil
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry {
        children[rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry[i].GetSegmentPath()] = &rsvpsenderoutinterfacetable.Rsvpsenderoutinterfaceentry[i]
    }
    return children
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetYangName() string { return "rsvpSenderOutInterfaceTable" }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) SetParent(parent types.Entity) { rsvpsenderoutinterfacetable.parent = parent }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetParent() types.Entity { return rsvpsenderoutinterfacetable.parent }

func (rsvpsenderoutinterfacetable *RSVPMIB_Rsvpsenderoutinterfacetable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry
// List of outgoing interfaces	that a	particular
// PATH	message	has.
type RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry struct {
    parent types.Entity
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

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetFilter() yfilter.YFilter { return rsvpsenderoutinterfaceentry.YFilter }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) SetFilter(yf yfilter.YFilter) { rsvpsenderoutinterfaceentry.YFilter = yf }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetGoName(yname string) string {
    if yname == "rsvpSessionNumber" { return "Rsvpsessionnumber" }
    if yname == "rsvpSenderNumber" { return "Rsvpsendernumber" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "rsvpSenderOutInterfaceStatus" { return "Rsvpsenderoutinterfacestatus" }
    return ""
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetSegmentPath() string {
    return "rsvpSenderOutInterfaceEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Rsvpsessionnumber) + "']" + "[rsvpSenderNumber='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Rsvpsendernumber) + "']" + "[ifIndex='" + fmt.Sprintf("%v", rsvpsenderoutinterfaceentry.Ifindex) + "']"
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpSessionNumber"] = rsvpsenderoutinterfaceentry.Rsvpsessionnumber
    leafs["rsvpSenderNumber"] = rsvpsenderoutinterfaceentry.Rsvpsendernumber
    leafs["ifIndex"] = rsvpsenderoutinterfaceentry.Ifindex
    leafs["rsvpSenderOutInterfaceStatus"] = rsvpsenderoutinterfaceentry.Rsvpsenderoutinterfacestatus
    return leafs
}

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetYangName() string { return "rsvpSenderOutInterfaceEntry" }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) SetParent(parent types.Entity) { rsvpsenderoutinterfaceentry.parent = parent }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetParent() types.Entity { return rsvpsenderoutinterfaceentry.parent }

func (rsvpsenderoutinterfaceentry *RSVPMIB_Rsvpsenderoutinterfacetable_Rsvpsenderoutinterfaceentry) GetParentYangName() string { return "rsvpSenderOutInterfaceTable" }

// RSVPMIB_Rsvpresvtable
// Information	describing the	state  information
// displayed by	receivers in RESV messages.
type RSVPMIB_Rsvpresvtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed  by  a single
    // receiver's RESV message concerning a	single sender. The type is slice of
    // RSVPMIB_Rsvpresvtable_Rsvpresventry.
    Rsvpresventry []RSVPMIB_Rsvpresvtable_Rsvpresventry
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetFilter() yfilter.YFilter { return rsvpresvtable.YFilter }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) SetFilter(yf yfilter.YFilter) { rsvpresvtable.YFilter = yf }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetGoName(yname string) string {
    if yname == "rsvpResvEntry" { return "Rsvpresventry" }
    return ""
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetSegmentPath() string {
    return "rsvpResvTable"
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpResvEntry" {
        for _, c := range rsvpresvtable.Rsvpresventry {
            if rsvpresvtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpresvtable_Rsvpresventry{}
        rsvpresvtable.Rsvpresventry = append(rsvpresvtable.Rsvpresventry, child)
        return &rsvpresvtable.Rsvpresventry[len(rsvpresvtable.Rsvpresventry)-1]
    }
    return nil
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpresvtable.Rsvpresventry {
        children[rsvpresvtable.Rsvpresventry[i].GetSegmentPath()] = &rsvpresvtable.Rsvpresventry[i]
    }
    return children
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetYangName() string { return "rsvpResvTable" }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) SetParent(parent types.Entity) { rsvpresvtable.parent = parent }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetParent() types.Entity { return rsvpresvtable.parent }

func (rsvpresvtable *RSVPMIB_Rsvpresvtable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpresvtable_Rsvpresventry
// Information	describing the	state  information
// displayed  by  a single receiver's RESV message
// concerning a	single sender.
type RSVPMIB_Rsvpresvtable_Rsvpresventry struct {
    parent types.Entity
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

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetFilter() yfilter.YFilter { return rsvpresventry.YFilter }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) SetFilter(yf yfilter.YFilter) { rsvpresventry.YFilter = yf }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetGoName(yname string) string {
    if yname == "rsvpSessionNumber" { return "Rsvpsessionnumber" }
    if yname == "rsvpResvNumber" { return "Rsvpresvnumber" }
    if yname == "rsvpResvType" { return "Rsvpresvtype" }
    if yname == "rsvpResvDestAddr" { return "Rsvpresvdestaddr" }
    if yname == "rsvpResvSenderAddr" { return "Rsvpresvsenderaddr" }
    if yname == "rsvpResvDestAddrLength" { return "Rsvpresvdestaddrlength" }
    if yname == "rsvpResvSenderAddrLength" { return "Rsvpresvsenderaddrlength" }
    if yname == "rsvpResvProtocol" { return "Rsvpresvprotocol" }
    if yname == "rsvpResvDestPort" { return "Rsvpresvdestport" }
    if yname == "rsvpResvPort" { return "Rsvpresvport" }
    if yname == "rsvpResvHopAddr" { return "Rsvpresvhopaddr" }
    if yname == "rsvpResvHopLih" { return "Rsvpresvhoplih" }
    if yname == "rsvpResvInterface" { return "Rsvpresvinterface" }
    if yname == "rsvpResvService" { return "Rsvpresvservice" }
    if yname == "rsvpResvTSpecRate" { return "Rsvpresvtspecrate" }
    if yname == "rsvpResvTSpecPeakRate" { return "Rsvpresvtspecpeakrate" }
    if yname == "rsvpResvTSpecBurst" { return "Rsvpresvtspecburst" }
    if yname == "rsvpResvTSpecMinTU" { return "Rsvpresvtspecmintu" }
    if yname == "rsvpResvTSpecMaxTU" { return "Rsvpresvtspecmaxtu" }
    if yname == "rsvpResvRSpecRate" { return "Rsvpresvrspecrate" }
    if yname == "rsvpResvRSpecSlack" { return "Rsvpresvrspecslack" }
    if yname == "rsvpResvInterval" { return "Rsvpresvinterval" }
    if yname == "rsvpResvScope" { return "Rsvpresvscope" }
    if yname == "rsvpResvShared" { return "Rsvpresvshared" }
    if yname == "rsvpResvExplicit" { return "Rsvpresvexplicit" }
    if yname == "rsvpResvRSVPHop" { return "Rsvpresvrsvphop" }
    if yname == "rsvpResvLastChange" { return "Rsvpresvlastchange" }
    if yname == "rsvpResvPolicy" { return "Rsvpresvpolicy" }
    if yname == "rsvpResvStatus" { return "Rsvpresvstatus" }
    if yname == "rsvpResvTTL" { return "Rsvpresvttl" }
    if yname == "rsvpResvFlowId" { return "Rsvpresvflowid" }
    return ""
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetSegmentPath() string {
    return "rsvpResvEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpresventry.Rsvpsessionnumber) + "']" + "[rsvpResvNumber='" + fmt.Sprintf("%v", rsvpresventry.Rsvpresvnumber) + "']"
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpSessionNumber"] = rsvpresventry.Rsvpsessionnumber
    leafs["rsvpResvNumber"] = rsvpresventry.Rsvpresvnumber
    leafs["rsvpResvType"] = rsvpresventry.Rsvpresvtype
    leafs["rsvpResvDestAddr"] = rsvpresventry.Rsvpresvdestaddr
    leafs["rsvpResvSenderAddr"] = rsvpresventry.Rsvpresvsenderaddr
    leafs["rsvpResvDestAddrLength"] = rsvpresventry.Rsvpresvdestaddrlength
    leafs["rsvpResvSenderAddrLength"] = rsvpresventry.Rsvpresvsenderaddrlength
    leafs["rsvpResvProtocol"] = rsvpresventry.Rsvpresvprotocol
    leafs["rsvpResvDestPort"] = rsvpresventry.Rsvpresvdestport
    leafs["rsvpResvPort"] = rsvpresventry.Rsvpresvport
    leafs["rsvpResvHopAddr"] = rsvpresventry.Rsvpresvhopaddr
    leafs["rsvpResvHopLih"] = rsvpresventry.Rsvpresvhoplih
    leafs["rsvpResvInterface"] = rsvpresventry.Rsvpresvinterface
    leafs["rsvpResvService"] = rsvpresventry.Rsvpresvservice
    leafs["rsvpResvTSpecRate"] = rsvpresventry.Rsvpresvtspecrate
    leafs["rsvpResvTSpecPeakRate"] = rsvpresventry.Rsvpresvtspecpeakrate
    leafs["rsvpResvTSpecBurst"] = rsvpresventry.Rsvpresvtspecburst
    leafs["rsvpResvTSpecMinTU"] = rsvpresventry.Rsvpresvtspecmintu
    leafs["rsvpResvTSpecMaxTU"] = rsvpresventry.Rsvpresvtspecmaxtu
    leafs["rsvpResvRSpecRate"] = rsvpresventry.Rsvpresvrspecrate
    leafs["rsvpResvRSpecSlack"] = rsvpresventry.Rsvpresvrspecslack
    leafs["rsvpResvInterval"] = rsvpresventry.Rsvpresvinterval
    leafs["rsvpResvScope"] = rsvpresventry.Rsvpresvscope
    leafs["rsvpResvShared"] = rsvpresventry.Rsvpresvshared
    leafs["rsvpResvExplicit"] = rsvpresventry.Rsvpresvexplicit
    leafs["rsvpResvRSVPHop"] = rsvpresventry.Rsvpresvrsvphop
    leafs["rsvpResvLastChange"] = rsvpresventry.Rsvpresvlastchange
    leafs["rsvpResvPolicy"] = rsvpresventry.Rsvpresvpolicy
    leafs["rsvpResvStatus"] = rsvpresventry.Rsvpresvstatus
    leafs["rsvpResvTTL"] = rsvpresventry.Rsvpresvttl
    leafs["rsvpResvFlowId"] = rsvpresventry.Rsvpresvflowid
    return leafs
}

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetYangName() string { return "rsvpResvEntry" }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) SetParent(parent types.Entity) { rsvpresventry.parent = parent }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetParent() types.Entity { return rsvpresventry.parent }

func (rsvpresventry *RSVPMIB_Rsvpresvtable_Rsvpresventry) GetParentYangName() string { return "rsvpResvTable" }

// RSVPMIB_Rsvpresvfwdtable
// Information	describing the	state  information
// displayed upstream in RESV messages.
type RSVPMIB_Rsvpresvfwdtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information	describing the	state  information displayed   upstream	  in  
    // an   RESV   message concerning a	single sender. The type is slice of
    // RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry.
    Rsvpresvfwdentry []RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetFilter() yfilter.YFilter { return rsvpresvfwdtable.YFilter }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) SetFilter(yf yfilter.YFilter) { rsvpresvfwdtable.YFilter = yf }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetGoName(yname string) string {
    if yname == "rsvpResvFwdEntry" { return "Rsvpresvfwdentry" }
    return ""
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetSegmentPath() string {
    return "rsvpResvFwdTable"
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpResvFwdEntry" {
        for _, c := range rsvpresvfwdtable.Rsvpresvfwdentry {
            if rsvpresvfwdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry{}
        rsvpresvfwdtable.Rsvpresvfwdentry = append(rsvpresvfwdtable.Rsvpresvfwdentry, child)
        return &rsvpresvfwdtable.Rsvpresvfwdentry[len(rsvpresvfwdtable.Rsvpresvfwdentry)-1]
    }
    return nil
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpresvfwdtable.Rsvpresvfwdentry {
        children[rsvpresvfwdtable.Rsvpresvfwdentry[i].GetSegmentPath()] = &rsvpresvfwdtable.Rsvpresvfwdentry[i]
    }
    return children
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetYangName() string { return "rsvpResvFwdTable" }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) SetParent(parent types.Entity) { rsvpresvfwdtable.parent = parent }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetParent() types.Entity { return rsvpresvfwdtable.parent }

func (rsvpresvfwdtable *RSVPMIB_Rsvpresvfwdtable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry
// Information	describing the	state  information
// displayed   upstream	  in   an   RESV   message
// concerning a	single sender.
type RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry struct {
    parent types.Entity
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

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetFilter() yfilter.YFilter { return rsvpresvfwdentry.YFilter }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) SetFilter(yf yfilter.YFilter) { rsvpresvfwdentry.YFilter = yf }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetGoName(yname string) string {
    if yname == "rsvpSessionNumber" { return "Rsvpsessionnumber" }
    if yname == "rsvpResvFwdNumber" { return "Rsvpresvfwdnumber" }
    if yname == "rsvpResvFwdType" { return "Rsvpresvfwdtype" }
    if yname == "rsvpResvFwdDestAddr" { return "Rsvpresvfwddestaddr" }
    if yname == "rsvpResvFwdSenderAddr" { return "Rsvpresvfwdsenderaddr" }
    if yname == "rsvpResvFwdDestAddrLength" { return "Rsvpresvfwddestaddrlength" }
    if yname == "rsvpResvFwdSenderAddrLength" { return "Rsvpresvfwdsenderaddrlength" }
    if yname == "rsvpResvFwdProtocol" { return "Rsvpresvfwdprotocol" }
    if yname == "rsvpResvFwdDestPort" { return "Rsvpresvfwddestport" }
    if yname == "rsvpResvFwdPort" { return "Rsvpresvfwdport" }
    if yname == "rsvpResvFwdHopAddr" { return "Rsvpresvfwdhopaddr" }
    if yname == "rsvpResvFwdHopLih" { return "Rsvpresvfwdhoplih" }
    if yname == "rsvpResvFwdInterface" { return "Rsvpresvfwdinterface" }
    if yname == "rsvpResvFwdService" { return "Rsvpresvfwdservice" }
    if yname == "rsvpResvFwdTSpecRate" { return "Rsvpresvfwdtspecrate" }
    if yname == "rsvpResvFwdTSpecPeakRate" { return "Rsvpresvfwdtspecpeakrate" }
    if yname == "rsvpResvFwdTSpecBurst" { return "Rsvpresvfwdtspecburst" }
    if yname == "rsvpResvFwdTSpecMinTU" { return "Rsvpresvfwdtspecmintu" }
    if yname == "rsvpResvFwdTSpecMaxTU" { return "Rsvpresvfwdtspecmaxtu" }
    if yname == "rsvpResvFwdRSpecRate" { return "Rsvpresvfwdrspecrate" }
    if yname == "rsvpResvFwdRSpecSlack" { return "Rsvpresvfwdrspecslack" }
    if yname == "rsvpResvFwdInterval" { return "Rsvpresvfwdinterval" }
    if yname == "rsvpResvFwdScope" { return "Rsvpresvfwdscope" }
    if yname == "rsvpResvFwdShared" { return "Rsvpresvfwdshared" }
    if yname == "rsvpResvFwdExplicit" { return "Rsvpresvfwdexplicit" }
    if yname == "rsvpResvFwdRSVPHop" { return "Rsvpresvfwdrsvphop" }
    if yname == "rsvpResvFwdLastChange" { return "Rsvpresvfwdlastchange" }
    if yname == "rsvpResvFwdPolicy" { return "Rsvpresvfwdpolicy" }
    if yname == "rsvpResvFwdStatus" { return "Rsvpresvfwdstatus" }
    if yname == "rsvpResvFwdTTL" { return "Rsvpresvfwdttl" }
    if yname == "rsvpResvFwdFlowId" { return "Rsvpresvfwdflowid" }
    return ""
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetSegmentPath() string {
    return "rsvpResvFwdEntry" + "[rsvpSessionNumber='" + fmt.Sprintf("%v", rsvpresvfwdentry.Rsvpsessionnumber) + "']" + "[rsvpResvFwdNumber='" + fmt.Sprintf("%v", rsvpresvfwdentry.Rsvpresvfwdnumber) + "']"
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rsvpSessionNumber"] = rsvpresvfwdentry.Rsvpsessionnumber
    leafs["rsvpResvFwdNumber"] = rsvpresvfwdentry.Rsvpresvfwdnumber
    leafs["rsvpResvFwdType"] = rsvpresvfwdentry.Rsvpresvfwdtype
    leafs["rsvpResvFwdDestAddr"] = rsvpresvfwdentry.Rsvpresvfwddestaddr
    leafs["rsvpResvFwdSenderAddr"] = rsvpresvfwdentry.Rsvpresvfwdsenderaddr
    leafs["rsvpResvFwdDestAddrLength"] = rsvpresvfwdentry.Rsvpresvfwddestaddrlength
    leafs["rsvpResvFwdSenderAddrLength"] = rsvpresvfwdentry.Rsvpresvfwdsenderaddrlength
    leafs["rsvpResvFwdProtocol"] = rsvpresvfwdentry.Rsvpresvfwdprotocol
    leafs["rsvpResvFwdDestPort"] = rsvpresvfwdentry.Rsvpresvfwddestport
    leafs["rsvpResvFwdPort"] = rsvpresvfwdentry.Rsvpresvfwdport
    leafs["rsvpResvFwdHopAddr"] = rsvpresvfwdentry.Rsvpresvfwdhopaddr
    leafs["rsvpResvFwdHopLih"] = rsvpresvfwdentry.Rsvpresvfwdhoplih
    leafs["rsvpResvFwdInterface"] = rsvpresvfwdentry.Rsvpresvfwdinterface
    leafs["rsvpResvFwdService"] = rsvpresvfwdentry.Rsvpresvfwdservice
    leafs["rsvpResvFwdTSpecRate"] = rsvpresvfwdentry.Rsvpresvfwdtspecrate
    leafs["rsvpResvFwdTSpecPeakRate"] = rsvpresvfwdentry.Rsvpresvfwdtspecpeakrate
    leafs["rsvpResvFwdTSpecBurst"] = rsvpresvfwdentry.Rsvpresvfwdtspecburst
    leafs["rsvpResvFwdTSpecMinTU"] = rsvpresvfwdentry.Rsvpresvfwdtspecmintu
    leafs["rsvpResvFwdTSpecMaxTU"] = rsvpresvfwdentry.Rsvpresvfwdtspecmaxtu
    leafs["rsvpResvFwdRSpecRate"] = rsvpresvfwdentry.Rsvpresvfwdrspecrate
    leafs["rsvpResvFwdRSpecSlack"] = rsvpresvfwdentry.Rsvpresvfwdrspecslack
    leafs["rsvpResvFwdInterval"] = rsvpresvfwdentry.Rsvpresvfwdinterval
    leafs["rsvpResvFwdScope"] = rsvpresvfwdentry.Rsvpresvfwdscope
    leafs["rsvpResvFwdShared"] = rsvpresvfwdentry.Rsvpresvfwdshared
    leafs["rsvpResvFwdExplicit"] = rsvpresvfwdentry.Rsvpresvfwdexplicit
    leafs["rsvpResvFwdRSVPHop"] = rsvpresvfwdentry.Rsvpresvfwdrsvphop
    leafs["rsvpResvFwdLastChange"] = rsvpresvfwdentry.Rsvpresvfwdlastchange
    leafs["rsvpResvFwdPolicy"] = rsvpresvfwdentry.Rsvpresvfwdpolicy
    leafs["rsvpResvFwdStatus"] = rsvpresvfwdentry.Rsvpresvfwdstatus
    leafs["rsvpResvFwdTTL"] = rsvpresvfwdentry.Rsvpresvfwdttl
    leafs["rsvpResvFwdFlowId"] = rsvpresvfwdentry.Rsvpresvfwdflowid
    return leafs
}

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetYangName() string { return "rsvpResvFwdEntry" }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) SetParent(parent types.Entity) { rsvpresvfwdentry.parent = parent }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetParent() types.Entity { return rsvpresvfwdentry.parent }

func (rsvpresvfwdentry *RSVPMIB_Rsvpresvfwdtable_Rsvpresvfwdentry) GetParentYangName() string { return "rsvpResvFwdTable" }

// RSVPMIB_Rsvpiftable
// The	RSVP-specific attributes of  the  system's
// interfaces.
type RSVPMIB_Rsvpiftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The	RSVP-specific attributes of  the  a  given interface. The type is slice
    // of RSVPMIB_Rsvpiftable_Rsvpifentry.
    Rsvpifentry []RSVPMIB_Rsvpiftable_Rsvpifentry
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetFilter() yfilter.YFilter { return rsvpiftable.YFilter }

func (rsvpiftable *RSVPMIB_Rsvpiftable) SetFilter(yf yfilter.YFilter) { rsvpiftable.YFilter = yf }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetGoName(yname string) string {
    if yname == "rsvpIfEntry" { return "Rsvpifentry" }
    return ""
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetSegmentPath() string {
    return "rsvpIfTable"
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpIfEntry" {
        for _, c := range rsvpiftable.Rsvpifentry {
            if rsvpiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpiftable_Rsvpifentry{}
        rsvpiftable.Rsvpifentry = append(rsvpiftable.Rsvpifentry, child)
        return &rsvpiftable.Rsvpifentry[len(rsvpiftable.Rsvpifentry)-1]
    }
    return nil
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpiftable.Rsvpifentry {
        children[rsvpiftable.Rsvpifentry[i].GetSegmentPath()] = &rsvpiftable.Rsvpifentry[i]
    }
    return children
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetYangName() string { return "rsvpIfTable" }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpiftable *RSVPMIB_Rsvpiftable) SetParent(parent types.Entity) { rsvpiftable.parent = parent }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetParent() types.Entity { return rsvpiftable.parent }

func (rsvpiftable *RSVPMIB_Rsvpiftable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpiftable_Rsvpifentry
// The	RSVP-specific attributes of  the  a  given
// interface.
type RSVPMIB_Rsvpiftable_Rsvpifentry struct {
    parent types.Entity
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

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetFilter() yfilter.YFilter { return rsvpifentry.YFilter }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) SetFilter(yf yfilter.YFilter) { rsvpifentry.YFilter = yf }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "rsvpIfUdpNbrs" { return "Rsvpifudpnbrs" }
    if yname == "rsvpIfIpNbrs" { return "Rsvpifipnbrs" }
    if yname == "rsvpIfNbrs" { return "Rsvpifnbrs" }
    if yname == "rsvpIfRefreshBlockadeMultiple" { return "Rsvpifrefreshblockademultiple" }
    if yname == "rsvpIfRefreshMultiple" { return "Rsvpifrefreshmultiple" }
    if yname == "rsvpIfTTL" { return "Rsvpifttl" }
    if yname == "rsvpIfRefreshInterval" { return "Rsvpifrefreshinterval" }
    if yname == "rsvpIfRouteDelay" { return "Rsvpifroutedelay" }
    if yname == "rsvpIfEnabled" { return "Rsvpifenabled" }
    if yname == "rsvpIfUdpRequired" { return "Rsvpifudprequired" }
    if yname == "rsvpIfStatus" { return "Rsvpifstatus" }
    return ""
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetSegmentPath() string {
    return "rsvpIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", rsvpifentry.Ifindex) + "']"
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = rsvpifentry.Ifindex
    leafs["rsvpIfUdpNbrs"] = rsvpifentry.Rsvpifudpnbrs
    leafs["rsvpIfIpNbrs"] = rsvpifentry.Rsvpifipnbrs
    leafs["rsvpIfNbrs"] = rsvpifentry.Rsvpifnbrs
    leafs["rsvpIfRefreshBlockadeMultiple"] = rsvpifentry.Rsvpifrefreshblockademultiple
    leafs["rsvpIfRefreshMultiple"] = rsvpifentry.Rsvpifrefreshmultiple
    leafs["rsvpIfTTL"] = rsvpifentry.Rsvpifttl
    leafs["rsvpIfRefreshInterval"] = rsvpifentry.Rsvpifrefreshinterval
    leafs["rsvpIfRouteDelay"] = rsvpifentry.Rsvpifroutedelay
    leafs["rsvpIfEnabled"] = rsvpifentry.Rsvpifenabled
    leafs["rsvpIfUdpRequired"] = rsvpifentry.Rsvpifudprequired
    leafs["rsvpIfStatus"] = rsvpifentry.Rsvpifstatus
    return leafs
}

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetYangName() string { return "rsvpIfEntry" }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) SetParent(parent types.Entity) { rsvpifentry.parent = parent }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetParent() types.Entity { return rsvpifentry.parent }

func (rsvpifentry *RSVPMIB_Rsvpiftable_Rsvpifentry) GetParentYangName() string { return "rsvpIfTable" }

// RSVPMIB_Rsvpnbrtable
// Information	describing  the	 Neighbors  of	an
// RSVP	system.
type RSVPMIB_Rsvpnbrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information	  describing   a    single    RSVP Neighbor. The type is slice
    // of RSVPMIB_Rsvpnbrtable_Rsvpnbrentry.
    Rsvpnbrentry []RSVPMIB_Rsvpnbrtable_Rsvpnbrentry
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetFilter() yfilter.YFilter { return rsvpnbrtable.YFilter }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) SetFilter(yf yfilter.YFilter) { rsvpnbrtable.YFilter = yf }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetGoName(yname string) string {
    if yname == "rsvpNbrEntry" { return "Rsvpnbrentry" }
    return ""
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetSegmentPath() string {
    return "rsvpNbrTable"
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvpNbrEntry" {
        for _, c := range rsvpnbrtable.Rsvpnbrentry {
            if rsvpnbrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RSVPMIB_Rsvpnbrtable_Rsvpnbrentry{}
        rsvpnbrtable.Rsvpnbrentry = append(rsvpnbrtable.Rsvpnbrentry, child)
        return &rsvpnbrtable.Rsvpnbrentry[len(rsvpnbrtable.Rsvpnbrentry)-1]
    }
    return nil
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rsvpnbrtable.Rsvpnbrentry {
        children[rsvpnbrtable.Rsvpnbrentry[i].GetSegmentPath()] = &rsvpnbrtable.Rsvpnbrentry[i]
    }
    return children
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetYangName() string { return "rsvpNbrTable" }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) SetParent(parent types.Entity) { rsvpnbrtable.parent = parent }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetParent() types.Entity { return rsvpnbrtable.parent }

func (rsvpnbrtable *RSVPMIB_Rsvpnbrtable) GetParentYangName() string { return "RSVP-MIB" }

// RSVPMIB_Rsvpnbrtable_Rsvpnbrentry
// Information	  describing   a    single    RSVP
// Neighbor.
type RSVPMIB_Rsvpnbrtable_Rsvpnbrentry struct {
    parent types.Entity
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

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetFilter() yfilter.YFilter { return rsvpnbrentry.YFilter }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) SetFilter(yf yfilter.YFilter) { rsvpnbrentry.YFilter = yf }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "rsvpNbrAddress" { return "Rsvpnbraddress" }
    if yname == "rsvpNbrProtocol" { return "Rsvpnbrprotocol" }
    if yname == "rsvpNbrStatus" { return "Rsvpnbrstatus" }
    return ""
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetSegmentPath() string {
    return "rsvpNbrEntry" + "[ifIndex='" + fmt.Sprintf("%v", rsvpnbrentry.Ifindex) + "']" + "[rsvpNbrAddress='" + fmt.Sprintf("%v", rsvpnbrentry.Rsvpnbraddress) + "']"
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = rsvpnbrentry.Ifindex
    leafs["rsvpNbrAddress"] = rsvpnbrentry.Rsvpnbraddress
    leafs["rsvpNbrProtocol"] = rsvpnbrentry.Rsvpnbrprotocol
    leafs["rsvpNbrStatus"] = rsvpnbrentry.Rsvpnbrstatus
    return leafs
}

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetBundleName() string { return "cisco_ios_xe" }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetYangName() string { return "rsvpNbrEntry" }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) SetParent(parent types.Entity) { rsvpnbrentry.parent = parent }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetParent() types.Entity { return rsvpnbrentry.parent }

func (rsvpnbrentry *RSVPMIB_Rsvpnbrtable_Rsvpnbrentry) GetParentYangName() string { return "rsvpNbrTable" }

