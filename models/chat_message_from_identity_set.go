package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageFromIdentitySet 
type ChatMessageFromIdentitySet struct {
    IdentitySet
}
// NewChatMessageFromIdentitySet instantiates a new ChatMessageFromIdentitySet and sets the default values.
func NewChatMessageFromIdentitySet()(*ChatMessageFromIdentitySet) {
    m := &ChatMessageFromIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// CreateChatMessageFromIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageFromIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageFromIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageFromIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ChatMessageFromIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
