// Code generated by smithy-go-codegen DO NOT EDIT.

// Package oam provides the API client, operations, and parameter types for
// CloudWatch Observability Access Manager.
//
// Use Amazon CloudWatch Observability Access Manager to create and manage links
// between source accounts and monitoring accounts by using CloudWatch
// cross-account observability. With CloudWatch cross-account observability, you
// can monitor and troubleshoot applications that span multiple accounts within a
// Region. Seamlessly search, visualize, and analyze your metrics, logs, traces,
// Application Signals services and service level objectives (SLOs), Application
// Insights applications, and internet monitors in any of the linked accounts
// without account boundaries.
//
// Set up one or more Amazon Web Services accounts as monitoring accounts and link
// them with multiple source accounts. A monitoring account is a central Amazon Web
// Services account that can view and interact with observability data generated
// from source accounts. A source account is an individual Amazon Web Services
// account that generates observability data for the resources that reside in it.
// Source accounts share their observability data with the monitoring account. The
// shared observability data can include metrics in Amazon CloudWatch, logs in
// Amazon CloudWatch Logs, traces in X-Ray, Application Signals services and
// service level objectives (SLOs), applications in Amazon CloudWatch Application
// Insights, and internet monitors in CloudWatch Internet Monitor.
//
// When you set up a link, you can choose to share the metrics from all namespaces
// with the monitoring account, or filter to a subset of namespaces. And for
// CloudWatch Logs, you can choose to share all log groups with the monitoring
// account, or filter to a subset of log groups.
package oam
