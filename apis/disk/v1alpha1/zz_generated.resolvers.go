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

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Volume.
func (mg *Volume) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ImageIDRef,
		Selector:     mg.Spec.ForProvider.ImageIDSelector,
		To: reference.To{
			List:    &ImageList{},
			Managed: &Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageID")
	}
	mg.Spec.ForProvider.ImageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotID")
	}
	mg.Spec.ForProvider.SnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceVolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SourceVolIDRef,
		Selector:     mg.Spec.ForProvider.SourceVolIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceVolID")
	}
	mg.Spec.ForProvider.SourceVolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceVolIDRef = rsp.ResolvedReference

	return nil
}
