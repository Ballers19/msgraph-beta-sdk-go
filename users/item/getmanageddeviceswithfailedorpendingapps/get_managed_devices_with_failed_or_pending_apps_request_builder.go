package getmanageddeviceswithfailedorpendingapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetManagedDevicesWithFailedOrPendingAppsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    m := &GetManagedDevicesWithFailedOrPendingAppsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/microsoft.graph.getManagedDevicesWithFailedOrPendingApps()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) CreateGetRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) Get(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)([]GetManagedDevicesWithFailedOrPendingApps, error) {
    requestInfo, err := m.CreateGetRequestInformation(h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetManagedDevicesWithFailedOrPendingApps() }, responseHandler)
    if err != nil {
        return nil, err
    }
    val := make([]GetManagedDevicesWithFailedOrPendingApps, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetManagedDevicesWithFailedOrPendingApps))
    }
    return val, nil
}
