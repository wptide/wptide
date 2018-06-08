
Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.define "go-tide"
  config.vm.hostname = "tide"
  config.vm.network "private_network", type: "dhcp"
  config.vm.synced_folder ".", "/vagrant", create: true, mount_options: ["dmode=777,fmode=777"]

  # Remove ubuntu-xenial-16.04-cloudimg-console.log
  config.vm.provider "virtualbox" do |vb|
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]
  end

  # Provision the box.
  config.vm.provision :shell, privileged: false, path: "bin/vagrant-provision.sh"
  config.vm.provision :shell, privileged: false, path: "bin/vagrant-up.sh", run: "always"
end
