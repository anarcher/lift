package app

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
)

type BuildOpts struct {
	Path        string
	OutputField string
}

func NewBuildOpts() BuildOpts {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return BuildOpts{
		Path:        path,
		OutputField: "output",
	}
}

func Build(opts BuildOpts) error {
	lc := &load.Config{
		Dir: opts.Path,
	}
	binst := load.Instances([]string{}, lc)
	insts := cue.Build(binst)
	inst := cue.Merge(insts...)
	if inst.Err != nil {
		return inst.Err
	}

	if err := inst.Value().Validate(cue.All(), cue.Concrete(true)); err != nil {
		return err
	}

	iter, err := inst.Value().Fields()
	if err != nil {
		return err
	}

	var idx int
	for iter.Next() {
		i := iter.Value()
		v := i.LookupPath(cue.ParsePath(opts.OutputField))
		if !v.Exists() {
			continue
		}
		if v.Err() != nil {
			return v.Err()
		}

		if idx > 0 {
			fmt.Fprintln(os.Stdout, "---")
		}

		if err := printValue(v); err != nil {
			return err
		}
		idx++
	}

	return nil
}
func printValue(val cue.Value) error {
	if val.Kind() == cue.ListKind {
		iter, err := val.List()
		if err != nil {
			return err
		}

		bs, err := yaml.EncodeStream(iter)
		if err != nil {
			return err
		}
		fmt.Fprint(os.Stdout, string(bs))

	} else {
		bs, err := yaml.Encode(val)
		if err != nil {
			return err
		}
		fmt.Fprint(os.Stdout, string(bs))
	}

	return nil
}
