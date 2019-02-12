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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    IntSrvGenObjects INTSERVMIB_IntSrvGenObjects

    // The reservable attributes of the system's  in- terfaces.
    IntSrvIfAttribTable INTSERVMIB_IntSrvIfAttribTable

    // Information describing the reserved flows  us- ing the system's interfaces.
    IntSrvFlowTable INTSERVMIB_IntSrvFlowTable
}

func (iNTSERVMIB *INTSERVMIB) GetEntityData() *types.CommonEntityData {
    iNTSERVMIB.EntityData.YFilter = iNTSERVMIB.YFilter
    iNTSERVMIB.EntityData.YangName = "INT-SERV-MIB"
    iNTSERVMIB.EntityData.BundleName = "cisco_ios_xe"
    iNTSERVMIB.EntityData.ParentYangName = "INT-SERV-MIB"
    iNTSERVMIB.EntityData.SegmentPath = "INT-SERV-MIB:INT-SERV-MIB"
    iNTSERVMIB.EntityData.AbsolutePath = iNTSERVMIB.EntityData.SegmentPath
    iNTSERVMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iNTSERVMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iNTSERVMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iNTSERVMIB.EntityData.Children = types.NewOrderedMap()
    iNTSERVMIB.EntityData.Children.Append("intSrvGenObjects", types.YChild{"IntSrvGenObjects", &iNTSERVMIB.IntSrvGenObjects})
    iNTSERVMIB.EntityData.Children.Append("intSrvIfAttribTable", types.YChild{"IntSrvIfAttribTable", &iNTSERVMIB.IntSrvIfAttribTable})
    iNTSERVMIB.EntityData.Children.Append("intSrvFlowTable", types.YChild{"IntSrvFlowTable", &iNTSERVMIB.IntSrvFlowTable})
    iNTSERVMIB.EntityData.Leafs = types.NewOrderedMap()

    iNTSERVMIB.EntityData.YListKeys = []string {}

    return &(iNTSERVMIB.EntityData)
}

// INTSERVMIB_IntSrvGenObjects
type INTSERVMIB_IntSrvGenObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This  object  is  used  to  assign  values  to intSrvFlowNumber  as
    // described in 'Textual Con- ventions  for  SNMPv2'.   The  network  manager
    // reads  the  object,  and  then writes the value back in the SET that
    // creates a new instance  of intSrvFlowEntry.   If  the  SET  fails with the
    // code 'inconsistentValue', then the process must be  repeated; If the SET
    // succeeds, then the ob- ject is incremented, and the  new  instance  is
    // created according to the manager's directions. The type is interface{} with
    // range: 0..2147483647.
    IntSrvFlowNewIndex interface{}
}

func (intSrvGenObjects *INTSERVMIB_IntSrvGenObjects) GetEntityData() *types.CommonEntityData {
    intSrvGenObjects.EntityData.YFilter = intSrvGenObjects.YFilter
    intSrvGenObjects.EntityData.YangName = "intSrvGenObjects"
    intSrvGenObjects.EntityData.BundleName = "cisco_ios_xe"
    intSrvGenObjects.EntityData.ParentYangName = "INT-SERV-MIB"
    intSrvGenObjects.EntityData.SegmentPath = "intSrvGenObjects"
    intSrvGenObjects.EntityData.AbsolutePath = "INT-SERV-MIB:INT-SERV-MIB/" + intSrvGenObjects.EntityData.SegmentPath
    intSrvGenObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intSrvGenObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intSrvGenObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intSrvGenObjects.EntityData.Children = types.NewOrderedMap()
    intSrvGenObjects.EntityData.Leafs = types.NewOrderedMap()
    intSrvGenObjects.EntityData.Leafs.Append("intSrvFlowNewIndex", types.YLeaf{"IntSrvFlowNewIndex", intSrvGenObjects.IntSrvFlowNewIndex})

    intSrvGenObjects.EntityData.YListKeys = []string {}

    return &(intSrvGenObjects.EntityData)
}

// INTSERVMIB_IntSrvIfAttribTable
// The reservable attributes of the system's  in-
// terfaces.
type INTSERVMIB_IntSrvIfAttribTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The reservable attributes of  a  given  inter- face. The type is slice of
    // INTSERVMIB_IntSrvIfAttribTable_IntSrvIfAttribEntry.
    IntSrvIfAttribEntry []*INTSERVMIB_IntSrvIfAttribTable_IntSrvIfAttribEntry
}

func (intSrvIfAttribTable *INTSERVMIB_IntSrvIfAttribTable) GetEntityData() *types.CommonEntityData {
    intSrvIfAttribTable.EntityData.YFilter = intSrvIfAttribTable.YFilter
    intSrvIfAttribTable.EntityData.YangName = "intSrvIfAttribTable"
    intSrvIfAttribTable.EntityData.BundleName = "cisco_ios_xe"
    intSrvIfAttribTable.EntityData.ParentYangName = "INT-SERV-MIB"
    intSrvIfAttribTable.EntityData.SegmentPath = "intSrvIfAttribTable"
    intSrvIfAttribTable.EntityData.AbsolutePath = "INT-SERV-MIB:INT-SERV-MIB/" + intSrvIfAttribTable.EntityData.SegmentPath
    intSrvIfAttribTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intSrvIfAttribTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intSrvIfAttribTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intSrvIfAttribTable.EntityData.Children = types.NewOrderedMap()
    intSrvIfAttribTable.EntityData.Children.Append("intSrvIfAttribEntry", types.YChild{"IntSrvIfAttribEntry", nil})
    for i := range intSrvIfAttribTable.IntSrvIfAttribEntry {
        intSrvIfAttribTable.EntityData.Children.Append(types.GetSegmentPath(intSrvIfAttribTable.IntSrvIfAttribEntry[i]), types.YChild{"IntSrvIfAttribEntry", intSrvIfAttribTable.IntSrvIfAttribEntry[i]})
    }
    intSrvIfAttribTable.EntityData.Leafs = types.NewOrderedMap()

    intSrvIfAttribTable.EntityData.YListKeys = []string {}

    return &(intSrvIfAttribTable.EntityData)
}

// INTSERVMIB_IntSrvIfAttribTable_IntSrvIfAttribEntry
// The reservable attributes of  a  given  inter-
// face.
type INTSERVMIB_IntSrvIfAttribTable_IntSrvIfAttribEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of bits/second currently  allocated to reserved sessions on the
    // interface. The type is interface{} with range: 0..2147483647. Units are
    // Bits per second.
    IntSrvIfAttribAllocatedBits interface{}

    // The maximum number of bits/second that may  be allocated  to  reserved 
    // sessions on the inter- face. The type is interface{} with range:
    // 0..2147483647. Units are Bits per second.
    IntSrvIfAttribMaxAllocatedBits interface{}

    // The amount of buffer space  required  to  hold the simultaneous burst of
    // all reserved flows on the interface. The type is interface{} with range:
    // 0..2147483647. Units are Bytes.
    IntSrvIfAttribAllocatedBuffer interface{}

    // The number of reserved flows currently  active on  this  interface.  A flow
    // can be created ei- ther from a reservation protocol (such as  RSVP or
    // ST-II) or via configuration information. The type is interface{} with
    // range: 0..4294967295.
    IntSrvIfAttribFlows interface{}

    // The amount of propagation delay that this  in- terface  introduces  in
    // addition to that intro- diced by bit propagation delays. The type is
    // interface{} with range: -2147483648..2147483647. Units are microseconds.
    IntSrvIfAttribPropagationDelay interface{}

    // 'active' on interfaces that are configured for RSVP. The type is RowStatus.
    IntSrvIfAttribStatus interface{}
}

func (intSrvIfAttribEntry *INTSERVMIB_IntSrvIfAttribTable_IntSrvIfAttribEntry) GetEntityData() *types.CommonEntityData {
    intSrvIfAttribEntry.EntityData.YFilter = intSrvIfAttribEntry.YFilter
    intSrvIfAttribEntry.EntityData.YangName = "intSrvIfAttribEntry"
    intSrvIfAttribEntry.EntityData.BundleName = "cisco_ios_xe"
    intSrvIfAttribEntry.EntityData.ParentYangName = "intSrvIfAttribTable"
    intSrvIfAttribEntry.EntityData.SegmentPath = "intSrvIfAttribEntry" + types.AddKeyToken(intSrvIfAttribEntry.IfIndex, "ifIndex")
    intSrvIfAttribEntry.EntityData.AbsolutePath = "INT-SERV-MIB:INT-SERV-MIB/intSrvIfAttribTable/" + intSrvIfAttribEntry.EntityData.SegmentPath
    intSrvIfAttribEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intSrvIfAttribEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intSrvIfAttribEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intSrvIfAttribEntry.EntityData.Children = types.NewOrderedMap()
    intSrvIfAttribEntry.EntityData.Leafs = types.NewOrderedMap()
    intSrvIfAttribEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", intSrvIfAttribEntry.IfIndex})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribAllocatedBits", types.YLeaf{"IntSrvIfAttribAllocatedBits", intSrvIfAttribEntry.IntSrvIfAttribAllocatedBits})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribMaxAllocatedBits", types.YLeaf{"IntSrvIfAttribMaxAllocatedBits", intSrvIfAttribEntry.IntSrvIfAttribMaxAllocatedBits})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribAllocatedBuffer", types.YLeaf{"IntSrvIfAttribAllocatedBuffer", intSrvIfAttribEntry.IntSrvIfAttribAllocatedBuffer})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribFlows", types.YLeaf{"IntSrvIfAttribFlows", intSrvIfAttribEntry.IntSrvIfAttribFlows})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribPropagationDelay", types.YLeaf{"IntSrvIfAttribPropagationDelay", intSrvIfAttribEntry.IntSrvIfAttribPropagationDelay})
    intSrvIfAttribEntry.EntityData.Leafs.Append("intSrvIfAttribStatus", types.YLeaf{"IntSrvIfAttribStatus", intSrvIfAttribEntry.IntSrvIfAttribStatus})

    intSrvIfAttribEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(intSrvIfAttribEntry.EntityData)
}

// INTSERVMIB_IntSrvFlowTable
// Information describing the reserved flows  us-
// ing the system's interfaces.
type INTSERVMIB_IntSrvFlowTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information describing the use of a given  in- terface   by   a   given  
    // flow.   The  counter intSrvFlowPoliced starts counting  at  the  in-
    // stallation of the flow. The type is slice of
    // INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry.
    IntSrvFlowEntry []*INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry
}

func (intSrvFlowTable *INTSERVMIB_IntSrvFlowTable) GetEntityData() *types.CommonEntityData {
    intSrvFlowTable.EntityData.YFilter = intSrvFlowTable.YFilter
    intSrvFlowTable.EntityData.YangName = "intSrvFlowTable"
    intSrvFlowTable.EntityData.BundleName = "cisco_ios_xe"
    intSrvFlowTable.EntityData.ParentYangName = "INT-SERV-MIB"
    intSrvFlowTable.EntityData.SegmentPath = "intSrvFlowTable"
    intSrvFlowTable.EntityData.AbsolutePath = "INT-SERV-MIB:INT-SERV-MIB/" + intSrvFlowTable.EntityData.SegmentPath
    intSrvFlowTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intSrvFlowTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intSrvFlowTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intSrvFlowTable.EntityData.Children = types.NewOrderedMap()
    intSrvFlowTable.EntityData.Children.Append("intSrvFlowEntry", types.YChild{"IntSrvFlowEntry", nil})
    for i := range intSrvFlowTable.IntSrvFlowEntry {
        intSrvFlowTable.EntityData.Children.Append(types.GetSegmentPath(intSrvFlowTable.IntSrvFlowEntry[i]), types.YChild{"IntSrvFlowEntry", intSrvFlowTable.IntSrvFlowEntry[i]})
    }
    intSrvFlowTable.EntityData.Leafs = types.NewOrderedMap()

    intSrvFlowTable.EntityData.YListKeys = []string {}

    return &(intSrvFlowTable.EntityData)
}

// INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry
// Information describing the use of a given  in-
// terface   by   a   given   flow.   The  counter
// intSrvFlowPoliced starts counting  at  the  in-
// stallation of the flow.
type INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The number of this flow.  This is for SNMP In-
    // dexing purposes only and has no relation to any protocol value. The type is
    // interface{} with range: 0..2147483647.
    IntSrvFlowNumber interface{}

    // The type of session (IP4, IP6, IP6  with  flow information, etc). The type
    // is interface{} with range: 1..255.
    IntSrvFlowType interface{}

    // The process that installed this  flow  in  the queue policy database. The
    // type is IntSrvFlowOwner.
    IntSrvFlowOwner interface{}

    // The destination address used by all senders in this  session.   This object
    // may not be changed when the value of the RowStatus object is  'ac- tive'.
    // The type is string with length: 4..16.
    IntSrvFlowDestAddr interface{}

    // The source address of the sender  selected  by this  reservation.  The
    // value of all zeroes in- dicates 'all senders'.  This object may not  be
    // changed  when the value of the RowStatus object is 'active'. The type is
    // string with length: 4..16.
    IntSrvFlowSenderAddr interface{}

    // The length of the destination address in bits. This  is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    IntSrvFlowDestAddrLength interface{}

    // The length of the sender's  address  in  bits. This  is  the CIDR Prefix
    // Length, which for IP4 hosts and multicast addresses is 32 bits.  This
    // object may not be changed when the value of the RowStatus object is
    // 'active'. The type is interface{} with range: 0..128.
    IntSrvFlowSenderAddrLength interface{}

    // The IP Protocol used by a session.   This  ob- ject  may  not be changed
    // when the value of the RowStatus object is 'active'. The type is interface{}
    // with range: 1..255.
    IntSrvFlowProtocol interface{}

    // The UDP or TCP port number used as a  destina- tion  port for all senders
    // in this session.  If the  IP   protocol   in   use,   specified   by
    // intSrvResvFwdProtocol,  is 50 (ESP) or 51 (AH), this  represents  a 
    // virtual  destination  port number.   A value of zero indicates that the IP
    // protocol in use does not have ports.  This  ob- ject  may  not be changed
    // when the value of the RowStatus object is 'active'. The type is string with
    // length: 2..4.
    IntSrvFlowDestPort interface{}

    // The UDP or TCP port number used  as  a  source port  for  this sender in
    // this session.  If the IP    protocol    in    use,    specified    by
    // intSrvResvFwdProtocol  is  50 (ESP) or 51 (AH), this represents a
    // generalized  port  identifier (GPI).   A  value of zero indicates that the
    // IP protocol in use does not have ports.  This  ob- ject  may  not be
    // changed when the value of the RowStatus object is 'active'. The type is
    // string with length: 2..4.
    IntSrvFlowPort interface{}

    // The flow ID that  this  sender  is  using,  if this  is  an IPv6 session.
    // The type is interface{} with range: 0..16777215.
    IntSrvFlowFlowId interface{}

    // The ifIndex value of the  interface  on  which this reservation exists. The
    // type is interface{} with range: 1..2147483647.
    IntSrvFlowInterface interface{}

    // The IP Address on the ifEntry  on  which  this reservation  exists.  This
    // is present primarily to support those interfaces which layer  multi- ple IP
    // Addresses on the interface. The type is string with length: 4..16.
    IntSrvFlowIfAddr interface{}

    // The Reserved Rate of the sender's data stream. If this is a Controlled Load
    // service flow, this rate is derived from the Tspec  rate  parameter (r).  
    // If  this  is  a Guaranteed service flow, this rate is derived from  the 
    // Rspec  clearing rate parameter (R). The type is interface{} with range:
    // 0..2147483647. Units are bits per second.
    IntSrvFlowRate interface{}

    // The size of the largest  burst  expected  from the sender at a time.  If
    // this is less than  the  sender's  advertised burst  size, the receiver is
    // asking the network to provide flow pacing  beyond  what  would  be provided
    // under normal circumstances. Such pac- ing is at the network's option. The
    // type is interface{} with range: 0..2147483647. Units are bytes.
    IntSrvFlowBurst interface{}

    // The weight used  to  prioritize  the  traffic. Note  that the
    // interpretation of this object is implementation-specific,   as  
    // implementations vary in their use of weighting procedures. The type is
    // interface{} with range: -2147483648..2147483647.
    IntSrvFlowWeight interface{}

    // The number of the queue used by this  traffic. Note  that the
    // interpretation of this object is implementation-specific,   as  
    // implementations vary in their use of queue identifiers. The type is
    // interface{} with range: -2147483648..2147483647.
    IntSrvFlowQueue interface{}

    // The minimum message size for  this  flow.  The policing  algorithm will
    // treat smaller messages as though they are this size. The type is
    // interface{} with range: 0..2147483647.
    IntSrvFlowMinTU interface{}

    // The maximum datagram size for this  flow  that will conform to the traffic
    // specification. This value cannot exceed the MTU of the interface. The type
    // is interface{} with range: 0..2147483647.
    IntSrvFlowMaxTU interface{}

    // The number of packets that  were  remanded  to best effort service. The
    // type is interface{} with range: 0..4294967295.
    IntSrvFlowBestEffort interface{}

    // The number of packets policed since the incep- tion of the flow's service.
    // The type is interface{} with range: 0..4294967295.
    IntSrvFlowPoliced interface{}

    // If 'true', the flow  is  to  incur  loss  when traffic is policed.  If
    // 'false', policed traff- ic is treated as best effort traffic. The type is
    // bool.
    IntSrvFlowDiscard interface{}

    // The QoS service being applied to this flow. The type is QosService.
    IntSrvFlowService interface{}

    // In the event of ambiguity, the order in  which the  classifier  should 
    // make  its comparisons. The row with intSrvFlowOrder=0 is tried  first, and 
    // comparisons  proceed  in  the order of in- creasing value.  Non-serial
    // implementations  of the classifier should emulate this behavior. The type
    // is interface{} with range: 0..65535.
    IntSrvFlowOrder interface{}

    // 'active' for all active  flows.   This  object may be used to install
    // static classifier infor- mation, delete classifier information,  or  au-
    // thorize such. The type is RowStatus.
    IntSrvFlowStatus interface{}
}

func (intSrvFlowEntry *INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry) GetEntityData() *types.CommonEntityData {
    intSrvFlowEntry.EntityData.YFilter = intSrvFlowEntry.YFilter
    intSrvFlowEntry.EntityData.YangName = "intSrvFlowEntry"
    intSrvFlowEntry.EntityData.BundleName = "cisco_ios_xe"
    intSrvFlowEntry.EntityData.ParentYangName = "intSrvFlowTable"
    intSrvFlowEntry.EntityData.SegmentPath = "intSrvFlowEntry" + types.AddKeyToken(intSrvFlowEntry.IntSrvFlowNumber, "intSrvFlowNumber")
    intSrvFlowEntry.EntityData.AbsolutePath = "INT-SERV-MIB:INT-SERV-MIB/intSrvFlowTable/" + intSrvFlowEntry.EntityData.SegmentPath
    intSrvFlowEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intSrvFlowEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intSrvFlowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intSrvFlowEntry.EntityData.Children = types.NewOrderedMap()
    intSrvFlowEntry.EntityData.Leafs = types.NewOrderedMap()
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowNumber", types.YLeaf{"IntSrvFlowNumber", intSrvFlowEntry.IntSrvFlowNumber})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowType", types.YLeaf{"IntSrvFlowType", intSrvFlowEntry.IntSrvFlowType})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowOwner", types.YLeaf{"IntSrvFlowOwner", intSrvFlowEntry.IntSrvFlowOwner})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowDestAddr", types.YLeaf{"IntSrvFlowDestAddr", intSrvFlowEntry.IntSrvFlowDestAddr})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowSenderAddr", types.YLeaf{"IntSrvFlowSenderAddr", intSrvFlowEntry.IntSrvFlowSenderAddr})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowDestAddrLength", types.YLeaf{"IntSrvFlowDestAddrLength", intSrvFlowEntry.IntSrvFlowDestAddrLength})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowSenderAddrLength", types.YLeaf{"IntSrvFlowSenderAddrLength", intSrvFlowEntry.IntSrvFlowSenderAddrLength})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowProtocol", types.YLeaf{"IntSrvFlowProtocol", intSrvFlowEntry.IntSrvFlowProtocol})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowDestPort", types.YLeaf{"IntSrvFlowDestPort", intSrvFlowEntry.IntSrvFlowDestPort})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowPort", types.YLeaf{"IntSrvFlowPort", intSrvFlowEntry.IntSrvFlowPort})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowFlowId", types.YLeaf{"IntSrvFlowFlowId", intSrvFlowEntry.IntSrvFlowFlowId})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowInterface", types.YLeaf{"IntSrvFlowInterface", intSrvFlowEntry.IntSrvFlowInterface})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowIfAddr", types.YLeaf{"IntSrvFlowIfAddr", intSrvFlowEntry.IntSrvFlowIfAddr})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowRate", types.YLeaf{"IntSrvFlowRate", intSrvFlowEntry.IntSrvFlowRate})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowBurst", types.YLeaf{"IntSrvFlowBurst", intSrvFlowEntry.IntSrvFlowBurst})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowWeight", types.YLeaf{"IntSrvFlowWeight", intSrvFlowEntry.IntSrvFlowWeight})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowQueue", types.YLeaf{"IntSrvFlowQueue", intSrvFlowEntry.IntSrvFlowQueue})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowMinTU", types.YLeaf{"IntSrvFlowMinTU", intSrvFlowEntry.IntSrvFlowMinTU})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowMaxTU", types.YLeaf{"IntSrvFlowMaxTU", intSrvFlowEntry.IntSrvFlowMaxTU})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowBestEffort", types.YLeaf{"IntSrvFlowBestEffort", intSrvFlowEntry.IntSrvFlowBestEffort})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowPoliced", types.YLeaf{"IntSrvFlowPoliced", intSrvFlowEntry.IntSrvFlowPoliced})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowDiscard", types.YLeaf{"IntSrvFlowDiscard", intSrvFlowEntry.IntSrvFlowDiscard})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowService", types.YLeaf{"IntSrvFlowService", intSrvFlowEntry.IntSrvFlowService})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowOrder", types.YLeaf{"IntSrvFlowOrder", intSrvFlowEntry.IntSrvFlowOrder})
    intSrvFlowEntry.EntityData.Leafs.Append("intSrvFlowStatus", types.YLeaf{"IntSrvFlowStatus", intSrvFlowEntry.IntSrvFlowStatus})

    intSrvFlowEntry.EntityData.YListKeys = []string {"IntSrvFlowNumber"}

    return &(intSrvFlowEntry.EntityData)
}

// INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner represents queue policy database.
type INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner string

const (
    INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner_other INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner = "other"

    INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner_rsvp INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner = "rsvp"

    INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner_management INTSERVMIB_IntSrvFlowTable_IntSrvFlowEntry_IntSrvFlowOwner = "management"
)

