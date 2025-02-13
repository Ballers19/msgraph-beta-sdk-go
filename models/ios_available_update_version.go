package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosAvailableUpdateVersion iOS available update version details
type IosAvailableUpdateVersion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expiration date of the update.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The posting date of the update.
    postingDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version of the update.
    productVersion *string
    // List of supported devices for the update.
    supportedDevices []string
}
// NewIosAvailableUpdateVersion instantiates a new iosAvailableUpdateVersion and sets the default values.
func NewIosAvailableUpdateVersion()(*IosAvailableUpdateVersion) {
    m := &IosAvailableUpdateVersion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosAvailableUpdateVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosAvailableUpdateVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosAvailableUpdateVersion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosAvailableUpdateVersion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date of the update.
func (m *IosAvailableUpdateVersion) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosAvailableUpdateVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["postingDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostingDateTime(val)
        }
        return nil
    }
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    res["supportedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedDevices(res)
        }
        return nil
    }
    return res
}
// GetPostingDateTime gets the postingDateTime property value. The posting date of the update.
func (m *IosAvailableUpdateVersion) GetPostingDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.postingDateTime
    }
}
// GetProductVersion gets the productVersion property value. The version of the update.
func (m *IosAvailableUpdateVersion) GetProductVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productVersion
    }
}
// GetSupportedDevices gets the supportedDevices property value. List of supported devices for the update.
func (m *IosAvailableUpdateVersion) GetSupportedDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedDevices
    }
}
// Serialize serializes information the current object
func (m *IosAvailableUpdateVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("postingDateTime", m.GetPostingDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedDevices() != nil {
        err := writer.WriteCollectionOfStringValues("supportedDevices", m.GetSupportedDevices())
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
func (m *IosAvailableUpdateVersion) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date of the update.
func (m *IosAvailableUpdateVersion) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetPostingDateTime sets the postingDateTime property value. The posting date of the update.
func (m *IosAvailableUpdateVersion) SetPostingDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.postingDateTime = value
    }
}
// SetProductVersion sets the productVersion property value. The version of the update.
func (m *IosAvailableUpdateVersion) SetProductVersion(value *string)() {
    if m != nil {
        m.productVersion = value
    }
}
// SetSupportedDevices sets the supportedDevices property value. List of supported devices for the update.
func (m *IosAvailableUpdateVersion) SetSupportedDevices(value []string)() {
    if m != nil {
        m.supportedDevices = value
    }
}
