// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ImageInitParameters struct {

	// optional string →  The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported.New since v0.4.2.
	// The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported._new_since_v0.4.2_.
	ArchivingFormat *string `json:"archivingFormat,omitempty" tf:"archiving_format,omitempty"`

	// Type header will be used to detect compression format.New since v0.4.2.
	// The format of compressed image. Use this attribute to decompress image when downloading it from source. Must be one of "auto", "bzip2", "gzip", "xz". If set to "auto", response Content-Type header will be used to detect compression format._new_since_v0.4.2_.
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// required string →  The container format. Must be one of "bare".
	// The container format. Must be one of "bare".
	ContainerFormat *string `json:"containerFormat,omitempty" tf:"container_format,omitempty"`

	// required string →  The disk format. Must be one of "raw", "iso".
	// The disk format. Must be one of "raw", "iso".
	DiskFormat *string `json:"diskFormat,omitempty" tf:"disk_format,omitempty"`

	// optional string →  This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/
	ImageCachePath *string `json:"imageCachePath,omitempty" tf:"image_cache_path,omitempty"`

	// optional string →  This is the url of the raw image. The image will be downloaded in the image_cache_path before being uploaded to VKCS. Conflicts with local_file_path.
	// This is the url of the raw image. The image will be downloaded in the `image_cache_path` before being uploaded to VKCS. Conflicts with `local_file_path`.
	ImageSourceURL *string `json:"imageSourceUrl,omitempty" tf:"image_source_url,omitempty"`

	// optional string →  The username of basic auth to download image_source_url.
	// The username of basic auth to download `image_source_url`.
	ImageSourceUsername *string `json:"imageSourceUsername,omitempty" tf:"image_source_username,omitempty"`

	// optional string →  This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with image_source_url
	// This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with `image_source_url`
	LocalFilePath *string `json:"localFilePath,omitempty" tf:"local_file_path,omitempty"`

	// optional number →  Amount of disk space (in GB) required to boot VM. Defaults to 0.
	// Amount of disk space (in GB) required to boot VM. Defaults to 0.
	MinDiskGb *float64 `json:"minDiskGb,omitempty" tf:"min_disk_gb,omitempty"`

	// optional number →  Amount of ram (in MB) required to boot VM. Defauts to 0.
	// Amount of ram (in MB) required to boot VM. Defauts to 0.
	MinRAMMb *float64 `json:"minRamMb,omitempty" tf:"min_ram_mb,omitempty"`

	// required string →  The name of the image.
	// The name of the image.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// optional map of string →  A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	// A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// optional boolean →  If true, image will not be deletable. Defaults to false.
	// If true, image will not be deletable. Defaults to false.
	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	// optional string →  The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the region argument of the provider is used. Changing this creates a new Image.
	// The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the `region` argument of the provider is used. Changing this creates a new Image.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional set of string →  The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	// The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// optional boolean →  If false, the checksum will not be verified once the image is finished uploading.
	// If false, the checksum will not be verified once the image is finished uploading.
	VerifyChecksum *bool `json:"verifyChecksum,omitempty" tf:"verify_checksum,omitempty"`

	// optional string →  The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	// The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ImageObservation struct {

	// optional string →  The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported.New since v0.4.2.
	// The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported._new_since_v0.4.2_.
	ArchivingFormat *string `json:"archivingFormat,omitempty" tf:"archiving_format,omitempty"`

	// string →  The checksum of the data associated with the image.
	// The checksum of the data associated with the image.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// Type header will be used to detect compression format.New since v0.4.2.
	// The format of compressed image. Use this attribute to decompress image when downloading it from source. Must be one of "auto", "bzip2", "gzip", "xz". If set to "auto", response Content-Type header will be used to detect compression format._new_since_v0.4.2_.
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// required string →  The container format. Must be one of "bare".
	// The container format. Must be one of "bare".
	ContainerFormat *string `json:"containerFormat,omitempty" tf:"container_format,omitempty"`

	// string →  The date the image was created.
	// The date the image was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// required string →  The disk format. Must be one of "raw", "iso".
	// The disk format. Must be one of "raw", "iso".
	DiskFormat *string `json:"diskFormat,omitempty" tf:"disk_format,omitempty"`

	// string →  The trailing path after the image endpoint that represent the location of the image or the path to retrieve it.
	// The trailing path after the image endpoint that represent the location of the image or the path to retrieve it.
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// optional string →  This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/
	ImageCachePath *string `json:"imageCachePath,omitempty" tf:"image_cache_path,omitempty"`

	// optional string →  This is the url of the raw image. The image will be downloaded in the image_cache_path before being uploaded to VKCS. Conflicts with local_file_path.
	// This is the url of the raw image. The image will be downloaded in the `image_cache_path` before being uploaded to VKCS. Conflicts with `local_file_path`.
	ImageSourceURL *string `json:"imageSourceUrl,omitempty" tf:"image_source_url,omitempty"`

	// optional string →  The username of basic auth to download image_source_url.
	// The username of basic auth to download `image_source_url`.
	ImageSourceUsername *string `json:"imageSourceUsername,omitempty" tf:"image_source_username,omitempty"`

	// optional string →  This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with image_source_url
	// This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with `image_source_url`
	LocalFilePath *string `json:"localFilePath,omitempty" tf:"local_file_path,omitempty"`

	// concepts.html.
	// The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// optional number →  Amount of disk space (in GB) required to boot VM. Defaults to 0.
	// Amount of disk space (in GB) required to boot VM. Defaults to 0.
	MinDiskGb *float64 `json:"minDiskGb,omitempty" tf:"min_disk_gb,omitempty"`

	// optional number →  Amount of ram (in MB) required to boot VM. Defauts to 0.
	// Amount of ram (in MB) required to boot VM. Defauts to 0.
	MinRAMMb *float64 `json:"minRamMb,omitempty" tf:"min_ram_mb,omitempty"`

	// required string →  The name of the image.
	// The name of the image.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// string →  The id of the vkcs user who owns the image.
	// The id of the vkcs user who owns the image.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// optional map of string →  A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	// A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// optional boolean →  If true, image will not be deletable. Defaults to false.
	// If true, image will not be deletable. Defaults to false.
	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	// optional string →  The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the region argument of the provider is used. Changing this creates a new Image.
	// The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the `region` argument of the provider is used. Changing this creates a new Image.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// schema that represent the image or image
	// The path to the JSON-schema that represent the image or image
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// number →  The size in bytes of the data associated with the image.
	// The size in bytes of the data associated with the image.
	SizeBytes *float64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`

	// string →  The status of the image. It can be "queued", "active" or "saving".
	// The status of the image. It can be "queued", "active" or "saving".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// optional set of string →  The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	// The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// string →  The date the image was last updated.
	// The date the image was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// optional boolean →  If false, the checksum will not be verified once the image is finished uploading.
	// If false, the checksum will not be verified once the image is finished uploading.
	VerifyChecksum *bool `json:"verifyChecksum,omitempty" tf:"verify_checksum,omitempty"`

	// optional string →  The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	// The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ImageParameters struct {

	// optional string →  The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported.New since v0.4.2.
	// The format of archived image file. Use this to unzip image file when downloading an archive. Currently only "tar" format is supported._new_since_v0.4.2_.
	// +kubebuilder:validation:Optional
	ArchivingFormat *string `json:"archivingFormat,omitempty" tf:"archiving_format,omitempty"`

	// Type header will be used to detect compression format.New since v0.4.2.
	// The format of compressed image. Use this attribute to decompress image when downloading it from source. Must be one of "auto", "bzip2", "gzip", "xz". If set to "auto", response Content-Type header will be used to detect compression format._new_since_v0.4.2_.
	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// required string →  The container format. Must be one of "bare".
	// The container format. Must be one of "bare".
	// +kubebuilder:validation:Optional
	ContainerFormat *string `json:"containerFormat,omitempty" tf:"container_format,omitempty"`

	// required string →  The disk format. Must be one of "raw", "iso".
	// The disk format. Must be one of "raw", "iso".
	// +kubebuilder:validation:Optional
	DiskFormat *string `json:"diskFormat,omitempty" tf:"disk_format,omitempty"`

	// optional string →  This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/
	// +kubebuilder:validation:Optional
	ImageCachePath *string `json:"imageCachePath,omitempty" tf:"image_cache_path,omitempty"`

	// optional sensitive string →  The password of basic auth to download image_source_url.
	// The password of basic auth to download `image_source_url`.
	// +kubebuilder:validation:Optional
	ImageSourcePasswordSecretRef *v1.SecretKeySelector `json:"imageSourcePasswordSecretRef,omitempty" tf:"-"`

	// optional string →  This is the url of the raw image. The image will be downloaded in the image_cache_path before being uploaded to VKCS. Conflicts with local_file_path.
	// This is the url of the raw image. The image will be downloaded in the `image_cache_path` before being uploaded to VKCS. Conflicts with `local_file_path`.
	// +kubebuilder:validation:Optional
	ImageSourceURL *string `json:"imageSourceUrl,omitempty" tf:"image_source_url,omitempty"`

	// optional string →  The username of basic auth to download image_source_url.
	// The username of basic auth to download `image_source_url`.
	// +kubebuilder:validation:Optional
	ImageSourceUsername *string `json:"imageSourceUsername,omitempty" tf:"image_source_username,omitempty"`

	// optional string →  This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with image_source_url
	// This is the filepath of the raw image file that will be uploaded to VKCS. Conflicts with `image_source_url`
	// +kubebuilder:validation:Optional
	LocalFilePath *string `json:"localFilePath,omitempty" tf:"local_file_path,omitempty"`

	// optional number →  Amount of disk space (in GB) required to boot VM. Defaults to 0.
	// Amount of disk space (in GB) required to boot VM. Defaults to 0.
	// +kubebuilder:validation:Optional
	MinDiskGb *float64 `json:"minDiskGb,omitempty" tf:"min_disk_gb,omitempty"`

	// optional number →  Amount of ram (in MB) required to boot VM. Defauts to 0.
	// Amount of ram (in MB) required to boot VM. Defauts to 0.
	// +kubebuilder:validation:Optional
	MinRAMMb *float64 `json:"minRamMb,omitempty" tf:"min_ram_mb,omitempty"`

	// required string →  The name of the image.
	// The name of the image.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// optional map of string →  A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	// A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// optional boolean →  If true, image will not be deletable. Defaults to false.
	// If true, image will not be deletable. Defaults to false.
	// +kubebuilder:validation:Optional
	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	// optional string →  The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the region argument of the provider is used. Changing this creates a new Image.
	// The region in which to obtain the Image client. An Image client is needed to create an Image that can be used with a compute instance. If omitted, the `region` argument of the provider is used. Changing this creates a new Image.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional set of string →  The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	// The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// optional boolean →  If false, the checksum will not be verified once the image is finished uploading.
	// If false, the checksum will not be verified once the image is finished uploading.
	// +kubebuilder:validation:Optional
	VerifyChecksum *bool `json:"verifyChecksum,omitempty" tf:"verify_checksum,omitempty"`

	// optional string →  The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	// The visibility of the image. Must be one of "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the VKCS cloud.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ImageInitParameters `json:"initProvider,omitempty"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Manages an Image resource within VKCS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.containerFormat) || (has(self.initProvider) && has(self.initProvider.containerFormat))",message="spec.forProvider.containerFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.diskFormat) || (has(self.initProvider) && has(self.initProvider.diskFormat))",message="spec.forProvider.diskFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ImageSpec   `json:"spec"`
	Status ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}