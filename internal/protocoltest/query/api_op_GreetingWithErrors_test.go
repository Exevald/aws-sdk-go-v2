// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GreetingWithErrors_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *GreetingWithErrorsOutput
	}{
		// Ensures that operations with errors successfully know how to deserialize the
		// successful response
		"QueryGreetingWithErrors": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<GreetingWithErrorsResponse xmlns="https://example.com/">
			    <GreetingWithErrorsResult>
			        <greeting>Hello</greeting>
			    </GreetingWithErrorsResult>
			</GreetingWithErrorsResponse>
			`),
			ExpectResult: &GreetingWithErrorsOutput{
				Greeting: ptr.String("Hello"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_InvalidGreeting_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.InvalidGreeting
	}{
		// Parses simple XML errors
		"QueryInvalidGreetingError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<ErrorResponse>
			   <Error>
			      <Type>Sender</Type>
			      <Code>InvalidGreeting</Code>
			      <Message>Hi</Message>
			   </Error>
			   <RequestId>foo-id</RequestId>
			</ErrorResponse>
			`),
			ExpectError: &types.InvalidGreeting{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.InvalidGreeting operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.InvalidGreeting
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.InvalidGreeting result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_ComplexError_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.ComplexError
	}{
		"QueryComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<ErrorResponse>
			   <Error>
			      <Type>Sender</Type>
			      <Code>ComplexError</Code>
			      <TopLevel>Top level</TopLevel>
			      <Nested>
			          <Foo>bar</Foo>
			      </Nested>
			   </Error>
			   <RequestId>foo-id</RequestId>
			</ErrorResponse>
			`),
			ExpectError: &types.ComplexError{
				TopLevel: ptr.String("Top level"),
				Nested: &types.ComplexNestedErrorData{
					Foo: ptr.String("bar"),
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.ComplexError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.ComplexError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.ComplexError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_CustomCodeError_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.CustomCodeError
	}{
		// Parses customized XML errors
		"QueryCustomizedError": {
			StatusCode: 402,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<ErrorResponse>
			   <Error>
			      <Type>Sender</Type>
			      <Code>Customized</Code>
			      <Message>Hi</Message>
			   </Error>
			   <RequestId>foo-id</RequestId>
			</ErrorResponse>
			`),
			ExpectError: &types.CustomCodeError{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.CustomCodeError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.CustomCodeError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.CustomCodeError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}
