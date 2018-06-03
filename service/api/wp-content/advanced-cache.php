<?php
/**
 * Bootstraps Redis Page Cache.
 *
 * @package WPTide
 */

// Windows-friendly symlink.
if ( defined( 'WP_CACHE' ) && true === WP_CACHE ) {
	require_once WP_CONTENT_DIR . '/plugins/pj-page-cache-red/advanced-cache.php';
}