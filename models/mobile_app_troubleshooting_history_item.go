package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingHistoryItem history Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingHistoryItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Time when the history item occurred.
    occurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Object containing detailed information about the error and its remediation.
    troubleshootingErrorDetails DeviceManagementTroubleshootingErrorDetailsable
    // The type property
    type_escaped *string
}
// NewMobileAppTroubleshootingHistoryItem instantiates a new mobileAppTroubleshootingHistoryItem and sets the default values.
func NewMobileAppTroubleshootingHistoryItem()(*MobileAppTroubleshootingHistoryItem) {
    m := &MobileAppTroubleshootingHistoryItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.mobileAppTroubleshootingHistoryItem";
    m.SetType(&odatatypeValue);
    return m
}
// CreateMobileAppTroubleshootingHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.mobileAppTroubleshootingAppPolicyCreationHistory":
                        return NewMobileAppTroubleshootingAppPolicyCreationHistory(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingAppStateHistory":
                        return NewMobileAppTroubleshootingAppStateHistory(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingAppTargetHistory":
                        return NewMobileAppTroubleshootingAppTargetHistory(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingAppUpdateHistory":
                        return NewMobileAppTroubleshootingAppUpdateHistory(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingDeviceCheckinHistory":
                        return NewMobileAppTroubleshootingDeviceCheckinHistory(), nil
                }
            }
        }
    }
    return NewMobileAppTroubleshootingHistoryItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppTroubleshootingHistoryItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["occurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrenceDateTime(val)
        }
        return nil
    }
    res["troubleshootingErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingErrorDetails(val.(DeviceManagementTroubleshootingErrorDetailsable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetOccurrenceDateTime gets the occurrenceDateTime property value. Time when the history item occurred.
func (m *MobileAppTroubleshootingHistoryItem) GetOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.occurrenceDateTime
    }
}
// GetTroubleshootingErrorDetails gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *MobileAppTroubleshootingHistoryItem) GetTroubleshootingErrorDetails()(DeviceManagementTroubleshootingErrorDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingErrorDetails
    }
}
// GetType gets the @odata.type property value. The type property
func (m *MobileAppTroubleshootingHistoryItem) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("occurrenceDateTime", m.GetOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("troubleshootingErrorDetails", m.GetTroubleshootingErrorDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *MobileAppTroubleshootingHistoryItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOccurrenceDateTime sets the occurrenceDateTime property value. Time when the history item occurred.
func (m *MobileAppTroubleshootingHistoryItem) SetOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.occurrenceDateTime = value
    }
}
// SetTroubleshootingErrorDetails sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *MobileAppTroubleshootingHistoryItem) SetTroubleshootingErrorDetails(value DeviceManagementTroubleshootingErrorDetailsable)() {
    if m != nil {
        m.troubleshootingErrorDetails = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *MobileAppTroubleshootingHistoryItem) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
