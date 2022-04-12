package revokealllicenses

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeAllLicensesRequestBody provides operations to call the revokeAllLicenses method.
type RevokeAllLicensesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The notifyManagedDevices property
    notifyManagedDevices *bool
}
// NewRevokeAllLicensesRequestBody instantiates a new revokeAllLicensesRequestBody and sets the default values.
func NewRevokeAllLicensesRequestBody()(*RevokeAllLicensesRequestBody) {
    m := &RevokeAllLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRevokeAllLicensesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeAllLicensesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeAllLicensesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeAllLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeAllLicensesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notifyManagedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyManagedDevices(val)
        }
        return nil
    }
    return res
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeAllLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
// Serialize serializes information the current object
func (m *RevokeAllLicensesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
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
func (m *RevokeAllLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeAllLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    if m != nil {
        m.notifyManagedDevices = value
    }
}
