package rules

import (
	"fmt"
	"regexp"

	"github.com/zclconf/go-cty/cty"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type RequiredTagsFormatRule struct {
	tflint.DefaultRule
}

func NewRequiredTagsFormatRule() *RequiredTagsFormatRule {
	return &RequiredTagsFormatRule{}
}

func (r *RequiredTagsFormatRule) Name() string {
	return "custom_required_tags_format"
}

func (r *RequiredTagsFormatRule) Enabled() bool {
	return true
}

func (r *RequiredTagsFormatRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *RequiredTagsFormatRule) Check(runner tflint.Runner) error {
	schema := &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "resource"},
		},
	}

	content, err := runner.GetModuleContent(schema, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Type != "resource" || len(block.Labels) < 2 || block.Labels[0] != "null_resource" {
			continue
		}

		attrs := block.Body.Attributes
		attr, ok := attrs["triggers"]
		if !ok {
			_ = runner.EmitIssue(r, "No se encontraron etiquetas obligatorias (triggers)", block.DefRange)
			continue
		}

		var obj map[string]cty.Value
		err := runner.EvaluateExpr(attr.Expr, &obj, nil)
		if err != nil {
			_ = runner.EmitIssue(r, "Error evaluando triggers como objeto", attr.Expr.Range())
			continue
		}

		required := []string{"Name", "Owner", "Env"}
		regex := regexp.MustCompile(`^[a-z0-9-]+$`)

		for _, key := range required {
			val, ok := obj[key]
			if !ok {
				_ = runner.EmitIssue(r, fmt.Sprintf("Falta la etiqueta obligatoria: %s", key), attr.Expr.Range())
				continue
			}
			if val.Type() != cty.String || !regex.MatchString(val.AsString()) {
				_ = runner.EmitIssue(r, fmt.Sprintf("Formato inválido en %s. Debe cumplir ^[a-z0-9-]+$", key), attr.Expr.Range())
			}
		}
	}

	return nil
}


