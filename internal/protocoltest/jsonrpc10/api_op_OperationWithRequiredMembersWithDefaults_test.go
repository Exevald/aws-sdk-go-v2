// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_OperationWithRequiredMembersWithDefaults_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *OperationWithRequiredMembersWithDefaultsOutput
	}{
		// Client error corrects with default values when server fails to serialize
		// required values.
		"AwsJson10ClientErrorCorrectsWithDefaultValuesWhenServerFailsToSerializeRequiredValues": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectResult: &OperationWithRequiredMembersWithDefaultsOutput{
				RequiredString:    ptr.String("hi"),
				RequiredBoolean:   ptr.Bool(true),
				RequiredList:      []string{},
				RequiredTimestamp: ptr.Time(smithytime.ParseEpochSeconds(1)),
				RequiredBlob:      []byte("blob"),
				RequiredByte:      ptr.Int8(1),
				RequiredShort:     ptr.Int16(1),
				RequiredInteger:   ptr.Int32(10),
				RequiredLong:      ptr.Int64(100),
				RequiredFloat:     ptr.Float32(1.0),
				RequiredDouble:    ptr.Float64(1.0),
				RequiredMap:       map[string]string{},
				RequiredEnum:      types.RequiredEnum("FOO"),
				RequiredIntEnum:   1,
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "AwsJson10ClientErrorCorrectsWithDefaultValuesWhenServerFailsToSerializeRequiredValues" {
				t.Skip("disabled test aws.protocoltests.json10#JsonRpc10 aws.protocoltests.json10#OperationWithRequiredMembersWithDefaults")
			}

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
				Region: "us-west-2",
			})
			var params OperationWithRequiredMembersWithDefaultsInput
			result, err := client.OperationWithRequiredMembersWithDefaults(context.Background(), &params)
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
