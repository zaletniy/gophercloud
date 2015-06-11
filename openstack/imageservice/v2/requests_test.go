package v2

// TODO
// compare with openstack/compute/v2/servers/requests_test.go

import (
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fakeclient "github.com/rackspace/gophercloud/testhelper/client"
)

func TestCreateImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageCreationSuccessfully(t)

	id := "e7db3b45-8db7-47ad-8109-3fb55c2c24fd"
	name := "Ubuntu 12.10"

	actualImage, err := Create(fakeclient.ServiceClient(), CreateOpts{
		Id: &id,
		Name: &name,
		Tags: []string{"ubuntu", "quantal"},
	}).Extract()

	th.AssertNoErr(t, err)

	expectedImage := Image{
		Id: "e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
		Name: "Ubuntu 12.10",
		Status: ImageStatusQueued,
		Tags: []string{"ubuntu", "quantal"},
	}

	th.AssertDeepEquals(t, &expectedImage, actualImage)
}

func TestGetImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageGetSuccessfully(t)

	actualImage, err := Get(fakeclient.ServiceClient(), "1bea47ed-f6a9-463b-b423-14b9cca9ad27").Extract()

	th.AssertNoErr(t, err)

	expectedImage := Image{
		Id: "1bea47ed-f6a9-463b-b423-14b9cca9ad27",
		Name: "cirros-0.3.2-x86_64-disk",
		Status: ImageStatusActive,
		Tags: []string{},

		ContainerFormat: "bare",
		DiskFormat: "qcow2",

		MinDiskGigabytes: 0,
		MinRamMegabytes: 0,

		Owner: "5ef70662f8b34079a6eddb8da9d75fe8",

		Protected: false,
		Visibility: ImageVisibilityPublic,

		Checksum: "64d7c1cd2b6f60c92c14662941cb7913",
		SizeBytes: 13167616,
	}

	th.AssertDeepEquals(t, &expectedImage, actualImage)
}
