runtime: php
env: flex

beta_settings:
  cloud_sql_instances: ${GCP_PROJECT}:${GCP_REGION}:${API_CLOUD_SQL_DB}

runtime_config:
  document_root: wordpress

readiness_check:
  path: "/health-check.php"
  timeout_sec: 4
  check_interval_sec: 5
  failure_threshold: 2
  success_threshold: 2
  app_start_timeout_sec: 300

automatic_scaling:
  min_num_instances: 2
  max_num_instances: 20
  cool_down_period_sec: 120
  cpu_utilization:
    target_utilization: 0.5

env_variables:
  WHITELIST_FUNCTIONS: escapeshellarg,escapeshellcmd,exec,pclose,popen,shell_exec,phpversion,php_uname

skip_files:
- tpl/
- health-check.php
- Makefile
- service-account.json
- wp-config.php
- wptests.sql