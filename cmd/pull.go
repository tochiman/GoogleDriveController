/*
Copyright © 2023 tochiman development@tochiman.com
*/
package cmd

import (
	"os"
	"io"
	"fmt"
	"log"
	"sync"
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Able to download specified files or folders",
	Long: `Able to download specified files or folders`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		srv, err := drive.NewService(ctx)
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		r, err := srv.Files.List().PageSize(1000).
			Fields("files(id, name, mimeType)").
			Q(fmt.Sprintf("'%s' in parents or name = '%s'", folder, query)).
			Context(ctx).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		var wg sync.WaitGroup
		
		for _, f := range r.Files {
			if f.MimeType == "application/vnd.google-apps.folder" {
				continue //フォルダーの場合はスキップ
			}
			wg.Add(1)
			go func(f *drive.File){
				defer wg.Done()
				if err := download(ctx, srv, f.Name, f.Id); err != nil {
					log.Fatalf("Unable to download: %v", err)
				}
			}(f)
		}
		wg.Wait()
		
	},
	
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringVarP(&query, "query", "q", "", "specifies the file name")
	pullCmd.Flags().StringVarP(&path, "path", "p", "", "Specify file path when downloading")
	pullCmd.Flags().StringVarP(&folder, "folder", "f", "", "Specify folder-id when downloading.")
		
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func download(ctx context.Context, srv *drive.Service, name, id string) error {
	if path != "" {
		name = path + name
	}
	create, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer create.Close()

	resp, err := srv.Files.Get(id).Context(ctx).Download()
	if err != nil {
		return fmt.Errorf("get drive file: %w", err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(create, resp.Body); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}