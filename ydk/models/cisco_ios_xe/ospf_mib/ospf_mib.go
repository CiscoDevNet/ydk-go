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

    
    OspfGeneralGroup OSPFMIB_OspfGeneralGroup

    // Information describing the configured parameters and cumulative statistics
    // of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.
    OspfAreaTable OSPFMIB_OspfAreaTable

    // The set of metrics that will be advertised by a default Area Border Router
    // into a stub area.
    OspfStubAreaTable OSPFMIB_OspfStubAreaTable

    // The OSPF Process's link state database (LSDB). The LSDB contains the link
    // state advertisements from throughout the areas that the device is attached
    // to.
    OspfLsdbTable OSPFMIB_OspfLsdbTable

    // The Address Range Table acts as an adjunct to the Area Table.  It describes
    // those Address Range Summaries that are configured to be propagated from an
    // Area to reduce the amount of information about it that is known beyond its
    // borders.  It contains a set of IP address ranges specified by an IP
    // address/IP network mask pair. For example, class B address range of X.X.X.X
    // with a network mask of 255.255.0.0 includes all IP addresses from X.X.0.0
    // to X.X.255.255.  Note that this table is obsoleted and is replaced by the
    // Area Aggregate Table.
    OspfAreaRangeTable OSPFMIB_OspfAreaRangeTable

    // The Host/Metric Table indicates what hosts are directly  attached to the
    // router, what metrics and types of service should be advertised for them,
    // and what areas they are found within.
    OspfHostTable OSPFMIB_OspfHostTable

    // The OSPF Interface Table describes the interfaces from the viewpoint of
    // OSPF. It augments the ipAddrTable with OSPF specific information.
    OspfIfTable OSPFMIB_OspfIfTable

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
    OspfIfMetricTable OSPFMIB_OspfIfMetricTable

    // Information about this router's virtual interfaces that the OSPF Process is
    // configured to carry on.
    OspfVirtIfTable OSPFMIB_OspfVirtIfTable

    // A table describing all non-virtual neighbors in the locality of the OSPF
    // router.
    OspfNbrTable OSPFMIB_OspfNbrTable

    // This table describes all virtual neighbors. Since virtual links are
    // configured in the Virtual Interface Table, this table is read-only.
    OspfVirtNbrTable OSPFMIB_OspfVirtNbrTable

    // The OSPF Process's external LSA link state database.  This table is
    // identical to the OSPF LSDB Table in format, but contains only external link
    // state advertisements.  The purpose is to allow external  LSAs to be
    // displayed once for the router rather than once in each non-stub area.  Note
    // that external LSAs are also in the AS-scope link state database.
    OspfExtLsdbTable OSPFMIB_OspfExtLsdbTable

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
    OspfAreaAggregateTable OSPFMIB_OspfAreaAggregateTable

    // The OSPF Process's link-local link state database for non-virtual links.
    // This table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for non-virtual links.  The purpose is
    // to allow link-local LSAs to be displayed for each non-virtual interface. 
    // This table is implemented to support type-9 LSAs that are defined in 'The
    // OSPF Opaque LSA Option'.
    OspfLocalLsdbTable OSPFMIB_OspfLocalLsdbTable

    // The OSPF Process's link-local link state database for virtual links.  This
    // table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for virtual links.  The purpose is to
    // allow link-local LSAs to be displayed for each virtual interface.  This
    // table is implemented to support type-9 LSAs that are defined in 'The OSPF
    // Opaque LSA Option'.
    OspfVirtLocalLsdbTable OSPFMIB_OspfVirtLocalLsdbTable

    // The OSPF Process's AS-scope LSA link state database. The database contains
    // the AS-scope Link State Advertisements from throughout the areas that the
    // device is attached to.  This table is identical to the OSPF LSDB Table in
    // format, but contains only AS-scope Link State Advertisements.  The purpose
    // is to allow AS-scope LSAs to be displayed once for the router rather than
    // once in each non-stub area.
    OspfAsLsdbTable OSPFMIB_OspfAsLsdbTable

    // This table maintains per-area, per-LSA-type counters.
    OspfAreaLsaCountTable OSPFMIB_OspfAreaLsaCountTable
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

    oSPFMIB.EntityData.Children = types.NewOrderedMap()
    oSPFMIB.EntityData.Children.Append("ospfGeneralGroup", types.YChild{"OspfGeneralGroup", &oSPFMIB.OspfGeneralGroup})
    oSPFMIB.EntityData.Children.Append("ospfAreaTable", types.YChild{"OspfAreaTable", &oSPFMIB.OspfAreaTable})
    oSPFMIB.EntityData.Children.Append("ospfStubAreaTable", types.YChild{"OspfStubAreaTable", &oSPFMIB.OspfStubAreaTable})
    oSPFMIB.EntityData.Children.Append("ospfLsdbTable", types.YChild{"OspfLsdbTable", &oSPFMIB.OspfLsdbTable})
    oSPFMIB.EntityData.Children.Append("ospfAreaRangeTable", types.YChild{"OspfAreaRangeTable", &oSPFMIB.OspfAreaRangeTable})
    oSPFMIB.EntityData.Children.Append("ospfHostTable", types.YChild{"OspfHostTable", &oSPFMIB.OspfHostTable})
    oSPFMIB.EntityData.Children.Append("ospfIfTable", types.YChild{"OspfIfTable", &oSPFMIB.OspfIfTable})
    oSPFMIB.EntityData.Children.Append("ospfIfMetricTable", types.YChild{"OspfIfMetricTable", &oSPFMIB.OspfIfMetricTable})
    oSPFMIB.EntityData.Children.Append("ospfVirtIfTable", types.YChild{"OspfVirtIfTable", &oSPFMIB.OspfVirtIfTable})
    oSPFMIB.EntityData.Children.Append("ospfNbrTable", types.YChild{"OspfNbrTable", &oSPFMIB.OspfNbrTable})
    oSPFMIB.EntityData.Children.Append("ospfVirtNbrTable", types.YChild{"OspfVirtNbrTable", &oSPFMIB.OspfVirtNbrTable})
    oSPFMIB.EntityData.Children.Append("ospfExtLsdbTable", types.YChild{"OspfExtLsdbTable", &oSPFMIB.OspfExtLsdbTable})
    oSPFMIB.EntityData.Children.Append("ospfAreaAggregateTable", types.YChild{"OspfAreaAggregateTable", &oSPFMIB.OspfAreaAggregateTable})
    oSPFMIB.EntityData.Children.Append("ospfLocalLsdbTable", types.YChild{"OspfLocalLsdbTable", &oSPFMIB.OspfLocalLsdbTable})
    oSPFMIB.EntityData.Children.Append("ospfVirtLocalLsdbTable", types.YChild{"OspfVirtLocalLsdbTable", &oSPFMIB.OspfVirtLocalLsdbTable})
    oSPFMIB.EntityData.Children.Append("ospfAsLsdbTable", types.YChild{"OspfAsLsdbTable", &oSPFMIB.OspfAsLsdbTable})
    oSPFMIB.EntityData.Children.Append("ospfAreaLsaCountTable", types.YChild{"OspfAreaLsaCountTable", &oSPFMIB.OspfAreaLsaCountTable})
    oSPFMIB.EntityData.Leafs = types.NewOrderedMap()

    oSPFMIB.EntityData.YListKeys = []string {}

    return &(oSPFMIB.EntityData)
}

// OSPFMIB_OspfGeneralGroup
type OSPFMIB_OspfGeneralGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A 32-bit integer uniquely identifying the router in the Autonomous System.
    // By convention, to ensure uniqueness, this should default to the value of
    // one of the router's IP interface addresses.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfRouterId interface{}

    // The administrative status of OSPF in the router.  The value 'enabled'
    // denotes that the OSPF Process is active on at least one interface;
    // 'disabled' disables it on all interfaces.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is Status.
    OspfAdminStat interface{}

    // The current version number of the OSPF protocol is 2. The type is
    // OspfVersionNumber.
    OspfVersionNumber interface{}

    // A flag to note whether this router is an Area Border Router. The type is
    // bool.
    OspfAreaBdrRtrStatus interface{}

    // A flag to note whether this router is configured as an Autonomous System
    // Border Router.  This object is persistent and when written the entity
    // SHOULD save the change to non-volatile storage. The type is bool.
    OspfASBdrRtrStatus interface{}

    // The number of external (LS type-5) link state advertisements in the link
    // state database. The type is interface{} with range: 0..4294967295.
    OspfExternLsaCount interface{}

    // The 32-bit sum of the LS checksums of the external link state
    // advertisements contained in the link state database.  This sum can be used
    // to determine if there has been a change in a router's link state database
    // and to compare the link state database of two routers.  The value should be
    // treated as unsigned when comparing two sums of checksums. The type is
    // interface{} with range: -2147483648..2147483647.
    OspfExternLsaCksumSum interface{}

    // The router's support for type-of-service routing.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    OspfTOSSupport interface{}

    // The number of new link state advertisements that have been originated. 
    // This number is incremented each time the router originates a new LSA. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    OspfOriginateNewLsas interface{}

    // The number of link state advertisements received that are determined to be
    // new instantiations. This number does not include newer instantiations of
    // self-originated link state advertisements.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ospfDiscontinuityTime. The type
    // is interface{} with range: 0..4294967295.
    OspfRxNewLsas interface{}

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
    OspfExtLsdbLimit interface{}

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
    OspfMulticastExtensions interface{}

    // The number of seconds that, after entering OverflowState, a router will
    // attempt to leave OverflowState.  This allows the router to again originate
    // non-default AS-external LSAs.  When set to 0, the router will not leave
    // overflow state until restarted.  This object is persistent and when written
    // the entity SHOULD save the change to non-volatile storage. The type is
    // interface{} with range: 0..2147483647.
    OspfExitOverflowInterval interface{}

    // The router's support for demand routing. This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is bool.
    OspfDemandExtensions interface{}

    // Indicates metrics used to choose among multiple AS-external LSAs.  When
    // RFC1583Compatibility is set to enabled, only cost will be used when
    // choosing among multiple AS-external LSAs advertising the same destination. 
    // When RFC1583Compatibility is set to disabled, preference will be driven
    // first by type of path using cost only to break ties.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    OspfRFC1583Compatibility interface{}

    // The router's support for Opaque LSA types. The type is bool.
    OspfOpaqueLsaSupport interface{}

    // Reference bandwidth in kilobits/second for  calculating default interface
    // metrics.  The default value is 100,000 KBPS (100 MBPS).  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 0..4294967295.
    // Units are kilobits per second.
    OspfReferenceBandwidth interface{}

    // The router's support for OSPF graceful restart. Options include: no restart
    // support, only planned restarts, or both planned and unplanned restarts. 
    // This object is persistent and when written the entity SHOULD save the
    // change to non-volatile storage. The type is OspfRestartSupport.
    OspfRestartSupport interface{}

    // Configured OSPF graceful restart timeout interval.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 1..1800. Units
    // are seconds.
    OspfRestartInterval interface{}

    // Indicates if strict LSA checking is enabled for graceful restart.  This
    // object is persistent and when written the entity SHOULD save the change to
    // non-volatile  storage. The type is bool.
    OspfRestartStrictLsaChecking interface{}

    // Current status of OSPF graceful restart. The type is OspfRestartStatus.
    OspfRestartStatus interface{}

    // Remaining time in current OSPF graceful restart interval. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    OspfRestartAge interface{}

    // Describes the outcome of the last attempt at a graceful restart.  If the
    // value is 'none', no restart has yet been attempted.  If the value is
    // 'inProgress', a restart attempt is currently underway. The type is
    // OspfRestartExitReason.
    OspfRestartExitReason interface{}

    // The number of AS-scope link state advertisements in the AS-scope link state
    // database. The type is interface{} with range: 0..4294967295.
    OspfAsLsaCount interface{}

    // The 32-bit unsigned sum of the LS checksums of the AS link state
    // advertisements contained in the AS-scope link state database.  This sum can
    // be used to determine if there has been a change in a router's AS-scope link
    // state database, and to compare the AS-scope link state database of two
    // routers. The type is interface{} with range: 0..4294967295.
    OspfAsLsaCksumSum interface{}

    // The router's support for stub router functionality. The type is bool.
    OspfStubRouterSupport interface{}

    // This object controls the advertisement of stub router LSAs by the router. 
    // The value doNotAdvertise will result in the advertisement of a standard
    // router LSA and is the default value.  This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is OspfStubRouterAdvertisement.
    OspfStubRouterAdvertisement interface{}

    // The value of sysUpTime on the most recent occasion at which any one of this
    // MIB's counters suffered a discontinuity.  If no such discontinuities have
    // occurred since the last re-initialization of the local management
    // subsystem, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    OspfDiscontinuityTime interface{}
}

func (ospfGeneralGroup *OSPFMIB_OspfGeneralGroup) GetEntityData() *types.CommonEntityData {
    ospfGeneralGroup.EntityData.YFilter = ospfGeneralGroup.YFilter
    ospfGeneralGroup.EntityData.YangName = "ospfGeneralGroup"
    ospfGeneralGroup.EntityData.BundleName = "cisco_ios_xe"
    ospfGeneralGroup.EntityData.ParentYangName = "OSPF-MIB"
    ospfGeneralGroup.EntityData.SegmentPath = "ospfGeneralGroup"
    ospfGeneralGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfGeneralGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfGeneralGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfGeneralGroup.EntityData.Children = types.NewOrderedMap()
    ospfGeneralGroup.EntityData.Leafs = types.NewOrderedMap()
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRouterId", types.YLeaf{"OspfRouterId", ospfGeneralGroup.OspfRouterId})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfAdminStat", types.YLeaf{"OspfAdminStat", ospfGeneralGroup.OspfAdminStat})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfVersionNumber", types.YLeaf{"OspfVersionNumber", ospfGeneralGroup.OspfVersionNumber})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfAreaBdrRtrStatus", types.YLeaf{"OspfAreaBdrRtrStatus", ospfGeneralGroup.OspfAreaBdrRtrStatus})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfASBdrRtrStatus", types.YLeaf{"OspfASBdrRtrStatus", ospfGeneralGroup.OspfASBdrRtrStatus})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfExternLsaCount", types.YLeaf{"OspfExternLsaCount", ospfGeneralGroup.OspfExternLsaCount})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfExternLsaCksumSum", types.YLeaf{"OspfExternLsaCksumSum", ospfGeneralGroup.OspfExternLsaCksumSum})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfTOSSupport", types.YLeaf{"OspfTOSSupport", ospfGeneralGroup.OspfTOSSupport})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfOriginateNewLsas", types.YLeaf{"OspfOriginateNewLsas", ospfGeneralGroup.OspfOriginateNewLsas})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRxNewLsas", types.YLeaf{"OspfRxNewLsas", ospfGeneralGroup.OspfRxNewLsas})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfExtLsdbLimit", types.YLeaf{"OspfExtLsdbLimit", ospfGeneralGroup.OspfExtLsdbLimit})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfMulticastExtensions", types.YLeaf{"OspfMulticastExtensions", ospfGeneralGroup.OspfMulticastExtensions})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfExitOverflowInterval", types.YLeaf{"OspfExitOverflowInterval", ospfGeneralGroup.OspfExitOverflowInterval})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfDemandExtensions", types.YLeaf{"OspfDemandExtensions", ospfGeneralGroup.OspfDemandExtensions})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRFC1583Compatibility", types.YLeaf{"OspfRFC1583Compatibility", ospfGeneralGroup.OspfRFC1583Compatibility})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfOpaqueLsaSupport", types.YLeaf{"OspfOpaqueLsaSupport", ospfGeneralGroup.OspfOpaqueLsaSupport})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfReferenceBandwidth", types.YLeaf{"OspfReferenceBandwidth", ospfGeneralGroup.OspfReferenceBandwidth})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartSupport", types.YLeaf{"OspfRestartSupport", ospfGeneralGroup.OspfRestartSupport})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartInterval", types.YLeaf{"OspfRestartInterval", ospfGeneralGroup.OspfRestartInterval})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartStrictLsaChecking", types.YLeaf{"OspfRestartStrictLsaChecking", ospfGeneralGroup.OspfRestartStrictLsaChecking})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartStatus", types.YLeaf{"OspfRestartStatus", ospfGeneralGroup.OspfRestartStatus})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartAge", types.YLeaf{"OspfRestartAge", ospfGeneralGroup.OspfRestartAge})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfRestartExitReason", types.YLeaf{"OspfRestartExitReason", ospfGeneralGroup.OspfRestartExitReason})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfAsLsaCount", types.YLeaf{"OspfAsLsaCount", ospfGeneralGroup.OspfAsLsaCount})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfAsLsaCksumSum", types.YLeaf{"OspfAsLsaCksumSum", ospfGeneralGroup.OspfAsLsaCksumSum})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfStubRouterSupport", types.YLeaf{"OspfStubRouterSupport", ospfGeneralGroup.OspfStubRouterSupport})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfStubRouterAdvertisement", types.YLeaf{"OspfStubRouterAdvertisement", ospfGeneralGroup.OspfStubRouterAdvertisement})
    ospfGeneralGroup.EntityData.Leafs.Append("ospfDiscontinuityTime", types.YLeaf{"OspfDiscontinuityTime", ospfGeneralGroup.OspfDiscontinuityTime})

    ospfGeneralGroup.EntityData.YListKeys = []string {}

    return &(ospfGeneralGroup.EntityData)
}

// OSPFMIB_OspfGeneralGroup_OspfRestartExitReason represents a restart attempt is currently underway.
type OSPFMIB_OspfGeneralGroup_OspfRestartExitReason string

const (
    OSPFMIB_OspfGeneralGroup_OspfRestartExitReason_none OSPFMIB_OspfGeneralGroup_OspfRestartExitReason = "none"

    OSPFMIB_OspfGeneralGroup_OspfRestartExitReason_inProgress OSPFMIB_OspfGeneralGroup_OspfRestartExitReason = "inProgress"

    OSPFMIB_OspfGeneralGroup_OspfRestartExitReason_completed OSPFMIB_OspfGeneralGroup_OspfRestartExitReason = "completed"

    OSPFMIB_OspfGeneralGroup_OspfRestartExitReason_timedOut OSPFMIB_OspfGeneralGroup_OspfRestartExitReason = "timedOut"

    OSPFMIB_OspfGeneralGroup_OspfRestartExitReason_topologyChanged OSPFMIB_OspfGeneralGroup_OspfRestartExitReason = "topologyChanged"
)

// OSPFMIB_OspfGeneralGroup_OspfRestartStatus represents Current status of OSPF graceful restart.
type OSPFMIB_OspfGeneralGroup_OspfRestartStatus string

const (
    OSPFMIB_OspfGeneralGroup_OspfRestartStatus_notRestarting OSPFMIB_OspfGeneralGroup_OspfRestartStatus = "notRestarting"

    OSPFMIB_OspfGeneralGroup_OspfRestartStatus_plannedRestart OSPFMIB_OspfGeneralGroup_OspfRestartStatus = "plannedRestart"

    OSPFMIB_OspfGeneralGroup_OspfRestartStatus_unplannedRestart OSPFMIB_OspfGeneralGroup_OspfRestartStatus = "unplannedRestart"
)

// OSPFMIB_OspfGeneralGroup_OspfRestartSupport represents storage.
type OSPFMIB_OspfGeneralGroup_OspfRestartSupport string

const (
    OSPFMIB_OspfGeneralGroup_OspfRestartSupport_none OSPFMIB_OspfGeneralGroup_OspfRestartSupport = "none"

    OSPFMIB_OspfGeneralGroup_OspfRestartSupport_plannedOnly OSPFMIB_OspfGeneralGroup_OspfRestartSupport = "plannedOnly"

    OSPFMIB_OspfGeneralGroup_OspfRestartSupport_plannedAndUnplanned OSPFMIB_OspfGeneralGroup_OspfRestartSupport = "plannedAndUnplanned"
)

// OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement represents storage.
type OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement string

const (
    OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement_doNotAdvertise OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement = "doNotAdvertise"

    OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement_advertise OSPFMIB_OspfGeneralGroup_OspfStubRouterAdvertisement = "advertise"
)

// OSPFMIB_OspfGeneralGroup_OspfVersionNumber represents The current version number of the OSPF protocol is 2.
type OSPFMIB_OspfGeneralGroup_OspfVersionNumber string

const (
    OSPFMIB_OspfGeneralGroup_OspfVersionNumber_version2 OSPFMIB_OspfGeneralGroup_OspfVersionNumber = "version2"
)

// OSPFMIB_OspfAreaTable
// Information describing the configured parameters and
// cumulative statistics of the router's attached areas.
// The interfaces and virtual links are configured
// as part of these areas.  Area 0.0.0.0, by definition,
// is the backbone area.
type OSPFMIB_OspfAreaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information describing the configured parameters and cumulative statistics
    // of one of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_OspfAreaTable_OspfAreaEntry.
    OspfAreaEntry []*OSPFMIB_OspfAreaTable_OspfAreaEntry
}

func (ospfAreaTable *OSPFMIB_OspfAreaTable) GetEntityData() *types.CommonEntityData {
    ospfAreaTable.EntityData.YFilter = ospfAreaTable.YFilter
    ospfAreaTable.EntityData.YangName = "ospfAreaTable"
    ospfAreaTable.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfAreaTable.EntityData.SegmentPath = "ospfAreaTable"
    ospfAreaTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaTable.EntityData.Children = types.NewOrderedMap()
    ospfAreaTable.EntityData.Children.Append("ospfAreaEntry", types.YChild{"OspfAreaEntry", nil})
    for i := range ospfAreaTable.OspfAreaEntry {
        ospfAreaTable.EntityData.Children.Append(types.GetSegmentPath(ospfAreaTable.OspfAreaEntry[i]), types.YChild{"OspfAreaEntry", ospfAreaTable.OspfAreaEntry[i]})
    }
    ospfAreaTable.EntityData.Leafs = types.NewOrderedMap()

    ospfAreaTable.EntityData.YListKeys = []string {}

    return &(ospfAreaTable.EntityData)
}

// OSPFMIB_OspfAreaTable_OspfAreaEntry
// Information describing the configured parameters and
// cumulative statistics of one of the router's attached areas.
// The interfaces and virtual links are configured as part of
// these areas.  Area 0.0.0.0, by definition, is the backbone
// area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfAreaTable_OspfAreaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A 32-bit integer uniquely identifying an area.
    // Area ID 0.0.0.0 is used for the OSPF backbone. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaId interface{}

    // The authentication type specified for an area. The type is
    // OspfAuthenticationType.
    OspfAuthType interface{}

    // Indicates if an area is a stub area, NSSA, or standard area.  Type-5
    // AS-external LSAs and type-11 Opaque LSAs are not imported into stub areas
    // or NSSAs.  NSSAs import AS-external data as type-7 LSAs. The type is
    // OspfImportAsExtern.
    OspfImportAsExtern interface{}

    // The number of times that the intra-area route table has been calculated
    // using this area's link state database.  This is typically done using
    // Dijkstra's algorithm.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    OspfSpfRuns interface{}

    // The total number of Area Border Routers reachable within this area.  This
    // is initially zero and is calculated in each Shortest Path First (SPF) pass.
    // The type is interface{} with range: 0..4294967295.
    OspfAreaBdrRtrCount interface{}

    // The total number of Autonomous System Border Routers reachable within this
    // area.  This is initially zero and is calculated in each SPF pass. The type
    // is interface{} with range: 0..4294967295.
    OspfAsBdrRtrCount interface{}

    // The total number of link state advertisements in this area's link state
    // database, excluding AS-external LSAs. The type is interface{} with range:
    // 0..4294967295.
    OspfAreaLsaCount interface{}

    // The 32-bit sum of the link state advertisements' LS checksums contained in
    // this area's link state database.  This sum excludes external (LS type-5)
    // link state advertisements. The sum can be used to determine if there has
    // been a change in a router's link state database, and to compare the link
    // state database of two routers.  The value should be treated as unsigned
    // when comparing two sums of checksums. The type is interface{} with range:
    // -2147483648..2147483647.
    OspfAreaLsaCksumSum interface{}

    // The variable ospfAreaSummary controls the import of summary LSAs into stub
    // and NSSA areas. It has no effect on other areas.  If it is noAreaSummary,
    // the router will not originate summary LSAs into the stub or NSSA area. It
    // will rely entirely on its default route.  If it is sendAreaSummary, the
    // router will both summarize and propagate summary LSAs. The type is
    // OspfAreaSummary.
    OspfAreaSummary interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfAreaStatus interface{}

    // Indicates an NSSA border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is OspfAreaNssaTranslatorRole.
    OspfAreaNssaTranslatorRole interface{}

    // Indicates if and how an NSSA border router is performing NSSA translation
    // of type-7 LSAs into type-5  LSAs.  When this object is set to enabled, the
    // NSSA Border router's OspfAreaNssaExtTranslatorRole has been set to always. 
    // When this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA border router is NOT translating type-7 LSAs into type-5.
    // The type is OspfAreaNssaTranslatorState.
    OspfAreaNssaTranslatorState interface{}

    // The number of seconds after an elected translator determines its services
    // are no longer required, that it should continue to perform its translation
    // duties. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    OspfAreaNssaTranslatorStabilityInterval interface{}

    // Indicates the number of translator state changes that have occurred since
    // the last boot-up.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    OspfAreaNssaTranslatorEvents interface{}

    // The total number of Opaque Area and AS link-state  advertisements in the
    // link state database of this area. The type is interface{} with range:
    // 0..4294967295.
    CospfOpaqueAreaLsaCount interface{}

    // The 32-bit unsigned sum of the Opaque Area and AS  link-state
    // advertisements' LS checksums contained  link state database of this area. 
    // The sum can be  used to determine if there has been a change in the  link
    // state database for Opaque Area and AS link-state advertisements. The type
    // is interface{} with range: 0..4294967295.
    CospfOpaqueAreaLsaCksumSum interface{}

    // Indicates an NSSA Border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is CospfAreaNssaTranslatorRole.
    CospfAreaNssaTranslatorRole interface{}

    // Indicates if and how an NSSA Border router is performing NSSA translation
    // of type-7 LSAs into type-5 LSAs. When this object set to enabled, the NSSA
    // Border router's cospfAreaNssaExtTranslatorRole has been set to always. When
    // this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA Border router is NOT translating type-7 LSAs into type-5.
    // The type is CospfAreaNssaTranslatorState.
    CospfAreaNssaTranslatorState interface{}

    // Indicates the number of Translator State changes that have occurred since
    // the last boot-up. The type is interface{} with range: 0..4294967295.
    CospfAreaNssaTranslatorEvents interface{}
}

func (ospfAreaEntry *OSPFMIB_OspfAreaTable_OspfAreaEntry) GetEntityData() *types.CommonEntityData {
    ospfAreaEntry.EntityData.YFilter = ospfAreaEntry.YFilter
    ospfAreaEntry.EntityData.YangName = "ospfAreaEntry"
    ospfAreaEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaEntry.EntityData.ParentYangName = "ospfAreaTable"
    ospfAreaEntry.EntityData.SegmentPath = "ospfAreaEntry" + types.AddKeyToken(ospfAreaEntry.OspfAreaId, "ospfAreaId")
    ospfAreaEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaEntry.EntityData.Children = types.NewOrderedMap()
    ospfAreaEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaId", types.YLeaf{"OspfAreaId", ospfAreaEntry.OspfAreaId})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAuthType", types.YLeaf{"OspfAuthType", ospfAreaEntry.OspfAuthType})
    ospfAreaEntry.EntityData.Leafs.Append("ospfImportAsExtern", types.YLeaf{"OspfImportAsExtern", ospfAreaEntry.OspfImportAsExtern})
    ospfAreaEntry.EntityData.Leafs.Append("ospfSpfRuns", types.YLeaf{"OspfSpfRuns", ospfAreaEntry.OspfSpfRuns})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaBdrRtrCount", types.YLeaf{"OspfAreaBdrRtrCount", ospfAreaEntry.OspfAreaBdrRtrCount})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAsBdrRtrCount", types.YLeaf{"OspfAsBdrRtrCount", ospfAreaEntry.OspfAsBdrRtrCount})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaLsaCount", types.YLeaf{"OspfAreaLsaCount", ospfAreaEntry.OspfAreaLsaCount})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaLsaCksumSum", types.YLeaf{"OspfAreaLsaCksumSum", ospfAreaEntry.OspfAreaLsaCksumSum})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaSummary", types.YLeaf{"OspfAreaSummary", ospfAreaEntry.OspfAreaSummary})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaStatus", types.YLeaf{"OspfAreaStatus", ospfAreaEntry.OspfAreaStatus})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaNssaTranslatorRole", types.YLeaf{"OspfAreaNssaTranslatorRole", ospfAreaEntry.OspfAreaNssaTranslatorRole})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaNssaTranslatorState", types.YLeaf{"OspfAreaNssaTranslatorState", ospfAreaEntry.OspfAreaNssaTranslatorState})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaNssaTranslatorStabilityInterval", types.YLeaf{"OspfAreaNssaTranslatorStabilityInterval", ospfAreaEntry.OspfAreaNssaTranslatorStabilityInterval})
    ospfAreaEntry.EntityData.Leafs.Append("ospfAreaNssaTranslatorEvents", types.YLeaf{"OspfAreaNssaTranslatorEvents", ospfAreaEntry.OspfAreaNssaTranslatorEvents})
    ospfAreaEntry.EntityData.Leafs.Append("cospfOpaqueAreaLsaCount", types.YLeaf{"CospfOpaqueAreaLsaCount", ospfAreaEntry.CospfOpaqueAreaLsaCount})
    ospfAreaEntry.EntityData.Leafs.Append("cospfOpaqueAreaLsaCksumSum", types.YLeaf{"CospfOpaqueAreaLsaCksumSum", ospfAreaEntry.CospfOpaqueAreaLsaCksumSum})
    ospfAreaEntry.EntityData.Leafs.Append("cospfAreaNssaTranslatorRole", types.YLeaf{"CospfAreaNssaTranslatorRole", ospfAreaEntry.CospfAreaNssaTranslatorRole})
    ospfAreaEntry.EntityData.Leafs.Append("cospfAreaNssaTranslatorState", types.YLeaf{"CospfAreaNssaTranslatorState", ospfAreaEntry.CospfAreaNssaTranslatorState})
    ospfAreaEntry.EntityData.Leafs.Append("cospfAreaNssaTranslatorEvents", types.YLeaf{"CospfAreaNssaTranslatorEvents", ospfAreaEntry.CospfAreaNssaTranslatorEvents})

    ospfAreaEntry.EntityData.YListKeys = []string {"OspfAreaId"}

    return &(ospfAreaEntry.EntityData)
}

// OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole represents type-5 LSAs.
type OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole_always OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole = "always"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole_candidate OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorRole = "candidate"
)

// OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState represents Border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState_enabled OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState = "enabled"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState_elected OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState = "elected"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState_disabled OSPFMIB_OspfAreaTable_OspfAreaEntry_CospfAreaNssaTranslatorState = "disabled"
)

// OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole represents type-5 LSAs.
type OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole_always OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole = "always"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole_candidate OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorRole = "candidate"
)

// OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState represents border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState_enabled OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState = "enabled"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState_elected OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState = "elected"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState_disabled OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaNssaTranslatorState = "disabled"
)

// OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary represents summarize and propagate summary LSAs.
type OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary_noAreaSummary OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary = "noAreaSummary"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary_sendAreaSummary OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfAreaSummary = "sendAreaSummary"
)

// OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern represents AS-external data as type-7 LSAs
type OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern string

const (
    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern_importExternal OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern = "importExternal"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern_importNoExternal OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern = "importNoExternal"

    OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern_importNssa OSPFMIB_OspfAreaTable_OspfAreaEntry_OspfImportAsExtern = "importNssa"
)

// OSPFMIB_OspfStubAreaTable
// The set of metrics that will be advertised
// by a default Area Border Router into a stub area.
type OSPFMIB_OspfStubAreaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The metric for a given Type of Service that will be advertised by a default
    // Area Border Router into a stub area.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry.
    OspfStubAreaEntry []*OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry
}

func (ospfStubAreaTable *OSPFMIB_OspfStubAreaTable) GetEntityData() *types.CommonEntityData {
    ospfStubAreaTable.EntityData.YFilter = ospfStubAreaTable.YFilter
    ospfStubAreaTable.EntityData.YangName = "ospfStubAreaTable"
    ospfStubAreaTable.EntityData.BundleName = "cisco_ios_xe"
    ospfStubAreaTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfStubAreaTable.EntityData.SegmentPath = "ospfStubAreaTable"
    ospfStubAreaTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfStubAreaTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfStubAreaTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfStubAreaTable.EntityData.Children = types.NewOrderedMap()
    ospfStubAreaTable.EntityData.Children.Append("ospfStubAreaEntry", types.YChild{"OspfStubAreaEntry", nil})
    for i := range ospfStubAreaTable.OspfStubAreaEntry {
        ospfStubAreaTable.EntityData.Children.Append(types.GetSegmentPath(ospfStubAreaTable.OspfStubAreaEntry[i]), types.YChild{"OspfStubAreaEntry", ospfStubAreaTable.OspfStubAreaEntry[i]})
    }
    ospfStubAreaTable.EntityData.Leafs = types.NewOrderedMap()

    ospfStubAreaTable.EntityData.YListKeys = []string {}

    return &(ospfStubAreaTable.EntityData)
}

// OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry
// The metric for a given Type of Service that
// will be advertised by a default Area Border
// Router into a stub area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier for the stub area.  On
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfStubAreaId interface{}

    // This attribute is a key. The Type of Service associated with the metric. 
    // On creation, this can be derived from  the instance. The type is
    // interface{} with range: 0..30.
    OspfStubTOS interface{}

    // The metric value applied at the indicated Type of Service.  By default,
    // this equals the least metric at the Type of Service among the interfaces to
    // other areas. The type is interface{} with range: 0..16777215.
    OspfStubMetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfStubStatus interface{}

    // This variable displays the type of metric advertised as a default route.
    // The type is OspfStubMetricType.
    OspfStubMetricType interface{}
}

func (ospfStubAreaEntry *OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry) GetEntityData() *types.CommonEntityData {
    ospfStubAreaEntry.EntityData.YFilter = ospfStubAreaEntry.YFilter
    ospfStubAreaEntry.EntityData.YangName = "ospfStubAreaEntry"
    ospfStubAreaEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfStubAreaEntry.EntityData.ParentYangName = "ospfStubAreaTable"
    ospfStubAreaEntry.EntityData.SegmentPath = "ospfStubAreaEntry" + types.AddKeyToken(ospfStubAreaEntry.OspfStubAreaId, "ospfStubAreaId") + types.AddKeyToken(ospfStubAreaEntry.OspfStubTOS, "ospfStubTOS")
    ospfStubAreaEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfStubAreaEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfStubAreaEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfStubAreaEntry.EntityData.Children = types.NewOrderedMap()
    ospfStubAreaEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfStubAreaEntry.EntityData.Leafs.Append("ospfStubAreaId", types.YLeaf{"OspfStubAreaId", ospfStubAreaEntry.OspfStubAreaId})
    ospfStubAreaEntry.EntityData.Leafs.Append("ospfStubTOS", types.YLeaf{"OspfStubTOS", ospfStubAreaEntry.OspfStubTOS})
    ospfStubAreaEntry.EntityData.Leafs.Append("ospfStubMetric", types.YLeaf{"OspfStubMetric", ospfStubAreaEntry.OspfStubMetric})
    ospfStubAreaEntry.EntityData.Leafs.Append("ospfStubStatus", types.YLeaf{"OspfStubStatus", ospfStubAreaEntry.OspfStubStatus})
    ospfStubAreaEntry.EntityData.Leafs.Append("ospfStubMetricType", types.YLeaf{"OspfStubMetricType", ospfStubAreaEntry.OspfStubMetricType})

    ospfStubAreaEntry.EntityData.YListKeys = []string {"OspfStubAreaId", "OspfStubTOS"}

    return &(ospfStubAreaEntry.EntityData)
}

// OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType represents advertised as a default route.
type OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType string

const (
    OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType_ospfMetric OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType = "ospfMetric"

    OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType_comparableCost OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType = "comparableCost"

    OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType_nonComparable OSPFMIB_OspfStubAreaTable_OspfStubAreaEntry_OspfStubMetricType = "nonComparable"
)

// OSPFMIB_OspfLsdbTable
// The OSPF Process's link state database (LSDB).
// The LSDB contains the link state advertisements
// from throughout the areas that the device is attached to.
type OSPFMIB_OspfLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_OspfLsdbTable_OspfLsdbEntry.
    OspfLsdbEntry []*OSPFMIB_OspfLsdbTable_OspfLsdbEntry
}

func (ospfLsdbTable *OSPFMIB_OspfLsdbTable) GetEntityData() *types.CommonEntityData {
    ospfLsdbTable.EntityData.YFilter = ospfLsdbTable.YFilter
    ospfLsdbTable.EntityData.YangName = "ospfLsdbTable"
    ospfLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    ospfLsdbTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfLsdbTable.EntityData.SegmentPath = "ospfLsdbTable"
    ospfLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfLsdbTable.EntityData.Children = types.NewOrderedMap()
    ospfLsdbTable.EntityData.Children.Append("ospfLsdbEntry", types.YChild{"OspfLsdbEntry", nil})
    for i := range ospfLsdbTable.OspfLsdbEntry {
        ospfLsdbTable.EntityData.Children.Append(types.GetSegmentPath(ospfLsdbTable.OspfLsdbEntry[i]), types.YChild{"OspfLsdbEntry", ospfLsdbTable.OspfLsdbEntry[i]})
    }
    ospfLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    ospfLsdbTable.EntityData.YListKeys = []string {}

    return &(ospfLsdbTable.EntityData)
}

// OSPFMIB_OspfLsdbTable_OspfLsdbEntry
// A single link state advertisement.
type OSPFMIB_OspfLsdbTable_OspfLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier of the area from which the
    // LSA was received. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLsdbAreaId interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format.  Note: External link
    // state advertisements are permitted for backward compatibility, but should
    // be displayed in the ospfAsLsdbTable rather than here. The type is
    // OspfLsdbType.
    OspfLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLsdbLsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate Link State Advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    OspfLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    OspfLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless  datagrams;
    // it is commonly referred to as the Fletcher checksum. The type is
    // interface{} with range: -2147483648..2147483647.
    OspfLsdbChecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    OspfLsdbAdvertisement interface{}
}

func (ospfLsdbEntry *OSPFMIB_OspfLsdbTable_OspfLsdbEntry) GetEntityData() *types.CommonEntityData {
    ospfLsdbEntry.EntityData.YFilter = ospfLsdbEntry.YFilter
    ospfLsdbEntry.EntityData.YangName = "ospfLsdbEntry"
    ospfLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfLsdbEntry.EntityData.ParentYangName = "ospfLsdbTable"
    ospfLsdbEntry.EntityData.SegmentPath = "ospfLsdbEntry" + types.AddKeyToken(ospfLsdbEntry.OspfLsdbAreaId, "ospfLsdbAreaId") + types.AddKeyToken(ospfLsdbEntry.OspfLsdbType, "ospfLsdbType") + types.AddKeyToken(ospfLsdbEntry.OspfLsdbLsid, "ospfLsdbLsid") + types.AddKeyToken(ospfLsdbEntry.OspfLsdbRouterId, "ospfLsdbRouterId")
    ospfLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfLsdbEntry.EntityData.Children = types.NewOrderedMap()
    ospfLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbAreaId", types.YLeaf{"OspfLsdbAreaId", ospfLsdbEntry.OspfLsdbAreaId})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbType", types.YLeaf{"OspfLsdbType", ospfLsdbEntry.OspfLsdbType})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbLsid", types.YLeaf{"OspfLsdbLsid", ospfLsdbEntry.OspfLsdbLsid})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbRouterId", types.YLeaf{"OspfLsdbRouterId", ospfLsdbEntry.OspfLsdbRouterId})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbSequence", types.YLeaf{"OspfLsdbSequence", ospfLsdbEntry.OspfLsdbSequence})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbAge", types.YLeaf{"OspfLsdbAge", ospfLsdbEntry.OspfLsdbAge})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbChecksum", types.YLeaf{"OspfLsdbChecksum", ospfLsdbEntry.OspfLsdbChecksum})
    ospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbAdvertisement", types.YLeaf{"OspfLsdbAdvertisement", ospfLsdbEntry.OspfLsdbAdvertisement})

    ospfLsdbEntry.EntityData.YListKeys = []string {"OspfLsdbAreaId", "OspfLsdbType", "OspfLsdbLsid", "OspfLsdbRouterId"}

    return &(ospfLsdbEntry.EntityData)
}

// OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType represents in the ospfAsLsdbTable rather than here.
type OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType string

const (
    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_routerLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "routerLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_networkLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "networkLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_summaryLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "summaryLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_asSummaryLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "asSummaryLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_asExternalLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "asExternalLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_multicastLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "multicastLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_nssaExternalLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "nssaExternalLink"

    OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType_areaOpaqueLink OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbType = "areaOpaqueLink"
)

// OSPFMIB_OspfAreaRangeTable
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
type OSPFMIB_OspfAreaRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single area address range.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry.
    OspfAreaRangeEntry []*OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry
}

func (ospfAreaRangeTable *OSPFMIB_OspfAreaRangeTable) GetEntityData() *types.CommonEntityData {
    ospfAreaRangeTable.EntityData.YFilter = ospfAreaRangeTable.YFilter
    ospfAreaRangeTable.EntityData.YangName = "ospfAreaRangeTable"
    ospfAreaRangeTable.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaRangeTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfAreaRangeTable.EntityData.SegmentPath = "ospfAreaRangeTable"
    ospfAreaRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaRangeTable.EntityData.Children = types.NewOrderedMap()
    ospfAreaRangeTable.EntityData.Children.Append("ospfAreaRangeEntry", types.YChild{"OspfAreaRangeEntry", nil})
    for i := range ospfAreaRangeTable.OspfAreaRangeEntry {
        ospfAreaRangeTable.EntityData.Children.Append(types.GetSegmentPath(ospfAreaRangeTable.OspfAreaRangeEntry[i]), types.YChild{"OspfAreaRangeEntry", ospfAreaRangeTable.OspfAreaRangeEntry[i]})
    }
    ospfAreaRangeTable.EntityData.Leafs = types.NewOrderedMap()

    ospfAreaRangeTable.EntityData.YListKeys = []string {}

    return &(ospfAreaRangeTable.EntityData)
}

// OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry
// A single area address range.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The area that the address range is to be found
    // within. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaRangeAreaId interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaRangeNet interface{}

    // The subnet mask that pertains to the net or subnet. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaRangeMask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfAreaRangeStatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated summary (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is OspfAreaRangeEffect.
    OspfAreaRangeEffect interface{}
}

func (ospfAreaRangeEntry *OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry) GetEntityData() *types.CommonEntityData {
    ospfAreaRangeEntry.EntityData.YFilter = ospfAreaRangeEntry.YFilter
    ospfAreaRangeEntry.EntityData.YangName = "ospfAreaRangeEntry"
    ospfAreaRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaRangeEntry.EntityData.ParentYangName = "ospfAreaRangeTable"
    ospfAreaRangeEntry.EntityData.SegmentPath = "ospfAreaRangeEntry" + types.AddKeyToken(ospfAreaRangeEntry.OspfAreaRangeAreaId, "ospfAreaRangeAreaId") + types.AddKeyToken(ospfAreaRangeEntry.OspfAreaRangeNet, "ospfAreaRangeNet")
    ospfAreaRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaRangeEntry.EntityData.Children = types.NewOrderedMap()
    ospfAreaRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfAreaRangeEntry.EntityData.Leafs.Append("ospfAreaRangeAreaId", types.YLeaf{"OspfAreaRangeAreaId", ospfAreaRangeEntry.OspfAreaRangeAreaId})
    ospfAreaRangeEntry.EntityData.Leafs.Append("ospfAreaRangeNet", types.YLeaf{"OspfAreaRangeNet", ospfAreaRangeEntry.OspfAreaRangeNet})
    ospfAreaRangeEntry.EntityData.Leafs.Append("ospfAreaRangeMask", types.YLeaf{"OspfAreaRangeMask", ospfAreaRangeEntry.OspfAreaRangeMask})
    ospfAreaRangeEntry.EntityData.Leafs.Append("ospfAreaRangeStatus", types.YLeaf{"OspfAreaRangeStatus", ospfAreaRangeEntry.OspfAreaRangeStatus})
    ospfAreaRangeEntry.EntityData.Leafs.Append("ospfAreaRangeEffect", types.YLeaf{"OspfAreaRangeEffect", ospfAreaRangeEntry.OspfAreaRangeEffect})

    ospfAreaRangeEntry.EntityData.YListKeys = []string {"OspfAreaRangeAreaId", "OspfAreaRangeNet"}

    return &(ospfAreaRangeEntry.EntityData)
}

// OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect represents being advertised at all outside the area.
type OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect string

const (
    OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect_advertiseMatching OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect = "advertiseMatching"

    OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect_doNotAdvertiseMatching OSPFMIB_OspfAreaRangeTable_OspfAreaRangeEntry_OspfAreaRangeEffect = "doNotAdvertiseMatching"
)

// OSPFMIB_OspfHostTable
// The Host/Metric Table indicates what hosts are directly
// 
// attached to the router, what metrics and types
// of service should be advertised for them,
// and what areas they are found within.
type OSPFMIB_OspfHostTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A metric to be advertised, for a given type of service, when a given host
    // is reachable.  Information in this table is persistent and when this object
    // is written the entity SHOULD save the change to non-volatile storage. The
    // type is slice of OSPFMIB_OspfHostTable_OspfHostEntry.
    OspfHostEntry []*OSPFMIB_OspfHostTable_OspfHostEntry
}

func (ospfHostTable *OSPFMIB_OspfHostTable) GetEntityData() *types.CommonEntityData {
    ospfHostTable.EntityData.YFilter = ospfHostTable.YFilter
    ospfHostTable.EntityData.YangName = "ospfHostTable"
    ospfHostTable.EntityData.BundleName = "cisco_ios_xe"
    ospfHostTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfHostTable.EntityData.SegmentPath = "ospfHostTable"
    ospfHostTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfHostTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfHostTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfHostTable.EntityData.Children = types.NewOrderedMap()
    ospfHostTable.EntityData.Children.Append("ospfHostEntry", types.YChild{"OspfHostEntry", nil})
    for i := range ospfHostTable.OspfHostEntry {
        ospfHostTable.EntityData.Children.Append(types.GetSegmentPath(ospfHostTable.OspfHostEntry[i]), types.YChild{"OspfHostEntry", ospfHostTable.OspfHostEntry[i]})
    }
    ospfHostTable.EntityData.Leafs = types.NewOrderedMap()

    ospfHostTable.EntityData.YListKeys = []string {}

    return &(ospfHostTable.EntityData)
}

// OSPFMIB_OspfHostTable_OspfHostEntry
// A metric to be advertised, for a given type of
// service, when a given host is reachable.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfHostTable_OspfHostEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the host. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfHostIpAddress interface{}

    // This attribute is a key. The Type of Service of the route being configured.
    // The type is interface{} with range: 0..30.
    OspfHostTOS interface{}

    // The metric to be advertised. The type is interface{} with range: 0..65535.
    OspfHostMetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfHostStatus interface{}

    // The OSPF area to which the host belongs. Deprecated by ospfHostCfgAreaID.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfHostAreaID interface{}

    // To configure the OSPF area to which the host belongs. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfHostCfgAreaID interface{}
}

func (ospfHostEntry *OSPFMIB_OspfHostTable_OspfHostEntry) GetEntityData() *types.CommonEntityData {
    ospfHostEntry.EntityData.YFilter = ospfHostEntry.YFilter
    ospfHostEntry.EntityData.YangName = "ospfHostEntry"
    ospfHostEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfHostEntry.EntityData.ParentYangName = "ospfHostTable"
    ospfHostEntry.EntityData.SegmentPath = "ospfHostEntry" + types.AddKeyToken(ospfHostEntry.OspfHostIpAddress, "ospfHostIpAddress") + types.AddKeyToken(ospfHostEntry.OspfHostTOS, "ospfHostTOS")
    ospfHostEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfHostEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfHostEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfHostEntry.EntityData.Children = types.NewOrderedMap()
    ospfHostEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfHostEntry.EntityData.Leafs.Append("ospfHostIpAddress", types.YLeaf{"OspfHostIpAddress", ospfHostEntry.OspfHostIpAddress})
    ospfHostEntry.EntityData.Leafs.Append("ospfHostTOS", types.YLeaf{"OspfHostTOS", ospfHostEntry.OspfHostTOS})
    ospfHostEntry.EntityData.Leafs.Append("ospfHostMetric", types.YLeaf{"OspfHostMetric", ospfHostEntry.OspfHostMetric})
    ospfHostEntry.EntityData.Leafs.Append("ospfHostStatus", types.YLeaf{"OspfHostStatus", ospfHostEntry.OspfHostStatus})
    ospfHostEntry.EntityData.Leafs.Append("ospfHostAreaID", types.YLeaf{"OspfHostAreaID", ospfHostEntry.OspfHostAreaID})
    ospfHostEntry.EntityData.Leafs.Append("ospfHostCfgAreaID", types.YLeaf{"OspfHostCfgAreaID", ospfHostEntry.OspfHostCfgAreaID})

    ospfHostEntry.EntityData.YListKeys = []string {"OspfHostIpAddress", "OspfHostTOS"}

    return &(ospfHostEntry.EntityData)
}

// OSPFMIB_OspfIfTable
// The OSPF Interface Table describes the interfaces
// from the viewpoint of OSPF.
// It augments the ipAddrTable with OSPF specific information.
type OSPFMIB_OspfIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The OSPF interface entry describes one interface from the viewpoint of
    // OSPF.  Information in this table is persistent and when this object is
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is slice of OSPFMIB_OspfIfTable_OspfIfEntry.
    OspfIfEntry []*OSPFMIB_OspfIfTable_OspfIfEntry
}

func (ospfIfTable *OSPFMIB_OspfIfTable) GetEntityData() *types.CommonEntityData {
    ospfIfTable.EntityData.YFilter = ospfIfTable.YFilter
    ospfIfTable.EntityData.YangName = "ospfIfTable"
    ospfIfTable.EntityData.BundleName = "cisco_ios_xe"
    ospfIfTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfIfTable.EntityData.SegmentPath = "ospfIfTable"
    ospfIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfIfTable.EntityData.Children = types.NewOrderedMap()
    ospfIfTable.EntityData.Children.Append("ospfIfEntry", types.YChild{"OspfIfEntry", nil})
    for i := range ospfIfTable.OspfIfEntry {
        ospfIfTable.EntityData.Children.Append(types.GetSegmentPath(ospfIfTable.OspfIfEntry[i]), types.YChild{"OspfIfEntry", ospfIfTable.OspfIfEntry[i]})
    }
    ospfIfTable.EntityData.Leafs = types.NewOrderedMap()

    ospfIfTable.EntityData.YListKeys = []string {}

    return &(ospfIfTable.EntityData)
}

// OSPFMIB_OspfIfTable_OspfIfEntry
// The OSPF interface entry describes one interface
// from the viewpoint of OSPF.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfIfTable_OspfIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfIpAddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the corresponding value of ifIndex for
    // interfaces having no IP address. The type is interface{} with range:
    // 0..2147483647.
    OspfAddressLessIf interface{}

    // A 32-bit integer uniquely identifying the area to which the interface
    // connects.  Area ID 0.0.0.0 is used for the OSPF backbone. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfAreaId interface{}

    // The OSPF interface type. By way of a default, this field may be intuited
    // from the corresponding value of ifType. Broadcast LANs, such as Ethernet
    // and IEEE 802.5, take the value 'broadcast', X.25 and similar technologies
    // take the value 'nbma', and links that are definitively point to point take
    // the value 'pointToPoint'. The type is OspfIfType.
    OspfIfType interface{}

    // The OSPF interface's administrative status. The value formed on the
    // interface, and the interface will be advertised as an internal route to
    // some area. The value 'disabled' denotes that the interface is external to
    // OSPF. The type is Status.
    OspfIfAdminStat interface{}

    // The priority of this interface.  Used in multi-access networks, this field
    // is used in the designated router election algorithm.  The value 0 signifies
    // that the router is not eligible to become the designated router on this
    // particular network.  In the event of a tie in this value, routers will use
    // their Router ID as a tie breaker. The type is interface{} with range:
    // 0..255.
    OspfIfRtrPriority interface{}

    // The estimated number of seconds it takes to transmit a link state update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    OspfIfTransitDelay interface{}

    // The number of seconds between link state advertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting  database description and Link State request packets. Note
    // that minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    OspfIfRetransInterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range: 1..65535.
    // Units are seconds.
    OspfIfHelloInterval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down. This should be some multiple
    // of the Hello interval.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    OspfIfRtrDeadInterval interface{}

    // The larger time interval, in seconds, between the Hello packets sent to an
    // inactive non-broadcast multi-access neighbor. The type is interface{} with
    // range: 0..2147483647. Units are seconds.
    OspfIfPollInterval interface{}

    // The OSPF Interface State. The type is OspfIfState.
    OspfIfState interface{}

    // The IP address of the designated router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfDesignatedRouter interface{}

    // The IP address of the backup designated router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfBackupDesignatedRouter interface{}

    // The number of times this OSPF interface has changed its state or an error
    // has occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    OspfIfEvents interface{}

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
    OspfIfAuthKey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfIfStatus interface{}

    // The way multicasts should be forwarded on this interface: not forwarded,
    // forwarded as data link multicasts, or forwarded as data link unicasts. 
    // Data link multicasting is not meaningful on point-to-point and NBMA
    // interfaces, and setting ospfMulticastForwarding to 0 effectively disables
    // all multicast forwarding. The type is OspfIfMulticastForwarding.
    OspfIfMulticastForwarding interface{}

    // Indicates whether Demand OSPF procedures (hello suppression to FULL
    // neighbors and setting the DoNotAge flag on propagated LSAs) should be
    // performed on this interface. The type is bool.
    OspfIfDemand interface{}

    // The authentication type specified for an interface.  Note that this object
    // can be used to engage in significant attacks against an OSPF router. The
    // type is OspfAuthenticationType.
    OspfIfAuthType interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    OspfIfLsaCount interface{}

    // The 32-bit unsigned sum of the Link State Advertisements' LS checksums
    // contained in this interface's link-local link state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database and to compare the interface link state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    OspfIfLsaCksumSum interface{}

    // The Router ID of the designated router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfDesignatedRouterId interface{}

    // The Router ID of the backup designated router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfBackupDesignatedRouterId interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    CospfIfLsaCount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this interface's link-local link  state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database, and to compare the interface link-state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    CospfIfLsaCksumSum interface{}
}

func (ospfIfEntry *OSPFMIB_OspfIfTable_OspfIfEntry) GetEntityData() *types.CommonEntityData {
    ospfIfEntry.EntityData.YFilter = ospfIfEntry.YFilter
    ospfIfEntry.EntityData.YangName = "ospfIfEntry"
    ospfIfEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfIfEntry.EntityData.ParentYangName = "ospfIfTable"
    ospfIfEntry.EntityData.SegmentPath = "ospfIfEntry" + types.AddKeyToken(ospfIfEntry.OspfIfIpAddress, "ospfIfIpAddress") + types.AddKeyToken(ospfIfEntry.OspfAddressLessIf, "ospfAddressLessIf")
    ospfIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfIfEntry.EntityData.Children = types.NewOrderedMap()
    ospfIfEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfIfEntry.EntityData.Leafs.Append("ospfIfIpAddress", types.YLeaf{"OspfIfIpAddress", ospfIfEntry.OspfIfIpAddress})
    ospfIfEntry.EntityData.Leafs.Append("ospfAddressLessIf", types.YLeaf{"OspfAddressLessIf", ospfIfEntry.OspfAddressLessIf})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfAreaId", types.YLeaf{"OspfIfAreaId", ospfIfEntry.OspfIfAreaId})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfType", types.YLeaf{"OspfIfType", ospfIfEntry.OspfIfType})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfAdminStat", types.YLeaf{"OspfIfAdminStat", ospfIfEntry.OspfIfAdminStat})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfRtrPriority", types.YLeaf{"OspfIfRtrPriority", ospfIfEntry.OspfIfRtrPriority})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfTransitDelay", types.YLeaf{"OspfIfTransitDelay", ospfIfEntry.OspfIfTransitDelay})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfRetransInterval", types.YLeaf{"OspfIfRetransInterval", ospfIfEntry.OspfIfRetransInterval})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfHelloInterval", types.YLeaf{"OspfIfHelloInterval", ospfIfEntry.OspfIfHelloInterval})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfRtrDeadInterval", types.YLeaf{"OspfIfRtrDeadInterval", ospfIfEntry.OspfIfRtrDeadInterval})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfPollInterval", types.YLeaf{"OspfIfPollInterval", ospfIfEntry.OspfIfPollInterval})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfState", types.YLeaf{"OspfIfState", ospfIfEntry.OspfIfState})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfDesignatedRouter", types.YLeaf{"OspfIfDesignatedRouter", ospfIfEntry.OspfIfDesignatedRouter})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfBackupDesignatedRouter", types.YLeaf{"OspfIfBackupDesignatedRouter", ospfIfEntry.OspfIfBackupDesignatedRouter})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfEvents", types.YLeaf{"OspfIfEvents", ospfIfEntry.OspfIfEvents})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfAuthKey", types.YLeaf{"OspfIfAuthKey", ospfIfEntry.OspfIfAuthKey})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfStatus", types.YLeaf{"OspfIfStatus", ospfIfEntry.OspfIfStatus})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfMulticastForwarding", types.YLeaf{"OspfIfMulticastForwarding", ospfIfEntry.OspfIfMulticastForwarding})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfDemand", types.YLeaf{"OspfIfDemand", ospfIfEntry.OspfIfDemand})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfAuthType", types.YLeaf{"OspfIfAuthType", ospfIfEntry.OspfIfAuthType})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfLsaCount", types.YLeaf{"OspfIfLsaCount", ospfIfEntry.OspfIfLsaCount})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfLsaCksumSum", types.YLeaf{"OspfIfLsaCksumSum", ospfIfEntry.OspfIfLsaCksumSum})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfDesignatedRouterId", types.YLeaf{"OspfIfDesignatedRouterId", ospfIfEntry.OspfIfDesignatedRouterId})
    ospfIfEntry.EntityData.Leafs.Append("ospfIfBackupDesignatedRouterId", types.YLeaf{"OspfIfBackupDesignatedRouterId", ospfIfEntry.OspfIfBackupDesignatedRouterId})
    ospfIfEntry.EntityData.Leafs.Append("cospfIfLsaCount", types.YLeaf{"CospfIfLsaCount", ospfIfEntry.CospfIfLsaCount})
    ospfIfEntry.EntityData.Leafs.Append("cospfIfLsaCksumSum", types.YLeaf{"CospfIfLsaCksumSum", ospfIfEntry.CospfIfLsaCksumSum})

    ospfIfEntry.EntityData.YListKeys = []string {"OspfIfIpAddress", "OspfAddressLessIf"}

    return &(ospfIfEntry.EntityData)
}

// OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding represents disables all multicast forwarding.
type OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding string

const (
    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding_blocked OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding = "blocked"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding_multicast OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding = "multicast"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding_unicast OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfMulticastForwarding = "unicast"
)

// OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState represents The OSPF Interface State.
type OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState string

const (
    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_down OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "down"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_loopback OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "loopback"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_waiting OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "waiting"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_pointToPoint OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "pointToPoint"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_designatedRouter OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "designatedRouter"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_backupDesignatedRouter OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "backupDesignatedRouter"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState_otherDesignatedRouter OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfState = "otherDesignatedRouter"
)

// OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType represents value 'pointToPoint'.
type OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType string

const (
    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType_broadcast OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType = "broadcast"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType_nbma OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType = "nbma"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType_pointToPoint OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType = "pointToPoint"

    OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType_pointToMultipoint OSPFMIB_OspfIfTable_OspfIfEntry_OspfIfType = "pointToMultipoint"
)

// OSPFMIB_OspfIfMetricTable
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
type OSPFMIB_OspfIfMetricTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular TOS metric for a non-virtual interface identified by the
    // interface index.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_OspfIfMetricTable_OspfIfMetricEntry.
    OspfIfMetricEntry []*OSPFMIB_OspfIfMetricTable_OspfIfMetricEntry
}

func (ospfIfMetricTable *OSPFMIB_OspfIfMetricTable) GetEntityData() *types.CommonEntityData {
    ospfIfMetricTable.EntityData.YFilter = ospfIfMetricTable.YFilter
    ospfIfMetricTable.EntityData.YangName = "ospfIfMetricTable"
    ospfIfMetricTable.EntityData.BundleName = "cisco_ios_xe"
    ospfIfMetricTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfIfMetricTable.EntityData.SegmentPath = "ospfIfMetricTable"
    ospfIfMetricTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfIfMetricTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfIfMetricTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfIfMetricTable.EntityData.Children = types.NewOrderedMap()
    ospfIfMetricTable.EntityData.Children.Append("ospfIfMetricEntry", types.YChild{"OspfIfMetricEntry", nil})
    for i := range ospfIfMetricTable.OspfIfMetricEntry {
        ospfIfMetricTable.EntityData.Children.Append(types.GetSegmentPath(ospfIfMetricTable.OspfIfMetricEntry[i]), types.YChild{"OspfIfMetricEntry", ospfIfMetricTable.OspfIfMetricEntry[i]})
    }
    ospfIfMetricTable.EntityData.Leafs = types.NewOrderedMap()

    ospfIfMetricTable.EntityData.YListKeys = []string {}

    return &(ospfIfMetricTable.EntityData)
}

// OSPFMIB_OspfIfMetricTable_OspfIfMetricEntry
// A particular TOS metric for a non-virtual interface
// identified by the interface index.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfIfMetricTable_OspfIfMetricEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface.  On row
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfIfMetricIpAddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the value of ifIndex for interfaces having
    // no IP address.  On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    OspfIfMetricAddressLessIf interface{}

    // This attribute is a key. The Type of Service metric being referenced. On
    // row creation, this can be derived from the instance. The type is
    // interface{} with range: 0..30.
    OspfIfMetricTOS interface{}

    // The metric of using this Type of Service on this interface.  The default
    // value of the TOS 0 metric is 10^8 / ifSpeed. The type is interface{} with
    // range: 0..65535.
    OspfIfMetricValue interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfIfMetricStatus interface{}
}

func (ospfIfMetricEntry *OSPFMIB_OspfIfMetricTable_OspfIfMetricEntry) GetEntityData() *types.CommonEntityData {
    ospfIfMetricEntry.EntityData.YFilter = ospfIfMetricEntry.YFilter
    ospfIfMetricEntry.EntityData.YangName = "ospfIfMetricEntry"
    ospfIfMetricEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfIfMetricEntry.EntityData.ParentYangName = "ospfIfMetricTable"
    ospfIfMetricEntry.EntityData.SegmentPath = "ospfIfMetricEntry" + types.AddKeyToken(ospfIfMetricEntry.OspfIfMetricIpAddress, "ospfIfMetricIpAddress") + types.AddKeyToken(ospfIfMetricEntry.OspfIfMetricAddressLessIf, "ospfIfMetricAddressLessIf") + types.AddKeyToken(ospfIfMetricEntry.OspfIfMetricTOS, "ospfIfMetricTOS")
    ospfIfMetricEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfIfMetricEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfIfMetricEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfIfMetricEntry.EntityData.Children = types.NewOrderedMap()
    ospfIfMetricEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfIfMetricEntry.EntityData.Leafs.Append("ospfIfMetricIpAddress", types.YLeaf{"OspfIfMetricIpAddress", ospfIfMetricEntry.OspfIfMetricIpAddress})
    ospfIfMetricEntry.EntityData.Leafs.Append("ospfIfMetricAddressLessIf", types.YLeaf{"OspfIfMetricAddressLessIf", ospfIfMetricEntry.OspfIfMetricAddressLessIf})
    ospfIfMetricEntry.EntityData.Leafs.Append("ospfIfMetricTOS", types.YLeaf{"OspfIfMetricTOS", ospfIfMetricEntry.OspfIfMetricTOS})
    ospfIfMetricEntry.EntityData.Leafs.Append("ospfIfMetricValue", types.YLeaf{"OspfIfMetricValue", ospfIfMetricEntry.OspfIfMetricValue})
    ospfIfMetricEntry.EntityData.Leafs.Append("ospfIfMetricStatus", types.YLeaf{"OspfIfMetricStatus", ospfIfMetricEntry.OspfIfMetricStatus})

    ospfIfMetricEntry.EntityData.YListKeys = []string {"OspfIfMetricIpAddress", "OspfIfMetricAddressLessIf", "OspfIfMetricTOS"}

    return &(ospfIfMetricEntry.EntityData)
}

// OSPFMIB_OspfVirtIfTable
// Information about this router's virtual interfaces
// that the OSPF Process is configured to carry on.
type OSPFMIB_OspfVirtIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single virtual interface.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry.
    OspfVirtIfEntry []*OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry
}

func (ospfVirtIfTable *OSPFMIB_OspfVirtIfTable) GetEntityData() *types.CommonEntityData {
    ospfVirtIfTable.EntityData.YFilter = ospfVirtIfTable.YFilter
    ospfVirtIfTable.EntityData.YangName = "ospfVirtIfTable"
    ospfVirtIfTable.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtIfTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfVirtIfTable.EntityData.SegmentPath = "ospfVirtIfTable"
    ospfVirtIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtIfTable.EntityData.Children = types.NewOrderedMap()
    ospfVirtIfTable.EntityData.Children.Append("ospfVirtIfEntry", types.YChild{"OspfVirtIfEntry", nil})
    for i := range ospfVirtIfTable.OspfVirtIfEntry {
        ospfVirtIfTable.EntityData.Children.Append(types.GetSegmentPath(ospfVirtIfTable.OspfVirtIfEntry[i]), types.YChild{"OspfVirtIfEntry", ospfVirtIfTable.OspfVirtIfEntry[i]})
    }
    ospfVirtIfTable.EntityData.Leafs = types.NewOrderedMap()

    ospfVirtIfTable.EntityData.YListKeys = []string {}

    return &(ospfVirtIfTable.EntityData)
}

// OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry
// Information about a single virtual interface.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtIfAreaId interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtIfNeighbor interface{}

    // The estimated number of seconds it takes to transmit a Link State update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    OspfVirtIfTransitDelay interface{}

    // The number of seconds between link state avertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting database description and Link State request packets.  This
    // value should be well over the expected round-trip time.  Note that the
    // minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    OspfVirtIfRetransInterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 1..65535. Units are seconds.
    OspfVirtIfHelloInterval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down.  This should be some multiple
    // of the Hello interval.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    OspfVirtIfRtrDeadInterval interface{}

    // OSPF virtual interface states. The type is OspfVirtIfState.
    OspfVirtIfState interface{}

    // The number of state changes or error events on this virtual link. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    OspfVirtIfEvents interface{}

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
    OspfVirtIfAuthKey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfVirtIfStatus interface{}

    // The authentication type specified for a virtual interface.  Note that this
    // object can be used to engage in significant attacks against an OSPF router.
    // The type is OspfAuthenticationType.
    OspfVirtIfAuthType interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    OspfVirtIfLsaCount interface{}

    // The 32-bit unsigned sum of the link state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface link
    // state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    OspfVirtIfLsaCksumSum interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    CospfVirtIfLsaCount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface
    // link-state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    CospfVirtIfLsaCksumSum interface{}
}

func (ospfVirtIfEntry *OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry) GetEntityData() *types.CommonEntityData {
    ospfVirtIfEntry.EntityData.YFilter = ospfVirtIfEntry.YFilter
    ospfVirtIfEntry.EntityData.YangName = "ospfVirtIfEntry"
    ospfVirtIfEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtIfEntry.EntityData.ParentYangName = "ospfVirtIfTable"
    ospfVirtIfEntry.EntityData.SegmentPath = "ospfVirtIfEntry" + types.AddKeyToken(ospfVirtIfEntry.OspfVirtIfAreaId, "ospfVirtIfAreaId") + types.AddKeyToken(ospfVirtIfEntry.OspfVirtIfNeighbor, "ospfVirtIfNeighbor")
    ospfVirtIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtIfEntry.EntityData.Children = types.NewOrderedMap()
    ospfVirtIfEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfAreaId", types.YLeaf{"OspfVirtIfAreaId", ospfVirtIfEntry.OspfVirtIfAreaId})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfNeighbor", types.YLeaf{"OspfVirtIfNeighbor", ospfVirtIfEntry.OspfVirtIfNeighbor})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfTransitDelay", types.YLeaf{"OspfVirtIfTransitDelay", ospfVirtIfEntry.OspfVirtIfTransitDelay})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfRetransInterval", types.YLeaf{"OspfVirtIfRetransInterval", ospfVirtIfEntry.OspfVirtIfRetransInterval})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfHelloInterval", types.YLeaf{"OspfVirtIfHelloInterval", ospfVirtIfEntry.OspfVirtIfHelloInterval})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfRtrDeadInterval", types.YLeaf{"OspfVirtIfRtrDeadInterval", ospfVirtIfEntry.OspfVirtIfRtrDeadInterval})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfState", types.YLeaf{"OspfVirtIfState", ospfVirtIfEntry.OspfVirtIfState})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfEvents", types.YLeaf{"OspfVirtIfEvents", ospfVirtIfEntry.OspfVirtIfEvents})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfAuthKey", types.YLeaf{"OspfVirtIfAuthKey", ospfVirtIfEntry.OspfVirtIfAuthKey})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfStatus", types.YLeaf{"OspfVirtIfStatus", ospfVirtIfEntry.OspfVirtIfStatus})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfAuthType", types.YLeaf{"OspfVirtIfAuthType", ospfVirtIfEntry.OspfVirtIfAuthType})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfLsaCount", types.YLeaf{"OspfVirtIfLsaCount", ospfVirtIfEntry.OspfVirtIfLsaCount})
    ospfVirtIfEntry.EntityData.Leafs.Append("ospfVirtIfLsaCksumSum", types.YLeaf{"OspfVirtIfLsaCksumSum", ospfVirtIfEntry.OspfVirtIfLsaCksumSum})
    ospfVirtIfEntry.EntityData.Leafs.Append("cospfVirtIfLsaCount", types.YLeaf{"CospfVirtIfLsaCount", ospfVirtIfEntry.CospfVirtIfLsaCount})
    ospfVirtIfEntry.EntityData.Leafs.Append("cospfVirtIfLsaCksumSum", types.YLeaf{"CospfVirtIfLsaCksumSum", ospfVirtIfEntry.CospfVirtIfLsaCksumSum})

    ospfVirtIfEntry.EntityData.YListKeys = []string {"OspfVirtIfAreaId", "OspfVirtIfNeighbor"}

    return &(ospfVirtIfEntry.EntityData)
}

// OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState represents OSPF virtual interface states.
type OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState string

const (
    OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState_down OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState = "down"

    OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState_pointToPoint OSPFMIB_OspfVirtIfTable_OspfVirtIfEntry_OspfVirtIfState = "pointToPoint"
)

// OSPFMIB_OspfNbrTable
// A table describing all non-virtual neighbors
// in the locality of the OSPF router.
type OSPFMIB_OspfNbrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single neighbor.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile  storage. The type is slice of
    // OSPFMIB_OspfNbrTable_OspfNbrEntry.
    OspfNbrEntry []*OSPFMIB_OspfNbrTable_OspfNbrEntry
}

func (ospfNbrTable *OSPFMIB_OspfNbrTable) GetEntityData() *types.CommonEntityData {
    ospfNbrTable.EntityData.YFilter = ospfNbrTable.YFilter
    ospfNbrTable.EntityData.YangName = "ospfNbrTable"
    ospfNbrTable.EntityData.BundleName = "cisco_ios_xe"
    ospfNbrTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfNbrTable.EntityData.SegmentPath = "ospfNbrTable"
    ospfNbrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfNbrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfNbrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfNbrTable.EntityData.Children = types.NewOrderedMap()
    ospfNbrTable.EntityData.Children.Append("ospfNbrEntry", types.YChild{"OspfNbrEntry", nil})
    for i := range ospfNbrTable.OspfNbrEntry {
        ospfNbrTable.EntityData.Children.Append(types.GetSegmentPath(ospfNbrTable.OspfNbrEntry[i]), types.YChild{"OspfNbrEntry", ospfNbrTable.OspfNbrEntry[i]})
    }
    ospfNbrTable.EntityData.Leafs = types.NewOrderedMap()

    ospfNbrTable.EntityData.YListKeys = []string {}

    return &(ospfNbrTable.EntityData)
}

// OSPFMIB_OspfNbrTable_OspfNbrEntry
// The information regarding a single neighbor.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// 
// storage.
type OSPFMIB_OspfNbrTable_OspfNbrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address this neighbor is using in its IP
    // source address.  Note that, on addressless links, this will not be 0.0.0.0
    // but the  address of another of the neighbor's interfaces. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfNbrIpAddr interface{}

    // This attribute is a key. On an interface having an IP address, zero. On
    // addressless interfaces, the corresponding value of ifIndex in the Internet
    // Standard MIB. On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    OspfNbrAddressLessIndex interface{}

    // A 32-bit integer (represented as a type IpAddress) uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfNbrRtrId interface{}

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
    OspfNbrOptions interface{}

    // The priority of this neighbor in the designated router election algorithm. 
    // The value 0 signifies that the neighbor is not eligible to become the
    // designated router on this particular network. The type is interface{} with
    // range: 0..255.
    OspfNbrPriority interface{}

    // The state of the relationship with this neighbor. The type is OspfNbrState.
    OspfNbrState interface{}

    // The number of times this neighbor relationship has changed state or an
    // error has occurred.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    OspfNbrEvents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    OspfNbrLsRetransQLen interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfNbmaNbrStatus interface{}

    // This variable displays the status of the entry; 'dynamic' and 'permanent'
    // refer to how the neighbor became known. The type is OspfNbmaNbrPermanence.
    OspfNbmaNbrPermanence interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    OspfNbrHelloSuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is OspfNbrRestartHelperStatus.
    OspfNbrRestartHelperStatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    OspfNbrRestartHelperAge interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is OspfNbrRestartHelperExitReason.
    OspfNbrRestartHelperExitReason interface{}
}

func (ospfNbrEntry *OSPFMIB_OspfNbrTable_OspfNbrEntry) GetEntityData() *types.CommonEntityData {
    ospfNbrEntry.EntityData.YFilter = ospfNbrEntry.YFilter
    ospfNbrEntry.EntityData.YangName = "ospfNbrEntry"
    ospfNbrEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfNbrEntry.EntityData.ParentYangName = "ospfNbrTable"
    ospfNbrEntry.EntityData.SegmentPath = "ospfNbrEntry" + types.AddKeyToken(ospfNbrEntry.OspfNbrIpAddr, "ospfNbrIpAddr") + types.AddKeyToken(ospfNbrEntry.OspfNbrAddressLessIndex, "ospfNbrAddressLessIndex")
    ospfNbrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfNbrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfNbrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfNbrEntry.EntityData.Children = types.NewOrderedMap()
    ospfNbrEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrIpAddr", types.YLeaf{"OspfNbrIpAddr", ospfNbrEntry.OspfNbrIpAddr})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrAddressLessIndex", types.YLeaf{"OspfNbrAddressLessIndex", ospfNbrEntry.OspfNbrAddressLessIndex})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrRtrId", types.YLeaf{"OspfNbrRtrId", ospfNbrEntry.OspfNbrRtrId})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrOptions", types.YLeaf{"OspfNbrOptions", ospfNbrEntry.OspfNbrOptions})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrPriority", types.YLeaf{"OspfNbrPriority", ospfNbrEntry.OspfNbrPriority})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrState", types.YLeaf{"OspfNbrState", ospfNbrEntry.OspfNbrState})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrEvents", types.YLeaf{"OspfNbrEvents", ospfNbrEntry.OspfNbrEvents})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrLsRetransQLen", types.YLeaf{"OspfNbrLsRetransQLen", ospfNbrEntry.OspfNbrLsRetransQLen})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbmaNbrStatus", types.YLeaf{"OspfNbmaNbrStatus", ospfNbrEntry.OspfNbmaNbrStatus})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbmaNbrPermanence", types.YLeaf{"OspfNbmaNbrPermanence", ospfNbrEntry.OspfNbmaNbrPermanence})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrHelloSuppressed", types.YLeaf{"OspfNbrHelloSuppressed", ospfNbrEntry.OspfNbrHelloSuppressed})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrRestartHelperStatus", types.YLeaf{"OspfNbrRestartHelperStatus", ospfNbrEntry.OspfNbrRestartHelperStatus})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrRestartHelperAge", types.YLeaf{"OspfNbrRestartHelperAge", ospfNbrEntry.OspfNbrRestartHelperAge})
    ospfNbrEntry.EntityData.Leafs.Append("ospfNbrRestartHelperExitReason", types.YLeaf{"OspfNbrRestartHelperExitReason", ospfNbrEntry.OspfNbrRestartHelperExitReason})

    ospfNbrEntry.EntityData.YListKeys = []string {"OspfNbrIpAddr", "OspfNbrAddressLessIndex"}

    return &(ospfNbrEntry.EntityData)
}

// OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence represents became known.
type OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence string

const (
    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence_dynamic OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence = "dynamic"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence_permanent OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbmaNbrPermanence = "permanent"
)

// OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason represents as a graceful restart helper for the neighbor.
type OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason string

const (
    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason_none OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason = "none"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason_inProgress OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason = "inProgress"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason_completed OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason = "completed"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason_timedOut OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason = "timedOut"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason_topologyChanged OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperExitReason = "topologyChanged"
)

// OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus string

const (
    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus_notHelping OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus = "notHelping"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus_helping OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrRestartHelperStatus = "helping"
)

// OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState represents The state of the relationship with this neighbor.
type OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState string

const (
    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_down OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "down"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_attempt OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "attempt"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_init OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "init"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_twoWay OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "twoWay"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_exchangeStart OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "exchangeStart"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_exchange OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "exchange"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_loading OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "loading"

    OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState_full OSPFMIB_OspfNbrTable_OspfNbrEntry_OspfNbrState = "full"
)

// OSPFMIB_OspfVirtNbrTable
// This table describes all virtual neighbors.
// Since virtual links are configured
// in the Virtual Interface Table, this table is read-only.
type OSPFMIB_OspfVirtNbrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual neighbor information. The type is slice of
    // OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry.
    OspfVirtNbrEntry []*OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry
}

func (ospfVirtNbrTable *OSPFMIB_OspfVirtNbrTable) GetEntityData() *types.CommonEntityData {
    ospfVirtNbrTable.EntityData.YFilter = ospfVirtNbrTable.YFilter
    ospfVirtNbrTable.EntityData.YangName = "ospfVirtNbrTable"
    ospfVirtNbrTable.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtNbrTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfVirtNbrTable.EntityData.SegmentPath = "ospfVirtNbrTable"
    ospfVirtNbrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtNbrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtNbrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtNbrTable.EntityData.Children = types.NewOrderedMap()
    ospfVirtNbrTable.EntityData.Children.Append("ospfVirtNbrEntry", types.YChild{"OspfVirtNbrEntry", nil})
    for i := range ospfVirtNbrTable.OspfVirtNbrEntry {
        ospfVirtNbrTable.EntityData.Children.Append(types.GetSegmentPath(ospfVirtNbrTable.OspfVirtNbrEntry[i]), types.YChild{"OspfVirtNbrEntry", ospfVirtNbrTable.OspfVirtNbrEntry[i]})
    }
    ospfVirtNbrTable.EntityData.Leafs = types.NewOrderedMap()

    ospfVirtNbrTable.EntityData.YListKeys = []string {}

    return &(ospfVirtNbrTable.EntityData)
}

// OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry
// Virtual neighbor information.
type OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Transit Area Identifier. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtNbrArea interface{}

    // This attribute is a key. A 32-bit integer uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtNbrRtrId interface{}

    // The IP address this virtual neighbor is using. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtNbrIpAddr interface{}

    // A bit mask corresponding to the neighbor's options field.  Bit 1, if set,
    // indicates that the system will operate on Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0
    // metric.  Bit 2, if set, indicates that the system is network multicast
    // capable, i.e., that it implements OSPF multicast routing. The type is
    // interface{} with range: -2147483648..2147483647.
    OspfVirtNbrOptions interface{}

    // The state of the virtual neighbor relationship. The type is
    // OspfVirtNbrState.
    OspfVirtNbrState interface{}

    // The number of times this virtual link has changed its state or an error has
    // occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    OspfVirtNbrEvents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    OspfVirtNbrLsRetransQLen interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    OspfVirtNbrHelloSuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is OspfVirtNbrRestartHelperStatus.
    OspfVirtNbrRestartHelperStatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    OspfVirtNbrRestartHelperAge interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is OspfVirtNbrRestartHelperExitReason.
    OspfVirtNbrRestartHelperExitReason interface{}
}

func (ospfVirtNbrEntry *OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry) GetEntityData() *types.CommonEntityData {
    ospfVirtNbrEntry.EntityData.YFilter = ospfVirtNbrEntry.YFilter
    ospfVirtNbrEntry.EntityData.YangName = "ospfVirtNbrEntry"
    ospfVirtNbrEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtNbrEntry.EntityData.ParentYangName = "ospfVirtNbrTable"
    ospfVirtNbrEntry.EntityData.SegmentPath = "ospfVirtNbrEntry" + types.AddKeyToken(ospfVirtNbrEntry.OspfVirtNbrArea, "ospfVirtNbrArea") + types.AddKeyToken(ospfVirtNbrEntry.OspfVirtNbrRtrId, "ospfVirtNbrRtrId")
    ospfVirtNbrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtNbrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtNbrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtNbrEntry.EntityData.Children = types.NewOrderedMap()
    ospfVirtNbrEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrArea", types.YLeaf{"OspfVirtNbrArea", ospfVirtNbrEntry.OspfVirtNbrArea})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrRtrId", types.YLeaf{"OspfVirtNbrRtrId", ospfVirtNbrEntry.OspfVirtNbrRtrId})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrIpAddr", types.YLeaf{"OspfVirtNbrIpAddr", ospfVirtNbrEntry.OspfVirtNbrIpAddr})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrOptions", types.YLeaf{"OspfVirtNbrOptions", ospfVirtNbrEntry.OspfVirtNbrOptions})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrState", types.YLeaf{"OspfVirtNbrState", ospfVirtNbrEntry.OspfVirtNbrState})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrEvents", types.YLeaf{"OspfVirtNbrEvents", ospfVirtNbrEntry.OspfVirtNbrEvents})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrLsRetransQLen", types.YLeaf{"OspfVirtNbrLsRetransQLen", ospfVirtNbrEntry.OspfVirtNbrLsRetransQLen})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrHelloSuppressed", types.YLeaf{"OspfVirtNbrHelloSuppressed", ospfVirtNbrEntry.OspfVirtNbrHelloSuppressed})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrRestartHelperStatus", types.YLeaf{"OspfVirtNbrRestartHelperStatus", ospfVirtNbrEntry.OspfVirtNbrRestartHelperStatus})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrRestartHelperAge", types.YLeaf{"OspfVirtNbrRestartHelperAge", ospfVirtNbrEntry.OspfVirtNbrRestartHelperAge})
    ospfVirtNbrEntry.EntityData.Leafs.Append("ospfVirtNbrRestartHelperExitReason", types.YLeaf{"OspfVirtNbrRestartHelperExitReason", ospfVirtNbrEntry.OspfVirtNbrRestartHelperExitReason})

    ospfVirtNbrEntry.EntityData.YListKeys = []string {"OspfVirtNbrArea", "OspfVirtNbrRtrId"}

    return &(ospfVirtNbrEntry.EntityData)
}

// OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason represents as a graceful restart helper for the neighbor.
type OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason string

const (
    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason_none OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason = "none"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason_inProgress OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason = "inProgress"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason_completed OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason = "completed"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason_timedOut OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason = "timedOut"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason_topologyChanged OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperExitReason = "topologyChanged"
)

// OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus string

const (
    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus_notHelping OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus = "notHelping"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus_helping OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrRestartHelperStatus = "helping"
)

// OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState represents The state of the virtual neighbor relationship.
type OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState string

const (
    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_down OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "down"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_attempt OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "attempt"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_init OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "init"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_twoWay OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "twoWay"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_exchangeStart OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "exchangeStart"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_exchange OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "exchange"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_loading OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "loading"

    OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState_full OSPFMIB_OspfVirtNbrTable_OspfVirtNbrEntry_OspfVirtNbrState = "full"
)

// OSPFMIB_OspfExtLsdbTable
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
type OSPFMIB_OspfExtLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry.
    OspfExtLsdbEntry []*OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry
}

func (ospfExtLsdbTable *OSPFMIB_OspfExtLsdbTable) GetEntityData() *types.CommonEntityData {
    ospfExtLsdbTable.EntityData.YFilter = ospfExtLsdbTable.YFilter
    ospfExtLsdbTable.EntityData.YangName = "ospfExtLsdbTable"
    ospfExtLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    ospfExtLsdbTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfExtLsdbTable.EntityData.SegmentPath = "ospfExtLsdbTable"
    ospfExtLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfExtLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfExtLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfExtLsdbTable.EntityData.Children = types.NewOrderedMap()
    ospfExtLsdbTable.EntityData.Children.Append("ospfExtLsdbEntry", types.YChild{"OspfExtLsdbEntry", nil})
    for i := range ospfExtLsdbTable.OspfExtLsdbEntry {
        ospfExtLsdbTable.EntityData.Children.Append(types.GetSegmentPath(ospfExtLsdbTable.OspfExtLsdbEntry[i]), types.YChild{"OspfExtLsdbEntry", ospfExtLsdbTable.OspfExtLsdbEntry[i]})
    }
    ospfExtLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    ospfExtLsdbTable.EntityData.YListKeys = []string {}

    return &(ospfExtLsdbTable.EntityData)
}

// OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry
// A single link state advertisement.
type OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // OspfExtLsdbType.
    OspfExtLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfExtLsdbLsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfExtLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    OspfExtLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    OspfExtLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    OspfExtLsdbChecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 36.
    OspfExtLsdbAdvertisement interface{}
}

func (ospfExtLsdbEntry *OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry) GetEntityData() *types.CommonEntityData {
    ospfExtLsdbEntry.EntityData.YFilter = ospfExtLsdbEntry.YFilter
    ospfExtLsdbEntry.EntityData.YangName = "ospfExtLsdbEntry"
    ospfExtLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfExtLsdbEntry.EntityData.ParentYangName = "ospfExtLsdbTable"
    ospfExtLsdbEntry.EntityData.SegmentPath = "ospfExtLsdbEntry" + types.AddKeyToken(ospfExtLsdbEntry.OspfExtLsdbType, "ospfExtLsdbType") + types.AddKeyToken(ospfExtLsdbEntry.OspfExtLsdbLsid, "ospfExtLsdbLsid") + types.AddKeyToken(ospfExtLsdbEntry.OspfExtLsdbRouterId, "ospfExtLsdbRouterId")
    ospfExtLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfExtLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfExtLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfExtLsdbEntry.EntityData.Children = types.NewOrderedMap()
    ospfExtLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbType", types.YLeaf{"OspfExtLsdbType", ospfExtLsdbEntry.OspfExtLsdbType})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbLsid", types.YLeaf{"OspfExtLsdbLsid", ospfExtLsdbEntry.OspfExtLsdbLsid})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbRouterId", types.YLeaf{"OspfExtLsdbRouterId", ospfExtLsdbEntry.OspfExtLsdbRouterId})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbSequence", types.YLeaf{"OspfExtLsdbSequence", ospfExtLsdbEntry.OspfExtLsdbSequence})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbAge", types.YLeaf{"OspfExtLsdbAge", ospfExtLsdbEntry.OspfExtLsdbAge})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbChecksum", types.YLeaf{"OspfExtLsdbChecksum", ospfExtLsdbEntry.OspfExtLsdbChecksum})
    ospfExtLsdbEntry.EntityData.Leafs.Append("ospfExtLsdbAdvertisement", types.YLeaf{"OspfExtLsdbAdvertisement", ospfExtLsdbEntry.OspfExtLsdbAdvertisement})

    ospfExtLsdbEntry.EntityData.YListKeys = []string {"OspfExtLsdbType", "OspfExtLsdbLsid", "OspfExtLsdbRouterId"}

    return &(ospfExtLsdbEntry.EntityData)
}

// OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry_OspfExtLsdbType represents format.
type OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry_OspfExtLsdbType string

const (
    OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry_OspfExtLsdbType_asExternalLink OSPFMIB_OspfExtLsdbTable_OspfExtLsdbEntry_OspfExtLsdbType = "asExternalLink"
)

// OSPFMIB_OspfAreaAggregateTable
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
type OSPFMIB_OspfAreaAggregateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single area aggregate entry.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry.
    OspfAreaAggregateEntry []*OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry
}

func (ospfAreaAggregateTable *OSPFMIB_OspfAreaAggregateTable) GetEntityData() *types.CommonEntityData {
    ospfAreaAggregateTable.EntityData.YFilter = ospfAreaAggregateTable.YFilter
    ospfAreaAggregateTable.EntityData.YangName = "ospfAreaAggregateTable"
    ospfAreaAggregateTable.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaAggregateTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfAreaAggregateTable.EntityData.SegmentPath = "ospfAreaAggregateTable"
    ospfAreaAggregateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaAggregateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaAggregateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaAggregateTable.EntityData.Children = types.NewOrderedMap()
    ospfAreaAggregateTable.EntityData.Children.Append("ospfAreaAggregateEntry", types.YChild{"OspfAreaAggregateEntry", nil})
    for i := range ospfAreaAggregateTable.OspfAreaAggregateEntry {
        ospfAreaAggregateTable.EntityData.Children.Append(types.GetSegmentPath(ospfAreaAggregateTable.OspfAreaAggregateEntry[i]), types.YChild{"OspfAreaAggregateEntry", ospfAreaAggregateTable.OspfAreaAggregateEntry[i]})
    }
    ospfAreaAggregateTable.EntityData.Leafs = types.NewOrderedMap()

    ospfAreaAggregateTable.EntityData.YListKeys = []string {}

    return &(ospfAreaAggregateTable.EntityData)
}

// OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry
// A single area aggregate entry.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The area within which the address aggregate is to
    // be found. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaAggregateAreaID interface{}

    // This attribute is a key. The type of the address aggregate.  This field
    // specifies the Lsdb type that this address aggregate applies to. The type is
    // OspfAreaAggregateLsdbType.
    OspfAreaAggregateLsdbType interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaAggregateNet interface{}

    // This attribute is a key. The subnet mask that pertains to the net or
    // subnet. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaAggregateMask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    OspfAreaAggregateStatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated aggregate (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is OspfAreaAggregateEffect.
    OspfAreaAggregateEffect interface{}

    // External route tag to be included in NSSA (type-7) LSAs. The type is
    // interface{} with range: 0..4294967295.
    OspfAreaAggregateExtRouteTag interface{}
}

func (ospfAreaAggregateEntry *OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry) GetEntityData() *types.CommonEntityData {
    ospfAreaAggregateEntry.EntityData.YFilter = ospfAreaAggregateEntry.YFilter
    ospfAreaAggregateEntry.EntityData.YangName = "ospfAreaAggregateEntry"
    ospfAreaAggregateEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaAggregateEntry.EntityData.ParentYangName = "ospfAreaAggregateTable"
    ospfAreaAggregateEntry.EntityData.SegmentPath = "ospfAreaAggregateEntry" + types.AddKeyToken(ospfAreaAggregateEntry.OspfAreaAggregateAreaID, "ospfAreaAggregateAreaID") + types.AddKeyToken(ospfAreaAggregateEntry.OspfAreaAggregateLsdbType, "ospfAreaAggregateLsdbType") + types.AddKeyToken(ospfAreaAggregateEntry.OspfAreaAggregateNet, "ospfAreaAggregateNet") + types.AddKeyToken(ospfAreaAggregateEntry.OspfAreaAggregateMask, "ospfAreaAggregateMask")
    ospfAreaAggregateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaAggregateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaAggregateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaAggregateEntry.EntityData.Children = types.NewOrderedMap()
    ospfAreaAggregateEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateAreaID", types.YLeaf{"OspfAreaAggregateAreaID", ospfAreaAggregateEntry.OspfAreaAggregateAreaID})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateLsdbType", types.YLeaf{"OspfAreaAggregateLsdbType", ospfAreaAggregateEntry.OspfAreaAggregateLsdbType})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateNet", types.YLeaf{"OspfAreaAggregateNet", ospfAreaAggregateEntry.OspfAreaAggregateNet})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateMask", types.YLeaf{"OspfAreaAggregateMask", ospfAreaAggregateEntry.OspfAreaAggregateMask})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateStatus", types.YLeaf{"OspfAreaAggregateStatus", ospfAreaAggregateEntry.OspfAreaAggregateStatus})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateEffect", types.YLeaf{"OspfAreaAggregateEffect", ospfAreaAggregateEntry.OspfAreaAggregateEffect})
    ospfAreaAggregateEntry.EntityData.Leafs.Append("ospfAreaAggregateExtRouteTag", types.YLeaf{"OspfAreaAggregateExtRouteTag", ospfAreaAggregateEntry.OspfAreaAggregateExtRouteTag})

    ospfAreaAggregateEntry.EntityData.YListKeys = []string {"OspfAreaAggregateAreaID", "OspfAreaAggregateLsdbType", "OspfAreaAggregateNet", "OspfAreaAggregateMask"}

    return &(ospfAreaAggregateEntry.EntityData)
}

// OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect represents being advertised at all outside the area.
type OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect string

const (
    OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect_advertiseMatching OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect = "advertiseMatching"

    OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect_doNotAdvertiseMatching OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateEffect = "doNotAdvertiseMatching"
)

// OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType represents aggregate applies to.
type OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType string

const (
    OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType_summaryLink OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType = "summaryLink"

    OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType_nssaExternalLink OSPFMIB_OspfAreaAggregateTable_OspfAreaAggregateEntry_OspfAreaAggregateLsdbType = "nssaExternalLink"
)

// OSPFMIB_OspfLocalLsdbTable
// The OSPF Process's link-local link state database
// for non-virtual links.
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for non-virtual links.  The purpose is
// to allow link-local LSAs to be displayed for each
// non-virtual interface.  This table is implemented to
// support type-9 LSAs that are defined
// in 'The OSPF Opaque LSA Option'.
type OSPFMIB_OspfLocalLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry.
    OspfLocalLsdbEntry []*OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry
}

func (ospfLocalLsdbTable *OSPFMIB_OspfLocalLsdbTable) GetEntityData() *types.CommonEntityData {
    ospfLocalLsdbTable.EntityData.YFilter = ospfLocalLsdbTable.YFilter
    ospfLocalLsdbTable.EntityData.YangName = "ospfLocalLsdbTable"
    ospfLocalLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    ospfLocalLsdbTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfLocalLsdbTable.EntityData.SegmentPath = "ospfLocalLsdbTable"
    ospfLocalLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfLocalLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfLocalLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfLocalLsdbTable.EntityData.Children = types.NewOrderedMap()
    ospfLocalLsdbTable.EntityData.Children.Append("ospfLocalLsdbEntry", types.YChild{"OspfLocalLsdbEntry", nil})
    for i := range ospfLocalLsdbTable.OspfLocalLsdbEntry {
        ospfLocalLsdbTable.EntityData.Children.Append(types.GetSegmentPath(ospfLocalLsdbTable.OspfLocalLsdbEntry[i]), types.YChild{"OspfLocalLsdbEntry", ospfLocalLsdbTable.OspfLocalLsdbEntry[i]})
    }
    ospfLocalLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    ospfLocalLsdbTable.EntityData.YListKeys = []string {}

    return &(ospfLocalLsdbTable.EntityData)
}

// OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry
// A single link state advertisement.
type OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLocalLsdbIpAddress interface{}

    // This attribute is a key. The interface index of the interface from which
    // the LSA was received if the interface is unnumbered. The type is
    // interface{} with range: 0..2147483647.
    OspfLocalLsdbAddressLessIf interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // OspfLocalLsdbType.
    OspfLocalLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLocalLsdbLsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfLocalLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    OspfLocalLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    OspfLocalLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    OspfLocalLsdbChecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    OspfLocalLsdbAdvertisement interface{}
}

func (ospfLocalLsdbEntry *OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry) GetEntityData() *types.CommonEntityData {
    ospfLocalLsdbEntry.EntityData.YFilter = ospfLocalLsdbEntry.YFilter
    ospfLocalLsdbEntry.EntityData.YangName = "ospfLocalLsdbEntry"
    ospfLocalLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfLocalLsdbEntry.EntityData.ParentYangName = "ospfLocalLsdbTable"
    ospfLocalLsdbEntry.EntityData.SegmentPath = "ospfLocalLsdbEntry" + types.AddKeyToken(ospfLocalLsdbEntry.OspfLocalLsdbIpAddress, "ospfLocalLsdbIpAddress") + types.AddKeyToken(ospfLocalLsdbEntry.OspfLocalLsdbAddressLessIf, "ospfLocalLsdbAddressLessIf") + types.AddKeyToken(ospfLocalLsdbEntry.OspfLocalLsdbType, "ospfLocalLsdbType") + types.AddKeyToken(ospfLocalLsdbEntry.OspfLocalLsdbLsid, "ospfLocalLsdbLsid") + types.AddKeyToken(ospfLocalLsdbEntry.OspfLocalLsdbRouterId, "ospfLocalLsdbRouterId")
    ospfLocalLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfLocalLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfLocalLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfLocalLsdbEntry.EntityData.Children = types.NewOrderedMap()
    ospfLocalLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbIpAddress", types.YLeaf{"OspfLocalLsdbIpAddress", ospfLocalLsdbEntry.OspfLocalLsdbIpAddress})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbAddressLessIf", types.YLeaf{"OspfLocalLsdbAddressLessIf", ospfLocalLsdbEntry.OspfLocalLsdbAddressLessIf})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbType", types.YLeaf{"OspfLocalLsdbType", ospfLocalLsdbEntry.OspfLocalLsdbType})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbLsid", types.YLeaf{"OspfLocalLsdbLsid", ospfLocalLsdbEntry.OspfLocalLsdbLsid})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbRouterId", types.YLeaf{"OspfLocalLsdbRouterId", ospfLocalLsdbEntry.OspfLocalLsdbRouterId})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbSequence", types.YLeaf{"OspfLocalLsdbSequence", ospfLocalLsdbEntry.OspfLocalLsdbSequence})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbAge", types.YLeaf{"OspfLocalLsdbAge", ospfLocalLsdbEntry.OspfLocalLsdbAge})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbChecksum", types.YLeaf{"OspfLocalLsdbChecksum", ospfLocalLsdbEntry.OspfLocalLsdbChecksum})
    ospfLocalLsdbEntry.EntityData.Leafs.Append("ospfLocalLsdbAdvertisement", types.YLeaf{"OspfLocalLsdbAdvertisement", ospfLocalLsdbEntry.OspfLocalLsdbAdvertisement})

    ospfLocalLsdbEntry.EntityData.YListKeys = []string {"OspfLocalLsdbIpAddress", "OspfLocalLsdbAddressLessIf", "OspfLocalLsdbType", "OspfLocalLsdbLsid", "OspfLocalLsdbRouterId"}

    return &(ospfLocalLsdbEntry.EntityData)
}

// OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry_OspfLocalLsdbType represents advertisement format.
type OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry_OspfLocalLsdbType string

const (
    OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry_OspfLocalLsdbType_localOpaqueLink OSPFMIB_OspfLocalLsdbTable_OspfLocalLsdbEntry_OspfLocalLsdbType = "localOpaqueLink"
)

// OSPFMIB_OspfVirtLocalLsdbTable
// The OSPF Process's link-local link state database
// for virtual links.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for virtual links.  The purpose is to
// allow link-local LSAs to be displayed for each virtual
// interface.  This table is implemented to support type-9 LSAs
// that are defined in 'The OSPF Opaque LSA Option'.
type OSPFMIB_OspfVirtLocalLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry.
    OspfVirtLocalLsdbEntry []*OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry
}

func (ospfVirtLocalLsdbTable *OSPFMIB_OspfVirtLocalLsdbTable) GetEntityData() *types.CommonEntityData {
    ospfVirtLocalLsdbTable.EntityData.YFilter = ospfVirtLocalLsdbTable.YFilter
    ospfVirtLocalLsdbTable.EntityData.YangName = "ospfVirtLocalLsdbTable"
    ospfVirtLocalLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtLocalLsdbTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfVirtLocalLsdbTable.EntityData.SegmentPath = "ospfVirtLocalLsdbTable"
    ospfVirtLocalLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtLocalLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtLocalLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtLocalLsdbTable.EntityData.Children = types.NewOrderedMap()
    ospfVirtLocalLsdbTable.EntityData.Children.Append("ospfVirtLocalLsdbEntry", types.YChild{"OspfVirtLocalLsdbEntry", nil})
    for i := range ospfVirtLocalLsdbTable.OspfVirtLocalLsdbEntry {
        ospfVirtLocalLsdbTable.EntityData.Children.Append(types.GetSegmentPath(ospfVirtLocalLsdbTable.OspfVirtLocalLsdbEntry[i]), types.YChild{"OspfVirtLocalLsdbEntry", ospfVirtLocalLsdbTable.OspfVirtLocalLsdbEntry[i]})
    }
    ospfVirtLocalLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    ospfVirtLocalLsdbTable.EntityData.YListKeys = []string {}

    return &(ospfVirtLocalLsdbTable.EntityData)
}

// OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry
// A single link state advertisement.
type OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtLocalLsdbTransitArea interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtLocalLsdbNeighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // OspfVirtLocalLsdbType.
    OspfVirtLocalLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtLocalLsdbLsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfVirtLocalLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    OspfVirtLocalLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    OspfVirtLocalLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that  an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    OspfVirtLocalLsdbChecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    OspfVirtLocalLsdbAdvertisement interface{}
}

func (ospfVirtLocalLsdbEntry *OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry) GetEntityData() *types.CommonEntityData {
    ospfVirtLocalLsdbEntry.EntityData.YFilter = ospfVirtLocalLsdbEntry.YFilter
    ospfVirtLocalLsdbEntry.EntityData.YangName = "ospfVirtLocalLsdbEntry"
    ospfVirtLocalLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfVirtLocalLsdbEntry.EntityData.ParentYangName = "ospfVirtLocalLsdbTable"
    ospfVirtLocalLsdbEntry.EntityData.SegmentPath = "ospfVirtLocalLsdbEntry" + types.AddKeyToken(ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbTransitArea, "ospfVirtLocalLsdbTransitArea") + types.AddKeyToken(ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbNeighbor, "ospfVirtLocalLsdbNeighbor") + types.AddKeyToken(ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbType, "ospfVirtLocalLsdbType") + types.AddKeyToken(ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbLsid, "ospfVirtLocalLsdbLsid") + types.AddKeyToken(ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbRouterId, "ospfVirtLocalLsdbRouterId")
    ospfVirtLocalLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfVirtLocalLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfVirtLocalLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfVirtLocalLsdbEntry.EntityData.Children = types.NewOrderedMap()
    ospfVirtLocalLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbTransitArea", types.YLeaf{"OspfVirtLocalLsdbTransitArea", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbTransitArea})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbNeighbor", types.YLeaf{"OspfVirtLocalLsdbNeighbor", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbNeighbor})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbType", types.YLeaf{"OspfVirtLocalLsdbType", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbType})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbLsid", types.YLeaf{"OspfVirtLocalLsdbLsid", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbLsid})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbRouterId", types.YLeaf{"OspfVirtLocalLsdbRouterId", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbRouterId})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbSequence", types.YLeaf{"OspfVirtLocalLsdbSequence", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbSequence})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbAge", types.YLeaf{"OspfVirtLocalLsdbAge", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbAge})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbChecksum", types.YLeaf{"OspfVirtLocalLsdbChecksum", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbChecksum})
    ospfVirtLocalLsdbEntry.EntityData.Leafs.Append("ospfVirtLocalLsdbAdvertisement", types.YLeaf{"OspfVirtLocalLsdbAdvertisement", ospfVirtLocalLsdbEntry.OspfVirtLocalLsdbAdvertisement})

    ospfVirtLocalLsdbEntry.EntityData.YListKeys = []string {"OspfVirtLocalLsdbTransitArea", "OspfVirtLocalLsdbNeighbor", "OspfVirtLocalLsdbType", "OspfVirtLocalLsdbLsid", "OspfVirtLocalLsdbRouterId"}

    return &(ospfVirtLocalLsdbEntry.EntityData)
}

// OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry_OspfVirtLocalLsdbType represents advertisement format.
type OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry_OspfVirtLocalLsdbType string

const (
    OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry_OspfVirtLocalLsdbType_localOpaqueLink OSPFMIB_OspfVirtLocalLsdbTable_OspfVirtLocalLsdbEntry_OspfVirtLocalLsdbType = "localOpaqueLink"
)

// OSPFMIB_OspfAsLsdbTable
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
type OSPFMIB_OspfAsLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry.
    OspfAsLsdbEntry []*OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry
}

func (ospfAsLsdbTable *OSPFMIB_OspfAsLsdbTable) GetEntityData() *types.CommonEntityData {
    ospfAsLsdbTable.EntityData.YFilter = ospfAsLsdbTable.YFilter
    ospfAsLsdbTable.EntityData.YangName = "ospfAsLsdbTable"
    ospfAsLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    ospfAsLsdbTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfAsLsdbTable.EntityData.SegmentPath = "ospfAsLsdbTable"
    ospfAsLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAsLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAsLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAsLsdbTable.EntityData.Children = types.NewOrderedMap()
    ospfAsLsdbTable.EntityData.Children.Append("ospfAsLsdbEntry", types.YChild{"OspfAsLsdbEntry", nil})
    for i := range ospfAsLsdbTable.OspfAsLsdbEntry {
        ospfAsLsdbTable.EntityData.Children.Append(types.GetSegmentPath(ospfAsLsdbTable.OspfAsLsdbEntry[i]), types.YChild{"OspfAsLsdbEntry", ospfAsLsdbTable.OspfAsLsdbEntry[i]})
    }
    ospfAsLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    ospfAsLsdbTable.EntityData.YListKeys = []string {}

    return &(ospfAsLsdbTable.EntityData)
}

// OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry
// A single link state advertisement.
type OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // OspfAsLsdbType.
    OspfAsLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address;  it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAsLsdbLsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAsLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    OspfAsLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    OspfAsLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    OspfAsLsdbChecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    OspfAsLsdbAdvertisement interface{}
}

func (ospfAsLsdbEntry *OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry) GetEntityData() *types.CommonEntityData {
    ospfAsLsdbEntry.EntityData.YFilter = ospfAsLsdbEntry.YFilter
    ospfAsLsdbEntry.EntityData.YangName = "ospfAsLsdbEntry"
    ospfAsLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfAsLsdbEntry.EntityData.ParentYangName = "ospfAsLsdbTable"
    ospfAsLsdbEntry.EntityData.SegmentPath = "ospfAsLsdbEntry" + types.AddKeyToken(ospfAsLsdbEntry.OspfAsLsdbType, "ospfAsLsdbType") + types.AddKeyToken(ospfAsLsdbEntry.OspfAsLsdbLsid, "ospfAsLsdbLsid") + types.AddKeyToken(ospfAsLsdbEntry.OspfAsLsdbRouterId, "ospfAsLsdbRouterId")
    ospfAsLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAsLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAsLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAsLsdbEntry.EntityData.Children = types.NewOrderedMap()
    ospfAsLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbType", types.YLeaf{"OspfAsLsdbType", ospfAsLsdbEntry.OspfAsLsdbType})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbLsid", types.YLeaf{"OspfAsLsdbLsid", ospfAsLsdbEntry.OspfAsLsdbLsid})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbRouterId", types.YLeaf{"OspfAsLsdbRouterId", ospfAsLsdbEntry.OspfAsLsdbRouterId})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbSequence", types.YLeaf{"OspfAsLsdbSequence", ospfAsLsdbEntry.OspfAsLsdbSequence})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbAge", types.YLeaf{"OspfAsLsdbAge", ospfAsLsdbEntry.OspfAsLsdbAge})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbChecksum", types.YLeaf{"OspfAsLsdbChecksum", ospfAsLsdbEntry.OspfAsLsdbChecksum})
    ospfAsLsdbEntry.EntityData.Leafs.Append("ospfAsLsdbAdvertisement", types.YLeaf{"OspfAsLsdbAdvertisement", ospfAsLsdbEntry.OspfAsLsdbAdvertisement})

    ospfAsLsdbEntry.EntityData.YListKeys = []string {"OspfAsLsdbType", "OspfAsLsdbLsid", "OspfAsLsdbRouterId"}

    return &(ospfAsLsdbEntry.EntityData)
}

// OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType represents advertisement format.
type OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType string

const (
    OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType_asExternalLink OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType = "asExternalLink"

    OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType_asOpaqueLink OSPFMIB_OspfAsLsdbTable_OspfAsLsdbEntry_OspfAsLsdbType = "asOpaqueLink"
)

// OSPFMIB_OspfAreaLsaCountTable
// This table maintains per-area, per-LSA-type counters
type OSPFMIB_OspfAreaLsaCountTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with a number of link advertisements  of a given type for a given
    // area. The type is slice of
    // OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry.
    OspfAreaLsaCountEntry []*OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry
}

func (ospfAreaLsaCountTable *OSPFMIB_OspfAreaLsaCountTable) GetEntityData() *types.CommonEntityData {
    ospfAreaLsaCountTable.EntityData.YFilter = ospfAreaLsaCountTable.YFilter
    ospfAreaLsaCountTable.EntityData.YangName = "ospfAreaLsaCountTable"
    ospfAreaLsaCountTable.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaLsaCountTable.EntityData.ParentYangName = "OSPF-MIB"
    ospfAreaLsaCountTable.EntityData.SegmentPath = "ospfAreaLsaCountTable"
    ospfAreaLsaCountTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaLsaCountTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaLsaCountTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaLsaCountTable.EntityData.Children = types.NewOrderedMap()
    ospfAreaLsaCountTable.EntityData.Children.Append("ospfAreaLsaCountEntry", types.YChild{"OspfAreaLsaCountEntry", nil})
    for i := range ospfAreaLsaCountTable.OspfAreaLsaCountEntry {
        ospfAreaLsaCountTable.EntityData.Children.Append(types.GetSegmentPath(ospfAreaLsaCountTable.OspfAreaLsaCountEntry[i]), types.YChild{"OspfAreaLsaCountEntry", ospfAreaLsaCountTable.OspfAreaLsaCountEntry[i]})
    }
    ospfAreaLsaCountTable.EntityData.Leafs = types.NewOrderedMap()

    ospfAreaLsaCountTable.EntityData.YListKeys = []string {}

    return &(ospfAreaLsaCountTable.EntityData)
}

// OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry
// An entry with a number of link advertisements
// 
// of a given type for a given area.
type OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This entry Area ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfAreaLsaCountAreaId interface{}

    // This attribute is a key. This entry LSA type. The type is
    // OspfAreaLsaCountLsaType.
    OspfAreaLsaCountLsaType interface{}

    // Number of LSAs of a given type for a given area. The type is interface{}
    // with range: 0..4294967295.
    OspfAreaLsaCountNumber interface{}
}

func (ospfAreaLsaCountEntry *OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry) GetEntityData() *types.CommonEntityData {
    ospfAreaLsaCountEntry.EntityData.YFilter = ospfAreaLsaCountEntry.YFilter
    ospfAreaLsaCountEntry.EntityData.YangName = "ospfAreaLsaCountEntry"
    ospfAreaLsaCountEntry.EntityData.BundleName = "cisco_ios_xe"
    ospfAreaLsaCountEntry.EntityData.ParentYangName = "ospfAreaLsaCountTable"
    ospfAreaLsaCountEntry.EntityData.SegmentPath = "ospfAreaLsaCountEntry" + types.AddKeyToken(ospfAreaLsaCountEntry.OspfAreaLsaCountAreaId, "ospfAreaLsaCountAreaId") + types.AddKeyToken(ospfAreaLsaCountEntry.OspfAreaLsaCountLsaType, "ospfAreaLsaCountLsaType")
    ospfAreaLsaCountEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfAreaLsaCountEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfAreaLsaCountEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfAreaLsaCountEntry.EntityData.Children = types.NewOrderedMap()
    ospfAreaLsaCountEntry.EntityData.Leafs = types.NewOrderedMap()
    ospfAreaLsaCountEntry.EntityData.Leafs.Append("ospfAreaLsaCountAreaId", types.YLeaf{"OspfAreaLsaCountAreaId", ospfAreaLsaCountEntry.OspfAreaLsaCountAreaId})
    ospfAreaLsaCountEntry.EntityData.Leafs.Append("ospfAreaLsaCountLsaType", types.YLeaf{"OspfAreaLsaCountLsaType", ospfAreaLsaCountEntry.OspfAreaLsaCountLsaType})
    ospfAreaLsaCountEntry.EntityData.Leafs.Append("ospfAreaLsaCountNumber", types.YLeaf{"OspfAreaLsaCountNumber", ospfAreaLsaCountEntry.OspfAreaLsaCountNumber})

    ospfAreaLsaCountEntry.EntityData.YListKeys = []string {"OspfAreaLsaCountAreaId", "OspfAreaLsaCountLsaType"}

    return &(ospfAreaLsaCountEntry.EntityData)
}

// OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType represents This entry LSA type.
type OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType string

const (
    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_routerLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "routerLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_networkLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "networkLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_summaryLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "summaryLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_asSummaryLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "asSummaryLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_multicastLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "multicastLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_nssaExternalLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "nssaExternalLink"

    OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType_areaOpaqueLink OSPFMIB_OspfAreaLsaCountTable_OspfAreaLsaCountEntry_OspfAreaLsaCountLsaType = "areaOpaqueLink"
)

