package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallParticipantInfo 
type CallParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identity of the call participant.
    participant IdentitySetable
}
// NewCallParticipantInfo instantiates a new callParticipantInfo and sets the default values.
func NewCallParticipantInfo()(*CallParticipantInfo) {
    m := &CallParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallParticipantInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallParticipantInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallParticipantInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallParticipantInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["participant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipant(val.(IdentitySetable))
        }
        return nil
    }
    return res
}
// GetParticipant gets the participant property value. Identity of the call participant.
func (m *CallParticipantInfo) GetParticipant()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.participant
    }
}
// Serialize serializes information the current object
func (m *CallParticipantInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("participant", m.GetParticipant())
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
func (m *CallParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetParticipant sets the participant property value. Identity of the call participant.
func (m *CallParticipantInfo) SetParticipant(value IdentitySetable)() {
    if m != nil {
        m.participant = value
    }
}
