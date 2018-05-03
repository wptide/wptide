package process

import (
	"github.com/wptide/pkg/process"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/storage"
	"github.com/wptide/pkg/payload"
	"errors"
	"strings"
	"github.com/wptide/pkg/tide"
)

var (
	//PHPCompatibilityWPPath = "/root/.composer/vendor/wimg/php-compatibility/framework-rulesets/wordpress.xml"
	PHPCompatibilityWPPath = "/Users/rheinardkorf/Utilities/phpcs-standards/phpcompatibility/framework-rulesets/wordpress.xml"
)

func DoIngest(msg *message.Message, result *process.Result, tempFolder string) error {
	proc := process.Ingest{
		TempFolder: tempFolder,
	}
	proc.Result = result
	proc.Message = *msg

	return proc.Do()
}

func DoInfo(msg *message.Message, result *process.Result) error {
	proc := process.Info{}
	proc.Result = result
	proc.Message = *msg

	return proc.Do()
}

func DoLighthouse(msg *message.Message, result *process.Result, tempFolder string, storageProvider storage.StorageProvider) error {
	proc := process.Lighthouse{
		TempFolder: tempFolder,
		StorageProvider: storageProvider,
	}
	proc.Result = result
	proc.Message = *msg

	return proc.Do()
}

func DoPhpcs(msg *message.Message, result *process.Result, tempFolder string, storageProvider storage.StorageProvider, config map[string]interface{}) error {

	proc := &process.Phpcs{
		TempFolder: tempFolder,
		Config: config,
		StorageProvider: storageProvider,
	}
	proc.Result = result

	res := *result
	codeInfo, ok := res["info"].(tide.CodeInfo)

	// If this is a theme or plugin then we need to tweak the incoming message
	// to use the WordPress PHPCompatibility standards file.
	if ok && (codeInfo.Type == "theme" || codeInfo.Type == "plugin") {
		var newAuditsArray []message.Audit
		for _, audit := range *msg.Audits {
			newAudit := audit
			if audit.Type == "phpcs" && audit.Options.Standard == "phpcompatibility" {
				newAudit.Options.StandardOverride = PHPCompatibilityWPPath
			}
			newAuditsArray = append(newAuditsArray, newAudit)
		}
		*msg.Audits = newAuditsArray
	}

	proc.Message = *msg

	errs := []string{}
	for _, audit := range *proc.Message.Audits {
		if audit.Type == "phpcs" {
			if err := proc.Do(audit); err != nil {
				errs = append(errs, "PHPCS Error: " + err.Error())
			}
		}
	}

	concatErrs := strings.Join(errs,"\n")

	if concatErrs == "" {
		return nil
	} else {
		return errors.New(concatErrs)
	}
}

func DoResponse(msg *message.Message, result *process.Result, payloaders map[string]payload.Payloader) error {
	proc := process.Response{
		Payloaders: payloaders,
	}
	proc.Result = result
	proc.Message = *msg

	return proc.Do()
}
