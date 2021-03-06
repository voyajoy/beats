// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "eJzElsFu4zYQhu9+ih977uqQow8Fukbvi7aLPTpjcmwNTJEEZ2jB+/QFZSexU2ebBlGqoyjOfN8MRfIz9nxcgkZdACYWeIlPNOqnBeBZXZFskuISvy4A4J5GvceQfA0Ml0JgZ4rfvv+JIUWxVCTuMLAVcYptScM0tgqp+pHM9d0CKByYlJfYsNEC2AoHr8sp/mdEGviBpz12zLzErqSaz29uYF0HuQzE7u7x3a1gLwY8Pffs7u7hUjSSqLCeH+2sJ8PIhaGuUGb/zPd788XYi+ufAtyoknI0bI7TxN9Xd91F/udSl2Iu186SUeiys6svHjTVUWC/3oZEzz/4iXF7/uoZmYvjaLRjpC0ohOTI2DdEuDTkaowaxc6FoMJwtRSOFo6QiKqMFKeKSVSj6Lh7UcQV9mLrqrTjGVxiHTZcmsfq6zeckik0nyt/yYhtKtNX1STID2ph/5V7Q6HNnZWcqUT2VwKnwscn9p4U5Fyp7KHS3ohhJEWgGl3PHqlAjYqxf1lKa8mh6voD5c4pr816OjA2zPGpUxRRY5BB2kp81B57jmjTVl+/raYIX07MOFCoDFH84JJea6xr11PZsZ9XeXK6Kd7+pZgMmcTDpzE29X/2/xdQ9OcNxvqqkOhqaTUi76VRUMBJ5bZ6ZBtT2XcSu0xuz6azGp9zoLBjObTFGNu+8oABicZlS471+U/5c/xU7UP5pw07VXsvfond5mg8L/yUYZbSfxT7e5Xdi+4ldYXJz8L+5VxpOl8IGuvjVqWWCuOQQh1YQQeSQJvAsPR68rGI8YzoLb5xbEzvzj5VPeX3Bl+lIQduh8JU95S5TCe3vr0F7Q5DbZd2shX27T4kybflaDK8pkFzWk4ZLjXf2Ku3SKqRVe1cz26/3pKEFw7KkOLuv/n9wTkV03aeW8/lmrTdbTKpsscmWX89eGLCxDSdim1Uj2o8XI/J6UYaSA2DxGqvl1yf4n2w6xwiD3n+B5XbHbsl83cAAAD//7kBWzA="
}
