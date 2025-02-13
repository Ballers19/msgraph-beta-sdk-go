package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTeamsClientConfiguration 
type TeamworkTeamsClientConfiguration struct {
    // The configuration of the Microsoft Teams client user account for a device.
    accountConfiguration TeamworkAccountConfigurationable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The configuration of Microsoft Teams client features for a device.
    featuresConfiguration TeamworkFeaturesConfigurationable
}
// NewTeamworkTeamsClientConfiguration instantiates a new teamworkTeamsClientConfiguration and sets the default values.
func NewTeamworkTeamsClientConfiguration()(*TeamworkTeamsClientConfiguration) {
    m := &TeamworkTeamsClientConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamsClientConfiguration(), nil
}
// GetAccountConfiguration gets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) GetAccountConfiguration()(TeamworkAccountConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.accountConfiguration
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeaturesConfiguration gets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) GetFeaturesConfiguration()(TeamworkFeaturesConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.featuresConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamsClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkAccountConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountConfiguration(val.(TeamworkAccountConfigurationable))
        }
        return nil
    }
    res["featuresConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkFeaturesConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeaturesConfiguration(val.(TeamworkFeaturesConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamworkTeamsClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accountConfiguration", m.GetAccountConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("featuresConfiguration", m.GetFeaturesConfiguration())
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
// SetAccountConfiguration sets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) SetAccountConfiguration(value TeamworkAccountConfigurationable)() {
    if m != nil {
        m.accountConfiguration = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeaturesConfiguration sets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) SetFeaturesConfiguration(value TeamworkFeaturesConfigurationable)() {
    if m != nil {
        m.featuresConfiguration = value
    }
}
