package match

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type MatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    lookupArray *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
    // 
    lookupValue *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
    // 
    matchType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
}
// Instantiates a new matchRequestBody and sets the default values.
func NewMatchRequestBody()(*MatchRequestBody) {
    m := &MatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the lookupArray property value. 
func (m *MatchRequestBody) GetLookupArray()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.lookupArray
    }
}
// Gets the lookupValue property value. 
func (m *MatchRequestBody) GetLookupValue()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.lookupValue
    }
}
// Gets the matchType property value. 
func (m *MatchRequestBody) GetMatchType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.matchType
    }
}
// The deserialization information for the current model
func (m *MatchRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["lookupArray"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLookupArray(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    res["lookupValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLookupValue(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    res["matchType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    return res
}
func (m *MatchRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MatchRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("lookupArray", m.GetLookupArray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lookupValue", m.GetLookupValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("matchType", m.GetMatchType())
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
func (m *MatchRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the lookupArray property value. 
// Parameters:
//  - value : Value to set for the lookupArray property.
func (m *MatchRequestBody) SetLookupArray(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.lookupArray = value
}
// Sets the lookupValue property value. 
// Parameters:
//  - value : Value to set for the lookupValue property.
func (m *MatchRequestBody) SetLookupValue(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.lookupValue = value
}
// Sets the matchType property value. 
// Parameters:
//  - value : Value to set for the matchType property.
func (m *MatchRequestBody) SetMatchType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.matchType = value
}
