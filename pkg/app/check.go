package app

import (
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

type CheckOpts struct {
	Path        string
	OutputField string
}

func NewCheckOpts() CheckOpts {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return CheckOpts{
		Path:        path,
		OutputField: "output",
	}
}

func Check(opts CheckOpts) error {
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

	for iter.Next() {
		i := iter.Value()
		v := i.LookupPath(cue.ParsePath(opts.OutputField))
		if !v.Exists() {
			continue
		}
		if v.Err() != nil {
			return v.Err()
		}

	}

	return nil
}
