package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewNotificationRecipientItem 
type AccessReviewNotificationRecipientItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Determines the recipient of the notification email.
    notificationRecipientScope AccessReviewNotificationRecipientScopeable
    // Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients which sends review completion notifications to the recipients.
    notificationTemplateType *string
}
// NewAccessReviewNotificationRecipientItem instantiates a new accessReviewNotificationRecipientItem and sets the default values.
func NewAccessReviewNotificationRecipientItem()(*AccessReviewNotificationRecipientItem) {
    m := &AccessReviewNotificationRecipientItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewNotificationRecipientItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewNotificationRecipientItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewNotificationRecipientItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewNotificationRecipientItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewNotificationRecipientItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationRecipientScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewNotificationRecipientScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationRecipientScope(val.(AccessReviewNotificationRecipientScopeable))
        }
        return nil
    }
    res["notificationTemplateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTemplateType(val)
        }
        return nil
    }
    return res
}
// GetNotificationRecipientScope gets the notificationRecipientScope property value. Determines the recipient of the notification email.
func (m *AccessReviewNotificationRecipientItem) GetNotificationRecipientScope()(AccessReviewNotificationRecipientScopeable) {
    if m == nil {
        return nil
    } else {
        return m.notificationRecipientScope
    }
}
// GetNotificationTemplateType gets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients which sends review completion notifications to the recipients.
func (m *AccessReviewNotificationRecipientItem) GetNotificationTemplateType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateType
    }
}
// Serialize serializes information the current object
func (m *AccessReviewNotificationRecipientItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("notificationRecipientScope", m.GetNotificationRecipientScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTemplateType", m.GetNotificationTemplateType())
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
func (m *AccessReviewNotificationRecipientItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotificationRecipientScope sets the notificationRecipientScope property value. Determines the recipient of the notification email.
func (m *AccessReviewNotificationRecipientItem) SetNotificationRecipientScope(value AccessReviewNotificationRecipientScopeable)() {
    if m != nil {
        m.notificationRecipientScope = value
    }
}
// SetNotificationTemplateType sets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients which sends review completion notifications to the recipients.
func (m *AccessReviewNotificationRecipientItem) SetNotificationTemplateType(value *string)() {
    if m != nil {
        m.notificationTemplateType = value
    }
}
