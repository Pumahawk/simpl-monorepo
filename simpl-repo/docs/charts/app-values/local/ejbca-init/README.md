# EJBCA Init Chart for local development

This is a convenience chart that can be used to easily bootstrap EJBCA in order to get a database configuration suitable for local development.

The resulting ejbca database can then be dumped into the [ejbca.dump](../local-authority/files/ejbca.dump) file, which is used by the local development chart to initialize EJBCA without running the [ejbca-preconfig](https://code.europa.eu/simpl/simpl-open/development/iaa/ejbca-preconfig) script and ensuring the same EJBCA configuration on all local development setups.

## Usage

1. `helm install ejbca-init --create-namespace -n ejbca-init  .`

    **Note:** EJBCA init script will fail if unable to create the secret, but the secret is not required for this setup

2. Download logs of EJBCA init script e.g. to `ejbca-init.log`

    **Note:** if secret creation has failed, the container will restart so be quick to gather the logs

3. Find the dump of the keystore and truststore base64 and convert them respectively to:

    * `SuperAdmin.p12.local` (keystore)

    * `truststore.jks.local` (truststore)

    These are the files that must be configured in the local `Identity Provider` project in order to communicate with EJBCA.

4. Dump the EJBCA database to [ejbca.dump](../local-authority/files/ejbca.dump) file

    * Use Plain SQL dump with INSERT syntax to populate tables

5. Remove the installation and/or the namespace if wanted

    * `helm uninstall ejbca-init -n ejbca-init`

    * `kubectl delete namespace ejbca-init`