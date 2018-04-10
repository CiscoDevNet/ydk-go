// The MIB module provides management support for Dialer
// Number Information Service (DNIS) mapping.  A DNIS
// entry is associated with a Voice XML (VXML) page to
// provide audio play back features. Multiple DNIS
// entries can be grouped together to form a DNIS
// mapping with a unique map name.
// 
// *** ABBREVIATIONS, ACRONYMS, AND SYMBOLS ***
// 
// DNIS - Dialer Number Information Service
// 
// XML  - Extensible Markup Language
// 
// VXML - Voice XML
// 
// URL  - Uniform Resource Locator  
package cisco_voice_dnis_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_voice_dnis_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VOICE-DNIS-MIB CISCO-VOICE-DNIS-MIB}", reflect.TypeOf(CISCOVOICEDNISMIB{}))
    ydk.RegisterEntity("CISCO-VOICE-DNIS-MIB:CISCO-VOICE-DNIS-MIB", reflect.TypeOf(CISCOVOICEDNISMIB{}))
}

// CISCOVOICEDNISMIB
type CISCOVOICEDNISMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The table contains the map name and a url specifying a file name. The file
    // contains DNIS entries that belong to the DNIS mapping.
    Cvdnismappingtable CISCOVOICEDNISMIB_Cvdnismappingtable

    // The table contains a DNIS name and a url. The url is a pointer to a VXML
    // page for the DNIS name. .
    Cvdnisnodetable CISCOVOICEDNISMIB_Cvdnisnodetable
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetEntityData() *types.CommonEntityData {
    cISCOVOICEDNISMIB.EntityData.YFilter = cISCOVOICEDNISMIB.YFilter
    cISCOVOICEDNISMIB.EntityData.YangName = "CISCO-VOICE-DNIS-MIB"
    cISCOVOICEDNISMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVOICEDNISMIB.EntityData.ParentYangName = "CISCO-VOICE-DNIS-MIB"
    cISCOVOICEDNISMIB.EntityData.SegmentPath = "CISCO-VOICE-DNIS-MIB:CISCO-VOICE-DNIS-MIB"
    cISCOVOICEDNISMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVOICEDNISMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVOICEDNISMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVOICEDNISMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVOICEDNISMIB.EntityData.Children["cvDnisMappingTable"] = types.YChild{"Cvdnismappingtable", &cISCOVOICEDNISMIB.Cvdnismappingtable}
    cISCOVOICEDNISMIB.EntityData.Children["cvDnisNodeTable"] = types.YChild{"Cvdnisnodetable", &cISCOVOICEDNISMIB.Cvdnisnodetable}
    cISCOVOICEDNISMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVOICEDNISMIB.EntityData)
}

// CISCOVOICEDNISMIB_Cvdnismappingtable
// The table contains the map name and a url specifying
// a file name. The file contains DNIS entries that belong
// to the DNIS mapping.
type CISCOVOICEDNISMIB_Cvdnismappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single DNIS mapping. There is a unique DNIS map name.
    // New DNIS mapping can be created using cvDnisMappingStatus.  The entry can
    // be created with or without a file location  specified by cvDnisMappingUrl.
    // The mapping file contains DNIS name and VXML page per line. For example, a 
    // cvDnisMappingUrl could be tftp://someserver/dnismap.txt. This file is a
    // text file and the content format is   dnis <dnisname> url <urlname>. An
    // example of the contents of the file itself can be   dnis 18004251234 url
    // http://www.b.com/p/vwelcome.vxml   dnis 18004253421 url
    // http://www.c.com/j/vxmlintf.vxml If a mapping file location is specified,
    // then new rows corresponding to this map name are created and populated in
    // cvDnisNodeTable from the contents of the file. The rows corresponding to
    // this map name in cvDnisNodeTable cannot be created or modified or deleted
    // but can be read.   If a mapping file location is not specified in
    // cvDnisMappingUrl, then individual DNIS entries corresponding to this map
    // name can be created, modified and deleted in cvDnisNodeTable.   Deleting an
    // entry deletes all the related entries in cvDnisNodeTable. . The type is
    // slice of CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry.
    Cvdnismappingentry []CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetEntityData() *types.CommonEntityData {
    cvdnismappingtable.EntityData.YFilter = cvdnismappingtable.YFilter
    cvdnismappingtable.EntityData.YangName = "cvDnisMappingTable"
    cvdnismappingtable.EntityData.BundleName = "cisco_ios_xe"
    cvdnismappingtable.EntityData.ParentYangName = "CISCO-VOICE-DNIS-MIB"
    cvdnismappingtable.EntityData.SegmentPath = "cvDnisMappingTable"
    cvdnismappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvdnismappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvdnismappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvdnismappingtable.EntityData.Children = make(map[string]types.YChild)
    cvdnismappingtable.EntityData.Children["cvDnisMappingEntry"] = types.YChild{"Cvdnismappingentry", nil}
    for i := range cvdnismappingtable.Cvdnismappingentry {
        cvdnismappingtable.EntityData.Children[types.GetSegmentPath(&cvdnismappingtable.Cvdnismappingentry[i])] = types.YChild{"Cvdnismappingentry", &cvdnismappingtable.Cvdnismappingentry[i]}
    }
    cvdnismappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvdnismappingtable.EntityData)
}

// CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry
// Information about a single DNIS mapping. There is a
// unique DNIS map name. New DNIS mapping can be created
// using cvDnisMappingStatus.
// 
// The entry can be created with or without a file location 
// specified by cvDnisMappingUrl. The mapping file contains
// DNIS name and VXML page per line. For example, a 
// cvDnisMappingUrl could be tftp://someserver/dnismap.txt.
// This file is a text file and the content format is
//   dnis <dnisname> url <urlname>.
// An example of the contents of the file itself can be
//   dnis 18004251234 url http://www.b.com/p/vwelcome.vxml
//   dnis 18004253421 url http://www.c.com/j/vxmlintf.vxml
// If a mapping file location is specified, then new rows
// corresponding to this map name are created and populated
// in cvDnisNodeTable from the contents of the file. The
// rows corresponding to this map name in cvDnisNodeTable
// cannot be created or modified or deleted but can be
// read. 
// 
// If a mapping file location is not specified in
// cvDnisMappingUrl, then individual DNIS entries
// corresponding to this map name can be created, modified
// and deleted in cvDnisNodeTable. 
// 
// Deleting an entry deletes all the related entries in
// cvDnisNodeTable. 
type CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name which uniquely identifies a DNIS mapping.
    // . The type is string with length: 1..32.
    Cvdnismappingname interface{}

    // The url specifies a file location. The file contains individual DNIS
    // entries that belong to the DNIS map  name specified by cvDnisMappingName. 
    // Once a url is created and associated with a map name (the association is
    // complete when the row is made active(1)), it cannot be modified while
    // cvDnisMappingStatus is active. If a different url needs to be associated
    // with the current map name, the row status should be made notInService(2)
    // and this object has to be modified to associate a new url. When a new
    // association is made all the DNIS entries corresponding to the old
    // association will be deleted from the cvDnisNodeTable.  The url is read when
    // the row status is made active(1) or when the row status is active and the
    // object   cvDnisMappingRefresh is explicitly set to refresh(2).   If the url
    // is not accessible then a cvDnisMappingUrlInaccessible notification will be
    // generted. . The type is string.
    Cvdnismappingurl interface{}

    // Whenever there is a need to re-read the contents of the file specified by
    // cvDnisMappingUrl, this object can be set to refresh(2). This will cause the
    // contents of the file to be re-read and correspondingly update the
    // cvDnisNodeTable. After the completion of this operation, the value of this
    // object is reset to idle(1). The only operation allowed on this object is
    // setting it to refresh(2). This can only be done when the current value is
    // idle(1) and the rowstatus is active(1).  idle       - The refreshing
    // process is idle and the user              can modify this object to
    // refresh. refresh    - The refreshing process is currently busy and         
    // the user have to wait till the object              becomes idle to issue
    // new refresh. The type is Cvdnismappingrefresh.
    Cvdnismappingrefresh interface{}

    // ASCII text describing the error on last access of the url specified in
    // cvDnisMappingUrl.  If the url access does not succeed, then this object is
    // populated with an error message indicating the reason for failure. If the
    // url access succeeds, this object is set to null string. The type is string.
    Cvdnismappingurlaccesserror interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. When making the status active(1), if a valid
    // cvDnisMappingUrl is present the contents of the url is downloaded and
    // during that time cvDnisMappingRefresh is set to refresh(2). When
    // cvDnisMappingRefresh is set to refresh(2), only the destroy(6) operation is
    // allowed. The type is RowStatus.
    Cvdnismappingstatus interface{}
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetEntityData() *types.CommonEntityData {
    cvdnismappingentry.EntityData.YFilter = cvdnismappingentry.YFilter
    cvdnismappingentry.EntityData.YangName = "cvDnisMappingEntry"
    cvdnismappingentry.EntityData.BundleName = "cisco_ios_xe"
    cvdnismappingentry.EntityData.ParentYangName = "cvDnisMappingTable"
    cvdnismappingentry.EntityData.SegmentPath = "cvDnisMappingEntry" + "[cvDnisMappingName='" + fmt.Sprintf("%v", cvdnismappingentry.Cvdnismappingname) + "']"
    cvdnismappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvdnismappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvdnismappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvdnismappingentry.EntityData.Children = make(map[string]types.YChild)
    cvdnismappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvdnismappingentry.EntityData.Leafs["cvDnisMappingName"] = types.YLeaf{"Cvdnismappingname", cvdnismappingentry.Cvdnismappingname}
    cvdnismappingentry.EntityData.Leafs["cvDnisMappingUrl"] = types.YLeaf{"Cvdnismappingurl", cvdnismappingentry.Cvdnismappingurl}
    cvdnismappingentry.EntityData.Leafs["cvDnisMappingRefresh"] = types.YLeaf{"Cvdnismappingrefresh", cvdnismappingentry.Cvdnismappingrefresh}
    cvdnismappingentry.EntityData.Leafs["cvDnisMappingUrlAccessError"] = types.YLeaf{"Cvdnismappingurlaccesserror", cvdnismappingentry.Cvdnismappingurlaccesserror}
    cvdnismappingentry.EntityData.Leafs["cvDnisMappingStatus"] = types.YLeaf{"Cvdnismappingstatus", cvdnismappingentry.Cvdnismappingstatus}
    return &(cvdnismappingentry.EntityData)
}

// CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh represents              becomes idle to issue new refresh.
type CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh string

const (
    CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh_idle CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh = "idle"

    CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh_refresh CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh = "refresh"
)

// CISCOVOICEDNISMIB_Cvdnisnodetable
// The table contains a DNIS name and a url. The url is a
// pointer to a VXML page for the DNIS name. 
type CISCOVOICEDNISMIB_Cvdnisnodetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry is a DNIS name and the location of the associated VXML page. New
    // DNIS entries can be created or the existing entries can be modified or
    // deleted only if the corresponding map name (defined in cvDnisMappingTable)
    // does not have any file name provided in the cvDnisMappingUrl object.   If a
    // file name is provided in cvDnisMappingUrl corresponding to this entry's map
    // name, then this row will have read permission only. The type is slice of
    // CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry.
    Cvdnisnodeentry []CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetEntityData() *types.CommonEntityData {
    cvdnisnodetable.EntityData.YFilter = cvdnisnodetable.YFilter
    cvdnisnodetable.EntityData.YangName = "cvDnisNodeTable"
    cvdnisnodetable.EntityData.BundleName = "cisco_ios_xe"
    cvdnisnodetable.EntityData.ParentYangName = "CISCO-VOICE-DNIS-MIB"
    cvdnisnodetable.EntityData.SegmentPath = "cvDnisNodeTable"
    cvdnisnodetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvdnisnodetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvdnisnodetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvdnisnodetable.EntityData.Children = make(map[string]types.YChild)
    cvdnisnodetable.EntityData.Children["cvDnisNodeEntry"] = types.YChild{"Cvdnisnodeentry", nil}
    for i := range cvdnisnodetable.Cvdnisnodeentry {
        cvdnisnodetable.EntityData.Children[types.GetSegmentPath(&cvdnisnodetable.Cvdnisnodeentry[i])] = types.YChild{"Cvdnisnodeentry", &cvdnisnodetable.Cvdnisnodeentry[i]}
    }
    cvdnisnodetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvdnisnodetable.EntityData)
}

// CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry
// Each entry is a DNIS name and the location of the
// associated VXML page. New DNIS entries can be created or
// the existing entries can be modified or deleted only if
// the corresponding map name (defined in
// cvDnisMappingTable) does not have any file name provided
// in the cvDnisMappingUrl object. 
// 
// If a file name is provided in cvDnisMappingUrl
// corresponding to this entry's map name, then this row
// will have read permission only.
type CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // cisco_voice_dnis_mib.CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingname
    Cvdnismappingname interface{}

    // This attribute is a key. The individual DNIS name. It is unique within a
    // DNIS mapping. The type is string.
    Cvdnisnumber interface{}

    // The url specifies a VXML page. This page contains voice XML links to play
    // audio data.  This url which is a VXML page is not read immediately when the
    // row is made active(1), but only when a call that requires the use of this
    // DNIS comes through. The type is string.
    Cvdnisnodeurl interface{}

    // This object specifies whether the object in a particular row is modifiable.
    // The object is set to true(1) if the corresponding map name (defined in
    // cvDnisMappingTable) does not have any file name provided in the
    // cvDnisMappingUrl object. Otherwise this object is set to false(2) and the
    // row becomes read only.  . The type is bool.
    Cvdnisnodemodifiable interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The objects in a row can be modified or deleted while the
    // row status is active(1) and cvDnisNodeModifiable is true(1). The row status
    // cannot be set to notInService(2) or createAndWait(5). . The type is
    // RowStatus.
    Cvdnisnodestatus interface{}
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetEntityData() *types.CommonEntityData {
    cvdnisnodeentry.EntityData.YFilter = cvdnisnodeentry.YFilter
    cvdnisnodeentry.EntityData.YangName = "cvDnisNodeEntry"
    cvdnisnodeentry.EntityData.BundleName = "cisco_ios_xe"
    cvdnisnodeentry.EntityData.ParentYangName = "cvDnisNodeTable"
    cvdnisnodeentry.EntityData.SegmentPath = "cvDnisNodeEntry" + "[cvDnisMappingName='" + fmt.Sprintf("%v", cvdnisnodeentry.Cvdnismappingname) + "']" + "[cvDnisNumber='" + fmt.Sprintf("%v", cvdnisnodeentry.Cvdnisnumber) + "']"
    cvdnisnodeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvdnisnodeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvdnisnodeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvdnisnodeentry.EntityData.Children = make(map[string]types.YChild)
    cvdnisnodeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvdnisnodeentry.EntityData.Leafs["cvDnisMappingName"] = types.YLeaf{"Cvdnismappingname", cvdnisnodeentry.Cvdnismappingname}
    cvdnisnodeentry.EntityData.Leafs["cvDnisNumber"] = types.YLeaf{"Cvdnisnumber", cvdnisnodeentry.Cvdnisnumber}
    cvdnisnodeentry.EntityData.Leafs["cvDnisNodeUrl"] = types.YLeaf{"Cvdnisnodeurl", cvdnisnodeentry.Cvdnisnodeurl}
    cvdnisnodeentry.EntityData.Leafs["cvDnisNodeModifiable"] = types.YLeaf{"Cvdnisnodemodifiable", cvdnisnodeentry.Cvdnisnodemodifiable}
    cvdnisnodeentry.EntityData.Leafs["cvDnisNodeStatus"] = types.YLeaf{"Cvdnisnodestatus", cvdnisnodeentry.Cvdnisnodestatus}
    return &(cvdnisnodeentry.EntityData)
}

