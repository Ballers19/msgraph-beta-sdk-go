package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageQuestion 
type AccessPackageQuestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // ID of the question.
    id *string
    // Specifies whether the requestor is allowed to edit answers to questions.
    isAnswerEditable *bool
    // Whether the requestor is required to supply an answer or not.
    isRequired *bool
    // Relative position of this question when displaying a list of questions to the requestor.
    sequence *int32
    // The text of the question to show to the requestor.
    text AccessPackageLocalizedContentable
    // The type property
    type_escaped *string
}
// NewAccessPackageQuestion instantiates a new accessPackageQuestion and sets the default values.
func NewAccessPackageQuestion()(*AccessPackageQuestion) {
    m := &AccessPackageQuestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.accessPackageQuestion";
    m.SetType(&odatatypeValue);
    return m
}
// CreateAccessPackageQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.accessPackageMultipleChoiceQuestion":
                        return NewAccessPackageMultipleChoiceQuestion(), nil
                    case "#microsoft.graph.accessPackageTextInputQuestion":
                        return NewAccessPackageTextInputQuestion(), nil
                }
            }
        }
    }
    return NewAccessPackageQuestion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isAnswerEditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAnswerEditable(val)
        }
        return nil
    }
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["sequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageLocalizedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(AccessPackageLocalizedContentable))
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
// GetId gets the id property value. ID of the question.
func (m *AccessPackageQuestion) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIsAnswerEditable gets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) GetIsAnswerEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAnswerEditable
    }
}
// GetIsRequired gets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// GetSequence gets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// GetText gets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) GetText()(AccessPackageLocalizedContentable) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetType gets the @odata.type property value. The type property
func (m *AccessPackageQuestion) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AccessPackageQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAnswerEditable", m.GetIsAnswerEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("text", m.GetText())
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
func (m *AccessPackageQuestion) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. ID of the question.
func (m *AccessPackageQuestion) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetIsAnswerEditable sets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) SetIsAnswerEditable(value *bool)() {
    if m != nil {
        m.isAnswerEditable = value
    }
}
// SetIsRequired sets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) SetIsRequired(value *bool)() {
    if m != nil {
        m.isRequired = value
    }
}
// SetSequence sets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) SetSequence(value *int32)() {
    if m != nil {
        m.sequence = value
    }
}
// SetText sets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) SetText(value AccessPackageLocalizedContentable)() {
    if m != nil {
        m.text = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *AccessPackageQuestion) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
