/*
Copyright 2016 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	fv1 "github.com/fission/fission/pkg/apis/fission.io/v1"
)

//
// Fission-Environment interface. The following types are not
// exposed in the Fission API, but rather used by Fission to
// talk to environments.
//
type (
	FetchRequestType int

	FunctionSpecializeRequest struct {
		FetchReq FunctionFetchRequest
		LoadReq  FunctionLoadRequest
	}

	FunctionFetchRequest struct {
		FetchType     FetchRequestType         `json:"fetchType"`
		Package       metav1.ObjectMeta        `json:"package"`
		Url           string                   `json:"url"`
		StorageSvcUrl string                   `json:"storagesvcurl"`
		Filename      string                   `json:"filename"`
		Secrets       []fv1.SecretReference    `json:"secretList"`
		ConfigMaps    []fv1.ConfigMapReference `json:"configMapList"`
		KeepArchive   bool                     `json:"keeparchive"`
	}

	FunctionLoadRequest struct {
		// FilePath is an absolute filesystem path to the
		// function. What exactly is stored here is
		// env-specific. Optional.
		FilePath string `json:"filepath"`

		// FunctionName has an environment-specific meaning;
		// usually, it defines a function within a module
		// containing multiple functions. Optional; default is
		// environment-specific.
		FunctionName string `json:"functionName"`

		// URL to expose this function at. Optional; defaults
		// to "/".
		URL string `json:"url"`

		// Metatdata
		FunctionMetadata *metav1.ObjectMeta

		EnvVersion int `json:"envVersion"`
	}

	// ArchiveUploadRequest send from builder manager describes which
	// deployment package should be upload to storage service.
	ArchiveUploadRequest struct {
		Filename       string `json:"filename"`
		StorageSvcUrl  string `json:"storagesvcurl"`
		ArchivePackage bool   `json:"archivepackage"`
	}

	// ArchiveUploadResponse defines the download url of an archive and
	// its checksum.
	ArchiveUploadResponse struct {
		ArchiveDownloadUrl string       `json:"archiveDownloadUrl"`
		Checksum           fv1.Checksum `json:"checksum"`
	}
)

const (
	FETCH_SOURCE = iota
	FETCH_DEPLOYMENT
	FETCH_URL // remove this?
)

const EXECUTOR_INSTANCEID_LABEL = fv1.EXECUTOR_INSTANCEID_LABEL
const POOLMGR_INSTANCEID_LABEL = fv1.POOLMGR_INSTANCEID_LABEL

const (
	ChecksumTypeSHA256 = fv1.ChecksumTypeSHA256
)

const (
	// ArchiveTypeLiteral means the package contents are specified in the Literal field of
	// resource itself.
	ArchiveTypeLiteral = fv1.ArchiveTypeLiteral

	// ArchiveTypeUrl means the package contents are at the specified URL.
	ArchiveTypeUrl = fv1.ArchiveTypeUrl
)

const (
	BuildStatusPending   = fv1.BuildStatusPending
	BuildStatusRunning   = fv1.BuildStatusRunning
	BuildStatusSucceeded = fv1.BuildStatusSucceeded
	BuildStatusFailed    = fv1.BuildStatusFailed
	BuildStatusNone      = fv1.BuildStatusNone
)

const (
	AllowedFunctionsPerContainerSingle   = fv1.AllowedFunctionsPerContainerSingle
	AllowedFunctionsPerContainerInfinite = fv1.AllowedFunctionsPerContainerInfinite
)

// executor kubernetes object label key
const (
	ENVIRONMENT_NAMESPACE = "environmentNamespace"
	ENVIRONMENT_NAME      = "environmentName"
	ENVIRONMENT_UID       = "environmentUid"
	FUNCTION_NAMESPACE    = "functionNamespace"
	FUNCTION_NAME         = "functionName"
	FUNCTION_UID          = "functionUid"
	EXECUTOR_TYPE         = "executorType"
)

const (
	SharedVolumeUserfunc   = fv1.SharedVolumeUserfunc
	SharedVolumePackages   = fv1.SharedVolumePackages
	SharedVolumeSecrets    = fv1.SharedVolumeSecrets
	SharedVolumeConfigmaps = fv1.SharedVolumeConfigmaps
)

const (
	MessageQueueTypeNats  = fv1.MessageQueueTypeNats
	MessageQueueTypeASQ   = fv1.MessageQueueTypeASQ
	MessageQueueTypeKafka = fv1.MessageQueueTypeKafka
)

const (
	// FunctionReferenceFunctionName means that the function
	// reference is simply by function name.
	FunctionReferenceTypeFunctionName = fv1.FunctionReferenceTypeFunctionName

	//   Set of function references (recursively), by percentage of traffic
	FunctionReferenceTypeFunctionWeights = fv1.FunctionReferenceTypeFunctionWeights

	// Other function reference types we'd like to support:
	//   Versioned function, latest version
	//   Versioned function. by semver "latest compatible"

)

const (
	ArchiveLiteralSizeLimit int64 = 256 * 1024
)

const (
	FissionBuilderSA = "fission-builder"
	FissionFetcherSA = "fission-fetcher"

	SecretConfigMapGetterCR = "secret-configmap-getter"
	SecretConfigMapGetterRB = "secret-configmap-getter-binding"

	PackageGetterCR = "package-getter"
	PackageGetterRB = "package-getter-binding"

	ClusterRole = "ClusterRole"
)

const (
	FailureTypeStatusCode        = fv1.FailureTypeStatusCode
	CanaryConfigStatusPending    = fv1.CanaryConfigStatusPending
	CanaryConfigStatusSucceeded  = fv1.CanaryConfigStatusSucceeded
	CanaryConfigStatusFailed     = fv1.CanaryConfigStatusFailed
	CanaryConfigStatusAborted    = fv1.CanaryConfigStatusAborted
	MaxIterationsForCanaryConfig = fv1.MaxIterationsForCanaryConfig
)
