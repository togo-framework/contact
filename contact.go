// Package contact is the Contact content pack for togo: ContactSubmission resource (spam/read flags, audit) + a contact form.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package contact

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("contact", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
