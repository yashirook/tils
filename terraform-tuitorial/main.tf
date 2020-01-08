provider "google" {
  credentials = file("yashiroken-dev-10204a57e1c4.json")

  project = "yashiroken-dev"
  region = "us-west1"
  zone = "us-west1-b"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}
