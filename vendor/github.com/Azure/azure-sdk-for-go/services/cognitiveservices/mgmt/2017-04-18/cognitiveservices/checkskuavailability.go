package cognitiveservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CheckSkuAvailabilityClient is the cognitive Services Management Client
type CheckSkuAvailabilityClient struct {
	BaseClient
}

// NewCheckSkuAvailabilityClient creates an instance of the CheckSkuAvailabilityClient client.
func NewCheckSkuAvailabilityClient(subscriptionID string) CheckSkuAvailabilityClient {
	return NewCheckSkuAvailabilityClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCheckSkuAvailabilityClientWithBaseURI creates an instance of the CheckSkuAvailabilityClient client.
func NewCheckSkuAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckSkuAvailabilityClient {
	return CheckSkuAvailabilityClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List check available SKUs.
// Parameters:
// location - resource location.
// parameters - check SKU Availability POST body.
func (client CheckSkuAvailabilityClient) List(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (result CheckSkuAvailabilityResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CheckSkuAvailabilityClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Skus", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Kind", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.CheckSkuAvailabilityClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CheckSkuAvailabilityClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.CheckSkuAvailabilityClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CheckSkuAvailabilityClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client CheckSkuAvailabilityClient) ListPreparer(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CheckSkuAvailabilityClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CheckSkuAvailabilityClient) ListResponder(resp *http.Response) (result CheckSkuAvailabilityResultList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
