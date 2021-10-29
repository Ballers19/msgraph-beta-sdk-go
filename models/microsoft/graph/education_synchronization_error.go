package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationSynchronizationError struct {
    Entity
    // Represents the sync entity (school, section, student, teacher).
    entryType *string;
    // Represents the error code for this error.
    errorCode *string;
    // Contains a description of the error.
    errorMessage *string;
    // The unique identifier for the entry.
    joiningValue *string;
    // The time of occurrence of this error.
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier of this error entry.
    reportableIdentifier *string;
}
// Instantiates a new educationSynchronizationError and sets the default values.
func NewEducationSynchronizationError()(*EducationSynchronizationError) {
    m := &EducationSynchronizationError{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the entryType property value. Represents the sync entity (school, section, student, teacher).
func (m *EducationSynchronizationError) GetEntryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entryType
    }
}
// Gets the errorCode property value. Represents the error code for this error.
func (m *EducationSynchronizationError) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the errorMessage property value. Contains a description of the error.
func (m *EducationSynchronizationError) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// Gets the joiningValue property value. The unique identifier for the entry.
func (m *EducationSynchronizationError) GetJoiningValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joiningValue
    }
}
// Gets the recordedDateTime property value. The time of occurrence of this error.
func (m *EducationSynchronizationError) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recordedDateTime
    }
}
// Gets the reportableIdentifier property value. The identifier of this error entry.
func (m *EducationSynchronizationError) GetReportableIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportableIdentifier
    }
}
// The deserialization information for the current model
func (m *EducationSynchronizationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEntryType(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorMessage(val)
        return nil
    }
    res["joiningValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoiningValue(val)
        return nil
    }
    res["recordedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRecordedDateTime(val)
        return nil
    }
    res["reportableIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportableIdentifier(val)
        return nil
    }
    return res
}
func (m *EducationSynchronizationError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationSynchronizationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("entryType", m.GetEntryType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joiningValue", m.GetJoiningValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("recordedDateTime", m.GetRecordedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportableIdentifier", m.GetReportableIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the entryType property value. Represents the sync entity (school, section, student, teacher).
// Parameters:
//  - value : Value to set for the entryType property.
func (m *EducationSynchronizationError) SetEntryType(value *string)() {
    m.entryType = value
}
// Sets the errorCode property value. Represents the error code for this error.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *EducationSynchronizationError) SetErrorCode(value *string)() {
    m.errorCode = value
}
// Sets the errorMessage property value. Contains a description of the error.
// Parameters:
//  - value : Value to set for the errorMessage property.
func (m *EducationSynchronizationError) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// Sets the joiningValue property value. The unique identifier for the entry.
// Parameters:
//  - value : Value to set for the joiningValue property.
func (m *EducationSynchronizationError) SetJoiningValue(value *string)() {
    m.joiningValue = value
}
// Sets the recordedDateTime property value. The time of occurrence of this error.
// Parameters:
//  - value : Value to set for the recordedDateTime property.
func (m *EducationSynchronizationError) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recordedDateTime = value
}
// Sets the reportableIdentifier property value. The identifier of this error entry.
// Parameters:
//  - value : Value to set for the reportableIdentifier property.
func (m *EducationSynchronizationError) SetReportableIdentifier(value *string)() {
    m.reportableIdentifier = value
}
