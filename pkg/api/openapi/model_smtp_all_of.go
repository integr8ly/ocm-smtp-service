/*
 * OCM SMTP Service API
 *
 * Manages SMTP credentials
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// SmtpAllOf struct for SmtpAllOf
type SmtpAllOf struct {
	ClusterID string    `json:"clusterID,omitempty"`
	Host      string    `json:"host,omitempty"`
	Port      string    `json:"port,omitempty"`
	Tls       string    `json:"tls,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
