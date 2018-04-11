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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Intsrvgenobjects INTEGRATEDSERVICESMIB_Intsrvgenobjects

    // The reservable attributes of the system's  in- terfaces.
    Intsrvifattribtable INTEGRATEDSERVICESMIB_Intsrvifattribtable

    // Information describing the reserved flows  us- ing the system's interfaces.
    Intsrvflowtable INTEGRATEDSERVICESMIB_Intsrvflowtable
}

func (iNTEGRATEDSERVICESMIB *INTEGRATEDSERVICESMIB) GetEntityData() *types.CommonEntityData {
    iNTEGRATEDSERVICESMIB.EntityData.YFilter = iNTEGRATEDSERVICESMIB.YFilter
    iNTEGRATEDSERVICESMIB.EntityData.YangName = "INTEGRATED-SERVICES-MIB"
    iNTEGRATEDSERVICESMIB.EntityData.BundleName = "cisco_ios_xe"
    iNTEGRATEDSERVICESMIB.EntityData.ParentYangName = "INTEGRATED-SERVICES-MIB"
    iNTEGRATEDSERVICESMIB.EntityData.SegmentPath = "INTEGRATED-SERVICES-MIB:INTEGRATED-SERVICES-MIB"
    iNTEGRATEDSERVICESMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iNTEGRATEDSERVICESMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iNTEGRATEDSERVICESMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iNTEGRATEDSERVICESMIB.EntityData.Children = make(map[string]types.YChild)
    iNTEGRATEDSERVICESMIB.EntityData.Children["intSrvGenObjects"] = types.YChild{"Intsrvgenobjects", &iNTEGRATEDSERVICESMIB.Intsrvgenobjects}
    iNTEGRATEDSERVICESMIB.EntityData.Children["intSrvIfAttribTable"] = types.YChild{"Intsrvifattribtable", &iNTEGRATEDSERVICESMIB.Intsrvifattribtable}
    iNTEGRATEDSERVICESMIB.EntityData.Children["intSrvFlowTable"] = types.YChild{"Intsrvflowtable", &iNTEGRATEDSERVICESMIB.Intsrvflowtable}
    iNTEGRATEDSERVICESMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iNTEGRATEDSERVICESMIB.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvgenobjects
type INTEGRATEDSERVICESMIB_Intsrvgenobjects struct {
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
    Intsrvflownewindex interface{}
}

func (intsrvgenobjects *INTEGRATEDSERVICESMIB_Intsrvgenobjects) GetEntityData() *types.CommonEntityData {
    intsrvgenobjects.EntityData.YFilter = intsrvgenobjects.YFilter
    intsrvgenobjects.EntityData.YangName = "intSrvGenObjects"
    intsrvgenobjects.EntityData.BundleName = "cisco_ios_xe"
    intsrvgenobjects.EntityData.ParentYangName = "INTEGRATED-SERVICES-MIB"
    intsrvgenobjects.EntityData.SegmentPath = "intSrvGenObjects"
    intsrvgenobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intsrvgenobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intsrvgenobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intsrvgenobjects.EntityData.Children = make(map[string]types.YChild)
    intsrvgenobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    intsrvgenobjects.EntityData.Leafs["intSrvFlowNewIndex"] = types.YLeaf{"Intsrvflownewindex", intsrvgenobjects.Intsrvflownewindex}
    return &(intsrvgenobjects.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvifattribtable
// The reservable attributes of the system's  in-
// terfaces.
type INTEGRATEDSERVICESMIB_Intsrvifattribtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The reservable attributes of  a  given  inter- face. The type is slice of
    // INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry.
    Intsrvifattribentry []INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry
}

func (intsrvifattribtable *INTEGRATEDSERVICESMIB_Intsrvifattribtable) GetEntityData() *types.CommonEntityData {
    intsrvifattribtable.EntityData.YFilter = intsrvifattribtable.YFilter
    intsrvifattribtable.EntityData.YangName = "intSrvIfAttribTable"
    intsrvifattribtable.EntityData.BundleName = "cisco_ios_xe"
    intsrvifattribtable.EntityData.ParentYangName = "INTEGRATED-SERVICES-MIB"
    intsrvifattribtable.EntityData.SegmentPath = "intSrvIfAttribTable"
    intsrvifattribtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intsrvifattribtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intsrvifattribtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intsrvifattribtable.EntityData.Children = make(map[string]types.YChild)
    intsrvifattribtable.EntityData.Children["intSrvIfAttribEntry"] = types.YChild{"Intsrvifattribentry", nil}
    for i := range intsrvifattribtable.Intsrvifattribentry {
        intsrvifattribtable.EntityData.Children[types.GetSegmentPath(&intsrvifattribtable.Intsrvifattribentry[i])] = types.YChild{"Intsrvifattribentry", &intsrvifattribtable.Intsrvifattribentry[i]}
    }
    intsrvifattribtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(intsrvifattribtable.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry
// The reservable attributes of  a  given  inter-
// face.
type INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry struct {
    EntityData types.CommonEntityData
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

func (intsrvifattribentry *INTEGRATEDSERVICESMIB_Intsrvifattribtable_Intsrvifattribentry) GetEntityData() *types.CommonEntityData {
    intsrvifattribentry.EntityData.YFilter = intsrvifattribentry.YFilter
    intsrvifattribentry.EntityData.YangName = "intSrvIfAttribEntry"
    intsrvifattribentry.EntityData.BundleName = "cisco_ios_xe"
    intsrvifattribentry.EntityData.ParentYangName = "intSrvIfAttribTable"
    intsrvifattribentry.EntityData.SegmentPath = "intSrvIfAttribEntry" + "[ifIndex='" + fmt.Sprintf("%v", intsrvifattribentry.Ifindex) + "']"
    intsrvifattribentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intsrvifattribentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intsrvifattribentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intsrvifattribentry.EntityData.Children = make(map[string]types.YChild)
    intsrvifattribentry.EntityData.Leafs = make(map[string]types.YLeaf)
    intsrvifattribentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", intsrvifattribentry.Ifindex}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribAllocatedBits"] = types.YLeaf{"Intsrvifattriballocatedbits", intsrvifattribentry.Intsrvifattriballocatedbits}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribMaxAllocatedBits"] = types.YLeaf{"Intsrvifattribmaxallocatedbits", intsrvifattribentry.Intsrvifattribmaxallocatedbits}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribAllocatedBuffer"] = types.YLeaf{"Intsrvifattriballocatedbuffer", intsrvifattribentry.Intsrvifattriballocatedbuffer}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribFlows"] = types.YLeaf{"Intsrvifattribflows", intsrvifattribentry.Intsrvifattribflows}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribPropagationDelay"] = types.YLeaf{"Intsrvifattribpropagationdelay", intsrvifattribentry.Intsrvifattribpropagationdelay}
    intsrvifattribentry.EntityData.Leafs["intSrvIfAttribStatus"] = types.YLeaf{"Intsrvifattribstatus", intsrvifattribentry.Intsrvifattribstatus}
    return &(intsrvifattribentry.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvflowtable
// Information describing the reserved flows  us-
// ing the system's interfaces.
type INTEGRATEDSERVICESMIB_Intsrvflowtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information describing the use of a given  in- terface   by   a   given  
    // flow.   The  counter intSrvFlowPoliced starts counting  at  the  in-
    // stallation of the flow. The type is slice of
    // INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry.
    Intsrvflowentry []INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry
}

func (intsrvflowtable *INTEGRATEDSERVICESMIB_Intsrvflowtable) GetEntityData() *types.CommonEntityData {
    intsrvflowtable.EntityData.YFilter = intsrvflowtable.YFilter
    intsrvflowtable.EntityData.YangName = "intSrvFlowTable"
    intsrvflowtable.EntityData.BundleName = "cisco_ios_xe"
    intsrvflowtable.EntityData.ParentYangName = "INTEGRATED-SERVICES-MIB"
    intsrvflowtable.EntityData.SegmentPath = "intSrvFlowTable"
    intsrvflowtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intsrvflowtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intsrvflowtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intsrvflowtable.EntityData.Children = make(map[string]types.YChild)
    intsrvflowtable.EntityData.Children["intSrvFlowEntry"] = types.YChild{"Intsrvflowentry", nil}
    for i := range intsrvflowtable.Intsrvflowentry {
        intsrvflowtable.EntityData.Children[types.GetSegmentPath(&intsrvflowtable.Intsrvflowentry[i])] = types.YChild{"Intsrvflowentry", &intsrvflowtable.Intsrvflowentry[i]}
    }
    intsrvflowtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(intsrvflowtable.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry
// Information describing the use of a given  in-
// terface   by   a   given   flow.   The  counter
// intSrvFlowPoliced starts counting  at  the  in-
// stallation of the flow.
type INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry struct {
    EntityData types.CommonEntityData
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

func (intsrvflowentry *INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry) GetEntityData() *types.CommonEntityData {
    intsrvflowentry.EntityData.YFilter = intsrvflowentry.YFilter
    intsrvflowentry.EntityData.YangName = "intSrvFlowEntry"
    intsrvflowentry.EntityData.BundleName = "cisco_ios_xe"
    intsrvflowentry.EntityData.ParentYangName = "intSrvFlowTable"
    intsrvflowentry.EntityData.SegmentPath = "intSrvFlowEntry" + "[intSrvFlowNumber='" + fmt.Sprintf("%v", intsrvflowentry.Intsrvflownumber) + "']"
    intsrvflowentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intsrvflowentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intsrvflowentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intsrvflowentry.EntityData.Children = make(map[string]types.YChild)
    intsrvflowentry.EntityData.Leafs = make(map[string]types.YLeaf)
    intsrvflowentry.EntityData.Leafs["intSrvFlowNumber"] = types.YLeaf{"Intsrvflownumber", intsrvflowentry.Intsrvflownumber}
    intsrvflowentry.EntityData.Leafs["intSrvFlowType"] = types.YLeaf{"Intsrvflowtype", intsrvflowentry.Intsrvflowtype}
    intsrvflowentry.EntityData.Leafs["intSrvFlowOwner"] = types.YLeaf{"Intsrvflowowner", intsrvflowentry.Intsrvflowowner}
    intsrvflowentry.EntityData.Leafs["intSrvFlowDestAddr"] = types.YLeaf{"Intsrvflowdestaddr", intsrvflowentry.Intsrvflowdestaddr}
    intsrvflowentry.EntityData.Leafs["intSrvFlowSenderAddr"] = types.YLeaf{"Intsrvflowsenderaddr", intsrvflowentry.Intsrvflowsenderaddr}
    intsrvflowentry.EntityData.Leafs["intSrvFlowDestAddrLength"] = types.YLeaf{"Intsrvflowdestaddrlength", intsrvflowentry.Intsrvflowdestaddrlength}
    intsrvflowentry.EntityData.Leafs["intSrvFlowSenderAddrLength"] = types.YLeaf{"Intsrvflowsenderaddrlength", intsrvflowentry.Intsrvflowsenderaddrlength}
    intsrvflowentry.EntityData.Leafs["intSrvFlowProtocol"] = types.YLeaf{"Intsrvflowprotocol", intsrvflowentry.Intsrvflowprotocol}
    intsrvflowentry.EntityData.Leafs["intSrvFlowDestPort"] = types.YLeaf{"Intsrvflowdestport", intsrvflowentry.Intsrvflowdestport}
    intsrvflowentry.EntityData.Leafs["intSrvFlowPort"] = types.YLeaf{"Intsrvflowport", intsrvflowentry.Intsrvflowport}
    intsrvflowentry.EntityData.Leafs["intSrvFlowFlowId"] = types.YLeaf{"Intsrvflowflowid", intsrvflowentry.Intsrvflowflowid}
    intsrvflowentry.EntityData.Leafs["intSrvFlowInterface"] = types.YLeaf{"Intsrvflowinterface", intsrvflowentry.Intsrvflowinterface}
    intsrvflowentry.EntityData.Leafs["intSrvFlowIfAddr"] = types.YLeaf{"Intsrvflowifaddr", intsrvflowentry.Intsrvflowifaddr}
    intsrvflowentry.EntityData.Leafs["intSrvFlowRate"] = types.YLeaf{"Intsrvflowrate", intsrvflowentry.Intsrvflowrate}
    intsrvflowentry.EntityData.Leafs["intSrvFlowBurst"] = types.YLeaf{"Intsrvflowburst", intsrvflowentry.Intsrvflowburst}
    intsrvflowentry.EntityData.Leafs["intSrvFlowWeight"] = types.YLeaf{"Intsrvflowweight", intsrvflowentry.Intsrvflowweight}
    intsrvflowentry.EntityData.Leafs["intSrvFlowQueue"] = types.YLeaf{"Intsrvflowqueue", intsrvflowentry.Intsrvflowqueue}
    intsrvflowentry.EntityData.Leafs["intSrvFlowMinTU"] = types.YLeaf{"Intsrvflowmintu", intsrvflowentry.Intsrvflowmintu}
    intsrvflowentry.EntityData.Leafs["intSrvFlowMaxTU"] = types.YLeaf{"Intsrvflowmaxtu", intsrvflowentry.Intsrvflowmaxtu}
    intsrvflowentry.EntityData.Leafs["intSrvFlowBestEffort"] = types.YLeaf{"Intsrvflowbesteffort", intsrvflowentry.Intsrvflowbesteffort}
    intsrvflowentry.EntityData.Leafs["intSrvFlowPoliced"] = types.YLeaf{"Intsrvflowpoliced", intsrvflowentry.Intsrvflowpoliced}
    intsrvflowentry.EntityData.Leafs["intSrvFlowDiscard"] = types.YLeaf{"Intsrvflowdiscard", intsrvflowentry.Intsrvflowdiscard}
    intsrvflowentry.EntityData.Leafs["intSrvFlowService"] = types.YLeaf{"Intsrvflowservice", intsrvflowentry.Intsrvflowservice}
    intsrvflowentry.EntityData.Leafs["intSrvFlowOrder"] = types.YLeaf{"Intsrvfloworder", intsrvflowentry.Intsrvfloworder}
    intsrvflowentry.EntityData.Leafs["intSrvFlowStatus"] = types.YLeaf{"Intsrvflowstatus", intsrvflowentry.Intsrvflowstatus}
    return &(intsrvflowentry.EntityData)
}

// INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner represents queue policy database.
type INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner string

const (
    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_other INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "other"

    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_rsvp INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "rsvp"

    INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner_management INTEGRATEDSERVICESMIB_Intsrvflowtable_Intsrvflowentry_Intsrvflowowner = "management"
)

