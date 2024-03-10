// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	spanmetricsconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"
	debugexporter "go.opentelemetry.io/collector/exporter/debugexporter"
	loggingexporter "go.opentelemetry.io/collector/exporter/loggingexporter"
	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	prometheusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	kafkaexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	awscloudwatchlogsexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"
	zpagesextension "go.opentelemetry.io/collector/extension/zpagesextension"
	ballastextension "go.opentelemetry.io/collector/extension/ballastextension"
	memorylimiterextension "go.opentelemetry.io/collector/extension/memorylimiterextension"
	jaegerremotesampling "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	healthcheckextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	pprofextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	oauth2clientauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	bearertokenauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	filestorage "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	attributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	resourceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	spanprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	k8sattributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	resourcedetectionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	filterprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	routingprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"
	jaegerreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	hostmetricsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"
	opencensusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	prometheusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	zipkinreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	kafkareceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	filelogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	journaldreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"
	k8seventsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
)

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		zpagesextension.NewFactory(),
		ballastextension.NewFactory(),
		memorylimiterextension.NewFactory(),
		jaegerremotesampling.NewFactory(),
		healthcheckextension.NewFactory(),
		pprofextension.NewFactory(),
		oauth2clientauthextension.NewFactory(),
		bearertokenauthextension.NewFactory(),
		filestorage.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Receivers, err = receiver.MakeFactoryMap(
		otlpreceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		hostmetricsreceiver.NewFactory(),
		opencensusreceiver.NewFactory(),
		prometheusreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
		kafkareceiver.NewFactory(),
		filelogreceiver.NewFactory(),
		journaldreceiver.NewFactory(),
		k8seventsreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Exporters, err = exporter.MakeFactoryMap(
		debugexporter.NewFactory(),
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		kafkaexporter.NewFactory(),
		awscloudwatchlogsexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		attributesprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		routingprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Connectors, err = connector.MakeFactoryMap(
		spanmetricsconnector.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	return factories, nil
}
