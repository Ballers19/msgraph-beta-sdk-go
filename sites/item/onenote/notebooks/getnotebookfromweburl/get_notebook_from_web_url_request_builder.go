package getnotebookfromweburl

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetNotebookFromWebUrlRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GetNotebookFromWebUrlResponse struct {
    additionalData map[string]interface{};
    copyNotebookModel *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CopyNotebookModel;
}
func NewGetNotebookFromWebUrlResponse()(*GetNotebookFromWebUrlResponse) {
    m := &GetNotebookFromWebUrlResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetNotebookFromWebUrlResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetNotebookFromWebUrlResponse) GetCopyNotebookModel()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CopyNotebookModel) {
    if m == nil {
        return nil
    } else {
        return m.copyNotebookModel
    }
}
func (m *GetNotebookFromWebUrlResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["copyNotebookModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCopyNotebookModel() })
        if err != nil {
            return err
        }
        m.SetCopyNotebookModel(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CopyNotebookModel))
        return nil
    }
    return res
}
func (m *GetNotebookFromWebUrlResponse) IsNil()(bool) {
    return m == nil
}
func (m *GetNotebookFromWebUrlResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("copyNotebookModel", m.GetCopyNotebookModel())
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
func (m *GetNotebookFromWebUrlResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GetNotebookFromWebUrlResponse) SetCopyNotebookModel(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CopyNotebookModel)() {
    m.copyNotebookModel = value
}
func NewGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    m := &GetNotebookFromWebUrlRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/onenote/notebooks/microsoft.graph.getNotebookFromWebUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GetNotebookFromWebUrlRequestBuilder) CreatePostRequestInformation(body *GetNotebookFromWebUrlRequestBody, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *GetNotebookFromWebUrlRequestBuilder) Post(body *GetNotebookFromWebUrlRequestBody, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*GetNotebookFromWebUrlResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetNotebookFromWebUrlResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*GetNotebookFromWebUrlResponse), nil
}
