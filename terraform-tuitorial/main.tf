provider "google" {
  credentials = file("yashiroken-dev-10204a57e1c4.json")

  project = "yashiroken-dev"
  region = "us-west1"
  zone = "us-west1-b"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}

resource "google_compute_instance" "vm_instance" {
  name = "terraform-instance"
  machine_type = "f1-micro"
  tags = ["web", "dev"]

  boot_disk {
    initialize_params {
      image = "cos-cloud/cos-stable"
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.self_link
    access_config {
      nat_ip = google_compute_address.vm_static_ip.address
    }
  }
}

resource "google_compute_address" "vm_static_ip" {
  name = "terraform-static-ip"
}

/* resource "google_storage_bucket" "example_bucket" {
  name = "terraform-example-bucket-yashiroken-20200109"
  location = "US"

  website {
    main_page_suffix = "index.html"
    not_found_page = "404.html"
  }
}

resource "google_compute_instance" "another_instance" {
  depends_on = [google_storage_bucket.example_bucket]

  name = "terraform-instance-2"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "cos-cloud/cos-stable"
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.self_link
    access_config {
    }
  }
} */
