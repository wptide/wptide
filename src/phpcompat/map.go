package phpcompat

var PhpCompatibilityMap = map[string]*Compatibility{

	"PHPCompatibility.PHP.NewFunctions.intdivFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intdivFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_least_recently_usedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_least_recently_usedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_block_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_block_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.datetimeimmutableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.datetimeimmutableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_ssl_enginesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_ssl_enginesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_chunk_length_penalty_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_chunk_length_penalty_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_1Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_1Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_unescapeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_unescapeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.inflate_initFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.inflate_initFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_dhfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_dhfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.password_argon2iFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.password_argon2iFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_sasl_usernameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_sasl_usernameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_stoppedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_stoppedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.stream_context_get_paramsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.stream_context_get_paramsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_use_daylight_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_use_daylight_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_entropy_fileRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_entropy_fileRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.1.13",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_ecbDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_ecbDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_traitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_traitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_crop_whiteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_crop_whiteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_ftp_entry_pathFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_ftp_entry_pathFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocisavelobfileDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocisavelobfileDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_modeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_modeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_http_version_2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_http_version_2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_actual_maximumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_actual_maximumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_localeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_localeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"Generic.NamingConventions.CamelCapsFunctionName.MethodDoubleUnderscore": &Compatibility{
		Source: "Generic.NamingConventions.CamelCapsFunctionName.MethodDoubleUnderscore",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.call_user_methodDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.call_user_methodDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.assertionerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.assertionerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_sigioFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_sigioFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_utf8Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_utf8Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.2",
			Reported: "5.3.2",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_join_source_groupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_join_source_groupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_fd_setsizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_fd_setsizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.filter_flag_email_unicodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.filter_flag_email_unicodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_get_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_get_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.xsl_security_prefsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.xsl_security_prefsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_blobinfileRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_blobinfileRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.splitiDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.splitiDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_majorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_majorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.endforeachFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.endforeachFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.hash_copyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.hash_copyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.iterableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.iterableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_exec_dirDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_exec_dirDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_zero_paddingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_zero_paddingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_notice_allFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_notice_allFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_dirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_dirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_keyword_values_for_localeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_keyword_values_for_localeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_exop_whoamiFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_exop_whoamiFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagegetclipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagegetclipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_slowestRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_slowestRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.InternalInterfaces.throwableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.InternalInterfaces.throwableFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mbstring_http_output_conv_mimetypeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mbstring_http_output_conv_mimetypeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_detect_unicodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_detect_unicodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlcalendarFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlcalendarFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_maxpathlenFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_maxpathlenFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_expect_100_timeout_msFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_expect_100_timeout_msFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.xsl_security_prefsRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.xsl_security_prefsRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.0.26",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.sql_regcaseDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.sql_regcaseDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.preg_filterFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.preg_filterFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.get_declared_traitsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.get_declared_traitsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.assert_exceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.assert_exceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_round_half_downFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_round_half_downFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.snmp_oid_output_suffixFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.snmp_oid_output_suffixFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__unsetMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__unsetMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.PregReplaceEModifier.Removed": &Compatibility{
		Source: "PHPCompatibility.PHP.PregReplaceEModifier.Removed",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_256_cbcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_256_cbcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_leading_combining_markFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_leading_combining_markFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_push_okFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_push_okFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.ircgRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.ircgRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.objectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_decryptDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_decryptDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_affine_translateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_affine_translateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_quick_checkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_quick_checkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_display_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_display_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.intl_idna_variant_2003Deprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.intl_idna_variant_2003Deprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.session_registerDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.session_registerDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splpriorityqueueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splpriorityqueueFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_bigint_as_stringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_bigint_as_stringFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_redir_post_302Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_redir_post_302Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlftp_create_dir_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlftp_create_dir_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproto_smbFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproto_smbFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ConstantArraysUsingConst.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ConstantArraysUsingConst.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundSuperglobal": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundSuperglobal",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_idleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_idleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RequiredOptionalFunctionParameters.preg_match_all_matchesMissing": &Compatibility{
		Source: "PHPCompatibility.PHP.RequiredOptionalFunctionParameters.preg_match_all_matchesMissing",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_connect_pollFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_connect_pollFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_addrinfo_explainFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_addrinfo_explainFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_minorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_minorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_crop_transparentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_crop_transparentFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_blackmanFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_blackmanFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.InvalidTypeHintFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.InvalidTypeHintFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_is_in_sessionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_is_in_sessionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewArrayStringDereferencing.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewArrayStringDereferencing.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociplogonDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociplogonDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_initDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_initDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_version_http2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_version_http2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.23",
			Reported: "5.5.23",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_dml_escapeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_dml_escapeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_sasl_nocanonFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_sasl_nocanonFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_control_paged_result_responseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_control_paged_result_responseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.always_populate_raw_post_dataDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.always_populate_raw_post_dataDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_max_recv_speed_largeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_max_recv_speed_largeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_stdflagsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_stdflagsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_nontransitional_to_unicodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_nontransitional_to_unicodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.joaatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.joaatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.InternalInterfaces.traversableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.InternalInterfaces.traversableFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewClasses.sessionhandlerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.sessionhandlerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_bug_compat_warnDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_bug_compat_warnDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.iterableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.iterableFound",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.encodingInvalidValueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.encodingInvalidValueFound",
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.pcre_jitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.pcre_jitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociparseDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociparseDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsbboxRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsbboxRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_parsehugeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_parsehugeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.11",
			Reported: "5.2.11",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_sp_majorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_sp_majorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_yieldFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_yieldFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.ypRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.ypRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.forward_static_call_arrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.forward_static_call_arrayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_spki_verifyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_spki_verifyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_fetchDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_fetchDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connect_asyncFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connect_asyncFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_rtprioFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_rtprioFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_intervalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_intervalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_float_epsilonFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_float_epsilonFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.get_resourcesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.get_resourcesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_sybaseDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_sybaseDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollgetelemDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollgetelemDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splmaxheapFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splmaxheapFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.scandir_sort_ascendingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.scandir_sort_ascendingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_affine_scaleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_affine_scaleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mssqlRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mssqlRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.constFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.constFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.fnv164Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.fnv164Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.max_input_varsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.max_input_varsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.8",
			Reported: "5.3.8",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.8",
			Reported:   "5.3.8",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_collect_memory_statisticsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_collect_memory_statisticsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociloadlobDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociloadlobDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.catchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.catchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_opt_int_and_float_nativeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_opt_int_and_float_nativeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewUseConstFunction.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewUseConstFunction.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pcntl_sigtimedwaitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pcntl_sigtimedwaitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.numericFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.numericFound",
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.2.1",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mail_add_x_headerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mail_add_x_headerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.oci_no_auto_commitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.oci_no_auto_commitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.1",
			Reported: "5.3.1",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_modeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_modeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlbreakiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlbreakiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.newFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.newFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.cli_promptFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.cli_promptFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.log1pFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.log1pFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_get_links_statsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_get_links_statsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_runtimeDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_runtimeDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.voidFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.voidFound",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenCallTimePassByReference.NotAllowed": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenCallTimePassByReference.NotAllowed",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.tcp_nodelayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.tcp_nodelayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_nt_serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_nt_serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_cmsg_spaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_cmsg_spaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.allow_call_time_pass_referenceDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.allow_call_time_pass_referenceDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.insteadofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.insteadofFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_diffFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_diffFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.encodingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.encodingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.curl_cainfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.curl_cainfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.6",
			Reported: "5.3.6",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.6",
			Reported:   "5.3.6",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifetchintoDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifetchintoDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_encryptDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_encryptDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_coprocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_coprocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_startedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_startedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_suggestFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_suggestFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.stream_set_chunk_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.stream_set_chunk_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocibindbynameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocibindbynameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.eachDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.eachDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.defaultFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.defaultFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysql_escape_stringDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysql_escape_stringDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlrulebasedbreakiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlrulebasedbreakiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_session_disabledFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_session_disabledFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_escape_dnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_escape_dnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intl_get_error_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intl_get_error_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_share_strerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_share_strerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.InternalInterfaces.datetimeinterfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.InternalInterfaces.datetimeinterfaceFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_touchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_touchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mcrypt_modes_dirDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mcrypt_modes_dirDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.sql_safe_modeRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.sql_safe_modeRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.2.0",
			High: "7.2.1",
			Reported: "7.2",
			MajorMinor: "7.2",
		},
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._getFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._getFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.EmptyNonVariable.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.EmptyNonVariable.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4.45",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_ofbDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_ofbDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_128_cbcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_128_cbcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_hostRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_hostRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curle_sshFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curle_sshFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.acoshFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.acoshFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_freeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_freeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.ripemd320Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.ripemd320Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewInterfaces.jsonserializableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewInterfaces.jsonserializableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociwritelobtofileDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociwritelobtofileDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.ui_exception_invalidargumentexceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.ui_exception_invalidargumentexceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_any_clientFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_any_clientFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_full_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_full_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__dir__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__dir__Found",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.hash_pbkdf2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.hash_pbkdf2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.md2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.md2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_count_equivalent_idsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_count_equivalent_idsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_least_trafficRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_least_trafficRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocisetprefetchDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocisetprefetchDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_disabledFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_disabledFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_coreFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_coreFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_set_repeated_wall_time_optionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_set_repeated_wall_time_optionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_exop_passwdFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_exop_passwdFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.salsa10Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.salsa10Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltdivFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltdivFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.u_idna_domain_name_too_long_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.u_idna_domain_name_too_long_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.json_last_error_msgFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.json_last_error_msgFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctionParameters.dirname_levelsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctionParameters.dirname_levelsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.oci_cred_extFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.oci_cred_extFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_intdivFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_intdivFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.implementsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.implementsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.implementsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.implementsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_time_zoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_time_zoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_error_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_error_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.password_argon2_default_time_costFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.password_argon2_default_time_costFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_free_dictFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_free_dictFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_checkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_checkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imageopenpolygonFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imageopenpolygonFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__tostringMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__tostringMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__issetMethodStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__issetMethodStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_hermiteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_hermiteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_yield_fromFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_yield_fromFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.abstractFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.abstractFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.intl_use_exceptionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.intl_use_exceptionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_collect_statisticsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_collect_statisticsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_get_metadataDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_get_metadataDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_trailing_hyphenFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_trailing_hyphenFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_algo_rmd160Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_algo_rmd160Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.7",
			Reported: "5.4.7",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.globalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.globalFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagecropautoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagecropautoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_least_maximumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_least_maximumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.define_syslog_variablesDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.define_syslog_variablesDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.odbc_default_cursortypeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.odbc_default_cursortypeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_dirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_dirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_encoding_rawFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_encoding_rawFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifetchstatementDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifetchstatementDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_algorithm_modeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_algorithm_modeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_hyphen_3_4Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_hyphen_3_4Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_create_from_rulesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_create_from_rulesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_list_idsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_list_idsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_afterFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_afterFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.fileinfo_compressRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.fileinfo_compressRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctionArrayDereferencing.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctionArrayDereferencing.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.transliteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.transliteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bsplineFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bsplineFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_cseq_recvFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_cseq_recvFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.pfproRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.pfproRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.asFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.asFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_subFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_subFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundThis": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundThis",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_textasvarcharRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_textasvarcharRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.pgsql_attr_disable_native_prepared_statementRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.pgsql_attr_disable_native_prepared_statementRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_iv_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_iv_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.doFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.doFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ValidIntegers.BinaryIntegerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ValidIntegers.BinaryIntegerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.header_register_callbackFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.header_register_callbackFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gmp_rootremFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gmp_rootremFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.sqlite3_extension_dirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.sqlite3_extension_dirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.10",
			Reported: "5.3.10",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.10",
			Reported:   "5.3.10",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_start_nowdocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_start_nowdocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_share_errnoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_share_errnoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_http_outputDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_http_outputDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_round_half_upFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_round_half_upFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_contiunedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_contiunedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_algo_sha384Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_algo_sha384Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.7",
			Reported: "5.4.7",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gc_enableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gc_enableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._sessionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._sessionFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocirowcountDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocirowcountDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.callableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.callableFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_get_connection_statusFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_get_connection_statusFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_read_async_queryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_read_async_queryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__issetMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__issetMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_cipher_nameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_cipher_nameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlftp_create_dirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlftp_create_dirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.dbxRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.dbxRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.finallyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.finallyFound",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_allowed_env_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_allowed_env_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_internal_encodingDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_internal_encodingDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewHeredocInitialize.propertyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHeredocInitialize.propertyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_privateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_privateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_depthFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_depthFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ipproto_ipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ipproto_ipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.snmpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.snmpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bicubic_fixedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bicubic_fixedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_contFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_contFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.shm_has_varsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.shm_has_varsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.deflate_initFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.deflate_initFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagebmpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagebmpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fileinfo_mime_typeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fileinfo_mime_typeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_escape_identifierFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_escape_identifierFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.callbackfilteriteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.callbackfilteriteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_send_contFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_send_contFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.SelfOutsideClassScopeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.SelfOutsideClassScopeFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedNewReference.Forbidden": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedNewReference.Forbidden",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqli_rollback_on_cached_plinkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqli_rollback_on_cached_plinkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.5.37",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewHeredocInitialize.constFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHeredocInitialize.constFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_groupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_groupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.header_removeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.header_removeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.zlib_encodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.zlib_encodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_socketFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_socketFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_allow_persistentRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_allow_persistentRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.spldoublylinkedlistFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.spldoublylinkedlistFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_ns_separatorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_ns_separatorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.debug_backtrace_ignore_argsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.debug_backtrace_ignore_argsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.5",
			Reported: "5.3.5",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.namespaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.namespaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewInterfaces.sessionhandlerinterfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewInterfaces.sessionhandlerinterfaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_pow_equalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_pow_equalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.strict_typesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.strict_typesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltovfFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltovfFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_flip_horizontalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_flip_horizontalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_condition_unmetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_condition_unmetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_pipewaitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_pipewaitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_abortFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_abortFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewClasses.dateintervalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.dateintervalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_xoauth2_bearerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_xoauth2_bearerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_exop_modify_passwdFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_exop_modify_passwdFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.CaseSensitiveKeywords.NonLowercaseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.CaseSensitiveKeywords.NonLowercaseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_get_cache_statsRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_get_cache_statsRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_addrinfo_bindFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_addrinfo_bindFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumntypeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumntypeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocierrorDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocierrorDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_session_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_session_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.sys_temp_dirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.sys_temp_dirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicancelDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicancelDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_0Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_0Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.18",
			Reported: "5.5.18",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_blockFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_blockFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_float_maxFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_float_maxFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_effect_multiplyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_effect_multiplyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_file_createFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_file_createFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_create_time_zoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_create_time_zoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.set_magic_quotes_runtimeDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.set_magic_quotes_runtimeDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splheapFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splheapFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.bus_adralnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.bus_adralnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_trappedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_trappedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_sendFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_sendFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ValidIntegers.HexNumericStringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ValidIntegers.HexNumericStringFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.finfo_bufferFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.finfo_bufferFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._postFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._postFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_mempool_default_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_mempool_default_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_param_countDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_param_countDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_addFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_addFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_transliterateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_transliterateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_regionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_regionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_spki_exportFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_spki_exportFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_output_encodingDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_output_encodingDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_prefixFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_prefixFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.e_deprecatedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.e_deprecatedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_punycodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_punycodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlheader_separateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlheader_separateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.caseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.caseFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_use_strict_modeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_use_strict_modeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.1",
			Reported: "5.5.1",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.5.1",
			Reported:   "5.5.1",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.abstractFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.abstractFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.boolFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.boolFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mingRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mingRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollappendDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollappendDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.file_binaryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.file_binaryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_notice_clearFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_notice_clearFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__namespace__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__namespace__Found",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mcrypt_algorithms_dirDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mcrypt_algorithms_dirDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._filesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._filesFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_notice_lastFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_notice_lastFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedAlternativePHPTags.ScriptOpenTagFound": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedAlternativePHPTags.ScriptOpenTagFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.breakFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.breakFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_oldestRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_oldestRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_algorithmDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_is_block_algorithmDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_object_as_arrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_object_as_arrayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_algo_sha256Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_algo_sha256Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.7",
			Reported: "5.4.7",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_headeroptFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_headeroptFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenFunctionParametersWithSameName.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenFunctionParametersWithSameName.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_spki_newFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_spki_newFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__unsetMethodStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__unsetMethodStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.track_errorsDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.track_errorsDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_endDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_endDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.6.32",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifreecollectionDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifreecollectionDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.parseerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.parseerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewMultiCatch.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewMultiCatch.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pcre_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pcre_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ip_multicast_ttlFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ip_multicast_ttlFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_quadraticFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_quadraticFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_recv_contFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_recv_contFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_parse_from_formatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_parse_from_formatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_get_cert_locationsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_get_cert_locationsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gd_extra_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gd_extra_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_buildFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_buildFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_proxy_service_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_proxy_service_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_version_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_version_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_beforeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_beforeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_skipped_wall_time_optionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_skipped_wall_time_optionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_round_half_evenFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_round_half_evenFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_biglinesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_biglinesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_version_pslFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_version_pslFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_ssl_falsestartFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_ssl_falsestartFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenBreakContinueVariableArguments.zeroArgumentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenBreakContinueVariableArguments.zeroArgumentFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_get_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_get_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_field_differenceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_field_differenceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHeredocInitialize.staticvarFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHeredocInitialize.staticvarFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewNullableTypes.typeDeclarationFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewNullableTypes.typeDeclarationFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imap_utf8_to_mutf7Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imap_utf8_to_mutf7Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_list_dictsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_list_dictsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.floatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.floatFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.cli_server_colorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.cli_server_colorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splminheapFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splminheapFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_session_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_session_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_ssl_startupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_ssl_startupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlheader_unifiedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlheader_unifiedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedTypeCasts.t_unset_castDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedTypeCasts.t_unset_castDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.finfo_fileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.finfo_fileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_1_clientFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_1_clientFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_ns_separatorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_ns_separatorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.asinhFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.asinhFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.intFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.intFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.floatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.floatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlftp_create_dir_retryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlftp_create_dir_retryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_namespaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_namespaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClosure.ThisFoundInStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClosure.ThisFoundInStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.cpdfRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.cpdfRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_server_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_server_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_yieldFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_yieldFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_sid_lengthFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_sid_lengthFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.0.26",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollsizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollsizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.soap_ssl_method_tlsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.soap_ssl_method_tlsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_session_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_session_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.foreachFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.foreachFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.iterableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.iterableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_setenvFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_setenvFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_max_pipeline_lengthFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_max_pipeline_lengthFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_offsetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_offsetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_invalid_ace_labelFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_invalid_ace_labelFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.enddeclareFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.enddeclareFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_has_same_rulesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_has_same_rulesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.stringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.stringFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_resetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_resetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_primary_portFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_primary_portFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.6",
			Reported: "5.4.6",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_client_cseqFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_client_cseqFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_rttimeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_rttimeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlauth_negotiateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlauth_negotiateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_push_denyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_push_denyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.stringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.stringFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_sync_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_sync_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.fbsqlRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.fbsqlRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.catchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.catchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.array_replaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.array_replaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_begin_transactionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_begin_transactionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_parse_exopFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_parse_exopFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_shut_rdFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_shut_rdFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_ztsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_ztsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_opt_ssl_verify_server_certFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_opt_ssl_verify_server_certFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.staticFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.staticFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_consume_inputFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_consume_inputFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.y2k_complianceDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.y2k_complianceDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_input_encodingDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_input_encodingDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.zend_logo_guidDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.zend_logo_guidDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.4.45",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ldap_sortDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ldap_sortDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.2.1",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_local_ipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_local_ipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.6",
			Reported: "5.4.6",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bicubicFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bicubicFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_memlockFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_memlockFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_redir_post_303Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_redir_post_303Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_dns_interfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_dns_interfaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fileinfo_extensionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fileinfo_extensionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_set_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_set_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.falseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.falseFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_callableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_callableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_share_closeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_share_closeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_cmd_buffer_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_cmd_buffer_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.session_unregisterDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.session_unregisterDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_algorithms_nameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_algorithms_nameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_escape_filterFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_escape_filterFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewClasses.recursivecallbackfilteriteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.recursivecallbackfilteriteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_packageFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_packageFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.finalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.finalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_create_inverseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_create_inverseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.pharFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.pharFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.file_textFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.file_textFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_label_has_dotFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_label_has_dotFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_polling_writingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_polling_writingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_nprocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_nprocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_peerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_peerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_typeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_typeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_tz_data_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_tz_data_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnisnullDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnisnullDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocilogonDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocilogonDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mdecrypt_genericDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mdecrypt_genericDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_rssFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_rssFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.intl_error_levelFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.intl_error_levelFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlsslopt_allow_beastFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlsslopt_allow_beastFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.endforFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.endforFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.array_columnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.array_columnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_strerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_strerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__sleepMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__sleepMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.sha224Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.sha224Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_freqFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_freqFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_key_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_key_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.dateperiodFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.dateperiodFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_certinfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_certinfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_owner_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_owner_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_boxFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_boxFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.finalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.finalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.intFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.intFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_openDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_openDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_encoding_gzipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_encoding_gzipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_default_strategyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_default_strategyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_rollFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_rollFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.eregiDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.eregiDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_syntaxFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_syntaxFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_server_cseqFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_rtsp_server_cseqFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_0_clientFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_0_clientFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_constFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_constFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.insteadofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.insteadofFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.session_is_registeredDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.session_is_registeredDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltundFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltundFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_label_too_longFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_label_too_longFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_algo_sha512Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_algo_sha512Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.7",
			Reported: "5.4.7",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_pushfunctionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_pushfunctionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_getFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_getFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_gmtFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_gmtFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.boolFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.boolFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_disallowedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_disallowedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_recvFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_recvFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpipe_http1Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpipe_http1Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_connect_toFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_connect_toFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_ns_cFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_ns_cFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedPHP4StyleConstructors.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedPHP4StyleConstructors.Found",
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.2.1",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_get_cache_statsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_get_cache_statsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_multi_strerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_multi_strerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagecropFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagecropFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_http_inputDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_http_inputDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.cgi_discard_pathFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.cgi_discard_pathFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_login_optionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_login_optionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.andFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.andFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_trait_cFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_trait_cFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_closeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_closeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.bus_objerrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.bus_objerrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpipe_nothingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpipe_nothingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_keyfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_keyfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.stream_context_set_defaultFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.stream_context_set_defaultFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_create_instanceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_create_instanceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.gost_cryptoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.gost_cryptoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociwritetemporarylobDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociwritetemporarylobDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ini_scanner_normalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ini_scanner_normalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproxy_http_1_0Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproxy_http_1_0Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.LateStaticBinding.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.LateStaticBinding.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.nullFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.nullFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_ellipsisFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_ellipsisFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.windows_show_crt_warningFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.windows_show_crt_warningFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_force_objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_force_objectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_setFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_setFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.asp_tagsRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.asp_tagsRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.0.26",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_redir_protocolsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_redir_protocolsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.9",
			Reported: "5.2.9",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_timerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_timerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_check_contextjFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_check_contextjFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_catmullromFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_catmullromFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_raw_post_dataDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_raw_post_dataDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.zlib_decodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.zlib_decodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.register_globalsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.register_globalsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._envFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._envFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.gmp_randomDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.gmp_randomDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_nt_workstationFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_nt_workstationFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_inf_or_nanFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_inf_or_nanFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.dbaseRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.dbaseRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqli_allow_persistentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqli_allow_persistentFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_prvregFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_prvregFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__callstaticMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__callstaticMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.magic_quotes_runtimeDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.magic_quotes_runtimeDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepstextRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepstextRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.trap_traceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.trap_traceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.preg_unmatched_as_nullFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.preg_unmatched_as_nullFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.cloneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.cloneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedFunctionParameters.ldap_first_attribute_ber_identifierRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedFunctionParameters.ldap_first_attribute_ber_identifierRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.4",
			High: "7.2.1",
			Reported: "5.2.4",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.datefmt_format_objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.datefmt_format_objectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.parentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.parentFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewClasses.pharfileinfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.pharfileinfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.filter_validate_macFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.filter_validate_macFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.fnv132Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.fnv132Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_intovfFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_intovfFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_webpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_webpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.24",
			Reported: "5.6.24",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_store_replacementFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_store_replacementFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_get_error_messageFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_get_error_messageFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.intl_default_localeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.intl_default_localeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_version_kerberos5Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_version_kerberos5Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_ssl2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_ssl2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.atanhFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.atanhFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_to_date_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_to_date_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gmp_random_seedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gmp_random_seedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_protected_env_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_protected_env_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.RemovedAlternativePHPTags.MaybeASPOpenTagFound": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedAlternativePHPTags.MaybeASPOpenTagFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.0.26",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imap_gcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imap_gcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_minimumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_minimumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gmp_rootFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gmp_rootFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_modify_batchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_modify_batchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.25",
			Reported: "5.4.25",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_ssl_verifystatusFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_ssl_verifystatusFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mt_rand_phpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mt_rand_phpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.interfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.interfaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.hash_equalsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.hash_equalsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_inFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_inFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.class_usesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.class_usesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_day_of_week_typeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_day_of_week_typeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_is_equivalent_toFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_is_equivalent_toFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_block_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_get_block_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.recursivetreeiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.recursivetreeiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splstackFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splstackFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.ncursesRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.ncursesRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClosure.ThisFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClosure.ThisFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_escapeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_escapeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.boolFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.boolFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_ownerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_ownerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intl_is_failureFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intl_is_failureFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_algorithmDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_algorithmDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.18",
			Reported: "5.5.18",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.stream_supports_lockFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.stream_supports_lockFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagesetclipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagesetclipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.mixedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.mixedFound",
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.2.1",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewClasses.typeerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.typeerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_session_activeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_session_activeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_rleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_rleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.fdfRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.fdfRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.quoted_printable_encodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.quoted_printable_encodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_filteredFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_filteredFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_finishFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_finishFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_canonical_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_canonical_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.objectFound",
		Breaks: &CompatibilityRange{
			Low: "7.2.0",
			High: "7.2.1",
			Reported: "7.2",
			MajorMinor: "7.2",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.1.13",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_badstkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_badstkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.cli_set_process_titleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.cli_set_process_titleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_spki_export_challengeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_spki_export_challengeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.deflate_addFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.deflate_addFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_opt_net_cmd_buffer_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_opt_net_cmd_buffer_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_asyncioFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_asyncioFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.array_filter_use_keyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.array_filter_use_keyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_partial_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_partial_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mcveRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mcveRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.endswitchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.endswitchFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bellFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bellFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClosure.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClosure.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.password_hashFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.password_hashFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.bus_adrerrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.bus_adrerrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_userFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_userFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_encoding_deflateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_encoding_deflateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crlfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crlfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ShortArray.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ShortArray.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "5.3.29",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intl_error_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intl_error_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.cli_get_process_titleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.cli_get_process_titleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.random_intFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.random_intFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gc_mem_cachesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gc_mem_cachesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_supported_key_sizesDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_supported_key_sizesDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_prvopcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_prvopcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RequiredOptionalFunctionParameters.stream_socket_enable_crypto_crypto_typeMissing": &Compatibility{
		Source: "PHPCompatibility.PHP.RequiredOptionalFunctionParameters.stream_socket_enable_crypto_crypto_typeMissing",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__file__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__file__Found",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.snmp_oid_output_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.snmp_oid_output_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_sigpendingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_sigpendingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.request_orderFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.request_orderFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.varFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.varFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_available_localesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_available_localesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_first_day_of_weekFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_first_day_of_weekFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.TernaryOperators.MiddleMissing": &Compatibility{
		Source: "PHPCompatibility.PHP.TernaryOperators.MiddleMissing",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_hex_aposFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_hex_aposFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_unescaped_line_terminatorsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_unescaped_line_terminatorsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_initFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_initFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"Generic.NamingConventions.CamelCapsFunctionName.FunctionDoubleUnderscore": &Compatibility{
		Source: "Generic.NamingConventions.CamelCapsFunctionName.FunctionDoubleUnderscore",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewClasses.ui_exception_runtimeexceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.ui_exception_runtimeexceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_ns_cFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_ns_cFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_join_groupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_join_groupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_html_nodefdtdFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_html_nodefdtdFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_crop_sidesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_crop_sidesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_unknownFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_unknownFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocirollbackDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocirollbackDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_ignoreFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_ignoreFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.scandir_sort_descendingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.scandir_sort_descendingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_niceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_niceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.MbstringReplaceEModifier.Deprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.MbstringReplaceEModifier.Deprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.e_user_deprecatedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.e_user_deprecatedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_generalized_cubicFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_generalized_cubicFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_server_public_keyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_server_public_keyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_clearFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_clearFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gd_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gd_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_appconnect_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_appconnect_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.error_clear_lastFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.error_clear_lastFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_script_encodingRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_script_encodingRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.eregi_replaceDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.eregi_replaceDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_extra_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_extra_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_loaded_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_loaded_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cacertfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cacertfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.user_ini_cache_ttlFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.user_ini_cache_ttlFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splqueueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splqueueFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_check_bidiFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_check_bidiFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_allFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_allFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.privateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.privateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gmp_testbitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gmp_testbitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pcntl_signal_dispatchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pcntl_signal_dispatchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mb_ordFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mb_ordFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_exitedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_exitedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_queueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_queueFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.snmp_oid_output_moduleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.snmp_oid_output_moduleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.objectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.OptionalRequiredFunctionParameters.parse_str_resultSoftRequired": &Compatibility{
		Source: "PHPCompatibility.PHP.OptionalRequiredFunctionParameters.parse_str_resultSoftRequired",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_maximumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_maximumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mysql_DeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mysql_DeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.oracleRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.oracleRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.classFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.classFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.endifFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.endifFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundShadowParam": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenClosureUseVariableNames.FoundShadowParam",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.msg_queue_existsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.msg_queue_existsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_cleanupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_cleanupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_sp_minorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_sp_minorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.imagetype_icoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.imagetype_icoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_hex_ampFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_hex_ampFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltinvFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltinvFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_tkillFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_tkillFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_awaiting_responseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_awaiting_responseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewClosure.StaticFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClosure.StaticFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_dict_existsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_dict_existsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals.globalsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals.globalsFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocinumcolsDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocinumcolsDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociserverversionDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociserverversionDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_version_unix_socketsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_version_unix_socketsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.password_verifyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.password_verifyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_tlsext_server_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_tlsext_server_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.1",
			Reported: "5.3.1",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_pedanticFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_pedanticFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_affine_shear_verticalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_affine_shear_verticalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.callableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.callableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewTypeCasts.t_binary_castFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewTypeCasts.t_binary_castFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.getimagesizefromstringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.getimagesizefromstringFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_shut_wrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_shut_wrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.class_aliasFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.class_aliasFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.forward_static_callFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.forward_static_callFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mail_logFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mail_logFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_gaussianFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_gaussianFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_0_serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_0_serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicloselobDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicloselobDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_utf16Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_utf16Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_dns_local_ip4Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_dns_local_ip4Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.password_argon2_default_memory_costFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.password_argon2_default_memory_costFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gc_collect_cyclesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gc_collect_cyclesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ip_multicast_loopFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ip_multicast_loopFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_polling_readingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_polling_readingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_add_to_sessionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_add_to_sessionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlopt_closepolicyRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlopt_closepolicyRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_enabledFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_upload_progress_enabledFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_read_timeoutFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_read_timeoutFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.define_syslog_variablesDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.define_syslog_variablesDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.reflectionzendextensionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.reflectionzendextensionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.sapi_windows_vt100_supportFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.sapi_windows_vt100_supportFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_assertionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_assertionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.php_egg_logo_guidDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.php_egg_logo_guidDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.4.45",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifetchDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifetchDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv23Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv23Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ini_scanner_typedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ini_scanner_typedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.0",
			Reported: "5.6.0",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_cpuFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_cpuFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pcntl_sigwaitinfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pcntl_sigwaitinfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intl_get_error_messageFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intl_get_error_messageFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_error_messageFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_error_messageFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_debugFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_debugFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.ovrimosRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.ovrimosRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewAnonymousClasses.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewAnonymousClasses.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.json_last_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.json_last_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__callMethodStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__callMethodStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewClasses.phardataFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.phardataFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_hupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_hupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_namespaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_namespaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_hanningFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_hanningFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__function__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__function__Found",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.floatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.floatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqli_cache_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqli_cache_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocinlogonDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocinlogonDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gmp_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gmp_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.segv_accerrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.segv_accerrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_disallowedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_disallowedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_leave_source_groupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_leave_source_groupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_httpauth_availFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_httpauth_availFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_polling_activeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_polling_activeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_multi_setoptFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_multi_setoptFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.pharexceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.pharexceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intltimezoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intltimezoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_affine_shear_horizontalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_affine_shear_horizontalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproxy_socks4aFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproxy_socks4aFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.22",
			Reported: "5.5.22",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imageaffinematrixgetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imageaffinematrixgetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_error_messageFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_error_messageFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_script_encodingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_script_encodingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocistatementtypeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocistatementtypeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_accessFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_accessFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_powerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_powerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mnogosearchRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mnogosearchRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.ripemd256Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.ripemd256Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_1_serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_1_serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_msgqueueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_msgqueueFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_exop_who_am_iFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_exop_who_am_iFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.declareFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.declareFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.endwhileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.endwhileFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_random_pseudo_bytesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_random_pseudo_bytesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__setMethodStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__setMethodStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_probesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_keepalive_probesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.useFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.useFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__getMethodStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__getMethodStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._cookieFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._cookieFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicommitDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicommitDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_illadrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_illadrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_raw_dataFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_raw_dataFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ValidIntegers.InvalidBinaryIntegerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ValidIntegers.InvalidBinaryIntegerFound",
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_fetch_allFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_fetch_allFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.hash_hmac_algosFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.hash_hmac_algosFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.class_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.class_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_leading_hyphenFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_leading_hyphenFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_http_version_2_prior_knowledgeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_http_version_2_prior_knowledgeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.http_response_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.http_response_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_register_shutdownFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_register_shutdownFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagepalettetotruecolorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagepalettetotruecolorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_savepointFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_savepointFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_addrinfo_connectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_addrinfo_connectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.intl_idna_variant_2003Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.intl_idna_variant_2003Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.png2wbmpDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.png2wbmpDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_madeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_madeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_pow_equalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_pow_equalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_nofileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_nofileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_max_host_connectionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_max_host_connectionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cacertdirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cacertdirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagecreatefrombmpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagecreatefrombmpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifreedescDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifreedescDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gc_enabledFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gc_enabledFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.user_ini_filenameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.user_ini_filenameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.trap_brkptFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.trap_brkptFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_insteadofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_insteadofFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_recursionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_recursionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_invalid_utf8_ignoreFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_invalid_utf8_ignoreFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.sybaseRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.sybaseRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_userRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_userRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.ticksInvalidValueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.ticksInvalidValueFound",
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_key_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_key_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_shut_rdwrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_shut_rdwrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_auth_okFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_auth_okFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imageflipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imageflipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_weekend_transitionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_weekend_transitionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_sha256_server_public_keyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_sha256_server_public_keyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_unblock_sourceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_unblock_sourceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_cleanFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_cleanFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bilinear_fixedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bilinear_fixedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_describeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_describeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_trace_allocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_trace_allocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.php_check_syntaxRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.php_check_syntaxRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0.5",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.curlfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.curlfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.preg_bad_utf8_offset_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.preg_bad_utf8_offset_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_illopcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_illopcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_errFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_errFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_max_persistentRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_max_persistentRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_illtrpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_illtrpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_tcp_fastopenFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_tcp_fastopenFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.gotoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.gotoFound",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_set_orderingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_set_orderingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_priFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_priFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_noinfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_noinfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_maxconnectsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_maxconnectsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_proxyheaderFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_proxyheaderFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.throwFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.throwFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imageaffinematrixconcatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imageaffinematrixconcatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_bug_compat_42DeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_bug_compat_42DeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_infinityFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_infinityFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlssh_auth_agentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlssh_auth_agentFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_request_pwl_dictFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_request_pwl_dictFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.LateStaticBinding.OutsideClassScope": &Compatibility{
		Source: "PHPCompatibility.PHP.LateStaticBinding.OutsideClassScope",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.ingresRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.ingresRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_traitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_traitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.max_file_uploadsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.max_file_uploadsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.11",
			Reported: "5.2.11",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.11",
			Reported:   "5.2.11",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifreecursorDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifreecursorDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.arithmeticerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.arithmeticerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv3Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.soap_ssl_method_sslv3Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.ifFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.ifFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_addFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_addFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_equivalent_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_equivalent_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_include_dirDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_include_dirDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.enable_post_data_readingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.enable_post_data_readingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollassignDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollassignDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_addrinfo_lookupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_addrinfo_lookupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.max_input_nesting_levelFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.max_input_nesting_levelFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.2",
			Reported: "5.2.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.2",
			Reported:   "5.2.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_timeout_msFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_timeout_msFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.2",
			Reported: "5.2.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cipher_suiteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_cipher_suiteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenEmptyListAssignment.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenEmptyListAssignment.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gc_disableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gc_disableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_gcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_gcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_html_noimpliedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_html_noimpliedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_invalid_property_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_invalid_property_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.orFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.orFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.parse_ini_stringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.parse_ini_stringFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproto_smbsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproto_smbsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_sincFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_sincFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_int_minFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_int_minFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_float_digFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_float_digFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.password_argon2_default_threadsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.password_argon2_default_threadsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.tryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.tryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlgregcal_get_gregorian_changeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlgregcal_get_gregorian_changeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_gpcDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.magic_quotes_gpcDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.fbsql_batchsizeRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.fbsql_batchsizeRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.0",
			Reported:   "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_nullformatRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_nullformatRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_statusFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_statusFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_repeated_wall_time_optionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_repeated_wall_time_optionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_removableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_removableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_os_errnoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_os_errnoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.libxml_schema_createFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.libxml_schema_createFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.1",
			Reported: "5.5.1",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.array_filter_use_bothFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.array_filter_use_bothFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlsslopt_no_revokeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlsslopt_no_revokeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.preg_replace_callback_arrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.preg_replace_callback_arrayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_cfbDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_cfbDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocinewdescriptorDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocinewdescriptorDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.jpeg2wbmpDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.jpeg2wbmpDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ip_multicast_ifFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ip_multicast_ifFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.RemovedHashAlgorithms.salsa10Removed": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedHashAlgorithms.salsa10Removed",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_finalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_finalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_polling_failedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_polling_failedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstVisibility.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstVisibility.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.trait_existsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.trait_existsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ParameterShadowSuperGlobals._requestFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ParameterShadowSuperGlobals._requestFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_callbackRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedConstants.curlclosepolicy_callbackRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.6.0",
			High: "7.2.1",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqli_max_persistentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqli_max_persistentFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_genericDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_genericDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.famRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.famRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.instanceofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.instanceofFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_bind_paramDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_bind_paramDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocifreestatementDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocifreestatementDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gd_minor_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gd_minor_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_domain_name_too_longFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_domain_name_too_longFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_get_error_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_get_error_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_raw_offsetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_raw_offsetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewInterfaces.throwableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewInterfaces.throwableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_algorithm_modeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_is_block_algorithm_modeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_exop_turnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_exop_turnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_cookie_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_cookie_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.continueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.continueFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.finallyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.finallyFound",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.call_user_method_arrayDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.call_user_method_arrayDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_finallyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_finallyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_any_serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_any_serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.preg_jit_stacklimit_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.preg_jit_stacklimit_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_os_familyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_os_familyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_bmpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_bmpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.protectedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.protectedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenSwitchWithMultipleDefaultBlocks.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenSwitchWithMultipleDefaultBlocks.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.cli_pagerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.cli_pagerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_client_encodingDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_client_encodingDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.php_logo_guidDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.php_logo_guidDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.4.45",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.scandir_sort_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.scandir_sort_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_unescaped_slashesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_unescaped_slashesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__setMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__setMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctionParameters.filter_input_array_add_emptyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctionParameters.filter_input_array_add_emptyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_multibyteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_multibyteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollassignelemDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollassignelemDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.finfo_closeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.finfo_closeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_read_buffer_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_net_read_buffer_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_sid_bits_per_characterFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_sid_bits_per_characterFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.0.26",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_debugFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_debugFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_pretty_printFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_pretty_printFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_hammingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_hammingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_asFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_asFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_finallyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_finallyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.finfo_openFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.finfo_openFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.boolvalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.boolvalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewMagicMethods.__callstaticFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewMagicMethods.__callstaticFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.dioRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.dioRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_useFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_useFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imageresolutionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imageresolutionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.cgi_check_shebang_lineFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.cgi_check_shebang_lineFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.0",
			Reported:   "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewClasses.splfixedarrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.splfixedarrayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_nt_domain_controllerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_nt_domain_controllerFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.transliteral_createFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.transliteral_createFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_recvmsgFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_recvmsgFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_spaceshipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_spaceshipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_enable_gcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_enable_gcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlauth_ntlm_wbFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlauth_ntlm_wbFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_exop_refreshFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_exop_refreshFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.publicFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.publicFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_share_initFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_share_initFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocinewcollectionDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocinewcollectionDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gd_release_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gd_release_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_xhtmlFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_xhtmlFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_leave_groupFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_leave_groupFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.whileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.whileFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.lcfirstFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.lcfirstFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_opt_net_read_buffer_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_opt_net_read_buffer_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_writeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_writeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_1Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_sslversion_tlsv1_1Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.18",
			Reported: "5.5.18",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.InvalidDirectiveFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.InvalidDirectiveFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_privateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_privateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_major_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_major_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_nearest_neighbourFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_nearest_neighbourFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.random_bytesFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.random_bytesFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.selfFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.selfFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_coalesce_equalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_coalesce_equalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnsizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnsizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_protocolsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_protocolsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.9",
			Reported: "5.2.9",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNegativeBitshift.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNegativeBitshift.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_pauseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_pauseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_nowFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_nowFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_x509_fingerprintFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_x509_fingerprintFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsfreefontRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsfreefontRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_hex_tagFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_hex_tagFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_bidiFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_bidiFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_primary_ipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_primary_ipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.6",
			Reported: "5.4.6",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_http_version_2tlsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_http_version_2tlsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_start_heredocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_start_heredocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_escapeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_escapeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_crop_blackFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_crop_blackFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpause_allFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpause_allFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproxy_socks5_hostnameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproxy_socks5_hostnameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.22",
			Reported: "5.5.22",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_connection_startedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_connection_startedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_0Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_tls1_0Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.msqlRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.msqlRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.elseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.elseFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewClosure.ThisFoundOutsideClass": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClosure.ThisFoundOutsideClass",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlgregcal_create_instanceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlgregcal_create_instanceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_hex_quotFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_hex_quotFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_cookielistFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_cookielistFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.fileproRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.fileproRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.w32apiRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.w32apiRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.extendsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.extendsFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ValidIntegers.InvalidOctalIntegerFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ValidIntegers.InvalidOctalIntegerFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.register_long_arraysDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.register_long_arraysDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnprecisionDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnprecisionDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sort_naturalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sort_naturalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_dataFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_dataFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_equalsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_equalsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolltrimDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolltrimDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ForbiddenGlobalVariableVariable.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenGlobalVariableVariable.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ipv6_multicast_loopFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ipv6_multicast_loopFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.openssl_pbkdf2Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.openssl_pbkdf2Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_entropy_lengthRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_entropy_lengthRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.1.13",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_2_serverFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_2_serverFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_actual_minumumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_actual_minumumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_from_date_time_zoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_from_date_time_zoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_html401Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_html401Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.datefmt_get_timezoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.datefmt_get_timezoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_exop_start_tlsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_exop_start_tlsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.gethostnameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.gethostnameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_in_daylight_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_in_daylight_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_modeDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_modeDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsslantfontRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsslantfontRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewClasses.globiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.globiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_fixedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_fixedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.gotoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.gotoFound",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.privateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.privateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_charasvarcharRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_charasvarcharRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_passwordRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_default_passwordRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_bind_resultDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_bind_resultDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocidefinebynameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocidefinebynameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_algo_key_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_algo_key_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlexceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlexceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.0",
			Reported: "5.5.0",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_certfileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_certfileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.str_getcsvFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.str_getcsvFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_create_idFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_create_idFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_proxyauth_availFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_proxyauth_availFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.imagetype_webpFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.imagetype_webpFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_env_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_env_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.elseifFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.elseifFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_multi_errnoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_multi_errnoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.stream_isattyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.stream_isattyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_self_testDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_self_testDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.snmp_oid_output_ucdFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.snmp_oid_output_ucdFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_content_length_penalty_sizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_content_length_penalty_sizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_stream_weightFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_stream_weightFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.arrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.arrayFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.array_replace_recursiveFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.array_replace_recursiveFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_useFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_useFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_nontransitional_to_asciiFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_nontransitional_to_asciiFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_2_clientFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_crypto_method_tlsv1_2_clientFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_ssl_enable_npnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_ssl_enable_npnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.get_called_classFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.get_called_classFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltsubFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltsubFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_html5Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_html5Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_pinnedpublickeyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_pinnedpublickeyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.hw_apiRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.hw_apiRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__method__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__method__Found",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.hex2binFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.hex2binFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.divisionbyzeroerrorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.divisionbyzeroerrorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_pipeliningFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_pipeliningFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_redir_post_allFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_redir_post_allFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_ssl_enable_alpnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_ssl_enable_alpnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.posix_setrlimitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.posix_setrlimitFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.dlDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.dlDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "7.2.1",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewClasses.errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_release_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_release_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.stringFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.stringFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imagescaleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imagescaleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.mcryptDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.mcryptDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnnameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnnameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.create_functionDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.create_functionDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ini_scanner_rawFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ini_scanner_rawFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_platformFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_platformFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_besselFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_besselFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_default_stream_ciphersFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_default_stream_ciphersFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_preserve_zero_fractionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_preserve_zero_fractionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.5",
			Reported: "5.6.5",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sqlite3_deterministicFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sqlite3_deterministicFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.3",
			Reported: "7.1.3",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_dict_add_to_personalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_dict_add_to_personalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocilogoffDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocilogoffDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_producttypeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_producttypeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_ctrl_charFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_ctrl_charFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_192_cbcFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_cipher_aes_192_cbcFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_mitchellFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_mitchellFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_response_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_response_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewScalarTypeDeclarations.intFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewScalarTypeDeclarations.intFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.RemovedFunctionParameters.gmmktime_is_dstDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedFunctionParameters.gmmktime_is_dstDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.6.32",
			Reported:   "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedHashAlgorithms.salsa20Removed": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedHashAlgorithms.salsa20Removed",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnscaleDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumnscaleDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociexecuteDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociexecuteDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ipv6_multicast_ifFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ipv6_multicast_ifFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.namespaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.namespaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_powFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_powFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewMagicClassConstant.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewMagicClassConstant.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_hash_functionRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_hash_functionRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.1.13",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.session_hash_bits_per_characterRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.session_hash_bits_per_characterRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.1.0",
			High: "7.2.1",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.1.13",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysql_list_dbsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysql_list_dbsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.6.32",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_algo_block_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_algo_block_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ipproto_ipv6Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ipproto_ipv6Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_flushableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_flushableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlgregcal_is_leap_yearFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlgregcal_is_leap_yearFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__getMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__getMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_list_algorithmsDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_list_algorithmsDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlgregoriancalendarFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlgregoriancalendarFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sig_unblockFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sig_unblockFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sort_flag_caseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sort_flag_caseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_usernameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_usernameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_default_protocolFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_default_protocolFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.zend_signal_checkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.zend_signal_checkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fileinfo_mime_encodingFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fileinfo_mime_encodingFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.callableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.callableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysql_db_queryDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysql_db_queryDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_msgFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_msgFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_http_connectcodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_http_connectcodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.sqliteRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.sqliteRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.finfo_set_flagsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.finfo_set_flagsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqlnd_log_maskFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqlnd_log_maskFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.token_parseFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.token_parseFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_no_flushFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_no_flushFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_minimal_days_in_first_weekFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_minimal_days_in_first_weekFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.inflate_addFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.inflate_addFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_cbcDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_cbcDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_network_timeoutFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_network_timeoutFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewNullableTypes.returnTypeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewNullableTypes.returnTypeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.switchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.switchFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_describeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_describeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_get_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_get_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.eregDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.eregDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_self_testDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_self_testDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_output_handler_cleanableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_output_handler_cleanableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.pgsql_polling_okFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.pgsql_polling_okFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_modes_nameDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_modes_nameDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_xml1Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_xml1Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.intl_idna_variant_uts46Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.intl_idna_variant_uts46Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.password_get_infoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.password_get_infoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.snefru256Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.snefru256Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ereg_replaceDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ereg_replaceDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_contextjFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_contextjFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_partial_output_on_errorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_partial_output_on_errorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.imap_mutf7_to_utf8Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.imap_mutf7_to_utf8Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_dst_savingsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_dst_savingsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_max_linksRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_max_linksRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewMagicMethods.__invokeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewMagicMethods.__invokeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.openssl_algo_sha224Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.openssl_algo_sha224Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.7",
			Reported: "5.4.7",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlmopt_max_total_connectionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlmopt_max_total_connectionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_dns_local_ip6Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_dns_local_ip6Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.password_needs_rehashFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.password_needs_rehashFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_sendmsgFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_sendmsgFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.gd_major_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.gd_major_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.idna_error_empty_labelFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.idna_error_empty_labelFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_insteadofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_insteadofFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.traitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.traitFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_release_savepointFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_release_savepointFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ftp_appendFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ftp_appendFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_gidDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.safe_mode_gidDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_byteasvarcharRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.ifx_byteasvarcharRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.1",
			High: "7.2.1",
			Reported: "5.2.1",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.1",
			Reported:   "5.2.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_tcp_nodelayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_tcp_nodelayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.0",
			Reported: "5.2.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_windows_version_suitemaskFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_windows_version_suitemaskFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curl_redir_post_301Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curl_redir_post_301Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewInterfaces.datetimeinterfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewInterfaces.datetimeinterfaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.mysqli_allow_local_infileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.mysqli_allow_local_infileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.3",
			Reported: "5.2.3",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.3",
			Reported:   "5.2.3",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.fpe_fltresFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.fpe_fltresFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_path_as_isFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_path_as_isFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_end_nowdocFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_end_nowdocFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_is_lenientFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_is_lenientFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_state_mismatchFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_state_mismatchFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_affine_rotateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_affine_rotateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_triangleFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_triangleFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.zlib_huffman_onlyFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.zlib_huffman_onlyFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_noneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crl_noneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_minFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_minFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.exit_on_timeoutFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.exit_on_timeoutFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.2.17",
			Reported:   "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mysqli_send_long_dataDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mysqli_send_long_dataDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicolumntyperawDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicolumntyperawDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.multipleiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.multipleiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sig_blockFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sig_blockFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pcntl_sigprocmaskFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pcntl_sigprocmaskFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociresultDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociresultDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewClasses.intlcodepointbreakiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.intlcodepointbreakiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_minor_versionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_minor_versionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.6",
			Reported: "5.2.6",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.sig_setmaskFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.sig_setmaskFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_flip_bothFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_flip_bothFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_stackFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_stackFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_random_fileFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_random_fileFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_create_time_zone_id_enumerationFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_create_time_zone_id_enumerationFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocinewcursorDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocinewcursorDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_filter_pixelateFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_filter_pixelateFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.callableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.callableFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_get_last_errorsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_get_last_errorsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.enchant_broker_request_dictFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.enchant_broker_request_dictFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_set_time_zoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_set_time_zoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_from_date_timeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_from_date_timeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__callstaticMethodNonStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__callstaticMethodNonStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.import_request_variablesDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.import_request_variablesDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_list_modesDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_list_modesDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_dumpedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_dumpedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.instanceofFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.instanceofFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenBreakContinueVariableArguments.variableArgumentFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenBreakContinueVariableArguments.variableArgumentFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_import_streamFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_import_streamFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.socket_export_streamFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.socket_export_streamFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_local_portFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_local_portFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.6",
			Reported: "5.4.6",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewHashAlgorithms.salsa20Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewHashAlgorithms.salsa20Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ociinternaldebugDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ociinternaldebugDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_progressfunctionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_progressfunctionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_ellipsisFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_ellipsisFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_stmt_error_listFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_stmt_error_listFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.arrayFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.arrayFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewMagicMethods.__debuginfoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewMagicMethods.__debuginfoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.5.37",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"Internal.Exception": &Compatibility{
		Source: "Internal.Exception",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.php_real_logo_guidDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.php_real_logo_guidDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.5.0",
			High: "7.2.1",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.4.45",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.datefmt_set_timezone_idDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.datefmt_set_timezone_idDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.5.0",
			High:       "5.6.32",
			Reported:   "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsencodefontRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsencodefontRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_num_connectsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_num_connectsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewGroupUseDeclarations.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewGroupUseDeclarations.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.curl_share_setoptFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.curl_share_setoptFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewInterfaces.serializableUnsupportedMethod": &Compatibility{
		Source: "PHPCompatibility.PHP.NewInterfaces.serializableUnsupportedMethod",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_error_listFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_error_listFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsloadfontRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsloadfontRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_weighted4Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_weighted4Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_sasl_irFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_sasl_irFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_invalid_utf8_substituteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_invalid_utf8_substituteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.libxml_set_external_entity_loaderFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.libxml_set_external_entity_loaderFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.datefmt_set_timezoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.datefmt_set_timezoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.zend_ze1_compatibility_modeRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.zend_ze1_compatibility_modeRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.3.0",
			High: "7.2.1",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.eregDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.eregDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_iv_sizeDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_enc_get_iv_sizeDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_deinitDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_generic_deinitDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_numeric_checkFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_numeric_checkFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.2",
			Reported: "5.3.2",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crlcheckFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_crlcheckFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenBreakContinueOutsideLoop.FatalError": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenBreakContinueOutsideLoop.FatalError",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.RemovedFunctionParameters.ldap_next_attribute_ber_identifierRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedFunctionParameters.ldap_next_attribute_ber_identifierRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.4",
			High: "7.2.1",
			Reported: "5.2.4",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_exopFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_exopFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_func_overloadDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.mbstring_func_overloadDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.2.0",
			High:       "7.2.1",
			Reported:   "7.2",
			MajorMinor: "7.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_round_half_oddFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_round_half_oddFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_msggqFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_msggqFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_post_varsDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_post_varsDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.forFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.forFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.RemovedFunctionParameters.mktime_is_dstDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedFunctionParameters.mktime_is_dstDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.2.0",
			High:       "5.6.32",
			Reported:   "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_max_send_speed_largeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_max_send_speed_largeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_unix_socket_pathFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_unix_socket_pathFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mb_scrubFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mb_scrubFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewGroupUseDeclarations.TrailingCommaFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewGroupUseDeclarations.TrailingCommaFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocicollmaxDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocicollmaxDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ipv6_multicast_hopsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ipv6_multicast_hopsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_shareFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_shareFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mysqli_pollFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mysqli_pollFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_connecttimeout_msFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_connecttimeout_msFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.2",
			Reported: "5.2.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.poll_outFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.poll_outFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_gotoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_gotoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_locksFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_locksFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_create_defaultFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_create_defaultFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.highlight_bgDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.highlight_bgDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.resourceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.resourceFound",
		Warns: &CompatibilityRange{
			Low:        "7.0.0",
			High:       "7.2.1",
			Reported:   "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewClasses.filesystemiteratorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.filesystemiteratorFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.si_kernelFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.si_kernelFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_binaryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_binaryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.functionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.functionFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.__class__Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.__class__Found",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewClasses.snmpexceptionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewClasses.snmpexceptionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.session_resetFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.session_resetFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewLanguageConstructs.t_coalesceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewLanguageConstructs.t_coalesceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_supported_key_sizesDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_module_get_supported_key_sizesDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_crop_defaultFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_crop_defaultFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.img_flip_verticalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.img_flip_verticalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.filter_validate_domainFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.filter_validate_domainFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.publicFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.publicFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.mb_chrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.mb_chrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__set_stateMethodNonStatic": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__set_stateMethodNonStatic",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlproxy_socks4Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlproxy_socks4Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.9",
			Reported: "5.2.9",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.segv_maperrFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.segv_maperrFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_powFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_powFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.5.37",
			Reported: "5.5",
			MajorMinor: "5.5",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.datefmt_get_calendar_objectFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.datefmt_get_calendar_objectFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_internal_encodingDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.iconv_internal_encodingDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "7.2.1",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.ocisavelobDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.ocisavelobDeprecated",
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "7.2.1",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mysqli_client_ssl_dont_verify_server_certFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mysqli_client_ssl_dont_verify_server_certFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.15",
			Reported: "5.6.15",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlpipe_multiplexFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlpipe_multiplexFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.posix_rlimit_fsizeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.posix_rlimit_fsizeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_create_enumerationFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_create_enumerationFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.VariableVariables.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.VariableVariables.Found",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ent_substituteFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ent_substituteFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.stream_meta_group_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.stream_meta_group_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_unescaped_unicodeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_unescaped_unicodeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_trait_cFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_trait_cFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.json_error_unsupported_typeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.json_error_unsupported_typeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.RemovedExtensions.activescriptRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedExtensions.activescriptRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.1",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewKeywords.t_gotoFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewKeywords.t_gotoFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.expm1Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.expm1Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pg_escape_literalFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pg_escape_literalFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlgregcal_set_gregorian_changeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlgregcal_set_gregorian_changeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.date_create_from_formatFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.date_create_from_formatFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_to_date_time_zoneFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_to_date_time_zoneFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.is_iterableFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.is_iterableFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.session_lazy_writeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.session_lazy_writeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
		Warns: &CompatibilityRange{
			Low:        "5.6.0",
			High:       "5.6.32",
			Reported:   "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ill_illopnFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ill_illopnFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_mandirFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_mandirFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.6",
			Reported: "5.3.6",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.interfaceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.interfaceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.traitFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsInvokedFunctions.traitFound",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.ConstantArraysUsingDefine.Found": &Compatibility{
		Source: "PHPCompatibility.PHP.ConstantArraysUsingDefine.Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.timezone_version_getFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.timezone_version_getFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_set_skipped_wall_time_optionFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_set_skipped_wall_time_optionFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.DeprecatedIniDirectives.detect_unicodeRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedIniDirectives.detect_unicodeRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.4.0",
			High:       "5.4.45",
			Reported:   "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewIniDirectives.phar_cache_listFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewIniDirectives.phar_cache_listFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_ssl3Found": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.ldap_opt_x_tls_protocol_ssl3Found",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlz_get_error_codeFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlz_get_error_codeFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewReturnTypeDeclarations.voidFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewReturnTypeDeclarations.voidFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.splitDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.splitDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.mcast_block_sourceFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.mcast_block_sourceFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_ssl_optionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_ssl_optionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NewConstants.t_spaceshipFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.t_spaceshipFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.6.32",
			Reported: "5.6",
			MajorMinor: "5.6",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.xorFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.xorFound",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "7.2.1",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.set_socket_blockingDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.set_socket_blockingDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.6.32",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlinfo_redirect_urlFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlinfo_redirect_urlFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.6",
			Reported: "5.3.6",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewConstants.php_float_minFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.php_float_minFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.1.13",
			Reported: "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.RemovedGlobalVariables.http_post_filesDeprecatedRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.RemovedGlobalVariables.http_post_filesDeprecatedRemoved",
		Breaks: &CompatibilityRange{
			Low: "5.4.0",
			High: "7.2.1",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
		Warns: &CompatibilityRange{
			Low:        "5.3.0",
			High:       "5.3.29",
			Reported:   "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.protectedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.protectedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.intlcal_get_greatest_minimumFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.intlcal_get_greatest_minimumFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.4.45",
			Reported: "5.4",
			MajorMinor: "5.4",
		},
	},
	"PHPCompatibility.PHP.NonStaticMagicMethods.__callMethodVisibility": &Compatibility{
		Source: "PHPCompatibility.PHP.NonStaticMagicMethods.__callMethodVisibility",
		Breaks: &CompatibilityRange{
			Low: "all",
			High: "all",
			Reported: "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_tftp_no_optionsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_tftp_no_optionsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNames.tryFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNames.tryFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.2.1",
			Reported: "5.0",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.ldap_control_paged_resultFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.ldap_control_paged_resultFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.3.29",
			Reported: "5.3",
			MajorMinor: "5.3",
		},
	},
	"PHPCompatibility.PHP.NewFunctions.pcntl_async_signalsFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewFunctions.pcntl_async_signalsFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.26",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewExecutionDirectives.strict_typesInvalidValueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewExecutionDirectives.strict_typesInvalidValueFound",
		Warns: &CompatibilityRange{
			Low:        "all",
			High:       "all",
			Reported:   "all",
			MajorMinor: "all",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.imagepsextendfontRemoved": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.imagepsextendfontRemoved",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_create_ivDeprecated": &Compatibility{
		Source: "PHPCompatibility.PHP.DeprecatedFunctions.mcrypt_create_ivDeprecated",
		Warns: &CompatibilityRange{
			Low:        "7.1.0",
			High:       "7.2.1",
			Reported:   "7.1",
			MajorMinor: "7.1",
		},
	},
	"PHPCompatibility.PHP.NewConstants.cld_killedFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.cld_killedFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "5.2.17",
			Reported: "5.2",
			MajorMinor: "5.2",
		},
	},
	"PHPCompatibility.PHP.ForbiddenNamesAsDeclared.trueFound": &Compatibility{
		Source: "PHPCompatibility.PHP.ForbiddenNamesAsDeclared.trueFound",
		Breaks: &CompatibilityRange{
			Low: "7.0.0",
			High: "7.2.1",
			Reported: "7.0",
			MajorMinor: "7.0",
		},
	},
	"PHPCompatibility.PHP.NewConstants.curlopt_service_nameFound": &Compatibility{
		Source: "PHPCompatibility.PHP.NewConstants.curlopt_service_nameFound",
		Breaks: &CompatibilityRange{
			Low: "5.2.0",
			High: "7.0.6",
			Reported: "7.0.6",
			MajorMinor: "7.0",
		},
	},
}
