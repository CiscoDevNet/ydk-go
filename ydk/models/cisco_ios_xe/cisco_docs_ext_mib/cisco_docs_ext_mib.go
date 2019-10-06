// This is the MIB module for the Cisco specific extension 
// objects of Data Over Cable Service, Radio Frequency 
// interface.  There is a standard MIB for Data-Over-Cable 
// Service Interface Specifications (DOCSIS) and in Cisco,
// it is called DOCS-IF-MIB. Besides the objects in 
// DOCS-IF-MIB, this MIB module contains the extension 
// objects to manage the Cable Modem Termination Systems 
// (CMTS).
// 
// This MIB module includes objects for the scheduler 
// that supports Quality of Service (QoS) of MCNS/DOCSIS 
// compliant Radio Frequency (RF) interfaces in Cable Modem 
// Termination Systems (CMTS). And the purpose is to let 
// users configure attributes of the schedulers in 
// order to ensure the Quality of Service and fairness for 
// modem requests according to users' business needs. 
// Also this MIB shows various states of the schedulers 
// for users to monitor of the schedulers' current status. 
// 
// This MIB module also includes connection status objects
// for cable modems and Customer Premise Equipment (CPE) 
// and the purpose is to let users easily get the connection 
// status and manage access group information about cable 
// modems and CPE.
// 
// This MIB module also includes objects for upstream 
// configuration for automated spectrum management in 
// order to mitigate upstream impairment.
// 
// This MIB module also includes objects to keep count of
// the total # of modems, # of registered and # of active
// modems on the mac interface as well as each 
// upstream. 
// 
// Glossary:
// BE       Best Effort
// 
// CPE      Customer Premise Equipment
// 
// CM       Cable Modem
// 
// CMTS     Cable Modem Termination Systems
// 
// DMIC     Dynamic Message Integrity Check
// 
// DOCSIS   Data Over Cable Service Interface Specifications
// 
// IE       Information Element
// 
// NIC      Network Interface Card
// 
// QoS      Quality of Service
// 
// RF       Radio Frequency
// 
// RTPS     Real-Time Polling Service
// 
// SFID     Service Flow ID
// 
// SID      Service Id
// 
// TOD      Time of the Day
// 
// UGS      Unsolicited Grant Service
// 
// UGS-AD   Unsolicited Grant Service with Activity Detection
package cisco_docs_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_docs_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-DOCS-EXT-MIB CISCO-DOCS-EXT-MIB}", reflect.TypeOf(CISCODOCSEXTMIB{}))
    ydk.RegisterEntity("CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB", reflect.TypeOf(CISCODOCSEXTMIB{}))
}

// CISCODOCSEXTMIB
type CISCODOCSEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CdxCmtsCmCpeObjects CISCODOCSEXTMIB_CdxCmtsCmCpeObjects

    
    CdxWBResilObjects CISCODOCSEXTMIB_CdxWBResilObjects

    
    CdxCmtsDocsisLBObjects CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects

    // For each upstream interface, this table maintains a number  of objects
    // related to Quality of Service scheduler which uses these attributes to
    // control cable modem registration. .
    CdxQosCtrlUpTable CISCODOCSEXTMIB_CdxQosCtrlUpTable

    // This table describes the attributes of rate limiting for  schedulers in
    // cable upstream and downstream interfaces that  support Quality of Service. 
    // The rate limiting process is  to ensure the Quality of Service and
    // fairness. .
    CdxQosIfRateLimitTable CISCODOCSEXTMIB_CdxQosIfRateLimitTable

    // The list contains the additional attributes of a single Service ID that
    // provided by docsIfCmtsServiceEntry. .
    CdxCmtsServiceExtTable CISCODOCSEXTMIB_CdxCmtsServiceExtTable

    // The table contains the attributes of a particular  Information Element type
    // for each instance of the MAC  scheduler. It is indexed by upstream ifIndex.
    // An Entry exists for each ifEntry with ifType of docsCableUpstream(129)
    // Since each upstream has an instance of a MAC Scheduler so  this table has
    // the per MAC scheduler slot information on a per Information Element basis.
    // .
    CdxUpInfoElemStatsTable CISCODOCSEXTMIB_CdxUpInfoElemStatsTable

    // This table describes the attributes of queues   in cable interfaces
    // schedulers that support  Quality of Service.
    CdxBWQueueTable CISCODOCSEXTMIB_CdxBWQueueTable

    // This table contains information about cable modems (CM) or  Customer
    // Premises Equipments (CPE). .
    CdxCmCpeTable CISCODOCSEXTMIB_CdxCmCpeTable

    // The list contains the additional CM status information. .
    CdxCmtsCmStatusExtTable CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable

    // This table contains the additions attributes of a CMTS MAC interface that
    // provided by docsIfCmtsMacTable. .
    CdxCmtsMacExtTable CISCODOCSEXTMIB_CdxCmtsMacExtTable

    // A table of CMTS operation entries to instruct cable modems to move to a new
    // downstream and/or upstream channel.   An entry in this table is an
    // operation that has been  initiated from CMTS to generates downstream
    // frequency and/or  upstream channel override fields in the RNG-RSP message
    // sent  to a cable modem.  A RNG-RSP message is sent to a cable  modem during
    // initial maintenance opportunity.   This operation causes the uBR to place
    // an entry for the cable  modem specified into the override request queue. 
    // The link is  then broken by deleting the modem from its polling list.  When
    // the modem attempts initial ranging, the override request  causes downstream
    // frequency and upstream channel override  fields to be inserted into the
    // RNG-RSP message.  .
    CdxCmtsCmChOverTable CISCODOCSEXTMIB_CdxCmtsCmChOverTable

    // This table contains attributes or configurable parameters  for cable modems
    // from a CMTS. .
    CdxCmtsCmTable CISCODOCSEXTMIB_CdxCmtsCmTable

    // This table contains the list of modems which failed the CMTS Dynamic
    // Message Integrity Check (DMIC). The modems are  either Marked: The modems
    // failed the DMIC check but were still          allowed to come online.
    // Locked: The modems failed the DMIC check and hence were          allowed to
    // come online with a restrictive QoS          profile as defined in 
    // cdxCmtsCmDMICLockQos. Rejected: The modems failed the DMIC check and hence 
    // were not allowed to come online. Another objective of the objects in this
    // table is to clear the lock on the modems.
    CdxCmtsCmStatusDMICTable CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable

    // This table contains information about CPE connects behind cable modem. It
    // will return IP address and IP address type of each CPE connect to a CM.  It
    // is not intended to walk the whole table. An application would need to query
    // this table based on the specific indices. Otherwise, it will impact the
    // CMTS performance due to the  huge size of this table.  The agent
    // creates/destroys/modifies an entry whenever there is a CPE connect to a
    // cable modem or disconnect from a cable modem.
    CdxCmToCpeTable CISCODOCSEXTMIB_CdxCmToCpeTable

    // This table contains information about cable modems with CPE connects to. 
    // It is not intended to walk the whole table. An application would need to
    // query this table base on the specific index. Otherwise, it will impact the
    // CMTS performance due to the huge size of this table.  The agent
    // creates/destroys/modifies an entry whenever there is update for the cable
    // modem that CPE connects to.
    CdxCpeToCmTable CISCODOCSEXTMIB_CdxCpeToCmTable

    // The table contains a list CPE's IP Prefix management information.
    CdxCpeIpPrefixTable CISCODOCSEXTMIB_CdxCpeIpPrefixTable

    // This table contains upstream channel attributes for   automated Spectrum
    // management, in addition to the ones provided by docsIfUpstreamChannelEntry.
    // It also contains upstream channel attributes to count  the number of
    // active,registered and total number of cable  modems on this upstream. .
    CdxIfUpstreamChannelExtTable CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable

    // This table contains information about partial service cable modems (CM),
    // including both downstream and upstream partial service modems.
    CdxWBResilCmTable CISCODOCSEXTMIB_CdxWBResilCmTable

    // This table contains information of the mapping of the physical RF channels
    // to the primary RF channels.
    CdxRFtoPrimaryChannelMappingTable CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable

    // This table contains information of the mapping of the primary RF channels
    // to the physical RF channels.
    CdxPrimaryChanneltoRFMappingTable CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable

    // This table contains CM management information of Transmit Channel Set(TCS)
    // in the system.
    CdxCmtsMtcCmTable CISCODOCSEXTMIB_CdxCmtsMtcCmTable

    // This table contains the Upstream Channel Bonding Service Flow management
    // information.
    CdxCmtsUscbSflowTable CISCODOCSEXTMIB_CdxCmtsUscbSflowTable

    // The cdxRPDGS7KTable contains the attributes of GS7K.  An Entry exists for
    // each instance.  It is indexed by GS7K's MacAddress.
    CdxRPDGS7KTable CISCODOCSEXTMIB_CdxRPDGS7KTable

    // A list of cable helper entries on Bundle/Sub-Bundle interface.
    CdxBundleIpHelperTable CISCODOCSEXTMIB_CdxBundleIpHelperTable

    // Ipv6 dhcp relay configurations on Bundle/Sub-Bundle interface.
    CdxBundleIPv6DHCPRelayTable CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable

    // A list of IPv6 DHCP relay destination entries on Cable Bundle/Sub-Bundle
    // interfaces.
    CdxBundleIPv6DHCPRelayDestTable CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable
}

func (cISCODOCSEXTMIB *CISCODOCSEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCODOCSEXTMIB.EntityData.YFilter = cISCODOCSEXTMIB.YFilter
    cISCODOCSEXTMIB.EntityData.YangName = "CISCO-DOCS-EXT-MIB"
    cISCODOCSEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODOCSEXTMIB.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cISCODOCSEXTMIB.EntityData.SegmentPath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB"
    cISCODOCSEXTMIB.EntityData.AbsolutePath = cISCODOCSEXTMIB.EntityData.SegmentPath
    cISCODOCSEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODOCSEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODOCSEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODOCSEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsCmCpeObjects", types.YChild{"CdxCmtsCmCpeObjects", &cISCODOCSEXTMIB.CdxCmtsCmCpeObjects})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxWBResilObjects", types.YChild{"CdxWBResilObjects", &cISCODOCSEXTMIB.CdxWBResilObjects})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsDocsisLBObjects", types.YChild{"CdxCmtsDocsisLBObjects", &cISCODOCSEXTMIB.CdxCmtsDocsisLBObjects})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxQosCtrlUpTable", types.YChild{"CdxQosCtrlUpTable", &cISCODOCSEXTMIB.CdxQosCtrlUpTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxQosIfRateLimitTable", types.YChild{"CdxQosIfRateLimitTable", &cISCODOCSEXTMIB.CdxQosIfRateLimitTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsServiceExtTable", types.YChild{"CdxCmtsServiceExtTable", &cISCODOCSEXTMIB.CdxCmtsServiceExtTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxUpInfoElemStatsTable", types.YChild{"CdxUpInfoElemStatsTable", &cISCODOCSEXTMIB.CdxUpInfoElemStatsTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxBWQueueTable", types.YChild{"CdxBWQueueTable", &cISCODOCSEXTMIB.CdxBWQueueTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmCpeTable", types.YChild{"CdxCmCpeTable", &cISCODOCSEXTMIB.CdxCmCpeTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsCmStatusExtTable", types.YChild{"CdxCmtsCmStatusExtTable", &cISCODOCSEXTMIB.CdxCmtsCmStatusExtTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsMacExtTable", types.YChild{"CdxCmtsMacExtTable", &cISCODOCSEXTMIB.CdxCmtsMacExtTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsCmChOverTable", types.YChild{"CdxCmtsCmChOverTable", &cISCODOCSEXTMIB.CdxCmtsCmChOverTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsCmTable", types.YChild{"CdxCmtsCmTable", &cISCODOCSEXTMIB.CdxCmtsCmTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsCmStatusDMICTable", types.YChild{"CdxCmtsCmStatusDMICTable", &cISCODOCSEXTMIB.CdxCmtsCmStatusDMICTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmToCpeTable", types.YChild{"CdxCmToCpeTable", &cISCODOCSEXTMIB.CdxCmToCpeTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCpeToCmTable", types.YChild{"CdxCpeToCmTable", &cISCODOCSEXTMIB.CdxCpeToCmTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCpeIpPrefixTable", types.YChild{"CdxCpeIpPrefixTable", &cISCODOCSEXTMIB.CdxCpeIpPrefixTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxIfUpstreamChannelExtTable", types.YChild{"CdxIfUpstreamChannelExtTable", &cISCODOCSEXTMIB.CdxIfUpstreamChannelExtTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxWBResilCmTable", types.YChild{"CdxWBResilCmTable", &cISCODOCSEXTMIB.CdxWBResilCmTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxRFtoPrimaryChannelMappingTable", types.YChild{"CdxRFtoPrimaryChannelMappingTable", &cISCODOCSEXTMIB.CdxRFtoPrimaryChannelMappingTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxPrimaryChanneltoRFMappingTable", types.YChild{"CdxPrimaryChanneltoRFMappingTable", &cISCODOCSEXTMIB.CdxPrimaryChanneltoRFMappingTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsMtcCmTable", types.YChild{"CdxCmtsMtcCmTable", &cISCODOCSEXTMIB.CdxCmtsMtcCmTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxCmtsUscbSflowTable", types.YChild{"CdxCmtsUscbSflowTable", &cISCODOCSEXTMIB.CdxCmtsUscbSflowTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxRPDGS7KTable", types.YChild{"CdxRPDGS7KTable", &cISCODOCSEXTMIB.CdxRPDGS7KTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxBundleIpHelperTable", types.YChild{"CdxBundleIpHelperTable", &cISCODOCSEXTMIB.CdxBundleIpHelperTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxBundleIPv6DHCPRelayTable", types.YChild{"CdxBundleIPv6DHCPRelayTable", &cISCODOCSEXTMIB.CdxBundleIPv6DHCPRelayTable})
    cISCODOCSEXTMIB.EntityData.Children.Append("cdxBundleIPv6DHCPRelayDestTable", types.YChild{"CdxBundleIPv6DHCPRelayDestTable", &cISCODOCSEXTMIB.CdxBundleIPv6DHCPRelayDestTable})
    cISCODOCSEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCODOCSEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCODOCSEXTMIB.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmCpeObjects
type CISCODOCSEXTMIB_CdxCmtsCmCpeObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time period to expire a CMTS channel override operation.  Within the
    // time period, if the CMTS cannot send out a  RNG-RSP message with channel
    // override fields to a cable  modem specified in the operation, the CMTS will
    // abort  the operation. The possible reason is that the cable  modem does not
    // repeat the initial ranging.   The change to this object will not affect the
    // already active  operations in this cdxCmtsCmChOverTable.      Once the
    // operation completes, the management station should retrieve the values of
    // the cdxCmtsCmChOverState  object of interest, and should then delete the
    // entry from cdxCmtsCmChOverTable.  In order to prevent old  entries from
    // clogging the table, entries will be aged out,  but an entry will never be
    // deleted within 15 minutes of  completing. . The type is interface{} with
    // range: 1..86400. Units are minutes.
    CdxCmtsCmChOverTimeExpiration interface{}
}

func (cdxCmtsCmCpeObjects *CISCODOCSEXTMIB_CdxCmtsCmCpeObjects) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmCpeObjects.EntityData.YFilter = cdxCmtsCmCpeObjects.YFilter
    cdxCmtsCmCpeObjects.EntityData.YangName = "cdxCmtsCmCpeObjects"
    cdxCmtsCmCpeObjects.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmCpeObjects.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsCmCpeObjects.EntityData.SegmentPath = "cdxCmtsCmCpeObjects"
    cdxCmtsCmCpeObjects.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsCmCpeObjects.EntityData.SegmentPath
    cdxCmtsCmCpeObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmCpeObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmCpeObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmCpeObjects.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmCpeObjects.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsCmCpeObjects.EntityData.Leafs.Append("cdxCmtsCmChOverTimeExpiration", types.YLeaf{"CdxCmtsCmChOverTimeExpiration", cdxCmtsCmCpeObjects.CdxCmtsCmChOverTimeExpiration})

    cdxCmtsCmCpeObjects.EntityData.YListKeys = []string {}

    return &(cdxCmtsCmCpeObjects.EntityData)
}

// CISCODOCSEXTMIB_CdxWBResilObjects
type CISCODOCSEXTMIB_CdxWBResilObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the amount of time an RF channel must remain in its
    // new state, either UP or DOWN, before the transition is considered valid. 
    // This value applies to all non-primary RF channels in the CMTS. The type is
    // interface{} with range: 1..65535. Units are Second.
    CdxWBResilRFChangeDampenTime interface{}

    // This object specifies the percentage of cable modems (CMs) that must report
    // that a particular Non Primary RF channel is DOWN, before that channel is
    // removed from any/all bonding groups with that Non Primary RF channel
    // configured. The value of 0 will prevent from any bonding group
    // modifications. In order to dampen state's changes for an RF channel, the
    // trigger for  a channel being restored is one half of this object's value. .
    // The type is interface{} with range: 0..100. Units are Percentage.
    CdxWBResilRFChangeTriggerPercentage interface{}

    // This object specifies the count of cable modems (CMs) that must report that
    // a particular Non Primary RF channel is DOWN, before that channel is removed
    // from any/all bonding groups with that Non Primary RF channel configured.
    // The value of 0 will prevent from any bonding group modifications. In order
    // to dampen state's changes for an RF channel, the trigger for  a channel
    // being restored is one half of this object's value. The type is interface{}
    // with range: 0..65535.
    CdxWBResilRFChangeTriggerCount interface{}

    // This object specifies whether the secondary service flows are allowed to be
    // moved and created on the narrowband interface. The type is bool.
    CdxWBResilRFChangeTriggerMoveSecondary interface{}

    // An indication of whether the cdxWBResilRFDown, cdxWBResilRFUp, 
    // cdxWBResilCMPartialServiceNotif, cdxWBResilCMFullServiceNotif, 
    // cdxWBResilEvent, cdxWBResilUsFullServiceNotif and
    // cdxWBResilUsPartialServiceNotif are enabled. The type is map[string]bool.
    CdxWBResilNotificationEnable interface{}

    // This object specifies the interval that cdxWBResilEvent traps  could be
    // sent per cable modem. It is to avoid too many cdxWBResilEvent traps sent
    // for a cable modem during a short period of time. The default value is 1
    // (second). If the value is 0, the trap  cdxWBResilEvent will be sent for
    // every wideband resiliency event. If the value is set to any value greater
    // than 0, for the wideband  resiliency events which occurred in the same
    // specific period of time,  the CMTS will send only one trap. The type is
    // interface{} with range: 0..86400. Units are Second.
    CdxWBResilNotificationsInterval interface{}
}

func (cdxWBResilObjects *CISCODOCSEXTMIB_CdxWBResilObjects) GetEntityData() *types.CommonEntityData {
    cdxWBResilObjects.EntityData.YFilter = cdxWBResilObjects.YFilter
    cdxWBResilObjects.EntityData.YangName = "cdxWBResilObjects"
    cdxWBResilObjects.EntityData.BundleName = "cisco_ios_xe"
    cdxWBResilObjects.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxWBResilObjects.EntityData.SegmentPath = "cdxWBResilObjects"
    cdxWBResilObjects.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxWBResilObjects.EntityData.SegmentPath
    cdxWBResilObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxWBResilObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxWBResilObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxWBResilObjects.EntityData.Children = types.NewOrderedMap()
    cdxWBResilObjects.EntityData.Leafs = types.NewOrderedMap()
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilRFChangeDampenTime", types.YLeaf{"CdxWBResilRFChangeDampenTime", cdxWBResilObjects.CdxWBResilRFChangeDampenTime})
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilRFChangeTriggerPercentage", types.YLeaf{"CdxWBResilRFChangeTriggerPercentage", cdxWBResilObjects.CdxWBResilRFChangeTriggerPercentage})
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilRFChangeTriggerCount", types.YLeaf{"CdxWBResilRFChangeTriggerCount", cdxWBResilObjects.CdxWBResilRFChangeTriggerCount})
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilRFChangeTriggerMoveSecondary", types.YLeaf{"CdxWBResilRFChangeTriggerMoveSecondary", cdxWBResilObjects.CdxWBResilRFChangeTriggerMoveSecondary})
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilNotificationEnable", types.YLeaf{"CdxWBResilNotificationEnable", cdxWBResilObjects.CdxWBResilNotificationEnable})
    cdxWBResilObjects.EntityData.Leafs.Append("cdxWBResilNotificationsInterval", types.YLeaf{"CdxWBResilNotificationsInterval", cdxWBResilObjects.CdxWBResilNotificationsInterval})

    cdxWBResilObjects.EntityData.YListKeys = []string {}

    return &(cdxWBResilObjects.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a cisco private object. Setting this object to true(1) enables d2.0
    // loadbalancing on CMTS and allows user to further config other options for
    // d3.0 loadbalancing(cdxCmtsD30LBEnable, cdxCmtsD30LBUpstreamEnable
    // cdxCmtsD30LBStaticEnable  		and cdxCmtsD30LBDynEnable).  Setting it to
    // false(2) disables all oad balancing operations. The type is
    // CdxCmtsDocsisLBEnable.
    CdxCmtsDocsisLBEnable interface{}

    // Setting this object to true(1) enables d3.0 static loadbalancing on CMTS
    // and allows user to further config other objects for d3.0
    // loadbalancing(cdxCmtsD30LBUpstreamEnable and cdxCmtsD30LBStaticEnable and
    // cdxCmtsD30LBDynEnable) . Setting it to false(2) disables d3.0
    // loadbalancing. The type is CdxCmtsD30LBEnable.
    CdxCmtsD30LBEnable interface{}

    // Setting this object to true(1) enables upstream loadbalancing in docsis 3.0
    // static loadbalancing. Default is false(2).Only if docsis-enable and
    // docsis30-enable set to true can this object take effect. The type is
    // CdxCmtsD30LBUpstreamEnable.
    CdxCmtsD30LBUpstreamEnable interface{}

    // Setting this to true(1) enables autonomous D30 LB to move wideband modems
    // if LB group is not in a balancing state.Default is false(2). Only if
    // docsis-enable and docsis30-enable is set to true can this object set to
    // true(1). The type is CdxCmtsD30LBStaticEnable.
    CdxCmtsD30LBStaticEnable interface{}

    // Setting this to true(1) enables autonomous D30 LB to move wideband modems
    // if LB group is not in a balancing state.Default is false(2). Only if
    // docsis-enable and docsis30-enable is set to true can this object set to
    // true(1). The type is CdxCmtsD30LBDynEnable.
    CdxCmtsD30LBDynEnable interface{}
}

func (cdxCmtsDocsisLBObjects *CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects) GetEntityData() *types.CommonEntityData {
    cdxCmtsDocsisLBObjects.EntityData.YFilter = cdxCmtsDocsisLBObjects.YFilter
    cdxCmtsDocsisLBObjects.EntityData.YangName = "cdxCmtsDocsisLBObjects"
    cdxCmtsDocsisLBObjects.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsDocsisLBObjects.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsDocsisLBObjects.EntityData.SegmentPath = "cdxCmtsDocsisLBObjects"
    cdxCmtsDocsisLBObjects.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsDocsisLBObjects.EntityData.SegmentPath
    cdxCmtsDocsisLBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsDocsisLBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsDocsisLBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsDocsisLBObjects.EntityData.Children = types.NewOrderedMap()
    cdxCmtsDocsisLBObjects.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsDocsisLBObjects.EntityData.Leafs.Append("cdxCmtsDocsisLBEnable", types.YLeaf{"CdxCmtsDocsisLBEnable", cdxCmtsDocsisLBObjects.CdxCmtsDocsisLBEnable})
    cdxCmtsDocsisLBObjects.EntityData.Leafs.Append("cdxCmtsD30LBEnable", types.YLeaf{"CdxCmtsD30LBEnable", cdxCmtsDocsisLBObjects.CdxCmtsD30LBEnable})
    cdxCmtsDocsisLBObjects.EntityData.Leafs.Append("cdxCmtsD30LBUpstreamEnable", types.YLeaf{"CdxCmtsD30LBUpstreamEnable", cdxCmtsDocsisLBObjects.CdxCmtsD30LBUpstreamEnable})
    cdxCmtsDocsisLBObjects.EntityData.Leafs.Append("cdxCmtsD30LBStaticEnable", types.YLeaf{"CdxCmtsD30LBStaticEnable", cdxCmtsDocsisLBObjects.CdxCmtsD30LBStaticEnable})
    cdxCmtsDocsisLBObjects.EntityData.Leafs.Append("cdxCmtsD30LBDynEnable", types.YLeaf{"CdxCmtsD30LBDynEnable", cdxCmtsDocsisLBObjects.CdxCmtsD30LBDynEnable})

    cdxCmtsDocsisLBObjects.EntityData.YListKeys = []string {}

    return &(cdxCmtsDocsisLBObjects.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable represents this object set to true(1)
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable string

const (
    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable_true_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable = "true"

    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable_false_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBDynEnable = "false"
)

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable represents Setting it to false(2) disables d3.0 loadbalancing.
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable string

const (
    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable_true_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable = "true"

    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable_false_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBEnable = "false"
)

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable represents this object set to true(1)
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable string

const (
    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable_true_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable = "true"

    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable_false_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBStaticEnable = "false"
)

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable represents take effect.
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable string

const (
    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable_true_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable = "true"

    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable_false_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsD30LBUpstreamEnable = "false"
)

// CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable represents Setting it to false(2) disables all oad balancing operations.
type CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable string

const (
    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable_true_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable = "true"

    CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable_false_ CISCODOCSEXTMIB_CdxCmtsDocsisLBObjects_CdxCmtsDocsisLBEnable = "false"
)

// CISCODOCSEXTMIB_CdxQosCtrlUpTable
// For each upstream interface, this table maintains a number 
// of objects related to Quality of Service scheduler which uses
// these attributes to control cable modem registration. 
type CISCODOCSEXTMIB_CdxQosCtrlUpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of attributes for each upstream MAC scheduler  that supports Quality
    // of Service.  Entries in this table exist for each ifEntry with ifType of
    // docsCableUpstream(129). . The type is slice of
    // CISCODOCSEXTMIB_CdxQosCtrlUpTable_CdxQosCtrlUpEntry.
    CdxQosCtrlUpEntry []*CISCODOCSEXTMIB_CdxQosCtrlUpTable_CdxQosCtrlUpEntry
}

func (cdxQosCtrlUpTable *CISCODOCSEXTMIB_CdxQosCtrlUpTable) GetEntityData() *types.CommonEntityData {
    cdxQosCtrlUpTable.EntityData.YFilter = cdxQosCtrlUpTable.YFilter
    cdxQosCtrlUpTable.EntityData.YangName = "cdxQosCtrlUpTable"
    cdxQosCtrlUpTable.EntityData.BundleName = "cisco_ios_xe"
    cdxQosCtrlUpTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxQosCtrlUpTable.EntityData.SegmentPath = "cdxQosCtrlUpTable"
    cdxQosCtrlUpTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxQosCtrlUpTable.EntityData.SegmentPath
    cdxQosCtrlUpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxQosCtrlUpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxQosCtrlUpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxQosCtrlUpTable.EntityData.Children = types.NewOrderedMap()
    cdxQosCtrlUpTable.EntityData.Children.Append("cdxQosCtrlUpEntry", types.YChild{"CdxQosCtrlUpEntry", nil})
    for i := range cdxQosCtrlUpTable.CdxQosCtrlUpEntry {
        cdxQosCtrlUpTable.EntityData.Children.Append(types.GetSegmentPath(cdxQosCtrlUpTable.CdxQosCtrlUpEntry[i]), types.YChild{"CdxQosCtrlUpEntry", cdxQosCtrlUpTable.CdxQosCtrlUpEntry[i]})
    }
    cdxQosCtrlUpTable.EntityData.Leafs = types.NewOrderedMap()

    cdxQosCtrlUpTable.EntityData.YListKeys = []string {}

    return &(cdxQosCtrlUpTable.EntityData)
}

// CISCODOCSEXTMIB_CdxQosCtrlUpTable_CdxQosCtrlUpEntry
// A list of attributes for each upstream MAC scheduler 
// that supports Quality of Service.  Entries in this table
// exist for each ifEntry with ifType of docsCableUpstream(129). 
type CISCODOCSEXTMIB_CdxQosCtrlUpTable_CdxQosCtrlUpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The admission control status for minimum guaranteed upstream  bandwidth
    // scheduling service requests for this upstream.  When this object is set to
    // 'true', if there is a new modem  with minimum guaranteed upstream bandwidth
    // scheduling service in its QoS class requesting to be supported in this
    // upstream, the upstream scheduler will check the virtual reserved  bandwidth
    // remaining capacity before giving admission to this  new modem. If there is
    // not enough reserved upstream bandwidth to serve the modem's minimum
    // guaranteed bandwidth, the  registration request will be rejected.    This
    // object is set to 'false' to disable admission control. That is, there will
    // be no checking for bandwidth capacity and the upstream interface scheduler
    // just admits modem registration requests.   This object is not meant for
    // Unsolicited Grant Service(UGS)  scheduling service as admission control is
    // a requirement in  this case. . The type is bool.
    CdxQosCtrlUpAdmissionCtrl interface{}

    // The percentage of upstream maximum reserved bandwidth to the  raw bandwidth
    // if the admission control is enabled on this  upstream.   For example, if
    // the upstream interface has raw bandwidth  1,600,000 bits/second and
    // cdxQosCtrlUpMaxRsvdBWPercent is 200  percent, then this upstream scheduler
    // will set the maximum of  virtual reserved bandwidth capacity to 3,200,000
    // bits/second  (1,600,000 * 2) to serve cable modems with minimum guaranteed 
    // upstream bandwidth.    The default value is 100 percent (that is, maximum
    // reserved  bandwidth is the raw bandwidth.) Whenever the admission control 
    // is changed (on to off, off to on), this value will be reset to  the default
    // value 100.    If the admission control is disabled, the value will be reset
    // to 100 (the default value). . The type is interface{} with range: 10..1000.
    // Units are percent.
    CdxQosCtrlUpMaxRsvdBWPercent interface{}

    // The count of cable modem registration requests rejected on  this upstream
    // interface due to insufficient reserved  bandwidth for serving the cable
    // modems with Unsolicited  Grant Service (UGS) scheduling service when UGS is
    // supported and for serving the cable modems with minimum  guaranteed
    // bandwidth in its Quality of Service class when admission control is enabled
    // on this upstream interface . The type is interface{} with range:
    // 0..4294967295.
    CdxQosCtrlUpAdmissionRejects interface{}

    // The current total reserved bandwidth in bits per second of  this upstream
    // interface.  It is the sum of all cable modems' minimum guaranteed bandwidth
    // in bits per second currently  supported on this upstream. . The type is
    // interface{} with range: 0..102400000. Units are bits/second.
    CdxQosCtrlUpReservedBW interface{}

    // The maximum virtual bandwidth capacity of this upstream interface if the
    // admission control is enabled. It is the raw bandwidth  in bits per second
    // times the percentage. If the admission  control is disabled, then this
    // object will contain the value  zero. . The type is interface{} with range:
    // 0..102400000. Units are bits/second.
    CdxQosCtrlUpMaxVirtualBW interface{}
}

func (cdxQosCtrlUpEntry *CISCODOCSEXTMIB_CdxQosCtrlUpTable_CdxQosCtrlUpEntry) GetEntityData() *types.CommonEntityData {
    cdxQosCtrlUpEntry.EntityData.YFilter = cdxQosCtrlUpEntry.YFilter
    cdxQosCtrlUpEntry.EntityData.YangName = "cdxQosCtrlUpEntry"
    cdxQosCtrlUpEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxQosCtrlUpEntry.EntityData.ParentYangName = "cdxQosCtrlUpTable"
    cdxQosCtrlUpEntry.EntityData.SegmentPath = "cdxQosCtrlUpEntry" + types.AddKeyToken(cdxQosCtrlUpEntry.IfIndex, "ifIndex")
    cdxQosCtrlUpEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxQosCtrlUpTable/" + cdxQosCtrlUpEntry.EntityData.SegmentPath
    cdxQosCtrlUpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxQosCtrlUpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxQosCtrlUpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxQosCtrlUpEntry.EntityData.Children = types.NewOrderedMap()
    cdxQosCtrlUpEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxQosCtrlUpEntry.IfIndex})
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("cdxQosCtrlUpAdmissionCtrl", types.YLeaf{"CdxQosCtrlUpAdmissionCtrl", cdxQosCtrlUpEntry.CdxQosCtrlUpAdmissionCtrl})
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("cdxQosCtrlUpMaxRsvdBWPercent", types.YLeaf{"CdxQosCtrlUpMaxRsvdBWPercent", cdxQosCtrlUpEntry.CdxQosCtrlUpMaxRsvdBWPercent})
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("cdxQosCtrlUpAdmissionRejects", types.YLeaf{"CdxQosCtrlUpAdmissionRejects", cdxQosCtrlUpEntry.CdxQosCtrlUpAdmissionRejects})
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("cdxQosCtrlUpReservedBW", types.YLeaf{"CdxQosCtrlUpReservedBW", cdxQosCtrlUpEntry.CdxQosCtrlUpReservedBW})
    cdxQosCtrlUpEntry.EntityData.Leafs.Append("cdxQosCtrlUpMaxVirtualBW", types.YLeaf{"CdxQosCtrlUpMaxVirtualBW", cdxQosCtrlUpEntry.CdxQosCtrlUpMaxVirtualBW})

    cdxQosCtrlUpEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxQosCtrlUpEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxQosIfRateLimitTable
// This table describes the attributes of rate limiting for 
// schedulers in cable upstream and downstream interfaces that 
// support Quality of Service.  The rate limiting process is 
// to ensure the Quality of Service and fairness. 
type CISCODOCSEXTMIB_CdxQosIfRateLimitTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the rate limiting attributes for cable upstream and  downstream
    // interfaces schedulers that support Quality of  Service. Entries in this
    // table exist for each ifEntry with  ifType of docsCableUpstream(129) and
    // docsCableDownstream(128). The type is slice of
    // CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry.
    CdxQosIfRateLimitEntry []*CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry
}

func (cdxQosIfRateLimitTable *CISCODOCSEXTMIB_CdxQosIfRateLimitTable) GetEntityData() *types.CommonEntityData {
    cdxQosIfRateLimitTable.EntityData.YFilter = cdxQosIfRateLimitTable.YFilter
    cdxQosIfRateLimitTable.EntityData.YangName = "cdxQosIfRateLimitTable"
    cdxQosIfRateLimitTable.EntityData.BundleName = "cisco_ios_xe"
    cdxQosIfRateLimitTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxQosIfRateLimitTable.EntityData.SegmentPath = "cdxQosIfRateLimitTable"
    cdxQosIfRateLimitTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxQosIfRateLimitTable.EntityData.SegmentPath
    cdxQosIfRateLimitTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxQosIfRateLimitTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxQosIfRateLimitTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxQosIfRateLimitTable.EntityData.Children = types.NewOrderedMap()
    cdxQosIfRateLimitTable.EntityData.Children.Append("cdxQosIfRateLimitEntry", types.YChild{"CdxQosIfRateLimitEntry", nil})
    for i := range cdxQosIfRateLimitTable.CdxQosIfRateLimitEntry {
        cdxQosIfRateLimitTable.EntityData.Children.Append(types.GetSegmentPath(cdxQosIfRateLimitTable.CdxQosIfRateLimitEntry[i]), types.YChild{"CdxQosIfRateLimitEntry", cdxQosIfRateLimitTable.CdxQosIfRateLimitEntry[i]})
    }
    cdxQosIfRateLimitTable.EntityData.Leafs = types.NewOrderedMap()

    cdxQosIfRateLimitTable.EntityData.YListKeys = []string {}

    return &(cdxQosIfRateLimitTable.EntityData)
}

// CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry
// List of the rate limiting attributes for cable upstream and 
// downstream interfaces schedulers that support Quality of 
// Service. Entries in this table exist for each ifEntry with 
// ifType of docsCableUpstream(129) and docsCableDownstream(128).
type CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // To ensure fairness, the CMTS will throttle the rate for bandwidth  request
    // (upstream)/packet sent (downstream) at which CMTS issues  grants(upstream)
    // or allow packet to be send(downstream) such that  the flow never gets more
    // than its provisioned peak rate in bps.   There are two directions for every
    // Service Id (Sid) traffic:  downstream and upstream. Each direction is
    // called a service flow  here and assigned one token bucket with chosen
    // algorithm.   The enumerations for the rate limiting algorithm are: 
    // noRateLimit(1): The rate limiting is disabled. No rate limiting. 
    // oneSecBurst(2): Bursty 1 second token bucket algorithm.  carLike(3)    :
    // Average token usage (CAR-like) algorithm   wtExPacketDiscard(4) : Weighted
    // excess packet discard algorithm.  shaping(5): token bucket algorithm with
    // shaping  Upstream supports the following:    No rate limiting (1),   
    // Bursty 1 second token bucket algorithm(2),    Average token usage
    // (CAR-like) algorithm(3),   Token bucket algorithm with shaping(5). 
    // Downstream supports the following:   No rate limiting (1),    Bursty 1
    // second token bucket algorithm(2),   Average token usage (CAR-like)
    // algorithm(3),   Weighted excess packet discard algorithm(4), and   Token
    // bucket algorithm with shaping(5).  Token bucket algorithm with shaping is
    // the default algorithm for upstream if CMTS is in DOCSIS 1.0 mode or DOCSIS
    // 1.1 mode.   Bursty 1 second token bucket algorithm is the  default
    // algorithm for downstream if the CMTS is in DOCSIS 1.0 mode. If it is in
    // DOCSIS 1.1 mode the default algorithm for downstream is  Token bucket
    // algorithm with shaping .  Each algorithm is described as below:   No rate
    // limiting:      The rate limiting process is disabled and no checking     
    // against the maximum bandwidth allowed.     Bursty 1 second token bucket
    // rate limiting algorithm:      In this algorithm, at the start of every 1
    // second interval,      a service flow's token usage is reset to 0, and every
    // time      the modem for that service flow sends a request (upstream) /     
    // packet (downstream) the upstream/downstream bandwidth      token usage is
    // incremented by the size of the      request/packet sent. As long as the
    // service flow's bandwidth      token usage is less than the maximum
    // bandwidth in bits      per second (peak rate limit) its QoS service class  
    // allows, the request/packets will not be restricted.      Once the service
    // flow has sent more than its peak rate in the      one second interval, it
    // is prevented from sending more      data by rejecting request (upstream) or
    // dropping      packets (downstream). This is expected to slow down     the
    // higher layer sources. The token usage counter gets      reset to 0 after
    // the 1 second interval has elapsed. The      modem for that service flow is
    // free to send more data up to the      peak rate limit in the new 1 second
    // interval that follows.      Average token usage (Cisco CAR like) algorithm:
    // This algorithm maintains a continuous average of the      burst token usage
    // of a service flow. There is no sudden      refilling of tokens every 1
    // second interval. Every time a      request/packet is to be handled, the
    // scheduler tries to see      how much time has elapsed since last
    // transmission, and      computes the number of tokens accumulated by this
    // service flow      at its QoS class peak rate. If burst usage of the service
    // flow      is less than tokens accumulated, the burst usage is reset to 0   
    // and request/packet is forwarded. If the service flow has      accumulated
    // fewer tokens than its burst usage, the burst usage      shows an
    // outstanding balance usage after decrementing by the      tokens
    // accumulated. In such cases, the request/packet is still      forwarded,
    // provided the service flow's outstanding usage does      not exceed peak
    // rate limit of its QoS class. If outstanding      burst usage exceeds the
    // peak rate of the class, the service      flow is given some token credit up
    // to a certain maximum credit      limit and the request/packet is forwarded.
    // The request/packet      is dropped when outstanding usage exceeds peak rate
    // and maximum      credit has been used up by this service flow. This
    // algorithm      tracks long term average bandwidth usage of the service flow
    // and controls this average usage at the peak rate limit.    Weighted excess
    // packet discard algorithm:     This rate limiting algorithm is only
    // available as an option      for downstream rate limiting. The algorithm is
    // to maintain an      weighted exponential moving average of the loss rate of
    // a      service flow over time. The loss rate, expressed in packets,     
    // represents the number of packets that can be sent from this      service
    // flow in a one second interval before a packet will      be dropped. At
    // every one second interval, the loss rate gets      updated using the ratio
    // between the flow peak rate (in bps)      in its QoS profile and the service
    // flow actual usage (in bps).      If the service flow begins to send more
    // than its peak rate      continuously, the number of packets it can send in
    // an one      second interval before experiencing a drop will slowly keep    
    // reducing until cable modem for that service flow slows down      as
    // indicated by actual usage less or equal to the peak rate.     Token bucket
    // algorithm with shaping:      If there is no QoS class peak rate limit,
    // forward the       request/packet without delay. If there is a QoS peak rate
    // limit, every time a request/packet is to be handled, the       scheduler
    // determines the number of bandwidth tokens that this       service flow has
    // accumulated over the elapsed time at its       QoS class peak rate and
    // increments the tokens counter of the       service flow accordingly.  The
    // scheduler limits the token       count to the maximum transmit burst (token
    // bucket depth).        If token count is greater than the number of tokens
    // required       to handle current request/packet, decrement token count by  
    // size of request/packet and forwards the request/packet       without delay.
    // If token count is less than the size of       request/packet, compute the
    // shaping delay time after       which the deficit number of tokens would be
    // available. If       shaping delay time is less than the maximum shaping
    // delay,       decrement tokens count by size of request/packet and      
    // forward this request/packet with the shaping delay in the       shaping
    // delay queue. When the delay time expires, the       request/packet is
    // forwarded. If shaping delay time is       greater than the maximum shaping
    // delay that the subsequent       shaper can handle, the request/packet is
    // dropped. Users can      use cdxQosIfRateLimitShpMaxDelay to configure the
    // the maximum       shaping delay and cdxQosIfRateLimitShpGranularity to     
    // configure the shaping granularity.  . The type is CdxQosIfRateLimitAlgm.
    CdxQosIfRateLimitAlgm interface{}

    // Weight for exponential moving average of loss rate for  weighted excess
    // packet discard algorithm to maintain. The higher value of the weight makes
    // the algorithm more sensitive to the recent bandwidth usage by the Sid.  
    // The default value is 1 and whenever the rate limiting algorithm is changed
    // to weighted excess packet discard  algorithm, this value will be reset to
    // the default 1.  If the rate limiting algorithm is not weighted excess 
    // packet discard algorithm, the value will be always the  default value 1. .
    // The type is interface{} with range: 1..4.
    CdxQosIfRateLimitExpWt interface{}

    // The maximum shaping delay in milliseconds. That is, the maximum  amount
    // time of buffering the CMTS will allow for any rate exceeded  flow.  If the
    // max buffering delay is large, the grants/packets of  the flow will be
    // buffered for a longer period of time even though  the flow is rate
    // exceeded. This means fewer chances of drops for such rate exceeded flow.
    // However, too large a max shaping delay  can result is quick drainage of
    // packet buffers at the CMTS, since  several packets will be in the shaping
    // (delay) queue waiting for  their proper transmission time. Another
    // important point to be noted  is that delaying a flows packets (especially
    // TCP flows) for  extended periods of time is useless, since the higher
    // protocol  layers may assume a packet loss after a certain amount of time. 
    // The maximum shaping delay is only applied to rate limit algorithm:  Token
    // bucket algorithm with shaping.  If the rate limit  algorithm is not Token
    // bucket algorithm with shaping, the  value is always na(1) which is not
    // applicable.  If the token count is less than the size of request/packet,
    // CMTS  computes the shaping delay time after which the deficit number of 
    // tokens would be available. If the shaping delay time is greater  than the
    // maximum shaping delay, the request/packet will be dropped.    The
    // enumerations for maximum shaping delay are:   na(1): maximum shaping delay
    // is not applied to the current           rate limit algorithm  msec128(2):
    // maximum shaping delay is 128 milliseconds    msec256(3): maximum shaping
    // delay is 256 milliseconds   msec512(4): maximum shaping delay is 512
    // milliseconds  msec1024(5): maximum shaping delay is 1024 milliseconds   The
    // downstream maximum shaping delay is configurable and the default value is
    // msec128(2). Whenever the downstream rate  limit algorithm is changed to
    // Token bucket algorithm with  shaping from other rate limit algorithm, the
    // value will  be reset to the default value.   The upstream maximum shaping
    // delay is not configurable and it  is read-only value.  . The type is
    // CdxQosIfRateLimitShpMaxDelay.
    CdxQosIfRateLimitShpMaxDelay interface{}

    // The width in milliseconds of each element in shaping  delay queue, that is,
    // the shaping granularity.  The shaping granularity is only applied to rate
    // limit algorithm: Token bucket algorithm with shaping. It  controls how
    // accurately the algorithm quantizes the shaping  delay for a rate exceeded
    // flow. If granularity is large, several  shaping delay values will all be
    // quantized to the same element  in the queue resulting in less accurate rate
    // shaping for the flows  in bits/sec. On the other hand, choosing too small
    // granularity  causes more memory to be used for the shaper block, and also 
    // can cost a bit more in runtime overhead.  If the rate limit algorithm is
    // not Token bucket algorithm with  shaping, the value is always na(1) which
    // is not applicable.  The enumerations for shaping granularity are:   na(1):
    // shaping granularity is not applied to the current           rate limit
    // algorithm    msec1(2): shaping granularity in 1 milliseconds     msec2(3):
    // shaping granularity in 2 milliseconds     msec4(4): shaping granularity in
    // 4 milliseconds     msec8(5): shaping granularity in 8 milliseconds   
    // msec16(6): shaping granularity in 16 milliseconds    The downstream shaping
    // granularity is configurable and the default value is msec4(4). Whenever the
    // downstream rate limit  algorithm is changed to Token bucket algorithm with
    // shaping  from other rate limit algorithm, the value will be reset to the 
    // default value.   The upstream shaping granularity is not configurable and 
    // it is read-only value.  . The type is CdxQosIfRateLimitShpGranularity.
    CdxQosIfRateLimitShpGranularity interface{}
}

func (cdxQosIfRateLimitEntry *CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry) GetEntityData() *types.CommonEntityData {
    cdxQosIfRateLimitEntry.EntityData.YFilter = cdxQosIfRateLimitEntry.YFilter
    cdxQosIfRateLimitEntry.EntityData.YangName = "cdxQosIfRateLimitEntry"
    cdxQosIfRateLimitEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxQosIfRateLimitEntry.EntityData.ParentYangName = "cdxQosIfRateLimitTable"
    cdxQosIfRateLimitEntry.EntityData.SegmentPath = "cdxQosIfRateLimitEntry" + types.AddKeyToken(cdxQosIfRateLimitEntry.IfIndex, "ifIndex")
    cdxQosIfRateLimitEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxQosIfRateLimitTable/" + cdxQosIfRateLimitEntry.EntityData.SegmentPath
    cdxQosIfRateLimitEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxQosIfRateLimitEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxQosIfRateLimitEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxQosIfRateLimitEntry.EntityData.Children = types.NewOrderedMap()
    cdxQosIfRateLimitEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxQosIfRateLimitEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxQosIfRateLimitEntry.IfIndex})
    cdxQosIfRateLimitEntry.EntityData.Leafs.Append("cdxQosIfRateLimitAlgm", types.YLeaf{"CdxQosIfRateLimitAlgm", cdxQosIfRateLimitEntry.CdxQosIfRateLimitAlgm})
    cdxQosIfRateLimitEntry.EntityData.Leafs.Append("cdxQosIfRateLimitExpWt", types.YLeaf{"CdxQosIfRateLimitExpWt", cdxQosIfRateLimitEntry.CdxQosIfRateLimitExpWt})
    cdxQosIfRateLimitEntry.EntityData.Leafs.Append("cdxQosIfRateLimitShpMaxDelay", types.YLeaf{"CdxQosIfRateLimitShpMaxDelay", cdxQosIfRateLimitEntry.CdxQosIfRateLimitShpMaxDelay})
    cdxQosIfRateLimitEntry.EntityData.Leafs.Append("cdxQosIfRateLimitShpGranularity", types.YLeaf{"CdxQosIfRateLimitShpGranularity", cdxQosIfRateLimitEntry.CdxQosIfRateLimitShpGranularity})

    cdxQosIfRateLimitEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxQosIfRateLimitEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm represents      configure the shaping granularity.  
type CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm string

const (
    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm_noRateLimit CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm = "noRateLimit"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm_oneSecBurst CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm = "oneSecBurst"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm_carLike CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm = "carLike"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm_wtExPacketDiscard CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm = "wtExPacketDiscard"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm_shaping CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitAlgm = "shaping"
)

// CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity represents it is read-only value.  
type CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity string

const (
    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_na CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "na"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_msec1 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "msec1"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_msec2 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "msec2"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_msec4 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "msec4"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_msec8 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "msec8"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity_msec16 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpGranularity = "msec16"
)

// CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay represents is read-only value.  
type CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay string

const (
    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay_na CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay = "na"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay_msec128 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay = "msec128"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay_msec256 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay = "msec256"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay_msec512 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay = "msec512"

    CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay_msec1024 CISCODOCSEXTMIB_CdxQosIfRateLimitTable_CdxQosIfRateLimitEntry_CdxQosIfRateLimitShpMaxDelay = "msec1024"
)

// CISCODOCSEXTMIB_CdxCmtsServiceExtTable
// The list contains the additional attributes of a single Service
// ID that provided by docsIfCmtsServiceEntry. 
type CISCODOCSEXTMIB_CdxCmtsServiceExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional objects for docsIfCmtsServiceTable entry including  downstream
    // traffic statistics and excess counts against the  Quality of Service limits
    // for each Service ID. From DOCSIS 1.1 onwards statistics are maintained for
    // each  Service Flow(instead of the Service ID) in the DOCS-QOS-MIB  in
    // docsQosServiceFlowStatsTable objects. For Cable modems not running in
    // DOCSIS 1.0 mode, the objects   cdxIfCmtsServiceOutOctets and
    // cdxIfCmtsServiceOutPackets will only support primary service flow. . The
    // type is slice of
    // CISCODOCSEXTMIB_CdxCmtsServiceExtTable_CdxCmtsServiceExtEntry.
    CdxCmtsServiceExtEntry []*CISCODOCSEXTMIB_CdxCmtsServiceExtTable_CdxCmtsServiceExtEntry
}

func (cdxCmtsServiceExtTable *CISCODOCSEXTMIB_CdxCmtsServiceExtTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsServiceExtTable.EntityData.YFilter = cdxCmtsServiceExtTable.YFilter
    cdxCmtsServiceExtTable.EntityData.YangName = "cdxCmtsServiceExtTable"
    cdxCmtsServiceExtTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsServiceExtTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsServiceExtTable.EntityData.SegmentPath = "cdxCmtsServiceExtTable"
    cdxCmtsServiceExtTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsServiceExtTable.EntityData.SegmentPath
    cdxCmtsServiceExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsServiceExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsServiceExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsServiceExtTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsServiceExtTable.EntityData.Children.Append("cdxCmtsServiceExtEntry", types.YChild{"CdxCmtsServiceExtEntry", nil})
    for i := range cdxCmtsServiceExtTable.CdxCmtsServiceExtEntry {
        cdxCmtsServiceExtTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsServiceExtTable.CdxCmtsServiceExtEntry[i]), types.YChild{"CdxCmtsServiceExtEntry", cdxCmtsServiceExtTable.CdxCmtsServiceExtEntry[i]})
    }
    cdxCmtsServiceExtTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsServiceExtTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsServiceExtTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsServiceExtTable_CdxCmtsServiceExtEntry
// Additional objects for docsIfCmtsServiceTable entry including 
// downstream traffic statistics and excess counts against the 
// Quality of Service limits for each Service ID.
// From DOCSIS 1.1 onwards statistics are maintained for each 
// Service Flow(instead of the Service ID) in the DOCS-QOS-MIB 
// in docsQosServiceFlowStatsTable objects. For Cable modems
// not running in DOCSIS 1.0 mode, the objects  
// cdxIfCmtsServiceOutOctets and cdxIfCmtsServiceOutPackets
// will only support primary service flow. 
type CISCODOCSEXTMIB_CdxCmtsServiceExtTable_CdxCmtsServiceExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..16383. Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceId
    DocsIfCmtsServiceId interface{}

    // The cumulative number of Packet Data octets sent for this  Service ID. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfCmtsServiceOutOctets interface{}

    // The cumulative number of Packet data packets sent for this  Service ID. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfCmtsServiceOutPackets interface{}

    // The number of upstream bandwidth requests which exceeds the maximum
    // upstream bandwidth allowed for a service defined  in the Quality of Service
    // profile associated with this Sid.  The request which exceeds the maximum
    // upstream bandwidth  allowed will be rejected by the upstream's rate
    // limiting  process using one of the rate limiting algorithm.   Note that the
    // value of this counter cannot be directly used  to know the number of
    // upstream packets that got dropped at  the cable modem.  A single upstream
    // packet drop of a modem  can result in up to 16 increments in this counter,
    // since the  modem keeps retrying and keeps getting bandwidth request  drops
    // at CMTS if it has consumed its peak rate.  . The type is interface{} with
    // range: 0..4294967295.
    CdxQosMaxUpBWExcessRequests interface{}

    // The number of downstream bandwidth packets which exceeds the maximum
    // downstream bandwidth allowed for a service defined in the Quality of
    // Service profile associated with this Sid.  The packet which exceeds the
    // maximum downstream bandwidth  allowed will be dropped by the downstream's
    // rate limiting  process using one of the rate limiting algorithm. . The type
    // is interface{} with range: 0..4294967295.
    CdxQosMaxDownBWExcessPackets interface{}

    // The cumulative number of Packet Data octets received on this Service ID.
    // The count does not include the size of the Cable MAC header. This object is
    // a 64-bit version of docsIfCmtsServiceInOctets. The type is interface{} with
    // range: 0..18446744073709551615.
    CdxIfCmtsServiceHCInOctets interface{}

    // The cumulative number of Packet Data packets received on this Service ID.
    // This object is a 64-bit version of docsIfCmtsServiceInPackets. The type is
    // interface{} with range: 0..18446744073709551615.
    CdxIfCmtsServiceHCInPackets interface{}

    // The cumulative number of Packet Data octets sent for this Service ID. This
    // object is a 64-bit version of cdxIfCmtsServiceOutOctets. The type is
    // interface{} with range: 0..18446744073709551615.
    CdxIfCmtsServiceHCOutOctets interface{}

    // The cumulative number of Packet Data packets sent for this Service ID. This
    // object is a 64-bit version of cdxIfCmtsServiceOutPackets. The type is
    // interface{} with range: 0..18446744073709551615.
    CdxIfCmtsServiceHCOutPackets interface{}
}

func (cdxCmtsServiceExtEntry *CISCODOCSEXTMIB_CdxCmtsServiceExtTable_CdxCmtsServiceExtEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsServiceExtEntry.EntityData.YFilter = cdxCmtsServiceExtEntry.YFilter
    cdxCmtsServiceExtEntry.EntityData.YangName = "cdxCmtsServiceExtEntry"
    cdxCmtsServiceExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsServiceExtEntry.EntityData.ParentYangName = "cdxCmtsServiceExtTable"
    cdxCmtsServiceExtEntry.EntityData.SegmentPath = "cdxCmtsServiceExtEntry" + types.AddKeyToken(cdxCmtsServiceExtEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxCmtsServiceExtEntry.DocsIfCmtsServiceId, "docsIfCmtsServiceId")
    cdxCmtsServiceExtEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsServiceExtTable/" + cdxCmtsServiceExtEntry.EntityData.SegmentPath
    cdxCmtsServiceExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsServiceExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsServiceExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsServiceExtEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsServiceExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxCmtsServiceExtEntry.IfIndex})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("docsIfCmtsServiceId", types.YLeaf{"DocsIfCmtsServiceId", cdxCmtsServiceExtEntry.DocsIfCmtsServiceId})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceOutOctets", types.YLeaf{"CdxIfCmtsServiceOutOctets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceOutOctets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceOutPackets", types.YLeaf{"CdxIfCmtsServiceOutPackets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceOutPackets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxQosMaxUpBWExcessRequests", types.YLeaf{"CdxQosMaxUpBWExcessRequests", cdxCmtsServiceExtEntry.CdxQosMaxUpBWExcessRequests})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxQosMaxDownBWExcessPackets", types.YLeaf{"CdxQosMaxDownBWExcessPackets", cdxCmtsServiceExtEntry.CdxQosMaxDownBWExcessPackets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceHCInOctets", types.YLeaf{"CdxIfCmtsServiceHCInOctets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceHCInOctets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceHCInPackets", types.YLeaf{"CdxIfCmtsServiceHCInPackets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceHCInPackets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceHCOutOctets", types.YLeaf{"CdxIfCmtsServiceHCOutOctets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceHCOutOctets})
    cdxCmtsServiceExtEntry.EntityData.Leafs.Append("cdxIfCmtsServiceHCOutPackets", types.YLeaf{"CdxIfCmtsServiceHCOutPackets", cdxCmtsServiceExtEntry.CdxIfCmtsServiceHCOutPackets})

    cdxCmtsServiceExtEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIfCmtsServiceId"}

    return &(cdxCmtsServiceExtEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxUpInfoElemStatsTable
// The table contains the attributes of a particular 
// Information Element type for each instance of the MAC 
// scheduler. It is indexed by upstream ifIndex. An Entry
// exists for each ifEntry with ifType of docsCableUpstream(129)
// Since each upstream has an instance of a MAC Scheduler so 
// this table has the per MAC scheduler slot information on a
// per Information Element basis. 
type CISCODOCSEXTMIB_CdxUpInfoElemStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of statistics for a type of Information Element(IE) which defines
    // the allowed usage for a range of upstream mini slots. One entry exists for
    // each Information Element (IE) in a upstream which ifType is
    // docsCableUpstream (12). The type is slice of
    // CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry.
    CdxUpInfoElemStatsEntry []*CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry
}

func (cdxUpInfoElemStatsTable *CISCODOCSEXTMIB_CdxUpInfoElemStatsTable) GetEntityData() *types.CommonEntityData {
    cdxUpInfoElemStatsTable.EntityData.YFilter = cdxUpInfoElemStatsTable.YFilter
    cdxUpInfoElemStatsTable.EntityData.YangName = "cdxUpInfoElemStatsTable"
    cdxUpInfoElemStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cdxUpInfoElemStatsTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxUpInfoElemStatsTable.EntityData.SegmentPath = "cdxUpInfoElemStatsTable"
    cdxUpInfoElemStatsTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxUpInfoElemStatsTable.EntityData.SegmentPath
    cdxUpInfoElemStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxUpInfoElemStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxUpInfoElemStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxUpInfoElemStatsTable.EntityData.Children = types.NewOrderedMap()
    cdxUpInfoElemStatsTable.EntityData.Children.Append("cdxUpInfoElemStatsEntry", types.YChild{"CdxUpInfoElemStatsEntry", nil})
    for i := range cdxUpInfoElemStatsTable.CdxUpInfoElemStatsEntry {
        cdxUpInfoElemStatsTable.EntityData.Children.Append(types.GetSegmentPath(cdxUpInfoElemStatsTable.CdxUpInfoElemStatsEntry[i]), types.YChild{"CdxUpInfoElemStatsEntry", cdxUpInfoElemStatsTable.CdxUpInfoElemStatsEntry[i]})
    }
    cdxUpInfoElemStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cdxUpInfoElemStatsTable.EntityData.YListKeys = []string {}

    return &(cdxUpInfoElemStatsTable.EntityData)
}

// CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry
// The list of statistics for a type of Information Element(IE)
// which defines the allowed usage for a range of upstream mini
// slots. One entry exists for each Information Element (IE) in
// a upstream which ifType is docsCableUpstream (12).
type CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This entry describes the Information Element (IE)
    // type. Enumerations are : reqIE(1)          : Request Information Element,  
    // The request Information Element provides                     an upstream
    // interval in which a CM can                     request the CMTS for
    // bandwidth on the                      upstream channel. reqOrDataIE(2)    :
    // Request/Data Information Element,                     The Request/data
    // Information Element                      provides an upstream interval in
    // which                      request may be made by the CM to the            
    // CMTS for bandwidth or short data                      packets may be
    // transmitted on the                     upstream channel. initMtnIE(3)     
    // : Initial Maintenance Information Element,                     The Initial
    // Maintenance Information                      Element provides an interval
    // in which                      new stations may join the network.
    // stnMtnIE(4)       : Station Maintenance Information Element,               
    // The Station Maintenance Information                      Element provides
    // an interval in which                      stations are expected to perform
    // some                      aspect of routine network maintenance ,          
    // such as ranging or power adjustment. shortGrantIE(5)   : Short Data Grant
    // Information Element,                     Short data grant Information
    // Element                     provides the CM an opportunity to              
    // transmit one or more upstream PDU's.                     Short data grants
    // are used with                      intervals equal to or less than the     
    // maximum burst size for the usage                      specified in the
    // Upstream Channel                      Descriptor. longGrantIE(6)    : Long
    // Data Grant Information Element,                     The long data grant
    // Information Element                     also provides the CM an opportunity
    // to                     transmit one or more upstream PDU's.                
    // All long data grant Information Elements                     must have a
    // larger number of mini-slots                     than that defined for a
    // short data grant                     Information Element profile defined in
    // the Upstream Channel Descriptor. . The type is CdxUpInfoElemStatsNameCode.
    CdxUpInfoElemStatsNameCode interface{}

    // The current number of mini-slots used for the Information  Element type.
    // The value is only a snapshot since the  current number of mini-slots of
    // this IE type could be changing rapidly. . The type is interface{} with
    // range: -2147483648..2147483647.
    CdxUpInfoElemStatsIEType interface{}
}

func (cdxUpInfoElemStatsEntry *CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry) GetEntityData() *types.CommonEntityData {
    cdxUpInfoElemStatsEntry.EntityData.YFilter = cdxUpInfoElemStatsEntry.YFilter
    cdxUpInfoElemStatsEntry.EntityData.YangName = "cdxUpInfoElemStatsEntry"
    cdxUpInfoElemStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxUpInfoElemStatsEntry.EntityData.ParentYangName = "cdxUpInfoElemStatsTable"
    cdxUpInfoElemStatsEntry.EntityData.SegmentPath = "cdxUpInfoElemStatsEntry" + types.AddKeyToken(cdxUpInfoElemStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxUpInfoElemStatsEntry.CdxUpInfoElemStatsNameCode, "cdxUpInfoElemStatsNameCode")
    cdxUpInfoElemStatsEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxUpInfoElemStatsTable/" + cdxUpInfoElemStatsEntry.EntityData.SegmentPath
    cdxUpInfoElemStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxUpInfoElemStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxUpInfoElemStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxUpInfoElemStatsEntry.EntityData.Children = types.NewOrderedMap()
    cdxUpInfoElemStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxUpInfoElemStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxUpInfoElemStatsEntry.IfIndex})
    cdxUpInfoElemStatsEntry.EntityData.Leafs.Append("cdxUpInfoElemStatsNameCode", types.YLeaf{"CdxUpInfoElemStatsNameCode", cdxUpInfoElemStatsEntry.CdxUpInfoElemStatsNameCode})
    cdxUpInfoElemStatsEntry.EntityData.Leafs.Append("cdxUpInfoElemStatsIEType", types.YLeaf{"CdxUpInfoElemStatsIEType", cdxUpInfoElemStatsEntry.CdxUpInfoElemStatsIEType})

    cdxUpInfoElemStatsEntry.EntityData.YListKeys = []string {"IfIndex", "CdxUpInfoElemStatsNameCode"}

    return &(cdxUpInfoElemStatsEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode represents                     the Upstream Channel Descriptor. 
type CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode string

const (
    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_reqIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "reqIE"

    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_reqOrDataIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "reqOrDataIE"

    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_initMtnIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "initMtnIE"

    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_stnMtnIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "stnMtnIE"

    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_shortGrantIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "shortGrantIE"

    CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode_longGrantIE CISCODOCSEXTMIB_CdxUpInfoElemStatsTable_CdxUpInfoElemStatsEntry_CdxUpInfoElemStatsNameCode = "longGrantIE"
)

// CISCODOCSEXTMIB_CdxBWQueueTable
// This table describes the attributes of queues  
// in cable interfaces schedulers that support 
// Quality of Service.
type CISCODOCSEXTMIB_CdxBWQueueTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of queue attributes in cable upstream and downstream interfaces
    // schedulers that supports Quality of Service.  Entries in this table exist
    // for each ifEntry with ifType of  docsCableUpstream(129) and
    // docsCableDownstream(128). . The type is slice of
    // CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry.
    CdxBWQueueEntry []*CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry
}

func (cdxBWQueueTable *CISCODOCSEXTMIB_CdxBWQueueTable) GetEntityData() *types.CommonEntityData {
    cdxBWQueueTable.EntityData.YFilter = cdxBWQueueTable.YFilter
    cdxBWQueueTable.EntityData.YangName = "cdxBWQueueTable"
    cdxBWQueueTable.EntityData.BundleName = "cisco_ios_xe"
    cdxBWQueueTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxBWQueueTable.EntityData.SegmentPath = "cdxBWQueueTable"
    cdxBWQueueTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxBWQueueTable.EntityData.SegmentPath
    cdxBWQueueTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBWQueueTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBWQueueTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBWQueueTable.EntityData.Children = types.NewOrderedMap()
    cdxBWQueueTable.EntityData.Children.Append("cdxBWQueueEntry", types.YChild{"CdxBWQueueEntry", nil})
    for i := range cdxBWQueueTable.CdxBWQueueEntry {
        cdxBWQueueTable.EntityData.Children.Append(types.GetSegmentPath(cdxBWQueueTable.CdxBWQueueEntry[i]), types.YChild{"CdxBWQueueEntry", cdxBWQueueTable.CdxBWQueueEntry[i]})
    }
    cdxBWQueueTable.EntityData.Leafs = types.NewOrderedMap()

    cdxBWQueueTable.EntityData.YListKeys = []string {}

    return &(cdxBWQueueTable.EntityData)
}

// CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry
// The list of queue attributes in cable upstream and downstream
// interfaces schedulers that supports Quality of Service. 
// Entries in this table exist for each ifEntry with ifType of 
// docsCableUpstream(129) and docsCableDownstream(128). 
type CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The name code for the queue.  cirQ :CIR queue. The
    // queue is for Committed Information Rate        (CIR) type of service which
    // serves Service IDs which       have minimum guaranteed rate in its QoS
    // profile.       It is applicable if CMTS is running is either of       
    // DOCSIS 1.0 or 1.1 modes.For DOCSIS 1.1 it has        priority 8.           
    // tbeQ :TBE Queue.The queue is for TIERED BEST EFFORT type        service
    // which serves Service IDs which does not have        minimum guaranteed rate
    // in its QoS profile. It is        only applicable if CMTS is running in
    // DOCSIS 1.0       mode.  p0BEGrantQ-p7BEGrantQ : BEST EFFORT Queue       The
    // queues p0BEGrantQ to P7BEGrantQ are for TIERED        BEST EFFORT type
    // service which serves Service IDs        which do not have minimum
    // guaranteed rate specified       in the QoS parameters. P0 has lowest
    // priority and P7       has highest.Best Effort type is purely handled with  
    // prioritization in FIFO's and hence demands more        number of queues.
    // These queues are applicable only if       CMTS is running under mode DOCSIS
    // 1.1.                               rngPollQ : RngPoll queue.       The
    // queue is for the ranging SID's.It has the highest       priority. This
    // queue information is only provided if       CMTS is running under mode
    // DOCSIS 1.1. . The type is CdxBWQueueNameCode.
    CdxBWQueueNameCode interface{}

    // The relative order of this queue to the other queues within the  cable
    // interface. The smaller number has higher order. That is, 0 is the highest
    // order and 10 is the lowest order.  The  scheduler will serve the requests
    // in higher order queue up to  the number of requests defined in
    // cdxBWQueueNumServedBeforeYield before serving requests in the next higher
    // order queue.    If there are n queues on this interface, the queue order
    // will  be 0 to n-1 and maximum number of requests defined as 
    // cdxBWQueueNumServedBeforeYield in order 0 queue will be served  before the
    // requests in order 1 queue to be served. . The type is interface{} with
    // range: 0..10.
    CdxBWQueueOrder interface{}

    // The maximum number of requests/packets the scheduler can  serve before
    // yielding to another queue. The value 0 means all  requests must be served
    // before yielding to another queue. The  range is 0-50 for DOCSIS 1.0 and for
    // DOCSIS 1.1 it is 0-64. . The type is interface{} with range: 0..64.
    CdxBWQueueNumServedBeforeYield interface{}

    // The queuing type which decides the position of a request/packet within the
    // queue.   unknown : queue type unknown.    other   : not fifo, and not
    // priority.   fifo    : first in first out.   priority: each bandwidth
    // request has a priority and the              position of the request within
    // the queue depends              on its priority.   For DOCSIS1.1 all the
    // priority queues are fifo queues. . The type is CdxBWQueueType.
    CdxBWQueueType interface{}

    // The maximum number of requests/packets which the queue can  support.The
    // range is 0-50 for DOCSIS1.0 and for DOCSIS1.1 it is 0-64. . The type is
    // interface{} with range: 0..64.
    CdxBWQueueMaxDepth interface{}

    // The current number of requests/packets in the queue. The range is 0-50 for
    // DOCSIS1.0 and for DOCSIS1.1 it is 0-64. . The type is interface{} with
    // range: 0..64.
    CdxBWQueueDepth interface{}

    // The number of requests/packets discarded because of queue overflow (queue
    // depth > queue maximum depth). . The type is interface{} with range:
    // 0..4294967295.
    CdxBWQueueDiscards interface{}
}

func (cdxBWQueueEntry *CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry) GetEntityData() *types.CommonEntityData {
    cdxBWQueueEntry.EntityData.YFilter = cdxBWQueueEntry.YFilter
    cdxBWQueueEntry.EntityData.YangName = "cdxBWQueueEntry"
    cdxBWQueueEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxBWQueueEntry.EntityData.ParentYangName = "cdxBWQueueTable"
    cdxBWQueueEntry.EntityData.SegmentPath = "cdxBWQueueEntry" + types.AddKeyToken(cdxBWQueueEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxBWQueueEntry.CdxBWQueueNameCode, "cdxBWQueueNameCode")
    cdxBWQueueEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxBWQueueTable/" + cdxBWQueueEntry.EntityData.SegmentPath
    cdxBWQueueEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBWQueueEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBWQueueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBWQueueEntry.EntityData.Children = types.NewOrderedMap()
    cdxBWQueueEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxBWQueueEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxBWQueueEntry.IfIndex})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueNameCode", types.YLeaf{"CdxBWQueueNameCode", cdxBWQueueEntry.CdxBWQueueNameCode})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueOrder", types.YLeaf{"CdxBWQueueOrder", cdxBWQueueEntry.CdxBWQueueOrder})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueNumServedBeforeYield", types.YLeaf{"CdxBWQueueNumServedBeforeYield", cdxBWQueueEntry.CdxBWQueueNumServedBeforeYield})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueType", types.YLeaf{"CdxBWQueueType", cdxBWQueueEntry.CdxBWQueueType})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueMaxDepth", types.YLeaf{"CdxBWQueueMaxDepth", cdxBWQueueEntry.CdxBWQueueMaxDepth})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueDepth", types.YLeaf{"CdxBWQueueDepth", cdxBWQueueEntry.CdxBWQueueDepth})
    cdxBWQueueEntry.EntityData.Leafs.Append("cdxBWQueueDiscards", types.YLeaf{"CdxBWQueueDiscards", cdxBWQueueEntry.CdxBWQueueDiscards})

    cdxBWQueueEntry.EntityData.YListKeys = []string {"IfIndex", "CdxBWQueueNameCode"}

    return &(cdxBWQueueEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode represents       CMTS is running under mode DOCSIS 1.1. 
type CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode string

const (
    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_cirQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "cirQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_tbeQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "tbeQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p0BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p0BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p1BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p1BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p2BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p2BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p3BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p3BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p4BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p4BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p5BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p5BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p6BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p6BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_p7BEGrantQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "p7BEGrantQ"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode_rngPollQ CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueNameCode = "rngPollQ"
)

// CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType represents   For DOCSIS1.1 all the priority queues are fifo queues. 
type CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType string

const (
    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType_unknown CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType = "unknown"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType_other CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType = "other"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType_fifo CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType = "fifo"

    CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType_priority CISCODOCSEXTMIB_CdxBWQueueTable_CdxBWQueueEntry_CdxBWQueueType = "priority"
)

// CISCODOCSEXTMIB_CdxCmCpeTable
// This table contains information about cable modems (CM) or 
// Customer Premises Equipments (CPE). 
type CISCODOCSEXTMIB_CdxCmCpeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list contains information for a cable modem (CM) or a Customer Premises
    // Equipment (CPE). An entry exist for  each cable modem supported by CMTS and
    // each Customer Premises  Equipment connected to a cable modem supported by
    // CMTS. . The type is slice of CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry.
    CdxCmCpeEntry []*CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry
}

func (cdxCmCpeTable *CISCODOCSEXTMIB_CdxCmCpeTable) GetEntityData() *types.CommonEntityData {
    cdxCmCpeTable.EntityData.YFilter = cdxCmCpeTable.YFilter
    cdxCmCpeTable.EntityData.YangName = "cdxCmCpeTable"
    cdxCmCpeTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmCpeTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmCpeTable.EntityData.SegmentPath = "cdxCmCpeTable"
    cdxCmCpeTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmCpeTable.EntityData.SegmentPath
    cdxCmCpeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmCpeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmCpeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmCpeTable.EntityData.Children = types.NewOrderedMap()
    cdxCmCpeTable.EntityData.Children.Append("cdxCmCpeEntry", types.YChild{"CdxCmCpeEntry", nil})
    for i := range cdxCmCpeTable.CdxCmCpeEntry {
        cdxCmCpeTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmCpeTable.CdxCmCpeEntry[i]), types.YChild{"CdxCmCpeEntry", cdxCmCpeTable.CdxCmCpeEntry[i]})
    }
    cdxCmCpeTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmCpeTable.EntityData.YListKeys = []string {}

    return &(cdxCmCpeTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry
// The list contains information for a cable modem (CM) or a
// Customer Premises Equipment (CPE). An entry exist for 
// each cable modem supported by CMTS and each Customer Premises 
// Equipment connected to a cable modem supported by CMTS. 
type CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Mac address to identify a cable modem or a
    // Customer  Premises Equipment. . The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCmCpeMacAddress interface{}

    // Indicate this entry is for cable modem or Customer Premises  Equipment. 
    // The enumerations are:   cm(1): cable modem  cpe(2): Customer Premises
    // Equipment . The type is CdxCmCpeType.
    CdxCmCpeType interface{}

    // Ip address of the cable modem or Customer Premises Equipment. . The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CdxCmCpeIpAddress interface{}

    // The CMTS cable MAC interface index (ifType of  docsCableMaclayer(127)) that
    // cable modem or Customer Premises  Equipment connects to.    Use
    // cdxCmCpeIfIndex and cdxCmCpeCmtsServiceId to identify an  entry in
    // docsIfCmtsServiceTable.  . The type is interface{} with range:
    // 1..2147483647.
    CdxCmCpeIfIndex interface{}

    // The cable modem's primary Service ID if the type is cm.  The primary
    // Service ID for the CM which the CPE connects if the  type is cpe.  Use
    // cdxCmCpeIfIndex and cdxCmCpeCmtsServiceId to identify an  entry in
    // docsIfCmtsServiceTable.  . The type is interface{} with range: 1..16383.
    CdxCmCpeCmtsServiceId interface{}

    // Pointer to an entry in docsIfCmtsCmStatusTable identifying  status of the
    // CM (which the CPE connects to.) . The type is interface{} with range:
    // 1..2147483647.
    CdxCmCpeCmStatusIndex interface{}

    // ASCII text to identify the Access Group for a CM or CPE.  Access Group is
    // to filter the upstream traffic for that CM or CPE.  . The type is string.
    CdxCmCpeAccessGroup interface{}

    // Setting this object to true(1) causes the device to  reset.  Reading this
    // object always returns false(2).  For cdxCmCpeType value cm(1),  CMTS
    // removes the  CM from the Station Maintenance List and would cause  the CM
    // to reset its interface.  For cdxCmCpeType value cpe(2), CMTS removes the 
    // CPE's MAC address from the internal address table.   It then rediscovers
    // and associates the CPE with the  correct CM during the next DHCP lease
    // cycle.  By resetting  the CPE, the user can replace an existing CPE or
    // change  its network interface card (NIC). The type is bool.
    CdxCmCpeResetNow interface{}

    // Setting this object to true(1) causes the CM/CPE to be deleted.  Reading
    // this object always returns false(2).  For cdxCmCpeType value cm(1),  CMTS
    // delete CM from  its interface.  For cdxCmCpeType value cpe(2), CMTS delete
    // CPE from  its associated CM. The type is bool.
    CdxCmCpeDeleteNow interface{}
}

func (cdxCmCpeEntry *CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry) GetEntityData() *types.CommonEntityData {
    cdxCmCpeEntry.EntityData.YFilter = cdxCmCpeEntry.YFilter
    cdxCmCpeEntry.EntityData.YangName = "cdxCmCpeEntry"
    cdxCmCpeEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmCpeEntry.EntityData.ParentYangName = "cdxCmCpeTable"
    cdxCmCpeEntry.EntityData.SegmentPath = "cdxCmCpeEntry" + types.AddKeyToken(cdxCmCpeEntry.CdxCmCpeMacAddress, "cdxCmCpeMacAddress")
    cdxCmCpeEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmCpeTable/" + cdxCmCpeEntry.EntityData.SegmentPath
    cdxCmCpeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmCpeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmCpeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmCpeEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmCpeEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeMacAddress", types.YLeaf{"CdxCmCpeMacAddress", cdxCmCpeEntry.CdxCmCpeMacAddress})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeType", types.YLeaf{"CdxCmCpeType", cdxCmCpeEntry.CdxCmCpeType})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeIpAddress", types.YLeaf{"CdxCmCpeIpAddress", cdxCmCpeEntry.CdxCmCpeIpAddress})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeIfIndex", types.YLeaf{"CdxCmCpeIfIndex", cdxCmCpeEntry.CdxCmCpeIfIndex})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeCmtsServiceId", types.YLeaf{"CdxCmCpeCmtsServiceId", cdxCmCpeEntry.CdxCmCpeCmtsServiceId})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeCmStatusIndex", types.YLeaf{"CdxCmCpeCmStatusIndex", cdxCmCpeEntry.CdxCmCpeCmStatusIndex})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeAccessGroup", types.YLeaf{"CdxCmCpeAccessGroup", cdxCmCpeEntry.CdxCmCpeAccessGroup})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeResetNow", types.YLeaf{"CdxCmCpeResetNow", cdxCmCpeEntry.CdxCmCpeResetNow})
    cdxCmCpeEntry.EntityData.Leafs.Append("cdxCmCpeDeleteNow", types.YLeaf{"CdxCmCpeDeleteNow", cdxCmCpeEntry.CdxCmCpeDeleteNow})

    cdxCmCpeEntry.EntityData.YListKeys = []string {"CdxCmCpeMacAddress"}

    return &(cdxCmCpeEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType represents  cpe(2): Customer Premises Equipment 
type CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType string

const (
    CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType_cm CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType = "cm"

    CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType_cpe CISCODOCSEXTMIB_CdxCmCpeTable_CdxCmCpeEntry_CdxCmCpeType = "cpe"
)

// CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable
// The list contains the additional CM status information. 
type CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional objects for docsIfCmtsCmStatusTable entry. . The type is slice
    // of CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry.
    CdxCmtsCmStatusExtEntry []*CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry
}

func (cdxCmtsCmStatusExtTable *CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmStatusExtTable.EntityData.YFilter = cdxCmtsCmStatusExtTable.YFilter
    cdxCmtsCmStatusExtTable.EntityData.YangName = "cdxCmtsCmStatusExtTable"
    cdxCmtsCmStatusExtTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmStatusExtTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsCmStatusExtTable.EntityData.SegmentPath = "cdxCmtsCmStatusExtTable"
    cdxCmtsCmStatusExtTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsCmStatusExtTable.EntityData.SegmentPath
    cdxCmtsCmStatusExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmStatusExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmStatusExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmStatusExtTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmStatusExtTable.EntityData.Children.Append("cdxCmtsCmStatusExtEntry", types.YChild{"CdxCmtsCmStatusExtEntry", nil})
    for i := range cdxCmtsCmStatusExtTable.CdxCmtsCmStatusExtEntry {
        cdxCmtsCmStatusExtTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsCmStatusExtTable.CdxCmtsCmStatusExtEntry[i]), types.YChild{"CdxCmtsCmStatusExtEntry", cdxCmtsCmStatusExtTable.CdxCmtsCmStatusExtEntry[i]})
    }
    cdxCmtsCmStatusExtTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsCmStatusExtTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsCmStatusExtTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry
// Additional objects for docsIfCmtsCmStatusTable entry. 
type CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // Current Cable Modem connectivity state. The object extends  states in
    // docsIfCmtsCmStatusValue in more details.   The enumerations are: offline(1)
    // : modem considered offline. others(2)                 : states is in       
    // docsIfCmtsCmStatusValue. initRangingRcvd(3)        : modem sent initial
    // ranging. initDhcpReqRcvd(4)        : dhcp request received.
    // onlineNetAccessDisabled(5): modem registered, but network                  
    // access for the CM is disabled. onlineKekAssigned(6)      : modem
    // registered, BPI enabled                             and KEK assigned.
    // onlineTekAssigned(7)      : modem registered, BPI enabled                  
    // and TEK assigned. rejectBadMic(8)           : modem did attempt to register
    // but                             registration was refused due to            
    // bad mic. rejectBadCos(9)           : modem did attempt to register but     
    // registration was refused due to                             bad COS.
    // kekRejected(10)           : KEK modem key assignment rejected.
    // tekRejected(11)           : TEK modem key assignment rejected. online(12)  
    // : modem registered, enabled for data. initTftpPacketRcvd(13)    : tftp
    // packet received and option                             file transfer
    // started.  initTodRquestRcvd(14)     : Time of the Day (TOD) request        
    // received. reset(15)                 : modem is resetting.
    // rangingInProgress(16)     : initial ranging is in progress. --           
    // rangingCompleted(17)      : initial ranging is completed. (deprecated)
    // dhcpGotIpAddr(18)         : modem has got an IP address                    
    // from the DHCP server. rejStaleConfig(19)        : modem did attempt to
    // register                             but registration was refused          
    // due to stale Config. rejIpSpoof(20)            : modem did attempt to
    // register but                             registration was refused due to IP
    // Spoof. rejClassFail(21)          : modem did attempt to register but       
    // registration was refused due to                              Class failure.
    // rejRegNack(22)            : modem did attempt to register but              
    // no acknowledgement was                              received.
    // bpiKekExpired(23)         : KEK modem key assignment expired.
    // bpiTekExpired(24)         : TEK modem key assignment expired. shutdown(25) 
    // : modem is in shutdown state. channelChgInitRangingRcvd(26)  : modem sent
    // initial ranging                                   during channel change.
    // channelChgRangingInProgress(27) : initial ranging is in                    
    // progress during channel                                   change.  This
    // cdxCmtsCmStatusValue could return initRangingRcvd(3) or
    // rangingInProgress(16) when the docsIfCmtsCmStatusValue is ranging(2).  This
    // cdxCmtsCmStatusValue will return others(2) when the docsIfCmtsCmStatusValue
    // states is either rangingAborted(3), rangingComplete(4), and ipComplete(5). 
    // This cdxCmtsCmStatusValue could return online(12), or
    // onlineNetAccessDisabled(5) for CM with BPI disabled, or
    // onlineNetAccessDisabled(5) or onlineTekAssigned(7) for CM with BPI enabled,
    // when the docsIfCmtsCmStatusValue is registrationComplete(6).  This
    // cdxCmtsCmStatusValue could return either rejectBadMic(8), rejectBadCos(9)
    // rejStaleConfig(19) or rejRegNack(22) when the docsIfCmtsCmStatusValue is
    // accessDenied(7) for possible reasons of cable modem registration abort. 
    // This cdxCmtsCmStatusValue could return either onlineKekAssigned(6),
    // kekRejected(10), tekRejected(11), or online(12) for the CM with BPI enabled
    // when the docsIfCmtsCmStatusValue is registeredBPIInitializing(9).    FOR
    // DOCSIS 1.0 -------------- The ranging, rangingAborted, rangingComplete, and
    // ipComplete  states in docsIfCmtsCmStatusValue is others in this object
    // since this object is extension of docsIfCmtsCmStatusValue.   The
    // registrationComplete state in docsIfCmtsCmStatusValue  could be online,
    // onlineNetAccessDisabled, onlineKekAssigned, or  onlineTekAssigned in this
    // object.    The accessDenied state in docsIfCmtsCmStatusValue could be 
    // rejectBadMic, rejectBadCos in this object for the possible reasons of cable
    // modem registration abort.  The states 15 to 25 are not applicable.  FOR
    // DOCSIS 1.1 --------------               The registrationComplete state in
    // docsIfCmtsCmStatusValue  could be online, onlineNetAccessDisabled, 
    // onlineBpiKekAssigned,or onlineBpiTekAssigned in this  object.    The
    // accessDenied state in docsIfCmtsCmStatusValue could be  rejMicFail,
    // rejStaleConfig, rejIpSpoof, rejClassFail,  rejRegNack in this object for
    // the possible reasons of cable modem registration abort.  The CMTS only
    // reports states it is able to detect. States Online(2) and  rejectBadCos(9)
    // are not applicable for  DOCSIS1.1 modems. . The type is
    // CdxCmtsCmStatusValue.
    CdxCmtsCmStatusValue interface{}

    // The number of times that the modem changes the connectivity  state from
    // 'offline' to 'online' over the time period from  the modem's first ranging
    // message received by CMTS until  now.  The modem is considered as 'online'
    // when the value for  cdxCmtsCmStatusValue is any of the values: online(5), 
    // onlineNetAccessDisabled(6), onlineKekAssigned(7), and 
    // onlineTekAssigned(8), and the modem is considered as 'offline' for other
    // values for cdxCmtsCmStatusValue.  . The type is interface{} with range:
    // 0..4294967295.
    CdxIfCmtsCmStatusOnlineTimes interface{}

    // The percentage of time that the modem stays 'online' over  the time period
    // from the modem's first ranging message  received by CMTS until now.   The
    // value for this object is 100 times bigger than the real  percentage value.
    // For example, 32.15% will be value 3215.  The modem is considered as
    // 'online' when the value for  cdxCmtsCmStatusValue is any of the values:
    // online(5),  onlineNetAccessDisabled(6), onlineKekAssigned(7), and 
    // onlineTekAssigned(8), and the modem is considered as  'offline' for other
    // values for cdxCmtsCmStatusValue.  . The type is interface{} with range:
    // 0..10000.
    CdxIfCmtsCmStatusPercentOnline interface{}

    // The minimum period of time the modem stayed 'online' over the time period
    // from the modem's first ranging message  received by CMTS until now.  The
    // modem is considered as 'online' when the value for  cdxCmtsCmStatusValue is
    // any of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusMinOnlineTime interface{}

    // The average period of time the modem stayed 'online' over the time period
    // from the modem's first ranging message  received by CMTS until now.  The
    // modem is considered as 'online' when the value for  cdxCmtsCmStatusValue is
    // any of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusAvgOnlineTime interface{}

    // The maximum period of time the modem stayed 'online' over the time period
    // from the modem's first ranging message  received by CMTS until now.  The
    // modem is considered as 'online' when the value for  cdxCmtsCmStatusValue is
    // any of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusMaxOnlineTime interface{}

    // The minimum period of time modem stayed 'offline' over the time period from
    // the modem's first ranging message  received by CMTS until now.  The modem
    // is considered as 'online' when the value for  cdxCmtsCmStatusValue is any
    // of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusMinOfflineTime interface{}

    // The average period of time the modem stayed 'offline' over the time period
    // from the modem's first ranging message  received by CMTS until now.  The
    // modem is considered as 'online' when the value for  cdxCmtsCmStatusValue is
    // any of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusAvgOfflineTime interface{}

    // The maximum period of time the modem stayed 'offline' over the time period
    // from the modem's first ranging message  received by CMTS until now.  The
    // modem is considered as 'online' when the value for  cdxCmtsCmStatusValue is
    // any of the values: online(5),  onlineNetAccessDisabled(6),
    // onlineKekAssigned(7), and  onlineTekAssigned(8), and the modem is
    // considered as  'offline' for other values for cdxCmtsCmStatusValue.  . The
    // type is interface{} with range: 0..2147483647.
    CdxIfCmtsCmStatusMaxOfflineTime interface{}

    // The number of active dynamic SIDs on this modem. Prior to getting the
    // assigned the Service Flow IDs(SFID) the CM must must complete a number of
    // protocol  transactions. The CMTS assigns a temporary Service ID (SID) to
    // complete these steps. The type is interface{} with range: 0..16383.
    CdxIfCmtsCmStatusDynSidCount interface{}

    // This object indicates additional attributes regarding the CM. 1. noisyPlant
    // indicates that the CM connection is noisy. 2. modemPowerMaxOut indicates
    // that the modem has reached its maximum power level.  A bit of of this
    // object is set to 1 if the condition indicated by the bit is satisfied. 
    // Note that BITS are encoded most significant bit first. . The type is
    // map[string]bool.
    CdxIfCmtsCmStatusAddlInfo interface{}

    // The number of times that the modem changes the connectivity state from
    // 'offline' to 'online' over the time period from the modem's first ranging
    // message received by CMTS until now.  The modem is considered as 'online'
    // when the value for cdxCmtsCmStatusValue is any of the values: online(5),
    // onlineNetAccessDisabled(6), onlineKekAssigned(7), and onlineTekAssigned(8),
    // and the modem is considered as 'offline' for other values for
    // cdxCmtsCmStatusValue.  The value of this object is reset to 0 if the value
    // in cdxIfCmtsCmStatusLastResetTime is changed. . The type is interface{}
    // with range: 0..4294967295.
    CdxIfCmtsCmStatusOnlineTimesNum interface{}

    // The last cable modem connectivity statistics reset time. If the value of 
    // this object is '0', then the cable modem connectivity statistics had not
    // been reset. The type is interface{} with range: 0..4294967295.
    CdxIfCmtsCmStatusLastResetTime interface{}
}

func (cdxCmtsCmStatusExtEntry *CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmStatusExtEntry.EntityData.YFilter = cdxCmtsCmStatusExtEntry.YFilter
    cdxCmtsCmStatusExtEntry.EntityData.YangName = "cdxCmtsCmStatusExtEntry"
    cdxCmtsCmStatusExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmStatusExtEntry.EntityData.ParentYangName = "cdxCmtsCmStatusExtTable"
    cdxCmtsCmStatusExtEntry.EntityData.SegmentPath = "cdxCmtsCmStatusExtEntry" + types.AddKeyToken(cdxCmtsCmStatusExtEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    cdxCmtsCmStatusExtEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsCmStatusExtTable/" + cdxCmtsCmStatusExtEntry.EntityData.SegmentPath
    cdxCmtsCmStatusExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmStatusExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmStatusExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmStatusExtEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmStatusExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", cdxCmtsCmStatusExtEntry.DocsIfCmtsCmStatusIndex})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxCmtsCmStatusValue", types.YLeaf{"CdxCmtsCmStatusValue", cdxCmtsCmStatusExtEntry.CdxCmtsCmStatusValue})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusOnlineTimes", types.YLeaf{"CdxIfCmtsCmStatusOnlineTimes", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusOnlineTimes})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusPercentOnline", types.YLeaf{"CdxIfCmtsCmStatusPercentOnline", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusPercentOnline})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusMinOnlineTime", types.YLeaf{"CdxIfCmtsCmStatusMinOnlineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusMinOnlineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusAvgOnlineTime", types.YLeaf{"CdxIfCmtsCmStatusAvgOnlineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusAvgOnlineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusMaxOnlineTime", types.YLeaf{"CdxIfCmtsCmStatusMaxOnlineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusMaxOnlineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusMinOfflineTime", types.YLeaf{"CdxIfCmtsCmStatusMinOfflineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusMinOfflineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusAvgOfflineTime", types.YLeaf{"CdxIfCmtsCmStatusAvgOfflineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusAvgOfflineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusMaxOfflineTime", types.YLeaf{"CdxIfCmtsCmStatusMaxOfflineTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusMaxOfflineTime})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusDynSidCount", types.YLeaf{"CdxIfCmtsCmStatusDynSidCount", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusDynSidCount})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusAddlInfo", types.YLeaf{"CdxIfCmtsCmStatusAddlInfo", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusAddlInfo})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusOnlineTimesNum", types.YLeaf{"CdxIfCmtsCmStatusOnlineTimesNum", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusOnlineTimesNum})
    cdxCmtsCmStatusExtEntry.EntityData.Leafs.Append("cdxIfCmtsCmStatusLastResetTime", types.YLeaf{"CdxIfCmtsCmStatusLastResetTime", cdxCmtsCmStatusExtEntry.CdxIfCmtsCmStatusLastResetTime})

    cdxCmtsCmStatusExtEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(cdxCmtsCmStatusExtEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue represents DOCSIS1.1 modems. 
type CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue string

const (
    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_offline CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "offline"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_others CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "others"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_initRangingRcvd CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "initRangingRcvd"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_initDhcpReqRcvd CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "initDhcpReqRcvd"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_onlineNetAccessDisabled CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "onlineNetAccessDisabled"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_onlineKekAssigned CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "onlineKekAssigned"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_onlineTekAssigned CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "onlineTekAssigned"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejectBadMic CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejectBadMic"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejectBadCos CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejectBadCos"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_kekRejected CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "kekRejected"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_tekRejected CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "tekRejected"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_online CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "online"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_initTftpPacketRcvd CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "initTftpPacketRcvd"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_initTodRequestRcvd CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "initTodRequestRcvd"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_reset CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "reset"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rangingInProgress CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rangingInProgress"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rangingCompleted CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rangingCompleted"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_dhcpGotIpAddr CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "dhcpGotIpAddr"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejStaleConfig CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejStaleConfig"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejIpSpoof CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejIpSpoof"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejClassFail CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejClassFail"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_rejRegNack CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "rejRegNack"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_bpiKekExpired CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "bpiKekExpired"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_bpiTekExpired CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "bpiTekExpired"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_shutdown CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "shutdown"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_channelChgInitRangingRcvd CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "channelChgInitRangingRcvd"

    CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue_channelChgRangingInProgress CISCODOCSEXTMIB_CdxCmtsCmStatusExtTable_CdxCmtsCmStatusExtEntry_CdxCmtsCmStatusValue = "channelChgRangingInProgress"
)

// CISCODOCSEXTMIB_CdxCmtsMacExtTable
// This table contains the additions attributes of a CMTS MAC
// interface that provided by docsIfCmtsMacTable. 
type CISCODOCSEXTMIB_CdxCmtsMacExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional objects for docsIfCmtsMacTable entry including the cable modem
    // notification enable/disable and the interval  of cable modem notification
    // sent by the CMTS for a cable modem that the Mac interface supports. An
    // entry in this table exists  for each ifEntry with an ifType of
    // docsCableMaclayer(127). Additional objects added to determine the number of
    // active/registered/total cable modems on this cable mac  interface since
    // boot. Also, it contains the object to set the Dynamic Message Integrity
    // Check (DMIC) which users  can configure how cable modems are handled if CMs
    // fail  the DMIC. The type is slice of
    // CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry.
    CdxCmtsMacExtEntry []*CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry
}

func (cdxCmtsMacExtTable *CISCODOCSEXTMIB_CdxCmtsMacExtTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsMacExtTable.EntityData.YFilter = cdxCmtsMacExtTable.YFilter
    cdxCmtsMacExtTable.EntityData.YangName = "cdxCmtsMacExtTable"
    cdxCmtsMacExtTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsMacExtTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsMacExtTable.EntityData.SegmentPath = "cdxCmtsMacExtTable"
    cdxCmtsMacExtTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsMacExtTable.EntityData.SegmentPath
    cdxCmtsMacExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsMacExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsMacExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsMacExtTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsMacExtTable.EntityData.Children.Append("cdxCmtsMacExtEntry", types.YChild{"CdxCmtsMacExtEntry", nil})
    for i := range cdxCmtsMacExtTable.CdxCmtsMacExtEntry {
        cdxCmtsMacExtTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsMacExtTable.CdxCmtsMacExtEntry[i]), types.YChild{"CdxCmtsMacExtEntry", cdxCmtsMacExtTable.CdxCmtsMacExtEntry[i]})
    }
    cdxCmtsMacExtTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsMacExtTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsMacExtTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry
// Additional objects for docsIfCmtsMacTable entry including
// the cable modem notification enable/disable and the interval 
// of cable modem notification sent by the CMTS for a cable modem
// that the Mac interface supports. An entry in this table exists 
// for each ifEntry with an ifType of docsCableMaclayer(127).
// Additional objects added to determine the number of 
// active/registered/total cable modems on this cable mac 
// interface since boot. Also, it contains the object to set
// the Dynamic Message Integrity Check (DMIC) which users 
// can configure how cable modems are handled if CMs fail 
// the DMIC.
type CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // An indication of whether the cdxCmtsCmOnOffNotification  is enabled. The
    // default value is false(2). . The type is bool.
    CdxCmtsCmOnOffTrapEnable interface{}

    // The interval for cdxCmtsCmOnOffNotification sent by CMTS for one
    // online/offline state change if cdxCmtsCmOnOffTrapEnable  is true.   If
    // there are more than one state changes to online/offline  for a cable modem
    // during this interval, only one  cdxCmtsCmOnOffNotification is sent by CMTS
    // for the first  state change to online and one cdxCmtsCmOnOffNotification 
    // for the first state changing to offline if  cdxCmtsCmOnOffTrapEnable is
    // true.  This is to avoid too many notifications sent for a cable  modem
    // online/offline state changes during a short period of time.   If the value
    // is 0, then cdxCmtsCmOnOffNotification will be  sent for every state changes
    // to online/offline for a cable  modem if cdxCmtsCmOnOffTrapEnable is true.  
    // If cdxCmtsCmOnOffTrapEnable value changes from true to false  or from false
    // to true, this value will remain no change as  before.   The default value
    // is 600 seconds. . The type is interface{} with range: 0..86400. Units are
    // seconds.
    CdxCmtsCmOnOffTrapInterval interface{}

    // The default maximum number of permitted CPEs per modem  in this cable
    // interface. A modem can override this  value by setting the object
    // cdxCmtsCmMaxCpeNumber in the cdxCmtsCmTable.    The value -1 means the
    // default value of maximum hosts  per modem in this cable interface is not
    // specified.  The value 0 means no maximum limit.  Setting the value will not
    // affect the already connected CPEs to the modems in this cable interface. .
    // The type is interface{} with range: -1..255.
    CdxCmtsCmDefaultMaxCpes interface{}

    // The total count of cable modems on this cable mac interface since boot. The
    // type is interface{} with range: 0..2147483647.
    CdxCmtsCmTotal interface{}

    // The count of cable modems that are active. Active cable  modems are
    // recognized by the cdxCmtsCmStatusValue  other than offline(1). . The type
    // is interface{} with range: 0..2147483647.
    CdxCmtsCmActive interface{}

    // The count of cable modems that are registered and online  on this cable mac
    // interface. Registered cable modems are  those with one of the following
    // values.  registrationComplete(6) of docsIfCmtsCmStatusValue OR  either of
    // online(12), kekRejected(10),  onlineKekAssigned(6),tekRejected(11),
    // onlineTekAssigned(7) of cdxCmtsCmStatusValue. The type is interface{} with
    // range: 0..2147483647.
    CdxCmtsCmRegistered interface{}

    // The Dynamic Shared Secret feature can operate in three  different modes,
    // depending on what action should be taken  for cable modems that fail the
    // CMTS MIC verification check: notConfigured(1): It indicates that the DMIC
    // is not                    configured for this cable interface. mark(2): By
    // default, the Dynamic Shared Secret feature           is enabled on all
    // cable interfaces using the           mark option. In this mode, the CMTS
    // allows           cable modems to come online even if they fail          
    // the CMTS MIC validity check. However, for          this modem
    // cdxCmtsCmStatusDMICMode will          be labeled as marked. lock(3): When
    // the lock option is used, the CMTS assigns           a restrictive QoS
    // configuration to CMs that           fail the MIC validity check twice in a
    // row. A           particular QoS profile to be used for locked          
    // cable modems can be specified by setting           cdxCmtsCmDMICLockQos.   
    // If a customer resets their CM, the CM will           reregister but still
    // uses the restricted QoS           profile. A locked CM continues with the  
    // restricted QoS profile until it goes offline           and remains offline
    // for at least 24 hours, at           which point it is allowed to reregister
    // with a           valid DOCSIS configuration file. A system          
    // operator can manually clear the lock on a CM by           setting
    // cdxCmtsCmStatusDMICUnLock object. reject(4):  In the reject mode, the CMTS
    // refuses to allow              CMs to come online if they fail the CMTS MIC 
    // validity check. The type is CdxCmtsCmDMICMode.
    CdxCmtsCmDMICMode interface{}

    // If cdxCmtsCmDMICMode is set to lockingMode(3), this object would contain
    // the restrictive QoS profile number as  indicated by docsIfQosProfIndex if
    // set and it will  have 0 if not applicable or not defined. In case,
    // cdxCmtsCmDMICMode is set to lockingMode(3) and this object is not defined
    // then the CMTS defaults to special QoS profile that limits the downstream
    // and upstream  service flows to a maximum rate of 10 kbps. However, for this
    // to happen the modems should have the  permission to create QoS profile. The
    // type is interface{} with range: 0..16383.
    CdxCmtsCmDMICLockQos interface{}
}

func (cdxCmtsMacExtEntry *CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsMacExtEntry.EntityData.YFilter = cdxCmtsMacExtEntry.YFilter
    cdxCmtsMacExtEntry.EntityData.YangName = "cdxCmtsMacExtEntry"
    cdxCmtsMacExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsMacExtEntry.EntityData.ParentYangName = "cdxCmtsMacExtTable"
    cdxCmtsMacExtEntry.EntityData.SegmentPath = "cdxCmtsMacExtEntry" + types.AddKeyToken(cdxCmtsMacExtEntry.IfIndex, "ifIndex")
    cdxCmtsMacExtEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsMacExtTable/" + cdxCmtsMacExtEntry.EntityData.SegmentPath
    cdxCmtsMacExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsMacExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsMacExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsMacExtEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsMacExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxCmtsMacExtEntry.IfIndex})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmOnOffTrapEnable", types.YLeaf{"CdxCmtsCmOnOffTrapEnable", cdxCmtsMacExtEntry.CdxCmtsCmOnOffTrapEnable})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmOnOffTrapInterval", types.YLeaf{"CdxCmtsCmOnOffTrapInterval", cdxCmtsMacExtEntry.CdxCmtsCmOnOffTrapInterval})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmDefaultMaxCpes", types.YLeaf{"CdxCmtsCmDefaultMaxCpes", cdxCmtsMacExtEntry.CdxCmtsCmDefaultMaxCpes})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmTotal", types.YLeaf{"CdxCmtsCmTotal", cdxCmtsMacExtEntry.CdxCmtsCmTotal})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmActive", types.YLeaf{"CdxCmtsCmActive", cdxCmtsMacExtEntry.CdxCmtsCmActive})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmRegistered", types.YLeaf{"CdxCmtsCmRegistered", cdxCmtsMacExtEntry.CdxCmtsCmRegistered})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmDMICMode", types.YLeaf{"CdxCmtsCmDMICMode", cdxCmtsMacExtEntry.CdxCmtsCmDMICMode})
    cdxCmtsMacExtEntry.EntityData.Leafs.Append("cdxCmtsCmDMICLockQos", types.YLeaf{"CdxCmtsCmDMICLockQos", cdxCmtsMacExtEntry.CdxCmtsCmDMICLockQos})

    cdxCmtsMacExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxCmtsMacExtEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode represents             validity check.
type CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode string

const (
    CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode_notConfigured CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode = "notConfigured"

    CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode_mark CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode = "mark"

    CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode_lock CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode = "lock"

    CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode_reject CISCODOCSEXTMIB_CdxCmtsMacExtTable_CdxCmtsMacExtEntry_CdxCmtsCmDMICMode = "reject"
)

// CISCODOCSEXTMIB_CdxCmtsCmChOverTable
// A table of CMTS operation entries to instruct cable modems
// to move to a new downstream and/or upstream channel. 
// 
// An entry in this table is an operation that has been 
// initiated from CMTS to generates downstream frequency and/or 
// upstream channel override fields in the RNG-RSP message sent 
// to a cable modem.  A RNG-RSP message is sent to a cable 
// modem during initial maintenance opportunity. 
// 
// This operation causes the uBR to place an entry for the cable 
// modem specified into the override request queue.  The link is 
// then broken by deleting the modem from its polling list.  When 
// the modem attempts initial ranging, the override request 
// causes downstream frequency and upstream channel override 
// fields to be inserted into the RNG-RSP message.  
type CISCODOCSEXTMIB_CdxCmtsCmChOverTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A CMTS operation entry to instruct a cable modem to move to a new
    // downstream and/or upstream channel.  A CMTS operator can use this to
    // initiate an operation in CMTS to instruct a cable modem to move to a new
    // downstream, or upstream channel or both.   Each entry consists of the mac
    // address of the cable modem to be moved, a new downstream frequency, a new
    // upstream channel  id etc..  More than one entries could have for a cable
    // modem, so there is a time stamp for each entry to show the time when this
    // operation is initiated.   A management station wishing to create an entry
    // should first generate a pseudo-random serial number to be used as the index
    // to this sparse table.  The station should then create the associated
    // instance of the row status object. It must also, either in the same or in
    // successive PDUs, create the associated instance of the command and
    // parameter objects. It should also modify the default values for any of the
    // parameter objects if the defaults are not appropriate.  Once the
    // appropriate instances of all the command objects have been created, either
    // by an explicit SNMP set request or by default, the row status should be set
    // to active to initiate the operation. Note that this entire procedure may be
    // initiated via a single set request which specifies a row status  of
    // createAndGo as well as specifies valid values for the non-defaulted
    // parameter objects.  Once an operation has been activated, it cannot be
    // stopped. That is, it will run until either the CMTS has generated 
    // downstream frequency and/or upstream channel override fields  in the
    // RNG-RSP message sent to a cable modem or time out.  In either case, the
    // operation is completed.  Once the operation is completed, the real result
    // of the  operation to the cable modem cannot be known from this table. The
    // result of the cable modem's downstream frequency and the  upstream channel
    // id can be checked from other MIB tables.   For example,
    // docsIfCmtsServiceTable from DOCS-IF-MIB can be  used to check whether the
    // cable modem's downstream frequency and upstream channel id are changed. 
    // Please note that even the CMTS has generated downstream frequency and/or
    // upstream  channel override fields in the RNG-RSP message sent to a  cable
    // modems, if the cable modem cannot lock the instructed  downstream frequency
    // or no upstream channel id could be used,  it may reconnect back to the
    // original downstream frequency and upstream channel id.   Once the operation
    // completes, the management station should retrieve the values of the
    // cdxCmtsCmChOverState  objects of interest, and should then delete the
    // entry.   In order to prevent old entries from clogging the table,  entries
    // will be aged out, but an entry will never be deleted  within 15 minutes of
    // completing. . The type is slice of
    // CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry.
    CdxCmtsCmChOverEntry []*CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry
}

func (cdxCmtsCmChOverTable *CISCODOCSEXTMIB_CdxCmtsCmChOverTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmChOverTable.EntityData.YFilter = cdxCmtsCmChOverTable.YFilter
    cdxCmtsCmChOverTable.EntityData.YangName = "cdxCmtsCmChOverTable"
    cdxCmtsCmChOverTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmChOverTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsCmChOverTable.EntityData.SegmentPath = "cdxCmtsCmChOverTable"
    cdxCmtsCmChOverTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsCmChOverTable.EntityData.SegmentPath
    cdxCmtsCmChOverTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmChOverTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmChOverTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmChOverTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmChOverTable.EntityData.Children.Append("cdxCmtsCmChOverEntry", types.YChild{"CdxCmtsCmChOverEntry", nil})
    for i := range cdxCmtsCmChOverTable.CdxCmtsCmChOverEntry {
        cdxCmtsCmChOverTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsCmChOverTable.CdxCmtsCmChOverEntry[i]), types.YChild{"CdxCmtsCmChOverEntry", cdxCmtsCmChOverTable.CdxCmtsCmChOverEntry[i]})
    }
    cdxCmtsCmChOverTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsCmChOverTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsCmChOverTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry
// A CMTS operation entry to instruct a cable modem to move to
// a new downstream and/or upstream channel.
// 
// A CMTS operator can use this to initiate an operation
// in CMTS to instruct a cable modem to move to a new
// downstream, or upstream channel or both. 
// 
// Each entry consists of the mac address of the cable modem
// to be moved, a new downstream frequency, a new upstream channel 
// id etc..  More than one entries could have for a cable modem,
// so there is a time stamp for each entry to show the time
// when this operation is initiated. 
// 
// A management station wishing to create an entry should
// first generate a pseudo-random serial number to be used
// as the index to this sparse table.  The station should
// then create the associated instance of the row status
// object. It must also, either in the same or in successive
// PDUs, create the associated instance of the command and
// parameter objects. It should also modify the default values
// for any of the parameter objects if the defaults are not
// appropriate.
// 
// Once the appropriate instances of all the command
// objects have been created, either by an explicit SNMP
// set request or by default, the row status should be set
// to active to initiate the operation. Note that this entire
// procedure may be initiated via a single set request which
// specifies a row status  of createAndGo as well as specifies
// valid values for the non-defaulted parameter objects.
// 
// Once an operation has been activated, it cannot be stopped.
// That is, it will run until either the CMTS has generated 
// downstream frequency and/or upstream channel override fields 
// in the RNG-RSP message sent to a cable modem or time out. 
// In either case, the operation is completed.
// 
// Once the operation is completed, the real result of the 
// operation to the cable modem cannot be known from this table.
// The result of the cable modem's downstream frequency and the 
// upstream channel id can be checked from other MIB tables.  
// For example, docsIfCmtsServiceTable from DOCS-IF-MIB can be 
// used to check whether the cable modem's downstream frequency
// and upstream channel id are changed.  Please note that even
// the CMTS has generated downstream frequency and/or upstream 
// channel override fields in the RNG-RSP message sent to a 
// cable modems, if the cable modem cannot lock the instructed 
// downstream frequency or no upstream channel id could be used, 
// it may reconnect back to the original downstream frequency
// and upstream channel id. 
// 
// Once the operation completes, the management station should
// retrieve the values of the cdxCmtsCmChOverState 
// objects of interest, and should then delete the entry.  
// In order to prevent old entries from clogging the table, 
// entries will be aged out, but an entry will never be deleted 
// within 15 minutes of completing. 
type CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // table. A management station wishing to initiate a channel override
    // operation should use a pseudo-random  value for this object when creating
    // or modifying an  instance of a cdxCmtsCmChOverEntry.  . The type is
    // interface{} with range: 1..2147483647.
    CdxCmtsCmChOverSerialNumber interface{}

    // The mac address of the cable modem that the CMTS instructs to move to a new
    // downstream and/or upstream channel.    This column must be set to a valid
    // Mac address currently in the CMTS in order for this entry's row status to
    // be set to active successfully. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCmtsCmChOverMacAddress interface{}

    // The new downstream frequency which the cable modem is  instructed to move
    // to.  The value 0 is to ask the CMTS not to override the downstream
    // frequency. . The type is interface{} with range: 0..1000000000. Units are
    // hertz.
    CdxCmtsCmChOverDownFrequency interface{}

    // The new channel Id which the cable modem is instructed to  move to.  The
    // value -1 is to ask the CMTS not to override the upstream channel Id. . The
    // type is interface{} with range: -1..255.
    CdxCmtsCmChOverUpChannelId interface{}

    // Specifies whether or not a cdxCmtsCmChOverNotification  should be issued on
    // completion of the operation.  If such a  notification is desired, it is the
    // responsibility of the  management entity to ensure that the SNMP
    // administrative model  is configured in such a way as to allow the
    // notification to be  delivered. . The type is bool.
    CdxCmtsCmChOverTrapOnCompletion interface{}

    // The value of sysUpTime at which the operation was initiated.   Since it is
    // possible to have more than one entry in this  table for a cable modem, this
    // object can help to distinguish  the entries for the same cable modem. . The
    // type is interface{} with range: 0..4294967295.
    CdxCmtsCmChOverOpInitiatedTime interface{}

    // The status of the specified channel override operation. The enumerations
    // are:   messageSent(1): the CMTS has sent a RNG-RSP message               
    // with channel override to the cable modem.    commandNotActive(2): the
    // command is not in active mode                        due to this entry's
    // row status is not                        in active yet.   noOpNeed(3): The
    // downstream frequency and the upstream                 channel Id in this
    // entry are the same as                 original ones when this entry's row
    // status                is set to active, so CMTS does not need to           
    // do any operation.     modemNotFound(4): The modem is not found in the CMTS 
    // at the time when the command becomes                     active.  
    // waitToSendMessage(5): specified the operation is active                    
    // and CMTS is waiting to send                         a RNG-RSP message with
    // channel                          override to the cable modem.   timeOut(6):
    // specified the operation is timed out.               That is, the CMTS
    // cannot send a RNG-RSP message                with channel override to the
    // cable modem within                the time specified in the object of      
    // cdxCmtsCmChOverTimeExpiration.                The possible reason is that
    // the cable modem               does not repeat the initial ranging.      The
    // possible state change diagram is as below:     [commandNotActive ->]
    // waitToSendMessage ->         messageSent or timeOut.     [commandNotActive
    // ->] noOpNeeded or modemNotFound. . The type is CdxCmtsCmChOverState.
    CdxCmtsCmChOverState interface{}

    // The status of this table entry.    This value for cdxCmtsCmChOverMacAddress
    // must be valid Mac  address currently in the CMTS in order for the row 
    // status to be set to active successfully.      Once the row status becomes
    // active and state becomes  waitToSendMessage, the entry cannot not be
    // changed except  to delete the entry by setting the row status to destroy(6)
    // and since the operation cannot be stopped, the destroy(6)  will just cause
    // the SNMP agent to hide the entry from  application and the SNMP agent will
    // delete the entry  right after the operation is completed. . The type is
    // RowStatus.
    CdxCmtsCmChOverRowStatus interface{}
}

func (cdxCmtsCmChOverEntry *CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmChOverEntry.EntityData.YFilter = cdxCmtsCmChOverEntry.YFilter
    cdxCmtsCmChOverEntry.EntityData.YangName = "cdxCmtsCmChOverEntry"
    cdxCmtsCmChOverEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmChOverEntry.EntityData.ParentYangName = "cdxCmtsCmChOverTable"
    cdxCmtsCmChOverEntry.EntityData.SegmentPath = "cdxCmtsCmChOverEntry" + types.AddKeyToken(cdxCmtsCmChOverEntry.CdxCmtsCmChOverSerialNumber, "cdxCmtsCmChOverSerialNumber")
    cdxCmtsCmChOverEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsCmChOverTable/" + cdxCmtsCmChOverEntry.EntityData.SegmentPath
    cdxCmtsCmChOverEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmChOverEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmChOverEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmChOverEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmChOverEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverSerialNumber", types.YLeaf{"CdxCmtsCmChOverSerialNumber", cdxCmtsCmChOverEntry.CdxCmtsCmChOverSerialNumber})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverMacAddress", types.YLeaf{"CdxCmtsCmChOverMacAddress", cdxCmtsCmChOverEntry.CdxCmtsCmChOverMacAddress})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverDownFrequency", types.YLeaf{"CdxCmtsCmChOverDownFrequency", cdxCmtsCmChOverEntry.CdxCmtsCmChOverDownFrequency})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverUpChannelId", types.YLeaf{"CdxCmtsCmChOverUpChannelId", cdxCmtsCmChOverEntry.CdxCmtsCmChOverUpChannelId})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverTrapOnCompletion", types.YLeaf{"CdxCmtsCmChOverTrapOnCompletion", cdxCmtsCmChOverEntry.CdxCmtsCmChOverTrapOnCompletion})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverOpInitiatedTime", types.YLeaf{"CdxCmtsCmChOverOpInitiatedTime", cdxCmtsCmChOverEntry.CdxCmtsCmChOverOpInitiatedTime})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverState", types.YLeaf{"CdxCmtsCmChOverState", cdxCmtsCmChOverEntry.CdxCmtsCmChOverState})
    cdxCmtsCmChOverEntry.EntityData.Leafs.Append("cdxCmtsCmChOverRowStatus", types.YLeaf{"CdxCmtsCmChOverRowStatus", cdxCmtsCmChOverEntry.CdxCmtsCmChOverRowStatus})

    cdxCmtsCmChOverEntry.EntityData.YListKeys = []string {"CdxCmtsCmChOverSerialNumber"}

    return &(cdxCmtsCmChOverEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState represents    [commandNotActive ->] noOpNeeded or modemNotFound. 
type CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState string

const (
    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_messageSent CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "messageSent"

    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_commandNotActive CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "commandNotActive"

    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_noOpNeeded CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "noOpNeeded"

    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_modemNotFound CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "modemNotFound"

    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_waitToSendMessage CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "waitToSendMessage"

    CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState_timeOut CISCODOCSEXTMIB_CdxCmtsCmChOverTable_CdxCmtsCmChOverEntry_CdxCmtsCmChOverState = "timeOut"
)

// CISCODOCSEXTMIB_CdxCmtsCmTable
// This table contains attributes or configurable parameters 
// for cable modems from a CMTS. 
type CISCODOCSEXTMIB_CdxCmtsCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list contains a cable modem's attributes or  configurable parameters
    // from a CMTS. . The type is slice of
    // CISCODOCSEXTMIB_CdxCmtsCmTable_CdxCmtsCmEntry.
    CdxCmtsCmEntry []*CISCODOCSEXTMIB_CdxCmtsCmTable_CdxCmtsCmEntry
}

func (cdxCmtsCmTable *CISCODOCSEXTMIB_CdxCmtsCmTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmTable.EntityData.YFilter = cdxCmtsCmTable.YFilter
    cdxCmtsCmTable.EntityData.YangName = "cdxCmtsCmTable"
    cdxCmtsCmTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsCmTable.EntityData.SegmentPath = "cdxCmtsCmTable"
    cdxCmtsCmTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsCmTable.EntityData.SegmentPath
    cdxCmtsCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmTable.EntityData.Children.Append("cdxCmtsCmEntry", types.YChild{"CdxCmtsCmEntry", nil})
    for i := range cdxCmtsCmTable.CdxCmtsCmEntry {
        cdxCmtsCmTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsCmTable.CdxCmtsCmEntry[i]), types.YChild{"CdxCmtsCmEntry", cdxCmtsCmTable.CdxCmtsCmEntry[i]})
    }
    cdxCmtsCmTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsCmTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsCmTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmTable_CdxCmtsCmEntry
// The list contains a cable modem's attributes or 
// configurable parameters from a CMTS. 
type CISCODOCSEXTMIB_CdxCmtsCmTable_CdxCmtsCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // The maximum number of permitted CPEs connecting to the modem.   The value
    // -1 means to use the default value of maximum  hosts per modem in the CMTS
    // cable interface which the modem  connects to and the value is defined in 
    // cdxCmtsCmDefaultMaxCpes in the cdxCmtsMacExtTable.   The value 0 means no
    // maximum limit.  Setting the value will not affect the already connected
    // CPEs to the modem. . The type is interface{} with range: -1..255.
    CdxCmtsCmMaxCpeNumber interface{}

    // The current number of CPEs connecting to the modem.  The value 0 means no
    // hosts connecting to the modem. The type is interface{} with range: 0..255.
    CdxCmtsCmCurrCpeNumber interface{}

    // The index in docsIfQosProfileTable describing the quality of service
    // attributes associated with this particular modem's primary SID.   When
    // trying to change the value, if the new value is not  a valid index in the
    // docsIfQosProfileTable, the modem will  retain the old docsIfQosProfileTable
    // entry. If no associated  docsIfQosProfileTable entry exists for this modem,
    // this object returns a value of zero on read.  This object has meaning only
    // for DOCSIS1.0 cable modems. For cable modems in DOCSIS1.1 or above mode,
    // this object will  report 0 and cannot be changed to any other values since 
    // there is no QoS profile associated with cable modems in  DOCSIS1.1 or above
    // mode. The type is interface{} with range: 0..16383.
    CdxCmtsCmQosProfile interface{}
}

func (cdxCmtsCmEntry *CISCODOCSEXTMIB_CdxCmtsCmTable_CdxCmtsCmEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmEntry.EntityData.YFilter = cdxCmtsCmEntry.YFilter
    cdxCmtsCmEntry.EntityData.YangName = "cdxCmtsCmEntry"
    cdxCmtsCmEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmEntry.EntityData.ParentYangName = "cdxCmtsCmTable"
    cdxCmtsCmEntry.EntityData.SegmentPath = "cdxCmtsCmEntry" + types.AddKeyToken(cdxCmtsCmEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    cdxCmtsCmEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsCmTable/" + cdxCmtsCmEntry.EntityData.SegmentPath
    cdxCmtsCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsCmEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", cdxCmtsCmEntry.DocsIfCmtsCmStatusIndex})
    cdxCmtsCmEntry.EntityData.Leafs.Append("cdxCmtsCmMaxCpeNumber", types.YLeaf{"CdxCmtsCmMaxCpeNumber", cdxCmtsCmEntry.CdxCmtsCmMaxCpeNumber})
    cdxCmtsCmEntry.EntityData.Leafs.Append("cdxCmtsCmCurrCpeNumber", types.YLeaf{"CdxCmtsCmCurrCpeNumber", cdxCmtsCmEntry.CdxCmtsCmCurrCpeNumber})
    cdxCmtsCmEntry.EntityData.Leafs.Append("cdxCmtsCmQosProfile", types.YLeaf{"CdxCmtsCmQosProfile", cdxCmtsCmEntry.CdxCmtsCmQosProfile})

    cdxCmtsCmEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(cdxCmtsCmEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable
// This table contains the list of modems which failed the CMTS
// Dynamic Message Integrity Check (DMIC). The modems are 
// either
// Marked: The modems failed the DMIC check but were still 
//         allowed to come online.
// Locked: The modems failed the DMIC check and hence were 
//         allowed to come online with a restrictive QoS 
//         profile as defined in  cdxCmtsCmDMICLockQos.
// Rejected: The modems failed the DMIC check and hence
//           were not allowed to come online.
// Another objective of the objects in this table is to clear
// the lock on the modems.
type CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional DMIC objects for docsIfCmtsCmStatusTable  entry. . The type is
    // slice of CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry.
    CdxCmtsCmStatusDMICEntry []*CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry
}

func (cdxCmtsCmStatusDMICTable *CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmStatusDMICTable.EntityData.YFilter = cdxCmtsCmStatusDMICTable.YFilter
    cdxCmtsCmStatusDMICTable.EntityData.YangName = "cdxCmtsCmStatusDMICTable"
    cdxCmtsCmStatusDMICTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmStatusDMICTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsCmStatusDMICTable.EntityData.SegmentPath = "cdxCmtsCmStatusDMICTable"
    cdxCmtsCmStatusDMICTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsCmStatusDMICTable.EntityData.SegmentPath
    cdxCmtsCmStatusDMICTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmStatusDMICTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmStatusDMICTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmStatusDMICTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmStatusDMICTable.EntityData.Children.Append("cdxCmtsCmStatusDMICEntry", types.YChild{"CdxCmtsCmStatusDMICEntry", nil})
    for i := range cdxCmtsCmStatusDMICTable.CdxCmtsCmStatusDMICEntry {
        cdxCmtsCmStatusDMICTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsCmStatusDMICTable.CdxCmtsCmStatusDMICEntry[i]), types.YChild{"CdxCmtsCmStatusDMICEntry", cdxCmtsCmStatusDMICTable.CdxCmtsCmStatusDMICEntry[i]})
    }
    cdxCmtsCmStatusDMICTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsCmStatusDMICTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsCmStatusDMICTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry
// Additional DMIC objects for docsIfCmtsCmStatusTable 
// entry. 
type CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // This shows all the cable modems that are online or offline and that had
    // failed the Dynamic CMTS MIC verification check. The state is mentioned as
    // follows: mark(1): The modem was allowed to come online. lock(2): The modem
    // was allowed to come online but with            a restrictive QoS profile as
    // defined by             cdxCmtsCmDMICLockQos. reject(3): The modem was not
    // allowed to come online. The type is CdxCmtsCmStatusDMICMode.
    CdxCmtsCmStatusDMICMode interface{}

    // When set to TRUE, it forces the cable modems to  reinitialize, and the
    // cable modems must re-register with a valid DOCSIS configuration file before
    // being allowed online. Otherwise, the cable modem is locked  in its current
    // restricted QoS profile and cannot  reregister with a different profile
    // until it has  been offline for at least 24 hours. If
    // cdxCmtsCmStatusDMICUnLock is set to TRUE, and re-init succeeds, that modem
    // row is removed from the cdxCmtsCmStatusDMICTable. And if re-init again
    // fails, the row remains in that table, possibly with a new value for
    // cdxCmtsCmStatusDMICMode When polled, it will always return FALSE. The type
    // is bool.
    CdxCmtsCmStatusDMICUnLock interface{}
}

func (cdxCmtsCmStatusDMICEntry *CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsCmStatusDMICEntry.EntityData.YFilter = cdxCmtsCmStatusDMICEntry.YFilter
    cdxCmtsCmStatusDMICEntry.EntityData.YangName = "cdxCmtsCmStatusDMICEntry"
    cdxCmtsCmStatusDMICEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsCmStatusDMICEntry.EntityData.ParentYangName = "cdxCmtsCmStatusDMICTable"
    cdxCmtsCmStatusDMICEntry.EntityData.SegmentPath = "cdxCmtsCmStatusDMICEntry" + types.AddKeyToken(cdxCmtsCmStatusDMICEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    cdxCmtsCmStatusDMICEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsCmStatusDMICTable/" + cdxCmtsCmStatusDMICEntry.EntityData.SegmentPath
    cdxCmtsCmStatusDMICEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsCmStatusDMICEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsCmStatusDMICEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsCmStatusDMICEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsCmStatusDMICEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsCmStatusDMICEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", cdxCmtsCmStatusDMICEntry.DocsIfCmtsCmStatusIndex})
    cdxCmtsCmStatusDMICEntry.EntityData.Leafs.Append("cdxCmtsCmStatusDMICMode", types.YLeaf{"CdxCmtsCmStatusDMICMode", cdxCmtsCmStatusDMICEntry.CdxCmtsCmStatusDMICMode})
    cdxCmtsCmStatusDMICEntry.EntityData.Leafs.Append("cdxCmtsCmStatusDMICUnLock", types.YLeaf{"CdxCmtsCmStatusDMICUnLock", cdxCmtsCmStatusDMICEntry.CdxCmtsCmStatusDMICUnLock})

    cdxCmtsCmStatusDMICEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(cdxCmtsCmStatusDMICEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode represents reject(3): The modem was not allowed to come online.
type CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode string

const (
    CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode_mark CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode = "mark"

    CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode_lock CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode = "lock"

    CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode_reject CISCODOCSEXTMIB_CdxCmtsCmStatusDMICTable_CdxCmtsCmStatusDMICEntry_CdxCmtsCmStatusDMICMode = "reject"
)

// CISCODOCSEXTMIB_CdxCmToCpeTable
// This table contains information about CPE connects behind
// cable modem. It will return IP address and IP address type
// of each CPE connect to a CM.
// 
// It is not intended to walk the whole table. An application
// would need to query this table based on the specific indices.
// Otherwise, it will impact the CMTS performance due to the 
// huge size of this table.
// 
// The agent creates/destroys/modifies an entry whenever there
// is a CPE connect to a cable modem or disconnect from a cable
// modem.
type CISCODOCSEXTMIB_CdxCmToCpeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Represents an entry in the table. Each entry is created if there is CPE
    // connects to a cable modem.  The indices uniquely identify a CPE. It is
    // never the intent for an application to perform a SNMP Get operation against
    // a table of this nature, rather it is the intent to merely enumberate
    // mappings.   An application would determine the CPEs behind all cable modems
    // by performing a SNMP GetNext starting with the variable bindings: -
    // cdxCmToCpeInetAddressType.0 - cdxCmToCpeInetAddress.0  It will return the
    // IP address type and value tuple corresponding to the CPE with lowest IP
    // address behind the cable modem with the lowest MAC address. An application
    // can perform a SNMP GetNext operation with the following variable bindings:
    // - cdxCmToCpeInetAddressType.x.y.z - cdxCmToCpeInetAddress.x.y.z where x is
    // MAC address of cable modem, and y.z is IP address type and value tuple of
    // the reported CPE. An application can repeat this process until it has
    // traversed the entire table.  If the application only wants to know the CPEs
    // behind a given cable modem, it can perform a SNMP GetNext opertaion with
    // the following: - cdxCmToCpeInetAddressType.x - cdxCmToCpeInetAddress.x
    // where x is MAC address of cable modem. The type is slice of
    // CISCODOCSEXTMIB_CdxCmToCpeTable_CdxCmToCpeEntry.
    CdxCmToCpeEntry []*CISCODOCSEXTMIB_CdxCmToCpeTable_CdxCmToCpeEntry
}

func (cdxCmToCpeTable *CISCODOCSEXTMIB_CdxCmToCpeTable) GetEntityData() *types.CommonEntityData {
    cdxCmToCpeTable.EntityData.YFilter = cdxCmToCpeTable.YFilter
    cdxCmToCpeTable.EntityData.YangName = "cdxCmToCpeTable"
    cdxCmToCpeTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmToCpeTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmToCpeTable.EntityData.SegmentPath = "cdxCmToCpeTable"
    cdxCmToCpeTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmToCpeTable.EntityData.SegmentPath
    cdxCmToCpeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmToCpeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmToCpeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmToCpeTable.EntityData.Children = types.NewOrderedMap()
    cdxCmToCpeTable.EntityData.Children.Append("cdxCmToCpeEntry", types.YChild{"CdxCmToCpeEntry", nil})
    for i := range cdxCmToCpeTable.CdxCmToCpeEntry {
        cdxCmToCpeTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmToCpeTable.CdxCmToCpeEntry[i]), types.YChild{"CdxCmToCpeEntry", cdxCmToCpeTable.CdxCmToCpeEntry[i]})
    }
    cdxCmToCpeTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmToCpeTable.EntityData.YListKeys = []string {}

    return &(cdxCmToCpeTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmToCpeTable_CdxCmToCpeEntry
// Represents an entry in the table. Each entry is created if
// there is CPE connects to a cable modem.
// 
// The indices uniquely identify a CPE. It is never the intent
// for an application to perform a SNMP Get operation against
// a table of this nature, rather it is the intent to merely
// enumberate mappings. 
// 
// An application would determine the CPEs behind all cable
// modems by performing a SNMP GetNext starting with the
// variable bindings:
// - cdxCmToCpeInetAddressType.0
// - cdxCmToCpeInetAddress.0
// 
// It will return the IP address type and value tuple
// corresponding to the CPE with lowest IP address behind the
// cable modem with the lowest MAC address. An application can
// perform a SNMP GetNext operation with the following variable
// bindings:
// - cdxCmToCpeInetAddressType.x.y.z
// - cdxCmToCpeInetAddress.x.y.z
// where x is MAC address of cable modem, and y.z is IP address
// type and value tuple of the reported CPE.
// An application can repeat this process until it has
// traversed the entire table.
// 
// If the application only wants to know the CPEs behind a
// given cable modem, it can perform a SNMP GetNext opertaion
// with the following:
// - cdxCmToCpeInetAddressType.x
// - cdxCmToCpeInetAddress.x
// where x is MAC address of cable modem.
type CISCODOCSEXTMIB_CdxCmToCpeTable_CdxCmToCpeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The MAC address that uniquely identifies a cable
    // modem that CPEs connects to. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCmToCpeCmMacAddress interface{}

    // This attribute is a key. The type of Internet address of the
    // cdxCmToCpeInetAddress. The type is InetAddressType.
    CdxCmToCpeInetAddressType interface{}

    // This attribute is a key. This object identifies the address assigned to
    // this CPE. The type is string with length: 0..255.
    CdxCmToCpeInetAddress interface{}
}

func (cdxCmToCpeEntry *CISCODOCSEXTMIB_CdxCmToCpeTable_CdxCmToCpeEntry) GetEntityData() *types.CommonEntityData {
    cdxCmToCpeEntry.EntityData.YFilter = cdxCmToCpeEntry.YFilter
    cdxCmToCpeEntry.EntityData.YangName = "cdxCmToCpeEntry"
    cdxCmToCpeEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmToCpeEntry.EntityData.ParentYangName = "cdxCmToCpeTable"
    cdxCmToCpeEntry.EntityData.SegmentPath = "cdxCmToCpeEntry" + types.AddKeyToken(cdxCmToCpeEntry.CdxCmToCpeCmMacAddress, "cdxCmToCpeCmMacAddress") + types.AddKeyToken(cdxCmToCpeEntry.CdxCmToCpeInetAddressType, "cdxCmToCpeInetAddressType") + types.AddKeyToken(cdxCmToCpeEntry.CdxCmToCpeInetAddress, "cdxCmToCpeInetAddress")
    cdxCmToCpeEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmToCpeTable/" + cdxCmToCpeEntry.EntityData.SegmentPath
    cdxCmToCpeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmToCpeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmToCpeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmToCpeEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmToCpeEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmToCpeEntry.EntityData.Leafs.Append("cdxCmToCpeCmMacAddress", types.YLeaf{"CdxCmToCpeCmMacAddress", cdxCmToCpeEntry.CdxCmToCpeCmMacAddress})
    cdxCmToCpeEntry.EntityData.Leafs.Append("cdxCmToCpeInetAddressType", types.YLeaf{"CdxCmToCpeInetAddressType", cdxCmToCpeEntry.CdxCmToCpeInetAddressType})
    cdxCmToCpeEntry.EntityData.Leafs.Append("cdxCmToCpeInetAddress", types.YLeaf{"CdxCmToCpeInetAddress", cdxCmToCpeEntry.CdxCmToCpeInetAddress})

    cdxCmToCpeEntry.EntityData.YListKeys = []string {"CdxCmToCpeCmMacAddress", "CdxCmToCpeInetAddressType", "CdxCmToCpeInetAddress"}

    return &(cdxCmToCpeEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCpeToCmTable
// This table contains information about cable modems with CPE
// connects to.
// 
// It is not intended to walk the whole table. An application
// would need to query this table base on the specific index.
// Otherwise, it will impact the CMTS performance due to the
// huge size of this table.
// 
// The agent creates/destroys/modifies an entry whenever there
// is update for the cable modem that CPE connects to.
type CISCODOCSEXTMIB_CdxCpeToCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in cdxCpeToCmTable. Each entry contains information on the MAC
    // address, IP Address, and status index for the  cable modem with a specific
    // CPE connects to. Each entry is created if there is any cable modem with CPE
    // connects to. Entries are ordered by cdxCpeToCmCpeMacAddress. The type is
    // slice of CISCODOCSEXTMIB_CdxCpeToCmTable_CdxCpeToCmEntry.
    CdxCpeToCmEntry []*CISCODOCSEXTMIB_CdxCpeToCmTable_CdxCpeToCmEntry
}

func (cdxCpeToCmTable *CISCODOCSEXTMIB_CdxCpeToCmTable) GetEntityData() *types.CommonEntityData {
    cdxCpeToCmTable.EntityData.YFilter = cdxCpeToCmTable.YFilter
    cdxCpeToCmTable.EntityData.YangName = "cdxCpeToCmTable"
    cdxCpeToCmTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCpeToCmTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCpeToCmTable.EntityData.SegmentPath = "cdxCpeToCmTable"
    cdxCpeToCmTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCpeToCmTable.EntityData.SegmentPath
    cdxCpeToCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCpeToCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCpeToCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCpeToCmTable.EntityData.Children = types.NewOrderedMap()
    cdxCpeToCmTable.EntityData.Children.Append("cdxCpeToCmEntry", types.YChild{"CdxCpeToCmEntry", nil})
    for i := range cdxCpeToCmTable.CdxCpeToCmEntry {
        cdxCpeToCmTable.EntityData.Children.Append(types.GetSegmentPath(cdxCpeToCmTable.CdxCpeToCmEntry[i]), types.YChild{"CdxCpeToCmEntry", cdxCpeToCmTable.CdxCpeToCmEntry[i]})
    }
    cdxCpeToCmTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCpeToCmTable.EntityData.YListKeys = []string {}

    return &(cdxCpeToCmTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCpeToCmTable_CdxCpeToCmEntry
// An entry in cdxCpeToCmTable. Each entry contains information
// on the MAC address, IP Address, and status index for the 
// cable modem with a specific CPE connects to. Each entry is
// created if there is any cable modem with CPE connects to.
// Entries are ordered by cdxCpeToCmCpeMacAddress.
type CISCODOCSEXTMIB_CdxCpeToCmTable_CdxCpeToCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object identifies the MAC address of the CPE.
    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCpeToCmCpeMacAddress interface{}

    // This object identifies the MAC address of the cable modem. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCpeToCmMacAddress interface{}

    // The type of Internet address of the cdxCpeToCmInetAddress object. The type
    // is InetAddressType.
    CdxCpeToCmInetAddressType interface{}

    // This object identifies the address assigned to this cable modem. The type
    // is string with length: 0..255.
    CdxCpeToCmInetAddress interface{}

    // An entry in docsIfCmtsCmStatusTable identifying status index of the cable
    // modem which the CPE connects to. The type is interface{} with range:
    // 1..2147483647.
    CdxCpeToCmStatusIndex interface{}
}

func (cdxCpeToCmEntry *CISCODOCSEXTMIB_CdxCpeToCmTable_CdxCpeToCmEntry) GetEntityData() *types.CommonEntityData {
    cdxCpeToCmEntry.EntityData.YFilter = cdxCpeToCmEntry.YFilter
    cdxCpeToCmEntry.EntityData.YangName = "cdxCpeToCmEntry"
    cdxCpeToCmEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCpeToCmEntry.EntityData.ParentYangName = "cdxCpeToCmTable"
    cdxCpeToCmEntry.EntityData.SegmentPath = "cdxCpeToCmEntry" + types.AddKeyToken(cdxCpeToCmEntry.CdxCpeToCmCpeMacAddress, "cdxCpeToCmCpeMacAddress")
    cdxCpeToCmEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCpeToCmTable/" + cdxCpeToCmEntry.EntityData.SegmentPath
    cdxCpeToCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCpeToCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCpeToCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCpeToCmEntry.EntityData.Children = types.NewOrderedMap()
    cdxCpeToCmEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCpeToCmEntry.EntityData.Leafs.Append("cdxCpeToCmCpeMacAddress", types.YLeaf{"CdxCpeToCmCpeMacAddress", cdxCpeToCmEntry.CdxCpeToCmCpeMacAddress})
    cdxCpeToCmEntry.EntityData.Leafs.Append("cdxCpeToCmMacAddress", types.YLeaf{"CdxCpeToCmMacAddress", cdxCpeToCmEntry.CdxCpeToCmMacAddress})
    cdxCpeToCmEntry.EntityData.Leafs.Append("cdxCpeToCmInetAddressType", types.YLeaf{"CdxCpeToCmInetAddressType", cdxCpeToCmEntry.CdxCpeToCmInetAddressType})
    cdxCpeToCmEntry.EntityData.Leafs.Append("cdxCpeToCmInetAddress", types.YLeaf{"CdxCpeToCmInetAddress", cdxCpeToCmEntry.CdxCpeToCmInetAddress})
    cdxCpeToCmEntry.EntityData.Leafs.Append("cdxCpeToCmStatusIndex", types.YLeaf{"CdxCpeToCmStatusIndex", cdxCpeToCmEntry.CdxCpeToCmStatusIndex})

    cdxCpeToCmEntry.EntityData.YListKeys = []string {"CdxCpeToCmCpeMacAddress"}

    return &(cdxCpeToCmEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCpeIpPrefixTable
// The table contains a list CPE's IP Prefix management
// information.
type CISCODOCSEXTMIB_CdxCpeIpPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry contains information of CM's MAC, CPE's IP prefix type, CPE's IP
    // prefix address, CPE's IP prefix length and CPE's MAC address. An entry is
    // created if CPE is associated with a prefix. The type is slice of
    // CISCODOCSEXTMIB_CdxCpeIpPrefixTable_CdxCpeIpPrefixEntry.
    CdxCpeIpPrefixEntry []*CISCODOCSEXTMIB_CdxCpeIpPrefixTable_CdxCpeIpPrefixEntry
}

func (cdxCpeIpPrefixTable *CISCODOCSEXTMIB_CdxCpeIpPrefixTable) GetEntityData() *types.CommonEntityData {
    cdxCpeIpPrefixTable.EntityData.YFilter = cdxCpeIpPrefixTable.YFilter
    cdxCpeIpPrefixTable.EntityData.YangName = "cdxCpeIpPrefixTable"
    cdxCpeIpPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCpeIpPrefixTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCpeIpPrefixTable.EntityData.SegmentPath = "cdxCpeIpPrefixTable"
    cdxCpeIpPrefixTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCpeIpPrefixTable.EntityData.SegmentPath
    cdxCpeIpPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCpeIpPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCpeIpPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCpeIpPrefixTable.EntityData.Children = types.NewOrderedMap()
    cdxCpeIpPrefixTable.EntityData.Children.Append("cdxCpeIpPrefixEntry", types.YChild{"CdxCpeIpPrefixEntry", nil})
    for i := range cdxCpeIpPrefixTable.CdxCpeIpPrefixEntry {
        cdxCpeIpPrefixTable.EntityData.Children.Append(types.GetSegmentPath(cdxCpeIpPrefixTable.CdxCpeIpPrefixEntry[i]), types.YChild{"CdxCpeIpPrefixEntry", cdxCpeIpPrefixTable.CdxCpeIpPrefixEntry[i]})
    }
    cdxCpeIpPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCpeIpPrefixTable.EntityData.YListKeys = []string {}

    return &(cdxCpeIpPrefixTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCpeIpPrefixTable_CdxCpeIpPrefixEntry
// An entry contains information of CM's MAC,
// CPE's IP prefix type, CPE's IP prefix address,
// CPE's IP prefix length and CPE's MAC address.
// An entry is created if CPE is associated with a prefix.
type CISCODOCSEXTMIB_CdxCpeIpPrefixTable_CdxCpeIpPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates the MAC address of the cable
    // modem. The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCpeIpPrefixCmMacAddress interface{}

    // This attribute is a key. This object indicates the IP prefix type of the
    // CPE. This is the type of cdxCpeIpPrefixAddress object. The type is
    // InetAddressType.
    CdxCpeIpPrefixType interface{}

    // This attribute is a key. This object indicates the IP prefix address. The
    // type of this address is determined by the value of  cdxCpeIpPrefixType
    // object. The type is string with length: 1..96.
    CdxCpeIpPrefixAddress interface{}

    // This attribute is a key. This object indicates the IP prefix length of the
    // CPE. This is the length of cdxCpeIpPrefixAddress object. The type is
    // interface{} with range: 0..2040.
    CdxCpeIpPrefixLen interface{}

    // This object indicates the MAC address of CPE. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxCpeIpPrefixCpeMacAddress interface{}

    // This object indicates the type of CPE. Device Type: B - CM Bridge, R - CM
    // Router IP Assignment Method: D - DHCP the format looks like 'R/D'. The type
    // is string.
    CdxCpeIpPrefixCpeType interface{}
}

func (cdxCpeIpPrefixEntry *CISCODOCSEXTMIB_CdxCpeIpPrefixTable_CdxCpeIpPrefixEntry) GetEntityData() *types.CommonEntityData {
    cdxCpeIpPrefixEntry.EntityData.YFilter = cdxCpeIpPrefixEntry.YFilter
    cdxCpeIpPrefixEntry.EntityData.YangName = "cdxCpeIpPrefixEntry"
    cdxCpeIpPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCpeIpPrefixEntry.EntityData.ParentYangName = "cdxCpeIpPrefixTable"
    cdxCpeIpPrefixEntry.EntityData.SegmentPath = "cdxCpeIpPrefixEntry" + types.AddKeyToken(cdxCpeIpPrefixEntry.CdxCpeIpPrefixCmMacAddress, "cdxCpeIpPrefixCmMacAddress") + types.AddKeyToken(cdxCpeIpPrefixEntry.CdxCpeIpPrefixType, "cdxCpeIpPrefixType") + types.AddKeyToken(cdxCpeIpPrefixEntry.CdxCpeIpPrefixAddress, "cdxCpeIpPrefixAddress") + types.AddKeyToken(cdxCpeIpPrefixEntry.CdxCpeIpPrefixLen, "cdxCpeIpPrefixLen")
    cdxCpeIpPrefixEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCpeIpPrefixTable/" + cdxCpeIpPrefixEntry.EntityData.SegmentPath
    cdxCpeIpPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCpeIpPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCpeIpPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCpeIpPrefixEntry.EntityData.Children = types.NewOrderedMap()
    cdxCpeIpPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixCmMacAddress", types.YLeaf{"CdxCpeIpPrefixCmMacAddress", cdxCpeIpPrefixEntry.CdxCpeIpPrefixCmMacAddress})
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixType", types.YLeaf{"CdxCpeIpPrefixType", cdxCpeIpPrefixEntry.CdxCpeIpPrefixType})
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixAddress", types.YLeaf{"CdxCpeIpPrefixAddress", cdxCpeIpPrefixEntry.CdxCpeIpPrefixAddress})
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixLen", types.YLeaf{"CdxCpeIpPrefixLen", cdxCpeIpPrefixEntry.CdxCpeIpPrefixLen})
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixCpeMacAddress", types.YLeaf{"CdxCpeIpPrefixCpeMacAddress", cdxCpeIpPrefixEntry.CdxCpeIpPrefixCpeMacAddress})
    cdxCpeIpPrefixEntry.EntityData.Leafs.Append("cdxCpeIpPrefixCpeType", types.YLeaf{"CdxCpeIpPrefixCpeType", cdxCpeIpPrefixEntry.CdxCpeIpPrefixCpeType})

    cdxCpeIpPrefixEntry.EntityData.YListKeys = []string {"CdxCpeIpPrefixCmMacAddress", "CdxCpeIpPrefixType", "CdxCpeIpPrefixAddress", "CdxCpeIpPrefixLen"}

    return &(cdxCpeIpPrefixEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable
// This table contains upstream channel attributes for  
// automated Spectrum management, in addition to the ones
// provided by docsIfUpstreamChannelEntry.
// It also contains upstream channel attributes to count 
// the number of active,registered and total number of cable 
// modems on this upstream. 
type CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional objects for docsIfUpstreamChannelEntry,including  the secondary
    // upstream channel modulation profile,the  lower bound for the channel width
    // and the number of active, registered and total number of cable modems on
    // this  upstream channel. . The type is slice of
    // CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable_CdxIfUpstreamChannelExtEntry.
    CdxIfUpstreamChannelExtEntry []*CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable_CdxIfUpstreamChannelExtEntry
}

func (cdxIfUpstreamChannelExtTable *CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable) GetEntityData() *types.CommonEntityData {
    cdxIfUpstreamChannelExtTable.EntityData.YFilter = cdxIfUpstreamChannelExtTable.YFilter
    cdxIfUpstreamChannelExtTable.EntityData.YangName = "cdxIfUpstreamChannelExtTable"
    cdxIfUpstreamChannelExtTable.EntityData.BundleName = "cisco_ios_xe"
    cdxIfUpstreamChannelExtTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxIfUpstreamChannelExtTable.EntityData.SegmentPath = "cdxIfUpstreamChannelExtTable"
    cdxIfUpstreamChannelExtTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxIfUpstreamChannelExtTable.EntityData.SegmentPath
    cdxIfUpstreamChannelExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxIfUpstreamChannelExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxIfUpstreamChannelExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxIfUpstreamChannelExtTable.EntityData.Children = types.NewOrderedMap()
    cdxIfUpstreamChannelExtTable.EntityData.Children.Append("cdxIfUpstreamChannelExtEntry", types.YChild{"CdxIfUpstreamChannelExtEntry", nil})
    for i := range cdxIfUpstreamChannelExtTable.CdxIfUpstreamChannelExtEntry {
        cdxIfUpstreamChannelExtTable.EntityData.Children.Append(types.GetSegmentPath(cdxIfUpstreamChannelExtTable.CdxIfUpstreamChannelExtEntry[i]), types.YChild{"CdxIfUpstreamChannelExtEntry", cdxIfUpstreamChannelExtTable.CdxIfUpstreamChannelExtEntry[i]})
    }
    cdxIfUpstreamChannelExtTable.EntityData.Leafs = types.NewOrderedMap()

    cdxIfUpstreamChannelExtTable.EntityData.YListKeys = []string {}

    return &(cdxIfUpstreamChannelExtTable.EntityData)
}

// CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable_CdxIfUpstreamChannelExtEntry
// Additional objects for docsIfUpstreamChannelEntry,including 
// the secondary upstream channel modulation profile,the 
// lower bound for the channel width and the number of active,
// registered and total number of cable modems on this 
// upstream channel. 
type CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable_CdxIfUpstreamChannelExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The lower bound for the bandwidth of this upstream channel.  The bandwidth
    // specified by docsIfUpChannelWidth is used as the upper bound of the
    // upstream channel. The two objects, docsIfUpChannelWidth and
    // cdxIfUpChannelWidth, in  conjunction, define the upstream channel width
    // range to be used for the automated spectrum management.  This object
    // returns 0 if the channel width is undefined  or unknown.  For those
    // upstreams in the linecards which do not have the automated spectrum
    // management feature, this channel width is undefined and always has value 0.
    // . The type is interface{} with range: 0..16000000. Units are hertz.
    CdxIfUpChannelWidth interface{}

    // The secondary modulation profile for the upstream channel. This should be a
    // QPSK modulation profile if the primary profile  is QAM-16. The CMTS will
    // switch from primary profile (QAM16) to  secondary profile (QPSK) depending
    // on the noise level of a  particular spectrum band.  This is an entry
    // identical to the docsIfModIndex in the  docsIfCmtsModulationTable that
    // describes this channel. This channel is further instantiated there by a
    // grouping of interval usage codes which together fully describe the channel
    // modulation. This object returns 0 if the docsIfCmtsModulationTable does not
    // exist or is empty. . The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelModulationProfile interface{}

    // The total count of cable modems on this upstream channel since boot. The
    // type is interface{} with range: 0..8191.
    CdxIfUpChannelCmTotal interface{}

    // The count of cable modems that are active.Active cable  modems are
    // recognized by the cdxCmtsCmStatusValue other   than offline(1). . The type
    // is interface{} with range: 0..8191.
    CdxIfUpChannelCmActive interface{}

    // The count of cable modems that are registered and online on this upstream.
    // Registered cable modems are those with one of the following values:
    // registrationComplete(6) of docsIfCmtsCmStatusValue OR online(12),
    // kekRejected(10), onlineKekAssigned(6), tekRejected(11),onlineTekAssigned(7)
    // of  cdxCmtsCmStatusValue. The type is interface{} with range: 0..8191.
    CdxIfUpChannelCmRegistered interface{}

    // The Upstream Input power level at the CMTS interface. This is the expected
    // power level and is different from the actual power received. If not
    // configured the default value is 0 dBmV and is also the optimum setting
    // power level for the upstream. For FPGA line cards, the valid range is <-10
    // to 10> dBmV and for ASIC Line cards, it is  <-10  to 25> dBmV. . The type
    // is interface{} with range: -100..250.
    CdxIfUpChannelInputPowerLevel interface{}

    // The average percentage of upstream channel utilization.  This item
    // indicates the running average of percent channel utilization in CMTS
    // upstream Mac scheduler. . The type is interface{} with range: 0..100. Units
    // are percent.
    CdxIfUpChannelAvgUtil interface{}

    // The average percentage of contention mini-slots. This item indicates the
    // running average of percent contention mini-slots in CMTS upstream Mac
    // scheduler. . The type is interface{} with range: 0..100. Units are percent.
    CdxIfUpChannelAvgContSlots interface{}

    // The average percentage of initial ranging mini-slots.  This item indicates
    // the running average of percent initial ranging mini-slots in CMTS upstream
    // Mac scheduler. . The type is interface{} with range: 0..100. Units are
    // percent.
    CdxIfUpChannelRangeSlots interface{}

    // This object indicates the number of active  Unsolicited Grant Service (UGS)
    // on a given upstream. This would be used for the user to evaluate traffic 
    // load at any given time of the day.  The Unsolicited Grant Service (UGS) is
    // designed to  support real-time service flows that generate fixed size data
    // packets on a periodic basis. . The type is interface{} with range:
    // 0..4294967295.
    CdxIfUpChannelNumActiveUGS interface{}

    // This object indicates the maximum number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last one hour. This would be
    // used for the user to evaluate traffic load at any given time of the day. 
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelMaxUGSLastOneHour interface{}

    // This object indicates the minimum number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last one hour. This would be
    // used for the user to evaluate traffic load at any given time of the day. 
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelMinUGSLastOneHour interface{}

    // This object indicates the average number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last one hour. This would be
    // used for the user to evaluate traffic load at any given time of the day. 
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelAvgUGSLastOneHour interface{}

    // This object indicates the maximum number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last five minutes. This would 
    // be used for the user to evaluate traffic load at any given time of the day.
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelMaxUGSLastFiveMins interface{}

    // This object indicates the minimum number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last five minutes. This would 
    // be used for the user to evaluate traffic load at any given time of the day.
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelMinUGSLastFiveMins interface{}

    // This object indicates the average number of  Unsolicited Grant Service
    // (UGS) allocated on a given upstream in the last five minutes. This would 
    // be used for the user to evaluate traffic load at any given time of the day.
    // The Unsolicited Grant Service (UGS) is designed to support real-time
    // service flows that generate fixed size data packets on a periodic basis. .
    // The type is interface{} with range: 0..4294967295.
    CdxIfUpChannelAvgUGSLastFiveMins interface{}
}

func (cdxIfUpstreamChannelExtEntry *CISCODOCSEXTMIB_CdxIfUpstreamChannelExtTable_CdxIfUpstreamChannelExtEntry) GetEntityData() *types.CommonEntityData {
    cdxIfUpstreamChannelExtEntry.EntityData.YFilter = cdxIfUpstreamChannelExtEntry.YFilter
    cdxIfUpstreamChannelExtEntry.EntityData.YangName = "cdxIfUpstreamChannelExtEntry"
    cdxIfUpstreamChannelExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxIfUpstreamChannelExtEntry.EntityData.ParentYangName = "cdxIfUpstreamChannelExtTable"
    cdxIfUpstreamChannelExtEntry.EntityData.SegmentPath = "cdxIfUpstreamChannelExtEntry" + types.AddKeyToken(cdxIfUpstreamChannelExtEntry.IfIndex, "ifIndex")
    cdxIfUpstreamChannelExtEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxIfUpstreamChannelExtTable/" + cdxIfUpstreamChannelExtEntry.EntityData.SegmentPath
    cdxIfUpstreamChannelExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxIfUpstreamChannelExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxIfUpstreamChannelExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxIfUpstreamChannelExtEntry.EntityData.Children = types.NewOrderedMap()
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxIfUpstreamChannelExtEntry.IfIndex})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelWidth", types.YLeaf{"CdxIfUpChannelWidth", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelWidth})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelModulationProfile", types.YLeaf{"CdxIfUpChannelModulationProfile", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelModulationProfile})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelCmTotal", types.YLeaf{"CdxIfUpChannelCmTotal", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelCmTotal})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelCmActive", types.YLeaf{"CdxIfUpChannelCmActive", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelCmActive})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelCmRegistered", types.YLeaf{"CdxIfUpChannelCmRegistered", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelCmRegistered})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelInputPowerLevel", types.YLeaf{"CdxIfUpChannelInputPowerLevel", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelInputPowerLevel})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelAvgUtil", types.YLeaf{"CdxIfUpChannelAvgUtil", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelAvgUtil})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelAvgContSlots", types.YLeaf{"CdxIfUpChannelAvgContSlots", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelAvgContSlots})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelRangeSlots", types.YLeaf{"CdxIfUpChannelRangeSlots", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelRangeSlots})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelNumActiveUGS", types.YLeaf{"CdxIfUpChannelNumActiveUGS", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelNumActiveUGS})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelMaxUGSLastOneHour", types.YLeaf{"CdxIfUpChannelMaxUGSLastOneHour", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelMaxUGSLastOneHour})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelMinUGSLastOneHour", types.YLeaf{"CdxIfUpChannelMinUGSLastOneHour", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelMinUGSLastOneHour})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelAvgUGSLastOneHour", types.YLeaf{"CdxIfUpChannelAvgUGSLastOneHour", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelAvgUGSLastOneHour})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelMaxUGSLastFiveMins", types.YLeaf{"CdxIfUpChannelMaxUGSLastFiveMins", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelMaxUGSLastFiveMins})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelMinUGSLastFiveMins", types.YLeaf{"CdxIfUpChannelMinUGSLastFiveMins", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelMinUGSLastFiveMins})
    cdxIfUpstreamChannelExtEntry.EntityData.Leafs.Append("cdxIfUpChannelAvgUGSLastFiveMins", types.YLeaf{"CdxIfUpChannelAvgUGSLastFiveMins", cdxIfUpstreamChannelExtEntry.CdxIfUpChannelAvgUGSLastFiveMins})

    cdxIfUpstreamChannelExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxIfUpstreamChannelExtEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxWBResilCmTable
// This table contains information about partial service cable
// modems (CM), including both downstream and upstream partial
// service modems.
type CISCODOCSEXTMIB_CdxWBResilCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list contains information for a partial service cable modem (CM). 
    // Provided the following information for a partial service cable modem: How
    // many downstream channels in total; How many upstream channels in total; How
    // many active downstream channels; How many active upstream channels; Which
    // downstream channels are in partial service mode; Which upstream channels
    // are in partial service mode;. The type is slice of
    // CISCODOCSEXTMIB_CdxWBResilCmTable_CdxWBResilCmEntry.
    CdxWBResilCmEntry []*CISCODOCSEXTMIB_CdxWBResilCmTable_CdxWBResilCmEntry
}

func (cdxWBResilCmTable *CISCODOCSEXTMIB_CdxWBResilCmTable) GetEntityData() *types.CommonEntityData {
    cdxWBResilCmTable.EntityData.YFilter = cdxWBResilCmTable.YFilter
    cdxWBResilCmTable.EntityData.YangName = "cdxWBResilCmTable"
    cdxWBResilCmTable.EntityData.BundleName = "cisco_ios_xe"
    cdxWBResilCmTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxWBResilCmTable.EntityData.SegmentPath = "cdxWBResilCmTable"
    cdxWBResilCmTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxWBResilCmTable.EntityData.SegmentPath
    cdxWBResilCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxWBResilCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxWBResilCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxWBResilCmTable.EntityData.Children = types.NewOrderedMap()
    cdxWBResilCmTable.EntityData.Children.Append("cdxWBResilCmEntry", types.YChild{"CdxWBResilCmEntry", nil})
    for i := range cdxWBResilCmTable.CdxWBResilCmEntry {
        cdxWBResilCmTable.EntityData.Children.Append(types.GetSegmentPath(cdxWBResilCmTable.CdxWBResilCmEntry[i]), types.YChild{"CdxWBResilCmEntry", cdxWBResilCmTable.CdxWBResilCmEntry[i]})
    }
    cdxWBResilCmTable.EntityData.Leafs = types.NewOrderedMap()

    cdxWBResilCmTable.EntityData.YListKeys = []string {}

    return &(cdxWBResilCmTable.EntityData)
}

// CISCODOCSEXTMIB_CdxWBResilCmTable_CdxWBResilCmEntry
// The list contains information for a partial service cable modem
// (CM).
// 
// Provided the following information for a partial service cable
// modem:
// How many downstream channels in total;
// How many upstream channels in total;
// How many active downstream channels;
// How many active upstream channels;
// Which downstream channels are in partial service mode;
// Which upstream channels are in partial service mode;
type CISCODOCSEXTMIB_CdxWBResilCmTable_CdxWBResilCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This attribute uniquely identifies a CM.  The CMTS
    // must assign a single id value for each CM MAC address seen by the CMTS. 
    // The CMTS should ensure that the association between an Id and MAC Address
    // remains constant during CMTS uptime. The type is interface{} with range:
    // 1..4294967295.
    CdxWBResilCmIndex interface{}

    // This attribute represents the MAC address of the CM. If the CM has multiple
    // MAC addresses, this is the MAC address associated with the MAC Domain
    // interface. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxWBResilCmMacAddr interface{}

    // Total downstream channel number of the CM. The type is interface{} with
    // range: 0..4294967295.
    CdxWBResilCmTotalDsNum interface{}

    // Total upstream channel number of the CM. The type is interface{} with
    // range: 0..4294967295.
    CdxWBResilCmTotalUsNum interface{}

    // Current active downstream channel number, it's the total downstream channel
    // minus downstream partial service channel number. The type is interface{}
    // with range: 0..4294967295.
    CdxWBResilCmCurrentDsNum interface{}

    // Current active upstream channel number, it's the total upstream channel
    // minus upstream partial service channel number. The type is interface{} with
    // range: 0..4294967295.
    CdxWBResilCmCurrentUsNum interface{}

    // Impaired downstream channel index list. The index in list is rf channel
    // ifIndex. If there's no downstream channel impaired, return empty. The
    // output looks like: '137000 137001 137002'. The type is string.
    CdxWBResilCmImpairedDsChIndex interface{}

    // Impaired upstream channel index list. The index in list is upstream channel
    // ifIndex. If there's no upstream channel impaired, return empty. The output
    // looks like: '196408 196409 196410'. The type is string.
    CdxWBResilCmImpairedUsChIndex interface{}
}

func (cdxWBResilCmEntry *CISCODOCSEXTMIB_CdxWBResilCmTable_CdxWBResilCmEntry) GetEntityData() *types.CommonEntityData {
    cdxWBResilCmEntry.EntityData.YFilter = cdxWBResilCmEntry.YFilter
    cdxWBResilCmEntry.EntityData.YangName = "cdxWBResilCmEntry"
    cdxWBResilCmEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxWBResilCmEntry.EntityData.ParentYangName = "cdxWBResilCmTable"
    cdxWBResilCmEntry.EntityData.SegmentPath = "cdxWBResilCmEntry" + types.AddKeyToken(cdxWBResilCmEntry.CdxWBResilCmIndex, "cdxWBResilCmIndex")
    cdxWBResilCmEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxWBResilCmTable/" + cdxWBResilCmEntry.EntityData.SegmentPath
    cdxWBResilCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxWBResilCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxWBResilCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxWBResilCmEntry.EntityData.Children = types.NewOrderedMap()
    cdxWBResilCmEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmIndex", types.YLeaf{"CdxWBResilCmIndex", cdxWBResilCmEntry.CdxWBResilCmIndex})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmMacAddr", types.YLeaf{"CdxWBResilCmMacAddr", cdxWBResilCmEntry.CdxWBResilCmMacAddr})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmTotalDsNum", types.YLeaf{"CdxWBResilCmTotalDsNum", cdxWBResilCmEntry.CdxWBResilCmTotalDsNum})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmTotalUsNum", types.YLeaf{"CdxWBResilCmTotalUsNum", cdxWBResilCmEntry.CdxWBResilCmTotalUsNum})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmCurrentDsNum", types.YLeaf{"CdxWBResilCmCurrentDsNum", cdxWBResilCmEntry.CdxWBResilCmCurrentDsNum})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmCurrentUsNum", types.YLeaf{"CdxWBResilCmCurrentUsNum", cdxWBResilCmEntry.CdxWBResilCmCurrentUsNum})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmImpairedDsChIndex", types.YLeaf{"CdxWBResilCmImpairedDsChIndex", cdxWBResilCmEntry.CdxWBResilCmImpairedDsChIndex})
    cdxWBResilCmEntry.EntityData.Leafs.Append("cdxWBResilCmImpairedUsChIndex", types.YLeaf{"CdxWBResilCmImpairedUsChIndex", cdxWBResilCmEntry.CdxWBResilCmImpairedUsChIndex})

    cdxWBResilCmEntry.EntityData.YListKeys = []string {"CdxWBResilCmIndex"}

    return &(cdxWBResilCmEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable
// This table contains information of the mapping of
// the physical RF channels to the primary RF channels.
type CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An Entry provides the association between the physical RF channels and the
    // primary RF Channels. The type is slice of
    // CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable_CdxRFtoPrimaryChannelMappingEntry.
    CdxRFtoPrimaryChannelMappingEntry []*CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable_CdxRFtoPrimaryChannelMappingEntry
}

func (cdxRFtoPrimaryChannelMappingTable *CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable) GetEntityData() *types.CommonEntityData {
    cdxRFtoPrimaryChannelMappingTable.EntityData.YFilter = cdxRFtoPrimaryChannelMappingTable.YFilter
    cdxRFtoPrimaryChannelMappingTable.EntityData.YangName = "cdxRFtoPrimaryChannelMappingTable"
    cdxRFtoPrimaryChannelMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cdxRFtoPrimaryChannelMappingTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxRFtoPrimaryChannelMappingTable.EntityData.SegmentPath = "cdxRFtoPrimaryChannelMappingTable"
    cdxRFtoPrimaryChannelMappingTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxRFtoPrimaryChannelMappingTable.EntityData.SegmentPath
    cdxRFtoPrimaryChannelMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxRFtoPrimaryChannelMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxRFtoPrimaryChannelMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxRFtoPrimaryChannelMappingTable.EntityData.Children = types.NewOrderedMap()
    cdxRFtoPrimaryChannelMappingTable.EntityData.Children.Append("cdxRFtoPrimaryChannelMappingEntry", types.YChild{"CdxRFtoPrimaryChannelMappingEntry", nil})
    for i := range cdxRFtoPrimaryChannelMappingTable.CdxRFtoPrimaryChannelMappingEntry {
        cdxRFtoPrimaryChannelMappingTable.EntityData.Children.Append(types.GetSegmentPath(cdxRFtoPrimaryChannelMappingTable.CdxRFtoPrimaryChannelMappingEntry[i]), types.YChild{"CdxRFtoPrimaryChannelMappingEntry", cdxRFtoPrimaryChannelMappingTable.CdxRFtoPrimaryChannelMappingEntry[i]})
    }
    cdxRFtoPrimaryChannelMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cdxRFtoPrimaryChannelMappingTable.EntityData.YListKeys = []string {}

    return &(cdxRFtoPrimaryChannelMappingTable.EntityData)
}

// CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable_CdxRFtoPrimaryChannelMappingEntry
// An Entry provides the association between the physical
// RF channels and the primary RF Channels.
type CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable_CdxRFtoPrimaryChannelMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The ifIndex of the primary channel interface. The type is interface{} with
    // range: 1..2147483647.
    CdxPrimaryChannelIfIndex interface{}
}

func (cdxRFtoPrimaryChannelMappingEntry *CISCODOCSEXTMIB_CdxRFtoPrimaryChannelMappingTable_CdxRFtoPrimaryChannelMappingEntry) GetEntityData() *types.CommonEntityData {
    cdxRFtoPrimaryChannelMappingEntry.EntityData.YFilter = cdxRFtoPrimaryChannelMappingEntry.YFilter
    cdxRFtoPrimaryChannelMappingEntry.EntityData.YangName = "cdxRFtoPrimaryChannelMappingEntry"
    cdxRFtoPrimaryChannelMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxRFtoPrimaryChannelMappingEntry.EntityData.ParentYangName = "cdxRFtoPrimaryChannelMappingTable"
    cdxRFtoPrimaryChannelMappingEntry.EntityData.SegmentPath = "cdxRFtoPrimaryChannelMappingEntry" + types.AddKeyToken(cdxRFtoPrimaryChannelMappingEntry.IfIndex, "ifIndex")
    cdxRFtoPrimaryChannelMappingEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxRFtoPrimaryChannelMappingTable/" + cdxRFtoPrimaryChannelMappingEntry.EntityData.SegmentPath
    cdxRFtoPrimaryChannelMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxRFtoPrimaryChannelMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxRFtoPrimaryChannelMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxRFtoPrimaryChannelMappingEntry.EntityData.Children = types.NewOrderedMap()
    cdxRFtoPrimaryChannelMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxRFtoPrimaryChannelMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxRFtoPrimaryChannelMappingEntry.IfIndex})
    cdxRFtoPrimaryChannelMappingEntry.EntityData.Leafs.Append("cdxPrimaryChannelIfIndex", types.YLeaf{"CdxPrimaryChannelIfIndex", cdxRFtoPrimaryChannelMappingEntry.CdxPrimaryChannelIfIndex})

    cdxRFtoPrimaryChannelMappingEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxRFtoPrimaryChannelMappingEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable
// This table contains information of the mapping of
// the primary RF channels to the physical RF channels.
type CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An Entry provides the association between the primary RF channels and the
    // physical RF Channels. The type is slice of
    // CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable_CdxPrimaryChanneltoRFMappingEntry.
    CdxPrimaryChanneltoRFMappingEntry []*CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable_CdxPrimaryChanneltoRFMappingEntry
}

func (cdxPrimaryChanneltoRFMappingTable *CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable) GetEntityData() *types.CommonEntityData {
    cdxPrimaryChanneltoRFMappingTable.EntityData.YFilter = cdxPrimaryChanneltoRFMappingTable.YFilter
    cdxPrimaryChanneltoRFMappingTable.EntityData.YangName = "cdxPrimaryChanneltoRFMappingTable"
    cdxPrimaryChanneltoRFMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cdxPrimaryChanneltoRFMappingTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxPrimaryChanneltoRFMappingTable.EntityData.SegmentPath = "cdxPrimaryChanneltoRFMappingTable"
    cdxPrimaryChanneltoRFMappingTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxPrimaryChanneltoRFMappingTable.EntityData.SegmentPath
    cdxPrimaryChanneltoRFMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxPrimaryChanneltoRFMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxPrimaryChanneltoRFMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxPrimaryChanneltoRFMappingTable.EntityData.Children = types.NewOrderedMap()
    cdxPrimaryChanneltoRFMappingTable.EntityData.Children.Append("cdxPrimaryChanneltoRFMappingEntry", types.YChild{"CdxPrimaryChanneltoRFMappingEntry", nil})
    for i := range cdxPrimaryChanneltoRFMappingTable.CdxPrimaryChanneltoRFMappingEntry {
        cdxPrimaryChanneltoRFMappingTable.EntityData.Children.Append(types.GetSegmentPath(cdxPrimaryChanneltoRFMappingTable.CdxPrimaryChanneltoRFMappingEntry[i]), types.YChild{"CdxPrimaryChanneltoRFMappingEntry", cdxPrimaryChanneltoRFMappingTable.CdxPrimaryChanneltoRFMappingEntry[i]})
    }
    cdxPrimaryChanneltoRFMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cdxPrimaryChanneltoRFMappingTable.EntityData.YListKeys = []string {}

    return &(cdxPrimaryChanneltoRFMappingTable.EntityData)
}

// CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable_CdxPrimaryChanneltoRFMappingEntry
// An Entry provides the association between the primary
// RF channels and the physical RF Channels.
type CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable_CdxPrimaryChanneltoRFMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The ifIndex of the physical RF channel interface. The type is interface{}
    // with range: 1..2147483647.
    CdxPhysicalRFIfIndex interface{}
}

func (cdxPrimaryChanneltoRFMappingEntry *CISCODOCSEXTMIB_CdxPrimaryChanneltoRFMappingTable_CdxPrimaryChanneltoRFMappingEntry) GetEntityData() *types.CommonEntityData {
    cdxPrimaryChanneltoRFMappingEntry.EntityData.YFilter = cdxPrimaryChanneltoRFMappingEntry.YFilter
    cdxPrimaryChanneltoRFMappingEntry.EntityData.YangName = "cdxPrimaryChanneltoRFMappingEntry"
    cdxPrimaryChanneltoRFMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxPrimaryChanneltoRFMappingEntry.EntityData.ParentYangName = "cdxPrimaryChanneltoRFMappingTable"
    cdxPrimaryChanneltoRFMappingEntry.EntityData.SegmentPath = "cdxPrimaryChanneltoRFMappingEntry" + types.AddKeyToken(cdxPrimaryChanneltoRFMappingEntry.IfIndex, "ifIndex")
    cdxPrimaryChanneltoRFMappingEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxPrimaryChanneltoRFMappingTable/" + cdxPrimaryChanneltoRFMappingEntry.EntityData.SegmentPath
    cdxPrimaryChanneltoRFMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxPrimaryChanneltoRFMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxPrimaryChanneltoRFMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxPrimaryChanneltoRFMappingEntry.EntityData.Children = types.NewOrderedMap()
    cdxPrimaryChanneltoRFMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxPrimaryChanneltoRFMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxPrimaryChanneltoRFMappingEntry.IfIndex})
    cdxPrimaryChanneltoRFMappingEntry.EntityData.Leafs.Append("cdxPhysicalRFIfIndex", types.YLeaf{"CdxPhysicalRFIfIndex", cdxPrimaryChanneltoRFMappingEntry.CdxPhysicalRFIfIndex})

    cdxPrimaryChanneltoRFMappingEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxPrimaryChanneltoRFMappingEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsMtcCmTable
// This table contains CM management information of Transmit
// Channel Set(TCS) in the system.
type CISCODOCSEXTMIB_CdxCmtsMtcCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides the CM statistics and management information of a
    // specific TCS. The interface populated in this table is of ifType =
    // docsCableMaclayer(127). The type is slice of
    // CISCODOCSEXTMIB_CdxCmtsMtcCmTable_CdxCmtsMtcCmEntry.
    CdxCmtsMtcCmEntry []*CISCODOCSEXTMIB_CdxCmtsMtcCmTable_CdxCmtsMtcCmEntry
}

func (cdxCmtsMtcCmTable *CISCODOCSEXTMIB_CdxCmtsMtcCmTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsMtcCmTable.EntityData.YFilter = cdxCmtsMtcCmTable.YFilter
    cdxCmtsMtcCmTable.EntityData.YangName = "cdxCmtsMtcCmTable"
    cdxCmtsMtcCmTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsMtcCmTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsMtcCmTable.EntityData.SegmentPath = "cdxCmtsMtcCmTable"
    cdxCmtsMtcCmTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsMtcCmTable.EntityData.SegmentPath
    cdxCmtsMtcCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsMtcCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsMtcCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsMtcCmTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsMtcCmTable.EntityData.Children.Append("cdxCmtsMtcCmEntry", types.YChild{"CdxCmtsMtcCmEntry", nil})
    for i := range cdxCmtsMtcCmTable.CdxCmtsMtcCmEntry {
        cdxCmtsMtcCmTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsMtcCmTable.CdxCmtsMtcCmEntry[i]), types.YChild{"CdxCmtsMtcCmEntry", cdxCmtsMtcCmTable.CdxCmtsMtcCmEntry[i]})
    }
    cdxCmtsMtcCmTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsMtcCmTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsMtcCmTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsMtcCmTable_CdxCmtsMtcCmEntry
// An entry provides the CM statistics and management
// information of a specific TCS. The interface populated in this
// table is of ifType = docsCableMaclayer(127).
type CISCODOCSEXTMIB_CdxCmtsMtcCmTable_CdxCmtsMtcCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This object indicates the Id of the Transmit
    // Channel Set. The type is interface{} with range: 0..4294967295.
    CdxCmtsMtcTcsId interface{}

    // This object indicates the total number of cable modems which use this TCS
    // in the MAC domain. The type is interface{} with range: 0..4294967295.
    CdxCmtsMtcCmTotal interface{}

    // This object indicates the number of operational cable modems which uses
    // this TCS in the MAC domain. The type is interface{} with range:
    // 0..4294967295.
    CdxCMtsMtcCmOperational interface{}

    // This object indicates the number of registered cable modems which use this
    // TCS in the MAC domain. The type is interface{} with range: 0..4294967295.
    CdxCmtsMtcCmRegistered interface{}

    // This object indicates the number of unregistered cable modem which use this
    // TCS in the MAC domain. The type is interface{} with range: 0..4294967295.
    CdxCmtsMtcCmUnregistered interface{}

    // This object indicates the number of offline cable modems which uses this
    // TCS in the MAC domain. The type is interface{} with range: 0..4294967295.
    CdxCmtsMtcCmOffline interface{}

    // This object indicates the number of operational cable modems which are in
    // wideband state and use this TCS in the MAC domain. The type is interface{}
    // with range: 0..4294967295.
    CdxCmtsMtcCmWideband interface{}

    // This object specifies the upstream channel bonding group. The type is
    // string.
    CdxCmtsMtcUpstreamBondGrp interface{}
}

func (cdxCmtsMtcCmEntry *CISCODOCSEXTMIB_CdxCmtsMtcCmTable_CdxCmtsMtcCmEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsMtcCmEntry.EntityData.YFilter = cdxCmtsMtcCmEntry.YFilter
    cdxCmtsMtcCmEntry.EntityData.YangName = "cdxCmtsMtcCmEntry"
    cdxCmtsMtcCmEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsMtcCmEntry.EntityData.ParentYangName = "cdxCmtsMtcCmTable"
    cdxCmtsMtcCmEntry.EntityData.SegmentPath = "cdxCmtsMtcCmEntry" + types.AddKeyToken(cdxCmtsMtcCmEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxCmtsMtcCmEntry.CdxCmtsMtcTcsId, "cdxCmtsMtcTcsId")
    cdxCmtsMtcCmEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsMtcCmTable/" + cdxCmtsMtcCmEntry.EntityData.SegmentPath
    cdxCmtsMtcCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsMtcCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsMtcCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsMtcCmEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsMtcCmEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxCmtsMtcCmEntry.IfIndex})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcTcsId", types.YLeaf{"CdxCmtsMtcTcsId", cdxCmtsMtcCmEntry.CdxCmtsMtcTcsId})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcCmTotal", types.YLeaf{"CdxCmtsMtcCmTotal", cdxCmtsMtcCmEntry.CdxCmtsMtcCmTotal})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCMtsMtcCmOperational", types.YLeaf{"CdxCMtsMtcCmOperational", cdxCmtsMtcCmEntry.CdxCMtsMtcCmOperational})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcCmRegistered", types.YLeaf{"CdxCmtsMtcCmRegistered", cdxCmtsMtcCmEntry.CdxCmtsMtcCmRegistered})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcCmUnregistered", types.YLeaf{"CdxCmtsMtcCmUnregistered", cdxCmtsMtcCmEntry.CdxCmtsMtcCmUnregistered})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcCmOffline", types.YLeaf{"CdxCmtsMtcCmOffline", cdxCmtsMtcCmEntry.CdxCmtsMtcCmOffline})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcCmWideband", types.YLeaf{"CdxCmtsMtcCmWideband", cdxCmtsMtcCmEntry.CdxCmtsMtcCmWideband})
    cdxCmtsMtcCmEntry.EntityData.Leafs.Append("cdxCmtsMtcUpstreamBondGrp", types.YLeaf{"CdxCmtsMtcUpstreamBondGrp", cdxCmtsMtcCmEntry.CdxCmtsMtcUpstreamBondGrp})

    cdxCmtsMtcCmEntry.EntityData.YListKeys = []string {"IfIndex", "CdxCmtsMtcTcsId"}

    return &(cdxCmtsMtcCmEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsUscbSflowTable
// This table contains the Upstream Channel Bonding
// Service Flow management information.
type CISCODOCSEXTMIB_CdxCmtsUscbSflowTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A entry contains the Service Flow statistics for a specific Upstream
    // Channel Bonding group. The interface populated in this table is of ifType =
    // docsCableMaclayer(127). The type is slice of
    // CISCODOCSEXTMIB_CdxCmtsUscbSflowTable_CdxCmtsUscbSflowEntry.
    CdxCmtsUscbSflowEntry []*CISCODOCSEXTMIB_CdxCmtsUscbSflowTable_CdxCmtsUscbSflowEntry
}

func (cdxCmtsUscbSflowTable *CISCODOCSEXTMIB_CdxCmtsUscbSflowTable) GetEntityData() *types.CommonEntityData {
    cdxCmtsUscbSflowTable.EntityData.YFilter = cdxCmtsUscbSflowTable.YFilter
    cdxCmtsUscbSflowTable.EntityData.YangName = "cdxCmtsUscbSflowTable"
    cdxCmtsUscbSflowTable.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsUscbSflowTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxCmtsUscbSflowTable.EntityData.SegmentPath = "cdxCmtsUscbSflowTable"
    cdxCmtsUscbSflowTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxCmtsUscbSflowTable.EntityData.SegmentPath
    cdxCmtsUscbSflowTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsUscbSflowTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsUscbSflowTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsUscbSflowTable.EntityData.Children = types.NewOrderedMap()
    cdxCmtsUscbSflowTable.EntityData.Children.Append("cdxCmtsUscbSflowEntry", types.YChild{"CdxCmtsUscbSflowEntry", nil})
    for i := range cdxCmtsUscbSflowTable.CdxCmtsUscbSflowEntry {
        cdxCmtsUscbSflowTable.EntityData.Children.Append(types.GetSegmentPath(cdxCmtsUscbSflowTable.CdxCmtsUscbSflowEntry[i]), types.YChild{"CdxCmtsUscbSflowEntry", cdxCmtsUscbSflowTable.CdxCmtsUscbSflowEntry[i]})
    }
    cdxCmtsUscbSflowTable.EntityData.Leafs = types.NewOrderedMap()

    cdxCmtsUscbSflowTable.EntityData.YListKeys = []string {}

    return &(cdxCmtsUscbSflowTable.EntityData)
}

// CISCODOCSEXTMIB_CdxCmtsUscbSflowTable_CdxCmtsUscbSflowEntry
// A entry contains the Service Flow statistics for a specific
// Upstream Channel Bonding group. The interface populated in this
// table is of ifType = docsCableMaclayer(127).
type CISCODOCSEXTMIB_CdxCmtsUscbSflowTable_CdxCmtsUscbSflowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This object indicates upstream bonding group
    // identifier within the MAC Domain. The type is interface{} with range:
    // 1..4294967295.
    CdxCmtsUsBondingGrpId interface{}

    // This object indicates the total number of service flows which use this
    // upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbSfTotal interface{}

    // This object indicates the total number of  primary service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbSfPri interface{}

    // This object indicates the number of static BE service flows which use this
    // upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbStaticSfBe interface{}

    // This object indicates the number of static UGS service flows which use this
    // upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbStaticSfUgs interface{}

    // This object indicates the number of static UGS-AD service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbStaticSfUgsad interface{}

    // This object indicates the number of static RTPS service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbStaticSfRtps interface{}

    // This object indicates the number of static NRTPS service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbStaticSfNrtps interface{}

    // This object indicates the number of dynamic BE service flows which use this
    // upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbDynSfBe interface{}

    // This object indicates the number of dynamic UGS service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbDynSfUgs interface{}

    // This object indicates the number of dynamic UGS-Ad service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbDynSfUgsad interface{}

    // This object indicates the number of dynamic RTPS service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbDynSfRtps interface{}

    // This object indicates the number of dynamic NRTPS service flows which use
    // this upstream channel bonding group. The type is interface{} with range:
    // 0..4294967295.
    CdxCmtsUscbDynSfNrtps interface{}

    // This object indicates the description of upstream channel bonding group.
    // The type is string.
    CdxCmtsUscbDescr interface{}
}

func (cdxCmtsUscbSflowEntry *CISCODOCSEXTMIB_CdxCmtsUscbSflowTable_CdxCmtsUscbSflowEntry) GetEntityData() *types.CommonEntityData {
    cdxCmtsUscbSflowEntry.EntityData.YFilter = cdxCmtsUscbSflowEntry.YFilter
    cdxCmtsUscbSflowEntry.EntityData.YangName = "cdxCmtsUscbSflowEntry"
    cdxCmtsUscbSflowEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxCmtsUscbSflowEntry.EntityData.ParentYangName = "cdxCmtsUscbSflowTable"
    cdxCmtsUscbSflowEntry.EntityData.SegmentPath = "cdxCmtsUscbSflowEntry" + types.AddKeyToken(cdxCmtsUscbSflowEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxCmtsUscbSflowEntry.CdxCmtsUsBondingGrpId, "cdxCmtsUsBondingGrpId")
    cdxCmtsUscbSflowEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxCmtsUscbSflowTable/" + cdxCmtsUscbSflowEntry.EntityData.SegmentPath
    cdxCmtsUscbSflowEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxCmtsUscbSflowEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxCmtsUscbSflowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxCmtsUscbSflowEntry.EntityData.Children = types.NewOrderedMap()
    cdxCmtsUscbSflowEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxCmtsUscbSflowEntry.IfIndex})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUsBondingGrpId", types.YLeaf{"CdxCmtsUsBondingGrpId", cdxCmtsUscbSflowEntry.CdxCmtsUsBondingGrpId})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbSfTotal", types.YLeaf{"CdxCmtsUscbSfTotal", cdxCmtsUscbSflowEntry.CdxCmtsUscbSfTotal})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbSfPri", types.YLeaf{"CdxCmtsUscbSfPri", cdxCmtsUscbSflowEntry.CdxCmtsUscbSfPri})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbStaticSfBe", types.YLeaf{"CdxCmtsUscbStaticSfBe", cdxCmtsUscbSflowEntry.CdxCmtsUscbStaticSfBe})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbStaticSfUgs", types.YLeaf{"CdxCmtsUscbStaticSfUgs", cdxCmtsUscbSflowEntry.CdxCmtsUscbStaticSfUgs})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbStaticSfUgsad", types.YLeaf{"CdxCmtsUscbStaticSfUgsad", cdxCmtsUscbSflowEntry.CdxCmtsUscbStaticSfUgsad})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbStaticSfRtps", types.YLeaf{"CdxCmtsUscbStaticSfRtps", cdxCmtsUscbSflowEntry.CdxCmtsUscbStaticSfRtps})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbStaticSfNrtps", types.YLeaf{"CdxCmtsUscbStaticSfNrtps", cdxCmtsUscbSflowEntry.CdxCmtsUscbStaticSfNrtps})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDynSfBe", types.YLeaf{"CdxCmtsUscbDynSfBe", cdxCmtsUscbSflowEntry.CdxCmtsUscbDynSfBe})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDynSfUgs", types.YLeaf{"CdxCmtsUscbDynSfUgs", cdxCmtsUscbSflowEntry.CdxCmtsUscbDynSfUgs})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDynSfUgsad", types.YLeaf{"CdxCmtsUscbDynSfUgsad", cdxCmtsUscbSflowEntry.CdxCmtsUscbDynSfUgsad})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDynSfRtps", types.YLeaf{"CdxCmtsUscbDynSfRtps", cdxCmtsUscbSflowEntry.CdxCmtsUscbDynSfRtps})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDynSfNrtps", types.YLeaf{"CdxCmtsUscbDynSfNrtps", cdxCmtsUscbSflowEntry.CdxCmtsUscbDynSfNrtps})
    cdxCmtsUscbSflowEntry.EntityData.Leafs.Append("cdxCmtsUscbDescr", types.YLeaf{"CdxCmtsUscbDescr", cdxCmtsUscbSflowEntry.CdxCmtsUscbDescr})

    cdxCmtsUscbSflowEntry.EntityData.YListKeys = []string {"IfIndex", "CdxCmtsUsBondingGrpId"}

    return &(cdxCmtsUscbSflowEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxRPDGS7KTable
// The cdxRPDGS7KTable contains the attributes of GS7K. 
// An Entry exists for each instance. 
// It is indexed by GS7K's MacAddress.
type CISCODOCSEXTMIB_CdxRPDGS7KTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of statistics for all the sensor,  such as volatage, the state of
    // TriSwitch. The type is slice of
    // CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry.
    CdxRPDGS7KEntry []*CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry
}

func (cdxRPDGS7KTable *CISCODOCSEXTMIB_CdxRPDGS7KTable) GetEntityData() *types.CommonEntityData {
    cdxRPDGS7KTable.EntityData.YFilter = cdxRPDGS7KTable.YFilter
    cdxRPDGS7KTable.EntityData.YangName = "cdxRPDGS7KTable"
    cdxRPDGS7KTable.EntityData.BundleName = "cisco_ios_xe"
    cdxRPDGS7KTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxRPDGS7KTable.EntityData.SegmentPath = "cdxRPDGS7KTable"
    cdxRPDGS7KTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxRPDGS7KTable.EntityData.SegmentPath
    cdxRPDGS7KTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxRPDGS7KTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxRPDGS7KTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxRPDGS7KTable.EntityData.Children = types.NewOrderedMap()
    cdxRPDGS7KTable.EntityData.Children.Append("cdxRPDGS7KEntry", types.YChild{"CdxRPDGS7KEntry", nil})
    for i := range cdxRPDGS7KTable.CdxRPDGS7KEntry {
        cdxRPDGS7KTable.EntityData.Children.Append(types.GetSegmentPath(cdxRPDGS7KTable.CdxRPDGS7KEntry[i]), types.YChild{"CdxRPDGS7KEntry", cdxRPDGS7KTable.CdxRPDGS7KEntry[i]})
    }
    cdxRPDGS7KTable.EntityData.Leafs = types.NewOrderedMap()

    cdxRPDGS7KTable.EntityData.YListKeys = []string {}

    return &(cdxRPDGS7KTable.EntityData)
}

// CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry
// The list of statistics for all the sensor, 
// such as volatage, the state of TriSwitch.
type CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is MacAddress of RPDGS7K which is used for
    // index. The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdxRPDGS7KMacAddress interface{}

    // This is the Object of RPDGS7KPS1p24v. The type is interface{} with range:
    // 0..3000. Units are 0.01VDC.
    CdxRPDGS7KPS1p24v interface{}

    // This is the Object of RPDGS7KPS1p8v. The type is interface{} with range:
    // 0..1000. Units are 0.01VDC.
    CdxRPDGS7KPS1p8v interface{}

    // This is the Object of RPDGS7KPS1p5v. The type is interface{} with range:
    // 0..625. Units are 0.01VDC.
    CdxRPDGS7KPS1p5v interface{}

    // This is the Object of RPDGS7KPS1n6v. The type is interface{} with range:
    // 0..800. Units are 0.01VDC.
    CdxRPDGS7KPS1n6v interface{}

    // This is the Object of RPDGS7KPS1AC. The type is interface{} with range:
    // 300..2000. Units are 0.1VAC.
    CdxRPDGS7KPS1AC interface{}

    // RPDGS7KPS2p24v. The type is interface{} with range: 0..3000. Units are
    // 0.01VDC.
    CdxRPDGS7KPS2p24v interface{}

    // This is the Object of RPDGS7KPS2p8v. The type is interface{} with range:
    // 0..1000. Units are 0.01VDC.
    CdxRPDGS7KPS2p8v interface{}

    // This is the Object of RPDGS7KPS2p5v. The type is interface{} with range:
    // 0..625. Units are 0.01VDC.
    CdxRPDGS7KPS2p5v interface{}

    // This is the Object of RPDGS7KPS2n6v. The type is interface{} with range:
    // 0..800. Units are 0.01VDC.
    CdxRPDGS7KPS2n6v interface{}

    // This is the Object of RPDGS7KPS2AC. The type is interface{} with range:
    // 300..2000. Units are 0.1VAC.
    CdxRPDGS7KPS2AC interface{}

    // This is the Object of RPDGS7K Tx4 Opt Power. The type is interface{} with
    // range: 0..300. Units are 0.01mW.
    CdxRPDGS7KTx1OptPower interface{}

    // This is the Object of RPDGS7K Rx4 Opt Power. The type is interface{} with
    // range: 0..300. Units are 0.01mW.
    CdxRPDGS7KRx1OptPower interface{}

    // This is the Object of RPDGS7K TriSwitch, The relationship which the number
    // indicates is    low(1) for -6dB high(2) for 0dB pad(3) for off. The type is
    // CdxRPDGS7KTriSwitch.
    CdxRPDGS7KTriSwitch interface{}

    // This is the Object of RPDGS7K Tamp. The type is CdxRPDGS7KTamp.
    CdxRPDGS7KTamp interface{}
}

func (cdxRPDGS7KEntry *CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry) GetEntityData() *types.CommonEntityData {
    cdxRPDGS7KEntry.EntityData.YFilter = cdxRPDGS7KEntry.YFilter
    cdxRPDGS7KEntry.EntityData.YangName = "cdxRPDGS7KEntry"
    cdxRPDGS7KEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxRPDGS7KEntry.EntityData.ParentYangName = "cdxRPDGS7KTable"
    cdxRPDGS7KEntry.EntityData.SegmentPath = "cdxRPDGS7KEntry" + types.AddKeyToken(cdxRPDGS7KEntry.CdxRPDGS7KMacAddress, "cdxRPDGS7KMacAddress")
    cdxRPDGS7KEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxRPDGS7KTable/" + cdxRPDGS7KEntry.EntityData.SegmentPath
    cdxRPDGS7KEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxRPDGS7KEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxRPDGS7KEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxRPDGS7KEntry.EntityData.Children = types.NewOrderedMap()
    cdxRPDGS7KEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KMacAddress", types.YLeaf{"CdxRPDGS7KMacAddress", cdxRPDGS7KEntry.CdxRPDGS7KMacAddress})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS1p24v", types.YLeaf{"CdxRPDGS7KPS1p24v", cdxRPDGS7KEntry.CdxRPDGS7KPS1p24v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS1p8v", types.YLeaf{"CdxRPDGS7KPS1p8v", cdxRPDGS7KEntry.CdxRPDGS7KPS1p8v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS1p5v", types.YLeaf{"CdxRPDGS7KPS1p5v", cdxRPDGS7KEntry.CdxRPDGS7KPS1p5v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS1n6v", types.YLeaf{"CdxRPDGS7KPS1n6v", cdxRPDGS7KEntry.CdxRPDGS7KPS1n6v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS1AC", types.YLeaf{"CdxRPDGS7KPS1AC", cdxRPDGS7KEntry.CdxRPDGS7KPS1AC})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS2p24v", types.YLeaf{"CdxRPDGS7KPS2p24v", cdxRPDGS7KEntry.CdxRPDGS7KPS2p24v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS2p8v", types.YLeaf{"CdxRPDGS7KPS2p8v", cdxRPDGS7KEntry.CdxRPDGS7KPS2p8v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS2p5v", types.YLeaf{"CdxRPDGS7KPS2p5v", cdxRPDGS7KEntry.CdxRPDGS7KPS2p5v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS2n6v", types.YLeaf{"CdxRPDGS7KPS2n6v", cdxRPDGS7KEntry.CdxRPDGS7KPS2n6v})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KPS2AC", types.YLeaf{"CdxRPDGS7KPS2AC", cdxRPDGS7KEntry.CdxRPDGS7KPS2AC})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KTx1OptPower", types.YLeaf{"CdxRPDGS7KTx1OptPower", cdxRPDGS7KEntry.CdxRPDGS7KTx1OptPower})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KRx1OptPower", types.YLeaf{"CdxRPDGS7KRx1OptPower", cdxRPDGS7KEntry.CdxRPDGS7KRx1OptPower})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KTriSwitch", types.YLeaf{"CdxRPDGS7KTriSwitch", cdxRPDGS7KEntry.CdxRPDGS7KTriSwitch})
    cdxRPDGS7KEntry.EntityData.Leafs.Append("cdxRPDGS7KTamp", types.YLeaf{"CdxRPDGS7KTamp", cdxRPDGS7KEntry.CdxRPDGS7KTamp})

    cdxRPDGS7KEntry.EntityData.YListKeys = []string {"CdxRPDGS7KMacAddress"}

    return &(cdxRPDGS7KEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp represents This is the Object of RPDGS7K Tamp
type CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp string

const (
    CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp_intact CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp = "intact"

    CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp_compromised CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTamp = "compromised"
)

// CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch represents pad(3) for off
type CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch string

const (
    CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch_low CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch = "low"

    CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch_high CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch = "high"

    CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch_pad CISCODOCSEXTMIB_CdxRPDGS7KTable_CdxRPDGS7KEntry_CdxRPDGS7KTriSwitch = "pad"
)

// CISCODOCSEXTMIB_CdxBundleIpHelperTable
// A list of cable helper entries on Bundle/Sub-Bundle interface.
type CISCODOCSEXTMIB_CdxBundleIpHelperTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of cdxBundleIpHelperEntry. An instance exists for Cable
    // Bundle/Sub-Bundle Interface. The type is slice of
    // CISCODOCSEXTMIB_CdxBundleIpHelperTable_CdxBundleIpHelperEntry.
    CdxBundleIpHelperEntry []*CISCODOCSEXTMIB_CdxBundleIpHelperTable_CdxBundleIpHelperEntry
}

func (cdxBundleIpHelperTable *CISCODOCSEXTMIB_CdxBundleIpHelperTable) GetEntityData() *types.CommonEntityData {
    cdxBundleIpHelperTable.EntityData.YFilter = cdxBundleIpHelperTable.YFilter
    cdxBundleIpHelperTable.EntityData.YangName = "cdxBundleIpHelperTable"
    cdxBundleIpHelperTable.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIpHelperTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxBundleIpHelperTable.EntityData.SegmentPath = "cdxBundleIpHelperTable"
    cdxBundleIpHelperTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxBundleIpHelperTable.EntityData.SegmentPath
    cdxBundleIpHelperTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIpHelperTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIpHelperTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIpHelperTable.EntityData.Children = types.NewOrderedMap()
    cdxBundleIpHelperTable.EntityData.Children.Append("cdxBundleIpHelperEntry", types.YChild{"CdxBundleIpHelperEntry", nil})
    for i := range cdxBundleIpHelperTable.CdxBundleIpHelperEntry {
        cdxBundleIpHelperTable.EntityData.Children.Append(types.GetSegmentPath(cdxBundleIpHelperTable.CdxBundleIpHelperEntry[i]), types.YChild{"CdxBundleIpHelperEntry", cdxBundleIpHelperTable.CdxBundleIpHelperEntry[i]})
    }
    cdxBundleIpHelperTable.EntityData.Leafs = types.NewOrderedMap()

    cdxBundleIpHelperTable.EntityData.YListKeys = []string {}

    return &(cdxBundleIpHelperTable.EntityData)
}

// CISCODOCSEXTMIB_CdxBundleIpHelperTable_CdxBundleIpHelperEntry
// The conceptual row of cdxBundleIpHelperEntry.
// An instance exists for Cable Bundle/Sub-Bundle Interface.
type CISCODOCSEXTMIB_CdxBundleIpHelperTable_CdxBundleIpHelperEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Cable helper IP address. The type is string with
    // length: 0..255.
    CdxBundleHelperAddr interface{}

    // This object describes which kind of device will be associated to a cable
    // helper. The entity may support more than one device  class. For example,
    // the entity supports both host and mta. Therefore, bit 1 and bit 3 are set
    // to 1 for this object. If  all bits are cleared, the entity supports all
    // device types. Note that BITS are encoded most significant bit first. The
    // type is map[string]bool.
    CdxBundleHelperType interface{}
}

func (cdxBundleIpHelperEntry *CISCODOCSEXTMIB_CdxBundleIpHelperTable_CdxBundleIpHelperEntry) GetEntityData() *types.CommonEntityData {
    cdxBundleIpHelperEntry.EntityData.YFilter = cdxBundleIpHelperEntry.YFilter
    cdxBundleIpHelperEntry.EntityData.YangName = "cdxBundleIpHelperEntry"
    cdxBundleIpHelperEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIpHelperEntry.EntityData.ParentYangName = "cdxBundleIpHelperTable"
    cdxBundleIpHelperEntry.EntityData.SegmentPath = "cdxBundleIpHelperEntry" + types.AddKeyToken(cdxBundleIpHelperEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxBundleIpHelperEntry.CdxBundleHelperAddr, "cdxBundleHelperAddr")
    cdxBundleIpHelperEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxBundleIpHelperTable/" + cdxBundleIpHelperEntry.EntityData.SegmentPath
    cdxBundleIpHelperEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIpHelperEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIpHelperEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIpHelperEntry.EntityData.Children = types.NewOrderedMap()
    cdxBundleIpHelperEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxBundleIpHelperEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxBundleIpHelperEntry.IfIndex})
    cdxBundleIpHelperEntry.EntityData.Leafs.Append("cdxBundleHelperAddr", types.YLeaf{"CdxBundleHelperAddr", cdxBundleIpHelperEntry.CdxBundleHelperAddr})
    cdxBundleIpHelperEntry.EntityData.Leafs.Append("cdxBundleHelperType", types.YLeaf{"CdxBundleHelperType", cdxBundleIpHelperEntry.CdxBundleHelperType})

    cdxBundleIpHelperEntry.EntityData.YListKeys = []string {"IfIndex", "CdxBundleHelperAddr"}

    return &(cdxBundleIpHelperEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable
// Ipv6 dhcp relay configurations on Bundle/Sub-Bundle interface.
type CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of cdxBundleIPv6DHCPRelayTable. An instance exist for
    // the Bundle/Sub-Bundle interface. The type is slice of
    // CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable_CdxBundleIPv6DHCPRelayEntry.
    CdxBundleIPv6DHCPRelayEntry []*CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable_CdxBundleIPv6DHCPRelayEntry
}

func (cdxBundleIPv6DHCPRelayTable *CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable) GetEntityData() *types.CommonEntityData {
    cdxBundleIPv6DHCPRelayTable.EntityData.YFilter = cdxBundleIPv6DHCPRelayTable.YFilter
    cdxBundleIPv6DHCPRelayTable.EntityData.YangName = "cdxBundleIPv6DHCPRelayTable"
    cdxBundleIPv6DHCPRelayTable.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIPv6DHCPRelayTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxBundleIPv6DHCPRelayTable.EntityData.SegmentPath = "cdxBundleIPv6DHCPRelayTable"
    cdxBundleIPv6DHCPRelayTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxBundleIPv6DHCPRelayTable.EntityData.SegmentPath
    cdxBundleIPv6DHCPRelayTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIPv6DHCPRelayTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIPv6DHCPRelayTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIPv6DHCPRelayTable.EntityData.Children = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayTable.EntityData.Children.Append("cdxBundleIPv6DHCPRelayEntry", types.YChild{"CdxBundleIPv6DHCPRelayEntry", nil})
    for i := range cdxBundleIPv6DHCPRelayTable.CdxBundleIPv6DHCPRelayEntry {
        cdxBundleIPv6DHCPRelayTable.EntityData.Children.Append(types.GetSegmentPath(cdxBundleIPv6DHCPRelayTable.CdxBundleIPv6DHCPRelayEntry[i]), types.YChild{"CdxBundleIPv6DHCPRelayEntry", cdxBundleIPv6DHCPRelayTable.CdxBundleIPv6DHCPRelayEntry[i]})
    }
    cdxBundleIPv6DHCPRelayTable.EntityData.Leafs = types.NewOrderedMap()

    cdxBundleIPv6DHCPRelayTable.EntityData.YListKeys = []string {}

    return &(cdxBundleIPv6DHCPRelayTable.EntityData)
}

// CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable_CdxBundleIPv6DHCPRelayEntry
// The conceptual row of cdxBundleIPv6DHCPRelayTable.
// An instance exist for the Bundle/Sub-Bundle interface.
type CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable_CdxBundleIPv6DHCPRelayEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Insert VSS option in Relay-Forward Messages. The type is bool.
    CdxBundleIPv6DHCPRelayInsertVSSOption interface{}

    // Interface is trusted to process relay-replies. The type is bool.
    CdxBundleIPv6DHCPRelayTrustToRelayReply interface{}

    // Source interface name for IPv6 DHCP relayed messages. The type is string.
    CdxBundleIPv6DHDPRelaySourceIfname interface{}
}

func (cdxBundleIPv6DHCPRelayEntry *CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayTable_CdxBundleIPv6DHCPRelayEntry) GetEntityData() *types.CommonEntityData {
    cdxBundleIPv6DHCPRelayEntry.EntityData.YFilter = cdxBundleIPv6DHCPRelayEntry.YFilter
    cdxBundleIPv6DHCPRelayEntry.EntityData.YangName = "cdxBundleIPv6DHCPRelayEntry"
    cdxBundleIPv6DHCPRelayEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIPv6DHCPRelayEntry.EntityData.ParentYangName = "cdxBundleIPv6DHCPRelayTable"
    cdxBundleIPv6DHCPRelayEntry.EntityData.SegmentPath = "cdxBundleIPv6DHCPRelayEntry" + types.AddKeyToken(cdxBundleIPv6DHCPRelayEntry.IfIndex, "ifIndex")
    cdxBundleIPv6DHCPRelayEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxBundleIPv6DHCPRelayTable/" + cdxBundleIPv6DHCPRelayEntry.EntityData.SegmentPath
    cdxBundleIPv6DHCPRelayEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIPv6DHCPRelayEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIPv6DHCPRelayEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIPv6DHCPRelayEntry.EntityData.Children = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxBundleIPv6DHCPRelayEntry.IfIndex})
    cdxBundleIPv6DHCPRelayEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayInsertVSSOption", types.YLeaf{"CdxBundleIPv6DHCPRelayInsertVSSOption", cdxBundleIPv6DHCPRelayEntry.CdxBundleIPv6DHCPRelayInsertVSSOption})
    cdxBundleIPv6DHCPRelayEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayTrustToRelayReply", types.YLeaf{"CdxBundleIPv6DHCPRelayTrustToRelayReply", cdxBundleIPv6DHCPRelayEntry.CdxBundleIPv6DHCPRelayTrustToRelayReply})
    cdxBundleIPv6DHCPRelayEntry.EntityData.Leafs.Append("cdxBundleIPv6DHDPRelaySourceIfname", types.YLeaf{"CdxBundleIPv6DHDPRelaySourceIfname", cdxBundleIPv6DHCPRelayEntry.CdxBundleIPv6DHDPRelaySourceIfname})

    cdxBundleIPv6DHCPRelayEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdxBundleIPv6DHCPRelayEntry.EntityData)
}

// CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable
//  A list of IPv6 DHCP relay destination entries
// on Cable Bundle/Sub-Bundle interfaces.
type CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of cdxBundleIPv6DHCPRelayDestTable. An instance exists
    // for the Cable Bundle/Sub-Bundle interface. The type is slice of
    // CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable_CdxBundleIPv6DHCPRelayDestEntry.
    CdxBundleIPv6DHCPRelayDestEntry []*CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable_CdxBundleIPv6DHCPRelayDestEntry
}

func (cdxBundleIPv6DHCPRelayDestTable *CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable) GetEntityData() *types.CommonEntityData {
    cdxBundleIPv6DHCPRelayDestTable.EntityData.YFilter = cdxBundleIPv6DHCPRelayDestTable.YFilter
    cdxBundleIPv6DHCPRelayDestTable.EntityData.YangName = "cdxBundleIPv6DHCPRelayDestTable"
    cdxBundleIPv6DHCPRelayDestTable.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIPv6DHCPRelayDestTable.EntityData.ParentYangName = "CISCO-DOCS-EXT-MIB"
    cdxBundleIPv6DHCPRelayDestTable.EntityData.SegmentPath = "cdxBundleIPv6DHCPRelayDestTable"
    cdxBundleIPv6DHCPRelayDestTable.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/" + cdxBundleIPv6DHCPRelayDestTable.EntityData.SegmentPath
    cdxBundleIPv6DHCPRelayDestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIPv6DHCPRelayDestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIPv6DHCPRelayDestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIPv6DHCPRelayDestTable.EntityData.Children = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayDestTable.EntityData.Children.Append("cdxBundleIPv6DHCPRelayDestEntry", types.YChild{"CdxBundleIPv6DHCPRelayDestEntry", nil})
    for i := range cdxBundleIPv6DHCPRelayDestTable.CdxBundleIPv6DHCPRelayDestEntry {
        cdxBundleIPv6DHCPRelayDestTable.EntityData.Children.Append(types.GetSegmentPath(cdxBundleIPv6DHCPRelayDestTable.CdxBundleIPv6DHCPRelayDestEntry[i]), types.YChild{"CdxBundleIPv6DHCPRelayDestEntry", cdxBundleIPv6DHCPRelayDestTable.CdxBundleIPv6DHCPRelayDestEntry[i]})
    }
    cdxBundleIPv6DHCPRelayDestTable.EntityData.Leafs = types.NewOrderedMap()

    cdxBundleIPv6DHCPRelayDestTable.EntityData.YListKeys = []string {}

    return &(cdxBundleIPv6DHCPRelayDestTable.EntityData)
}

// CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable_CdxBundleIPv6DHCPRelayDestEntry
// The conceptual row of cdxBundleIPv6DHCPRelayDestTable.
// An instance exists for the Cable Bundle/Sub-Bundle interface.
type CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable_CdxBundleIPv6DHCPRelayDestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The vrf identifier that the
    // cdxBundleIPv6DHCPRelayDestOutIfIndex belongs to, it is assigned to each VRF
    // and is used to uniquely identify it, if it is zero, means in global vrf.
    // The type is interface{} with range: 0..65535.
    CdxBundleIPv6DHCPRelayDestOutIfVrfIndex interface{}

    // This attribute is a key. IPv6 DHCP relay destination address. The type is
    // string with length: 0..255.
    CdxBundleIPv6DHCPRelayDestAddr interface{}

    // This attribute is a key. The snmp ifIndex of the IPv6 DHCP relay
    // destination output interface. If the ifIndex is 0, it means there is no
    // output interface specified. The type is interface{} with range:
    // 0..2147483647.
    CdxBundleIPv6DHCPRelayDestOutIfIndex interface{}

    // IPv6 DHCP relay destination source address. The type is string.
    CdxBundleIPv6DHCPRelayDestSourceAddress interface{}

    // IPv6 DHCP relay destination link address. The type is string.
    CdxBundleIPv6DHCPRelayDestLinkAddress interface{}
}

func (cdxBundleIPv6DHCPRelayDestEntry *CISCODOCSEXTMIB_CdxBundleIPv6DHCPRelayDestTable_CdxBundleIPv6DHCPRelayDestEntry) GetEntityData() *types.CommonEntityData {
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.YFilter = cdxBundleIPv6DHCPRelayDestEntry.YFilter
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.YangName = "cdxBundleIPv6DHCPRelayDestEntry"
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.BundleName = "cisco_ios_xe"
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.ParentYangName = "cdxBundleIPv6DHCPRelayDestTable"
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.SegmentPath = "cdxBundleIPv6DHCPRelayDestEntry" + types.AddKeyToken(cdxBundleIPv6DHCPRelayDestEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestOutIfVrfIndex, "cdxBundleIPv6DHCPRelayDestOutIfVrfIndex") + types.AddKeyToken(cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestAddr, "cdxBundleIPv6DHCPRelayDestAddr") + types.AddKeyToken(cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestOutIfIndex, "cdxBundleIPv6DHCPRelayDestOutIfIndex")
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.AbsolutePath = "CISCO-DOCS-EXT-MIB:CISCO-DOCS-EXT-MIB/cdxBundleIPv6DHCPRelayDestTable/" + cdxBundleIPv6DHCPRelayDestEntry.EntityData.SegmentPath
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Children = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs = types.NewOrderedMap()
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdxBundleIPv6DHCPRelayDestEntry.IfIndex})
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayDestOutIfVrfIndex", types.YLeaf{"CdxBundleIPv6DHCPRelayDestOutIfVrfIndex", cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestOutIfVrfIndex})
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayDestAddr", types.YLeaf{"CdxBundleIPv6DHCPRelayDestAddr", cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestAddr})
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayDestOutIfIndex", types.YLeaf{"CdxBundleIPv6DHCPRelayDestOutIfIndex", cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestOutIfIndex})
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayDestSourceAddress", types.YLeaf{"CdxBundleIPv6DHCPRelayDestSourceAddress", cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestSourceAddress})
    cdxBundleIPv6DHCPRelayDestEntry.EntityData.Leafs.Append("cdxBundleIPv6DHCPRelayDestLinkAddress", types.YLeaf{"CdxBundleIPv6DHCPRelayDestLinkAddress", cdxBundleIPv6DHCPRelayDestEntry.CdxBundleIPv6DHCPRelayDestLinkAddress})

    cdxBundleIPv6DHCPRelayDestEntry.EntityData.YListKeys = []string {"IfIndex", "CdxBundleIPv6DHCPRelayDestOutIfVrfIndex", "CdxBundleIPv6DHCPRelayDestAddr", "CdxBundleIPv6DHCPRelayDestOutIfIndex"}

    return &(cdxBundleIPv6DHCPRelayDestEntry.EntityData)
}

