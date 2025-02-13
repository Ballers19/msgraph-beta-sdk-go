package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i09fd74140caf51c542e318ca00111fee98f39a8fd68d7b344f57441e6823aaf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/allowedusers/item/ref"
)

// UserItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\allowedUsers\{user-id}
type UserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/shares/{printerShare%2Did}/allowedUsers/{user%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the ref property
func (m *UserItemRequestBuilder) Ref()(*i09fd74140caf51c542e318ca00111fee98f39a8fd68d7b344f57441e6823aaf4.RefRequestBuilder) {
    return i09fd74140caf51c542e318ca00111fee98f39a8fd68d7b344f57441e6823aaf4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
