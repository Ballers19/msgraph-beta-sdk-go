package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalculatedColumn 
type CalculatedColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
    format *string
    // The formula used to compute the value for this column.
    formula *string
    // The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
    outputType *string
}
// NewCalculatedColumn instantiates a new calculatedColumn and sets the default values.
func NewCalculatedColumn()(*CalculatedColumn) {
    m := &CalculatedColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCalculatedColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalculatedColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalculatedColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalculatedColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalculatedColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["formula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormula(val)
        }
        return nil
    }
    res["outputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputType(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
func (m *CalculatedColumn) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetFormula gets the formula property value. The formula used to compute the value for this column.
func (m *CalculatedColumn) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
// GetOutputType gets the outputType property value. The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
func (m *CalculatedColumn) GetOutputType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputType
    }
}
// Serialize serializes information the current object
func (m *CalculatedColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputType", m.GetOutputType())
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
func (m *CalculatedColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFormat sets the format property value. For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
func (m *CalculatedColumn) SetFormat(value *string)() {
    if m != nil {
        m.format = value
    }
}
// SetFormula sets the formula property value. The formula used to compute the value for this column.
func (m *CalculatedColumn) SetFormula(value *string)() {
    if m != nil {
        m.formula = value
    }
}
// SetOutputType sets the outputType property value. The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
func (m *CalculatedColumn) SetOutputType(value *string)() {
    if m != nil {
        m.outputType = value
    }
}
