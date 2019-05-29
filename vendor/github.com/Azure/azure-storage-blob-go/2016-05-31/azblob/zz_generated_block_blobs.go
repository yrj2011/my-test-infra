package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// blockBlobsClient is the client for the BlockBlobs methods of the Azblob service.
type blockBlobsClient struct {
	managementClient
}

// newBlockBlobsClient creates an instance of the blockBlobsClient client.
func newBlockBlobsClient(url url.URL, p pipeline.Pipeline) blockBlobsClient {
	return blockBlobsClient{newManagementClient(url, p)}
}

// GetBlockList the Get Block List operation retrieves the list of blocks that have been uploaded as part of a block
// blob
//
// listType is specifies whether to return the list of committed blocks, the list of uncommitted blocks, or both lists
// together. snapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the blob
// snapshot to retrieve. For more information on working with blob snapshots, see <a
// href="http://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/creating-a-snapshot-of-a-blob">Creating
// a Snapshot of a Blob.</a> timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="http://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. requestID is provides a client-generated, opaque value with a 1 KB character
// limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client blockBlobsClient) GetBlockList(ctx context.Context, listType BlockListType, snapshot *time.Time, timeout *int32, leaseID *string, requestID *string) (*BlockList, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getBlockListPreparer(listType, snapshot, timeout, leaseID, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getBlockListResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockList), err
}

// getBlockListPreparer prepares the GetBlockList request.
func (client blockBlobsClient) getBlockListPreparer(listType BlockListType, snapshot *time.Time, timeout *int32, leaseID *string, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if snapshot != nil {
		params.Set("snapshot", (*snapshot).Format(rfc3339Format))
	}
	params.Set("blocklisttype", fmt.Sprintf("%v", listType))
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "blocklist")
	req.URL.RawQuery = params.Encode()
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// getBlockListResponder handles the response to the GetBlockList request.
func (client blockBlobsClient) getBlockListResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &BlockList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// PutBlock the Put Block operation creates a new block to be committed as part of a blob
//
// blockID is a valid Base64 string value that identifies the block. Prior to encoding, the string must be less than or
// equal to 64 bytes in size. For a given blob, the length of the value specified for the blockid parameter must be the
// same size for each block. body is initial data body will be closed upon successful return. Callers should ensure
// closure when receiving an error.timeout is the timeout parameter is expressed in seconds. For more information, see
// <a
// href="http://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. requestID is provides a client-generated, opaque value with a 1 KB character
// limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client blockBlobsClient) PutBlock(ctx context.Context, blockID string, body io.ReadSeeker, timeout *int32, leaseID *string, requestID *string) (*BlockBlobsPutBlockResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.putBlockPreparer(blockID, body, timeout, leaseID, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putBlockResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobsPutBlockResponse), err
}

// putBlockPreparer prepares the PutBlock request.
func (client blockBlobsClient) putBlockPreparer(blockID string, body io.ReadSeeker, timeout *int32, leaseID *string, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("blockid", blockID)
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "block")
	req.URL.RawQuery = params.Encode()
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// putBlockResponder handles the response to the PutBlock request.
func (client blockBlobsClient) putBlockResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	return &BlockBlobsPutBlockResponse{rawResponse: resp.Response()}, err
}

// PutBlockList the Put Block List operation writes a blob by specifying the list of block IDs that make up the blob.
// In order to be written as part of a blob, a block must have been successfully written to the server in a prior Put
// Block operation. You can call Put Block List to update a blob by uploading only those blocks that have changed, then
// committing the new and existing blocks together. You can do this by specifying whether to commit a block from the
// committed block list or from the uncommitted block list, or to commit the most recently uploaded version of the
// block, whichever list it may belong to.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="http://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> blobCacheControl is optional. Sets the blob's cache control. If specified,
// this property is stored with the blob and returned with a read request. blobContentType is optional. Sets the blob's
// content type. If specified, this property is stored with the blob and returned with a read request.
// blobContentEncoding is optional. Sets the blob's content encoding. If specified, this property is stored with the
// blob and returned with a read request. blobContentLanguage is optional. Set the blob's content language. If
// specified, this property is stored with the blob and returned with a read request. blobContentMD5 is optional. An
// MD5 hash of the blob content. Note that this hash is not validated, as the hashes for the individual blocks were
// validated when each was uploaded. metadata is optional. Specifies a user-defined name-value pair associated with the
// blob. If no name-value pairs are specified, the operation will copy the metadata from the source blob or file to the
// destination blob. If one or more name-value pairs are specified, the destination blob is created with the specified
// metadata, and metadata is not copied from the source blob or file. Note that beginning with version 2009-09-19,
// metadata names must adhere to the naming rules for C# identifiers. See Naming and Referencing Containers, Blobs, and
// Metadata for more information. leaseID is if specified, the operation only succeeds if the container's lease is
// active and matches this ID. blobContentDisposition is optional. Sets the blob's Content-Disposition header.
// ifModifiedSince is specify this header value to operate only on a blob if it has been modified since the specified
// date/time. ifUnmodifiedSince is specify this header value to operate only on a blob if it has not been modified
// since the specified date/time. ifMatches is specify an ETag value to operate only on blobs with a matching value.
// ifNoneMatch is specify an ETag value to operate only on blobs without a matching value. requestID is provides a
// client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage
// analytics logging is enabled.
func (client blockBlobsClient) PutBlockList(ctx context.Context, blocks BlockLookupList, timeout *int32, blobCacheControl *string, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (*BlockBlobsPutBlockListResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: metadata,
			constraints: []constraint{{target: "metadata", name: null, rule: false,
				chain: []constraint{{target: "metadata", name: pattern, rule: `^[a-zA-Z]+$`, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.putBlockListPreparer(blocks, timeout, blobCacheControl, blobContentType, blobContentEncoding, blobContentLanguage, blobContentMD5, metadata, leaseID, blobContentDisposition, ifModifiedSince, ifUnmodifiedSince, ifMatches, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putBlockListResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobsPutBlockListResponse), err
}

// putBlockListPreparer prepares the PutBlockList request.
func (client blockBlobsClient) putBlockListPreparer(blocks BlockLookupList, timeout *int32, blobCacheControl *string, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "blocklist")
	req.URL.RawQuery = params.Encode()
	if blobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobCacheControl)
	}
	if blobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobContentType)
	}
	if blobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobContentEncoding)
	}
	if blobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobContentLanguage)
	}
	if blobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", *blobContentMD5)
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if blobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobContentDisposition)
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatches != nil {
		req.Header.Set("If-Match", string(*ifMatches))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	b, err := xml.Marshal(blocks)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putBlockListResponder handles the response to the PutBlockList request.
func (client blockBlobsClient) putBlockListResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	return &BlockBlobsPutBlockListResponse{rawResponse: resp.Response()}, err
}
