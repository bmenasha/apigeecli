// Copyright 2020 Google LLC
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

package env

import (
	"fmt"
	"internal/apiclient"
	"internal/clilog"

	environments "internal/client/env"

	"github.com/spf13/cobra"
)

// SetAxCmd to set role on env
var SetAxCmd = &cobra.Command{
	Use:   "setax",
	Short: "Set Analytics Agent role for a member on an environment",
	Long:  "Set Analytics Agent role for a member an Environment",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if role != "analyticsAgent" && role != "analyticsViewer" {
			return fmt.Errorf("invalid memberRole. Member role must be analyticsViewer or analyticsAgent")
		}
		apiclient.SetApigeeEnv(environment)
		apiclient.SetRegion(region)
		return apiclient.SetApigeeOrg(org)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cmd.SilenceUsage = true

		err = environments.SetIAM(memberName, role, memberType)
		if err != nil {
			return err
		}
		clilog.Info.Printf("Member %s granted access to %s role\n", memberName, role)
		return nil
	},
}

func init() {
	SetAxCmd.Flags().StringVarP(&memberName, "name", "n",
		"", "Member Name, example Service Account Name")
	SetAxCmd.Flags().StringVarP(&memberType, "member-type", "m",
		"serviceAccount", "memberType must be serviceAccount, user or group")
	SetAxCmd.Flags().StringVarP(&role, "member-role", "",
		"analyticsAgent", "memberRole must be analyticsViewer or analyticsAgent")
	_ = SetAxCmd.MarkFlagRequired("name")
}
