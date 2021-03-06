// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package android

import (
	"github.com/google/blueprint"
	_ "github.com/google/blueprint/bootstrap"
)

var (
	pctx = NewPackageContext("android/soong/common")

	cpPreserveSymlinks = pctx.VariableConfigMethod("cpPreserveSymlinks",
		Config.CpPreserveSymlinksFlags)

	// A phony rule that is not the built-in Ninja phony rule.  The built-in
	// phony rule has special behavior that is sometimes not desired.  See the
	// Ninja docs for more details.
	Phony = pctx.AndroidStaticRule("Phony",
		blueprint.RuleParams{
			Command:     "# phony $out",
			Description: "phony $out",
		})

	// GeneratedFile is a rule for indicating that a given file was generated
	// while running soong.  This allows the file to be cleaned up if it ever
	// stops being generated by soong.
	GeneratedFile = pctx.AndroidStaticRule("GeneratedFile",
		blueprint.RuleParams{
			Command:     "# generated $out",
			Description: "generated $out",
			Generator:   true,
		})

	// A copy rule.
	Cp = pctx.AndroidStaticRule("Cp",
		blueprint.RuleParams{
			Command:     "cp $cpPreserveSymlinks $cpFlags $in $out",
			Description: "cp $out",
		},
		"cpFlags")

	// A timestamp touch rule.
	Touch = pctx.AndroidStaticRule("Touch",
		blueprint.RuleParams{
			Command:     "touch $out",
			Description: "touch $out",
		})

	// A symlink rule.
	Symlink = pctx.AndroidStaticRule("Symlink",
		blueprint.RuleParams{
			Command:     "ln -f -s $fromPath $out",
			Description: "symlink $out",
		},
		"fromPath")

	ErrorRule = pctx.AndroidStaticRule("Error",
		blueprint.RuleParams{
			Command:     `echo "$error" && false`,
			Description: "error building $out",
		},
		"error")

	Cat = pctx.AndroidStaticRule("Cat",
		blueprint.RuleParams{
			Command:     "cat $in > $out",
			Description: "concatenate licenses $out",
		})

	WriteFile = pctx.AndroidStaticRule("WriteFile",
		blueprint.RuleParams{
			Command:     "echo '$content' > $out",
			Description: "writing file $out",
		},
		"content")

	// Used only when USE_GOMA=true is set, to restrict non-goma jobs to the local parallelism value
	localPool = blueprint.NewBuiltinPool("local_pool")
)

func init() {
	pctx.Import("github.com/google/blueprint/bootstrap")
}
