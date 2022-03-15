terraform {
required_providers {
dsd = {
source = "seb/dsd"
}
}
}
provider "dsd"{
    dsd_api = "https://dsd-external-api.dcp.acc.sebank.se/"
}