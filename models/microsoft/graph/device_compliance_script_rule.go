package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceComplianceScriptRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
    dataType *DataType;
    // Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
    deviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDataType;
    // Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
    deviceComplianceScriptRulOperator *DeviceComplianceScriptRulOperator;
    // Operand specified in the rule.
    operand *string;
    // Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
    operator *Operator;
    // Setting name specified in the rule.
    settingName *string;
}
// Instantiates a new deviceComplianceScriptRule and sets the default values.
func NewDeviceComplianceScriptRule()(*DeviceComplianceScriptRule) {
    m := &DeviceComplianceScriptRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) GetDataType()(*DataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// Gets the deviceComplianceScriptRuleDataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRuleDataType
    }
}
// Gets the deviceComplianceScriptRulOperator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRulOperator
    }
}
// Gets the operand property value. Operand specified in the rule.
func (m *DeviceComplianceScriptRule) GetOperand()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operand
    }
}
// Gets the operator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) GetOperator()(*Operator) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// Gets the settingName property value. Setting name specified in the rule.
func (m *DeviceComplianceScriptRule) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// The deserialization information for the current model
func (m *DeviceComplianceScriptRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataType)
        if err != nil {
            return err
        }
        cast := val.(DataType)
        m.SetDataType(&cast)
        return nil
    }
    res["deviceComplianceScriptRuleDataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRuleDataType)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceScriptRuleDataType)
        m.SetDeviceComplianceScriptRuleDataType(&cast)
        return nil
    }
    res["deviceComplianceScriptRulOperator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulOperator)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceScriptRulOperator)
        m.SetDeviceComplianceScriptRulOperator(&cast)
        return nil
    }
    res["operand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperand(val)
        return nil
    }
    res["operator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperator)
        if err != nil {
            return err
        }
        cast := val.(Operator)
        m.SetOperator(&cast)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptRule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceComplianceScriptRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := m.GetDataType().String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRuleDataType() != nil {
        cast := m.GetDeviceComplianceScriptRuleDataType().String()
        err := writer.WriteStringValue("deviceComplianceScriptRuleDataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulOperator() != nil {
        cast := m.GetDeviceComplianceScriptRulOperator().String()
        err := writer.WriteStringValue("deviceComplianceScriptRulOperator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operand", m.GetOperand())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := m.GetOperator().String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceComplianceScriptRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
// Parameters:
//  - value : Value to set for the dataType property.
func (m *DeviceComplianceScriptRule) SetDataType(value *DataType)() {
    m.dataType = value
}
// Sets the deviceComplianceScriptRuleDataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
// Parameters:
//  - value : Value to set for the deviceComplianceScriptRuleDataType property.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)() {
    m.deviceComplianceScriptRuleDataType = value
}
// Sets the deviceComplianceScriptRulOperator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
// Parameters:
//  - value : Value to set for the deviceComplianceScriptRulOperator property.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)() {
    m.deviceComplianceScriptRulOperator = value
}
// Sets the operand property value. Operand specified in the rule.
// Parameters:
//  - value : Value to set for the operand property.
func (m *DeviceComplianceScriptRule) SetOperand(value *string)() {
    m.operand = value
}
// Sets the operator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
// Parameters:
//  - value : Value to set for the operator property.
func (m *DeviceComplianceScriptRule) SetOperator(value *Operator)() {
    m.operator = value
}
// Sets the settingName property value. Setting name specified in the rule.
// Parameters:
//  - value : Value to set for the settingName property.
func (m *DeviceComplianceScriptRule) SetSettingName(value *string)() {
    m.settingName = value
}
