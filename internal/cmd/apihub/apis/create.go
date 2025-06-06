// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apis

import (
	"internal/apiclient"
	"internal/client/hub"
	"os"

	"github.com/spf13/cobra"
)

// CrtCmd
var CrtCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new API Hub API",
	Long:  "Create a new API Hub API",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetRegion(region)
		return apiclient.SetApigeeOrg(org)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cmd.SilenceUsage = true
		var apiFileBytes []byte

		if apiFileBytes, err = os.ReadFile(apiFilePath); err != nil {
			return err
		}
		_, err = hub.CreateApi(id, apiFileBytes)
		return
	},
	Example: `Create a new API: ` + GetExample(0),
}

var id, apiFilePath string

func init() {
	CrtCmd.Flags().StringVarP(&id, "id", "i",
		"", "API ID")
	CrtCmd.Flags().StringVarP(&apiFilePath, "file", "f",
		"", "Path to a file containing the API defintion")

	_ = CrtCmd.MarkFlagRequired("id")
	_ = CrtCmd.MarkFlagRequired("file")
}
