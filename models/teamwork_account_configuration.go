package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkAccountConfiguration 
type TeamworkAccountConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The account used to sync the calendar.
    onPremisesCalendarSyncConfiguration TeamworkOnPremisesCalendarSyncConfigurationable
    // The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams, teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
    supportedClient *TeamworkSupportedClient
}
// NewTeamworkAccountConfiguration instantiates a new teamworkAccountConfiguration and sets the default values.
func NewTeamworkAccountConfiguration()(*TeamworkAccountConfiguration) {
    m := &TeamworkAccountConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkAccountConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkAccountConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkAccountConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkAccountConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkAccountConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["onPremisesCalendarSyncConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkOnPremisesCalendarSyncConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesCalendarSyncConfiguration(val.(TeamworkOnPremisesCalendarSyncConfigurationable))
        }
        return nil
    }
    res["supportedClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkSupportedClient)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedClient(val.(*TeamworkSupportedClient))
        }
        return nil
    }
    return res
}
// GetOnPremisesCalendarSyncConfiguration gets the onPremisesCalendarSyncConfiguration property value. The account used to sync the calendar.
func (m *TeamworkAccountConfiguration) GetOnPremisesCalendarSyncConfiguration()(TeamworkOnPremisesCalendarSyncConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesCalendarSyncConfiguration
    }
}
// GetSupportedClient gets the supportedClient property value. The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams, teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
func (m *TeamworkAccountConfiguration) GetSupportedClient()(*TeamworkSupportedClient) {
    if m == nil {
        return nil
    } else {
        return m.supportedClient
    }
}
// Serialize serializes information the current object
func (m *TeamworkAccountConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("onPremisesCalendarSyncConfiguration", m.GetOnPremisesCalendarSyncConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedClient() != nil {
        cast := (*m.GetSupportedClient()).String()
        err := writer.WriteStringValue("supportedClient", &cast)
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
func (m *TeamworkAccountConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOnPremisesCalendarSyncConfiguration sets the onPremisesCalendarSyncConfiguration property value. The account used to sync the calendar.
func (m *TeamworkAccountConfiguration) SetOnPremisesCalendarSyncConfiguration(value TeamworkOnPremisesCalendarSyncConfigurationable)() {
    if m != nil {
        m.onPremisesCalendarSyncConfiguration = value
    }
}
// SetSupportedClient sets the supportedClient property value. The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams, teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
func (m *TeamworkAccountConfiguration) SetSupportedClient(value *TeamworkSupportedClient)() {
    if m != nil {
        m.supportedClient = value
    }
}
