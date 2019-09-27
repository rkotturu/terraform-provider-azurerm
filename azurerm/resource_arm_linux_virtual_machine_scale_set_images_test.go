package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

// TODO: from a custom image, dynamically captured `source_image_id`

func TestAccAzureRMLinuxVirtualMachineScaleSet_imagesAutomaticUpdate(t *testing.T) {
	resourceName := "azurerm_linux_virtual_machine_scale_set.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLinuxVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesAutomaticUpdate(ri, location, "16.04-LTS"),
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
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesAutomaticUpdate(ri, location, "18.04-LTS"),
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

func TestAccAzureRMLinuxVirtualMachineScaleSet_imagesManualUpdate(t *testing.T) {
	resourceName := "azurerm_linux_virtual_machine_scale_set.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLinuxVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesManualUpdate(ri, location, "16.04-LTS"),
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
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesManualUpdate(ri, location, "18.04-LTS"),
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

func TestAccAzureRMLinuxVirtualMachineScaleSet_imagesManualUpdateExternalRoll(t *testing.T) {
	// TODO: this requires the feature flag to roll the images outside of TF
}

func TestAccAzureRMLinuxVirtualMachineScaleSet_imagesRollingUpdate(t *testing.T) {
	resourceName := "azurerm_linux_virtual_machine_scale_set.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLinuxVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesRollingUpdate(ri, location, "16.04-LTS"),
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
				Config: testAccAzureRMLinuxVirtualMachineScaleSet_imagesRollingUpdate(ri, location, "18.04-LTS"),
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
		},
	})
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesAutomaticUpdate(rInt int, location, version string) string {
	template := testAccAzureRMLinuxVirtualMachineScaleSet_template(rInt, location)
	return fmt.Sprintf(`
%s


resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "acctestlb-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  frontend_ip_configuration {
    name                 = "internal"
    public_ip_address_id = azurerm_public_ip.test.id
  }
}

resource "azurerm_lb_backend_address_pool" "test" {
  name                = "test"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  loadbalancer_id     = azurerm_lb.test.id
}

resource "azurerm_lb_nat_pool" "test" {
  name                           = "test"
  resource_group_name            = azurerm_resource_group.test.name
  loadbalancer_id                = azurerm_lb.test.id
  frontend_ip_configuration_name = "internal"
  protocol                       = "Tcp"
  frontend_port_start            = 80
  frontend_port_end              = 81
  backend_port                   = 8080
}

resource "azurerm_lb_probe" "test" {
  resource_group_name = azurerm_resource_group.test.name
  loadbalancer_id     = azurerm_lb.test.id
  name                = "acctest-lb-probe"
  port                = 22
  protocol            = "Tcp"
}

resource "azurerm_linux_virtual_machine_scale_set" "test" {
  name                = "acctestvmss-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"
  upgrade_mode        = "Automatic"
  disable_password_authentication = false

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "%s"
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
	  load_balancer_backend_address_pool_ids = [ azurerm_lb_backend_address_pool.test.id ]
	  load_balancer_inbound_nat_rules_ids = [ azurerm_lb_nat_pool.test.id ]
    }
  }

  automatic_os_upgrade_policy {
	disable_automatic_rollback  = true
	enable_automatic_os_upgrade = true
	health_probe_id             = azurerm_lb_probe.test.id
  }
}
`, template, rInt, rInt, rInt, version)
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesManualUpdate(rInt int, location, version string) string {
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
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "%s"
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
}
`, template, rInt, version)
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesRollingUpdate(rInt int, location, version string) string {
	template := testAccAzureRMLinuxVirtualMachineScaleSet_template(rInt, location)
	return fmt.Sprintf(`
%s

resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "acctestlb-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  frontend_ip_configuration {
    name                 = "internal"
    public_ip_address_id = azurerm_public_ip.test.id
  }
}

resource "azurerm_lb_backend_address_pool" "test" {
  name                = "test"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  loadbalancer_id     = azurerm_lb.test.id
}

resource "azurerm_lb_nat_pool" "test" {
  name                           = "test"
  resource_group_name            = azurerm_resource_group.test.name
  loadbalancer_id                = azurerm_lb.test.id
  frontend_ip_configuration_name = "internal"
  protocol                       = "Tcp"
  frontend_port_start            = 80
  frontend_port_end              = 81
  backend_port                   = 8080
}

resource "azurerm_lb_probe" "test" {
  resource_group_name = azurerm_resource_group.test.name
  loadbalancer_id     = azurerm_lb.test.id
  name                = "acctest-lb-probe"
  port                = 22
  protocol            = "Tcp"
}

resource "azurerm_linux_virtual_machine_scale_set" "test" {
  name                = "acctestvmss-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"
  upgrade_mode        = "Rolling"
  disable_password_authentication = false

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "%s"
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
      load_balancer_backend_address_pool_ids = [ azurerm_lb_backend_address_pool.test.id ]
      load_balancer_inbound_nat_rules_ids = [ azurerm_lb_nat_pool.test.id ]
    }
  }

  rolling_upgrade_policy {
    health_probe_id                         = azurerm_lb_probe.test.id
    max_batch_instance_percent              = 21
    max_unhealthy_instance_percent          = 22
    max_unhealthy_upgraded_instance_percent = 23
    pause_time_between_batches              = "PT30S"
  }
}
`, template, rInt, rInt, rInt, version)
}

func testAccAzureRMLinuxVirtualMachineScaleSet_imagesPlan(rInt int, location string) string {
	template := testAccAzureRMLinuxVirtualMachineScaleSet_template(rInt, location)
	return fmt.Sprintf(`
%s

resource "azurerm_marketplace_agreement" "test" {
  publisher = "cloudbees"
  offer     = "jenkins-operations-center"
  plan      = "jenkins-operations-center-solo"
}

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

  depends_on = [ "azurerm_marketplace_agreement.test" ]
}
`, template, rInt)
}
