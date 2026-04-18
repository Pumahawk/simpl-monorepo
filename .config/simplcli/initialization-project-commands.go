package main

import (
	"simplcli/internal/ex"
)

var InitializationProjectCmd = command{
	Name: "initialization:project",
	Descr: "" +
		"Create kubernetes cluster, install dependencies, " +
		" autoconfigure keycloak and ejbca, starts simpl " +
		"microservices and start initialization authority",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var InitializationCodeCmd = command{
	Name: "initialization:code",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		return runList(
			GitInitCmd,
			GitUpdateAllCmd,
		)
	},
}

var InitializationCluster = command{
	Name: "initialization:cluster",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		return runList(
			ClusterCreateCmd,
			InitializationZscalerCmd,
			ClusterForwardNodeUpCmd,
			ClusterRedpandaUpCmd,
			ClusterAuthorityInstallOrUpgradeCmd,
			CodeTestAutomationAutoconfigureClientSecretCmd,
			ClusterInitializationEjbcaCmd,
			ClusterDownloadEjbcaPemCmd,
			ConfigPrepareFilesCmd,
			SimplServicesTier1AuthorityUpCmd,
			ClusterKeycloakAutoconfigureCmd,
			SimplServicesUpCmd,
			SimplServicesFeUpCmd,
			ClusterAuthorityInitializationCmd,
		)
	},
}

var InitializationBuildCodeCmd = command{
	Name: "initialization:build:code",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		if err := ex.RunList(
			ex.New("mvnd", "--stop").Run,
			ex.New("mvnd", "clean", "install", "-N", "-f", "simpl-repo/backend/libs/common").Run,
			ex.New("mvnd", "clean", "install", "-N", "-f", "simpl-repo/backend/libs/simpl-http-client").Run,
		); err != nil {
			return 1
		}
		return CodeBuildNoTestCmd.Func()
	},
}

var InitializationClusterCmd = command{
	Name: "initialization:cluster",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var InitializationZscalerCmd = command{
	Name: "initialization:zscaler",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}
