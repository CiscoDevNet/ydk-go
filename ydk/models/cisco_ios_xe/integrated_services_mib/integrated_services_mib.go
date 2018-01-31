// The MIB module to describe the Integrated Services
// Protocol
package integrated_services_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package integrated_services_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:INTEGRATED-SERVICES-MIB INTEGRATED-SERVICES-MIB}", reflect.TypeOf(INTEGRATEDSERVICESMIB{}))
    ydk.RegisterEntity("INTEGRATED-SERVICES-MIB:INTEGRATED-SERVICES-MIB", reflect.TypeOf(INTEGRATEDSERVICESMIB{}))
}

// QosService represents The class of service in use by a flow.
type QosService string

const (
    QosService_bestEffort QosService = "bestEffort"

    QosService_guaranteedDelay QosService = "guaranteedDelay"

    QosService_controlledLoad QosService = "controlledLoad"
)

// INTEGRATEDSERVICESMIB
type INTEGRATEDSERVICESMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Intsrvgenobjects INTEGRATEDSERVICESMIB_Intsrvgenobjects

    // The reservable attributes of the system's  in- terfaces.
    Intsrvifattribtable INTEGRATEDSERVICESMIB_Intsrvifattribtable

    // Information describing the reserved flows  us- ing the system's interfaces.
    Intsrvflowtable INTEGRATEDSERVICESMIB_Intsrvflowtable
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetFilter() yfilter.YFilter { return iNTEGRATEDSERVICESMIB.YFilter }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) SetFilter(yf yfilter.YFilter) { iNTEGRATEDSERVICESMIB.YFilter = yf }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetGoName(yname string) string {
    if yname == "intSrvGenObjects" { return "Intsrvgenobjects" }
    if yname == "intSrvIfAttribTable" { return "Intsrvifattribtable" }
    if yname == "intSrvFlowTable" { return "Intsrvflowtable" }
    return ""
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetSegmentPath() string {
    return "INTEGRATED-SERVICES-MIB:INTEGRATED-SERVICES-MIB"
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvGenObjects" {
        return &iNTEGRATEDSERVICESMIB.Intsrvgenobjects
    }
    if childYangName == "intSrvIfAttribTable" {
        return &iNTEGRATEDSERVICESMIB.Intsrvifattribtable
    }
    if childYangName == "intSrvFlowTable" {
        return &iNTEGRATEDSERVICESMIB.Intsrvflowtable
    }
    return nil
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intSrvGenObjects"] = &iNTEGRATEDSERVICESMIB.Intsrvgenobjects
    children["intSrvIfAttribTable"] = &iNTEGRATEDSERVICESMIB.Intsrvifattribtable
    children["intSrvFlowTable"] = &iNTEGRATEDSERVICESMIB.Intsrvflowtable
    return children
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetYangName() string { return "INTEGRATED-SERVICES-MIB" }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) SetParent(parent types.Entity) { iNTEGRATEDSERVICESMIB.parent = parent }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetParent() types.Entity { return iNTEGRATEDSERVICESMIB.parent }

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetParentYangName() string { return "INTEGRATED-SERVICES-MIB" }

// INTEGRATEDSERVICESMIB_Intsrvgenobjects
type INTEGRATEDSERVICESMIB_Intsrvgenobjects struct {
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

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetFilter() yfilter.YFilter { return intsrvgenobjects.YFilter }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) SetFilter(yf yfilter.YFilter) { intsrvgenobjects.YFilter = yf }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetGoName(yname string) string {
    if yname == "intSrvFlowNewIndex" { return "Intsrvflownewindex" }
    return ""
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetSegmentPath() string {
    return "intSrvGenObjects"
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intSrvFlowNewIndex"] = intsrvgenobjects.Intsrvflownewindex
    return leafs
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetYangName() string { return "intSrvGenObjects" }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) SetParent(parent types.Entity) { intsrvgenobjects.parent = parent }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetParent() types.Entity { return intsrvgenobjects.parent }

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetParentYangName() string { return "INTEGRATED-SERVICES-MIB" }

// INTEGRATEDSERVICESMIB_Intsrvifattribtable
// The reservable attributes of the system's  in-
// terfaces.
type INTEGRATEDSERVICESMIB_Intsrvifattribtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The reservable attributes of  a  given  inter- face. The type is slice of
    // INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry.
    Intsrvifattribentry []INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetFilter() yfilter.YFilter { return intsrvifattribtable.YFilter }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) SetFilter(yf yfilter.YFilter) { intsrvifattribtable.YFilter = yf }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetGoName(yname string) string {
    if yname == "intSrvIfAttribEntry" { return "Intsrvifattribentry" }
    return ""
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetSegmentPath() string {
    return "intSrvIfAttribTable"
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvIfAttribEntry" {
        for _, c := range intsrvifattribtable.Intsrvifattribentry {
            if intsrvifattribtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry{}
        intsrvifattribtable.Intsrvifattribentry = append(intsrvifattribtable.Intsrvifattribentry, child)
        return &intsrvifattribtable.Intsrvifattribentry[len(intsrvifattribtable.Intsrvifattribentry)-1]
    }
    return nil
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intsrvifattribtable.Intsrvifattribentry {
        children[intsrvifattribtable.Intsrvifattribentry[i].GetSegmentPath()] = &intsrvifattribtable.Intsrvifattribentry[i]
    }
    return children
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetYangName() string { return "intSrvIfAttribTable" }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) SetParent(parent types.Entity) { intsrvifattribtable.parent = parent }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetParent() types.Entity { return intsrvifattribtable.parent }

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetParentYangName() string { return "INTEGRATED-SERVICES-MIB" }

// INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry
// The reservable attributes of  a  given  inter-
// face.
type INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry struct {
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

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetFilter() yfilter.YFilter { return intsrvifattribentry.YFilter }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) SetFilter(yf yfilter.YFilter) { intsrvifattribentry.YFilter = yf }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "intSrvIfAttribAllocatedBits" { return "Intsrvifattriballocatedbits" }
    if yname == "intSrvIfAttribMaxAllocatedBits" { return "Intsrvifattribmaxallocatedbits" }
    if yname == "intSrvIfAttribAllocatedBuffer" { return "Intsrvifattriballocatedbuffer" }
    if yname == "intSrvIfAttribFlows" { return "Intsrvifattribflows" }
    if yname == "intSrvIfAttribPropagationDelay" { return "Intsrvifattribpropagationdelay" }
    if yname == "intSrvIfAttribStatus" { return "Intsrvifattribstatus" }
    return ""
}

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetSegmentPath() string {
    return "intSrvIfAttribEntry" + "[ifIndex='" + fmt.Sprintf("%v", intsrvifattribentry.Ifindex) + "']"
}

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetLeafs() map[string]interface{} {
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

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetYangName() string { return "intSrvIfAttribEntry" }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) SetParent(parent types.Entity) { intsrvifattribentry.parent = parent }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetParent() types.Entity { return intsrvifattribentry.parent }

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetParentYangName() string { return "intSrvIfAttribTable" }

// INTEGRATEDSERVICESMIB_Intsrvflowtable
// Information describing the reserved flows  us-
// ing the system's interfaces.
type INTEGRATEDSERVICESMIB_Intsrvflowtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information describing the use of a given  in- terface   by   a   given  
    // flow.   The  counter intSrvFlowPoliced starts counting  at  the  in-
    // stallation of the flow. The type is slice of
    // INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry.
    Intsrvflowentry []INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetFilter() yfilter.YFilter { return intsrvflowtable.YFilter }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) SetFilter(yf yfilter.YFilter) { intsrvflowtable.YFilter = yf }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetGoName(yname string) string {
    if yname == "intSrvFlowEntry" { return "Intsrvflowentry" }
    return ""
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetSegmentPath() string {
    return "intSrvFlowTable"
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intSrvFlowEntry" {
        for _, c := range intsrvflowtable.Intsrvflowentry {
            if intsrvflowtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry{}
        intsrvflowtable.Intsrvflowentry = append(intsrvflowtable.Intsrvflowentry, child)
        return &intsrvflowtable.Intsrvflowentry[len(intsrvflowtable.Intsrvflowentry)-1]
    }
    return nil
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intsrvflowtable.Intsrvflowentry {
        children[intsrvflowtable.Intsrvflowentry[i].GetSegmentPath()] = &intsrvflowtable.Intsrvflowentry[i]
    }
    return children
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetYangName() string { return "intSrvFlowTable" }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) SetParent(parent types.Entity) { intsrvflowtable.parent = parent }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetParent() types.Entity { return intsrvflowtable.parent }

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetParentYangName() string { return "INTEGRATED-SERVICES-MIB" }

// INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry
// Information describing the use of a given  in-
// terface   by   a   given   flow.   The  counter
// intSrvFlowPoliced starts counting  at  the  in-
// stallation of the flow.
type INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry struct {
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

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetFilter() yfilter.YFilter { return intsrvflowentry.YFilter }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) SetFilter(yf yfilter.YFilter) { intsrvflowentry.YFilter = yf }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetGoName(yname string) string {
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

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetSegmentPath() string {
    return "intSrvFlowEntry" + "[intSrvFlowNumber='" + fmt.Sprintf("%v", intsrvflowentry.Intsrvflownumber) + "']"
}

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetLeafs() map[string]interface{} {
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

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetBundleName() string { return "cisco_ios_xe" }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetYangName() string { return "intSrvFlowEntry" }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) SetParent(parent types.Entity) { intsrvflowentry.parent = parent }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetParent() types.Entity { return intsrvflowentry.parent }

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetParentYangName() string { return "intSrvFlowTable" }

// INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner represents queue policy database.
type INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner string

const (
    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_other INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "other"

    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_rsvp INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "rsvp"

    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_management INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "management"
)

