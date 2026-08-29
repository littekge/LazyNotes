package markdown

import (
	"io"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/renderer"
	"github.com/yuin/goldmark/v2/util"
)

type Renderer = renderer.Renderer[io.Writer]

type NodeRenderer = renderer.NodeRenderer[io.Writer]

func NodeRendererFunc(f func(w io.Writer, source []byte,
	n ast.Node, entering bool, rc renderer.Context) (ast.WalkStatus, error),
) NodeRenderer {
	return renderer.NodeRendererFunc(f)
}

func WithNodeRenderers(nodeRenderers map[ast.NodeKind]NodeRenderer) Option {
	return renderer.WithNodeRenderers[io.Writer, Config](nodeRenderers)
}

type Option = renderer.Option[Config]

type Extension = renderer.Extension[Config]

func WithExtensions(extensions ...Extension) Option {
	return renderer.WithExtensions[io.Writer](extensions...)
}

type Config struct {
	renderer.Config[io.Writer, Config]
}

/* ---------- markdownRenderer ---------- */
type markdownRenderer struct {
	*renderer.Helper[io.Writer, Config]
}

func New(opts ...Option) Renderer {
	bm := NewBaseMark()
	opts = append([]Option{WithExtensions(bm)}, opts...)
	var hb renderer.HelperBuilder[io.Writer, Config]
	helper := hb.Options(opts...).Build()
	r := &markdownRenderer{
		Helper: helper,
	}
	return r
}

func (md *markdownRenderer) Render(w io.Writer, source []byte, n ast.Node, opts ...renderer.RenderOption) error {
	return md.Helper.Render(w, source, n, opts...)
}

func (r *markdownRenderer) RenderStringSource(w io.Writer, source string, n ast.Node, opts ...renderer.RenderOption) error {
	return r.Render(w, util.StringToReadOnlyBytes(source), n, opts...)
}

/* ---------- end markdownRenderer ---------- */

/* ---------- baseMark ---------- */
type baseMark struct{}

func NewBaseMark() Extension {
	return &baseMark{}
}

func (bm *baseMark) RendererOptions(cfg *Config) []Option {
	return []Option{
		WithNodeRenderers(
			map[ast.NodeKind]NodeRenderer{
				ast.KindDocument: NodeRendererFunc(bm.renderDocument),
			},
		),
	}
}

func (bm *baseMark) renderDocument(
	_ io.Writer, _ []byte, _ ast.Node, _ bool, _ renderer.Context,
) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (bm *baseMark) renderStrong(
	w io.Writer, _ []byte, n ast.Node, entering bool, _ renderer.Context,
) (ast.WalkStatus, error) {
}

/* ---------- end baseMark ---------- */
