package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i279637f29307e0ba29e5c84dd61c33cc691ddb4f87370139cbc03ce0aa2f627f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/members/item/ref"
)

// EducationUserItemRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}\members\{educationUser-id}
type EducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUserItemRequestBuilder) {
    m := &EducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/members/{educationUser%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the ref property
func (m *EducationUserItemRequestBuilder) Ref()(*i279637f29307e0ba29e5c84dd61c33cc691ddb4f87370139cbc03ce0aa2f627f.RefRequestBuilder) {
    return i279637f29307e0ba29e5c84dd61c33cc691ddb4f87370139cbc03ce0aa2f627f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
