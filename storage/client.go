/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package storage

import "github.com/bpg/proxmox-api/rest"

// Client is an interface for accessing the Proxmox storage API.
type Client struct {
	rest.Client
}
