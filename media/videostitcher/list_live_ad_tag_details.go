// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package videostitcher

// [START videostitcher_list_live_ad_tag_details]
import (
	"context"
	"fmt"
	"io"

	"google.golang.org/api/iterator"

	stitcher "cloud.google.com/go/video/stitcher/apiv1"
	stitcherstreampb "google.golang.org/genproto/googleapis/cloud/video/stitcher/v1"
)

// listLiveAdTagDetails lists the ad tag details for the specified live session.
func listLiveAdTagDetails(w io.Writer, projectID, sessionID string) error {
	// projectID := "my-project-id"
	// sessionID := "my-session-id"
	location := "us-central1"
	ctx := context.Background()
	client, err := stitcher.NewVideoStitcherClient(ctx)
	if err != nil {
		return fmt.Errorf("stitcher.NewVideoStitcherClient: %v", err)
	}
	defer client.Close()

	req := &stitcherstreampb.ListLiveAdTagDetailsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/liveSessions/%s", projectID, location, sessionID),
	}

	it := client.ListLiveAdTagDetails(ctx, req)
	fmt.Fprintln(w, "Live ad tag details:")
	for {
		response, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("it.Next(): %v", err)
		}
		fmt.Fprintln(w, response.GetName())
	}
	return nil
}

// [END videostitcher_list_live_ad_tag_details]
