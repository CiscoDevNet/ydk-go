// This MIB module contains the management objects for the 
// CMTS control of the IP4 and IPv6 traffic with origin and 
// destination to CMs and/or CPEs behind the CM.
package docs_submgt3_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_submgt3_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-SUBMGT3-MIB DOCS-SUBMGT3-MIB}", reflect.TypeOf(DOCSSUBMGT3MIB{}))
    ydk.RegisterEntity("DOCS-SUBMGT3-MIB:DOCS-SUBMGT3-MIB", reflect.TypeOf(DOCSSUBMGT3MIB{}))
}

// DOCSSUBMGT3MIB
type DOCSSUBMGT3MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DocsSubmgt3Base DOCSSUBMGT3MIB_DocsSubmgt3Base

    // This object maintains per-CM traffic policies enforced  by the CMTS. The
    // CMTS acquires the CM traffic policies  through the CM registration process,
    // or in the  absence of some or all of those parameters, from the  Base
    // object. The CM information and controls are meaningful  and used by the
    // CMTS, but only after the CM is  operational.
    DocsSubmgt3CpeCtrlTable DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable

    // This object defines the list of IP Addresses behind  the CM known by the
    // CMTS.   If the Active attribute of the CpeCtrl object associated  with a CM
    // is set to 'true' and the CMTS receives an  IP packet from a CM that
    // contains a source IP address that  does not match one of the CPE IP
    // addresses associated  with this CM, one of two things occurs. If the number
    // of CPE IPs is less than the MaxCpeIp of the CpeCtrl object  for that CM,
    // the source IP address is added to this  object and the packet is forwarded;
    // otherwise, the  packet is dropped.
    DocsSubmgt3CpeIpTable DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable

    // This object defines the set of downstream and upstream  filter groups that
    // the CMTS applies to traffic associated  with that CM.
    DocsSubmgt3GrpTable DOCSSUBMGT3MIB_DocsSubmgt3GrpTable

    // This object describes a set of filter or classifier  criteria. Classifiers
    // are assigned by group to the  individual CMs. That assignment is made via
    // the 'Subscriber  Management TLVs' encodings sent upstream from  the CM to
    // the CMTS during registration or in their  absence, default values
    // configured in the CMTS.  A Filter Group ID (GrpId) is a set of rules that
    // correspond  to the expansion of a UDC Group ID into UDC individual  
    // classification rules.  The Filter Group Ids are generated   whenever the
    // CMTS is configured to send UDCs during the CM   registration process.
    // Implementation of L2 classification   criteria is optional for the CMTS;
    // LLC/MAC upstream and   downstream filter criteria can be ignored during the
    // packet  matching process.
    DocsSubmgt3FilterGrpTable DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable
}

func (dOCSSUBMGT3MIB *DOCSSUBMGT3MIB) GetEntityData() *types.CommonEntityData {
    dOCSSUBMGT3MIB.EntityData.YFilter = dOCSSUBMGT3MIB.YFilter
    dOCSSUBMGT3MIB.EntityData.YangName = "DOCS-SUBMGT3-MIB"
    dOCSSUBMGT3MIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSSUBMGT3MIB.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    dOCSSUBMGT3MIB.EntityData.SegmentPath = "DOCS-SUBMGT3-MIB:DOCS-SUBMGT3-MIB"
    dOCSSUBMGT3MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSSUBMGT3MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSSUBMGT3MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSSUBMGT3MIB.EntityData.Children = types.NewOrderedMap()
    dOCSSUBMGT3MIB.EntityData.Children.Append("docsSubmgt3Base", types.YChild{"DocsSubmgt3Base", &dOCSSUBMGT3MIB.DocsSubmgt3Base})
    dOCSSUBMGT3MIB.EntityData.Children.Append("docsSubmgt3CpeCtrlTable", types.YChild{"DocsSubmgt3CpeCtrlTable", &dOCSSUBMGT3MIB.DocsSubmgt3CpeCtrlTable})
    dOCSSUBMGT3MIB.EntityData.Children.Append("docsSubmgt3CpeIpTable", types.YChild{"DocsSubmgt3CpeIpTable", &dOCSSUBMGT3MIB.DocsSubmgt3CpeIpTable})
    dOCSSUBMGT3MIB.EntityData.Children.Append("docsSubmgt3GrpTable", types.YChild{"DocsSubmgt3GrpTable", &dOCSSUBMGT3MIB.DocsSubmgt3GrpTable})
    dOCSSUBMGT3MIB.EntityData.Children.Append("docsSubmgt3FilterGrpTable", types.YChild{"DocsSubmgt3FilterGrpTable", &dOCSSUBMGT3MIB.DocsSubmgt3FilterGrpTable})
    dOCSSUBMGT3MIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSSUBMGT3MIB.EntityData.YListKeys = []string {}

    return &(dOCSSUBMGT3MIB.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3Base
type DOCSSUBMGT3MIB_DocsSubmgt3Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute represents the maximum number of IPv4   Addresses allowed
    // for the CM's CPEs if not signaled in the   registration process. The type
    // is interface{} with range: 0..1023.
    DocsSubmgt3BaseCpeMaxIpv4Def interface{}

    // This attribute represents the maximum number of IPv6   prefixes allowed for
    // the CM's CPEs if not signaled in  the registration process. The type is
    // interface{} with range: 0..1023.
    DocsSubmgt3BaseCpeMaxIpv6PrefixDef interface{}

    // This attribute represents the default value for enabling  Subscriber
    // Management filters and controls  in the CM if the parameter is not signaled
    // in the DOCSIS  Registration process. The type is bool.
    DocsSubmgt3BaseCpeActiveDef interface{}

    // This attribute represents the default value for enabling  the CPE learning
    // process for the CM if the parameter  is not signaled in the DOCSIS
    // Registration process. The type is bool.
    DocsSubmgt3BaseCpeLearnableDef interface{}

    // This attribute represents the default value for the  subscriber (CPE)
    // downstream filter group for the  CM if the parameter is not signaled in the
    // DOCSIS Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseSubFilterDownDef interface{}

    // This attribute represents the default value for the  subscriber (CPE)
    // upstream filter group for the CM  if the parameter is not signaled in the
    // DOCSIS Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseSubFilterUpDef interface{}

    // This attribute represents the default value for the  CM stack downstream
    // filter group applying to the CM  if the parameter is not signaled in the
    // DOCSIS Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseCmFilterDownDef interface{}

    // This attribute represents the default value for the  CM stack upstream
    // filter group applying to the CM if  the parameter is not signaled in the
    // DOCSIS Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseCmFilterUpDef interface{}

    // This attribute represents the default value for the  PS or eRouter
    // downstream filter group for the CM if  the parameter is not signaled in the
    // DOCSIS Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BasePsFilterDownDef interface{}

    // This attribute represents the default value for the  PS or eRouter upstream
    // filter group for the CM if the  parameter is not signaled in the DOCSIS
    // Registration  process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BasePsFilterUpDef interface{}

    // This attribute represents the default value for the  MTA downstream filter
    // group for the CM if the parameter  is not signaled in the DOCSIS
    // Registration process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseMtaFilterDownDef interface{}

    // This attribute represents the default value for the  MTA upstream filter
    // group for the CM if the parameter  is not signaled in the DOCSIS
    // Registration process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseMtaFilterUpDef interface{}

    // This attribute represents the default value for the  STB downstream filter
    // group for the CM if the parameter  is not signaled in the DOCSIS
    // Registration process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseStbFilterDownDef interface{}

    // This attribute represents the default value for the  STB upstream filter
    // group for the CM if the parameter  is not signaled in the DOCSIS
    // Registration process. The type is interface{} with range: 0..1024.
    DocsSubmgt3BaseStbFilterUpDef interface{}
}

func (docsSubmgt3Base *DOCSSUBMGT3MIB_DocsSubmgt3Base) GetEntityData() *types.CommonEntityData {
    docsSubmgt3Base.EntityData.YFilter = docsSubmgt3Base.YFilter
    docsSubmgt3Base.EntityData.YangName = "docsSubmgt3Base"
    docsSubmgt3Base.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3Base.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    docsSubmgt3Base.EntityData.SegmentPath = "docsSubmgt3Base"
    docsSubmgt3Base.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3Base.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3Base.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3Base.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3Base.EntityData.Leafs = types.NewOrderedMap()
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCpeMaxIpv4Def", types.YLeaf{"DocsSubmgt3BaseCpeMaxIpv4Def", docsSubmgt3Base.DocsSubmgt3BaseCpeMaxIpv4Def})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCpeMaxIpv6PrefixDef", types.YLeaf{"DocsSubmgt3BaseCpeMaxIpv6PrefixDef", docsSubmgt3Base.DocsSubmgt3BaseCpeMaxIpv6PrefixDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCpeActiveDef", types.YLeaf{"DocsSubmgt3BaseCpeActiveDef", docsSubmgt3Base.DocsSubmgt3BaseCpeActiveDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCpeLearnableDef", types.YLeaf{"DocsSubmgt3BaseCpeLearnableDef", docsSubmgt3Base.DocsSubmgt3BaseCpeLearnableDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseSubFilterDownDef", types.YLeaf{"DocsSubmgt3BaseSubFilterDownDef", docsSubmgt3Base.DocsSubmgt3BaseSubFilterDownDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseSubFilterUpDef", types.YLeaf{"DocsSubmgt3BaseSubFilterUpDef", docsSubmgt3Base.DocsSubmgt3BaseSubFilterUpDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCmFilterDownDef", types.YLeaf{"DocsSubmgt3BaseCmFilterDownDef", docsSubmgt3Base.DocsSubmgt3BaseCmFilterDownDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseCmFilterUpDef", types.YLeaf{"DocsSubmgt3BaseCmFilterUpDef", docsSubmgt3Base.DocsSubmgt3BaseCmFilterUpDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BasePsFilterDownDef", types.YLeaf{"DocsSubmgt3BasePsFilterDownDef", docsSubmgt3Base.DocsSubmgt3BasePsFilterDownDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BasePsFilterUpDef", types.YLeaf{"DocsSubmgt3BasePsFilterUpDef", docsSubmgt3Base.DocsSubmgt3BasePsFilterUpDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseMtaFilterDownDef", types.YLeaf{"DocsSubmgt3BaseMtaFilterDownDef", docsSubmgt3Base.DocsSubmgt3BaseMtaFilterDownDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseMtaFilterUpDef", types.YLeaf{"DocsSubmgt3BaseMtaFilterUpDef", docsSubmgt3Base.DocsSubmgt3BaseMtaFilterUpDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseStbFilterDownDef", types.YLeaf{"DocsSubmgt3BaseStbFilterDownDef", docsSubmgt3Base.DocsSubmgt3BaseStbFilterDownDef})
    docsSubmgt3Base.EntityData.Leafs.Append("docsSubmgt3BaseStbFilterUpDef", types.YLeaf{"DocsSubmgt3BaseStbFilterUpDef", docsSubmgt3Base.DocsSubmgt3BaseStbFilterUpDef})

    docsSubmgt3Base.EntityData.YListKeys = []string {}

    return &(docsSubmgt3Base.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable
// This object maintains per-CM traffic policies enforced 
// by the CMTS. The CMTS acquires the CM traffic policies 
// through the CM registration process, or in the 
// absence of some or all of those parameters, from the 
// Base object. The CM information and controls are meaningful 
// and used by the CMTS, but only after the CM is 
// operational.
type DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsSubmgt3CpeCtrlTable.  The CMTS does not persist
    // the instances of the CpeCtrl  object across reinitializations. The type is
    // slice of DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable_DocsSubmgt3CpeCtrlEntry.
    DocsSubmgt3CpeCtrlEntry []*DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable_DocsSubmgt3CpeCtrlEntry
}

func (docsSubmgt3CpeCtrlTable *DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable) GetEntityData() *types.CommonEntityData {
    docsSubmgt3CpeCtrlTable.EntityData.YFilter = docsSubmgt3CpeCtrlTable.YFilter
    docsSubmgt3CpeCtrlTable.EntityData.YangName = "docsSubmgt3CpeCtrlTable"
    docsSubmgt3CpeCtrlTable.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3CpeCtrlTable.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    docsSubmgt3CpeCtrlTable.EntityData.SegmentPath = "docsSubmgt3CpeCtrlTable"
    docsSubmgt3CpeCtrlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3CpeCtrlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3CpeCtrlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3CpeCtrlTable.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3CpeCtrlTable.EntityData.Children.Append("docsSubmgt3CpeCtrlEntry", types.YChild{"DocsSubmgt3CpeCtrlEntry", nil})
    for i := range docsSubmgt3CpeCtrlTable.DocsSubmgt3CpeCtrlEntry {
        docsSubmgt3CpeCtrlTable.EntityData.Children.Append(types.GetSegmentPath(docsSubmgt3CpeCtrlTable.DocsSubmgt3CpeCtrlEntry[i]), types.YChild{"DocsSubmgt3CpeCtrlEntry", docsSubmgt3CpeCtrlTable.DocsSubmgt3CpeCtrlEntry[i]})
    }
    docsSubmgt3CpeCtrlTable.EntityData.Leafs = types.NewOrderedMap()

    docsSubmgt3CpeCtrlTable.EntityData.YListKeys = []string {}

    return &(docsSubmgt3CpeCtrlTable.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable_DocsSubmgt3CpeCtrlEntry
// The conceptual row of docsSubmgt3CpeCtrlTable. 
// The CMTS does not persist the instances of the CpeCtrl 
// object across reinitializations.
type DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable_DocsSubmgt3CpeCtrlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry_DocsIf3CmtsCmRegStatusId
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute represents the number of simultaneous  IP v4 addresses
    // permitted for CPE connected to the CM.  When the MaxCpeIpv4 attribute is
    // set to zero (0), all Ipv4 CPE  traffic from the CM is dropped. The CMTS
    // configures this  attribute with whichever of the 'Subscriber Management CPE
    // IPv4  List' or 'Subscriber Management Control-Max_CpeIPv4' signaled 
    // encodings is greater, or in the absence of all of those  provisioning
    // parameters, with the CpeMaxIp v4Def  from the Base object. This limit
    // applies to learned  and DOCSIS-provisioned entries but not to entries added
    // through some administrative process at the CMTS.  Note that this attribute
    // is only meaningful when the  Active attribute of the CM is set to 'true'.
    // The type is interface{} with range: 0..1023.
    DocsSubmgt3CpeCtrlMaxCpeIpv4 interface{}

    // This attribute represents the number of simultaneous  IPv6 prefixes
    // permitted for CPE connected to the  CM.   When the MaxCpeIpv6Prefix is set
    // to zero (0), all IPv6 CPE   traffic from the CM is dropped. The CMTS
    // configures this   attribute with whichever of the 'Subscriber Management
    // CPE IPv6 List'   or'Subscriber Management Control Max Cpe IPv6 Prefix'  
    // signaled encodings is greater, or in the absence of all of those 
    // provisioning parameters, with the CpeMaxIpv6PrefixDef  from the Base
    // object. This limit applies to learned  and DOCSIS-provisioned entries but
    // not to entries added  through some administrative process at the CMTS. 
    // Note that this attribute is only meaningful when the  Active attribute of
    // the CM is set to 'true'. The type is interface{} with range: 0..1023.
    DocsSubmgt3CpeCtrlMaxCpeIpv6Prefix interface{}

    // This attribute controls the application of subscriber  management to this
    // CM. If this is set to 'true',  CMTS-based CPE control is active, and all
    // the actions  required by the various filter policies and controls  apply at
    // the CMTS. If this is set to false, no subscriber  management filtering is
    // done at the CMTS (but other  filters may apply). If not set through DOCSIS
    // provisioning,  this object defaults to the value of the Active  attribute
    // of the Base object. The type is bool.
    DocsSubmgt3CpeCtrlActive interface{}

    // This attribute controls whether the CMTS may learn  (and pass traffic for)
    // CPE IP addresses associated  with a CM. If this is set to 'true', the CMTS
    // may learn up  to the CM MaxCpeIp value less any DOCSIS-provisioned  entries
    // related to this CM. The nature of the learning  mechanism is not specified
    // here. If not set through  DOCSIS provisioning, this object defaults to the 
    // value of the CpeLearnableDef attribute from the Base  object. Note that
    // this attribute is only meaningful  if docsSubMgtCpeControlActive is 'true'
    // to enforce  a limit in the number of CPEs learned. CPE learning  is always
    // performed for the CMTS for security reasons. The type is bool.
    DocsSubmgt3CpeCtrlLearnable interface{}

    // If set to 'true', this attribute commands the CMTS  to delete the instances
    // denoted as 'learned' addresses  in the CpeIp object. This attribute always
    // returns  false on read. The type is bool.
    DocsSubmgt3CpeCtrlReset interface{}

    // This attribute represents the system Up Time of the  last set to 'true' of
    // the Reset attribute of this instance.  Zero if never reset. The type is
    // interface{} with range: 0..4294967295.
    DocsSubmgt3CpeCtrlLastReset interface{}
}

func (docsSubmgt3CpeCtrlEntry *DOCSSUBMGT3MIB_DocsSubmgt3CpeCtrlTable_DocsSubmgt3CpeCtrlEntry) GetEntityData() *types.CommonEntityData {
    docsSubmgt3CpeCtrlEntry.EntityData.YFilter = docsSubmgt3CpeCtrlEntry.YFilter
    docsSubmgt3CpeCtrlEntry.EntityData.YangName = "docsSubmgt3CpeCtrlEntry"
    docsSubmgt3CpeCtrlEntry.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3CpeCtrlEntry.EntityData.ParentYangName = "docsSubmgt3CpeCtrlTable"
    docsSubmgt3CpeCtrlEntry.EntityData.SegmentPath = "docsSubmgt3CpeCtrlEntry" + types.AddKeyToken(docsSubmgt3CpeCtrlEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId")
    docsSubmgt3CpeCtrlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3CpeCtrlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3CpeCtrlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3CpeCtrlEntry.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs = types.NewOrderedMap()
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsSubmgt3CpeCtrlEntry.DocsIf3CmtsCmRegStatusId})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlMaxCpeIpv4", types.YLeaf{"DocsSubmgt3CpeCtrlMaxCpeIpv4", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlMaxCpeIpv4})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlMaxCpeIpv6Prefix", types.YLeaf{"DocsSubmgt3CpeCtrlMaxCpeIpv6Prefix", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlMaxCpeIpv6Prefix})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlActive", types.YLeaf{"DocsSubmgt3CpeCtrlActive", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlActive})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlLearnable", types.YLeaf{"DocsSubmgt3CpeCtrlLearnable", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlLearnable})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlReset", types.YLeaf{"DocsSubmgt3CpeCtrlReset", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlReset})
    docsSubmgt3CpeCtrlEntry.EntityData.Leafs.Append("docsSubmgt3CpeCtrlLastReset", types.YLeaf{"DocsSubmgt3CpeCtrlLastReset", docsSubmgt3CpeCtrlEntry.DocsSubmgt3CpeCtrlLastReset})

    docsSubmgt3CpeCtrlEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId"}

    return &(docsSubmgt3CpeCtrlEntry.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable
// This object defines the list of IP Addresses behind 
// the CM known by the CMTS. 
// 
// If the Active attribute of the CpeCtrl object associated 
// with a CM is set to 'true' and the CMTS receives an 
// IP packet from a CM that contains a source IP address that 
// does not match one of the CPE IP addresses associated 
// with this CM, one of two things occurs. If the number 
// of CPE IPs is less than the MaxCpeIp of the CpeCtrl object 
// for that CM, the source IP address is added to this 
// object and the packet is forwarded; otherwise, the 
// packet is dropped.
type DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsSubmgt3CpeIpTable. The type is slice of
    // DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry.
    DocsSubmgt3CpeIpEntry []*DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry
}

func (docsSubmgt3CpeIpTable *DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable) GetEntityData() *types.CommonEntityData {
    docsSubmgt3CpeIpTable.EntityData.YFilter = docsSubmgt3CpeIpTable.YFilter
    docsSubmgt3CpeIpTable.EntityData.YangName = "docsSubmgt3CpeIpTable"
    docsSubmgt3CpeIpTable.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3CpeIpTable.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    docsSubmgt3CpeIpTable.EntityData.SegmentPath = "docsSubmgt3CpeIpTable"
    docsSubmgt3CpeIpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3CpeIpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3CpeIpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3CpeIpTable.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3CpeIpTable.EntityData.Children.Append("docsSubmgt3CpeIpEntry", types.YChild{"DocsSubmgt3CpeIpEntry", nil})
    for i := range docsSubmgt3CpeIpTable.DocsSubmgt3CpeIpEntry {
        docsSubmgt3CpeIpTable.EntityData.Children.Append(types.GetSegmentPath(docsSubmgt3CpeIpTable.DocsSubmgt3CpeIpEntry[i]), types.YChild{"DocsSubmgt3CpeIpEntry", docsSubmgt3CpeIpTable.DocsSubmgt3CpeIpEntry[i]})
    }
    docsSubmgt3CpeIpTable.EntityData.Leafs = types.NewOrderedMap()

    docsSubmgt3CpeIpTable.EntityData.YListKeys = []string {}

    return &(docsSubmgt3CpeIpTable.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry
// The conceptual row of docsSubmgt3CpeIpTable.
type DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry_DocsIf3CmtsCmRegStatusId
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute is a key. This attribute represents a unique identifier for 
    // a CPE IP of the CM. An instance of this attribute exists  for each CPE
    // provisioned in the 'Subscriber Management  CPE IPv4 Table' or 'Subscriber
    // Management CPE  IPv6 Table' encodings. An entry is created either through 
    // the included CPE IP addresses in the provisioning  object, or CPEs learned
    // from traffic sourced from the  CM. The type is interface{} with range:
    // 1..1023.
    DocsSubmgt3CpeIpId interface{}

    // The type of Internet address of the Addr attribute. The type is
    // InetAddressType.
    DocsSubmgt3CpeIpAddrType interface{}

    // This attribute represents the IP address either set  from provisioning or
    // learned via address gleaning  or other forwarding means. The type is string
    // with length: 0..255.
    DocsSubmgt3CpeIpAddr interface{}

    // This attribute represents the prefix length associated with  the IP subnet
    // prefix either set from provisioning or learned  via address gleaning or
    // other forwarding means. For IPv4 CPE  addresses this attribute generally
    // reports the value 32   (32 bits) to indicate a unicast IPv4 address. For
    // IPv6, this  attribute represents either an IPv6 unicast address  (128 bits,
    // equal to /128 prefix length) or a subnet prefix   length (for example 56
    // bits, equal to /56 prefix length). The type is interface{} with range:
    // 0..2040.
    DocsSubmgt3CpeIpAddrPrefixLen interface{}

    // This attribute is set to 'true' when the IP address  was learned from IP
    // packets sent upstream rather than  via the CM provisioning process. The
    // type is bool.
    DocsSubmgt3CpeIpLearned interface{}

    // This attribute represents the type of CPE based on  the following
    // classification below:               'cpe' Regular CPE clients.             
    // 'ps'  CableHome Portal Server (PS)               'mta' PacketCable
    // Multimedia Terminal Adapter (MTA)               'stb' Digital Set-top Box
    // (STB).               'tea' T1 Emulation adapter (TEA)              
    // 'erouter' Embedded Router (eRouter). The type is DocsSubmgt3CpeIpType.
    DocsSubmgt3CpeIpType interface{}
}

func (docsSubmgt3CpeIpEntry *DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry) GetEntityData() *types.CommonEntityData {
    docsSubmgt3CpeIpEntry.EntityData.YFilter = docsSubmgt3CpeIpEntry.YFilter
    docsSubmgt3CpeIpEntry.EntityData.YangName = "docsSubmgt3CpeIpEntry"
    docsSubmgt3CpeIpEntry.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3CpeIpEntry.EntityData.ParentYangName = "docsSubmgt3CpeIpTable"
    docsSubmgt3CpeIpEntry.EntityData.SegmentPath = "docsSubmgt3CpeIpEntry" + types.AddKeyToken(docsSubmgt3CpeIpEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId") + types.AddKeyToken(docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpId, "docsSubmgt3CpeIpId")
    docsSubmgt3CpeIpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3CpeIpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3CpeIpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3CpeIpEntry.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3CpeIpEntry.EntityData.Leafs = types.NewOrderedMap()
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsSubmgt3CpeIpEntry.DocsIf3CmtsCmRegStatusId})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpId", types.YLeaf{"DocsSubmgt3CpeIpId", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpId})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpAddrType", types.YLeaf{"DocsSubmgt3CpeIpAddrType", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpAddrType})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpAddr", types.YLeaf{"DocsSubmgt3CpeIpAddr", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpAddr})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpAddrPrefixLen", types.YLeaf{"DocsSubmgt3CpeIpAddrPrefixLen", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpAddrPrefixLen})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpLearned", types.YLeaf{"DocsSubmgt3CpeIpLearned", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpLearned})
    docsSubmgt3CpeIpEntry.EntityData.Leafs.Append("docsSubmgt3CpeIpType", types.YLeaf{"DocsSubmgt3CpeIpType", docsSubmgt3CpeIpEntry.DocsSubmgt3CpeIpType})

    docsSubmgt3CpeIpEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId", "DocsSubmgt3CpeIpId"}

    return &(docsSubmgt3CpeIpEntry.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType represents              'erouter' Embedded Router (eRouter)
type DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType string

const (
    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_cpe DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "cpe"

    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_ps DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "ps"

    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_mta DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "mta"

    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_stb DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "stb"

    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_tea DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "tea"

    DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType_erouter DOCSSUBMGT3MIB_DocsSubmgt3CpeIpTable_DocsSubmgt3CpeIpEntry_DocsSubmgt3CpeIpType = "erouter"
)

// DOCSSUBMGT3MIB_DocsSubmgt3GrpTable
// This object defines the set of downstream and upstream 
// filter groups that the CMTS applies to traffic associated 
// with that CM.
type DOCSSUBMGT3MIB_DocsSubmgt3GrpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsSubmgt3GrpTable.  The CMTS does not persist the
    // instances of the Grp  object across reinitializations. The type is slice of
    // DOCSSUBMGT3MIB_DocsSubmgt3GrpTable_DocsSubmgt3GrpEntry.
    DocsSubmgt3GrpEntry []*DOCSSUBMGT3MIB_DocsSubmgt3GrpTable_DocsSubmgt3GrpEntry
}

func (docsSubmgt3GrpTable *DOCSSUBMGT3MIB_DocsSubmgt3GrpTable) GetEntityData() *types.CommonEntityData {
    docsSubmgt3GrpTable.EntityData.YFilter = docsSubmgt3GrpTable.YFilter
    docsSubmgt3GrpTable.EntityData.YangName = "docsSubmgt3GrpTable"
    docsSubmgt3GrpTable.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3GrpTable.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    docsSubmgt3GrpTable.EntityData.SegmentPath = "docsSubmgt3GrpTable"
    docsSubmgt3GrpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3GrpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3GrpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3GrpTable.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3GrpTable.EntityData.Children.Append("docsSubmgt3GrpEntry", types.YChild{"DocsSubmgt3GrpEntry", nil})
    for i := range docsSubmgt3GrpTable.DocsSubmgt3GrpEntry {
        docsSubmgt3GrpTable.EntityData.Children.Append(types.GetSegmentPath(docsSubmgt3GrpTable.DocsSubmgt3GrpEntry[i]), types.YChild{"DocsSubmgt3GrpEntry", docsSubmgt3GrpTable.DocsSubmgt3GrpEntry[i]})
    }
    docsSubmgt3GrpTable.EntityData.Leafs = types.NewOrderedMap()

    docsSubmgt3GrpTable.EntityData.YListKeys = []string {}

    return &(docsSubmgt3GrpTable.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3GrpTable_DocsSubmgt3GrpEntry
// The conceptual row of docsSubmgt3GrpTable. 
// The CMTS does not persist the instances of the Grp 
// object across reinitializations.
type DOCSSUBMGT3MIB_DocsSubmgt3GrpTable_DocsSubmgt3GrpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry_DocsIf3CmtsCmRegStatusId
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute represents the filter group(s) associated  with the CM
    // signaled 'Upstream Drop Classifier Group ID'   encodings during the
    // registration process. UDC Group IDs are  integer values and this attribute
    // reports them as decimal   numbers that are space-separated. The zero-length
    // string indicates   that the CM didn't signal UDC Group IDs.   This
    // attribute provides two functions:   - Communicate the CM the configured UDC
    // Group ID(s), irrespective   of the CM being provisioned to filter upstream
    // traffic based   on IP Filters or UDCs.   - Optionally, and with regards to
    // the CMTS, if the value of the  attribute UdcSentInReqRsp is 'true',
    // indicates that the filtering   rules associated with the Subscriber
    // Management Group ID(s) will  be sent during registration to the CM. It is
    // vendor specific   whether the CMTS updates individual CM UDCs after
    // registration  when rules are changed in the Grp object. The type is string.
    DocsSubMgt3GrpUdcGroupIds interface{}

    // This attribute represents the CMTS upstream filtering   status for this CM.
    // The value 'true' indicates that the   CMTS has sent UDCs to the CM during
    // registration process.     In order for a CMTS to send UDCs to a CM, the
    // CMTS MAC Domain  needed to be enabled via the MAC Domain attribute  
    // SendUdcRulesEnabled and the CM had indicated the UDC capability   support
    // during the registration process. The value 'false'   indicates that the
    // CMTS was not enabled to sent UDCs to the   CMs in the MAC Domain, or the CM
    // does not advertised UDC   support in its capabilities encodings, or both.
    // Since the   CMTS capability to sent UDCs to CMs during the registration 
    // process is optional, the CMTS is not required to implement   the value
    // 'true'. The type is bool.
    DocsSubMgt3GrpUdcSentInRegRsp interface{}

    // This attribute represents the filter group applied  to traffic destined for
    // subscriber's CPE attached to the  referenced CM (attached to CM CPE
    // interfaces). This  value corresponds to the 'Subscriber Downstream  Group'
    // value of the 'Subscriber Management Filter Groups'  encoding signaled
    // during the CM registration  or in its absence, to the SubFilterDownDef
    // attribute  of the Base object. The value zero or a filter group  ID not
    // configured in the CMTS means no filtering is applied  to traffic destined
    // to hosts attached to this CM. The type is interface{} with range: 0..1024.
    DocsSubmgt3GrpSubFilterDs interface{}

    // This attribute represents the filter group applied  to traffic originating
    // from subscriber's CPE attached  to the referenced CM (attached to CM CPE
    // interfaces).  This value corresponds to the 'Subscriber Upstream  Group'
    // value of the 'Subscriber Management Filter  Groups' encoding signaled
    // during the CM registration  or in its absence, to the SubFilterUpDef
    // attribute  of the Base object. The value zero or a filter group  ID not
    // configured in the CMTS means no filtering  is applied to traffic
    // originating from hosts attached  to this CM. The type is interface{} with
    // range: 0..1024.
    DocsSubmgt3GrpSubFilterUs interface{}

    // This attribute represents the filter group applied  to traffic destined for
    // the CM itself. This value corresponds  to the 'CM Downstream Group' value
    // of the  'Subscriber Management Filter Groups' encoding signaled  during the
    // CM registration or in its absence,  to the CmFilterDownDef attribute of the
    // Base object.  The value zero or a filter group ID not configured in  the
    // CMTS means no filtering is applied to traffic destined  to the CM. The type
    // is interface{} with range: 0..1024.
    DocsSubmgt3GrpCmFilterDs interface{}

    // This attribute represents the filter group applied  to traffic originating
    // from the CM itself. This value  corresponds to the 'Subscriber Upstream
    // Group'  value of the 'Subscriber Management Filter Groups'  encoding
    // signaled during the CM registration or in its  absence, to the
    // SubFilterUpDef attribute of the Base  object. The value zero or a filter
    // group ID not configured  in the CMTS means no filtering is applied to
    // traffic  originating from this CM. The type is interface{} with range:
    // 0..1024.
    DocsSubmgt3GrpCmFilterUs interface{}

    // This attribute represents the filter group applied  to traffic destined to
    // the Embedded CableHome Portal  Services Element or the Embedded Router on
    // the referenced  CM. This value corresponds to the 'PS Downstream  Group'
    // value of the 'Subscriber Management Filter  Groups' encoding signaled
    // during the CM registration  or in its absence, to the SubFilterDownDef
    // attribute  of the Base object. The value zero or a filter  group ID not
    // configured in the CMTS means no filtering  is applied to traffic destined
    // to the Embedded CableHome  Portal Services Element or Embedded Router on 
    // this CM. The type is interface{} with range: 0..1024.
    DocsSubmgt3GrpPsFilterDs interface{}

    // This attribute represents the filter group applied  to traffic originating
    // from the Embedded CableHome  Portal Services Element or Embedded Router on
    // the  referenced CM. This value corresponds to the 'PS Upstream  Group'
    // value of the 'Subscriber Management Filter  Groups' encoding signaled
    // during the CM registration  or in its absence, to the SubFilterUpDef
    // attribute  of the Base object. The value zero or a filter group  ID not
    // configured in the CMTS means no filtering is  applied to traffic
    // originating from the Embedded CableHome  Portal Services Element or
    // Embedded Router  on this CM. The type is interface{} with range: 0..1024.
    DocsSubmgt3GrpPsFilterUs interface{}

    // This attribute represents the filter group applied  to traffic destined to
    // the Embedded Multimedia Terminal  Adapter on the referenced CM. This value
    // corresponds  to the 'MTA Downstream Group' value of the 'Subscriber 
    // Management Filter Groups' encoding signaled  during the CM registration or
    // in its absence, to  the SubFilterDownDef attribute of the Base object.  The
    // value zero or a filter group ID not configured in the  CMTS means no
    // filtering is applied to traffic destined  to the Embedded Multimedia
    // Terminal Adapter on  this CM. The type is interface{} with range: 0..1024.
    DocsSubmgt3GrpMtaFilterDs interface{}

    // This attribute represents the filter group applied  to traffic originating
    // from the Embedded Multimedia  Terminal Adapter on the referenced CM. This
    // value  corresponds to the 'MTA Upstream Group' value of the  'Subscriber
    // Management Filter Groups' encoding signaled  during the CM registration or
    // in its absence,  to the SubFilterUpDef attribute of the Base object.  The
    // value zero or a filter group ID not configured in  the CMTS means no
    // filtering is applied to traffic originating  from the Embedded Multimedia
    // Terminal Adapter  on this CM. The type is interface{} with range: 0..1024.
    DocsSubmgt3GrpMtaFilterUs interface{}

    // This attribute represents the filter group applied  to traffic destined for
    // the Embedded Set-Top Box on  the referenced CM. This value corresponds to
    // the 'STB  Downstream Group' value of the 'Subscriber Management  Filter
    // Groups' encoding signaled during the CM  registration or in its absence, to
    // the SubFilterDownDef  attribute of the Base object. The value zero or  a
    // filter group ID not configured in the CMTS means no filtering  is applied
    // to traffic destined to the Embedded  Set-Top Box on this CM. The type is
    // interface{} with range: 0..1024.
    DocsSubmgt3GrpStbFilterDs interface{}

    // This attribute represents the filter group applied  to traffic originating
    // from the Embedded Set-Top  Box on the referenced CM. This value corresponds
    // to the  'STB Upstream Group' value of the 'Subscriber Management  Filter
    // Groups' encoding signaled during the  CM registration or in its absence, to
    // the SubFilterUpDef  attribute of the Base object. The value zero or  a
    // filter group ID not configured in the CMTS means no filtering  is applied
    // to traffic originating from the  Embedded Set-Top Box on this CM. The type
    // is interface{} with range: 0..1024.
    DocsSubmgt3GrpStbFilterUs interface{}
}

func (docsSubmgt3GrpEntry *DOCSSUBMGT3MIB_DocsSubmgt3GrpTable_DocsSubmgt3GrpEntry) GetEntityData() *types.CommonEntityData {
    docsSubmgt3GrpEntry.EntityData.YFilter = docsSubmgt3GrpEntry.YFilter
    docsSubmgt3GrpEntry.EntityData.YangName = "docsSubmgt3GrpEntry"
    docsSubmgt3GrpEntry.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3GrpEntry.EntityData.ParentYangName = "docsSubmgt3GrpTable"
    docsSubmgt3GrpEntry.EntityData.SegmentPath = "docsSubmgt3GrpEntry" + types.AddKeyToken(docsSubmgt3GrpEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId")
    docsSubmgt3GrpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3GrpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3GrpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3GrpEntry.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3GrpEntry.EntityData.Leafs = types.NewOrderedMap()
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsSubmgt3GrpEntry.DocsIf3CmtsCmRegStatusId})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubMgt3GrpUdcGroupIds", types.YLeaf{"DocsSubMgt3GrpUdcGroupIds", docsSubmgt3GrpEntry.DocsSubMgt3GrpUdcGroupIds})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubMgt3GrpUdcSentInRegRsp", types.YLeaf{"DocsSubMgt3GrpUdcSentInRegRsp", docsSubmgt3GrpEntry.DocsSubMgt3GrpUdcSentInRegRsp})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpSubFilterDs", types.YLeaf{"DocsSubmgt3GrpSubFilterDs", docsSubmgt3GrpEntry.DocsSubmgt3GrpSubFilterDs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpSubFilterUs", types.YLeaf{"DocsSubmgt3GrpSubFilterUs", docsSubmgt3GrpEntry.DocsSubmgt3GrpSubFilterUs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpCmFilterDs", types.YLeaf{"DocsSubmgt3GrpCmFilterDs", docsSubmgt3GrpEntry.DocsSubmgt3GrpCmFilterDs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpCmFilterUs", types.YLeaf{"DocsSubmgt3GrpCmFilterUs", docsSubmgt3GrpEntry.DocsSubmgt3GrpCmFilterUs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpPsFilterDs", types.YLeaf{"DocsSubmgt3GrpPsFilterDs", docsSubmgt3GrpEntry.DocsSubmgt3GrpPsFilterDs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpPsFilterUs", types.YLeaf{"DocsSubmgt3GrpPsFilterUs", docsSubmgt3GrpEntry.DocsSubmgt3GrpPsFilterUs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpMtaFilterDs", types.YLeaf{"DocsSubmgt3GrpMtaFilterDs", docsSubmgt3GrpEntry.DocsSubmgt3GrpMtaFilterDs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpMtaFilterUs", types.YLeaf{"DocsSubmgt3GrpMtaFilterUs", docsSubmgt3GrpEntry.DocsSubmgt3GrpMtaFilterUs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpStbFilterDs", types.YLeaf{"DocsSubmgt3GrpStbFilterDs", docsSubmgt3GrpEntry.DocsSubmgt3GrpStbFilterDs})
    docsSubmgt3GrpEntry.EntityData.Leafs.Append("docsSubmgt3GrpStbFilterUs", types.YLeaf{"DocsSubmgt3GrpStbFilterUs", docsSubmgt3GrpEntry.DocsSubmgt3GrpStbFilterUs})

    docsSubmgt3GrpEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId"}

    return &(docsSubmgt3GrpEntry.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable
// This object describes a set of filter or classifier 
// criteria. Classifiers are assigned by group to the 
// individual CMs. That assignment is made via the 'Subscriber 
// Management TLVs' encodings sent upstream from 
// the CM to the CMTS during registration or in their 
// absence, default values configured in the CMTS. 
// A Filter Group ID (GrpId) is a set of rules that correspond 
// to the expansion of a UDC Group ID into UDC individual  
// classification rules.  The Filter Group Ids are generated  
// whenever the CMTS is configured to send UDCs during the CM  
// registration process. Implementation of L2 classification  
// criteria is optional for the CMTS; LLC/MAC upstream and  
// downstream filter criteria can be ignored during the packet 
// matching process.
type DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsSubmgt3FilterGrpTable.  The CMTS persists all
    // instances of the FilterGrp object  across reinitializations. The type is
    // slice of
    // DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry.
    DocsSubmgt3FilterGrpEntry []*DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry
}

func (docsSubmgt3FilterGrpTable *DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable) GetEntityData() *types.CommonEntityData {
    docsSubmgt3FilterGrpTable.EntityData.YFilter = docsSubmgt3FilterGrpTable.YFilter
    docsSubmgt3FilterGrpTable.EntityData.YangName = "docsSubmgt3FilterGrpTable"
    docsSubmgt3FilterGrpTable.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3FilterGrpTable.EntityData.ParentYangName = "DOCS-SUBMGT3-MIB"
    docsSubmgt3FilterGrpTable.EntityData.SegmentPath = "docsSubmgt3FilterGrpTable"
    docsSubmgt3FilterGrpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3FilterGrpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3FilterGrpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3FilterGrpTable.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3FilterGrpTable.EntityData.Children.Append("docsSubmgt3FilterGrpEntry", types.YChild{"DocsSubmgt3FilterGrpEntry", nil})
    for i := range docsSubmgt3FilterGrpTable.DocsSubmgt3FilterGrpEntry {
        docsSubmgt3FilterGrpTable.EntityData.Children.Append(types.GetSegmentPath(docsSubmgt3FilterGrpTable.DocsSubmgt3FilterGrpEntry[i]), types.YChild{"DocsSubmgt3FilterGrpEntry", docsSubmgt3FilterGrpTable.DocsSubmgt3FilterGrpEntry[i]})
    }
    docsSubmgt3FilterGrpTable.EntityData.Leafs = types.NewOrderedMap()

    docsSubmgt3FilterGrpTable.EntityData.YListKeys = []string {}

    return &(docsSubmgt3FilterGrpTable.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry
// The conceptual row of docsSubmgt3FilterGrpTable. 
// The CMTS persists all instances of the FilterGrp object 
// across reinitializations.
type DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This key is an identifier for a set of classifiers
    // known  as a filter group. Each CM may be associated with  several filter
    // groups for its upstream and downstream  traffic, one group per target end
    // point on the CM as  defined in the Grp object. Typically, many CMs share  a
    // common set of filter groups. The type is interface{} with range: 1..65535.
    DocsSubmgt3FilterGrpGrpId interface{}

    // This attribute is a key. This key represents an ordered classifier
    // identifier  within the group.  Filters are applied in order if  the
    // Priority attribute is not supported. The type is interface{} with range:
    // 1..65535.
    DocsSubmgt3FilterGrpRuleId interface{}

    // This attribute represents the action to take upon  this filter matching. 
    // 'permit' means to stop the classification  matching and accept the packet
    // for further  processing.  'deny' means to drop the packet. The type is
    // DocsSubmgt3FilterGrpAction.
    DocsSubmgt3FilterGrpAction interface{}

    // This attribute defines the order in which classifiers  are compared against
    // packets. The higher the value,  the higher the priority. The type is
    // interface{} with range: 0..65535.
    DocsSubmgt3FilterGrpPriority interface{}

    // This attribute represents the low value of a range  of ToS (Type of
    // Service) octet values. This object is  defined as an 8-bit octet as per the
    // DOCSIS Specification  for packet classification.  The IP ToS octet, as
    // originally defined in RFC 791, has  been superseded by the 6-bit
    // Differentiated Services  Field (DSField, RFC 3260) and the 2-bit Explicit 
    // Congestion Notification Field (ECN field, RFC 3168). The type is string
    // with length: 1.
    DocsSubmgt3FilterGrpIpTosLow interface{}

    // This attribute represents the high value of a range  of ToS octet values.
    // This object is defined as an 8-bit  octet as per the DOCSIS Specification
    // for packet classification.  The IP ToS octet, as originally defined in RFC
    // 791, has  been superseded by the 6-bit Differentiated Services  Field
    // (DSField, RFC 3260) and the 2-bit Explicit  Congestion Notification Field
    // (ECN field, RFC 3168). The type is string with length: 1.
    DocsSubmgt3FilterGrpIpTosHigh interface{}

    // This attribute represents the mask value that is bitwise  ANDed with ToS
    // octet in an IP packet, and the resulting value  is used for range checking
    // of IpTosLow and IpTosHigh. The type is string with length: 1.
    DocsSubmgt3FilterGrpIpTosMask interface{}

    // This attribute represents the value of the IP Protocol  field required for
    // IP packets to match this rule.  The value 256 matches traffic with any IP
    // Protocol value.  The value 257 by convention matches both TCP and  UDP. The
    // type is interface{} with range: 0..257.
    DocsSubmgt3FilterGrpIpProtocol interface{}

    // The type of the Internet address for InetSrcAddr,  InetSrcMask,
    // InetDestAddr, and InetDestMask. The type is InetAddressType.
    DocsSubmgt3FilterGrpInetAddrType interface{}

    // This attribute specifies the value of the IP Source Address  required for
    // packets to match this rule. An IP packet  matches the rule when the
    // packet's IP Source Address  bitwise ANDed with the InetSrcMask value equals
    // the InetSrcAddr value. The address type of this object  is specified by the
    // InetAddressType attribute. The type is string with length: 0..255.
    DocsSubmgt3FilterGrpInetSrcAddr interface{}

    // This attribute represents which bits of a packet's  IP Source Address are
    // compared to match this rule. An  IP packet matches the rule when the
    // packet's IP Source  Address bitwise ANDed with the InetSrcMask value equals
    // the InetSrcAddr value. The address type of this  object is specified by
    // InetAddrType. The type is string with length: 0..255.
    DocsSubmgt3FilterGrpInetSrcMask interface{}

    // This attribute specifies the value of the IP Destination  Address required
    // for packets to match this rule.  An IP packet matches the rule when the
    // packet's IP Destination  Address bitwise ANDed with the InetSrcMask  value
    // equals the InetDestAddr value. The address type  of this object is
    // specified by the InetAddrType attribute. The type is string with length:
    // 0..255.
    DocsSubmgt3FilterGrpInetDestAddr interface{}

    // This attribute represents which bits of a packet's  IP Destination Address
    // are compared to match this rule.  An IP packet matches the rule when the
    // packet's IP Destination  Address bitwise ANDed with the InetDestMask value 
    // equals the InetDestAddr value. The address type  of this object is
    // specified by InetAddrType. The type is string with length: 0..255.
    DocsSubmgt3FilterGrpInetDestMask interface{}

    // This attribute represents the low-end inclusive  range of TCP/UDP source
    // port numbers to which a packet  is compared. This attribute is irrelevant
    // for non-TCP/UDP  IP packets. The type is interface{} with range: 0..65535.
    DocsSubmgt3FilterGrpSrcPortStart interface{}

    // This attribute represents the high-end inclusive  range of TCP/UDP source
    // port numbers to which a packet  is compared. This attribute is irrelevant
    // for non-TCP/UDP  IP packets. The type is interface{} with range: 0..65535.
    DocsSubmgt3FilterGrpSrcPortEnd interface{}

    // This attribute represents the low-end inclusive  range of TCP/UDP
    // destination port numbers to which a  packet is compared. This attribute is
    // irrelevant for  non-TCP/UDP IP packets. The type is interface{} with range:
    // 0..65535.
    DocsSubmgt3FilterGrpDestPortStart interface{}

    // This attribute represents the high-end inclusive  range of TCP/UDP
    // destination port numbers to which a packet  is compared. This attribute is
    // irrelevant for non-TCP/UDP  IP packets. The type is interface{} with range:
    // 0..65535.
    DocsSubmgt3FilterGrpDestPortEnd interface{}

    // This attribute represents the criteria to match against  an Ethernet packet
    // MAC address bitwise ANDed  with DestMacMask. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsSubmgt3FilterGrpDestMacAddr interface{}

    // An Ethernet packet matches an entry when its  destination MAC address
    // bitwise ANDed with  the DestMacMask attribute equals the value of  the
    // DestMacAddr attribute. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsSubmgt3FilterGrpDestMacMask interface{}

    // This attribute represents the value to match against  an Ethernet packet
    // source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsSubmgt3FilterGrpSrcMacAddr interface{}

    // This attribute indicates the format of the layer 3  protocol ID in the
    // Ethernet packet. A value of 'none'  means that the rule does not use the
    // layer 3 protocol  type as a matching criteria. A value of 'ethertype' 
    // means that the rule applies only to frames that contain  an EtherType
    // value. EtherType values are contained  in packets using the DEC-Intel-Xerox
    // (DIX) encapsulation  or the RFC 1042 Sub-Network Access Protocol  (SNAP)
    // encapsulation formats. A value of 'dsap' means  that the rule applies only
    // to frames using the IEEE802.3  encapsulation format with a Destination
    // Service  Access Point (DSAP) other than 0xAA (which is reserved  for SNAP).
    // A value of 'mac' means that the rule  applies only to MAC management
    // messages for MAC management  messages. A value of 'all' means that the rule
    // matches all Ethernet packets. If the Ethernet frame  contains an 802.1P/Q
    // Tag header (i.e., EtherType  0x8100), this attribute applies to the
    // embedded EtherType  field within the 802.1p/Q header.  The value 'mac' is
    // only used for passing UDCs to CMs during   Registration. The CMTS ignores
    // filter rules that include   the value of this attribute set to 'mac' for
    // CMTS enforced   upstream and downstream subscriber management filter group 
    // rules. The type is DocsSubmgt3FilterGrpEnetProtocolType.
    DocsSubmgt3FilterGrpEnetProtocolType interface{}

    // This attribute represents the Ethernet protocol  type to be matched against
    // the packets. For EnetProtocolType  set to 'none', this attribute is ignored
    // when considering  whether a packet matches the current rule.  If the
    // attribute EnetProtocolType is 'ethertype',  this attribute gives the 16-bit
    // value of the EtherType  that the packet must match in order to match the
    // rule.  If the attribute EnetProtocolType is 'dsap', the lower  8 bits of
    // this attribute's value must match the DSAP  octet of the packet in order to
    // match the rule. If the Ethernet  frame contains an 802.1p/Q Tag header
    // (i.e.,  EtherType 0x8100), this attribute applies to the embedded 
    // EtherType field within the 802.1p/Q header. The type is interface{} with
    // range: 0..65535.
    DocsSubmgt3FilterGrpEnetProtocol interface{}

    // This attribute applies only to Ethernet frames using  the 802.1p/Q tag
    // header (indicated with EtherType  0x8100). Such frames include a 16-bit Tag
    // that contains  a 3-bit Priority field and a 12-bit VLAN number.  Tagged
    // Ethernet packets must have a 3-bit Priority  field within the range of
    // PriLow to PriHigh in order to  match this rule. The type is interface{}
    // with range: 0..7.
    DocsSubmgt3FilterGrpUserPriLow interface{}

    // This attribute applies only to Ethernet frames using  the 802.1p/Q tag
    // header (indicated with EtherType  0x8100). Such frames include a 16-bit Tag
    // that contains  a 3-bit Priority field and a 12-bit VLAN number.  Tagged
    // Ethernet packets must have a 3-bit Priority  field within the range of
    // PriLow to PriHigh in order to  match this rule. The type is interface{}
    // with range: 0..7.
    DocsSubmgt3FilterGrpUserPriHigh interface{}

    // This attribute applies only to Ethernet frames using  the 802.1p/Q tag
    // header. Tagged packets must have  a VLAN Identifier that matches the value
    // in order to  match the rule. The type is interface{} with range: 0..4094.
    DocsSubmgt3FilterGrpVlanId interface{}

    // This attribute counts the number of packets  that have been classified
    // (matched) using this rule  entry. This includes all packets delivered to a
    // Service  Flow maximum rate policing function, whether  or not that function
    // drops the packets. Discontinuities  in the value of this counter can occur
    // at re-initialization  of the managed system, and at other times  as
    // indicated by the value of ifCounterDiscontinuityTime  for the CM MAC Domain
    // interface. The type is interface{} with range: 0..18446744073709551615.
    DocsSubmgt3FilterGrpClassPkts interface{}

    // This attribute represents the Flow Label field in  the IPv6 header to be
    // matched by the classifier.  The value zero indicates that the Flow Label is
    // not specified  as part of the classifier and is not matched against
    // packets. The type is interface{} with range: 0..1048575.
    DocsSubmgt3FilterGrpFlowLabel interface{}

    // This attribute represents a bit-mask of the CM in-bound  interfaces to
    // which this classifier applies.  This attribute only applies to upstream
    // Drop Classifiers   being sent to CMs during the registration process. The
    // type is map[string]bool.
    DocsSubmgt3FilterGrpCmInterfaceMask interface{}

    // The conceptual row status of this object. The type is RowStatus.
    DocsSubmgt3FilterGrpRowStatus interface{}
}

func (docsSubmgt3FilterGrpEntry *DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry) GetEntityData() *types.CommonEntityData {
    docsSubmgt3FilterGrpEntry.EntityData.YFilter = docsSubmgt3FilterGrpEntry.YFilter
    docsSubmgt3FilterGrpEntry.EntityData.YangName = "docsSubmgt3FilterGrpEntry"
    docsSubmgt3FilterGrpEntry.EntityData.BundleName = "cisco_ios_xe"
    docsSubmgt3FilterGrpEntry.EntityData.ParentYangName = "docsSubmgt3FilterGrpTable"
    docsSubmgt3FilterGrpEntry.EntityData.SegmentPath = "docsSubmgt3FilterGrpEntry" + types.AddKeyToken(docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpGrpId, "docsSubmgt3FilterGrpGrpId") + types.AddKeyToken(docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpRuleId, "docsSubmgt3FilterGrpRuleId")
    docsSubmgt3FilterGrpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsSubmgt3FilterGrpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsSubmgt3FilterGrpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsSubmgt3FilterGrpEntry.EntityData.Children = types.NewOrderedMap()
    docsSubmgt3FilterGrpEntry.EntityData.Leafs = types.NewOrderedMap()
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpGrpId", types.YLeaf{"DocsSubmgt3FilterGrpGrpId", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpGrpId})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpRuleId", types.YLeaf{"DocsSubmgt3FilterGrpRuleId", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpRuleId})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpAction", types.YLeaf{"DocsSubmgt3FilterGrpAction", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpAction})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpPriority", types.YLeaf{"DocsSubmgt3FilterGrpPriority", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpPriority})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpIpTosLow", types.YLeaf{"DocsSubmgt3FilterGrpIpTosLow", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpIpTosLow})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpIpTosHigh", types.YLeaf{"DocsSubmgt3FilterGrpIpTosHigh", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpIpTosHigh})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpIpTosMask", types.YLeaf{"DocsSubmgt3FilterGrpIpTosMask", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpIpTosMask})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpIpProtocol", types.YLeaf{"DocsSubmgt3FilterGrpIpProtocol", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpIpProtocol})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpInetAddrType", types.YLeaf{"DocsSubmgt3FilterGrpInetAddrType", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpInetAddrType})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpInetSrcAddr", types.YLeaf{"DocsSubmgt3FilterGrpInetSrcAddr", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpInetSrcAddr})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpInetSrcMask", types.YLeaf{"DocsSubmgt3FilterGrpInetSrcMask", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpInetSrcMask})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpInetDestAddr", types.YLeaf{"DocsSubmgt3FilterGrpInetDestAddr", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpInetDestAddr})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpInetDestMask", types.YLeaf{"DocsSubmgt3FilterGrpInetDestMask", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpInetDestMask})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpSrcPortStart", types.YLeaf{"DocsSubmgt3FilterGrpSrcPortStart", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpSrcPortStart})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpSrcPortEnd", types.YLeaf{"DocsSubmgt3FilterGrpSrcPortEnd", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpSrcPortEnd})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpDestPortStart", types.YLeaf{"DocsSubmgt3FilterGrpDestPortStart", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpDestPortStart})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpDestPortEnd", types.YLeaf{"DocsSubmgt3FilterGrpDestPortEnd", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpDestPortEnd})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpDestMacAddr", types.YLeaf{"DocsSubmgt3FilterGrpDestMacAddr", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpDestMacAddr})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpDestMacMask", types.YLeaf{"DocsSubmgt3FilterGrpDestMacMask", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpDestMacMask})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpSrcMacAddr", types.YLeaf{"DocsSubmgt3FilterGrpSrcMacAddr", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpSrcMacAddr})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpEnetProtocolType", types.YLeaf{"DocsSubmgt3FilterGrpEnetProtocolType", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpEnetProtocolType})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpEnetProtocol", types.YLeaf{"DocsSubmgt3FilterGrpEnetProtocol", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpEnetProtocol})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpUserPriLow", types.YLeaf{"DocsSubmgt3FilterGrpUserPriLow", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpUserPriLow})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpUserPriHigh", types.YLeaf{"DocsSubmgt3FilterGrpUserPriHigh", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpUserPriHigh})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpVlanId", types.YLeaf{"DocsSubmgt3FilterGrpVlanId", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpVlanId})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpClassPkts", types.YLeaf{"DocsSubmgt3FilterGrpClassPkts", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpClassPkts})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpFlowLabel", types.YLeaf{"DocsSubmgt3FilterGrpFlowLabel", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpFlowLabel})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpCmInterfaceMask", types.YLeaf{"DocsSubmgt3FilterGrpCmInterfaceMask", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpCmInterfaceMask})
    docsSubmgt3FilterGrpEntry.EntityData.Leafs.Append("docsSubmgt3FilterGrpRowStatus", types.YLeaf{"DocsSubmgt3FilterGrpRowStatus", docsSubmgt3FilterGrpEntry.DocsSubmgt3FilterGrpRowStatus})

    docsSubmgt3FilterGrpEntry.EntityData.YListKeys = []string {"DocsSubmgt3FilterGrpGrpId", "DocsSubmgt3FilterGrpRuleId"}

    return &(docsSubmgt3FilterGrpEntry.EntityData)
}

// DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction represents processing.  'deny' means to drop the packet.
type DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction string

const (
    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction_permit DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction = "permit"

    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction_deny DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpAction = "deny"
)

// DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType represents rules.
type DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType string

const (
    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType_none DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType = "none"

    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType_ethertype DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType = "ethertype"

    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType_dsap DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType = "dsap"

    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType_mac DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType = "mac"

    DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType_all DOCSSUBMGT3MIB_DocsSubmgt3FilterGrpTable_DocsSubmgt3FilterGrpEntry_DocsSubmgt3FilterGrpEnetProtocolType = "all"
)

