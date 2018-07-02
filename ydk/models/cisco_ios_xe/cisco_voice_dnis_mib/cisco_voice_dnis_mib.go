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
    CvDnisMappingTable CISCOVOICEDNISMIB_CvDnisMappingTable

    // The table contains a DNIS name and a url. The url is a pointer to a VXML
    // page for the DNIS name. .
    CvDnisNodeTable CISCOVOICEDNISMIB_CvDnisNodeTable
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

    cISCOVOICEDNISMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVOICEDNISMIB.EntityData.Children.Append("cvDnisMappingTable", types.YChild{"CvDnisMappingTable", &cISCOVOICEDNISMIB.CvDnisMappingTable})
    cISCOVOICEDNISMIB.EntityData.Children.Append("cvDnisNodeTable", types.YChild{"CvDnisNodeTable", &cISCOVOICEDNISMIB.CvDnisNodeTable})
    cISCOVOICEDNISMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVOICEDNISMIB.EntityData.YListKeys = []string {}

    return &(cISCOVOICEDNISMIB.EntityData)
}

// CISCOVOICEDNISMIB_CvDnisMappingTable
// The table contains the map name and a url specifying
// a file name. The file contains DNIS entries that belong
// to the DNIS mapping.
type CISCOVOICEDNISMIB_CvDnisMappingTable struct {
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
    // slice of CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry.
    CvDnisMappingEntry []*CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry
}

func (cvDnisMappingTable *CISCOVOICEDNISMIB_CvDnisMappingTable) GetEntityData() *types.CommonEntityData {
    cvDnisMappingTable.EntityData.YFilter = cvDnisMappingTable.YFilter
    cvDnisMappingTable.EntityData.YangName = "cvDnisMappingTable"
    cvDnisMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cvDnisMappingTable.EntityData.ParentYangName = "CISCO-VOICE-DNIS-MIB"
    cvDnisMappingTable.EntityData.SegmentPath = "cvDnisMappingTable"
    cvDnisMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvDnisMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvDnisMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvDnisMappingTable.EntityData.Children = types.NewOrderedMap()
    cvDnisMappingTable.EntityData.Children.Append("cvDnisMappingEntry", types.YChild{"CvDnisMappingEntry", nil})
    for i := range cvDnisMappingTable.CvDnisMappingEntry {
        cvDnisMappingTable.EntityData.Children.Append(types.GetSegmentPath(cvDnisMappingTable.CvDnisMappingEntry[i]), types.YChild{"CvDnisMappingEntry", cvDnisMappingTable.CvDnisMappingEntry[i]})
    }
    cvDnisMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cvDnisMappingTable.EntityData.YListKeys = []string {}

    return &(cvDnisMappingTable.EntityData)
}

// CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry
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
type CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name which uniquely identifies a DNIS mapping.
    // . The type is string with length: 1..32.
    CvDnisMappingName interface{}

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
    CvDnisMappingUrl interface{}

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
    // new refresh. The type is CvDnisMappingRefresh.
    CvDnisMappingRefresh interface{}

    // ASCII text describing the error on last access of the url specified in
    // cvDnisMappingUrl.  If the url access does not succeed, then this object is
    // populated with an error message indicating the reason for failure. If the
    // url access succeeds, this object is set to null string. The type is string.
    CvDnisMappingUrlAccessError interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. When making the status active(1), if a valid
    // cvDnisMappingUrl is present the contents of the url is downloaded and
    // during that time cvDnisMappingRefresh is set to refresh(2). When
    // cvDnisMappingRefresh is set to refresh(2), only the destroy(6) operation is
    // allowed. The type is RowStatus.
    CvDnisMappingStatus interface{}
}

func (cvDnisMappingEntry *CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry) GetEntityData() *types.CommonEntityData {
    cvDnisMappingEntry.EntityData.YFilter = cvDnisMappingEntry.YFilter
    cvDnisMappingEntry.EntityData.YangName = "cvDnisMappingEntry"
    cvDnisMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cvDnisMappingEntry.EntityData.ParentYangName = "cvDnisMappingTable"
    cvDnisMappingEntry.EntityData.SegmentPath = "cvDnisMappingEntry" + types.AddKeyToken(cvDnisMappingEntry.CvDnisMappingName, "cvDnisMappingName")
    cvDnisMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvDnisMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvDnisMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvDnisMappingEntry.EntityData.Children = types.NewOrderedMap()
    cvDnisMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cvDnisMappingEntry.EntityData.Leafs.Append("cvDnisMappingName", types.YLeaf{"CvDnisMappingName", cvDnisMappingEntry.CvDnisMappingName})
    cvDnisMappingEntry.EntityData.Leafs.Append("cvDnisMappingUrl", types.YLeaf{"CvDnisMappingUrl", cvDnisMappingEntry.CvDnisMappingUrl})
    cvDnisMappingEntry.EntityData.Leafs.Append("cvDnisMappingRefresh", types.YLeaf{"CvDnisMappingRefresh", cvDnisMappingEntry.CvDnisMappingRefresh})
    cvDnisMappingEntry.EntityData.Leafs.Append("cvDnisMappingUrlAccessError", types.YLeaf{"CvDnisMappingUrlAccessError", cvDnisMappingEntry.CvDnisMappingUrlAccessError})
    cvDnisMappingEntry.EntityData.Leafs.Append("cvDnisMappingStatus", types.YLeaf{"CvDnisMappingStatus", cvDnisMappingEntry.CvDnisMappingStatus})

    cvDnisMappingEntry.EntityData.YListKeys = []string {"CvDnisMappingName"}

    return &(cvDnisMappingEntry.EntityData)
}

// CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh represents              becomes idle to issue new refresh.
type CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh string

const (
    CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh_idle CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh = "idle"

    CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh_refresh CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingRefresh = "refresh"
)

// CISCOVOICEDNISMIB_CvDnisNodeTable
// The table contains a DNIS name and a url. The url is a
// pointer to a VXML page for the DNIS name. 
type CISCOVOICEDNISMIB_CvDnisNodeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry is a DNIS name and the location of the associated VXML page. New
    // DNIS entries can be created or the existing entries can be modified or
    // deleted only if the corresponding map name (defined in cvDnisMappingTable)
    // does not have any file name provided in the cvDnisMappingUrl object.   If a
    // file name is provided in cvDnisMappingUrl corresponding to this entry's map
    // name, then this row will have read permission only. The type is slice of
    // CISCOVOICEDNISMIB_CvDnisNodeTable_CvDnisNodeEntry.
    CvDnisNodeEntry []*CISCOVOICEDNISMIB_CvDnisNodeTable_CvDnisNodeEntry
}

func (cvDnisNodeTable *CISCOVOICEDNISMIB_CvDnisNodeTable) GetEntityData() *types.CommonEntityData {
    cvDnisNodeTable.EntityData.YFilter = cvDnisNodeTable.YFilter
    cvDnisNodeTable.EntityData.YangName = "cvDnisNodeTable"
    cvDnisNodeTable.EntityData.BundleName = "cisco_ios_xe"
    cvDnisNodeTable.EntityData.ParentYangName = "CISCO-VOICE-DNIS-MIB"
    cvDnisNodeTable.EntityData.SegmentPath = "cvDnisNodeTable"
    cvDnisNodeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvDnisNodeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvDnisNodeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvDnisNodeTable.EntityData.Children = types.NewOrderedMap()
    cvDnisNodeTable.EntityData.Children.Append("cvDnisNodeEntry", types.YChild{"CvDnisNodeEntry", nil})
    for i := range cvDnisNodeTable.CvDnisNodeEntry {
        cvDnisNodeTable.EntityData.Children.Append(types.GetSegmentPath(cvDnisNodeTable.CvDnisNodeEntry[i]), types.YChild{"CvDnisNodeEntry", cvDnisNodeTable.CvDnisNodeEntry[i]})
    }
    cvDnisNodeTable.EntityData.Leafs = types.NewOrderedMap()

    cvDnisNodeTable.EntityData.YListKeys = []string {}

    return &(cvDnisNodeTable.EntityData)
}

// CISCOVOICEDNISMIB_CvDnisNodeTable_CvDnisNodeEntry
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
type CISCOVOICEDNISMIB_CvDnisNodeTable_CvDnisNodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // cisco_voice_dnis_mib.CISCOVOICEDNISMIB_CvDnisMappingTable_CvDnisMappingEntry_CvDnisMappingName
    CvDnisMappingName interface{}

    // This attribute is a key. The individual DNIS name. It is unique within a
    // DNIS mapping. The type is string.
    CvDnisNumber interface{}

    // The url specifies a VXML page. This page contains voice XML links to play
    // audio data.  This url which is a VXML page is not read immediately when the
    // row is made active(1), but only when a call that requires the use of this
    // DNIS comes through. The type is string.
    CvDnisNodeUrl interface{}

    // This object specifies whether the object in a particular row is modifiable.
    // The object is set to true(1) if the corresponding map name (defined in
    // cvDnisMappingTable) does not have any file name provided in the
    // cvDnisMappingUrl object. Otherwise this object is set to false(2) and the
    // row becomes read only.  . The type is bool.
    CvDnisNodeModifiable interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The objects in a row can be modified or deleted while the
    // row status is active(1) and cvDnisNodeModifiable is true(1). The row status
    // cannot be set to notInService(2) or createAndWait(5). . The type is
    // RowStatus.
    CvDnisNodeStatus interface{}
}

func (cvDnisNodeEntry *CISCOVOICEDNISMIB_CvDnisNodeTable_CvDnisNodeEntry) GetEntityData() *types.CommonEntityData {
    cvDnisNodeEntry.EntityData.YFilter = cvDnisNodeEntry.YFilter
    cvDnisNodeEntry.EntityData.YangName = "cvDnisNodeEntry"
    cvDnisNodeEntry.EntityData.BundleName = "cisco_ios_xe"
    cvDnisNodeEntry.EntityData.ParentYangName = "cvDnisNodeTable"
    cvDnisNodeEntry.EntityData.SegmentPath = "cvDnisNodeEntry" + types.AddKeyToken(cvDnisNodeEntry.CvDnisMappingName, "cvDnisMappingName") + types.AddKeyToken(cvDnisNodeEntry.CvDnisNumber, "cvDnisNumber")
    cvDnisNodeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvDnisNodeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvDnisNodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvDnisNodeEntry.EntityData.Children = types.NewOrderedMap()
    cvDnisNodeEntry.EntityData.Leafs = types.NewOrderedMap()
    cvDnisNodeEntry.EntityData.Leafs.Append("cvDnisMappingName", types.YLeaf{"CvDnisMappingName", cvDnisNodeEntry.CvDnisMappingName})
    cvDnisNodeEntry.EntityData.Leafs.Append("cvDnisNumber", types.YLeaf{"CvDnisNumber", cvDnisNodeEntry.CvDnisNumber})
    cvDnisNodeEntry.EntityData.Leafs.Append("cvDnisNodeUrl", types.YLeaf{"CvDnisNodeUrl", cvDnisNodeEntry.CvDnisNodeUrl})
    cvDnisNodeEntry.EntityData.Leafs.Append("cvDnisNodeModifiable", types.YLeaf{"CvDnisNodeModifiable", cvDnisNodeEntry.CvDnisNodeModifiable})
    cvDnisNodeEntry.EntityData.Leafs.Append("cvDnisNodeStatus", types.YLeaf{"CvDnisNodeStatus", cvDnisNodeEntry.CvDnisNodeStatus})

    cvDnisNodeEntry.EntityData.YListKeys = []string {"CvDnisMappingName", "CvDnisNumber"}

    return &(cvDnisNodeEntry.EntityData)
}

