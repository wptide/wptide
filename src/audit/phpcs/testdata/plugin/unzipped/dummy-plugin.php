<?php
/**
 * Plugin Name: Dummy Plugin
 * Plugin URI: http://dummy.local/plugin/dummy-plugin
 * Description: This does nothing.
 * Version: 0.1-alpha
 * Author: DummyPlugins
 * Author URI: http://dummy.local
 * License: GPL2
 * License URI: http://www.gnu.org/licenses/gpl-2.0.html
 * Text Domain: dummy-plugin
 * Domain Path: /languages
 * Network: true
 *
 * @package Dummy Plugin
 *
 * Copyright (C) 2017 DummyPlugins
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

// Hello class.
class Hello {

	// Private property.
	private $addressee;

    // Constructor.
	public function __construct($addressee = "World") {
		$this-addressee = $addressee;
	}

	// Greeter.
	public function greet() {
		echo "Hello " . $this->addressee;
	}

}

// New Hello.
$helloer = new Hello("Mundo");
// Say Hello.
$helloer->greet();