/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package cluster

import (
	"fmt"

	clusterfirewall "github.com/bpg/proxmox-api/cluster/firewall"
	"github.com/bpg/proxmox-api/firewall"
	"github.com/bpg/proxmox-api/rest"
)

// Client is an interface for accessing the Proxmox cluster API.
type Client struct {
	rest.Client
}

// ExpandPath expands a relative path to a full cluster API path.
func (c *Client) ExpandPath(path string) string {
	return fmt.Sprintf("cluster/%s", path)
}

// Firewall returns a client for managing the cluster firewall.
func (c *Client) Firewall() clusterfirewall.API {
	return &clusterfirewall.Client{
		Client: firewall.Client{Client: c},
	}
}
