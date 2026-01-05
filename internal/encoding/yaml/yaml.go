package yaml

import "sigs.k8s.io/yaml"

func Marshal(obj any) ([]byte, error) {
	return yaml.Marshal(obj)
}

func Unmarshal(yamlBytes []byte, obj any) error {
	return yaml.Unmarshal(yamlBytes, obj)
}
