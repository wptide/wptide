<?php
/**
 * Plugin Name: Google Cloud Storage Upload
 * Description: Upload media files to Google Cloud Storage. Based on https://wordpress.org/plugins/gcs/
 * Version: 1.0.0
 */

namespace Google\Cloud\Storage\WordPress;

/**
 * Google Cloud Storage Upload
 */
class Upload {

	/**
	 * Use HTTPS Setting.
	 *
	 * @type string
	 */
	const USE_HTTPS = 'gcs_upload_use_https';

	/**
	 * Bucket Name Setting.
	 *
	 * @type string
	 */
	const BUCKET_NAME = 'gcs_upload_bucket_name';

	/**
	 * Plugin constructor.
	 */
	public function __construct() {

		// The Storage Client is not being referenced in the wp-config.php
		if ( ! class_exists( '\Google\Cloud\Storage\StorageClient' ) ) {
			return;
		}

		// Use the local file system on the host machine.
		if ( 'gcs' !== getenv( 'API_UPLOAD_HANDLER' ) && false === getenv( 'GAE_VERSION' ) ) {
			return;
		}

		// Add plugin hooks.
		add_action( 'admin_menu', array( $this, 'options_page' ) );
		add_action( 'admin_init', array( $this, 'register_settings' ) );
		add_filter( 'plugin_action_links', array( $this, 'settings_link' ), 10, 2 );
		add_filter( 'upload_dir', array( $this, 'filter_upload_dir' ) );
		add_filter( 'wp_delete_file', array( $this, 'filter_delete_file' ) );
	}

	/**
	 * Callback for defining options.
	 */
	public function options_page() {
		add_options_page(
			__( 'Google Cloud Storage Upload', 'gcs-upload' ),
			__( 'GCS Upload', 'gcs-upload' ),
			'manage_options',
			'gcs-upload',
			array( $this, 'options_page_view' )
		);
	}

	/**
	 * Render the options page.
	 */
	public function options_page_view() {

		// Check user capabilities.
		if ( ! current_user_can( 'manage_options' ) ) {
			return;
		}
		?>
		<div class="wrap">
			<h1><?= esc_html( get_admin_page_title() ); ?></h1>
			<form action="options.php" method="POST">
				<?php

				// Output security fields for the registered setting "gcs_settings".
				settings_fields( 'gcs_settings' );

				/*
				 * Output setting sections and their fields.
				 *
				 * Sections are registered to the "gcs-upload" page, each field is
				 * registered to a specific section.
				 */
				do_settings_sections( 'gcs-upload' );

				// Output save settings button.
				submit_button( __( 'Save Settings', 'gcs-upload' ) );
				?>
			</form>
		</div>
		<?php
	}

	/**
	 * Callback for registering the setting.
	 */
	public function register_settings() {
		add_option(self::USE_HTTPS, true);

		register_setting(
			'gcs_settings',
			self::BUCKET_NAME,
			array( $this, 'validate_bucket' )
		);

		register_setting(
			'gcs_settings',
			self::USE_HTTPS,
			array( $this, 'validate_use_https' )
		);

		add_settings_section(
			'gcs-upload-config',
			__( 'Configurations', 'gcs-upload' ),
			null,
			'gcs-upload'
		);

		add_settings_field(
			self::BUCKET_NAME,
			__( 'Bucket Name', 'gcs-upload' ),
			array( $this, 'bucket_form' ),
			'gcs-upload',
			'gcs-upload-config'
		);

		add_settings_field(
			self::USE_HTTPS,
			__( 'Use HTTPS', 'gcs-upload' ),
			array( $this, 'use_https_form' ),
			'gcs-upload',
			'gcs-upload-config'
		);
	}

	/**
	 * Add a settings link.
	 *
	 * @action plugin_action_links
	 *
	 * @param array  $links Array of links in HTML format.
	 * @param string $file  The mu-plugin file.
	 */
	function settings_link( $links, $file ) {
		if ( $file === basename( __FILE__ ) ) {
			$links[] = sprintf( '<a href="%s">%s</a>', admin_url( 'options-general.php?page=gcs-upload' ), __( 'Settings', 'gcs-upload' ) );
		}

		return $links;
	}

	/**
	 * Swap the upload dir with gs:// path in the GCS bucket.
	 *
	 * @param array $uploads Array of upload directory data with keys of 'path',
	 * 'url', 'subdir, 'basedir', and 'error'.
	 */
	public function filter_upload_dir( $uploads ) {
		$bucket = get_option( self::BUCKET_NAME, '' );

		// Do nothing if this is a local audit report request.
		if ( ! empty( $_REQUEST['post_id'] ) && ! empty( $_REQUEST['type'] ) && ! empty( $_REQUEST['standard'] ) ) {
			return $uploads;
		}

		// Do nothing without the bucket name.
		if ( '' === $bucket ) {
			return $uploads;
		}

		$basedir = sprintf( 'gs://%s/%s', $bucket, get_current_blog_id() );
		$use_https = get_option( self::USE_HTTPS, false );
		$baseurl = sprintf(
			'%s://storage.googleapis.com/%s/%s',
			$use_https ? 'https' : 'http',
			$bucket,
			get_current_blog_id()
		);

		$uploads = array(
			'path' => $basedir . $uploads['subdir'],
			'subdir' => $uploads['subdir'],
			'error' => false,
		);

		$uploads['url'] = rtrim( $baseurl . $uploads['subdir'], '/' );
		$uploads['basedir'] = $basedir;
		$uploads['baseurl'] = $baseurl;

		return $uploads;
	}

	/**
	 * Unlink files starts with 'gs://'
	 *
	 * This is needed because WordPress thinks a path starts with 'gs://' is
	 * not an absolute path and manipulate it in a wrong way before unlinking
	 * intermediate files.
	 *
	 * TODO: Use `path_is_absolute` filter when a bug below is resolved:
	 * https://core.trac.wordpress.org/ticket/38907#ticket
	 *
	 * @param string $file The file path.
	 */
	public function filter_delete_file( $file ) {
		$prefix = 'gs://';
		if ( substr( $file, 0, strlen( $prefix ) ) === $prefix ) {
			@ unlink( $file );
		}

		return $file;
	}

	/**
	 * Display the input form for the bucket.
	 */
	public function bucket_form() {
		$bucket = get_option( self::BUCKET_NAME, '' );

		echo sprintf(
			'<input id="%s" name="%s" type="text" value="%s">',
			self::BUCKET_NAME,
			self::BUCKET_NAME,
			esc_attr($bucket)
		);
	}

	/**
	 * Display the input form for use_https_for_media.
	 */
	public function use_https_form() {
		$enabled = get_option( self::USE_HTTPS, false );
		echo sprintf(
			'<input id="%s", name="%s" type="checkbox" %s>',
			self::USE_HTTPS,
			self::USE_HTTPS,
			checked( $enabled, true, false )
		);

		$desc = __( 'Only affects new uploads, it will not change the HTTP scheme for files previously uploaded', 'gcs-upload' );
		echo '<p class="description">' . esc_html( $desc ) . '</p>';
	}

	/**
	 * Validate the bucket name in the form.
	 *
	 * @param string $input The input value.
	 */
	public function validate_bucket( $input ) {
		$path = sprintf( 'gs://%s/', $input );

		if ( ! is_writable( $path ) ) {
			add_settings_error(
				'gcs_settings',
				'invalid-bucket',
				__( 'The bucket does not exist, or is not writable', 'gcs-upload' )
			);

			return get_option( self::BUCKET_NAME, '' );
		}

		return $input;
	}

	/**
	 * Validate the value for the use_https form.
	 *
	 * @param bool $input The input value.
	 */
	public function validate_use_https( $input ) {
		return (bool) $input;
	}
}

$gcs_upload = new Upload();
