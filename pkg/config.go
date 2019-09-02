package pkg

// FIXME must be part of the configuration in the target deployment. (ConfigMap)
var (
	SecretReferenceOfPartition = map[string]string{
		"fra-equ01": "fra-equ01-seed",
		"nbg-w8101": "seed-nbg-gardener-test-01",
	}

	NetworksOfPartition = map[string]map[string]string{
		"fra-equ01": {
			"internet": "internet-fra-equ01",
		},
		"nbg-w8101": {
			"internet": "internet-nbg-w8101",
		},
	}

	DefaultMachineTypesOfPartition = map[string]map[string]string{
		"fra-equ01": {
			"firewall": "c1-xlarge-x86",
			"worker":   "c1-large-x86",
		},
		"nbg-w8101": {
			"firewall": "c1-xlarge-x86",
			"worker":   "c1-xlarge-x86",
		},
	}
)
