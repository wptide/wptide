<?php
/**
 * The base configuration for WordPress
 *
 * The wp-config.php creation script uses this file during the
 * installation. You don't have to use the web site, you can
 * copy this file to "wp-config.php" and fill in the values.
 *
 * This file contains the following configurations:
 *
 * * MySQL settings
 * * Secret keys
 * * Database table prefix
 * * ABSPATH
 *
 * @link https://codex.wordpress.org/Editing_wp-config.php
 *
 * @package WordPress
 */

// $is_gae is true on Google App Engine.
$is_gae = ( getenv( 'GAE_VERSION' ) !== false );

// Register GCS stream wrapper
if ( $is_gae || 'gcs' === getenv( 'API_UPLOAD_HANDLER' ) ) {
    require_once( __DIR__ . '/../vendor/autoload.php' );
    $storageClient = new Google\Cloud\Storage\StorageClient();
    $storageClient->registerStreamWrapper();
}

// Disable pseudo cron behavior
define( 'DISABLE_WP_CRON', true );

// Use HTTPS on production.
$protocol = $is_gae ? 'https://' : 'http://';

// Define the host.
if ( isset( $_SERVER['HTTP_HOST'] ) ) {
    define( 'HTTP_HOST', $_SERVER['HTTP_HOST'] );
} else {
    define( 'HTTP_HOST', 'localhost' );
}

// Set the URLs.
define( 'WP_HOME', $protocol . HTTP_HOST );
define( 'WP_SITEURL', $protocol . HTTP_HOST );

// Force SSL for admin pages
define( 'FORCE_SSL_ADMIN', $is_gae );

// Security settings.
define( 'DISALLOW_FILE_MODS', true );
define( 'AUTOMATIC_UPDATER_DISABLED', true );

// Set the content directory.
define( 'WP_CONTENT_DIR', dirname( dirname( __FILE__ ) ) . '/wp-content' );

// ** MySQL settings ** //
if ( $is_gae ) {
    define( 'DB_HOST',     ':/cloudsql/${GCP_PROJECT}:${GCP_REGION}:${GCSQL_API_INSTANCE}' );
    define( 'DB_NAME',     '${GCSQL_API_DB_NAME}' );
    define( 'DB_USER',     '${GCSQL_API_DB_USER}' );
    define( 'DB_PASSWORD', '${GCSQL_API_DB_PASSWORD}' );
} else {
    define( 'DB_HOST',     '${API_DB_HOST}' );
    define( 'DB_NAME',     '${API_DB_NAME}' );
    define( 'DB_USER',     '${API_DB_USER}' );
    define( 'DB_PASSWORD', '${API_DB_PASSWORD}' );
}

/** Database Charset to use in creating database tables. */
define( 'DB_CHARSET', 'utf8' );

/** The Database Collate type. Don't change this if in doubt. */
define( 'DB_COLLATE', '' );

/**#@+
 * Authentication Unique Keys and Salts.
 *
 * Change these to different unique phrases!
 * You can generate these using the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}
 * You can change these at any point in time to invalidate all existing cookies. This will force all users to have to log in again.
 *
 * @since 2.6.0
 */

define( 'AUTH_KEY',         '${AUTH_KEY}' );
define( 'SECURE_AUTH_KEY',  '${SECURE_AUTH_KEY}' );
define( 'LOGGED_IN_KEY',    '${LOGGED_IN_KEY}' );
define( 'NONCE_KEY',        '${NONCE_KEY}' );
define( 'AUTH_SALT',        '${AUTH_SALT}' );
define( 'SECURE_AUTH_SALT', '${SECURE_AUTH_SALT}' );
define( 'LOGGED_IN_SALT',   '${LOGGED_IN_SALT}' );
define( 'NONCE_SALT',       '${NONCE_SALT}' );

/**#@-*/

/**
 * WordPress Database Table prefix.
 *
 * You can have multiple installations in one database if you give each
 * a unique prefix. Only numbers, letters, and underscores please!
 */
$table_prefix  = 'wp_';

/**
 * WordPress Cache.
 *
 * Page and Object caching are both backed by Redis.
 */
define( 'WP_CACHE',          ${API_CACHE} );
define( 'WP_CACHE_KEY_SALT', '${API_CACHE_KEY_SALT}' );

// Redis settings.
if ( true === WP_CACHE ) {
    global $redis_server, $redis_page_cache_config;

    $redis_server = array(
        'host'     => getenv( 'API_REDIS_HOST' ),
        'port'     => (int) getenv( 'API_REDIS_PORT' ),
        'auth'     => getenv( 'API_REDIS_AUTH' ),
        'database' => (int) getenv( 'API_REDIS_DATABASE' ),
    );

    $redis_page_cache_config = array(
        'redis_host'          => getenv( 'API_REDIS_HOST' ),
        'redis_port'          => (int) getenv( 'API_REDIS_PORT' ),
        'redis_auth'          => getenv( 'API_REDIS_AUTH' ),
        'redis_db'            => (int) getenv( 'API_REDIS_DATABASE' ),
        'ttl'                 => getenv( 'API_CACHE_TTL' ),
        'ignore_cookies'      => array( 'wordpress_test_cookie', '__utmt', '__utma', '__utmb', '__utmc', '__utmz', '__gads', '__qca', '_ga' ),
        'ignore_request_keys' => array( 'utm_source', 'utm_medium', 'utm_term', 'utm_content', 'utm_campaign' ),
        'debug'               => getenv( 'API_CACHE_DEBUG' ),
        'gzip'                => true,
    );

    // Cloud Memorystore does not need to authorize.
    if ( $is_gae ) {
        unset( $redis_server['auth'] );
        unset( $redis_page_cache_config['redis_auth'] );
    }
}

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 *
 * For information on other constants that can be used for debugging,
 * visit the Codex.
 *
 * @link https://codex.wordpress.org/Debugging_in_WordPress
 */
define( 'WP_DEBUG', ! $is_gae );
define( 'WP_DEBUG_LOG', ! $is_gae );

// API configuration settings.
define( 'API_MESSAGE_PROVIDER', getenv( 'API_MESSAGE_PROVIDER' ) );

// AWS API settings.
define( 'AWS_API_KEY',         getenv( 'AWS_API_KEY' ) );
define( 'AWS_API_SECRET',      getenv( 'AWS_API_SECRET' ) );

// AWS S3 settings.
define( 'AWS_S3_BUCKET_NAME',  getenv( 'AWS_S3_BUCKET_NAME' ) );
define( 'AWS_S3_REGION',       getenv( 'AWS_S3_REGION' ) );
define( 'AWS_S3_VERSION',      getenv( 'AWS_S3_VERSION' ) );

// AWS SQS settings.
define( 'AWS_SQS_QUEUE_LH',    getenv( 'AWS_SQS_QUEUE_LH' ) );
define( 'AWS_SQS_QUEUE_PHPCS', getenv( 'AWS_SQS_QUEUE_PHPCS' ) );
define( 'AWS_SQS_REGION',      getenv( 'AWS_SQS_REGION' ) );
define( 'AWS_SQS_VERSION',     getenv( 'AWS_SQS_VERSION' ) );

// Cloud Firestore settings.
define( 'GCF_QUEUE_LH',    getenv( 'GCF_QUEUE_LH' ) );
define( 'GCF_QUEUE_PHPCS', getenv( 'GCF_QUEUE_PHPCS' ) );
define( 'GCP_PROJECT',     getenv( 'GCP_PROJECT' ) );

// MongoDB settings.
if ( 'local' === API_MESSAGE_PROVIDER ) {
	define( 'MONGO_DATABASE_NAME',     getenv( 'MONGO_DATABASE_NAME' ) );
	define( 'MONGO_DATABASE_PASSWORD', getenv( 'MONGO_DATABASE_PASSWORD' ) );
	define( 'MONGO_DATABASE_USERNAME', getenv( 'MONGO_DATABASE_USERNAME' ) );
	define( 'MONGO_HOST',              'mongo:27017' );
	define( 'MONGO_QUEUE_LH',          getenv( 'MONGO_QUEUE_LH' ) );
	define( 'MONGO_QUEUE_PHPCS',       getenv( 'MONGO_QUEUE_PHPCS' ) );
}

/* That's all, stop editing! Happy blogging. */

/** Absolute path to the WordPress directory. */
if ( ! defined( 'ABSPATH' ) ) {
    define( 'ABSPATH', dirname( __FILE__ ) . '/' );
}

/** Sets up WordPress vars and included files. */
require_once( ABSPATH . 'wp-settings.php' );
