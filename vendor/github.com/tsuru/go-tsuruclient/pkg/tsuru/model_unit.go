/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Unit struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Appname     string `json:"appname,omitempty"`
	Processname string `json:"processname,omitempty"`
	Type        string `json:"type,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Status      string `json:"status,omitempty"`
	Address     Url    `json:"address,omitempty"`
}
