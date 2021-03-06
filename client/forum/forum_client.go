// Code generated by go-swagger; DO NOT EDIT.

package forum

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new forum API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for forum API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ForumApproveFireteamThread Allows the owner of a fireteam thread to approve all joined members and start a private message conversation with them.
*/
func (a *Client) ForumApproveFireteamThread(params *ForumApproveFireteamThreadParams, authInfo runtime.ClientAuthInfoWriter) (*ForumApproveFireteamThreadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumApproveFireteamThreadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.ApproveFireteamThread",
		Method:             "POST",
		PathPattern:        "/Forum/Recruit/Approve/{topicId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumApproveFireteamThreadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumApproveFireteamThreadOK), nil

}

/*
ForumGetCoreTopicsPaged Gets a listing of all topics marked as part of the core group.
*/
func (a *Client) ForumGetCoreTopicsPaged(params *ForumGetCoreTopicsPagedParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetCoreTopicsPagedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetCoreTopicsPagedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetCoreTopicsPaged",
		Method:             "GET",
		PathPattern:        "/Forum/GetCoreTopicsPaged/{page}/{sort}/{quickDate}/{categoryFilter}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetCoreTopicsPagedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetCoreTopicsPagedOK), nil

}

/*
ForumGetForumTagSuggestions Gets tag suggestions based on partial text entry, matching them with other tags previously used in the forums.
*/
func (a *Client) ForumGetForumTagSuggestions(params *ForumGetForumTagSuggestionsParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetForumTagSuggestionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetForumTagSuggestionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetForumTagSuggestions",
		Method:             "GET",
		PathPattern:        "/Forum/GetForumTagSuggestions/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetForumTagSuggestionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetForumTagSuggestionsOK), nil

}

/*
ForumGetPoll Gets the specified forum poll.
*/
func (a *Client) ForumGetPoll(params *ForumGetPollParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetPollOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetPollParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetPoll",
		Method:             "GET",
		PathPattern:        "/Forum/Poll/{topicId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetPollReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetPollOK), nil

}

/*
ForumGetPostAndParent Returns the post specified and its immediate parent.
*/
func (a *Client) ForumGetPostAndParent(params *ForumGetPostAndParentParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetPostAndParentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetPostAndParentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetPostAndParent",
		Method:             "GET",
		PathPattern:        "/Forum/GetPostAndParent/{childPostId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetPostAndParentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetPostAndParentOK), nil

}

/*
ForumGetPostAndParentAwaitingApproval Returns the post specified and its immediate parent of posts that are awaiting approval.
*/
func (a *Client) ForumGetPostAndParentAwaitingApproval(params *ForumGetPostAndParentAwaitingApprovalParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetPostAndParentAwaitingApprovalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetPostAndParentAwaitingApprovalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetPostAndParentAwaitingApproval",
		Method:             "GET",
		PathPattern:        "/Forum/GetPostAndParentAwaitingApproval/{childPostId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetPostAndParentAwaitingApprovalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetPostAndParentAwaitingApprovalOK), nil

}

/*
ForumGetPostsThreadedPaged Returns a thread of posts at the given parent, optionally returning replies to those posts as well as the original parent.
*/
func (a *Client) ForumGetPostsThreadedPaged(params *ForumGetPostsThreadedPagedParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetPostsThreadedPagedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetPostsThreadedPagedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetPostsThreadedPaged",
		Method:             "GET",
		PathPattern:        "/Forum/GetPostsThreadedPaged/{parentPostId}/{page}/{pageSize}/{replySize}/{getParentPost}/{rootThreadMode}/{sortMode}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetPostsThreadedPagedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetPostsThreadedPagedOK), nil

}

/*
ForumGetPostsThreadedPagedFromChild Returns a thread of posts starting at the topicId of the input childPostId, optionally returning replies to those posts as well as the original parent.
*/
func (a *Client) ForumGetPostsThreadedPagedFromChild(params *ForumGetPostsThreadedPagedFromChildParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetPostsThreadedPagedFromChildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetPostsThreadedPagedFromChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetPostsThreadedPagedFromChild",
		Method:             "GET",
		PathPattern:        "/Forum/GetPostsThreadedPagedFromChild/{childPostId}/{page}/{pageSize}/{replySize}/{rootThreadMode}/{sortMode}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetPostsThreadedPagedFromChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetPostsThreadedPagedFromChildOK), nil

}

/*
ForumGetRecruitmentThreadSummaries Allows the caller to get a list of to 25 recruitment thread summary information objects.
*/
func (a *Client) ForumGetRecruitmentThreadSummaries(params *ForumGetRecruitmentThreadSummariesParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetRecruitmentThreadSummariesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetRecruitmentThreadSummariesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetRecruitmentThreadSummaries",
		Method:             "POST",
		PathPattern:        "/Forum/Recruit/Summaries/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetRecruitmentThreadSummariesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetRecruitmentThreadSummariesOK), nil

}

/*
ForumGetTopicForContent Gets the post Id for the given content item's comments, if it exists.
*/
func (a *Client) ForumGetTopicForContent(params *ForumGetTopicForContentParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetTopicForContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetTopicForContentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetTopicForContent",
		Method:             "GET",
		PathPattern:        "/Forum/GetTopicForContent/{contentId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetTopicForContentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetTopicForContentOK), nil

}

/*
ForumGetTopicsPaged Get topics from any forum.
*/
func (a *Client) ForumGetTopicsPaged(params *ForumGetTopicsPagedParams, authInfo runtime.ClientAuthInfoWriter) (*ForumGetTopicsPagedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumGetTopicsPagedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.GetTopicsPaged",
		Method:             "GET",
		PathPattern:        "/Forum/GetTopicsPaged/{page}/{pageSize}/{group}/{sort}/{quickDate}/{categoryFilter}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumGetTopicsPagedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumGetTopicsPagedOK), nil

}

/*
ForumJoinFireteamThread Allows a user to slot themselves into a recruitment thread fireteam slot. Returns the new state of the fireteam.
*/
func (a *Client) ForumJoinFireteamThread(params *ForumJoinFireteamThreadParams, authInfo runtime.ClientAuthInfoWriter) (*ForumJoinFireteamThreadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumJoinFireteamThreadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.JoinFireteamThread",
		Method:             "POST",
		PathPattern:        "/Forum/Recruit/Join/{topicId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumJoinFireteamThreadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumJoinFireteamThreadOK), nil

}

/*
ForumKickBanFireteamApplicant Allows a recruitment thread owner to kick a join user from the fireteam. Returns the new state of the fireteam.
*/
func (a *Client) ForumKickBanFireteamApplicant(params *ForumKickBanFireteamApplicantParams, authInfo runtime.ClientAuthInfoWriter) (*ForumKickBanFireteamApplicantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumKickBanFireteamApplicantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.KickBanFireteamApplicant",
		Method:             "POST",
		PathPattern:        "/Forum/Recruit/KickBan/{topicId}/{targetMembershipId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumKickBanFireteamApplicantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumKickBanFireteamApplicantOK), nil

}

/*
ForumLeaveFireteamThread Allows a user to remove themselves from a recruitment thread fireteam slot. Returns the new state of the fireteam.
*/
func (a *Client) ForumLeaveFireteamThread(params *ForumLeaveFireteamThreadParams, authInfo runtime.ClientAuthInfoWriter) (*ForumLeaveFireteamThreadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForumLeaveFireteamThreadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Forum.LeaveFireteamThread",
		Method:             "POST",
		PathPattern:        "/Forum/Recruit/Leave/{topicId}/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForumLeaveFireteamThreadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForumLeaveFireteamThreadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
