<?php
/**
 * Health Check
 *
 * @package WordPress
 */

// If the WordPress config file is missing, fail the health check.
if ( ! file_exists( dirname( __FILE__ ) . '/wp-config.php' ) ) {
	header( 'HTTP/1.1 499 NOT OK' );
	return false;
}

/**
 * Ensure the Database is connected or respond with header '499'.
 */
register_shutdown_function( function () {
	$output = ob_get_contents();
	if ( false !== stripos( $output, 'Error establishing a database connection' ) ) {
		header( 'HTTP/1.1 499 NOT OK' );
	}
	ob_end_clean();
} );

ob_start();
require_once( dirname( __FILE__ ) . '/wp-load.php' );