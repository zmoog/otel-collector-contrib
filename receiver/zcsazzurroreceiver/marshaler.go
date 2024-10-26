package zcsazzurroreceiver

import (
	"github.com/zmoog/zcs/azzurro"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const (
	scopeName   = "github.com/zmoog/otel-collector-contrib/receiver/zcsazzurroreceiver"
	scopeVerion = "v0.1.0"
)

type azzurroRealtimeDataMarshaler struct {
	logger *zap.Logger
}

func (m *azzurroRealtimeDataMarshaler) UnmarshalMetrics(response azzurro.RealtimeDataResponse) (pmetric.Metrics, error) {
	m.logger.Info("Unmarshalling azzurro realtime data response", zap.Any("response", response))
	md := pmetric.NewMetrics()

	resourceMetrics := md.ResourceMetrics().AppendEmpty()
	resource := resourceMetrics.Resource()

	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName(scopeName)
	scopeMetrics.Scope().SetVersion(scopeVerion)

	if !response.RealtimeData.Success {
		m.logger.Error("Failed to fetch realtime data", zap.Any("response", response))
		return md, nil
	}

	for _, v := range response.RealtimeData.Params.Value {
		for thingKey, value := range v {
			timestamp := pcommon.Timestamp(value.LastUpdate.UnixNano())

			// ----------------------------------------------------------------
			// Power metrics
			// ----------------------------------------------------------------

			powerAutoconsuming := scopeMetrics.Metrics().AppendEmpty()
			powerAutoconsuming.SetName("power_autoconsuming")
			powerAutoconsuming.SetDescription("Power autoconsuming")
			powerAutoconsuming.SetUnit("W")
			powerAutoconsumingDataPoint := powerAutoconsuming.SetEmptyGauge().DataPoints().AppendEmpty()
			powerAutoconsumingDataPoint.SetIntValue(int64(value.PowerAutoconsuming))
			powerAutoconsumingDataPoint.SetTimestamp(timestamp)

			powerCharging := scopeMetrics.Metrics().AppendEmpty()
			powerCharging.SetName("power_charging")
			powerCharging.SetDescription("Power charging")
			powerCharging.SetUnit("W")
			powerChargingDataPoint := powerCharging.SetEmptyGauge().DataPoints().AppendEmpty()
			powerChargingDataPoint.SetIntValue(int64(value.PowerCharging))
			powerChargingDataPoint.SetTimestamp(timestamp)

			powerConsuming := scopeMetrics.Metrics().AppendEmpty()
			powerConsuming.SetName("power_consuming")
			powerConsuming.SetDescription("Power consuming")
			powerConsuming.SetUnit("W")
			powerConsumingDataPoint := powerConsuming.SetEmptyGauge().DataPoints().AppendEmpty()
			powerConsumingDataPoint.SetIntValue(int64(value.PowerConsuming))
			powerConsumingDataPoint.SetTimestamp(timestamp)

			powerDischarging := scopeMetrics.Metrics().AppendEmpty()
			powerDischarging.SetName("power_discharging")
			powerDischarging.SetDescription("Power discharging")
			powerDischarging.SetUnit("W")
			powerDischargingDataPoint := powerDischarging.SetEmptyGauge().DataPoints().AppendEmpty()
			powerDischargingDataPoint.SetIntValue(int64(value.PowerDischarging))
			powerDischargingDataPoint.SetTimestamp(timestamp)

			powerExporting := scopeMetrics.Metrics().AppendEmpty()
			powerExporting.SetName("power_exporting")
			powerExporting.SetDescription("Power exporting")
			powerExporting.SetUnit("W")
			powerExportingDataPoint := powerExporting.SetEmptyGauge().DataPoints().AppendEmpty()
			powerExportingDataPoint.SetIntValue(int64(value.PowerExporting))
			powerExportingDataPoint.SetTimestamp(timestamp)

			powerGenerating := scopeMetrics.Metrics().AppendEmpty()
			powerGenerating.SetName("power_generating")
			powerGenerating.SetDescription("Power generating")
			powerGenerating.SetUnit("W")
			powerGeneratingDataPoint := powerGenerating.SetEmptyGauge().DataPoints().AppendEmpty()
			powerGeneratingDataPoint.SetIntValue(int64(value.PowerGenerating))
			powerGeneratingDataPoint.SetTimestamp(timestamp)

			powerImporting := scopeMetrics.Metrics().AppendEmpty()
			powerImporting.SetName("power_importing")
			powerImporting.SetDescription("Power importing")
			powerImporting.SetUnit("W")
			powerImportingDataPoint := powerImporting.SetEmptyGauge().DataPoints().AppendEmpty()
			powerImportingDataPoint.SetIntValue(int64(value.PowerImporting))
			powerImportingDataPoint.SetTimestamp(timestamp)

			// ----------------------------------------------------------------
			// Battery metrics
			// ----------------------------------------------------------------

			batterySoC := scopeMetrics.Metrics().AppendEmpty()
			batterySoC.SetName("battery_soc")
			batterySoC.SetDescription("Battery SOC")
			batterySoC.SetUnit("%")
			batterySoCDataPoint := batterySoC.SetEmptyGauge().DataPoints().AppendEmpty()
			batterySoCDataPoint.SetIntValue(int64(value.BatterySoC))
			batterySoCDataPoint.SetTimestamp(timestamp)

			batteryCycletimeTotal := scopeMetrics.Metrics().AppendEmpty()
			batteryCycletimeTotal.SetName("battery_cycletime_total")
			batteryCycletimeTotal.SetDescription("Total battery cycletime")
			batteryCycletimeTotal.SetUnit("cycles")
			batteryCycletimeTotalDataPoint := batteryCycletimeTotal.SetEmptySum().DataPoints().AppendEmpty()
			batteryCycletimeTotalDataPoint.SetIntValue(int64(value.BatteryCycletime))
			batteryCycletimeTotalDataPoint.SetTimestamp(timestamp)

			// ----------------------------------------------------------------
			// Energy metrics
			// ----------------------------------------------------------------

			energyAutoconsuming := scopeMetrics.Metrics().AppendEmpty()
			energyAutoconsuming.SetName("energy_autoconsuming")
			energyAutoconsuming.SetDescription("Energy autoconsuming")
			energyAutoconsuming.SetUnit("kWh")
			energyAutoconsumingDataPoint := energyAutoconsuming.SetEmptyGauge().DataPoints().AppendEmpty()
			energyAutoconsumingDataPoint.SetDoubleValue(value.EnergyAutoconsuming)
			energyAutoconsumingDataPoint.SetTimestamp(timestamp)

			energyAutoconsumingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyAutoconsumingTotal.SetName("energy_autoconsuming_total")
			energyAutoconsumingTotal.SetDescription("Energy autoconsuming total")
			energyAutoconsumingTotal.SetUnit("kWh")
			energyAutoconsumingTotalDataPoint := energyAutoconsumingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyAutoconsumingTotalDataPoint.SetDoubleValue(value.EnergyAutoconsumingTotal)
			energyAutoconsumingTotalDataPoint.SetTimestamp(timestamp)

			energyCharging := scopeMetrics.Metrics().AppendEmpty()
			energyCharging.SetName("energy_charging")
			energyCharging.SetDescription("Energy charging")
			energyCharging.SetUnit("kWh")
			energyChargingDataPoint := energyCharging.SetEmptyGauge().DataPoints().AppendEmpty()
			energyChargingDataPoint.SetDoubleValue(value.EnergyCharging)
			energyChargingDataPoint.SetTimestamp(timestamp)

			energyChargingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyChargingTotal.SetName("energy_charging_total")
			energyChargingTotal.SetDescription("Energy charging total")
			energyChargingTotal.SetUnit("kWh")
			energyChargingTotalDataPoint := energyChargingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyChargingTotalDataPoint.SetDoubleValue(value.EnergyChargingTotal)
			energyChargingTotalDataPoint.SetTimestamp(timestamp)

			energyConsuming := scopeMetrics.Metrics().AppendEmpty()
			energyConsuming.SetName("energy_consuming")
			energyConsuming.SetDescription("Energy consuming")
			energyConsuming.SetUnit("kWh")
			energyConsumingDataPoint := energyConsuming.SetEmptyGauge().DataPoints().AppendEmpty()
			energyConsumingDataPoint.SetDoubleValue(value.EnergyConsuming)
			energyConsumingDataPoint.SetTimestamp(timestamp)

			energyConsumingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyConsumingTotal.SetName("energy_consuming_total")
			energyConsumingTotal.SetDescription("Energy consuming total")
			energyConsumingTotal.SetUnit("kWh")
			energyConsumingTotalDataPoint := energyConsumingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyConsumingTotalDataPoint.SetDoubleValue(value.EnergyConsumingTotal)
			energyConsumingTotalDataPoint.SetTimestamp(timestamp)

			energyDischarging := scopeMetrics.Metrics().AppendEmpty()
			energyDischarging.SetName("energy_discharging")
			energyDischarging.SetDescription("Energy discharging")
			energyDischarging.SetUnit("kWh")
			energyDischargingDataPoint := energyDischarging.SetEmptyGauge().DataPoints().AppendEmpty()
			energyDischargingDataPoint.SetDoubleValue(value.EnergyDischarging)
			energyDischargingDataPoint.SetTimestamp(timestamp)

			energyDischargingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyDischargingTotal.SetName("energy_discharging_total")
			energyDischargingTotal.SetDescription("Energy discharging total")
			energyDischargingTotal.SetUnit("kWh")
			energyDischargingTotalDataPoint := energyDischargingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyDischargingTotalDataPoint.SetDoubleValue(value.EnergyDischargingTotal)
			energyDischargingTotalDataPoint.SetTimestamp(timestamp)

			energyExporting := scopeMetrics.Metrics().AppendEmpty()
			energyExporting.SetName("energy_exporting")
			energyExporting.SetDescription("Energy exporting")
			energyExporting.SetUnit("kWh")
			energyExportingDataPoint := energyExporting.SetEmptyGauge().DataPoints().AppendEmpty()
			energyExportingDataPoint.SetDoubleValue(value.EnergyExporting)
			energyExportingDataPoint.SetTimestamp(timestamp)

			energyExportingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyExportingTotal.SetName("energy_exporting_total")
			energyExportingTotal.SetDescription("Energy exporting total")
			energyExportingTotal.SetUnit("kWh")
			energyExportingTotalDataPoint := energyExportingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyExportingTotalDataPoint.SetDoubleValue(value.EnergyExportingTotal)
			energyExportingTotalDataPoint.SetTimestamp(timestamp)

			energyGenerating := scopeMetrics.Metrics().AppendEmpty()
			energyGenerating.SetName("energy_generating")
			energyGenerating.SetDescription("Energy generating")
			energyGenerating.SetUnit("kWh")
			energyGeneratingDataPoint := energyGenerating.SetEmptyGauge().DataPoints().AppendEmpty()
			energyGeneratingDataPoint.SetDoubleValue(value.EnergyGenerating)
			energyGeneratingDataPoint.SetTimestamp(timestamp)

			energyGeneratingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyGeneratingTotal.SetName("energy_generating_total")
			energyGeneratingTotal.SetDescription("Energy generating total")
			energyGeneratingTotal.SetUnit("kWh")
			energyGeneratingTotalDataPoint := energyGeneratingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyGeneratingTotalDataPoint.SetDoubleValue(value.EnergyGeneratingTotal)
			energyGeneratingTotalDataPoint.SetTimestamp(timestamp)

			energyImporting := scopeMetrics.Metrics().AppendEmpty()
			energyImporting.SetName("energy_importing")
			energyImporting.SetDescription("Energy importing")
			energyImporting.SetUnit("kWh")
			energyImportingDataPoint := energyImporting.SetEmptyGauge().DataPoints().AppendEmpty()
			energyImportingDataPoint.SetDoubleValue(value.EnergyImporting)
			energyImportingDataPoint.SetTimestamp(timestamp)

			energyImportingTotal := scopeMetrics.Metrics().AppendEmpty()
			energyImportingTotal.SetName("energy_importing_total")
			energyImportingTotal.SetDescription("Energy importing total")
			energyImportingTotal.SetUnit("kWh")
			energyImportingTotalDataPoint := energyImportingTotal.SetEmptySum().DataPoints().AppendEmpty()
			energyImportingTotalDataPoint.SetDoubleValue(value.EnergyImportingTotal)
			energyImportingTotalDataPoint.SetTimestamp(timestamp)

			resource.Attributes().PutStr("thing_key", thingKey)
		}
	}

	return md, nil
}
