package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateUserBinding 
type X509CertificateUserBinding struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
    priority *int32
    // Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, email. Required.
    userProperty *string
    // The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name.
    x509CertificateField *string
}
// NewX509CertificateUserBinding instantiates a new x509CertificateUserBinding and sets the default values.
func NewX509CertificateUserBinding()(*X509CertificateUserBinding) {
    m := &X509CertificateUserBinding{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateX509CertificateUserBindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateUserBindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateUserBinding(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *X509CertificateUserBinding) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateUserBinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["userProperty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserProperty(val)
        }
        return nil
    }
    res["x509CertificateField"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateField(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
func (m *X509CertificateUserBinding) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetUserProperty gets the userProperty property value. Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, email. Required.
func (m *X509CertificateUserBinding) GetUserProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userProperty
    }
}
// GetX509CertificateField gets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name.
func (m *X509CertificateUserBinding) GetX509CertificateField()(*string) {
    if m == nil {
        return nil
    } else {
        return m.x509CertificateField
    }
}
// Serialize serializes information the current object
func (m *X509CertificateUserBinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userProperty", m.GetUserProperty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("x509CertificateField", m.GetX509CertificateField())
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
func (m *X509CertificateUserBinding) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPriority sets the priority property value. The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
func (m *X509CertificateUserBinding) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetUserProperty sets the userProperty property value. Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, email. Required.
func (m *X509CertificateUserBinding) SetUserProperty(value *string)() {
    if m != nil {
        m.userProperty = value
    }
}
// SetX509CertificateField sets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name.
func (m *X509CertificateUserBinding) SetX509CertificateField(value *string)() {
    if m != nil {
        m.x509CertificateField = value
    }
}
