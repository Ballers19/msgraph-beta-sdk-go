package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkConfiguredPeripheral 
type TeamworkConfiguredPeripheral struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
    isOptional *bool
    // The peripheral property
    peripheral TeamworkPeripheralable
}
// NewTeamworkConfiguredPeripheral instantiates a new teamworkConfiguredPeripheral and sets the default values.
func NewTeamworkConfiguredPeripheral()(*TeamworkConfiguredPeripheral) {
    m := &TeamworkConfiguredPeripheral{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkConfiguredPeripheralFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkConfiguredPeripheralFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkConfiguredPeripheral(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkConfiguredPeripheral) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConfiguredPeripheral) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["peripheral"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeripheral(val.(TeamworkPeripheralable))
        }
        return nil
    }
    return res
}
// GetIsOptional gets the isOptional property value. True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
func (m *TeamworkConfiguredPeripheral) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// GetPeripheral gets the peripheral property value. The peripheral property
func (m *TeamworkConfiguredPeripheral) GetPeripheral()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.peripheral
    }
}
// Serialize serializes information the current object
func (m *TeamworkConfiguredPeripheral) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("peripheral", m.GetPeripheral())
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
func (m *TeamworkConfiguredPeripheral) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsOptional sets the isOptional property value. True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
func (m *TeamworkConfiguredPeripheral) SetIsOptional(value *bool)() {
    if m != nil {
        m.isOptional = value
    }
}
// SetPeripheral sets the peripheral property value. The peripheral property
func (m *TeamworkConfiguredPeripheral) SetPeripheral(value TeamworkPeripheralable)() {
    if m != nil {
        m.peripheral = value
    }
}
