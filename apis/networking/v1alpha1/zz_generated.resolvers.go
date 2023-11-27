/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Port.
func (mg *Port) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.FixedIP); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FixedIP[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.FixedIP[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.FixedIP[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.FixedIP[i3].SubnetID")
		}
		mg.Spec.ForProvider.FixedIP[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.FixedIP[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &SecgroupList{},
			Managed: &Secgroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this PortSecgroupAssociate.
func (mg *PortSecgroupAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &PortList{},
			Managed: &Port{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &SecgroupList{},
			Managed: &Secgroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Router.
func (mg *Router) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExternalNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ExternalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.ExternalNetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExternalNetworkID")
	}
	mg.Spec.ForProvider.ExternalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExternalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecgroupRule.
func (mg *SecgroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RemoteGroupIDRef,
		Selector:     mg.Spec.ForProvider.RemoteGroupIDSelector,
		To: reference.To{
			List:    &SecgroupList{},
			Managed: &Secgroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteGroupID")
	}
	mg.Spec.ForProvider.RemoteGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecgroupList{},
			Managed: &Secgroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetRoute.
func (mg *SubnetRoute) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}
