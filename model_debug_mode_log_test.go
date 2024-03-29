/*
 * Merge CRM API
 *
 * The unified API for building rich integrations with multiple CRM platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_crm_client

import (
    "testing"
)

func TestDeserializeDebugModeLog(t *testing.T) {
    var (
        rawJson     = []byte(`{"log_id": "99433219-8017-4acd-bb3c-ceb23d663832", "dashboard_view": "https://app.merge.dev/logs/99433219-8017-4acd-bb3c-ceb23d663832", "log_summary": {"url": "https://harvest.greenhouse.io/v1/candidates/", "method": "POST", "status_code": 200}}`)
        result      DebugModeLog
    )

    err := NewAPIClient(NewConfiguration()).decode(&result, rawJson, "application/json")

    if err != nil {
        t.Errorf("Failed to decode: %s", err.Error())
    }
}
