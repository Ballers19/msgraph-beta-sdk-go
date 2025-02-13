package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InviteParticipantsOperation 
type InviteParticipantsOperation struct {
    CommsOperation
    // The participants to invite.
    participants []InvitationParticipantInfoable
}
// NewInviteParticipantsOperation instantiates a new InviteParticipantsOperation and sets the default values.
func NewInviteParticipantsOperation()(*InviteParticipantsOperation) {
    m := &InviteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateInviteParticipantsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInviteParticipantsOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInviteParticipantsOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InviteParticipantsOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(InvitationParticipantInfoable)
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants to invite.
func (m *InviteParticipantsOperation) GetParticipants()([]InvitationParticipantInfoable) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// Serialize serializes information the current object
func (m *InviteParticipantsOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParticipants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParticipants sets the participants property value. The participants to invite.
func (m *InviteParticipantsOperation) SetParticipants(value []InvitationParticipantInfoable)() {
    if m != nil {
        m.participants = value
    }
}
