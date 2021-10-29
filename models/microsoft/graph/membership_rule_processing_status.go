package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MembershipRuleProcessingStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
    errorMessage *string;
    // Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
    lastMembershipUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
    status *MembershipRuleProcessingStatusDetails;
}
// Instantiates a new membershipRuleProcessingStatus and sets the default values.
func NewMembershipRuleProcessingStatus()(*MembershipRuleProcessingStatus) {
    m := &MembershipRuleProcessingStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MembershipRuleProcessingStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// Gets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetLastMembershipUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastMembershipUpdated
    }
}
// Gets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
func (m *MembershipRuleProcessingStatus) GetStatus()(*MembershipRuleProcessingStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *MembershipRuleProcessingStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorMessage(val)
        return nil
    }
    res["lastMembershipUpdated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastMembershipUpdated(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMembershipRuleProcessingStatusDetails)
        if err != nil {
            return err
        }
        cast := val.(MembershipRuleProcessingStatusDetails)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *MembershipRuleProcessingStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MembershipRuleProcessingStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastMembershipUpdated", m.GetLastMembershipUpdated())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *MembershipRuleProcessingStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
// Parameters:
//  - value : Value to set for the errorMessage property.
func (m *MembershipRuleProcessingStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// Sets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastMembershipUpdated property.
func (m *MembershipRuleProcessingStatus) SetLastMembershipUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMembershipUpdated = value
}
// Sets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
// Parameters:
//  - value : Value to set for the status property.
func (m *MembershipRuleProcessingStatus) SetStatus(value *MembershipRuleProcessingStatusDetails)() {
    m.status = value
}
