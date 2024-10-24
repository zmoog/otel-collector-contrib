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

			powerExportingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerExportingTotal.SetName("power_exporting_total")
			powerExportingTotal.SetDescription("Total power exporting")
			powerExportingTotal.SetUnit("W")
			powerExporting := powerExportingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerExporting.SetIntValue(int64(value.PowerExporting))
			powerExporting.SetTimestamp(timestamp)

			powerImportingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerImportingTotal.SetName("power_importing_total")
			powerImportingTotal.SetDescription("Total power importing")
			powerImportingTotal.SetUnit("W")
			powerImporting := powerImportingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerImporting.SetIntValue(int64(value.PowerImporting))
			powerImporting.SetTimestamp(timestamp)

			powerConsumingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerConsumingTotal.SetName("power_consuming_total")
			powerConsumingTotal.SetDescription("Total power consuming")
			powerConsumingTotal.SetUnit("W")
			powerConsuming := powerConsumingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerConsuming.SetIntValue(int64(value.PowerConsuming))
			powerConsuming.SetTimestamp(timestamp)

			powerGeneratingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerGeneratingTotal.SetName("power_generating_total")
			powerGeneratingTotal.SetDescription("Total power generating")
			powerGeneratingTotal.SetUnit("W")
			powerGenerating := powerGeneratingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerGenerating.SetIntValue(int64(value.PowerGenerating))
			powerGenerating.SetTimestamp(timestamp)

			powerChargingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerChargingTotal.SetName("power_charging_total")
			powerChargingTotal.SetDescription("Total power charging")
			powerChargingTotal.SetUnit("W")
			powerCharging := powerChargingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerCharging.SetIntValue(int64(value.PowerCharging))
			powerCharging.SetTimestamp(timestamp)

			powerAutoconsumingTotal := scopeMetrics.Metrics().AppendEmpty()
			powerAutoconsumingTotal.SetName("power_autoconsuming_total")
			powerAutoconsumingTotal.SetDescription("Total power autoconsuming")
			powerAutoconsumingTotal.SetUnit("W")
			powerAutoconsuming := powerAutoconsumingTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			powerAutoconsuming.SetIntValue(int64(value.PowerAutoconsuming))
			powerAutoconsuming.SetTimestamp(timestamp)

			batterySoCTotal := scopeMetrics.Metrics().AppendEmpty()
			batterySoCTotal.SetName("battery_soc_total")
			batterySoCTotal.SetDescription("Total battery SOC")
			batterySoCTotal.SetUnit("%")
			batterySoC := batterySoCTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			batterySoC.SetIntValue(int64(value.BatterySoC))
			batterySoC.SetTimestamp(timestamp)

			batteryCycletimeTotal := scopeMetrics.Metrics().AppendEmpty()
			batteryCycletimeTotal.SetName("battery_cycletime_total")
			batteryCycletimeTotal.SetDescription("Total battery cycletime")
			batteryCycletimeTotal.SetUnit("s")
			batteryCycletime := batteryCycletimeTotal.SetEmptyGauge().DataPoints().AppendEmpty()
			batteryCycletime.SetIntValue(int64(value.BatteryCycletime))
			batteryCycletime.SetTimestamp(timestamp)

			resource.Attributes().PutStr("thing_key", thingKey)
		}
	}

	return md, nil
}
