package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedundancyDetectionSettings 
type RedundancyDetectionSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether email threading and near duplicate detection are enabled.
    isEnabled *bool
    // Specifies the maximum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
    maxWords *int32
    // Specifies the minimum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
    minWords *int32
    // Specifies the similarity level for documents to be put in the same near duplicate set. To learn more, see Document and email similarity threshold.
    similarityThreshold *int32
}
// NewRedundancyDetectionSettings instantiates a new redundancyDetectionSettings and sets the default values.
func NewRedundancyDetectionSettings()(*RedundancyDetectionSettings) {
    m := &RedundancyDetectionSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRedundancyDetectionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedundancyDetectionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedundancyDetectionSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedundancyDetectionSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedundancyDetectionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["maxWords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxWords(val)
        }
        return nil
    }
    res["minWords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinWords(val)
        }
        return nil
    }
    res["similarityThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimilarityThreshold(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
func (m *RedundancyDetectionSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetMaxWords gets the maxWords property value. Specifies the maximum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
func (m *RedundancyDetectionSettings) GetMaxWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxWords
    }
}
// GetMinWords gets the minWords property value. Specifies the minimum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
func (m *RedundancyDetectionSettings) GetMinWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minWords
    }
}
// GetSimilarityThreshold gets the similarityThreshold property value. Specifies the similarity level for documents to be put in the same near duplicate set. To learn more, see Document and email similarity threshold.
func (m *RedundancyDetectionSettings) GetSimilarityThreshold()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.similarityThreshold
    }
}
// Serialize serializes information the current object
func (m *RedundancyDetectionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxWords", m.GetMaxWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minWords", m.GetMinWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("similarityThreshold", m.GetSimilarityThreshold())
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
func (m *RedundancyDetectionSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
func (m *RedundancyDetectionSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetMaxWords sets the maxWords property value. Specifies the maximum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
func (m *RedundancyDetectionSettings) SetMaxWords(value *int32)() {
    if m != nil {
        m.maxWords = value
    }
}
// SetMinWords sets the minWords property value. Specifies the minimum number of words used for email threading and near duplicate detection. To learn more, see Minimum/maximum number of words.
func (m *RedundancyDetectionSettings) SetMinWords(value *int32)() {
    if m != nil {
        m.minWords = value
    }
}
// SetSimilarityThreshold sets the similarityThreshold property value. Specifies the similarity level for documents to be put in the same near duplicate set. To learn more, see Document and email similarity threshold.
func (m *RedundancyDetectionSettings) SetSimilarityThreshold(value *int32)() {
    if m != nil {
        m.similarityThreshold = value
    }
}
