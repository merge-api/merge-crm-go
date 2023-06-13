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

func TestDeserializeAccountDetailsAndActions(t *testing.T) {
    var (
        rawJson     = []byte(`{"id": "e59b1821-f85c-4e28-a6b3-1804156f3563", "category": "hris", "status": "COMPLETE", "status_detail": null, "end_user_origin_id": "3ac95cde-6c7f-4eef-afec-be710b42308d", "end_user_organization_name": "Foo Bar, LLC", "end_user_email_address": "hradmin@foobar.dev", "webhook_listener_url": "https://api.merge.dev/api/integrations/webhook-listener/7fc3mee0UW8ecV4", "is_duplicate": true, "integration": {"name": "SAP SuccessFactors", "categories": ["hris", "ats"], "image": "https://cdn.merge.dev/SuccessFactors_Logo.png", "square_image": "https://cdn.merge.dev/SuccessFactors_Square_Logo.jpg", "color": "#F6A704", "slug": "sap-successfactors", "passthrough_available": true, "available_model_operations": [{"model_name": "Candidate", "available_operations": ["FETCH", "CREATE"], "required_post_parameters": ["remote_user_id"], "supported_fields": ["first_name", "last_name", "company", "title"]}]}}`)
        result      AccountDetailsAndActions
    )

    err := NewAPIClient(NewConfiguration()).decode(&result, rawJson, "application/json")

    if err != nil {
        t.Errorf("Failed to decode: %s", err.Error())
    }
}
