package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SynchronizationQuarantine struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    currentBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Describes the error(s) that occurred when putting the synchronization job into quarantine.
    error *SynchronizationError;
    // Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    nextAttempt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A code that signifies why the quarantine was imposed. Possible values are: EncounteredBaseEscrowThreshold, EncounteredTotalEscrowThreshold, EncounteredEscrowProportionThreshold, EncounteredQuarantineException, QuarantinedOnDemand, TooManyDeletes, Unknown.
    reason *QuarantineReason;
    // Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    seriesBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
    seriesCount *int64;
}
// Instantiates a new synchronizationQuarantine and sets the default values.
func NewSynchronizationQuarantine()(*SynchronizationQuarantine) {
    m := &SynchronizationQuarantine{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationQuarantine) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the currentBegan property value. Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetCurrentBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.currentBegan
    }
}
// Gets the error property value. Describes the error(s) that occurred when putting the synchronization job into quarantine.
func (m *SynchronizationQuarantine) GetError()(*SynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the nextAttempt property value. Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetNextAttempt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.nextAttempt
    }
}
// Gets the reason property value. A code that signifies why the quarantine was imposed. Possible values are: EncounteredBaseEscrowThreshold, EncounteredTotalEscrowThreshold, EncounteredEscrowProportionThreshold, EncounteredQuarantineException, QuarantinedOnDemand, TooManyDeletes, Unknown.
func (m *SynchronizationQuarantine) GetReason()(*QuarantineReason) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Gets the seriesBegan property value. Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetSeriesBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.seriesBegan
    }
}
// Gets the seriesCount property value. Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
func (m *SynchronizationQuarantine) GetSeriesCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.seriesCount
    }
}
// The deserialization information for the current model
func (m *SynchronizationQuarantine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentBegan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCurrentBegan(val)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*SynchronizationError))
        return nil
    }
    res["nextAttempt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetNextAttempt(val)
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseQuarantineReason)
        if err != nil {
            return err
        }
        cast := val.(QuarantineReason)
        m.SetReason(&cast)
        return nil
    }
    res["seriesBegan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSeriesBegan(val)
        return nil
    }
    res["seriesCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSeriesCount(val)
        return nil
    }
    return res
}
func (m *SynchronizationQuarantine) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SynchronizationQuarantine) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("currentBegan", m.GetCurrentBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("nextAttempt", m.GetNextAttempt())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := m.GetReason().String()
        err := writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("seriesBegan", m.GetSeriesBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("seriesCount", m.GetSeriesCount())
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
func (m *SynchronizationQuarantine) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the currentBegan property value. Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the currentBegan property.
func (m *SynchronizationQuarantine) SetCurrentBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.currentBegan = value
}
// Sets the error property value. Describes the error(s) that occurred when putting the synchronization job into quarantine.
// Parameters:
//  - value : Value to set for the error property.
func (m *SynchronizationQuarantine) SetError(value *SynchronizationError)() {
    m.error = value
}
// Sets the nextAttempt property value. Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the nextAttempt property.
func (m *SynchronizationQuarantine) SetNextAttempt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextAttempt = value
}
// Sets the reason property value. A code that signifies why the quarantine was imposed. Possible values are: EncounteredBaseEscrowThreshold, EncounteredTotalEscrowThreshold, EncounteredEscrowProportionThreshold, EncounteredQuarantineException, QuarantinedOnDemand, TooManyDeletes, Unknown.
// Parameters:
//  - value : Value to set for the reason property.
func (m *SynchronizationQuarantine) SetReason(value *QuarantineReason)() {
    m.reason = value
}
// Sets the seriesBegan property value. Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the seriesBegan property.
func (m *SynchronizationQuarantine) SetSeriesBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.seriesBegan = value
}
// Sets the seriesCount property value. Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
// Parameters:
//  - value : Value to set for the seriesCount property.
func (m *SynchronizationQuarantine) SetSeriesCount(value *int64)() {
    m.seriesCount = value
}
