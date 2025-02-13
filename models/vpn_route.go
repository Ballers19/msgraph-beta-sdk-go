package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnRoute vPN Route definition.
type VpnRoute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Destination prefix (IPv4/v6 address).
    destinationPrefix *string
    // Prefix size. (1-32). Valid values 1 to 32
    prefixSize *int32
}
// NewVpnRoute instantiates a new vpnRoute and sets the default values.
func NewVpnRoute()(*VpnRoute) {
    m := &VpnRoute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVpnRouteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnRouteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnRoute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnRoute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDestinationPrefix gets the destinationPrefix property value. Destination prefix (IPv4/v6 address).
func (m *VpnRoute) GetDestinationPrefix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationPrefix
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnRoute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationPrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationPrefix(val)
        }
        return nil
    }
    res["prefixSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefixSize(val)
        }
        return nil
    }
    return res
}
// GetPrefixSize gets the prefixSize property value. Prefix size. (1-32). Valid values 1 to 32
func (m *VpnRoute) GetPrefixSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.prefixSize
    }
}
// Serialize serializes information the current object
func (m *VpnRoute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationPrefix", m.GetDestinationPrefix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("prefixSize", m.GetPrefixSize())
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
func (m *VpnRoute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDestinationPrefix sets the destinationPrefix property value. Destination prefix (IPv4/v6 address).
func (m *VpnRoute) SetDestinationPrefix(value *string)() {
    if m != nil {
        m.destinationPrefix = value
    }
}
// SetPrefixSize sets the prefixSize property value. Prefix size. (1-32). Valid values 1 to 32
func (m *VpnRoute) SetPrefixSize(value *int32)() {
    if m != nil {
        m.prefixSize = value
    }
}
