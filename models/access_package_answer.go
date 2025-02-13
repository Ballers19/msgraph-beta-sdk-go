package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAnswer 
type AccessPackageAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The question the answer is for. Required and Read-only.
    answeredQuestion AccessPackageQuestionable
    // The display value of the answer. Required.
    displayValue *string
    // The type property
    type_escaped *string
}
// NewAccessPackageAnswer instantiates a new accessPackageAnswer and sets the default values.
func NewAccessPackageAnswer()(*AccessPackageAnswer) {
    m := &AccessPackageAnswer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.accessPackageAnswer";
    m.SetType(&odatatypeValue);
    return m
}
// CreateAccessPackageAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.accessPackageAnswerString":
                        return NewAccessPackageAnswerString(), nil
                }
            }
        }
    }
    return NewAccessPackageAnswer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAnswer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAnsweredQuestion gets the answeredQuestion property value. The question the answer is for. Required and Read-only.
func (m *AccessPackageAnswer) GetAnsweredQuestion()(AccessPackageQuestionable) {
    if m == nil {
        return nil
    } else {
        return m.answeredQuestion
    }
}
// GetDisplayValue gets the displayValue property value. The display value of the answer. Required.
func (m *AccessPackageAnswer) GetDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["answeredQuestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnsweredQuestion(val.(AccessPackageQuestionable))
        }
        return nil
    }
    res["displayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayValue(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetType gets the @odata.type property value. The type property
func (m *AccessPackageAnswer) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("answeredQuestion", m.GetAnsweredQuestion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayValue", m.GetDisplayValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *AccessPackageAnswer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAnsweredQuestion sets the answeredQuestion property value. The question the answer is for. Required and Read-only.
func (m *AccessPackageAnswer) SetAnsweredQuestion(value AccessPackageQuestionable)() {
    if m != nil {
        m.answeredQuestion = value
    }
}
// SetDisplayValue sets the displayValue property value. The display value of the answer. Required.
func (m *AccessPackageAnswer) SetDisplayValue(value *string)() {
    if m != nil {
        m.displayValue = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *AccessPackageAnswer) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
