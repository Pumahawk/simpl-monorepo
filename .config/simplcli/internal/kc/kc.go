package kc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"simplcli/internal/ex"
	"time"
)

var NotFound = fmt.Errorf("not found")

func Wait(t time.Duration) error {
	cmd := ex.New(
		"kubectl",
		"wait", "-n", "local-authority",
		"--for=jsonpath={.status.availableReplicas}=1",
		fmt.Sprintf("--timeout=%ds", int(t.Seconds())),
		"statefulset.apps/local-authority-keycloak",
	)
	return cmd.Run()
}

func Login() error {
	cmd := kexec(`
    /opt/keycloak/bin/kcadm.sh config credentials \
      --config /tmp/kcadm.config \
      --server http://localhost:8080 \
      --realm master \
      --client admin-cli \
      --user admin \
      --password admin`)
	return cmd.Run()
}

func GetSecret(realm, clientId string) (string, error) {
	out := &bytes.Buffer{}
	cmd := kexec(`
	  realm="$1"
	  client_id="$2"
    /opt/keycloak/bin/kcadm.sh get clients \
      -r "$realm" \
      -q clientId="$client_id" \
      --config /tmp/kcadm.config`, realm, clientId)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Decode expected json output
	outs := out.String()
	type rt struct{ Secret string }
	outj := make([]rt, 0, 1)
	if err := json.NewDecoder(out).Decode(&outj); err != nil {
		return "", fmt.Errorf("decode response get clients %q: %w", outs, err)
	}

	if len(outj) == 0 {
		return "", NotFound
	} else if len(outj) > 1 {
		return "", fmt.Errorf("expected 1 secret, got %d", len(outj))
	}
	return outj[0].Secret, nil
}

func kexec(script string, args ...string) *exec.Cmd {
	return ex.New("kubectl", append([]string{
		"exec",
		"-n", "local-authority",
		"local-authority-keycloak-0",
		"--",
		"bash", "-e", "-c", script, "--"}, args...)...)
}
