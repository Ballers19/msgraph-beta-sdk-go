package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharingLink 
type SharingLink struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The app the link is associated with.
    application Identityable
    // The configuratorUrl property
    configuratorUrl *string
    // If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
    preventsDownload *bool
    // The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
    scope *string
    // The type of the link created.
    type_escaped *string
    // For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
    webHtml *string
    // A URL that opens the item in the browser on the OneDrive website.
    webUrl *string
}
// NewSharingLink instantiates a new sharingLink and sets the default values.
func NewSharingLink()(*SharingLink) {
    m := &SharingLink{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSharingLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharingLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharingLink(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingLink) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplication gets the application property value. The app the link is associated with.
func (m *SharingLink) GetApplication()(Identityable) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetConfiguratorUrl gets the configuratorUrl property value. The configuratorUrl property
func (m *SharingLink) GetConfiguratorUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configuratorUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharingLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Identityable))
        }
        return nil
    }
    res["configuratorUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguratorUrl(val)
        }
        return nil
    }
    res["preventsDownload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventsDownload(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["webHtml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebHtml(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetPreventsDownload gets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
func (m *SharingLink) GetPreventsDownload()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preventsDownload
    }
}
// GetScope gets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
func (m *SharingLink) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetType gets the type property value. The type of the link created.
func (m *SharingLink) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetWebHtml gets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
func (m *SharingLink) GetWebHtml()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webHtml
    }
}
// GetWebUrl gets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
func (m *SharingLink) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *SharingLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("configuratorUrl", m.GetConfiguratorUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preventsDownload", m.GetPreventsDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webHtml", m.GetWebHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *SharingLink) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplication sets the application property value. The app the link is associated with.
func (m *SharingLink) SetApplication(value Identityable)() {
    if m != nil {
        m.application = value
    }
}
// SetConfiguratorUrl sets the configuratorUrl property value. The configuratorUrl property
func (m *SharingLink) SetConfiguratorUrl(value *string)() {
    if m != nil {
        m.configuratorUrl = value
    }
}
// SetPreventsDownload sets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
func (m *SharingLink) SetPreventsDownload(value *bool)() {
    if m != nil {
        m.preventsDownload = value
    }
}
// SetScope sets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
func (m *SharingLink) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
// SetType sets the type property value. The type of the link created.
func (m *SharingLink) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetWebHtml sets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
func (m *SharingLink) SetWebHtml(value *string)() {
    if m != nil {
        m.webHtml = value
    }
}
// SetWebUrl sets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
func (m *SharingLink) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
