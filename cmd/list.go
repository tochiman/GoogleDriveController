/*
Copyright © 2023 tochiman development@tochiman.com
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/crackcomm/go-clitable"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"

	"github.com/tochiman/DriveManegement/exe"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "This command checks list of file or directory",
	Long: `This command checks list of file or directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		srv, err := drive.NewService(ctx)
		if err != nil {
			fmt.Printf("Unable to retrieve Drive client: %v", err)
		}


		for {
			r, err := srv.Files.List().PageSize(1000).
			Fields("nextPageToken, files(id, name, fileExtension, size, mimeType)").
			Context(ctx).
			Q(fmt.Sprintf("name contains '%s'", query)).Do()
			if err != nil {
				fmt.Printf("Unable to retrieve files: %v", err)
			}

			table := clitable.New([]string{"ID","Name", "Extention", "Size"})
			for _, f := range r.Files {
				if f.FileExtension == "" { f.FileExtension = "dir" }
				size := exe.Conversion(float64(f.Size))
				table.AddRow(map[string]interface{}{"ID":f.Id, "Name": f.Name, "Extention": f.FileExtension, "Size": size})
			}
			table.Print()

			paging = r.NextPageToken
			if len(paging) == 0 {
				break
			}

		}
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&query, "query", "q","", "This flag specifies the file name")
	// listCmd.Flags().StringVarP(&extention, "extention", "e", "", "Specify file-Extention type when downloading. ")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}