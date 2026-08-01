# 🏪 Merchant Service

> **Version:** 1.0.0  
> **Project:** KongPay  
> **Document Type:** Merchant Service  
> **Last Updated:** August 2026

---

# Executive Summary

The Merchant Service is responsible for managing merchant accounts, business profiles, payment capabilities, API credentials, operational settings, and merchant lifecycle management within the KongPay platform.

It provides the necessary foundation for organizations and businesses to securely integrate KongPay into their own applications while maintaining clear ownership, configurable permissions, and standardized payment interfaces.

The Merchant Service operates independently from payment processing and wallet management, enabling modular development and simplified platform maintenance.

---

# Objectives

The Merchant Service is designed to:

- Register merchant organizations
- Manage merchant profiles
- Store merchant configuration
- Issue API credentials
- Manage operational status
- Support payment integration
- Enable secure merchant authentication
- Maintain merchant-related audit records

---

# Design Principles

The Merchant Service follows several architectural principles:

- API-First Design
- Security by Design
- Modular Components
- Stateless Operations
- Clear Ownership
- Configuration Isolation
- Auditability

---

# Merchant Lifecycle

Every merchant follows a standardized lifecycle.

```text
Application
      │
      ▼
Registration
      │
      ▼
Verification
      │
      ▼
Activation
      │
      ▼
Operational
      │
 ┌────┴───────────┐
 ▼                ▼
Suspended      Closed
```

---

# Merchant Status

| Status | Description |
|---------|-------------|
| PENDING | Registration submitted |
| VERIFIED | Identity verification completed |
| ACTIVE | Merchant can process payments |
| SUSPENDED | Temporarily disabled |
| CLOSED | Permanently closed |

---

# Merchant Profile

Each merchant maintains a profile containing business-related information.

Typical profile attributes include:

- Merchant Identifier
- Business Name
- Legal Name
- Contact Information
- Business Category
- Operational Status
- Time Zone
- Country
- Supported Currency
- Configuration Metadata

Exact fields may vary depending on deployment requirements.

---

# Merchant Components

## Merchant Registration

Responsible for onboarding new merchants.

Typical responsibilities include:

- Registration
- Identity information
- Business information
- Contact details

---

## Merchant Configuration

Responsible for configurable merchant settings.

Examples include:

- Notification preferences
- Webhook configuration
- Payment options
- API access
- Operational preferences

---

## API Credential Management

The Merchant Service manages credentials used by merchant applications.

Typical responsibilities include:

- API Key issuance
- Credential rotation
- Credential revocation
- Usage tracking

Credential storage should follow secure secret management practices.

---

## Merchant Dashboard

The Merchant Dashboard provides operational visibility.

Typical capabilities include:

- Profile management
- Payment overview
- Transaction history
- API credential management
- Operational settings
- Reporting

The dashboard implementation is outside the scope of this document.

---

# Merchant Capabilities

A merchant may perform operations such as:

- Create payment requests
- View transaction history
- Retrieve reports
- Manage webhooks
- Configure notifications
- Manage API credentials

Available capabilities depend on assigned permissions.

---

# Merchant Authentication

Merchant requests must be authenticated before accessing protected resources.

Authentication mechanisms may include:

- API Keys
- JWT
- OAuth 2.0
- Session Tokens

Authentication details are described in **authentication.md**.

---

# Merchant Authorization

Authorization determines what actions a merchant is permitted to perform.

Examples include:

- Read-only access
- Payment creation
- Refund requests
- Configuration updates
- Administrative actions

Authorization details are documented in **authorization.md**.

---

# Merchant Payment Flow

A simplified payment request follows this sequence:

```text
Merchant

    │

API Gateway

    │

Authentication

    │

Authorization

    │

Payment Service

    │

Wallet Service

    │

Settlement

    │

Notification

    │

Audit
```

---

# Webhooks

Merchants may receive asynchronous notifications for important events.

Examples include:

- Payment completed
- Payment failed
- Refund processed
- Settlement completed
- Account updates

Webhook delivery should support retry mechanisms and signature verification.

---

# Reporting

The Merchant Service supports operational reporting.

Typical reports include:

- Payment summary
- Transaction history
- Daily activity
- Monthly activity
- Settlement summary

Reporting requirements may differ by deployment.

---

# Security Considerations

Merchant operations should implement security controls such as:

- HTTPS
- Authentication
- Authorization
- Secure API Keys
- Audit Logging
- Input Validation
- Rate Limiting

Sensitive operations should require appropriate authorization.

---

# Monitoring

Recommended metrics include:

- Registered merchants
- Active merchants
- API requests
- Failed requests
- Payment volume
- Authentication failures
- Average response time

These metrics support operational monitoring.

---

# Scalability

The Merchant Service is designed to support:

- Stateless application servers
- Horizontal scaling
- Distributed deployments
- Container orchestration
- Independent service updates

---

# Future Enhancements

Potential future capabilities include:

- Multi-organization support
- Team management
- Fine-grained permissions
- Merchant analytics
- Business verification workflows
- Risk scoring
- Fraud monitoring
- Marketplace integrations

Future enhancements will remain aligned with the overall KongPay architecture.

---

# Related Documentation

- architecture.md
- system-overview.md
- payment-lifecycle.md
- wallet.md
- authentication.md
- authorization.md
- api-reference.md
- security.md

---

# Conclusion

The Merchant Service provides the operational foundation for businesses integrating with KongPay.

By separating merchant management from payment processing, the platform achieves greater modularity, maintainability, and scalability while enabling secure and standardized merchant interactions across the KongPay ecosystem.
