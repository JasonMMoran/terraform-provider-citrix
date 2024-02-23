// Copyright © 2023. Citrix Systems, Inc.

package hypervisor

import (
	citrixorchestration "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
	"github.com/citrix/terraform-provider-citrix/internal/util"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HypervisorResourceModel maps the resource schema data.
type NutanixHypervisorResourceModel struct {
	/**** Connection Details ****/
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Zone types.String `tfsdk:"zone"`
	/** Nutanix Connection **/
	Username                            types.String   `tfsdk:"username"`
	Password                            types.String   `tfsdk:"password"`
	PasswordFormat                      types.String   `tfsdk:"password_format"`
	Addresses                           []types.String `tfsdk:"addresses"`
	MaxAbsoluteActiveActions            types.Int64    `tfsdk:"max_absolute_active_actions"`
	MaxAbsoluteNewActionsPerMinute      types.Int64    `tfsdk:"max_absolute_new_actions_per_minute"`
	MaxPowerActionsPercentageOfMachines types.Int64    `tfsdk:"max_power_actions_percentage_of_machines"`
}

func (r NutanixHypervisorResourceModel) RefreshPropertyValues(hypervisor *citrixorchestration.HypervisorDetailResponseModel) NutanixHypervisorResourceModel {
	r.Id = types.StringValue(hypervisor.GetId())
	r.Name = types.StringValue(hypervisor.GetName())
	r.Username = types.StringValue(hypervisor.GetUserName())
	r.Addresses = util.RefreshList(r.Addresses, hypervisor.GetAddresses())
	r.MaxAbsoluteActiveActions = types.Int64Value(int64(hypervisor.GetMaxAbsoluteActiveActions()))
	r.MaxAbsoluteNewActionsPerMinute = types.Int64Value(int64(hypervisor.GetMaxAbsoluteNewActionsPerMinute()))
	r.MaxPowerActionsPercentageOfMachines = types.Int64Value(int64(hypervisor.GetMaxPowerActionsPercentageOfMachines()))

	hypZone := hypervisor.GetZone()
	r.Zone = types.StringValue(hypZone.GetId())
	return r
}
