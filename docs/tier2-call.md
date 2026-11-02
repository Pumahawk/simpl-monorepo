Perform a tier2 call

The documentation for performing a tier2 call is available here: https://code.europa.eu/simpl/simpl-open/security/access-control-and-trust/supporting-security-services/gateway/tier-2-gateway/-/tree/main/documents/user-manual?ref_type=heads
The documentation describes the requirements and the endpoints that must be called to perform a tier2 call.
Since tier2 calls use mTLS, the following steps explain how to easily retrieve the private key and certificate and provide examples using curl commands.

Prerequisites:
- Working Authority (verify that the echo page does not return errors - https://${AUTHORITY_URL}/participant-utility/echo user: m.m)
- Working Participant (verify that the echo page does not return errors - https://${PARTICIPANT_URL}/participant-utility/echo user: a.w)
- kubectl configured with access to the participant and with the namespace configured to access the target participant. kubectl will be required to create a port-forward to the authentication-provider-be service.
- bash, curl, jq: Optional tools. They are recommended if you want to execute the commands exactly as written in this email. Their absence requires modifying the commands.

Steps to perform a tier2 call:

Port-forward to the authentication-provider service.
After ensuring that kubectl has been configured to access both the cluster and the namespace of the target participant, run the following command:

kubectl port-forward svc/authentication-provider-be 8081:8080

This command will enable port forwarding from local port 8081 to port 8080 of the authentication-provider-be service and must remain active for all subsequent commands.
Open a new terminal window so that the port-forward command can remain active.

Download the private key using curl:

curl localhost:8081/tier1/v2/keypairs/active | jq -r .privateKey > private.pem

Note that the call uses the port-forward and jq to extract the private key from the JSON response and save it into a file named private.pem.

To retrieve the certificate, follow these steps:

1 - Open the participant echo page (https://${PARTICIPANT_URL}/participant-utility/echo user: a.w) and note the value of "Common name". It is a UUID.
2 - Access the authority using this URL: https://${AUTHORITY_URL}/identity-provider/participant/${Common name}, making sure to replace ${Common name} with the value retrieved from the participant page, and log in with user m.m.
3 - Click the Credentials tab, locate the active credential (the one marked as active), and download the certificate by clicking the download icon.
4 - At this point the certificate has been downloaded and must be used for mTLS calls. To continue using the commands provided in this email, move the downloaded certificate from the Downloads folder to the directory containing the private.pem file.

Retrieve the participant tier2 address. This information is available within the service.
It is stored in the annotations. The easiest way to retrieve it is with the following command:

kubectl get service tier2-gateway -oyaml | grep tls

The output is a single line containing the TLS address for tier2 calls, for example:

tls.participant.iaadsdevconsumer.sandbox-iaa.simpl-europe.eu

The following curl commands can then be used to perform tier2 calls.

Use the port-forward to retrieve the ephemeral proof and save it in the EPHEMERAL_PROOF variable:

EPHEMERAL_PROOF=$(curl -S -s http://localhost:8081/tier1/v2/ephemeralProof | jq -r .proof)

Perform the call to create the tier2 session:

curl --key private.pem --cert /tmp/certificate.pem --cacert certificate.pem -X POST https://${ENDPOINT_TLS}/authApi/tier2/v2/ephemeralProof --data '{"proof": "'"$EPHEMERAL_PROOF"'"}' -H "Content-Type: application/json"

Perform a ping request to ensure that the tier2 call is working correctly:

curl --key private.pem --cert certificate.pem --cacert /tmp/certificate.pem -X GET https://${ENDPOINT_TLS}/authApi/tier2/v2/ping