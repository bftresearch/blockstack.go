// Copyright © 2017 Jack Zampolin <jack.zampolin@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/spf13/cobra"
)

// getNameBlockchainRecordCmd represents the getNameBlockchainRecord command
var getNameBlockchainRecordCmd = &cobra.Command{
	Use:   "get_name_blockchain_record [name]",
	Short: "[name]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		res, err := client.GetNameBlockchainRecord(args[0])
		handleResult(res, err)
	},
}

func init() {
	RootCmd.AddCommand(getNameBlockchainRecordCmd)
}
