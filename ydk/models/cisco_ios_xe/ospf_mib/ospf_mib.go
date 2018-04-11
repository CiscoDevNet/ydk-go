// The MIB module to describe the OSPF Version 2
// Protocol.  Note that some objects in this MIB
// module may pose a significant security risk.
// Refer to the Security Considerations section
// in RFC 4750 for more information.
// 
// Copyright (C) The IETF Trust (2006).
// This version of this MIB module is part of
// RFC 4750;  see the RFC itself for full legal
// notices.
package ospf_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:OSPF-MIB OSPF-MIB}", reflect.TypeOf(OSPFMIB{}))
    ydk.RegisterEntity("OSPF-MIB:OSPF-MIB", reflect.TypeOf(OSPFMIB{}))
}

// Status represents and 'disabled' indicates that it is not.
type Status string

const (
    Status_enabled Status = "enabled"

    Status_disabled Status = "disabled"
)

// OspfAuthenticationType represents The authentication type.
type OspfAuthenticationType string

const (
    OspfAuthenticationType_none OspfAuthenticationType = "none"

    OspfAuthenticationType_simplePassword OspfAuthenticationType = "simplePassword"

    OspfAuthenticationType_md5 OspfAuthenticationType = "md5"
)

// OSPFMIB
type OSPFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ospfgeneralgroup OSPFMIB_Ospfgeneralgroup

    // Information describing the configured parameters and cumulative statistics
    // of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.
    Ospfareatable OSPFMIB_Ospfareatable

    // The set of metrics that will be advertised by a default Area Border Router
    // into a stub area.
    Ospfstubareatable OSPFMIB_Ospfstubareatable

    // The OSPF Process's link state database (LSDB). The LSDB contains the link
    // state advertisements from throughout the areas that the device is attached
    // to.
    Ospflsdbtable OSPFMIB_Ospflsdbtable

    // The Address Range Table acts as an adjunct to the Area Table.  It describes
    // those Address Range Summaries that are configured to be propagated from an
    // Area to reduce the amount of information about it that is known beyond its
    // borders.  It contains a set of IP address ranges specified by an IP
    // address/IP network mask pair. For example, class B address range of X.X.X.X
    // with a network mask of 255.255.0.0 includes all IP addresses from X.X.0.0
    // to X.X.255.255.  Note that this table is obsoleted and is replaced by the
    // Area Aggregate Table.
    Ospfarearangetable OSPFMIB_Ospfarearangetable

    // The Host/Metric Table indicates what hosts are directly  attached to the
    // router, what metrics and types of service should be advertised for them,
    // and what areas they are found within.
    Ospfhosttable OSPFMIB_Ospfhosttable

    // The OSPF Interface Table describes the interfaces from the viewpoint of
    // OSPF. It augments the ipAddrTable with OSPF specific information.
    Ospfiftable OSPFMIB_Ospfiftable

    // The Metric Table describes the metrics to be advertised for a specified
    // interface at the various types of service. As such, this table is an
    // adjunct of the OSPF Interface Table.  Types of service, as defined by RFC
    // 791, have the ability to request low delay, high bandwidth, or reliable
    // linkage.  For the purposes of this specification, the measure of bandwidth:
    // Metric = referenceBandwidth / ifSpeed  is the default value. The default
    // reference bandwidth is 10^8. For multiple link interfaces, note that
    // ifSpeed is the sum of the individual link speeds.  This yields a number
    // having the following typical values:  Network Type/bit rate   Metric  >=
    // 100 MBPS                 1 Ethernet/802.3             10 E1                
    // 48 T1 (ESF)                   65 64 KBPS                    1562 56 KBPS   
    // 1785 19.2 KBPS                  5208 9.6 KBPS                   10416 
    // Routes that are not specified use the default (TOS 0) metric.  Note that
    // the default reference bandwidth can be configured using the general group
    // object ospfReferenceBandwidth.
    Ospfifmetrictable OSPFMIB_Ospfifmetrictable

    // Information about this router's virtual interfaces that the OSPF Process is
    // configured to carry on.
    Ospfvirtiftable OSPFMIB_Ospfvirtiftable

    // A table describing all non-virtual neighbors in the locality of the OSPF
    // router.
    Ospfnbrtable OSPFMIB_Ospfnbrtable

    // This table describes all virtual neighbors. Since virtual links are
    // configured in the Virtual Interface Table, this table is read-only.
    Ospfvirtnbrtable OSPFMIB_Ospfvirtnbrtable

    // The OSPF Process's external LSA link state database.  This table is
    // identical to the OSPF LSDB Table in format, but contains only external link
    // state advertisements.  The purpose is to allow external  LSAs to be
    // displayed once for the router rather than once in each non-stub area.  Note
    // that external LSAs are also in the AS-scope link state database.
    Ospfextlsdbtable OSPFMIB_Ospfextlsdbtable

    // The Area Aggregate Table acts as an adjunct to the Area Table.  It
    // describes those address aggregates that are configured to be propagated
    // from an area. Its purpose is to reduce the amount of information that is
    // known beyond an Area's borders.  It contains a set of IP address ranges
    // specified by an IP address/IP network mask pair. For example, a class B
    // address range of X.X.X.X with a network mask of 255.255.0.0 includes all IP
    // addresses from X.X.0.0 to X.X.255.255.  Note that if ranges are configured
    // such that one range subsumes another range (e.g., 10.0.0.0 mask 255.0.0.0
    // and 10.1.0.0 mask 255.255.0.0), the most specific match is the preferred
    // one.
    Ospfareaaggregatetable OSPFMIB_Ospfareaaggregatetable

    // The OSPF Process's link-local link state database for non-virtual links.
    // This table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for non-virtual links.  The purpose is
    // to allow link-local LSAs to be displayed for each non-virtual interface. 
    // This table is implemented to support type-9 LSAs that are defined in 'The
    // OSPF Opaque LSA Option'.
    Ospflocallsdbtable OSPFMIB_Ospflocallsdbtable

    // The OSPF Process's link-local link state database for virtual links.  This
    // table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for virtual links.  The purpose is to
    // allow link-local LSAs to be displayed for each virtual interface.  This
    // table is implemented to support type-9 LSAs that are defined in 'The OSPF
    // Opaque LSA Option'.
    Ospfvirtlocallsdbtable OSPFMIB_Ospfvirtlocallsdbtable

    // The OSPF Process's AS-scope LSA link state database. The database contains
    // the AS-scope Link State Advertisements from throughout the areas that the
    // device is attached to.  This table is identical to the OSPF LSDB Table in
    // format, but contains only AS-scope Link State Advertisements.  The purpose
    // is to allow AS-scope LSAs to be displayed once for the router rather than
    // once in each non-stub area.
    Ospfaslsdbtable OSPFMIB_Ospfaslsdbtable

    // This table maintains per-area, per-LSA-type counters.
    Ospfarealsacounttable OSPFMIB_Ospfarealsacounttable
}

func (oSPFMIB *OSPFMIB) GetEntityData() *types.CommonEntityData {
    oSPFMIB.EntityData.YFilter = oSPFMIB.YFilter
    oSPFMIB.EntityData.YangName = "OSPF-MIB"
    oSPFMIB.EntityData.BundleName = "cisco_ios_xe"
    oSPFMIB.EntityData.ParentYangName = "OSPF-MIB"
    oSPFMIB.EntityData.SegmentPath = "OSPF-MIB:OSPF-MIB"
    oSPFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    oSPFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    oSPFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    oSPFMIB.EntityData.Children = make(map[string]types.YChild)
    oSPFMIB.EntityData.Children["ospfGeneralGroup"] = types.YChild{"Ospfgeneralgroup", &oSPFMIB.Ospfgeneralgroup}
    oSPFMIB.EntityData.Children["ospfAreaTable"] = types.YChild{"Ospfareatable", &oSPFMIB.Ospfareatable}
    oSPFMIB.EntityData.Children["ospfStubAreaTable"] = types.YChild{"Ospfstubareatable", &oSPFMIB.Ospfstubareatable}
    oSPFMIB.EntityData.Children["ospfLsdbTable"] = types.YChild{"Ospflsdbtable", &oSPFMIB.Ospflsdbtable}
    oSPFMIB.EntityData.Children["ospfAreaRangeTable"] = types.YChild{"Ospfarearangetable", &oSPFMIB.Ospfarearangetable}
    oSPFMIB.EntityData.Children["ospfHostTable"] = types.YChild{"Ospfhosttable", &oSPFMIB.Ospfhosttable}
    oSPFMIB.EntityData.Children["ospfIfTable"] = types.YChild{"Ospfiftable", &oSPFMIB.Ospfiftable}
    oSPFMIB.EntityData.Children["ospfIfMetricTable"] = types.YChild{"Ospfifmetrictable", &oSPFMIB.Ospfifmetrictable}
    oSPFMIB.EntityData.Children["ospfVirtIfTable"] = types.YChild{"Ospfvirtiftable", &oSPFMIB.Ospfvirtiftable}
    oSPFMIB.EntityData.Children["ospfNbrTable"] = types.YChild{"Ospfnbrtable", &oSPFMIB.Ospfnbrtable}
    oSPFMIB.EntityData.Children["ospfVirtNbrTable"] = types.YChild{"Ospfvirtnbrtable", &oSPFMIB.Ospfvirtnbrtable}
    oSPFMIB.EntityData.Children["ospfExtLsdbTable"] = types.YChild{"Ospfextlsdbtable", &oSPFMIB.Ospfextlsdbtable}
    oSPFMIB.EntityData.Children["ospfAreaAggregateTable"] = types.YChild{"Ospfareaaggregatetable", &oSPFMIB.Ospfareaaggregatetable}
    oSPFMIB.EntityData.Children["ospfLocalLsdbTable"] = types.YChild{"Ospflocallsdbtable", &oSPFMIB.Ospflocallsdbtable}
    oSPFMIB.EntityData.Children["ospfVirtLocalLsdbTable"] = types.YChild{"Ospfvirtlocallsdbtable", &oSPFMIB.Ospfvirtlocallsdbtable}
    oSPFMIB.EntityData.Children["ospfAsLsdbTable"] = types.YChild{"Ospfaslsdbtable", &oSPFMIB.Ospfaslsdbtable}
    oSPFMIB.EntityData.Children["ospfAreaLsaCountTable"] = types.YChild{"Ospfarealsacounttable", &oSPFMIB.Ospfarealsacounttable}
    oSPFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oSPFMIB.EntityData)
}

// OSPFMIB_Ospfgeneralgroup
type OSPFMIB_Ospfgeneralgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A 32-bit integer uniquely identifying the router in the Autonomous System.
    // By convention, to ensure uniqueness, this should default to the value of
    // one of the router's IP interface addresses.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfrouterid interface{}

    // The administrative status of OSPF in the router.  The value 'enabled'
    // denotes that the OSPF Process is active on at least one interface;
    // 'disabled' disables it on all interfaces.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is Status.
    Ospfadminstat interface{}

    // The current version number of the OSPF protocol is 2. The type is
    // Ospfversionnumber.
    Ospfversionnumber interface{}

    // A flag to note whether this router is an Area Border Router. The type is
    // bool.
    Ospfareabdrrtrstatus interface{}

    // A flag to note whether this router is configured as an Autonomous System
    // Border Router.  This object is persistent and when written the entity
    // SHOULD save the change to non-volatile storage. The type is bool.
    Ospfasbdrrtrstatus interface{}

    // The number of external (LS type-5) link state advertisements in the link
    // state database. The type is interface{} with range: 0..4294967295.
    Ospfexternlsacount interface{}

    // The 32-bit sum of the LS checksums of the external link state
    // advertisements contained in the link state database.  This sum can be used
    // to determine if there has been a change in a router's link state database
    // and to compare the link state database of two routers.  The value should be
    // treated as unsigned when comparing two sums of checksums. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfexternlsacksumsum interface{}

    // The router's support for type-of-service routing.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    Ospftossupport interface{}

    // The number of new link state advertisements that have been originated. 
    // This number is incremented each time the router originates a new LSA. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    Ospforiginatenewlsas interface{}

    // The number of link state advertisements received that are determined to be
    // new instantiations. This number does not include newer instantiations of
    // self-originated link state advertisements.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ospfDiscontinuityTime. The type
    // is interface{} with range: 0..4294967295.
    Ospfrxnewlsas interface{}

    // The maximum number of non-default AS-external LSAs entries that can be
    // stored in the link state database.  If the value is -1, then there is no
    // limit.  When the number of non-default AS-external LSAs in a router's link
    // state database reaches ospfExtLsdbLimit, the router enters overflow state. 
    // The router never holds more than ospfExtLsdbLimit non-default AS-external
    // LSAs in its database.  OspfExtLsdbLimit MUST be set identically in all
    // routers attached to the OSPF backbone and/or any regular OSPF area (i.e.,
    // OSPF stub areas and NSSAs are excluded).  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is interface{} with range: -1..2147483647.
    Ospfextlsdblimit interface{}

    // A bit mask indicating whether the router is forwarding IP multicast (Class
    // D) datagrams based on the algorithms defined in the multicast extensions to
    // OSPF.  Bit 0, if set, indicates that the router can  forward IP multicast
    // datagrams in the router's directly attached areas (called intra-area
    // multicast routing).  Bit 1, if set, indicates that the router can forward
    // IP multicast datagrams between OSPF areas (called inter-area multicast
    // routing).  Bit 2, if set, indicates that the router can forward IP
    // multicast datagrams between Autonomous Systems (called inter-AS multicast
    // routing).  Only certain combinations of bit settings are allowed, namely: 0
    // (no multicast forwarding is enabled), 1 (intra-area multicasting only), 3
    // (intra-area and inter-area multicasting), 5 (intra-area and inter-AS
    // multicasting), and 7 (multicasting everywhere).  By default, no multicast
    // forwarding is enabled.  This object is persistent and when written the
    // entity SHOULD save the change to non-volatile storage. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfmulticastextensions interface{}

    // The number of seconds that, after entering OverflowState, a router will
    // attempt to leave OverflowState.  This allows the router to again originate
    // non-default AS-external LSAs.  When set to 0, the router will not leave
    // overflow state until restarted.  This object is persistent and when written
    // the entity SHOULD save the change to non-volatile storage. The type is
    // interface{} with range: 0..2147483647.
    Ospfexitoverflowinterval interface{}

    // The router's support for demand routing. This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is bool.
    Ospfdemandextensions interface{}

    // Indicates metrics used to choose among multiple AS-external LSAs.  When
    // RFC1583Compatibility is set to enabled, only cost will be used when
    // choosing among multiple AS-external LSAs advertising the same destination. 
    // When RFC1583Compatibility is set to disabled, preference will be driven
    // first by type of path using cost only to break ties.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    Ospfrfc1583Compatibility interface{}

    // The router's support for Opaque LSA types. The type is bool.
    Ospfopaquelsasupport interface{}

    // Reference bandwidth in kilobits/second for  calculating default interface
    // metrics.  The default value is 100,000 KBPS (100 MBPS).  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 0..4294967295.
    // Units are kilobits per second.
    Ospfreferencebandwidth interface{}

    // The router's support for OSPF graceful restart. Options include: no restart
    // support, only planned restarts, or both planned and unplanned restarts. 
    // This object is persistent and when written the entity SHOULD save the
    // change to non-volatile storage. The type is Ospfrestartsupport.
    Ospfrestartsupport interface{}

    // Configured OSPF graceful restart timeout interval.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 1..1800. Units
    // are seconds.
    Ospfrestartinterval interface{}

    // Indicates if strict LSA checking is enabled for graceful restart.  This
    // object is persistent and when written the entity SHOULD save the change to
    // non-volatile  storage. The type is bool.
    Ospfrestartstrictlsachecking interface{}

    // Current status of OSPF graceful restart. The type is Ospfrestartstatus.
    Ospfrestartstatus interface{}

    // Remaining time in current OSPF graceful restart interval. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Ospfrestartage interface{}

    // Describes the outcome of the last attempt at a graceful restart.  If the
    // value is 'none', no restart has yet been attempted.  If the value is
    // 'inProgress', a restart attempt is currently underway. The type is
    // Ospfrestartexitreason.
    Ospfrestartexitreason interface{}

    // The number of AS-scope link state advertisements in the AS-scope link state
    // database. The type is interface{} with range: 0..4294967295.
    Ospfaslsacount interface{}

    // The 32-bit unsigned sum of the LS checksums of the AS link state
    // advertisements contained in the AS-scope link state database.  This sum can
    // be used to determine if there has been a change in a router's AS-scope link
    // state database, and to compare the AS-scope link state database of two
    // routers. The type is interface{} with range: 0..4294967295.
    Ospfaslsacksumsum interface{}

    // The router's support for stub router functionality. The type is bool.
    Ospfstubroutersupport interface{}

    // This object controls the advertisement of stub router LSAs by the router. 
    // The value doNotAdvertise will result in the advertisement of a standard
    // router LSA and is the default value.  This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is Ospfstubrouteradvertisement.
    Ospfstubrouteradvertisement interface{}

    // The value of sysUpTime on the most recent occasion at which any one of this
    // MIB's counters suffered a discontinuity.  If no such discontinuities have
    // occurred since the last re-initialization of the local management
    // subsystem, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    Ospfdiscontinuitytime interface{}
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetEntityData() *types.CommonEntityData {
    ospfgeneralgroup.EntityData.YFilter = ospfgeneralgroup.YFilter
    ospfgeneralgroup.EntityData.YangName = "ospfGeneralGroup"
    ospfgeneralgroup.EntityData.BundleName = "cisco_ios_xe"
    ospfgeneralgroup.EntityData.ParentYangName = "OSPF-MIB"
    ospfgeneralgroup.EntityData.SegmentPath = "ospfGeneralGroup"
    ospfgeneralgroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfgeneralgroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfgeneralgroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfgeneralgroup.EntityData.Children = make(map[string]types.YChild)
    ospfgeneralgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfgeneralgroup.EntityData.Leafs["ospfRouterId"] = types.YLeaf{"Ospfrouterid", ospfgeneralgroup.Ospfrouterid}
    ospfgeneralgroup.EntityData.Leafs["ospfAdminStat"] = types.YLeaf{"Ospfadminstat", ospfgeneralgroup.Ospfadminstat}
    ospfgeneralgroup.EntityData.Leafs["ospfVersionNumber"] = types.YLeaf{"Ospfversionnumber", ospfgeneralgroup.Ospfversionnumber}
    ospfgeneralgroup.EntityData.Leafs["ospfAreaBdrRtrStatus"] = types.YLeaf{"Ospfareabdrrtrstatus", ospfgeneralgroup.Ospfareabdrrtrstatus}
    ospfgeneralgroup.EntityData.Leafs["ospfASBdrRtrStatus"] = types.YLeaf{"Ospfasbdrrtrstatus", ospfgeneralgroup.Ospfasbdrrtrstatus}
    ospfgeneralgroup.EntityData.Leafs["ospfExternLsaCount"] = types.YLeaf{"Ospfexternlsacount", ospfgeneralgroup.Ospfexternlsacount}
    ospfgeneralgroup.EntityData.Leafs["ospfExternLsaCksumSum"] = types.YLeaf{"Ospfexternlsacksumsum", ospfgeneralgroup.Ospfexternlsacksumsum}
    ospfgeneralgroup.EntityData.Leafs["ospfTOSSupport"] = types.YLeaf{"Ospftossupport", ospfgeneralgroup.Ospftossupport}
    ospfgeneralgroup.EntityData.Leafs["ospfOriginateNewLsas"] = types.YLeaf{"Ospforiginatenewlsas", ospfgeneralgroup.Ospforiginatenewlsas}
    ospfgeneralgroup.EntityData.Leafs["ospfRxNewLsas"] = types.YLeaf{"Ospfrxnewlsas", ospfgeneralgroup.Ospfrxnewlsas}
    ospfgeneralgroup.EntityData.Leafs["ospfExtLsdbLimit"] = types.YLeaf{"Ospfextlsdblimit", ospfgeneralgroup.Ospfextlsdblimit}
    ospfgeneralgroup.EntityData.Leafs["ospfMulticastExtensions"] = types.YLeaf{"Ospfmulticastextensions", ospfgeneralgroup.Ospfmulticastextensions}
    ospfgeneralgroup.EntityData.Leafs["ospfExitOverflowInterval"] = types.YLeaf{"Ospfexitoverflowinterval", ospfgeneralgroup.Ospfexitoverflowinterval}
    ospfgeneralgroup.EntityData.Leafs["ospfDemandExtensions"] = types.YLeaf{"Ospfdemandextensions", ospfgeneralgroup.Ospfdemandextensions}
    ospfgeneralgroup.EntityData.Leafs["ospfRFC1583Compatibility"] = types.YLeaf{"Ospfrfc1583Compatibility", ospfgeneralgroup.Ospfrfc1583Compatibility}
    ospfgeneralgroup.EntityData.Leafs["ospfOpaqueLsaSupport"] = types.YLeaf{"Ospfopaquelsasupport", ospfgeneralgroup.Ospfopaquelsasupport}
    ospfgeneralgroup.EntityData.Leafs["ospfReferenceBandwidth"] = types.YLeaf{"Ospfreferencebandwidth", ospfgeneralgroup.Ospfreferencebandwidth}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartSupport"] = types.YLeaf{"Ospfrestartsupport", ospfgeneralgroup.Ospfrestartsupport}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartInterval"] = types.YLeaf{"Ospfrestartinterval", ospfgeneralgroup.Ospfrestartinterval}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartStrictLsaChecking"] = types.YLeaf{"Ospfrestartstrictlsachecking", ospfgeneralgroup.Ospfrestartstrictlsachecking}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartStatus"] = types.YLeaf{"Ospfrestartstatus", ospfgeneralgroup.Ospfrestartstatus}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartAge"] = types.YLeaf{"Ospfrestartage", ospfgeneralgroup.Ospfrestartage}
    ospfgeneralgroup.EntityData.Leafs["ospfRestartExitReason"] = types.YLeaf{"Ospfrestartexitreason", ospfgeneralgroup.Ospfrestartexitreason}
    ospfgeneralgroup.EntityData.Leafs["ospfAsLsaCount"] = types.YLeaf{"Ospfaslsacount", ospfgeneralgroup.Ospfaslsacount}
    ospfgeneralgroup.EntityData.Leafs["ospfAsLsaCksumSum"] = types.YLeaf{"Ospfaslsacksumsum", ospfgeneralgroup.Ospfaslsacksumsum}
    ospfgeneralgroup.EntityData.Leafs["ospfStubRouterSupport"] = types.YLeaf{"Ospfstubroutersupport", ospfgeneralgroup.Ospfstubroutersupport}
    ospfgeneralgroup.EntityData.Leafs["ospfStubRouterAdvertisement"] = types.YLeaf{"Ospfstubrouteradvertisement", ospfgeneralgroup.Ospfstubrouteradvertisement}
    ospfgeneralgroup.EntityData.Leafs["ospfDiscontinuityTime"] = types.YLeaf{"Ospfdiscontinuitytime", ospfgeneralgroup.Ospfdiscontinuitytime}
    return &(ospfgeneralgroup.EntityData)
}

// OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason represents a restart attempt is currently underway.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_none OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "none"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_inProgress OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "inProgress"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_completed OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "completed"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_timedOut OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "timedOut"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_topologyChanged OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "topologyChanged"
)

// OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus represents Current status of OSPF graceful restart.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_notRestarting OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "notRestarting"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_plannedRestart OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "plannedRestart"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_unplannedRestart OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "unplannedRestart"
)

// OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport represents storage.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_none OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "none"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_plannedOnly OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "plannedOnly"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_plannedAndUnplanned OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "plannedAndUnplanned"
)

// OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement represents storage.
type OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement_doNotAdvertise OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement = "doNotAdvertise"

    OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement_advertise OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement = "advertise"
)

// OSPFMIB_Ospfgeneralgroup_Ospfversionnumber represents The current version number of the OSPF protocol is 2.
type OSPFMIB_Ospfgeneralgroup_Ospfversionnumber string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfversionnumber_version2 OSPFMIB_Ospfgeneralgroup_Ospfversionnumber = "version2"
)

// OSPFMIB_Ospfareatable
// Information describing the configured parameters and
// cumulative statistics of the router's attached areas.
// The interfaces and virtual links are configured
// as part of these areas.  Area 0.0.0.0, by definition,
// is the backbone area.
type OSPFMIB_Ospfareatable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information describing the configured parameters and cumulative statistics
    // of one of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_Ospfareatable_Ospfareaentry.
    Ospfareaentry []OSPFMIB_Ospfareatable_Ospfareaentry
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetEntityData() *types.CommonEntityData {
    ospfareatable.EntityData.YFilter = ospfareatable.YFilter
    ospfareatable.EntityData.YangName = "ospfAreaTable"
    ospfareatable.EntityData.BundleName = "cisco_ios_xe"
    ospfareatable.EntityData.ParentYangName = "OSPF-MIB"
    ospfareatable.EntityData.SegmentPath = "ospfAreaTable"
    ospfareatable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfareatable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfareatable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfareatable.EntityData.Children = make(map[string]types.YChild)
    ospfareatable.EntityData.Children["ospfAreaEntry"] = types.YChild{"Ospfareaentry", nil}
    for i := range ospfareatable.Ospfareaentry {
        ospfareatable.EntityData.Children[types.GetSegmentPath(&ospfareatable.Ospfareaentry[i])] = types.YChild{"Ospfareaentry", &ospfareatable.Ospfareaentry[i]}
    }
    ospfareatable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfareatable.EntityData)
}

// OSPFMIB_Ospfareatable_Ospfareaentry
// Information describing the configured parameters and
// cumulative statistics of one of the router's attached areas.
// The interfaces and virtual links are configured as part of
// these areas.  Area 0.0.0.0, by definition, is the backbone
// area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfareatable_Ospfareaentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A 32-bit integer uniquely identifying an area.
    // Area ID 0.0.0.0 is used for the OSPF backbone. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfareaid interface{}

    // The authentication type specified for an area. The type is
    // OspfAuthenticationType.
    Ospfauthtype interface{}

    // Indicates if an area is a stub area, NSSA, or standard area.  Type-5
    // AS-external LSAs and type-11 Opaque LSAs are not imported into stub areas
    // or NSSAs.  NSSAs import AS-external data as type-7 LSAs. The type is
    // Ospfimportasextern.
    Ospfimportasextern interface{}

    // The number of times that the intra-area route table has been calculated
    // using this area's link state database.  This is typically done using
    // Dijkstra's algorithm.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfspfruns interface{}

    // The total number of Area Border Routers reachable within this area.  This
    // is initially zero and is calculated in each Shortest Path First (SPF) pass.
    // The type is interface{} with range: 0..4294967295.
    Ospfareabdrrtrcount interface{}

    // The total number of Autonomous System Border Routers reachable within this
    // area.  This is initially zero and is calculated in each SPF pass. The type
    // is interface{} with range: 0..4294967295.
    Ospfasbdrrtrcount interface{}

    // The total number of link state advertisements in this area's link state
    // database, excluding AS-external LSAs. The type is interface{} with range:
    // 0..4294967295.
    Ospfarealsacount interface{}

    // The 32-bit sum of the link state advertisements' LS checksums contained in
    // this area's link state database.  This sum excludes external (LS type-5)
    // link state advertisements. The sum can be used to determine if there has
    // been a change in a router's link state database, and to compare the link
    // state database of two routers.  The value should be treated as unsigned
    // when comparing two sums of checksums. The type is interface{} with range:
    // -2147483648..2147483647.
    Ospfarealsacksumsum interface{}

    // The variable ospfAreaSummary controls the import of summary LSAs into stub
    // and NSSA areas. It has no effect on other areas.  If it is noAreaSummary,
    // the router will not originate summary LSAs into the stub or NSSA area. It
    // will rely entirely on its default route.  If it is sendAreaSummary, the
    // router will both summarize and propagate summary LSAs. The type is
    // Ospfareasummary.
    Ospfareasummary interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfareastatus interface{}

    // Indicates an NSSA border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is Ospfareanssatranslatorrole.
    Ospfareanssatranslatorrole interface{}

    // Indicates if and how an NSSA border router is performing NSSA translation
    // of type-7 LSAs into type-5  LSAs.  When this object is set to enabled, the
    // NSSA Border router's OspfAreaNssaExtTranslatorRole has been set to always. 
    // When this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA border router is NOT translating type-7 LSAs into type-5.
    // The type is Ospfareanssatranslatorstate.
    Ospfareanssatranslatorstate interface{}

    // The number of seconds after an elected translator determines its services
    // are no longer required, that it should continue to perform its translation
    // duties. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    Ospfareanssatranslatorstabilityinterval interface{}

    // Indicates the number of translator state changes that have occurred since
    // the last boot-up.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfareanssatranslatorevents interface{}

    // The total number of Opaque Area and AS link-state  advertisements in the
    // link state database of this area. The type is interface{} with range:
    // 0..4294967295.
    Cospfopaquearealsacount interface{}

    // The 32-bit unsigned sum of the Opaque Area and AS  link-state
    // advertisements' LS checksums contained  link state database of this area. 
    // The sum can be  used to determine if there has been a change in the  link
    // state database for Opaque Area and AS link-state advertisements. The type
    // is interface{} with range: 0..4294967295.
    Cospfopaquearealsacksumsum interface{}

    // Indicates an NSSA Border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is Cospfareanssatranslatorrole.
    Cospfareanssatranslatorrole interface{}

    // Indicates if and how an NSSA Border router is performing NSSA translation
    // of type-7 LSAs into type-5 LSAs. When this object set to enabled, the NSSA
    // Border router's cospfAreaNssaExtTranslatorRole has been set to always. When
    // this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA Border router is NOT translating type-7 LSAs into type-5.
    // The type is Cospfareanssatranslatorstate.
    Cospfareanssatranslatorstate interface{}

    // Indicates the number of Translator State changes that have occurred since
    // the last boot-up. The type is interface{} with range: 0..4294967295.
    Cospfareanssatranslatorevents interface{}
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetEntityData() *types.CommonEntityData {
    ospfareaentry.EntityData.YFilter = ospfareaentry.YFilter
    ospfareaentry.EntityData.YangName = "ospfAreaEntry"
    ospfareaentry.EntityData.BundleName = "cisco_ios_xe"
    ospfareaentry.EntityData.ParentYangName = "ospfAreaTable"
    ospfareaentry.EntityData.SegmentPath = "ospfAreaEntry" + "[ospfAreaId='" + fmt.Sprintf("%v", ospfareaentry.Ospfareaid) + "']"
    ospfareaentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfareaentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfareaentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfareaentry.EntityData.Children = make(map[string]types.YChild)
    ospfareaentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfareaentry.EntityData.Leafs["ospfAreaId"] = types.YLeaf{"Ospfareaid", ospfareaentry.Ospfareaid}
    ospfareaentry.EntityData.Leafs["ospfAuthType"] = types.YLeaf{"Ospfauthtype", ospfareaentry.Ospfauthtype}
    ospfareaentry.EntityData.Leafs["ospfImportAsExtern"] = types.YLeaf{"Ospfimportasextern", ospfareaentry.Ospfimportasextern}
    ospfareaentry.EntityData.Leafs["ospfSpfRuns"] = types.YLeaf{"Ospfspfruns", ospfareaentry.Ospfspfruns}
    ospfareaentry.EntityData.Leafs["ospfAreaBdrRtrCount"] = types.YLeaf{"Ospfareabdrrtrcount", ospfareaentry.Ospfareabdrrtrcount}
    ospfareaentry.EntityData.Leafs["ospfAsBdrRtrCount"] = types.YLeaf{"Ospfasbdrrtrcount", ospfareaentry.Ospfasbdrrtrcount}
    ospfareaentry.EntityData.Leafs["ospfAreaLsaCount"] = types.YLeaf{"Ospfarealsacount", ospfareaentry.Ospfarealsacount}
    ospfareaentry.EntityData.Leafs["ospfAreaLsaCksumSum"] = types.YLeaf{"Ospfarealsacksumsum", ospfareaentry.Ospfarealsacksumsum}
    ospfareaentry.EntityData.Leafs["ospfAreaSummary"] = types.YLeaf{"Ospfareasummary", ospfareaentry.Ospfareasummary}
    ospfareaentry.EntityData.Leafs["ospfAreaStatus"] = types.YLeaf{"Ospfareastatus", ospfareaentry.Ospfareastatus}
    ospfareaentry.EntityData.Leafs["ospfAreaNssaTranslatorRole"] = types.YLeaf{"Ospfareanssatranslatorrole", ospfareaentry.Ospfareanssatranslatorrole}
    ospfareaentry.EntityData.Leafs["ospfAreaNssaTranslatorState"] = types.YLeaf{"Ospfareanssatranslatorstate", ospfareaentry.Ospfareanssatranslatorstate}
    ospfareaentry.EntityData.Leafs["ospfAreaNssaTranslatorStabilityInterval"] = types.YLeaf{"Ospfareanssatranslatorstabilityinterval", ospfareaentry.Ospfareanssatranslatorstabilityinterval}
    ospfareaentry.EntityData.Leafs["ospfAreaNssaTranslatorEvents"] = types.YLeaf{"Ospfareanssatranslatorevents", ospfareaentry.Ospfareanssatranslatorevents}
    ospfareaentry.EntityData.Leafs["cospfOpaqueAreaLsaCount"] = types.YLeaf{"Cospfopaquearealsacount", ospfareaentry.Cospfopaquearealsacount}
    ospfareaentry.EntityData.Leafs["cospfOpaqueAreaLsaCksumSum"] = types.YLeaf{"Cospfopaquearealsacksumsum", ospfareaentry.Cospfopaquearealsacksumsum}
    ospfareaentry.EntityData.Leafs["cospfAreaNssaTranslatorRole"] = types.YLeaf{"Cospfareanssatranslatorrole", ospfareaentry.Cospfareanssatranslatorrole}
    ospfareaentry.EntityData.Leafs["cospfAreaNssaTranslatorState"] = types.YLeaf{"Cospfareanssatranslatorstate", ospfareaentry.Cospfareanssatranslatorstate}
    ospfareaentry.EntityData.Leafs["cospfAreaNssaTranslatorEvents"] = types.YLeaf{"Cospfareanssatranslatorevents", ospfareaentry.Cospfareanssatranslatorevents}
    return &(ospfareaentry.EntityData)
}

// OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole represents type-5 LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole_always OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole = "always"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole_candidate OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole = "candidate"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate represents Border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_enabled OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "enabled"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_elected OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "elected"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_disabled OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "disabled"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole represents type-5 LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole_always OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole = "always"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole_candidate OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole = "candidate"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate represents border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_enabled OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "enabled"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_elected OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "elected"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_disabled OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "disabled"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary represents summarize and propagate summary LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary_noAreaSummary OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary = "noAreaSummary"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary_sendAreaSummary OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary = "sendAreaSummary"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern represents AS-external data as type-7 LSAs
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importExternal OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importExternal"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importNoExternal OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importNoExternal"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importNssa OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importNssa"
)

// OSPFMIB_Ospfstubareatable
// The set of metrics that will be advertised
// by a default Area Border Router into a stub area.
type OSPFMIB_Ospfstubareatable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The metric for a given Type of Service that will be advertised by a default
    // Area Border Router into a stub area.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_Ospfstubareatable_Ospfstubareaentry.
    Ospfstubareaentry []OSPFMIB_Ospfstubareatable_Ospfstubareaentry
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetEntityData() *types.CommonEntityData {
    ospfstubareatable.EntityData.YFilter = ospfstubareatable.YFilter
    ospfstubareatable.EntityData.YangName = "ospfStubAreaTable"
    ospfstubareatable.EntityData.BundleName = "cisco_ios_xe"
    ospfstubareatable.EntityData.ParentYangName = "OSPF-MIB"
    ospfstubareatable.EntityData.SegmentPath = "ospfStubAreaTable"
    ospfstubareatable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfstubareatable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfstubareatable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfstubareatable.EntityData.Children = make(map[string]types.YChild)
    ospfstubareatable.EntityData.Children["ospfStubAreaEntry"] = types.YChild{"Ospfstubareaentry", nil}
    for i := range ospfstubareatable.Ospfstubareaentry {
        ospfstubareatable.EntityData.Children[types.GetSegmentPath(&ospfstubareatable.Ospfstubareaentry[i])] = types.YChild{"Ospfstubareaentry", &ospfstubareatable.Ospfstubareaentry[i]}
    }
    ospfstubareatable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfstubareatable.EntityData)
}

// OSPFMIB_Ospfstubareatable_Ospfstubareaentry
// The metric for a given Type of Service that
// will be advertised by a default Area Border
// Router into a stub area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfstubareatable_Ospfstubareaentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier for the stub area.  On
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfstubareaid interface{}

    // This attribute is a key. The Type of Service associated with the metric. 
    // On creation, this can be derived from  the instance. The type is
    // interface{} with range: 0..30.
    Ospfstubtos interface{}

    // The metric value applied at the indicated Type of Service.  By default,
    // this equals the least metric at the Type of Service among the interfaces to
    // other areas. The type is interface{} with range: 0..16777215.
    Ospfstubmetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfstubstatus interface{}

    // This variable displays the type of metric advertised as a default route.
    // The type is Ospfstubmetrictype.
    Ospfstubmetrictype interface{}
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetEntityData() *types.CommonEntityData {
    ospfstubareaentry.EntityData.YFilter = ospfstubareaentry.YFilter
    ospfstubareaentry.EntityData.YangName = "ospfStubAreaEntry"
    ospfstubareaentry.EntityData.BundleName = "cisco_ios_xe"
    ospfstubareaentry.EntityData.ParentYangName = "ospfStubAreaTable"
    ospfstubareaentry.EntityData.SegmentPath = "ospfStubAreaEntry" + "[ospfStubAreaId='" + fmt.Sprintf("%v", ospfstubareaentry.Ospfstubareaid) + "']" + "[ospfStubTOS='" + fmt.Sprintf("%v", ospfstubareaentry.Ospfstubtos) + "']"
    ospfstubareaentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfstubareaentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfstubareaentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfstubareaentry.EntityData.Children = make(map[string]types.YChild)
    ospfstubareaentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfstubareaentry.EntityData.Leafs["ospfStubAreaId"] = types.YLeaf{"Ospfstubareaid", ospfstubareaentry.Ospfstubareaid}
    ospfstubareaentry.EntityData.Leafs["ospfStubTOS"] = types.YLeaf{"Ospfstubtos", ospfstubareaentry.Ospfstubtos}
    ospfstubareaentry.EntityData.Leafs["ospfStubMetric"] = types.YLeaf{"Ospfstubmetric", ospfstubareaentry.Ospfstubmetric}
    ospfstubareaentry.EntityData.Leafs["ospfStubStatus"] = types.YLeaf{"Ospfstubstatus", ospfstubareaentry.Ospfstubstatus}
    ospfstubareaentry.EntityData.Leafs["ospfStubMetricType"] = types.YLeaf{"Ospfstubmetrictype", ospfstubareaentry.Ospfstubmetrictype}
    return &(ospfstubareaentry.EntityData)
}

// OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype represents advertised as a default route.
type OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype string

const (
    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_ospfMetric OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "ospfMetric"

    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_comparableCost OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "comparableCost"

    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_nonComparable OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "nonComparable"
)

// OSPFMIB_Ospflsdbtable
// The OSPF Process's link state database (LSDB).
// The LSDB contains the link state advertisements
// from throughout the areas that the device is attached to.
type OSPFMIB_Ospflsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospflsdbtable_Ospflsdbentry.
    Ospflsdbentry []OSPFMIB_Ospflsdbtable_Ospflsdbentry
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetEntityData() *types.CommonEntityData {
    ospflsdbtable.EntityData.YFilter = ospflsdbtable.YFilter
    ospflsdbtable.EntityData.YangName = "ospfLsdbTable"
    ospflsdbtable.EntityData.BundleName = "cisco_ios_xe"
    ospflsdbtable.EntityData.ParentYangName = "OSPF-MIB"
    ospflsdbtable.EntityData.SegmentPath = "ospfLsdbTable"
    ospflsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospflsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospflsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospflsdbtable.EntityData.Children = make(map[string]types.YChild)
    ospflsdbtable.EntityData.Children["ospfLsdbEntry"] = types.YChild{"Ospflsdbentry", nil}
    for i := range ospflsdbtable.Ospflsdbentry {
        ospflsdbtable.EntityData.Children[types.GetSegmentPath(&ospflsdbtable.Ospflsdbentry[i])] = types.YChild{"Ospflsdbentry", &ospflsdbtable.Ospflsdbentry[i]}
    }
    ospflsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospflsdbtable.EntityData)
}

// OSPFMIB_Ospflsdbtable_Ospflsdbentry
// A single link state advertisement.
type OSPFMIB_Ospflsdbtable_Ospflsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier of the area from which the
    // LSA was received. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflsdbareaid interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format.  Note: External link
    // state advertisements are permitted for backward compatibility, but should
    // be displayed in the ospfAsLsdbTable rather than here. The type is
    // Ospflsdbtype.
    Ospflsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate Link State Advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospflsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospflsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless  datagrams;
    // it is commonly referred to as the Fletcher checksum. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospflsdbchecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    Ospflsdbadvertisement interface{}
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetEntityData() *types.CommonEntityData {
    ospflsdbentry.EntityData.YFilter = ospflsdbentry.YFilter
    ospflsdbentry.EntityData.YangName = "ospfLsdbEntry"
    ospflsdbentry.EntityData.BundleName = "cisco_ios_xe"
    ospflsdbentry.EntityData.ParentYangName = "ospfLsdbTable"
    ospflsdbentry.EntityData.SegmentPath = "ospfLsdbEntry" + "[ospfLsdbAreaId='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbareaid) + "']" + "[ospfLsdbType='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbtype) + "']" + "[ospfLsdbLsid='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdblsid) + "']" + "[ospfLsdbRouterId='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbrouterid) + "']"
    ospflsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospflsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospflsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospflsdbentry.EntityData.Children = make(map[string]types.YChild)
    ospflsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospflsdbentry.EntityData.Leafs["ospfLsdbAreaId"] = types.YLeaf{"Ospflsdbareaid", ospflsdbentry.Ospflsdbareaid}
    ospflsdbentry.EntityData.Leafs["ospfLsdbType"] = types.YLeaf{"Ospflsdbtype", ospflsdbentry.Ospflsdbtype}
    ospflsdbentry.EntityData.Leafs["ospfLsdbLsid"] = types.YLeaf{"Ospflsdblsid", ospflsdbentry.Ospflsdblsid}
    ospflsdbentry.EntityData.Leafs["ospfLsdbRouterId"] = types.YLeaf{"Ospflsdbrouterid", ospflsdbentry.Ospflsdbrouterid}
    ospflsdbentry.EntityData.Leafs["ospfLsdbSequence"] = types.YLeaf{"Ospflsdbsequence", ospflsdbentry.Ospflsdbsequence}
    ospflsdbentry.EntityData.Leafs["ospfLsdbAge"] = types.YLeaf{"Ospflsdbage", ospflsdbentry.Ospflsdbage}
    ospflsdbentry.EntityData.Leafs["ospfLsdbChecksum"] = types.YLeaf{"Ospflsdbchecksum", ospflsdbentry.Ospflsdbchecksum}
    ospflsdbentry.EntityData.Leafs["ospfLsdbAdvertisement"] = types.YLeaf{"Ospflsdbadvertisement", ospflsdbentry.Ospflsdbadvertisement}
    return &(ospflsdbentry.EntityData)
}

// OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype represents in the ospfAsLsdbTable rather than here.
type OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype string

const (
    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_routerLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "routerLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_networkLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "networkLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_summaryLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "summaryLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_asSummaryLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "asSummaryLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_asExternalLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "asExternalLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_multicastLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "multicastLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_nssaExternalLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "nssaExternalLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_areaOpaqueLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "areaOpaqueLink"
)

// OSPFMIB_Ospfarearangetable
// The Address Range Table acts as an adjunct to the Area
// Table.  It describes those Address Range Summaries that
// are configured to be propagated from an Area to reduce
// the amount of information about it that is known beyond
// its borders.  It contains a set of IP address ranges
// specified by an IP address/IP network mask pair.
// For example, class B address range of X.X.X.X
// with a network mask of 255.255.0.0 includes all IP
// addresses from X.X.0.0 to X.X.255.255.
// 
// Note that this table is obsoleted and is replaced
// by the Area Aggregate Table.
type OSPFMIB_Ospfarearangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single area address range.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_Ospfarearangetable_Ospfarearangeentry.
    Ospfarearangeentry []OSPFMIB_Ospfarearangetable_Ospfarearangeentry
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetEntityData() *types.CommonEntityData {
    ospfarearangetable.EntityData.YFilter = ospfarearangetable.YFilter
    ospfarearangetable.EntityData.YangName = "ospfAreaRangeTable"
    ospfarearangetable.EntityData.BundleName = "cisco_ios_xe"
    ospfarearangetable.EntityData.ParentYangName = "OSPF-MIB"
    ospfarearangetable.EntityData.SegmentPath = "ospfAreaRangeTable"
    ospfarearangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfarearangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfarearangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfarearangetable.EntityData.Children = make(map[string]types.YChild)
    ospfarearangetable.EntityData.Children["ospfAreaRangeEntry"] = types.YChild{"Ospfarearangeentry", nil}
    for i := range ospfarearangetable.Ospfarearangeentry {
        ospfarearangetable.EntityData.Children[types.GetSegmentPath(&ospfarearangetable.Ospfarearangeentry[i])] = types.YChild{"Ospfarearangeentry", &ospfarearangetable.Ospfarearangeentry[i]}
    }
    ospfarearangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfarearangetable.EntityData)
}

// OSPFMIB_Ospfarearangetable_Ospfarearangeentry
// A single area address range.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfarearangetable_Ospfarearangeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The area that the address range is to be found
    // within. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfarearangeareaid interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfarearangenet interface{}

    // The subnet mask that pertains to the net or subnet. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfarearangemask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfarearangestatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated summary (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is Ospfarearangeeffect.
    Ospfarearangeeffect interface{}
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetEntityData() *types.CommonEntityData {
    ospfarearangeentry.EntityData.YFilter = ospfarearangeentry.YFilter
    ospfarearangeentry.EntityData.YangName = "ospfAreaRangeEntry"
    ospfarearangeentry.EntityData.BundleName = "cisco_ios_xe"
    ospfarearangeentry.EntityData.ParentYangName = "ospfAreaRangeTable"
    ospfarearangeentry.EntityData.SegmentPath = "ospfAreaRangeEntry" + "[ospfAreaRangeAreaId='" + fmt.Sprintf("%v", ospfarearangeentry.Ospfarearangeareaid) + "']" + "[ospfAreaRangeNet='" + fmt.Sprintf("%v", ospfarearangeentry.Ospfarearangenet) + "']"
    ospfarearangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfarearangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfarearangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfarearangeentry.EntityData.Children = make(map[string]types.YChild)
    ospfarearangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfarearangeentry.EntityData.Leafs["ospfAreaRangeAreaId"] = types.YLeaf{"Ospfarearangeareaid", ospfarearangeentry.Ospfarearangeareaid}
    ospfarearangeentry.EntityData.Leafs["ospfAreaRangeNet"] = types.YLeaf{"Ospfarearangenet", ospfarearangeentry.Ospfarearangenet}
    ospfarearangeentry.EntityData.Leafs["ospfAreaRangeMask"] = types.YLeaf{"Ospfarearangemask", ospfarearangeentry.Ospfarearangemask}
    ospfarearangeentry.EntityData.Leafs["ospfAreaRangeStatus"] = types.YLeaf{"Ospfarearangestatus", ospfarearangeentry.Ospfarearangestatus}
    ospfarearangeentry.EntityData.Leafs["ospfAreaRangeEffect"] = types.YLeaf{"Ospfarearangeeffect", ospfarearangeentry.Ospfarearangeeffect}
    return &(ospfarearangeentry.EntityData)
}

// OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect represents being advertised at all outside the area.
type OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect string

const (
    OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect_advertiseMatching OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect = "advertiseMatching"

    OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect_doNotAdvertiseMatching OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect = "doNotAdvertiseMatching"
)

// OSPFMIB_Ospfhosttable
// The Host/Metric Table indicates what hosts are directly
// 
// attached to the router, what metrics and types
// of service should be advertised for them,
// and what areas they are found within.
type OSPFMIB_Ospfhosttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A metric to be advertised, for a given type of service, when a given host
    // is reachable.  Information in this table is persistent and when this object
    // is written the entity SHOULD save the change to non-volatile storage. The
    // type is slice of OSPFMIB_Ospfhosttable_Ospfhostentry.
    Ospfhostentry []OSPFMIB_Ospfhosttable_Ospfhostentry
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetEntityData() *types.CommonEntityData {
    ospfhosttable.EntityData.YFilter = ospfhosttable.YFilter
    ospfhosttable.EntityData.YangName = "ospfHostTable"
    ospfhosttable.EntityData.BundleName = "cisco_ios_xe"
    ospfhosttable.EntityData.ParentYangName = "OSPF-MIB"
    ospfhosttable.EntityData.SegmentPath = "ospfHostTable"
    ospfhosttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfhosttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfhosttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfhosttable.EntityData.Children = make(map[string]types.YChild)
    ospfhosttable.EntityData.Children["ospfHostEntry"] = types.YChild{"Ospfhostentry", nil}
    for i := range ospfhosttable.Ospfhostentry {
        ospfhosttable.EntityData.Children[types.GetSegmentPath(&ospfhosttable.Ospfhostentry[i])] = types.YChild{"Ospfhostentry", &ospfhosttable.Ospfhostentry[i]}
    }
    ospfhosttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfhosttable.EntityData)
}

// OSPFMIB_Ospfhosttable_Ospfhostentry
// A metric to be advertised, for a given type of
// service, when a given host is reachable.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfhosttable_Ospfhostentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the host. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfhostipaddress interface{}

    // This attribute is a key. The Type of Service of the route being configured.
    // The type is interface{} with range: 0..30.
    Ospfhosttos interface{}

    // The metric to be advertised. The type is interface{} with range: 0..65535.
    Ospfhostmetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfhoststatus interface{}

    // The OSPF area to which the host belongs. Deprecated by ospfHostCfgAreaID.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfhostareaid interface{}

    // To configure the OSPF area to which the host belongs. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfhostcfgareaid interface{}
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetEntityData() *types.CommonEntityData {
    ospfhostentry.EntityData.YFilter = ospfhostentry.YFilter
    ospfhostentry.EntityData.YangName = "ospfHostEntry"
    ospfhostentry.EntityData.BundleName = "cisco_ios_xe"
    ospfhostentry.EntityData.ParentYangName = "ospfHostTable"
    ospfhostentry.EntityData.SegmentPath = "ospfHostEntry" + "[ospfHostIpAddress='" + fmt.Sprintf("%v", ospfhostentry.Ospfhostipaddress) + "']" + "[ospfHostTOS='" + fmt.Sprintf("%v", ospfhostentry.Ospfhosttos) + "']"
    ospfhostentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfhostentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfhostentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfhostentry.EntityData.Children = make(map[string]types.YChild)
    ospfhostentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfhostentry.EntityData.Leafs["ospfHostIpAddress"] = types.YLeaf{"Ospfhostipaddress", ospfhostentry.Ospfhostipaddress}
    ospfhostentry.EntityData.Leafs["ospfHostTOS"] = types.YLeaf{"Ospfhosttos", ospfhostentry.Ospfhosttos}
    ospfhostentry.EntityData.Leafs["ospfHostMetric"] = types.YLeaf{"Ospfhostmetric", ospfhostentry.Ospfhostmetric}
    ospfhostentry.EntityData.Leafs["ospfHostStatus"] = types.YLeaf{"Ospfhoststatus", ospfhostentry.Ospfhoststatus}
    ospfhostentry.EntityData.Leafs["ospfHostAreaID"] = types.YLeaf{"Ospfhostareaid", ospfhostentry.Ospfhostareaid}
    ospfhostentry.EntityData.Leafs["ospfHostCfgAreaID"] = types.YLeaf{"Ospfhostcfgareaid", ospfhostentry.Ospfhostcfgareaid}
    return &(ospfhostentry.EntityData)
}

// OSPFMIB_Ospfiftable
// The OSPF Interface Table describes the interfaces
// from the viewpoint of OSPF.
// It augments the ipAddrTable with OSPF specific information.
type OSPFMIB_Ospfiftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The OSPF interface entry describes one interface from the viewpoint of
    // OSPF.  Information in this table is persistent and when this object is
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is slice of OSPFMIB_Ospfiftable_Ospfifentry.
    Ospfifentry []OSPFMIB_Ospfiftable_Ospfifentry
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetEntityData() *types.CommonEntityData {
    ospfiftable.EntityData.YFilter = ospfiftable.YFilter
    ospfiftable.EntityData.YangName = "ospfIfTable"
    ospfiftable.EntityData.BundleName = "cisco_ios_xe"
    ospfiftable.EntityData.ParentYangName = "OSPF-MIB"
    ospfiftable.EntityData.SegmentPath = "ospfIfTable"
    ospfiftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfiftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfiftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfiftable.EntityData.Children = make(map[string]types.YChild)
    ospfiftable.EntityData.Children["ospfIfEntry"] = types.YChild{"Ospfifentry", nil}
    for i := range ospfiftable.Ospfifentry {
        ospfiftable.EntityData.Children[types.GetSegmentPath(&ospfiftable.Ospfifentry[i])] = types.YChild{"Ospfifentry", &ospfiftable.Ospfifentry[i]}
    }
    ospfiftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfiftable.EntityData)
}

// OSPFMIB_Ospfiftable_Ospfifentry
// The OSPF interface entry describes one interface
// from the viewpoint of OSPF.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfiftable_Ospfifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifipaddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the corresponding value of ifIndex for
    // interfaces having no IP address. The type is interface{} with range:
    // 0..2147483647.
    Ospfaddresslessif interface{}

    // A 32-bit integer uniquely identifying the area to which the interface
    // connects.  Area ID 0.0.0.0 is used for the OSPF backbone. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifareaid interface{}

    // The OSPF interface type. By way of a default, this field may be intuited
    // from the corresponding value of ifType. Broadcast LANs, such as Ethernet
    // and IEEE 802.5, take the value 'broadcast', X.25 and similar technologies
    // take the value 'nbma', and links that are definitively point to point take
    // the value 'pointToPoint'. The type is Ospfiftype.
    Ospfiftype interface{}

    // The OSPF interface's administrative status. The value formed on the
    // interface, and the interface will be advertised as an internal route to
    // some area. The value 'disabled' denotes that the interface is external to
    // OSPF. The type is Status.
    Ospfifadminstat interface{}

    // The priority of this interface.  Used in multi-access networks, this field
    // is used in the designated router election algorithm.  The value 0 signifies
    // that the router is not eligible to become the designated router on this
    // particular network.  In the event of a tie in this value, routers will use
    // their Router ID as a tie breaker. The type is interface{} with range:
    // 0..255.
    Ospfifrtrpriority interface{}

    // The estimated number of seconds it takes to transmit a link state update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    Ospfiftransitdelay interface{}

    // The number of seconds between link state advertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting  database description and Link State request packets. Note
    // that minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    Ospfifretransinterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range: 1..65535.
    // Units are seconds.
    Ospfifhellointerval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down. This should be some multiple
    // of the Hello interval.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    Ospfifrtrdeadinterval interface{}

    // The larger time interval, in seconds, between the Hello packets sent to an
    // inactive non-broadcast multi-access neighbor. The type is interface{} with
    // range: 0..2147483647. Units are seconds.
    Ospfifpollinterval interface{}

    // The OSPF Interface State. The type is Ospfifstate.
    Ospfifstate interface{}

    // The IP address of the designated router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifdesignatedrouter interface{}

    // The IP address of the backup designated router. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifbackupdesignatedrouter interface{}

    // The number of times this OSPF interface has changed its state or an error
    // has occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ospfifevents interface{}

    // The cleartext password used as an OSPF authentication key when
    // simplePassword security is enabled.  This object does not access any OSPF
    // cryptogaphic (e.g., MD5) authentication key under any circumstance.  If the
    // key length is shorter than 8 octets, the agent will left adjust and zero
    // fill to 8 octets.  Unauthenticated interfaces need no authentication key,
    // and simple password authentication cannot use a key of more than 8 octets. 
    // Note that the use of simplePassword authentication is NOT recommended when
    // there is concern regarding attack upon the OSPF system.  SimplePassword
    // authentication is only sufficient to protect against accidental
    // misconfigurations because it re-uses cleartext passwords [RFC1704].  When
    // read, ospfIfAuthKey always returns an octet string of length zero. The type
    // is string with length: 0..256.
    Ospfifauthkey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfifstatus interface{}

    // The way multicasts should be forwarded on this interface: not forwarded,
    // forwarded as data link multicasts, or forwarded as data link unicasts. 
    // Data link multicasting is not meaningful on point-to-point and NBMA
    // interfaces, and setting ospfMulticastForwarding to 0 effectively disables
    // all multicast forwarding. The type is Ospfifmulticastforwarding.
    Ospfifmulticastforwarding interface{}

    // Indicates whether Demand OSPF procedures (hello suppression to FULL
    // neighbors and setting the DoNotAge flag on propagated LSAs) should be
    // performed on this interface. The type is bool.
    Ospfifdemand interface{}

    // The authentication type specified for an interface.  Note that this object
    // can be used to engage in significant attacks against an OSPF router. The
    // type is OspfAuthenticationType.
    Ospfifauthtype interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Ospfiflsacount interface{}

    // The 32-bit unsigned sum of the Link State Advertisements' LS checksums
    // contained in this interface's link-local link state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database and to compare the interface link state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    Ospfiflsacksumsum interface{}

    // The Router ID of the designated router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifdesignatedrouterid interface{}

    // The Router ID of the backup designated router. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifbackupdesignatedrouterid interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Cospfiflsacount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this interface's link-local link  state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database, and to compare the interface link-state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    Cospfiflsacksumsum interface{}
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetEntityData() *types.CommonEntityData {
    ospfifentry.EntityData.YFilter = ospfifentry.YFilter
    ospfifentry.EntityData.YangName = "ospfIfEntry"
    ospfifentry.EntityData.BundleName = "cisco_ios_xe"
    ospfifentry.EntityData.ParentYangName = "ospfIfTable"
    ospfifentry.EntityData.SegmentPath = "ospfIfEntry" + "[ospfIfIpAddress='" + fmt.Sprintf("%v", ospfifentry.Ospfifipaddress) + "']" + "[ospfAddressLessIf='" + fmt.Sprintf("%v", ospfifentry.Ospfaddresslessif) + "']"
    ospfifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfifentry.EntityData.Children = make(map[string]types.YChild)
    ospfifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfifentry.EntityData.Leafs["ospfIfIpAddress"] = types.YLeaf{"Ospfifipaddress", ospfifentry.Ospfifipaddress}
    ospfifentry.EntityData.Leafs["ospfAddressLessIf"] = types.YLeaf{"Ospfaddresslessif", ospfifentry.Ospfaddresslessif}
    ospfifentry.EntityData.Leafs["ospfIfAreaId"] = types.YLeaf{"Ospfifareaid", ospfifentry.Ospfifareaid}
    ospfifentry.EntityData.Leafs["ospfIfType"] = types.YLeaf{"Ospfiftype", ospfifentry.Ospfiftype}
    ospfifentry.EntityData.Leafs["ospfIfAdminStat"] = types.YLeaf{"Ospfifadminstat", ospfifentry.Ospfifadminstat}
    ospfifentry.EntityData.Leafs["ospfIfRtrPriority"] = types.YLeaf{"Ospfifrtrpriority", ospfifentry.Ospfifrtrpriority}
    ospfifentry.EntityData.Leafs["ospfIfTransitDelay"] = types.YLeaf{"Ospfiftransitdelay", ospfifentry.Ospfiftransitdelay}
    ospfifentry.EntityData.Leafs["ospfIfRetransInterval"] = types.YLeaf{"Ospfifretransinterval", ospfifentry.Ospfifretransinterval}
    ospfifentry.EntityData.Leafs["ospfIfHelloInterval"] = types.YLeaf{"Ospfifhellointerval", ospfifentry.Ospfifhellointerval}
    ospfifentry.EntityData.Leafs["ospfIfRtrDeadInterval"] = types.YLeaf{"Ospfifrtrdeadinterval", ospfifentry.Ospfifrtrdeadinterval}
    ospfifentry.EntityData.Leafs["ospfIfPollInterval"] = types.YLeaf{"Ospfifpollinterval", ospfifentry.Ospfifpollinterval}
    ospfifentry.EntityData.Leafs["ospfIfState"] = types.YLeaf{"Ospfifstate", ospfifentry.Ospfifstate}
    ospfifentry.EntityData.Leafs["ospfIfDesignatedRouter"] = types.YLeaf{"Ospfifdesignatedrouter", ospfifentry.Ospfifdesignatedrouter}
    ospfifentry.EntityData.Leafs["ospfIfBackupDesignatedRouter"] = types.YLeaf{"Ospfifbackupdesignatedrouter", ospfifentry.Ospfifbackupdesignatedrouter}
    ospfifentry.EntityData.Leafs["ospfIfEvents"] = types.YLeaf{"Ospfifevents", ospfifentry.Ospfifevents}
    ospfifentry.EntityData.Leafs["ospfIfAuthKey"] = types.YLeaf{"Ospfifauthkey", ospfifentry.Ospfifauthkey}
    ospfifentry.EntityData.Leafs["ospfIfStatus"] = types.YLeaf{"Ospfifstatus", ospfifentry.Ospfifstatus}
    ospfifentry.EntityData.Leafs["ospfIfMulticastForwarding"] = types.YLeaf{"Ospfifmulticastforwarding", ospfifentry.Ospfifmulticastforwarding}
    ospfifentry.EntityData.Leafs["ospfIfDemand"] = types.YLeaf{"Ospfifdemand", ospfifentry.Ospfifdemand}
    ospfifentry.EntityData.Leafs["ospfIfAuthType"] = types.YLeaf{"Ospfifauthtype", ospfifentry.Ospfifauthtype}
    ospfifentry.EntityData.Leafs["ospfIfLsaCount"] = types.YLeaf{"Ospfiflsacount", ospfifentry.Ospfiflsacount}
    ospfifentry.EntityData.Leafs["ospfIfLsaCksumSum"] = types.YLeaf{"Ospfiflsacksumsum", ospfifentry.Ospfiflsacksumsum}
    ospfifentry.EntityData.Leafs["ospfIfDesignatedRouterId"] = types.YLeaf{"Ospfifdesignatedrouterid", ospfifentry.Ospfifdesignatedrouterid}
    ospfifentry.EntityData.Leafs["ospfIfBackupDesignatedRouterId"] = types.YLeaf{"Ospfifbackupdesignatedrouterid", ospfifentry.Ospfifbackupdesignatedrouterid}
    ospfifentry.EntityData.Leafs["cospfIfLsaCount"] = types.YLeaf{"Cospfiflsacount", ospfifentry.Cospfiflsacount}
    ospfifentry.EntityData.Leafs["cospfIfLsaCksumSum"] = types.YLeaf{"Cospfiflsacksumsum", ospfifentry.Cospfiflsacksumsum}
    return &(ospfifentry.EntityData)
}

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding represents disables all multicast forwarding.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_blocked OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "blocked"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_multicast OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "multicast"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_unicast OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "unicast"
)

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate represents The OSPF Interface State.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_down OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "down"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_loopback OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "loopback"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_waiting OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "waiting"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_pointToPoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "pointToPoint"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_designatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "designatedRouter"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_backupDesignatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "backupDesignatedRouter"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_otherDesignatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "otherDesignatedRouter"
)

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype represents value 'pointToPoint'.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_broadcast OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "broadcast"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_nbma OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "nbma"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_pointToPoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "pointToPoint"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_pointToMultipoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "pointToMultipoint"
)

// OSPFMIB_Ospfifmetrictable
// The Metric Table describes the metrics to be advertised
// for a specified interface at the various types of service.
// As such, this table is an adjunct of the OSPF Interface
// Table.
// 
// Types of service, as defined by RFC 791, have the ability
// to request low delay, high bandwidth, or reliable linkage.
// 
// For the purposes of this specification, the measure of
// bandwidth:
// 
// Metric = referenceBandwidth / ifSpeed
// 
// is the default value.
// The default reference bandwidth is 10^8.
// For multiple link interfaces, note that ifSpeed is the sum
// of the individual link speeds.  This yields a number having
// the following typical values:
// 
// Network Type/bit rate   Metric
// 
// >= 100 MBPS                 1
// Ethernet/802.3             10
// E1                         48
// T1 (ESF)                   65
// 64 KBPS                    1562
// 56 KBPS                    1785
// 19.2 KBPS                  5208
// 9.6 KBPS                   10416
// 
// Routes that are not specified use the default
// (TOS 0) metric.
// 
// Note that the default reference bandwidth can be configured
// using the general group object ospfReferenceBandwidth.
type OSPFMIB_Ospfifmetrictable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular TOS metric for a non-virtual interface identified by the
    // interface index.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_Ospfifmetrictable_Ospfifmetricentry.
    Ospfifmetricentry []OSPFMIB_Ospfifmetrictable_Ospfifmetricentry
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetEntityData() *types.CommonEntityData {
    ospfifmetrictable.EntityData.YFilter = ospfifmetrictable.YFilter
    ospfifmetrictable.EntityData.YangName = "ospfIfMetricTable"
    ospfifmetrictable.EntityData.BundleName = "cisco_ios_xe"
    ospfifmetrictable.EntityData.ParentYangName = "OSPF-MIB"
    ospfifmetrictable.EntityData.SegmentPath = "ospfIfMetricTable"
    ospfifmetrictable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfifmetrictable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfifmetrictable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfifmetrictable.EntityData.Children = make(map[string]types.YChild)
    ospfifmetrictable.EntityData.Children["ospfIfMetricEntry"] = types.YChild{"Ospfifmetricentry", nil}
    for i := range ospfifmetrictable.Ospfifmetricentry {
        ospfifmetrictable.EntityData.Children[types.GetSegmentPath(&ospfifmetrictable.Ospfifmetricentry[i])] = types.YChild{"Ospfifmetricentry", &ospfifmetrictable.Ospfifmetricentry[i]}
    }
    ospfifmetrictable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfifmetrictable.EntityData)
}

// OSPFMIB_Ospfifmetrictable_Ospfifmetricentry
// A particular TOS metric for a non-virtual interface
// identified by the interface index.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfifmetrictable_Ospfifmetricentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface.  On row
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfifmetricipaddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the value of ifIndex for interfaces having
    // no IP address.  On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    Ospfifmetricaddresslessif interface{}

    // This attribute is a key. The Type of Service metric being referenced. On
    // row creation, this can be derived from the instance. The type is
    // interface{} with range: 0..30.
    Ospfifmetrictos interface{}

    // The metric of using this Type of Service on this interface.  The default
    // value of the TOS 0 metric is 10^8 / ifSpeed. The type is interface{} with
    // range: 0..65535.
    Ospfifmetricvalue interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfifmetricstatus interface{}
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetEntityData() *types.CommonEntityData {
    ospfifmetricentry.EntityData.YFilter = ospfifmetricentry.YFilter
    ospfifmetricentry.EntityData.YangName = "ospfIfMetricEntry"
    ospfifmetricentry.EntityData.BundleName = "cisco_ios_xe"
    ospfifmetricentry.EntityData.ParentYangName = "ospfIfMetricTable"
    ospfifmetricentry.EntityData.SegmentPath = "ospfIfMetricEntry" + "[ospfIfMetricIpAddress='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetricipaddress) + "']" + "[ospfIfMetricAddressLessIf='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetricaddresslessif) + "']" + "[ospfIfMetricTOS='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetrictos) + "']"
    ospfifmetricentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfifmetricentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfifmetricentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfifmetricentry.EntityData.Children = make(map[string]types.YChild)
    ospfifmetricentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfifmetricentry.EntityData.Leafs["ospfIfMetricIpAddress"] = types.YLeaf{"Ospfifmetricipaddress", ospfifmetricentry.Ospfifmetricipaddress}
    ospfifmetricentry.EntityData.Leafs["ospfIfMetricAddressLessIf"] = types.YLeaf{"Ospfifmetricaddresslessif", ospfifmetricentry.Ospfifmetricaddresslessif}
    ospfifmetricentry.EntityData.Leafs["ospfIfMetricTOS"] = types.YLeaf{"Ospfifmetrictos", ospfifmetricentry.Ospfifmetrictos}
    ospfifmetricentry.EntityData.Leafs["ospfIfMetricValue"] = types.YLeaf{"Ospfifmetricvalue", ospfifmetricentry.Ospfifmetricvalue}
    ospfifmetricentry.EntityData.Leafs["ospfIfMetricStatus"] = types.YLeaf{"Ospfifmetricstatus", ospfifmetricentry.Ospfifmetricstatus}
    return &(ospfifmetricentry.EntityData)
}

// OSPFMIB_Ospfvirtiftable
// Information about this router's virtual interfaces
// that the OSPF Process is configured to carry on.
type OSPFMIB_Ospfvirtiftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single virtual interface.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_Ospfvirtiftable_Ospfvirtifentry.
    Ospfvirtifentry []OSPFMIB_Ospfvirtiftable_Ospfvirtifentry
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetEntityData() *types.CommonEntityData {
    ospfvirtiftable.EntityData.YFilter = ospfvirtiftable.YFilter
    ospfvirtiftable.EntityData.YangName = "ospfVirtIfTable"
    ospfvirtiftable.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtiftable.EntityData.ParentYangName = "OSPF-MIB"
    ospfvirtiftable.EntityData.SegmentPath = "ospfVirtIfTable"
    ospfvirtiftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtiftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtiftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtiftable.EntityData.Children = make(map[string]types.YChild)
    ospfvirtiftable.EntityData.Children["ospfVirtIfEntry"] = types.YChild{"Ospfvirtifentry", nil}
    for i := range ospfvirtiftable.Ospfvirtifentry {
        ospfvirtiftable.EntityData.Children[types.GetSegmentPath(&ospfvirtiftable.Ospfvirtifentry[i])] = types.YChild{"Ospfvirtifentry", &ospfvirtiftable.Ospfvirtifentry[i]}
    }
    ospfvirtiftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfvirtiftable.EntityData)
}

// OSPFMIB_Ospfvirtiftable_Ospfvirtifentry
// Information about a single virtual interface.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfvirtiftable_Ospfvirtifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtifareaid interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtifneighbor interface{}

    // The estimated number of seconds it takes to transmit a Link State update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    Ospfvirtiftransitdelay interface{}

    // The number of seconds between link state avertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting database description and Link State request packets.  This
    // value should be well over the expected round-trip time.  Note that the
    // minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    Ospfvirtifretransinterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 1..65535. Units are seconds.
    Ospfvirtifhellointerval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down.  This should be some multiple
    // of the Hello interval.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    Ospfvirtifrtrdeadinterval interface{}

    // OSPF virtual interface states. The type is Ospfvirtifstate.
    Ospfvirtifstate interface{}

    // The number of state changes or error events on this virtual link. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    Ospfvirtifevents interface{}

    // The cleartext password used as an OSPF authentication key when
    // simplePassword security is enabled.  This object does not access any OSPF
    // cryptogaphic (e.g., MD5) authentication key under any circumstance.  If the
    // key length is shorter than 8 octets, the agent will left adjust and zero
    // fill to 8 octets.  Unauthenticated interfaces need no authentication key,
    // and simple password authentication cannot use a key of more than 8 octets. 
    // Note that the use of simplePassword authentication is NOT recommended when
    // there is concern regarding attack upon the OSPF system.  SimplePassword
    // authentication is only sufficient to protect against accidental
    // misconfigurations because it re-uses cleartext passwords.  [RFC1704]  When
    // read, ospfIfAuthKey always returns an octet string of length zero. The type
    // is string with length: 0..256.
    Ospfvirtifauthkey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfvirtifstatus interface{}

    // The authentication type specified for a virtual interface.  Note that this
    // object can be used to engage in significant attacks against an OSPF router.
    // The type is OspfAuthenticationType.
    Ospfvirtifauthtype interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Ospfvirtiflsacount interface{}

    // The 32-bit unsigned sum of the link state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface link
    // state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    Ospfvirtiflsacksumsum interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Cospfvirtiflsacount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface
    // link-state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    Cospfvirtiflsacksumsum interface{}
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetEntityData() *types.CommonEntityData {
    ospfvirtifentry.EntityData.YFilter = ospfvirtifentry.YFilter
    ospfvirtifentry.EntityData.YangName = "ospfVirtIfEntry"
    ospfvirtifentry.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtifentry.EntityData.ParentYangName = "ospfVirtIfTable"
    ospfvirtifentry.EntityData.SegmentPath = "ospfVirtIfEntry" + "[ospfVirtIfAreaId='" + fmt.Sprintf("%v", ospfvirtifentry.Ospfvirtifareaid) + "']" + "[ospfVirtIfNeighbor='" + fmt.Sprintf("%v", ospfvirtifentry.Ospfvirtifneighbor) + "']"
    ospfvirtifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtifentry.EntityData.Children = make(map[string]types.YChild)
    ospfvirtifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfAreaId"] = types.YLeaf{"Ospfvirtifareaid", ospfvirtifentry.Ospfvirtifareaid}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfNeighbor"] = types.YLeaf{"Ospfvirtifneighbor", ospfvirtifentry.Ospfvirtifneighbor}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfTransitDelay"] = types.YLeaf{"Ospfvirtiftransitdelay", ospfvirtifentry.Ospfvirtiftransitdelay}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfRetransInterval"] = types.YLeaf{"Ospfvirtifretransinterval", ospfvirtifentry.Ospfvirtifretransinterval}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfHelloInterval"] = types.YLeaf{"Ospfvirtifhellointerval", ospfvirtifentry.Ospfvirtifhellointerval}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfRtrDeadInterval"] = types.YLeaf{"Ospfvirtifrtrdeadinterval", ospfvirtifentry.Ospfvirtifrtrdeadinterval}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfState"] = types.YLeaf{"Ospfvirtifstate", ospfvirtifentry.Ospfvirtifstate}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfEvents"] = types.YLeaf{"Ospfvirtifevents", ospfvirtifentry.Ospfvirtifevents}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfAuthKey"] = types.YLeaf{"Ospfvirtifauthkey", ospfvirtifentry.Ospfvirtifauthkey}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfStatus"] = types.YLeaf{"Ospfvirtifstatus", ospfvirtifentry.Ospfvirtifstatus}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfAuthType"] = types.YLeaf{"Ospfvirtifauthtype", ospfvirtifentry.Ospfvirtifauthtype}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfLsaCount"] = types.YLeaf{"Ospfvirtiflsacount", ospfvirtifentry.Ospfvirtiflsacount}
    ospfvirtifentry.EntityData.Leafs["ospfVirtIfLsaCksumSum"] = types.YLeaf{"Ospfvirtiflsacksumsum", ospfvirtifentry.Ospfvirtiflsacksumsum}
    ospfvirtifentry.EntityData.Leafs["cospfVirtIfLsaCount"] = types.YLeaf{"Cospfvirtiflsacount", ospfvirtifentry.Cospfvirtiflsacount}
    ospfvirtifentry.EntityData.Leafs["cospfVirtIfLsaCksumSum"] = types.YLeaf{"Cospfvirtiflsacksumsum", ospfvirtifentry.Cospfvirtiflsacksumsum}
    return &(ospfvirtifentry.EntityData)
}

// OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate represents OSPF virtual interface states.
type OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate string

const (
    OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate_down OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate = "down"

    OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate_pointToPoint OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate = "pointToPoint"
)

// OSPFMIB_Ospfnbrtable
// A table describing all non-virtual neighbors
// in the locality of the OSPF router.
type OSPFMIB_Ospfnbrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single neighbor.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile  storage. The type is slice of
    // OSPFMIB_Ospfnbrtable_Ospfnbrentry.
    Ospfnbrentry []OSPFMIB_Ospfnbrtable_Ospfnbrentry
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetEntityData() *types.CommonEntityData {
    ospfnbrtable.EntityData.YFilter = ospfnbrtable.YFilter
    ospfnbrtable.EntityData.YangName = "ospfNbrTable"
    ospfnbrtable.EntityData.BundleName = "cisco_ios_xe"
    ospfnbrtable.EntityData.ParentYangName = "OSPF-MIB"
    ospfnbrtable.EntityData.SegmentPath = "ospfNbrTable"
    ospfnbrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfnbrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfnbrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfnbrtable.EntityData.Children = make(map[string]types.YChild)
    ospfnbrtable.EntityData.Children["ospfNbrEntry"] = types.YChild{"Ospfnbrentry", nil}
    for i := range ospfnbrtable.Ospfnbrentry {
        ospfnbrtable.EntityData.Children[types.GetSegmentPath(&ospfnbrtable.Ospfnbrentry[i])] = types.YChild{"Ospfnbrentry", &ospfnbrtable.Ospfnbrentry[i]}
    }
    ospfnbrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfnbrtable.EntityData)
}

// OSPFMIB_Ospfnbrtable_Ospfnbrentry
// The information regarding a single neighbor.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// 
// storage.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address this neighbor is using in its IP
    // source address.  Note that, on addressless links, this will not be 0.0.0.0
    // but the  address of another of the neighbor's interfaces. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfnbripaddr interface{}

    // This attribute is a key. On an interface having an IP address, zero. On
    // addressless interfaces, the corresponding value of ifIndex in the Internet
    // Standard MIB. On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    Ospfnbraddresslessindex interface{}

    // A 32-bit integer (represented as a type IpAddress) uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfnbrrtrid interface{}

    // A bit mask corresponding to the neighbor's options field.  Bit 0, if set,
    // indicates that the system will operate on Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0
    // metric.  Bit 1, if set, indicates that the associated area accepts and
    // operates on external information; if zero, it is a stub area.  Bit 2, if
    // set, indicates that the system is capable of routing IP multicast
    // datagrams, that is that it implements the multicast extensions to OSPF. 
    // Bit 3, if set, indicates that the associated area is an NSSA.  These areas
    // are capable of carrying type-7 external advertisements, which are
    // translated into type-5 external advertisements at NSSA borders. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfnbroptions interface{}

    // The priority of this neighbor in the designated router election algorithm. 
    // The value 0 signifies that the neighbor is not eligible to become the
    // designated router on this particular network. The type is interface{} with
    // range: 0..255.
    Ospfnbrpriority interface{}

    // The state of the relationship with this neighbor. The type is Ospfnbrstate.
    Ospfnbrstate interface{}

    // The number of times this neighbor relationship has changed state or an
    // error has occurred.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfnbrevents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    Ospfnbrlsretransqlen interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfnbmanbrstatus interface{}

    // This variable displays the status of the entry; 'dynamic' and 'permanent'
    // refer to how the neighbor became known. The type is Ospfnbmanbrpermanence.
    Ospfnbmanbrpermanence interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    Ospfnbrhellosuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is Ospfnbrrestarthelperstatus.
    Ospfnbrrestarthelperstatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Ospfnbrrestarthelperage interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is Ospfnbrrestarthelperexitreason.
    Ospfnbrrestarthelperexitreason interface{}
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetEntityData() *types.CommonEntityData {
    ospfnbrentry.EntityData.YFilter = ospfnbrentry.YFilter
    ospfnbrentry.EntityData.YangName = "ospfNbrEntry"
    ospfnbrentry.EntityData.BundleName = "cisco_ios_xe"
    ospfnbrentry.EntityData.ParentYangName = "ospfNbrTable"
    ospfnbrentry.EntityData.SegmentPath = "ospfNbrEntry" + "[ospfNbrIpAddr='" + fmt.Sprintf("%v", ospfnbrentry.Ospfnbripaddr) + "']" + "[ospfNbrAddressLessIndex='" + fmt.Sprintf("%v", ospfnbrentry.Ospfnbraddresslessindex) + "']"
    ospfnbrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfnbrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfnbrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfnbrentry.EntityData.Children = make(map[string]types.YChild)
    ospfnbrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfnbrentry.EntityData.Leafs["ospfNbrIpAddr"] = types.YLeaf{"Ospfnbripaddr", ospfnbrentry.Ospfnbripaddr}
    ospfnbrentry.EntityData.Leafs["ospfNbrAddressLessIndex"] = types.YLeaf{"Ospfnbraddresslessindex", ospfnbrentry.Ospfnbraddresslessindex}
    ospfnbrentry.EntityData.Leafs["ospfNbrRtrId"] = types.YLeaf{"Ospfnbrrtrid", ospfnbrentry.Ospfnbrrtrid}
    ospfnbrentry.EntityData.Leafs["ospfNbrOptions"] = types.YLeaf{"Ospfnbroptions", ospfnbrentry.Ospfnbroptions}
    ospfnbrentry.EntityData.Leafs["ospfNbrPriority"] = types.YLeaf{"Ospfnbrpriority", ospfnbrentry.Ospfnbrpriority}
    ospfnbrentry.EntityData.Leafs["ospfNbrState"] = types.YLeaf{"Ospfnbrstate", ospfnbrentry.Ospfnbrstate}
    ospfnbrentry.EntityData.Leafs["ospfNbrEvents"] = types.YLeaf{"Ospfnbrevents", ospfnbrentry.Ospfnbrevents}
    ospfnbrentry.EntityData.Leafs["ospfNbrLsRetransQLen"] = types.YLeaf{"Ospfnbrlsretransqlen", ospfnbrentry.Ospfnbrlsretransqlen}
    ospfnbrentry.EntityData.Leafs["ospfNbmaNbrStatus"] = types.YLeaf{"Ospfnbmanbrstatus", ospfnbrentry.Ospfnbmanbrstatus}
    ospfnbrentry.EntityData.Leafs["ospfNbmaNbrPermanence"] = types.YLeaf{"Ospfnbmanbrpermanence", ospfnbrentry.Ospfnbmanbrpermanence}
    ospfnbrentry.EntityData.Leafs["ospfNbrHelloSuppressed"] = types.YLeaf{"Ospfnbrhellosuppressed", ospfnbrentry.Ospfnbrhellosuppressed}
    ospfnbrentry.EntityData.Leafs["ospfNbrRestartHelperStatus"] = types.YLeaf{"Ospfnbrrestarthelperstatus", ospfnbrentry.Ospfnbrrestarthelperstatus}
    ospfnbrentry.EntityData.Leafs["ospfNbrRestartHelperAge"] = types.YLeaf{"Ospfnbrrestarthelperage", ospfnbrentry.Ospfnbrrestarthelperage}
    ospfnbrentry.EntityData.Leafs["ospfNbrRestartHelperExitReason"] = types.YLeaf{"Ospfnbrrestarthelperexitreason", ospfnbrentry.Ospfnbrrestarthelperexitreason}
    return &(ospfnbrentry.EntityData)
}

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence represents became known.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence_dynamic OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence = "dynamic"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence_permanent OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence = "permanent"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_none OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "none"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_inProgress OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "inProgress"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_completed OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "completed"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_timedOut OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "timedOut"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_topologyChanged OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "topologyChanged"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus_notHelping OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus = "notHelping"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus_helping OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus = "helping"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate represents The state of the relationship with this neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_down OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "down"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_attempt OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "attempt"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_init OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "init"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_twoWay OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "twoWay"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_exchangeStart OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "exchangeStart"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_exchange OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "exchange"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_loading OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "loading"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_full OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "full"
)

// OSPFMIB_Ospfvirtnbrtable
// This table describes all virtual neighbors.
// Since virtual links are configured
// in the Virtual Interface Table, this table is read-only.
type OSPFMIB_Ospfvirtnbrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual neighbor information. The type is slice of
    // OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry.
    Ospfvirtnbrentry []OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetEntityData() *types.CommonEntityData {
    ospfvirtnbrtable.EntityData.YFilter = ospfvirtnbrtable.YFilter
    ospfvirtnbrtable.EntityData.YangName = "ospfVirtNbrTable"
    ospfvirtnbrtable.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtnbrtable.EntityData.ParentYangName = "OSPF-MIB"
    ospfvirtnbrtable.EntityData.SegmentPath = "ospfVirtNbrTable"
    ospfvirtnbrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtnbrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtnbrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtnbrtable.EntityData.Children = make(map[string]types.YChild)
    ospfvirtnbrtable.EntityData.Children["ospfVirtNbrEntry"] = types.YChild{"Ospfvirtnbrentry", nil}
    for i := range ospfvirtnbrtable.Ospfvirtnbrentry {
        ospfvirtnbrtable.EntityData.Children[types.GetSegmentPath(&ospfvirtnbrtable.Ospfvirtnbrentry[i])] = types.YChild{"Ospfvirtnbrentry", &ospfvirtnbrtable.Ospfvirtnbrentry[i]}
    }
    ospfvirtnbrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfvirtnbrtable.EntityData)
}

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry
// Virtual neighbor information.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Transit Area Identifier. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtnbrarea interface{}

    // This attribute is a key. A 32-bit integer uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtnbrrtrid interface{}

    // The IP address this virtual neighbor is using. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtnbripaddr interface{}

    // A bit mask corresponding to the neighbor's options field.  Bit 1, if set,
    // indicates that the system will operate on Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0
    // metric.  Bit 2, if set, indicates that the system is network multicast
    // capable, i.e., that it implements OSPF multicast routing. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfvirtnbroptions interface{}

    // The state of the virtual neighbor relationship. The type is
    // Ospfvirtnbrstate.
    Ospfvirtnbrstate interface{}

    // The number of times this virtual link has changed its state or an error has
    // occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ospfvirtnbrevents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    Ospfvirtnbrlsretransqlen interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    Ospfvirtnbrhellosuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is Ospfvirtnbrrestarthelperstatus.
    Ospfvirtnbrrestarthelperstatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Ospfvirtnbrrestarthelperage interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is Ospfvirtnbrrestarthelperexitreason.
    Ospfvirtnbrrestarthelperexitreason interface{}
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetEntityData() *types.CommonEntityData {
    ospfvirtnbrentry.EntityData.YFilter = ospfvirtnbrentry.YFilter
    ospfvirtnbrentry.EntityData.YangName = "ospfVirtNbrEntry"
    ospfvirtnbrentry.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtnbrentry.EntityData.ParentYangName = "ospfVirtNbrTable"
    ospfvirtnbrentry.EntityData.SegmentPath = "ospfVirtNbrEntry" + "[ospfVirtNbrArea='" + fmt.Sprintf("%v", ospfvirtnbrentry.Ospfvirtnbrarea) + "']" + "[ospfVirtNbrRtrId='" + fmt.Sprintf("%v", ospfvirtnbrentry.Ospfvirtnbrrtrid) + "']"
    ospfvirtnbrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtnbrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtnbrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtnbrentry.EntityData.Children = make(map[string]types.YChild)
    ospfvirtnbrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrArea"] = types.YLeaf{"Ospfvirtnbrarea", ospfvirtnbrentry.Ospfvirtnbrarea}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrRtrId"] = types.YLeaf{"Ospfvirtnbrrtrid", ospfvirtnbrentry.Ospfvirtnbrrtrid}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrIpAddr"] = types.YLeaf{"Ospfvirtnbripaddr", ospfvirtnbrentry.Ospfvirtnbripaddr}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrOptions"] = types.YLeaf{"Ospfvirtnbroptions", ospfvirtnbrentry.Ospfvirtnbroptions}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrState"] = types.YLeaf{"Ospfvirtnbrstate", ospfvirtnbrentry.Ospfvirtnbrstate}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrEvents"] = types.YLeaf{"Ospfvirtnbrevents", ospfvirtnbrentry.Ospfvirtnbrevents}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrLsRetransQLen"] = types.YLeaf{"Ospfvirtnbrlsretransqlen", ospfvirtnbrentry.Ospfvirtnbrlsretransqlen}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrHelloSuppressed"] = types.YLeaf{"Ospfvirtnbrhellosuppressed", ospfvirtnbrentry.Ospfvirtnbrhellosuppressed}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrRestartHelperStatus"] = types.YLeaf{"Ospfvirtnbrrestarthelperstatus", ospfvirtnbrentry.Ospfvirtnbrrestarthelperstatus}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrRestartHelperAge"] = types.YLeaf{"Ospfvirtnbrrestarthelperage", ospfvirtnbrentry.Ospfvirtnbrrestarthelperage}
    ospfvirtnbrentry.EntityData.Leafs["ospfVirtNbrRestartHelperExitReason"] = types.YLeaf{"Ospfvirtnbrrestarthelperexitreason", ospfvirtnbrentry.Ospfvirtnbrrestarthelperexitreason}
    return &(ospfvirtnbrentry.EntityData)
}

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_none OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "none"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_inProgress OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "inProgress"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_completed OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "completed"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_timedOut OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "timedOut"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_topologyChanged OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "topologyChanged"
)

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus_notHelping OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus = "notHelping"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus_helping OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus = "helping"
)

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate represents The state of the virtual neighbor relationship.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_down OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "down"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_attempt OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "attempt"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_init OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "init"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_twoWay OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "twoWay"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_exchangeStart OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "exchangeStart"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_exchange OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "exchange"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_loading OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "loading"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_full OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "full"
)

// OSPFMIB_Ospfextlsdbtable
// The OSPF Process's external LSA link state database.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only external link state
// advertisements.  The purpose is to allow external
// 
// LSAs to be displayed once for the router rather
// than once in each non-stub area.
// 
// Note that external LSAs are also in the AS-scope link state
// database.
type OSPFMIB_Ospfextlsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry.
    Ospfextlsdbentry []OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetEntityData() *types.CommonEntityData {
    ospfextlsdbtable.EntityData.YFilter = ospfextlsdbtable.YFilter
    ospfextlsdbtable.EntityData.YangName = "ospfExtLsdbTable"
    ospfextlsdbtable.EntityData.BundleName = "cisco_ios_xe"
    ospfextlsdbtable.EntityData.ParentYangName = "OSPF-MIB"
    ospfextlsdbtable.EntityData.SegmentPath = "ospfExtLsdbTable"
    ospfextlsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfextlsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfextlsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfextlsdbtable.EntityData.Children = make(map[string]types.YChild)
    ospfextlsdbtable.EntityData.Children["ospfExtLsdbEntry"] = types.YChild{"Ospfextlsdbentry", nil}
    for i := range ospfextlsdbtable.Ospfextlsdbentry {
        ospfextlsdbtable.EntityData.Children[types.GetSegmentPath(&ospfextlsdbtable.Ospfextlsdbentry[i])] = types.YChild{"Ospfextlsdbentry", &ospfextlsdbtable.Ospfextlsdbentry[i]}
    }
    ospfextlsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfextlsdbtable.EntityData)
}

// OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfextlsdbtype.
    Ospfextlsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfextlsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfextlsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfextlsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfextlsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfextlsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 36.
    Ospfextlsdbadvertisement interface{}
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetEntityData() *types.CommonEntityData {
    ospfextlsdbentry.EntityData.YFilter = ospfextlsdbentry.YFilter
    ospfextlsdbentry.EntityData.YangName = "ospfExtLsdbEntry"
    ospfextlsdbentry.EntityData.BundleName = "cisco_ios_xe"
    ospfextlsdbentry.EntityData.ParentYangName = "ospfExtLsdbTable"
    ospfextlsdbentry.EntityData.SegmentPath = "ospfExtLsdbEntry" + "[ospfExtLsdbType='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdbtype) + "']" + "[ospfExtLsdbLsid='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdblsid) + "']" + "[ospfExtLsdbRouterId='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdbrouterid) + "']"
    ospfextlsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfextlsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfextlsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfextlsdbentry.EntityData.Children = make(map[string]types.YChild)
    ospfextlsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbType"] = types.YLeaf{"Ospfextlsdbtype", ospfextlsdbentry.Ospfextlsdbtype}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbLsid"] = types.YLeaf{"Ospfextlsdblsid", ospfextlsdbentry.Ospfextlsdblsid}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbRouterId"] = types.YLeaf{"Ospfextlsdbrouterid", ospfextlsdbentry.Ospfextlsdbrouterid}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbSequence"] = types.YLeaf{"Ospfextlsdbsequence", ospfextlsdbentry.Ospfextlsdbsequence}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbAge"] = types.YLeaf{"Ospfextlsdbage", ospfextlsdbentry.Ospfextlsdbage}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbChecksum"] = types.YLeaf{"Ospfextlsdbchecksum", ospfextlsdbentry.Ospfextlsdbchecksum}
    ospfextlsdbentry.EntityData.Leafs["ospfExtLsdbAdvertisement"] = types.YLeaf{"Ospfextlsdbadvertisement", ospfextlsdbentry.Ospfextlsdbadvertisement}
    return &(ospfextlsdbentry.EntityData)
}

// OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype represents format.
type OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype string

const (
    OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype_asExternalLink OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype = "asExternalLink"
)

// OSPFMIB_Ospfareaaggregatetable
// The Area Aggregate Table acts as an adjunct
// to the Area Table.  It describes those address aggregates
// that are configured to be propagated from an area.
// Its purpose is to reduce the amount of information
// that is known beyond an Area's borders.
// 
// It contains a set of IP address ranges
// specified by an IP address/IP network mask pair.
// For example, a class B address range of X.X.X.X
// with a network mask of 255.255.0.0 includes all IP
// addresses from X.X.0.0 to X.X.255.255.
// 
// Note that if ranges are configured such that one range
// subsumes another range (e.g., 10.0.0.0 mask 255.0.0.0
// and 10.1.0.0 mask 255.255.0.0),
// the most specific match is the preferred one.
type OSPFMIB_Ospfareaaggregatetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single area aggregate entry.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry.
    Ospfareaaggregateentry []OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetEntityData() *types.CommonEntityData {
    ospfareaaggregatetable.EntityData.YFilter = ospfareaaggregatetable.YFilter
    ospfareaaggregatetable.EntityData.YangName = "ospfAreaAggregateTable"
    ospfareaaggregatetable.EntityData.BundleName = "cisco_ios_xe"
    ospfareaaggregatetable.EntityData.ParentYangName = "OSPF-MIB"
    ospfareaaggregatetable.EntityData.SegmentPath = "ospfAreaAggregateTable"
    ospfareaaggregatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfareaaggregatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfareaaggregatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfareaaggregatetable.EntityData.Children = make(map[string]types.YChild)
    ospfareaaggregatetable.EntityData.Children["ospfAreaAggregateEntry"] = types.YChild{"Ospfareaaggregateentry", nil}
    for i := range ospfareaaggregatetable.Ospfareaaggregateentry {
        ospfareaaggregatetable.EntityData.Children[types.GetSegmentPath(&ospfareaaggregatetable.Ospfareaaggregateentry[i])] = types.YChild{"Ospfareaaggregateentry", &ospfareaaggregatetable.Ospfareaaggregateentry[i]}
    }
    ospfareaaggregatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfareaaggregatetable.EntityData)
}

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry
// A single area aggregate entry.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The area within which the address aggregate is to
    // be found. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfareaaggregateareaid interface{}

    // This attribute is a key. The type of the address aggregate.  This field
    // specifies the Lsdb type that this address aggregate applies to. The type is
    // Ospfareaaggregatelsdbtype.
    Ospfareaaggregatelsdbtype interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfareaaggregatenet interface{}

    // This attribute is a key. The subnet mask that pertains to the net or
    // subnet. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfareaaggregatemask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfareaaggregatestatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated aggregate (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is Ospfareaaggregateeffect.
    Ospfareaaggregateeffect interface{}

    // External route tag to be included in NSSA (type-7) LSAs. The type is
    // interface{} with range: 0..4294967295.
    Ospfareaaggregateextroutetag interface{}
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetEntityData() *types.CommonEntityData {
    ospfareaaggregateentry.EntityData.YFilter = ospfareaaggregateentry.YFilter
    ospfareaaggregateentry.EntityData.YangName = "ospfAreaAggregateEntry"
    ospfareaaggregateentry.EntityData.BundleName = "cisco_ios_xe"
    ospfareaaggregateentry.EntityData.ParentYangName = "ospfAreaAggregateTable"
    ospfareaaggregateentry.EntityData.SegmentPath = "ospfAreaAggregateEntry" + "[ospfAreaAggregateAreaID='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregateareaid) + "']" + "[ospfAreaAggregateLsdbType='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatelsdbtype) + "']" + "[ospfAreaAggregateNet='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatenet) + "']" + "[ospfAreaAggregateMask='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatemask) + "']"
    ospfareaaggregateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfareaaggregateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfareaaggregateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfareaaggregateentry.EntityData.Children = make(map[string]types.YChild)
    ospfareaaggregateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateAreaID"] = types.YLeaf{"Ospfareaaggregateareaid", ospfareaaggregateentry.Ospfareaaggregateareaid}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateLsdbType"] = types.YLeaf{"Ospfareaaggregatelsdbtype", ospfareaaggregateentry.Ospfareaaggregatelsdbtype}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateNet"] = types.YLeaf{"Ospfareaaggregatenet", ospfareaaggregateentry.Ospfareaaggregatenet}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateMask"] = types.YLeaf{"Ospfareaaggregatemask", ospfareaaggregateentry.Ospfareaaggregatemask}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateStatus"] = types.YLeaf{"Ospfareaaggregatestatus", ospfareaaggregateentry.Ospfareaaggregatestatus}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateEffect"] = types.YLeaf{"Ospfareaaggregateeffect", ospfareaaggregateentry.Ospfareaaggregateeffect}
    ospfareaaggregateentry.EntityData.Leafs["ospfAreaAggregateExtRouteTag"] = types.YLeaf{"Ospfareaaggregateextroutetag", ospfareaaggregateentry.Ospfareaaggregateextroutetag}
    return &(ospfareaaggregateentry.EntityData)
}

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect represents being advertised at all outside the area.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect string

const (
    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect_advertiseMatching OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect = "advertiseMatching"

    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect_doNotAdvertiseMatching OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect = "doNotAdvertiseMatching"
)

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype represents aggregate applies to.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype string

const (
    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype_summaryLink OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype = "summaryLink"

    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype_nssaExternalLink OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype = "nssaExternalLink"
)

// OSPFMIB_Ospflocallsdbtable
// The OSPF Process's link-local link state database
// for non-virtual links.
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for non-virtual links.  The purpose is
// to allow link-local LSAs to be displayed for each
// non-virtual interface.  This table is implemented to
// support type-9 LSAs that are defined
// in 'The OSPF Opaque LSA Option'.
type OSPFMIB_Ospflocallsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry.
    Ospflocallsdbentry []OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetEntityData() *types.CommonEntityData {
    ospflocallsdbtable.EntityData.YFilter = ospflocallsdbtable.YFilter
    ospflocallsdbtable.EntityData.YangName = "ospfLocalLsdbTable"
    ospflocallsdbtable.EntityData.BundleName = "cisco_ios_xe"
    ospflocallsdbtable.EntityData.ParentYangName = "OSPF-MIB"
    ospflocallsdbtable.EntityData.SegmentPath = "ospfLocalLsdbTable"
    ospflocallsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospflocallsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospflocallsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospflocallsdbtable.EntityData.Children = make(map[string]types.YChild)
    ospflocallsdbtable.EntityData.Children["ospfLocalLsdbEntry"] = types.YChild{"Ospflocallsdbentry", nil}
    for i := range ospflocallsdbtable.Ospflocallsdbentry {
        ospflocallsdbtable.EntityData.Children[types.GetSegmentPath(&ospflocallsdbtable.Ospflocallsdbentry[i])] = types.YChild{"Ospflocallsdbentry", &ospflocallsdbtable.Ospflocallsdbentry[i]}
    }
    ospflocallsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospflocallsdbtable.EntityData)
}

// OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry
// A single link state advertisement.
type OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflocallsdbipaddress interface{}

    // This attribute is a key. The interface index of the interface from which
    // the LSA was received if the interface is unnumbered. The type is
    // interface{} with range: 0..2147483647.
    Ospflocallsdbaddresslessif interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospflocallsdbtype.
    Ospflocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflocallsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospflocallsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospflocallsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospflocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospflocallsdbchecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    Ospflocallsdbadvertisement interface{}
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetEntityData() *types.CommonEntityData {
    ospflocallsdbentry.EntityData.YFilter = ospflocallsdbentry.YFilter
    ospflocallsdbentry.EntityData.YangName = "ospfLocalLsdbEntry"
    ospflocallsdbentry.EntityData.BundleName = "cisco_ios_xe"
    ospflocallsdbentry.EntityData.ParentYangName = "ospfLocalLsdbTable"
    ospflocallsdbentry.EntityData.SegmentPath = "ospfLocalLsdbEntry" + "[ospfLocalLsdbIpAddress='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbipaddress) + "']" + "[ospfLocalLsdbAddressLessIf='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbaddresslessif) + "']" + "[ospfLocalLsdbType='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbtype) + "']" + "[ospfLocalLsdbLsid='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdblsid) + "']" + "[ospfLocalLsdbRouterId='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbrouterid) + "']"
    ospflocallsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospflocallsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospflocallsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospflocallsdbentry.EntityData.Children = make(map[string]types.YChild)
    ospflocallsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbIpAddress"] = types.YLeaf{"Ospflocallsdbipaddress", ospflocallsdbentry.Ospflocallsdbipaddress}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbAddressLessIf"] = types.YLeaf{"Ospflocallsdbaddresslessif", ospflocallsdbentry.Ospflocallsdbaddresslessif}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbType"] = types.YLeaf{"Ospflocallsdbtype", ospflocallsdbentry.Ospflocallsdbtype}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbLsid"] = types.YLeaf{"Ospflocallsdblsid", ospflocallsdbentry.Ospflocallsdblsid}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbRouterId"] = types.YLeaf{"Ospflocallsdbrouterid", ospflocallsdbentry.Ospflocallsdbrouterid}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbSequence"] = types.YLeaf{"Ospflocallsdbsequence", ospflocallsdbentry.Ospflocallsdbsequence}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbAge"] = types.YLeaf{"Ospflocallsdbage", ospflocallsdbentry.Ospflocallsdbage}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbChecksum"] = types.YLeaf{"Ospflocallsdbchecksum", ospflocallsdbentry.Ospflocallsdbchecksum}
    ospflocallsdbentry.EntityData.Leafs["ospfLocalLsdbAdvertisement"] = types.YLeaf{"Ospflocallsdbadvertisement", ospflocallsdbentry.Ospflocallsdbadvertisement}
    return &(ospflocallsdbentry.EntityData)
}

// OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype represents advertisement format.
type OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype string

const (
    OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype_localOpaqueLink OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype = "localOpaqueLink"
)

// OSPFMIB_Ospfvirtlocallsdbtable
// The OSPF Process's link-local link state database
// for virtual links.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for virtual links.  The purpose is to
// allow link-local LSAs to be displayed for each virtual
// interface.  This table is implemented to support type-9 LSAs
// that are defined in 'The OSPF Opaque LSA Option'.
type OSPFMIB_Ospfvirtlocallsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry.
    Ospfvirtlocallsdbentry []OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetEntityData() *types.CommonEntityData {
    ospfvirtlocallsdbtable.EntityData.YFilter = ospfvirtlocallsdbtable.YFilter
    ospfvirtlocallsdbtable.EntityData.YangName = "ospfVirtLocalLsdbTable"
    ospfvirtlocallsdbtable.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtlocallsdbtable.EntityData.ParentYangName = "OSPF-MIB"
    ospfvirtlocallsdbtable.EntityData.SegmentPath = "ospfVirtLocalLsdbTable"
    ospfvirtlocallsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtlocallsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtlocallsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtlocallsdbtable.EntityData.Children = make(map[string]types.YChild)
    ospfvirtlocallsdbtable.EntityData.Children["ospfVirtLocalLsdbEntry"] = types.YChild{"Ospfvirtlocallsdbentry", nil}
    for i := range ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry {
        ospfvirtlocallsdbtable.EntityData.Children[types.GetSegmentPath(&ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry[i])] = types.YChild{"Ospfvirtlocallsdbentry", &ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry[i]}
    }
    ospfvirtlocallsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfvirtlocallsdbtable.EntityData)
}

// OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtlocallsdbtransitarea interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtlocallsdbneighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfvirtlocallsdbtype.
    Ospfvirtlocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtlocallsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfvirtlocallsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfvirtlocallsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfvirtlocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that  an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfvirtlocallsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    Ospfvirtlocallsdbadvertisement interface{}
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetEntityData() *types.CommonEntityData {
    ospfvirtlocallsdbentry.EntityData.YFilter = ospfvirtlocallsdbentry.YFilter
    ospfvirtlocallsdbentry.EntityData.YangName = "ospfVirtLocalLsdbEntry"
    ospfvirtlocallsdbentry.EntityData.BundleName = "cisco_ios_xe"
    ospfvirtlocallsdbentry.EntityData.ParentYangName = "ospfVirtLocalLsdbTable"
    ospfvirtlocallsdbentry.EntityData.SegmentPath = "ospfVirtLocalLsdbEntry" + "[ospfVirtLocalLsdbTransitArea='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtransitarea) + "']" + "[ospfVirtLocalLsdbNeighbor='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbneighbor) + "']" + "[ospfVirtLocalLsdbType='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtype) + "']" + "[ospfVirtLocalLsdbLsid='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdblsid) + "']" + "[ospfVirtLocalLsdbRouterId='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbrouterid) + "']"
    ospfvirtlocallsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfvirtlocallsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfvirtlocallsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfvirtlocallsdbentry.EntityData.Children = make(map[string]types.YChild)
    ospfvirtlocallsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbTransitArea"] = types.YLeaf{"Ospfvirtlocallsdbtransitarea", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtransitarea}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbNeighbor"] = types.YLeaf{"Ospfvirtlocallsdbneighbor", ospfvirtlocallsdbentry.Ospfvirtlocallsdbneighbor}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbType"] = types.YLeaf{"Ospfvirtlocallsdbtype", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtype}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbLsid"] = types.YLeaf{"Ospfvirtlocallsdblsid", ospfvirtlocallsdbentry.Ospfvirtlocallsdblsid}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbRouterId"] = types.YLeaf{"Ospfvirtlocallsdbrouterid", ospfvirtlocallsdbentry.Ospfvirtlocallsdbrouterid}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbSequence"] = types.YLeaf{"Ospfvirtlocallsdbsequence", ospfvirtlocallsdbentry.Ospfvirtlocallsdbsequence}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbAge"] = types.YLeaf{"Ospfvirtlocallsdbage", ospfvirtlocallsdbentry.Ospfvirtlocallsdbage}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbChecksum"] = types.YLeaf{"Ospfvirtlocallsdbchecksum", ospfvirtlocallsdbentry.Ospfvirtlocallsdbchecksum}
    ospfvirtlocallsdbentry.EntityData.Leafs["ospfVirtLocalLsdbAdvertisement"] = types.YLeaf{"Ospfvirtlocallsdbadvertisement", ospfvirtlocallsdbentry.Ospfvirtlocallsdbadvertisement}
    return &(ospfvirtlocallsdbentry.EntityData)
}

// OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype represents advertisement format.
type OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype string

const (
    OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype_localOpaqueLink OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype = "localOpaqueLink"
)

// OSPFMIB_Ospfaslsdbtable
// The OSPF Process's AS-scope LSA link state database.
// The database contains the AS-scope Link State
// Advertisements from throughout the areas that
// the device is attached to.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only AS-scope Link State
// Advertisements.  The purpose is to allow AS-scope
// LSAs to be displayed once for the router rather
// than once in each non-stub area.
type OSPFMIB_Ospfaslsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry.
    Ospfaslsdbentry []OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetEntityData() *types.CommonEntityData {
    ospfaslsdbtable.EntityData.YFilter = ospfaslsdbtable.YFilter
    ospfaslsdbtable.EntityData.YangName = "ospfAsLsdbTable"
    ospfaslsdbtable.EntityData.BundleName = "cisco_ios_xe"
    ospfaslsdbtable.EntityData.ParentYangName = "OSPF-MIB"
    ospfaslsdbtable.EntityData.SegmentPath = "ospfAsLsdbTable"
    ospfaslsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfaslsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfaslsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfaslsdbtable.EntityData.Children = make(map[string]types.YChild)
    ospfaslsdbtable.EntityData.Children["ospfAsLsdbEntry"] = types.YChild{"Ospfaslsdbentry", nil}
    for i := range ospfaslsdbtable.Ospfaslsdbentry {
        ospfaslsdbtable.EntityData.Children[types.GetSegmentPath(&ospfaslsdbtable.Ospfaslsdbentry[i])] = types.YChild{"Ospfaslsdbentry", &ospfaslsdbtable.Ospfaslsdbentry[i]}
    }
    ospfaslsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfaslsdbtable.EntityData)
}

// OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfaslsdbtype.
    Ospfaslsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address;  it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfaslsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfaslsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfaslsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfaslsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfaslsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    Ospfaslsdbadvertisement interface{}
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetEntityData() *types.CommonEntityData {
    ospfaslsdbentry.EntityData.YFilter = ospfaslsdbentry.YFilter
    ospfaslsdbentry.EntityData.YangName = "ospfAsLsdbEntry"
    ospfaslsdbentry.EntityData.BundleName = "cisco_ios_xe"
    ospfaslsdbentry.EntityData.ParentYangName = "ospfAsLsdbTable"
    ospfaslsdbentry.EntityData.SegmentPath = "ospfAsLsdbEntry" + "[ospfAsLsdbType='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdbtype) + "']" + "[ospfAsLsdbLsid='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdblsid) + "']" + "[ospfAsLsdbRouterId='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdbrouterid) + "']"
    ospfaslsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfaslsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfaslsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfaslsdbentry.EntityData.Children = make(map[string]types.YChild)
    ospfaslsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbType"] = types.YLeaf{"Ospfaslsdbtype", ospfaslsdbentry.Ospfaslsdbtype}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbLsid"] = types.YLeaf{"Ospfaslsdblsid", ospfaslsdbentry.Ospfaslsdblsid}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbRouterId"] = types.YLeaf{"Ospfaslsdbrouterid", ospfaslsdbentry.Ospfaslsdbrouterid}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbSequence"] = types.YLeaf{"Ospfaslsdbsequence", ospfaslsdbentry.Ospfaslsdbsequence}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbAge"] = types.YLeaf{"Ospfaslsdbage", ospfaslsdbentry.Ospfaslsdbage}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbChecksum"] = types.YLeaf{"Ospfaslsdbchecksum", ospfaslsdbentry.Ospfaslsdbchecksum}
    ospfaslsdbentry.EntityData.Leafs["ospfAsLsdbAdvertisement"] = types.YLeaf{"Ospfaslsdbadvertisement", ospfaslsdbentry.Ospfaslsdbadvertisement}
    return &(ospfaslsdbentry.EntityData)
}

// OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype represents advertisement format.
type OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype string

const (
    OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype_asExternalLink OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype = "asExternalLink"

    OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype_asOpaqueLink OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype = "asOpaqueLink"
)

// OSPFMIB_Ospfarealsacounttable
// This table maintains per-area, per-LSA-type counters
type OSPFMIB_Ospfarealsacounttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with a number of link advertisements  of a given type for a given
    // area. The type is slice of
    // OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry.
    Ospfarealsacountentry []OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetEntityData() *types.CommonEntityData {
    ospfarealsacounttable.EntityData.YFilter = ospfarealsacounttable.YFilter
    ospfarealsacounttable.EntityData.YangName = "ospfAreaLsaCountTable"
    ospfarealsacounttable.EntityData.BundleName = "cisco_ios_xe"
    ospfarealsacounttable.EntityData.ParentYangName = "OSPF-MIB"
    ospfarealsacounttable.EntityData.SegmentPath = "ospfAreaLsaCountTable"
    ospfarealsacounttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfarealsacounttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfarealsacounttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfarealsacounttable.EntityData.Children = make(map[string]types.YChild)
    ospfarealsacounttable.EntityData.Children["ospfAreaLsaCountEntry"] = types.YChild{"Ospfarealsacountentry", nil}
    for i := range ospfarealsacounttable.Ospfarealsacountentry {
        ospfarealsacounttable.EntityData.Children[types.GetSegmentPath(&ospfarealsacounttable.Ospfarealsacountentry[i])] = types.YChild{"Ospfarealsacountentry", &ospfarealsacounttable.Ospfarealsacountentry[i]}
    }
    ospfarealsacounttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfarealsacounttable.EntityData)
}

// OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry
// An entry with a number of link advertisements
// 
// of a given type for a given area.
type OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This entry Area ID. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfarealsacountareaid interface{}

    // This attribute is a key. This entry LSA type. The type is
    // Ospfarealsacountlsatype.
    Ospfarealsacountlsatype interface{}

    // Number of LSAs of a given type for a given area. The type is interface{}
    // with range: 0..4294967295.
    Ospfarealsacountnumber interface{}
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetEntityData() *types.CommonEntityData {
    ospfarealsacountentry.EntityData.YFilter = ospfarealsacountentry.YFilter
    ospfarealsacountentry.EntityData.YangName = "ospfAreaLsaCountEntry"
    ospfarealsacountentry.EntityData.BundleName = "cisco_ios_xe"
    ospfarealsacountentry.EntityData.ParentYangName = "ospfAreaLsaCountTable"
    ospfarealsacountentry.EntityData.SegmentPath = "ospfAreaLsaCountEntry" + "[ospfAreaLsaCountAreaId='" + fmt.Sprintf("%v", ospfarealsacountentry.Ospfarealsacountareaid) + "']" + "[ospfAreaLsaCountLsaType='" + fmt.Sprintf("%v", ospfarealsacountentry.Ospfarealsacountlsatype) + "']"
    ospfarealsacountentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfarealsacountentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfarealsacountentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfarealsacountentry.EntityData.Children = make(map[string]types.YChild)
    ospfarealsacountentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfarealsacountentry.EntityData.Leafs["ospfAreaLsaCountAreaId"] = types.YLeaf{"Ospfarealsacountareaid", ospfarealsacountentry.Ospfarealsacountareaid}
    ospfarealsacountentry.EntityData.Leafs["ospfAreaLsaCountLsaType"] = types.YLeaf{"Ospfarealsacountlsatype", ospfarealsacountentry.Ospfarealsacountlsatype}
    ospfarealsacountentry.EntityData.Leafs["ospfAreaLsaCountNumber"] = types.YLeaf{"Ospfarealsacountnumber", ospfarealsacountentry.Ospfarealsacountnumber}
    return &(ospfarealsacountentry.EntityData)
}

// OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype represents This entry LSA type.
type OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype string

const (
    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_routerLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "routerLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_networkLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "networkLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_summaryLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "summaryLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_asSummaryLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "asSummaryLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_multicastLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "multicastLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_nssaExternalLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "nssaExternalLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_areaOpaqueLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "areaOpaqueLink"
)

