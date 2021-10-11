package consumption_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type ConsumptionBudgetResourceGroupDataSource struct{}

func TestAccDataSourceBudget_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_consumption_budget_resource_group", "test")
	r := ConsumptionBudgetResourceGroupDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").Exists(),
				// TODO fill in remaining checks
				check.That(data.ResourceName).Key("resource_group_id").Exists(),
				check.That(data.ResourceName).Key("amount").Exists(),
				check.That(data.ResourceName).Key("notification").Exists(),
				check.That(data.ResourceName).Key("notification.0.threshold").Exists(),
				check.That(data.ResourceName).Key("notification.0.operator").Exists(),
				check.That(data.ResourceName).Key("time_period.start_date").Exists(),
			),
		},
	})
}

func (d ConsumptionBudgetResourceGroupDataSource) basic(data acceptance.TestData) string {
	config := ConsumptionBudgetResourceGroupResource{}.basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_consumption_budget_resource_group" "test" {
  name              = azurerm_consumption_budget_resource_group.test.name
  resource_group_id = azurerm_consumption_budget_resource_group.test.resource_group_id
}
`, config)
}
