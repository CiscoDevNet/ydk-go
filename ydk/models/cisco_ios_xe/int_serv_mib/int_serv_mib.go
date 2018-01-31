// The MIB module to describe the Integrated Services
// Protocol
package int_serv_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package int_serv_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:INT-SERV-MIB INT-SERV-MIB}", reflect.TypeOf(INTSERVMIB{}))
    ydk.RegisterEntity("INT-SERV-MIB:INT-SERV-MIB", reflect.TypeOf(INTSERVMIB{}))
}

// QosService represents The class of service in use by a flow.
type QosService string

const (
    QosService_bestEffort QosService = "bestEffort"

    QosService_guaranteedDelay QosService = "guaranteedDelay"

    QosService_controlledLoad QosService = "controlledLoad"
)

// INTSERVMIB
type INTSERVMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Intsrvgenobjects INTSERVMIB_Intsrvgenobjects

    // The reservable attributes of the system's  in- terfaces.
    Intsrvifattribtable INTSERVMIB_Intsrvifattribtable

    // Information describing the reserved flows  us- ing the system's interfaces.
    Intsrvflowtable INTSERVMIB_Intsrvflowtable
}

func (iNTSERVMIB *INTSERVMIB) GetFilter() yfilter.YFilter { return iNTSERVMIB.YFilter }

func (iNTSERVMIB *INTSERVMIB) SetFilter(yf yfilter.YFilter) { iNTSERVMIB.YFilter = yf }

func (iNTSERVMIB *INTSERVMIB) GetGoName(yname string) string {
    if yname == "intSrvGenObjects" { return "Intsrvgenobjects" }
    if yname == "intSrvIfAttribTable" { return "Intsrvifattribtable" }
    if yname == "intSrvFlowTable" { return "Intsrvflowtable" }
    return ""
}

func (iNTSERVMIB *INTSERVMIB) GetSegmentPath() string {
    return "INT-SERV-MIB:INT-SERV-MIB"
}

func (iNTSERVMIB *INTSERVMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvGenObjects" {
        return &iNTSERVMIB.Intsrvgenobjects
    }
    if childYangName == "intSrvIfAttribTable" {
        return &iNTSERVMIB.Intsrvifattribtable
    }
    if childYangName == "intSrvFlowTable" {
        return &iNTSERVMIB.Intsrvflowtable
    }
    return nil
}

func (iNTSERVMIB *INTSERVMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intSrvGenObjects"] = &iNTSERVMIB.Intsrvgenobjects
    children["intSrvIfAttribTable"] = &iNTSERVMIB.Intsrvifattribtable
    children["intSrvFlowTable"] = &iNTSERVMIB.Intsrvflowtable
    return children
}

func (iNTSERVMIB *INTSERVMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iNTSERVMIB *INTSERVMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iNTSERVMIB *INTSERVMIB) GetYangName() string { return "INT-SERV-MIB" }

func (iNTSERVMIB *INTSERVMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iNTSERVMIB *INTSERVMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iNTSERVMIB *INTSERVMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iNTSERVMIB *INTSERVMIB) SetParent(parent types.Entity) { iNTSERVMIB.parent = parent }

func (iNTSERVMIB *INTSERVMIB) GetParent() types.Entity { return iNTSERVMIB.parent }

func (iNTSERVMIB *INTSERVMIB) GetParentYangName() string { return "INT-SERV-MIB" }

// INTSERVMIB_Intsrvgenobjects
type INTSERVMIB_Intsrvgenobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This  object  is  used  to  assign  values  to intSrvFlowNumber  as
    // described in 'Textual Con- ventions  for  SNMPv2'.   The  network  manager
    // reads  the  object,  and  then writes the value back in the SET that
    // creates a new instance  of intSrvFlowEntry.   If  the  SET  fails with the
    // code 'inconsistentValue', then the process must be  repeated; If the SET
    // succeeds, then the ob- ject is incremented, and the  new  instance  is
    // created according to the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    Intsrvflownewindex interface{}
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetFilter() yfilter.YFilter { return intsrvgenobjects.YFilter }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) SetFilter(yf yfilter.YFilter) { intsrvgenobjects.YFilter = yf }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetGoName(yname string) string {
    if yname == "intSrvFlowNewIndex" { return "Intsrvflownewindex" }
    return ""
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetSegmentPath() string {
    return "intSrvGenObjects"
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intSrvFlowNewIndex"] = intsrvgenobjects.Intsrvflownewindex
    return leafs
}

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetYangName() string { return "intSrvGenObjects" }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) SetParent(parent types.Entity) { intsrvgenobjects.parent = parent }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetParent() types.Entity { return intsrvgenobjects.parent }

func (intsrvgenobjects *INTSERVMIB_Intsrvgenobjects) GetParentYangName() string { return "INT-SERV-MIB" }

// INTSERVMIB_Intsrvifattribtable
// The reservable attributes of the system's  in-
// terfaces.
type INTSERVMIB_Intsrvifattribtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The reservable attributes of  a  given  inter- face. The type is slice of
    // INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry.
    Intsrvifattribentry []INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetFilter() yfilter.YFilter { return intsrvifattribtable.YFilter }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) SetFilter(yf yfilter.YFilter) { intsrvifattribtable.YFilter = yf }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetGoName(yname string) string {
    if yname == "intSrvIfAttribEntry" { return "Intsrvifattribentry" }
    return ""
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetSegmentPath() string {
    return "intSrvIfAttribTable"
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvIfAttribEntry" {
        for _, c := range intsrvifattribtable.Intsrvifattribentry {
            if intsrvifattribtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry{}
        intsrvifattribtable.Intsrvifattribentry = append(intsrvifattribtable.Intsrvifattribentry, child)
        return &intsrvifattribtable.Intsrvifattribentry[len(intsrvifattribtable.Intsrvifattribentry)-1]
    }
    return nil
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intsrvifattribtable.Intsrvifattribentry {
        children[intsrvifattribtable.Intsrvifattribentry[i].GetSegmentPath()] = &intsrvifattribtable.Intsrvifattribentry[i]
    }
    return children
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetYangName() string { return "intSrvIfAttribTable" }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) SetParent(parent types.Entity) { intsrvifattribtable.parent = parent }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetParent() types.Entity { return intsrvifattribtable.parent }

func (intsrvifattribtable *INTSERVMIB_Intsrvifattribtable) GetParentYangName() string { return "INT-SERV-MIB" }

// INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry
// The reservable attributes of  a  given  inter-
// face.
type INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of bits/second currently  allocated to reserved sessions on the
    // interface. The type is interface{} with range: 0..2147483647. Units are
    // Bits per second.
    Intsrvifattriballocatedbits interface{}

    // The maximum number of bits/second that may  be allocated  to  reserved 
    // sessions on the inter- face. The type is interface{} with range:
    // 0..2147483647. Units are Bits per second.
    Intsrvifattribmaxallocatedbits interface{}

    // The amount of buffer space  required  to  hold the simultaneous burst of
    // all reserved flows on the interface. The type is interface{} with range:
    // 0..2147483647. Units are Bytes.
    Intsrvifattriballocatedbuffer interface{}

    // The number of reserved flows currently  active on  this  interface.  A flow
    // can be created ei- ther from a reservation protocol (such as  RSVP or
    // ST-II) or via configuration information. The type is interface{} with
    // range: 0..4294967295.
    Intsrvifattribflows interface{}

    // The amount of propagation delay that this  in- terface  introduces  in
    // addition to that intro- diced by bit propagation delays. The type is
    // interface{} with range: -2147483648..2147483647. Units are microseconds.
    Intsrvifattribpropagationdelay interface{}

    // 'active' on interfaces that are configured for RSVP. The type is RowStatus.
    Intsrvifattribstatus interface{}
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetFilter() yfilter.YFilter { return intsrvifattribentry.YFilter }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) SetFilter(yf yfilter.YFilter) { intsrvifattribentry.YFilter = yf }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "intSrvIfAttribAllocatedBits" { return "Intsrvifattriballocatedbits" }
    if yname == "intSrvIfAttribMaxAllocatedBits" { return "Intsrvifattribmaxallocatedbits" }
    if yname == "intSrvIfAttribAllocatedBuffer" { return "Intsrvifattriballocatedbuffer" }
    if yname == "intSrvIfAttribFlows" { return "Intsrvifattribflows" }
    if yname == "intSrvIfAttribPropagationDelay" { return "Intsrvifattribpropagationdelay" }
    if yname == "intSrvIfAttribStatus" { return "Intsrvifattribstatus" }
    return ""
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetSegmentPath() string {
    return "intSrvIfAttribEntry" + "[ifIndex='" + fmt.Sprintf("%v", intsrvifattribentry.Ifindex) + "']"
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = intsrvifattribentry.Ifindex
    leafs["intSrvIfAttribAllocatedBits"] = intsrvifattribentry.Intsrvifattriballocatedbits
    leafs["intSrvIfAttribMaxAllocatedBits"] = intsrvifattribentry.Intsrvifattribmaxallocatedbits
    leafs["intSrvIfAttribAllocatedBuffer"] = intsrvifattribentry.Intsrvifattriballocatedbuffer
    leafs["intSrvIfAttribFlows"] = intsrvifattribentry.Intsrvifattribflows
    leafs["intSrvIfAttribPropagationDelay"] = intsrvifattribentry.Intsrvifattribpropagationdelay
    leafs["intSrvIfAttribStatus"] = intsrvifattribentry.Intsrvifattribstatus
    return leafs
}

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetYangName() string { return "intSrvIfAttribEntry" }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) SetParent(parent types.Entity) { intsrvifattribentry.parent = parent }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetParent() types.Entity { return intsrvifattribentry.parent }

func (intsrvifattribentry *INTSERVMIB_Intsrvifattribtable_Intsrvifattribentry) GetParentYangName() string { return "intSrvIfAttribTable" }

// INTSERVMIB_Intsrvflowtable
// Information describing the reserved flows  us-
// ing the system's interfaces.
type INTSERVMIB_Intsrvflowtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information describing the use of a given  in- terface   by   a   given  
    // flow.   The  counter intSrvFlowPoliced starts counting  at  the  in-
    // stallation of the flow. The type is slice of
    // INTSERVMIB_Intsrvflowtable_Intsrvflowentry.
    Intsrvflowentry []INTSERVMIB_Intsrvflowtable_Intsrvflowentry
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetFilter() yfilter.YFilter { return intsrvflowtable.YFilter }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) SetFilter(yf yfilter.YFilter) { intsrvflowtable.YFilter = yf }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetGoName(yname string) string {
    if yname == "intSrvFlowEntry" { return "Intsrvflowentry" }
    return ""
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetSegmentPath() string {
    return "intSrvFlowTable"
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvFlowEntry" {
        for _, c := range intsrvflowtable.Intsrvflowentry {
            if intsrvflowtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := INTSERVMIB_Intsrvflowtable_Intsrvflowentry{}
        intsrvflowtable.Intsrvflowentry = append(intsrvflowtable.Intsrvflowentry, child)
        return &intsrvflowtable.Intsrvflowentry[len(intsrvflowtable.Intsrvflowentry)-1]
    }
    return nil
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intsrvflowtable.Intsrvflowentry {
        children[intsrvflowtable.Intsrvflowentry[i].GetSegmentPath()] = &intsrvflowtable.Intsrvflowentry[i]
    }
    return children
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetYangName() string { return "intSrvFlowTable" }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) SetParent(parent types.Entity) { intsrvflowtable.parent = parent }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetParent() types.Entity { return intsrvflowtable.parent }

func (intsrvflowtable *INTSERVMIB_Intsrvflowtable) GetParentYangName() string { return "INT-SERV-MIB" }

// INTSERVMIB_Intsrvflowtable_Intsrvflowentry
// Information describing the use of a given  in-
// terface   by   a   given   flow.   The  counter
// intSrvFlowPoliced starts counting  at  the  in-
// stallation of the flow.
type INTSERVMIB_Intsrvflowtable_Intsrvflowentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The number of this flow.  This is for SNMP In-
    // dexing purposes only and has no relation to any protocol value. The type is
    // interface{} with range: 0..2147483647.
    Intsrvflownumber interface{}

    // The type of session (IP4, IP6, IP6  with  flow information, etc). The type
    // is interface{} with range: 1..255.
    Intsrvflowtype interface{}

    // The process that installed this  flow  in  the queue policy database. The
    // type is Intsrvflowowner.
    Intsrvflowowner interface{}

    // The destination address used by all senders in this  session.   This object
    // may not be changed when the value of the RowStatus object is  'ac- tive'.
    // The type is string with length: 4..16.
    Intsrvflowdestaddr interface{}

    // The source address of the sender  selected  by this  reservation.  The
    // value of all zeroes in- dicates 'all senders'.  This object may not  be
    // changed  when the value of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    Intsrvflowsenderaddr interface{}

    // The length of the destination address in bits. This  is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Intsrvflowdestaddrlength interface{}

    // The length of the sender's  address  in  bits. This  is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    Intsrvflowsenderaddrlength interface{}

    // The IP Protocol used by a session.   This  ob- ject  may  not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    Intsrvflowprotocol interface{}

    // The UDP or TCP port number used as a  destina- tion  port for all senders
    // in this session.  If the  IP   protocol   in   use,   specified   by
    // intSrvResvFwdProtocol,  is 50 (ESP) or 51 (AH), this  represents  a 
    // virtual  destination  port number.   A value of zero indicates that the IP
    // protocol in use does not have ports.  This  ob- ject  may  not be changed
    // when the value of the RowStatus object is 'active'. The type is string with
    // length: 2..4.
    Intsrvflowdestport interface{}

    // The UDP or TCP port number used  as  a  source port  for  this sender in
    // this session.  If the IP    protocol    in    use,    specified    by
    // intSrvResvFwdProtocol  is  50 (ESP) or 51 (AH), this represents a
    // generalized  port  identifier (GPI).   A  value of zero indicates that the
    // IP protocol in use does not have ports.  This  ob- ject  may  not be
    // changed when the value of the RowStatus object is 'active'. The type is
    // string with length: 2..4.
    Intsrvflowport interface{}

    // The flow ID that  this  sender  is  using,  if this  is  an IPv6 session.
    // The type is interface{} with range: 0..16777215.
    Intsrvflowflowid interface{}

    // The ifIndex value of the  interface  on  which this reservation exists. The
    // type is interface{} with range: 1..2147483647.
    Intsrvflowinterface interface{}

    // The IP Address on the ifEntry  on  which  this reservation  exists.  This
    // is present primarily to support those interfaces which layer  multi- ple IP
    // Addresses on the interface. The type is string with length: 4..16.
    Intsrvflowifaddr interface{}

    // The Reserved Rate of the sender's data stream. If this is a Controlled Load
    // service flow, this rate is derived from the Tspec  rate  parameter (r).  
    // If  this  is  a Guaranteed service flow, this rate is derived from  the 
    // Rspec  clearing rate parameter (R). The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    Intsrvflowrate interface{}

    // The size of the largest  burst  expected  from the sender at a time.  If
    // this is less than  the  sender's  advertised burst  size, the receiver is
    // asking the network to provide flow pacing  beyond  what  would  be provided
    // under normal circumstances. Such pac- ing is at the network's option. The
    // type is interface{} with range: 0..2147483647. Units are bytes.
    Intsrvflowburst interface{}

    // The weight used  to  prioritize  the  traffic. Note  that the
    // interpretation of this object is implementation-specific,   as  
    // implementations vary in their use of weighting procedures. The type is
    // interface{} with range: -2147483648..2147483647.
    Intsrvflowweight interface{}

    // The number of the queue used by this  traffic. Note  that the
    // interpretation of this object is implementation-specific,   as  
    // implementations vary in their use of queue identifiers. The type is
    // interface{} with range: -2147483648..2147483647.
    Intsrvflowqueue interface{}

    // The minimum message size for  this  flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    Intsrvflowmintu interface{}

    // The maximum datagram size for this  flow  that will conform to the traffic
    // specification. This value cannot exceed the MTU of the interface. The type
    // is interface{} with range: 0..2147483647.
    Intsrvflowmaxtu interface{}

    // The number of packets that  were  remanded  to best effort service. The
    // type is interface{} with range: 0..4294967295.
    Intsrvflowbesteffort interface{}

    // The number of packets policed since the incep- tion of the flow's service.
    // The type is interface{} with range: 0..4294967295.
    Intsrvflowpoliced interface{}

    // If 'true', the flow  is  to  incur  loss  when traffic is policed.  If
    // 'false', policed traff- ic is treated as best effort traffic. The type is
    // bool.
    Intsrvflowdiscard interface{}

    // The QoS service being applied to this flow. The type is QosService.
    Intsrvflowservice interface{}

    // In the event of ambiguity, the order in  which the  classifier  should 
    // make  its comparisons. The row with intSrvFlowOrder=0 is tried  first, and 
    // comparisons  proceed  in  the order of in- creasing value.  Non-serial
    // implementations  of the classifier should emulate this behavior. The type
    // is interface{} with range: 0..65535.
    Intsrvfloworder interface{}

    // 'active' for all active  flows.   This  object may be used to install
    // static classifier infor- mation, delete classifier information,  or  au-
    // thorize such. The type is RowStatus.
    Intsrvflowstatus interface{}
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetFilter() yfilter.YFilter { return intsrvflowentry.YFilter }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) SetFilter(yf yfilter.YFilter) { intsrvflowentry.YFilter = yf }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetGoName(yname string) string {
    if yname == "intSrvFlowNumber" { return "Intsrvflownumber" }
    if yname == "intSrvFlowType" { return "Intsrvflowtype" }
    if yname == "intSrvFlowOwner" { return "Intsrvflowowner" }
    if yname == "intSrvFlowDestAddr" { return "Intsrvflowdestaddr" }
    if yname == "intSrvFlowSenderAddr" { return "Intsrvflowsenderaddr" }
    if yname == "intSrvFlowDestAddrLength" { return "Intsrvflowdestaddrlength" }
    if yname == "intSrvFlowSenderAddrLength" { return "Intsrvflowsenderaddrlength" }
    if yname == "intSrvFlowProtocol" { return "Intsrvflowprotocol" }
    if yname == "intSrvFlowDestPort" { return "Intsrvflowdestport" }
    if yname == "intSrvFlowPort" { return "Intsrvflowport" }
    if yname == "intSrvFlowFlowId" { return "Intsrvflowflowid" }
    if yname == "intSrvFlowInterface" { return "Intsrvflowinterface" }
    if yname == "intSrvFlowIfAddr" { return "Intsrvflowifaddr" }
    if yname == "intSrvFlowRate" { return "Intsrvflowrate" }
    if yname == "intSrvFlowBurst" { return "Intsrvflowburst" }
    if yname == "intSrvFlowWeight" { return "Intsrvflowweight" }
    if yname == "intSrvFlowQueue" { return "Intsrvflowqueue" }
    if yname == "intSrvFlowMinTU" { return "Intsrvflowmintu" }
    if yname == "intSrvFlowMaxTU" { return "Intsrvflowmaxtu" }
    if yname == "intSrvFlowBestEffort" { return "Intsrvflowbesteffort" }
    if yname == "intSrvFlowPoliced" { return "Intsrvflowpoliced" }
    if yname == "intSrvFlowDiscard" { return "Intsrvflowdiscard" }
    if yname == "intSrvFlowService" { return "Intsrvflowservice" }
    if yname == "intSrvFlowOrder" { return "Intsrvfloworder" }
    if yname == "intSrvFlowStatus" { return "Intsrvflowstatus" }
    return ""
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetSegmentPath() string {
    return "intSrvFlowEntry" + "[intSrvFlowNumber='" + fmt.Sprintf("%v", intsrvflowentry.Intsrvflownumber) + "']"
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intSrvFlowNumber"] = intsrvflowentry.Intsrvflownumber
    leafs["intSrvFlowType"] = intsrvflowentry.Intsrvflowtype
    leafs["intSrvFlowOwner"] = intsrvflowentry.Intsrvflowowner
    leafs["intSrvFlowDestAddr"] = intsrvflowentry.Intsrvflowdestaddr
    leafs["intSrvFlowSenderAddr"] = intsrvflowentry.Intsrvflowsenderaddr
    leafs["intSrvFlowDestAddrLength"] = intsrvflowentry.Intsrvflowdestaddrlength
    leafs["intSrvFlowSenderAddrLength"] = intsrvflowentry.Intsrvflowsenderaddrlength
    leafs["intSrvFlowProtocol"] = intsrvflowentry.Intsrvflowprotocol
    leafs["intSrvFlowDestPort"] = intsrvflowentry.Intsrvflowdestport
    leafs["intSrvFlowPort"] = intsrvflowentry.Intsrvflowport
    leafs["intSrvFlowFlowId"] = intsrvflowentry.Intsrvflowflowid
    leafs["intSrvFlowInterface"] = intsrvflowentry.Intsrvflowinterface
    leafs["intSrvFlowIfAddr"] = intsrvflowentry.Intsrvflowifaddr
    leafs["intSrvFlowRate"] = intsrvflowentry.Intsrvflowrate
    leafs["intSrvFlowBurst"] = intsrvflowentry.Intsrvflowburst
    leafs["intSrvFlowWeight"] = intsrvflowentry.Intsrvflowweight
    leafs["intSrvFlowQueue"] = intsrvflowentry.Intsrvflowqueue
    leafs["intSrvFlowMinTU"] = intsrvflowentry.Intsrvflowmintu
    leafs["intSrvFlowMaxTU"] = intsrvflowentry.Intsrvflowmaxtu
    leafs["intSrvFlowBestEffort"] = intsrvflowentry.Intsrvflowbesteffort
    leafs["intSrvFlowPoliced"] = intsrvflowentry.Intsrvflowpoliced
    leafs["intSrvFlowDiscard"] = intsrvflowentry.Intsrvflowdiscard
    leafs["intSrvFlowService"] = intsrvflowentry.Intsrvflowservice
    leafs["intSrvFlowOrder"] = intsrvflowentry.Intsrvfloworder
    leafs["intSrvFlowStatus"] = intsrvflowentry.Intsrvflowstatus
    return leafs
}

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetYangName() string { return "intSrvFlowEntry" }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) SetParent(parent types.Entity) { intsrvflowentry.parent = parent }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetParent() types.Entity { return intsrvflowentry.parent }

func (intsrvflowentry *INTSERVMIB_Intsrvflowtable_Intsrvflowentry) GetParentYangName() string { return "intSrvFlowTable" }

// INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner represents queue policy database.
type INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner string

const (
    INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_other INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "other"

    INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_rsvp INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "rsvp"

    INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_management INTSERVMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "management"
)

