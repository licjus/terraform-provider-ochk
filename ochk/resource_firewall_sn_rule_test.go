package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccFirewallSNRuleResource_create_update(t *testing.T) {
	resourceName := "ochk_firewall_sn_rule.default"
	dataSourceRouter := "data.ochk_router.default"
	dataSourceService := "data.ochk_service.http"
	resourceRuleSource := "ochk_security_group.source"
	resourceRuleDestination := "ochk_security_group.destination"

	sourceDisplayName := generateRandName()
	destinationDisplayName := generateRandName()
	ruleDisplayName := generateRandName()

	routerName := "T1"
	routerNameUpdated := "T0"

	ruleDisplayNameUpdated := ruleDisplayName + "-updated"
	actionUpdated := "DROP"
	directionUpdated := "OUT"
	ipProtocolUpdated := "IPV4"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallSNRuleResourceConfig(routerName, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigWithPriority(ruleDisplayName, 100, "data.ochk_custom_service.custom_service1.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", ruleDisplayName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "direction", "IN_OUT"),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", "IPV4_IPV6"),
					resource.TestCheckResourceAttrPair(resourceName, "source.0", resourceRuleSource, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "destination.0", resourceRuleDestination, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "services.0", dataSourceService, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "scope.0", dataSourceRouter, "id"),
					resource.TestCheckResourceAttr(resourceName, "priority", "100"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
			{
				Config: testAccFirewallSNRuleResourceConfig(routerNameUpdated, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigWithPriority(ruleDisplayName, 100, "data.ochk_custom_service.custom_service1.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "scope.0", dataSourceRouter, "id"),
				),
			},
			{
				Config: testAccFirewallSNRuleResourceConfig(routerNameUpdated, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigNoPositionUpdate(ruleDisplayNameUpdated, actionUpdated, ipProtocolUpdated, directionUpdated, 200, "data.ochk_custom_service.custom_service2.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", ruleDisplayNameUpdated),
					resource.TestCheckResourceAttr(resourceName, "action", actionUpdated),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocolUpdated),
					resource.TestCheckResourceAttr(resourceName, "direction", directionUpdated),
					resource.TestCheckResourceAttr(resourceName, "priority", "200"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service2", "id"),
				),
			},
		},
		CheckDestroy: testAccFirewallSNRuleResourceNotExists(ruleDisplayName),
	})
}

func testAccFirewallSNRuleResourceConfig(router string, source string, destination string) string {
	return fmt.Sprintf(`
locals {
	routerDisplayName = %[1]q
}

data "ochk_gateway_policy" "default" {
  display_name = local.routerDisplayName
}

data "ochk_router" "default" {
  display_name = local.routerDisplayName
}

data "ochk_service" "http" {
  display_name = "http"
}

data "ochk_virtual_machine" "default" {
	display_name = %[4]q
}

data "ochk_custom_service" "custom_service1" {
	display_name = %[5]q
}

data "ochk_custom_service" "custom_service2" {
	display_name = %[6]q
}

resource "ochk_security_group" "source" {
  display_name = %[2]q

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination" {
  display_name = %[3]q
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}
`, router, source, destination, testData.LegacyVirtualMachineDisplayName, testData.CustomService1DisplayName, testData.CustomService2DisplayName)
}

func testAccFirewallSNRuleResourceConfigWithPriority(displayName string, priority int64, customServiceResourceID string) string {
	return fmt.Sprintf(`
resource "ochk_firewall_sn_rule" "default" {
  display_name = %[1]q
  gateway_policy_id = data.ochk_gateway_policy.default.id

  services = [data.ochk_service.http.id]
  custom_services = [%[3]s]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
  scope = [data.ochk_router.default.id]
  priority = %[2]d
}
`, displayName, priority, customServiceResourceID)
}

func testAccFirewallSNRuleResourceConfigNoPositionUpdate(displayName string, action string, ipProtocol string, direction string, priority int64, customServiceResourceID string) string {
	return fmt.Sprintf(`
resource "ochk_firewall_sn_rule" "default" {
  display_name = %[1]q
  action = %[2]q 
  ip_protocol = %[3]q
  direction = %[4]q
  gateway_policy_id = data.ochk_gateway_policy.default.id

  services = [data.ochk_service.http.id]
  custom_services = [%[6]s]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
  scope = [data.ochk_router.default.id]
  priority = %[5]d 
}
`, displayName, action, ipProtocol, direction, priority, customServiceResourceID)
}

func testAccFirewallSNRuleResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		gatewayPolicies, err := client.GatewayPolicy.ListByDisplayName(ctx, "T1")
		if err != nil {
			return err
		}

		if len(gatewayPolicies) != 1 {
			return fmt.Errorf("wrong number of gateway policies")
		}

		firewallRule, err := client.FirewallSNRules.ListByDisplayName(ctx, gatewayPolicies[0].GatewayPolicyID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall SN rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
