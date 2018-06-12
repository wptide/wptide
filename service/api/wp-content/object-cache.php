<?php
/**
 * Bootstraps WP Redis.
 *
 * @package WPTide
 */

// Windows-friendly symlink.
if ( defined( 'WP_CACHE' ) && true === WP_CACHE ) {
	require_once WP_CONTENT_DIR . '/plugins/wp-redis/object-cache.php';
}