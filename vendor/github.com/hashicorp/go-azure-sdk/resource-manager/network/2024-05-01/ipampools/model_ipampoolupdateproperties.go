package ipampools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPamPoolUpdateProperties struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}
