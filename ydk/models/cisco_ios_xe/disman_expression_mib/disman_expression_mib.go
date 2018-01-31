// The MIB module for defining expressions of MIB objects for
// management purposes.
package disman_expression_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package disman_expression_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DISMAN-EXPRESSION-MIB DISMAN-EXPRESSION-MIB}", reflect.TypeOf(DISMANEXPRESSIONMIB{}))
    ydk.RegisterEntity("DISMAN-EXPRESSION-MIB:DISMAN-EXPRESSION-MIB", reflect.TypeOf(DISMANEXPRESSIONMIB{}))
}

// DISMANEXPRESSIONMIB
type DISMANEXPRESSIONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Expresource DISMANEXPRESSIONMIB_Expresource

    // A table of expression definitions.
    Expexpressiontable DISMANEXPRESSIONMIB_Expexpressiontable

    // A table of expression errors.
    Experrortable DISMANEXPRESSIONMIB_Experrortable

    // A table of object definitions for each expExpression.  Wildcarding instance
    // IDs:  It is legal to omit all or part of the instance portion for some or
    // all of the objects in an expression. (See the DESCRIPTION of expObjectID
    // for details.  However, note that if more than one object in the same
    // expression is wildcarded in this way, they all must be objects where that
    // portion of the instance is the same.  In other words, all objects may be in
    // the same SEQUENCE or in different SEQUENCEs but with the same semantic
    // index value (e.g., a value of ifIndex) for the wildcarded portion.
    Expobjecttable DISMANEXPRESSIONMIB_Expobjecttable

    // A table of values from evaluated expressions.
    Expvaluetable DISMANEXPRESSIONMIB_Expvaluetable
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetFilter() yfilter.YFilter { return dISMANEXPRESSIONMIB.YFilter }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) SetFilter(yf yfilter.YFilter) { dISMANEXPRESSIONMIB.YFilter = yf }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetGoName(yname string) string {
    if yname == "expResource" { return "Expresource" }
    if yname == "expExpressionTable" { return "Expexpressiontable" }
    if yname == "expErrorTable" { return "Experrortable" }
    if yname == "expObjectTable" { return "Expobjecttable" }
    if yname == "expValueTable" { return "Expvaluetable" }
    return ""
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetSegmentPath() string {
    return "DISMAN-EXPRESSION-MIB:DISMAN-EXPRESSION-MIB"
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expResource" {
        return &dISMANEXPRESSIONMIB.Expresource
    }
    if childYangName == "expExpressionTable" {
        return &dISMANEXPRESSIONMIB.Expexpressiontable
    }
    if childYangName == "expErrorTable" {
        return &dISMANEXPRESSIONMIB.Experrortable
    }
    if childYangName == "expObjectTable" {
        return &dISMANEXPRESSIONMIB.Expobjecttable
    }
    if childYangName == "expValueTable" {
        return &dISMANEXPRESSIONMIB.Expvaluetable
    }
    return nil
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["expResource"] = &dISMANEXPRESSIONMIB.Expresource
    children["expExpressionTable"] = &dISMANEXPRESSIONMIB.Expexpressiontable
    children["expErrorTable"] = &dISMANEXPRESSIONMIB.Experrortable
    children["expObjectTable"] = &dISMANEXPRESSIONMIB.Expobjecttable
    children["expValueTable"] = &dISMANEXPRESSIONMIB.Expvaluetable
    return children
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetYangName() string { return "DISMAN-EXPRESSION-MIB" }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) SetParent(parent types.Entity) { dISMANEXPRESSIONMIB.parent = parent }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetParent() types.Entity { return dISMANEXPRESSIONMIB.parent }

func (dISMANEXPRESSIONMIB *DISMANEXPRESSIONMIB) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Expresource
type DISMANEXPRESSIONMIB_Expresource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The minimum expExpressionDeltaInterval this system will accept.  A system
    // may use the larger values of this minimum to lessen the impact of
    // constantly computing deltas.  For larger delta sampling intervals the
    // system samples less often and suffers less overhead.  This object provides
    // a way to enforce such lower overhead for all expressions created after it
    // is set.  The value -1 indicates that expResourceDeltaMinimum is irrelevant
    // as the system will not accept 'deltaValue' as a value for
    // expObjectSampleType.  Unless explicitly resource limited, a system's value
    // for this object should be 1, allowing as small as a 1 second interval for
    // ongoing delta sampling.  Changing this value will not invalidate an
    // existing setting of expObjectSampleType. The type is interface{} with
    // range: -1..None | 1..600. Units are seconds.
    Expresourcedeltaminimum interface{}

    // For every instance of a deltaValue object, one dynamic instance entry is
    // needed for holding the instance value from the previous sample, i.e. to
    // maintain state.  This object limits maximum number of dynamic instance
    // entries this system will support for wildcarded delta objects in
    // expressions. For a given delta expression, the number of dynamic instances
    // is the number of values that meet all criteria to exist times the number of
    // delta values in the expression.  A value of 0 indicates no preset limit,
    // that is, the limit is dynamic based on system operation and resources. 
    // Unless explicitly resource limited, a system's value for this object should
    // be 0.  Changing this value will not eliminate or inhibit existing delta
    // wildcard instance objects but will prevent the creation of more such
    // objects.  An attempt to allocate beyond the limit results in expErrorCode
    // being tooManyWildcardValues for that evaluation attempt. The type is
    // interface{} with range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstancemaximum interface{}

    // The number of currently active instance entries as defined for
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstances interface{}

    // The highest value of expResourceDeltaWildcardInstances that has occurred
    // since initialization of the managed system. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstanceshigh interface{}

    // The number of times this system could not evaluate an expression because
    // that would have created a value instance in excess of
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstanceresourcelacks interface{}
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetFilter() yfilter.YFilter { return expresource.YFilter }

func (expresource *DISMANEXPRESSIONMIB_Expresource) SetFilter(yf yfilter.YFilter) { expresource.YFilter = yf }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetGoName(yname string) string {
    if yname == "expResourceDeltaMinimum" { return "Expresourcedeltaminimum" }
    if yname == "expResourceDeltaWildcardInstanceMaximum" { return "Expresourcedeltawildcardinstancemaximum" }
    if yname == "expResourceDeltaWildcardInstances" { return "Expresourcedeltawildcardinstances" }
    if yname == "expResourceDeltaWildcardInstancesHigh" { return "Expresourcedeltawildcardinstanceshigh" }
    if yname == "expResourceDeltaWildcardInstanceResourceLacks" { return "Expresourcedeltawildcardinstanceresourcelacks" }
    return ""
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetSegmentPath() string {
    return "expResource"
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expResourceDeltaMinimum"] = expresource.Expresourcedeltaminimum
    leafs["expResourceDeltaWildcardInstanceMaximum"] = expresource.Expresourcedeltawildcardinstancemaximum
    leafs["expResourceDeltaWildcardInstances"] = expresource.Expresourcedeltawildcardinstances
    leafs["expResourceDeltaWildcardInstancesHigh"] = expresource.Expresourcedeltawildcardinstanceshigh
    leafs["expResourceDeltaWildcardInstanceResourceLacks"] = expresource.Expresourcedeltawildcardinstanceresourcelacks
    return leafs
}

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetBundleName() string { return "cisco_ios_xe" }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetYangName() string { return "expResource" }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expresource *DISMANEXPRESSIONMIB_Expresource) SetParent(parent types.Entity) { expresource.parent = parent }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetParent() types.Entity { return expresource.parent }

func (expresource *DISMANEXPRESSIONMIB_Expresource) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Expexpressiontable
// A table of expression definitions.
type DISMANEXPRESSIONMIB_Expexpressiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single expression.  New expressions can be created
    // using expExpressionRowStatus.  To create an expression first create the
    // named entry in this table.  Then use expExpressionName to populate
    // expObjectTable. For expression evaluation to succeed all related entries in
    // expExpressionTable and expObjectTable must be 'active'.  If these
    // conditions are not met the corresponding values in expValue simply are not
    // instantiated.  Deleting an entry deletes all related entries in
    // expObjectTable and expErrorTable.  Because of the relationships among the
    // multiple tables for an expression (expExpressionTable, expObjectTable, and
    // expValueTable) and the SNMP rules for independence in setting object
    // values, it is necessary to do final error checking when an expression is
    // evaluated, that is, when one of its instances in expValueTable is read or a
    // delta interval expires.  Earlier checking need not be done and an
    // implementation may not impose any ordering on the creation of objects
    // related to an expression.  To maintain security of MIB information, when
    // creating a new row in this table, the managed system must record the
    // security credentials of the requester.  These security credentials are the
    // parameters necessary as inputs to isAccessAllowed from the Architecture for
    // Describing SNMP Management Frameworks.  When obtaining the objects that
    // make up the expression, the system must (conceptually) use isAccessAllowed
    // to ensure that it does not violate security.  The evaluation of the
    // expression takes place under the security credentials of the creator of its
    // expExpressionEntry.  Values of read-write objects in this table may be
    // changed at any time. The type is slice of
    // DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry.
    Expexpressionentry []DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetFilter() yfilter.YFilter { return expexpressiontable.YFilter }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) SetFilter(yf yfilter.YFilter) { expexpressiontable.YFilter = yf }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetGoName(yname string) string {
    if yname == "expExpressionEntry" { return "Expexpressionentry" }
    return ""
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetSegmentPath() string {
    return "expExpressionTable"
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expExpressionEntry" {
        for _, c := range expexpressiontable.Expexpressionentry {
            if expexpressiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry{}
        expexpressiontable.Expexpressionentry = append(expexpressiontable.Expexpressionentry, child)
        return &expexpressiontable.Expexpressionentry[len(expexpressiontable.Expexpressionentry)-1]
    }
    return nil
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expexpressiontable.Expexpressionentry {
        children[expexpressiontable.Expexpressionentry[i].GetSegmentPath()] = &expexpressiontable.Expexpressionentry[i]
    }
    return children
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetBundleName() string { return "cisco_ios_xe" }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetYangName() string { return "expExpressionTable" }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) SetParent(parent types.Entity) { expexpressiontable.parent = parent }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetParent() types.Entity { return expexpressiontable.parent }

func (expexpressiontable *DISMANEXPRESSIONMIB_Expexpressiontable) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry
// Information about a single expression.  New expressions
// can be created using expExpressionRowStatus.
// 
// To create an expression first create the named entry in this
// table.  Then use expExpressionName to populate expObjectTable.
// For expression evaluation to succeed all related entries in
// expExpressionTable and expObjectTable must be 'active'.  If
// these conditions are not met the corresponding values in
// expValue simply are not instantiated.
// 
// Deleting an entry deletes all related entries in expObjectTable
// and expErrorTable.
// 
// Because of the relationships among the multiple tables for an
// expression (expExpressionTable, expObjectTable, and
// expValueTable) and the SNMP rules for independence in setting
// object values, it is necessary to do final error checking when
// an expression is evaluated, that is, when one of its instances
// in expValueTable is read or a delta interval expires.  Earlier
// checking need not be done and an implementation may not impose
// any ordering on the creation of objects related to an
// expression.
// 
// To maintain security of MIB information, when creating a new row in
// this table, the managed system must record the security credentials
// of the requester.  These security credentials are the parameters
// necessary as inputs to isAccessAllowed from the Architecture for
// 
// Describing SNMP Management Frameworks.  When obtaining the objects
// that make up the expression, the system must (conceptually) use
// isAccessAllowed to ensure that it does not violate security.
// 
// The evaluation of the expression takes place under the
// security credentials of the creator of its expExpressionEntry.
// 
// Values of read-write objects in this table may be changed
// at any time.
type DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The owner of this entry. The exact semantics of
    // this string are subject to the security policy defined by the security
    // administrator. The type is string with length: 0..32.
    Expexpressionowner interface{}

    // This attribute is a key. The name of the expression.  This is locally
    // unique, within the scope of an expExpressionOwner. The type is string with
    // length: 1..32.
    Expexpressionname interface{}

    // The expression to be evaluated.  This object is the same as a DisplayString
    // (RFC 1903) except for its maximum length.  Except for the variable names
    // the expression is in ANSI C syntax.  Only the subset of ANSI C operators
    // and functions listed here is allowed.  Variables are expressed as a dollar
    // sign ('$') and an integer that corresponds to an expObjectIndex.  An
    // example of a valid expression is:       ($1-$5)*100  Expressions must not
    // be recursive, that is although an expression may use the results of another
    // expression, it must not contain any variable that is directly or indirectly
    // a result of its own evaluation. The managed system must check for recursive
    // expressions.  The only allowed operators are:       ( )      - (unary)     
    // + - * / %      & | ^ << >> ~      ! && || == != > >= < <=  Note the
    // parentheses are included for parenthesizing the expression, not for casting
    // data types.  The only constant types defined are:       int (32-bit signed)
    // long (64-bit signed)      unsigned int      unsigned long      hexadecimal 
    // character      string      oid  The default type for a positive integer is
    // int unless it is too large in which case it is long.  All but oid are as
    // defined for ANSI C.  Note that a hexadecimal constant may end up as a
    // scalar or an array of 8-bit integers.  A string constant is enclosed in
    // double quotes and may contain back-slashed individual characters as in ANSI
    // C.  An oid constant comprises 32-bit, unsigned integers and at least one
    // period, for example:       0.      .0      1.3.6.1  No additional leading
    // or trailing subidentifiers are automatically added to an OID constant.  The
    // constant is taken as expressed.  Integer-typed objects are treated as 32-
    // or 64-bit, signed or unsigned integers, as appropriate.  The results of
    // mixing them are as for ANSI C, including the type of the result.  Note that
    // a 32-bit value is thus promoted to 64 bits only in an operation with a
    // 64-bit value.  There is no provision for larger values to handle overflow. 
    // Relative to SNMP data types, a resulting value becomes unsigned when
    // calculating it uses any unsigned value, including a counter.  To force the
    // final value to be of data type counter the expression must explicitly use
    // the counter32() or counter64() function (defined below).  OCTET STRINGS and
    // OBJECT IDENTIFIERs are treated as one-dimensioned arrays of unsigned 8-bit
    // integers and unsigned 32-bit integers, respectively.  IpAddresses are
    // treated as 32-bit, unsigned integers in network byte order, that is, the
    // hex version of 255.0.0.0 is 0xff000000.  Conditional expressions result in
    // a 32-bit, unsigned integer of value 0 for false or 1 for true. When an
    // arbitrary value is used as a boolean 0 is false and non-zero is true. 
    // Rules for the resulting data type from an operation, based on the operator:
    // For << and >> the result is the same as the left hand operand.  For &&, ||,
    // ==, !=, <, <=, >, and >= the result is always Unsigned32.  For unary - the
    // result is always Integer32.  For +, -, *, /, %, &, |, and ^ the result is
    // promoted according to the following rules, in order from most to least
    // preferred:       If left hand and right hand operands are the same type,   
    // use that.       If either side is Counter64, use that.       If either side
    // is IpAddress, use that.       If either side is TimeTicks, use that.      
    // If either side is Counter32, use that.       Otherwise use Unsigned32.  The
    // following rules say what operators apply with what data types.  Any
    // combination not explicitly defined does not work.  For all operators any of
    // the following can be the left hand or right hand operand: Integer32,
    // Counter32, Unsigned32, Counter64.  The operators +, -, *, /, %, <, <=, >,
    // and >= work with TimeTicks.  The operators &, |, and ^ work with IpAddress.
    // The operators << and >> work with IpAddress but only as the left hand
    // operand.  The + operator performs a concatenation of two OCTET STRINGs or
    // two OBJECT IDENTIFIERs.  The operators &, | perform bitwise operations on
    // OCTET STRINGs. If the OCTET STRING happens to be a DisplayString the
    // results may be meaningless, but the agent system does not check this as
    // some such systems do not have this information.  The operators << and >>
    // perform bitwise operations on OCTET STRINGs appearing as the left hand
    // operand.  The only functions defined are:       counter32      counter64   
    // arraySection      stringBegins      stringEnds      stringContains     
    // oidBegins      oidEnds      oidContains      average      maximum     
    // minimum      sum      exists  The following function definitions indicate
    // their parameters by naming the data type of the parameter in the
    // parameter's position in the parameter list.  The parameter must be of the
    // type indicated and generally may be a constant, a MIB object, a function,
    // or an expression.  counter32(integer) - wrapped around an integer value
    // counter32 forces Counter32 as a data type.  counter64(integer) - similar to
    // counter32 except that the resulting data type is 'counter64'. 
    // arraySection(array, integer, integer) - selects a piece of an array (i.e.
    // part of an OCTET STRING or OBJECT IDENTIFIER).  The integer arguments are
    // in the range 0 to 4,294,967,295.  The first is an initial array index
    // (one-dimensioned) and the second is an ending array index.  A value of 0
    // indicates first or last element, respectively.  If the first element is
    // larger than the array length the result is 0 length.  If the second integer
    // is less than or equal to the first, the result is 0 length.  If the second
    // is larger than the array length it indicates last element. 
    // stringBegins/Ends/Contains(octetString, octetString) - looks for the second
    // string (which can be a string constant) in the first and returns the
    // one-dimensioned arrayindex where the match began. A return value of 0
    // indicates no match (i.e. boolean false).  oidBegins/Ends/Contains(oid, oid)
    // - looks for the second OID (which can be an OID constant) in the first and
    // returns the the one-dimensioned index where the match began. A return value
    // of 0 indicates no match (i.e. boolean false). 
    // average/maximum/minimum(integer) - calculates the average, minimum, or
    // maximum value of the integer valued object over multiple sample times.  If
    // the object disappears for any sample period, the accumulation and the
    // resulting value object cease to exist until the object reappears at which
    // point the calculation starts over.  sum(integerObject*) - sums all
    // available values of the wildcarded integer object, resulting in an integer
    // scalar.  Must be used with caution as it wraps on overflow with no
    // notification.  exists(anyTypeObject) - verifies the object instance exists.
    // A return value of 0 indicates NoSuchInstance (i.e. boolean false). The type
    // is string with length: 1..1024.
    Expexpression interface{}

    // The type of the expression value.  One and only one of the value objects in
    // expValueTable will be instantiated to match this type.  If the result of
    // the expression can not be made into this type, an invalidOperandType error
    // will occur. The type is Expexpressionvaluetype.
    Expexpressionvaluetype interface{}

    // A comment to explain the use or meaning of the expression. The type is
    // string.
    Expexpressioncomment interface{}

    // Sampling interval for objects in this expression with expObjectSampleType
    // 'deltaValue'.  This object has no effect if the the expression has no
    // deltaValue objects.  A value of 0 indicates no automated sampling.  In this
    // case the delta is the difference from the last time the expression was
    // evaluated.  Note that this is subject to unpredictable delta times in the
    // face of retries or multiple managers.  A value greater than zero is the
    // number of seconds between automated samples.  Until the delta interval has
    // expired once the delta for the object is effectively not instantiated and
    // evaluating the expression has results as if the object itself were not
    // instantiated.  Note that delta values potentially consume large amounts of
    // system CPU and memory.  Delta state and processing must continue constantly
    // even if the expression is not being used. That is, the expression is being
    // evaluated every delta interval, even if no application is reading those
    // values.  For wildcarded objects this can be substantial overhead.  Note
    // that delta intervals, external expression value sampling intervals and
    // delta intervals for expressions within other expressions can have unusual
    // interactions as they are impossible to synchronize accurately.  In general
    // one interval embedded below another must be enough shorter that the higher
    // sample sees relatively smooth, predictable behavior.  So, for example, to
    // avoid the higher level getting the same sample twice, the lower level
    // should sample at least twice as fast as the higher level does. The type is
    // interface{} with range: 0..86400. Units are seconds.
    Expexpressiondeltainterval interface{}

    // An object prefix to assist an application in determining the instance
    // indexing to use in expValueTable, relieving the application of the need to
    // scan the expObjectTable to determine such a prefix.  See expObjectTable for
    // information on wildcarded objects.  If the expValueInstance portion of the
    // value OID may be treated as a scalar (that is, normally, 0) the value of
    // expExpressionPrefix is zero length, that is, no OID at all. Note that zero
    // length implies a null OID, not the OID 0.0.  Otherwise, the value of
    // expExpressionPrefix is the expObjectID value of any one of the wildcarded
    // objects for the expression. This is sufficient, as the remainder, that is,
    // the instance fragment relevant to instancing the values, must be the same
    // for all wildcarded objects in the expression. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expexpressionprefix interface{}

    // The number of errors encountered while evaluating this expression.  Note
    // that an object in the expression not being accessible, is not considered an
    // error. An example of an inaccessible object is when the object is excluded
    // from the view of the user whose security credentials are used in the
    // expression evaluation. In such cases, it is a legitimate condition that
    // causes the corresponding expression value not to be instantiated. The type
    // is interface{} with range: 0..4294967295.
    Expexpressionerrors interface{}

    // The control that allows creation and deletion of entries. The type is
    // RowStatus.
    Expexpressionentrystatus interface{}
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetFilter() yfilter.YFilter { return expexpressionentry.YFilter }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) SetFilter(yf yfilter.YFilter) { expexpressionentry.YFilter = yf }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetGoName(yname string) string {
    if yname == "expExpressionOwner" { return "Expexpressionowner" }
    if yname == "expExpressionName" { return "Expexpressionname" }
    if yname == "expExpression" { return "Expexpression" }
    if yname == "expExpressionValueType" { return "Expexpressionvaluetype" }
    if yname == "expExpressionComment" { return "Expexpressioncomment" }
    if yname == "expExpressionDeltaInterval" { return "Expexpressiondeltainterval" }
    if yname == "expExpressionPrefix" { return "Expexpressionprefix" }
    if yname == "expExpressionErrors" { return "Expexpressionerrors" }
    if yname == "expExpressionEntryStatus" { return "Expexpressionentrystatus" }
    return ""
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetSegmentPath() string {
    return "expExpressionEntry" + "[expExpressionOwner='" + fmt.Sprintf("%v", expexpressionentry.Expexpressionowner) + "']" + "[expExpressionName='" + fmt.Sprintf("%v", expexpressionentry.Expexpressionname) + "']"
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionOwner"] = expexpressionentry.Expexpressionowner
    leafs["expExpressionName"] = expexpressionentry.Expexpressionname
    leafs["expExpression"] = expexpressionentry.Expexpression
    leafs["expExpressionValueType"] = expexpressionentry.Expexpressionvaluetype
    leafs["expExpressionComment"] = expexpressionentry.Expexpressioncomment
    leafs["expExpressionDeltaInterval"] = expexpressionentry.Expexpressiondeltainterval
    leafs["expExpressionPrefix"] = expexpressionentry.Expexpressionprefix
    leafs["expExpressionErrors"] = expexpressionentry.Expexpressionerrors
    leafs["expExpressionEntryStatus"] = expexpressionentry.Expexpressionentrystatus
    return leafs
}

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetBundleName() string { return "cisco_ios_xe" }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetYangName() string { return "expExpressionEntry" }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) SetParent(parent types.Entity) { expexpressionentry.parent = parent }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetParent() types.Entity { return expexpressionentry.parent }

func (expexpressionentry *DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetParentYangName() string { return "expExpressionTable" }

// DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype represents an invalidOperandType error will occur.
type DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype string

const (
    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_counter32 DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "counter32"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_unsigned32 DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "unsigned32"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_timeTicks DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "timeTicks"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_integer32 DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "integer32"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_ipAddress DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "ipAddress"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_octetString DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "octetString"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_objectId DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "objectId"

    DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_counter64 DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "counter64"
)

// DISMANEXPRESSIONMIB_Experrortable
// A table of expression errors.
type DISMANEXPRESSIONMIB_Experrortable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about errors in processing an expression.  Entries appear in
    // this table only when there is a matching expExpressionEntry and then only
    // when there has been an error for that expression as reflected by the error
    // codes defined for expErrorCode. The type is slice of
    // DISMANEXPRESSIONMIB_Experrortable_Experrorentry.
    Experrorentry []DISMANEXPRESSIONMIB_Experrortable_Experrorentry
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetFilter() yfilter.YFilter { return experrortable.YFilter }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) SetFilter(yf yfilter.YFilter) { experrortable.YFilter = yf }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetGoName(yname string) string {
    if yname == "expErrorEntry" { return "Experrorentry" }
    return ""
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetSegmentPath() string {
    return "expErrorTable"
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expErrorEntry" {
        for _, c := range experrortable.Experrorentry {
            if experrortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEXPRESSIONMIB_Experrortable_Experrorentry{}
        experrortable.Experrorentry = append(experrortable.Experrorentry, child)
        return &experrortable.Experrorentry[len(experrortable.Experrorentry)-1]
    }
    return nil
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range experrortable.Experrorentry {
        children[experrortable.Experrorentry[i].GetSegmentPath()] = &experrortable.Experrorentry[i]
    }
    return children
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetBundleName() string { return "cisco_ios_xe" }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetYangName() string { return "expErrorTable" }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) SetParent(parent types.Entity) { experrortable.parent = parent }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetParent() types.Entity { return experrortable.parent }

func (experrortable *DISMANEXPRESSIONMIB_Experrortable) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Experrortable_Experrorentry
// Information about errors in processing an expression.
// 
// Entries appear in this table only when there is a matching
// expExpressionEntry and then only when there has been an
// error for that expression as reflected by the error codes
// defined for expErrorCode.
type DISMANEXPRESSIONMIB_Experrortable_Experrorentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionowner
    Expexpressionowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionname
    Expexpressionname interface{}

    // The value of sysUpTime the last time an error caused a failure to evaluate
    // this expression. The type is interface{} with range: 0..4294967295.
    Experrortime interface{}

    // The one-dimensioned character array index into expExpression for where the
    // error occurred.  The value zero indicates irrelevance. The type is
    // interface{} with range: -2147483648..2147483647.
    Experrorindex interface{}

    // The error that occurred.  In the following explanations the expected timing
    // of the error is in parentheses.  'S' means the error occurs on a Set
    // request.  'E' means the error occurs on the attempt to evaluate the
    // expression either due to Get from expValueTable or in ongoing delta
    // processing.  invalidSyntax       the value sent for expExpression is not   
    // valid Expression MIB expression syntax                (S)
    // undefinedObjectIndex     an object reference ($n) in               
    // expExpression does not have a matching                instance in
    // expObjectTable (E) unrecognizedOperator     the value sent for
    // expExpression held an                unrecognized operator (S)
    // unrecognizedFunction     the value sent for expExpression held an          
    // unrecognized function name (S) invalidOperandType  an operand in
    // expExpression is not the                right type for the associated
    // operator                or result (SE) unmatchedParenthesis     the value
    // sent for expExpression is not                correctly parenthesized (S)
    // tooManyWildcardValues    evaluating the expression exceeded the            
    // limit set by                expResourceDeltaWildcardInstanceMaximum        
    // (E) recursion      through some chain of embedded               
    // expressions the expression invokes itself                (E) deltaTooShort 
    // the delta for the next evaluation passed                before the system
    // could evaluate the                present sample (E) resourceUnavailable
    // some resource, typically dynamic memory,                was unavailable
    // (SE) divideByZero        an attempt to divide by zero occurred             
    // (E)  For the errors that occur when the attempt is made to set
    // expExpression Set request fails with the SNMP error code 'wrongValue'. 
    // Such failures refer to the most recent failure to Set expExpression, not to
    // the present value of expExpression which must be either unset or
    // syntactically correct.  Errors that occur during evaluation for a Get*
    // operation return the SNMP error code 'genErr' except for
    // 'tooManyWildcardValues' and 'resourceUnavailable' which return the SNMP
    // error code 'resourceUnavailable'. The type is Experrorcode.
    Experrorcode interface{}

    // The expValueInstance being evaluated when the error occurred.  A
    // zero-length indicates irrelevance. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Experrorinstance interface{}
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetFilter() yfilter.YFilter { return experrorentry.YFilter }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) SetFilter(yf yfilter.YFilter) { experrorentry.YFilter = yf }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetGoName(yname string) string {
    if yname == "expExpressionOwner" { return "Expexpressionowner" }
    if yname == "expExpressionName" { return "Expexpressionname" }
    if yname == "expErrorTime" { return "Experrortime" }
    if yname == "expErrorIndex" { return "Experrorindex" }
    if yname == "expErrorCode" { return "Experrorcode" }
    if yname == "expErrorInstance" { return "Experrorinstance" }
    return ""
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetSegmentPath() string {
    return "expErrorEntry" + "[expExpressionOwner='" + fmt.Sprintf("%v", experrorentry.Expexpressionowner) + "']" + "[expExpressionName='" + fmt.Sprintf("%v", experrorentry.Expexpressionname) + "']"
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionOwner"] = experrorentry.Expexpressionowner
    leafs["expExpressionName"] = experrorentry.Expexpressionname
    leafs["expErrorTime"] = experrorentry.Experrortime
    leafs["expErrorIndex"] = experrorentry.Experrorindex
    leafs["expErrorCode"] = experrorentry.Experrorcode
    leafs["expErrorInstance"] = experrorentry.Experrorinstance
    return leafs
}

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetBundleName() string { return "cisco_ios_xe" }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetYangName() string { return "expErrorEntry" }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) SetParent(parent types.Entity) { experrorentry.parent = parent }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetParent() types.Entity { return experrorentry.parent }

func (experrorentry *DISMANEXPRESSIONMIB_Experrortable_Experrorentry) GetParentYangName() string { return "expErrorTable" }

// DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode represents 'resourceUnavailable'.
type DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode string

const (
    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_invalidSyntax DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "invalidSyntax"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_undefinedObjectIndex DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "undefinedObjectIndex"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_unrecognizedOperator DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "unrecognizedOperator"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_unrecognizedFunction DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "unrecognizedFunction"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_invalidOperandType DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "invalidOperandType"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_unmatchedParenthesis DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "unmatchedParenthesis"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_tooManyWildcardValues DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "tooManyWildcardValues"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_recursion DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "recursion"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_deltaTooShort DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "deltaTooShort"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_resourceUnavailable DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "resourceUnavailable"

    DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode_divideByZero DISMANEXPRESSIONMIB_Experrortable_Experrorentry_Experrorcode = "divideByZero"
)

// DISMANEXPRESSIONMIB_Expobjecttable
// A table of object definitions for each expExpression.
// 
// Wildcarding instance IDs:
// 
// It is legal to omit all or part of the instance portion for
// some or all of the objects in an expression. (See the
// DESCRIPTION of expObjectID for details.  However, note that
// if more than one object in the same expression is wildcarded
// in this way, they all must be objects where that portion of
// the instance is the same.  In other words, all objects may be
// in the same SEQUENCE or in different SEQUENCEs but with the
// same semantic index value (e.g., a value of ifIndex)
// for the wildcarded portion.
type DISMANEXPRESSIONMIB_Expobjecttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an object.  An application uses expObjectEntryStatus to
    // create entries in this table while in the process of defining an
    // expression.  Values of read-create objects in this table may be changed at
    // any time. The type is slice of
    // DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry.
    Expobjectentry []DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetFilter() yfilter.YFilter { return expobjecttable.YFilter }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) SetFilter(yf yfilter.YFilter) { expobjecttable.YFilter = yf }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetGoName(yname string) string {
    if yname == "expObjectEntry" { return "Expobjectentry" }
    return ""
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetSegmentPath() string {
    return "expObjectTable"
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expObjectEntry" {
        for _, c := range expobjecttable.Expobjectentry {
            if expobjecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry{}
        expobjecttable.Expobjectentry = append(expobjecttable.Expobjectentry, child)
        return &expobjecttable.Expobjectentry[len(expobjecttable.Expobjectentry)-1]
    }
    return nil
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expobjecttable.Expobjectentry {
        children[expobjecttable.Expobjectentry[i].GetSegmentPath()] = &expobjecttable.Expobjectentry[i]
    }
    return children
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetBundleName() string { return "cisco_ios_xe" }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetYangName() string { return "expObjectTable" }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) SetParent(parent types.Entity) { expobjecttable.parent = parent }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetParent() types.Entity { return expobjecttable.parent }

func (expobjecttable *DISMANEXPRESSIONMIB_Expobjecttable) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry
// Information about an object.  An application uses
// expObjectEntryStatus to create entries in this table while
// in the process of defining an expression.
// 
// Values of read-create objects in this table may be
// changed at any time.
type DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionowner
    Expexpressionowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionname
    Expexpressionname interface{}

    // This attribute is a key. Within an expression, a unique, numeric
    // identification for an object.  Prefixed with a dollar sign ('$') this is
    // used to reference the object in the corresponding expExpression. The type
    // is interface{} with range: 1..4294967295.
    Expobjectindex interface{}

    // The OBJECT IDENTIFIER (OID) of this object.  The OID may be fully
    // qualified, meaning it includes a complete instance identifier part (e.g.,
    // ifInOctets.1 or sysUpTime.0), or it may not be fully qualified, meaning it
    // may lack all or part of the instance identifier.  If the expObjectID is not
    // fully qualified, then expObjectWildcard must be set to true(1). The value
    // of the expression will be multiple values, as if done for a GetNext sweep
    // of the object.  An object here may itself be the result of an expression
    // but recursion is not allowed.  NOTE:  The simplest implementations of this
    // MIB may not allow wildcards. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectid interface{}

    // A true value indicates the expObjecID of this row is a wildcard object.
    // False indicates that expObjectID is fully instanced. If all
    // expObjectWildcard values for a given expression are FALSE,
    // expExpressionPrefix will reflect a scalar object (i.e. will be 0.0).  NOTE:
    // The simplest implementations of this MIB may not allow wildcards. The type
    // is bool.
    Expobjectidwildcard interface{}

    // The method of sampling the selected variable.  An 'absoluteValue' is simply
    // the present value of the object.  A 'deltaValue' is the present value minus
    // the previous value, which was sampled expExpressionDeltaInterval seconds
    // ago. This is intended primarily for use with SNMP counters, which are
    // meaningless as an 'absoluteValue', but may be used with any integer-based
    // value.  A 'changedValue' is a boolean for whether the present value is
    // different from the previous value.  It is applicable to any data type and
    // results in an Unsigned32 with value 1 if the object's value is changed and
    // 0 if not.  In all other respects it is as a 'deltaValue' and all statements
    // and operation regarding delta values apply to changed values.  When an
    // expression contains both delta and absolute values the absolute values are
    // obtained at the end of the delta period. The type is Expobjectsampletype.
    Expobjectsampletype interface{}

    // The OBJECT IDENTIFIER (OID) of a TimeTicks, TimeStamp, or DateAndTime
    // object that indicates a discontinuity in the value at expObjectID.  This
    // object is instantiated only if expObjectSampleType is 'deltaValue' or
    // 'changedValue'.  The OID may be for a leaf object (e.g. sysUpTime.0) or may
    // be wildcarded to match expObjectID.  This object supports normal checking
    // for a discontinuity in a counter.  Note that if this object does not point
    // to sysUpTime discontinuity checking must still check sysUpTime for an
    // overall discontinuity.  If the object identified is not accessible no
    // discontinuity check will be made. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectdeltadiscontinuityid interface{}

    // A true value indicates the expObjectDeltaDiscontinuityID of this row is a
    // wildcard object.  False indicates that expObjectDeltaDiscontinuityID is
    // fully instanced.  This object is instantiated only if expObjectSampleType
    // is 'deltaValue' or 'changedValue'.  NOTE:  The simplest implementations of
    // this MIB may not allow wildcards. The type is bool.
    Expobjectdiscontinuityidwildcard interface{}

    // The value 'timeTicks' indicates the expObjectDeltaDiscontinuityID of this
    // row is of syntax TimeTicks.  The value 'timeStamp' indicates syntax
    // TimeStamp.  The value 'dateAndTime indicates syntax DateAndTime.  This
    // object is instantiated only if expObjectSampleType is 'deltaValue' or
    // 'changedValue'. The type is Expobjectdiscontinuityidtype.
    Expobjectdiscontinuityidtype interface{}

    // The OBJECT IDENTIFIER (OID) of an object that overrides whether the
    // instance of expObjectID is to be considered usable.  If the value of the
    // object at expObjectConditional is 0 or not instantiated, the object at
    // expObjectID is treated as if it is not instantiated.  In other words,
    // expObjectConditional is a filter that controls whether or not to use the
    // value at expObjectID.  The OID may be for a leaf object (e.g.
    // sysObjectID.0) or may be wildcarded to match expObjectID.  If expObject is
    // wildcarded and expObjectID in the same row is not, the wild portion of
    // expObjectConditional must match the wildcarding of the rest of the
    // expression.  If no object in the expression is wildcarded but
    // expObjectConditional is, use the lexically first instance (if any) of
    // expObjectConditional.  If the value of expObjectConditional is 0.0
    // operation is as if the value pointed to by expObjectConditional is a
    // non-zero (true) value.  Note that expObjectConditional can not trivially
    // use an object of syntax TruthValue, since the underlying value is not 0 or
    // 1. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectconditional interface{}

    // A true value indicates the expObjectConditional of this row is a wildcard
    // object. False indicates that expObjectConditional is fully instanced. 
    // NOTE: The simplest implementations of this MIB may not allow wildcards. The
    // type is bool.
    Expobjectconditionalwildcard interface{}

    // The control that allows creation/deletion of entries.  Objects in this
    // table may be changed while expObjectEntryStatus is in any state. The type
    // is RowStatus.
    Expobjectentrystatus interface{}
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetFilter() yfilter.YFilter { return expobjectentry.YFilter }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) SetFilter(yf yfilter.YFilter) { expobjectentry.YFilter = yf }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetGoName(yname string) string {
    if yname == "expExpressionOwner" { return "Expexpressionowner" }
    if yname == "expExpressionName" { return "Expexpressionname" }
    if yname == "expObjectIndex" { return "Expobjectindex" }
    if yname == "expObjectID" { return "Expobjectid" }
    if yname == "expObjectIDWildcard" { return "Expobjectidwildcard" }
    if yname == "expObjectSampleType" { return "Expobjectsampletype" }
    if yname == "expObjectDeltaDiscontinuityID" { return "Expobjectdeltadiscontinuityid" }
    if yname == "expObjectDiscontinuityIDWildcard" { return "Expobjectdiscontinuityidwildcard" }
    if yname == "expObjectDiscontinuityIDType" { return "Expobjectdiscontinuityidtype" }
    if yname == "expObjectConditional" { return "Expobjectconditional" }
    if yname == "expObjectConditionalWildcard" { return "Expobjectconditionalwildcard" }
    if yname == "expObjectEntryStatus" { return "Expobjectentrystatus" }
    return ""
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetSegmentPath() string {
    return "expObjectEntry" + "[expExpressionOwner='" + fmt.Sprintf("%v", expobjectentry.Expexpressionowner) + "']" + "[expExpressionName='" + fmt.Sprintf("%v", expobjectentry.Expexpressionname) + "']" + "[expObjectIndex='" + fmt.Sprintf("%v", expobjectentry.Expobjectindex) + "']"
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionOwner"] = expobjectentry.Expexpressionowner
    leafs["expExpressionName"] = expobjectentry.Expexpressionname
    leafs["expObjectIndex"] = expobjectentry.Expobjectindex
    leafs["expObjectID"] = expobjectentry.Expobjectid
    leafs["expObjectIDWildcard"] = expobjectentry.Expobjectidwildcard
    leafs["expObjectSampleType"] = expobjectentry.Expobjectsampletype
    leafs["expObjectDeltaDiscontinuityID"] = expobjectentry.Expobjectdeltadiscontinuityid
    leafs["expObjectDiscontinuityIDWildcard"] = expobjectentry.Expobjectdiscontinuityidwildcard
    leafs["expObjectDiscontinuityIDType"] = expobjectentry.Expobjectdiscontinuityidtype
    leafs["expObjectConditional"] = expobjectentry.Expobjectconditional
    leafs["expObjectConditionalWildcard"] = expobjectentry.Expobjectconditionalwildcard
    leafs["expObjectEntryStatus"] = expobjectentry.Expobjectentrystatus
    return leafs
}

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetBundleName() string { return "cisco_ios_xe" }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetYangName() string { return "expObjectEntry" }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) SetParent(parent types.Entity) { expobjectentry.parent = parent }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetParent() types.Entity { return expobjectentry.parent }

func (expobjectentry *DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry) GetParentYangName() string { return "expObjectTable" }

// DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype represents 'deltaValue' or 'changedValue'.
type DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype string

const (
    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype_timeTicks DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype = "timeTicks"

    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype_timeStamp DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype = "timeStamp"

    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype_dateAndTime DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype = "dateAndTime"
)

// DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype represents period.
type DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype string

const (
    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype_absoluteValue DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype = "absoluteValue"

    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype_deltaValue DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype = "deltaValue"

    DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype_changedValue DISMANEXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype = "changedValue"
)

// DISMANEXPRESSIONMIB_Expvaluetable
// A table of values from evaluated expressions.
type DISMANEXPRESSIONMIB_Expvaluetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single value from an evaluated expression.  For a given instance, only
    // one 'Val' object in the conceptual row will be instantiated, that is, the
    // one with the appropriate type for the value.  For values that contain no
    // objects of expObjectSampleType 'deltaValue' or 'changedValue', reading a
    // value from the table causes the evaluation of the expression for that
    // value.  For those that contain a 'deltaValue' or 'changedValue' the value
    // read is as of the last sampling interval.  If in the attempt to evaluate
    // the expression one or more of the necessary objects is not available, the
    // corresponding entry in this table is effectively not instantiated.  To
    // maintain security of MIB information, when creating a new row in this
    // table, the managed system must record the security credentials of the
    // requester.  These security credentials are the parameters necessary as
    // inputs to isAccessAllowed from [RFC2571]. When obtaining the objects that
    // make up the expression, the system must (conceptually) use isAccessAllowed
    // to ensure that it does not violate security.  The evaluation of that
    // expression takes place under the security credentials of the creator of its
    // expExpressionEntry.  To maintain security of MIB information, expression
    // evaluation must take place using security credentials for the implied Gets
    // of the objects in the expression as inputs (conceptually) to
    // isAccessAllowed from the Architecture for Describing SNMP Management
    // Frameworks.  These are the security credentials of the creator of the
    // corresponding expExpressionEntry. The type is slice of
    // DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry.
    Expvalueentry []DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetFilter() yfilter.YFilter { return expvaluetable.YFilter }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) SetFilter(yf yfilter.YFilter) { expvaluetable.YFilter = yf }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetGoName(yname string) string {
    if yname == "expValueEntry" { return "Expvalueentry" }
    return ""
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetSegmentPath() string {
    return "expValueTable"
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expValueEntry" {
        for _, c := range expvaluetable.Expvalueentry {
            if expvaluetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry{}
        expvaluetable.Expvalueentry = append(expvaluetable.Expvalueentry, child)
        return &expvaluetable.Expvalueentry[len(expvaluetable.Expvalueentry)-1]
    }
    return nil
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expvaluetable.Expvalueentry {
        children[expvaluetable.Expvalueentry[i].GetSegmentPath()] = &expvaluetable.Expvalueentry[i]
    }
    return children
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetBundleName() string { return "cisco_ios_xe" }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetYangName() string { return "expValueTable" }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) SetParent(parent types.Entity) { expvaluetable.parent = parent }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetParent() types.Entity { return expvaluetable.parent }

func (expvaluetable *DISMANEXPRESSIONMIB_Expvaluetable) GetParentYangName() string { return "DISMAN-EXPRESSION-MIB" }

// DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry
// A single value from an evaluated expression.  For a given
// instance, only one 'Val' object in the conceptual row will be
// instantiated, that is, the one with the appropriate type for
// the value.  For values that contain no objects of
// expObjectSampleType 'deltaValue' or 'changedValue', reading a
// value from the table causes the evaluation of the expression
// for that value.  For those that contain a 'deltaValue' or
// 'changedValue' the value read is as of the last sampling
// interval.
// 
// If in the attempt to evaluate the expression one or more
// of the necessary objects is not available, the corresponding
// entry in this table is effectively not instantiated.
// 
// To maintain security of MIB information, when creating a new
// row in this table, the managed system must record the security
// credentials of the requester.  These security credentials are
// the parameters necessary as inputs to isAccessAllowed from
// [RFC2571]. When obtaining the objects that make up the
// expression, the system must (conceptually) use isAccessAllowed to
// ensure that it does not violate security.
// 
// The evaluation of that expression takes place under the
// security credentials of the creator of its expExpressionEntry.
// 
// To maintain security of MIB information, expression evaluation must
// take place using security credentials for the implied Gets of the
// objects in the expression as inputs (conceptually) to
// isAccessAllowed from the Architecture for Describing SNMP
// Management Frameworks.  These are the security credentials of the
// creator of the corresponding expExpressionEntry.
type DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionowner
    Expexpressionowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_expression_mib.DISMANEXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionname
    Expexpressionname interface{}

    // This attribute is a key. The final instance portion of a value's OID
    // according to the wildcarding in instances of expObjectID for the
    // expression.  The prefix of this OID fragment is 0.0, leading to the
    // following behavior.  If there is no wildcarding, the value is 0.0.0.  In
    // other words, there is one value which standing alone would have been a
    // scalar with a 0 at the end of its OID.  If there is wildcarding, the value
    // is 0.0 followed by a value that the wildcard can take, thus defining one
    // value instance for each real, possible value of the wildcard. So, for
    // example, if the wildcard worked out to be an ifIndex, there is an
    // expValueInstance for each applicable ifIndex. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expvalueinstance interface{}

    // The value when expExpressionValueType is 'counter32'. The type is
    // interface{} with range: 0..4294967295.
    Expvaluecounter32Val interface{}

    // The value when expExpressionValueType is 'unsigned32'. The type is
    // interface{} with range: 0..4294967295.
    Expvalueunsigned32Val interface{}

    // The value when expExpressionValueType is 'timeTicks'. The type is
    // interface{} with range: 0..4294967295.
    Expvaluetimeticksval interface{}

    // The value when expExpressionValueType is 'integer32'. The type is
    // interface{} with range: -2147483648..2147483647.
    Expvalueinteger32Val interface{}

    // The value when expExpressionValueType is 'ipAddress'. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Expvalueipaddressval interface{}

    // The value when expExpressionValueType is 'octetString'. The type is string
    // with length: 0..65536.
    Expvalueoctetstringval interface{}

    // The value when expExpressionValueType is 'objectId'. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expvalueoidval interface{}

    // The value when expExpressionValueType is 'counter64'. The type is
    // interface{} with range: 0..18446744073709551615.
    Expvaluecounter64Val interface{}
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetFilter() yfilter.YFilter { return expvalueentry.YFilter }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) SetFilter(yf yfilter.YFilter) { expvalueentry.YFilter = yf }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetGoName(yname string) string {
    if yname == "expExpressionOwner" { return "Expexpressionowner" }
    if yname == "expExpressionName" { return "Expexpressionname" }
    if yname == "expValueInstance" { return "Expvalueinstance" }
    if yname == "expValueCounter32Val" { return "Expvaluecounter32Val" }
    if yname == "expValueUnsigned32Val" { return "Expvalueunsigned32Val" }
    if yname == "expValueTimeTicksVal" { return "Expvaluetimeticksval" }
    if yname == "expValueInteger32Val" { return "Expvalueinteger32Val" }
    if yname == "expValueIpAddressVal" { return "Expvalueipaddressval" }
    if yname == "expValueOctetStringVal" { return "Expvalueoctetstringval" }
    if yname == "expValueOidVal" { return "Expvalueoidval" }
    if yname == "expValueCounter64Val" { return "Expvaluecounter64Val" }
    return ""
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetSegmentPath() string {
    return "expValueEntry" + "[expExpressionOwner='" + fmt.Sprintf("%v", expvalueentry.Expexpressionowner) + "']" + "[expExpressionName='" + fmt.Sprintf("%v", expvalueentry.Expexpressionname) + "']" + "[expValueInstance='" + fmt.Sprintf("%v", expvalueentry.Expvalueinstance) + "']"
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionOwner"] = expvalueentry.Expexpressionowner
    leafs["expExpressionName"] = expvalueentry.Expexpressionname
    leafs["expValueInstance"] = expvalueentry.Expvalueinstance
    leafs["expValueCounter32Val"] = expvalueentry.Expvaluecounter32Val
    leafs["expValueUnsigned32Val"] = expvalueentry.Expvalueunsigned32Val
    leafs["expValueTimeTicksVal"] = expvalueentry.Expvaluetimeticksval
    leafs["expValueInteger32Val"] = expvalueentry.Expvalueinteger32Val
    leafs["expValueIpAddressVal"] = expvalueentry.Expvalueipaddressval
    leafs["expValueOctetStringVal"] = expvalueentry.Expvalueoctetstringval
    leafs["expValueOidVal"] = expvalueentry.Expvalueoidval
    leafs["expValueCounter64Val"] = expvalueentry.Expvaluecounter64Val
    return leafs
}

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetBundleName() string { return "cisco_ios_xe" }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetYangName() string { return "expValueEntry" }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) SetParent(parent types.Entity) { expvalueentry.parent = parent }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetParent() types.Entity { return expvalueentry.parent }

func (expvalueentry *DISMANEXPRESSIONMIB_Expvaluetable_Expvalueentry) GetParentYangName() string { return "expValueTable" }

