// Code generated by go-swagger; DO NOT EDIT.

package trending

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new trending API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trending API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TrendingGetTrendingCategories Returns trending items for Bungie.net, collapsed into the first page of items per category. For pagination within a category, call GetTrendingCategory.
*/
func (a *Client) TrendingGetTrendingCategories(params *TrendingGetTrendingCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*TrendingGetTrendingCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrendingGetTrendingCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Trending.GetTrendingCategories",
		Method:             "GET",
		PathPattern:        "/Trending/Categories/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrendingGetTrendingCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TrendingGetTrendingCategoriesOK), nil

}

/*
TrendingGetTrendingCategory Returns paginated lists of trending items for a category.
*/
func (a *Client) TrendingGetTrendingCategory(params *TrendingGetTrendingCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*TrendingGetTrendingCategoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrendingGetTrendingCategoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Trending.GetTrendingCategory",
		Method:             "GET",
		PathPattern:        "/Trending/Categories/{categoryId}/{pageNumber}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrendingGetTrendingCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TrendingGetTrendingCategoryOK), nil

}

/*
TrendingGetTrendingEntryDetail Returns the detailed results for a specific trending entry. Note that trending entries are uniquely identified by a combination of *both* the TrendingEntryType *and* the identifier: the identifier alone is not guaranteed to be globally unique.
*/
func (a *Client) TrendingGetTrendingEntryDetail(params *TrendingGetTrendingEntryDetailParams, authInfo runtime.ClientAuthInfoWriter) (*TrendingGetTrendingEntryDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrendingGetTrendingEntryDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Trending.GetTrendingEntryDetail",
		Method:             "GET",
		PathPattern:        "/Trending/Details/{trendingEntryType}/{identifier}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrendingGetTrendingEntryDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TrendingGetTrendingEntryDetailOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
