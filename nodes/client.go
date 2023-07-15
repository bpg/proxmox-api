/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package nodes

import (
	"fmt"
	"net/url"

	"github.com/bpg/proxmox-api/nodes/containers"
	"github.com/bpg/proxmox-api/nodes/tasks"
	"github.com/bpg/proxmox-api/nodes/vms"
	"github.com/bpg/proxmox-api/rest"
)

// Client is an interface for accessing the Proxmox node API.
type Client struct {
	rest.Client
	NodeName string
}

// ExpandPath expands a relative path to a full node API path.
func (c *Client) ExpandPath(path string) string {
	return fmt.Sprintf("nodes/%s/%s", url.PathEscape(c.NodeName), path)
}

// Container returns a client for managing a specific container.
func (c *Client) Container(vmID int) *containers.Client {
	return &containers.Client{
		Client: c,
		VMID:   vmID,
	}
}

// VM returns a client for managing a specific VM.
func (c *Client) VM(vmID int) *vms.Client {
	return &vms.Client{
		Client: c,
		VMID:   vmID,
	}
}

// Tasks returns a client for managing VM tasks.
func (c *Client) Tasks() *tasks.Client {
	return &tasks.Client{
		Client: c,
	}
}
