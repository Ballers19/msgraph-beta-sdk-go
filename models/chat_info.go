package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatInfo 
type ChatInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier for a message in a Microsoft Teams channel.
    messageId *string
    // The ID of the reply message.
    replyChainMessageId *string
    // The unique identifier for a thread in Microsoft Teams.
    threadId *string
}
// NewChatInfo instantiates a new chatInfo and sets the default values.
func NewChatInfo()(*ChatInfo) {
    m := &ChatInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["replyChainMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplyChainMessageId(val)
        }
        return nil
    }
    res["threadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreadId(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the messageId property value. The unique identifier for a message in a Microsoft Teams channel.
func (m *ChatInfo) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
// GetReplyChainMessageId gets the replyChainMessageId property value. The ID of the reply message.
func (m *ChatInfo) GetReplyChainMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyChainMessageId
    }
}
// GetThreadId gets the threadId property value. The unique identifier for a thread in Microsoft Teams.
func (m *ChatInfo) GetThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threadId
    }
}
// Serialize serializes information the current object
func (m *ChatInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("replyChainMessageId", m.GetReplyChainMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("threadId", m.GetThreadId())
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
func (m *ChatInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessageId sets the messageId property value. The unique identifier for a message in a Microsoft Teams channel.
func (m *ChatInfo) SetMessageId(value *string)() {
    if m != nil {
        m.messageId = value
    }
}
// SetReplyChainMessageId sets the replyChainMessageId property value. The ID of the reply message.
func (m *ChatInfo) SetReplyChainMessageId(value *string)() {
    if m != nil {
        m.replyChainMessageId = value
    }
}
// SetThreadId sets the threadId property value. The unique identifier for a thread in Microsoft Teams.
func (m *ChatInfo) SetThreadId(value *string)() {
    if m != nil {
        m.threadId = value
    }
}
