package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccFirewallEWRuleResource_create_update(t *testing.T) {
	resourceName := "ochk_firewall_ew_rule.default"
	displayName := generateRandName()
	displayNameUpdated := displayName + "-updated"
	action := "ALLOW"
	actionUpdated := "DROP"
	direction := "IN_OUT"
	directionUpdated := "OUT"
	ipProtocol := "IPV4_IPV6"
	ipProtocolUpdated := "IPV4"

	source := generateRandName()
	destination := generateRandName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfig(displayName, source, destination, action, ipProtocol, direction, 1000, "data.ochk_custom_service.custom_service1.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "direction", direction),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocol),
					resource.TestCheckResourceAttr(resourceName, "priority", "1000"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
			{
				Config: testAccFirewallEWRuleResourceConfig(displayNameUpdated, source, destination, actionUpdated, ipProtocolUpdated, directionUpdated, 2000, "data.ochk_custom_service.custom_service2.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayNameUpdated),
					resource.TestCheckResourceAttr(resourceName, "action", actionUpdated),
					resource.TestCheckResourceAttr(resourceName, "direction", directionUpdated),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocolUpdated),
					resource.TestCheckResourceAttr(resourceName, "priority", "2000"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service2", "id"),
				),
			},
		},
		CheckDestroy: testAccFirewallEWRuleResourceDoesntExist(displayName),
	})
}

func TestAccFirewallEWRuleResource_withPriority(t *testing.T) {
	randDisplayName := generateRandName()
	displayNameMiddle := randDisplayName + "-middle"
	displayNameBefore := randDisplayName + "-before"
	displayNameAfter := randDisplayName + "-after"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore, displayNameMiddle, displayNameAfter),
				Check: resource.ComposeTestCheckFunc(
					testAccFirewallEWRuleCheckRulesOrder("data.ochk_security_policy.default", displayNameBefore, displayNameMiddle, displayNameAfter),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccFirewallEWRuleResourceDoesntExist(displayNameBefore),
			testAccFirewallEWRuleResourceDoesntExist(displayNameMiddle),
			testAccFirewallEWRuleResourceDoesntExist(displayNameAfter)),
	})
}

func testAccFirewallEWRuleCheckRulesOrder(securityPolicyResourceName string, displayNames ...string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		securityPolicyResource, ok := state.RootModule().Resources[securityPolicyResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", securityPolicyResourceName)
		}

		securityPolicies, err := client.FirewallEWRules.List(ctx, securityPolicyResource.Primary.ID)
		if err != nil {
			return err
		}

		if !checkOrderOfSecurityPolicies(securityPolicies, displayNames) {
			securityPoliciesDisplayNames := transformToStringSlice(len(securityPolicies), func(idx int) string {
				return securityPolicies[idx].DisplayName
			})
			return fmt.Errorf("security policies not found in expected order: %+v, security polices: %+v", displayNames, securityPoliciesDisplayNames)
		}

		return nil
	}
}

func TestCheckOrderOfSecurityPolicies(t *testing.T) {
	assert.NotPanics(t, func() {
		expectedOrder := []string{"a", "b", "c"}

		matchingCases := [][]string{
			{"a", "b", "c"},
			{"0", "a", "b", "c"},
			{"a", "b", "c", "d"},
			{"1", "2", "3", "4", "a", "b", "c"},
			{"1", "2", "3", "4", "a", "b", "c", "5", "6", "7"},
		}

		notMatchingCases := [][]string{
			{"a", "b"},
			{"0", "a", "b"},
			{"b", "c", "d"},
			{"1", "2", "3", "4"},
			{},
		}

		stringArrayToDFWRule := func(strArr []string) []*models.DFWRule {
			var result []*models.DFWRule
			for i := 0; i < len(strArr); i++ {
				result = append(result, &models.DFWRule{DisplayName: strArr[i]})
			}
			return result
		}

		for i := 0; i < len(matchingCases); i++ {
			assert.True(t, checkOrderOfSecurityPolicies(stringArrayToDFWRule(matchingCases[i]), expectedOrder))
		}

		for i := 0; i < len(notMatchingCases); i++ {
			assert.False(t, checkOrderOfSecurityPolicies(stringArrayToDFWRule(notMatchingCases[i]), expectedOrder))
		}
	})
}

func checkOrderOfSecurityPolicies(securityPolicies []*models.DFWRule, displayNames []string) bool {
	indexOfFirstElement := findIndexOfFirstMatch(len(securityPolicies), func(idx int) bool {
		return securityPolicies[idx].DisplayName == displayNames[0]
	})

	if indexOfFirstElement == -1 {
		return false
	}

	if indexOfFirstElement+len(displayNames) > len(securityPolicies) {
		return false
	}

	leftSlice := securityPolicies[indexOfFirstElement : indexOfFirstElement+len(displayNames)]
	return slicesEqual(len(leftSlice), len(displayNames), func(lhsIdx int, rhsIdx int) bool {
		return leftSlice[lhsIdx].DisplayName == displayNames[rhsIdx]
	})
}

func testAccFirewallEWRuleResourceConfig(displayName string, source string, destination string, action string, ipProtocol string, direction string, priority int64, customServiceResourceID string) string {
	return fmt.Sprintf(`
data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_service" "http" {
  display_name = "http"
}

data "ochk_virtual_machine" "default" {
	display_name = %[7]q
}

data "ochk_custom_service" "custom_service1" {
	display_name = %[9]q
}

data "ochk_custom_service" "custom_service2" {
	display_name = %[10]q
}

resource "ochk_security_group" "source" {
  display_name = %[1]q

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination" {
  display_name = %[2]q
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_ew_rule" "default" {
  display_name = %[3]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  custom_services = [%[11]s]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]

  action = %[4]q 
  ip_protocol = %[5]q
  direction = %[6]q

  priority = %[8]d
}
`, source, destination, displayName, action, ipProtocol, direction, testData.LegacyVirtualMachineDisplayName, priority, testData.CustomService1DisplayName, testData.CustomService2DisplayName, customServiceResourceID)
}

func testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore string, displayNameMiddle string, displayNameAfter string) string {
	source := generateRandName()
	destination := generateRandName()

	return fmt.Sprintf(`
data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_service" "http" {
  display_name = "http"
}

data "ochk_virtual_machine" "default" {
	display_name = %[6]q
}

resource "ochk_security_group" "source-before" {
  display_name = "%[1]s-before"

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "source-middle" {
  display_name = "%[1]s-middle"

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "source-after" {
  display_name = "%[1]s-after"

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-before" {
  display_name = "%[2]s-before"
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-middle" {
  display_name = "%[2]s-middle"
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-after" {
  display_name = "%[2]s-after"
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_ew_rule" "middle" {
  display_name = %[3]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-middle.id]
  destination = [ochk_security_group.destination-middle.id]
  priority = 20001
}

resource "ochk_firewall_ew_rule" "before" {
  display_name = %[4]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-before.id]
  destination = [ochk_security_group.destination-before.id]
  priority = 20000
}

resource "ochk_firewall_ew_rule" "after" {
  display_name = %[5]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-after.id]
  destination = [ochk_security_group.destination-after.id]
  priority = 20002
}
`, source, destination, displayNameMiddle, displayNameBefore, displayNameAfter, testData.LegacyVirtualMachineDisplayName)
}

func testAccFirewallEWRuleResourceDoesntExist(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		securityPolicies, err := client.SecurityPolicy.ListByDisplayName(ctx, "devel")
		if err != nil {
			return err
		}

		if len(securityPolicies) != 1 {
			return fmt.Errorf("wrong number of security policies")
		}

		firewallRule, err := client.FirewallEWRules.ListByDisplayName(ctx, securityPolicies[0].SecurityPolicyID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall EW rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
