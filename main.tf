provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "resource_group" {
  name     = "BlogSource"
  location = "eastasia"
}

resource "azurerm_mssql_server" "sql_server" {
  name                         = "blog-db-server"
  resource_group_name          = azurerm_resource_group.resource_group.name
  location                     = azurerm_resource_group.resource_group.location
  version                      = "12.0"
  administrator_login          = var.id
  administrator_login_password = var.password
}
resource "azurerm_mssql_database" "db" {
  name         = "BlogDB"
  server_id    = azurerm_mssql_server.sql_server.id
  collation    = "SQL_Latin1_General_CP1_CI_AS"
  license_type = "LicenseIncluded"
  max_size_gb  = 2
  sku_name     = "Basic"
}
resource "azurerm_container_registry" "container_registry" {
  name                = "BlogBackendContainer"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  sku                 = "Basic"
  admin_enabled       = true
}
resource "azurerm_service_plan" "service_plan" {
  name                = "Blog-service-plan"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  os_type             = "Linux"
  sku_name            = "B1"
}

resource "azurerm_linux_web_app" "blog_app" {
  name                = "app-linux-Blog-web-backend"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  service_plan_id     = azurerm_service_plan.service_plan.id

  site_config {
    always_on = false
    application_stack {
      docker_image     = var.dockerImageName
      docker_image_tag = "latest"
    }
  }
  app_settings = {
    DOCKER_REGISTRY_SERVER_URL      = var.BlogBackendContainerServerUrl,
    DOCKER_REGISTRY_SERVER_USERNAME = var.BlogBackendContainerServerUsername,
    DOCKER_REGISTRY_SERVER_PASSWORD = var.BlogBackendContainerServerPassword
  }
  connection_string ={
    SQL_CONNECTION = var.BlogBackendConnectString
  }
}


 resource "azurerm_mssql_firewall_rule" "sql_firewall" {
  for_each         = toset(azurerm_linux_web_app.blog_app.outbound_ip_address_list)
  name             = "blogDBfirewall${each.key}"
  server_id        = azurerm_mssql_server.sql_server.id
  start_ip_address = each.key
  end_ip_address   = each.key
}
 