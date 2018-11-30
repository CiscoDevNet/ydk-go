// This is the management information for
// Quality Of Service (QOS) for DOCSIS 1.1.
package docs_qos_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_qos_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-QOS-MIB DOCS-QOS-MIB}", reflect.TypeOf(DOCSQOSMIB{}))
    ydk.RegisterEntity("DOCS-QOS-MIB:DOCS-QOS-MIB", reflect.TypeOf(DOCSQOSMIB{}))
}

// IfDirection represents Cable Modem Termination System.
type IfDirection string

const (
    IfDirection_downstream IfDirection = "downstream"

    IfDirection_upstream IfDirection = "upstream"
)

// SchedulingType represents Sets.
type SchedulingType string

const (
    SchedulingType_undefined SchedulingType = "undefined"

    SchedulingType_bestEffort SchedulingType = "bestEffort"

    SchedulingType_nonRealTimePollingService SchedulingType = "nonRealTimePollingService"

    SchedulingType_realTimePollingService SchedulingType = "realTimePollingService"

    SchedulingType_unsolictedGrantServiceWithAD SchedulingType = "unsolictedGrantServiceWithAD"

    SchedulingType_unsolictedGrantService SchedulingType = "unsolictedGrantService"
)

// DOCSQOSMIB
type DOCSQOSMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table describes the packet classification configured on the CM or
    // CMTS.   The model is that a packet either received as input from an
    // interface or transmitted  for output on an interface may be compared 
    // against an ordered list of rules pertaining to the packet contents. Each
    // rule is a row of this table. A matching rule provides a service flow id to
    // to which the packet is classified.  All rules need to match for a packet to
    // match  a classifier.   The objects in this row correspond to a set of
    // Classifier Encoding parameters in a DOCSIS MAC management message. The
    // docsQosPktClassBitMap indicates which particular parameters were present in
    // the classifier as signalled in the DOCSIS message. If the referenced
    // parameter was not present in the signalled DOCSIS 1.1 Classifier, the
    // corresponding object in this row reports a  value as specified in the
    // DESCRIPTION section.
    DocsQosPktClassTable DOCSQOSMIB_DocsQosPktClassTable

    // This table describes the set of DOCSIS 1.1 QOS  parameters defined in a
    // managed device.  The ifIndex index specifies a DOCSIS MAC Domain. The
    // docsQosServiceFlowId index specifies a particular Service Flow.  The
    // docsQosParamSetType index indicates whether the active, admitted, or
    // provisioned QOS Parameter  Set is being described by the row.  Only the QOS
    // Parameter Sets of Docsis 1.1 service flows are represented in this table. 
    // Docsis 1.0 QOS service profiles are not represented in this table.  Each
    // row corresponds to a DOCSIS QOS Parameter Set as signaled via DOCSIS MAC
    // management messages. Each object in the row corresponds to one or  part of
    // one DOCSIS 1.1 Service Flow Encoding. The docsQosParamSetBitMap object in
    // the row indicates which particular parameters were signalled in  the
    // original registration or dynamic service request message that created the
    // QOS Parameter Set.  In many cases, even if a QOS Parameter Set parameter
    // was not signalled, the DOCSIS specification calls for a default value to be
    // used. That default value is reported as the value of the corresponding
    // object in this row.  Many objects are not applicable depending on the
    // service flow direction or upstream scheduling type.  The object value
    // reported in this case is specified in the DESCRIPTION clause.
    DocsQosParamSetTable DOCSQOSMIB_DocsQosParamSetTable

    // This table describes the set of Docsis-QOS  Service Flows in a managed
    // device. .
    DocsQosServiceFlowTable DOCSQOSMIB_DocsQosServiceFlowTable

    // This table describes statistics associated with the   Service Flows in a
    // managed device. .
    DocsQosServiceFlowStatsTable DOCSQOSMIB_DocsQosServiceFlowStatsTable

    // This table describes statistics associated with   upstream service flows.
    // All counted frames must  be received without an FCS error.
    DocsQosUpstreamStatsTable DOCSQOSMIB_DocsQosUpstreamStatsTable

    // This table describes statistics associated with the   Dynamic Service Flows
    // in a managed device. .
    DocsQosDynamicServiceStatsTable DOCSQOSMIB_DocsQosDynamicServiceStatsTable

    // This table contains a log of the disconnected Service Flows in a managed
    // device.
    DocsQosServiceFlowLogTable DOCSQOSMIB_DocsQosServiceFlowLogTable

    // This table describes the set of Docsis-QOS  Service Classes in a CMTS. .
    DocsQosServiceClassTable DOCSQOSMIB_DocsQosServiceClassTable

    // This table describes the set of Docsis-QOS  Service Class Policies.    This
    // table is an adjunct to the docsDevFilterPolicy table.  Entries in 
    // docsDevFilterPolicy table can  point to  specific rows in this table.  This
    // table permits mapping a packet to a service class name of an active service
    // flow so long as  a classifier does not exist at a higher priority.
    DocsQosServiceClassPolicyTable DOCSQOSMIB_DocsQosServiceClassPolicyTable

    // This table describes set of payload header suppression entries.
    DocsQosPHSTable DOCSQOSMIB_DocsQosPHSTable

    // This table provide for referencing the service flows  associated with a
    // particular cable modem. This allows  for indexing into other docsQos tables
    // that are  indexed by docsQosServiceFlowId and ifIndex.
    DocsQosCmtsMacToSrvFlowTable DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable
}

func (dOCSQOSMIB *DOCSQOSMIB) GetEntityData() *types.CommonEntityData {
    dOCSQOSMIB.EntityData.YFilter = dOCSQOSMIB.YFilter
    dOCSQOSMIB.EntityData.YangName = "DOCS-QOS-MIB"
    dOCSQOSMIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSQOSMIB.EntityData.ParentYangName = "DOCS-QOS-MIB"
    dOCSQOSMIB.EntityData.SegmentPath = "DOCS-QOS-MIB:DOCS-QOS-MIB"
    dOCSQOSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSQOSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSQOSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSQOSMIB.EntityData.Children = types.NewOrderedMap()
    dOCSQOSMIB.EntityData.Children.Append("docsQosPktClassTable", types.YChild{"DocsQosPktClassTable", &dOCSQOSMIB.DocsQosPktClassTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosParamSetTable", types.YChild{"DocsQosParamSetTable", &dOCSQOSMIB.DocsQosParamSetTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosServiceFlowTable", types.YChild{"DocsQosServiceFlowTable", &dOCSQOSMIB.DocsQosServiceFlowTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosServiceFlowStatsTable", types.YChild{"DocsQosServiceFlowStatsTable", &dOCSQOSMIB.DocsQosServiceFlowStatsTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosUpstreamStatsTable", types.YChild{"DocsQosUpstreamStatsTable", &dOCSQOSMIB.DocsQosUpstreamStatsTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosDynamicServiceStatsTable", types.YChild{"DocsQosDynamicServiceStatsTable", &dOCSQOSMIB.DocsQosDynamicServiceStatsTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosServiceFlowLogTable", types.YChild{"DocsQosServiceFlowLogTable", &dOCSQOSMIB.DocsQosServiceFlowLogTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosServiceClassTable", types.YChild{"DocsQosServiceClassTable", &dOCSQOSMIB.DocsQosServiceClassTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosServiceClassPolicyTable", types.YChild{"DocsQosServiceClassPolicyTable", &dOCSQOSMIB.DocsQosServiceClassPolicyTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosPHSTable", types.YChild{"DocsQosPHSTable", &dOCSQOSMIB.DocsQosPHSTable})
    dOCSQOSMIB.EntityData.Children.Append("docsQosCmtsMacToSrvFlowTable", types.YChild{"DocsQosCmtsMacToSrvFlowTable", &dOCSQOSMIB.DocsQosCmtsMacToSrvFlowTable})
    dOCSQOSMIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSQOSMIB.EntityData.YListKeys = []string {}

    return &(dOCSQOSMIB.EntityData)
}

// DOCSQOSMIB_DocsQosPktClassTable
// This table describes the packet classification
// configured on the CM or CMTS.  
// The model is that a packet either received
// as input from an interface or transmitted 
// for output on an interface may be compared 
// against an ordered list of rules pertaining to
// the packet contents. Each rule is a row of this
// table. A matching rule provides a service flow
// id to to which the packet is classified. 
// All rules need to match for a packet to match 
// a classifier. 
// 
// The objects in this row correspond to a set of
// Classifier Encoding parameters in a DOCSIS
// MAC management message. The docsQosPktClassBitMap
// indicates which particular parameters were present
// in the classifier as signalled in the DOCSIS message.
// If the referenced parameter was not present
// in the signalled DOCSIS 1.1 Classifier, the
// corresponding object in this row reports a 
// value as specified in the DESCRIPTION section.
type DOCSQOSMIB_DocsQosPktClassTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table provides a single packet classifier rule. The index
    // ifIndex is an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry.
    DocsQosPktClassEntry []*DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry
}

func (docsQosPktClassTable *DOCSQOSMIB_DocsQosPktClassTable) GetEntityData() *types.CommonEntityData {
    docsQosPktClassTable.EntityData.YFilter = docsQosPktClassTable.YFilter
    docsQosPktClassTable.EntityData.YangName = "docsQosPktClassTable"
    docsQosPktClassTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosPktClassTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosPktClassTable.EntityData.SegmentPath = "docsQosPktClassTable"
    docsQosPktClassTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosPktClassTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosPktClassTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosPktClassTable.EntityData.Children = types.NewOrderedMap()
    docsQosPktClassTable.EntityData.Children.Append("docsQosPktClassEntry", types.YChild{"DocsQosPktClassEntry", nil})
    for i := range docsQosPktClassTable.DocsQosPktClassEntry {
        docsQosPktClassTable.EntityData.Children.Append(types.GetSegmentPath(docsQosPktClassTable.DocsQosPktClassEntry[i]), types.YChild{"DocsQosPktClassEntry", docsQosPktClassTable.DocsQosPktClassEntry[i]})
    }
    docsQosPktClassTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosPktClassTable.EntityData.YListKeys = []string {}

    return &(docsQosPktClassTable.EntityData)
}

// DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry
// An entry in this table provides a single packet
// classifier rule. The index ifIndex is an ifType
// of docsCableMaclayer(127).
type DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // This attribute is a key. Index assigned to packet classifier entry by  the
    // CMTS which is unique per service flow. The type is interface{} with range:
    // 1..65535.
    DocsQosPktClassId interface{}

    // Indicates the direction to which the classifier  is applied. The type is
    // IfDirection.
    DocsQosPktClassDirection interface{}

    // The value specifies the order of evaluation of the classifiers. The higher
    // the value the higher the priority. The value of 0 is used as default in 
    // provisioned service flows classifiers.  The default value of 64 is used for
    // dynamic  service flow classifiers. If the referenced parameter is not
    // present in a classifier, this object reports the default value as defined
    // above. The type is interface{} with range: 0..255.
    DocsQosPktClassPriority interface{}

    // The low value of a range of TOS byte values. If the referenced parameter is
    // not present in a classifier, this object reports the value of 0. The type
    // is string with length: 1.
    DocsQosPktClassIpTosLow interface{}

    // The 8-bit high value of a range of TOS byte values.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // 0. The type is string with length: 1.
    DocsQosPktClassIpTosHigh interface{}

    // The mask value is bitwise ANDed with TOS byte  in an IP packet and this
    // value is used check  range checking of TosLow and TosHigh.  If the
    // referenced parameter is not present in a classifier, this object reports
    // the value of 0. The type is string with length: 1.
    DocsQosPktClassIpTosMask interface{}

    // This object indicates the value of the IP Protocol field required for IP
    // packets to match this rule.   The value 256 matches traffic with any IP
    // Protocol  value. The value 257 by convention matches both TCP and UDP.   If
    // the referenced parameter is not present in a classifier, this object
    // reports the value of 258. The type is interface{} with range: 0..258.
    DocsQosPktClassIpProtocol interface{}

    // This object is deprecated in favor of the object pair
    // docsQosPktClassInetSourceAddrType and docsQosPktClassInetSourceAddr. Agents
    // that choose to implement this object MUST report an address that matches
    // docsQosPktClassInetSourceAddr object as long as
    // docsQosPktClassInetSourceAddrType is ipv4(1). Otherwise, the value of this
    // object shall be 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsQosPktClassIpSourceAddr interface{}

    // This object is deprecated in favor of the object pair
    // docsQosPktClassInetSourceMaskType and docsQosPktClassInetSourceMask. Agents
    // that choose to implement this object MUST report an address that matches
    // docsQosPktClassInetSourceMask object as long as
    // docsQosPktClassInetSourceMaskType is ipv4(1). Otherwise, the value of this
    // object shall be 255.255.255.255.  SNMP mangers should note that agent
    // implementation of previous versions of this MIB report 0.0.0.0 as the value
    // when the reference parameter is not present, rather than 255.255.255.255.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsQosPktClassIpSourceMask interface{}

    // This object is deprecated in favor of the object pair
    // docsQosPktClassInetDestAddrType and docsQosPktClassInetDestAddr. Agents
    // that choose to implement this object MUST report an address that matches
    // docsQosPktClassInetDestAddr object as long as
    // docsQosPktClassInetDestAddrType is ipv4(1). Otherwise, the value of this
    // object shall be 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsQosPktClassIpDestAddr interface{}

    // This object is deprecated in favor of the object pair
    // docsQosPktClassInetDestMaskType and docsQosPktClassInetDestMask. Agents
    // that choose to implement this object MUST report an address that matches
    // docsQosPktClassInetDestMask object as long as
    // docsQosPktClassInetDestMaskType is ipv4(1). Otherwise, the value of this
    // object shall be 255.255.255.255.  SNMP mangers should note that agent
    // implementation of previous versions of this MIB report 0.0.0.0 as the value
    // when the reference parameter is not present, rather than 255.255.255.255.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsQosPktClassIpDestMask interface{}

    // This object specifies the low end inclusive range of TCP/UDP source port
    // numbers to which a packet is compared. This object is irrelevant for
    // non-TCP/UDP IP packets.  If the referenced parameter is not present in a
    // classifier, this object reports the value of 0. The type is interface{}
    // with range: 0..65535.
    DocsQosPktClassSourcePortStart interface{}

    // This object specifies the high end inclusive range of TCP/UDP source port
    // numbers to which a packet is compared. This object is irrelevant for
    // non-TCP/UDP IP packets.  If the referenced parameter is not present in a
    // classifier, this object reports the value of  65535. The type is
    // interface{} with range: 0..65535.
    DocsQosPktClassSourcePortEnd interface{}

    // This object specifies the low end inclusive range of TCP/UDP destination
    // port numbers to which a packet is compared.  If the referenced parameter is
    // not present in a classifier, this object reports the value of 0. The type
    // is interface{} with range: 0..65535.
    DocsQosPktClassDestPortStart interface{}

    // This object specifies the high end inclusive range of TCP/UDP destination
    // port numbers to which a packet is compared.  If the referenced parameter is
    // not present in a classifier, this object reports the value of  65535. The
    // type is interface{} with range: 0..65535.
    DocsQosPktClassDestPortEnd interface{}

    // An Ethernet packet matches an entry when its  destination MAC address
    // bitwise ANDed with  docsQosPktClassDestMacMask equals the value of 
    // docsQosPktClassDestMacAddr.   If the referenced parameter is not present in
    // a classifier, this object reports the value of  '000000000000'H. The type
    // is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsQosPktClassDestMacAddr interface{}

    // An Ethernet packet matches an entry when its  destination MAC address
    // bitwise ANDed with  docsQosPktClassDestMacMask equals the value of 
    // docsQosPktClassDestMacAddr.  If the referenced parameter is not present in
    // a classifier, this object reports the value of  '000000000000'H. The type
    // is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsQosPktClassDestMacMask interface{}

    // An Ethernet packet matches this entry when its  source MAC address equals
    // the value of  this object.  If the referenced parameter is not present in a
    // classifier, this object reports the value of  'FFFFFFFFFFFF'H. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsQosPktClassSourceMacAddr interface{}

    // This object indicates the format of the layer 3  protocol id in the
    // Ethernet packet. A value of  none(0) means that the rule does not use the 
    // layer 3 protocol type as a matching criteria.  A value of ethertype(1)
    // means that the rule applies only to frames which contains an  EtherType
    // value. Ethertype values are contained in packets using the Dec-Intel-Xerox
    // (DIX) encapsulation or the RFC1042 Sub-Network Access Protocol (SNAP)
    // encapsulation formats.  A value of dsap(2) means that the rule applies only
    // to frames using the IEEE802.3 encapsulation format with a Destination
    // Service Access Point (DSAP) other  than 0xAA (which is reserved for SNAP). 
    // A value of mac(3) means that the rule applies  only to MAC management
    // messages for MAC management messages.  A value of all(4) means that the
    // rule matches all Ethernet packets.   If the Ethernet frame contains an
    // 802.1P/Q Tag  header (i.e. EtherType 0x8100), this object applies to the
    // embedded EtherType field within  the 802.1P/Q header.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // 0. The type is DocsQosPktClassEnetProtocolType.
    DocsQosPktClassEnetProtocolType interface{}

    // If docsQosEthPktClassProtocolType is none(0),  this object is ignored when
    // considering whether  a packet matches the current rule.  If
    // dosQosPktClassEnetProtocolType is ethertype(1), this object gives the
    // 16-bit value of the EtherType that the packet must match in order to match
    // the rule.  If docsQosPktClassEnetProtocolType is dsap(2), the lower 8 bits
    // of this object's value must match the DSAP byte of the packet in order to
    // match the rule.  If docsQosPktClassEnetProtocolType is mac(3), the lower 8
    // bits of this object value represent a lower bound (inclusive) of MAC
    // management message type codes matched, and the upper 8 bits of this object
    // value represent the upper bound (inclusive) of matched MAC message type
    // codes.  Certain message type codes are excluded from matching, as specified
    // in the reference.  If the Ethernet frame contains an 802.1P/Q Tag header 
    // (i.e. EtherType 0x8100), this object applies to the  embedded EtherType
    // field within the 802.1P/Q header.  If the referenced parameter is not
    // present in the classifier, the value of this object is reported as 0. The
    // type is interface{} with range: 0..65535.
    DocsQosPktClassEnetProtocol interface{}

    // This object is obsolete. The type is interface{} with range: 0..1.
    DocsQosPktClassUserPriApplies interface{}

    // This object applies only to Ethernet frames  using the 802.1P/Q tag header
    // (indicated with  EtherType 0x8100). Such frames include a 16-bit  Tag that
    // contains a 3 bit Priority field and a 12 bit VLAN number.  Tagged Ethernet
    // packets must have a 3-bit Priority field within the range of 
    // docsQosPktClassPriLow and docsQosPktClassPriHigh in  order to match this
    // rule.  If the referenced parameter is not present in the classifier, the
    // value of this object is reported as 0. The type is interface{} with range:
    // 0..7.
    DocsQosPktClassUserPriLow interface{}

    // This object applies only to Ethernet frames  using the 802.1P/Qtag header
    // (indicated with  EtherType 0x8100). Such frames include a 16-bit  Tag that
    // contains a 3 bit Priority field and a 12 bit VLAN number.  Tagged Ethernet
    // packets must have a 3-bit Priority field within the range of 
    // docsQosPktClassPriLow and  docsQosPktClassPriHigh in order to match this
    // rule.  If the referenced parameter is not present in the classifier, the
    // value of this object is reported  as 7. The type is interface{} with range:
    // 0..7.
    DocsQosPktClassUserPriHigh interface{}

    // This object applies only to Ethernet frames  using the 802.1P/Q tag header.
    // If this object's value is nonzero, tagged packets must have a VLAN
    // Identifier that matches the value in order to match the rule.  Only the
    // least significant 12 bits of this object's value are valid.   If the
    // referenced parameter is not present in the classifier, the value of this
    // object is reported  as 0. The type is interface{} with range: 0..4095.
    DocsQosPktClassVlanId interface{}

    // This object indicates whether or not the classifier is enabled to classify
    // packets to a Service Flow.  If the referenced parameter is not present in
    // the classifier, the value of this object is reported  as active(1). The
    // type is DocsQosPktClassState.
    DocsQosPktClassState interface{}

    // This object counts the number of packets that have been classified using
    // this entry. The type is interface{} with range: 0..4294967295.
    DocsQosPktClassPkts interface{}

    // This object indicates which parameter encodings were actually present in
    // the DOCSIS packet classifier encoding signaled in the DOCSIS message that
    // created or modified the classifier. Note that Dynamic Service Change
    // messages have replace semantics, so that all non-default parameters must be
    // present whether the classifier is being created or changed.  A bit of of
    // this object is set to 1 if the parameter indicated by the comment was
    // present in the classifier  encoding, and 0 otherwise.  Note that BITS are
    // encoded most significant bit first, so that if e.g. bits 6 and 7 are set,
    // this object is encoded as the octet string '030000'H. The type is
    // map[string]bool.
    DocsQosPktClassBitMap interface{}

    // The type of the internet address for docsQosPktClassInetSourceAddr. This
    // type must be the same as the docsQosPktClassInetSourceMaskType.  If the
    // referenced parameter is not present in a classifier, this object reports
    // the value of ipv4(1). The type is InetAddressType.
    DocsQosPktClassInetSourceAddrType interface{}

    // This object specifies the value of the IP Source Address required for
    // packets to match this rule. An IP packet matches the rule when the packet
    // ip source address bitwise ANDed with the docsQosPktClassInetSourceMask
    // value equals the docsQosPktClassInetSourceAddr value.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // '00000000'H. The type is string with length: 0..255.
    DocsQosPktClassInetSourceAddr interface{}

    // The type of the internet address for docsQosPktClassInetSourceMask. This
    // type must be the same as the docsQosPktClassInetSourceAddrType.  If the
    // referenced parameter is not present in a classifier, this object reports
    // the value of ipv4(1). The type is InetAddressType.
    DocsQosPktClassInetSourceMaskType interface{}

    // This object specifies which bits of a packet's IP Source Address that are
    // compared to match this rule. An IP packet matches the rule when the packet
    // source address bitwise ANDed with the docsQosPktClassInetSourceMask value
    // equals the docsQosIpPktClassInetSourceAddr value.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // 'FFFFFFFF'H. The type is string with length: 0..255.
    DocsQosPktClassInetSourceMask interface{}

    // The type of the internet address for docsQosPktClassInetDestAddr. This type
    // must be the same as the docsQosPktClassInetDestMaskType.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // ipv4(1). The type is InetAddressType.
    DocsQosPktClassInetDestAddrType interface{}

    // This object specifies the value of the IP Destination Address required for
    // packets to match this rule. An IP packet matches the rule when the packet
    // ip destination address bitwise ANDed with the docsQosPktClassInetDestMask
    // value equals the docsQosPktClassInetDestAddr value.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // '00000000'H. The type is string with length: 0..255.
    DocsQosPktClassInetDestAddr interface{}

    // The type of the internet address for docsQosPktClassInetDestMask. This type
    // must be the same as the docsQosPktClassInetDestAddrType.  If the referenced
    // parameter is not present in a classifier, this object reports the value of
    // ipv4(1). The type is InetAddressType.
    DocsQosPktClassInetDestMaskType interface{}

    // This object specifies which bits of a packet's IP Destination Address that
    // are compared to match this rule. An IP packet matches the rule when the
    // packet destination address bitwise ANDed with the
    // docsQosPktClassInetDestMask value equals the docsQosIpPktClassInetDestAddr
    // value.  If the referenced parameter is not present in a classifier, this
    // object reports the value of 'FFFFFFFF'H. The type is string with length:
    // 0..255.
    DocsQosPktClassInetDestMask interface{}
}

func (docsQosPktClassEntry *DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry) GetEntityData() *types.CommonEntityData {
    docsQosPktClassEntry.EntityData.YFilter = docsQosPktClassEntry.YFilter
    docsQosPktClassEntry.EntityData.YangName = "docsQosPktClassEntry"
    docsQosPktClassEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosPktClassEntry.EntityData.ParentYangName = "docsQosPktClassTable"
    docsQosPktClassEntry.EntityData.SegmentPath = "docsQosPktClassEntry" + types.AddKeyToken(docsQosPktClassEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosPktClassEntry.DocsQosServiceFlowId, "docsQosServiceFlowId") + types.AddKeyToken(docsQosPktClassEntry.DocsQosPktClassId, "docsQosPktClassId")
    docsQosPktClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosPktClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosPktClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosPktClassEntry.EntityData.Children = types.NewOrderedMap()
    docsQosPktClassEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosPktClassEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosPktClassEntry.IfIndex})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsQosPktClassEntry.DocsQosServiceFlowId})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassId", types.YLeaf{"DocsQosPktClassId", docsQosPktClassEntry.DocsQosPktClassId})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassDirection", types.YLeaf{"DocsQosPktClassDirection", docsQosPktClassEntry.DocsQosPktClassDirection})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassPriority", types.YLeaf{"DocsQosPktClassPriority", docsQosPktClassEntry.DocsQosPktClassPriority})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpTosLow", types.YLeaf{"DocsQosPktClassIpTosLow", docsQosPktClassEntry.DocsQosPktClassIpTosLow})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpTosHigh", types.YLeaf{"DocsQosPktClassIpTosHigh", docsQosPktClassEntry.DocsQosPktClassIpTosHigh})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpTosMask", types.YLeaf{"DocsQosPktClassIpTosMask", docsQosPktClassEntry.DocsQosPktClassIpTosMask})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpProtocol", types.YLeaf{"DocsQosPktClassIpProtocol", docsQosPktClassEntry.DocsQosPktClassIpProtocol})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpSourceAddr", types.YLeaf{"DocsQosPktClassIpSourceAddr", docsQosPktClassEntry.DocsQosPktClassIpSourceAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpSourceMask", types.YLeaf{"DocsQosPktClassIpSourceMask", docsQosPktClassEntry.DocsQosPktClassIpSourceMask})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpDestAddr", types.YLeaf{"DocsQosPktClassIpDestAddr", docsQosPktClassEntry.DocsQosPktClassIpDestAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassIpDestMask", types.YLeaf{"DocsQosPktClassIpDestMask", docsQosPktClassEntry.DocsQosPktClassIpDestMask})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassSourcePortStart", types.YLeaf{"DocsQosPktClassSourcePortStart", docsQosPktClassEntry.DocsQosPktClassSourcePortStart})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassSourcePortEnd", types.YLeaf{"DocsQosPktClassSourcePortEnd", docsQosPktClassEntry.DocsQosPktClassSourcePortEnd})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassDestPortStart", types.YLeaf{"DocsQosPktClassDestPortStart", docsQosPktClassEntry.DocsQosPktClassDestPortStart})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassDestPortEnd", types.YLeaf{"DocsQosPktClassDestPortEnd", docsQosPktClassEntry.DocsQosPktClassDestPortEnd})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassDestMacAddr", types.YLeaf{"DocsQosPktClassDestMacAddr", docsQosPktClassEntry.DocsQosPktClassDestMacAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassDestMacMask", types.YLeaf{"DocsQosPktClassDestMacMask", docsQosPktClassEntry.DocsQosPktClassDestMacMask})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassSourceMacAddr", types.YLeaf{"DocsQosPktClassSourceMacAddr", docsQosPktClassEntry.DocsQosPktClassSourceMacAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassEnetProtocolType", types.YLeaf{"DocsQosPktClassEnetProtocolType", docsQosPktClassEntry.DocsQosPktClassEnetProtocolType})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassEnetProtocol", types.YLeaf{"DocsQosPktClassEnetProtocol", docsQosPktClassEntry.DocsQosPktClassEnetProtocol})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassUserPriApplies", types.YLeaf{"DocsQosPktClassUserPriApplies", docsQosPktClassEntry.DocsQosPktClassUserPriApplies})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassUserPriLow", types.YLeaf{"DocsQosPktClassUserPriLow", docsQosPktClassEntry.DocsQosPktClassUserPriLow})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassUserPriHigh", types.YLeaf{"DocsQosPktClassUserPriHigh", docsQosPktClassEntry.DocsQosPktClassUserPriHigh})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassVlanId", types.YLeaf{"DocsQosPktClassVlanId", docsQosPktClassEntry.DocsQosPktClassVlanId})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassState", types.YLeaf{"DocsQosPktClassState", docsQosPktClassEntry.DocsQosPktClassState})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassPkts", types.YLeaf{"DocsQosPktClassPkts", docsQosPktClassEntry.DocsQosPktClassPkts})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassBitMap", types.YLeaf{"DocsQosPktClassBitMap", docsQosPktClassEntry.DocsQosPktClassBitMap})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetSourceAddrType", types.YLeaf{"DocsQosPktClassInetSourceAddrType", docsQosPktClassEntry.DocsQosPktClassInetSourceAddrType})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetSourceAddr", types.YLeaf{"DocsQosPktClassInetSourceAddr", docsQosPktClassEntry.DocsQosPktClassInetSourceAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetSourceMaskType", types.YLeaf{"DocsQosPktClassInetSourceMaskType", docsQosPktClassEntry.DocsQosPktClassInetSourceMaskType})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetSourceMask", types.YLeaf{"DocsQosPktClassInetSourceMask", docsQosPktClassEntry.DocsQosPktClassInetSourceMask})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetDestAddrType", types.YLeaf{"DocsQosPktClassInetDestAddrType", docsQosPktClassEntry.DocsQosPktClassInetDestAddrType})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetDestAddr", types.YLeaf{"DocsQosPktClassInetDestAddr", docsQosPktClassEntry.DocsQosPktClassInetDestAddr})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetDestMaskType", types.YLeaf{"DocsQosPktClassInetDestMaskType", docsQosPktClassEntry.DocsQosPktClassInetDestMaskType})
    docsQosPktClassEntry.EntityData.Leafs.Append("docsQosPktClassInetDestMask", types.YLeaf{"DocsQosPktClassInetDestMask", docsQosPktClassEntry.DocsQosPktClassInetDestMask})

    docsQosPktClassEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId", "DocsQosPktClassId"}

    return &(docsQosPktClassEntry.EntityData)
}

// DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType represents in a classifier, this object reports the value of 0.
type DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType string

const (
    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType_none DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType = "none"

    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType_ethertype DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType = "ethertype"

    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType_dsap DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType = "dsap"

    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType_mac DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType = "mac"

    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType_all DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassEnetProtocolType = "all"
)

// DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState represents as active(1).
type DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState string

const (
    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState_active DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState = "active"

    DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState_inactive DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassState = "inactive"
)

// DOCSQOSMIB_DocsQosParamSetTable
// This table describes the set of DOCSIS 1.1 QOS 
// parameters defined in a managed device.
// 
// The ifIndex index specifies a DOCSIS MAC Domain.
// The docsQosServiceFlowId index specifies a particular
// Service Flow. 
// The docsQosParamSetType index indicates whether
// the active, admitted, or provisioned QOS Parameter 
// Set is being described by the row.
// 
// Only the QOS Parameter Sets of Docsis 1.1 service
// flows are represented in this table.  Docsis 1.0
// QOS service profiles are not represented in this
// table.
// 
// Each row corresponds to a DOCSIS QOS Parameter Set
// as signaled via DOCSIS MAC management messages.
// Each object in the row corresponds to one or 
// part of one DOCSIS 1.1 Service Flow Encoding.
// The docsQosParamSetBitMap object in the row indicates
// which particular parameters were signalled in 
// the original registration or dynamic service
// request message that created the QOS Parameter Set.
// 
// In many cases, even if a QOS Parameter Set parameter
// was not signalled, the DOCSIS specification calls
// for a default value to be used. That default value
// is reported as the value of the corresponding object
// in this row.
// 
// Many objects are not applicable depending on
// the service flow direction or upstream scheduling
// type.  The object value reported in this case
// is specified in the DESCRIPTION clause.
type DOCSQOSMIB_DocsQosParamSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A unique set of QOS parameters. The type is slice of
    // DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry.
    DocsQosParamSetEntry []*DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry
}

func (docsQosParamSetTable *DOCSQOSMIB_DocsQosParamSetTable) GetEntityData() *types.CommonEntityData {
    docsQosParamSetTable.EntityData.YFilter = docsQosParamSetTable.YFilter
    docsQosParamSetTable.EntityData.YangName = "docsQosParamSetTable"
    docsQosParamSetTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosParamSetTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosParamSetTable.EntityData.SegmentPath = "docsQosParamSetTable"
    docsQosParamSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosParamSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosParamSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosParamSetTable.EntityData.Children = types.NewOrderedMap()
    docsQosParamSetTable.EntityData.Children.Append("docsQosParamSetEntry", types.YChild{"DocsQosParamSetEntry", nil})
    for i := range docsQosParamSetTable.DocsQosParamSetEntry {
        docsQosParamSetTable.EntityData.Children.Append(types.GetSegmentPath(docsQosParamSetTable.DocsQosParamSetEntry[i]), types.YChild{"DocsQosParamSetEntry", docsQosParamSetTable.DocsQosParamSetEntry[i]})
    }
    docsQosParamSetTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosParamSetTable.EntityData.YListKeys = []string {}

    return &(docsQosParamSetTable.EntityData)
}

// DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry
// A unique set of QOS parameters.
type DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // This attribute is a key. Defines the type of the QOS parameter set defined
    // by this row. active(1) indicates the Active QOS parameter set, describing
    // the service currently being provided by the Docsis MAC domain to the 
    // service flow. admitted(2) indicates the Admitted QOS Parameter Set,
    // describing services reserved by by the Docsis MAC domain for use by the
    // service flow. provisioned (3) describes the QOS Parameter Set defined in
    // the DOCSIS CM Configuration file for the service flow. The type is
    // DocsQosParamSetType.
    DocsQosParamSetType interface{}

    // Refers to the Service Class Name that the  parameter set values were
    // derived.  If the referenced parameter is not present in the corresponding
    // DOCSIS QOS Parameter Set, the default  value of this object is a zero
    // length string. The type is string.
    DocsQosParamSetServiceClassName interface{}

    // The relative priority of a service flow. Higher numbers indicate higher
    // priority. This priority should only be used to differentiate service flow
    // with identical parameter sets.  If the referenced parameter is not present
    // in the corresponding DOCSIS QOS Parameter Set, the default  value of this
    // object is 0.  If the parameter is not applicable, the reported value is 0.
    // The type is interface{} with range: 0..7.
    DocsQosParamSetPriority interface{}

    // Maximum sustained traffic rate allowed for this  service flow in bits/sec.
    // Must count all MAC frame  data PDU from the bytes following the MAC header
    // HCS to the end of the CRC. The number of bytes  forwarded is limited during
    // any time interval. The value 0 means no maximum traffic rate is  enforced.
    // This object applies to both upstream and downstream service flows.  If the
    // referenced parameter is not present in the corresponding DOCSIS QOS
    // Parameter Set, the default value of this object is 0. If the parameter is
    // not applicable, it is reported as 0. The type is interface{} with range:
    // 0..4294967295.
    DocsQosParamSetMaxTrafficRate interface{}

    // Specifies the token bucket size in bytes for this parameter set. The value
    // is calculated  from the byte following the MAC header HCS to  the end of
    // the CRC. This object is applied in  conjunction with
    // docsQosParamSetMaxTrafficRate to  calculate maximum sustained traffic rate.
    // If the referenced parameter is not present in the corresponding DOCSIS QOS
    // Parameter Set, the default value of this object for scheduling types
    // bestEffort (2), nonRealTimePollingService(3), and realTimePollingService(4)
    // is 1522.   If this parameter is not applicable, it is reported as 0. The
    // type is interface{} with range: 0..4294967295.
    DocsQosParamSetMaxTrafficBurst interface{}

    // Specifies the guaranteed minimum rate in bits/sec for this parameter set.
    // The value is  calculated from the byte following the MAC  header HCS to the
    // end of the CRC. The default value of 0 has the meaning that no bandwidth 
    // is reserved. If the referenced parameter is not present in the
    // corresponding DOCSIS QOS Parameter Set, the default value of this object is
    // 0. If the parameter is not applicable, it is reported as 0. The type is
    // interface{} with range: 0..4294967295.
    DocsQosParamSetMinReservedRate interface{}

    // Specifies an assumed minimum packet size in  bytes for which the
    // docsQosParamSetMinReservedRate  will be provided. The value is calculated
    // from the byte following the MAC header HCS to the  end of the CRC.       If
    // the referenced parameter is omitted from a DOCSIS QOS parameter set, the
    // default value is CMTS implementation dependent. In this case, the CMTS
    // reports the default value it is using and the CM reports a value of 0. If
    // the referenced parameter is not applicable to the direction or scheduling
    // type of the service flow, both CMTS and CM report this object's value as 0.
    // The type is interface{} with range: 0..65535.
    DocsQosParamSetMinReservedPkt interface{}

    // Specifies the maximum duration in seconds that  resources remain unused on
    // an active service flow before CMTS signals that both active and admitted
    // parameters set are null. The default value of 0 signifies an infinite
    // amount of time.  If the referenced parameter is not present in the
    // corresponding DOCSIS QOS Parameter Set, the default value of this object is
    // 0. The type is interface{} with range: 0..65535. Units are seconds.
    DocsQosParamSetActiveTimeout interface{}

    // Specifies the maximum duration in seconds that  resources remain in
    // admitted state before  resources must be released. The value of 0 signifies
    // an infinite amount  of time.  If the referenced parameter is not present in
    // the corresponding DOCSIS QOS Parameter Set, the  default value of this
    // object is 200. The type is interface{} with range: 0..65535. Units are
    // seconds.
    DocsQosParamSetAdmittedTimeout interface{}

    // Specifies the maximum concatenated burst in bytes which an upstream 
    // service flow is allowed.  The value is calculated from the FC byte of the
    // Concatenation MAC Header to the last CRC byte in  of the last concatenated
    // MAC frame, inclusive. The value of 0 specifies no maximum burst.  If the
    // referenced parameter is not present in the corresponding DOCSIS QOS
    // Parameter Set, the default value of this object is 0. If the parameter is
    // not applicable, this object's value is reported as 0. The type is
    // interface{} with range: 0..65535.
    DocsQosParamSetMaxConcatBurst interface{}

    // Specifies the upstream scheduling service used for  upstream service flow. 
    // If the referenced parameter is not present in the corresponding DOCSIS QOS
    // Parameter Set of an upstream service flow, the default value of this object
    // is bestEffort(2). For QOS parameter sets of downstream service flows, this
    // object's value is reported as undefined(1). The type is SchedulingType.
    DocsQosParamSetSchedulingType interface{}

    // Specifies the nominal interval in microseconds  between successive unicast
    // request opportunities on an upstream service flow.  This object applies
    // only to upstream service flows with schedulingType of value
    // nonRealTimePollingService(3), realTimePollingService(4), and
    // unsolictedGrantServiceWithAD(5).  The parameter is mandatory for
    // realTimePollingService(4).  If the parameter is omitted with
    // nonRealTimePollingService(3), the CMTS uses an implementation dependent
    // value.  If the parameter is omitted with unsolictedGrantServiceWithAD(5),
    // the CMTS uses as a default value the value of the Nominal Grant Interval
    // parameter.  In all cases, the CMTS reports the value it is using when the
    // parameter is applicable.  The CM reports the  signaled parameter value if
    // it was signaled,  and 0 otherwise.  If the referenced parameter is not
    // applicable to the direction or scheduling type of the corresponding DOCSIS
    // QOS Parameter Set, both CMTS and CM report this object's value as 0. The
    // type is interface{} with range: 0..4294967295. Units are microseconds.
    DocsQosParamSetNomPollInterval interface{}

    // Specifies the maximum amount of time in microseconds that the unicast
    // request interval may be delayed from the nominal periodic schedule on an
    // upstream service flow.  This parameter is applicable only to upstream
    // service flows with a Schedulingtype of realTimePollingService(4) or
    // unsolictedGrantServiceWithAD(5).  If the referenced parameter is applicable
    // but not present in the corresponding DOCSIS QOS Parameter Set, the CMTS
    // uses an implementation dependent  value and reports the value it is using.
    // The CM reports a value of 0 in this case.  If the parameter is not
    // applicable to the direction or upstream scheduling type of the service
    // flow, both CMTS and CM report this object's value as 0. The type is
    // interface{} with range: 0..4294967295. Units are microseconds.
    DocsQosParamSetTolPollJitter interface{}

    // Specifies the unsolicited grant size in bytes.  The grant size includes the
    // entire MAC frame  data PDU from the Frame Control byte to end of the MAC
    // frame.  The referenced parameter is applicable only  for upstream flows
    // with a SchedulingType of of unsolicitedGrantServicewithAD(5) or
    // unsolicitedGrantService(6), and is mandatory when applicable. Both CMTS and
    // CM report the signaled value of the parameter in this case.                
    // If the referenced parameter is not applicable to the direction or
    // scheduling type of the corresponding DOCSIS QOS Parameter Set, both CMTS
    // and CM report this object's value as 0. The type is interface{} with range:
    // 0..65535.
    DocsQosParamSetUnsolicitGrantSize interface{}

    // Specifies the nominal interval in microseconds  between successive data
    // grant opportunities  on an upstream service flow.  The referenced parameter
    // is applicable only  for upstream flows with a SchedulingType of of
    // unsolicitedGrantServicewithAD(5) or unsolicitedGrantService(6), and is
    // mandatory when applicable. Both CMTS and CM report the signaled value of
    // the parameter in this case.  If the referenced parameter is not applicable
    // to the direction or scheduling type of the corresponding DOCSIS QOS
    // Parameter Set, both CMTS and CM report this object's value as 0. The type
    // is interface{} with range: 0..4294967295. Units are microseconds.
    DocsQosParamSetNomGrantInterval interface{}

    // Specifies the maximum amount of time in microseconds that the transmission
    // opportunities may be delayed from the nominal periodic schedule.   The
    // referenced parameter is applicable only  for upstream flows with a
    // SchedulingType of of unsolicitedGrantServicewithAD(5) or
    // unsolicitedGrantService(6), and is mandatory when applicable. Both CMTS and
    // CM report the signaled value of the parameter in this case.  If the
    // referenced parameter is not applicable to the direction or scheduling type
    // of the corresponding DOCSIS QOS Parameter Set, both CMTS and CM report this
    // object's value as 0. The type is interface{} with range: 0..4294967295.
    // Units are microseconds.
    DocsQosParamSetTolGrantJitter interface{}

    // Specifies the number of data grants per Nominal  Grant Interval 
    // (docsQosParamSetNomGrantInterval).  The referenced parameter is applicable
    // only  for upstream flows with a SchedulingType of of
    // unsolicitedGrantServicewithAD(5) or unsolicitedGrantService(6), and is
    // mandatory when applicable. Both CMTS and CM report the signaled value of
    // the parameter in this case.  If the referenced parameter is not applicable
    // to the direction or scheduling type of the corresponding DOCSIS QOS
    // Parameter Set, both CMTS and CM report this object's value as 0. The type
    // is interface{} with range: 0..127.
    DocsQosParamSetGrantsPerInterval interface{}

    // Specifies the AND mask for IP TOS byte for overwriting IP packets TOS
    // value.  The IP packets TOS byte is  bitwise ANDed with
    // docsQosParamSetTosAndMask and  result is bitwise ORed with
    // docsQosParamSetTosORMask and result is written to IP packet TOS byte.  A
    // value of 'FF'H for docsQosParamSetTosAndMask and a value of '00'H for
    // docsQosParamSetTosOrMask means  that IP Packet TOS byte is not overwritten.
    // Even though the this object is only enforced by the Cable Modem Termination
    // System (CMTS), Cable Modems must report the value as signaled in the
    // referenced parameter.  This combination is reported if the referenced
    // parameter is not present in a QOS Parameter Set. The type is string with
    // length: 1.
    DocsQosParamSetTosAndMask interface{}

    // Specifies the OR mask for IP TOS byte. See the description of
    // docsQosParamSetTosAndMask for further details. The type is string with
    // length: 1.
    DocsQosParamSetTosOrMask interface{}

    // Specifies the maximum latency between the reception of a packet by the CMTS
    // on its NSI  and the forwarding of the packet to the RF interface. A value
    // of 0 signifies no maximum latency enforced. This object only applies to
    // downstream service flows.  If the referenced parameter is not present in
    // the corresponding downstream DOCSIS QOS Parameter Set,  the default value
    // is 0. This parameter is not applicable to upstream DOCSIS QOS Parameter
    // Sets, and its value is reported as 0 in this case. The type is interface{}
    // with range: 0..4294967295. Units are microseconds.
    DocsQosParamSetMaxLatency interface{}

    // Specifies which transmit interval opportunities  the CM omits for upstream
    // transmission requests and  packet transmissions. This object takes its
    // default value for downstream service flows.  Unless otherwise indicated, a
    // bit value of 1 means that a CM must *not* use that opportunity for 
    // upstream transmission.  Calling bit 0 the least significant bit of the 
    // least significant (4th) octet, and increasing bit number with significance,
    // the bit definitions are as defined below:  broadcastReqOpp(0):      all CMs
    // broadcast request opportunities  priorityReqMulticastReq(1):      priority
    // request multicast request opportunities  reqDataForReq(2):     
    // request/data opportunities for requests  reqDataForData(3):     
    // request/data opportunities for data  piggybackReqWithData(4):     
    // piggyback requests with data  concatenateData(5):      concatenate data 
    // fragmentData(6):      fragment data  suppresspayloadheaders(7):      
    // suppress payload headers  dropPktsExceedUGSize(8):      A value of 1 mean
    // that service flow must drop      packet that do not fit in the Unsolicited 
    // Grant size   If the referenced parameter is not present in  a QOS Parameter
    // Set, the value of this object is reported as '00000000'H. The type is
    // string with length: 4.
    DocsQosParamSetRequestPolicyOct interface{}

    // This object indicates the set of QOS Parameter Set parameters actually
    // signaled in the  DOCSIS registration or dynamic service request message
    // that created or modified the QOS Parameter Set.  A bit is set to 1 when the
    // parameter described by the indicated reference section is present in the
    // original request.    Note that when Service Class names are expanded, the
    // registration or dynamic response message may contain parameters as expanded
    // by the CMTS based on a stored service class. These expanded parameters are
    // *not* indicated by a 1 bit in this object.  Note that even though some QOS
    // Parameter Set  parameters may not be signalled in a message (so that the
    // paramater's bit in this object is 0) the DOCSIS specification calls for
    // default values to be used. These default values are reported as the
    // corresponding object's value in the row.   Note that BITS objects are
    // encoded most significant bit first. For example, if bits 1 and 16 are set,
    // the value of this object  is the octet string '400080'H. The type is
    // map[string]bool.
    DocsQosParamSetBitMap interface{}
}

func (docsQosParamSetEntry *DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry) GetEntityData() *types.CommonEntityData {
    docsQosParamSetEntry.EntityData.YFilter = docsQosParamSetEntry.YFilter
    docsQosParamSetEntry.EntityData.YangName = "docsQosParamSetEntry"
    docsQosParamSetEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosParamSetEntry.EntityData.ParentYangName = "docsQosParamSetTable"
    docsQosParamSetEntry.EntityData.SegmentPath = "docsQosParamSetEntry" + types.AddKeyToken(docsQosParamSetEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosParamSetEntry.DocsQosServiceFlowId, "docsQosServiceFlowId") + types.AddKeyToken(docsQosParamSetEntry.DocsQosParamSetType, "docsQosParamSetType")
    docsQosParamSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosParamSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosParamSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosParamSetEntry.EntityData.Children = types.NewOrderedMap()
    docsQosParamSetEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosParamSetEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosParamSetEntry.IfIndex})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsQosParamSetEntry.DocsQosServiceFlowId})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetType", types.YLeaf{"DocsQosParamSetType", docsQosParamSetEntry.DocsQosParamSetType})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetServiceClassName", types.YLeaf{"DocsQosParamSetServiceClassName", docsQosParamSetEntry.DocsQosParamSetServiceClassName})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetPriority", types.YLeaf{"DocsQosParamSetPriority", docsQosParamSetEntry.DocsQosParamSetPriority})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMaxTrafficRate", types.YLeaf{"DocsQosParamSetMaxTrafficRate", docsQosParamSetEntry.DocsQosParamSetMaxTrafficRate})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMaxTrafficBurst", types.YLeaf{"DocsQosParamSetMaxTrafficBurst", docsQosParamSetEntry.DocsQosParamSetMaxTrafficBurst})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMinReservedRate", types.YLeaf{"DocsQosParamSetMinReservedRate", docsQosParamSetEntry.DocsQosParamSetMinReservedRate})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMinReservedPkt", types.YLeaf{"DocsQosParamSetMinReservedPkt", docsQosParamSetEntry.DocsQosParamSetMinReservedPkt})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetActiveTimeout", types.YLeaf{"DocsQosParamSetActiveTimeout", docsQosParamSetEntry.DocsQosParamSetActiveTimeout})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetAdmittedTimeout", types.YLeaf{"DocsQosParamSetAdmittedTimeout", docsQosParamSetEntry.DocsQosParamSetAdmittedTimeout})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMaxConcatBurst", types.YLeaf{"DocsQosParamSetMaxConcatBurst", docsQosParamSetEntry.DocsQosParamSetMaxConcatBurst})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetSchedulingType", types.YLeaf{"DocsQosParamSetSchedulingType", docsQosParamSetEntry.DocsQosParamSetSchedulingType})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetNomPollInterval", types.YLeaf{"DocsQosParamSetNomPollInterval", docsQosParamSetEntry.DocsQosParamSetNomPollInterval})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetTolPollJitter", types.YLeaf{"DocsQosParamSetTolPollJitter", docsQosParamSetEntry.DocsQosParamSetTolPollJitter})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetUnsolicitGrantSize", types.YLeaf{"DocsQosParamSetUnsolicitGrantSize", docsQosParamSetEntry.DocsQosParamSetUnsolicitGrantSize})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetNomGrantInterval", types.YLeaf{"DocsQosParamSetNomGrantInterval", docsQosParamSetEntry.DocsQosParamSetNomGrantInterval})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetTolGrantJitter", types.YLeaf{"DocsQosParamSetTolGrantJitter", docsQosParamSetEntry.DocsQosParamSetTolGrantJitter})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetGrantsPerInterval", types.YLeaf{"DocsQosParamSetGrantsPerInterval", docsQosParamSetEntry.DocsQosParamSetGrantsPerInterval})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetTosAndMask", types.YLeaf{"DocsQosParamSetTosAndMask", docsQosParamSetEntry.DocsQosParamSetTosAndMask})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetTosOrMask", types.YLeaf{"DocsQosParamSetTosOrMask", docsQosParamSetEntry.DocsQosParamSetTosOrMask})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetMaxLatency", types.YLeaf{"DocsQosParamSetMaxLatency", docsQosParamSetEntry.DocsQosParamSetMaxLatency})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetRequestPolicyOct", types.YLeaf{"DocsQosParamSetRequestPolicyOct", docsQosParamSetEntry.DocsQosParamSetRequestPolicyOct})
    docsQosParamSetEntry.EntityData.Leafs.Append("docsQosParamSetBitMap", types.YLeaf{"DocsQosParamSetBitMap", docsQosParamSetEntry.DocsQosParamSetBitMap})

    docsQosParamSetEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId", "DocsQosParamSetType"}

    return &(docsQosParamSetEntry.EntityData)
}

// DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType represents the service flow.
type DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType string

const (
    DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType_active DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType = "active"

    DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType_admitted DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType = "admitted"

    DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType_provisioned DOCSQOSMIB_DocsQosParamSetTable_DocsQosParamSetEntry_DocsQosParamSetType = "provisioned"
)

// DOCSQOSMIB_DocsQosServiceFlowTable
// This table describes the set of Docsis-QOS 
// Service Flows in a managed device. 
type DOCSQOSMIB_DocsQosServiceFlowTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a service flow. An entry in the table exists for each  Service
    // Flow ID. The ifIndex is an   ifType of docsCableMaclayer(127). The type is
    // slice of DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry.
    DocsQosServiceFlowEntry []*DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry
}

func (docsQosServiceFlowTable *DOCSQOSMIB_DocsQosServiceFlowTable) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowTable.EntityData.YFilter = docsQosServiceFlowTable.YFilter
    docsQosServiceFlowTable.EntityData.YangName = "docsQosServiceFlowTable"
    docsQosServiceFlowTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosServiceFlowTable.EntityData.SegmentPath = "docsQosServiceFlowTable"
    docsQosServiceFlowTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowTable.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowTable.EntityData.Children.Append("docsQosServiceFlowEntry", types.YChild{"DocsQosServiceFlowEntry", nil})
    for i := range docsQosServiceFlowTable.DocsQosServiceFlowEntry {
        docsQosServiceFlowTable.EntityData.Children.Append(types.GetSegmentPath(docsQosServiceFlowTable.DocsQosServiceFlowEntry[i]), types.YChild{"DocsQosServiceFlowEntry", docsQosServiceFlowTable.DocsQosServiceFlowEntry[i]})
    }
    docsQosServiceFlowTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosServiceFlowTable.EntityData.YListKeys = []string {}

    return &(docsQosServiceFlowTable.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry
// Describes a service flow.
// An entry in the table exists for each 
// Service Flow ID. The ifIndex is an  
// ifType of docsCableMaclayer(127).
type DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An index assigned to a service flow by CMTS. The
    // type is interface{} with range: 1..4294967295.
    DocsQosServiceFlowId interface{}

    // This object is obsolete. The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowProvisionedParamSetIndex interface{}

    // This object is obsolete. The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowAdmittedParamSetIndex interface{}

    // This object is obsolete. The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowActiveParamSetIndex interface{}

    // Service Identifier (SID) assigned to an  admitted or active service flow.
    // This object reports a value of 0 if a Service Id is not  associated with
    // the service flow. Only active  or admitted upstream service flows will have
    // a Service Id (SID). The type is interface{} with range: 0..16383.
    DocsQosServiceFlowSID interface{}

    // The direction of the service flow. The type is IfDirection.
    DocsQosServiceFlowDirection interface{}

    // Object reflects whether service flow is the primary  or a secondary service
    // flow.  A primary service flow is the default service flow for otherwise
    // unclassified traffic and all MAC  messages. The type is bool.
    DocsQosServiceFlowPrimary interface{}

    // This object is obsolete. The type is interface{} with range: 0..65535.
    // Units are seconds.
    DocsQosServiceFlowActiveTimeout interface{}

    // This object is obsolete. The type is interface{} with range: 0..65535.
    // Units are seconds.
    DocsQosServiceFlowAdmittedTimeout interface{}

    // This object is obsolete. The type is SchedulingType.
    DocsQosServiceFlowSchedulingType interface{}

    // This object is obsolete. The type is string with length: 4.
    DocsQosServiceFlowRequestPolicy interface{}

    // This object is obsolete. The type is string with length: 1.
    DocsQosServiceFlowTosAndMask interface{}

    // This object is obsolete. The type is string with length: 1.
    DocsQosServiceFlowTosOrMask interface{}
}

func (docsQosServiceFlowEntry *DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowEntry.EntityData.YFilter = docsQosServiceFlowEntry.YFilter
    docsQosServiceFlowEntry.EntityData.YangName = "docsQosServiceFlowEntry"
    docsQosServiceFlowEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowEntry.EntityData.ParentYangName = "docsQosServiceFlowTable"
    docsQosServiceFlowEntry.EntityData.SegmentPath = "docsQosServiceFlowEntry" + types.AddKeyToken(docsQosServiceFlowEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosServiceFlowEntry.DocsQosServiceFlowId, "docsQosServiceFlowId")
    docsQosServiceFlowEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowEntry.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosServiceFlowEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosServiceFlowEntry.IfIndex})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsQosServiceFlowEntry.DocsQosServiceFlowId})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowProvisionedParamSetIndex", types.YLeaf{"DocsQosServiceFlowProvisionedParamSetIndex", docsQosServiceFlowEntry.DocsQosServiceFlowProvisionedParamSetIndex})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowAdmittedParamSetIndex", types.YLeaf{"DocsQosServiceFlowAdmittedParamSetIndex", docsQosServiceFlowEntry.DocsQosServiceFlowAdmittedParamSetIndex})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowActiveParamSetIndex", types.YLeaf{"DocsQosServiceFlowActiveParamSetIndex", docsQosServiceFlowEntry.DocsQosServiceFlowActiveParamSetIndex})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowSID", types.YLeaf{"DocsQosServiceFlowSID", docsQosServiceFlowEntry.DocsQosServiceFlowSID})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowDirection", types.YLeaf{"DocsQosServiceFlowDirection", docsQosServiceFlowEntry.DocsQosServiceFlowDirection})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowPrimary", types.YLeaf{"DocsQosServiceFlowPrimary", docsQosServiceFlowEntry.DocsQosServiceFlowPrimary})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowActiveTimeout", types.YLeaf{"DocsQosServiceFlowActiveTimeout", docsQosServiceFlowEntry.DocsQosServiceFlowActiveTimeout})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowAdmittedTimeout", types.YLeaf{"DocsQosServiceFlowAdmittedTimeout", docsQosServiceFlowEntry.DocsQosServiceFlowAdmittedTimeout})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowSchedulingType", types.YLeaf{"DocsQosServiceFlowSchedulingType", docsQosServiceFlowEntry.DocsQosServiceFlowSchedulingType})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowRequestPolicy", types.YLeaf{"DocsQosServiceFlowRequestPolicy", docsQosServiceFlowEntry.DocsQosServiceFlowRequestPolicy})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowTosAndMask", types.YLeaf{"DocsQosServiceFlowTosAndMask", docsQosServiceFlowEntry.DocsQosServiceFlowTosAndMask})
    docsQosServiceFlowEntry.EntityData.Leafs.Append("docsQosServiceFlowTosOrMask", types.YLeaf{"DocsQosServiceFlowTosOrMask", docsQosServiceFlowEntry.DocsQosServiceFlowTosOrMask})

    docsQosServiceFlowEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId"}

    return &(docsQosServiceFlowEntry.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowStatsTable
// This table describes statistics associated with the  
// Service Flows in a managed device. 
type DOCSQOSMIB_DocsQosServiceFlowStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a set of service flow statistics. An entry in the table exists
    // for each  Service Flow ID. The ifIndex is an   ifType of
    // docsCableMaclayer(127). The type is slice of
    // DOCSQOSMIB_DocsQosServiceFlowStatsTable_DocsQosServiceFlowStatsEntry.
    DocsQosServiceFlowStatsEntry []*DOCSQOSMIB_DocsQosServiceFlowStatsTable_DocsQosServiceFlowStatsEntry
}

func (docsQosServiceFlowStatsTable *DOCSQOSMIB_DocsQosServiceFlowStatsTable) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowStatsTable.EntityData.YFilter = docsQosServiceFlowStatsTable.YFilter
    docsQosServiceFlowStatsTable.EntityData.YangName = "docsQosServiceFlowStatsTable"
    docsQosServiceFlowStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowStatsTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosServiceFlowStatsTable.EntityData.SegmentPath = "docsQosServiceFlowStatsTable"
    docsQosServiceFlowStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowStatsTable.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowStatsTable.EntityData.Children.Append("docsQosServiceFlowStatsEntry", types.YChild{"DocsQosServiceFlowStatsEntry", nil})
    for i := range docsQosServiceFlowStatsTable.DocsQosServiceFlowStatsEntry {
        docsQosServiceFlowStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsQosServiceFlowStatsTable.DocsQosServiceFlowStatsEntry[i]), types.YChild{"DocsQosServiceFlowStatsEntry", docsQosServiceFlowStatsTable.DocsQosServiceFlowStatsEntry[i]})
    }
    docsQosServiceFlowStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosServiceFlowStatsTable.EntityData.YListKeys = []string {}

    return &(docsQosServiceFlowStatsTable.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowStatsTable_DocsQosServiceFlowStatsEntry
// Describes a set of service flow statistics.
// An entry in the table exists for each 
// Service Flow ID. The ifIndex is an  
// ifType of docsCableMaclayer(127).
type DOCSQOSMIB_DocsQosServiceFlowStatsTable_DocsQosServiceFlowStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // The number of Packet Data PDUs classified to this service flow. This object
    // does not count MAC-specific management messages. CMs not classifying
    // downstream packets may report this object's value as 0. . The type is
    // interface{} with range: 0..4294967295.
    DocsQosServiceFlowPkts interface{}

    // The number of octets transmitted on the Docsis RF network from the byte
    // after the MAC header HCS to the end of the CRC for all packets counted in
    // the docsQosServiceFlowPkts object for this row. Note that this counts the
    // octets after payload header suppression has been applied. CMs not
    // classifying to a downstream service flow may report this object's value as
    // 0 for that flow. . The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowOctets interface{}

    // The value of sysUpTime when the service flow  was created. The type is
    // interface{} with range: 0..4294967295.
    DocsQosServiceFlowTimeCreated interface{}

    // The total time that service flow has been active. The type is interface{}
    // with range: 0..4294967295. Units are seconds.
    DocsQosServiceFlowTimeActive interface{}

    // The number of packets received on the service flow with an unknown payload
    // header suppression index. The type is interface{} with range:
    // 0..4294967295.
    DocsQosServiceFlowPHSUnknowns interface{}

    // The number of packets dropped due to policing of  the service flow,
    // especially to limit the maximum  rate of the flow. The type is interface{}
    // with range: 0..4294967295.
    DocsQosServiceFlowPolicedDropPkts interface{}

    // The number of packet delayed due to policing of  the service flow,
    // especially to limit the maximum rate of the flow. The type is interface{}
    // with range: 0..4294967295.
    DocsQosServiceFlowPolicedDelayPkts interface{}
}

func (docsQosServiceFlowStatsEntry *DOCSQOSMIB_DocsQosServiceFlowStatsTable_DocsQosServiceFlowStatsEntry) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowStatsEntry.EntityData.YFilter = docsQosServiceFlowStatsEntry.YFilter
    docsQosServiceFlowStatsEntry.EntityData.YangName = "docsQosServiceFlowStatsEntry"
    docsQosServiceFlowStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowStatsEntry.EntityData.ParentYangName = "docsQosServiceFlowStatsTable"
    docsQosServiceFlowStatsEntry.EntityData.SegmentPath = "docsQosServiceFlowStatsEntry" + types.AddKeyToken(docsQosServiceFlowStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosServiceFlowStatsEntry.DocsQosServiceFlowId, "docsQosServiceFlowId")
    docsQosServiceFlowStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosServiceFlowStatsEntry.IfIndex})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsQosServiceFlowStatsEntry.DocsQosServiceFlowId})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowPkts", types.YLeaf{"DocsQosServiceFlowPkts", docsQosServiceFlowStatsEntry.DocsQosServiceFlowPkts})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowOctets", types.YLeaf{"DocsQosServiceFlowOctets", docsQosServiceFlowStatsEntry.DocsQosServiceFlowOctets})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowTimeCreated", types.YLeaf{"DocsQosServiceFlowTimeCreated", docsQosServiceFlowStatsEntry.DocsQosServiceFlowTimeCreated})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowTimeActive", types.YLeaf{"DocsQosServiceFlowTimeActive", docsQosServiceFlowStatsEntry.DocsQosServiceFlowTimeActive})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowPHSUnknowns", types.YLeaf{"DocsQosServiceFlowPHSUnknowns", docsQosServiceFlowStatsEntry.DocsQosServiceFlowPHSUnknowns})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowPolicedDropPkts", types.YLeaf{"DocsQosServiceFlowPolicedDropPkts", docsQosServiceFlowStatsEntry.DocsQosServiceFlowPolicedDropPkts})
    docsQosServiceFlowStatsEntry.EntityData.Leafs.Append("docsQosServiceFlowPolicedDelayPkts", types.YLeaf{"DocsQosServiceFlowPolicedDelayPkts", docsQosServiceFlowStatsEntry.DocsQosServiceFlowPolicedDelayPkts})

    docsQosServiceFlowStatsEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId"}

    return &(docsQosServiceFlowStatsEntry.EntityData)
}

// DOCSQOSMIB_DocsQosUpstreamStatsTable
// This table describes statistics associated with  
// upstream service flows. All counted frames must 
// be received without an FCS error.
type DOCSQOSMIB_DocsQosUpstreamStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a set of upstream service flow statistics. An entry in the table
    // exists for each  upstream Service Flow in a managed device.  The ifIndex is
    // an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSQOSMIB_DocsQosUpstreamStatsTable_DocsQosUpstreamStatsEntry.
    DocsQosUpstreamStatsEntry []*DOCSQOSMIB_DocsQosUpstreamStatsTable_DocsQosUpstreamStatsEntry
}

func (docsQosUpstreamStatsTable *DOCSQOSMIB_DocsQosUpstreamStatsTable) GetEntityData() *types.CommonEntityData {
    docsQosUpstreamStatsTable.EntityData.YFilter = docsQosUpstreamStatsTable.YFilter
    docsQosUpstreamStatsTable.EntityData.YangName = "docsQosUpstreamStatsTable"
    docsQosUpstreamStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosUpstreamStatsTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosUpstreamStatsTable.EntityData.SegmentPath = "docsQosUpstreamStatsTable"
    docsQosUpstreamStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosUpstreamStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosUpstreamStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosUpstreamStatsTable.EntityData.Children = types.NewOrderedMap()
    docsQosUpstreamStatsTable.EntityData.Children.Append("docsQosUpstreamStatsEntry", types.YChild{"DocsQosUpstreamStatsEntry", nil})
    for i := range docsQosUpstreamStatsTable.DocsQosUpstreamStatsEntry {
        docsQosUpstreamStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsQosUpstreamStatsTable.DocsQosUpstreamStatsEntry[i]), types.YChild{"DocsQosUpstreamStatsEntry", docsQosUpstreamStatsTable.DocsQosUpstreamStatsEntry[i]})
    }
    docsQosUpstreamStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosUpstreamStatsTable.EntityData.YListKeys = []string {}

    return &(docsQosUpstreamStatsTable.EntityData)
}

// DOCSQOSMIB_DocsQosUpstreamStatsTable_DocsQosUpstreamStatsEntry
// Describes a set of upstream service flow statistics.
// An entry in the table exists for each 
// upstream Service Flow in a managed device. 
// The ifIndex is an ifType of docsCableMaclayer(127).
type DOCSQOSMIB_DocsQosUpstreamStatsTable_DocsQosUpstreamStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Identifies a service id for an admitted or active 
    // upstream service flow. The type is interface{} with range: 1..16383.
    DocsQosSID interface{}

    // The number of fragmentation headers received on an upstream  service flow,
    // regardless of whether the fragment was correctly reassembled into a  valid
    // packet. . The type is interface{} with range: 0..4294967295.
    DocsQosUpstreamFragments interface{}

    // The number of upstream fragments discarded and not  assembled into a valid
    // upstream packet. The type is interface{} with range: 0..4294967295.
    DocsQosUpstreamFragDiscards interface{}

    // The number of concatenation headers received on an  upstream service flow.
    // The type is interface{} with range: 0..4294967295.
    DocsQosUpstreamConcatBursts interface{}
}

func (docsQosUpstreamStatsEntry *DOCSQOSMIB_DocsQosUpstreamStatsTable_DocsQosUpstreamStatsEntry) GetEntityData() *types.CommonEntityData {
    docsQosUpstreamStatsEntry.EntityData.YFilter = docsQosUpstreamStatsEntry.YFilter
    docsQosUpstreamStatsEntry.EntityData.YangName = "docsQosUpstreamStatsEntry"
    docsQosUpstreamStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosUpstreamStatsEntry.EntityData.ParentYangName = "docsQosUpstreamStatsTable"
    docsQosUpstreamStatsEntry.EntityData.SegmentPath = "docsQosUpstreamStatsEntry" + types.AddKeyToken(docsQosUpstreamStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosUpstreamStatsEntry.DocsQosSID, "docsQosSID")
    docsQosUpstreamStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosUpstreamStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosUpstreamStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosUpstreamStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsQosUpstreamStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosUpstreamStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosUpstreamStatsEntry.IfIndex})
    docsQosUpstreamStatsEntry.EntityData.Leafs.Append("docsQosSID", types.YLeaf{"DocsQosSID", docsQosUpstreamStatsEntry.DocsQosSID})
    docsQosUpstreamStatsEntry.EntityData.Leafs.Append("docsQosUpstreamFragments", types.YLeaf{"DocsQosUpstreamFragments", docsQosUpstreamStatsEntry.DocsQosUpstreamFragments})
    docsQosUpstreamStatsEntry.EntityData.Leafs.Append("docsQosUpstreamFragDiscards", types.YLeaf{"DocsQosUpstreamFragDiscards", docsQosUpstreamStatsEntry.DocsQosUpstreamFragDiscards})
    docsQosUpstreamStatsEntry.EntityData.Leafs.Append("docsQosUpstreamConcatBursts", types.YLeaf{"DocsQosUpstreamConcatBursts", docsQosUpstreamStatsEntry.DocsQosUpstreamConcatBursts})

    docsQosUpstreamStatsEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosSID"}

    return &(docsQosUpstreamStatsEntry.EntityData)
}

// DOCSQOSMIB_DocsQosDynamicServiceStatsTable
// This table describes statistics associated with the  
// Dynamic Service Flows in a managed device. 
type DOCSQOSMIB_DocsQosDynamicServiceStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a set of dynamic service flow statistics. Two entries exist for
    // each Docsis mac layer  interface for the upstream and downstream direction.
    // On the CMTS, the downstream direction row indicates messages transmitted or
    // transactions originated by the CMTS. The upstream direction row indicates
    // messages received or transaction originated by the CM. On the CM, the
    // downstream direction row  indicates messages received or transactions
    // originated by the CMTS. The upstream direction  row indicates messages
    // transmitted by the CM or transactions originated by the CM. The ifIndex is
    // an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSQOSMIB_DocsQosDynamicServiceStatsTable_DocsQosDynamicServiceStatsEntry.
    DocsQosDynamicServiceStatsEntry []*DOCSQOSMIB_DocsQosDynamicServiceStatsTable_DocsQosDynamicServiceStatsEntry
}

func (docsQosDynamicServiceStatsTable *DOCSQOSMIB_DocsQosDynamicServiceStatsTable) GetEntityData() *types.CommonEntityData {
    docsQosDynamicServiceStatsTable.EntityData.YFilter = docsQosDynamicServiceStatsTable.YFilter
    docsQosDynamicServiceStatsTable.EntityData.YangName = "docsQosDynamicServiceStatsTable"
    docsQosDynamicServiceStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosDynamicServiceStatsTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosDynamicServiceStatsTable.EntityData.SegmentPath = "docsQosDynamicServiceStatsTable"
    docsQosDynamicServiceStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosDynamicServiceStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosDynamicServiceStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosDynamicServiceStatsTable.EntityData.Children = types.NewOrderedMap()
    docsQosDynamicServiceStatsTable.EntityData.Children.Append("docsQosDynamicServiceStatsEntry", types.YChild{"DocsQosDynamicServiceStatsEntry", nil})
    for i := range docsQosDynamicServiceStatsTable.DocsQosDynamicServiceStatsEntry {
        docsQosDynamicServiceStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsQosDynamicServiceStatsTable.DocsQosDynamicServiceStatsEntry[i]), types.YChild{"DocsQosDynamicServiceStatsEntry", docsQosDynamicServiceStatsTable.DocsQosDynamicServiceStatsEntry[i]})
    }
    docsQosDynamicServiceStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosDynamicServiceStatsTable.EntityData.YListKeys = []string {}

    return &(docsQosDynamicServiceStatsTable.EntityData)
}

// DOCSQOSMIB_DocsQosDynamicServiceStatsTable_DocsQosDynamicServiceStatsEntry
// Describes a set of dynamic service flow statistics.
// Two entries exist for each Docsis mac layer 
// interface for the upstream and downstream direction.
// On the CMTS, the downstream direction row indicates
// messages transmitted or transactions originated
// by the CMTS. The upstream direction row indicates
// messages received or transaction originated by the
// CM. On the CM, the downstream direction row 
// indicates messages received or transactions
// originated by the CMTS. The upstream direction 
// row indicates messages transmitted by the CM or
// transactions originated by the CM.
// The ifIndex is an ifType of docsCableMaclayer(127).
type DOCSQOSMIB_DocsQosDynamicServiceStatsTable_DocsQosDynamicServiceStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The direction of interface. The type is
    // IfDirection.
    DocsQosIfDirection interface{}

    // The number of Dynamic Service Addition Requests, including retries. The
    // type is interface{} with range: 0..4294967295.
    DocsQosDSAReqs interface{}

    // The number of Dynamic Service Addition Responses, including retries. The
    // type is interface{} with range: 0..4294967295.
    DocsQosDSARsps interface{}

    // The number of Dynamic Service Addition Acknowledgements, including retries.
    // The type is interface{} with range: 0..4294967295.
    DocsQosDSAAcks interface{}

    // The number of Dynamic Service Change Requests, including retries. The type
    // is interface{} with range: 0..4294967295.
    DocsQosDSCReqs interface{}

    // The number of Dynamic Service Change Responses, including retries. The type
    // is interface{} with range: 0..4294967295.
    DocsQosDSCRsps interface{}

    // The number of Dynamic Service Change Acknowledgements, including retries.
    // The type is interface{} with range: 0..4294967295.
    DocsQosDSCAcks interface{}

    // The number of Dynamic Service Delete Requests, including retries. The type
    // is interface{} with range: 0..4294967295.
    DocsQosDSDReqs interface{}

    // The number of Dynamic Service Delete Responses, including retries. The type
    // is interface{} with range: 0..4294967295.
    DocsQosDSDRsps interface{}

    // The number of successful Dynamic Service Addition transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicAdds interface{}

    // The number of failed Dynamic Service Addition transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicAddFails interface{}

    // The number of successful Dynamic Service Change transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicChanges interface{}

    // The number of failed Dynamic Service Change transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicChangeFails interface{}

    // The number of successful Dynamic Service Delete transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicDeletes interface{}

    // The number of failed Dynamic Service Delete transactions. The type is
    // interface{} with range: 0..4294967295.
    DocsQosDynamicDeleteFails interface{}

    // The number of Dynamic Channel Change Request messages traversing an
    // interface. This count is nonzero only on downstream direction rows. This
    // count should include number of retries. The type is interface{} with range:
    // 0..4294967295.
    DocsQosDCCReqs interface{}

    // The number of Dynamic Channel Change Response messages traversing an
    // interface. This count is nonzero only on upstream direction rows. This
    // count should include number of retries. The type is interface{} with range:
    // 0..4294967295.
    DocsQosDCCRsps interface{}

    // The number of Dynamic Channel Change Acknowledgement messages traversing an
    // interface. This count is nonzero only on downstream direction rows. This
    // count should include number of retries. The type is interface{} with range:
    // 0..4294967295.
    DocsQosDCCAcks interface{}

    // The number of successful Dynamic Channel Change transactions. This count is
    // nonzero only on downstream direction rows. The type is interface{} with
    // range: 0..4294967295.
    DocsQosDCCs interface{}

    // The number of failed Dynamic Channel Change transactions. This count is
    // nonzero only on downstream direction rows. The type is interface{} with
    // range: 0..4294967295.
    DocsQosDCCFails interface{}

    // The number of Dynamic Channel Change Response (depart) messages traversing
    // an interface. This count is only counted  on upstream direction rows. The
    // type is interface{} with range: 0..4294967295.
    DocsQosDCCRspDeparts interface{}

    // The number of Dynamic Channel Change Response (arrive) messages traversing
    // an interface. This count is only counted on upstream direction rows. This
    // count should include number of retries. The type is interface{} with range:
    // 0..4294967295.
    DocsQosDCCRspArrives interface{}
}

func (docsQosDynamicServiceStatsEntry *DOCSQOSMIB_DocsQosDynamicServiceStatsTable_DocsQosDynamicServiceStatsEntry) GetEntityData() *types.CommonEntityData {
    docsQosDynamicServiceStatsEntry.EntityData.YFilter = docsQosDynamicServiceStatsEntry.YFilter
    docsQosDynamicServiceStatsEntry.EntityData.YangName = "docsQosDynamicServiceStatsEntry"
    docsQosDynamicServiceStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosDynamicServiceStatsEntry.EntityData.ParentYangName = "docsQosDynamicServiceStatsTable"
    docsQosDynamicServiceStatsEntry.EntityData.SegmentPath = "docsQosDynamicServiceStatsEntry" + types.AddKeyToken(docsQosDynamicServiceStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosDynamicServiceStatsEntry.DocsQosIfDirection, "docsQosIfDirection")
    docsQosDynamicServiceStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosDynamicServiceStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosDynamicServiceStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosDynamicServiceStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsQosDynamicServiceStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosDynamicServiceStatsEntry.IfIndex})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosIfDirection", types.YLeaf{"DocsQosIfDirection", docsQosDynamicServiceStatsEntry.DocsQosIfDirection})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSAReqs", types.YLeaf{"DocsQosDSAReqs", docsQosDynamicServiceStatsEntry.DocsQosDSAReqs})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSARsps", types.YLeaf{"DocsQosDSARsps", docsQosDynamicServiceStatsEntry.DocsQosDSARsps})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSAAcks", types.YLeaf{"DocsQosDSAAcks", docsQosDynamicServiceStatsEntry.DocsQosDSAAcks})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSCReqs", types.YLeaf{"DocsQosDSCReqs", docsQosDynamicServiceStatsEntry.DocsQosDSCReqs})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSCRsps", types.YLeaf{"DocsQosDSCRsps", docsQosDynamicServiceStatsEntry.DocsQosDSCRsps})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSCAcks", types.YLeaf{"DocsQosDSCAcks", docsQosDynamicServiceStatsEntry.DocsQosDSCAcks})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSDReqs", types.YLeaf{"DocsQosDSDReqs", docsQosDynamicServiceStatsEntry.DocsQosDSDReqs})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDSDRsps", types.YLeaf{"DocsQosDSDRsps", docsQosDynamicServiceStatsEntry.DocsQosDSDRsps})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicAdds", types.YLeaf{"DocsQosDynamicAdds", docsQosDynamicServiceStatsEntry.DocsQosDynamicAdds})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicAddFails", types.YLeaf{"DocsQosDynamicAddFails", docsQosDynamicServiceStatsEntry.DocsQosDynamicAddFails})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicChanges", types.YLeaf{"DocsQosDynamicChanges", docsQosDynamicServiceStatsEntry.DocsQosDynamicChanges})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicChangeFails", types.YLeaf{"DocsQosDynamicChangeFails", docsQosDynamicServiceStatsEntry.DocsQosDynamicChangeFails})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicDeletes", types.YLeaf{"DocsQosDynamicDeletes", docsQosDynamicServiceStatsEntry.DocsQosDynamicDeletes})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDynamicDeleteFails", types.YLeaf{"DocsQosDynamicDeleteFails", docsQosDynamicServiceStatsEntry.DocsQosDynamicDeleteFails})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCReqs", types.YLeaf{"DocsQosDCCReqs", docsQosDynamicServiceStatsEntry.DocsQosDCCReqs})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCRsps", types.YLeaf{"DocsQosDCCRsps", docsQosDynamicServiceStatsEntry.DocsQosDCCRsps})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCAcks", types.YLeaf{"DocsQosDCCAcks", docsQosDynamicServiceStatsEntry.DocsQosDCCAcks})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCs", types.YLeaf{"DocsQosDCCs", docsQosDynamicServiceStatsEntry.DocsQosDCCs})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCFails", types.YLeaf{"DocsQosDCCFails", docsQosDynamicServiceStatsEntry.DocsQosDCCFails})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCRspDeparts", types.YLeaf{"DocsQosDCCRspDeparts", docsQosDynamicServiceStatsEntry.DocsQosDCCRspDeparts})
    docsQosDynamicServiceStatsEntry.EntityData.Leafs.Append("docsQosDCCRspArrives", types.YLeaf{"DocsQosDCCRspArrives", docsQosDynamicServiceStatsEntry.DocsQosDCCRspArrives})

    docsQosDynamicServiceStatsEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosIfDirection"}

    return &(docsQosDynamicServiceStatsEntry.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowLogTable
// This table contains a log of the disconnected
// Service Flows in a managed device.
type DOCSQOSMIB_DocsQosServiceFlowLogTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single disconnected  service flow. The type is
    // slice of DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry.
    DocsQosServiceFlowLogEntry []*DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry
}

func (docsQosServiceFlowLogTable *DOCSQOSMIB_DocsQosServiceFlowLogTable) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowLogTable.EntityData.YFilter = docsQosServiceFlowLogTable.YFilter
    docsQosServiceFlowLogTable.EntityData.YangName = "docsQosServiceFlowLogTable"
    docsQosServiceFlowLogTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowLogTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosServiceFlowLogTable.EntityData.SegmentPath = "docsQosServiceFlowLogTable"
    docsQosServiceFlowLogTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowLogTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowLogTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowLogTable.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowLogTable.EntityData.Children.Append("docsQosServiceFlowLogEntry", types.YChild{"DocsQosServiceFlowLogEntry", nil})
    for i := range docsQosServiceFlowLogTable.DocsQosServiceFlowLogEntry {
        docsQosServiceFlowLogTable.EntityData.Children.Append(types.GetSegmentPath(docsQosServiceFlowLogTable.DocsQosServiceFlowLogEntry[i]), types.YChild{"DocsQosServiceFlowLogEntry", docsQosServiceFlowLogTable.DocsQosServiceFlowLogEntry[i]})
    }
    docsQosServiceFlowLogTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosServiceFlowLogTable.EntityData.YListKeys = []string {}

    return &(docsQosServiceFlowLogTable.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry
// The information regarding a single disconnected 
// service flow.
type DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for a logged service flow. The type
    // is interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogIndex interface{}

    // The ifIndex of ifType docsCableMaclayer(127)  on the CMTS where the service
    // flow was present. The type is interface{} with range: 1..2147483647.
    DocsQosServiceFlowLogIfIndex interface{}

    // The index assigned to the service flow by the CMTS. The type is interface{}
    // with range: 1..4294967295.
    DocsQosServiceFlowLogSFID interface{}

    // The MAC address for the cable modem associated with  the service flow. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsQosServiceFlowLogCmMac interface{}

    // The number of packets counted on this service flow  after payload header
    // suppression. The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogPkts interface{}

    // The number of octets counted on this service flow  after payload header
    // suppression. The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogOctets interface{}

    // The value of sysUpTime when the service flow  was deleted. The type is
    // interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogTimeDeleted interface{}

    // The value of sysUpTime when the service flow  was created. The type is
    // interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogTimeCreated interface{}

    // The total time that service flow was active. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    DocsQosServiceFlowLogTimeActive interface{}

    // The value of docsQosServiceFlowDirection  for the service flow. The type is
    // IfDirection.
    DocsQosServiceFlowLogDirection interface{}

    // The value of docsQosServiceFlowPrimary for the  service flow. The type is
    // bool.
    DocsQosServiceFlowLogPrimary interface{}

    // The value of docsQosParamSetServiceClassName for the provisioned QOS
    // Parameter Set of the  service flow. The type is string.
    DocsQosServiceFlowLogServiceClassName interface{}

    // The final value of docsQosServiceFlowPolicedDropPkts for the service flow.
    // The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogPolicedDropPkts interface{}

    // The final value of docsQosServiceFlowPolicedDelayPkts for the service flow.
    // The type is interface{} with range: 0..4294967295.
    DocsQosServiceFlowLogPolicedDelayPkts interface{}

    // Setting this object to the value destroy(6) removes this entry from the
    // table.  Reading this object return the value active(1). The type is
    // DocsQosServiceFlowLogControl.
    DocsQosServiceFlowLogControl interface{}
}

func (docsQosServiceFlowLogEntry *DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry) GetEntityData() *types.CommonEntityData {
    docsQosServiceFlowLogEntry.EntityData.YFilter = docsQosServiceFlowLogEntry.YFilter
    docsQosServiceFlowLogEntry.EntityData.YangName = "docsQosServiceFlowLogEntry"
    docsQosServiceFlowLogEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceFlowLogEntry.EntityData.ParentYangName = "docsQosServiceFlowLogTable"
    docsQosServiceFlowLogEntry.EntityData.SegmentPath = "docsQosServiceFlowLogEntry" + types.AddKeyToken(docsQosServiceFlowLogEntry.DocsQosServiceFlowLogIndex, "docsQosServiceFlowLogIndex")
    docsQosServiceFlowLogEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceFlowLogEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceFlowLogEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceFlowLogEntry.EntityData.Children = types.NewOrderedMap()
    docsQosServiceFlowLogEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogIndex", types.YLeaf{"DocsQosServiceFlowLogIndex", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogIndex})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogIfIndex", types.YLeaf{"DocsQosServiceFlowLogIfIndex", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogIfIndex})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogSFID", types.YLeaf{"DocsQosServiceFlowLogSFID", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogSFID})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogCmMac", types.YLeaf{"DocsQosServiceFlowLogCmMac", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogCmMac})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogPkts", types.YLeaf{"DocsQosServiceFlowLogPkts", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogPkts})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogOctets", types.YLeaf{"DocsQosServiceFlowLogOctets", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogOctets})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogTimeDeleted", types.YLeaf{"DocsQosServiceFlowLogTimeDeleted", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogTimeDeleted})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogTimeCreated", types.YLeaf{"DocsQosServiceFlowLogTimeCreated", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogTimeCreated})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogTimeActive", types.YLeaf{"DocsQosServiceFlowLogTimeActive", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogTimeActive})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogDirection", types.YLeaf{"DocsQosServiceFlowLogDirection", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogDirection})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogPrimary", types.YLeaf{"DocsQosServiceFlowLogPrimary", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogPrimary})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogServiceClassName", types.YLeaf{"DocsQosServiceFlowLogServiceClassName", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogServiceClassName})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogPolicedDropPkts", types.YLeaf{"DocsQosServiceFlowLogPolicedDropPkts", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogPolicedDropPkts})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogPolicedDelayPkts", types.YLeaf{"DocsQosServiceFlowLogPolicedDelayPkts", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogPolicedDelayPkts})
    docsQosServiceFlowLogEntry.EntityData.Leafs.Append("docsQosServiceFlowLogControl", types.YLeaf{"DocsQosServiceFlowLogControl", docsQosServiceFlowLogEntry.DocsQosServiceFlowLogControl})

    docsQosServiceFlowLogEntry.EntityData.YListKeys = []string {"DocsQosServiceFlowLogIndex"}

    return &(docsQosServiceFlowLogEntry.EntityData)
}

// DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl represents Reading this object return the value active(1).
type DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl string

const (
    DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl_active DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl = "active"

    DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl_destroy DOCSQOSMIB_DocsQosServiceFlowLogTable_DocsQosServiceFlowLogEntry_DocsQosServiceFlowLogControl = "destroy"
)

// DOCSQOSMIB_DocsQosServiceClassTable
// This table describes the set of Docsis-QOS 
// Service Classes in a CMTS. 
type DOCSQOSMIB_DocsQosServiceClassTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A provisioned service class on a CMTS.  Each entry defines a template for
    // certain  DOCSIS QOS Parameter Set values. When a CM  creates or modifies an
    // Admitted QOS Parameter Set for a Service Flow, it may reference a Service
    // Class Name instead of providing explicit QOS Parameter Set values. In this
    // case, the CMTS populates the QOS Parameter Set with the applicable 
    // corresponding values from the named Service Class. Subsequent changes to a
    // Service Class row do *not*  affect the QOS Parameter Set values of any
    // service flows already admitted.  A service class template applies to only a
    // single direction, as indicated in the  docsQosServiceClassDirection object.
    // The type is slice of
    // DOCSQOSMIB_DocsQosServiceClassTable_DocsQosServiceClassEntry.
    DocsQosServiceClassEntry []*DOCSQOSMIB_DocsQosServiceClassTable_DocsQosServiceClassEntry
}

func (docsQosServiceClassTable *DOCSQOSMIB_DocsQosServiceClassTable) GetEntityData() *types.CommonEntityData {
    docsQosServiceClassTable.EntityData.YFilter = docsQosServiceClassTable.YFilter
    docsQosServiceClassTable.EntityData.YangName = "docsQosServiceClassTable"
    docsQosServiceClassTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceClassTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosServiceClassTable.EntityData.SegmentPath = "docsQosServiceClassTable"
    docsQosServiceClassTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceClassTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceClassTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceClassTable.EntityData.Children = types.NewOrderedMap()
    docsQosServiceClassTable.EntityData.Children.Append("docsQosServiceClassEntry", types.YChild{"DocsQosServiceClassEntry", nil})
    for i := range docsQosServiceClassTable.DocsQosServiceClassEntry {
        docsQosServiceClassTable.EntityData.Children.Append(types.GetSegmentPath(docsQosServiceClassTable.DocsQosServiceClassEntry[i]), types.YChild{"DocsQosServiceClassEntry", docsQosServiceClassTable.DocsQosServiceClassEntry[i]})
    }
    docsQosServiceClassTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosServiceClassTable.EntityData.YListKeys = []string {}

    return &(docsQosServiceClassTable.EntityData)
}

// DOCSQOSMIB_DocsQosServiceClassTable_DocsQosServiceClassEntry
// A provisioned service class on a CMTS. 
// Each entry defines a template for certain 
// DOCSIS QOS Parameter Set values. When a CM 
// creates or modifies an Admitted QOS Parameter Set for a
// Service Flow, it may reference a Service Class
// Name instead of providing explicit QOS Parameter
// Set values. In this case, the CMTS populates
// the QOS Parameter Set with the applicable 
// corresponding values from the named Service Class.
// Subsequent changes to a Service Class row do *not* 
// affect the QOS Parameter Set values of any service flows
// already admitted.
// 
// A service class template applies to only
// a single direction, as indicated in the 
// docsQosServiceClassDirection object.
type DOCSQOSMIB_DocsQosServiceClassTable_DocsQosServiceClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Service Class Name. DOCSIS specifies that the
    // maximum size is 15 printable ASCII characters with  a terminating zero. The
    // terminating zero is not represented in this DisplayString syntax object.
    // The type is string with length: 1..15.
    DocsQosServiceClassName interface{}

    // This object is obsolete. The type is interface{} with range: 0..4294967295.
    DocsQosServiceClassParamSetIndex interface{}

    // Used to create or delete rows in this table. The type is RowStatus.
    DocsQosServiceClassStatus interface{}

    // Template for docsQosParamSetPriority. The type is interface{} with range:
    // 0..7.
    DocsQosServiceClassPriority interface{}

    // Template for docsQosParamSetMaxTrafficRate. The type is interface{} with
    // range: 0..4294967295.
    DocsQosServiceClassMaxTrafficRate interface{}

    // Template for docsQosParamSetMaxTrafficBurst. The type is interface{} with
    // range: 0..4294967295.
    DocsQosServiceClassMaxTrafficBurst interface{}

    // Template for docsQosParamSEtMinReservedRate. The type is interface{} with
    // range: 0..4294967295.
    DocsQosServiceClassMinReservedRate interface{}

    // Template for docsQosParamSetMinReservedPkt. The type is interface{} with
    // range: 0..65535.
    DocsQosServiceClassMinReservedPkt interface{}

    // Template for docsQosParamSetMaxConcatBurst. The type is interface{} with
    // range: 0..65535.
    DocsQosServiceClassMaxConcatBurst interface{}

    // Template for docsQosParamSetNomPollInterval. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    DocsQosServiceClassNomPollInterval interface{}

    // Template for docsQosParamSetTolPollJitter. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    DocsQosServiceClassTolPollJitter interface{}

    // Template for docsQosParamSetUnsolicitGrantSize. The type is interface{}
    // with range: 0..65535.
    DocsQosServiceClassUnsolicitGrantSize interface{}

    // Template for docsQosParamSetNomGrantInterval. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    DocsQosServiceClassNomGrantInterval interface{}

    // Template for docsQosParamSetTolGrantJitter. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    DocsQosServiceClassTolGrantJitter interface{}

    // Template for docsQosParamSetGrantsPerInterval. The type is interface{} with
    // range: 0..127.
    DocsQosServiceClassGrantsPerInterval interface{}

    // Template for docsQosParamSetClassMaxLatency. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    DocsQosServiceClassMaxLatency interface{}

    // Template for docsQosParamSetActiveTimeout. The type is interface{} with
    // range: 0..65535. Units are seconds.
    DocsQosServiceClassActiveTimeout interface{}

    // Template for docsQosParamSetAdmittedTimeout. The type is interface{} with
    // range: 0..65535. Units are seconds.
    DocsQosServiceClassAdmittedTimeout interface{}

    // Template for docsQosParamSetSchedulingType. The type is SchedulingType.
    DocsQosServiceClassSchedulingType interface{}

    // Template for docsQosParamSetRequestPolicyOct. The type is string with
    // length: 4.
    DocsQosServiceClassRequestPolicy interface{}

    // Template for docsQosParamSetTosAndMask. The type is string with length: 1.
    DocsQosServiceClassTosAndMask interface{}

    // Template for docsQosParamSetTosOrMask. The type is string with length: 1.
    DocsQosServiceClassTosOrMask interface{}

    // Specifies whether the service class template applies to upstream or
    // downstream service flows. The type is IfDirection.
    DocsQosServiceClassDirection interface{}
}

func (docsQosServiceClassEntry *DOCSQOSMIB_DocsQosServiceClassTable_DocsQosServiceClassEntry) GetEntityData() *types.CommonEntityData {
    docsQosServiceClassEntry.EntityData.YFilter = docsQosServiceClassEntry.YFilter
    docsQosServiceClassEntry.EntityData.YangName = "docsQosServiceClassEntry"
    docsQosServiceClassEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceClassEntry.EntityData.ParentYangName = "docsQosServiceClassTable"
    docsQosServiceClassEntry.EntityData.SegmentPath = "docsQosServiceClassEntry" + types.AddKeyToken(docsQosServiceClassEntry.DocsQosServiceClassName, "docsQosServiceClassName")
    docsQosServiceClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceClassEntry.EntityData.Children = types.NewOrderedMap()
    docsQosServiceClassEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassName", types.YLeaf{"DocsQosServiceClassName", docsQosServiceClassEntry.DocsQosServiceClassName})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassParamSetIndex", types.YLeaf{"DocsQosServiceClassParamSetIndex", docsQosServiceClassEntry.DocsQosServiceClassParamSetIndex})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassStatus", types.YLeaf{"DocsQosServiceClassStatus", docsQosServiceClassEntry.DocsQosServiceClassStatus})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassPriority", types.YLeaf{"DocsQosServiceClassPriority", docsQosServiceClassEntry.DocsQosServiceClassPriority})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMaxTrafficRate", types.YLeaf{"DocsQosServiceClassMaxTrafficRate", docsQosServiceClassEntry.DocsQosServiceClassMaxTrafficRate})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMaxTrafficBurst", types.YLeaf{"DocsQosServiceClassMaxTrafficBurst", docsQosServiceClassEntry.DocsQosServiceClassMaxTrafficBurst})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMinReservedRate", types.YLeaf{"DocsQosServiceClassMinReservedRate", docsQosServiceClassEntry.DocsQosServiceClassMinReservedRate})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMinReservedPkt", types.YLeaf{"DocsQosServiceClassMinReservedPkt", docsQosServiceClassEntry.DocsQosServiceClassMinReservedPkt})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMaxConcatBurst", types.YLeaf{"DocsQosServiceClassMaxConcatBurst", docsQosServiceClassEntry.DocsQosServiceClassMaxConcatBurst})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassNomPollInterval", types.YLeaf{"DocsQosServiceClassNomPollInterval", docsQosServiceClassEntry.DocsQosServiceClassNomPollInterval})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassTolPollJitter", types.YLeaf{"DocsQosServiceClassTolPollJitter", docsQosServiceClassEntry.DocsQosServiceClassTolPollJitter})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassUnsolicitGrantSize", types.YLeaf{"DocsQosServiceClassUnsolicitGrantSize", docsQosServiceClassEntry.DocsQosServiceClassUnsolicitGrantSize})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassNomGrantInterval", types.YLeaf{"DocsQosServiceClassNomGrantInterval", docsQosServiceClassEntry.DocsQosServiceClassNomGrantInterval})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassTolGrantJitter", types.YLeaf{"DocsQosServiceClassTolGrantJitter", docsQosServiceClassEntry.DocsQosServiceClassTolGrantJitter})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassGrantsPerInterval", types.YLeaf{"DocsQosServiceClassGrantsPerInterval", docsQosServiceClassEntry.DocsQosServiceClassGrantsPerInterval})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassMaxLatency", types.YLeaf{"DocsQosServiceClassMaxLatency", docsQosServiceClassEntry.DocsQosServiceClassMaxLatency})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassActiveTimeout", types.YLeaf{"DocsQosServiceClassActiveTimeout", docsQosServiceClassEntry.DocsQosServiceClassActiveTimeout})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassAdmittedTimeout", types.YLeaf{"DocsQosServiceClassAdmittedTimeout", docsQosServiceClassEntry.DocsQosServiceClassAdmittedTimeout})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassSchedulingType", types.YLeaf{"DocsQosServiceClassSchedulingType", docsQosServiceClassEntry.DocsQosServiceClassSchedulingType})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassRequestPolicy", types.YLeaf{"DocsQosServiceClassRequestPolicy", docsQosServiceClassEntry.DocsQosServiceClassRequestPolicy})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassTosAndMask", types.YLeaf{"DocsQosServiceClassTosAndMask", docsQosServiceClassEntry.DocsQosServiceClassTosAndMask})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassTosOrMask", types.YLeaf{"DocsQosServiceClassTosOrMask", docsQosServiceClassEntry.DocsQosServiceClassTosOrMask})
    docsQosServiceClassEntry.EntityData.Leafs.Append("docsQosServiceClassDirection", types.YLeaf{"DocsQosServiceClassDirection", docsQosServiceClassEntry.DocsQosServiceClassDirection})

    docsQosServiceClassEntry.EntityData.YListKeys = []string {"DocsQosServiceClassName"}

    return &(docsQosServiceClassEntry.EntityData)
}

// DOCSQOSMIB_DocsQosServiceClassPolicyTable
// This table describes the set of Docsis-QOS 
// Service Class Policies.  
// 
// This table is an adjunct to the
// docsDevFilterPolicy table.  Entries in 
// docsDevFilterPolicy table can  point to 
// specific rows in this table.
// 
// This table permits mapping a packet to a service
// class name of an active service flow so long as 
// a classifier does not exist at a higher
// priority.
type DOCSQOSMIB_DocsQosServiceClassPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A service class name policy entry. The type is slice of
    // DOCSQOSMIB_DocsQosServiceClassPolicyTable_DocsQosServiceClassPolicyEntry.
    DocsQosServiceClassPolicyEntry []*DOCSQOSMIB_DocsQosServiceClassPolicyTable_DocsQosServiceClassPolicyEntry
}

func (docsQosServiceClassPolicyTable *DOCSQOSMIB_DocsQosServiceClassPolicyTable) GetEntityData() *types.CommonEntityData {
    docsQosServiceClassPolicyTable.EntityData.YFilter = docsQosServiceClassPolicyTable.YFilter
    docsQosServiceClassPolicyTable.EntityData.YangName = "docsQosServiceClassPolicyTable"
    docsQosServiceClassPolicyTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceClassPolicyTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosServiceClassPolicyTable.EntityData.SegmentPath = "docsQosServiceClassPolicyTable"
    docsQosServiceClassPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceClassPolicyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceClassPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceClassPolicyTable.EntityData.Children = types.NewOrderedMap()
    docsQosServiceClassPolicyTable.EntityData.Children.Append("docsQosServiceClassPolicyEntry", types.YChild{"DocsQosServiceClassPolicyEntry", nil})
    for i := range docsQosServiceClassPolicyTable.DocsQosServiceClassPolicyEntry {
        docsQosServiceClassPolicyTable.EntityData.Children.Append(types.GetSegmentPath(docsQosServiceClassPolicyTable.DocsQosServiceClassPolicyEntry[i]), types.YChild{"DocsQosServiceClassPolicyEntry", docsQosServiceClassPolicyTable.DocsQosServiceClassPolicyEntry[i]})
    }
    docsQosServiceClassPolicyTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosServiceClassPolicyTable.EntityData.YListKeys = []string {}

    return &(docsQosServiceClassPolicyTable.EntityData)
}

// DOCSQOSMIB_DocsQosServiceClassPolicyTable_DocsQosServiceClassPolicyEntry
// A service class name policy entry.
type DOCSQOSMIB_DocsQosServiceClassPolicyTable_DocsQosServiceClassPolicyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index value to uniquely identify an entry in this
    // table. The type is interface{} with range: 1..2147483647.
    DocsQosServiceClassPolicyIndex interface{}

    // Service Class Name to identify the name of the  service class flow to which
    // the packet should be directed. The type is string.
    DocsQosServiceClassPolicyName interface{}

    // Service Class Policy rule priority for the entry. The type is interface{}
    // with range: 0..255.
    DocsQosServiceClassPolicyRulePriority interface{}

    // Used to create or delete rows in this table. This object should not be
    // deleted if it is reference by an entry in docsDevFilterPolicy. The
    // reference should be deleted first. The type is RowStatus.
    DocsQosServiceClassPolicyStatus interface{}
}

func (docsQosServiceClassPolicyEntry *DOCSQOSMIB_DocsQosServiceClassPolicyTable_DocsQosServiceClassPolicyEntry) GetEntityData() *types.CommonEntityData {
    docsQosServiceClassPolicyEntry.EntityData.YFilter = docsQosServiceClassPolicyEntry.YFilter
    docsQosServiceClassPolicyEntry.EntityData.YangName = "docsQosServiceClassPolicyEntry"
    docsQosServiceClassPolicyEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosServiceClassPolicyEntry.EntityData.ParentYangName = "docsQosServiceClassPolicyTable"
    docsQosServiceClassPolicyEntry.EntityData.SegmentPath = "docsQosServiceClassPolicyEntry" + types.AddKeyToken(docsQosServiceClassPolicyEntry.DocsQosServiceClassPolicyIndex, "docsQosServiceClassPolicyIndex")
    docsQosServiceClassPolicyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosServiceClassPolicyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosServiceClassPolicyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosServiceClassPolicyEntry.EntityData.Children = types.NewOrderedMap()
    docsQosServiceClassPolicyEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosServiceClassPolicyEntry.EntityData.Leafs.Append("docsQosServiceClassPolicyIndex", types.YLeaf{"DocsQosServiceClassPolicyIndex", docsQosServiceClassPolicyEntry.DocsQosServiceClassPolicyIndex})
    docsQosServiceClassPolicyEntry.EntityData.Leafs.Append("docsQosServiceClassPolicyName", types.YLeaf{"DocsQosServiceClassPolicyName", docsQosServiceClassPolicyEntry.DocsQosServiceClassPolicyName})
    docsQosServiceClassPolicyEntry.EntityData.Leafs.Append("docsQosServiceClassPolicyRulePriority", types.YLeaf{"DocsQosServiceClassPolicyRulePriority", docsQosServiceClassPolicyEntry.DocsQosServiceClassPolicyRulePriority})
    docsQosServiceClassPolicyEntry.EntityData.Leafs.Append("docsQosServiceClassPolicyStatus", types.YLeaf{"DocsQosServiceClassPolicyStatus", docsQosServiceClassPolicyEntry.DocsQosServiceClassPolicyStatus})

    docsQosServiceClassPolicyEntry.EntityData.YListKeys = []string {"DocsQosServiceClassPolicyIndex"}

    return &(docsQosServiceClassPolicyEntry.EntityData)
}

// DOCSQOSMIB_DocsQosPHSTable
// This table describes set of payload header
// suppression entries.
type DOCSQOSMIB_DocsQosPHSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A payload header suppression entry.  The ifIndex is an ifType of
    // docsCableMaclayer(127). The index docsQosServiceFlowId selects one service
    // flow from the cable MAC layer interface. The docsQosPktClassId index
    // matches an index of the docsQosPktClassTable. The type is slice of
    // DOCSQOSMIB_DocsQosPHSTable_DocsQosPHSEntry.
    DocsQosPHSEntry []*DOCSQOSMIB_DocsQosPHSTable_DocsQosPHSEntry
}

func (docsQosPHSTable *DOCSQOSMIB_DocsQosPHSTable) GetEntityData() *types.CommonEntityData {
    docsQosPHSTable.EntityData.YFilter = docsQosPHSTable.YFilter
    docsQosPHSTable.EntityData.YangName = "docsQosPHSTable"
    docsQosPHSTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosPHSTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosPHSTable.EntityData.SegmentPath = "docsQosPHSTable"
    docsQosPHSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosPHSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosPHSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosPHSTable.EntityData.Children = types.NewOrderedMap()
    docsQosPHSTable.EntityData.Children.Append("docsQosPHSEntry", types.YChild{"DocsQosPHSEntry", nil})
    for i := range docsQosPHSTable.DocsQosPHSEntry {
        docsQosPHSTable.EntityData.Children.Append(types.GetSegmentPath(docsQosPHSTable.DocsQosPHSEntry[i]), types.YChild{"DocsQosPHSEntry", docsQosPHSTable.DocsQosPHSEntry[i]})
    }
    docsQosPHSTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosPHSTable.EntityData.YListKeys = []string {}

    return &(docsQosPHSTable.EntityData)
}

// DOCSQOSMIB_DocsQosPHSTable_DocsQosPHSEntry
// A payload header suppression entry. 
// The ifIndex is an ifType of docsCableMaclayer(127).
// The index docsQosServiceFlowId selects one
// service flow from the cable MAC layer interface.
// The docsQosPktClassId index matches an
// index of the docsQosPktClassTable.
type DOCSQOSMIB_DocsQosPHSTable_DocsQosPHSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassId
    DocsQosPktClassId interface{}

    // Payload header suppression field defines the  bytes of the header which
    // must be  suppressed/restored by the sending/receiving  device.  The number
    // of octets in this object should be the same as the value of docsQosPHSSize.
    // The type is string with length: 0..255.
    DocsQosPHSField interface{}

    // Payload header suppression mask defines the  bit mask which used in
    // combination with the docsQosPHSField defines which bytes in header must be
    // suppressed/restored by the sending or receiving device.  Each bit of this
    // bit mask corresponds to a byte in the docsQosPHSField, with the least 
    // significant  bit corresponding to first byte of the docsQosPHSField.  Each
    // bit of the bit mask specifies whether of not the corresponding byte should
    // be suppressed in the packet. A bit value of '1' indicates that the byte
    // should be suppressed by the sending  device and restored by the receiving
    // device.  A bit value of '0' indicates that  the byte should not be
    // suppressed by the sending device or restored by the receiving device.  If
    // the bit mask does not contain a bit for each byte in the docsQosPHSField
    // then the bit mask is extended with bit values of '1' to be the necessary
    // length. The type is string with length: 0..32.
    DocsQosPHSMask interface{}

    // Payload header suppression size specifies the  number of bytes in the
    // header to be suppressed and restored.  The value of this object must match
    // the number of bytes in the docsQosPHSField. The type is interface{} with
    // range: 0..255.
    DocsQosPHSSize interface{}

    // Payload header suppression verification value of 'true' the sender must
    // verify docsQosPHSField  is the same as what is contained in the packet to
    // be suppressed. The type is bool.
    DocsQosPHSVerify interface{}

    // This object is obsolete. The type is interface{} with range: 1..65535.
    DocsQosPHSClassifierIndex interface{}

    // Payload header suppression index uniquely  references the PHS rule for a
    // given service flow. The type is interface{} with range: 1..255.
    DocsQosPHSIndex interface{}
}

func (docsQosPHSEntry *DOCSQOSMIB_DocsQosPHSTable_DocsQosPHSEntry) GetEntityData() *types.CommonEntityData {
    docsQosPHSEntry.EntityData.YFilter = docsQosPHSEntry.YFilter
    docsQosPHSEntry.EntityData.YangName = "docsQosPHSEntry"
    docsQosPHSEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosPHSEntry.EntityData.ParentYangName = "docsQosPHSTable"
    docsQosPHSEntry.EntityData.SegmentPath = "docsQosPHSEntry" + types.AddKeyToken(docsQosPHSEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsQosPHSEntry.DocsQosServiceFlowId, "docsQosServiceFlowId") + types.AddKeyToken(docsQosPHSEntry.DocsQosPktClassId, "docsQosPktClassId")
    docsQosPHSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosPHSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosPHSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosPHSEntry.EntityData.Children = types.NewOrderedMap()
    docsQosPHSEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosPHSEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsQosPHSEntry.IfIndex})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsQosPHSEntry.DocsQosServiceFlowId})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPktClassId", types.YLeaf{"DocsQosPktClassId", docsQosPHSEntry.DocsQosPktClassId})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSField", types.YLeaf{"DocsQosPHSField", docsQosPHSEntry.DocsQosPHSField})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSMask", types.YLeaf{"DocsQosPHSMask", docsQosPHSEntry.DocsQosPHSMask})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSSize", types.YLeaf{"DocsQosPHSSize", docsQosPHSEntry.DocsQosPHSSize})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSVerify", types.YLeaf{"DocsQosPHSVerify", docsQosPHSEntry.DocsQosPHSVerify})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSClassifierIndex", types.YLeaf{"DocsQosPHSClassifierIndex", docsQosPHSEntry.DocsQosPHSClassifierIndex})
    docsQosPHSEntry.EntityData.Leafs.Append("docsQosPHSIndex", types.YLeaf{"DocsQosPHSIndex", docsQosPHSEntry.DocsQosPHSIndex})

    docsQosPHSEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId", "DocsQosPktClassId"}

    return &(docsQosPHSEntry.EntityData)
}

// DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable
// This table provide for referencing the service flows 
// associated with a particular cable modem. This allows 
// for indexing into other docsQos tables that are 
// indexed by docsQosServiceFlowId and ifIndex.
type DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created by CMTS for each service flow  connected to this CMTS.
    // The type is slice of
    // DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable_DocsQosCmtsMacToSrvFlowEntry.
    DocsQosCmtsMacToSrvFlowEntry []*DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable_DocsQosCmtsMacToSrvFlowEntry
}

func (docsQosCmtsMacToSrvFlowTable *DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable) GetEntityData() *types.CommonEntityData {
    docsQosCmtsMacToSrvFlowTable.EntityData.YFilter = docsQosCmtsMacToSrvFlowTable.YFilter
    docsQosCmtsMacToSrvFlowTable.EntityData.YangName = "docsQosCmtsMacToSrvFlowTable"
    docsQosCmtsMacToSrvFlowTable.EntityData.BundleName = "cisco_ios_xe"
    docsQosCmtsMacToSrvFlowTable.EntityData.ParentYangName = "DOCS-QOS-MIB"
    docsQosCmtsMacToSrvFlowTable.EntityData.SegmentPath = "docsQosCmtsMacToSrvFlowTable"
    docsQosCmtsMacToSrvFlowTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosCmtsMacToSrvFlowTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosCmtsMacToSrvFlowTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosCmtsMacToSrvFlowTable.EntityData.Children = types.NewOrderedMap()
    docsQosCmtsMacToSrvFlowTable.EntityData.Children.Append("docsQosCmtsMacToSrvFlowEntry", types.YChild{"DocsQosCmtsMacToSrvFlowEntry", nil})
    for i := range docsQosCmtsMacToSrvFlowTable.DocsQosCmtsMacToSrvFlowEntry {
        docsQosCmtsMacToSrvFlowTable.EntityData.Children.Append(types.GetSegmentPath(docsQosCmtsMacToSrvFlowTable.DocsQosCmtsMacToSrvFlowEntry[i]), types.YChild{"DocsQosCmtsMacToSrvFlowEntry", docsQosCmtsMacToSrvFlowTable.DocsQosCmtsMacToSrvFlowEntry[i]})
    }
    docsQosCmtsMacToSrvFlowTable.EntityData.Leafs = types.NewOrderedMap()

    docsQosCmtsMacToSrvFlowTable.EntityData.YListKeys = []string {}

    return &(docsQosCmtsMacToSrvFlowTable.EntityData)
}

// DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable_DocsQosCmtsMacToSrvFlowEntry
// An entry is created by CMTS for each service flow 
// connected to this CMTS.
type DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable_DocsQosCmtsMacToSrvFlowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MAC address for the referenced CM. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsQosCmtsCmMac interface{}

    // This attribute is a key. An index assigned to a service flow by CMTS. The
    // type is interface{} with range: 1..4294967295.
    DocsQosCmtsServiceFlowId interface{}

    // The ifIndex of ifType docsCableMacLayter(127)  on the CMTS that is
    // connected to the Cable Modem. The type is interface{} with range:
    // 1..2147483647.
    DocsQosCmtsIfIndex interface{}
}

func (docsQosCmtsMacToSrvFlowEntry *DOCSQOSMIB_DocsQosCmtsMacToSrvFlowTable_DocsQosCmtsMacToSrvFlowEntry) GetEntityData() *types.CommonEntityData {
    docsQosCmtsMacToSrvFlowEntry.EntityData.YFilter = docsQosCmtsMacToSrvFlowEntry.YFilter
    docsQosCmtsMacToSrvFlowEntry.EntityData.YangName = "docsQosCmtsMacToSrvFlowEntry"
    docsQosCmtsMacToSrvFlowEntry.EntityData.BundleName = "cisco_ios_xe"
    docsQosCmtsMacToSrvFlowEntry.EntityData.ParentYangName = "docsQosCmtsMacToSrvFlowTable"
    docsQosCmtsMacToSrvFlowEntry.EntityData.SegmentPath = "docsQosCmtsMacToSrvFlowEntry" + types.AddKeyToken(docsQosCmtsMacToSrvFlowEntry.DocsQosCmtsCmMac, "docsQosCmtsCmMac") + types.AddKeyToken(docsQosCmtsMacToSrvFlowEntry.DocsQosCmtsServiceFlowId, "docsQosCmtsServiceFlowId")
    docsQosCmtsMacToSrvFlowEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsQosCmtsMacToSrvFlowEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsQosCmtsMacToSrvFlowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsQosCmtsMacToSrvFlowEntry.EntityData.Children = types.NewOrderedMap()
    docsQosCmtsMacToSrvFlowEntry.EntityData.Leafs = types.NewOrderedMap()
    docsQosCmtsMacToSrvFlowEntry.EntityData.Leafs.Append("docsQosCmtsCmMac", types.YLeaf{"DocsQosCmtsCmMac", docsQosCmtsMacToSrvFlowEntry.DocsQosCmtsCmMac})
    docsQosCmtsMacToSrvFlowEntry.EntityData.Leafs.Append("docsQosCmtsServiceFlowId", types.YLeaf{"DocsQosCmtsServiceFlowId", docsQosCmtsMacToSrvFlowEntry.DocsQosCmtsServiceFlowId})
    docsQosCmtsMacToSrvFlowEntry.EntityData.Leafs.Append("docsQosCmtsIfIndex", types.YLeaf{"DocsQosCmtsIfIndex", docsQosCmtsMacToSrvFlowEntry.DocsQosCmtsIfIndex})

    docsQosCmtsMacToSrvFlowEntry.EntityData.YListKeys = []string {"DocsQosCmtsCmMac", "DocsQosCmtsServiceFlowId"}

    return &(docsQosCmtsMacToSrvFlowEntry.EntityData)
}

