# Test Automation Project Documentation

## Overview

The **test-automation** project enables automated testing using [Cucumber](https://cucumber.io/). Tests are written in `.feature` files, following the principles of Behavior-Driven Development (BDD). 

Each Jira story is covered by:
- **One feature file** per story
- **One scenario** for each Acceptance Criteria (AC)

This approach ensures that automated tests directly reflect the business requirements tracked in Jira.

All API calls performed during test execution pass through the **Tier1 Gateway**. This makes it possible to test user authentication and authorization flows as well as API functionality.

---

## How Tests Work

Here’s the typical lifecycle of a test execution in the `test-automation` project:

1. **Write Feature Files**
   - Each feature represents a Jira story.
   - Each scenario represents an Acceptance Criteria (AC).
   - Feature files are written in Gherkin syntax.

2. **Run Tests with Cucumber**
   - Cucumber reads the `.feature` files.
   - Step definitions execute corresponding test logic.

3. **Calls Routed via Tier1 Gateway**
   - All requests (API calls, authentication) go through the Tier1 Gateway.
   - The gateway ensures that:
     - The proper user context is maintained.
     - Authorization checks are enforced.

4. **Authentication with Keycloak**
   - Tests authenticate users via Keycloak.
   - Tokens are issued and used for subsequent API calls.

5. **API Testing**
   - Real API endpoints are tested if services are available.
   - Alternatively, APIs can be mocked using Microcks (see below).

---

## Minimum Requirements

To run the `test-automation` project successfully, the following components must be available:

- **Tier1 Gateway**
  - Acts as the central entry point for API requests.
  - Handles routing, security, and user context.

- **Keycloak**
  - Provides authentication and authorization services.
  - Required for obtaining tokens used in API calls.

- **Microservices Under Test**
  - APIs or services that the tests will exercise.
  - Must be running and properly configured.

> **Note:** If running tests locally, all the above components (Tier1, Keycloak, and the microservices) must be accessible and properly configured in your local environment.

---

## Using Mocks with Microcks

To adopt a test-first approach and support parallel development between frontend and backend teams, the project integrates **Microcks** for mocking REST services.

### Why Use Microcks?

- Enables testing even if backend services are not fully implemented.
- Provides stable API responses for frontend development.
- Reduces dependencies between teams.

### How It Works

- Each microservice defines an OpenAPI specification.
- This OpenAPI file is used to configure a mock in Microcks.
- During test execution:
  - Tests can target real services if they are available.
  - Otherwise, tests can run against mocked services hosted in Microcks.

This strategy ensures that both API contracts and integration flows can be verified early in the development lifecycle.

---

## Benefits of This Approach

- Tests reflect real business requirements (Jira ACs).  
- Early detection of integration issues.  
- Ability to test authentication and authorization flows.  
- Supports parallel development (frontend and backend).  
- Flexible use of real or mocked services.

---

## Conclusion

The `test-automation` project provides a robust foundation for automated testing using BDD principles. By combining Cucumber, Tier1 Gateway, Keycloak, and Microcks, it ensures reliable, scalable, and maintainable test coverage for your applications.


