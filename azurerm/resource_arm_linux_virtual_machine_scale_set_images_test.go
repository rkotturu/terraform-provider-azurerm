package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

// TODO: updating the image manually, rolling & automatic?
// TODO: from a custom image, dynamically captured `source_image_id`

func TestAccAzureRMLinuxVirtualMachineScaleSet_imagesPlan(t *testing.T) {
	resourceName := "azurerm_linux_virtual_machine_scale_set.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLinuxVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesPlan(ri, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLinuxVirtualMachineScaleSetExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// not returned from the API
					"admin_password",
				},
			},
			{
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesPlanUpdated(ri, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLinuxVirtualMachineScaleSetExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// not returned from the API
					"admin_password",
				},
			},
		},
	})
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesPlan(rInt int, location string) string {
	template := testAccAzureRMLinuxVirtualMachineScaleSet_template(rInt, location)
	return fmt.Sprintf(`
%s

resource "azurerm_linux_virtual_machine_scale_set" "test" {
  name                = "acctestvmss-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"
  disable_password_authentication = false

  source_image_reference {
    publisher = "cloudbees"
    offer     = "jenkins-operations-center"
    sku       = "jenkins-operations-center-solo"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }

  plan {
    name      = "jenkins-operations-center-solo"
    product   = "jenkins-operations-center"
    publisher = "cloudbees"
  }
}
`, template, rInt)
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesPlanUpdated(rInt int, location string) string {
	template := testAccAzureRMLinuxVirtualMachineScaleSet_template(rInt, location)
	return fmt.Sprintf(`
%s

resource "azurerm_linux_virtual_machine_scale_set" "test" {
  name                = "acctestvmss-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"
  disable_password_authentication = false

  source_image_reference {
    publisher = "rancher"
    offer     = "rancheros"
    sku       = "os"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }

  plan {
    name      = "os"
    product   = "rancheros"
    publisher = "rancher"
  }
}
`, template, rInt)
}
