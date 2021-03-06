// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import acm "github.com/aws/aws-sdk-go/service/acm"

import context "context"
import ec2 "github.com/aws/aws-sdk-go/service/ec2"
import elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
import mock "github.com/stretchr/testify/mock"
import resourcegroupstaggingapi "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
import types "github.com/kubernetes-sigs/aws-alb-ingress-controller/pkg/util/types"
import waf "github.com/aws/aws-sdk-go/service/waf"
import wafregional "github.com/aws/aws-sdk-go/service/wafregional"

// CloudAPI is an autogenerated mock type for the CloudAPI type
type CloudAPI struct {
	mock.Mock
}

// ACMAvailable provides a mock function with given fields:
func (_m *CloudAPI) ACMAvailable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AddELBV2TagsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) AddELBV2TagsWithContext(_a0 context.Context, _a1 *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.AddTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.AddTagsInput) *elbv2.AddTagsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.AddTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.AddTagsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddListenerCertificates provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) AddListenerCertificates(_a0 context.Context, _a1 *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.AddListenerCertificatesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.AddListenerCertificatesInput) *elbv2.AddListenerCertificatesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.AddListenerCertificatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.AddListenerCertificatesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateWAF provides a mock function with given fields: ctx, resourceArn, webACLId
func (_m *CloudAPI) AssociateWAF(ctx context.Context, resourceArn *string, webACLId *string) (*wafregional.AssociateWebACLOutput, error) {
	ret := _m.Called(ctx, resourceArn, webACLId)

	var r0 *wafregional.AssociateWebACLOutput
	if rf, ok := ret.Get(0).(func(context.Context, *string, *string) *wafregional.AssociateWebACLOutput); ok {
		r0 = rf(ctx, resourceArn, webACLId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wafregional.AssociateWebACLOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string, *string) error); ok {
		r1 = rf(ctx, resourceArn, webACLId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthorizeSecurityGroupIngressWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) AuthorizeSecurityGroupIngressWithContext(_a0 context.Context, _a1 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.AuthorizeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.AuthorizeSecurityGroupIngressInput) *ec2.AuthorizeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.AuthorizeSecurityGroupIngressInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEC2TagsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateEC2TagsWithContext(_a0 context.Context, _a1 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.CreateTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.CreateTagsInput) *ec2.CreateTagsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.CreateTagsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateListenerWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateListenerWithContext(_a0 context.Context, _a1 *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.CreateListenerOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.CreateListenerInput) *elbv2.CreateListenerOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.CreateListenerOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.CreateListenerInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoadBalancerWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateLoadBalancerWithContext(_a0 context.Context, _a1 *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.CreateLoadBalancerOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.CreateLoadBalancerInput) *elbv2.CreateLoadBalancerOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.CreateLoadBalancerOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.CreateLoadBalancerInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRuleWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateRuleWithContext(_a0 context.Context, _a1 *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.CreateRuleOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.CreateRuleInput) *elbv2.CreateRuleOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.CreateRuleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.CreateRuleInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecurityGroupWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateSecurityGroupWithContext(_a0 context.Context, _a1 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.CreateSecurityGroupOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.CreateSecurityGroupInput) *ec2.CreateSecurityGroupOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSecurityGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.CreateSecurityGroupInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTargetGroupWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) CreateTargetGroupWithContext(_a0 context.Context, _a1 *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.CreateTargetGroupOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.CreateTargetGroupInput) *elbv2.CreateTargetGroupOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.CreateTargetGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.CreateTargetGroupInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEC2TagsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteEC2TagsWithContext(_a0 context.Context, _a1 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.DeleteTagsInput) *ec2.DeleteTagsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.DeleteTagsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteListenersByArn provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteListenersByArn(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteLoadBalancerByArn provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteLoadBalancerByArn(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRuleWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteRuleWithContext(_a0 context.Context, _a1 *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DeleteRuleOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DeleteRuleInput) *elbv2.DeleteRuleOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DeleteRuleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DeleteRuleInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecurityGroupByID provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteSecurityGroupByID(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTargetGroupByArn provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeleteTargetGroupByArn(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeregisterTargetsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DeregisterTargetsWithContext(_a0 context.Context, _a1 *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DeregisterTargetsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DeregisterTargetsInput) *elbv2.DeregisterTargetsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DeregisterTargetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DeregisterTargetsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCertificate provides a mock function with given fields: ctx, certArn
func (_m *CloudAPI) DescribeCertificate(ctx context.Context, certArn string) (*acm.CertificateDetail, error) {
	ret := _m.Called(ctx, certArn)

	var r0 *acm.CertificateDetail
	if rf, ok := ret.Get(0).(func(context.Context, string) *acm.CertificateDetail); ok {
		r0 = rf(ctx, certArn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.CertificateDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, certArn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeELBV2TagsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeELBV2TagsWithContext(_a0 context.Context, _a1 *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DescribeTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DescribeTagsInput) *elbv2.DescribeTagsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DescribeTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DescribeTagsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeListenerCertificates provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeListenerCertificates(_a0 context.Context, _a1 string) ([]*elbv2.Certificate, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*elbv2.Certificate
	if rf, ok := ret.Get(0).(func(context.Context, string) []*elbv2.Certificate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*elbv2.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLoadBalancerAttributesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeLoadBalancerAttributesWithContext(_a0 context.Context, _a1 *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DescribeLoadBalancerAttributesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DescribeLoadBalancerAttributesInput) *elbv2.DescribeLoadBalancerAttributesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DescribeLoadBalancerAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DescribeLoadBalancerAttributesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkInterfaces provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeNetworkInterfaces(_a0 context.Context, _a1 *ec2.DescribeNetworkInterfacesInput) ([]*ec2.NetworkInterface, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*ec2.NetworkInterface
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.DescribeNetworkInterfacesInput) []*ec2.NetworkInterface); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.NetworkInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.DescribeNetworkInterfacesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTargetGroupAttributesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeTargetGroupAttributesWithContext(_a0 context.Context, _a1 *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DescribeTargetGroupAttributesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DescribeTargetGroupAttributesInput) *elbv2.DescribeTargetGroupAttributesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DescribeTargetGroupAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DescribeTargetGroupAttributesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTargetHealthWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) DescribeTargetHealthWithContext(_a0 context.Context, _a1 *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.DescribeTargetHealthOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.DescribeTargetHealthInput) *elbv2.DescribeTargetHealthOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.DescribeTargetHealthOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.DescribeTargetHealthInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateWAF provides a mock function with given fields: ctx, resourceArn
func (_m *CloudAPI) DisassociateWAF(ctx context.Context, resourceArn *string) (*wafregional.DisassociateWebACLOutput, error) {
	ret := _m.Called(ctx, resourceArn)

	var r0 *wafregional.DisassociateWebACLOutput
	if rf, ok := ret.Get(0).(func(context.Context, *string) *wafregional.DisassociateWebACLOutput); ok {
		r0 = rf(ctx, resourceArn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wafregional.DisassociateWebACLOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, resourceArn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterSubnets provides a mock function with given fields:
func (_m *CloudAPI) GetClusterSubnets() (map[string]types.EC2Tags, error) {
	ret := _m.Called()

	var r0 map[string]types.EC2Tags
	if rf, ok := ret.Get(0).(func() map[string]types.EC2Tags); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]types.EC2Tags)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstancesByIDs provides a mock function with given fields: _a0
func (_m *CloudAPI) GetInstancesByIDs(_a0 []string) ([]*ec2.Instance, error) {
	ret := _m.Called(_a0)

	var r0 []*ec2.Instance
	if rf, ok := ret.Get(0).(func([]string) []*ec2.Instance); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoadBalancerByArn provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetLoadBalancerByArn(_a0 context.Context, _a1 string) (*elbv2.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.LoadBalancer
	if rf, ok := ret.Get(0).(func(context.Context, string) *elbv2.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoadBalancerByName provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetLoadBalancerByName(_a0 context.Context, _a1 string) (*elbv2.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.LoadBalancer
	if rf, ok := ret.Get(0).(func(context.Context, string) *elbv2.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcesByFilters provides a mock function with given fields: tagFilters, resourceTypeFilters
func (_m *CloudAPI) GetResourcesByFilters(tagFilters map[string][]string, resourceTypeFilters ...string) ([]string, error) {
	_va := make([]interface{}, len(resourceTypeFilters))
	for _i := range resourceTypeFilters {
		_va[_i] = resourceTypeFilters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, tagFilters)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(map[string][]string, ...string) []string); ok {
		r0 = rf(tagFilters, resourceTypeFilters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string][]string, ...string) error); ok {
		r1 = rf(tagFilters, resourceTypeFilters...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRules provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetRules(_a0 context.Context, _a1 string) ([]*elbv2.Rule, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*elbv2.Rule
	if rf, ok := ret.Get(0).(func(context.Context, string) []*elbv2.Rule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*elbv2.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecurityGroupByID provides a mock function with given fields: _a0
func (_m *CloudAPI) GetSecurityGroupByID(_a0 string) (*ec2.SecurityGroup, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.SecurityGroup
	if rf, ok := ret.Get(0).(func(string) *ec2.SecurityGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.SecurityGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecurityGroupByName provides a mock function with given fields: _a0
func (_m *CloudAPI) GetSecurityGroupByName(_a0 string) (*ec2.SecurityGroup, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.SecurityGroup
	if rf, ok := ret.Get(0).(func(string) *ec2.SecurityGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.SecurityGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecurityGroupsByName provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetSecurityGroupsByName(_a0 context.Context, _a1 []string) ([]*ec2.SecurityGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*ec2.SecurityGroup
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*ec2.SecurityGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.SecurityGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubnetsByNameOrID provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetSubnetsByNameOrID(_a0 context.Context, _a1 []string) ([]*ec2.Subnet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*ec2.Subnet
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*ec2.Subnet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Subnet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetGroupByArn provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetTargetGroupByArn(_a0 context.Context, _a1 string) (*elbv2.TargetGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.TargetGroup
	if rf, ok := ret.Get(0).(func(context.Context, string) *elbv2.TargetGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.TargetGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetGroupByName provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) GetTargetGroupByName(_a0 context.Context, _a1 string) (*elbv2.TargetGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.TargetGroup
	if rf, ok := ret.Get(0).(func(context.Context, string) *elbv2.TargetGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.TargetGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVpcWithContext provides a mock function with given fields: _a0
func (_m *CloudAPI) GetVpcWithContext(_a0 context.Context) (*ec2.Vpc, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Vpc
	if rf, ok := ret.Get(0).(func(context.Context) *ec2.Vpc); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Vpc)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWebACLSummary provides a mock function with given fields: ctx, resourceArn
func (_m *CloudAPI) GetWebACLSummary(ctx context.Context, resourceArn *string) (*waf.WebACLSummary, error) {
	ret := _m.Called(ctx, resourceArn)

	var r0 *waf.WebACLSummary
	if rf, ok := ret.Get(0).(func(context.Context, *string) *waf.WebACLSummary); ok {
		r0 = rf(ctx, resourceArn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*waf.WebACLSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, resourceArn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNodeHealthy provides a mock function with given fields: _a0
func (_m *CloudAPI) IsNodeHealthy(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCertificates provides a mock function with given fields: ctx, input
func (_m *CloudAPI) ListCertificates(ctx context.Context, input *acm.ListCertificatesInput) ([]*acm.CertificateSummary, error) {
	ret := _m.Called(ctx, input)

	var r0 []*acm.CertificateSummary
	if rf, ok := ret.Get(0).(func(context.Context, *acm.ListCertificatesInput) []*acm.CertificateSummary); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*acm.CertificateSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *acm.ListCertificatesInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListListenersByLoadBalancer provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ListListenersByLoadBalancer(_a0 context.Context, _a1 string) ([]*elbv2.Listener, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*elbv2.Listener
	if rf, ok := ret.Get(0).(func(context.Context, string) []*elbv2.Listener); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*elbv2.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyListenerWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyListenerWithContext(_a0 context.Context, _a1 *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.ModifyListenerOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.ModifyListenerInput) *elbv2.ModifyListenerOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.ModifyListenerOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.ModifyListenerInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyLoadBalancerAttributesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyLoadBalancerAttributesWithContext(_a0 context.Context, _a1 *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.ModifyLoadBalancerAttributesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.ModifyLoadBalancerAttributesInput) *elbv2.ModifyLoadBalancerAttributesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.ModifyLoadBalancerAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.ModifyLoadBalancerAttributesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyNetworkInterfaceAttributeWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyNetworkInterfaceAttributeWithContext(_a0 context.Context, _a1 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.ModifyNetworkInterfaceAttributeInput) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.ModifyNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyRuleWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyRuleWithContext(_a0 context.Context, _a1 *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.ModifyRuleOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.ModifyRuleInput) *elbv2.ModifyRuleOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.ModifyRuleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.ModifyRuleInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyTargetGroupAttributesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyTargetGroupAttributesWithContext(_a0 context.Context, _a1 *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.ModifyTargetGroupAttributesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.ModifyTargetGroupAttributesInput) *elbv2.ModifyTargetGroupAttributesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.ModifyTargetGroupAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.ModifyTargetGroupAttributesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyTargetGroupWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) ModifyTargetGroupWithContext(_a0 context.Context, _a1 *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.ModifyTargetGroupOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.ModifyTargetGroupInput) *elbv2.ModifyTargetGroupOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.ModifyTargetGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.ModifyTargetGroupInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterTargetsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) RegisterTargetsWithContext(_a0 context.Context, _a1 *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.RegisterTargetsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.RegisterTargetsInput) *elbv2.RegisterTargetsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.RegisterTargetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.RegisterTargetsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveELBV2TagsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) RemoveELBV2TagsWithContext(_a0 context.Context, _a1 *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.RemoveTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.RemoveTagsInput) *elbv2.RemoveTagsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.RemoveTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.RemoveTagsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveListenerCertificates provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) RemoveListenerCertificates(_a0 context.Context, _a1 *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.RemoveListenerCertificatesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.RemoveListenerCertificatesInput) *elbv2.RemoveListenerCertificatesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.RemoveListenerCertificatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.RemoveListenerCertificatesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeSecurityGroupIngressWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) RevokeSecurityGroupIngressWithContext(_a0 context.Context, _a1 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ec2.RevokeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.RevokeSecurityGroupIngressInput) *ec2.RevokeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ec2.RevokeSecurityGroupIngressInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetIpAddressTypeWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) SetIpAddressTypeWithContext(_a0 context.Context, _a1 *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.SetIpAddressTypeOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.SetIpAddressTypeInput) *elbv2.SetIpAddressTypeOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.SetIpAddressTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.SetIpAddressTypeInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSecurityGroupsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) SetSecurityGroupsWithContext(_a0 context.Context, _a1 *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.SetSecurityGroupsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.SetSecurityGroupsInput) *elbv2.SetSecurityGroupsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.SetSecurityGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.SetSecurityGroupsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSubnetsWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) SetSubnetsWithContext(_a0 context.Context, _a1 *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *elbv2.SetSubnetsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.SetSubnetsInput) *elbv2.SetSubnetsOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elbv2.SetSubnetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elbv2.SetSubnetsInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusACM provides a mock function with given fields:
func (_m *CloudAPI) StatusACM() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// StatusEC2 provides a mock function with given fields:
func (_m *CloudAPI) StatusEC2() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// StatusELBV2 provides a mock function with given fields:
func (_m *CloudAPI) StatusELBV2() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// StatusIAM provides a mock function with given fields:
func (_m *CloudAPI) StatusIAM() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// TagResourcesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) TagResourcesWithContext(_a0 context.Context, _a1 *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *resourcegroupstaggingapi.TagResourcesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.TagResourcesInput) *resourcegroupstaggingapi.TagResourcesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.TagResourcesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.TagResourcesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResourcesWithContext provides a mock function with given fields: _a0, _a1
func (_m *CloudAPI) UntagResourcesWithContext(_a0 context.Context, _a1 *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *resourcegroupstaggingapi.UntagResourcesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.UntagResourcesInput) *resourcegroupstaggingapi.UntagResourcesOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.UntagResourcesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.UntagResourcesInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WAFRegionalAvailable provides a mock function with given fields:
func (_m *CloudAPI) WAFRegionalAvailable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// WebACLExists provides a mock function with given fields: ctx, webACLId
func (_m *CloudAPI) WebACLExists(ctx context.Context, webACLId *string) (bool, error) {
	ret := _m.Called(ctx, webACLId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *string) bool); ok {
		r0 = rf(ctx, webACLId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, webACLId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
