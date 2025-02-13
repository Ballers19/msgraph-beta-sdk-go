package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamSummary 
type TeamSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Count of guests in a team.
    guestsCount *int32
    // Count of members in a team.
    membersCount *int32
    // Count of owners in a team.
    ownersCount *int32
}
// NewTeamSummary instantiates a new teamSummary and sets the default values.
func NewTeamSummary()(*TeamSummary) {
    m := &TeamSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guestsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestsCount(val)
        }
        return nil
    }
    res["membersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCount(val)
        }
        return nil
    }
    res["ownersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnersCount(val)
        }
        return nil
    }
    return res
}
// GetGuestsCount gets the guestsCount property value. Count of guests in a team.
func (m *TeamSummary) GetGuestsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.guestsCount
    }
}
// GetMembersCount gets the membersCount property value. Count of members in a team.
func (m *TeamSummary) GetMembersCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.membersCount
    }
}
// GetOwnersCount gets the ownersCount property value. Count of owners in a team.
func (m *TeamSummary) GetOwnersCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ownersCount
    }
}
// Serialize serializes information the current object
func (m *TeamSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("guestsCount", m.GetGuestsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("membersCount", m.GetMembersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ownersCount", m.GetOwnersCount())
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
func (m *TeamSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGuestsCount sets the guestsCount property value. Count of guests in a team.
func (m *TeamSummary) SetGuestsCount(value *int32)() {
    if m != nil {
        m.guestsCount = value
    }
}
// SetMembersCount sets the membersCount property value. Count of members in a team.
func (m *TeamSummary) SetMembersCount(value *int32)() {
    if m != nil {
        m.membersCount = value
    }
}
// SetOwnersCount sets the ownersCount property value. Count of owners in a team.
func (m *TeamSummary) SetOwnersCount(value *int32)() {
    if m != nil {
        m.ownersCount = value
    }
}
