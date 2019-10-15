// Code generated by pluginator on LabelTransformer; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/transformers"
	"sigs.k8s.io/kustomize/v3/pkg/transformers/config"
	"sigs.k8s.io/yaml"
)

// Add the given labels to the given field specifications.
type LabelTransformerPlugin struct {
	Labels     map[string]string  `json:"labels,omitempty" yaml:"labels,omitempty"`
	FieldSpecs []config.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *LabelTransformerPlugin) Config(
	h *resmap.PluginHelpers, c []byte) (err error) {
	p.Labels = nil
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *LabelTransformerPlugin) Transform(m resmap.ResMap) error {
	t, err := transformers.NewMapTransformer(
		p.FieldSpecs,
		p.Labels,
	)
	if err != nil {
		return err
	}
	return t.Transform(m)
}

func NewLabelTransformerPlugin() resmap.TransformerPlugin {
	return &LabelTransformerPlugin{}
}
