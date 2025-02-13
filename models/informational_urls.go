package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationalUrls 
type InformationalUrls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appSignUpUrl property
    appSignUpUrl *string
    // The singleSignOnDocumentationUrl property
    singleSignOnDocumentationUrl *string
}
// NewInformationalUrls instantiates a new informationalUrls and sets the default values.
func NewInformationalUrls()(*InformationalUrls) {
    m := &InformationalUrls{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInformationalUrlsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationalUrlsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationalUrls(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationalUrls) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppSignUpUrl gets the appSignUpUrl property value. The appSignUpUrl property
func (m *InformationalUrls) GetAppSignUpUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appSignUpUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationalUrls) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appSignUpUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSignUpUrl(val)
        }
        return nil
    }
    res["singleSignOnDocumentationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnDocumentationUrl(val)
        }
        return nil
    }
    return res
}
// GetSingleSignOnDocumentationUrl gets the singleSignOnDocumentationUrl property value. The singleSignOnDocumentationUrl property
func (m *InformationalUrls) GetSingleSignOnDocumentationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnDocumentationUrl
    }
}
// Serialize serializes information the current object
func (m *InformationalUrls) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appSignUpUrl", m.GetAppSignUpUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("singleSignOnDocumentationUrl", m.GetSingleSignOnDocumentationUrl())
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
func (m *InformationalUrls) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppSignUpUrl sets the appSignUpUrl property value. The appSignUpUrl property
func (m *InformationalUrls) SetAppSignUpUrl(value *string)() {
    if m != nil {
        m.appSignUpUrl = value
    }
}
// SetSingleSignOnDocumentationUrl sets the singleSignOnDocumentationUrl property value. The singleSignOnDocumentationUrl property
func (m *InformationalUrls) SetSingleSignOnDocumentationUrl(value *string)() {
    if m != nil {
        m.singleSignOnDocumentationUrl = value
    }
}
