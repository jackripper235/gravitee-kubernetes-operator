//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (C) 2015 The Gravitee team (http://gravitee.io)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package model

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Api) DeepCopyInto(out *Api) {
	*out = *in
	if in.DefinitionContext != nil {
		in, out := &in.DefinitionContext, &out.DefinitionContext
		*out = new(DefinitionContext)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		(*in).DeepCopyInto(*out)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(Services)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*Resource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Resource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]Flow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]*Property, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Property)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PathMappings != nil {
		in, out := &in.PathMappings, &out.PathMappings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResponseTemplates != nil {
		in, out := &in.ResponseTemplates, &out.ResponseTemplates
		*out = make(map[string]map[string]*ResponseTemplate, len(*in))
		for key, val := range *in {
			var outVal map[string]*ResponseTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]*ResponseTemplate, len(*in))
				for key, val := range *in {
					var outVal *ResponseTemplate
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(ResponseTemplate)
						(*in).DeepCopyInto(*out)
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make([]*Plan, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Plan)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]*Metadata, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Metadata)
				**out = **in
			}
		}
	}
	if in.PrimaryOwner != nil {
		in, out := &in.PrimaryOwner, &out.PrimaryOwner
		*out = new(Member)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Api.
func (in *Api) DeepCopy() *Api {
	if in == nil {
		return nil
	}
	out := new(Api)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(BasicAuth)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consumer) DeepCopyInto(out *Consumer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consumer.
func (in *Consumer) DeepCopy() *Consumer {
	if in == nil {
		return nil
	}
	out := new(Consumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(Auth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContextRef) DeepCopyInto(out *ContextRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContextRef.
func (in *ContextRef) DeepCopy() *ContextRef {
	if in == nil {
		return nil
	}
	out := new(ContextRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cors) DeepCopyInto(out *Cors) {
	*out = *in
	if in.AccessControlAllowOrigin != nil {
		in, out := &in.AccessControlAllowOrigin, &out.AccessControlAllowOrigin
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlExposeHeaders != nil {
		in, out := &in.AccessControlExposeHeaders, &out.AccessControlExposeHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowMethods != nil {
		in, out := &in.AccessControlAllowMethods, &out.AccessControlAllowMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowHeaders != nil {
		in, out := &in.AccessControlAllowHeaders, &out.AccessControlAllowHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cors.
func (in *Cors) DeepCopy() *Cors {
	if in == nil {
		return nil
	}
	out := new(Cors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionContext) DeepCopyInto(out *DefinitionContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionContext.
func (in *DefinitionContext) DeepCopy() *DefinitionContext {
	if in == nil {
		return nil
	}
	out := new(DefinitionContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicPropertyService) DeepCopyInto(out *DynamicPropertyService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicPropertyService.
func (in *DynamicPropertyService) DeepCopy() *DynamicPropertyService {
	if in == nil {
		return nil
	}
	out := new(DynamicPropertyService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointDiscoveryService) DeepCopyInto(out *EndpointDiscoveryService) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointDiscoveryService.
func (in *EndpointDiscoveryService) DeepCopy() *EndpointDiscoveryService {
	if in == nil {
		return nil
	}
	out := new(EndpointDiscoveryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroup) DeepCopyInto(out *EndpointGroup) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*HttpEndpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HttpEndpoint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.LoadBalancer = in.LoadBalancer
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(Services)
		(*in).DeepCopyInto(*out)
	}
	if in.HttpProxy != nil {
		in, out := &in.HttpProxy, &out.HttpProxy
		*out = new(HttpProxy)
		**out = **in
	}
	if in.HttpClientOptions != nil {
		in, out := &in.HttpClientOptions, &out.HttpClientOptions
		*out = new(HttpClientOptions)
		**out = **in
	}
	if in.HttpClientSslOptions != nil {
		in, out := &in.HttpClientSslOptions, &out.HttpClientSslOptions
		*out = new(HttpClientSslOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroup.
func (in *EndpointGroup) DeepCopy() *EndpointGroup {
	if in == nil {
		return nil
	}
	out := new(EndpointGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointHealthCheckService) DeepCopyInto(out *EndpointHealthCheckService) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointHealthCheckService.
func (in *EndpointHealthCheckService) DeepCopy() *EndpointHealthCheckService {
	if in == nil {
		return nil
	}
	out := new(EndpointHealthCheckService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Failover) DeepCopyInto(out *Failover) {
	*out = *in
	if in.Cases != nil {
		in, out := &in.Cases, &out.Cases
		*out = make([]FailoverCase, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Failover.
func (in *Failover) DeepCopy() *Failover {
	if in == nil {
		return nil
	}
	out := new(Failover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	if in.PathOperator != nil {
		in, out := &in.PathOperator, &out.PathOperator
		*out = new(PathOperator)
		**out = **in
	}
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = make([]FlowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = make([]FlowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]HttpMethod, len(*in))
		copy(*out, *in)
	}
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]Consumer, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowStep) DeepCopyInto(out *FlowStep) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowStep.
func (in *FlowStep) DeepCopy() *FlowStep {
	if in == nil {
		return nil
	}
	out := new(FlowStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericStringMap.
func (in *GenericStringMap) DeepCopy() *GenericStringMap {
	if in == nil {
		return nil
	}
	out := new(GenericStringMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckService) DeepCopyInto(out *HealthCheckService) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]*Step, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Step)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckService.
func (in *HealthCheckService) DeepCopy() *HealthCheckService {
	if in == nil {
		return nil
	}
	out := new(HealthCheckService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpClientOptions) DeepCopyInto(out *HttpClientOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpClientOptions.
func (in *HttpClientOptions) DeepCopy() *HttpClientOptions {
	if in == nil {
		return nil
	}
	out := new(HttpClientOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpClientSslOptions) DeepCopyInto(out *HttpClientSslOptions) {
	*out = *in
	if in.TrustStore != nil {
		in, out := &in.TrustStore, &out.TrustStore
		*out = new(TrustStore)
		**out = **in
	}
	if in.KeyStore != nil {
		in, out := &in.KeyStore, &out.KeyStore
		*out = new(KeyStore)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpClientSslOptions.
func (in *HttpClientSslOptions) DeepCopy() *HttpClientSslOptions {
	if in == nil {
		return nil
	}
	out := new(HttpClientSslOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpEndpoint) DeepCopyInto(out *HttpEndpoint) {
	*out = *in
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HttpProxy != nil {
		in, out := &in.HttpProxy, &out.HttpProxy
		*out = new(HttpProxy)
		**out = **in
	}
	if in.HttpClientOptions != nil {
		in, out := &in.HttpClientOptions, &out.HttpClientOptions
		*out = new(HttpClientOptions)
		**out = **in
	}
	if in.HttpClientSslOptions != nil {
		in, out := &in.HttpClientSslOptions, &out.HttpClientSslOptions
		*out = new(HttpClientSslOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HttpHeader, len(*in))
		copy(*out, *in)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(EndpointHealthCheckService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpEndpoint.
func (in *HttpEndpoint) DeepCopy() *HttpEndpoint {
	if in == nil {
		return nil
	}
	out := new(HttpEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpHeader) DeepCopyInto(out *HttpHeader) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpHeader.
func (in *HttpHeader) DeepCopy() *HttpHeader {
	if in == nil {
		return nil
	}
	out := new(HttpHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpProxy) DeepCopyInto(out *HttpProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpProxy.
func (in *HttpProxy) DeepCopy() *HttpProxy {
	if in == nil {
		return nil
	}
	out := new(HttpProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JKSKeyStore) DeepCopyInto(out *JKSKeyStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JKSKeyStore.
func (in *JKSKeyStore) DeepCopy() *JKSKeyStore {
	if in == nil {
		return nil
	}
	out := new(JKSKeyStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JKSTrustStore) DeepCopyInto(out *JKSTrustStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JKSTrustStore.
func (in *JKSTrustStore) DeepCopy() *JKSTrustStore {
	if in == nil {
		return nil
	}
	out := new(JKSTrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyStore) DeepCopyInto(out *KeyStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyStore.
func (in *KeyStore) DeepCopy() *KeyStore {
	if in == nil {
		return nil
	}
	out := new(KeyStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	out.Mode = in.Mode
	out.Scope = in.Scope
	out.Content = in.Content
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingContent) DeepCopyInto(out *LoggingContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingContent.
func (in *LoggingContent) DeepCopy() *LoggingContent {
	if in == nil {
		return nil
	}
	out := new(LoggingContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingMode) DeepCopyInto(out *LoggingMode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingMode.
func (in *LoggingMode) DeepCopy() *LoggingMode {
	if in == nil {
		return nil
	}
	out := new(LoggingMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingScope) DeepCopyInto(out *LoggingScope) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingScope.
func (in *LoggingScope) DeepCopy() *LoggingScope {
	if in == nil {
		return nil
	}
	out := new(LoggingScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Member) DeepCopyInto(out *Member) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Member.
func (in *Member) DeepCopy() *Member {
	if in == nil {
		return nil
	}
	out := new(Member)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PEMKeyStore) DeepCopyInto(out *PEMKeyStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PEMKeyStore.
func (in *PEMKeyStore) DeepCopy() *PEMKeyStore {
	if in == nil {
		return nil
	}
	out := new(PEMKeyStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PEMTrustStore) DeepCopyInto(out *PEMTrustStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PEMTrustStore.
func (in *PEMTrustStore) DeepCopy() *PEMTrustStore {
	if in == nil {
		return nil
	}
	out := new(PEMTrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKCS12KeyStore) DeepCopyInto(out *PKCS12KeyStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKCS12KeyStore.
func (in *PKCS12KeyStore) DeepCopy() *PKCS12KeyStore {
	if in == nil {
		return nil
	}
	out := new(PKCS12KeyStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKCS12TrustStore) DeepCopyInto(out *PKCS12TrustStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKCS12TrustStore.
func (in *PKCS12TrustStore) DeepCopy() *PKCS12TrustStore {
	if in == nil {
		return nil
	}
	out := new(PKCS12TrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Path) DeepCopyInto(out *Path) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*Rule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Rule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Path.
func (in *Path) DeepCopy() *Path {
	if in == nil {
		return nil
	}
	out := new(Path)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PathOperator) DeepCopyInto(out *PathOperator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathOperator.
func (in *PathOperator) DeepCopy() *PathOperator {
	if in == nil {
		return nil
	}
	out := new(PathOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plan) DeepCopyInto(out *Plan) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make(map[string][]Rule, len(*in))
		for key, val := range *in {
			var outVal []Rule
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]Rule, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]Flow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Characteristics != nil {
		in, out := &in.Characteristics, &out.Characteristics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in *Plan) DeepCopy() *Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginReference) DeepCopyInto(out *PluginReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginReference.
func (in *PluginReference) DeepCopy() *PluginReference {
	if in == nil {
		return nil
	}
	out := new(PluginReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginRevision) DeepCopyInto(out *PluginRevision) {
	*out = *in
	if in.PluginReference != nil {
		in, out := &in.PluginReference, &out.PluginReference
		*out = new(PluginReference)
		**out = **in
	}
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		*out = new(Plugin)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginRevision.
func (in *PluginRevision) DeepCopy() *PluginRevision {
	if in == nil {
		return nil
	}
	out := new(PluginRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
	if in.VirtualHosts != nil {
		in, out := &in.VirtualHosts, &out.VirtualHosts
		*out = make([]*VirtualHost, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VirtualHost)
				**out = **in
			}
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*EndpointGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EndpointGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = new(Failover)
		(*in).DeepCopyInto(*out)
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = new(Cors)
		(*in).DeepCopyInto(*out)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Request) DeepCopyInto(out *Request) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HttpHeader, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Request.
func (in *Request) DeepCopy() *Request {
	if in == nil {
		return nil
	}
	out := new(Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Response) DeepCopyInto(out *Response) {
	*out = *in
	if in.Assertions != nil {
		in, out := &in.Assertions, &out.Assertions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Response.
func (in *Response) DeepCopy() *Response {
	if in == nil {
		return nil
	}
	out := new(Response)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponseTemplate) DeepCopyInto(out *ResponseTemplate) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponseTemplate.
func (in *ResponseTemplate) DeepCopy() *ResponseTemplate {
	if in == nil {
		return nil
	}
	out := new(ResponseTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]HttpMethod, len(*in))
		copy(*out, *in)
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(Policy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Services) DeepCopyInto(out *Services) {
	*out = *in
	if in.EndpointDiscoveryService != nil {
		in, out := &in.EndpointDiscoveryService, &out.EndpointDiscoveryService
		*out = new(EndpointDiscoveryService)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheckService != nil {
		in, out := &in.HealthCheckService, &out.HealthCheckService
		*out = new(HealthCheckService)
		(*in).DeepCopyInto(*out)
	}
	if in.DynamicPropertyService != nil {
		in, out := &in.DynamicPropertyService, &out.DynamicPropertyService
		*out = new(DynamicPropertyService)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Services.
func (in *Services) DeepCopy() *Services {
	if in == nil {
		return nil
	}
	out := new(Services)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	in.Request.DeepCopyInto(&out.Request)
	in.Response.DeepCopyInto(&out.Response)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustStore) DeepCopyInto(out *TrustStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustStore.
func (in *TrustStore) DeepCopy() *TrustStore {
	if in == nil {
		return nil
	}
	out := new(TrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHost) DeepCopyInto(out *VirtualHost) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHost.
func (in *VirtualHost) DeepCopy() *VirtualHost {
	if in == nil {
		return nil
	}
	out := new(VirtualHost)
	in.DeepCopyInto(out)
	return out
}
