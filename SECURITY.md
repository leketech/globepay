# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within Globepay, please send an email to security@globepay.com. All security vulnerabilities will be promptly addressed.

Please do not publicly disclose the vulnerability until it has been addressed by the team.

## Security Measures

### Backend Security

- All data is encrypted at rest using AWS KMS
- TLS 1.2+ for all data in transit
- JWT-based authentication with secure token handling
- Input validation and sanitization on all endpoints
- SQL injection prevention through parameterized queries
- Rate limiting to prevent abuse
- Regular security scanning with Trivy and Snyk

### Frontend Security

- XSS protection headers
- CSRF tokens for state-changing operations
- Content Security Policy (CSP) headers
- Secure cookie settings
- Input validation and sanitization

### Infrastructure Security

- RBAC enabled on Kubernetes clusters
- Network policies for pod isolation
- Secrets stored in AWS Secrets Manager
- WAF rules for DDoS protection
- Regular security audits

## Compliance

Globepay is designed to be compliant with:

- PCI-DSS (Ready for Level 1 compliance)
- GDPR (Data protection controls in place)
- SOC 2 (Audit-ready logging and monitoring)

## Security Updates

We regularly update dependencies and release security patches. Users are encouraged to:

1. Keep their installations up to date
2. Monitor security advisories
3. Subscribe to our security mailing list

## Security Audits

We conduct regular security audits of our codebase and infrastructure. Third-party security firms are engaged annually to perform comprehensive penetration testing and code reviews.

## Incident Response

In the event of a security incident:

1. The incident is immediately contained
2. Affected systems are isolated
3. Forensic analysis is performed
4. Customers are notified if their data may have been compromised
5. A post-incident report is published
6. Preventive measures are implemented to avoid recurrence

## Contact

For any security-related questions or concerns, please contact:

- Email: security@globepay.com
- PGP Key: [Available upon request]