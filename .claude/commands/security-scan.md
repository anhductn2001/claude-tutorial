Perform a security audit of the current codebase or file: $ARGUMENTS

Check for:

**Secrets & Credentials**
- Hardcoded API keys, passwords, tokens
- .env files in git history
- Credentials in logs or error messages

**Injection Vulnerabilities**
- SQL injection (string concatenation in queries)
- Command injection (unsanitized shell exec)
- XSS (unescaped user content in HTML output)

**Authentication & Authorization**
- Missing auth checks on protected routes
- IDOR (user A accessing user B's data)
- Weak or missing session management

**Dependencies**
- Check package.json / go.mod / requirements.txt for known CVEs
- Flag outdated packages with security patches available

Output format:
## CRITICAL
- [file:line] Description + remediation

## HIGH
- [file:line] Description + remediation

## MEDIUM / LOW
- [file:line] Description
