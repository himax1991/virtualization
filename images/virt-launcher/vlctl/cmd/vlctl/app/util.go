/*
Copyright 2025 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"fmt"
	"os"
)

func marshalAndPrintOutput(baseOpts *BaseOptions, data interface{}) error {
	b, err := baseOpts.MarshalOutput(data)
	if err != nil {
		return fmt.Errorf("failed to marshal output: %w", err)
	}

	_, err = fmt.Fprintf(os.Stdout, "%s\n", string(b))
	return err
}
