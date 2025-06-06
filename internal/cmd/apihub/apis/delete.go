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

	"github.com/spf13/cobra"
)

// DelCmd to get a catalog items
var DelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an API from API Hub",
	Long:  "Delete an API from API Hub",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetRegion(region)
		return apiclient.SetApigeeOrg(org)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cmd.SilenceUsage = true
		_, err = hub.DeleteApi(apiID, force)
		return
	},
}

func init() {
	DelCmd.Flags().StringVarP(&apiID, "id", "",
		"", "API ID")
	DelCmd.Flags().BoolVarP(&force, "force", "",
		false, "force delete the API")

	_ = DelCmd.MarkFlagRequired("id")
}
