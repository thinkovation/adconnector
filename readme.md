# Readme
This is the result of some playing around with the Microsofr Graph API and might form the basis of a means to synchronise MS AD users with a third-party APP.

# Connector Setup
To set up the connector you need to create an App Registration in the Microsoft Entra admin center.

You should set it up as follows:

### Branding and Properties
Leave this blank - the only thing that should be set is the Name

### Authentication
Platform Configurations - leave blank

Supported Account Types -  the App should be set up as a single tenant application - "Accounts in this organisational directory only"

Advanced settings - leave as default

### Certificates and secrets
Create a new client secret - the client secret that is created here will be used with the tenant ID and client ID to authenticate the app with the directory

### Token Configuration
Leave blank

### API Permissions

The app should be granted the Directory.Read.All API permission on the Microsoft Graph API.

### Expose an API
Leave blank

### App Roles
Leave blank

### Owners
Make sure that the owner is assigned to someone with admin rights

### Roles and administrators 
Leave as is

### Manifest
Leave as is
