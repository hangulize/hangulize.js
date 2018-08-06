package main

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"github.com/hangulize/hangulize"
)

// packSpec packs a spec into {"spec": the-spec, "info": {"lang":, "config":,
// "test":, "source": ...}}.  It exposes some information
// to be used in JavaScript-side.
func packSpec(s *hangulize.Spec) *js.Object {
	// Pick lang, config, test, source for JavaScript-side usage.
	lang := js.Global.Get("Object").New()
	lang.Set("id", s.Lang.ID)
	lang.Set("codes", s.Lang.Codes)
	lang.Set("english", s.Lang.English)
	lang.Set("korean", s.Lang.Korean)
	lang.Set("script", s.Lang.Script)
	lang.Set("pronouncer", s.Lang.Pronouncer)

	config := js.Global.Get("Object").New()
	config.Set("authors", s.Config.Authors)
	config.Set("stage", s.Config.Stage)

	test := js.Global.Get("Array").New()
	for i, exm := range s.Test {
		o := js.Global.Get("Object").New()
		o.Set("word", exm[0])
		o.Set("transcribed", exm[1])
		test.SetIndex(i, &o)
	}

	o := js.Global.Get("Object").New()
	o.Set("lang", lang)
	o.Set("config", config)
	o.Set("test", test)
	o.Set("source", s.Source)

	// Keep the spec in $spec.
	o.Set("$spec", js.MakeWrapper(s))
	return o
}

var specs = make(map[string]*js.Object)

func init() {
	for _, lang := range hangulize.ListLangs() {
		spec, _ := hangulize.LoadSpec(lang)
		specs[lang] = packSpec(spec)
	}
}

type jsPronouncer struct {
	id        string
	pronounce *js.Object
}

func (p *jsPronouncer) ID() string {
	return p.id
}

func (p *jsPronouncer) Pronounce(word string) string {
	return p.pronounce.Invoke(word).String()
}

func main() {
	exports := map[string]interface{}{
		// hangulize is the easiest way to transcribe a loanword into Hangul.
		// Just call it with a lang ID and word.
		"hangulize": hangulize.Hangulize,

		// version is the hangulize package version.
		"version": hangulize.Version,

		// specs is pre-built {"lang": {"spec": ..., "info": ...}} object.
		// Use it to get a spec by a lang ID.
		"specs": specs,

		// parseSpec reads an HGL to pick a spec.
		"parseSpec": func(source string) *js.Object {
			r := strings.NewReader(source)
			spec, err := hangulize.ParseSpec(r)
			if err != nil {
				return nil
			}
			return packSpec(spec)
		},

		// newHangulizer wraps hangulize.NewHangulizer.  The returned object is
		// same with the Hangulizer struct.  Underlying methods are also Go
		// style.
		"newHangulizer": func(spec *js.Object) *js.Object {
			_spec := spec.Get("$spec").Interface().(*hangulize.Spec)
			h := hangulize.NewHangulizer(_spec)
			return js.MakeWrapper(h)
		},

		"usePronouncer": func(id string, pronounce *js.Object) bool {
			p := jsPronouncer{id, pronounce}
			return hangulize.UsePronouncer(&p)
		},

		"unusePronouncer": func(id string) bool {
			return hangulize.UnusePronouncer(id)
		},
	}

	js.Global.Set("hangulize", hangulize.Hangulize)

	for attr, val := range exports {
		js.Global.Get("hangulize").Set(attr, val)
	}

	if js.Module != js.Undefined {
		js.Module.Set("exports", exports)
	}
}
