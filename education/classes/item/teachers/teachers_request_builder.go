package teachers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0db9617499f9787a04f916beb72af7c5772ea91cc0cc7538b8e5a9bede4068c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/teachers/ref"
    i7405d2d738bd5ebf1bceb47127b6b5b893f3cec53fdbdfacbf09e62b40a7d5b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/teachers/delta"
)

type TeachersRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TeachersRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewTeachersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    m := &TeachersRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/classes/{educationClass_id}/teachers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewTeachersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeachersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TeachersRequestBuilder) CreateGetRequestInformation(q func (value *TeachersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TeachersRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
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
func (m *TeachersRequestBuilder) Delta()(*i7405d2d738bd5ebf1bceb47127b6b5b893f3cec53fdbdfacbf09e62b40a7d5b6.DeltaRequestBuilder) {
    return i7405d2d738bd5ebf1bceb47127b6b5b893f3cec53fdbdfacbf09e62b40a7d5b6.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeachersRequestBuilder) Get(q func (value *TeachersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*TeachersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeachersResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*TeachersResponse), nil
}
func (m *TeachersRequestBuilder) Ref()(*i0db9617499f9787a04f916beb72af7c5772ea91cc0cc7538b8e5a9bede4068c7.RefRequestBuilder) {
    return i0db9617499f9787a04f916beb72af7c5772ea91cc0cc7538b8e5a9bede4068c7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
