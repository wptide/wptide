runtime: php
env: flex

beta_settings:
  cloud_sql_instances: ${GCP_PROJECT}:${GCP_REGION}:${API_GAE_SQL_INSTANCE_NAME}

runtime_config:
  document_root: wordpress

readiness_check:
  path: "/health-check.php"
  timeout_sec: ${API_GAE_RC_TIMEOUT_SEC}
  check_interval_sec: ${API_GAE_RC_CHECK_INTERVAL_SEC}
  failure_threshold: ${API_GAE_RC_FAILURE_THRESHOLD}
  success_threshold: ${API_GAE_RC_SUCCESS_THRESHOLD}
  app_start_timeout_sec: ${API_GAE_RC_APP_START_TIMEOUT_SEC}

automatic_scaling:
  min_num_instances: ${API_GAE_AS_MIN_NUM_INSTANCES}
  max_num_instances: ${API_GAE_AS_MAX_NUM_INSTANCES}
  cool_down_period_sec: ${API_GAE_AS_COOL_DOWN_PERIOD_SEC}
  cpu_utilization:
    target_utilization: ${API_GAE_AS_CPU_TARGET_UTILIZATION}

env_variables:
  WHITELIST_FUNCTIONS: escapeshellarg,escapeshellcmd,exec,pclose,popen,shell_exec,phpversion,php_uname
  AWS_S3_BUCKET_NAME: ${AWS_S3_BUCKET_NAME}
  AWS_S3_KEY: ${AWS_S3_KEY}
  AWS_S3_REGION: ${AWS_S3_REGION}
  AWS_S3_SECRET: ${AWS_S3_SECRET}
  AWS_S3_VERSION: ${AWS_S3_VERSION}
  AWS_SQS_KEY: ${AWS_SQS_KEY}
  AWS_SQS_QUEUE_NAME: ${AWS_SQS_QUEUE_NAME}
  AWS_SQS_REGION: ${AWS_SQS_REGION}
  AWS_SQS_SECRET: ${AWS_SQS_SECRET}
  AWS_SQS_VERSION: ${AWS_SQS_VERSION}

skip_files:
- tpl/
- Makefile
- service-account.json
- setup.sh
- wptests.sql