// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	elasticsearchexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"
	basicauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	transformprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	azureeventhubreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"
	"github.com/zmoog/otel-collector-contrib/receiver/toggltrackreceiver"
	zcsazzurroreceiver "github.com/zmoog/otel-collector-contrib/receiver/zcsazzurroreceiver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	debugexporter "go.opentelemetry.io/collector/exporter/debugexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/receiver"
)

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		basicauthextension.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ExtensionModules = make(map[component.Type]string, len(factories.Extensions))
	factories.ExtensionModules[basicauthextension.NewFactory().Type()] = "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension v0.111.0"

	factories.Receivers, err = receiver.MakeFactoryMap(
		toggltrackreceiver.NewFactory(),
		zcsazzurroreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ReceiverModules = make(map[component.Type]string, len(factories.Receivers))
	factories.ReceiverModules[azureeventhubreceiver.NewFactory().Type()] = "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver v0.111.0"

	factories.Exporters, err = exporter.MakeFactoryMap(
		debugexporter.NewFactory(),
		elasticsearchexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ExporterModules = make(map[component.Type]string, len(factories.Exporters))
	factories.ExporterModules[debugexporter.NewFactory().Type()] = "go.opentelemetry.io/collector/exporter/debugexporter v0.111.0"
	factories.ExporterModules[elasticsearchexporter.NewFactory().Type()] = "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter v0.111.0"

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
		transformprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ProcessorModules = make(map[component.Type]string, len(factories.Processors))
	factories.ProcessorModules[batchprocessor.NewFactory().Type()] = "go.opentelemetry.io/collector/processor/batchprocessor v0.111.0"
	factories.ProcessorModules[transformprocessor.NewFactory().Type()] = "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor v0.111.0"

	factories.Connectors, err = connector.MakeFactoryMap(
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ConnectorModules = make(map[component.Type]string, len(factories.Connectors))

	return factories, nil
}
