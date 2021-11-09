package copytodefaultcontentlocation

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type CopyToDefaultContentLocationRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    destinationFileName *string;
    // 
    sourceFile *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReference;
}
// Instantiates a new copyToDefaultContentLocationRequestBody and sets the default values.
func NewCopyToDefaultContentLocationRequestBody()(*CopyToDefaultContentLocationRequestBody) {
    m := &CopyToDefaultContentLocationRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyToDefaultContentLocationRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the destinationFileName property value. 
func (m *CopyToDefaultContentLocationRequestBody) GetDestinationFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationFileName
    }
}
// Gets the sourceFile property value. 
func (m *CopyToDefaultContentLocationRequestBody) GetSourceFile()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.sourceFile
    }
}
// The deserialization information for the current model
func (m *CopyToDefaultContentLocationRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewItemReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReference))
        }
        return nil
    }
    return res
}
func (m *CopyToDefaultContentLocationRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CopyToDefaultContentLocationRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CopyToDefaultContentLocationRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the destinationFileName property value. 
// Parameters:
//  - value : Value to set for the destinationFileName property.
func (m *CopyToDefaultContentLocationRequestBody) SetDestinationFileName(value *string)() {
    m.destinationFileName = value
}
// Sets the sourceFile property value. 
// Parameters:
//  - value : Value to set for the sourceFile property.
func (m *CopyToDefaultContentLocationRequestBody) SetSourceFile(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReference)() {
    m.sourceFile = value
}
