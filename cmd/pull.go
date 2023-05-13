/*
Copyright © 2023 tochiman development@tochiman.com
*/
package cmd

import (
	"os"
	"io"
	"fmt"
	"log"
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"time"
    "github.com/schollz/progressbar"
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
			Fields("files(id, name, mimeType)").Context(ctx).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		count := len(r.Files)
		count64 := int64(count)
		bar := progressbar.Default(count64)
		for i := 0; i < count; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
		}
		
		for _, f := range r.Files {
			if f.MimeType == "application/vnd.google-apps.folder" {
				// フォルダの場合はスキップ
				continue
			}
	
			if err := download(ctx, srv, f.Name, f.Id); err != nil {
				log.Fatalf("Unable to download: %v", err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func download(ctx context.Context, srv *drive.Service, name, id string) error {
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