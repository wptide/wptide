package main

import (
	"github.com/wptide/pkg/process"
	"github.com/wptide/pkg/tide"
	"github.com/wptide/pkg/message"
)

var (
	// Location of wordpress.xml ruleset.
	PHPCompatibilityWPPath = "/root/.composer/vendor/wimg/php-compatibility/framework-rulesets/wordpress.xml"
)

// Intercept intercepts the message in the pipe and makes required alterations.
type Intercept struct {
	process.Process              // Assume all functions of Process.
	In  <-chan process.Processor // The Processor to intercept.
	Out chan process.Processor   // The outbound Processor.
}

func (ix *Intercept) Run(errc *chan error) error {

	go func() {
		for {
			select {
			case in := <-ix.In:
				ix.CopyFields(in)

				codeInfo, ok := ix.Result["info"].(tide.CodeInfo)

				// If this is a theme or plugin then we need to tweak the incoming message
				// to use the WordPress PHPCompatibility standards file.
				if ok && (codeInfo.Type == "theme" || codeInfo.Type == "plugin") {
					var newAuditsArray []message.Audit
					for _, audit := range *ix.Message.Audits {
						newAudit := audit
						if audit.Type == "phpcs" && audit.Options.Standard == "phpcompatibility" {
							newAudit.Options.StandardOverride = PHPCompatibilityWPPath
						}
						newAuditsArray = append(newAuditsArray, newAudit)
					}
					ix.Message.Audits = &newAuditsArray
				}
				ix.Out <- ix
			}
		}
	}()

	return nil
}
