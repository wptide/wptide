runtime: php
env: flex

beta_settings:
  cloud_sql_instances: ${GCP_PROJECT}:${GCP_REGION}:${GCSQL_API_INSTANCE}

runtime_config:
  document_root: wordpress

readiness_check:
  path: "/health-check.php"
  timeout_sec: ${GAE_API_RC_TIMEOUT_SEC}
  check_interval_sec: ${GAE_API_RC_CHECK_INTERVAL_SEC}
  failure_threshold: ${GAE_API_RC_FAILURE_THRESHOLD}
  success_threshold: ${GAE_API_RC_SUCCESS_THRESHOLD}
  app_start_timeout_sec: ${GAE_API_RC_APP_START_TIMEOUT_SEC}

automatic_scaling:
  min_num_instances: ${GAE_API_AS_MIN_NUM_INSTANCES}
  max_num_instances: ${GAE_API_AS_MAX_NUM_INSTANCES}
  cool_down_period_sec: ${GAE_API_AS_COOL_DOWN_PERIOD_SEC}
  cpu_utilization:
    target_utilization: ${GAE_API_AS_CPU_TARGET_UTILIZATION}

env_variables:
  WHITELIST_FUNCTIONS: escapeshellarg,escapeshellcmd,exec,pclose,popen,shell_exec,phpversion,php_uname
  API_CACHE_DEBUG: ${API_CACHE_DEBUG}
  API_CACHE_TTL: ${API_CACHE_TTL}
  API_REDIS_AUTH: ${API_REDIS_AUTH}
  API_REDIS_DATABASE: ${API_REDIS_DATABASE}
  API_REDIS_HOST: ${API_REDIS_HOST}
  API_REDIS_PORT: ${API_REDIS_PORT}
  AWS_API_KEY: ${AWS_API_KEY}
  AWS_API_SECRET: ${AWS_API_SECRET}
  AWS_S3_BUCKET_NAME: ${AWS_S3_BUCKET_NAME}
  AWS_S3_REGION: ${AWS_S3_REGION}
  AWS_S3_VERSION: ${AWS_S3_VERSION}
  AWS_SQS_QUEUE_LH: ${AWS_SQS_QUEUE_LH}
  AWS_SQS_QUEUE_PHPCS: ${AWS_SQS_QUEUE_PHPCS}
  AWS_SQS_REGION: ${AWS_SQS_REGION}
  AWS_SQS_VERSION: ${AWS_SQS_VERSION}

skip_files:
- tpl/
- Dockerfile
- Makefile
- setup.sh
- wptests.sql