package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOnPremisesConnectionHealthCheck 
type CloudPcOnPremisesConnectionHealthCheck struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Additional details about the health check or the recommended action.
    additionalDetails *string
    // The display name for this health check item.
    displayName *string
    // The end time of the health check item. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of error that occurred during this health check.
    errorType *CloudPcOnPremisesConnectionHealthCheckErrorType
    // The recommended action to fix the corresponding error.
    recommendedAction *string
    // The start time of the health check item. Read-only.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *CloudPcOnPremisesConnectionStatus
}
// NewCloudPcOnPremisesConnectionHealthCheck instantiates a new cloudPcOnPremisesConnectionHealthCheck and sets the default values.
func NewCloudPcOnPremisesConnectionHealthCheck()(*CloudPcOnPremisesConnectionHealthCheck) {
    m := &CloudPcOnPremisesConnectionHealthCheck{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcOnPremisesConnectionHealthCheckFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOnPremisesConnectionHealthCheckFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOnPremisesConnectionHealthCheck(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalDetails gets the additionalDetails property value. Additional details about the health check or the recommended action.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// GetDisplayName gets the displayName property value. The display name for this health check item.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndDateTime gets the endDateTime property value. The end time of the health check item. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetErrorType gets the errorType property value. The type of error that occurred during this health check.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetErrorType()(*CloudPcOnPremisesConnectionHealthCheckErrorType) {
    if m == nil {
        return nil
    } else {
        return m.errorType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOnPremisesConnectionHealthCheck) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["errorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionHealthCheckErrorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorType(val.(*CloudPcOnPremisesConnectionHealthCheckErrorType))
        }
        return nil
    }
    res["recommendedAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcOnPremisesConnectionStatus))
        }
        return nil
    }
    return res
}
// GetRecommendedAction gets the recommendedAction property value. The recommended action to fix the corresponding error.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
// GetStartDateTime gets the startDateTime property value. The start time of the health check item. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetStatus gets the status property value. The status property
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStatus()(*CloudPcOnPremisesConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *CloudPcOnPremisesConnectionHealthCheck) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetErrorType() != nil {
        cast := (*m.GetErrorType()).String()
        err := writer.WriteStringValue("errorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalDetails sets the additionalDetails property value. Additional details about the health check or the recommended action.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalDetails(value *string)() {
    if m != nil {
        m.additionalDetails = value
    }
}
// SetDisplayName sets the displayName property value. The display name for this health check item.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndDateTime sets the endDateTime property value. The end time of the health check item. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetErrorType sets the errorType property value. The type of error that occurred during this health check.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetErrorType(value *CloudPcOnPremisesConnectionHealthCheckErrorType)() {
    if m != nil {
        m.errorType = value
    }
}
// SetRecommendedAction sets the recommendedAction property value. The recommended action to fix the corresponding error.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetRecommendedAction(value *string)() {
    if m != nil {
        m.recommendedAction = value
    }
}
// SetStartDateTime sets the startDateTime property value. The start time of the health check item. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetStatus sets the status property value. The status property
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStatus(value *CloudPcOnPremisesConnectionStatus)() {
    if m != nil {
        m.status = value
    }
}
